// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20221201

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An object that represents a container registry.
type Registry struct {
	pulumi.CustomResourceState

	// The value that indicates whether the admin user is enabled.
	AdminUserEnabled pulumi.BoolPtrOutput `pulumi:"adminUserEnabled"`
	// The creation date of the container registry in ISO8601 format.
	CreationDate pulumi.StringOutput `pulumi:"creationDate"`
	// Enable a single data endpoint per region for serving data.
	DataEndpointEnabled pulumi.BoolPtrOutput `pulumi:"dataEndpointEnabled"`
	// List of host names that will serve data when dataEndpointEnabled is true.
	DataEndpointHostNames pulumi.StringArrayOutput `pulumi:"dataEndpointHostNames"`
	// The encryption settings of container registry.
	Encryption EncryptionPropertyResponsePtrOutput `pulumi:"encryption"`
	// The identity of the container registry.
	Identity IdentityPropertiesResponsePtrOutput `pulumi:"identity"`
	// The location of the resource. This cannot be changed after the resource is created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The URL that can be used to log into the container registry.
	LoginServer pulumi.StringOutput `pulumi:"loginServer"`
	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Whether to allow trusted Azure services to access a network restricted registry.
	NetworkRuleBypassOptions pulumi.StringPtrOutput `pulumi:"networkRuleBypassOptions"`
	// The network rule set for a container registry.
	NetworkRuleSet NetworkRuleSetResponsePtrOutput `pulumi:"networkRuleSet"`
	// The policies for a container registry.
	Policies PoliciesResponsePtrOutput `pulumi:"policies"`
	// List of private endpoint connections for a container registry.
	PrivateEndpointConnections PrivateEndpointConnectionResponseArrayOutput `pulumi:"privateEndpointConnections"`
	// The provisioning state of the container registry at the time the operation was called.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Whether or not public network access is allowed for the container registry.
	PublicNetworkAccess pulumi.StringPtrOutput `pulumi:"publicNetworkAccess"`
	// The SKU of the container registry.
	Sku SkuResponseOutput `pulumi:"sku"`
	// The status of the container registry at the time the operation was called.
	Status StatusResponseOutput `pulumi:"status"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The tags of the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
	// Whether or not zone redundancy is enabled for this container registry
	ZoneRedundancy pulumi.StringPtrOutput `pulumi:"zoneRedundancy"`
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
	if args.AdminUserEnabled == nil {
		args.AdminUserEnabled = pulumi.BoolPtr(false)
	}
	if args.NetworkRuleBypassOptions == nil {
		args.NetworkRuleBypassOptions = pulumi.StringPtr("AzureServices")
	}
	if args.NetworkRuleSet != nil {
		args.NetworkRuleSet = args.NetworkRuleSet.ToNetworkRuleSetPtrOutput().ApplyT(func(v *NetworkRuleSet) *NetworkRuleSet { return v.Defaults() }).(NetworkRuleSetPtrOutput)
	}
	if args.Policies != nil {
		args.Policies = args.Policies.ToPoliciesPtrOutput().ApplyT(func(v *Policies) *Policies { return v.Defaults() }).(PoliciesPtrOutput)
	}
	if args.PublicNetworkAccess == nil {
		args.PublicNetworkAccess = pulumi.StringPtr("Enabled")
	}
	if args.ZoneRedundancy == nil {
		args.ZoneRedundancy = pulumi.StringPtr("Disabled")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:containerregistry:Registry"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20170301:Registry"),
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
		{
			Type: pulumi.String("azure-native:containerregistry/v20231101preview:Registry"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Registry
	err := ctx.RegisterResource("azure-native:containerregistry/v20221201:Registry", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure-native:containerregistry/v20221201:Registry", name, id, state, &resource, opts...)
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
	// Enable a single data endpoint per region for serving data.
	DataEndpointEnabled *bool `pulumi:"dataEndpointEnabled"`
	// The encryption settings of container registry.
	Encryption *EncryptionProperty `pulumi:"encryption"`
	// The identity of the container registry.
	Identity *IdentityProperties `pulumi:"identity"`
	// The location of the resource. This cannot be changed after the resource is created.
	Location *string `pulumi:"location"`
	// Whether to allow trusted Azure services to access a network restricted registry.
	NetworkRuleBypassOptions *string `pulumi:"networkRuleBypassOptions"`
	// The network rule set for a container registry.
	NetworkRuleSet *NetworkRuleSet `pulumi:"networkRuleSet"`
	// The policies for a container registry.
	Policies *Policies `pulumi:"policies"`
	// Whether or not public network access is allowed for the container registry.
	PublicNetworkAccess *string `pulumi:"publicNetworkAccess"`
	// The name of the container registry.
	RegistryName *string `pulumi:"registryName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The SKU of the container registry.
	Sku Sku `pulumi:"sku"`
	// The tags of the resource.
	Tags map[string]string `pulumi:"tags"`
	// Whether or not zone redundancy is enabled for this container registry
	ZoneRedundancy *string `pulumi:"zoneRedundancy"`
}

// The set of arguments for constructing a Registry resource.
type RegistryArgs struct {
	// The value that indicates whether the admin user is enabled.
	AdminUserEnabled pulumi.BoolPtrInput
	// Enable a single data endpoint per region for serving data.
	DataEndpointEnabled pulumi.BoolPtrInput
	// The encryption settings of container registry.
	Encryption EncryptionPropertyPtrInput
	// The identity of the container registry.
	Identity IdentityPropertiesPtrInput
	// The location of the resource. This cannot be changed after the resource is created.
	Location pulumi.StringPtrInput
	// Whether to allow trusted Azure services to access a network restricted registry.
	NetworkRuleBypassOptions pulumi.StringPtrInput
	// The network rule set for a container registry.
	NetworkRuleSet NetworkRuleSetPtrInput
	// The policies for a container registry.
	Policies PoliciesPtrInput
	// Whether or not public network access is allowed for the container registry.
	PublicNetworkAccess pulumi.StringPtrInput
	// The name of the container registry.
	RegistryName pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The SKU of the container registry.
	Sku SkuInput
	// The tags of the resource.
	Tags pulumi.StringMapInput
	// Whether or not zone redundancy is enabled for this container registry
	ZoneRedundancy pulumi.StringPtrInput
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

// The value that indicates whether the admin user is enabled.
func (o RegistryOutput) AdminUserEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Registry) pulumi.BoolPtrOutput { return v.AdminUserEnabled }).(pulumi.BoolPtrOutput)
}

// The creation date of the container registry in ISO8601 format.
func (o RegistryOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v *Registry) pulumi.StringOutput { return v.CreationDate }).(pulumi.StringOutput)
}

// Enable a single data endpoint per region for serving data.
func (o RegistryOutput) DataEndpointEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Registry) pulumi.BoolPtrOutput { return v.DataEndpointEnabled }).(pulumi.BoolPtrOutput)
}

// List of host names that will serve data when dataEndpointEnabled is true.
func (o RegistryOutput) DataEndpointHostNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Registry) pulumi.StringArrayOutput { return v.DataEndpointHostNames }).(pulumi.StringArrayOutput)
}

// The encryption settings of container registry.
func (o RegistryOutput) Encryption() EncryptionPropertyResponsePtrOutput {
	return o.ApplyT(func(v *Registry) EncryptionPropertyResponsePtrOutput { return v.Encryption }).(EncryptionPropertyResponsePtrOutput)
}

// The identity of the container registry.
func (o RegistryOutput) Identity() IdentityPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *Registry) IdentityPropertiesResponsePtrOutput { return v.Identity }).(IdentityPropertiesResponsePtrOutput)
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

// Whether to allow trusted Azure services to access a network restricted registry.
func (o RegistryOutput) NetworkRuleBypassOptions() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Registry) pulumi.StringPtrOutput { return v.NetworkRuleBypassOptions }).(pulumi.StringPtrOutput)
}

// The network rule set for a container registry.
func (o RegistryOutput) NetworkRuleSet() NetworkRuleSetResponsePtrOutput {
	return o.ApplyT(func(v *Registry) NetworkRuleSetResponsePtrOutput { return v.NetworkRuleSet }).(NetworkRuleSetResponsePtrOutput)
}

// The policies for a container registry.
func (o RegistryOutput) Policies() PoliciesResponsePtrOutput {
	return o.ApplyT(func(v *Registry) PoliciesResponsePtrOutput { return v.Policies }).(PoliciesResponsePtrOutput)
}

// List of private endpoint connections for a container registry.
func (o RegistryOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v *Registry) PrivateEndpointConnectionResponseArrayOutput { return v.PrivateEndpointConnections }).(PrivateEndpointConnectionResponseArrayOutput)
}

// The provisioning state of the container registry at the time the operation was called.
func (o RegistryOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Registry) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Whether or not public network access is allowed for the container registry.
func (o RegistryOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Registry) pulumi.StringPtrOutput { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

// The SKU of the container registry.
func (o RegistryOutput) Sku() SkuResponseOutput {
	return o.ApplyT(func(v *Registry) SkuResponseOutput { return v.Sku }).(SkuResponseOutput)
}

// The status of the container registry at the time the operation was called.
func (o RegistryOutput) Status() StatusResponseOutput {
	return o.ApplyT(func(v *Registry) StatusResponseOutput { return v.Status }).(StatusResponseOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o RegistryOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Registry) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The tags of the resource.
func (o RegistryOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Registry) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource.
func (o RegistryOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Registry) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Whether or not zone redundancy is enabled for this container registry
func (o RegistryOutput) ZoneRedundancy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Registry) pulumi.StringPtrOutput { return v.ZoneRedundancy }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(RegistryOutput{})
}
