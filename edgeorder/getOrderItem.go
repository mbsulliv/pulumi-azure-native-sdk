// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package edgeorder

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get an order item.
// Azure REST API version: 2022-05-01-preview.
func LookupOrderItem(ctx *pulumi.Context, args *LookupOrderItemArgs, opts ...pulumi.InvokeOption) (*LookupOrderItemResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupOrderItemResult
	err := ctx.Invoke("azure-native:edgeorder:getOrderItem", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupOrderItemArgs struct {
	// $expand is supported on parent device details, device details, forward shipping details and reverse shipping details parameters. Each of these can be provided as a comma separated list. Parent Device Details for order item provides details on the devices of the product, Device Details for order item provides details on the devices of the child configurations of the product, Forward and Reverse Shipping details provide forward and reverse shipping details respectively.
	Expand *string `pulumi:"expand"`
	// The name of the order item.
	OrderItemName string `pulumi:"orderItemName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Represents order item resource.
type LookupOrderItemResult struct {
	// Represents shipping and return address for order item.
	AddressDetails AddressDetailsResponse `pulumi:"addressDetails"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Id of the order to which order item belongs to.
	OrderId string `pulumi:"orderId"`
	// Represents order item details.
	OrderItemDetails OrderItemDetailsResponse `pulumi:"orderItemDetails"`
	// Start time of order item.
	StartTime string `pulumi:"startTime"`
	// Represents resource creation and update time.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupOrderItemOutput(ctx *pulumi.Context, args LookupOrderItemOutputArgs, opts ...pulumi.InvokeOption) LookupOrderItemResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupOrderItemResult, error) {
			args := v.(LookupOrderItemArgs)
			r, err := LookupOrderItem(ctx, &args, opts...)
			var s LookupOrderItemResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupOrderItemResultOutput)
}

type LookupOrderItemOutputArgs struct {
	// $expand is supported on parent device details, device details, forward shipping details and reverse shipping details parameters. Each of these can be provided as a comma separated list. Parent Device Details for order item provides details on the devices of the product, Device Details for order item provides details on the devices of the child configurations of the product, Forward and Reverse Shipping details provide forward and reverse shipping details respectively.
	Expand pulumi.StringPtrInput `pulumi:"expand"`
	// The name of the order item.
	OrderItemName pulumi.StringInput `pulumi:"orderItemName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupOrderItemOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOrderItemArgs)(nil)).Elem()
}

// Represents order item resource.
type LookupOrderItemResultOutput struct{ *pulumi.OutputState }

func (LookupOrderItemResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOrderItemResult)(nil)).Elem()
}

func (o LookupOrderItemResultOutput) ToLookupOrderItemResultOutput() LookupOrderItemResultOutput {
	return o
}

func (o LookupOrderItemResultOutput) ToLookupOrderItemResultOutputWithContext(ctx context.Context) LookupOrderItemResultOutput {
	return o
}

// Represents shipping and return address for order item.
func (o LookupOrderItemResultOutput) AddressDetails() AddressDetailsResponseOutput {
	return o.ApplyT(func(v LookupOrderItemResult) AddressDetailsResponse { return v.AddressDetails }).(AddressDetailsResponseOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupOrderItemResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrderItemResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupOrderItemResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrderItemResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupOrderItemResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrderItemResult) string { return v.Name }).(pulumi.StringOutput)
}

// Id of the order to which order item belongs to.
func (o LookupOrderItemResultOutput) OrderId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrderItemResult) string { return v.OrderId }).(pulumi.StringOutput)
}

// Represents order item details.
func (o LookupOrderItemResultOutput) OrderItemDetails() OrderItemDetailsResponseOutput {
	return o.ApplyT(func(v LookupOrderItemResult) OrderItemDetailsResponse { return v.OrderItemDetails }).(OrderItemDetailsResponseOutput)
}

// Start time of order item.
func (o LookupOrderItemResultOutput) StartTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrderItemResult) string { return v.StartTime }).(pulumi.StringOutput)
}

// Represents resource creation and update time.
func (o LookupOrderItemResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupOrderItemResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupOrderItemResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupOrderItemResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupOrderItemResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrderItemResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupOrderItemResultOutput{})
}
