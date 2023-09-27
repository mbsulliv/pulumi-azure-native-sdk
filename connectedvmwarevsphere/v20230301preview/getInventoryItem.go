// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230301preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Implements InventoryItem GET method.
func LookupInventoryItem(ctx *pulumi.Context, args *LookupInventoryItemArgs, opts ...pulumi.InvokeOption) (*LookupInventoryItemResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupInventoryItemResult
	err := ctx.Invoke("azure-native:connectedvmwarevsphere/v20230301preview:getInventoryItem", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupInventoryItemArgs struct {
	// Name of the inventoryItem.
	InventoryItemName string `pulumi:"inventoryItemName"`
	// The Resource Group Name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the vCenter.
	VcenterName string `pulumi:"vcenterName"`
}

// Defines the inventory item.
type LookupInventoryItemResult struct {
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// They inventory type.
	InventoryType string `pulumi:"inventoryType"`
	// Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type; e.g. ApiApps are a kind of Microsoft.Web/sites type.  If supported, the resource provider must validate and persist this value.
	Kind *string `pulumi:"kind"`
	// Gets or sets the tracked resource id corresponding to the inventory resource.
	ManagedResourceId *string `pulumi:"managedResourceId"`
	// Gets or sets the vCenter Managed Object name for the inventory item.
	MoName *string `pulumi:"moName"`
	// Gets or sets the MoRef (Managed Object Reference) ID for the inventory item.
	MoRefId *string `pulumi:"moRefId"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Gets the provisioning state.
	ProvisioningState string `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupInventoryItemOutput(ctx *pulumi.Context, args LookupInventoryItemOutputArgs, opts ...pulumi.InvokeOption) LookupInventoryItemResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupInventoryItemResult, error) {
			args := v.(LookupInventoryItemArgs)
			r, err := LookupInventoryItem(ctx, &args, opts...)
			var s LookupInventoryItemResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupInventoryItemResultOutput)
}

type LookupInventoryItemOutputArgs struct {
	// Name of the inventoryItem.
	InventoryItemName pulumi.StringInput `pulumi:"inventoryItemName"`
	// The Resource Group Name.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name of the vCenter.
	VcenterName pulumi.StringInput `pulumi:"vcenterName"`
}

func (LookupInventoryItemOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInventoryItemArgs)(nil)).Elem()
}

// Defines the inventory item.
type LookupInventoryItemResultOutput struct{ *pulumi.OutputState }

func (LookupInventoryItemResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInventoryItemResult)(nil)).Elem()
}

func (o LookupInventoryItemResultOutput) ToLookupInventoryItemResultOutput() LookupInventoryItemResultOutput {
	return o
}

func (o LookupInventoryItemResultOutput) ToLookupInventoryItemResultOutputWithContext(ctx context.Context) LookupInventoryItemResultOutput {
	return o
}

func (o LookupInventoryItemResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupInventoryItemResult] {
	return pulumix.Output[LookupInventoryItemResult]{
		OutputState: o.OutputState,
	}
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupInventoryItemResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInventoryItemResult) string { return v.Id }).(pulumi.StringOutput)
}

// They inventory type.
func (o LookupInventoryItemResultOutput) InventoryType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInventoryItemResult) string { return v.InventoryType }).(pulumi.StringOutput)
}

// Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type; e.g. ApiApps are a kind of Microsoft.Web/sites type.  If supported, the resource provider must validate and persist this value.
func (o LookupInventoryItemResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInventoryItemResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// Gets or sets the tracked resource id corresponding to the inventory resource.
func (o LookupInventoryItemResultOutput) ManagedResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInventoryItemResult) *string { return v.ManagedResourceId }).(pulumi.StringPtrOutput)
}

// Gets or sets the vCenter Managed Object name for the inventory item.
func (o LookupInventoryItemResultOutput) MoName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInventoryItemResult) *string { return v.MoName }).(pulumi.StringPtrOutput)
}

// Gets or sets the MoRef (Managed Object Reference) ID for the inventory item.
func (o LookupInventoryItemResultOutput) MoRefId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInventoryItemResult) *string { return v.MoRefId }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o LookupInventoryItemResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInventoryItemResult) string { return v.Name }).(pulumi.StringOutput)
}

// Gets the provisioning state.
func (o LookupInventoryItemResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInventoryItemResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupInventoryItemResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupInventoryItemResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupInventoryItemResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInventoryItemResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupInventoryItemResultOutput{})
}
