// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230801preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The workspace manager group
type WorkspaceManagerGroup struct {
	pulumi.CustomResourceState

	// The description of the workspace manager group
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The display name of the workspace manager group
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Resource Etag.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The names of the workspace manager members participating in this group.
	MemberResourceNames pulumi.StringArrayOutput `pulumi:"memberResourceNames"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewWorkspaceManagerGroup registers a new resource with the given unique name, arguments, and options.
func NewWorkspaceManagerGroup(ctx *pulumi.Context,
	name string, args *WorkspaceManagerGroupArgs, opts ...pulumi.ResourceOption) (*WorkspaceManagerGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.MemberResourceNames == nil {
		return nil, errors.New("invalid value for required argument 'MemberResourceNames'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights:WorkspaceManagerGroup"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230401preview:WorkspaceManagerGroup"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230501preview:WorkspaceManagerGroup"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230601preview:WorkspaceManagerGroup"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230701preview:WorkspaceManagerGroup"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230901preview:WorkspaceManagerGroup"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource WorkspaceManagerGroup
	err := ctx.RegisterResource("azure-native:securityinsights/v20230801preview:WorkspaceManagerGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkspaceManagerGroup gets an existing WorkspaceManagerGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkspaceManagerGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkspaceManagerGroupState, opts ...pulumi.ResourceOption) (*WorkspaceManagerGroup, error) {
	var resource WorkspaceManagerGroup
	err := ctx.ReadResource("azure-native:securityinsights/v20230801preview:WorkspaceManagerGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WorkspaceManagerGroup resources.
type workspaceManagerGroupState struct {
}

type WorkspaceManagerGroupState struct {
}

func (WorkspaceManagerGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*workspaceManagerGroupState)(nil)).Elem()
}

type workspaceManagerGroupArgs struct {
	// The description of the workspace manager group
	Description *string `pulumi:"description"`
	// The display name of the workspace manager group
	DisplayName string `pulumi:"displayName"`
	// The names of the workspace manager members participating in this group.
	MemberResourceNames []string `pulumi:"memberResourceNames"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the workspace manager group
	WorkspaceManagerGroupName *string `pulumi:"workspaceManagerGroupName"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a WorkspaceManagerGroup resource.
type WorkspaceManagerGroupArgs struct {
	// The description of the workspace manager group
	Description pulumi.StringPtrInput
	// The display name of the workspace manager group
	DisplayName pulumi.StringInput
	// The names of the workspace manager members participating in this group.
	MemberResourceNames pulumi.StringArrayInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the workspace manager group
	WorkspaceManagerGroupName pulumi.StringPtrInput
	// The name of the workspace.
	WorkspaceName pulumi.StringInput
}

func (WorkspaceManagerGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workspaceManagerGroupArgs)(nil)).Elem()
}

type WorkspaceManagerGroupInput interface {
	pulumi.Input

	ToWorkspaceManagerGroupOutput() WorkspaceManagerGroupOutput
	ToWorkspaceManagerGroupOutputWithContext(ctx context.Context) WorkspaceManagerGroupOutput
}

func (*WorkspaceManagerGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceManagerGroup)(nil)).Elem()
}

func (i *WorkspaceManagerGroup) ToWorkspaceManagerGroupOutput() WorkspaceManagerGroupOutput {
	return i.ToWorkspaceManagerGroupOutputWithContext(context.Background())
}

func (i *WorkspaceManagerGroup) ToWorkspaceManagerGroupOutputWithContext(ctx context.Context) WorkspaceManagerGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceManagerGroupOutput)
}

func (i *WorkspaceManagerGroup) ToOutput(ctx context.Context) pulumix.Output[*WorkspaceManagerGroup] {
	return pulumix.Output[*WorkspaceManagerGroup]{
		OutputState: i.ToWorkspaceManagerGroupOutputWithContext(ctx).OutputState,
	}
}

type WorkspaceManagerGroupOutput struct{ *pulumi.OutputState }

func (WorkspaceManagerGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceManagerGroup)(nil)).Elem()
}

func (o WorkspaceManagerGroupOutput) ToWorkspaceManagerGroupOutput() WorkspaceManagerGroupOutput {
	return o
}

func (o WorkspaceManagerGroupOutput) ToWorkspaceManagerGroupOutputWithContext(ctx context.Context) WorkspaceManagerGroupOutput {
	return o
}

func (o WorkspaceManagerGroupOutput) ToOutput(ctx context.Context) pulumix.Output[*WorkspaceManagerGroup] {
	return pulumix.Output[*WorkspaceManagerGroup]{
		OutputState: o.OutputState,
	}
}

// The description of the workspace manager group
func (o WorkspaceManagerGroupOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkspaceManagerGroup) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The display name of the workspace manager group
func (o WorkspaceManagerGroupOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkspaceManagerGroup) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// Resource Etag.
func (o WorkspaceManagerGroupOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkspaceManagerGroup) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The names of the workspace manager members participating in this group.
func (o WorkspaceManagerGroupOutput) MemberResourceNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *WorkspaceManagerGroup) pulumi.StringArrayOutput { return v.MemberResourceNames }).(pulumi.StringArrayOutput)
}

// The name of the resource
func (o WorkspaceManagerGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkspaceManagerGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o WorkspaceManagerGroupOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *WorkspaceManagerGroup) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o WorkspaceManagerGroupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkspaceManagerGroup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WorkspaceManagerGroupOutput{})
}
