// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azurestack

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Linked Subscription information.
// Azure REST API version: 2020-06-01-preview. Prior API version in Azure Native 1.x: 2020-06-01-preview.
type LinkedSubscription struct {
	pulumi.CustomResourceState

	// The status of the remote management connection of the Azure Stack device.
	DeviceConnectionStatus pulumi.StringOutput `pulumi:"deviceConnectionStatus"`
	// The identifier of the Azure Stack device for remote management.
	DeviceId pulumi.StringOutput `pulumi:"deviceId"`
	// The connection state of the Azure Stack device.
	DeviceLinkState pulumi.StringOutput `pulumi:"deviceLinkState"`
	// The object identifier associated with the Azure Stack device connecting to Azure.
	DeviceObjectId pulumi.StringOutput `pulumi:"deviceObjectId"`
	// The entity tag used for optimistic concurrency when modifying the resource.
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// The kind of the resource.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// The last remote management connection time for the Azure Stack device connected to the linked subscription resource.
	LastConnectedTime pulumi.StringOutput `pulumi:"lastConnectedTime"`
	// The identifier associated with the device subscription.
	LinkedSubscriptionId pulumi.StringPtrOutput `pulumi:"linkedSubscriptionId"`
	// Location of the resource.
	Location pulumi.StringOutput `pulumi:"location"`
	// Name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The identifier associated with the device registration.
	RegistrationResourceId pulumi.StringPtrOutput `pulumi:"registrationResourceId"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Custom tags for the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Type of Resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewLinkedSubscription registers a new resource with the given unique name, arguments, and options.
func NewLinkedSubscription(ctx *pulumi.Context,
	name string, args *LinkedSubscriptionArgs, opts ...pulumi.ResourceOption) (*LinkedSubscription, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LinkedSubscriptionId == nil {
		return nil, errors.New("invalid value for required argument 'LinkedSubscriptionId'")
	}
	if args.RegistrationResourceId == nil {
		return nil, errors.New("invalid value for required argument 'RegistrationResourceId'")
	}
	if args.ResourceGroup == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroup'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:azurestack/v20200601preview:LinkedSubscription"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource LinkedSubscription
	err := ctx.RegisterResource("azure-native:azurestack:LinkedSubscription", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLinkedSubscription gets an existing LinkedSubscription resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLinkedSubscription(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LinkedSubscriptionState, opts ...pulumi.ResourceOption) (*LinkedSubscription, error) {
	var resource LinkedSubscription
	err := ctx.ReadResource("azure-native:azurestack:LinkedSubscription", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LinkedSubscription resources.
type linkedSubscriptionState struct {
}

type LinkedSubscriptionState struct {
}

func (LinkedSubscriptionState) ElementType() reflect.Type {
	return reflect.TypeOf((*linkedSubscriptionState)(nil)).Elem()
}

type linkedSubscriptionArgs struct {
	// The identifier associated with the device subscription.
	LinkedSubscriptionId string `pulumi:"linkedSubscriptionId"`
	// Name of the Linked Subscription resource.
	LinkedSubscriptionName *string `pulumi:"linkedSubscriptionName"`
	// Location of the resource.
	Location *string `pulumi:"location"`
	// The identifier associated with the device registration.
	RegistrationResourceId string `pulumi:"registrationResourceId"`
	// Name of the resource group.
	ResourceGroup string `pulumi:"resourceGroup"`
}

// The set of arguments for constructing a LinkedSubscription resource.
type LinkedSubscriptionArgs struct {
	// The identifier associated with the device subscription.
	LinkedSubscriptionId pulumi.StringInput
	// Name of the Linked Subscription resource.
	LinkedSubscriptionName pulumi.StringPtrInput
	// Location of the resource.
	Location pulumi.StringPtrInput
	// The identifier associated with the device registration.
	RegistrationResourceId pulumi.StringInput
	// Name of the resource group.
	ResourceGroup pulumi.StringInput
}

func (LinkedSubscriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*linkedSubscriptionArgs)(nil)).Elem()
}

type LinkedSubscriptionInput interface {
	pulumi.Input

	ToLinkedSubscriptionOutput() LinkedSubscriptionOutput
	ToLinkedSubscriptionOutputWithContext(ctx context.Context) LinkedSubscriptionOutput
}

func (*LinkedSubscription) ElementType() reflect.Type {
	return reflect.TypeOf((**LinkedSubscription)(nil)).Elem()
}

func (i *LinkedSubscription) ToLinkedSubscriptionOutput() LinkedSubscriptionOutput {
	return i.ToLinkedSubscriptionOutputWithContext(context.Background())
}

func (i *LinkedSubscription) ToLinkedSubscriptionOutputWithContext(ctx context.Context) LinkedSubscriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedSubscriptionOutput)
}

func (i *LinkedSubscription) ToOutput(ctx context.Context) pulumix.Output[*LinkedSubscription] {
	return pulumix.Output[*LinkedSubscription]{
		OutputState: i.ToLinkedSubscriptionOutputWithContext(ctx).OutputState,
	}
}

type LinkedSubscriptionOutput struct{ *pulumi.OutputState }

func (LinkedSubscriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LinkedSubscription)(nil)).Elem()
}

func (o LinkedSubscriptionOutput) ToLinkedSubscriptionOutput() LinkedSubscriptionOutput {
	return o
}

func (o LinkedSubscriptionOutput) ToLinkedSubscriptionOutputWithContext(ctx context.Context) LinkedSubscriptionOutput {
	return o
}

func (o LinkedSubscriptionOutput) ToOutput(ctx context.Context) pulumix.Output[*LinkedSubscription] {
	return pulumix.Output[*LinkedSubscription]{
		OutputState: o.OutputState,
	}
}

// The status of the remote management connection of the Azure Stack device.
func (o LinkedSubscriptionOutput) DeviceConnectionStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *LinkedSubscription) pulumi.StringOutput { return v.DeviceConnectionStatus }).(pulumi.StringOutput)
}

// The identifier of the Azure Stack device for remote management.
func (o LinkedSubscriptionOutput) DeviceId() pulumi.StringOutput {
	return o.ApplyT(func(v *LinkedSubscription) pulumi.StringOutput { return v.DeviceId }).(pulumi.StringOutput)
}

// The connection state of the Azure Stack device.
func (o LinkedSubscriptionOutput) DeviceLinkState() pulumi.StringOutput {
	return o.ApplyT(func(v *LinkedSubscription) pulumi.StringOutput { return v.DeviceLinkState }).(pulumi.StringOutput)
}

// The object identifier associated with the Azure Stack device connecting to Azure.
func (o LinkedSubscriptionOutput) DeviceObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *LinkedSubscription) pulumi.StringOutput { return v.DeviceObjectId }).(pulumi.StringOutput)
}

// The entity tag used for optimistic concurrency when modifying the resource.
func (o LinkedSubscriptionOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LinkedSubscription) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// The kind of the resource.
func (o LinkedSubscriptionOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *LinkedSubscription) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// The last remote management connection time for the Azure Stack device connected to the linked subscription resource.
func (o LinkedSubscriptionOutput) LastConnectedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *LinkedSubscription) pulumi.StringOutput { return v.LastConnectedTime }).(pulumi.StringOutput)
}

// The identifier associated with the device subscription.
func (o LinkedSubscriptionOutput) LinkedSubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LinkedSubscription) pulumi.StringPtrOutput { return v.LinkedSubscriptionId }).(pulumi.StringPtrOutput)
}

// Location of the resource.
func (o LinkedSubscriptionOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *LinkedSubscription) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Name of the resource.
func (o LinkedSubscriptionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *LinkedSubscription) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The identifier associated with the device registration.
func (o LinkedSubscriptionOutput) RegistrationResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LinkedSubscription) pulumi.StringPtrOutput { return v.RegistrationResourceId }).(pulumi.StringPtrOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LinkedSubscriptionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *LinkedSubscription) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Custom tags for the resource.
func (o LinkedSubscriptionOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *LinkedSubscription) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Type of Resource.
func (o LinkedSubscriptionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *LinkedSubscription) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LinkedSubscriptionOutput{})
}
