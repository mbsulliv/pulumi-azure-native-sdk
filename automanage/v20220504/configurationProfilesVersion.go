// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220504

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of the configuration profile.
type ConfigurationProfilesVersion struct {
	pulumi.CustomResourceState

	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties of the configuration profile.
	Properties ConfigurationProfilePropertiesResponseOutput `pulumi:"properties"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewConfigurationProfilesVersion registers a new resource with the given unique name, arguments, and options.
func NewConfigurationProfilesVersion(ctx *pulumi.Context,
	name string, args *ConfigurationProfilesVersionArgs, opts ...pulumi.ResourceOption) (*ConfigurationProfilesVersion, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConfigurationProfileName == nil {
		return nil, errors.New("invalid value for required argument 'ConfigurationProfileName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:automanage:ConfigurationProfilesVersion"),
		},
		{
			Type: pulumi.String("azure-native:automanage/v20210430preview:ConfigurationProfilesVersion"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ConfigurationProfilesVersion
	err := ctx.RegisterResource("azure-native:automanage/v20220504:ConfigurationProfilesVersion", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConfigurationProfilesVersion gets an existing ConfigurationProfilesVersion resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConfigurationProfilesVersion(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConfigurationProfilesVersionState, opts ...pulumi.ResourceOption) (*ConfigurationProfilesVersion, error) {
	var resource ConfigurationProfilesVersion
	err := ctx.ReadResource("azure-native:automanage/v20220504:ConfigurationProfilesVersion", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ConfigurationProfilesVersion resources.
type configurationProfilesVersionState struct {
}

type ConfigurationProfilesVersionState struct {
}

func (ConfigurationProfilesVersionState) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationProfilesVersionState)(nil)).Elem()
}

type configurationProfilesVersionArgs struct {
	// Name of the configuration profile.
	ConfigurationProfileName string `pulumi:"configurationProfileName"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// Properties of the configuration profile.
	Properties *ConfigurationProfileProperties `pulumi:"properties"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The configuration profile version name.
	VersionName *string `pulumi:"versionName"`
}

// The set of arguments for constructing a ConfigurationProfilesVersion resource.
type ConfigurationProfilesVersionArgs struct {
	// Name of the configuration profile.
	ConfigurationProfileName pulumi.StringInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// Properties of the configuration profile.
	Properties ConfigurationProfilePropertiesPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The configuration profile version name.
	VersionName pulumi.StringPtrInput
}

func (ConfigurationProfilesVersionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationProfilesVersionArgs)(nil)).Elem()
}

type ConfigurationProfilesVersionInput interface {
	pulumi.Input

	ToConfigurationProfilesVersionOutput() ConfigurationProfilesVersionOutput
	ToConfigurationProfilesVersionOutputWithContext(ctx context.Context) ConfigurationProfilesVersionOutput
}

func (*ConfigurationProfilesVersion) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationProfilesVersion)(nil)).Elem()
}

func (i *ConfigurationProfilesVersion) ToConfigurationProfilesVersionOutput() ConfigurationProfilesVersionOutput {
	return i.ToConfigurationProfilesVersionOutputWithContext(context.Background())
}

func (i *ConfigurationProfilesVersion) ToConfigurationProfilesVersionOutputWithContext(ctx context.Context) ConfigurationProfilesVersionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationProfilesVersionOutput)
}

type ConfigurationProfilesVersionOutput struct{ *pulumi.OutputState }

func (ConfigurationProfilesVersionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationProfilesVersion)(nil)).Elem()
}

func (o ConfigurationProfilesVersionOutput) ToConfigurationProfilesVersionOutput() ConfigurationProfilesVersionOutput {
	return o
}

func (o ConfigurationProfilesVersionOutput) ToConfigurationProfilesVersionOutputWithContext(ctx context.Context) ConfigurationProfilesVersionOutput {
	return o
}

// The geo-location where the resource lives
func (o ConfigurationProfilesVersionOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigurationProfilesVersion) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o ConfigurationProfilesVersionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigurationProfilesVersion) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Properties of the configuration profile.
func (o ConfigurationProfilesVersionOutput) Properties() ConfigurationProfilePropertiesResponseOutput {
	return o.ApplyT(func(v *ConfigurationProfilesVersion) ConfigurationProfilePropertiesResponseOutput {
		return v.Properties
	}).(ConfigurationProfilePropertiesResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o ConfigurationProfilesVersionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ConfigurationProfilesVersion) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o ConfigurationProfilesVersionOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ConfigurationProfilesVersion) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ConfigurationProfilesVersionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigurationProfilesVersion) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ConfigurationProfilesVersionOutput{})
}
