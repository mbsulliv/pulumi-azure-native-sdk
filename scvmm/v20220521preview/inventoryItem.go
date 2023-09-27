// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220521preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Defines the inventory item.
type InventoryItem struct {
	pulumi.CustomResourceState

	// Gets the Managed Object name in VMM for the inventory item.
	InventoryItemName pulumi.StringOutput `pulumi:"inventoryItemName"`
	// They inventory type.
	InventoryType pulumi.StringOutput `pulumi:"inventoryType"`
	// Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type; e.g. ApiApps are a kind of Microsoft.Web/sites type.  If supported, the resource provider must validate and persist this value.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Gets the tracked resource id corresponding to the inventory resource.
	ManagedResourceId pulumi.StringOutput `pulumi:"managedResourceId"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Gets the provisioning state.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// Gets the UUID (which is assigned by VMM) for the inventory item.
	Uuid pulumi.StringOutput `pulumi:"uuid"`
}

// NewInventoryItem registers a new resource with the given unique name, arguments, and options.
func NewInventoryItem(ctx *pulumi.Context,
	name string, args *InventoryItemArgs, opts ...pulumi.ResourceOption) (*InventoryItem, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InventoryType == nil {
		return nil, errors.New("invalid value for required argument 'InventoryType'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.VmmServerName == nil {
		return nil, errors.New("invalid value for required argument 'VmmServerName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:scvmm:InventoryItem"),
		},
		{
			Type: pulumi.String("azure-native:scvmm/v20200605preview:InventoryItem"),
		},
		{
			Type: pulumi.String("azure-native:scvmm/v20230401preview:InventoryItem"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource InventoryItem
	err := ctx.RegisterResource("azure-native:scvmm/v20220521preview:InventoryItem", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInventoryItem gets an existing InventoryItem resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInventoryItem(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InventoryItemState, opts ...pulumi.ResourceOption) (*InventoryItem, error) {
	var resource InventoryItem
	err := ctx.ReadResource("azure-native:scvmm/v20220521preview:InventoryItem", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InventoryItem resources.
type inventoryItemState struct {
}

type InventoryItemState struct {
}

func (InventoryItemState) ElementType() reflect.Type {
	return reflect.TypeOf((*inventoryItemState)(nil)).Elem()
}

type inventoryItemArgs struct {
	// Name of the inventoryItem.
	InventoryItemName *string `pulumi:"inventoryItemName"`
	// They inventory type.
	InventoryType string `pulumi:"inventoryType"`
	// Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type; e.g. ApiApps are a kind of Microsoft.Web/sites type.  If supported, the resource provider must validate and persist this value.
	Kind *string `pulumi:"kind"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the VMMServer.
	VmmServerName string `pulumi:"vmmServerName"`
}

// The set of arguments for constructing a InventoryItem resource.
type InventoryItemArgs struct {
	// Name of the inventoryItem.
	InventoryItemName pulumi.StringPtrInput
	// They inventory type.
	InventoryType pulumi.StringInput
	// Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type; e.g. ApiApps are a kind of Microsoft.Web/sites type.  If supported, the resource provider must validate and persist this value.
	Kind pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Name of the VMMServer.
	VmmServerName pulumi.StringInput
}

func (InventoryItemArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*inventoryItemArgs)(nil)).Elem()
}

type InventoryItemInput interface {
	pulumi.Input

	ToInventoryItemOutput() InventoryItemOutput
	ToInventoryItemOutputWithContext(ctx context.Context) InventoryItemOutput
}

func (*InventoryItem) ElementType() reflect.Type {
	return reflect.TypeOf((**InventoryItem)(nil)).Elem()
}

func (i *InventoryItem) ToInventoryItemOutput() InventoryItemOutput {
	return i.ToInventoryItemOutputWithContext(context.Background())
}

func (i *InventoryItem) ToInventoryItemOutputWithContext(ctx context.Context) InventoryItemOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InventoryItemOutput)
}

func (i *InventoryItem) ToOutput(ctx context.Context) pulumix.Output[*InventoryItem] {
	return pulumix.Output[*InventoryItem]{
		OutputState: i.ToInventoryItemOutputWithContext(ctx).OutputState,
	}
}

type InventoryItemOutput struct{ *pulumi.OutputState }

func (InventoryItemOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InventoryItem)(nil)).Elem()
}

func (o InventoryItemOutput) ToInventoryItemOutput() InventoryItemOutput {
	return o
}

func (o InventoryItemOutput) ToInventoryItemOutputWithContext(ctx context.Context) InventoryItemOutput {
	return o
}

func (o InventoryItemOutput) ToOutput(ctx context.Context) pulumix.Output[*InventoryItem] {
	return pulumix.Output[*InventoryItem]{
		OutputState: o.OutputState,
	}
}

// Gets the Managed Object name in VMM for the inventory item.
func (o InventoryItemOutput) InventoryItemName() pulumi.StringOutput {
	return o.ApplyT(func(v *InventoryItem) pulumi.StringOutput { return v.InventoryItemName }).(pulumi.StringOutput)
}

// They inventory type.
func (o InventoryItemOutput) InventoryType() pulumi.StringOutput {
	return o.ApplyT(func(v *InventoryItem) pulumi.StringOutput { return v.InventoryType }).(pulumi.StringOutput)
}

// Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type; e.g. ApiApps are a kind of Microsoft.Web/sites type.  If supported, the resource provider must validate and persist this value.
func (o InventoryItemOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InventoryItem) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Gets the tracked resource id corresponding to the inventory resource.
func (o InventoryItemOutput) ManagedResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *InventoryItem) pulumi.StringOutput { return v.ManagedResourceId }).(pulumi.StringOutput)
}

// The name of the resource
func (o InventoryItemOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *InventoryItem) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Gets the provisioning state.
func (o InventoryItemOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *InventoryItem) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o InventoryItemOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *InventoryItem) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o InventoryItemOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *InventoryItem) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Gets the UUID (which is assigned by VMM) for the inventory item.
func (o InventoryItemOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v *InventoryItem) pulumi.StringOutput { return v.Uuid }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(InventoryItemOutput{})
}
