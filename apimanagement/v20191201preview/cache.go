// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20191201preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Cache details.
type Cache struct {
	pulumi.CustomResourceState

	// Runtime connection string to cache
	ConnectionString pulumi.StringOutput `pulumi:"connectionString"`
	// Cache description
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Original uri of entity in external system cache points to
	ResourceId pulumi.StringPtrOutput `pulumi:"resourceId"`
	// Resource type for API Management resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewCache registers a new resource with the given unique name, arguments, and options.
func NewCache(ctx *pulumi.Context,
	name string, args *CacheArgs, opts ...pulumi.ResourceOption) (*Cache, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConnectionString == nil {
		return nil, errors.New("invalid value for required argument 'ConnectionString'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:apimanagement:Cache"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180601preview:Cache"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20190101:Cache"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201:Cache"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20200601preview:Cache"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20201201:Cache"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210101preview:Cache"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210401preview:Cache"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210801:Cache"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20211201preview:Cache"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220401preview:Cache"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220801:Cache"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220901preview:Cache"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20230301preview:Cache"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Cache
	err := ctx.RegisterResource("azure-native:apimanagement/v20191201preview:Cache", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCache gets an existing Cache resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCache(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CacheState, opts ...pulumi.ResourceOption) (*Cache, error) {
	var resource Cache
	err := ctx.ReadResource("azure-native:apimanagement/v20191201preview:Cache", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Cache resources.
type cacheState struct {
}

type CacheState struct {
}

func (CacheState) ElementType() reflect.Type {
	return reflect.TypeOf((*cacheState)(nil)).Elem()
}

type cacheArgs struct {
	// Identifier of the Cache entity. Cache identifier (should be either 'default' or valid Azure region identifier).
	CacheId *string `pulumi:"cacheId"`
	// Runtime connection string to cache
	ConnectionString string `pulumi:"connectionString"`
	// Cache description
	Description *string `pulumi:"description"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Original uri of entity in external system cache points to
	ResourceId *string `pulumi:"resourceId"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a Cache resource.
type CacheArgs struct {
	// Identifier of the Cache entity. Cache identifier (should be either 'default' or valid Azure region identifier).
	CacheId pulumi.StringPtrInput
	// Runtime connection string to cache
	ConnectionString pulumi.StringInput
	// Cache description
	Description pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Original uri of entity in external system cache points to
	ResourceId pulumi.StringPtrInput
	// The name of the API Management service.
	ServiceName pulumi.StringInput
}

func (CacheArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cacheArgs)(nil)).Elem()
}

type CacheInput interface {
	pulumi.Input

	ToCacheOutput() CacheOutput
	ToCacheOutputWithContext(ctx context.Context) CacheOutput
}

func (*Cache) ElementType() reflect.Type {
	return reflect.TypeOf((**Cache)(nil)).Elem()
}

func (i *Cache) ToCacheOutput() CacheOutput {
	return i.ToCacheOutputWithContext(context.Background())
}

func (i *Cache) ToCacheOutputWithContext(ctx context.Context) CacheOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheOutput)
}

func (i *Cache) ToOutput(ctx context.Context) pulumix.Output[*Cache] {
	return pulumix.Output[*Cache]{
		OutputState: i.ToCacheOutputWithContext(ctx).OutputState,
	}
}

type CacheOutput struct{ *pulumi.OutputState }

func (CacheOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Cache)(nil)).Elem()
}

func (o CacheOutput) ToCacheOutput() CacheOutput {
	return o
}

func (o CacheOutput) ToCacheOutputWithContext(ctx context.Context) CacheOutput {
	return o
}

func (o CacheOutput) ToOutput(ctx context.Context) pulumix.Output[*Cache] {
	return pulumix.Output[*Cache]{
		OutputState: o.OutputState,
	}
}

// Runtime connection string to cache
func (o CacheOutput) ConnectionString() pulumi.StringOutput {
	return o.ApplyT(func(v *Cache) pulumi.StringOutput { return v.ConnectionString }).(pulumi.StringOutput)
}

// Cache description
func (o CacheOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cache) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o CacheOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Cache) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Original uri of entity in external system cache points to
func (o CacheOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cache) pulumi.StringPtrOutput { return v.ResourceId }).(pulumi.StringPtrOutput)
}

// Resource type for API Management resource.
func (o CacheOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Cache) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(CacheOutput{})
}
