// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230401

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Bastion Host resource.
type BastionHost struct {
	pulumi.CustomResourceState

	// Enable/Disable Copy/Paste feature of the Bastion Host resource.
	DisableCopyPaste pulumi.BoolPtrOutput `pulumi:"disableCopyPaste"`
	// FQDN for the endpoint on which bastion host is accessible.
	DnsName pulumi.StringPtrOutput `pulumi:"dnsName"`
	// Enable/Disable File Copy feature of the Bastion Host resource.
	EnableFileCopy pulumi.BoolPtrOutput `pulumi:"enableFileCopy"`
	// Enable/Disable IP Connect feature of the Bastion Host resource.
	EnableIpConnect pulumi.BoolPtrOutput `pulumi:"enableIpConnect"`
	// Enable/Disable Kerberos feature of the Bastion Host resource.
	EnableKerberos pulumi.BoolPtrOutput `pulumi:"enableKerberos"`
	// Enable/Disable Shareable Link of the Bastion Host resource.
	EnableShareableLink pulumi.BoolPtrOutput `pulumi:"enableShareableLink"`
	// Enable/Disable Tunneling feature of the Bastion Host resource.
	EnableTunneling pulumi.BoolPtrOutput `pulumi:"enableTunneling"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// IP configuration of the Bastion Host resource.
	IpConfigurations BastionHostIPConfigurationResponseArrayOutput `pulumi:"ipConfigurations"`
	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioning state of the bastion host resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The scale units for the Bastion Host resource.
	ScaleUnits pulumi.IntPtrOutput `pulumi:"scaleUnits"`
	// The sku of this Bastion Host.
	Sku SkuResponsePtrOutput `pulumi:"sku"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewBastionHost registers a new resource with the given unique name, arguments, and options.
func NewBastionHost(ctx *pulumi.Context,
	name string, args *BastionHostArgs, opts ...pulumi.ResourceOption) (*BastionHost, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.DisableCopyPaste == nil {
		args.DisableCopyPaste = pulumi.BoolPtr(false)
	}
	if args.EnableFileCopy == nil {
		args.EnableFileCopy = pulumi.BoolPtr(false)
	}
	if args.EnableIpConnect == nil {
		args.EnableIpConnect = pulumi.BoolPtr(false)
	}
	if args.EnableKerberos == nil {
		args.EnableKerberos = pulumi.BoolPtr(false)
	}
	if args.EnableShareableLink == nil {
		args.EnableShareableLink = pulumi.BoolPtr(false)
	}
	if args.EnableTunneling == nil {
		args.EnableTunneling = pulumi.BoolPtr(false)
	}
	if args.Sku != nil {
		args.Sku = args.Sku.ToSkuPtrOutput().ApplyT(func(v *Sku) *Sku { return v.Defaults() }).(SkuPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:BastionHost"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190401:BastionHost"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:BastionHost"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:BastionHost"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:BastionHost"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:BastionHost"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:BastionHost"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:BastionHost"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:BastionHost"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:BastionHost"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:BastionHost"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:BastionHost"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:BastionHost"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:BastionHost"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:BastionHost"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:BastionHost"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:BastionHost"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:BastionHost"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:BastionHost"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:BastionHost"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:BastionHost"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220701:BastionHost"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220901:BastionHost"),
		},
		{
			Type: pulumi.String("azure-native:network/v20221101:BastionHost"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230201:BastionHost"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230501:BastionHost"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230601:BastionHost"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource BastionHost
	err := ctx.RegisterResource("azure-native:network/v20230401:BastionHost", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBastionHost gets an existing BastionHost resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBastionHost(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BastionHostState, opts ...pulumi.ResourceOption) (*BastionHost, error) {
	var resource BastionHost
	err := ctx.ReadResource("azure-native:network/v20230401:BastionHost", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BastionHost resources.
type bastionHostState struct {
}

type BastionHostState struct {
}

func (BastionHostState) ElementType() reflect.Type {
	return reflect.TypeOf((*bastionHostState)(nil)).Elem()
}

type bastionHostArgs struct {
	// The name of the Bastion Host.
	BastionHostName *string `pulumi:"bastionHostName"`
	// Enable/Disable Copy/Paste feature of the Bastion Host resource.
	DisableCopyPaste *bool `pulumi:"disableCopyPaste"`
	// FQDN for the endpoint on which bastion host is accessible.
	DnsName *string `pulumi:"dnsName"`
	// Enable/Disable File Copy feature of the Bastion Host resource.
	EnableFileCopy *bool `pulumi:"enableFileCopy"`
	// Enable/Disable IP Connect feature of the Bastion Host resource.
	EnableIpConnect *bool `pulumi:"enableIpConnect"`
	// Enable/Disable Kerberos feature of the Bastion Host resource.
	EnableKerberos *bool `pulumi:"enableKerberos"`
	// Enable/Disable Shareable Link of the Bastion Host resource.
	EnableShareableLink *bool `pulumi:"enableShareableLink"`
	// Enable/Disable Tunneling feature of the Bastion Host resource.
	EnableTunneling *bool `pulumi:"enableTunneling"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// IP configuration of the Bastion Host resource.
	IpConfigurations []BastionHostIPConfiguration `pulumi:"ipConfigurations"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The scale units for the Bastion Host resource.
	ScaleUnits *int `pulumi:"scaleUnits"`
	// The sku of this Bastion Host.
	Sku *Sku `pulumi:"sku"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a BastionHost resource.
type BastionHostArgs struct {
	// The name of the Bastion Host.
	BastionHostName pulumi.StringPtrInput
	// Enable/Disable Copy/Paste feature of the Bastion Host resource.
	DisableCopyPaste pulumi.BoolPtrInput
	// FQDN for the endpoint on which bastion host is accessible.
	DnsName pulumi.StringPtrInput
	// Enable/Disable File Copy feature of the Bastion Host resource.
	EnableFileCopy pulumi.BoolPtrInput
	// Enable/Disable IP Connect feature of the Bastion Host resource.
	EnableIpConnect pulumi.BoolPtrInput
	// Enable/Disable Kerberos feature of the Bastion Host resource.
	EnableKerberos pulumi.BoolPtrInput
	// Enable/Disable Shareable Link of the Bastion Host resource.
	EnableShareableLink pulumi.BoolPtrInput
	// Enable/Disable Tunneling feature of the Bastion Host resource.
	EnableTunneling pulumi.BoolPtrInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// IP configuration of the Bastion Host resource.
	IpConfigurations BastionHostIPConfigurationArrayInput
	// Resource location.
	Location pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The scale units for the Bastion Host resource.
	ScaleUnits pulumi.IntPtrInput
	// The sku of this Bastion Host.
	Sku SkuPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (BastionHostArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bastionHostArgs)(nil)).Elem()
}

type BastionHostInput interface {
	pulumi.Input

	ToBastionHostOutput() BastionHostOutput
	ToBastionHostOutputWithContext(ctx context.Context) BastionHostOutput
}

func (*BastionHost) ElementType() reflect.Type {
	return reflect.TypeOf((**BastionHost)(nil)).Elem()
}

func (i *BastionHost) ToBastionHostOutput() BastionHostOutput {
	return i.ToBastionHostOutputWithContext(context.Background())
}

func (i *BastionHost) ToBastionHostOutputWithContext(ctx context.Context) BastionHostOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BastionHostOutput)
}

type BastionHostOutput struct{ *pulumi.OutputState }

func (BastionHostOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BastionHost)(nil)).Elem()
}

func (o BastionHostOutput) ToBastionHostOutput() BastionHostOutput {
	return o
}

func (o BastionHostOutput) ToBastionHostOutputWithContext(ctx context.Context) BastionHostOutput {
	return o
}

// Enable/Disable Copy/Paste feature of the Bastion Host resource.
func (o BastionHostOutput) DisableCopyPaste() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BastionHost) pulumi.BoolPtrOutput { return v.DisableCopyPaste }).(pulumi.BoolPtrOutput)
}

// FQDN for the endpoint on which bastion host is accessible.
func (o BastionHostOutput) DnsName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BastionHost) pulumi.StringPtrOutput { return v.DnsName }).(pulumi.StringPtrOutput)
}

// Enable/Disable File Copy feature of the Bastion Host resource.
func (o BastionHostOutput) EnableFileCopy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BastionHost) pulumi.BoolPtrOutput { return v.EnableFileCopy }).(pulumi.BoolPtrOutput)
}

// Enable/Disable IP Connect feature of the Bastion Host resource.
func (o BastionHostOutput) EnableIpConnect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BastionHost) pulumi.BoolPtrOutput { return v.EnableIpConnect }).(pulumi.BoolPtrOutput)
}

// Enable/Disable Kerberos feature of the Bastion Host resource.
func (o BastionHostOutput) EnableKerberos() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BastionHost) pulumi.BoolPtrOutput { return v.EnableKerberos }).(pulumi.BoolPtrOutput)
}

// Enable/Disable Shareable Link of the Bastion Host resource.
func (o BastionHostOutput) EnableShareableLink() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BastionHost) pulumi.BoolPtrOutput { return v.EnableShareableLink }).(pulumi.BoolPtrOutput)
}

// Enable/Disable Tunneling feature of the Bastion Host resource.
func (o BastionHostOutput) EnableTunneling() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BastionHost) pulumi.BoolPtrOutput { return v.EnableTunneling }).(pulumi.BoolPtrOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o BastionHostOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *BastionHost) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// IP configuration of the Bastion Host resource.
func (o BastionHostOutput) IpConfigurations() BastionHostIPConfigurationResponseArrayOutput {
	return o.ApplyT(func(v *BastionHost) BastionHostIPConfigurationResponseArrayOutput { return v.IpConfigurations }).(BastionHostIPConfigurationResponseArrayOutput)
}

// Resource location.
func (o BastionHostOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BastionHost) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o BastionHostOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *BastionHost) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the bastion host resource.
func (o BastionHostOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *BastionHost) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The scale units for the Bastion Host resource.
func (o BastionHostOutput) ScaleUnits() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *BastionHost) pulumi.IntPtrOutput { return v.ScaleUnits }).(pulumi.IntPtrOutput)
}

// The sku of this Bastion Host.
func (o BastionHostOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *BastionHost) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

// Resource tags.
func (o BastionHostOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *BastionHost) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o BastionHostOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *BastionHost) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(BastionHostOutput{})
}
