// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package quota

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the GroupQuotas for the name passed. It will return the GroupQuotas properties only. The details on group quota can be access from the group quota APIs.
// Azure REST API version: 2023-06-01-preview.
func LookupGroupQuota(ctx *pulumi.Context, args *LookupGroupQuotaArgs, opts ...pulumi.InvokeOption) (*LookupGroupQuotaResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupGroupQuotaResult
	err := ctx.Invoke("azure-native:quota:getGroupQuota", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGroupQuotaArgs struct {
	// The GroupQuota name. The name should be unique for the provided context tenantId/MgId.
	GroupQuotaName string `pulumi:"groupQuotaName"`
	// Management Group Id.
	ManagementGroupId string `pulumi:"managementGroupId"`
}

// Properties and filters for ShareQuota. The request parameter is optional, if there are no filters specified.
type LookupGroupQuotaResult struct {
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Properties and filters for ShareQuota. The request parameter is optional, if there are no filters specified.
	Properties GroupQuotasEntityBaseResponse `pulumi:"properties"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupGroupQuotaOutput(ctx *pulumi.Context, args LookupGroupQuotaOutputArgs, opts ...pulumi.InvokeOption) LookupGroupQuotaResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGroupQuotaResult, error) {
			args := v.(LookupGroupQuotaArgs)
			r, err := LookupGroupQuota(ctx, &args, opts...)
			var s LookupGroupQuotaResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGroupQuotaResultOutput)
}

type LookupGroupQuotaOutputArgs struct {
	// The GroupQuota name. The name should be unique for the provided context tenantId/MgId.
	GroupQuotaName pulumi.StringInput `pulumi:"groupQuotaName"`
	// Management Group Id.
	ManagementGroupId pulumi.StringInput `pulumi:"managementGroupId"`
}

func (LookupGroupQuotaOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGroupQuotaArgs)(nil)).Elem()
}

// Properties and filters for ShareQuota. The request parameter is optional, if there are no filters specified.
type LookupGroupQuotaResultOutput struct{ *pulumi.OutputState }

func (LookupGroupQuotaResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGroupQuotaResult)(nil)).Elem()
}

func (o LookupGroupQuotaResultOutput) ToLookupGroupQuotaResultOutput() LookupGroupQuotaResultOutput {
	return o
}

func (o LookupGroupQuotaResultOutput) ToLookupGroupQuotaResultOutputWithContext(ctx context.Context) LookupGroupQuotaResultOutput {
	return o
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupGroupQuotaResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupQuotaResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupGroupQuotaResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupQuotaResult) string { return v.Name }).(pulumi.StringOutput)
}

// Properties and filters for ShareQuota. The request parameter is optional, if there are no filters specified.
func (o LookupGroupQuotaResultOutput) Properties() GroupQuotasEntityBaseResponseOutput {
	return o.ApplyT(func(v LookupGroupQuotaResult) GroupQuotasEntityBaseResponse { return v.Properties }).(GroupQuotasEntityBaseResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupGroupQuotaResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupGroupQuotaResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupGroupQuotaResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupQuotaResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGroupQuotaResultOutput{})
}
