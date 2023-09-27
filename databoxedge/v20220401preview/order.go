// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220401preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The order details.
type Order struct {
	pulumi.CustomResourceState

	// The contact details.
	ContactInformation ContactDetailsResponsePtrOutput `pulumi:"contactInformation"`
	// Current status of the order.
	CurrentStatus OrderStatusResponseOutput `pulumi:"currentStatus"`
	// Tracking information for the package delivered to the customer whether it has an original or a replacement device.
	DeliveryTrackingInfo TrackingInfoResponseArrayOutput `pulumi:"deliveryTrackingInfo"`
	// It specify the order api version.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// The object name.
	Name pulumi.StringOutput `pulumi:"name"`
	// List of status changes in the order.
	OrderHistory OrderStatusResponseArrayOutput `pulumi:"orderHistory"`
	// It specify the order resource id.
	OrderId pulumi.StringOutput `pulumi:"orderId"`
	// Tracking information for the package returned from the customer whether it has an original or a replacement device.
	ReturnTrackingInfo TrackingInfoResponseArrayOutput `pulumi:"returnTrackingInfo"`
	// Serial number of the device.
	SerialNumber pulumi.StringOutput `pulumi:"serialNumber"`
	// ShipmentType of the order
	ShipmentType pulumi.StringPtrOutput `pulumi:"shipmentType"`
	// The shipping address.
	ShippingAddress AddressResponsePtrOutput `pulumi:"shippingAddress"`
	// Metadata pertaining to creation and last modification of Order
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The hierarchical type of the object.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewOrder registers a new resource with the given unique name, arguments, and options.
func NewOrder(ctx *pulumi.Context,
	name string, args *OrderArgs, opts ...pulumi.ResourceOption) (*Order, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DeviceName == nil {
		return nil, errors.New("invalid value for required argument 'DeviceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:databoxedge:Order"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20190301:Order"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20190701:Order"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20190801:Order"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20200501preview:Order"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20200901:Order"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20200901preview:Order"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20201201:Order"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210201:Order"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210201preview:Order"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210601:Order"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210601preview:Order"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20220301:Order"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20221201preview:Order"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20230101preview:Order"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20230701:Order"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Order
	err := ctx.RegisterResource("azure-native:databoxedge/v20220401preview:Order", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOrder gets an existing Order resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOrder(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OrderState, opts ...pulumi.ResourceOption) (*Order, error) {
	var resource Order
	err := ctx.ReadResource("azure-native:databoxedge/v20220401preview:Order", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Order resources.
type orderState struct {
}

type OrderState struct {
}

func (OrderState) ElementType() reflect.Type {
	return reflect.TypeOf((*orderState)(nil)).Elem()
}

type orderArgs struct {
	// The contact details.
	ContactInformation *ContactDetails `pulumi:"contactInformation"`
	// The order details of a device.
	DeviceName string `pulumi:"deviceName"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// ShipmentType of the order
	ShipmentType *string `pulumi:"shipmentType"`
	// The shipping address.
	ShippingAddress *Address `pulumi:"shippingAddress"`
}

// The set of arguments for constructing a Order resource.
type OrderArgs struct {
	// The contact details.
	ContactInformation ContactDetailsPtrInput
	// The order details of a device.
	DeviceName pulumi.StringInput
	// The resource group name.
	ResourceGroupName pulumi.StringInput
	// ShipmentType of the order
	ShipmentType pulumi.StringPtrInput
	// The shipping address.
	ShippingAddress AddressPtrInput
}

func (OrderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*orderArgs)(nil)).Elem()
}

type OrderInput interface {
	pulumi.Input

	ToOrderOutput() OrderOutput
	ToOrderOutputWithContext(ctx context.Context) OrderOutput
}

func (*Order) ElementType() reflect.Type {
	return reflect.TypeOf((**Order)(nil)).Elem()
}

func (i *Order) ToOrderOutput() OrderOutput {
	return i.ToOrderOutputWithContext(context.Background())
}

func (i *Order) ToOrderOutputWithContext(ctx context.Context) OrderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrderOutput)
}

func (i *Order) ToOutput(ctx context.Context) pulumix.Output[*Order] {
	return pulumix.Output[*Order]{
		OutputState: i.ToOrderOutputWithContext(ctx).OutputState,
	}
}

type OrderOutput struct{ *pulumi.OutputState }

func (OrderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Order)(nil)).Elem()
}

func (o OrderOutput) ToOrderOutput() OrderOutput {
	return o
}

func (o OrderOutput) ToOrderOutputWithContext(ctx context.Context) OrderOutput {
	return o
}

func (o OrderOutput) ToOutput(ctx context.Context) pulumix.Output[*Order] {
	return pulumix.Output[*Order]{
		OutputState: o.OutputState,
	}
}

// The contact details.
func (o OrderOutput) ContactInformation() ContactDetailsResponsePtrOutput {
	return o.ApplyT(func(v *Order) ContactDetailsResponsePtrOutput { return v.ContactInformation }).(ContactDetailsResponsePtrOutput)
}

// Current status of the order.
func (o OrderOutput) CurrentStatus() OrderStatusResponseOutput {
	return o.ApplyT(func(v *Order) OrderStatusResponseOutput { return v.CurrentStatus }).(OrderStatusResponseOutput)
}

// Tracking information for the package delivered to the customer whether it has an original or a replacement device.
func (o OrderOutput) DeliveryTrackingInfo() TrackingInfoResponseArrayOutput {
	return o.ApplyT(func(v *Order) TrackingInfoResponseArrayOutput { return v.DeliveryTrackingInfo }).(TrackingInfoResponseArrayOutput)
}

// It specify the order api version.
func (o OrderOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *Order) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// The object name.
func (o OrderOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Order) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// List of status changes in the order.
func (o OrderOutput) OrderHistory() OrderStatusResponseArrayOutput {
	return o.ApplyT(func(v *Order) OrderStatusResponseArrayOutput { return v.OrderHistory }).(OrderStatusResponseArrayOutput)
}

// It specify the order resource id.
func (o OrderOutput) OrderId() pulumi.StringOutput {
	return o.ApplyT(func(v *Order) pulumi.StringOutput { return v.OrderId }).(pulumi.StringOutput)
}

// Tracking information for the package returned from the customer whether it has an original or a replacement device.
func (o OrderOutput) ReturnTrackingInfo() TrackingInfoResponseArrayOutput {
	return o.ApplyT(func(v *Order) TrackingInfoResponseArrayOutput { return v.ReturnTrackingInfo }).(TrackingInfoResponseArrayOutput)
}

// Serial number of the device.
func (o OrderOutput) SerialNumber() pulumi.StringOutput {
	return o.ApplyT(func(v *Order) pulumi.StringOutput { return v.SerialNumber }).(pulumi.StringOutput)
}

// ShipmentType of the order
func (o OrderOutput) ShipmentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Order) pulumi.StringPtrOutput { return v.ShipmentType }).(pulumi.StringPtrOutput)
}

// The shipping address.
func (o OrderOutput) ShippingAddress() AddressResponsePtrOutput {
	return o.ApplyT(func(v *Order) AddressResponsePtrOutput { return v.ShippingAddress }).(AddressResponsePtrOutput)
}

// Metadata pertaining to creation and last modification of Order
func (o OrderOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Order) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The hierarchical type of the object.
func (o OrderOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Order) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(OrderOutput{})
}
