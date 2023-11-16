// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package testbase

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The Customer Notification Event resource.
// Azure REST API version: 2022-04-01-preview. Prior API version in Azure Native 1.x: 2022-04-01-preview.
type CustomerEvent struct {
	pulumi.CustomResourceState

	// The name of the event subscribed to.
	EventName pulumi.StringOutput `pulumi:"eventName"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The notification event receivers.
	Receivers NotificationEventReceiverResponseArrayOutput `pulumi:"receivers"`
	// The system metadata relating to this resource
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewCustomerEvent registers a new resource with the given unique name, arguments, and options.
func NewCustomerEvent(ctx *pulumi.Context,
	name string, args *CustomerEventArgs, opts ...pulumi.ResourceOption) (*CustomerEvent, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EventName == nil {
		return nil, errors.New("invalid value for required argument 'EventName'")
	}
	if args.Receivers == nil {
		return nil, errors.New("invalid value for required argument 'Receivers'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.TestBaseAccountName == nil {
		return nil, errors.New("invalid value for required argument 'TestBaseAccountName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:testbase/v20201216preview:CustomerEvent"),
		},
		{
			Type: pulumi.String("azure-native:testbase/v20220401preview:CustomerEvent"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource CustomerEvent
	err := ctx.RegisterResource("azure-native:testbase:CustomerEvent", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCustomerEvent gets an existing CustomerEvent resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCustomerEvent(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CustomerEventState, opts ...pulumi.ResourceOption) (*CustomerEvent, error) {
	var resource CustomerEvent
	err := ctx.ReadResource("azure-native:testbase:CustomerEvent", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CustomerEvent resources.
type customerEventState struct {
}

type CustomerEventState struct {
}

func (CustomerEventState) ElementType() reflect.Type {
	return reflect.TypeOf((*customerEventState)(nil)).Elem()
}

type customerEventArgs struct {
	// The resource name of the Test Base Customer event.
	CustomerEventName *string `pulumi:"customerEventName"`
	// The name of the event subscribed to.
	EventName string `pulumi:"eventName"`
	// The notification event receivers.
	Receivers []NotificationEventReceiver `pulumi:"receivers"`
	// The name of the resource group that contains the resource.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The resource name of the Test Base Account.
	TestBaseAccountName string `pulumi:"testBaseAccountName"`
}

// The set of arguments for constructing a CustomerEvent resource.
type CustomerEventArgs struct {
	// The resource name of the Test Base Customer event.
	CustomerEventName pulumi.StringPtrInput
	// The name of the event subscribed to.
	EventName pulumi.StringInput
	// The notification event receivers.
	Receivers NotificationEventReceiverArrayInput
	// The name of the resource group that contains the resource.
	ResourceGroupName pulumi.StringInput
	// The resource name of the Test Base Account.
	TestBaseAccountName pulumi.StringInput
}

func (CustomerEventArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*customerEventArgs)(nil)).Elem()
}

type CustomerEventInput interface {
	pulumi.Input

	ToCustomerEventOutput() CustomerEventOutput
	ToCustomerEventOutputWithContext(ctx context.Context) CustomerEventOutput
}

func (*CustomerEvent) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomerEvent)(nil)).Elem()
}

func (i *CustomerEvent) ToCustomerEventOutput() CustomerEventOutput {
	return i.ToCustomerEventOutputWithContext(context.Background())
}

func (i *CustomerEvent) ToCustomerEventOutputWithContext(ctx context.Context) CustomerEventOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomerEventOutput)
}

type CustomerEventOutput struct{ *pulumi.OutputState }

func (CustomerEventOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomerEvent)(nil)).Elem()
}

func (o CustomerEventOutput) ToCustomerEventOutput() CustomerEventOutput {
	return o
}

func (o CustomerEventOutput) ToCustomerEventOutputWithContext(ctx context.Context) CustomerEventOutput {
	return o
}

// The name of the event subscribed to.
func (o CustomerEventOutput) EventName() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomerEvent) pulumi.StringOutput { return v.EventName }).(pulumi.StringOutput)
}

// Resource name.
func (o CustomerEventOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomerEvent) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The notification event receivers.
func (o CustomerEventOutput) Receivers() NotificationEventReceiverResponseArrayOutput {
	return o.ApplyT(func(v *CustomerEvent) NotificationEventReceiverResponseArrayOutput { return v.Receivers }).(NotificationEventReceiverResponseArrayOutput)
}

// The system metadata relating to this resource
func (o CustomerEventOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *CustomerEvent) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource type.
func (o CustomerEventOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomerEvent) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(CustomerEventOutput{})
}
