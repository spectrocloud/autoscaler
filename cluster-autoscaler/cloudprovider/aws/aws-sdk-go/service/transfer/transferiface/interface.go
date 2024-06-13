// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package transferiface provides an interface to enable mocking the AWS Transfer Family service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package transferiface

import (
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws/request"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/service/transfer"
)

// TransferAPI provides an interface to enable mocking the
// transfer.Transfer service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//	// myFunc uses an SDK service client to make a request to
//	// AWS Transfer Family.
//	func myFunc(svc transferiface.TransferAPI) bool {
//	    // Make svc.CreateAccess request
//	}
//
//	func main() {
//	    sess := session.New()
//	    svc := transfer.New(sess)
//
//	    myFunc(svc)
//	}
//
// In your _test.go file:
//
//	// Define a mock struct to be used in your unit tests of myFunc.
//	type mockTransferClient struct {
//	    transferiface.TransferAPI
//	}
//	func (m *mockTransferClient) CreateAccess(input *transfer.CreateAccessInput) (*transfer.CreateAccessOutput, error) {
//	    // mock response/functionality
//	}
//
//	func TestMyFunc(t *testing.T) {
//	    // Setup Test
//	    mockSvc := &mockTransferClient{}
//
//	    myfunc(mockSvc)
//
//	    // Verify myFunc's functionality
//	}
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type TransferAPI interface {
	CreateAccess(*transfer.CreateAccessInput) (*transfer.CreateAccessOutput, error)
	CreateAccessWithContext(aws.Context, *transfer.CreateAccessInput, ...request.Option) (*transfer.CreateAccessOutput, error)
	CreateAccessRequest(*transfer.CreateAccessInput) (*request.Request, *transfer.CreateAccessOutput)

	CreateAgreement(*transfer.CreateAgreementInput) (*transfer.CreateAgreementOutput, error)
	CreateAgreementWithContext(aws.Context, *transfer.CreateAgreementInput, ...request.Option) (*transfer.CreateAgreementOutput, error)
	CreateAgreementRequest(*transfer.CreateAgreementInput) (*request.Request, *transfer.CreateAgreementOutput)

	CreateConnector(*transfer.CreateConnectorInput) (*transfer.CreateConnectorOutput, error)
	CreateConnectorWithContext(aws.Context, *transfer.CreateConnectorInput, ...request.Option) (*transfer.CreateConnectorOutput, error)
	CreateConnectorRequest(*transfer.CreateConnectorInput) (*request.Request, *transfer.CreateConnectorOutput)

	CreateProfile(*transfer.CreateProfileInput) (*transfer.CreateProfileOutput, error)
	CreateProfileWithContext(aws.Context, *transfer.CreateProfileInput, ...request.Option) (*transfer.CreateProfileOutput, error)
	CreateProfileRequest(*transfer.CreateProfileInput) (*request.Request, *transfer.CreateProfileOutput)

	CreateServer(*transfer.CreateServerInput) (*transfer.CreateServerOutput, error)
	CreateServerWithContext(aws.Context, *transfer.CreateServerInput, ...request.Option) (*transfer.CreateServerOutput, error)
	CreateServerRequest(*transfer.CreateServerInput) (*request.Request, *transfer.CreateServerOutput)

	CreateUser(*transfer.CreateUserInput) (*transfer.CreateUserOutput, error)
	CreateUserWithContext(aws.Context, *transfer.CreateUserInput, ...request.Option) (*transfer.CreateUserOutput, error)
	CreateUserRequest(*transfer.CreateUserInput) (*request.Request, *transfer.CreateUserOutput)

	CreateWorkflow(*transfer.CreateWorkflowInput) (*transfer.CreateWorkflowOutput, error)
	CreateWorkflowWithContext(aws.Context, *transfer.CreateWorkflowInput, ...request.Option) (*transfer.CreateWorkflowOutput, error)
	CreateWorkflowRequest(*transfer.CreateWorkflowInput) (*request.Request, *transfer.CreateWorkflowOutput)

	DeleteAccess(*transfer.DeleteAccessInput) (*transfer.DeleteAccessOutput, error)
	DeleteAccessWithContext(aws.Context, *transfer.DeleteAccessInput, ...request.Option) (*transfer.DeleteAccessOutput, error)
	DeleteAccessRequest(*transfer.DeleteAccessInput) (*request.Request, *transfer.DeleteAccessOutput)

	DeleteAgreement(*transfer.DeleteAgreementInput) (*transfer.DeleteAgreementOutput, error)
	DeleteAgreementWithContext(aws.Context, *transfer.DeleteAgreementInput, ...request.Option) (*transfer.DeleteAgreementOutput, error)
	DeleteAgreementRequest(*transfer.DeleteAgreementInput) (*request.Request, *transfer.DeleteAgreementOutput)

	DeleteCertificate(*transfer.DeleteCertificateInput) (*transfer.DeleteCertificateOutput, error)
	DeleteCertificateWithContext(aws.Context, *transfer.DeleteCertificateInput, ...request.Option) (*transfer.DeleteCertificateOutput, error)
	DeleteCertificateRequest(*transfer.DeleteCertificateInput) (*request.Request, *transfer.DeleteCertificateOutput)

	DeleteConnector(*transfer.DeleteConnectorInput) (*transfer.DeleteConnectorOutput, error)
	DeleteConnectorWithContext(aws.Context, *transfer.DeleteConnectorInput, ...request.Option) (*transfer.DeleteConnectorOutput, error)
	DeleteConnectorRequest(*transfer.DeleteConnectorInput) (*request.Request, *transfer.DeleteConnectorOutput)

	DeleteHostKey(*transfer.DeleteHostKeyInput) (*transfer.DeleteHostKeyOutput, error)
	DeleteHostKeyWithContext(aws.Context, *transfer.DeleteHostKeyInput, ...request.Option) (*transfer.DeleteHostKeyOutput, error)
	DeleteHostKeyRequest(*transfer.DeleteHostKeyInput) (*request.Request, *transfer.DeleteHostKeyOutput)

	DeleteProfile(*transfer.DeleteProfileInput) (*transfer.DeleteProfileOutput, error)
	DeleteProfileWithContext(aws.Context, *transfer.DeleteProfileInput, ...request.Option) (*transfer.DeleteProfileOutput, error)
	DeleteProfileRequest(*transfer.DeleteProfileInput) (*request.Request, *transfer.DeleteProfileOutput)

	DeleteServer(*transfer.DeleteServerInput) (*transfer.DeleteServerOutput, error)
	DeleteServerWithContext(aws.Context, *transfer.DeleteServerInput, ...request.Option) (*transfer.DeleteServerOutput, error)
	DeleteServerRequest(*transfer.DeleteServerInput) (*request.Request, *transfer.DeleteServerOutput)

	DeleteSshPublicKey(*transfer.DeleteSshPublicKeyInput) (*transfer.DeleteSshPublicKeyOutput, error)
	DeleteSshPublicKeyWithContext(aws.Context, *transfer.DeleteSshPublicKeyInput, ...request.Option) (*transfer.DeleteSshPublicKeyOutput, error)
	DeleteSshPublicKeyRequest(*transfer.DeleteSshPublicKeyInput) (*request.Request, *transfer.DeleteSshPublicKeyOutput)

	DeleteUser(*transfer.DeleteUserInput) (*transfer.DeleteUserOutput, error)
	DeleteUserWithContext(aws.Context, *transfer.DeleteUserInput, ...request.Option) (*transfer.DeleteUserOutput, error)
	DeleteUserRequest(*transfer.DeleteUserInput) (*request.Request, *transfer.DeleteUserOutput)

	DeleteWorkflow(*transfer.DeleteWorkflowInput) (*transfer.DeleteWorkflowOutput, error)
	DeleteWorkflowWithContext(aws.Context, *transfer.DeleteWorkflowInput, ...request.Option) (*transfer.DeleteWorkflowOutput, error)
	DeleteWorkflowRequest(*transfer.DeleteWorkflowInput) (*request.Request, *transfer.DeleteWorkflowOutput)

	DescribeAccess(*transfer.DescribeAccessInput) (*transfer.DescribeAccessOutput, error)
	DescribeAccessWithContext(aws.Context, *transfer.DescribeAccessInput, ...request.Option) (*transfer.DescribeAccessOutput, error)
	DescribeAccessRequest(*transfer.DescribeAccessInput) (*request.Request, *transfer.DescribeAccessOutput)

	DescribeAgreement(*transfer.DescribeAgreementInput) (*transfer.DescribeAgreementOutput, error)
	DescribeAgreementWithContext(aws.Context, *transfer.DescribeAgreementInput, ...request.Option) (*transfer.DescribeAgreementOutput, error)
	DescribeAgreementRequest(*transfer.DescribeAgreementInput) (*request.Request, *transfer.DescribeAgreementOutput)

	DescribeCertificate(*transfer.DescribeCertificateInput) (*transfer.DescribeCertificateOutput, error)
	DescribeCertificateWithContext(aws.Context, *transfer.DescribeCertificateInput, ...request.Option) (*transfer.DescribeCertificateOutput, error)
	DescribeCertificateRequest(*transfer.DescribeCertificateInput) (*request.Request, *transfer.DescribeCertificateOutput)

	DescribeConnector(*transfer.DescribeConnectorInput) (*transfer.DescribeConnectorOutput, error)
	DescribeConnectorWithContext(aws.Context, *transfer.DescribeConnectorInput, ...request.Option) (*transfer.DescribeConnectorOutput, error)
	DescribeConnectorRequest(*transfer.DescribeConnectorInput) (*request.Request, *transfer.DescribeConnectorOutput)

	DescribeExecution(*transfer.DescribeExecutionInput) (*transfer.DescribeExecutionOutput, error)
	DescribeExecutionWithContext(aws.Context, *transfer.DescribeExecutionInput, ...request.Option) (*transfer.DescribeExecutionOutput, error)
	DescribeExecutionRequest(*transfer.DescribeExecutionInput) (*request.Request, *transfer.DescribeExecutionOutput)

	DescribeHostKey(*transfer.DescribeHostKeyInput) (*transfer.DescribeHostKeyOutput, error)
	DescribeHostKeyWithContext(aws.Context, *transfer.DescribeHostKeyInput, ...request.Option) (*transfer.DescribeHostKeyOutput, error)
	DescribeHostKeyRequest(*transfer.DescribeHostKeyInput) (*request.Request, *transfer.DescribeHostKeyOutput)

	DescribeProfile(*transfer.DescribeProfileInput) (*transfer.DescribeProfileOutput, error)
	DescribeProfileWithContext(aws.Context, *transfer.DescribeProfileInput, ...request.Option) (*transfer.DescribeProfileOutput, error)
	DescribeProfileRequest(*transfer.DescribeProfileInput) (*request.Request, *transfer.DescribeProfileOutput)

	DescribeSecurityPolicy(*transfer.DescribeSecurityPolicyInput) (*transfer.DescribeSecurityPolicyOutput, error)
	DescribeSecurityPolicyWithContext(aws.Context, *transfer.DescribeSecurityPolicyInput, ...request.Option) (*transfer.DescribeSecurityPolicyOutput, error)
	DescribeSecurityPolicyRequest(*transfer.DescribeSecurityPolicyInput) (*request.Request, *transfer.DescribeSecurityPolicyOutput)

	DescribeServer(*transfer.DescribeServerInput) (*transfer.DescribeServerOutput, error)
	DescribeServerWithContext(aws.Context, *transfer.DescribeServerInput, ...request.Option) (*transfer.DescribeServerOutput, error)
	DescribeServerRequest(*transfer.DescribeServerInput) (*request.Request, *transfer.DescribeServerOutput)

	DescribeUser(*transfer.DescribeUserInput) (*transfer.DescribeUserOutput, error)
	DescribeUserWithContext(aws.Context, *transfer.DescribeUserInput, ...request.Option) (*transfer.DescribeUserOutput, error)
	DescribeUserRequest(*transfer.DescribeUserInput) (*request.Request, *transfer.DescribeUserOutput)

	DescribeWorkflow(*transfer.DescribeWorkflowInput) (*transfer.DescribeWorkflowOutput, error)
	DescribeWorkflowWithContext(aws.Context, *transfer.DescribeWorkflowInput, ...request.Option) (*transfer.DescribeWorkflowOutput, error)
	DescribeWorkflowRequest(*transfer.DescribeWorkflowInput) (*request.Request, *transfer.DescribeWorkflowOutput)

	ImportCertificate(*transfer.ImportCertificateInput) (*transfer.ImportCertificateOutput, error)
	ImportCertificateWithContext(aws.Context, *transfer.ImportCertificateInput, ...request.Option) (*transfer.ImportCertificateOutput, error)
	ImportCertificateRequest(*transfer.ImportCertificateInput) (*request.Request, *transfer.ImportCertificateOutput)

	ImportHostKey(*transfer.ImportHostKeyInput) (*transfer.ImportHostKeyOutput, error)
	ImportHostKeyWithContext(aws.Context, *transfer.ImportHostKeyInput, ...request.Option) (*transfer.ImportHostKeyOutput, error)
	ImportHostKeyRequest(*transfer.ImportHostKeyInput) (*request.Request, *transfer.ImportHostKeyOutput)

	ImportSshPublicKey(*transfer.ImportSshPublicKeyInput) (*transfer.ImportSshPublicKeyOutput, error)
	ImportSshPublicKeyWithContext(aws.Context, *transfer.ImportSshPublicKeyInput, ...request.Option) (*transfer.ImportSshPublicKeyOutput, error)
	ImportSshPublicKeyRequest(*transfer.ImportSshPublicKeyInput) (*request.Request, *transfer.ImportSshPublicKeyOutput)

	ListAccesses(*transfer.ListAccessesInput) (*transfer.ListAccessesOutput, error)
	ListAccessesWithContext(aws.Context, *transfer.ListAccessesInput, ...request.Option) (*transfer.ListAccessesOutput, error)
	ListAccessesRequest(*transfer.ListAccessesInput) (*request.Request, *transfer.ListAccessesOutput)

	ListAccessesPages(*transfer.ListAccessesInput, func(*transfer.ListAccessesOutput, bool) bool) error
	ListAccessesPagesWithContext(aws.Context, *transfer.ListAccessesInput, func(*transfer.ListAccessesOutput, bool) bool, ...request.Option) error

	ListAgreements(*transfer.ListAgreementsInput) (*transfer.ListAgreementsOutput, error)
	ListAgreementsWithContext(aws.Context, *transfer.ListAgreementsInput, ...request.Option) (*transfer.ListAgreementsOutput, error)
	ListAgreementsRequest(*transfer.ListAgreementsInput) (*request.Request, *transfer.ListAgreementsOutput)

	ListAgreementsPages(*transfer.ListAgreementsInput, func(*transfer.ListAgreementsOutput, bool) bool) error
	ListAgreementsPagesWithContext(aws.Context, *transfer.ListAgreementsInput, func(*transfer.ListAgreementsOutput, bool) bool, ...request.Option) error

	ListCertificates(*transfer.ListCertificatesInput) (*transfer.ListCertificatesOutput, error)
	ListCertificatesWithContext(aws.Context, *transfer.ListCertificatesInput, ...request.Option) (*transfer.ListCertificatesOutput, error)
	ListCertificatesRequest(*transfer.ListCertificatesInput) (*request.Request, *transfer.ListCertificatesOutput)

	ListCertificatesPages(*transfer.ListCertificatesInput, func(*transfer.ListCertificatesOutput, bool) bool) error
	ListCertificatesPagesWithContext(aws.Context, *transfer.ListCertificatesInput, func(*transfer.ListCertificatesOutput, bool) bool, ...request.Option) error

	ListConnectors(*transfer.ListConnectorsInput) (*transfer.ListConnectorsOutput, error)
	ListConnectorsWithContext(aws.Context, *transfer.ListConnectorsInput, ...request.Option) (*transfer.ListConnectorsOutput, error)
	ListConnectorsRequest(*transfer.ListConnectorsInput) (*request.Request, *transfer.ListConnectorsOutput)

	ListConnectorsPages(*transfer.ListConnectorsInput, func(*transfer.ListConnectorsOutput, bool) bool) error
	ListConnectorsPagesWithContext(aws.Context, *transfer.ListConnectorsInput, func(*transfer.ListConnectorsOutput, bool) bool, ...request.Option) error

	ListExecutions(*transfer.ListExecutionsInput) (*transfer.ListExecutionsOutput, error)
	ListExecutionsWithContext(aws.Context, *transfer.ListExecutionsInput, ...request.Option) (*transfer.ListExecutionsOutput, error)
	ListExecutionsRequest(*transfer.ListExecutionsInput) (*request.Request, *transfer.ListExecutionsOutput)

	ListExecutionsPages(*transfer.ListExecutionsInput, func(*transfer.ListExecutionsOutput, bool) bool) error
	ListExecutionsPagesWithContext(aws.Context, *transfer.ListExecutionsInput, func(*transfer.ListExecutionsOutput, bool) bool, ...request.Option) error

	ListHostKeys(*transfer.ListHostKeysInput) (*transfer.ListHostKeysOutput, error)
	ListHostKeysWithContext(aws.Context, *transfer.ListHostKeysInput, ...request.Option) (*transfer.ListHostKeysOutput, error)
	ListHostKeysRequest(*transfer.ListHostKeysInput) (*request.Request, *transfer.ListHostKeysOutput)

	ListProfiles(*transfer.ListProfilesInput) (*transfer.ListProfilesOutput, error)
	ListProfilesWithContext(aws.Context, *transfer.ListProfilesInput, ...request.Option) (*transfer.ListProfilesOutput, error)
	ListProfilesRequest(*transfer.ListProfilesInput) (*request.Request, *transfer.ListProfilesOutput)

	ListProfilesPages(*transfer.ListProfilesInput, func(*transfer.ListProfilesOutput, bool) bool) error
	ListProfilesPagesWithContext(aws.Context, *transfer.ListProfilesInput, func(*transfer.ListProfilesOutput, bool) bool, ...request.Option) error

	ListSecurityPolicies(*transfer.ListSecurityPoliciesInput) (*transfer.ListSecurityPoliciesOutput, error)
	ListSecurityPoliciesWithContext(aws.Context, *transfer.ListSecurityPoliciesInput, ...request.Option) (*transfer.ListSecurityPoliciesOutput, error)
	ListSecurityPoliciesRequest(*transfer.ListSecurityPoliciesInput) (*request.Request, *transfer.ListSecurityPoliciesOutput)

	ListSecurityPoliciesPages(*transfer.ListSecurityPoliciesInput, func(*transfer.ListSecurityPoliciesOutput, bool) bool) error
	ListSecurityPoliciesPagesWithContext(aws.Context, *transfer.ListSecurityPoliciesInput, func(*transfer.ListSecurityPoliciesOutput, bool) bool, ...request.Option) error

	ListServers(*transfer.ListServersInput) (*transfer.ListServersOutput, error)
	ListServersWithContext(aws.Context, *transfer.ListServersInput, ...request.Option) (*transfer.ListServersOutput, error)
	ListServersRequest(*transfer.ListServersInput) (*request.Request, *transfer.ListServersOutput)

	ListServersPages(*transfer.ListServersInput, func(*transfer.ListServersOutput, bool) bool) error
	ListServersPagesWithContext(aws.Context, *transfer.ListServersInput, func(*transfer.ListServersOutput, bool) bool, ...request.Option) error

	ListTagsForResource(*transfer.ListTagsForResourceInput) (*transfer.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *transfer.ListTagsForResourceInput, ...request.Option) (*transfer.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*transfer.ListTagsForResourceInput) (*request.Request, *transfer.ListTagsForResourceOutput)

	ListTagsForResourcePages(*transfer.ListTagsForResourceInput, func(*transfer.ListTagsForResourceOutput, bool) bool) error
	ListTagsForResourcePagesWithContext(aws.Context, *transfer.ListTagsForResourceInput, func(*transfer.ListTagsForResourceOutput, bool) bool, ...request.Option) error

	ListUsers(*transfer.ListUsersInput) (*transfer.ListUsersOutput, error)
	ListUsersWithContext(aws.Context, *transfer.ListUsersInput, ...request.Option) (*transfer.ListUsersOutput, error)
	ListUsersRequest(*transfer.ListUsersInput) (*request.Request, *transfer.ListUsersOutput)

	ListUsersPages(*transfer.ListUsersInput, func(*transfer.ListUsersOutput, bool) bool) error
	ListUsersPagesWithContext(aws.Context, *transfer.ListUsersInput, func(*transfer.ListUsersOutput, bool) bool, ...request.Option) error

	ListWorkflows(*transfer.ListWorkflowsInput) (*transfer.ListWorkflowsOutput, error)
	ListWorkflowsWithContext(aws.Context, *transfer.ListWorkflowsInput, ...request.Option) (*transfer.ListWorkflowsOutput, error)
	ListWorkflowsRequest(*transfer.ListWorkflowsInput) (*request.Request, *transfer.ListWorkflowsOutput)

	ListWorkflowsPages(*transfer.ListWorkflowsInput, func(*transfer.ListWorkflowsOutput, bool) bool) error
	ListWorkflowsPagesWithContext(aws.Context, *transfer.ListWorkflowsInput, func(*transfer.ListWorkflowsOutput, bool) bool, ...request.Option) error

	SendWorkflowStepState(*transfer.SendWorkflowStepStateInput) (*transfer.SendWorkflowStepStateOutput, error)
	SendWorkflowStepStateWithContext(aws.Context, *transfer.SendWorkflowStepStateInput, ...request.Option) (*transfer.SendWorkflowStepStateOutput, error)
	SendWorkflowStepStateRequest(*transfer.SendWorkflowStepStateInput) (*request.Request, *transfer.SendWorkflowStepStateOutput)

	StartFileTransfer(*transfer.StartFileTransferInput) (*transfer.StartFileTransferOutput, error)
	StartFileTransferWithContext(aws.Context, *transfer.StartFileTransferInput, ...request.Option) (*transfer.StartFileTransferOutput, error)
	StartFileTransferRequest(*transfer.StartFileTransferInput) (*request.Request, *transfer.StartFileTransferOutput)

	StartServer(*transfer.StartServerInput) (*transfer.StartServerOutput, error)
	StartServerWithContext(aws.Context, *transfer.StartServerInput, ...request.Option) (*transfer.StartServerOutput, error)
	StartServerRequest(*transfer.StartServerInput) (*request.Request, *transfer.StartServerOutput)

	StopServer(*transfer.StopServerInput) (*transfer.StopServerOutput, error)
	StopServerWithContext(aws.Context, *transfer.StopServerInput, ...request.Option) (*transfer.StopServerOutput, error)
	StopServerRequest(*transfer.StopServerInput) (*request.Request, *transfer.StopServerOutput)

	TagResource(*transfer.TagResourceInput) (*transfer.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *transfer.TagResourceInput, ...request.Option) (*transfer.TagResourceOutput, error)
	TagResourceRequest(*transfer.TagResourceInput) (*request.Request, *transfer.TagResourceOutput)

	TestConnection(*transfer.TestConnectionInput) (*transfer.TestConnectionOutput, error)
	TestConnectionWithContext(aws.Context, *transfer.TestConnectionInput, ...request.Option) (*transfer.TestConnectionOutput, error)
	TestConnectionRequest(*transfer.TestConnectionInput) (*request.Request, *transfer.TestConnectionOutput)

	TestIdentityProvider(*transfer.TestIdentityProviderInput) (*transfer.TestIdentityProviderOutput, error)
	TestIdentityProviderWithContext(aws.Context, *transfer.TestIdentityProviderInput, ...request.Option) (*transfer.TestIdentityProviderOutput, error)
	TestIdentityProviderRequest(*transfer.TestIdentityProviderInput) (*request.Request, *transfer.TestIdentityProviderOutput)

	UntagResource(*transfer.UntagResourceInput) (*transfer.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *transfer.UntagResourceInput, ...request.Option) (*transfer.UntagResourceOutput, error)
	UntagResourceRequest(*transfer.UntagResourceInput) (*request.Request, *transfer.UntagResourceOutput)

	UpdateAccess(*transfer.UpdateAccessInput) (*transfer.UpdateAccessOutput, error)
	UpdateAccessWithContext(aws.Context, *transfer.UpdateAccessInput, ...request.Option) (*transfer.UpdateAccessOutput, error)
	UpdateAccessRequest(*transfer.UpdateAccessInput) (*request.Request, *transfer.UpdateAccessOutput)

	UpdateAgreement(*transfer.UpdateAgreementInput) (*transfer.UpdateAgreementOutput, error)
	UpdateAgreementWithContext(aws.Context, *transfer.UpdateAgreementInput, ...request.Option) (*transfer.UpdateAgreementOutput, error)
	UpdateAgreementRequest(*transfer.UpdateAgreementInput) (*request.Request, *transfer.UpdateAgreementOutput)

	UpdateCertificate(*transfer.UpdateCertificateInput) (*transfer.UpdateCertificateOutput, error)
	UpdateCertificateWithContext(aws.Context, *transfer.UpdateCertificateInput, ...request.Option) (*transfer.UpdateCertificateOutput, error)
	UpdateCertificateRequest(*transfer.UpdateCertificateInput) (*request.Request, *transfer.UpdateCertificateOutput)

	UpdateConnector(*transfer.UpdateConnectorInput) (*transfer.UpdateConnectorOutput, error)
	UpdateConnectorWithContext(aws.Context, *transfer.UpdateConnectorInput, ...request.Option) (*transfer.UpdateConnectorOutput, error)
	UpdateConnectorRequest(*transfer.UpdateConnectorInput) (*request.Request, *transfer.UpdateConnectorOutput)

	UpdateHostKey(*transfer.UpdateHostKeyInput) (*transfer.UpdateHostKeyOutput, error)
	UpdateHostKeyWithContext(aws.Context, *transfer.UpdateHostKeyInput, ...request.Option) (*transfer.UpdateHostKeyOutput, error)
	UpdateHostKeyRequest(*transfer.UpdateHostKeyInput) (*request.Request, *transfer.UpdateHostKeyOutput)

	UpdateProfile(*transfer.UpdateProfileInput) (*transfer.UpdateProfileOutput, error)
	UpdateProfileWithContext(aws.Context, *transfer.UpdateProfileInput, ...request.Option) (*transfer.UpdateProfileOutput, error)
	UpdateProfileRequest(*transfer.UpdateProfileInput) (*request.Request, *transfer.UpdateProfileOutput)

	UpdateServer(*transfer.UpdateServerInput) (*transfer.UpdateServerOutput, error)
	UpdateServerWithContext(aws.Context, *transfer.UpdateServerInput, ...request.Option) (*transfer.UpdateServerOutput, error)
	UpdateServerRequest(*transfer.UpdateServerInput) (*request.Request, *transfer.UpdateServerOutput)

	UpdateUser(*transfer.UpdateUserInput) (*transfer.UpdateUserOutput, error)
	UpdateUserWithContext(aws.Context, *transfer.UpdateUserInput, ...request.Option) (*transfer.UpdateUserOutput, error)
	UpdateUserRequest(*transfer.UpdateUserInput) (*request.Request, *transfer.UpdateUserOutput)

	WaitUntilServerOffline(*transfer.DescribeServerInput) error
	WaitUntilServerOfflineWithContext(aws.Context, *transfer.DescribeServerInput, ...request.WaiterOption) error

	WaitUntilServerOnline(*transfer.DescribeServerInput) error
	WaitUntilServerOnlineWithContext(aws.Context, *transfer.DescribeServerInput, ...request.WaiterOption) error
}

var _ TransferAPI = (*transfer.Transfer)(nil)
