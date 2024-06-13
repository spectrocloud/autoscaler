/*
Copyright 2015 The Kubernetes Authors.

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

package simulator

import (
	"fmt"
	"time"

	apiv1 "k8s.io/api/core/v1"
	policyv1 "k8s.io/api/policy/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/autoscaler/cluster-autoscaler/config"
	"k8s.io/autoscaler/cluster-autoscaler/simulator/drainability"
	"k8s.io/autoscaler/cluster-autoscaler/utils/drain"
	kube_util "k8s.io/autoscaler/cluster-autoscaler/utils/kubernetes"
	pod_util "k8s.io/autoscaler/cluster-autoscaler/utils/pod"
	schedulerframework "k8s.io/kubernetes/pkg/scheduler/framework"
)

// NodeDeleteOptions contains various options to customize how draining will behave
type NodeDeleteOptions struct {
	// SkipNodesWithSystemPods tells if nodes with pods from kube-system should be deleted (except for DaemonSet or mirror pods)
	SkipNodesWithSystemPods bool
	// SkipNodesWithLocalStorage tells if nodes with pods with local storage, e.g. EmptyDir or HostPath, should be deleted
	SkipNodesWithLocalStorage bool
	// SkipNodesWithCustomControllerPods tells if nodes with custom-controller owned pods should be skipped from deletion (skip if 'true')
	SkipNodesWithCustomControllerPods bool
	// MinReplicaCount controls the minimum number of replicas that a replica set or replication controller should have
	// to allow their pods deletion in scale down
	MinReplicaCount int
	// DrainabilityRules contain a list of checks that are used to verify whether a pod can be drained from node.
	DrainabilityRules []drainability.Rule
}

// NewNodeDeleteOptions returns new node delete options extracted from autoscaling options
func NewNodeDeleteOptions(opts config.AutoscalingOptions) NodeDeleteOptions {
	return NodeDeleteOptions{
		SkipNodesWithSystemPods:           opts.SkipNodesWithSystemPods,
		SkipNodesWithLocalStorage:         opts.SkipNodesWithLocalStorage,
		MinReplicaCount:                   opts.MinReplicaCount,
		SkipNodesWithCustomControllerPods: opts.SkipNodesWithCustomControllerPods,
	}
}

// GetPodsToMove returns a list of pods that should be moved elsewhere
// and a list of DaemonSet pods that should be evicted if the node
// is drained. Raises error if there is an unreplicated pod.
// Based on kubectl drain code. If listers is nil it makes an assumption that RC, DS, Jobs and RS were deleted
// along with their pods (no abandoned pods with dangling created-by annotation).
// If listers is not nil it checks whether RC, DS, Jobs and RS that created these pods
// still exist.
// TODO(x13n): Rewrite GetPodsForDeletionOnNodeDrain into a set of DrainabilityRules.
func GetPodsToMove(nodeInfo *schedulerframework.NodeInfo, deleteOptions NodeDeleteOptions, listers kube_util.ListerRegistry,
	pdbs []*policyv1.PodDisruptionBudget, timestamp time.Time) (pods []*apiv1.Pod, daemonSetPods []*apiv1.Pod, blockingPod *drain.BlockingPod, err error) {
	var drainPods, drainDs []*apiv1.Pod
	for _, podInfo := range nodeInfo.Pods {
		pod := podInfo.Pod
		d := drainabilityStatus(pod, deleteOptions.DrainabilityRules)
		if d.Matched {
			switch d.Reason {
			case drain.NoReason:
				if pod_util.IsDaemonSetPod(pod) {
					drainDs = append(drainDs, pod)
				} else {
					drainPods = append(drainPods, pod)
				}
				continue
			default:
				blockingPod = &drain.BlockingPod{pod, d.Reason}
				err = d.Error
				return
			}
		}
		pods = append(pods, podInfo.Pod)
	}
	pods, daemonSetPods, blockingPod, err = drain.GetPodsForDeletionOnNodeDrain(
		pods,
		pdbs,
		deleteOptions.SkipNodesWithSystemPods,
		deleteOptions.SkipNodesWithLocalStorage,
		deleteOptions.SkipNodesWithCustomControllerPods,
		listers,
		int32(deleteOptions.MinReplicaCount),
		timestamp)
	pods = append(pods, drainPods...)
	daemonSetPods = append(daemonSetPods, drainDs...)
	if err != nil {
		return pods, daemonSetPods, blockingPod, err
	}
	if pdbBlockingPod, err := checkPdbs(pods, pdbs); err != nil {
		return []*apiv1.Pod{}, []*apiv1.Pod{}, pdbBlockingPod, err
	}

	return pods, daemonSetPods, nil, nil
}

func checkPdbs(pods []*apiv1.Pod, pdbs []*policyv1.PodDisruptionBudget) (*drain.BlockingPod, error) {
	// TODO: remove it after deprecating legacy scale down.
	// RemainingPdbTracker.CanRemovePods() to replace this function.
	for _, pdb := range pdbs {
		selector, err := metav1.LabelSelectorAsSelector(pdb.Spec.Selector)
		if err != nil {
			return nil, err
		}
		for _, pod := range pods {
			if pod.Namespace == pdb.Namespace && selector.Matches(labels.Set(pod.Labels)) {
				if pdb.Status.DisruptionsAllowed < 1 {
					return &drain.BlockingPod{Pod: pod, Reason: drain.NotEnoughPdb}, fmt.Errorf("not enough pod disruption budget to move %s/%s", pod.Namespace, pod.Name)
				}
			}
		}
	}
	return nil, nil
}

func drainabilityStatus(pod *apiv1.Pod, dr []drainability.Rule) drainability.Status {
	for _, f := range dr {
		if d := f.Drainable(pod); d.Matched {
			return d
		}
	}
	return drainability.Status{}
}
