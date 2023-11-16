// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20200630preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of the configuration profile preference.
type ConfigurationProfilePreference struct {
	pulumi.CustomResourceState

	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties of the configuration profile preference.
	Properties ConfigurationProfilePreferencePropertiesResponseOutput `pulumi:"properties"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewConfigurationProfilePreference registers a new resource with the given unique name, arguments, and options.
func NewConfigurationProfilePreference(ctx *pulumi.Context,
	name string, args *ConfigurationProfilePreferenceArgs, opts ...pulumi.ResourceOption) (*ConfigurationProfilePreference, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:automanage:ConfigurationProfilePreference"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ConfigurationProfilePreference
	err := ctx.RegisterResource("azure-native:automanage/v20200630preview:ConfigurationProfilePreference", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConfigurationProfilePreference gets an existing ConfigurationProfilePreference resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConfigurationProfilePreference(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConfigurationProfilePreferenceState, opts ...pulumi.ResourceOption) (*ConfigurationProfilePreference, error) {
	var resource ConfigurationProfilePreference
	err := ctx.ReadResource("azure-native:automanage/v20200630preview:ConfigurationProfilePreference", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ConfigurationProfilePreference resources.
type configurationProfilePreferenceState struct {
}

type ConfigurationProfilePreferenceState struct {
}

func (ConfigurationProfilePreferenceState) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationProfilePreferenceState)(nil)).Elem()
}

type configurationProfilePreferenceArgs struct {
	// Name of the configuration profile preference.
	ConfigurationProfilePreferenceName *string `pulumi:"configurationProfilePreferenceName"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// Properties of the configuration profile preference.
	Properties *ConfigurationProfilePreferenceProperties `pulumi:"properties"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a ConfigurationProfilePreference resource.
type ConfigurationProfilePreferenceArgs struct {
	// Name of the configuration profile preference.
	ConfigurationProfilePreferenceName pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// Properties of the configuration profile preference.
	Properties ConfigurationProfilePreferencePropertiesPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (ConfigurationProfilePreferenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationProfilePreferenceArgs)(nil)).Elem()
}

type ConfigurationProfilePreferenceInput interface {
	pulumi.Input

	ToConfigurationProfilePreferenceOutput() ConfigurationProfilePreferenceOutput
	ToConfigurationProfilePreferenceOutputWithContext(ctx context.Context) ConfigurationProfilePreferenceOutput
}

func (*ConfigurationProfilePreference) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationProfilePreference)(nil)).Elem()
}

func (i *ConfigurationProfilePreference) ToConfigurationProfilePreferenceOutput() ConfigurationProfilePreferenceOutput {
	return i.ToConfigurationProfilePreferenceOutputWithContext(context.Background())
}

func (i *ConfigurationProfilePreference) ToConfigurationProfilePreferenceOutputWithContext(ctx context.Context) ConfigurationProfilePreferenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationProfilePreferenceOutput)
}

type ConfigurationProfilePreferenceOutput struct{ *pulumi.OutputState }

func (ConfigurationProfilePreferenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationProfilePreference)(nil)).Elem()
}

func (o ConfigurationProfilePreferenceOutput) ToConfigurationProfilePreferenceOutput() ConfigurationProfilePreferenceOutput {
	return o
}

func (o ConfigurationProfilePreferenceOutput) ToConfigurationProfilePreferenceOutputWithContext(ctx context.Context) ConfigurationProfilePreferenceOutput {
	return o
}

// The geo-location where the resource lives
func (o ConfigurationProfilePreferenceOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigurationProfilePreference) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o ConfigurationProfilePreferenceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigurationProfilePreference) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Properties of the configuration profile preference.
func (o ConfigurationProfilePreferenceOutput) Properties() ConfigurationProfilePreferencePropertiesResponseOutput {
	return o.ApplyT(func(v *ConfigurationProfilePreference) ConfigurationProfilePreferencePropertiesResponseOutput {
		return v.Properties
	}).(ConfigurationProfilePreferencePropertiesResponseOutput)
}

// Resource tags.
func (o ConfigurationProfilePreferenceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ConfigurationProfilePreference) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ConfigurationProfilePreferenceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigurationProfilePreference) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ConfigurationProfilePreferenceOutput{})
}
