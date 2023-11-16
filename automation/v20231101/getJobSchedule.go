// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20231101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieve the job schedule identified by job schedule name.
func LookupJobSchedule(ctx *pulumi.Context, args *LookupJobScheduleArgs, opts ...pulumi.InvokeOption) (*LookupJobScheduleResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupJobScheduleResult
	err := ctx.Invoke("azure-native:automation/v20231101:getJobSchedule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupJobScheduleArgs struct {
	// The name of the automation account.
	AutomationAccountName string `pulumi:"automationAccountName"`
	// The job schedule name.
	JobScheduleId string `pulumi:"jobScheduleId"`
	// Name of an Azure Resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Definition of the job schedule.
type LookupJobScheduleResult struct {
	// Gets the id of the resource.
	Id string `pulumi:"id"`
	// Gets or sets the id of job schedule.
	JobScheduleId *string `pulumi:"jobScheduleId"`
	// Gets the name of the variable.
	Name string `pulumi:"name"`
	// Gets or sets the parameters of the job schedule.
	Parameters map[string]string `pulumi:"parameters"`
	// Gets or sets the hybrid worker group that the scheduled job should run on.
	RunOn *string `pulumi:"runOn"`
	// Gets or sets the runbook.
	Runbook *RunbookAssociationPropertyResponse `pulumi:"runbook"`
	// Gets or sets the schedule.
	Schedule *ScheduleAssociationPropertyResponse `pulumi:"schedule"`
	// Resource type
	Type string `pulumi:"type"`
}

func LookupJobScheduleOutput(ctx *pulumi.Context, args LookupJobScheduleOutputArgs, opts ...pulumi.InvokeOption) LookupJobScheduleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupJobScheduleResult, error) {
			args := v.(LookupJobScheduleArgs)
			r, err := LookupJobSchedule(ctx, &args, opts...)
			var s LookupJobScheduleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupJobScheduleResultOutput)
}

type LookupJobScheduleOutputArgs struct {
	// The name of the automation account.
	AutomationAccountName pulumi.StringInput `pulumi:"automationAccountName"`
	// The job schedule name.
	JobScheduleId pulumi.StringInput `pulumi:"jobScheduleId"`
	// Name of an Azure Resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupJobScheduleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupJobScheduleArgs)(nil)).Elem()
}

// Definition of the job schedule.
type LookupJobScheduleResultOutput struct{ *pulumi.OutputState }

func (LookupJobScheduleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupJobScheduleResult)(nil)).Elem()
}

func (o LookupJobScheduleResultOutput) ToLookupJobScheduleResultOutput() LookupJobScheduleResultOutput {
	return o
}

func (o LookupJobScheduleResultOutput) ToLookupJobScheduleResultOutputWithContext(ctx context.Context) LookupJobScheduleResultOutput {
	return o
}

// Gets the id of the resource.
func (o LookupJobScheduleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobScheduleResult) string { return v.Id }).(pulumi.StringOutput)
}

// Gets or sets the id of job schedule.
func (o LookupJobScheduleResultOutput) JobScheduleId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupJobScheduleResult) *string { return v.JobScheduleId }).(pulumi.StringPtrOutput)
}

// Gets the name of the variable.
func (o LookupJobScheduleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobScheduleResult) string { return v.Name }).(pulumi.StringOutput)
}

// Gets or sets the parameters of the job schedule.
func (o LookupJobScheduleResultOutput) Parameters() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupJobScheduleResult) map[string]string { return v.Parameters }).(pulumi.StringMapOutput)
}

// Gets or sets the hybrid worker group that the scheduled job should run on.
func (o LookupJobScheduleResultOutput) RunOn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupJobScheduleResult) *string { return v.RunOn }).(pulumi.StringPtrOutput)
}

// Gets or sets the runbook.
func (o LookupJobScheduleResultOutput) Runbook() RunbookAssociationPropertyResponsePtrOutput {
	return o.ApplyT(func(v LookupJobScheduleResult) *RunbookAssociationPropertyResponse { return v.Runbook }).(RunbookAssociationPropertyResponsePtrOutput)
}

// Gets or sets the schedule.
func (o LookupJobScheduleResultOutput) Schedule() ScheduleAssociationPropertyResponsePtrOutput {
	return o.ApplyT(func(v LookupJobScheduleResult) *ScheduleAssociationPropertyResponse { return v.Schedule }).(ScheduleAssociationPropertyResponsePtrOutput)
}

// Resource type
func (o LookupJobScheduleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobScheduleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupJobScheduleResultOutput{})
}
