// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20190601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The Get ManagedNetworkGroups operation gets a Managed Network Group specified by the resource group, Managed Network name, and group name
func LookupManagedNetworkGroup(ctx *pulumi.Context, args *LookupManagedNetworkGroupArgs, opts ...pulumi.InvokeOption) (*LookupManagedNetworkGroupResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupManagedNetworkGroupResult
	err := ctx.Invoke("azure-native:managednetwork/v20190601preview:getManagedNetworkGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupManagedNetworkGroupArgs struct {
	// The name of the Managed Network Group.
	ManagedNetworkGroupName string `pulumi:"managedNetworkGroupName"`
	// The name of the Managed Network.
	ManagedNetworkName string `pulumi:"managedNetworkName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The Managed Network Group resource
type LookupManagedNetworkGroupResult struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// Responsibility role under which this Managed Network Group will be created
	Kind *string `pulumi:"kind"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The collection of management groups covered by the Managed Network
	ManagementGroups []ResourceIdResponse `pulumi:"managementGroups"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Provisioning state of the ManagedNetwork resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// The collection of  subnets covered by the Managed Network
	Subnets []ResourceIdResponse `pulumi:"subnets"`
	// The collection of subscriptions covered by the Managed Network
	Subscriptions []ResourceIdResponse `pulumi:"subscriptions"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type string `pulumi:"type"`
	// The collection of virtual nets covered by the Managed Network
	VirtualNetworks []ResourceIdResponse `pulumi:"virtualNetworks"`
}

func LookupManagedNetworkGroupOutput(ctx *pulumi.Context, args LookupManagedNetworkGroupOutputArgs, opts ...pulumi.InvokeOption) LookupManagedNetworkGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupManagedNetworkGroupResult, error) {
			args := v.(LookupManagedNetworkGroupArgs)
			r, err := LookupManagedNetworkGroup(ctx, &args, opts...)
			var s LookupManagedNetworkGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupManagedNetworkGroupResultOutput)
}

type LookupManagedNetworkGroupOutputArgs struct {
	// The name of the Managed Network Group.
	ManagedNetworkGroupName pulumi.StringInput `pulumi:"managedNetworkGroupName"`
	// The name of the Managed Network.
	ManagedNetworkName pulumi.StringInput `pulumi:"managedNetworkName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupManagedNetworkGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagedNetworkGroupArgs)(nil)).Elem()
}

// The Managed Network Group resource
type LookupManagedNetworkGroupResultOutput struct{ *pulumi.OutputState }

func (LookupManagedNetworkGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagedNetworkGroupResult)(nil)).Elem()
}

func (o LookupManagedNetworkGroupResultOutput) ToLookupManagedNetworkGroupResultOutput() LookupManagedNetworkGroupResultOutput {
	return o
}

func (o LookupManagedNetworkGroupResultOutput) ToLookupManagedNetworkGroupResultOutputWithContext(ctx context.Context) LookupManagedNetworkGroupResultOutput {
	return o
}

// A unique read-only string that changes whenever the resource is updated.
func (o LookupManagedNetworkGroupResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedNetworkGroupResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupManagedNetworkGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedNetworkGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

// Responsibility role under which this Managed Network Group will be created
func (o LookupManagedNetworkGroupResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedNetworkGroupResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// The geo-location where the resource lives
func (o LookupManagedNetworkGroupResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedNetworkGroupResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// The collection of management groups covered by the Managed Network
func (o LookupManagedNetworkGroupResultOutput) ManagementGroups() ResourceIdResponseArrayOutput {
	return o.ApplyT(func(v LookupManagedNetworkGroupResult) []ResourceIdResponse { return v.ManagementGroups }).(ResourceIdResponseArrayOutput)
}

// The name of the resource
func (o LookupManagedNetworkGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedNetworkGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

// Provisioning state of the ManagedNetwork resource.
func (o LookupManagedNetworkGroupResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedNetworkGroupResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The collection of  subnets covered by the Managed Network
func (o LookupManagedNetworkGroupResultOutput) Subnets() ResourceIdResponseArrayOutput {
	return o.ApplyT(func(v LookupManagedNetworkGroupResult) []ResourceIdResponse { return v.Subnets }).(ResourceIdResponseArrayOutput)
}

// The collection of subscriptions covered by the Managed Network
func (o LookupManagedNetworkGroupResultOutput) Subscriptions() ResourceIdResponseArrayOutput {
	return o.ApplyT(func(v LookupManagedNetworkGroupResult) []ResourceIdResponse { return v.Subscriptions }).(ResourceIdResponseArrayOutput)
}

// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
func (o LookupManagedNetworkGroupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedNetworkGroupResult) string { return v.Type }).(pulumi.StringOutput)
}

// The collection of virtual nets covered by the Managed Network
func (o LookupManagedNetworkGroupResultOutput) VirtualNetworks() ResourceIdResponseArrayOutput {
	return o.ApplyT(func(v LookupManagedNetworkGroupResult) []ResourceIdResponse { return v.VirtualNetworks }).(ResourceIdResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupManagedNetworkGroupResultOutput{})
}
