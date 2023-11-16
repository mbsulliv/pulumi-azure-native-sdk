// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20240301

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Metadata schema entity. Used to define metadata for the entities in API catalog.
type MetadataSchema struct {
	pulumi.CustomResourceState

	AssignedTo MetadataAssignmentResponseArrayOutput `pulumi:"assignedTo"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The schema defining the type.
	Schema pulumi.StringOutput `pulumi:"schema"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewMetadataSchema registers a new resource with the given unique name, arguments, and options.
func NewMetadataSchema(ctx *pulumi.Context,
	name string, args *MetadataSchemaArgs, opts ...pulumi.ResourceOption) (*MetadataSchema, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Schema == nil {
		return nil, errors.New("invalid value for required argument 'Schema'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:apicenter:MetadataSchema"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource MetadataSchema
	err := ctx.RegisterResource("azure-native:apicenter/v20240301:MetadataSchema", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMetadataSchema gets an existing MetadataSchema resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMetadataSchema(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MetadataSchemaState, opts ...pulumi.ResourceOption) (*MetadataSchema, error) {
	var resource MetadataSchema
	err := ctx.ReadResource("azure-native:apicenter/v20240301:MetadataSchema", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MetadataSchema resources.
type metadataSchemaState struct {
}

type MetadataSchemaState struct {
}

func (MetadataSchemaState) ElementType() reflect.Type {
	return reflect.TypeOf((*metadataSchemaState)(nil)).Elem()
}

type metadataSchemaArgs struct {
	AssignedTo []MetadataAssignment `pulumi:"assignedTo"`
	// The name of the metadata schema.
	MetadataSchemaName *string `pulumi:"metadataSchemaName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The schema defining the type.
	Schema string `pulumi:"schema"`
	// The name of Azure API Center service.
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a MetadataSchema resource.
type MetadataSchemaArgs struct {
	AssignedTo MetadataAssignmentArrayInput
	// The name of the metadata schema.
	MetadataSchemaName pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The schema defining the type.
	Schema pulumi.StringInput
	// The name of Azure API Center service.
	ServiceName pulumi.StringInput
}

func (MetadataSchemaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*metadataSchemaArgs)(nil)).Elem()
}

type MetadataSchemaInput interface {
	pulumi.Input

	ToMetadataSchemaOutput() MetadataSchemaOutput
	ToMetadataSchemaOutputWithContext(ctx context.Context) MetadataSchemaOutput
}

func (*MetadataSchema) ElementType() reflect.Type {
	return reflect.TypeOf((**MetadataSchema)(nil)).Elem()
}

func (i *MetadataSchema) ToMetadataSchemaOutput() MetadataSchemaOutput {
	return i.ToMetadataSchemaOutputWithContext(context.Background())
}

func (i *MetadataSchema) ToMetadataSchemaOutputWithContext(ctx context.Context) MetadataSchemaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetadataSchemaOutput)
}

type MetadataSchemaOutput struct{ *pulumi.OutputState }

func (MetadataSchemaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MetadataSchema)(nil)).Elem()
}

func (o MetadataSchemaOutput) ToMetadataSchemaOutput() MetadataSchemaOutput {
	return o
}

func (o MetadataSchemaOutput) ToMetadataSchemaOutputWithContext(ctx context.Context) MetadataSchemaOutput {
	return o
}

func (o MetadataSchemaOutput) AssignedTo() MetadataAssignmentResponseArrayOutput {
	return o.ApplyT(func(v *MetadataSchema) MetadataAssignmentResponseArrayOutput { return v.AssignedTo }).(MetadataAssignmentResponseArrayOutput)
}

// The name of the resource
func (o MetadataSchemaOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MetadataSchema) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The schema defining the type.
func (o MetadataSchemaOutput) Schema() pulumi.StringOutput {
	return o.ApplyT(func(v *MetadataSchema) pulumi.StringOutput { return v.Schema }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o MetadataSchemaOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *MetadataSchema) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o MetadataSchemaOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *MetadataSchema) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(MetadataSchemaOutput{})
}
