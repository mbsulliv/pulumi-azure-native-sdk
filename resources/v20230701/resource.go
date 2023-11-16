// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230701

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource information.
type Resource struct {
	pulumi.CustomResourceState

	// Resource extended location.
	ExtendedLocation ExtendedLocationResponsePtrOutput `pulumi:"extendedLocation"`
	// The identity of the resource.
	Identity IdentityResponsePtrOutput `pulumi:"identity"`
	// The kind of the resource.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Resource location
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// ID of the resource that manages this resource.
	ManagedBy pulumi.StringPtrOutput `pulumi:"managedBy"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// The plan of the resource.
	Plan PlanResponsePtrOutput `pulumi:"plan"`
	// The resource properties.
	Properties pulumi.AnyOutput `pulumi:"properties"`
	// The SKU of the resource.
	Sku SkuResponsePtrOutput `pulumi:"sku"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewResource registers a new resource with the given unique name, arguments, and options.
func NewResource(ctx *pulumi.Context,
	name string, args *ResourceArgs, opts ...pulumi.ResourceOption) (*Resource, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ParentResourcePath == nil {
		return nil, errors.New("invalid value for required argument 'ParentResourcePath'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ResourceProviderNamespace == nil {
		return nil, errors.New("invalid value for required argument 'ResourceProviderNamespace'")
	}
	if args.ResourceType == nil {
		return nil, errors.New("invalid value for required argument 'ResourceType'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:resources:Resource"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20151101:Resource"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20160201:Resource"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20160701:Resource"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20160901:Resource"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20170510:Resource"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20180201:Resource"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20180501:Resource"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20190301:Resource"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20190501:Resource"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20190510:Resource"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20190701:Resource"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20190801:Resource"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20191001:Resource"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20200601:Resource"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20200801:Resource"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20201001:Resource"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20210101:Resource"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20210401:Resource"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20220901:Resource"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Resource
	err := ctx.RegisterResource("azure-native:resources/v20230701:Resource", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResource gets an existing Resource resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResource(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResourceState, opts ...pulumi.ResourceOption) (*Resource, error) {
	var resource Resource
	err := ctx.ReadResource("azure-native:resources/v20230701:Resource", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Resource resources.
type resourceState struct {
}

type ResourceState struct {
}

func (ResourceState) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceState)(nil)).Elem()
}

type resourceArgs struct {
	// Resource extended location.
	ExtendedLocation *ExtendedLocation `pulumi:"extendedLocation"`
	// The identity of the resource.
	Identity *Identity `pulumi:"identity"`
	// The kind of the resource.
	Kind *string `pulumi:"kind"`
	// Resource location
	Location *string `pulumi:"location"`
	// ID of the resource that manages this resource.
	ManagedBy *string `pulumi:"managedBy"`
	// The parent resource identity.
	ParentResourcePath string `pulumi:"parentResourcePath"`
	// The plan of the resource.
	Plan *Plan `pulumi:"plan"`
	// The resource properties.
	Properties interface{} `pulumi:"properties"`
	// The name of the resource group for the resource. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the resource to create.
	ResourceName *string `pulumi:"resourceName"`
	// The namespace of the resource provider.
	ResourceProviderNamespace string `pulumi:"resourceProviderNamespace"`
	// The resource type of the resource to create.
	ResourceType string `pulumi:"resourceType"`
	// The SKU of the resource.
	Sku *Sku `pulumi:"sku"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Resource resource.
type ResourceArgs struct {
	// Resource extended location.
	ExtendedLocation ExtendedLocationPtrInput
	// The identity of the resource.
	Identity IdentityPtrInput
	// The kind of the resource.
	Kind pulumi.StringPtrInput
	// Resource location
	Location pulumi.StringPtrInput
	// ID of the resource that manages this resource.
	ManagedBy pulumi.StringPtrInput
	// The parent resource identity.
	ParentResourcePath pulumi.StringInput
	// The plan of the resource.
	Plan PlanPtrInput
	// The resource properties.
	Properties pulumi.Input
	// The name of the resource group for the resource. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the resource to create.
	ResourceName pulumi.StringPtrInput
	// The namespace of the resource provider.
	ResourceProviderNamespace pulumi.StringInput
	// The resource type of the resource to create.
	ResourceType pulumi.StringInput
	// The SKU of the resource.
	Sku SkuPtrInput
	// Resource tags
	Tags pulumi.StringMapInput
}

func (ResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceArgs)(nil)).Elem()
}

type ResourceInput interface {
	pulumi.Input

	ToResourceOutput() ResourceOutput
	ToResourceOutputWithContext(ctx context.Context) ResourceOutput
}

func (*Resource) ElementType() reflect.Type {
	return reflect.TypeOf((**Resource)(nil)).Elem()
}

func (i *Resource) ToResourceOutput() ResourceOutput {
	return i.ToResourceOutputWithContext(context.Background())
}

func (i *Resource) ToResourceOutputWithContext(ctx context.Context) ResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceOutput)
}

type ResourceOutput struct{ *pulumi.OutputState }

func (ResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Resource)(nil)).Elem()
}

func (o ResourceOutput) ToResourceOutput() ResourceOutput {
	return o
}

func (o ResourceOutput) ToResourceOutputWithContext(ctx context.Context) ResourceOutput {
	return o
}

// Resource extended location.
func (o ResourceOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v *Resource) ExtendedLocationResponsePtrOutput { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

// The identity of the resource.
func (o ResourceOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v *Resource) IdentityResponsePtrOutput { return v.Identity }).(IdentityResponsePtrOutput)
}

// The kind of the resource.
func (o ResourceOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Resource) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Resource location
func (o ResourceOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Resource) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// ID of the resource that manages this resource.
func (o ResourceOutput) ManagedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Resource) pulumi.StringPtrOutput { return v.ManagedBy }).(pulumi.StringPtrOutput)
}

// Resource name
func (o ResourceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Resource) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The plan of the resource.
func (o ResourceOutput) Plan() PlanResponsePtrOutput {
	return o.ApplyT(func(v *Resource) PlanResponsePtrOutput { return v.Plan }).(PlanResponsePtrOutput)
}

// The resource properties.
func (o ResourceOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v *Resource) pulumi.AnyOutput { return v.Properties }).(pulumi.AnyOutput)
}

// The SKU of the resource.
func (o ResourceOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *Resource) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

// Resource tags
func (o ResourceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Resource) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type
func (o ResourceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Resource) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ResourceOutput{})
}
