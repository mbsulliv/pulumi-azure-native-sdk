// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20221215preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Details of an Update run
type UpdateRun struct {
	pulumi.CustomResourceState

	// More detailed description of the step.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Duration of the update run.
	Duration pulumi.StringPtrOutput `pulumi:"duration"`
	// When the step reached a terminal state.
	EndTimeUtc pulumi.StringPtrOutput `pulumi:"endTimeUtc"`
	// Error message, specified if the step is in a failed state.
	ErrorMessage pulumi.StringPtrOutput `pulumi:"errorMessage"`
	// Expected execution time of a given step. This is optionally authored in the update action plan and can be empty.
	ExpectedExecutionTime pulumi.StringPtrOutput `pulumi:"expectedExecutionTime"`
	// Timestamp of the most recently completed step in the update run.
	LastUpdatedTime pulumi.StringPtrOutput `pulumi:"lastUpdatedTime"`
	// Completion time of this step or the last completed sub-step.
	LastUpdatedTimeUtc pulumi.StringPtrOutput `pulumi:"lastUpdatedTimeUtc"`
	// The geo-location where the resource lives
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Provisioning state of the UpdateRuns proxy resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// When the step started, or empty if it has not started executing.
	StartTimeUtc pulumi.StringPtrOutput `pulumi:"startTimeUtc"`
	// State of the update run.
	State pulumi.StringPtrOutput `pulumi:"state"`
	// Status of the step, bubbled up from the ECE action plan for installation attempts. Values are: 'Success', 'Error', 'InProgress', and 'Unknown status'.
	Status pulumi.StringPtrOutput `pulumi:"status"`
	// Recursive model for child steps of this step.
	Steps StepResponseArrayOutput `pulumi:"steps"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Timestamp of the update run was started.
	TimeStarted pulumi.StringPtrOutput `pulumi:"timeStarted"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewUpdateRun registers a new resource with the given unique name, arguments, and options.
func NewUpdateRun(ctx *pulumi.Context,
	name string, args *UpdateRunArgs, opts ...pulumi.ResourceOption) (*UpdateRun, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterName == nil {
		return nil, errors.New("invalid value for required argument 'ClusterName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.UpdateName == nil {
		return nil, errors.New("invalid value for required argument 'UpdateName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:azurestackhci:UpdateRun"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20221201:UpdateRun"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20230201:UpdateRun"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20230301:UpdateRun"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20230601:UpdateRun"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20230801:UpdateRun"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20230801preview:UpdateRun"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20231101preview:UpdateRun"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20240101:UpdateRun"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20240215preview:UpdateRun"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource UpdateRun
	err := ctx.RegisterResource("azure-native:azurestackhci/v20221215preview:UpdateRun", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUpdateRun gets an existing UpdateRun resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUpdateRun(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UpdateRunState, opts ...pulumi.ResourceOption) (*UpdateRun, error) {
	var resource UpdateRun
	err := ctx.ReadResource("azure-native:azurestackhci/v20221215preview:UpdateRun", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UpdateRun resources.
type updateRunState struct {
}

type UpdateRunState struct {
}

func (UpdateRunState) ElementType() reflect.Type {
	return reflect.TypeOf((*updateRunState)(nil)).Elem()
}

type updateRunArgs struct {
	// The name of the cluster.
	ClusterName string `pulumi:"clusterName"`
	// More detailed description of the step.
	Description *string `pulumi:"description"`
	// Duration of the update run.
	Duration *string `pulumi:"duration"`
	// When the step reached a terminal state.
	EndTimeUtc *string `pulumi:"endTimeUtc"`
	// Error message, specified if the step is in a failed state.
	ErrorMessage *string `pulumi:"errorMessage"`
	// Expected execution time of a given step. This is optionally authored in the update action plan and can be empty.
	ExpectedExecutionTime *string `pulumi:"expectedExecutionTime"`
	// Timestamp of the most recently completed step in the update run.
	LastUpdatedTime *string `pulumi:"lastUpdatedTime"`
	// Completion time of this step or the last completed sub-step.
	LastUpdatedTimeUtc *string `pulumi:"lastUpdatedTimeUtc"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// Name of the step.
	Name *string `pulumi:"name"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// When the step started, or empty if it has not started executing.
	StartTimeUtc *string `pulumi:"startTimeUtc"`
	// State of the update run.
	State *string `pulumi:"state"`
	// Status of the step, bubbled up from the ECE action plan for installation attempts. Values are: 'Success', 'Error', 'InProgress', and 'Unknown status'.
	Status *string `pulumi:"status"`
	// Recursive model for child steps of this step.
	Steps []Step `pulumi:"steps"`
	// Timestamp of the update run was started.
	TimeStarted *string `pulumi:"timeStarted"`
	// The name of the Update
	UpdateName string `pulumi:"updateName"`
	// The name of the Update Run
	UpdateRunName *string `pulumi:"updateRunName"`
}

// The set of arguments for constructing a UpdateRun resource.
type UpdateRunArgs struct {
	// The name of the cluster.
	ClusterName pulumi.StringInput
	// More detailed description of the step.
	Description pulumi.StringPtrInput
	// Duration of the update run.
	Duration pulumi.StringPtrInput
	// When the step reached a terminal state.
	EndTimeUtc pulumi.StringPtrInput
	// Error message, specified if the step is in a failed state.
	ErrorMessage pulumi.StringPtrInput
	// Expected execution time of a given step. This is optionally authored in the update action plan and can be empty.
	ExpectedExecutionTime pulumi.StringPtrInput
	// Timestamp of the most recently completed step in the update run.
	LastUpdatedTime pulumi.StringPtrInput
	// Completion time of this step or the last completed sub-step.
	LastUpdatedTimeUtc pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// Name of the step.
	Name pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// When the step started, or empty if it has not started executing.
	StartTimeUtc pulumi.StringPtrInput
	// State of the update run.
	State pulumi.StringPtrInput
	// Status of the step, bubbled up from the ECE action plan for installation attempts. Values are: 'Success', 'Error', 'InProgress', and 'Unknown status'.
	Status pulumi.StringPtrInput
	// Recursive model for child steps of this step.
	Steps StepArrayInput
	// Timestamp of the update run was started.
	TimeStarted pulumi.StringPtrInput
	// The name of the Update
	UpdateName pulumi.StringInput
	// The name of the Update Run
	UpdateRunName pulumi.StringPtrInput
}

func (UpdateRunArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*updateRunArgs)(nil)).Elem()
}

type UpdateRunInput interface {
	pulumi.Input

	ToUpdateRunOutput() UpdateRunOutput
	ToUpdateRunOutputWithContext(ctx context.Context) UpdateRunOutput
}

func (*UpdateRun) ElementType() reflect.Type {
	return reflect.TypeOf((**UpdateRun)(nil)).Elem()
}

func (i *UpdateRun) ToUpdateRunOutput() UpdateRunOutput {
	return i.ToUpdateRunOutputWithContext(context.Background())
}

func (i *UpdateRun) ToUpdateRunOutputWithContext(ctx context.Context) UpdateRunOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UpdateRunOutput)
}

type UpdateRunOutput struct{ *pulumi.OutputState }

func (UpdateRunOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UpdateRun)(nil)).Elem()
}

func (o UpdateRunOutput) ToUpdateRunOutput() UpdateRunOutput {
	return o
}

func (o UpdateRunOutput) ToUpdateRunOutputWithContext(ctx context.Context) UpdateRunOutput {
	return o
}

// More detailed description of the step.
func (o UpdateRunOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UpdateRun) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Duration of the update run.
func (o UpdateRunOutput) Duration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UpdateRun) pulumi.StringPtrOutput { return v.Duration }).(pulumi.StringPtrOutput)
}

// When the step reached a terminal state.
func (o UpdateRunOutput) EndTimeUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UpdateRun) pulumi.StringPtrOutput { return v.EndTimeUtc }).(pulumi.StringPtrOutput)
}

// Error message, specified if the step is in a failed state.
func (o UpdateRunOutput) ErrorMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UpdateRun) pulumi.StringPtrOutput { return v.ErrorMessage }).(pulumi.StringPtrOutput)
}

// Expected execution time of a given step. This is optionally authored in the update action plan and can be empty.
func (o UpdateRunOutput) ExpectedExecutionTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UpdateRun) pulumi.StringPtrOutput { return v.ExpectedExecutionTime }).(pulumi.StringPtrOutput)
}

// Timestamp of the most recently completed step in the update run.
func (o UpdateRunOutput) LastUpdatedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UpdateRun) pulumi.StringPtrOutput { return v.LastUpdatedTime }).(pulumi.StringPtrOutput)
}

// Completion time of this step or the last completed sub-step.
func (o UpdateRunOutput) LastUpdatedTimeUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UpdateRun) pulumi.StringPtrOutput { return v.LastUpdatedTimeUtc }).(pulumi.StringPtrOutput)
}

// The geo-location where the resource lives
func (o UpdateRunOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UpdateRun) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o UpdateRunOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *UpdateRun) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Provisioning state of the UpdateRuns proxy resource.
func (o UpdateRunOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *UpdateRun) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// When the step started, or empty if it has not started executing.
func (o UpdateRunOutput) StartTimeUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UpdateRun) pulumi.StringPtrOutput { return v.StartTimeUtc }).(pulumi.StringPtrOutput)
}

// State of the update run.
func (o UpdateRunOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UpdateRun) pulumi.StringPtrOutput { return v.State }).(pulumi.StringPtrOutput)
}

// Status of the step, bubbled up from the ECE action plan for installation attempts. Values are: 'Success', 'Error', 'InProgress', and 'Unknown status'.
func (o UpdateRunOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UpdateRun) pulumi.StringPtrOutput { return v.Status }).(pulumi.StringPtrOutput)
}

// Recursive model for child steps of this step.
func (o UpdateRunOutput) Steps() StepResponseArrayOutput {
	return o.ApplyT(func(v *UpdateRun) StepResponseArrayOutput { return v.Steps }).(StepResponseArrayOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o UpdateRunOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *UpdateRun) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Timestamp of the update run was started.
func (o UpdateRunOutput) TimeStarted() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UpdateRun) pulumi.StringPtrOutput { return v.TimeStarted }).(pulumi.StringPtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o UpdateRunOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *UpdateRun) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(UpdateRunOutput{})
}
