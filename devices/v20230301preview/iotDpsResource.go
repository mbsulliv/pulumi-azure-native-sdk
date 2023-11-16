// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230301preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The description of the provisioning service.
type IotDpsResource struct {
	pulumi.CustomResourceState

	// The Etag field is *not* required. If it is provided in the response body, it must also be provided as a header per the normal ETag convention.
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// The managed identities for a provisioning service.
	Identity ManagedServiceIdentityResponsePtrOutput `pulumi:"identity"`
	// The resource location.
	Location pulumi.StringOutput `pulumi:"location"`
	// The resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Service specific properties for a provisioning service
	Properties IotDpsPropertiesDescriptionResponseOutput `pulumi:"properties"`
	// The resource group of the resource.
	Resourcegroup pulumi.StringPtrOutput `pulumi:"resourcegroup"`
	// Sku info for a provisioning Service.
	Sku IotDpsSkuInfoResponseOutput `pulumi:"sku"`
	// The subscription id of the resource.
	Subscriptionid pulumi.StringPtrOutput `pulumi:"subscriptionid"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewIotDpsResource registers a new resource with the given unique name, arguments, and options.
func NewIotDpsResource(ctx *pulumi.Context,
	name string, args *IotDpsResourceArgs, opts ...pulumi.ResourceOption) (*IotDpsResource, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Sku == nil {
		return nil, errors.New("invalid value for required argument 'Sku'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:devices:IotDpsResource"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20170821preview:IotDpsResource"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20171115:IotDpsResource"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20180122:IotDpsResource"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20200101:IotDpsResource"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20200301:IotDpsResource"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20200901preview:IotDpsResource"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20211015:IotDpsResource"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20220205:IotDpsResource"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20221212:IotDpsResource"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource IotDpsResource
	err := ctx.RegisterResource("azure-native:devices/v20230301preview:IotDpsResource", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIotDpsResource gets an existing IotDpsResource resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIotDpsResource(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IotDpsResourceState, opts ...pulumi.ResourceOption) (*IotDpsResource, error) {
	var resource IotDpsResource
	err := ctx.ReadResource("azure-native:devices/v20230301preview:IotDpsResource", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IotDpsResource resources.
type iotDpsResourceState struct {
}

type IotDpsResourceState struct {
}

func (IotDpsResourceState) ElementType() reflect.Type {
	return reflect.TypeOf((*iotDpsResourceState)(nil)).Elem()
}

type iotDpsResourceArgs struct {
	// The managed identities for a provisioning service.
	Identity *ManagedServiceIdentity `pulumi:"identity"`
	// The resource location.
	Location *string `pulumi:"location"`
	// Service specific properties for a provisioning service
	Properties IotDpsPropertiesDescription `pulumi:"properties"`
	// Name of provisioning service to create or update.
	ProvisioningServiceName *string `pulumi:"provisioningServiceName"`
	// Resource group identifier.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The resource group of the resource.
	Resourcegroup *string `pulumi:"resourcegroup"`
	// Sku info for a provisioning Service.
	Sku IotDpsSkuInfo `pulumi:"sku"`
	// The subscription id of the resource.
	Subscriptionid *string `pulumi:"subscriptionid"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a IotDpsResource resource.
type IotDpsResourceArgs struct {
	// The managed identities for a provisioning service.
	Identity ManagedServiceIdentityPtrInput
	// The resource location.
	Location pulumi.StringPtrInput
	// Service specific properties for a provisioning service
	Properties IotDpsPropertiesDescriptionInput
	// Name of provisioning service to create or update.
	ProvisioningServiceName pulumi.StringPtrInput
	// Resource group identifier.
	ResourceGroupName pulumi.StringInput
	// The resource group of the resource.
	Resourcegroup pulumi.StringPtrInput
	// Sku info for a provisioning Service.
	Sku IotDpsSkuInfoInput
	// The subscription id of the resource.
	Subscriptionid pulumi.StringPtrInput
	// The resource tags.
	Tags pulumi.StringMapInput
}

func (IotDpsResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iotDpsResourceArgs)(nil)).Elem()
}

type IotDpsResourceInput interface {
	pulumi.Input

	ToIotDpsResourceOutput() IotDpsResourceOutput
	ToIotDpsResourceOutputWithContext(ctx context.Context) IotDpsResourceOutput
}

func (*IotDpsResource) ElementType() reflect.Type {
	return reflect.TypeOf((**IotDpsResource)(nil)).Elem()
}

func (i *IotDpsResource) ToIotDpsResourceOutput() IotDpsResourceOutput {
	return i.ToIotDpsResourceOutputWithContext(context.Background())
}

func (i *IotDpsResource) ToIotDpsResourceOutputWithContext(ctx context.Context) IotDpsResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IotDpsResourceOutput)
}

type IotDpsResourceOutput struct{ *pulumi.OutputState }

func (IotDpsResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IotDpsResource)(nil)).Elem()
}

func (o IotDpsResourceOutput) ToIotDpsResourceOutput() IotDpsResourceOutput {
	return o
}

func (o IotDpsResourceOutput) ToIotDpsResourceOutputWithContext(ctx context.Context) IotDpsResourceOutput {
	return o
}

// The Etag field is *not* required. If it is provided in the response body, it must also be provided as a header per the normal ETag convention.
func (o IotDpsResourceOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IotDpsResource) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// The managed identities for a provisioning service.
func (o IotDpsResourceOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *IotDpsResource) ManagedServiceIdentityResponsePtrOutput { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

// The resource location.
func (o IotDpsResourceOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *IotDpsResource) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The resource name.
func (o IotDpsResourceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *IotDpsResource) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Service specific properties for a provisioning service
func (o IotDpsResourceOutput) Properties() IotDpsPropertiesDescriptionResponseOutput {
	return o.ApplyT(func(v *IotDpsResource) IotDpsPropertiesDescriptionResponseOutput { return v.Properties }).(IotDpsPropertiesDescriptionResponseOutput)
}

// The resource group of the resource.
func (o IotDpsResourceOutput) Resourcegroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IotDpsResource) pulumi.StringPtrOutput { return v.Resourcegroup }).(pulumi.StringPtrOutput)
}

// Sku info for a provisioning Service.
func (o IotDpsResourceOutput) Sku() IotDpsSkuInfoResponseOutput {
	return o.ApplyT(func(v *IotDpsResource) IotDpsSkuInfoResponseOutput { return v.Sku }).(IotDpsSkuInfoResponseOutput)
}

// The subscription id of the resource.
func (o IotDpsResourceOutput) Subscriptionid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IotDpsResource) pulumi.StringPtrOutput { return v.Subscriptionid }).(pulumi.StringPtrOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o IotDpsResourceOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *IotDpsResource) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The resource tags.
func (o IotDpsResourceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *IotDpsResource) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The resource type.
func (o IotDpsResourceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *IotDpsResource) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(IotDpsResourceOutput{})
}
