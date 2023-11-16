// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230101preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Single item in List or Get Schema Group operation
type SchemaRegistry struct {
	pulumi.CustomResourceState

	// Exact time the Schema Group was created.
	CreatedAtUtc pulumi.StringOutput `pulumi:"createdAtUtc"`
	// The ETag value.
	ETag pulumi.StringOutput `pulumi:"eTag"`
	// dictionary object for SchemaGroup group properties
	GroupProperties pulumi.StringMapOutput `pulumi:"groupProperties"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name                pulumi.StringOutput    `pulumi:"name"`
	SchemaCompatibility pulumi.StringPtrOutput `pulumi:"schemaCompatibility"`
	SchemaType          pulumi.StringPtrOutput `pulumi:"schemaType"`
	// The system meta data relating to this resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.EventHub/Namespaces" or "Microsoft.EventHub/Namespaces/EventHubs"
	Type pulumi.StringOutput `pulumi:"type"`
	// Exact time the Schema Group was updated
	UpdatedAtUtc pulumi.StringOutput `pulumi:"updatedAtUtc"`
}

// NewSchemaRegistry registers a new resource with the given unique name, arguments, and options.
func NewSchemaRegistry(ctx *pulumi.Context,
	name string, args *SchemaRegistryArgs, opts ...pulumi.ResourceOption) (*SchemaRegistry, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NamespaceName == nil {
		return nil, errors.New("invalid value for required argument 'NamespaceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:eventhub:SchemaRegistry"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20211101:SchemaRegistry"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20220101preview:SchemaRegistry"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20221001preview:SchemaRegistry"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource SchemaRegistry
	err := ctx.RegisterResource("azure-native:eventhub/v20230101preview:SchemaRegistry", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSchemaRegistry gets an existing SchemaRegistry resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSchemaRegistry(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SchemaRegistryState, opts ...pulumi.ResourceOption) (*SchemaRegistry, error) {
	var resource SchemaRegistry
	err := ctx.ReadResource("azure-native:eventhub/v20230101preview:SchemaRegistry", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SchemaRegistry resources.
type schemaRegistryState struct {
}

type SchemaRegistryState struct {
}

func (SchemaRegistryState) ElementType() reflect.Type {
	return reflect.TypeOf((*schemaRegistryState)(nil)).Elem()
}

type schemaRegistryArgs struct {
	// dictionary object for SchemaGroup group properties
	GroupProperties map[string]string `pulumi:"groupProperties"`
	// The Namespace name
	NamespaceName string `pulumi:"namespaceName"`
	// Name of the resource group within the azure subscription.
	ResourceGroupName   string  `pulumi:"resourceGroupName"`
	SchemaCompatibility *string `pulumi:"schemaCompatibility"`
	// The Schema Group name
	SchemaGroupName *string `pulumi:"schemaGroupName"`
	SchemaType      *string `pulumi:"schemaType"`
}

// The set of arguments for constructing a SchemaRegistry resource.
type SchemaRegistryArgs struct {
	// dictionary object for SchemaGroup group properties
	GroupProperties pulumi.StringMapInput
	// The Namespace name
	NamespaceName pulumi.StringInput
	// Name of the resource group within the azure subscription.
	ResourceGroupName   pulumi.StringInput
	SchemaCompatibility pulumi.StringPtrInput
	// The Schema Group name
	SchemaGroupName pulumi.StringPtrInput
	SchemaType      pulumi.StringPtrInput
}

func (SchemaRegistryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*schemaRegistryArgs)(nil)).Elem()
}

type SchemaRegistryInput interface {
	pulumi.Input

	ToSchemaRegistryOutput() SchemaRegistryOutput
	ToSchemaRegistryOutputWithContext(ctx context.Context) SchemaRegistryOutput
}

func (*SchemaRegistry) ElementType() reflect.Type {
	return reflect.TypeOf((**SchemaRegistry)(nil)).Elem()
}

func (i *SchemaRegistry) ToSchemaRegistryOutput() SchemaRegistryOutput {
	return i.ToSchemaRegistryOutputWithContext(context.Background())
}

func (i *SchemaRegistry) ToSchemaRegistryOutputWithContext(ctx context.Context) SchemaRegistryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SchemaRegistryOutput)
}

type SchemaRegistryOutput struct{ *pulumi.OutputState }

func (SchemaRegistryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SchemaRegistry)(nil)).Elem()
}

func (o SchemaRegistryOutput) ToSchemaRegistryOutput() SchemaRegistryOutput {
	return o
}

func (o SchemaRegistryOutput) ToSchemaRegistryOutputWithContext(ctx context.Context) SchemaRegistryOutput {
	return o
}

// Exact time the Schema Group was created.
func (o SchemaRegistryOutput) CreatedAtUtc() pulumi.StringOutput {
	return o.ApplyT(func(v *SchemaRegistry) pulumi.StringOutput { return v.CreatedAtUtc }).(pulumi.StringOutput)
}

// The ETag value.
func (o SchemaRegistryOutput) ETag() pulumi.StringOutput {
	return o.ApplyT(func(v *SchemaRegistry) pulumi.StringOutput { return v.ETag }).(pulumi.StringOutput)
}

// dictionary object for SchemaGroup group properties
func (o SchemaRegistryOutput) GroupProperties() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SchemaRegistry) pulumi.StringMapOutput { return v.GroupProperties }).(pulumi.StringMapOutput)
}

// The geo-location where the resource lives
func (o SchemaRegistryOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *SchemaRegistry) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o SchemaRegistryOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SchemaRegistry) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SchemaRegistryOutput) SchemaCompatibility() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SchemaRegistry) pulumi.StringPtrOutput { return v.SchemaCompatibility }).(pulumi.StringPtrOutput)
}

func (o SchemaRegistryOutput) SchemaType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SchemaRegistry) pulumi.StringPtrOutput { return v.SchemaType }).(pulumi.StringPtrOutput)
}

// The system meta data relating to this resource.
func (o SchemaRegistryOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *SchemaRegistry) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.EventHub/Namespaces" or "Microsoft.EventHub/Namespaces/EventHubs"
func (o SchemaRegistryOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SchemaRegistry) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Exact time the Schema Group was updated
func (o SchemaRegistryOutput) UpdatedAtUtc() pulumi.StringOutput {
	return o.ApplyT(func(v *SchemaRegistry) pulumi.StringOutput { return v.UpdatedAtUtc }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SchemaRegistryOutput{})
}
