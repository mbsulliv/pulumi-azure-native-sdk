// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20170501preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// An object that represents a machine learning team account workspace.
type Workspace struct {
	pulumi.CustomResourceState

	// The immutable id of the team account which contains this workspace.
	AccountId pulumi.StringOutput `pulumi:"accountId"`
	// The creation date of the machine learning workspace in ISO8601 format.
	CreationDate pulumi.StringOutput `pulumi:"creationDate"`
	// The description of this workspace.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The friendly name for this workspace. This will be the workspace name in the arm id when the workspace object gets created
	FriendlyName pulumi.StringOutput `pulumi:"friendlyName"`
	// The location of the resource. This cannot be changed after the resource is created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The current deployment state of team account workspace resource. The provisioningState is to indicate states for resource provisioning.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The tags of the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
	// The immutable id of this workspace.
	WorkspaceId pulumi.StringOutput `pulumi:"workspaceId"`
}

// NewWorkspace registers a new resource with the given unique name, arguments, and options.
func NewWorkspace(ctx *pulumi.Context,
	name string, args *WorkspaceArgs, opts ...pulumi.ResourceOption) (*Workspace, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.FriendlyName == nil {
		return nil, errors.New("invalid value for required argument 'FriendlyName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:machinelearningexperimentation:Workspace"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Workspace
	err := ctx.RegisterResource("azure-native:machinelearningexperimentation/v20170501preview:Workspace", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkspace gets an existing Workspace resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkspace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkspaceState, opts ...pulumi.ResourceOption) (*Workspace, error) {
	var resource Workspace
	err := ctx.ReadResource("azure-native:machinelearningexperimentation/v20170501preview:Workspace", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Workspace resources.
type workspaceState struct {
}

type WorkspaceState struct {
}

func (WorkspaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*workspaceState)(nil)).Elem()
}

type workspaceArgs struct {
	// The name of the machine learning team account.
	AccountName string `pulumi:"accountName"`
	// The description of this workspace.
	Description *string `pulumi:"description"`
	// The friendly name for this workspace. This will be the workspace name in the arm id when the workspace object gets created
	FriendlyName string `pulumi:"friendlyName"`
	// The location of the resource. This cannot be changed after the resource is created.
	Location *string `pulumi:"location"`
	// The name of the resource group to which the machine learning team account belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The tags of the resource.
	Tags map[string]string `pulumi:"tags"`
	// The name of the machine learning team account workspace.
	WorkspaceName *string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a Workspace resource.
type WorkspaceArgs struct {
	// The name of the machine learning team account.
	AccountName pulumi.StringInput
	// The description of this workspace.
	Description pulumi.StringPtrInput
	// The friendly name for this workspace. This will be the workspace name in the arm id when the workspace object gets created
	FriendlyName pulumi.StringInput
	// The location of the resource. This cannot be changed after the resource is created.
	Location pulumi.StringPtrInput
	// The name of the resource group to which the machine learning team account belongs.
	ResourceGroupName pulumi.StringInput
	// The tags of the resource.
	Tags pulumi.StringMapInput
	// The name of the machine learning team account workspace.
	WorkspaceName pulumi.StringPtrInput
}

func (WorkspaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workspaceArgs)(nil)).Elem()
}

type WorkspaceInput interface {
	pulumi.Input

	ToWorkspaceOutput() WorkspaceOutput
	ToWorkspaceOutputWithContext(ctx context.Context) WorkspaceOutput
}

func (*Workspace) ElementType() reflect.Type {
	return reflect.TypeOf((**Workspace)(nil)).Elem()
}

func (i *Workspace) ToWorkspaceOutput() WorkspaceOutput {
	return i.ToWorkspaceOutputWithContext(context.Background())
}

func (i *Workspace) ToWorkspaceOutputWithContext(ctx context.Context) WorkspaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceOutput)
}

func (i *Workspace) ToOutput(ctx context.Context) pulumix.Output[*Workspace] {
	return pulumix.Output[*Workspace]{
		OutputState: i.ToWorkspaceOutputWithContext(ctx).OutputState,
	}
}

type WorkspaceOutput struct{ *pulumi.OutputState }

func (WorkspaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Workspace)(nil)).Elem()
}

func (o WorkspaceOutput) ToWorkspaceOutput() WorkspaceOutput {
	return o
}

func (o WorkspaceOutput) ToWorkspaceOutputWithContext(ctx context.Context) WorkspaceOutput {
	return o
}

func (o WorkspaceOutput) ToOutput(ctx context.Context) pulumix.Output[*Workspace] {
	return pulumix.Output[*Workspace]{
		OutputState: o.OutputState,
	}
}

// The immutable id of the team account which contains this workspace.
func (o WorkspaceOutput) AccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringOutput { return v.AccountId }).(pulumi.StringOutput)
}

// The creation date of the machine learning workspace in ISO8601 format.
func (o WorkspaceOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringOutput { return v.CreationDate }).(pulumi.StringOutput)
}

// The description of this workspace.
func (o WorkspaceOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The friendly name for this workspace. This will be the workspace name in the arm id when the workspace object gets created
func (o WorkspaceOutput) FriendlyName() pulumi.StringOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringOutput { return v.FriendlyName }).(pulumi.StringOutput)
}

// The location of the resource. This cannot be changed after the resource is created.
func (o WorkspaceOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource.
func (o WorkspaceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The current deployment state of team account workspace resource. The provisioningState is to indicate states for resource provisioning.
func (o WorkspaceOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The tags of the resource.
func (o WorkspaceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource.
func (o WorkspaceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The immutable id of this workspace.
func (o WorkspaceOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WorkspaceOutput{})
}
