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

// Represents a Configuration.
type Configuration struct {
	pulumi.CustomResourceState

	// Allowed values of the configuration.
	AllowedValues pulumi.StringOutput `pulumi:"allowedValues"`
	// Data type of the configuration.
	DataType pulumi.StringOutput `pulumi:"dataType"`
	// Default value of the configuration.
	DefaultValue pulumi.StringOutput `pulumi:"defaultValue"`
	// Description of the configuration.
	Description pulumi.StringOutput `pulumi:"description"`
	// Configuration documentation link.
	DocumentationLink pulumi.StringOutput `pulumi:"documentationLink"`
	// Configuration is pending restart or not.
	IsConfigPendingRestart pulumi.BoolOutput `pulumi:"isConfigPendingRestart"`
	// Configuration dynamic or static.
	IsDynamicConfig pulumi.BoolOutput `pulumi:"isDynamicConfig"`
	// Configuration read-only or not.
	IsReadOnly pulumi.BoolOutput `pulumi:"isReadOnly"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Source of the configuration.
	Source pulumi.StringPtrOutput `pulumi:"source"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// Configuration unit.
	Unit pulumi.StringOutput `pulumi:"unit"`
	// Value of the configuration.
	Value pulumi.StringPtrOutput `pulumi:"value"`
}

// NewConfiguration registers a new resource with the given unique name, arguments, and options.
func NewConfiguration(ctx *pulumi.Context,
	name string, args *ConfigurationArgs, opts ...pulumi.ResourceOption) (*Configuration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServerName == nil {
		return nil, errors.New("invalid value for required argument 'ServerName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:dbforpostgresql:Configuration"),
		},
		{
			Type: pulumi.String("azure-native:dbforpostgresql/v20200214preview:Configuration"),
		},
		{
			Type: pulumi.String("azure-native:dbforpostgresql/v20200214privatepreview:Configuration"),
		},
		{
			Type: pulumi.String("azure-native:dbforpostgresql/v20210410privatepreview:Configuration"),
		},
		{
			Type: pulumi.String("azure-native:dbforpostgresql/v20210601:Configuration"),
		},
		{
			Type: pulumi.String("azure-native:dbforpostgresql/v20210601preview:Configuration"),
		},
		{
			Type: pulumi.String("azure-native:dbforpostgresql/v20210615privatepreview:Configuration"),
		},
		{
			Type: pulumi.String("azure-native:dbforpostgresql/v20220120preview:Configuration"),
		},
		{
			Type: pulumi.String("azure-native:dbforpostgresql/v20220308preview:Configuration"),
		},
		{
			Type: pulumi.String("azure-native:dbforpostgresql/v20221201:Configuration"),
		},
		{
			Type: pulumi.String("azure-native:dbforpostgresql/v20230601preview:Configuration"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Configuration
	err := ctx.RegisterResource("azure-native:dbforpostgresql/v20230301preview:Configuration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConfiguration gets an existing Configuration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConfigurationState, opts ...pulumi.ResourceOption) (*Configuration, error) {
	var resource Configuration
	err := ctx.ReadResource("azure-native:dbforpostgresql/v20230301preview:Configuration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Configuration resources.
type configurationState struct {
}

type ConfigurationState struct {
}

func (ConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationState)(nil)).Elem()
}

type configurationArgs struct {
	// The name of the server configuration.
	ConfigurationName *string `pulumi:"configurationName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
	// Source of the configuration.
	Source *string `pulumi:"source"`
	// Value of the configuration.
	Value *string `pulumi:"value"`
}

// The set of arguments for constructing a Configuration resource.
type ConfigurationArgs struct {
	// The name of the server configuration.
	ConfigurationName pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the server.
	ServerName pulumi.StringInput
	// Source of the configuration.
	Source pulumi.StringPtrInput
	// Value of the configuration.
	Value pulumi.StringPtrInput
}

func (ConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationArgs)(nil)).Elem()
}

type ConfigurationInput interface {
	pulumi.Input

	ToConfigurationOutput() ConfigurationOutput
	ToConfigurationOutputWithContext(ctx context.Context) ConfigurationOutput
}

func (*Configuration) ElementType() reflect.Type {
	return reflect.TypeOf((**Configuration)(nil)).Elem()
}

func (i *Configuration) ToConfigurationOutput() ConfigurationOutput {
	return i.ToConfigurationOutputWithContext(context.Background())
}

func (i *Configuration) ToConfigurationOutputWithContext(ctx context.Context) ConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationOutput)
}

type ConfigurationOutput struct{ *pulumi.OutputState }

func (ConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Configuration)(nil)).Elem()
}

func (o ConfigurationOutput) ToConfigurationOutput() ConfigurationOutput {
	return o
}

func (o ConfigurationOutput) ToConfigurationOutputWithContext(ctx context.Context) ConfigurationOutput {
	return o
}

// Allowed values of the configuration.
func (o ConfigurationOutput) AllowedValues() pulumi.StringOutput {
	return o.ApplyT(func(v *Configuration) pulumi.StringOutput { return v.AllowedValues }).(pulumi.StringOutput)
}

// Data type of the configuration.
func (o ConfigurationOutput) DataType() pulumi.StringOutput {
	return o.ApplyT(func(v *Configuration) pulumi.StringOutput { return v.DataType }).(pulumi.StringOutput)
}

// Default value of the configuration.
func (o ConfigurationOutput) DefaultValue() pulumi.StringOutput {
	return o.ApplyT(func(v *Configuration) pulumi.StringOutput { return v.DefaultValue }).(pulumi.StringOutput)
}

// Description of the configuration.
func (o ConfigurationOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *Configuration) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Configuration documentation link.
func (o ConfigurationOutput) DocumentationLink() pulumi.StringOutput {
	return o.ApplyT(func(v *Configuration) pulumi.StringOutput { return v.DocumentationLink }).(pulumi.StringOutput)
}

// Configuration is pending restart or not.
func (o ConfigurationOutput) IsConfigPendingRestart() pulumi.BoolOutput {
	return o.ApplyT(func(v *Configuration) pulumi.BoolOutput { return v.IsConfigPendingRestart }).(pulumi.BoolOutput)
}

// Configuration dynamic or static.
func (o ConfigurationOutput) IsDynamicConfig() pulumi.BoolOutput {
	return o.ApplyT(func(v *Configuration) pulumi.BoolOutput { return v.IsDynamicConfig }).(pulumi.BoolOutput)
}

// Configuration read-only or not.
func (o ConfigurationOutput) IsReadOnly() pulumi.BoolOutput {
	return o.ApplyT(func(v *Configuration) pulumi.BoolOutput { return v.IsReadOnly }).(pulumi.BoolOutput)
}

// The name of the resource
func (o ConfigurationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Configuration) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Source of the configuration.
func (o ConfigurationOutput) Source() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Configuration) pulumi.StringPtrOutput { return v.Source }).(pulumi.StringPtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o ConfigurationOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Configuration) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ConfigurationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Configuration) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Configuration unit.
func (o ConfigurationOutput) Unit() pulumi.StringOutput {
	return o.ApplyT(func(v *Configuration) pulumi.StringOutput { return v.Unit }).(pulumi.StringOutput)
}

// Value of the configuration.
func (o ConfigurationOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Configuration) pulumi.StringPtrOutput { return v.Value }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ConfigurationOutput{})
}
