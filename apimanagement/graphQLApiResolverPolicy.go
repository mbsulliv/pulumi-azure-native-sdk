// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apimanagement

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Policy Contract details.
// Azure REST API version: 2022-08-01.
//
// Other available API versions: 2022-09-01-preview, 2023-03-01-preview.
type GraphQLApiResolverPolicy struct {
	pulumi.CustomResourceState

	// Format of the policyContent.
	Format pulumi.StringPtrOutput `pulumi:"format"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// Contents of the Policy as defined by the format.
	Value pulumi.StringOutput `pulumi:"value"`
}

// NewGraphQLApiResolverPolicy registers a new resource with the given unique name, arguments, and options.
func NewGraphQLApiResolverPolicy(ctx *pulumi.Context,
	name string, args *GraphQLApiResolverPolicyArgs, opts ...pulumi.ResourceOption) (*GraphQLApiResolverPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiId == nil {
		return nil, errors.New("invalid value for required argument 'ApiId'")
	}
	if args.ResolverId == nil {
		return nil, errors.New("invalid value for required argument 'ResolverId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	if args.Value == nil {
		return nil, errors.New("invalid value for required argument 'Value'")
	}
	if args.Format == nil {
		args.Format = pulumi.StringPtr("xml")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:apimanagement/v20220801:GraphQLApiResolverPolicy"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220901preview:GraphQLApiResolverPolicy"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20230301preview:GraphQLApiResolverPolicy"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource GraphQLApiResolverPolicy
	err := ctx.RegisterResource("azure-native:apimanagement:GraphQLApiResolverPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGraphQLApiResolverPolicy gets an existing GraphQLApiResolverPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGraphQLApiResolverPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GraphQLApiResolverPolicyState, opts ...pulumi.ResourceOption) (*GraphQLApiResolverPolicy, error) {
	var resource GraphQLApiResolverPolicy
	err := ctx.ReadResource("azure-native:apimanagement:GraphQLApiResolverPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GraphQLApiResolverPolicy resources.
type graphQLApiResolverPolicyState struct {
}

type GraphQLApiResolverPolicyState struct {
}

func (GraphQLApiResolverPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*graphQLApiResolverPolicyState)(nil)).Elem()
}

type graphQLApiResolverPolicyArgs struct {
	// API revision identifier. Must be unique in the current API Management service instance. Non-current revision has ;rev=n as a suffix where n is the revision number.
	ApiId string `pulumi:"apiId"`
	// Format of the policyContent.
	Format *string `pulumi:"format"`
	// The identifier of the Policy.
	PolicyId *string `pulumi:"policyId"`
	// Resolver identifier within a GraphQL API. Must be unique in the current API Management service instance.
	ResolverId string `pulumi:"resolverId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
	// Contents of the Policy as defined by the format.
	Value string `pulumi:"value"`
}

// The set of arguments for constructing a GraphQLApiResolverPolicy resource.
type GraphQLApiResolverPolicyArgs struct {
	// API revision identifier. Must be unique in the current API Management service instance. Non-current revision has ;rev=n as a suffix where n is the revision number.
	ApiId pulumi.StringInput
	// Format of the policyContent.
	Format pulumi.StringPtrInput
	// The identifier of the Policy.
	PolicyId pulumi.StringPtrInput
	// Resolver identifier within a GraphQL API. Must be unique in the current API Management service instance.
	ResolverId pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the API Management service.
	ServiceName pulumi.StringInput
	// Contents of the Policy as defined by the format.
	Value pulumi.StringInput
}

func (GraphQLApiResolverPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*graphQLApiResolverPolicyArgs)(nil)).Elem()
}

type GraphQLApiResolverPolicyInput interface {
	pulumi.Input

	ToGraphQLApiResolverPolicyOutput() GraphQLApiResolverPolicyOutput
	ToGraphQLApiResolverPolicyOutputWithContext(ctx context.Context) GraphQLApiResolverPolicyOutput
}

func (*GraphQLApiResolverPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**GraphQLApiResolverPolicy)(nil)).Elem()
}

func (i *GraphQLApiResolverPolicy) ToGraphQLApiResolverPolicyOutput() GraphQLApiResolverPolicyOutput {
	return i.ToGraphQLApiResolverPolicyOutputWithContext(context.Background())
}

func (i *GraphQLApiResolverPolicy) ToGraphQLApiResolverPolicyOutputWithContext(ctx context.Context) GraphQLApiResolverPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GraphQLApiResolverPolicyOutput)
}

type GraphQLApiResolverPolicyOutput struct{ *pulumi.OutputState }

func (GraphQLApiResolverPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GraphQLApiResolverPolicy)(nil)).Elem()
}

func (o GraphQLApiResolverPolicyOutput) ToGraphQLApiResolverPolicyOutput() GraphQLApiResolverPolicyOutput {
	return o
}

func (o GraphQLApiResolverPolicyOutput) ToGraphQLApiResolverPolicyOutputWithContext(ctx context.Context) GraphQLApiResolverPolicyOutput {
	return o
}

// Format of the policyContent.
func (o GraphQLApiResolverPolicyOutput) Format() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GraphQLApiResolverPolicy) pulumi.StringPtrOutput { return v.Format }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o GraphQLApiResolverPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *GraphQLApiResolverPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o GraphQLApiResolverPolicyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *GraphQLApiResolverPolicy) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Contents of the Policy as defined by the format.
func (o GraphQLApiResolverPolicyOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v *GraphQLApiResolverPolicy) pulumi.StringOutput { return v.Value }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GraphQLApiResolverPolicyOutput{})
}
