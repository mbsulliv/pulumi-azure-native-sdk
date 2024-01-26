// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220110preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The resource proxy definition object for quantum workspace.
type Workspace struct {
	pulumi.CustomResourceState

	// The URI of the workspace endpoint.
	EndpointUri pulumi.StringOutput `pulumi:"endpointUri"`
	// Managed Identity information.
	Identity QuantumWorkspaceResponseIdentityPtrOutput `pulumi:"identity"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// List of Providers selected for this Workspace
	Providers ProviderResponseArrayOutput `pulumi:"providers"`
	// Provisioning status field
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// ARM Resource Id of the storage account associated with this workspace.
	StorageAccount pulumi.StringPtrOutput `pulumi:"storageAccount"`
	// System metadata
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// Whether the current workspace is ready to accept Jobs.
	Usable pulumi.StringOutput `pulumi:"usable"`
}

// NewWorkspace registers a new resource with the given unique name, arguments, and options.
func NewWorkspace(ctx *pulumi.Context,
	name string, args *WorkspaceArgs, opts ...pulumi.ResourceOption) (*Workspace, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:quantum:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:quantum/v20191104preview:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:quantum/v20231113preview:Workspace"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Workspace
	err := ctx.RegisterResource("azure-native:quantum/v20220110preview:Workspace", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure-native:quantum/v20220110preview:Workspace", name, id, state, &resource, opts...)
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
	// Managed Identity information.
	Identity *QuantumWorkspaceIdentity `pulumi:"identity"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// List of Providers selected for this Workspace
	Providers []Provider `pulumi:"providers"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// ARM Resource Id of the storage account associated with this workspace.
	StorageAccount *string `pulumi:"storageAccount"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The name of the quantum workspace resource.
	WorkspaceName *string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a Workspace resource.
type WorkspaceArgs struct {
	// Managed Identity information.
	Identity QuantumWorkspaceIdentityPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// List of Providers selected for this Workspace
	Providers ProviderArrayInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// ARM Resource Id of the storage account associated with this workspace.
	StorageAccount pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The name of the quantum workspace resource.
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

// The URI of the workspace endpoint.
func (o WorkspaceOutput) EndpointUri() pulumi.StringOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringOutput { return v.EndpointUri }).(pulumi.StringOutput)
}

// Managed Identity information.
func (o WorkspaceOutput) Identity() QuantumWorkspaceResponseIdentityPtrOutput {
	return o.ApplyT(func(v *Workspace) QuantumWorkspaceResponseIdentityPtrOutput { return v.Identity }).(QuantumWorkspaceResponseIdentityPtrOutput)
}

// The geo-location where the resource lives
func (o WorkspaceOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o WorkspaceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// List of Providers selected for this Workspace
func (o WorkspaceOutput) Providers() ProviderResponseArrayOutput {
	return o.ApplyT(func(v *Workspace) ProviderResponseArrayOutput { return v.Providers }).(ProviderResponseArrayOutput)
}

// Provisioning status field
func (o WorkspaceOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// ARM Resource Id of the storage account associated with this workspace.
func (o WorkspaceOutput) StorageAccount() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringPtrOutput { return v.StorageAccount }).(pulumi.StringPtrOutput)
}

// System metadata
func (o WorkspaceOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Workspace) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o WorkspaceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o WorkspaceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Whether the current workspace is ready to accept Jobs.
func (o WorkspaceOutput) Usable() pulumi.StringOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringOutput { return v.Usable }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WorkspaceOutput{})
}
