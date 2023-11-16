// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220901preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Global Schema Contract details.
type WorkspaceGlobalSchema struct {
	pulumi.CustomResourceState

	// Free-form schema entity description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Schema Type. Immutable.
	SchemaType pulumi.StringOutput `pulumi:"schemaType"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// Json-encoded string for non json-based schema.
	Value pulumi.AnyOutput `pulumi:"value"`
}

// NewWorkspaceGlobalSchema registers a new resource with the given unique name, arguments, and options.
func NewWorkspaceGlobalSchema(ctx *pulumi.Context,
	name string, args *WorkspaceGlobalSchemaArgs, opts ...pulumi.ResourceOption) (*WorkspaceGlobalSchema, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SchemaType == nil {
		return nil, errors.New("invalid value for required argument 'SchemaType'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	if args.WorkspaceId == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceId'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:apimanagement:WorkspaceGlobalSchema"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20230301preview:WorkspaceGlobalSchema"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource WorkspaceGlobalSchema
	err := ctx.RegisterResource("azure-native:apimanagement/v20220901preview:WorkspaceGlobalSchema", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkspaceGlobalSchema gets an existing WorkspaceGlobalSchema resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkspaceGlobalSchema(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkspaceGlobalSchemaState, opts ...pulumi.ResourceOption) (*WorkspaceGlobalSchema, error) {
	var resource WorkspaceGlobalSchema
	err := ctx.ReadResource("azure-native:apimanagement/v20220901preview:WorkspaceGlobalSchema", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WorkspaceGlobalSchema resources.
type workspaceGlobalSchemaState struct {
}

type WorkspaceGlobalSchemaState struct {
}

func (WorkspaceGlobalSchemaState) ElementType() reflect.Type {
	return reflect.TypeOf((*workspaceGlobalSchemaState)(nil)).Elem()
}

type workspaceGlobalSchemaArgs struct {
	// Free-form schema entity description.
	Description *string `pulumi:"description"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Schema id identifier. Must be unique in the current API Management service instance.
	SchemaId *string `pulumi:"schemaId"`
	// Schema Type. Immutable.
	SchemaType string `pulumi:"schemaType"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
	// Json-encoded string for non json-based schema.
	Value interface{} `pulumi:"value"`
	// Workspace identifier. Must be unique in the current API Management service instance.
	WorkspaceId string `pulumi:"workspaceId"`
}

// The set of arguments for constructing a WorkspaceGlobalSchema resource.
type WorkspaceGlobalSchemaArgs struct {
	// Free-form schema entity description.
	Description pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Schema id identifier. Must be unique in the current API Management service instance.
	SchemaId pulumi.StringPtrInput
	// Schema Type. Immutable.
	SchemaType pulumi.StringInput
	// The name of the API Management service.
	ServiceName pulumi.StringInput
	// Json-encoded string for non json-based schema.
	Value pulumi.Input
	// Workspace identifier. Must be unique in the current API Management service instance.
	WorkspaceId pulumi.StringInput
}

func (WorkspaceGlobalSchemaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workspaceGlobalSchemaArgs)(nil)).Elem()
}

type WorkspaceGlobalSchemaInput interface {
	pulumi.Input

	ToWorkspaceGlobalSchemaOutput() WorkspaceGlobalSchemaOutput
	ToWorkspaceGlobalSchemaOutputWithContext(ctx context.Context) WorkspaceGlobalSchemaOutput
}

func (*WorkspaceGlobalSchema) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceGlobalSchema)(nil)).Elem()
}

func (i *WorkspaceGlobalSchema) ToWorkspaceGlobalSchemaOutput() WorkspaceGlobalSchemaOutput {
	return i.ToWorkspaceGlobalSchemaOutputWithContext(context.Background())
}

func (i *WorkspaceGlobalSchema) ToWorkspaceGlobalSchemaOutputWithContext(ctx context.Context) WorkspaceGlobalSchemaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceGlobalSchemaOutput)
}

type WorkspaceGlobalSchemaOutput struct{ *pulumi.OutputState }

func (WorkspaceGlobalSchemaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceGlobalSchema)(nil)).Elem()
}

func (o WorkspaceGlobalSchemaOutput) ToWorkspaceGlobalSchemaOutput() WorkspaceGlobalSchemaOutput {
	return o
}

func (o WorkspaceGlobalSchemaOutput) ToWorkspaceGlobalSchemaOutputWithContext(ctx context.Context) WorkspaceGlobalSchemaOutput {
	return o
}

// Free-form schema entity description.
func (o WorkspaceGlobalSchemaOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkspaceGlobalSchema) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o WorkspaceGlobalSchemaOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkspaceGlobalSchema) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Schema Type. Immutable.
func (o WorkspaceGlobalSchemaOutput) SchemaType() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkspaceGlobalSchema) pulumi.StringOutput { return v.SchemaType }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o WorkspaceGlobalSchemaOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkspaceGlobalSchema) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Json-encoded string for non json-based schema.
func (o WorkspaceGlobalSchemaOutput) Value() pulumi.AnyOutput {
	return o.ApplyT(func(v *WorkspaceGlobalSchema) pulumi.AnyOutput { return v.Value }).(pulumi.AnyOutput)
}

func init() {
	pulumi.RegisterOutputType(WorkspaceGlobalSchemaOutput{})
}
