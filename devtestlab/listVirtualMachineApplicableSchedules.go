// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package devtestlab

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Lists the applicable start/stop schedules, if any.
// Azure REST API version: 2018-09-15.
func ListVirtualMachineApplicableSchedules(ctx *pulumi.Context, args *ListVirtualMachineApplicableSchedulesArgs, opts ...pulumi.InvokeOption) (*ListVirtualMachineApplicableSchedulesResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListVirtualMachineApplicableSchedulesResult
	err := ctx.Invoke("azure-native:devtestlab:listVirtualMachineApplicableSchedules", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type ListVirtualMachineApplicableSchedulesArgs struct {
	// The name of the lab.
	LabName string `pulumi:"labName"`
	// The name of the virtual machine.
	Name string `pulumi:"name"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Schedules applicable to a virtual machine. The schedules may have been defined on a VM or on lab level.
type ListVirtualMachineApplicableSchedulesResult struct {
	// The identifier of the resource.
	Id string `pulumi:"id"`
	// The auto-shutdown schedule, if one has been set at the lab or lab resource level.
	LabVmsShutdown *ScheduleResponse `pulumi:"labVmsShutdown"`
	// The auto-startup schedule, if one has been set at the lab or lab resource level.
	LabVmsStartup *ScheduleResponse `pulumi:"labVmsStartup"`
	// The location of the resource.
	Location *string `pulumi:"location"`
	// The name of the resource.
	Name string `pulumi:"name"`
	// The tags of the resource.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource.
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for ListVirtualMachineApplicableSchedulesResult
func (val *ListVirtualMachineApplicableSchedulesResult) Defaults() *ListVirtualMachineApplicableSchedulesResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.LabVmsShutdown = tmp.LabVmsShutdown.Defaults()

	tmp.LabVmsStartup = tmp.LabVmsStartup.Defaults()

	return &tmp
}

func ListVirtualMachineApplicableSchedulesOutput(ctx *pulumi.Context, args ListVirtualMachineApplicableSchedulesOutputArgs, opts ...pulumi.InvokeOption) ListVirtualMachineApplicableSchedulesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListVirtualMachineApplicableSchedulesResult, error) {
			args := v.(ListVirtualMachineApplicableSchedulesArgs)
			r, err := ListVirtualMachineApplicableSchedules(ctx, &args, opts...)
			var s ListVirtualMachineApplicableSchedulesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListVirtualMachineApplicableSchedulesResultOutput)
}

type ListVirtualMachineApplicableSchedulesOutputArgs struct {
	// The name of the lab.
	LabName pulumi.StringInput `pulumi:"labName"`
	// The name of the virtual machine.
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListVirtualMachineApplicableSchedulesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListVirtualMachineApplicableSchedulesArgs)(nil)).Elem()
}

// Schedules applicable to a virtual machine. The schedules may have been defined on a VM or on lab level.
type ListVirtualMachineApplicableSchedulesResultOutput struct{ *pulumi.OutputState }

func (ListVirtualMachineApplicableSchedulesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListVirtualMachineApplicableSchedulesResult)(nil)).Elem()
}

func (o ListVirtualMachineApplicableSchedulesResultOutput) ToListVirtualMachineApplicableSchedulesResultOutput() ListVirtualMachineApplicableSchedulesResultOutput {
	return o
}

func (o ListVirtualMachineApplicableSchedulesResultOutput) ToListVirtualMachineApplicableSchedulesResultOutputWithContext(ctx context.Context) ListVirtualMachineApplicableSchedulesResultOutput {
	return o
}

func (o ListVirtualMachineApplicableSchedulesResultOutput) ToOutput(ctx context.Context) pulumix.Output[ListVirtualMachineApplicableSchedulesResult] {
	return pulumix.Output[ListVirtualMachineApplicableSchedulesResult]{
		OutputState: o.OutputState,
	}
}

// The identifier of the resource.
func (o ListVirtualMachineApplicableSchedulesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ListVirtualMachineApplicableSchedulesResult) string { return v.Id }).(pulumi.StringOutput)
}

// The auto-shutdown schedule, if one has been set at the lab or lab resource level.
func (o ListVirtualMachineApplicableSchedulesResultOutput) LabVmsShutdown() ScheduleResponsePtrOutput {
	return o.ApplyT(func(v ListVirtualMachineApplicableSchedulesResult) *ScheduleResponse { return v.LabVmsShutdown }).(ScheduleResponsePtrOutput)
}

// The auto-startup schedule, if one has been set at the lab or lab resource level.
func (o ListVirtualMachineApplicableSchedulesResultOutput) LabVmsStartup() ScheduleResponsePtrOutput {
	return o.ApplyT(func(v ListVirtualMachineApplicableSchedulesResult) *ScheduleResponse { return v.LabVmsStartup }).(ScheduleResponsePtrOutput)
}

// The location of the resource.
func (o ListVirtualMachineApplicableSchedulesResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListVirtualMachineApplicableSchedulesResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the resource.
func (o ListVirtualMachineApplicableSchedulesResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ListVirtualMachineApplicableSchedulesResult) string { return v.Name }).(pulumi.StringOutput)
}

// The tags of the resource.
func (o ListVirtualMachineApplicableSchedulesResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v ListVirtualMachineApplicableSchedulesResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource.
func (o ListVirtualMachineApplicableSchedulesResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ListVirtualMachineApplicableSchedulesResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListVirtualMachineApplicableSchedulesResultOutput{})
}
