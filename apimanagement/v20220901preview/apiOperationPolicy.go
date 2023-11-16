// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220901preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Policy Contract details.
type ApiOperationPolicy struct {
	pulumi.CustomResourceState

	// Format of the policyContent.
	Format pulumi.StringPtrOutput `pulumi:"format"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// Contents of the Policy as defined by the format.
	Value pulumi.StringOutput `pulumi:"value"`
}

// NewApiOperationPolicy registers a new resource with the given unique name, arguments, and options.
func NewApiOperationPolicy(ctx *pulumi.Context,
	name string, args *ApiOperationPolicyArgs, opts ...pulumi.ResourceOption) (*ApiOperationPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiId == nil {
		return nil, errors.New("invalid value for required argument 'ApiId'")
	}
	if args.OperationId == nil {
		return nil, errors.New("invalid value for required argument 'OperationId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	if args.Value == nil {
		return nil, errors.New("invalid value for required argument 'Value'")
	}
	if args.Format == nil {
		args.Format = pulumi.StringPtr("xml")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:apimanagement:ApiOperationPolicy"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20170301:ApiOperationPolicy"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180101:ApiOperationPolicy"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180601preview:ApiOperationPolicy"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20190101:ApiOperationPolicy"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201:ApiOperationPolicy"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201preview:ApiOperationPolicy"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20200601preview:ApiOperationPolicy"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20201201:ApiOperationPolicy"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210101preview:ApiOperationPolicy"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210401preview:ApiOperationPolicy"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210801:ApiOperationPolicy"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20211201preview:ApiOperationPolicy"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220401preview:ApiOperationPolicy"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220801:ApiOperationPolicy"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20230301preview:ApiOperationPolicy"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ApiOperationPolicy
	err := ctx.RegisterResource("azure-native:apimanagement/v20220901preview:ApiOperationPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApiOperationPolicy gets an existing ApiOperationPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApiOperationPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApiOperationPolicyState, opts ...pulumi.ResourceOption) (*ApiOperationPolicy, error) {
	var resource ApiOperationPolicy
	err := ctx.ReadResource("azure-native:apimanagement/v20220901preview:ApiOperationPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApiOperationPolicy resources.
type apiOperationPolicyState struct {
}

type ApiOperationPolicyState struct {
}

func (ApiOperationPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*apiOperationPolicyState)(nil)).Elem()
}

type apiOperationPolicyArgs struct {
	// API revision identifier. Must be unique in the current API Management service instance. Non-current revision has ;rev=n as a suffix where n is the revision number.
	ApiId string `pulumi:"apiId"`
	// Format of the policyContent.
	Format *string `pulumi:"format"`
	// Operation identifier within an API. Must be unique in the current API Management service instance.
	OperationId string `pulumi:"operationId"`
	// The identifier of the Policy.
	PolicyId *string `pulumi:"policyId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
	// Contents of the Policy as defined by the format.
	Value string `pulumi:"value"`
}

// The set of arguments for constructing a ApiOperationPolicy resource.
type ApiOperationPolicyArgs struct {
	// API revision identifier. Must be unique in the current API Management service instance. Non-current revision has ;rev=n as a suffix where n is the revision number.
	ApiId pulumi.StringInput
	// Format of the policyContent.
	Format pulumi.StringPtrInput
	// Operation identifier within an API. Must be unique in the current API Management service instance.
	OperationId pulumi.StringInput
	// The identifier of the Policy.
	PolicyId pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the API Management service.
	ServiceName pulumi.StringInput
	// Contents of the Policy as defined by the format.
	Value pulumi.StringInput
}

func (ApiOperationPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apiOperationPolicyArgs)(nil)).Elem()
}

type ApiOperationPolicyInput interface {
	pulumi.Input

	ToApiOperationPolicyOutput() ApiOperationPolicyOutput
	ToApiOperationPolicyOutputWithContext(ctx context.Context) ApiOperationPolicyOutput
}

func (*ApiOperationPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiOperationPolicy)(nil)).Elem()
}

func (i *ApiOperationPolicy) ToApiOperationPolicyOutput() ApiOperationPolicyOutput {
	return i.ToApiOperationPolicyOutputWithContext(context.Background())
}

func (i *ApiOperationPolicy) ToApiOperationPolicyOutputWithContext(ctx context.Context) ApiOperationPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiOperationPolicyOutput)
}

type ApiOperationPolicyOutput struct{ *pulumi.OutputState }

func (ApiOperationPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiOperationPolicy)(nil)).Elem()
}

func (o ApiOperationPolicyOutput) ToApiOperationPolicyOutput() ApiOperationPolicyOutput {
	return o
}

func (o ApiOperationPolicyOutput) ToApiOperationPolicyOutputWithContext(ctx context.Context) ApiOperationPolicyOutput {
	return o
}

// Format of the policyContent.
func (o ApiOperationPolicyOutput) Format() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiOperationPolicy) pulumi.StringPtrOutput { return v.Format }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o ApiOperationPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiOperationPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ApiOperationPolicyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiOperationPolicy) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Contents of the Policy as defined by the format.
func (o ApiOperationPolicyOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiOperationPolicy) pulumi.StringOutput { return v.Value }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ApiOperationPolicyOutput{})
}
