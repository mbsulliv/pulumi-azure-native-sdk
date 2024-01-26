// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20231201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The order details.
func LookupOrder(ctx *pulumi.Context, args *LookupOrderArgs, opts ...pulumi.InvokeOption) (*LookupOrderResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupOrderResult
	err := ctx.Invoke("azure-native:databoxedge/v20231201:getOrder", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupOrderArgs struct {
	// The device name.
	DeviceName string `pulumi:"deviceName"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The order details.
type LookupOrderResult struct {
	// The contact details.
	ContactInformation ContactDetailsResponse `pulumi:"contactInformation"`
	// Current status of the order.
	CurrentStatus OrderStatusResponse `pulumi:"currentStatus"`
	// Tracking information for the package delivered to the customer whether it has an original or a replacement device.
	DeliveryTrackingInfo []TrackingInfoResponse `pulumi:"deliveryTrackingInfo"`
	// The path ID that uniquely identifies the object.
	Id string `pulumi:"id"`
	// It specify the order api version.
	Kind string `pulumi:"kind"`
	// The object name.
	Name string `pulumi:"name"`
	// List of status changes in the order.
	OrderHistory []OrderStatusResponse `pulumi:"orderHistory"`
	// It specify the order resource id.
	OrderId string `pulumi:"orderId"`
	// Tracking information for the package returned from the customer whether it has an original or a replacement device.
	ReturnTrackingInfo []TrackingInfoResponse `pulumi:"returnTrackingInfo"`
	// Serial number of the device.
	SerialNumber string `pulumi:"serialNumber"`
	// ShipmentType of the order
	ShipmentType *string `pulumi:"shipmentType"`
	// The shipping address.
	ShippingAddress *AddressResponse `pulumi:"shippingAddress"`
	// Metadata pertaining to creation and last modification of Order
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The hierarchical type of the object.
	Type string `pulumi:"type"`
}

func LookupOrderOutput(ctx *pulumi.Context, args LookupOrderOutputArgs, opts ...pulumi.InvokeOption) LookupOrderResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupOrderResult, error) {
			args := v.(LookupOrderArgs)
			r, err := LookupOrder(ctx, &args, opts...)
			var s LookupOrderResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupOrderResultOutput)
}

type LookupOrderOutputArgs struct {
	// The device name.
	DeviceName pulumi.StringInput `pulumi:"deviceName"`
	// The resource group name.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupOrderOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOrderArgs)(nil)).Elem()
}

// The order details.
type LookupOrderResultOutput struct{ *pulumi.OutputState }

func (LookupOrderResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOrderResult)(nil)).Elem()
}

func (o LookupOrderResultOutput) ToLookupOrderResultOutput() LookupOrderResultOutput {
	return o
}

func (o LookupOrderResultOutput) ToLookupOrderResultOutputWithContext(ctx context.Context) LookupOrderResultOutput {
	return o
}

// The contact details.
func (o LookupOrderResultOutput) ContactInformation() ContactDetailsResponseOutput {
	return o.ApplyT(func(v LookupOrderResult) ContactDetailsResponse { return v.ContactInformation }).(ContactDetailsResponseOutput)
}

// Current status of the order.
func (o LookupOrderResultOutput) CurrentStatus() OrderStatusResponseOutput {
	return o.ApplyT(func(v LookupOrderResult) OrderStatusResponse { return v.CurrentStatus }).(OrderStatusResponseOutput)
}

// Tracking information for the package delivered to the customer whether it has an original or a replacement device.
func (o LookupOrderResultOutput) DeliveryTrackingInfo() TrackingInfoResponseArrayOutput {
	return o.ApplyT(func(v LookupOrderResult) []TrackingInfoResponse { return v.DeliveryTrackingInfo }).(TrackingInfoResponseArrayOutput)
}

// The path ID that uniquely identifies the object.
func (o LookupOrderResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrderResult) string { return v.Id }).(pulumi.StringOutput)
}

// It specify the order api version.
func (o LookupOrderResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrderResult) string { return v.Kind }).(pulumi.StringOutput)
}

// The object name.
func (o LookupOrderResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrderResult) string { return v.Name }).(pulumi.StringOutput)
}

// List of status changes in the order.
func (o LookupOrderResultOutput) OrderHistory() OrderStatusResponseArrayOutput {
	return o.ApplyT(func(v LookupOrderResult) []OrderStatusResponse { return v.OrderHistory }).(OrderStatusResponseArrayOutput)
}

// It specify the order resource id.
func (o LookupOrderResultOutput) OrderId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrderResult) string { return v.OrderId }).(pulumi.StringOutput)
}

// Tracking information for the package returned from the customer whether it has an original or a replacement device.
func (o LookupOrderResultOutput) ReturnTrackingInfo() TrackingInfoResponseArrayOutput {
	return o.ApplyT(func(v LookupOrderResult) []TrackingInfoResponse { return v.ReturnTrackingInfo }).(TrackingInfoResponseArrayOutput)
}

// Serial number of the device.
func (o LookupOrderResultOutput) SerialNumber() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrderResult) string { return v.SerialNumber }).(pulumi.StringOutput)
}

// ShipmentType of the order
func (o LookupOrderResultOutput) ShipmentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupOrderResult) *string { return v.ShipmentType }).(pulumi.StringPtrOutput)
}

// The shipping address.
func (o LookupOrderResultOutput) ShippingAddress() AddressResponsePtrOutput {
	return o.ApplyT(func(v LookupOrderResult) *AddressResponse { return v.ShippingAddress }).(AddressResponsePtrOutput)
}

// Metadata pertaining to creation and last modification of Order
func (o LookupOrderResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupOrderResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The hierarchical type of the object.
func (o LookupOrderResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrderResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupOrderResultOutput{})
}
