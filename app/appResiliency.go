// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package app

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configuration to setup App Resiliency
// Azure REST API version: 2023-08-01-preview.
type AppResiliency struct {
	pulumi.CustomResourceState

	// Policy that defines circuit breaker conditions
	CircuitBreakerPolicy CircuitBreakerPolicyResponsePtrOutput `pulumi:"circuitBreakerPolicy"`
	// Defines parameters for http connection pooling
	HttpConnectionPool HttpConnectionPoolResponsePtrOutput `pulumi:"httpConnectionPool"`
	// Policy that defines http request retry conditions
	HttpRetryPolicy HttpRetryPolicyResponsePtrOutput `pulumi:"httpRetryPolicy"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Defines parameters for tcp connection pooling
	TcpConnectionPool TcpConnectionPoolResponsePtrOutput `pulumi:"tcpConnectionPool"`
	// Policy that defines tcp request retry conditions
	TcpRetryPolicy TcpRetryPolicyResponsePtrOutput `pulumi:"tcpRetryPolicy"`
	// Policy to set request timeouts
	TimeoutPolicy TimeoutPolicyResponsePtrOutput `pulumi:"timeoutPolicy"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewAppResiliency registers a new resource with the given unique name, arguments, and options.
func NewAppResiliency(ctx *pulumi.Context,
	name string, args *AppResiliencyArgs, opts ...pulumi.ResourceOption) (*AppResiliency, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AppName == nil {
		return nil, errors.New("invalid value for required argument 'AppName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:app/v20230801preview:AppResiliency"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource AppResiliency
	err := ctx.RegisterResource("azure-native:app:AppResiliency", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAppResiliency gets an existing AppResiliency resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAppResiliency(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AppResiliencyState, opts ...pulumi.ResourceOption) (*AppResiliency, error) {
	var resource AppResiliency
	err := ctx.ReadResource("azure-native:app:AppResiliency", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AppResiliency resources.
type appResiliencyState struct {
}

type AppResiliencyState struct {
}

func (AppResiliencyState) ElementType() reflect.Type {
	return reflect.TypeOf((*appResiliencyState)(nil)).Elem()
}

type appResiliencyArgs struct {
	// Name of the Container App.
	AppName string `pulumi:"appName"`
	// Policy that defines circuit breaker conditions
	CircuitBreakerPolicy *CircuitBreakerPolicy `pulumi:"circuitBreakerPolicy"`
	// Defines parameters for http connection pooling
	HttpConnectionPool *HttpConnectionPool `pulumi:"httpConnectionPool"`
	// Policy that defines http request retry conditions
	HttpRetryPolicy *HttpRetryPolicy `pulumi:"httpRetryPolicy"`
	// Name of the resiliency policy.
	Name *string `pulumi:"name"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Defines parameters for tcp connection pooling
	TcpConnectionPool *TcpConnectionPool `pulumi:"tcpConnectionPool"`
	// Policy that defines tcp request retry conditions
	TcpRetryPolicy *TcpRetryPolicy `pulumi:"tcpRetryPolicy"`
	// Policy to set request timeouts
	TimeoutPolicy *TimeoutPolicy `pulumi:"timeoutPolicy"`
}

// The set of arguments for constructing a AppResiliency resource.
type AppResiliencyArgs struct {
	// Name of the Container App.
	AppName pulumi.StringInput
	// Policy that defines circuit breaker conditions
	CircuitBreakerPolicy CircuitBreakerPolicyPtrInput
	// Defines parameters for http connection pooling
	HttpConnectionPool HttpConnectionPoolPtrInput
	// Policy that defines http request retry conditions
	HttpRetryPolicy HttpRetryPolicyPtrInput
	// Name of the resiliency policy.
	Name pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Defines parameters for tcp connection pooling
	TcpConnectionPool TcpConnectionPoolPtrInput
	// Policy that defines tcp request retry conditions
	TcpRetryPolicy TcpRetryPolicyPtrInput
	// Policy to set request timeouts
	TimeoutPolicy TimeoutPolicyPtrInput
}

func (AppResiliencyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*appResiliencyArgs)(nil)).Elem()
}

type AppResiliencyInput interface {
	pulumi.Input

	ToAppResiliencyOutput() AppResiliencyOutput
	ToAppResiliencyOutputWithContext(ctx context.Context) AppResiliencyOutput
}

func (*AppResiliency) ElementType() reflect.Type {
	return reflect.TypeOf((**AppResiliency)(nil)).Elem()
}

func (i *AppResiliency) ToAppResiliencyOutput() AppResiliencyOutput {
	return i.ToAppResiliencyOutputWithContext(context.Background())
}

func (i *AppResiliency) ToAppResiliencyOutputWithContext(ctx context.Context) AppResiliencyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppResiliencyOutput)
}

type AppResiliencyOutput struct{ *pulumi.OutputState }

func (AppResiliencyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppResiliency)(nil)).Elem()
}

func (o AppResiliencyOutput) ToAppResiliencyOutput() AppResiliencyOutput {
	return o
}

func (o AppResiliencyOutput) ToAppResiliencyOutputWithContext(ctx context.Context) AppResiliencyOutput {
	return o
}

// Policy that defines circuit breaker conditions
func (o AppResiliencyOutput) CircuitBreakerPolicy() CircuitBreakerPolicyResponsePtrOutput {
	return o.ApplyT(func(v *AppResiliency) CircuitBreakerPolicyResponsePtrOutput { return v.CircuitBreakerPolicy }).(CircuitBreakerPolicyResponsePtrOutput)
}

// Defines parameters for http connection pooling
func (o AppResiliencyOutput) HttpConnectionPool() HttpConnectionPoolResponsePtrOutput {
	return o.ApplyT(func(v *AppResiliency) HttpConnectionPoolResponsePtrOutput { return v.HttpConnectionPool }).(HttpConnectionPoolResponsePtrOutput)
}

// Policy that defines http request retry conditions
func (o AppResiliencyOutput) HttpRetryPolicy() HttpRetryPolicyResponsePtrOutput {
	return o.ApplyT(func(v *AppResiliency) HttpRetryPolicyResponsePtrOutput { return v.HttpRetryPolicy }).(HttpRetryPolicyResponsePtrOutput)
}

// The name of the resource
func (o AppResiliencyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AppResiliency) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o AppResiliencyOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *AppResiliency) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Defines parameters for tcp connection pooling
func (o AppResiliencyOutput) TcpConnectionPool() TcpConnectionPoolResponsePtrOutput {
	return o.ApplyT(func(v *AppResiliency) TcpConnectionPoolResponsePtrOutput { return v.TcpConnectionPool }).(TcpConnectionPoolResponsePtrOutput)
}

// Policy that defines tcp request retry conditions
func (o AppResiliencyOutput) TcpRetryPolicy() TcpRetryPolicyResponsePtrOutput {
	return o.ApplyT(func(v *AppResiliency) TcpRetryPolicyResponsePtrOutput { return v.TcpRetryPolicy }).(TcpRetryPolicyResponsePtrOutput)
}

// Policy to set request timeouts
func (o AppResiliencyOutput) TimeoutPolicy() TimeoutPolicyResponsePtrOutput {
	return o.ApplyT(func(v *AppResiliency) TimeoutPolicyResponsePtrOutput { return v.TimeoutPolicy }).(TimeoutPolicyResponsePtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o AppResiliencyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AppResiliency) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AppResiliencyOutput{})
}
