// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230701preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get details of the specified quota rule
func LookupVolumeQuotaRule(ctx *pulumi.Context, args *LookupVolumeQuotaRuleArgs, opts ...pulumi.InvokeOption) (*LookupVolumeQuotaRuleResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupVolumeQuotaRuleResult
	err := ctx.Invoke("azure-native:netapp/v20230701preview:getVolumeQuotaRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVolumeQuotaRuleArgs struct {
	// The name of the NetApp account
	AccountName string `pulumi:"accountName"`
	// The name of the capacity pool
	PoolName string `pulumi:"poolName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the volume
	VolumeName string `pulumi:"volumeName"`
	// The name of volume quota rule
	VolumeQuotaRuleName string `pulumi:"volumeQuotaRuleName"`
}

// Quota Rule of a Volume
type LookupVolumeQuotaRuleResult struct {
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Gets the status of the VolumeQuotaRule at the time the operation was called.
	ProvisioningState string `pulumi:"provisioningState"`
	// Size of quota
	QuotaSizeInKiBs *float64 `pulumi:"quotaSizeInKiBs"`
	// UserID/GroupID/SID based on the quota target type. UserID and groupID can be found by running ‘id’ or ‘getent’ command for the user or group and SID can be found by running <wmic useraccount where name='user-name' get sid>
	QuotaTarget *string `pulumi:"quotaTarget"`
	// Type of quota
	QuotaType *string `pulumi:"quotaType"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupVolumeQuotaRuleOutput(ctx *pulumi.Context, args LookupVolumeQuotaRuleOutputArgs, opts ...pulumi.InvokeOption) LookupVolumeQuotaRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVolumeQuotaRuleResult, error) {
			args := v.(LookupVolumeQuotaRuleArgs)
			r, err := LookupVolumeQuotaRule(ctx, &args, opts...)
			var s LookupVolumeQuotaRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVolumeQuotaRuleResultOutput)
}

type LookupVolumeQuotaRuleOutputArgs struct {
	// The name of the NetApp account
	AccountName pulumi.StringInput `pulumi:"accountName"`
	// The name of the capacity pool
	PoolName pulumi.StringInput `pulumi:"poolName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the volume
	VolumeName pulumi.StringInput `pulumi:"volumeName"`
	// The name of volume quota rule
	VolumeQuotaRuleName pulumi.StringInput `pulumi:"volumeQuotaRuleName"`
}

func (LookupVolumeQuotaRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVolumeQuotaRuleArgs)(nil)).Elem()
}

// Quota Rule of a Volume
type LookupVolumeQuotaRuleResultOutput struct{ *pulumi.OutputState }

func (LookupVolumeQuotaRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVolumeQuotaRuleResult)(nil)).Elem()
}

func (o LookupVolumeQuotaRuleResultOutput) ToLookupVolumeQuotaRuleResultOutput() LookupVolumeQuotaRuleResultOutput {
	return o
}

func (o LookupVolumeQuotaRuleResultOutput) ToLookupVolumeQuotaRuleResultOutputWithContext(ctx context.Context) LookupVolumeQuotaRuleResultOutput {
	return o
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupVolumeQuotaRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeQuotaRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupVolumeQuotaRuleResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeQuotaRuleResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupVolumeQuotaRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeQuotaRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

// Gets the status of the VolumeQuotaRule at the time the operation was called.
func (o LookupVolumeQuotaRuleResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeQuotaRuleResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Size of quota
func (o LookupVolumeQuotaRuleResultOutput) QuotaSizeInKiBs() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupVolumeQuotaRuleResult) *float64 { return v.QuotaSizeInKiBs }).(pulumi.Float64PtrOutput)
}

// UserID/GroupID/SID based on the quota target type. UserID and groupID can be found by running ‘id’ or ‘getent’ command for the user or group and SID can be found by running <wmic useraccount where name='user-name' get sid>
func (o LookupVolumeQuotaRuleResultOutput) QuotaTarget() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVolumeQuotaRuleResult) *string { return v.QuotaTarget }).(pulumi.StringPtrOutput)
}

// Type of quota
func (o LookupVolumeQuotaRuleResultOutput) QuotaType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVolumeQuotaRuleResult) *string { return v.QuotaType }).(pulumi.StringPtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupVolumeQuotaRuleResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupVolumeQuotaRuleResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupVolumeQuotaRuleResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupVolumeQuotaRuleResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupVolumeQuotaRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeQuotaRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVolumeQuotaRuleResultOutput{})
}
