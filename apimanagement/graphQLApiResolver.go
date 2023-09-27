// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apimanagement

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// GraphQL API Resolver details.
// Azure REST API version: 2022-08-01.
type GraphQLApiResolver struct {
	pulumi.CustomResourceState

	// Description of the resolver. May include HTML formatting tags.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Resolver Name.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Path is type/field being resolved.
	Path pulumi.StringPtrOutput `pulumi:"path"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewGraphQLApiResolver registers a new resource with the given unique name, arguments, and options.
func NewGraphQLApiResolver(ctx *pulumi.Context,
	name string, args *GraphQLApiResolverArgs, opts ...pulumi.ResourceOption) (*GraphQLApiResolver, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiId == nil {
		return nil, errors.New("invalid value for required argument 'ApiId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:apimanagement/v20220801:GraphQLApiResolver"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220901preview:GraphQLApiResolver"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20230301preview:GraphQLApiResolver"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource GraphQLApiResolver
	err := ctx.RegisterResource("azure-native:apimanagement:GraphQLApiResolver", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGraphQLApiResolver gets an existing GraphQLApiResolver resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGraphQLApiResolver(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GraphQLApiResolverState, opts ...pulumi.ResourceOption) (*GraphQLApiResolver, error) {
	var resource GraphQLApiResolver
	err := ctx.ReadResource("azure-native:apimanagement:GraphQLApiResolver", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GraphQLApiResolver resources.
type graphQLApiResolverState struct {
}

type GraphQLApiResolverState struct {
}

func (GraphQLApiResolverState) ElementType() reflect.Type {
	return reflect.TypeOf((*graphQLApiResolverState)(nil)).Elem()
}

type graphQLApiResolverArgs struct {
	// API revision identifier. Must be unique in the current API Management service instance. Non-current revision has ;rev=n as a suffix where n is the revision number.
	ApiId string `pulumi:"apiId"`
	// Description of the resolver. May include HTML formatting tags.
	Description *string `pulumi:"description"`
	// Resolver Name.
	DisplayName *string `pulumi:"displayName"`
	// Path is type/field being resolved.
	Path *string `pulumi:"path"`
	// Resolver identifier within a GraphQL API. Must be unique in the current API Management service instance.
	ResolverId *string `pulumi:"resolverId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a GraphQLApiResolver resource.
type GraphQLApiResolverArgs struct {
	// API revision identifier. Must be unique in the current API Management service instance. Non-current revision has ;rev=n as a suffix where n is the revision number.
	ApiId pulumi.StringInput
	// Description of the resolver. May include HTML formatting tags.
	Description pulumi.StringPtrInput
	// Resolver Name.
	DisplayName pulumi.StringPtrInput
	// Path is type/field being resolved.
	Path pulumi.StringPtrInput
	// Resolver identifier within a GraphQL API. Must be unique in the current API Management service instance.
	ResolverId pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the API Management service.
	ServiceName pulumi.StringInput
}

func (GraphQLApiResolverArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*graphQLApiResolverArgs)(nil)).Elem()
}

type GraphQLApiResolverInput interface {
	pulumi.Input

	ToGraphQLApiResolverOutput() GraphQLApiResolverOutput
	ToGraphQLApiResolverOutputWithContext(ctx context.Context) GraphQLApiResolverOutput
}

func (*GraphQLApiResolver) ElementType() reflect.Type {
	return reflect.TypeOf((**GraphQLApiResolver)(nil)).Elem()
}

func (i *GraphQLApiResolver) ToGraphQLApiResolverOutput() GraphQLApiResolverOutput {
	return i.ToGraphQLApiResolverOutputWithContext(context.Background())
}

func (i *GraphQLApiResolver) ToGraphQLApiResolverOutputWithContext(ctx context.Context) GraphQLApiResolverOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GraphQLApiResolverOutput)
}

func (i *GraphQLApiResolver) ToOutput(ctx context.Context) pulumix.Output[*GraphQLApiResolver] {
	return pulumix.Output[*GraphQLApiResolver]{
		OutputState: i.ToGraphQLApiResolverOutputWithContext(ctx).OutputState,
	}
}

type GraphQLApiResolverOutput struct{ *pulumi.OutputState }

func (GraphQLApiResolverOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GraphQLApiResolver)(nil)).Elem()
}

func (o GraphQLApiResolverOutput) ToGraphQLApiResolverOutput() GraphQLApiResolverOutput {
	return o
}

func (o GraphQLApiResolverOutput) ToGraphQLApiResolverOutputWithContext(ctx context.Context) GraphQLApiResolverOutput {
	return o
}

func (o GraphQLApiResolverOutput) ToOutput(ctx context.Context) pulumix.Output[*GraphQLApiResolver] {
	return pulumix.Output[*GraphQLApiResolver]{
		OutputState: o.OutputState,
	}
}

// Description of the resolver. May include HTML formatting tags.
func (o GraphQLApiResolverOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GraphQLApiResolver) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Resolver Name.
func (o GraphQLApiResolverOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GraphQLApiResolver) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o GraphQLApiResolverOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *GraphQLApiResolver) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Path is type/field being resolved.
func (o GraphQLApiResolverOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GraphQLApiResolver) pulumi.StringPtrOutput { return v.Path }).(pulumi.StringPtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o GraphQLApiResolverOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *GraphQLApiResolver) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GraphQLApiResolverOutput{})
}
