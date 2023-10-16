// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20231001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets a schedule resource.
func LookupSchedule(ctx *pulumi.Context, args *LookupScheduleArgs, opts ...pulumi.InvokeOption) (*LookupScheduleResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupScheduleResult
	err := ctx.Invoke("azure-native:devcenter/v20231001preview:getSchedule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupScheduleArgs struct {
	// Name of the pool.
	PoolName string `pulumi:"poolName"`
	// The name of the project.
	ProjectName string `pulumi:"projectName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the schedule that uniquely identifies it.
	ScheduleName string `pulumi:"scheduleName"`
	// The maximum number of resources to return from the operation. Example: '$top=10'.
	Top *int `pulumi:"top"`
}

// Represents a Schedule to execute a task.
type LookupScheduleResult struct {
	// The frequency of this scheduled task.
	Frequency string `pulumi:"frequency"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The provisioning state of the resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Indicates whether or not this scheduled task is enabled.
	State *string `pulumi:"state"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The target time to trigger the action. The format is HH:MM.
	Time string `pulumi:"time"`
	// The IANA timezone id at which the schedule should execute.
	TimeZone string `pulumi:"timeZone"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupScheduleOutput(ctx *pulumi.Context, args LookupScheduleOutputArgs, opts ...pulumi.InvokeOption) LookupScheduleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupScheduleResult, error) {
			args := v.(LookupScheduleArgs)
			r, err := LookupSchedule(ctx, &args, opts...)
			var s LookupScheduleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupScheduleResultOutput)
}

type LookupScheduleOutputArgs struct {
	// Name of the pool.
	PoolName pulumi.StringInput `pulumi:"poolName"`
	// The name of the project.
	ProjectName pulumi.StringInput `pulumi:"projectName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the schedule that uniquely identifies it.
	ScheduleName pulumi.StringInput `pulumi:"scheduleName"`
	// The maximum number of resources to return from the operation. Example: '$top=10'.
	Top pulumi.IntPtrInput `pulumi:"top"`
}

func (LookupScheduleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupScheduleArgs)(nil)).Elem()
}

// Represents a Schedule to execute a task.
type LookupScheduleResultOutput struct{ *pulumi.OutputState }

func (LookupScheduleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupScheduleResult)(nil)).Elem()
}

func (o LookupScheduleResultOutput) ToLookupScheduleResultOutput() LookupScheduleResultOutput {
	return o
}

func (o LookupScheduleResultOutput) ToLookupScheduleResultOutputWithContext(ctx context.Context) LookupScheduleResultOutput {
	return o
}

func (o LookupScheduleResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupScheduleResult] {
	return pulumix.Output[LookupScheduleResult]{
		OutputState: o.OutputState,
	}
}

// The frequency of this scheduled task.
func (o LookupScheduleResultOutput) Frequency() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduleResult) string { return v.Frequency }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupScheduleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduleResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupScheduleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduleResult) string { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the resource.
func (o LookupScheduleResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduleResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Indicates whether or not this scheduled task is enabled.
func (o LookupScheduleResultOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScheduleResult) *string { return v.State }).(pulumi.StringPtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupScheduleResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupScheduleResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The target time to trigger the action. The format is HH:MM.
func (o LookupScheduleResultOutput) Time() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduleResult) string { return v.Time }).(pulumi.StringOutput)
}

// The IANA timezone id at which the schedule should execute.
func (o LookupScheduleResultOutput) TimeZone() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduleResult) string { return v.TimeZone }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupScheduleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupScheduleResultOutput{})
}