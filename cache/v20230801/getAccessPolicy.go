// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the detailed information about an access policy of a redis cache
func LookupAccessPolicy(ctx *pulumi.Context, args *LookupAccessPolicyArgs, opts ...pulumi.InvokeOption) (*LookupAccessPolicyResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupAccessPolicyResult
	err := ctx.Invoke("azure-native:cache/v20230801:getAccessPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAccessPolicyArgs struct {
	// The name of the access policy that is being added to the Redis cache.
	AccessPolicyName string `pulumi:"accessPolicyName"`
	// The name of the Redis cache.
	CacheName string `pulumi:"cacheName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Response to get/put access policy.
type LookupAccessPolicyResult struct {
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Permissions for the access policy. Learn how to configure permissions at https://aka.ms/redis/AADPreRequisites
	Permissions string `pulumi:"permissions"`
	// Provisioning state of access policy
	ProvisioningState string `pulumi:"provisioningState"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupAccessPolicyOutput(ctx *pulumi.Context, args LookupAccessPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupAccessPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAccessPolicyResult, error) {
			args := v.(LookupAccessPolicyArgs)
			r, err := LookupAccessPolicy(ctx, &args, opts...)
			var s LookupAccessPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAccessPolicyResultOutput)
}

type LookupAccessPolicyOutputArgs struct {
	// The name of the access policy that is being added to the Redis cache.
	AccessPolicyName pulumi.StringInput `pulumi:"accessPolicyName"`
	// The name of the Redis cache.
	CacheName pulumi.StringInput `pulumi:"cacheName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupAccessPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAccessPolicyArgs)(nil)).Elem()
}

// Response to get/put access policy.
type LookupAccessPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupAccessPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAccessPolicyResult)(nil)).Elem()
}

func (o LookupAccessPolicyResultOutput) ToLookupAccessPolicyResultOutput() LookupAccessPolicyResultOutput {
	return o
}

func (o LookupAccessPolicyResultOutput) ToLookupAccessPolicyResultOutputWithContext(ctx context.Context) LookupAccessPolicyResultOutput {
	return o
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupAccessPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccessPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupAccessPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccessPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

// Permissions for the access policy. Learn how to configure permissions at https://aka.ms/redis/AADPreRequisites
func (o LookupAccessPolicyResultOutput) Permissions() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccessPolicyResult) string { return v.Permissions }).(pulumi.StringOutput)
}

// Provisioning state of access policy
func (o LookupAccessPolicyResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccessPolicyResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupAccessPolicyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccessPolicyResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAccessPolicyResultOutput{})
}