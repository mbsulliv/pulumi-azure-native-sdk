// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apimanagement

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Policy Contract details.
// Azure REST API version: 2022-09-01-preview.
//
// Other available API versions: 2023-03-01-preview.
type WorkspacePolicy struct {
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

// NewWorkspacePolicy registers a new resource with the given unique name, arguments, and options.
func NewWorkspacePolicy(ctx *pulumi.Context,
	name string, args *WorkspacePolicyArgs, opts ...pulumi.ResourceOption) (*WorkspacePolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
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
			Type: pulumi.String("azure-native:apimanagement/v20220901preview:WorkspacePolicy"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20230301preview:WorkspacePolicy"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource WorkspacePolicy
	err := ctx.RegisterResource("azure-native:apimanagement:WorkspacePolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkspacePolicy gets an existing WorkspacePolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkspacePolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkspacePolicyState, opts ...pulumi.ResourceOption) (*WorkspacePolicy, error) {
	var resource WorkspacePolicy
	err := ctx.ReadResource("azure-native:apimanagement:WorkspacePolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WorkspacePolicy resources.
type workspacePolicyState struct {
}

type WorkspacePolicyState struct {
}

func (WorkspacePolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*workspacePolicyState)(nil)).Elem()
}

type workspacePolicyArgs struct {
	// Format of the policyContent.
	Format *string `pulumi:"format"`
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

// The set of arguments for constructing a WorkspacePolicy resource.
type WorkspacePolicyArgs struct {
	// Format of the policyContent.
	Format pulumi.StringPtrInput
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

func (WorkspacePolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workspacePolicyArgs)(nil)).Elem()
}

type WorkspacePolicyInput interface {
	pulumi.Input

	ToWorkspacePolicyOutput() WorkspacePolicyOutput
	ToWorkspacePolicyOutputWithContext(ctx context.Context) WorkspacePolicyOutput
}

func (*WorkspacePolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspacePolicy)(nil)).Elem()
}

func (i *WorkspacePolicy) ToWorkspacePolicyOutput() WorkspacePolicyOutput {
	return i.ToWorkspacePolicyOutputWithContext(context.Background())
}

func (i *WorkspacePolicy) ToWorkspacePolicyOutputWithContext(ctx context.Context) WorkspacePolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspacePolicyOutput)
}

type WorkspacePolicyOutput struct{ *pulumi.OutputState }

func (WorkspacePolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspacePolicy)(nil)).Elem()
}

func (o WorkspacePolicyOutput) ToWorkspacePolicyOutput() WorkspacePolicyOutput {
	return o
}

func (o WorkspacePolicyOutput) ToWorkspacePolicyOutputWithContext(ctx context.Context) WorkspacePolicyOutput {
	return o
}

// Format of the policyContent.
func (o WorkspacePolicyOutput) Format() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkspacePolicy) pulumi.StringPtrOutput { return v.Format }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o WorkspacePolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkspacePolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o WorkspacePolicyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkspacePolicy) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Contents of the Policy as defined by the format.
func (o WorkspacePolicyOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkspacePolicy) pulumi.StringOutput { return v.Value }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WorkspacePolicyOutput{})
}
