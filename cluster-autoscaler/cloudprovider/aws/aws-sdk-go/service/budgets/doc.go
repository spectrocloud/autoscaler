// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package budgets provides the client and types for making API
// requests to AWS Budgets.
//
// The Amazon Web Services Budgets API enables you to use Amazon Web Services
// Budgets to plan your service usage, service costs, and instance reservations.
// The API reference provides descriptions, syntax, and usage examples for each
// of the actions and data types for Amazon Web Services Budgets.
//
// Budgets provide you with a way to see the following information:
//
//   - How close your plan is to your budgeted amount or to the free tier limits
//
//   - Your usage-to-date, including how much you've used of your Reserved
//     Instances (RIs)
//
//   - Your current estimated charges from Amazon Web Services, and how much
//     your predicted usage will accrue in charges by the end of the month
//
//   - How much of your budget has been used
//
// Amazon Web Services updates your budget status several times a day. Budgets
// track your unblended costs, subscriptions, refunds, and RIs. You can create
// the following types of budgets:
//
//   - Cost budgets - Plan how much you want to spend on a service.
//
//   - Usage budgets - Plan how much you want to use one or more services.
//
//   - RI utilization budgets - Define a utilization threshold, and receive
//     alerts when your RI usage falls below that threshold. This lets you see
//     if your RIs are unused or under-utilized.
//
//   - RI coverage budgets - Define a coverage threshold, and receive alerts
//     when the number of your instance hours that are covered by RIs fall below
//     that threshold. This lets you see how much of your instance usage is covered
//     by a reservation.
//
// # Service Endpoint
//
// The Amazon Web Services Budgets API provides the following endpoint:
//
//   - https://budgets.amazonaws.com
//
// For information about costs that are associated with the Amazon Web Services
// Budgets API, see Amazon Web Services Cost Management Pricing (https://aws.amazon.com/aws-cost-management/pricing/).
//
// See budgets package documentation for more information.
// https://docs.aws.amazon.com/sdk-for-go/api/service/budgets/
//
// # Using the Client
//
// To contact AWS Budgets with the SDK use the New function to create
// a new service client. With that client you can make API requests to the service.
// These clients are safe to use concurrently.
//
// See the SDK's documentation for more information on how to use the SDK.
// https://docs.aws.amazon.com/sdk-for-go/api/
//
// See aws.Config documentation for more information on configuring SDK clients.
// https://docs.aws.amazon.com/sdk-for-go/api/aws/#Config
//
// See the AWS Budgets client Budgets for more
// information on creating client for this service.
// https://docs.aws.amazon.com/sdk-for-go/api/service/budgets/#New
package budgets