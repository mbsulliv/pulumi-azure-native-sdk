// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20221101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets a job.
func LookupJob(ctx *pulumi.Context, args *LookupJobArgs, opts ...pulumi.InvokeOption) (*LookupJobResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupJobResult
	err := ctx.Invoke("azure-native:sql/v20221101preview:getJob", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupJobArgs struct {
	// The name of the job agent.
	JobAgentName string `pulumi:"jobAgentName"`
	// The name of the job to get.
	JobName string `pulumi:"jobName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
}

// A job.
type LookupJobResult struct {
	// User-defined description of the job.
	Description *string `pulumi:"description"`
	// Resource ID.
	Id string `pulumi:"id"`
	// Resource name.
	Name string `pulumi:"name"`
	// Schedule properties of the job.
	Schedule *JobScheduleResponse `pulumi:"schedule"`
	// Resource type.
	Type string `pulumi:"type"`
	// The job version number.
	Version int `pulumi:"version"`
}

// Defaults sets the appropriate defaults for LookupJobResult
func (val *LookupJobResult) Defaults() *LookupJobResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.Description == nil {
		description_ := ""
		tmp.Description = &description_
	}
	tmp.Schedule = tmp.Schedule.Defaults()

	return &tmp
}

func LookupJobOutput(ctx *pulumi.Context, args LookupJobOutputArgs, opts ...pulumi.InvokeOption) LookupJobResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupJobResult, error) {
			args := v.(LookupJobArgs)
			r, err := LookupJob(ctx, &args, opts...)
			var s LookupJobResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupJobResultOutput)
}

type LookupJobOutputArgs struct {
	// The name of the job agent.
	JobAgentName pulumi.StringInput `pulumi:"jobAgentName"`
	// The name of the job to get.
	JobName pulumi.StringInput `pulumi:"jobName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName pulumi.StringInput `pulumi:"serverName"`
}

func (LookupJobOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupJobArgs)(nil)).Elem()
}

// A job.
type LookupJobResultOutput struct{ *pulumi.OutputState }

func (LookupJobResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupJobResult)(nil)).Elem()
}

func (o LookupJobResultOutput) ToLookupJobResultOutput() LookupJobResultOutput {
	return o
}

func (o LookupJobResultOutput) ToLookupJobResultOutputWithContext(ctx context.Context) LookupJobResultOutput {
	return o
}

func (o LookupJobResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupJobResult] {
	return pulumix.Output[LookupJobResult]{
		OutputState: o.OutputState,
	}
}

// User-defined description of the job.
func (o LookupJobResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupJobResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Resource ID.
func (o LookupJobResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource name.
func (o LookupJobResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobResult) string { return v.Name }).(pulumi.StringOutput)
}

// Schedule properties of the job.
func (o LookupJobResultOutput) Schedule() JobScheduleResponsePtrOutput {
	return o.ApplyT(func(v LookupJobResult) *JobScheduleResponse { return v.Schedule }).(JobScheduleResponsePtrOutput)
}

// Resource type.
func (o LookupJobResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobResult) string { return v.Type }).(pulumi.StringOutput)
}

// The job version number.
func (o LookupJobResultOutput) Version() pulumi.IntOutput {
	return o.ApplyT(func(v LookupJobResult) int { return v.Version }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupJobResultOutput{})
}
