// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package route53recoverycluster provides the client and types for making API
// requests to Route53 Recovery Cluster.
//
// Welcome to the Routing Control (Recovery Cluster) API Reference Guide for
// Amazon Route 53 Application Recovery Controller.
//
// With Route 53 ARC, you can use routing control with extreme reliability to
// recover applications by rerouting traffic across Availability Zones or Amazon
// Web Services Regions. Routing controls are simple on/off switches hosted
// on a highly available cluster in Route 53 ARC. A cluster provides a set of
// five redundant Regional endpoints against which you can run API calls to
// get or update the state of routing controls. To implement failover, you set
// one routing control On and another one Off, to reroute traffic from one Availability
// Zone or Amazon Web Services Region to another.
//
// Be aware that you must specify a Regional endpoint for a cluster when you
// work with API cluster operations to get or update routing control states
// in Route 53 ARC. In addition, you must specify the US West (Oregon) Region
// for Route 53 ARC API calls. For example, use the parameter --region us-west-2
// with AWS CLI commands. For more information, see Get and update routing control
// states using the API (https://docs.aws.amazon.com/r53recovery/latest/dg/routing-control.update.api.html)
// in the Amazon Route 53 Application Recovery Controller Developer Guide.
//
// This API guide includes information about the API operations for how to get
// and update routing control states in Route 53 ARC. To work with routing control
// in Route 53 ARC, you must first create the required components (clusters,
// control panels, and routing controls) using the recovery cluster configuration
// API.
//
// For more information about working with routing control in Route 53 ARC,
// see the following:
//
//   - Create clusters, control panels, and routing controls by using API operations.
//     For more information, see the Recovery Control Configuration API Reference
//     Guide for Amazon Route 53 Application Recovery Controller (https://docs.aws.amazon.com/recovery-cluster/latest/api/).
//
//   - Learn about the components in recovery control, including clusters,
//     routing controls, and control panels, and how to work with Route 53 ARC
//     in the Amazon Web Services console. For more information, see Recovery
//     control components (https://docs.aws.amazon.com/r53recovery/latest/dg/introduction-components.html#introduction-components-routing)
//     in the Amazon Route 53 Application Recovery Controller Developer Guide.
//
//   - Route 53 ARC also provides readiness checks that continually audit resources
//     to help make sure that your applications are scaled and ready to handle
//     failover traffic. For more information about the related API operations,
//     see the Recovery Readiness API Reference Guide for Amazon Route 53 Application
//     Recovery Controller (https://docs.aws.amazon.com/recovery-readiness/latest/api/).
//
//   - For more information about creating resilient applications and preparing
//     for recovery readiness with Route 53 ARC, see the Amazon Route 53 Application
//     Recovery Controller Developer Guide (https://docs.aws.amazon.com/r53recovery/latest/dg/).
//
// See https://docs.aws.amazon.com/goto/WebAPI/route53-recovery-cluster-2019-12-02 for more information on this service.
//
// See route53recoverycluster package documentation for more information.
// https://docs.aws.amazon.com/sdk-for-go/api/service/route53recoverycluster/
//
// # Using the Client
//
// To contact Route53 Recovery Cluster with the SDK use the New function to create
// a new service client. With that client you can make API requests to the service.
// These clients are safe to use concurrently.
//
// See the SDK's documentation for more information on how to use the SDK.
// https://docs.aws.amazon.com/sdk-for-go/api/
//
// See aws.Config documentation for more information on configuring SDK clients.
// https://docs.aws.amazon.com/sdk-for-go/api/aws/#Config
//
// See the Route53 Recovery Cluster client Route53RecoveryCluster for more
// information on creating client for this service.
// https://docs.aws.amazon.com/sdk-for-go/api/service/route53recoverycluster/#New
package route53recoverycluster
