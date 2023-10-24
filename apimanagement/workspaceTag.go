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

// Tag Contract details.
// Azure REST API version: 2022-09-01-preview.
//
// Other available API versions: 2023-03-01-preview.
type WorkspaceTag struct {
	pulumi.CustomResourceState

	// Tag name.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewWorkspaceTag registers a new resource with the given unique name, arguments, and options.
func NewWorkspaceTag(ctx *pulumi.Context,
	name string, args *WorkspaceTagArgs, opts ...pulumi.ResourceOption) (*WorkspaceTag, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	if args.WorkspaceId == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceId'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:apimanagement/v20220901preview:WorkspaceTag"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20230301preview:WorkspaceTag"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource WorkspaceTag
	err := ctx.RegisterResource("azure-native:apimanagement:WorkspaceTag", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkspaceTag gets an existing WorkspaceTag resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkspaceTag(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkspaceTagState, opts ...pulumi.ResourceOption) (*WorkspaceTag, error) {
	var resource WorkspaceTag
	err := ctx.ReadResource("azure-native:apimanagement:WorkspaceTag", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WorkspaceTag resources.
type workspaceTagState struct {
}

type WorkspaceTagState struct {
}

func (WorkspaceTagState) ElementType() reflect.Type {
	return reflect.TypeOf((*workspaceTagState)(nil)).Elem()
}

type workspaceTagArgs struct {
	// Tag name.
	DisplayName string `pulumi:"displayName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
	// Tag identifier. Must be unique in the current API Management service instance.
	TagId *string `pulumi:"tagId"`
	// Workspace identifier. Must be unique in the current API Management service instance.
	WorkspaceId string `pulumi:"workspaceId"`
}

// The set of arguments for constructing a WorkspaceTag resource.
type WorkspaceTagArgs struct {
	// Tag name.
	DisplayName pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the API Management service.
	ServiceName pulumi.StringInput
	// Tag identifier. Must be unique in the current API Management service instance.
	TagId pulumi.StringPtrInput
	// Workspace identifier. Must be unique in the current API Management service instance.
	WorkspaceId pulumi.StringInput
}

func (WorkspaceTagArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workspaceTagArgs)(nil)).Elem()
}

type WorkspaceTagInput interface {
	pulumi.Input

	ToWorkspaceTagOutput() WorkspaceTagOutput
	ToWorkspaceTagOutputWithContext(ctx context.Context) WorkspaceTagOutput
}

func (*WorkspaceTag) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceTag)(nil)).Elem()
}

func (i *WorkspaceTag) ToWorkspaceTagOutput() WorkspaceTagOutput {
	return i.ToWorkspaceTagOutputWithContext(context.Background())
}

func (i *WorkspaceTag) ToWorkspaceTagOutputWithContext(ctx context.Context) WorkspaceTagOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceTagOutput)
}

func (i *WorkspaceTag) ToOutput(ctx context.Context) pulumix.Output[*WorkspaceTag] {
	return pulumix.Output[*WorkspaceTag]{
		OutputState: i.ToWorkspaceTagOutputWithContext(ctx).OutputState,
	}
}

type WorkspaceTagOutput struct{ *pulumi.OutputState }

func (WorkspaceTagOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceTag)(nil)).Elem()
}

func (o WorkspaceTagOutput) ToWorkspaceTagOutput() WorkspaceTagOutput {
	return o
}

func (o WorkspaceTagOutput) ToWorkspaceTagOutputWithContext(ctx context.Context) WorkspaceTagOutput {
	return o
}

func (o WorkspaceTagOutput) ToOutput(ctx context.Context) pulumix.Output[*WorkspaceTag] {
	return pulumix.Output[*WorkspaceTag]{
		OutputState: o.OutputState,
	}
}

// Tag name.
func (o WorkspaceTagOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkspaceTag) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// The name of the resource
func (o WorkspaceTagOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkspaceTag) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o WorkspaceTagOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkspaceTag) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WorkspaceTagOutput{})
}
