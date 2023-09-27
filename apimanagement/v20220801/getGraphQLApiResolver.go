// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets the details of the GraphQL API Resolver specified by its identifier.
func LookupGraphQLApiResolver(ctx *pulumi.Context, args *LookupGraphQLApiResolverArgs, opts ...pulumi.InvokeOption) (*LookupGraphQLApiResolverResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupGraphQLApiResolverResult
	err := ctx.Invoke("azure-native:apimanagement/v20220801:getGraphQLApiResolver", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGraphQLApiResolverArgs struct {
	// API revision identifier. Must be unique in the current API Management service instance. Non-current revision has ;rev=n as a suffix where n is the revision number.
	ApiId string `pulumi:"apiId"`
	// Resolver identifier within a GraphQL API. Must be unique in the current API Management service instance.
	ResolverId string `pulumi:"resolverId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
}

// GraphQL API Resolver details.
type LookupGraphQLApiResolverResult struct {
	// Description of the resolver. May include HTML formatting tags.
	Description *string `pulumi:"description"`
	// Resolver Name.
	DisplayName *string `pulumi:"displayName"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Path is type/field being resolved.
	Path *string `pulumi:"path"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupGraphQLApiResolverOutput(ctx *pulumi.Context, args LookupGraphQLApiResolverOutputArgs, opts ...pulumi.InvokeOption) LookupGraphQLApiResolverResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGraphQLApiResolverResult, error) {
			args := v.(LookupGraphQLApiResolverArgs)
			r, err := LookupGraphQLApiResolver(ctx, &args, opts...)
			var s LookupGraphQLApiResolverResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGraphQLApiResolverResultOutput)
}

type LookupGraphQLApiResolverOutputArgs struct {
	// API revision identifier. Must be unique in the current API Management service instance. Non-current revision has ;rev=n as a suffix where n is the revision number.
	ApiId pulumi.StringInput `pulumi:"apiId"`
	// Resolver identifier within a GraphQL API. Must be unique in the current API Management service instance.
	ResolverId pulumi.StringInput `pulumi:"resolverId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupGraphQLApiResolverOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGraphQLApiResolverArgs)(nil)).Elem()
}

// GraphQL API Resolver details.
type LookupGraphQLApiResolverResultOutput struct{ *pulumi.OutputState }

func (LookupGraphQLApiResolverResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGraphQLApiResolverResult)(nil)).Elem()
}

func (o LookupGraphQLApiResolverResultOutput) ToLookupGraphQLApiResolverResultOutput() LookupGraphQLApiResolverResultOutput {
	return o
}

func (o LookupGraphQLApiResolverResultOutput) ToLookupGraphQLApiResolverResultOutputWithContext(ctx context.Context) LookupGraphQLApiResolverResultOutput {
	return o
}

func (o LookupGraphQLApiResolverResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupGraphQLApiResolverResult] {
	return pulumix.Output[LookupGraphQLApiResolverResult]{
		OutputState: o.OutputState,
	}
}

// Description of the resolver. May include HTML formatting tags.
func (o LookupGraphQLApiResolverResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGraphQLApiResolverResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Resolver Name.
func (o LookupGraphQLApiResolverResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGraphQLApiResolverResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupGraphQLApiResolverResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGraphQLApiResolverResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupGraphQLApiResolverResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGraphQLApiResolverResult) string { return v.Name }).(pulumi.StringOutput)
}

// Path is type/field being resolved.
func (o LookupGraphQLApiResolverResultOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGraphQLApiResolverResult) *string { return v.Path }).(pulumi.StringPtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupGraphQLApiResolverResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGraphQLApiResolverResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGraphQLApiResolverResultOutput{})
}
