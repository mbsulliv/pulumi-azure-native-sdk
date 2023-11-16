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

// API Schema Contract details.
type WorkspaceApiSchema struct {
	pulumi.CustomResourceState

	// Types definitions. Used for Swagger/OpenAPI v2/v3 schemas only, null otherwise.
	Components pulumi.AnyOutput `pulumi:"components"`
	// Must be a valid a media type used in a Content-Type header as defined in the RFC 2616. Media type of the schema document (e.g. application/json, application/xml). </br> - `Swagger` Schema use `application/vnd.ms-azure-apim.swagger.definitions+json` </br> - `WSDL` Schema use `application/vnd.ms-azure-apim.xsd+xml` </br> - `OpenApi` Schema use `application/vnd.oai.openapi.components+json` </br> - `WADL Schema` use `application/vnd.ms-azure-apim.wadl.grammars+xml`.
	ContentType pulumi.StringOutput `pulumi:"contentType"`
	// Types definitions. Used for Swagger/OpenAPI v1 schemas only, null otherwise.
	Definitions pulumi.AnyOutput `pulumi:"definitions"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// Json escaped string defining the document representing the Schema. Used for schemas other than Swagger/OpenAPI.
	Value pulumi.StringPtrOutput `pulumi:"value"`
}

// NewWorkspaceApiSchema registers a new resource with the given unique name, arguments, and options.
func NewWorkspaceApiSchema(ctx *pulumi.Context,
	name string, args *WorkspaceApiSchemaArgs, opts ...pulumi.ResourceOption) (*WorkspaceApiSchema, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiId == nil {
		return nil, errors.New("invalid value for required argument 'ApiId'")
	}
	if args.ContentType == nil {
		return nil, errors.New("invalid value for required argument 'ContentType'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	if args.WorkspaceId == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceId'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:apimanagement:WorkspaceApiSchema"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20230301preview:WorkspaceApiSchema"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource WorkspaceApiSchema
	err := ctx.RegisterResource("azure-native:apimanagement/v20220901preview:WorkspaceApiSchema", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkspaceApiSchema gets an existing WorkspaceApiSchema resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkspaceApiSchema(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkspaceApiSchemaState, opts ...pulumi.ResourceOption) (*WorkspaceApiSchema, error) {
	var resource WorkspaceApiSchema
	err := ctx.ReadResource("azure-native:apimanagement/v20220901preview:WorkspaceApiSchema", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WorkspaceApiSchema resources.
type workspaceApiSchemaState struct {
}

type WorkspaceApiSchemaState struct {
}

func (WorkspaceApiSchemaState) ElementType() reflect.Type {
	return reflect.TypeOf((*workspaceApiSchemaState)(nil)).Elem()
}

type workspaceApiSchemaArgs struct {
	// API revision identifier. Must be unique in the current API Management service instance. Non-current revision has ;rev=n as a suffix where n is the revision number.
	ApiId string `pulumi:"apiId"`
	// Types definitions. Used for Swagger/OpenAPI v2/v3 schemas only, null otherwise.
	Components interface{} `pulumi:"components"`
	// Must be a valid a media type used in a Content-Type header as defined in the RFC 2616. Media type of the schema document (e.g. application/json, application/xml). </br> - `Swagger` Schema use `application/vnd.ms-azure-apim.swagger.definitions+json` </br> - `WSDL` Schema use `application/vnd.ms-azure-apim.xsd+xml` </br> - `OpenApi` Schema use `application/vnd.oai.openapi.components+json` </br> - `WADL Schema` use `application/vnd.ms-azure-apim.wadl.grammars+xml`.
	ContentType string `pulumi:"contentType"`
	// Types definitions. Used for Swagger/OpenAPI v1 schemas only, null otherwise.
	Definitions interface{} `pulumi:"definitions"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Schema id identifier. Must be unique in the current API Management service instance.
	SchemaId *string `pulumi:"schemaId"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
	// Json escaped string defining the document representing the Schema. Used for schemas other than Swagger/OpenAPI.
	Value *string `pulumi:"value"`
	// Workspace identifier. Must be unique in the current API Management service instance.
	WorkspaceId string `pulumi:"workspaceId"`
}

// The set of arguments for constructing a WorkspaceApiSchema resource.
type WorkspaceApiSchemaArgs struct {
	// API revision identifier. Must be unique in the current API Management service instance. Non-current revision has ;rev=n as a suffix where n is the revision number.
	ApiId pulumi.StringInput
	// Types definitions. Used for Swagger/OpenAPI v2/v3 schemas only, null otherwise.
	Components pulumi.Input
	// Must be a valid a media type used in a Content-Type header as defined in the RFC 2616. Media type of the schema document (e.g. application/json, application/xml). </br> - `Swagger` Schema use `application/vnd.ms-azure-apim.swagger.definitions+json` </br> - `WSDL` Schema use `application/vnd.ms-azure-apim.xsd+xml` </br> - `OpenApi` Schema use `application/vnd.oai.openapi.components+json` </br> - `WADL Schema` use `application/vnd.ms-azure-apim.wadl.grammars+xml`.
	ContentType pulumi.StringInput
	// Types definitions. Used for Swagger/OpenAPI v1 schemas only, null otherwise.
	Definitions pulumi.Input
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Schema id identifier. Must be unique in the current API Management service instance.
	SchemaId pulumi.StringPtrInput
	// The name of the API Management service.
	ServiceName pulumi.StringInput
	// Json escaped string defining the document representing the Schema. Used for schemas other than Swagger/OpenAPI.
	Value pulumi.StringPtrInput
	// Workspace identifier. Must be unique in the current API Management service instance.
	WorkspaceId pulumi.StringInput
}

func (WorkspaceApiSchemaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workspaceApiSchemaArgs)(nil)).Elem()
}

type WorkspaceApiSchemaInput interface {
	pulumi.Input

	ToWorkspaceApiSchemaOutput() WorkspaceApiSchemaOutput
	ToWorkspaceApiSchemaOutputWithContext(ctx context.Context) WorkspaceApiSchemaOutput
}

func (*WorkspaceApiSchema) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceApiSchema)(nil)).Elem()
}

func (i *WorkspaceApiSchema) ToWorkspaceApiSchemaOutput() WorkspaceApiSchemaOutput {
	return i.ToWorkspaceApiSchemaOutputWithContext(context.Background())
}

func (i *WorkspaceApiSchema) ToWorkspaceApiSchemaOutputWithContext(ctx context.Context) WorkspaceApiSchemaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceApiSchemaOutput)
}

type WorkspaceApiSchemaOutput struct{ *pulumi.OutputState }

func (WorkspaceApiSchemaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceApiSchema)(nil)).Elem()
}

func (o WorkspaceApiSchemaOutput) ToWorkspaceApiSchemaOutput() WorkspaceApiSchemaOutput {
	return o
}

func (o WorkspaceApiSchemaOutput) ToWorkspaceApiSchemaOutputWithContext(ctx context.Context) WorkspaceApiSchemaOutput {
	return o
}

// Types definitions. Used for Swagger/OpenAPI v2/v3 schemas only, null otherwise.
func (o WorkspaceApiSchemaOutput) Components() pulumi.AnyOutput {
	return o.ApplyT(func(v *WorkspaceApiSchema) pulumi.AnyOutput { return v.Components }).(pulumi.AnyOutput)
}

// Must be a valid a media type used in a Content-Type header as defined in the RFC 2616. Media type of the schema document (e.g. application/json, application/xml). </br> - `Swagger` Schema use `application/vnd.ms-azure-apim.swagger.definitions+json` </br> - `WSDL` Schema use `application/vnd.ms-azure-apim.xsd+xml` </br> - `OpenApi` Schema use `application/vnd.oai.openapi.components+json` </br> - `WADL Schema` use `application/vnd.ms-azure-apim.wadl.grammars+xml`.
func (o WorkspaceApiSchemaOutput) ContentType() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkspaceApiSchema) pulumi.StringOutput { return v.ContentType }).(pulumi.StringOutput)
}

// Types definitions. Used for Swagger/OpenAPI v1 schemas only, null otherwise.
func (o WorkspaceApiSchemaOutput) Definitions() pulumi.AnyOutput {
	return o.ApplyT(func(v *WorkspaceApiSchema) pulumi.AnyOutput { return v.Definitions }).(pulumi.AnyOutput)
}

// The name of the resource
func (o WorkspaceApiSchemaOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkspaceApiSchema) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o WorkspaceApiSchemaOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkspaceApiSchema) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Json escaped string defining the document representing the Schema. Used for schemas other than Swagger/OpenAPI.
func (o WorkspaceApiSchemaOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkspaceApiSchema) pulumi.StringPtrOutput { return v.Value }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(WorkspaceApiSchemaOutput{})
}
