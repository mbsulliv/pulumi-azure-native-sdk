// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package storage

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get the object replication policy of the storage account by policy ID.
// Azure REST API version: 2022-09-01.
//
// Other available API versions: 2023-01-01, 2023-04-01.
func LookupObjectReplicationPolicy(ctx *pulumi.Context, args *LookupObjectReplicationPolicyArgs, opts ...pulumi.InvokeOption) (*LookupObjectReplicationPolicyResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupObjectReplicationPolicyResult
	err := ctx.Invoke("azure-native:storage:getObjectReplicationPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupObjectReplicationPolicyArgs struct {
	// The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
	AccountName string `pulumi:"accountName"`
	// For the destination account, provide the value 'default'. Configure the policy on the destination account first. For the source account, provide the value of the policy ID that is returned when you download the policy that was defined on the destination account. The policy is downloaded as a JSON file.
	ObjectReplicationPolicyId string `pulumi:"objectReplicationPolicyId"`
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The replication policy between two storage accounts. Multiple rules can be defined in one policy.
type LookupObjectReplicationPolicyResult struct {
	// Required. Destination account name. It should be full resource id if allowCrossTenantReplication set to false.
	DestinationAccount string `pulumi:"destinationAccount"`
	// Indicates when the policy is enabled on the source account.
	EnabledTime string `pulumi:"enabledTime"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// A unique id for object replication policy.
	PolicyId string `pulumi:"policyId"`
	// The storage account object replication rules.
	Rules []ObjectReplicationPolicyRuleResponse `pulumi:"rules"`
	// Required. Source account name. It should be full resource id if allowCrossTenantReplication set to false.
	SourceAccount string `pulumi:"sourceAccount"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupObjectReplicationPolicyOutput(ctx *pulumi.Context, args LookupObjectReplicationPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupObjectReplicationPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupObjectReplicationPolicyResult, error) {
			args := v.(LookupObjectReplicationPolicyArgs)
			r, err := LookupObjectReplicationPolicy(ctx, &args, opts...)
			var s LookupObjectReplicationPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupObjectReplicationPolicyResultOutput)
}

type LookupObjectReplicationPolicyOutputArgs struct {
	// The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
	AccountName pulumi.StringInput `pulumi:"accountName"`
	// For the destination account, provide the value 'default'. Configure the policy on the destination account first. For the source account, provide the value of the policy ID that is returned when you download the policy that was defined on the destination account. The policy is downloaded as a JSON file.
	ObjectReplicationPolicyId pulumi.StringInput `pulumi:"objectReplicationPolicyId"`
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupObjectReplicationPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupObjectReplicationPolicyArgs)(nil)).Elem()
}

// The replication policy between two storage accounts. Multiple rules can be defined in one policy.
type LookupObjectReplicationPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupObjectReplicationPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupObjectReplicationPolicyResult)(nil)).Elem()
}

func (o LookupObjectReplicationPolicyResultOutput) ToLookupObjectReplicationPolicyResultOutput() LookupObjectReplicationPolicyResultOutput {
	return o
}

func (o LookupObjectReplicationPolicyResultOutput) ToLookupObjectReplicationPolicyResultOutputWithContext(ctx context.Context) LookupObjectReplicationPolicyResultOutput {
	return o
}

// Required. Destination account name. It should be full resource id if allowCrossTenantReplication set to false.
func (o LookupObjectReplicationPolicyResultOutput) DestinationAccount() pulumi.StringOutput {
	return o.ApplyT(func(v LookupObjectReplicationPolicyResult) string { return v.DestinationAccount }).(pulumi.StringOutput)
}

// Indicates when the policy is enabled on the source account.
func (o LookupObjectReplicationPolicyResultOutput) EnabledTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupObjectReplicationPolicyResult) string { return v.EnabledTime }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupObjectReplicationPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupObjectReplicationPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupObjectReplicationPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupObjectReplicationPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

// A unique id for object replication policy.
func (o LookupObjectReplicationPolicyResultOutput) PolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupObjectReplicationPolicyResult) string { return v.PolicyId }).(pulumi.StringOutput)
}

// The storage account object replication rules.
func (o LookupObjectReplicationPolicyResultOutput) Rules() ObjectReplicationPolicyRuleResponseArrayOutput {
	return o.ApplyT(func(v LookupObjectReplicationPolicyResult) []ObjectReplicationPolicyRuleResponse { return v.Rules }).(ObjectReplicationPolicyRuleResponseArrayOutput)
}

// Required. Source account name. It should be full resource id if allowCrossTenantReplication set to false.
func (o LookupObjectReplicationPolicyResultOutput) SourceAccount() pulumi.StringOutput {
	return o.ApplyT(func(v LookupObjectReplicationPolicyResult) string { return v.SourceAccount }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupObjectReplicationPolicyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupObjectReplicationPolicyResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupObjectReplicationPolicyResultOutput{})
}
