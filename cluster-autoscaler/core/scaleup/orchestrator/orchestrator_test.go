/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package orchestrator

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"
	"time"

	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider"
	mockprovider "k8s.io/autoscaler/cluster-autoscaler/cloudprovider/mocks"
	testprovider "k8s.io/autoscaler/cluster-autoscaler/cloudprovider/test"
	"k8s.io/autoscaler/cluster-autoscaler/clusterstate"
	"k8s.io/autoscaler/cluster-autoscaler/config"
	"k8s.io/autoscaler/cluster-autoscaler/core/scaleup/resource"
	. "k8s.io/autoscaler/cluster-autoscaler/core/test"
	"k8s.io/autoscaler/cluster-autoscaler/core/utils"
	"k8s.io/autoscaler/cluster-autoscaler/estimator"
	"k8s.io/autoscaler/cluster-autoscaler/metrics"
	"k8s.io/autoscaler/cluster-autoscaler/processors/nodegroupset"
	"k8s.io/autoscaler/cluster-autoscaler/processors/nodeinfosprovider"
	"k8s.io/autoscaler/cluster-autoscaler/processors/status"
	"k8s.io/autoscaler/cluster-autoscaler/utils/errors"
	kube_util "k8s.io/autoscaler/cluster-autoscaler/utils/kubernetes"
	"k8s.io/autoscaler/cluster-autoscaler/utils/taints"
	. "k8s.io/autoscaler/cluster-autoscaler/utils/test"
	"k8s.io/autoscaler/cluster-autoscaler/utils/units"
	kube_record "k8s.io/client-go/tools/record"
	"k8s.io/component-base/metrics/legacyregistry"

	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes/fake"
	schedulerframework "k8s.io/kubernetes/pkg/scheduler/framework"

	"github.com/stretchr/testify/assert"
	"k8s.io/autoscaler/cluster-autoscaler/expander"
)

var defaultOptions = config.AutoscalingOptions{
	EstimatorName:  estimator.BinpackingEstimatorName,
	MaxCoresTotal:  config.DefaultMaxClusterCores,
	MaxMemoryTotal: config.DefaultMaxClusterMemory * units.GiB,
	MinCoresTotal:  0,
	MinMemoryTotal: 0,
}

// Scale up scenarios.
func TestScaleUpOK(t *testing.T) {
	config := &ScaleTestConfig{
		Nodes: []NodeConfig{
			{Name: "n1", Cpu: 100, Memory: 100, Gpu: 0, Ready: true, Group: "ng1"},
			{Name: "n2", Cpu: 1000, Memory: 1000, Gpu: 0, Ready: true, Group: "ng2"},
		},
		Pods: []PodConfig{
			{Name: "p1", Cpu: 80, Memory: 0, Gpu: 0, Node: "n1", ToleratesGpu: false},
			{Name: "p2", Cpu: 800, Memory: 0, Gpu: 0, Node: "n2", ToleratesGpu: false},
		},
		ExtraPods: []PodConfig{
			{Name: "p-new", Cpu: 500, Memory: 0, Gpu: 0, Node: "", ToleratesGpu: false},
		},
		Options:                 defaultOptions,
		ExpansionOptionToChoose: GroupSizeChange{GroupName: "ng2", SizeChange: 1},
	}
	expectedResults := &ScaleTestResults{
		FinalOption: GroupSizeChange{GroupName: "ng2", SizeChange: 1},
		ScaleUpStatus: ScaleUpStatusInfo{
			PodsTriggeredScaleUp: []string{"p-new"},
		},
	}

	simpleScaleUpTest(t, config, expectedResults)
}

// There are triggering, remaining & awaiting pods.
func TestMixedScaleUp(t *testing.T) {
	config := &ScaleTestConfig{
		Nodes: []NodeConfig{
			{Name: "n1", Cpu: 100, Memory: 1000, Gpu: 0, Ready: true, Group: "ng1"},
			{Name: "n2", Cpu: 1000, Memory: 100, Gpu: 0, Ready: true, Group: "ng2"},
		},
		Pods: []PodConfig{
			{Name: "p1", Cpu: 80, Memory: 0, Gpu: 0, Node: "n1", ToleratesGpu: false},
			{Name: "p2", Cpu: 800, Memory: 0, Gpu: 0, Node: "n2", ToleratesGpu: false},
		},
		ExtraPods: []PodConfig{
			// can only be scheduled on ng2
			{Name: "triggering", Cpu: 900, Memory: 0, Gpu: 0, Node: "", ToleratesGpu: false},
			// can't be scheduled
			{Name: "remaining", Cpu: 2000, Memory: 0, Gpu: 0, Node: "", ToleratesGpu: false},
			// can only be scheduled on ng1
			{Name: "awaiting", Cpu: 0, Memory: 200, Gpu: 0, Node: "", ToleratesGpu: false},
		},
		Options:                 defaultOptions,
		ExpansionOptionToChoose: GroupSizeChange{GroupName: "ng2", SizeChange: 1},
	}
	expectedResults := &ScaleTestResults{
		FinalOption: GroupSizeChange{GroupName: "ng2", SizeChange: 1},
		ScaleUpStatus: ScaleUpStatusInfo{
			PodsTriggeredScaleUp:    []string{"triggering"},
			PodsRemainUnschedulable: []string{"remaining"},
			PodsAwaitEvaluation:     []string{"awaiting"},
		},
	}

	simpleScaleUpTest(t, config, expectedResults)
}

func TestScaleUpMaxCoresLimitHit(t *testing.T) {
	options := defaultOptions
	options.MaxCoresTotal = 9
	config := &ScaleTestConfig{
		Nodes: []NodeConfig{
			{Name: "n1", Cpu: 2000, Memory: 100, Gpu: 0, Ready: true, Group: "ng1"},
			{Name: "n2", Cpu: 4000, Memory: 1000, Gpu: 0, Ready: true, Group: "ng2"},
		},
		Pods: []PodConfig{
			{Name: "p1", Cpu: 1000, Memory: 0, Gpu: 0, Node: "n1", ToleratesGpu: false},
			{Name: "p2", Cpu: 3000, Memory: 0, Gpu: 0, Node: "n2", ToleratesGpu: false},
		},
		ExtraPods: []PodConfig{
			{Name: "p-new-1", Cpu: 2000, Memory: 0, Gpu: 0, Node: "", ToleratesGpu: false},
			{Name: "p-new-2", Cpu: 2000, Memory: 0, Gpu: 0, Node: "", ToleratesGpu: false},
		},
		ExpansionOptionToChoose: GroupSizeChange{GroupName: "ng1", SizeChange: 2},
		Options:                 options,
	}
	results := &ScaleTestResults{
		FinalOption: GroupSizeChange{GroupName: "ng1", SizeChange: 1},
		ScaleUpStatus: ScaleUpStatusInfo{
			PodsTriggeredScaleUp: []string{"p-new-1", "p-new-2"},
		},
	}

	simpleScaleUpTest(t, config, results)
}

func TestScaleUpMaxCoresLimitHitWithNotAutoscaledGroup(t *testing.T) {
	options := defaultOptions
	options.MaxCoresTotal = 9
	config := &ScaleTestConfig{
		Nodes: []NodeConfig{
			{Name: "n1", Cpu: 2000, Memory: 100, Gpu: 0, Ready: true, Group: "ng1"},
			{Name: "n2", Cpu: 4000, Memory: 1000, Gpu: 0, Ready: true, Group: ""},
		},
		Pods: []PodConfig{
			{Name: "p1", Cpu: 1000, Memory: 0, Gpu: 0, Node: "n1", ToleratesGpu: false},
			{Name: "p2", Cpu: 3000, Memory: 0, Gpu: 0, Node: "n2", ToleratesGpu: false},
		},
		ExtraPods: []PodConfig{
			{Name: "p-new-1", Cpu: 2000, Memory: 0, Gpu: 0, Node: "", ToleratesGpu: false},
			{Name: "p-new-2", Cpu: 2000, Memory: 0, Gpu: 0, Node: "", ToleratesGpu: false},
		},
		ExpansionOptionToChoose: GroupSizeChange{GroupName: "ng1", SizeChange: 2},
		Options:                 options,
	}
	results := &ScaleTestResults{
		FinalOption: GroupSizeChange{GroupName: "ng1", SizeChange: 1},
		ScaleUpStatus: ScaleUpStatusInfo{
			PodsTriggeredScaleUp: []string{"p-new-1", "p-new-2"},
		},
	}

	simpleScaleUpTest(t, config, results)
}

func TestScaleUpMaxMemoryLimitHit(t *testing.T) {
	options := defaultOptions
	options.MaxMemoryTotal = 1300 * utils.MiB
	config := &ScaleTestConfig{
		Nodes: []NodeConfig{
			{Name: "n1", Cpu: 2000, Memory: 100 * utils.MiB, Gpu: 0, Ready: true, Group: "ng1"},
			{Name: "n2", Cpu: 4000, Memory: 1000 * utils.MiB, Gpu: 0, Ready: true, Group: "ng2"},
		},
		Pods: []PodConfig{
			{Name: "p1", Cpu: 1000, Memory: 0, Gpu: 0, Node: "n1", ToleratesGpu: false},
			{Name: "p2", Cpu: 3000, Memory: 0, Gpu: 0, Node: "n2", ToleratesGpu: false},
		},
		ExtraPods: []PodConfig{
			{Name: "p-new-1", Cpu: 2000, Memory: 100 * utils.MiB, Gpu: 0, Node: "", ToleratesGpu: false},
			{Name: "p-new-2", Cpu: 2000, Memory: 100 * utils.MiB, Gpu: 0, Node: "", ToleratesGpu: false},
			{Name: "p-new-3", Cpu: 2000, Memory: 100 * utils.MiB, Gpu: 0, Node: "", ToleratesGpu: false},
		},
		ExpansionOptionToChoose: GroupSizeChange{GroupName: "ng1", SizeChange: 3},
		Options:                 options,
	}
	results := &ScaleTestResults{
		FinalOption: GroupSizeChange{GroupName: "ng1", SizeChange: 2},
		ScaleUpStatus: ScaleUpStatusInfo{
			PodsTriggeredScaleUp: []string{"p-new-1", "p-new-2", "p-new-3"},
		},
	}

	simpleScaleUpTest(t, config, results)
}

func TestScaleUpMaxMemoryLimitHitWithNotAutoscaledGroup(t *testing.T) {
	options := defaultOptions
	options.MaxMemoryTotal = 1300 * utils.MiB
	config := &ScaleTestConfig{
		Nodes: []NodeConfig{
			{Name: "n1", Cpu: 2000, Memory: 100 * utils.MiB, Gpu: 0, Ready: true, Group: "ng1"},
			{Name: "n2", Cpu: 4000, Memory: 1000 * utils.MiB, Gpu: 0, Ready: true, Group: ""},
		},
		Pods: []PodConfig{
			{Name: "p1", Cpu: 1000, Memory: 0, Gpu: 0, Node: "n1", ToleratesGpu: false},
			{Name: "p2", Cpu: 3000, Memory: 0, Gpu: 0, Node: "n2", ToleratesGpu: false},
		},
		ExtraPods: []PodConfig{
			{Name: "p-new-1", Cpu: 2000, Memory: 100 * utils.MiB, Gpu: 0, Node: "", ToleratesGpu: false},
			{Name: "p-new-2", Cpu: 2000, Memory: 100 * utils.MiB, Gpu: 0, Node: "", ToleratesGpu: false},
			{Name: "p-new-3", Cpu: 2000, Memory: 100 * utils.MiB, Gpu: 0, Node: "", ToleratesGpu: false},
		},
		ExpansionOptionToChoose: GroupSizeChange{GroupName: "ng1", SizeChange: 3},
		Options:                 options,
	}
	results := &ScaleTestResults{
		FinalOption: GroupSizeChange{GroupName: "ng1", SizeChange: 2},
		ScaleUpStatus: ScaleUpStatusInfo{
			PodsTriggeredScaleUp: []string{"p-new-1", "p-new-2", "p-new-3"},
		},
	}

	simpleScaleUpTest(t, config, results)
}

func TestScaleUpCapToMaxTotalNodesLimit(t *testing.T) {
	options := defaultOptions
	options.MaxNodesTotal = 3
	config := &ScaleTestConfig{
		Nodes: []NodeConfig{
			{Name: "n1", Cpu: 2000, Memory: 100 * utils.MiB, Gpu: 0, Ready: true, Group: "ng1"},
			{Name: "n2", Cpu: 4000, Memory: 1000 * utils.MiB, Gpu: 0, Ready: true, Group: "ng2"},
		},
		Pods: []PodConfig{
			{Name: "p1", Cpu: 1000, Memory: 0, Gpu: 0, Node: "n1", ToleratesGpu: false},
			{Name: "p2", Cpu: 3000, Memory: 0, Gpu: 0, Node: "n2", ToleratesGpu: false},
		},
		ExtraPods: []PodConfig{
			{Name: "p-new-1", Cpu: 4000, Memory: 100 * utils.MiB, Gpu: 0, Node: "", ToleratesGpu: false},
			{Name: "p-new-2", Cpu: 4000, Memory: 100 * utils.MiB, Gpu: 0, Node: "", ToleratesGpu: false},
			{Name: "p-new-3", Cpu: 4000, Memory: 100 * utils.MiB, Gpu: 0, Node: "", ToleratesGpu: false},
		},
		ExpansionOptionToChoose: GroupSizeChange{GroupName: "ng2", SizeChange: 3},
		Options:                 options,
	}
	results := &ScaleTestResults{
		FinalOption: GroupSizeChange{GroupName: "ng2", SizeChange: 1},
		ScaleUpStatus: ScaleUpStatusInfo{
			PodsTriggeredScaleUp: []string{"p-new-1", "p-new-2", "p-new-3"},
		},
	}

	simpleScaleUpTest(t, config, results)
}

func TestScaleUpCapToMaxTotalNodesLimitWithNotAutoscaledGroup(t *testing.T) {
	options := defaultOptions
	options.MaxNodesTotal = 3
	config := &ScaleTestConfig{
		Nodes: []NodeConfig{
			{Name: "n1", Cpu: 2000, Memory: 100 * utils.MiB, Gpu: 0, Ready: true, Group: ""},
			{Name: "n2", Cpu: 4000, Memory: 1000 * utils.MiB, Gpu: 0, Ready: true, Group: "ng2"},
		},
		Pods: []PodConfig{
			{Name: "p1", Cpu: 1000, Memory: 0, Gpu: 0, Node: "n1", ToleratesGpu: false},
			{Name: "p2", Cpu: 3000, Memory: 0, Gpu: 0, Node: "n2", ToleratesGpu: false},
		},
		ExtraPods: []PodConfig{
			{Name: "p-new-1", Cpu: 4000, Memory: 100 * utils.MiB, Gpu: 0, Node: "", ToleratesGpu: false},
			{Name: "p-new-2", Cpu: 4000, Memory: 100 * utils.MiB, Gpu: 0, Node: "", ToleratesGpu: false},
			{Name: "p-new-3", Cpu: 4000, Memory: 100 * utils.MiB, Gpu: 0, Node: "", ToleratesGpu: false},
		},
		ExpansionOptionToChoose: GroupSizeChange{GroupName: "ng2", SizeChange: 3},
		Options:                 options,
	}
	results := &ScaleTestResults{
		FinalOption: GroupSizeChange{GroupName: "ng2", SizeChange: 1},
		ScaleUpStatus: ScaleUpStatusInfo{
			PodsTriggeredScaleUp: []string{"p-new-1", "p-new-2", "p-new-3"},
		},
	}

	simpleScaleUpTest(t, config, results)
}

func TestWillConsiderGpuAndStandardPoolForPodWhichDoesNotRequireGpu(t *testing.T) {
	options := defaultOptions
	options.MaxNodesTotal = 100
	config := &ScaleTestConfig{
		Nodes: []NodeConfig{
			{Name: "gpu-node-1", Cpu: 2000, Memory: 1000 * utils.MiB, Gpu: 1, Ready: true, Group: "gpu-pool"},
			{Name: "std-node-1", Cpu: 2000, Memory: 1000 * utils.MiB, Gpu: 0, Ready: true, Group: "std-pool"},
		},
		Pods: []PodConfig{
			{Name: "gpu-pod-1", Cpu: 2000, Memory: 1000 * utils.MiB, Gpu: 1, Node: "gpu-node-1", ToleratesGpu: true},
			{Name: "std-pod-1", Cpu: 2000, Memory: 1000 * utils.MiB, Gpu: 0, Node: "std-node-1", ToleratesGpu: false},
		},
		ExtraPods: []PodConfig{
			{Name: "extra-std-pod", Cpu: 2000, Memory: 1000 * utils.MiB, Gpu: 0, Node: "", ToleratesGpu: true},
		},
		ExpansionOptionToChoose: GroupSizeChange{GroupName: "std-pool", SizeChange: 1},
		Options:                 options,
	}
	results := &ScaleTestResults{
		FinalOption: GroupSizeChange{GroupName: "std-pool", SizeChange: 1},
		ExpansionOptions: []GroupSizeChange{
			{GroupName: "std-pool", SizeChange: 1},
			{GroupName: "gpu-pool", SizeChange: 1},
		},
		ScaleUpStatus: ScaleUpStatusInfo{
			PodsTriggeredScaleUp: []string{"extra-std-pod"},
		},
	}

	simpleScaleUpTest(t, config, results)
}

func TestWillConsiderOnlyGpuPoolForPodWhichDoesRequiresGpu(t *testing.T) {
	options := defaultOptions
	options.MaxNodesTotal = 100
	config := &ScaleTestConfig{
		Nodes: []NodeConfig{
			{Name: "gpu-node-1", Cpu: 2000, Memory: 1000 * utils.MiB, Gpu: 1, Ready: true, Group: "gpu-pool"},
			{Name: "std-node-1", Cpu: 2000, Memory: 1000 * utils.MiB, Gpu: 0, Ready: true, Group: "std-pool"},
		},
		Pods: []PodConfig{
			{Name: "gpu-pod-1", Cpu: 2000, Memory: 1000 * utils.MiB, Gpu: 1, Node: "gpu-node-1", ToleratesGpu: true},
			{Name: "std-pod-1", Cpu: 2000, Memory: 1000 * utils.MiB, Gpu: 0, Node: "std-node-1", ToleratesGpu: false},
		},
		ExtraPods: []PodConfig{
			{Name: "extra-gpu-pod", Cpu: 2000, Memory: 1000 * utils.MiB, Gpu: 1, Node: "", ToleratesGpu: true},
		},
		ExpansionOptionToChoose: GroupSizeChange{GroupName: "gpu-pool", SizeChange: 1},
		Options:                 options,
	}
	results := &ScaleTestResults{
		FinalOption: GroupSizeChange{GroupName: "gpu-pool", SizeChange: 1},
		ExpansionOptions: []GroupSizeChange{
			{GroupName: "gpu-pool", SizeChange: 1},
		},
		ScaleUpStatus: ScaleUpStatusInfo{
			PodsTriggeredScaleUp: []string{"extra-gpu-pod"},
		},
	}

	simpleScaleUpTest(t, config, results)
}

func TestWillConsiderAllPoolsWhichFitTwoPodsRequiringGpus(t *testing.T) {
	options := defaultOptions
	options.MaxNodesTotal = 100
	config := &ScaleTestConfig{
		Nodes: []NodeConfig{
			{Name: "gpu-1-node-1", Cpu: 2000, Memory: 1000 * utils.MiB, Gpu: 1, Ready: true, Group: "gpu-1-pool"},
			{Name: "gpu-2-node-1", Cpu: 2000, Memory: 1000 * utils.MiB, Gpu: 2, Ready: true, Group: "gpu-2-pool"},
			{Name: "gpu-4-node-1", Cpu: 2000, Memory: 1000 * utils.MiB, Gpu: 4, Ready: true, Group: "gpu-4-pool"},
			{Name: "std-node-1", Cpu: 2000, Memory: 1000 * utils.MiB, Gpu: 0, Ready: true, Group: "std-pool"},
		},
		Pods: []PodConfig{
			{Name: "gpu-pod-1", Cpu: 2000, Memory: 1000 * utils.MiB, Gpu: 1, Node: "gpu-1-node-1", ToleratesGpu: true},
			{Name: "gpu-pod-2", Cpu: 2000, Memory: 1000 * utils.MiB, Gpu: 2, Node: "gpu-2-node-1", ToleratesGpu: true},
			{Name: "gpu-pod-3", Cpu: 2000, Memory: 1000 * utils.MiB, Gpu: 4, Node: "gpu-4-node-1", ToleratesGpu: true},
			{Name: "std-pod-1", Cpu: 2000, Memory: 1000 * utils.MiB, Gpu: 0, Node: "std-node-1", ToleratesGpu: false},
		},
		ExtraPods: []PodConfig{
			{Name: "extra-gpu-pod-1", Cpu: 1, Memory: 1 * utils.MiB, Gpu: 1, Node: "", ToleratesGpu: true}, // CPU and mem negligible
			{Name: "extra-gpu-pod-2", Cpu: 1, Memory: 1 * utils.MiB, Gpu: 1, Node: "", ToleratesGpu: true}, // CPU and mem negligible
			{Name: "extra-gpu-pod-3", Cpu: 1, Memory: 1 * utils.MiB, Gpu: 1, Node: "", ToleratesGpu: true}, // CPU and mem negligible
		},
		ExpansionOptionToChoose: GroupSizeChange{GroupName: "gpu-1-pool", SizeChange: 3},
		Options:                 options,
	}
	results := &ScaleTestResults{
		FinalOption: GroupSizeChange{GroupName: "gpu-1-pool", SizeChange: 3},
		ExpansionOptions: []GroupSizeChange{
			{GroupName: "gpu-1-pool", SizeChange: 3},
			{GroupName: "gpu-2-pool", SizeChange: 2},
			{GroupName: "gpu-4-pool", SizeChange: 1},
		},
		ScaleUpStatus: ScaleUpStatusInfo{
			PodsTriggeredScaleUp: []string{"extra-gpu-pod-1", "extra-gpu-pod-2", "extra-gpu-pod-3"},
		},
	}

	simpleScaleUpTest(t, config, results)
}

// No scale up scenarios.
func TestNoScaleUpMaxCoresLimitHit(t *testing.T) {
	options := defaultOptions
	options.MaxCoresTotal = 7
	options.MaxMemoryTotal = 1150
	config := &ScaleTestConfig{
		Nodes: []NodeConfig{
			{Name: "n1", Cpu: 2000, Memory: 100, Gpu: 0, Ready: true, Group: "ng1"},
			{Name: "n2", Cpu: 4000, Memory: 1000, Gpu: 0, Ready: true, Group: "ng2"},
		},
		Pods: []PodConfig{
			{Name: "p1", Cpu: 1000, Memory: 0, Gpu: 0, Node: "n1", ToleratesGpu: false},
			{Name: "p2", Cpu: 3000, Memory: 0, Gpu: 0, Node: "n2", ToleratesGpu: false},
		},
		ExtraPods: []PodConfig{
			{Name: "p-new-1", Cpu: 2000, Memory: 0, Gpu: 0, Node: "", ToleratesGpu: false},
			{Name: "p-new-2", Cpu: 2000, Memory: 0, Gpu: 0, Node: "", ToleratesGpu: false},
		},
		Options: options,
	}
	results := &ScaleTestResults{
		NoScaleUpReason: "max cluster cpu, memory limit reached",
		ScaleUpStatus: ScaleUpStatusInfo{
			PodsRemainUnschedulable: []string{"p-new-1", "p-new-2"},
		},
	}

	simpleNoScaleUpTest(t, config, results)
}

// To implement expander.Strategy, BestOption method must have a struct receiver.
// This prevents it from modifying fields of reportingStrategy, so we need a thin
// pointer wrapper for mutable parts.
type expanderResults struct {
	inputOptions []GroupSizeChange
}

type reportingStrategy struct {
	initialNodeConfigs []NodeConfig
	optionToChoose     GroupSizeChange
	results            *expanderResults
	t                  *testing.T
}

func (r reportingStrategy) BestOption(options []expander.Option, nodeInfo map[string]*schedulerframework.NodeInfo) *expander.Option {
	r.results.inputOptions = expanderOptionsToGroupSizeChanges(options)
	for _, option := range options {
		GroupSizeChange := expanderOptionToGroupSizeChange(option)
		if GroupSizeChange == r.optionToChoose {
			return &option
		}
	}
	assert.Fail(r.t, "did not find expansionOptionToChoose %s", r.optionToChoose)
	return nil
}

func expanderOptionsToGroupSizeChanges(options []expander.Option) []GroupSizeChange {
	groupSizeChanges := make([]GroupSizeChange, 0, len(options))
	for _, option := range options {
		GroupSizeChange := expanderOptionToGroupSizeChange(option)
		groupSizeChanges = append(groupSizeChanges, GroupSizeChange)
	}
	return groupSizeChanges
}

func expanderOptionToGroupSizeChange(option expander.Option) GroupSizeChange {
	groupName := option.NodeGroup.Id()
	groupSizeIncrement := option.NodeCount
	scaleUpOption := GroupSizeChange{GroupName: groupName, SizeChange: groupSizeIncrement}
	return scaleUpOption
}

func runSimpleScaleUpTest(t *testing.T, config *ScaleTestConfig) *ScaleTestResults {
	expandedGroups := make(chan GroupSizeChange, 10)
	now := time.Now()

	groups := make(map[string][]*apiv1.Node)
	nodes := make([]*apiv1.Node, 0, len(config.Nodes))
	for _, n := range config.Nodes {
		node := BuildTestNode(n.Name, n.Cpu, n.Memory)
		if n.Gpu > 0 {
			AddGpusToNode(node, n.Gpu)
		}
		SetNodeReadyState(node, n.Ready, now.Add(-2*time.Minute))
		nodes = append(nodes, node)
		if n.Group != "" {
			groups[n.Group] = append(groups[n.Group], node)
		}
	}

	pods := make([]*apiv1.Pod, 0, len(config.Pods))
	for _, p := range config.Pods {
		pod := buildTestPod(p)
		pods = append(pods, pod)
	}

	podLister := kube_util.NewTestPodLister(pods)
	listers := kube_util.NewListerRegistry(nil, nil, podLister, nil, nil, nil, nil, nil, nil, nil)

	provider := testprovider.NewTestCloudProvider(func(nodeGroup string, increase int) error {
		expandedGroups <- GroupSizeChange{GroupName: nodeGroup, SizeChange: increase}
		return nil
	}, nil)

	for name, nodesInGroup := range groups {
		provider.AddNodeGroup(name, 1, 10, len(nodesInGroup))
		for _, n := range nodesInGroup {
			provider.AddNode(name, n)
		}
	}

	resourceLimiter := cloudprovider.NewResourceLimiter(
		map[string]int64{cloudprovider.ResourceNameCores: config.Options.MinCoresTotal, cloudprovider.ResourceNameMemory: config.Options.MinMemoryTotal},
		map[string]int64{cloudprovider.ResourceNameCores: config.Options.MaxCoresTotal, cloudprovider.ResourceNameMemory: config.Options.MaxMemoryTotal})
	provider.SetResourceLimiter(resourceLimiter)

	assert.NotNil(t, provider)

	// Create context with non-random expander strategy.
	context, err := NewScaleTestAutoscalingContext(config.Options, &fake.Clientset{}, listers, provider, nil, nil)
	assert.NoError(t, err)

	expander := reportingStrategy{
		initialNodeConfigs: config.Nodes,
		optionToChoose:     config.ExpansionOptionToChoose,
		results:            &expanderResults{},
		t:                  t,
	}
	context.ExpanderStrategy = expander

	nodeInfos, _ := nodeinfosprovider.NewDefaultTemplateNodeInfoProvider(nil, false).Process(&context, nodes, []*appsv1.DaemonSet{}, taints.TaintConfig{}, now)
	clusterState := clusterstate.NewClusterStateRegistry(provider, clusterstate.ClusterStateRegistryConfig{}, context.LogRecorder, NewBackoff(), clusterstate.NewStaticMaxNodeProvisionTimeProvider(15*time.Minute))
	clusterState.UpdateNodes(nodes, nodeInfos, time.Now())

	extraPods := make([]*apiv1.Pod, 0, len(config.ExtraPods))
	for _, p := range config.ExtraPods {
		pod := buildTestPod(p)
		extraPods = append(extraPods, pod)
	}

	processors := NewTestProcessors(&context)
	suOrchestrator := New()
	suOrchestrator.Initialize(&context, processors, clusterState, taints.TaintConfig{})
	scaleUpStatus, err := suOrchestrator.ScaleUp(extraPods, nodes, []*appsv1.DaemonSet{}, nodeInfos)
	processors.ScaleUpStatusProcessor.Process(&context, scaleUpStatus)

	assert.NoError(t, err)

	expandedGroup := getGroupSizeChangeFromChan(expandedGroups)
	var expandedGroupStruct GroupSizeChange
	if expandedGroup != nil {
		expandedGroupStruct = *expandedGroup
	}

	events := []string{}
	for eventsLeft := true; eventsLeft; {
		select {
		case event := <-context.Recorder.(*kube_record.FakeRecorder).Events:
			events = append(events, event)
		default:
			eventsLeft = false
		}
	}

	return &ScaleTestResults{
		ExpansionOptions: expander.results.inputOptions,
		FinalOption:      expandedGroupStruct,
		ScaleUpStatus:    simplifyScaleUpStatus(scaleUpStatus),
		Events:           events,
	}
}

func simpleNoScaleUpTest(t *testing.T, config *ScaleTestConfig, expectedResults *ScaleTestResults) {
	results := runSimpleScaleUpTest(t, config)

	assert.Equal(t, GroupSizeChange{}, results.FinalOption)
	assert.False(t, results.ScaleUpStatus.WasSuccessful())
	noScaleUpEventSeen := false
	for _, event := range results.Events {
		if strings.Contains(event, "NotTriggerScaleUp") {
			if strings.Contains(event, expectedResults.NoScaleUpReason) {
				noScaleUpEventSeen = true
			} else {
				// Surprisingly useful for debugging.
				fmt.Println("Event:", event)
			}
		}
		assert.NotRegexp(t, regexp.MustCompile("TriggeredScaleUp"), event)
	}
	assert.True(t, noScaleUpEventSeen)
	assert.ElementsMatch(t, results.ScaleUpStatus.PodsTriggeredScaleUp, expectedResults.ScaleUpStatus.PodsTriggeredScaleUp,
		"actual and expected triggering pods should be the same")
	assert.ElementsMatch(t, results.ScaleUpStatus.PodsRemainUnschedulable, expectedResults.ScaleUpStatus.PodsRemainUnschedulable,
		"actual and expected remaining pods should be the same")
	assert.ElementsMatch(t, results.ScaleUpStatus.PodsAwaitEvaluation, expectedResults.ScaleUpStatus.PodsAwaitEvaluation,
		"actual and expected awaiting evaluation pods should be the same")
}

func simpleScaleUpTest(t *testing.T, config *ScaleTestConfig, expectedResults *ScaleTestResults) {
	results := runSimpleScaleUpTest(t, config)

	assert.NotNil(t, results.FinalOption, "Expected scale up event")
	assert.Equal(t, expectedResults.FinalOption, results.FinalOption)
	assert.True(t, results.ScaleUpStatus.WasSuccessful())
	nodeEventSeen := false
	for _, event := range results.Events {
		if strings.Contains(event, "TriggeredScaleUp") && strings.Contains(event, expectedResults.FinalOption.GroupName) {
			nodeEventSeen = true
		}
		if len(expectedResults.ScaleUpStatus.PodsRemainUnschedulable) == 0 {
			assert.NotRegexp(t, regexp.MustCompile("NotTriggerScaleUp"), event)
		}
	}
	assert.True(t, nodeEventSeen)

	if len(expectedResults.ExpansionOptions) > 0 {
		// Empty ExpansionOptions means we do not want to do any assertions
		// on contents of actual scaleUp options

		// Check that option to choose is part of expected options.
		assert.Contains(t, expectedResults.ExpansionOptions, config.ExpansionOptionToChoose, "final expected expansion option must be in expected expansion options")
		assert.Contains(t, results.ExpansionOptions, config.ExpansionOptionToChoose, "final expected expansion option must be in expected expansion options")

		assert.ElementsMatch(t, results.ExpansionOptions, expectedResults.ExpansionOptions,
			"actual and expected expansion options should be the same")
	}

	assert.ElementsMatch(t, results.ScaleUpStatus.PodsTriggeredScaleUp, expectedResults.ScaleUpStatus.PodsTriggeredScaleUp,
		"actual and expected triggering pods should be the same")
	assert.ElementsMatch(t, results.ScaleUpStatus.PodsRemainUnschedulable, expectedResults.ScaleUpStatus.PodsRemainUnschedulable,
		"actual and expected remaining pods should be the same")
	assert.ElementsMatch(t, results.ScaleUpStatus.PodsAwaitEvaluation, expectedResults.ScaleUpStatus.PodsAwaitEvaluation,
		"actual and expected awaiting evaluation pods should be the same")
}

func getGroupSizeChangeFromChan(c chan GroupSizeChange) *GroupSizeChange {
	select {
	case val := <-c:
		return &val
	case <-time.After(100 * time.Millisecond):
		return nil
	}
}

func buildTestPod(p PodConfig) *apiv1.Pod {
	pod := BuildTestPod(p.Name, p.Cpu, p.Memory)
	if p.Gpu > 0 {
		RequestGpuForPod(pod, p.Gpu)
	}
	if p.ToleratesGpu {
		TolerateGpuForPod(pod)
	}
	if p.Node != "" {
		pod.Spec.NodeName = p.Node
	}
	return pod
}

func TestScaleUpUnhealthy(t *testing.T) {
	now := time.Now()
	someTimeAgo := now.Add(-2 * time.Minute)
	n1 := BuildTestNode("n1", 100, 1000)
	SetNodeReadyState(n1, true, someTimeAgo)
	n2 := BuildTestNode("n2", 1000, 1000)
	SetNodeReadyState(n2, true, someTimeAgo)

	p1 := BuildTestPod("p1", 80, 0)
	p2 := BuildTestPod("p2", 800, 0)
	p1.Spec.NodeName = "n1"
	p2.Spec.NodeName = "n2"

	podLister := kube_util.NewTestPodLister([]*apiv1.Pod{p1, p2})
	listers := kube_util.NewListerRegistry(nil, nil, podLister, nil, nil, nil, nil, nil, nil, nil)

	provider := testprovider.NewTestCloudProvider(func(nodeGroup string, increase int) error {
		t.Fatalf("No expansion is expected, but increased %s by %d", nodeGroup, increase)
		return nil
	}, nil)
	provider.AddNodeGroup("ng1", 1, 10, 1)
	provider.AddNodeGroup("ng2", 1, 10, 5)
	provider.AddNode("ng1", n1)
	provider.AddNode("ng2", n2)

	options := config.AutoscalingOptions{
		EstimatorName:  estimator.BinpackingEstimatorName,
		MaxCoresTotal:  config.DefaultMaxClusterCores,
		MaxMemoryTotal: config.DefaultMaxClusterMemory,
	}
	context, err := NewScaleTestAutoscalingContext(options, &fake.Clientset{}, listers, provider, nil, nil)
	assert.NoError(t, err)

	nodes := []*apiv1.Node{n1, n2}
	nodeInfos, _ := nodeinfosprovider.NewDefaultTemplateNodeInfoProvider(nil, false).Process(&context, nodes, []*appsv1.DaemonSet{}, taints.TaintConfig{}, now)
	clusterState := clusterstate.NewClusterStateRegistry(provider, clusterstate.ClusterStateRegistryConfig{}, context.LogRecorder, NewBackoff(), clusterstate.NewStaticMaxNodeProvisionTimeProvider(15*time.Minute))
	clusterState.UpdateNodes(nodes, nodeInfos, time.Now())
	p3 := BuildTestPod("p-new", 550, 0)

	processors := NewTestProcessors(&context)
	suOrchestrator := New()
	suOrchestrator.Initialize(&context, processors, clusterState, taints.TaintConfig{})
	scaleUpStatus, err := suOrchestrator.ScaleUp([]*apiv1.Pod{p3}, nodes, []*appsv1.DaemonSet{}, nodeInfos)

	assert.NoError(t, err)
	// Node group is unhealthy.
	assert.False(t, scaleUpStatus.WasSuccessful())
}

func TestScaleUpNoHelp(t *testing.T) {
	n1 := BuildTestNode("n1", 100, 1000)
	now := time.Now()
	SetNodeReadyState(n1, true, now.Add(-2*time.Minute))

	p1 := BuildTestPod("p1", 80, 0)
	p1.Spec.NodeName = "n1"

	podLister := kube_util.NewTestPodLister([]*apiv1.Pod{p1})
	listers := kube_util.NewListerRegistry(nil, nil, podLister, nil, nil, nil, nil, nil, nil, nil)

	provider := testprovider.NewTestCloudProvider(func(nodeGroup string, increase int) error {
		t.Fatalf("No expansion is expected")
		return nil
	}, nil)
	provider.AddNodeGroup("ng1", 1, 10, 1)
	provider.AddNode("ng1", n1)
	assert.NotNil(t, provider)

	options := config.AutoscalingOptions{
		EstimatorName:  estimator.BinpackingEstimatorName,
		MaxCoresTotal:  config.DefaultMaxClusterCores,
		MaxMemoryTotal: config.DefaultMaxClusterMemory,
	}
	context, err := NewScaleTestAutoscalingContext(options, &fake.Clientset{}, listers, provider, nil, nil)
	assert.NoError(t, err)

	nodes := []*apiv1.Node{n1}
	nodeInfos, _ := nodeinfosprovider.NewDefaultTemplateNodeInfoProvider(nil, false).Process(&context, nodes, []*appsv1.DaemonSet{}, taints.TaintConfig{}, now)
	clusterState := clusterstate.NewClusterStateRegistry(provider, clusterstate.ClusterStateRegistryConfig{}, context.LogRecorder, NewBackoff(), clusterstate.NewStaticMaxNodeProvisionTimeProvider(15*time.Minute))
	clusterState.UpdateNodes(nodes, nodeInfos, time.Now())
	p3 := BuildTestPod("p-new", 500, 0)

	processors := NewTestProcessors(&context)
	suOrchestrator := New()
	suOrchestrator.Initialize(&context, processors, clusterState, taints.TaintConfig{})
	scaleUpStatus, err := suOrchestrator.ScaleUp([]*apiv1.Pod{p3}, nodes, []*appsv1.DaemonSet{}, nodeInfos)
	processors.ScaleUpStatusProcessor.Process(&context, scaleUpStatus)

	assert.NoError(t, err)
	assert.False(t, scaleUpStatus.WasSuccessful())
	var event string
	select {
	case event = <-context.Recorder.(*kube_record.FakeRecorder).Events:
	default:
		t.Fatal("No Event recorded, expected NotTriggerScaleUp event")
	}
	assert.Regexp(t, regexp.MustCompile("NotTriggerScaleUp"), event)
}

func TestScaleUpBalanceGroups(t *testing.T) {
	provider := testprovider.NewTestCloudProvider(func(string, int) error {
		return nil
	}, nil)

	type ngInfo struct {
		min, max, size int
	}
	testCfg := map[string]ngInfo{
		"ng1": {min: 1, max: 1, size: 1},
		"ng2": {min: 1, max: 2, size: 1},
		"ng3": {min: 1, max: 5, size: 1},
		"ng4": {min: 1, max: 5, size: 3},
	}
	podList := make([]*apiv1.Pod, 0, len(testCfg))
	nodes := make([]*apiv1.Node, 0)

	now := time.Now()

	for gid, gconf := range testCfg {
		provider.AddNodeGroup(gid, gconf.min, gconf.max, gconf.size)
		for i := 0; i < gconf.size; i++ {
			nodeName := fmt.Sprintf("%v-node-%v", gid, i)
			node := BuildTestNode(nodeName, 100, 1000)
			SetNodeReadyState(node, true, now.Add(-2*time.Minute))
			nodes = append(nodes, node)

			pod := BuildTestPod(fmt.Sprintf("%v-pod-%v", gid, i), 80, 0)
			pod.Spec.NodeName = nodeName
			podList = append(podList, pod)

			provider.AddNode(gid, node)
		}
	}

	podLister := kube_util.NewTestPodLister(podList)
	listers := kube_util.NewListerRegistry(nil, nil, podLister, nil, nil, nil, nil, nil, nil, nil)

	options := config.AutoscalingOptions{
		EstimatorName:            estimator.BinpackingEstimatorName,
		BalanceSimilarNodeGroups: true,
		MaxCoresTotal:            config.DefaultMaxClusterCores,
		MaxMemoryTotal:           config.DefaultMaxClusterMemory,
	}
	context, err := NewScaleTestAutoscalingContext(options, &fake.Clientset{}, listers, provider, nil, nil)
	assert.NoError(t, err)

	nodeInfos, _ := nodeinfosprovider.NewDefaultTemplateNodeInfoProvider(nil, false).Process(&context, nodes, []*appsv1.DaemonSet{}, taints.TaintConfig{}, now)
	clusterState := clusterstate.NewClusterStateRegistry(provider, clusterstate.ClusterStateRegistryConfig{}, context.LogRecorder, NewBackoff(), clusterstate.NewStaticMaxNodeProvisionTimeProvider(15*time.Minute))
	clusterState.UpdateNodes(nodes, nodeInfos, time.Now())

	pods := make([]*apiv1.Pod, 0)
	for i := 0; i < 2; i++ {
		pods = append(pods, BuildTestPod(fmt.Sprintf("test-pod-%v", i), 80, 0))
	}

	processors := NewTestProcessors(&context)
	suOrchestrator := New()
	suOrchestrator.Initialize(&context, processors, clusterState, taints.TaintConfig{})
	scaleUpStatus, typedErr := suOrchestrator.ScaleUp(pods, nodes, []*appsv1.DaemonSet{}, nodeInfos)

	assert.NoError(t, typedErr)
	assert.True(t, scaleUpStatus.WasSuccessful())
	groupMap := make(map[string]cloudprovider.NodeGroup, 3)
	for _, group := range provider.NodeGroups() {
		groupMap[group.Id()] = group
	}

	ng2size, err := groupMap["ng2"].TargetSize()
	assert.NoError(t, err)
	ng3size, err := groupMap["ng3"].TargetSize()
	assert.NoError(t, err)
	assert.Equal(t, 2, ng2size)
	assert.Equal(t, 2, ng3size)
}

func TestScaleUpAutoprovisionedNodeGroup(t *testing.T) {
	createdGroups := make(chan string, 10)
	expandedGroups := make(chan string, 10)

	p1 := BuildTestPod("p1", 80, 0)

	fakeClient := &fake.Clientset{}

	t1 := BuildTestNode("t1", 4000, 1000000)
	SetNodeReadyState(t1, true, time.Time{})
	ti1 := schedulerframework.NewNodeInfo()
	ti1.SetNode(t1)

	provider := testprovider.NewTestAutoprovisioningCloudProvider(
		func(nodeGroup string, increase int) error {
			expandedGroups <- fmt.Sprintf("%s-%d", nodeGroup, increase)
			return nil
		}, nil, func(nodeGroup string) error {
			createdGroups <- nodeGroup
			return nil
		}, nil, []string{"T1"}, map[string]*schedulerframework.NodeInfo{"T1": ti1})

	options := config.AutoscalingOptions{
		EstimatorName:                    estimator.BinpackingEstimatorName,
		MaxCoresTotal:                    5000 * 64,
		MaxMemoryTotal:                   5000 * 64 * 20,
		NodeAutoprovisioningEnabled:      true,
		MaxAutoprovisionedNodeGroupCount: 10,
	}
	podLister := kube_util.NewTestPodLister([]*apiv1.Pod{})
	listers := kube_util.NewListerRegistry(nil, nil, podLister, nil, nil, nil, nil, nil, nil, nil)
	context, err := NewScaleTestAutoscalingContext(options, fakeClient, listers, provider, nil, nil)
	assert.NoError(t, err)

	clusterState := clusterstate.NewClusterStateRegistry(provider, clusterstate.ClusterStateRegistryConfig{}, context.LogRecorder, NewBackoff(), clusterstate.NewStaticMaxNodeProvisionTimeProvider(15*time.Minute))

	processors := NewTestProcessors(&context)
	processors.NodeGroupListProcessor = &MockAutoprovisioningNodeGroupListProcessor{T: t}
	processors.NodeGroupManager = &MockAutoprovisioningNodeGroupManager{T: t, ExtraGroups: 0}

	nodes := []*apiv1.Node{}
	nodeInfos, _ := nodeinfosprovider.NewDefaultTemplateNodeInfoProvider(nil, false).Process(&context, nodes, []*appsv1.DaemonSet{}, taints.TaintConfig{}, time.Now())

	suOrchestrator := New()
	suOrchestrator.Initialize(&context, processors, clusterState, taints.TaintConfig{})
	scaleUpStatus, err := suOrchestrator.ScaleUp([]*apiv1.Pod{p1}, nodes, []*appsv1.DaemonSet{}, nodeInfos)
	assert.NoError(t, err)
	assert.True(t, scaleUpStatus.WasSuccessful())
	assert.Equal(t, "autoprovisioned-T1", utils.GetStringFromChan(createdGroups))
	assert.Equal(t, "autoprovisioned-T1-1", utils.GetStringFromChan(expandedGroups))
}

func TestScaleUpBalanceAutoprovisionedNodeGroups(t *testing.T) {
	createdGroups := make(chan string, 10)
	expandedGroups := make(chan string, 10)

	p1 := BuildTestPod("p1", 80, 0)
	p2 := BuildTestPod("p2", 80, 0)
	p3 := BuildTestPod("p3", 80, 0)

	fakeClient := &fake.Clientset{}

	t1 := BuildTestNode("t1", 100, 1000000)
	SetNodeReadyState(t1, true, time.Time{})
	ti1 := schedulerframework.NewNodeInfo()
	ti1.SetNode(t1)

	provider := testprovider.NewTestAutoprovisioningCloudProvider(
		func(nodeGroup string, increase int) error {
			expandedGroups <- fmt.Sprintf("%s-%d", nodeGroup, increase)
			return nil
		}, nil, func(nodeGroup string) error {
			createdGroups <- nodeGroup
			return nil
		}, nil, []string{"T1"}, map[string]*schedulerframework.NodeInfo{"T1": ti1})

	options := config.AutoscalingOptions{
		BalanceSimilarNodeGroups:         true,
		EstimatorName:                    estimator.BinpackingEstimatorName,
		MaxCoresTotal:                    5000 * 64,
		MaxMemoryTotal:                   5000 * 64 * 20,
		NodeAutoprovisioningEnabled:      true,
		MaxAutoprovisionedNodeGroupCount: 10,
	}
	podLister := kube_util.NewTestPodLister([]*apiv1.Pod{})
	listers := kube_util.NewListerRegistry(nil, nil, podLister, nil, nil, nil, nil, nil, nil, nil)
	context, err := NewScaleTestAutoscalingContext(options, fakeClient, listers, provider, nil, nil)
	assert.NoError(t, err)

	clusterState := clusterstate.NewClusterStateRegistry(provider, clusterstate.ClusterStateRegistryConfig{}, context.LogRecorder, NewBackoff(), clusterstate.NewStaticMaxNodeProvisionTimeProvider(15*time.Minute))

	processors := NewTestProcessors(&context)
	processors.NodeGroupListProcessor = &MockAutoprovisioningNodeGroupListProcessor{T: t}
	processors.NodeGroupManager = &MockAutoprovisioningNodeGroupManager{T: t, ExtraGroups: 2}

	nodes := []*apiv1.Node{}
	nodeInfos, _ := nodeinfosprovider.NewDefaultTemplateNodeInfoProvider(nil, false).Process(&context, nodes, []*appsv1.DaemonSet{}, taints.TaintConfig{}, time.Now())

	suOrchestrator := New()
	suOrchestrator.Initialize(&context, processors, clusterState, taints.TaintConfig{})
	scaleUpStatus, err := suOrchestrator.ScaleUp([]*apiv1.Pod{p1, p2, p3}, nodes, []*appsv1.DaemonSet{}, nodeInfos)
	assert.NoError(t, err)
	assert.True(t, scaleUpStatus.WasSuccessful())
	assert.Equal(t, "autoprovisioned-T1", utils.GetStringFromChan(createdGroups))
	expandedGroupMap := map[string]bool{}
	for i := 0; i < 3; i++ {
		expandedGroupMap[utils.GetStringFromChan(expandedGroups)] = true
	}
	assert.True(t, expandedGroupMap["autoprovisioned-T1-1"])
	assert.True(t, expandedGroupMap["autoprovisioned-T1-1-1"])
	assert.True(t, expandedGroupMap["autoprovisioned-T1-2-1"])
}

func TestScaleUpToMeetNodeGroupMinSize(t *testing.T) {
	podLister := kube_util.NewTestPodLister([]*apiv1.Pod{})
	listers := kube_util.NewListerRegistry(nil, nil, podLister, nil, nil, nil, nil, nil, nil, nil)
	provider := testprovider.NewTestCloudProvider(func(nodeGroup string, increase int) error {
		assert.Equal(t, "ng1", nodeGroup)
		assert.Equal(t, 1, increase)
		return nil
	}, nil)
	resourceLimiter := cloudprovider.NewResourceLimiter(
		map[string]int64{cloudprovider.ResourceNameCores: 0, cloudprovider.ResourceNameMemory: 0},
		map[string]int64{cloudprovider.ResourceNameCores: 48, cloudprovider.ResourceNameMemory: 1000},
	)
	provider.SetResourceLimiter(resourceLimiter)

	// Test cases:
	// ng1: current size 1, min size 3, cores limit 48, memory limit 1000 => scale up with 1 new node.
	// ng2: current size 1, min size 1, cores limit 48, memory limit 1000 => no scale up.
	n1 := BuildTestNode("n1", 16000, 32)
	SetNodeReadyState(n1, true, time.Now())
	n2 := BuildTestNode("n2", 16000, 32)
	SetNodeReadyState(n2, true, time.Now())
	provider.AddNodeGroup("ng1", 3, 10, 1)
	provider.AddNode("ng1", n1)
	provider.AddNodeGroup("ng2", 1, 10, 1)
	provider.AddNode("ng2", n2)

	options := config.AutoscalingOptions{
		EstimatorName:  estimator.BinpackingEstimatorName,
		MaxCoresTotal:  config.DefaultMaxClusterCores,
		MaxMemoryTotal: config.DefaultMaxClusterMemory,
	}
	context, err := NewScaleTestAutoscalingContext(options, &fake.Clientset{}, listers, provider, nil, nil)
	assert.NoError(t, err)

	nodes := []*apiv1.Node{n1, n2}
	nodeInfos, _ := nodeinfosprovider.NewDefaultTemplateNodeInfoProvider(nil, false).Process(&context, nodes, []*appsv1.DaemonSet{}, taints.TaintConfig{}, time.Now())
	processors := NewTestProcessors(&context)
	clusterState := clusterstate.NewClusterStateRegistry(provider, clusterstate.ClusterStateRegistryConfig{}, context.LogRecorder, NewBackoff(), clusterstate.NewStaticMaxNodeProvisionTimeProvider(15*time.Minute))
	clusterState.UpdateNodes(nodes, nodeInfos, time.Now())

	suOrchestrator := New()
	suOrchestrator.Initialize(&context, processors, clusterState, taints.TaintConfig{})
	scaleUpStatus, err := suOrchestrator.ScaleUpToNodeGroupMinSize(nodes, nodeInfos)
	assert.NoError(t, err)
	assert.True(t, scaleUpStatus.WasSuccessful())
	assert.Equal(t, 1, len(scaleUpStatus.ScaleUpInfos))
	assert.Equal(t, 2, scaleUpStatus.ScaleUpInfos[0].NewSize)
	assert.Equal(t, "ng1", scaleUpStatus.ScaleUpInfos[0].Group.Id())
}

func TestCheckDeltaWithinLimits(t *testing.T) {
	type testcase struct {
		limits            resource.Limits
		delta             resource.Delta
		exceededResources []string
	}
	tests := []testcase{
		{
			limits:            resource.Limits{"a": 10},
			delta:             resource.Delta{"a": 10},
			exceededResources: []string{},
		},
		{
			limits:            resource.Limits{"a": 10},
			delta:             resource.Delta{"a": 11},
			exceededResources: []string{"a"},
		},
		{
			limits:            resource.Limits{"a": 10},
			delta:             resource.Delta{"b": 10},
			exceededResources: []string{},
		},
		{
			limits:            resource.Limits{"a": resource.LimitUnknown},
			delta:             resource.Delta{"a": 0},
			exceededResources: []string{},
		},
		{
			limits:            resource.Limits{"a": resource.LimitUnknown},
			delta:             resource.Delta{"a": 1},
			exceededResources: []string{"a"},
		},
		{
			limits:            resource.Limits{"a": 10, "b": 20, "c": 30},
			delta:             resource.Delta{"a": 11, "b": 20, "c": 31},
			exceededResources: []string{"a", "c"},
		},
	}

	for _, test := range tests {
		checkResult := resource.CheckDeltaWithinLimits(test.limits, test.delta)
		if len(test.exceededResources) == 0 {
			assert.Equal(t, resource.LimitsNotExceeded(), checkResult)
		} else {
			assert.Equal(t, resource.LimitsCheckResult{Exceeded: true, ExceededResources: test.exceededResources}, checkResult)
		}
	}
}

func TestAuthError(t *testing.T) {
	metrics.RegisterAll(false)
	context, err := NewScaleTestAutoscalingContext(config.AutoscalingOptions{}, &fake.Clientset{}, nil, nil, nil, nil)
	assert.NoError(t, err)

	nodeGroup := &mockprovider.NodeGroup{}
	info := nodegroupset.ScaleUpInfo{Group: nodeGroup}
	nodeGroup.On("Id").Return("A")
	nodeGroup.On("IncreaseSize", 0).Return(errors.NewAutoscalerError(errors.AutoscalerErrorType("abcd"), ""))

	processors := NewTestProcessors(&context)
	clusterStateRegistry := clusterstate.NewClusterStateRegistry(nil, clusterstate.ClusterStateRegistryConfig{}, context.LogRecorder, NewBackoff(), clusterstate.NewStaticMaxNodeProvisionTimeProvider(15*time.Minute))
	suOrchestrator := New()
	suOrchestrator.Initialize(&context, processors, clusterStateRegistry, taints.TaintConfig{})
	scaleUpOrchestrator := suOrchestrator.(*ScaleUpOrchestrator)
	aerr := scaleUpOrchestrator.executeScaleUp(info, "", "", time.Now())
	assert.Error(t, aerr)

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(legacyregistry.Handler().ServeHTTP)
	handler.ServeHTTP(rr, req)

	// Check that the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	// Check that the failed scale up reason is set correctly.
	assert.Contains(t, rr.Body.String(), "cluster_autoscaler_failed_scale_ups_total{reason=\"abcd\"} 1")
}

func simplifyScaleUpStatus(scaleUpStatus *status.ScaleUpStatus) ScaleUpStatusInfo {
	remainUnschedulable := []string{}
	for _, nsi := range scaleUpStatus.PodsRemainUnschedulable {
		remainUnschedulable = append(remainUnschedulable, nsi.Pod.Name)
	}
	return ScaleUpStatusInfo{
		Result:                  scaleUpStatus.Result,
		PodsTriggeredScaleUp:    ExtractPodNames(scaleUpStatus.PodsTriggeredScaleUp),
		PodsRemainUnschedulable: remainUnschedulable,
		PodsAwaitEvaluation:     ExtractPodNames(scaleUpStatus.PodsAwaitEvaluation),
	}
}
