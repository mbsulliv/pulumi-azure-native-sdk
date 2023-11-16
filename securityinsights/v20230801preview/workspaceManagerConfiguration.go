// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230801preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The workspace manager configuration
type WorkspaceManagerConfiguration struct {
	pulumi.CustomResourceState

	// Resource Etag.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The current mode of the workspace manager configuration
	Mode pulumi.StringOutput `pulumi:"mode"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewWorkspaceManagerConfiguration registers a new resource with the given unique name, arguments, and options.
func NewWorkspaceManagerConfiguration(ctx *pulumi.Context,
	name string, args *WorkspaceManagerConfigurationArgs, opts ...pulumi.ResourceOption) (*WorkspaceManagerConfiguration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Mode == nil {
		return nil, errors.New("invalid value for required argument 'Mode'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights:WorkspaceManagerConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230401preview:WorkspaceManagerConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230501preview:WorkspaceManagerConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230601preview:WorkspaceManagerConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230701preview:WorkspaceManagerConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230901preview:WorkspaceManagerConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20231001preview:WorkspaceManagerConfiguration"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource WorkspaceManagerConfiguration
	err := ctx.RegisterResource("azure-native:securityinsights/v20230801preview:WorkspaceManagerConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkspaceManagerConfiguration gets an existing WorkspaceManagerConfiguration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkspaceManagerConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkspaceManagerConfigurationState, opts ...pulumi.ResourceOption) (*WorkspaceManagerConfiguration, error) {
	var resource WorkspaceManagerConfiguration
	err := ctx.ReadResource("azure-native:securityinsights/v20230801preview:WorkspaceManagerConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WorkspaceManagerConfiguration resources.
type workspaceManagerConfigurationState struct {
}

type WorkspaceManagerConfigurationState struct {
}

func (WorkspaceManagerConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*workspaceManagerConfigurationState)(nil)).Elem()
}

type workspaceManagerConfigurationArgs struct {
	// The current mode of the workspace manager configuration
	Mode string `pulumi:"mode"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the workspace manager configuration
	WorkspaceManagerConfigurationName *string `pulumi:"workspaceManagerConfigurationName"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a WorkspaceManagerConfiguration resource.
type WorkspaceManagerConfigurationArgs struct {
	// The current mode of the workspace manager configuration
	Mode pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the workspace manager configuration
	WorkspaceManagerConfigurationName pulumi.StringPtrInput
	// The name of the workspace.
	WorkspaceName pulumi.StringInput
}

func (WorkspaceManagerConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workspaceManagerConfigurationArgs)(nil)).Elem()
}

type WorkspaceManagerConfigurationInput interface {
	pulumi.Input

	ToWorkspaceManagerConfigurationOutput() WorkspaceManagerConfigurationOutput
	ToWorkspaceManagerConfigurationOutputWithContext(ctx context.Context) WorkspaceManagerConfigurationOutput
}

func (*WorkspaceManagerConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceManagerConfiguration)(nil)).Elem()
}

func (i *WorkspaceManagerConfiguration) ToWorkspaceManagerConfigurationOutput() WorkspaceManagerConfigurationOutput {
	return i.ToWorkspaceManagerConfigurationOutputWithContext(context.Background())
}

func (i *WorkspaceManagerConfiguration) ToWorkspaceManagerConfigurationOutputWithContext(ctx context.Context) WorkspaceManagerConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceManagerConfigurationOutput)
}

type WorkspaceManagerConfigurationOutput struct{ *pulumi.OutputState }

func (WorkspaceManagerConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceManagerConfiguration)(nil)).Elem()
}

func (o WorkspaceManagerConfigurationOutput) ToWorkspaceManagerConfigurationOutput() WorkspaceManagerConfigurationOutput {
	return o
}

func (o WorkspaceManagerConfigurationOutput) ToWorkspaceManagerConfigurationOutputWithContext(ctx context.Context) WorkspaceManagerConfigurationOutput {
	return o
}

// Resource Etag.
func (o WorkspaceManagerConfigurationOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkspaceManagerConfiguration) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The current mode of the workspace manager configuration
func (o WorkspaceManagerConfigurationOutput) Mode() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkspaceManagerConfiguration) pulumi.StringOutput { return v.Mode }).(pulumi.StringOutput)
}

// The name of the resource
func (o WorkspaceManagerConfigurationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkspaceManagerConfiguration) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o WorkspaceManagerConfigurationOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *WorkspaceManagerConfiguration) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o WorkspaceManagerConfigurationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkspaceManagerConfiguration) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WorkspaceManagerConfigurationOutput{})
}
