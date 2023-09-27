// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package relay

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Description of the WCF relay resource.
// Azure REST API version: 2021-11-01. Prior API version in Azure Native 1.x: 2017-04-01
type WCFRelay struct {
	pulumi.CustomResourceState

	// The time the WCF relay was created.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Returns true if the relay is dynamic; otherwise, false.
	IsDynamic pulumi.BoolOutput `pulumi:"isDynamic"`
	// The number of listeners for this relay. Note that min :1 and max:25 are supported.
	ListenerCount pulumi.IntOutput `pulumi:"listenerCount"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// WCF relay type.
	RelayType pulumi.StringPtrOutput `pulumi:"relayType"`
	// Returns true if client authorization is needed for this relay; otherwise, false.
	RequiresClientAuthorization pulumi.BoolPtrOutput `pulumi:"requiresClientAuthorization"`
	// Returns true if transport security is needed for this relay; otherwise, false.
	RequiresTransportSecurity pulumi.BoolPtrOutput `pulumi:"requiresTransportSecurity"`
	// The system meta data relating to this resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.EventHub/Namespaces" or "Microsoft.EventHub/Namespaces/EventHubs"
	Type pulumi.StringOutput `pulumi:"type"`
	// The time the namespace was updated.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
	// The usermetadata is a placeholder to store user-defined string data for the WCF Relay endpoint. For example, it can be used to store descriptive data, such as list of teams and their contact information. Also, user-defined configuration settings can be stored.
	UserMetadata pulumi.StringPtrOutput `pulumi:"userMetadata"`
}

// NewWCFRelay registers a new resource with the given unique name, arguments, and options.
func NewWCFRelay(ctx *pulumi.Context,
	name string, args *WCFRelayArgs, opts ...pulumi.ResourceOption) (*WCFRelay, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NamespaceName == nil {
		return nil, errors.New("invalid value for required argument 'NamespaceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:relay/v20160701:WCFRelay"),
		},
		{
			Type: pulumi.String("azure-native:relay/v20170401:WCFRelay"),
		},
		{
			Type: pulumi.String("azure-native:relay/v20211101:WCFRelay"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource WCFRelay
	err := ctx.RegisterResource("azure-native:relay:WCFRelay", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWCFRelay gets an existing WCFRelay resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWCFRelay(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WCFRelayState, opts ...pulumi.ResourceOption) (*WCFRelay, error) {
	var resource WCFRelay
	err := ctx.ReadResource("azure-native:relay:WCFRelay", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WCFRelay resources.
type wcfrelayState struct {
}

type WCFRelayState struct {
}

func (WCFRelayState) ElementType() reflect.Type {
	return reflect.TypeOf((*wcfrelayState)(nil)).Elem()
}

type wcfrelayArgs struct {
	// The namespace name
	NamespaceName string `pulumi:"namespaceName"`
	// The relay name.
	RelayName *string `pulumi:"relayName"`
	// WCF relay type.
	RelayType *Relaytype `pulumi:"relayType"`
	// Returns true if client authorization is needed for this relay; otherwise, false.
	RequiresClientAuthorization *bool `pulumi:"requiresClientAuthorization"`
	// Returns true if transport security is needed for this relay; otherwise, false.
	RequiresTransportSecurity *bool `pulumi:"requiresTransportSecurity"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The usermetadata is a placeholder to store user-defined string data for the WCF Relay endpoint. For example, it can be used to store descriptive data, such as list of teams and their contact information. Also, user-defined configuration settings can be stored.
	UserMetadata *string `pulumi:"userMetadata"`
}

// The set of arguments for constructing a WCFRelay resource.
type WCFRelayArgs struct {
	// The namespace name
	NamespaceName pulumi.StringInput
	// The relay name.
	RelayName pulumi.StringPtrInput
	// WCF relay type.
	RelayType RelaytypePtrInput
	// Returns true if client authorization is needed for this relay; otherwise, false.
	RequiresClientAuthorization pulumi.BoolPtrInput
	// Returns true if transport security is needed for this relay; otherwise, false.
	RequiresTransportSecurity pulumi.BoolPtrInput
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput
	// The usermetadata is a placeholder to store user-defined string data for the WCF Relay endpoint. For example, it can be used to store descriptive data, such as list of teams and their contact information. Also, user-defined configuration settings can be stored.
	UserMetadata pulumi.StringPtrInput
}

func (WCFRelayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*wcfrelayArgs)(nil)).Elem()
}

type WCFRelayInput interface {
	pulumi.Input

	ToWCFRelayOutput() WCFRelayOutput
	ToWCFRelayOutputWithContext(ctx context.Context) WCFRelayOutput
}

func (*WCFRelay) ElementType() reflect.Type {
	return reflect.TypeOf((**WCFRelay)(nil)).Elem()
}

func (i *WCFRelay) ToWCFRelayOutput() WCFRelayOutput {
	return i.ToWCFRelayOutputWithContext(context.Background())
}

func (i *WCFRelay) ToWCFRelayOutputWithContext(ctx context.Context) WCFRelayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WCFRelayOutput)
}

func (i *WCFRelay) ToOutput(ctx context.Context) pulumix.Output[*WCFRelay] {
	return pulumix.Output[*WCFRelay]{
		OutputState: i.ToWCFRelayOutputWithContext(ctx).OutputState,
	}
}

type WCFRelayOutput struct{ *pulumi.OutputState }

func (WCFRelayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WCFRelay)(nil)).Elem()
}

func (o WCFRelayOutput) ToWCFRelayOutput() WCFRelayOutput {
	return o
}

func (o WCFRelayOutput) ToWCFRelayOutputWithContext(ctx context.Context) WCFRelayOutput {
	return o
}

func (o WCFRelayOutput) ToOutput(ctx context.Context) pulumix.Output[*WCFRelay] {
	return pulumix.Output[*WCFRelay]{
		OutputState: o.OutputState,
	}
}

// The time the WCF relay was created.
func (o WCFRelayOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *WCFRelay) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// Returns true if the relay is dynamic; otherwise, false.
func (o WCFRelayOutput) IsDynamic() pulumi.BoolOutput {
	return o.ApplyT(func(v *WCFRelay) pulumi.BoolOutput { return v.IsDynamic }).(pulumi.BoolOutput)
}

// The number of listeners for this relay. Note that min :1 and max:25 are supported.
func (o WCFRelayOutput) ListenerCount() pulumi.IntOutput {
	return o.ApplyT(func(v *WCFRelay) pulumi.IntOutput { return v.ListenerCount }).(pulumi.IntOutput)
}

// The geo-location where the resource lives
func (o WCFRelayOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *WCFRelay) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o WCFRelayOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WCFRelay) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// WCF relay type.
func (o WCFRelayOutput) RelayType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WCFRelay) pulumi.StringPtrOutput { return v.RelayType }).(pulumi.StringPtrOutput)
}

// Returns true if client authorization is needed for this relay; otherwise, false.
func (o WCFRelayOutput) RequiresClientAuthorization() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WCFRelay) pulumi.BoolPtrOutput { return v.RequiresClientAuthorization }).(pulumi.BoolPtrOutput)
}

// Returns true if transport security is needed for this relay; otherwise, false.
func (o WCFRelayOutput) RequiresTransportSecurity() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WCFRelay) pulumi.BoolPtrOutput { return v.RequiresTransportSecurity }).(pulumi.BoolPtrOutput)
}

// The system meta data relating to this resource.
func (o WCFRelayOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *WCFRelay) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.EventHub/Namespaces" or "Microsoft.EventHub/Namespaces/EventHubs"
func (o WCFRelayOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WCFRelay) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The time the namespace was updated.
func (o WCFRelayOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *WCFRelay) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

// The usermetadata is a placeholder to store user-defined string data for the WCF Relay endpoint. For example, it can be used to store descriptive data, such as list of teams and their contact information. Also, user-defined configuration settings can be stored.
func (o WCFRelayOutput) UserMetadata() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WCFRelay) pulumi.StringPtrOutput { return v.UserMetadata }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(WCFRelayOutput{})
}
