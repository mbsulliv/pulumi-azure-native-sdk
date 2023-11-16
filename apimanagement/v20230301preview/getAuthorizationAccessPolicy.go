// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230301preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the details of the authorization access policy specified by its identifier.
func LookupAuthorizationAccessPolicy(ctx *pulumi.Context, args *LookupAuthorizationAccessPolicyArgs, opts ...pulumi.InvokeOption) (*LookupAuthorizationAccessPolicyResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupAuthorizationAccessPolicyResult
	err := ctx.Invoke("azure-native:apimanagement/v20230301preview:getAuthorizationAccessPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAuthorizationAccessPolicyArgs struct {
	// Identifier of the authorization access policy.
	AuthorizationAccessPolicyId string `pulumi:"authorizationAccessPolicyId"`
	// Identifier of the authorization.
	AuthorizationId string `pulumi:"authorizationId"`
	// Identifier of the authorization provider.
	AuthorizationProviderId string `pulumi:"authorizationProviderId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
}

// Authorization access policy contract.
type LookupAuthorizationAccessPolicyResult struct {
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The Object Id
	ObjectId *string `pulumi:"objectId"`
	// The Tenant Id
	TenantId *string `pulumi:"tenantId"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupAuthorizationAccessPolicyOutput(ctx *pulumi.Context, args LookupAuthorizationAccessPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupAuthorizationAccessPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAuthorizationAccessPolicyResult, error) {
			args := v.(LookupAuthorizationAccessPolicyArgs)
			r, err := LookupAuthorizationAccessPolicy(ctx, &args, opts...)
			var s LookupAuthorizationAccessPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAuthorizationAccessPolicyResultOutput)
}

type LookupAuthorizationAccessPolicyOutputArgs struct {
	// Identifier of the authorization access policy.
	AuthorizationAccessPolicyId pulumi.StringInput `pulumi:"authorizationAccessPolicyId"`
	// Identifier of the authorization.
	AuthorizationId pulumi.StringInput `pulumi:"authorizationId"`
	// Identifier of the authorization provider.
	AuthorizationProviderId pulumi.StringInput `pulumi:"authorizationProviderId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupAuthorizationAccessPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAuthorizationAccessPolicyArgs)(nil)).Elem()
}

// Authorization access policy contract.
type LookupAuthorizationAccessPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupAuthorizationAccessPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAuthorizationAccessPolicyResult)(nil)).Elem()
}

func (o LookupAuthorizationAccessPolicyResultOutput) ToLookupAuthorizationAccessPolicyResultOutput() LookupAuthorizationAccessPolicyResultOutput {
	return o
}

func (o LookupAuthorizationAccessPolicyResultOutput) ToLookupAuthorizationAccessPolicyResultOutputWithContext(ctx context.Context) LookupAuthorizationAccessPolicyResultOutput {
	return o
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupAuthorizationAccessPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizationAccessPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupAuthorizationAccessPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizationAccessPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

// The Object Id
func (o LookupAuthorizationAccessPolicyResultOutput) ObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAuthorizationAccessPolicyResult) *string { return v.ObjectId }).(pulumi.StringPtrOutput)
}

// The Tenant Id
func (o LookupAuthorizationAccessPolicyResultOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAuthorizationAccessPolicyResult) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupAuthorizationAccessPolicyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizationAccessPolicyResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAuthorizationAccessPolicyResultOutput{})
}
