// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20221101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get a snapshot Policy
func LookupSnapshotPolicy(ctx *pulumi.Context, args *LookupSnapshotPolicyArgs, opts ...pulumi.InvokeOption) (*LookupSnapshotPolicyResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupSnapshotPolicyResult
	err := ctx.Invoke("azure-native:netapp/v20221101:getSnapshotPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSnapshotPolicyArgs struct {
	// The name of the NetApp account
	AccountName string `pulumi:"accountName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the snapshot policy
	SnapshotPolicyName string `pulumi:"snapshotPolicyName"`
}

// Snapshot policy information
type LookupSnapshotPolicyResult struct {
	// Schedule for daily snapshots
	DailySchedule *DailyScheduleResponse `pulumi:"dailySchedule"`
	// The property to decide policy is enabled or not
	Enabled *bool `pulumi:"enabled"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// Schedule for hourly snapshots
	HourlySchedule *HourlyScheduleResponse `pulumi:"hourlySchedule"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// Schedule for monthly snapshots
	MonthlySchedule *MonthlyScheduleResponse `pulumi:"monthlySchedule"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Azure lifecycle management
	ProvisioningState string `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// Schedule for weekly snapshots
	WeeklySchedule *WeeklyScheduleResponse `pulumi:"weeklySchedule"`
}

func LookupSnapshotPolicyOutput(ctx *pulumi.Context, args LookupSnapshotPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupSnapshotPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSnapshotPolicyResult, error) {
			args := v.(LookupSnapshotPolicyArgs)
			r, err := LookupSnapshotPolicy(ctx, &args, opts...)
			var s LookupSnapshotPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSnapshotPolicyResultOutput)
}

type LookupSnapshotPolicyOutputArgs struct {
	// The name of the NetApp account
	AccountName pulumi.StringInput `pulumi:"accountName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the snapshot policy
	SnapshotPolicyName pulumi.StringInput `pulumi:"snapshotPolicyName"`
}

func (LookupSnapshotPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSnapshotPolicyArgs)(nil)).Elem()
}

// Snapshot policy information
type LookupSnapshotPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupSnapshotPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSnapshotPolicyResult)(nil)).Elem()
}

func (o LookupSnapshotPolicyResultOutput) ToLookupSnapshotPolicyResultOutput() LookupSnapshotPolicyResultOutput {
	return o
}

func (o LookupSnapshotPolicyResultOutput) ToLookupSnapshotPolicyResultOutputWithContext(ctx context.Context) LookupSnapshotPolicyResultOutput {
	return o
}

// Schedule for daily snapshots
func (o LookupSnapshotPolicyResultOutput) DailySchedule() DailyScheduleResponsePtrOutput {
	return o.ApplyT(func(v LookupSnapshotPolicyResult) *DailyScheduleResponse { return v.DailySchedule }).(DailyScheduleResponsePtrOutput)
}

// The property to decide policy is enabled or not
func (o LookupSnapshotPolicyResultOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupSnapshotPolicyResult) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o LookupSnapshotPolicyResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSnapshotPolicyResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Schedule for hourly snapshots
func (o LookupSnapshotPolicyResultOutput) HourlySchedule() HourlyScheduleResponsePtrOutput {
	return o.ApplyT(func(v LookupSnapshotPolicyResult) *HourlyScheduleResponse { return v.HourlySchedule }).(HourlyScheduleResponsePtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupSnapshotPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSnapshotPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupSnapshotPolicyResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSnapshotPolicyResult) string { return v.Location }).(pulumi.StringOutput)
}

// Schedule for monthly snapshots
func (o LookupSnapshotPolicyResultOutput) MonthlySchedule() MonthlyScheduleResponsePtrOutput {
	return o.ApplyT(func(v LookupSnapshotPolicyResult) *MonthlyScheduleResponse { return v.MonthlySchedule }).(MonthlyScheduleResponsePtrOutput)
}

// The name of the resource
func (o LookupSnapshotPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSnapshotPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

// Azure lifecycle management
func (o LookupSnapshotPolicyResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSnapshotPolicyResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupSnapshotPolicyResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupSnapshotPolicyResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupSnapshotPolicyResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSnapshotPolicyResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupSnapshotPolicyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSnapshotPolicyResult) string { return v.Type }).(pulumi.StringOutput)
}

// Schedule for weekly snapshots
func (o LookupSnapshotPolicyResultOutput) WeeklySchedule() WeeklyScheduleResponsePtrOutput {
	return o.ApplyT(func(v LookupSnapshotPolicyResult) *WeeklyScheduleResponse { return v.WeeklySchedule }).(WeeklyScheduleResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSnapshotPolicyResultOutput{})
}
