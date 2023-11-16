// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package changeanalysis

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A profile object that contains change analysis configuration, such as notification settings, for this subscription
// Azure REST API version: 2020-04-01-preview. Prior API version in Azure Native 1.x: 2020-04-01-preview.
type ConfigurationProfile struct {
	pulumi.CustomResourceState

	// The identity block returned by ARM resource that supports managed identity.
	Identity ResourceIdentityResponsePtrOutput `pulumi:"identity"`
	// The location where the resource is to be deployed.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The properties of a configuration profile.
	Properties ConfigurationProfileResourcePropertiesResponseOutput `pulumi:"properties"`
	// Top level metadata https://github.com/Azure/azure-resource-manager-rpc/blob/master/v1.0/common-api-contracts.md#system-metadata-for-all-azure-resources
	SystemData SystemDataResponsePtrOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewConfigurationProfile registers a new resource with the given unique name, arguments, and options.
func NewConfigurationProfile(ctx *pulumi.Context,
	name string, args *ConfigurationProfileArgs, opts ...pulumi.ResourceOption) (*ConfigurationProfile, error) {
	if args == nil {
		args = &ConfigurationProfileArgs{}
	}

	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:changeanalysis/v20200401preview:ConfigurationProfile"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ConfigurationProfile
	err := ctx.RegisterResource("azure-native:changeanalysis:ConfigurationProfile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConfigurationProfile gets an existing ConfigurationProfile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConfigurationProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConfigurationProfileState, opts ...pulumi.ResourceOption) (*ConfigurationProfile, error) {
	var resource ConfigurationProfile
	err := ctx.ReadResource("azure-native:changeanalysis:ConfigurationProfile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ConfigurationProfile resources.
type configurationProfileState struct {
}

type ConfigurationProfileState struct {
}

func (ConfigurationProfileState) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationProfileState)(nil)).Elem()
}

type configurationProfileArgs struct {
	// The identity block returned by ARM resource that supports managed identity.
	Identity *ResourceIdentity `pulumi:"identity"`
	// The location where the resource is to be deployed.
	Location *string `pulumi:"location"`
	// The name of the configuration profile. The profile name should be set to 'default', all other names will be overwritten.
	ProfileName *string `pulumi:"profileName"`
	// The properties of a configuration profile.
	Properties *ConfigurationProfileResourceProperties `pulumi:"properties"`
}

// The set of arguments for constructing a ConfigurationProfile resource.
type ConfigurationProfileArgs struct {
	// The identity block returned by ARM resource that supports managed identity.
	Identity ResourceIdentityPtrInput
	// The location where the resource is to be deployed.
	Location pulumi.StringPtrInput
	// The name of the configuration profile. The profile name should be set to 'default', all other names will be overwritten.
	ProfileName pulumi.StringPtrInput
	// The properties of a configuration profile.
	Properties ConfigurationProfileResourcePropertiesPtrInput
}

func (ConfigurationProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationProfileArgs)(nil)).Elem()
}

type ConfigurationProfileInput interface {
	pulumi.Input

	ToConfigurationProfileOutput() ConfigurationProfileOutput
	ToConfigurationProfileOutputWithContext(ctx context.Context) ConfigurationProfileOutput
}

func (*ConfigurationProfile) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationProfile)(nil)).Elem()
}

func (i *ConfigurationProfile) ToConfigurationProfileOutput() ConfigurationProfileOutput {
	return i.ToConfigurationProfileOutputWithContext(context.Background())
}

func (i *ConfigurationProfile) ToConfigurationProfileOutputWithContext(ctx context.Context) ConfigurationProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationProfileOutput)
}

type ConfigurationProfileOutput struct{ *pulumi.OutputState }

func (ConfigurationProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationProfile)(nil)).Elem()
}

func (o ConfigurationProfileOutput) ToConfigurationProfileOutput() ConfigurationProfileOutput {
	return o
}

func (o ConfigurationProfileOutput) ToConfigurationProfileOutputWithContext(ctx context.Context) ConfigurationProfileOutput {
	return o
}

// The identity block returned by ARM resource that supports managed identity.
func (o ConfigurationProfileOutput) Identity() ResourceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *ConfigurationProfile) ResourceIdentityResponsePtrOutput { return v.Identity }).(ResourceIdentityResponsePtrOutput)
}

// The location where the resource is to be deployed.
func (o ConfigurationProfileOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigurationProfile) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o ConfigurationProfileOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigurationProfile) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The properties of a configuration profile.
func (o ConfigurationProfileOutput) Properties() ConfigurationProfileResourcePropertiesResponseOutput {
	return o.ApplyT(func(v *ConfigurationProfile) ConfigurationProfileResourcePropertiesResponseOutput {
		return v.Properties
	}).(ConfigurationProfileResourcePropertiesResponseOutput)
}

// Top level metadata https://github.com/Azure/azure-resource-manager-rpc/blob/master/v1.0/common-api-contracts.md#system-metadata-for-all-azure-resources
func (o ConfigurationProfileOutput) SystemData() SystemDataResponsePtrOutput {
	return o.ApplyT(func(v *ConfigurationProfile) SystemDataResponsePtrOutput { return v.SystemData }).(SystemDataResponsePtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ConfigurationProfileOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigurationProfile) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ConfigurationProfileOutput{})
}
