// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apimanagement

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Policy Contract details.
// Azure REST API version: 2022-09-01-preview.
type WorkspaceApiOperationPolicy struct {
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

// NewWorkspaceApiOperationPolicy registers a new resource with the given unique name, arguments, and options.
func NewWorkspaceApiOperationPolicy(ctx *pulumi.Context,
	name string, args *WorkspaceApiOperationPolicyArgs, opts ...pulumi.ResourceOption) (*WorkspaceApiOperationPolicy, error) {
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
	if args.WorkspaceId == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceId'")
	}
	if args.Format == nil {
		args.Format = pulumi.StringPtr("xml")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:apimanagement/v20220901preview:WorkspaceApiOperationPolicy"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20230301preview:WorkspaceApiOperationPolicy"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource WorkspaceApiOperationPolicy
	err := ctx.RegisterResource("azure-native:apimanagement:WorkspaceApiOperationPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkspaceApiOperationPolicy gets an existing WorkspaceApiOperationPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkspaceApiOperationPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkspaceApiOperationPolicyState, opts ...pulumi.ResourceOption) (*WorkspaceApiOperationPolicy, error) {
	var resource WorkspaceApiOperationPolicy
	err := ctx.ReadResource("azure-native:apimanagement:WorkspaceApiOperationPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WorkspaceApiOperationPolicy resources.
type workspaceApiOperationPolicyState struct {
}

type WorkspaceApiOperationPolicyState struct {
}

func (WorkspaceApiOperationPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*workspaceApiOperationPolicyState)(nil)).Elem()
}

type workspaceApiOperationPolicyArgs struct {
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
	// Workspace identifier. Must be unique in the current API Management service instance.
	WorkspaceId string `pulumi:"workspaceId"`
}

// The set of arguments for constructing a WorkspaceApiOperationPolicy resource.
type WorkspaceApiOperationPolicyArgs struct {
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
	// Workspace identifier. Must be unique in the current API Management service instance.
	WorkspaceId pulumi.StringInput
}

func (WorkspaceApiOperationPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workspaceApiOperationPolicyArgs)(nil)).Elem()
}

type WorkspaceApiOperationPolicyInput interface {
	pulumi.Input

	ToWorkspaceApiOperationPolicyOutput() WorkspaceApiOperationPolicyOutput
	ToWorkspaceApiOperationPolicyOutputWithContext(ctx context.Context) WorkspaceApiOperationPolicyOutput
}

func (*WorkspaceApiOperationPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceApiOperationPolicy)(nil)).Elem()
}

func (i *WorkspaceApiOperationPolicy) ToWorkspaceApiOperationPolicyOutput() WorkspaceApiOperationPolicyOutput {
	return i.ToWorkspaceApiOperationPolicyOutputWithContext(context.Background())
}

func (i *WorkspaceApiOperationPolicy) ToWorkspaceApiOperationPolicyOutputWithContext(ctx context.Context) WorkspaceApiOperationPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceApiOperationPolicyOutput)
}

func (i *WorkspaceApiOperationPolicy) ToOutput(ctx context.Context) pulumix.Output[*WorkspaceApiOperationPolicy] {
	return pulumix.Output[*WorkspaceApiOperationPolicy]{
		OutputState: i.ToWorkspaceApiOperationPolicyOutputWithContext(ctx).OutputState,
	}
}

type WorkspaceApiOperationPolicyOutput struct{ *pulumi.OutputState }

func (WorkspaceApiOperationPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceApiOperationPolicy)(nil)).Elem()
}

func (o WorkspaceApiOperationPolicyOutput) ToWorkspaceApiOperationPolicyOutput() WorkspaceApiOperationPolicyOutput {
	return o
}

func (o WorkspaceApiOperationPolicyOutput) ToWorkspaceApiOperationPolicyOutputWithContext(ctx context.Context) WorkspaceApiOperationPolicyOutput {
	return o
}

func (o WorkspaceApiOperationPolicyOutput) ToOutput(ctx context.Context) pulumix.Output[*WorkspaceApiOperationPolicy] {
	return pulumix.Output[*WorkspaceApiOperationPolicy]{
		OutputState: o.OutputState,
	}
}

// Format of the policyContent.
func (o WorkspaceApiOperationPolicyOutput) Format() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkspaceApiOperationPolicy) pulumi.StringPtrOutput { return v.Format }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o WorkspaceApiOperationPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkspaceApiOperationPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o WorkspaceApiOperationPolicyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkspaceApiOperationPolicy) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Contents of the Policy as defined by the format.
func (o WorkspaceApiOperationPolicyOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkspaceApiOperationPolicy) pulumi.StringOutput { return v.Value }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WorkspaceApiOperationPolicyOutput{})
}
