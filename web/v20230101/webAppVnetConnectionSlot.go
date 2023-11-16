// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230101

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Virtual Network information ARM resource.
type WebAppVnetConnectionSlot struct {
	pulumi.CustomResourceState

	// A certificate file (.cer) blob containing the public key of the private key used to authenticate a
	// Point-To-Site VPN connection.
	CertBlob pulumi.StringPtrOutput `pulumi:"certBlob"`
	// The client certificate thumbprint.
	CertThumbprint pulumi.StringOutput `pulumi:"certThumbprint"`
	// DNS servers to be used by this Virtual Network. This should be a comma-separated list of IP addresses.
	DnsServers pulumi.StringPtrOutput `pulumi:"dnsServers"`
	// Flag that is used to denote if this is VNET injection
	IsSwift pulumi.BoolPtrOutput `pulumi:"isSwift"`
	// Kind of resource.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Resource Name.
	Name pulumi.StringOutput `pulumi:"name"`
	// <code>true</code> if a resync is required; otherwise, <code>false</code>.
	ResyncRequired pulumi.BoolOutput `pulumi:"resyncRequired"`
	// The routes that this Virtual Network connection uses.
	Routes VnetRouteResponseArrayOutput `pulumi:"routes"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// The Virtual Network's resource ID.
	VnetResourceId pulumi.StringPtrOutput `pulumi:"vnetResourceId"`
}

// NewWebAppVnetConnectionSlot registers a new resource with the given unique name, arguments, and options.
func NewWebAppVnetConnectionSlot(ctx *pulumi.Context,
	name string, args *WebAppVnetConnectionSlotArgs, opts ...pulumi.ResourceOption) (*WebAppVnetConnectionSlot, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Slot == nil {
		return nil, errors.New("invalid value for required argument 'Slot'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:web:WebAppVnetConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20150801:WebAppVnetConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20160801:WebAppVnetConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:WebAppVnetConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20181101:WebAppVnetConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:WebAppVnetConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:WebAppVnetConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:WebAppVnetConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:WebAppVnetConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:WebAppVnetConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:WebAppVnetConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:WebAppVnetConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:WebAppVnetConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210301:WebAppVnetConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220301:WebAppVnetConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220901:WebAppVnetConnectionSlot"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource WebAppVnetConnectionSlot
	err := ctx.RegisterResource("azure-native:web/v20230101:WebAppVnetConnectionSlot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebAppVnetConnectionSlot gets an existing WebAppVnetConnectionSlot resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebAppVnetConnectionSlot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppVnetConnectionSlotState, opts ...pulumi.ResourceOption) (*WebAppVnetConnectionSlot, error) {
	var resource WebAppVnetConnectionSlot
	err := ctx.ReadResource("azure-native:web/v20230101:WebAppVnetConnectionSlot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebAppVnetConnectionSlot resources.
type webAppVnetConnectionSlotState struct {
}

type WebAppVnetConnectionSlotState struct {
}

func (WebAppVnetConnectionSlotState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppVnetConnectionSlotState)(nil)).Elem()
}

type webAppVnetConnectionSlotArgs struct {
	// A certificate file (.cer) blob containing the public key of the private key used to authenticate a
	// Point-To-Site VPN connection.
	CertBlob *string `pulumi:"certBlob"`
	// DNS servers to be used by this Virtual Network. This should be a comma-separated list of IP addresses.
	DnsServers *string `pulumi:"dnsServers"`
	// Flag that is used to denote if this is VNET injection
	IsSwift *bool `pulumi:"isSwift"`
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Name of the app.
	Name string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the deployment slot. If a slot is not specified, the API will add or update connections for the production slot.
	Slot string `pulumi:"slot"`
	// Name of an existing Virtual Network.
	VnetName *string `pulumi:"vnetName"`
	// The Virtual Network's resource ID.
	VnetResourceId *string `pulumi:"vnetResourceId"`
}

// The set of arguments for constructing a WebAppVnetConnectionSlot resource.
type WebAppVnetConnectionSlotArgs struct {
	// A certificate file (.cer) blob containing the public key of the private key used to authenticate a
	// Point-To-Site VPN connection.
	CertBlob pulumi.StringPtrInput
	// DNS servers to be used by this Virtual Network. This should be a comma-separated list of IP addresses.
	DnsServers pulumi.StringPtrInput
	// Flag that is used to denote if this is VNET injection
	IsSwift pulumi.BoolPtrInput
	// Kind of resource.
	Kind pulumi.StringPtrInput
	// Name of the app.
	Name pulumi.StringInput
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput
	// Name of the deployment slot. If a slot is not specified, the API will add or update connections for the production slot.
	Slot pulumi.StringInput
	// Name of an existing Virtual Network.
	VnetName pulumi.StringPtrInput
	// The Virtual Network's resource ID.
	VnetResourceId pulumi.StringPtrInput
}

func (WebAppVnetConnectionSlotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppVnetConnectionSlotArgs)(nil)).Elem()
}

type WebAppVnetConnectionSlotInput interface {
	pulumi.Input

	ToWebAppVnetConnectionSlotOutput() WebAppVnetConnectionSlotOutput
	ToWebAppVnetConnectionSlotOutputWithContext(ctx context.Context) WebAppVnetConnectionSlotOutput
}

func (*WebAppVnetConnectionSlot) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppVnetConnectionSlot)(nil)).Elem()
}

func (i *WebAppVnetConnectionSlot) ToWebAppVnetConnectionSlotOutput() WebAppVnetConnectionSlotOutput {
	return i.ToWebAppVnetConnectionSlotOutputWithContext(context.Background())
}

func (i *WebAppVnetConnectionSlot) ToWebAppVnetConnectionSlotOutputWithContext(ctx context.Context) WebAppVnetConnectionSlotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAppVnetConnectionSlotOutput)
}

type WebAppVnetConnectionSlotOutput struct{ *pulumi.OutputState }

func (WebAppVnetConnectionSlotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppVnetConnectionSlot)(nil)).Elem()
}

func (o WebAppVnetConnectionSlotOutput) ToWebAppVnetConnectionSlotOutput() WebAppVnetConnectionSlotOutput {
	return o
}

func (o WebAppVnetConnectionSlotOutput) ToWebAppVnetConnectionSlotOutputWithContext(ctx context.Context) WebAppVnetConnectionSlotOutput {
	return o
}

// A certificate file (.cer) blob containing the public key of the private key used to authenticate a
// Point-To-Site VPN connection.
func (o WebAppVnetConnectionSlotOutput) CertBlob() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppVnetConnectionSlot) pulumi.StringPtrOutput { return v.CertBlob }).(pulumi.StringPtrOutput)
}

// The client certificate thumbprint.
func (o WebAppVnetConnectionSlotOutput) CertThumbprint() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppVnetConnectionSlot) pulumi.StringOutput { return v.CertThumbprint }).(pulumi.StringOutput)
}

// DNS servers to be used by this Virtual Network. This should be a comma-separated list of IP addresses.
func (o WebAppVnetConnectionSlotOutput) DnsServers() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppVnetConnectionSlot) pulumi.StringPtrOutput { return v.DnsServers }).(pulumi.StringPtrOutput)
}

// Flag that is used to denote if this is VNET injection
func (o WebAppVnetConnectionSlotOutput) IsSwift() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WebAppVnetConnectionSlot) pulumi.BoolPtrOutput { return v.IsSwift }).(pulumi.BoolPtrOutput)
}

// Kind of resource.
func (o WebAppVnetConnectionSlotOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppVnetConnectionSlot) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Resource Name.
func (o WebAppVnetConnectionSlotOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppVnetConnectionSlot) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// <code>true</code> if a resync is required; otherwise, <code>false</code>.
func (o WebAppVnetConnectionSlotOutput) ResyncRequired() pulumi.BoolOutput {
	return o.ApplyT(func(v *WebAppVnetConnectionSlot) pulumi.BoolOutput { return v.ResyncRequired }).(pulumi.BoolOutput)
}

// The routes that this Virtual Network connection uses.
func (o WebAppVnetConnectionSlotOutput) Routes() VnetRouteResponseArrayOutput {
	return o.ApplyT(func(v *WebAppVnetConnectionSlot) VnetRouteResponseArrayOutput { return v.Routes }).(VnetRouteResponseArrayOutput)
}

// Resource type.
func (o WebAppVnetConnectionSlotOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppVnetConnectionSlot) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The Virtual Network's resource ID.
func (o WebAppVnetConnectionSlotOutput) VnetResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppVnetConnectionSlot) pulumi.StringPtrOutput { return v.VnetResourceId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(WebAppVnetConnectionSlotOutput{})
}
