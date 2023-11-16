// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package managednetworkfabric

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the Neighbor Group.
// Azure REST API version: 2023-06-15.
func LookupNeighborGroup(ctx *pulumi.Context, args *LookupNeighborGroupArgs, opts ...pulumi.InvokeOption) (*LookupNeighborGroupResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupNeighborGroupResult
	err := ctx.Invoke("azure-native:managednetworkfabric:getNeighborGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNeighborGroupArgs struct {
	// Name of the Neighbor Group.
	NeighborGroupName string `pulumi:"neighborGroupName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Defines the Neighbor Group.
type LookupNeighborGroupResult struct {
	// Switch configuration description.
	Annotation *string `pulumi:"annotation"`
	// An array of destination IPv4 Addresses or IPv6 Addresses.
	Destination NeighborGroupDestinationResponse `pulumi:"destination"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// List of NetworkTap IDs where neighbor group is associated.
	NetworkTapIds []string `pulumi:"networkTapIds"`
	// List of Network Tap Rule IDs where neighbor group is associated.
	NetworkTapRuleIds []string `pulumi:"networkTapRuleIds"`
	// The provisioning state of the resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupNeighborGroupOutput(ctx *pulumi.Context, args LookupNeighborGroupOutputArgs, opts ...pulumi.InvokeOption) LookupNeighborGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNeighborGroupResult, error) {
			args := v.(LookupNeighborGroupArgs)
			r, err := LookupNeighborGroup(ctx, &args, opts...)
			var s LookupNeighborGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNeighborGroupResultOutput)
}

type LookupNeighborGroupOutputArgs struct {
	// Name of the Neighbor Group.
	NeighborGroupName pulumi.StringInput `pulumi:"neighborGroupName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupNeighborGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNeighborGroupArgs)(nil)).Elem()
}

// Defines the Neighbor Group.
type LookupNeighborGroupResultOutput struct{ *pulumi.OutputState }

func (LookupNeighborGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNeighborGroupResult)(nil)).Elem()
}

func (o LookupNeighborGroupResultOutput) ToLookupNeighborGroupResultOutput() LookupNeighborGroupResultOutput {
	return o
}

func (o LookupNeighborGroupResultOutput) ToLookupNeighborGroupResultOutputWithContext(ctx context.Context) LookupNeighborGroupResultOutput {
	return o
}

// Switch configuration description.
func (o LookupNeighborGroupResultOutput) Annotation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNeighborGroupResult) *string { return v.Annotation }).(pulumi.StringPtrOutput)
}

// An array of destination IPv4 Addresses or IPv6 Addresses.
func (o LookupNeighborGroupResultOutput) Destination() NeighborGroupDestinationResponseOutput {
	return o.ApplyT(func(v LookupNeighborGroupResult) NeighborGroupDestinationResponse { return v.Destination }).(NeighborGroupDestinationResponseOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupNeighborGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNeighborGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupNeighborGroupResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNeighborGroupResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupNeighborGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNeighborGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

// List of NetworkTap IDs where neighbor group is associated.
func (o LookupNeighborGroupResultOutput) NetworkTapIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupNeighborGroupResult) []string { return v.NetworkTapIds }).(pulumi.StringArrayOutput)
}

// List of Network Tap Rule IDs where neighbor group is associated.
func (o LookupNeighborGroupResultOutput) NetworkTapRuleIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupNeighborGroupResult) []string { return v.NetworkTapRuleIds }).(pulumi.StringArrayOutput)
}

// The provisioning state of the resource.
func (o LookupNeighborGroupResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNeighborGroupResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupNeighborGroupResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupNeighborGroupResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupNeighborGroupResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupNeighborGroupResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupNeighborGroupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNeighborGroupResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNeighborGroupResultOutput{})
}
