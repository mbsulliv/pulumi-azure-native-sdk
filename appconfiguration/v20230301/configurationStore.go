// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230301

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The configuration store along with all resource properties. The Configuration Store will have all information to begin utilizing it.
type ConfigurationStore struct {
	pulumi.CustomResourceState

	// The creation date of configuration store.
	CreationDate pulumi.StringOutput `pulumi:"creationDate"`
	// Disables all authentication methods other than AAD authentication.
	DisableLocalAuth pulumi.BoolPtrOutput `pulumi:"disableLocalAuth"`
	// Property specifying whether protection against purge is enabled for this configuration store.
	EnablePurgeProtection pulumi.BoolPtrOutput `pulumi:"enablePurgeProtection"`
	// The encryption settings of the configuration store.
	Encryption EncryptionPropertiesResponsePtrOutput `pulumi:"encryption"`
	// The DNS endpoint where the configuration store API will be available.
	Endpoint pulumi.StringOutput `pulumi:"endpoint"`
	// The managed identity information, if configured.
	Identity ResourceIdentityResponsePtrOutput `pulumi:"identity"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The list of private endpoint connections that are set up for this resource.
	PrivateEndpointConnections PrivateEndpointConnectionReferenceResponseArrayOutput `pulumi:"privateEndpointConnections"`
	// The provisioning state of the configuration store.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Control permission for data plane traffic coming from public networks while private endpoint is enabled.
	PublicNetworkAccess pulumi.StringPtrOutput `pulumi:"publicNetworkAccess"`
	// The sku of the configuration store.
	Sku SkuResponseOutput `pulumi:"sku"`
	// The amount of time in days that the configuration store will be retained when it is soft deleted.
	SoftDeleteRetentionInDays pulumi.IntPtrOutput `pulumi:"softDeleteRetentionInDays"`
	// Resource system metadata.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewConfigurationStore registers a new resource with the given unique name, arguments, and options.
func NewConfigurationStore(ctx *pulumi.Context,
	name string, args *ConfigurationStoreArgs, opts ...pulumi.ResourceOption) (*ConfigurationStore, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Sku == nil {
		return nil, errors.New("invalid value for required argument 'Sku'")
	}
	if args.DisableLocalAuth == nil {
		args.DisableLocalAuth = pulumi.BoolPtr(false)
	}
	if args.EnablePurgeProtection == nil {
		args.EnablePurgeProtection = pulumi.BoolPtr(false)
	}
	if args.SoftDeleteRetentionInDays == nil {
		args.SoftDeleteRetentionInDays = pulumi.IntPtr(7)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:appconfiguration:ConfigurationStore"),
		},
		{
			Type: pulumi.String("azure-native:appconfiguration/v20190201preview:ConfigurationStore"),
		},
		{
			Type: pulumi.String("azure-native:appconfiguration/v20191001:ConfigurationStore"),
		},
		{
			Type: pulumi.String("azure-native:appconfiguration/v20191101preview:ConfigurationStore"),
		},
		{
			Type: pulumi.String("azure-native:appconfiguration/v20200601:ConfigurationStore"),
		},
		{
			Type: pulumi.String("azure-native:appconfiguration/v20200701preview:ConfigurationStore"),
		},
		{
			Type: pulumi.String("azure-native:appconfiguration/v20210301preview:ConfigurationStore"),
		},
		{
			Type: pulumi.String("azure-native:appconfiguration/v20211001preview:ConfigurationStore"),
		},
		{
			Type: pulumi.String("azure-native:appconfiguration/v20220301preview:ConfigurationStore"),
		},
		{
			Type: pulumi.String("azure-native:appconfiguration/v20220501:ConfigurationStore"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ConfigurationStore
	err := ctx.RegisterResource("azure-native:appconfiguration/v20230301:ConfigurationStore", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConfigurationStore gets an existing ConfigurationStore resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConfigurationStore(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConfigurationStoreState, opts ...pulumi.ResourceOption) (*ConfigurationStore, error) {
	var resource ConfigurationStore
	err := ctx.ReadResource("azure-native:appconfiguration/v20230301:ConfigurationStore", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ConfigurationStore resources.
type configurationStoreState struct {
}

type ConfigurationStoreState struct {
}

func (ConfigurationStoreState) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationStoreState)(nil)).Elem()
}

type configurationStoreArgs struct {
	// The name of the configuration store.
	ConfigStoreName *string `pulumi:"configStoreName"`
	// Indicates whether the configuration store need to be recovered.
	CreateMode *CreateMode `pulumi:"createMode"`
	// Disables all authentication methods other than AAD authentication.
	DisableLocalAuth *bool `pulumi:"disableLocalAuth"`
	// Property specifying whether protection against purge is enabled for this configuration store.
	EnablePurgeProtection *bool `pulumi:"enablePurgeProtection"`
	// The encryption settings of the configuration store.
	Encryption *EncryptionProperties `pulumi:"encryption"`
	// The managed identity information, if configured.
	Identity *ResourceIdentity `pulumi:"identity"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// Control permission for data plane traffic coming from public networks while private endpoint is enabled.
	PublicNetworkAccess *string `pulumi:"publicNetworkAccess"`
	// The name of the resource group to which the container registry belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The sku of the configuration store.
	Sku Sku `pulumi:"sku"`
	// The amount of time in days that the configuration store will be retained when it is soft deleted.
	SoftDeleteRetentionInDays *int `pulumi:"softDeleteRetentionInDays"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a ConfigurationStore resource.
type ConfigurationStoreArgs struct {
	// The name of the configuration store.
	ConfigStoreName pulumi.StringPtrInput
	// Indicates whether the configuration store need to be recovered.
	CreateMode CreateModePtrInput
	// Disables all authentication methods other than AAD authentication.
	DisableLocalAuth pulumi.BoolPtrInput
	// Property specifying whether protection against purge is enabled for this configuration store.
	EnablePurgeProtection pulumi.BoolPtrInput
	// The encryption settings of the configuration store.
	Encryption EncryptionPropertiesPtrInput
	// The managed identity information, if configured.
	Identity ResourceIdentityPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// Control permission for data plane traffic coming from public networks while private endpoint is enabled.
	PublicNetworkAccess pulumi.StringPtrInput
	// The name of the resource group to which the container registry belongs.
	ResourceGroupName pulumi.StringInput
	// The sku of the configuration store.
	Sku SkuInput
	// The amount of time in days that the configuration store will be retained when it is soft deleted.
	SoftDeleteRetentionInDays pulumi.IntPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (ConfigurationStoreArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationStoreArgs)(nil)).Elem()
}

type ConfigurationStoreInput interface {
	pulumi.Input

	ToConfigurationStoreOutput() ConfigurationStoreOutput
	ToConfigurationStoreOutputWithContext(ctx context.Context) ConfigurationStoreOutput
}

func (*ConfigurationStore) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationStore)(nil)).Elem()
}

func (i *ConfigurationStore) ToConfigurationStoreOutput() ConfigurationStoreOutput {
	return i.ToConfigurationStoreOutputWithContext(context.Background())
}

func (i *ConfigurationStore) ToConfigurationStoreOutputWithContext(ctx context.Context) ConfigurationStoreOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationStoreOutput)
}

func (i *ConfigurationStore) ToOutput(ctx context.Context) pulumix.Output[*ConfigurationStore] {
	return pulumix.Output[*ConfigurationStore]{
		OutputState: i.ToConfigurationStoreOutputWithContext(ctx).OutputState,
	}
}

type ConfigurationStoreOutput struct{ *pulumi.OutputState }

func (ConfigurationStoreOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationStore)(nil)).Elem()
}

func (o ConfigurationStoreOutput) ToConfigurationStoreOutput() ConfigurationStoreOutput {
	return o
}

func (o ConfigurationStoreOutput) ToConfigurationStoreOutputWithContext(ctx context.Context) ConfigurationStoreOutput {
	return o
}

func (o ConfigurationStoreOutput) ToOutput(ctx context.Context) pulumix.Output[*ConfigurationStore] {
	return pulumix.Output[*ConfigurationStore]{
		OutputState: o.OutputState,
	}
}

// The creation date of configuration store.
func (o ConfigurationStoreOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigurationStore) pulumi.StringOutput { return v.CreationDate }).(pulumi.StringOutput)
}

// Disables all authentication methods other than AAD authentication.
func (o ConfigurationStoreOutput) DisableLocalAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ConfigurationStore) pulumi.BoolPtrOutput { return v.DisableLocalAuth }).(pulumi.BoolPtrOutput)
}

// Property specifying whether protection against purge is enabled for this configuration store.
func (o ConfigurationStoreOutput) EnablePurgeProtection() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ConfigurationStore) pulumi.BoolPtrOutput { return v.EnablePurgeProtection }).(pulumi.BoolPtrOutput)
}

// The encryption settings of the configuration store.
func (o ConfigurationStoreOutput) Encryption() EncryptionPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *ConfigurationStore) EncryptionPropertiesResponsePtrOutput { return v.Encryption }).(EncryptionPropertiesResponsePtrOutput)
}

// The DNS endpoint where the configuration store API will be available.
func (o ConfigurationStoreOutput) Endpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigurationStore) pulumi.StringOutput { return v.Endpoint }).(pulumi.StringOutput)
}

// The managed identity information, if configured.
func (o ConfigurationStoreOutput) Identity() ResourceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *ConfigurationStore) ResourceIdentityResponsePtrOutput { return v.Identity }).(ResourceIdentityResponsePtrOutput)
}

// The geo-location where the resource lives
func (o ConfigurationStoreOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigurationStore) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o ConfigurationStoreOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigurationStore) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The list of private endpoint connections that are set up for this resource.
func (o ConfigurationStoreOutput) PrivateEndpointConnections() PrivateEndpointConnectionReferenceResponseArrayOutput {
	return o.ApplyT(func(v *ConfigurationStore) PrivateEndpointConnectionReferenceResponseArrayOutput {
		return v.PrivateEndpointConnections
	}).(PrivateEndpointConnectionReferenceResponseArrayOutput)
}

// The provisioning state of the configuration store.
func (o ConfigurationStoreOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigurationStore) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Control permission for data plane traffic coming from public networks while private endpoint is enabled.
func (o ConfigurationStoreOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigurationStore) pulumi.StringPtrOutput { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

// The sku of the configuration store.
func (o ConfigurationStoreOutput) Sku() SkuResponseOutput {
	return o.ApplyT(func(v *ConfigurationStore) SkuResponseOutput { return v.Sku }).(SkuResponseOutput)
}

// The amount of time in days that the configuration store will be retained when it is soft deleted.
func (o ConfigurationStoreOutput) SoftDeleteRetentionInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ConfigurationStore) pulumi.IntPtrOutput { return v.SoftDeleteRetentionInDays }).(pulumi.IntPtrOutput)
}

// Resource system metadata.
func (o ConfigurationStoreOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ConfigurationStore) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o ConfigurationStoreOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ConfigurationStore) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ConfigurationStoreOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigurationStore) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ConfigurationStoreOutput{})
}
