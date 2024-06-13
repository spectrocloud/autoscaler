// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2instanceconnect

import (
	"fmt"

	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws/awsutil"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws/request"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/private/protocol"
)

const opSendSSHPublicKey = "SendSSHPublicKey"

// SendSSHPublicKeyRequest generates a "aws/request.Request" representing the
// client's request for the SendSSHPublicKey operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See SendSSHPublicKey for more information on using the SendSSHPublicKey
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//	// Example sending a request using the SendSSHPublicKeyRequest method.
//	req, resp := client.SendSSHPublicKeyRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/ec2-instance-connect-2018-04-02/SendSSHPublicKey
func (c *EC2InstanceConnect) SendSSHPublicKeyRequest(input *SendSSHPublicKeyInput) (req *request.Request, output *SendSSHPublicKeyOutput) {
	op := &request.Operation{
		Name:       opSendSSHPublicKey,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &SendSSHPublicKeyInput{}
	}

	output = &SendSSHPublicKeyOutput{}
	req = c.newRequest(op, input, output)
	return
}

// SendSSHPublicKey API operation for AWS EC2 Instance Connect.
//
// Pushes an SSH public key to the specified EC2 instance for use by the specified
// user. The key remains for 60 seconds. For more information, see Connect to
// your Linux instance using EC2 Instance Connect (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Connect-using-EC2-Instance-Connect.html)
// in the Amazon EC2 User Guide.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for AWS EC2 Instance Connect's
// API operation SendSSHPublicKey for usage and error information.
//
// Returned Error Types:
//
//   - AuthException
//     Either your AWS credentials are not valid or you do not have access to the
//     EC2 instance.
//
//   - InvalidArgsException
//     One of the parameters is not valid.
//
//   - ServiceException
//     The service encountered an error. Follow the instructions in the error message
//     and try again.
//
//   - ThrottlingException
//     The requests were made too frequently and have been throttled. Wait a while
//     and try again. To increase the limit on your request frequency, contact AWS
//     Support.
//
//   - EC2InstanceNotFoundException
//     The specified instance was not found.
//
//   - EC2InstanceStateInvalidException
//     Unable to connect because the instance is not in a valid state. Connecting
//     to a stopped or terminated instance is not supported. If the instance is
//     stopped, start your instance, and try to connect again.
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/ec2-instance-connect-2018-04-02/SendSSHPublicKey
func (c *EC2InstanceConnect) SendSSHPublicKey(input *SendSSHPublicKeyInput) (*SendSSHPublicKeyOutput, error) {
	req, out := c.SendSSHPublicKeyRequest(input)
	return out, req.Send()
}

// SendSSHPublicKeyWithContext is the same as SendSSHPublicKey with the addition of
// the ability to pass a context and additional request options.
//
// See SendSSHPublicKey for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *EC2InstanceConnect) SendSSHPublicKeyWithContext(ctx aws.Context, input *SendSSHPublicKeyInput, opts ...request.Option) (*SendSSHPublicKeyOutput, error) {
	req, out := c.SendSSHPublicKeyRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opSendSerialConsoleSSHPublicKey = "SendSerialConsoleSSHPublicKey"

// SendSerialConsoleSSHPublicKeyRequest generates a "aws/request.Request" representing the
// client's request for the SendSerialConsoleSSHPublicKey operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See SendSerialConsoleSSHPublicKey for more information on using the SendSerialConsoleSSHPublicKey
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//	// Example sending a request using the SendSerialConsoleSSHPublicKeyRequest method.
//	req, resp := client.SendSerialConsoleSSHPublicKeyRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/ec2-instance-connect-2018-04-02/SendSerialConsoleSSHPublicKey
func (c *EC2InstanceConnect) SendSerialConsoleSSHPublicKeyRequest(input *SendSerialConsoleSSHPublicKeyInput) (req *request.Request, output *SendSerialConsoleSSHPublicKeyOutput) {
	op := &request.Operation{
		Name:       opSendSerialConsoleSSHPublicKey,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &SendSerialConsoleSSHPublicKeyInput{}
	}

	output = &SendSerialConsoleSSHPublicKeyOutput{}
	req = c.newRequest(op, input, output)
	return
}

// SendSerialConsoleSSHPublicKey API operation for AWS EC2 Instance Connect.
//
// Pushes an SSH public key to the specified EC2 instance. The key remains for
// 60 seconds, which gives you 60 seconds to establish a serial console connection
// to the instance using SSH. For more information, see EC2 Serial Console (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-serial-console.html)
// in the Amazon EC2 User Guide.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for AWS EC2 Instance Connect's
// API operation SendSerialConsoleSSHPublicKey for usage and error information.
//
// Returned Error Types:
//
//   - AuthException
//     Either your AWS credentials are not valid or you do not have access to the
//     EC2 instance.
//
//   - SerialConsoleAccessDisabledException
//     Your account is not authorized to use the EC2 Serial Console. To authorize
//     your account, run the EnableSerialConsoleAccess API. For more information,
//     see EnableSerialConsoleAccess (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_EnableSerialConsoleAccess.html)
//     in the Amazon EC2 API Reference.
//
//   - InvalidArgsException
//     One of the parameters is not valid.
//
//   - ServiceException
//     The service encountered an error. Follow the instructions in the error message
//     and try again.
//
//   - ThrottlingException
//     The requests were made too frequently and have been throttled. Wait a while
//     and try again. To increase the limit on your request frequency, contact AWS
//     Support.
//
//   - EC2InstanceNotFoundException
//     The specified instance was not found.
//
//   - EC2InstanceTypeInvalidException
//     The instance type is not supported for connecting via the serial console.
//     Only Nitro instance types are currently supported.
//
//   - SerialConsoleSessionLimitExceededException
//     The instance currently has 1 active serial console session. Only 1 session
//     is supported at a time.
//
//   - SerialConsoleSessionUnavailableException
//     Unable to start a serial console session. Please try again.
//
//   - EC2InstanceStateInvalidException
//     Unable to connect because the instance is not in a valid state. Connecting
//     to a stopped or terminated instance is not supported. If the instance is
//     stopped, start your instance, and try to connect again.
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/ec2-instance-connect-2018-04-02/SendSerialConsoleSSHPublicKey
func (c *EC2InstanceConnect) SendSerialConsoleSSHPublicKey(input *SendSerialConsoleSSHPublicKeyInput) (*SendSerialConsoleSSHPublicKeyOutput, error) {
	req, out := c.SendSerialConsoleSSHPublicKeyRequest(input)
	return out, req.Send()
}

// SendSerialConsoleSSHPublicKeyWithContext is the same as SendSerialConsoleSSHPublicKey with the addition of
// the ability to pass a context and additional request options.
//
// See SendSerialConsoleSSHPublicKey for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *EC2InstanceConnect) SendSerialConsoleSSHPublicKeyWithContext(ctx aws.Context, input *SendSerialConsoleSSHPublicKeyInput, opts ...request.Option) (*SendSerialConsoleSSHPublicKeyOutput, error) {
	req, out := c.SendSerialConsoleSSHPublicKeyRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

// Either your AWS credentials are not valid or you do not have access to the
// EC2 instance.
type AuthException struct {
	_            struct{}                  `type:"structure"`
	RespMetadata protocol.ResponseMetadata `json:"-" xml:"-"`

	Message_ *string `locationName:"Message" type:"string"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s AuthException) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s AuthException) GoString() string {
	return s.String()
}

func newErrorAuthException(v protocol.ResponseMetadata) error {
	return &AuthException{
		RespMetadata: v,
	}
}

// Code returns the exception type name.
func (s *AuthException) Code() string {
	return "AuthException"
}

// Message returns the exception's message.
func (s *AuthException) Message() string {
	if s.Message_ != nil {
		return *s.Message_
	}
	return ""
}

// OrigErr always returns nil, satisfies awserr.Error interface.
func (s *AuthException) OrigErr() error {
	return nil
}

func (s *AuthException) Error() string {
	return fmt.Sprintf("%s: %s", s.Code(), s.Message())
}

// Status code returns the HTTP status code for the request's response error.
func (s *AuthException) StatusCode() int {
	return s.RespMetadata.StatusCode
}

// RequestID returns the service's response RequestID for request.
func (s *AuthException) RequestID() string {
	return s.RespMetadata.RequestID
}

// The specified instance was not found.
type EC2InstanceNotFoundException struct {
	_            struct{}                  `type:"structure"`
	RespMetadata protocol.ResponseMetadata `json:"-" xml:"-"`

	Message_ *string `locationName:"Message" type:"string"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s EC2InstanceNotFoundException) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s EC2InstanceNotFoundException) GoString() string {
	return s.String()
}

func newErrorEC2InstanceNotFoundException(v protocol.ResponseMetadata) error {
	return &EC2InstanceNotFoundException{
		RespMetadata: v,
	}
}

// Code returns the exception type name.
func (s *EC2InstanceNotFoundException) Code() string {
	return "EC2InstanceNotFoundException"
}

// Message returns the exception's message.
func (s *EC2InstanceNotFoundException) Message() string {
	if s.Message_ != nil {
		return *s.Message_
	}
	return ""
}

// OrigErr always returns nil, satisfies awserr.Error interface.
func (s *EC2InstanceNotFoundException) OrigErr() error {
	return nil
}

func (s *EC2InstanceNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", s.Code(), s.Message())
}

// Status code returns the HTTP status code for the request's response error.
func (s *EC2InstanceNotFoundException) StatusCode() int {
	return s.RespMetadata.StatusCode
}

// RequestID returns the service's response RequestID for request.
func (s *EC2InstanceNotFoundException) RequestID() string {
	return s.RespMetadata.RequestID
}

// Unable to connect because the instance is not in a valid state. Connecting
// to a stopped or terminated instance is not supported. If the instance is
// stopped, start your instance, and try to connect again.
type EC2InstanceStateInvalidException struct {
	_            struct{}                  `type:"structure"`
	RespMetadata protocol.ResponseMetadata `json:"-" xml:"-"`

	Message_ *string `locationName:"Message" type:"string"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s EC2InstanceStateInvalidException) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s EC2InstanceStateInvalidException) GoString() string {
	return s.String()
}

func newErrorEC2InstanceStateInvalidException(v protocol.ResponseMetadata) error {
	return &EC2InstanceStateInvalidException{
		RespMetadata: v,
	}
}

// Code returns the exception type name.
func (s *EC2InstanceStateInvalidException) Code() string {
	return "EC2InstanceStateInvalidException"
}

// Message returns the exception's message.
func (s *EC2InstanceStateInvalidException) Message() string {
	if s.Message_ != nil {
		return *s.Message_
	}
	return ""
}

// OrigErr always returns nil, satisfies awserr.Error interface.
func (s *EC2InstanceStateInvalidException) OrigErr() error {
	return nil
}

func (s *EC2InstanceStateInvalidException) Error() string {
	return fmt.Sprintf("%s: %s", s.Code(), s.Message())
}

// Status code returns the HTTP status code for the request's response error.
func (s *EC2InstanceStateInvalidException) StatusCode() int {
	return s.RespMetadata.StatusCode
}

// RequestID returns the service's response RequestID for request.
func (s *EC2InstanceStateInvalidException) RequestID() string {
	return s.RespMetadata.RequestID
}

// The instance type is not supported for connecting via the serial console.
// Only Nitro instance types are currently supported.
type EC2InstanceTypeInvalidException struct {
	_            struct{}                  `type:"structure"`
	RespMetadata protocol.ResponseMetadata `json:"-" xml:"-"`

	Message_ *string `locationName:"Message" type:"string"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s EC2InstanceTypeInvalidException) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s EC2InstanceTypeInvalidException) GoString() string {
	return s.String()
}

func newErrorEC2InstanceTypeInvalidException(v protocol.ResponseMetadata) error {
	return &EC2InstanceTypeInvalidException{
		RespMetadata: v,
	}
}

// Code returns the exception type name.
func (s *EC2InstanceTypeInvalidException) Code() string {
	return "EC2InstanceTypeInvalidException"
}

// Message returns the exception's message.
func (s *EC2InstanceTypeInvalidException) Message() string {
	if s.Message_ != nil {
		return *s.Message_
	}
	return ""
}

// OrigErr always returns nil, satisfies awserr.Error interface.
func (s *EC2InstanceTypeInvalidException) OrigErr() error {
	return nil
}

func (s *EC2InstanceTypeInvalidException) Error() string {
	return fmt.Sprintf("%s: %s", s.Code(), s.Message())
}

// Status code returns the HTTP status code for the request's response error.
func (s *EC2InstanceTypeInvalidException) StatusCode() int {
	return s.RespMetadata.StatusCode
}

// RequestID returns the service's response RequestID for request.
func (s *EC2InstanceTypeInvalidException) RequestID() string {
	return s.RespMetadata.RequestID
}

// One of the parameters is not valid.
type InvalidArgsException struct {
	_            struct{}                  `type:"structure"`
	RespMetadata protocol.ResponseMetadata `json:"-" xml:"-"`

	Message_ *string `locationName:"Message" type:"string"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s InvalidArgsException) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s InvalidArgsException) GoString() string {
	return s.String()
}

func newErrorInvalidArgsException(v protocol.ResponseMetadata) error {
	return &InvalidArgsException{
		RespMetadata: v,
	}
}

// Code returns the exception type name.
func (s *InvalidArgsException) Code() string {
	return "InvalidArgsException"
}

// Message returns the exception's message.
func (s *InvalidArgsException) Message() string {
	if s.Message_ != nil {
		return *s.Message_
	}
	return ""
}

// OrigErr always returns nil, satisfies awserr.Error interface.
func (s *InvalidArgsException) OrigErr() error {
	return nil
}

func (s *InvalidArgsException) Error() string {
	return fmt.Sprintf("%s: %s", s.Code(), s.Message())
}

// Status code returns the HTTP status code for the request's response error.
func (s *InvalidArgsException) StatusCode() int {
	return s.RespMetadata.StatusCode
}

// RequestID returns the service's response RequestID for request.
func (s *InvalidArgsException) RequestID() string {
	return s.RespMetadata.RequestID
}

type SendSSHPublicKeyInput struct {
	_ struct{} `type:"structure"`

	// The Availability Zone in which the EC2 instance was launched.
	AvailabilityZone *string `min:"6" type:"string"`

	// The ID of the EC2 instance.
	//
	// InstanceId is a required field
	InstanceId *string `min:"10" type:"string" required:"true"`

	// The OS user on the EC2 instance for whom the key can be used to authenticate.
	//
	// InstanceOSUser is a required field
	InstanceOSUser *string `min:"1" type:"string" required:"true"`

	// The public key material. To use the public key, you must have the matching
	// private key.
	//
	// SSHPublicKey is a required field
	SSHPublicKey *string `min:"80" type:"string" required:"true"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s SendSSHPublicKeyInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s SendSSHPublicKeyInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SendSSHPublicKeyInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "SendSSHPublicKeyInput"}
	if s.AvailabilityZone != nil && len(*s.AvailabilityZone) < 6 {
		invalidParams.Add(request.NewErrParamMinLen("AvailabilityZone", 6))
	}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}
	if s.InstanceId != nil && len(*s.InstanceId) < 10 {
		invalidParams.Add(request.NewErrParamMinLen("InstanceId", 10))
	}
	if s.InstanceOSUser == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceOSUser"))
	}
	if s.InstanceOSUser != nil && len(*s.InstanceOSUser) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("InstanceOSUser", 1))
	}
	if s.SSHPublicKey == nil {
		invalidParams.Add(request.NewErrParamRequired("SSHPublicKey"))
	}
	if s.SSHPublicKey != nil && len(*s.SSHPublicKey) < 80 {
		invalidParams.Add(request.NewErrParamMinLen("SSHPublicKey", 80))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAvailabilityZone sets the AvailabilityZone field's value.
func (s *SendSSHPublicKeyInput) SetAvailabilityZone(v string) *SendSSHPublicKeyInput {
	s.AvailabilityZone = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *SendSSHPublicKeyInput) SetInstanceId(v string) *SendSSHPublicKeyInput {
	s.InstanceId = &v
	return s
}

// SetInstanceOSUser sets the InstanceOSUser field's value.
func (s *SendSSHPublicKeyInput) SetInstanceOSUser(v string) *SendSSHPublicKeyInput {
	s.InstanceOSUser = &v
	return s
}

// SetSSHPublicKey sets the SSHPublicKey field's value.
func (s *SendSSHPublicKeyInput) SetSSHPublicKey(v string) *SendSSHPublicKeyInput {
	s.SSHPublicKey = &v
	return s
}

type SendSSHPublicKeyOutput struct {
	_ struct{} `type:"structure"`

	// The ID of the request. Please provide this ID when contacting AWS Support
	// for assistance.
	RequestId *string `type:"string"`

	// Is true if the request succeeds and an error otherwise.
	Success *bool `type:"boolean"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s SendSSHPublicKeyOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s SendSSHPublicKeyOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *SendSSHPublicKeyOutput) SetRequestId(v string) *SendSSHPublicKeyOutput {
	s.RequestId = &v
	return s
}

// SetSuccess sets the Success field's value.
func (s *SendSSHPublicKeyOutput) SetSuccess(v bool) *SendSSHPublicKeyOutput {
	s.Success = &v
	return s
}

type SendSerialConsoleSSHPublicKeyInput struct {
	_ struct{} `type:"structure"`

	// The ID of the EC2 instance.
	//
	// InstanceId is a required field
	InstanceId *string `min:"10" type:"string" required:"true"`

	// The public key material. To use the public key, you must have the matching
	// private key. For information about the supported key formats and lengths,
	// see Requirements for key pairs (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-key-pairs.html#how-to-generate-your-own-key-and-import-it-to-aws)
	// in the Amazon EC2 User Guide.
	//
	// SSHPublicKey is a required field
	SSHPublicKey *string `min:"80" type:"string" required:"true"`

	// The serial port of the EC2 instance. Currently only port 0 is supported.
	//
	// Default: 0
	SerialPort *int64 `type:"integer"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s SendSerialConsoleSSHPublicKeyInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s SendSerialConsoleSSHPublicKeyInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SendSerialConsoleSSHPublicKeyInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "SendSerialConsoleSSHPublicKeyInput"}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}
	if s.InstanceId != nil && len(*s.InstanceId) < 10 {
		invalidParams.Add(request.NewErrParamMinLen("InstanceId", 10))
	}
	if s.SSHPublicKey == nil {
		invalidParams.Add(request.NewErrParamRequired("SSHPublicKey"))
	}
	if s.SSHPublicKey != nil && len(*s.SSHPublicKey) < 80 {
		invalidParams.Add(request.NewErrParamMinLen("SSHPublicKey", 80))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetInstanceId sets the InstanceId field's value.
func (s *SendSerialConsoleSSHPublicKeyInput) SetInstanceId(v string) *SendSerialConsoleSSHPublicKeyInput {
	s.InstanceId = &v
	return s
}

// SetSSHPublicKey sets the SSHPublicKey field's value.
func (s *SendSerialConsoleSSHPublicKeyInput) SetSSHPublicKey(v string) *SendSerialConsoleSSHPublicKeyInput {
	s.SSHPublicKey = &v
	return s
}

// SetSerialPort sets the SerialPort field's value.
func (s *SendSerialConsoleSSHPublicKeyInput) SetSerialPort(v int64) *SendSerialConsoleSSHPublicKeyInput {
	s.SerialPort = &v
	return s
}

type SendSerialConsoleSSHPublicKeyOutput struct {
	_ struct{} `type:"structure"`

	// The ID of the request. Please provide this ID when contacting AWS Support
	// for assistance.
	RequestId *string `type:"string"`

	// Is true if the request succeeds and an error otherwise.
	Success *bool `type:"boolean"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s SendSerialConsoleSSHPublicKeyOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s SendSerialConsoleSSHPublicKeyOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *SendSerialConsoleSSHPublicKeyOutput) SetRequestId(v string) *SendSerialConsoleSSHPublicKeyOutput {
	s.RequestId = &v
	return s
}

// SetSuccess sets the Success field's value.
func (s *SendSerialConsoleSSHPublicKeyOutput) SetSuccess(v bool) *SendSerialConsoleSSHPublicKeyOutput {
	s.Success = &v
	return s
}

// Your account is not authorized to use the EC2 Serial Console. To authorize
// your account, run the EnableSerialConsoleAccess API. For more information,
// see EnableSerialConsoleAccess (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_EnableSerialConsoleAccess.html)
// in the Amazon EC2 API Reference.
type SerialConsoleAccessDisabledException struct {
	_            struct{}                  `type:"structure"`
	RespMetadata protocol.ResponseMetadata `json:"-" xml:"-"`

	Message_ *string `locationName:"Message" type:"string"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s SerialConsoleAccessDisabledException) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s SerialConsoleAccessDisabledException) GoString() string {
	return s.String()
}

func newErrorSerialConsoleAccessDisabledException(v protocol.ResponseMetadata) error {
	return &SerialConsoleAccessDisabledException{
		RespMetadata: v,
	}
}

// Code returns the exception type name.
func (s *SerialConsoleAccessDisabledException) Code() string {
	return "SerialConsoleAccessDisabledException"
}

// Message returns the exception's message.
func (s *SerialConsoleAccessDisabledException) Message() string {
	if s.Message_ != nil {
		return *s.Message_
	}
	return ""
}

// OrigErr always returns nil, satisfies awserr.Error interface.
func (s *SerialConsoleAccessDisabledException) OrigErr() error {
	return nil
}

func (s *SerialConsoleAccessDisabledException) Error() string {
	return fmt.Sprintf("%s: %s", s.Code(), s.Message())
}

// Status code returns the HTTP status code for the request's response error.
func (s *SerialConsoleAccessDisabledException) StatusCode() int {
	return s.RespMetadata.StatusCode
}

// RequestID returns the service's response RequestID for request.
func (s *SerialConsoleAccessDisabledException) RequestID() string {
	return s.RespMetadata.RequestID
}

// The instance currently has 1 active serial console session. Only 1 session
// is supported at a time.
type SerialConsoleSessionLimitExceededException struct {
	_            struct{}                  `type:"structure"`
	RespMetadata protocol.ResponseMetadata `json:"-" xml:"-"`

	Message_ *string `locationName:"Message" type:"string"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s SerialConsoleSessionLimitExceededException) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s SerialConsoleSessionLimitExceededException) GoString() string {
	return s.String()
}

func newErrorSerialConsoleSessionLimitExceededException(v protocol.ResponseMetadata) error {
	return &SerialConsoleSessionLimitExceededException{
		RespMetadata: v,
	}
}

// Code returns the exception type name.
func (s *SerialConsoleSessionLimitExceededException) Code() string {
	return "SerialConsoleSessionLimitExceededException"
}

// Message returns the exception's message.
func (s *SerialConsoleSessionLimitExceededException) Message() string {
	if s.Message_ != nil {
		return *s.Message_
	}
	return ""
}

// OrigErr always returns nil, satisfies awserr.Error interface.
func (s *SerialConsoleSessionLimitExceededException) OrigErr() error {
	return nil
}

func (s *SerialConsoleSessionLimitExceededException) Error() string {
	return fmt.Sprintf("%s: %s", s.Code(), s.Message())
}

// Status code returns the HTTP status code for the request's response error.
func (s *SerialConsoleSessionLimitExceededException) StatusCode() int {
	return s.RespMetadata.StatusCode
}

// RequestID returns the service's response RequestID for request.
func (s *SerialConsoleSessionLimitExceededException) RequestID() string {
	return s.RespMetadata.RequestID
}

// Unable to start a serial console session. Please try again.
type SerialConsoleSessionUnavailableException struct {
	_            struct{}                  `type:"structure"`
	RespMetadata protocol.ResponseMetadata `json:"-" xml:"-"`

	Message_ *string `locationName:"Message" type:"string"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s SerialConsoleSessionUnavailableException) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s SerialConsoleSessionUnavailableException) GoString() string {
	return s.String()
}

func newErrorSerialConsoleSessionUnavailableException(v protocol.ResponseMetadata) error {
	return &SerialConsoleSessionUnavailableException{
		RespMetadata: v,
	}
}

// Code returns the exception type name.
func (s *SerialConsoleSessionUnavailableException) Code() string {
	return "SerialConsoleSessionUnavailableException"
}

// Message returns the exception's message.
func (s *SerialConsoleSessionUnavailableException) Message() string {
	if s.Message_ != nil {
		return *s.Message_
	}
	return ""
}

// OrigErr always returns nil, satisfies awserr.Error interface.
func (s *SerialConsoleSessionUnavailableException) OrigErr() error {
	return nil
}

func (s *SerialConsoleSessionUnavailableException) Error() string {
	return fmt.Sprintf("%s: %s", s.Code(), s.Message())
}

// Status code returns the HTTP status code for the request's response error.
func (s *SerialConsoleSessionUnavailableException) StatusCode() int {
	return s.RespMetadata.StatusCode
}

// RequestID returns the service's response RequestID for request.
func (s *SerialConsoleSessionUnavailableException) RequestID() string {
	return s.RespMetadata.RequestID
}

// The service encountered an error. Follow the instructions in the error message
// and try again.
type ServiceException struct {
	_            struct{}                  `type:"structure"`
	RespMetadata protocol.ResponseMetadata `json:"-" xml:"-"`

	Message_ *string `locationName:"Message" type:"string"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s ServiceException) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s ServiceException) GoString() string {
	return s.String()
}

func newErrorServiceException(v protocol.ResponseMetadata) error {
	return &ServiceException{
		RespMetadata: v,
	}
}

// Code returns the exception type name.
func (s *ServiceException) Code() string {
	return "ServiceException"
}

// Message returns the exception's message.
func (s *ServiceException) Message() string {
	if s.Message_ != nil {
		return *s.Message_
	}
	return ""
}

// OrigErr always returns nil, satisfies awserr.Error interface.
func (s *ServiceException) OrigErr() error {
	return nil
}

func (s *ServiceException) Error() string {
	return fmt.Sprintf("%s: %s", s.Code(), s.Message())
}

// Status code returns the HTTP status code for the request's response error.
func (s *ServiceException) StatusCode() int {
	return s.RespMetadata.StatusCode
}

// RequestID returns the service's response RequestID for request.
func (s *ServiceException) RequestID() string {
	return s.RespMetadata.RequestID
}

// The requests were made too frequently and have been throttled. Wait a while
// and try again. To increase the limit on your request frequency, contact AWS
// Support.
type ThrottlingException struct {
	_            struct{}                  `type:"structure"`
	RespMetadata protocol.ResponseMetadata `json:"-" xml:"-"`

	Message_ *string `locationName:"Message" type:"string"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s ThrottlingException) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s ThrottlingException) GoString() string {
	return s.String()
}

func newErrorThrottlingException(v protocol.ResponseMetadata) error {
	return &ThrottlingException{
		RespMetadata: v,
	}
}

// Code returns the exception type name.
func (s *ThrottlingException) Code() string {
	return "ThrottlingException"
}

// Message returns the exception's message.
func (s *ThrottlingException) Message() string {
	if s.Message_ != nil {
		return *s.Message_
	}
	return ""
}

// OrigErr always returns nil, satisfies awserr.Error interface.
func (s *ThrottlingException) OrigErr() error {
	return nil
}

func (s *ThrottlingException) Error() string {
	return fmt.Sprintf("%s: %s", s.Code(), s.Message())
}

// Status code returns the HTTP status code for the request's response error.
func (s *ThrottlingException) StatusCode() int {
	return s.RespMetadata.StatusCode
}

// RequestID returns the service's response RequestID for request.
func (s *ThrottlingException) RequestID() string {
	return s.RespMetadata.RequestID
}
