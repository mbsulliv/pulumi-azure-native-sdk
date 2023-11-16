// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package hybridnetwork

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configuration group schema resource.
// Azure REST API version: 2023-09-01.
type ConfigurationGroupSchema struct {
	pulumi.CustomResourceState

	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Configuration group schema properties.
	Properties ConfigurationGroupSchemaPropertiesFormatResponseOutput `pulumi:"properties"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewConfigurationGroupSchema registers a new resource with the given unique name, arguments, and options.
func NewConfigurationGroupSchema(ctx *pulumi.Context,
	name string, args *ConfigurationGroupSchemaArgs, opts ...pulumi.ResourceOption) (*ConfigurationGroupSchema, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PublisherName == nil {
		return nil, errors.New("invalid value for required argument 'PublisherName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:hybridnetwork/v20230901:ConfigurationGroupSchema"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ConfigurationGroupSchema
	err := ctx.RegisterResource("azure-native:hybridnetwork:ConfigurationGroupSchema", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConfigurationGroupSchema gets an existing ConfigurationGroupSchema resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConfigurationGroupSchema(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConfigurationGroupSchemaState, opts ...pulumi.ResourceOption) (*ConfigurationGroupSchema, error) {
	var resource ConfigurationGroupSchema
	err := ctx.ReadResource("azure-native:hybridnetwork:ConfigurationGroupSchema", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ConfigurationGroupSchema resources.
type configurationGroupSchemaState struct {
}

type ConfigurationGroupSchemaState struct {
}

func (ConfigurationGroupSchemaState) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationGroupSchemaState)(nil)).Elem()
}

type configurationGroupSchemaArgs struct {
	// The name of the configuration group schema.
	ConfigurationGroupSchemaName *string `pulumi:"configurationGroupSchemaName"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// Configuration group schema properties.
	Properties *ConfigurationGroupSchemaPropertiesFormat `pulumi:"properties"`
	// The name of the publisher.
	PublisherName string `pulumi:"publisherName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a ConfigurationGroupSchema resource.
type ConfigurationGroupSchemaArgs struct {
	// The name of the configuration group schema.
	ConfigurationGroupSchemaName pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// Configuration group schema properties.
	Properties ConfigurationGroupSchemaPropertiesFormatPtrInput
	// The name of the publisher.
	PublisherName pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (ConfigurationGroupSchemaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationGroupSchemaArgs)(nil)).Elem()
}

type ConfigurationGroupSchemaInput interface {
	pulumi.Input

	ToConfigurationGroupSchemaOutput() ConfigurationGroupSchemaOutput
	ToConfigurationGroupSchemaOutputWithContext(ctx context.Context) ConfigurationGroupSchemaOutput
}

func (*ConfigurationGroupSchema) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationGroupSchema)(nil)).Elem()
}

func (i *ConfigurationGroupSchema) ToConfigurationGroupSchemaOutput() ConfigurationGroupSchemaOutput {
	return i.ToConfigurationGroupSchemaOutputWithContext(context.Background())
}

func (i *ConfigurationGroupSchema) ToConfigurationGroupSchemaOutputWithContext(ctx context.Context) ConfigurationGroupSchemaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationGroupSchemaOutput)
}

type ConfigurationGroupSchemaOutput struct{ *pulumi.OutputState }

func (ConfigurationGroupSchemaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationGroupSchema)(nil)).Elem()
}

func (o ConfigurationGroupSchemaOutput) ToConfigurationGroupSchemaOutput() ConfigurationGroupSchemaOutput {
	return o
}

func (o ConfigurationGroupSchemaOutput) ToConfigurationGroupSchemaOutputWithContext(ctx context.Context) ConfigurationGroupSchemaOutput {
	return o
}

// The geo-location where the resource lives
func (o ConfigurationGroupSchemaOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigurationGroupSchema) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o ConfigurationGroupSchemaOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigurationGroupSchema) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Configuration group schema properties.
func (o ConfigurationGroupSchemaOutput) Properties() ConfigurationGroupSchemaPropertiesFormatResponseOutput {
	return o.ApplyT(func(v *ConfigurationGroupSchema) ConfigurationGroupSchemaPropertiesFormatResponseOutput {
		return v.Properties
	}).(ConfigurationGroupSchemaPropertiesFormatResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o ConfigurationGroupSchemaOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ConfigurationGroupSchema) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o ConfigurationGroupSchemaOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ConfigurationGroupSchema) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ConfigurationGroupSchemaOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigurationGroupSchema) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ConfigurationGroupSchemaOutput{})
}
