// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the managementpolicy associated with the specified storage account.
func LookupManagementPolicy(ctx *pulumi.Context, args *LookupManagementPolicyArgs, opts ...pulumi.InvokeOption) (*LookupManagementPolicyResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupManagementPolicyResult
	err := ctx.Invoke("azure-native:storage/v20230101:getManagementPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupManagementPolicyArgs struct {
	// The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
	AccountName string `pulumi:"accountName"`
	// The name of the Storage Account Management Policy. It should always be 'default'
	ManagementPolicyName string `pulumi:"managementPolicyName"`
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The Get Storage Account ManagementPolicies operation response.
type LookupManagementPolicyResult struct {
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// Returns the date and time the ManagementPolicies was last modified.
	LastModifiedTime string `pulumi:"lastModifiedTime"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The Storage Account ManagementPolicy, in JSON format. See more details in: https://docs.microsoft.com/en-us/azure/storage/common/storage-lifecycle-managment-concepts.
	Policy ManagementPolicySchemaResponse `pulumi:"policy"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupManagementPolicyOutput(ctx *pulumi.Context, args LookupManagementPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupManagementPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupManagementPolicyResult, error) {
			args := v.(LookupManagementPolicyArgs)
			r, err := LookupManagementPolicy(ctx, &args, opts...)
			var s LookupManagementPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupManagementPolicyResultOutput)
}

type LookupManagementPolicyOutputArgs struct {
	// The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
	AccountName pulumi.StringInput `pulumi:"accountName"`
	// The name of the Storage Account Management Policy. It should always be 'default'
	ManagementPolicyName pulumi.StringInput `pulumi:"managementPolicyName"`
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupManagementPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagementPolicyArgs)(nil)).Elem()
}

// The Get Storage Account ManagementPolicies operation response.
type LookupManagementPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupManagementPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagementPolicyResult)(nil)).Elem()
}

func (o LookupManagementPolicyResultOutput) ToLookupManagementPolicyResultOutput() LookupManagementPolicyResultOutput {
	return o
}

func (o LookupManagementPolicyResultOutput) ToLookupManagementPolicyResultOutputWithContext(ctx context.Context) LookupManagementPolicyResultOutput {
	return o
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupManagementPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagementPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

// Returns the date and time the ManagementPolicies was last modified.
func (o LookupManagementPolicyResultOutput) LastModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagementPolicyResult) string { return v.LastModifiedTime }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupManagementPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagementPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

// The Storage Account ManagementPolicy, in JSON format. See more details in: https://docs.microsoft.com/en-us/azure/storage/common/storage-lifecycle-managment-concepts.
func (o LookupManagementPolicyResultOutput) Policy() ManagementPolicySchemaResponseOutput {
	return o.ApplyT(func(v LookupManagementPolicyResult) ManagementPolicySchemaResponse { return v.Policy }).(ManagementPolicySchemaResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupManagementPolicyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagementPolicyResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupManagementPolicyResultOutput{})
}
