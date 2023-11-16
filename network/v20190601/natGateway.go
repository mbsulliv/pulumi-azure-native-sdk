// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20190601

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Nat Gateway resource.
type NatGateway struct {
	pulumi.CustomResourceState

	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// The idle timeout of the nat gateway.
	IdleTimeoutInMinutes pulumi.IntPtrOutput `pulumi:"idleTimeoutInMinutes"`
	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioning state of the NatGateway resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
	ProvisioningState pulumi.StringPtrOutput `pulumi:"provisioningState"`
	// An array of public ip addresses associated with the nat gateway resource.
	PublicIpAddresses SubResourceResponseArrayOutput `pulumi:"publicIpAddresses"`
	// An array of public ip prefixes associated with the nat gateway resource.
	PublicIpPrefixes SubResourceResponseArrayOutput `pulumi:"publicIpPrefixes"`
	// The resource GUID property of the nat gateway resource.
	ResourceGuid pulumi.StringPtrOutput `pulumi:"resourceGuid"`
	// The nat gateway SKU.
	Sku NatGatewaySkuResponsePtrOutput `pulumi:"sku"`
	// An array of references to the subnets using this nat gateway resource.
	Subnets SubResourceResponseArrayOutput `pulumi:"subnets"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// A list of availability zones denoting the zone in which Nat Gateway should be deployed.
	Zones pulumi.StringArrayOutput `pulumi:"zones"`
}

// NewNatGateway registers a new resource with the given unique name, arguments, and options.
func NewNatGateway(ctx *pulumi.Context,
	name string, args *NatGatewayArgs, opts ...pulumi.ResourceOption) (*NatGateway, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:NatGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190201:NatGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190401:NatGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:NatGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:NatGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:NatGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:NatGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:NatGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:NatGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:NatGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:NatGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:NatGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:NatGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:NatGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:NatGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:NatGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:NatGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:NatGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:NatGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:NatGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:NatGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220701:NatGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220901:NatGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20221101:NatGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230201:NatGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230401:NatGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230501:NatGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230601:NatGateway"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource NatGateway
	err := ctx.RegisterResource("azure-native:network/v20190601:NatGateway", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNatGateway gets an existing NatGateway resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNatGateway(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NatGatewayState, opts ...pulumi.ResourceOption) (*NatGateway, error) {
	var resource NatGateway
	err := ctx.ReadResource("azure-native:network/v20190601:NatGateway", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NatGateway resources.
type natGatewayState struct {
}

type NatGatewayState struct {
}

func (NatGatewayState) ElementType() reflect.Type {
	return reflect.TypeOf((*natGatewayState)(nil)).Elem()
}

type natGatewayArgs struct {
	// Resource ID.
	Id *string `pulumi:"id"`
	// The idle timeout of the nat gateway.
	IdleTimeoutInMinutes *int `pulumi:"idleTimeoutInMinutes"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The name of the nat gateway.
	NatGatewayName *string `pulumi:"natGatewayName"`
	// The provisioning state of the NatGateway resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
	ProvisioningState *string `pulumi:"provisioningState"`
	// An array of public ip addresses associated with the nat gateway resource.
	PublicIpAddresses []SubResource `pulumi:"publicIpAddresses"`
	// An array of public ip prefixes associated with the nat gateway resource.
	PublicIpPrefixes []SubResource `pulumi:"publicIpPrefixes"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The resource GUID property of the nat gateway resource.
	ResourceGuid *string `pulumi:"resourceGuid"`
	// The nat gateway SKU.
	Sku *NatGatewaySku `pulumi:"sku"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// A list of availability zones denoting the zone in which Nat Gateway should be deployed.
	Zones []string `pulumi:"zones"`
}

// The set of arguments for constructing a NatGateway resource.
type NatGatewayArgs struct {
	// Resource ID.
	Id pulumi.StringPtrInput
	// The idle timeout of the nat gateway.
	IdleTimeoutInMinutes pulumi.IntPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// The name of the nat gateway.
	NatGatewayName pulumi.StringPtrInput
	// The provisioning state of the NatGateway resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
	ProvisioningState pulumi.StringPtrInput
	// An array of public ip addresses associated with the nat gateway resource.
	PublicIpAddresses SubResourceArrayInput
	// An array of public ip prefixes associated with the nat gateway resource.
	PublicIpPrefixes SubResourceArrayInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The resource GUID property of the nat gateway resource.
	ResourceGuid pulumi.StringPtrInput
	// The nat gateway SKU.
	Sku NatGatewaySkuPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// A list of availability zones denoting the zone in which Nat Gateway should be deployed.
	Zones pulumi.StringArrayInput
}

func (NatGatewayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*natGatewayArgs)(nil)).Elem()
}

type NatGatewayInput interface {
	pulumi.Input

	ToNatGatewayOutput() NatGatewayOutput
	ToNatGatewayOutputWithContext(ctx context.Context) NatGatewayOutput
}

func (*NatGateway) ElementType() reflect.Type {
	return reflect.TypeOf((**NatGateway)(nil)).Elem()
}

func (i *NatGateway) ToNatGatewayOutput() NatGatewayOutput {
	return i.ToNatGatewayOutputWithContext(context.Background())
}

func (i *NatGateway) ToNatGatewayOutputWithContext(ctx context.Context) NatGatewayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NatGatewayOutput)
}

type NatGatewayOutput struct{ *pulumi.OutputState }

func (NatGatewayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NatGateway)(nil)).Elem()
}

func (o NatGatewayOutput) ToNatGatewayOutput() NatGatewayOutput {
	return o
}

func (o NatGatewayOutput) ToNatGatewayOutputWithContext(ctx context.Context) NatGatewayOutput {
	return o
}

// A unique read-only string that changes whenever the resource is updated.
func (o NatGatewayOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NatGateway) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// The idle timeout of the nat gateway.
func (o NatGatewayOutput) IdleTimeoutInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *NatGateway) pulumi.IntPtrOutput { return v.IdleTimeoutInMinutes }).(pulumi.IntPtrOutput)
}

// Resource location.
func (o NatGatewayOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NatGateway) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o NatGatewayOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NatGateway) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the NatGateway resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
func (o NatGatewayOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NatGateway) pulumi.StringPtrOutput { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

// An array of public ip addresses associated with the nat gateway resource.
func (o NatGatewayOutput) PublicIpAddresses() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v *NatGateway) SubResourceResponseArrayOutput { return v.PublicIpAddresses }).(SubResourceResponseArrayOutput)
}

// An array of public ip prefixes associated with the nat gateway resource.
func (o NatGatewayOutput) PublicIpPrefixes() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v *NatGateway) SubResourceResponseArrayOutput { return v.PublicIpPrefixes }).(SubResourceResponseArrayOutput)
}

// The resource GUID property of the nat gateway resource.
func (o NatGatewayOutput) ResourceGuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NatGateway) pulumi.StringPtrOutput { return v.ResourceGuid }).(pulumi.StringPtrOutput)
}

// The nat gateway SKU.
func (o NatGatewayOutput) Sku() NatGatewaySkuResponsePtrOutput {
	return o.ApplyT(func(v *NatGateway) NatGatewaySkuResponsePtrOutput { return v.Sku }).(NatGatewaySkuResponsePtrOutput)
}

// An array of references to the subnets using this nat gateway resource.
func (o NatGatewayOutput) Subnets() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v *NatGateway) SubResourceResponseArrayOutput { return v.Subnets }).(SubResourceResponseArrayOutput)
}

// Resource tags.
func (o NatGatewayOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NatGateway) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o NatGatewayOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *NatGateway) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// A list of availability zones denoting the zone in which Nat Gateway should be deployed.
func (o NatGatewayOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NatGateway) pulumi.StringArrayOutput { return v.Zones }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(NatGatewayOutput{})
}
