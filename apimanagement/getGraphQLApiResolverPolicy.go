// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apimanagement

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get the policy configuration at the GraphQL API Resolver level.
// Azure REST API version: 2022-08-01.
func LookupGraphQLApiResolverPolicy(ctx *pulumi.Context, args *LookupGraphQLApiResolverPolicyArgs, opts ...pulumi.InvokeOption) (*LookupGraphQLApiResolverPolicyResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupGraphQLApiResolverPolicyResult
	err := ctx.Invoke("azure-native:apimanagement:getGraphQLApiResolverPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupGraphQLApiResolverPolicyArgs struct {
	// API revision identifier. Must be unique in the current API Management service instance. Non-current revision has ;rev=n as a suffix where n is the revision number.
	ApiId string `pulumi:"apiId"`
	// Policy Export Format.
	Format *string `pulumi:"format"`
	// The identifier of the Policy.
	PolicyId string `pulumi:"policyId"`
	// Resolver identifier within a GraphQL API. Must be unique in the current API Management service instance.
	ResolverId string `pulumi:"resolverId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
}

// Policy Contract details.
type LookupGraphQLApiResolverPolicyResult struct {
	// Format of the policyContent.
	Format *string `pulumi:"format"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// Contents of the Policy as defined by the format.
	Value string `pulumi:"value"`
}

// Defaults sets the appropriate defaults for LookupGraphQLApiResolverPolicyResult
func (val *LookupGraphQLApiResolverPolicyResult) Defaults() *LookupGraphQLApiResolverPolicyResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.Format == nil {
		format_ := "xml"
		tmp.Format = &format_
	}
	return &tmp
}

func LookupGraphQLApiResolverPolicyOutput(ctx *pulumi.Context, args LookupGraphQLApiResolverPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupGraphQLApiResolverPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGraphQLApiResolverPolicyResult, error) {
			args := v.(LookupGraphQLApiResolverPolicyArgs)
			r, err := LookupGraphQLApiResolverPolicy(ctx, &args, opts...)
			var s LookupGraphQLApiResolverPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGraphQLApiResolverPolicyResultOutput)
}

type LookupGraphQLApiResolverPolicyOutputArgs struct {
	// API revision identifier. Must be unique in the current API Management service instance. Non-current revision has ;rev=n as a suffix where n is the revision number.
	ApiId pulumi.StringInput `pulumi:"apiId"`
	// Policy Export Format.
	Format pulumi.StringPtrInput `pulumi:"format"`
	// The identifier of the Policy.
	PolicyId pulumi.StringInput `pulumi:"policyId"`
	// Resolver identifier within a GraphQL API. Must be unique in the current API Management service instance.
	ResolverId pulumi.StringInput `pulumi:"resolverId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupGraphQLApiResolverPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGraphQLApiResolverPolicyArgs)(nil)).Elem()
}

// Policy Contract details.
type LookupGraphQLApiResolverPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupGraphQLApiResolverPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGraphQLApiResolverPolicyResult)(nil)).Elem()
}

func (o LookupGraphQLApiResolverPolicyResultOutput) ToLookupGraphQLApiResolverPolicyResultOutput() LookupGraphQLApiResolverPolicyResultOutput {
	return o
}

func (o LookupGraphQLApiResolverPolicyResultOutput) ToLookupGraphQLApiResolverPolicyResultOutputWithContext(ctx context.Context) LookupGraphQLApiResolverPolicyResultOutput {
	return o
}

func (o LookupGraphQLApiResolverPolicyResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupGraphQLApiResolverPolicyResult] {
	return pulumix.Output[LookupGraphQLApiResolverPolicyResult]{
		OutputState: o.OutputState,
	}
}

// Format of the policyContent.
func (o LookupGraphQLApiResolverPolicyResultOutput) Format() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGraphQLApiResolverPolicyResult) *string { return v.Format }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupGraphQLApiResolverPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGraphQLApiResolverPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupGraphQLApiResolverPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGraphQLApiResolverPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupGraphQLApiResolverPolicyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGraphQLApiResolverPolicyResult) string { return v.Type }).(pulumi.StringOutput)
}

// Contents of the Policy as defined by the format.
func (o LookupGraphQLApiResolverPolicyResultOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGraphQLApiResolverPolicyResult) string { return v.Value }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGraphQLApiResolverPolicyResultOutput{})
}
