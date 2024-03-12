// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20221101preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A job step.
type JobStep struct {
	pulumi.CustomResourceState

	// The action payload of the job step.
	Action JobStepActionResponseOutput `pulumi:"action"`
	// The resource ID of the job credential that will be used to connect to the targets.
	Credential pulumi.StringOutput `pulumi:"credential"`
	// Execution options for the job step.
	ExecutionOptions JobStepExecutionOptionsResponsePtrOutput `pulumi:"executionOptions"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Output destination properties of the job step.
	Output JobStepOutputResponsePtrOutput `pulumi:"output"`
	// The job step's index within the job. If not specified when creating the job step, it will be created as the last step. If not specified when updating the job step, the step id is not modified.
	StepId pulumi.IntPtrOutput `pulumi:"stepId"`
	// The resource ID of the target group that the job step will be executed on.
	TargetGroup pulumi.StringOutput `pulumi:"targetGroup"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewJobStep registers a new resource with the given unique name, arguments, and options.
func NewJobStep(ctx *pulumi.Context,
	name string, args *JobStepArgs, opts ...pulumi.ResourceOption) (*JobStep, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Action == nil {
		return nil, errors.New("invalid value for required argument 'Action'")
	}
	if args.Credential == nil {
		return nil, errors.New("invalid value for required argument 'Credential'")
	}
	if args.JobAgentName == nil {
		return nil, errors.New("invalid value for required argument 'JobAgentName'")
	}
	if args.JobName == nil {
		return nil, errors.New("invalid value for required argument 'JobName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServerName == nil {
		return nil, errors.New("invalid value for required argument 'ServerName'")
	}
	if args.TargetGroup == nil {
		return nil, errors.New("invalid value for required argument 'TargetGroup'")
	}
	args.Action = args.Action.ToJobStepActionOutput().ApplyT(func(v JobStepAction) JobStepAction { return *v.Defaults() }).(JobStepActionOutput)
	if args.ExecutionOptions != nil {
		args.ExecutionOptions = args.ExecutionOptions.ToJobStepExecutionOptionsPtrOutput().ApplyT(func(v *JobStepExecutionOptions) *JobStepExecutionOptions { return v.Defaults() }).(JobStepExecutionOptionsPtrOutput)
	}
	if args.Output != nil {
		args.Output = args.Output.ToJobStepOutputTypePtrOutput().ApplyT(func(v *JobStepOutputType) *JobStepOutputType { return v.Defaults() }).(JobStepOutputTypePtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:sql:JobStep"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20170301preview:JobStep"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200202preview:JobStep"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200801preview:JobStep"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20201101preview:JobStep"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210201preview:JobStep"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210501preview:JobStep"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210801preview:JobStep"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101:JobStep"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101preview:JobStep"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220201preview:JobStep"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220501preview:JobStep"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220801preview:JobStep"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20230201preview:JobStep"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20230501preview:JobStep"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20230801preview:JobStep"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource JobStep
	err := ctx.RegisterResource("azure-native:sql/v20221101preview:JobStep", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetJobStep gets an existing JobStep resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetJobStep(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *JobStepState, opts ...pulumi.ResourceOption) (*JobStep, error) {
	var resource JobStep
	err := ctx.ReadResource("azure-native:sql/v20221101preview:JobStep", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering JobStep resources.
type jobStepState struct {
}

type JobStepState struct {
}

func (JobStepState) ElementType() reflect.Type {
	return reflect.TypeOf((*jobStepState)(nil)).Elem()
}

type jobStepArgs struct {
	// The action payload of the job step.
	Action JobStepAction `pulumi:"action"`
	// The resource ID of the job credential that will be used to connect to the targets.
	Credential string `pulumi:"credential"`
	// Execution options for the job step.
	ExecutionOptions *JobStepExecutionOptions `pulumi:"executionOptions"`
	// The name of the job agent.
	JobAgentName string `pulumi:"jobAgentName"`
	// The name of the job.
	JobName string `pulumi:"jobName"`
	// Output destination properties of the job step.
	Output *JobStepOutputType `pulumi:"output"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
	// The job step's index within the job. If not specified when creating the job step, it will be created as the last step. If not specified when updating the job step, the step id is not modified.
	StepId *int `pulumi:"stepId"`
	// The name of the job step.
	StepName *string `pulumi:"stepName"`
	// The resource ID of the target group that the job step will be executed on.
	TargetGroup string `pulumi:"targetGroup"`
}

// The set of arguments for constructing a JobStep resource.
type JobStepArgs struct {
	// The action payload of the job step.
	Action JobStepActionInput
	// The resource ID of the job credential that will be used to connect to the targets.
	Credential pulumi.StringInput
	// Execution options for the job step.
	ExecutionOptions JobStepExecutionOptionsPtrInput
	// The name of the job agent.
	JobAgentName pulumi.StringInput
	// The name of the job.
	JobName pulumi.StringInput
	// Output destination properties of the job step.
	Output JobStepOutputTypePtrInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// The name of the server.
	ServerName pulumi.StringInput
	// The job step's index within the job. If not specified when creating the job step, it will be created as the last step. If not specified when updating the job step, the step id is not modified.
	StepId pulumi.IntPtrInput
	// The name of the job step.
	StepName pulumi.StringPtrInput
	// The resource ID of the target group that the job step will be executed on.
	TargetGroup pulumi.StringInput
}

func (JobStepArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*jobStepArgs)(nil)).Elem()
}

type JobStepInput interface {
	pulumi.Input

	ToJobStepOutput() JobStepOutput
	ToJobStepOutputWithContext(ctx context.Context) JobStepOutput
}

func (*JobStep) ElementType() reflect.Type {
	return reflect.TypeOf((**JobStep)(nil)).Elem()
}

func (i *JobStep) ToJobStepOutput() JobStepOutput {
	return i.ToJobStepOutputWithContext(context.Background())
}

func (i *JobStep) ToJobStepOutputWithContext(ctx context.Context) JobStepOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobStepOutput)
}

type JobStepOutput struct{ *pulumi.OutputState }

func (JobStepOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JobStep)(nil)).Elem()
}

func (o JobStepOutput) ToJobStepOutput() JobStepOutput {
	return o
}

func (o JobStepOutput) ToJobStepOutputWithContext(ctx context.Context) JobStepOutput {
	return o
}

// The action payload of the job step.
func (o JobStepOutput) Action() JobStepActionResponseOutput {
	return o.ApplyT(func(v *JobStep) JobStepActionResponseOutput { return v.Action }).(JobStepActionResponseOutput)
}

// The resource ID of the job credential that will be used to connect to the targets.
func (o JobStepOutput) Credential() pulumi.StringOutput {
	return o.ApplyT(func(v *JobStep) pulumi.StringOutput { return v.Credential }).(pulumi.StringOutput)
}

// Execution options for the job step.
func (o JobStepOutput) ExecutionOptions() JobStepExecutionOptionsResponsePtrOutput {
	return o.ApplyT(func(v *JobStep) JobStepExecutionOptionsResponsePtrOutput { return v.ExecutionOptions }).(JobStepExecutionOptionsResponsePtrOutput)
}

// Resource name.
func (o JobStepOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *JobStep) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Output destination properties of the job step.
func (o JobStepOutput) Output() JobStepOutputResponsePtrOutput {
	return o.ApplyT(func(v *JobStep) JobStepOutputResponsePtrOutput { return v.Output }).(JobStepOutputResponsePtrOutput)
}

// The job step's index within the job. If not specified when creating the job step, it will be created as the last step. If not specified when updating the job step, the step id is not modified.
func (o JobStepOutput) StepId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *JobStep) pulumi.IntPtrOutput { return v.StepId }).(pulumi.IntPtrOutput)
}

// The resource ID of the target group that the job step will be executed on.
func (o JobStepOutput) TargetGroup() pulumi.StringOutput {
	return o.ApplyT(func(v *JobStep) pulumi.StringOutput { return v.TargetGroup }).(pulumi.StringOutput)
}

// Resource type.
func (o JobStepOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *JobStep) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(JobStepOutput{})
}
