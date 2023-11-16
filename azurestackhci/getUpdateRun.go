// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azurestackhci

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get the Update run for a specified update
// Azure REST API version: 2023-03-01.
//
// Other available API versions: 2022-12-15-preview, 2023-06-01, 2023-08-01, 2023-08-01-preview.
func LookupUpdateRun(ctx *pulumi.Context, args *LookupUpdateRunArgs, opts ...pulumi.InvokeOption) (*LookupUpdateRunResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupUpdateRunResult
	err := ctx.Invoke("azure-native:azurestackhci:getUpdateRun", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupUpdateRunArgs struct {
	// The name of the cluster.
	ClusterName string `pulumi:"clusterName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Update
	UpdateName string `pulumi:"updateName"`
	// The name of the Update Run
	UpdateRunName string `pulumi:"updateRunName"`
}

// Details of an Update run
type LookupUpdateRunResult struct {
	// More detailed description of the step.
	Description *string `pulumi:"description"`
	// Duration of the update run.
	Duration *string `pulumi:"duration"`
	// When the step reached a terminal state.
	EndTimeUtc *string `pulumi:"endTimeUtc"`
	// Error message, specified if the step is in a failed state.
	ErrorMessage *string `pulumi:"errorMessage"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// Timestamp of the most recently completed step in the update run.
	LastUpdatedTime *string `pulumi:"lastUpdatedTime"`
	// Completion time of this step or the last completed sub-step.
	LastUpdatedTimeUtc *string `pulumi:"lastUpdatedTimeUtc"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Provisioning state of the UpdateRuns proxy resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// When the step started, or empty if it has not started executing.
	StartTimeUtc *string `pulumi:"startTimeUtc"`
	// State of the update run.
	State *string `pulumi:"state"`
	// Status of the step, bubbled up from the ECE action plan for installation attempts. Values are: 'Success', 'Error', 'InProgress', and 'Unknown status'.
	Status *string `pulumi:"status"`
	// Recursive model for child steps of this step.
	Steps []StepResponse `pulumi:"steps"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Timestamp of the update run was started.
	TimeStarted *string `pulumi:"timeStarted"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupUpdateRunOutput(ctx *pulumi.Context, args LookupUpdateRunOutputArgs, opts ...pulumi.InvokeOption) LookupUpdateRunResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupUpdateRunResult, error) {
			args := v.(LookupUpdateRunArgs)
			r, err := LookupUpdateRun(ctx, &args, opts...)
			var s LookupUpdateRunResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupUpdateRunResultOutput)
}

type LookupUpdateRunOutputArgs struct {
	// The name of the cluster.
	ClusterName pulumi.StringInput `pulumi:"clusterName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the Update
	UpdateName pulumi.StringInput `pulumi:"updateName"`
	// The name of the Update Run
	UpdateRunName pulumi.StringInput `pulumi:"updateRunName"`
}

func (LookupUpdateRunOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUpdateRunArgs)(nil)).Elem()
}

// Details of an Update run
type LookupUpdateRunResultOutput struct{ *pulumi.OutputState }

func (LookupUpdateRunResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUpdateRunResult)(nil)).Elem()
}

func (o LookupUpdateRunResultOutput) ToLookupUpdateRunResultOutput() LookupUpdateRunResultOutput {
	return o
}

func (o LookupUpdateRunResultOutput) ToLookupUpdateRunResultOutputWithContext(ctx context.Context) LookupUpdateRunResultOutput {
	return o
}

// More detailed description of the step.
func (o LookupUpdateRunResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUpdateRunResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Duration of the update run.
func (o LookupUpdateRunResultOutput) Duration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUpdateRunResult) *string { return v.Duration }).(pulumi.StringPtrOutput)
}

// When the step reached a terminal state.
func (o LookupUpdateRunResultOutput) EndTimeUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUpdateRunResult) *string { return v.EndTimeUtc }).(pulumi.StringPtrOutput)
}

// Error message, specified if the step is in a failed state.
func (o LookupUpdateRunResultOutput) ErrorMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUpdateRunResult) *string { return v.ErrorMessage }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupUpdateRunResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUpdateRunResult) string { return v.Id }).(pulumi.StringOutput)
}

// Timestamp of the most recently completed step in the update run.
func (o LookupUpdateRunResultOutput) LastUpdatedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUpdateRunResult) *string { return v.LastUpdatedTime }).(pulumi.StringPtrOutput)
}

// Completion time of this step or the last completed sub-step.
func (o LookupUpdateRunResultOutput) LastUpdatedTimeUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUpdateRunResult) *string { return v.LastUpdatedTimeUtc }).(pulumi.StringPtrOutput)
}

// The geo-location where the resource lives
func (o LookupUpdateRunResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUpdateRunResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o LookupUpdateRunResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUpdateRunResult) string { return v.Name }).(pulumi.StringOutput)
}

// Provisioning state of the UpdateRuns proxy resource.
func (o LookupUpdateRunResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUpdateRunResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// When the step started, or empty if it has not started executing.
func (o LookupUpdateRunResultOutput) StartTimeUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUpdateRunResult) *string { return v.StartTimeUtc }).(pulumi.StringPtrOutput)
}

// State of the update run.
func (o LookupUpdateRunResultOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUpdateRunResult) *string { return v.State }).(pulumi.StringPtrOutput)
}

// Status of the step, bubbled up from the ECE action plan for installation attempts. Values are: 'Success', 'Error', 'InProgress', and 'Unknown status'.
func (o LookupUpdateRunResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUpdateRunResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

// Recursive model for child steps of this step.
func (o LookupUpdateRunResultOutput) Steps() StepResponseArrayOutput {
	return o.ApplyT(func(v LookupUpdateRunResult) []StepResponse { return v.Steps }).(StepResponseArrayOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupUpdateRunResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupUpdateRunResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Timestamp of the update run was started.
func (o LookupUpdateRunResultOutput) TimeStarted() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUpdateRunResult) *string { return v.TimeStarted }).(pulumi.StringPtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupUpdateRunResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUpdateRunResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupUpdateRunResultOutput{})
}
