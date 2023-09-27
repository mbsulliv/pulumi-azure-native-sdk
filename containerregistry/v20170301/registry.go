// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20170301

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// An object that represents a container registry.
type Registry struct {
	pulumi.CustomResourceState

	// The value that indicates whether the admin user is enabled.
	AdminUserEnabled pulumi.BoolPtrOutput `pulumi:"adminUserEnabled"`
	// The creation date of the container registry in ISO8601 format.
	CreationDate pulumi.StringOutput `pulumi:"creationDate"`
	// The location of the resource. This cannot be changed after the resource is created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The URL that can be used to log into the container registry.
	LoginServer pulumi.StringOutput `pulumi:"loginServer"`
	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioning state of the container registry at the time the operation was called.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The SKU of the container registry.
	Sku SkuResponseOutput `pulumi:"sku"`
	// The properties of the storage account for the container registry.
	StorageAccount StorageAccountPropertiesResponsePtrOutput `pulumi:"storageAccount"`
	// The tags of the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewRegistry registers a new resource with the given unique name, arguments, and options.
func NewRegistry(ctx *pulumi.Context,
	name string, args *RegistryArgs, opts ...pulumi.ResourceOption) (*Registry, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Sku == nil {
		return nil, errors.New("invalid value for required argument 'Sku'")
	}
	if args.StorageAccount == nil {
		return nil, errors.New("invalid value for required argument 'StorageAccount'")
	}
	if args.AdminUserEnabled == nil {
		args.AdminUserEnabled = pulumi.BoolPtr(false)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:containerregistry:Registry"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20171001:Registry"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20190501:Registry"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20191201preview:Registry"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20201101preview:Registry"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20210601preview:Registry"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20210801preview:Registry"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20210901:Registry"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20211201preview:Registry"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20220201preview:Registry"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20221201:Registry"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20230101preview:Registry"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20230601preview:Registry"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20230701:Registry"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20230801preview:Registry"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Registry
	err := ctx.RegisterResource("azure-native:containerregistry/v20170301:Registry", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRegistry gets an existing Registry resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegistry(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegistryState, opts ...pulumi.ResourceOption) (*Registry, error) {
	var resource Registry
	err := ctx.ReadResource("azure-native:containerregistry/v20170301:Registry", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Registry resources.
type registryState struct {
}

type RegistryState struct {
}

func (RegistryState) ElementType() reflect.Type {
	return reflect.TypeOf((*registryState)(nil)).Elem()
}

type registryArgs struct {
	// The value that indicates whether the admin user is enabled.
	AdminUserEnabled *bool `pulumi:"adminUserEnabled"`
	// The location of the container registry. This cannot be changed after the resource is created.
	Location *string `pulumi:"location"`
	// The name of the container registry.
	RegistryName *string `pulumi:"registryName"`
	// The name of the resource group to which the container registry belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The SKU of the container registry.
	Sku Sku `pulumi:"sku"`
	// The parameters of a storage account for the container registry. If specified, the storage account must be in the same physical location as the container registry.
	StorageAccount StorageAccountParameters `pulumi:"storageAccount"`
	// The tags for the container registry.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Registry resource.
type RegistryArgs struct {
	// The value that indicates whether the admin user is enabled.
	AdminUserEnabled pulumi.BoolPtrInput
	// The location of the container registry. This cannot be changed after the resource is created.
	Location pulumi.StringPtrInput
	// The name of the container registry.
	RegistryName pulumi.StringPtrInput
	// The name of the resource group to which the container registry belongs.
	ResourceGroupName pulumi.StringInput
	// The SKU of the container registry.
	Sku SkuInput
	// The parameters of a storage account for the container registry. If specified, the storage account must be in the same physical location as the container registry.
	StorageAccount StorageAccountParametersInput
	// The tags for the container registry.
	Tags pulumi.StringMapInput
}

func (RegistryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*registryArgs)(nil)).Elem()
}

type RegistryInput interface {
	pulumi.Input

	ToRegistryOutput() RegistryOutput
	ToRegistryOutputWithContext(ctx context.Context) RegistryOutput
}

func (*Registry) ElementType() reflect.Type {
	return reflect.TypeOf((**Registry)(nil)).Elem()
}

func (i *Registry) ToRegistryOutput() RegistryOutput {
	return i.ToRegistryOutputWithContext(context.Background())
}

func (i *Registry) ToRegistryOutputWithContext(ctx context.Context) RegistryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistryOutput)
}

func (i *Registry) ToOutput(ctx context.Context) pulumix.Output[*Registry] {
	return pulumix.Output[*Registry]{
		OutputState: i.ToRegistryOutputWithContext(ctx).OutputState,
	}
}

type RegistryOutput struct{ *pulumi.OutputState }

func (RegistryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Registry)(nil)).Elem()
}

func (o RegistryOutput) ToRegistryOutput() RegistryOutput {
	return o
}

func (o RegistryOutput) ToRegistryOutputWithContext(ctx context.Context) RegistryOutput {
	return o
}

func (o RegistryOutput) ToOutput(ctx context.Context) pulumix.Output[*Registry] {
	return pulumix.Output[*Registry]{
		OutputState: o.OutputState,
	}
}

// The value that indicates whether the admin user is enabled.
func (o RegistryOutput) AdminUserEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Registry) pulumi.BoolPtrOutput { return v.AdminUserEnabled }).(pulumi.BoolPtrOutput)
}

// The creation date of the container registry in ISO8601 format.
func (o RegistryOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v *Registry) pulumi.StringOutput { return v.CreationDate }).(pulumi.StringOutput)
}

// The location of the resource. This cannot be changed after the resource is created.
func (o RegistryOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Registry) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The URL that can be used to log into the container registry.
func (o RegistryOutput) LoginServer() pulumi.StringOutput {
	return o.ApplyT(func(v *Registry) pulumi.StringOutput { return v.LoginServer }).(pulumi.StringOutput)
}

// The name of the resource.
func (o RegistryOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Registry) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the container registry at the time the operation was called.
func (o RegistryOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Registry) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The SKU of the container registry.
func (o RegistryOutput) Sku() SkuResponseOutput {
	return o.ApplyT(func(v *Registry) SkuResponseOutput { return v.Sku }).(SkuResponseOutput)
}

// The properties of the storage account for the container registry.
func (o RegistryOutput) StorageAccount() StorageAccountPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *Registry) StorageAccountPropertiesResponsePtrOutput { return v.StorageAccount }).(StorageAccountPropertiesResponsePtrOutput)
}

// The tags of the resource.
func (o RegistryOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Registry) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource.
func (o RegistryOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Registry) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(RegistryOutput{})
}
