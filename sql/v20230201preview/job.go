// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230201preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A job.
type Job struct {
	pulumi.CustomResourceState

	// User-defined description of the job.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Schedule properties of the job.
	Schedule JobScheduleResponsePtrOutput `pulumi:"schedule"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// The job version number.
	Version pulumi.IntOutput `pulumi:"version"`
}

// NewJob registers a new resource with the given unique name, arguments, and options.
func NewJob(ctx *pulumi.Context,
	name string, args *JobArgs, opts ...pulumi.ResourceOption) (*Job, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.JobAgentName == nil {
		return nil, errors.New("invalid value for required argument 'JobAgentName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServerName == nil {
		return nil, errors.New("invalid value for required argument 'ServerName'")
	}
	if args.Description == nil {
		args.Description = pulumi.StringPtr("")
	}
	if args.Schedule != nil {
		args.Schedule = args.Schedule.ToJobSchedulePtrOutput().ApplyT(func(v *JobSchedule) *JobSchedule { return v.Defaults() }).(JobSchedulePtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:sql:Job"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20170301preview:Job"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200202preview:Job"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200801preview:Job"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20201101preview:Job"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210201preview:Job"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210501preview:Job"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210801preview:Job"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101:Job"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101preview:Job"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220201preview:Job"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220501preview:Job"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220801preview:Job"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20221101preview:Job"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20230501preview:Job"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Job
	err := ctx.RegisterResource("azure-native:sql/v20230201preview:Job", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetJob gets an existing Job resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetJob(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *JobState, opts ...pulumi.ResourceOption) (*Job, error) {
	var resource Job
	err := ctx.ReadResource("azure-native:sql/v20230201preview:Job", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Job resources.
type jobState struct {
}

type JobState struct {
}

func (JobState) ElementType() reflect.Type {
	return reflect.TypeOf((*jobState)(nil)).Elem()
}

type jobArgs struct {
	// User-defined description of the job.
	Description *string `pulumi:"description"`
	// The name of the job agent.
	JobAgentName string `pulumi:"jobAgentName"`
	// The name of the job to get.
	JobName *string `pulumi:"jobName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Schedule properties of the job.
	Schedule *JobSchedule `pulumi:"schedule"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
}

// The set of arguments for constructing a Job resource.
type JobArgs struct {
	// User-defined description of the job.
	Description pulumi.StringPtrInput
	// The name of the job agent.
	JobAgentName pulumi.StringInput
	// The name of the job to get.
	JobName pulumi.StringPtrInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// Schedule properties of the job.
	Schedule JobSchedulePtrInput
	// The name of the server.
	ServerName pulumi.StringInput
}

func (JobArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*jobArgs)(nil)).Elem()
}

type JobInput interface {
	pulumi.Input

	ToJobOutput() JobOutput
	ToJobOutputWithContext(ctx context.Context) JobOutput
}

func (*Job) ElementType() reflect.Type {
	return reflect.TypeOf((**Job)(nil)).Elem()
}

func (i *Job) ToJobOutput() JobOutput {
	return i.ToJobOutputWithContext(context.Background())
}

func (i *Job) ToJobOutputWithContext(ctx context.Context) JobOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobOutput)
}

type JobOutput struct{ *pulumi.OutputState }

func (JobOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Job)(nil)).Elem()
}

func (o JobOutput) ToJobOutput() JobOutput {
	return o
}

func (o JobOutput) ToJobOutputWithContext(ctx context.Context) JobOutput {
	return o
}

// User-defined description of the job.
func (o JobOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Job) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o JobOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Job) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Schedule properties of the job.
func (o JobOutput) Schedule() JobScheduleResponsePtrOutput {
	return o.ApplyT(func(v *Job) JobScheduleResponsePtrOutput { return v.Schedule }).(JobScheduleResponsePtrOutput)
}

// Resource type.
func (o JobOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Job) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The job version number.
func (o JobOutput) Version() pulumi.IntOutput {
	return o.ApplyT(func(v *Job) pulumi.IntOutput { return v.Version }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(JobOutput{})
}
