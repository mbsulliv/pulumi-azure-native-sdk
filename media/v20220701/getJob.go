// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220701

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets a Job.
func LookupJob(ctx *pulumi.Context, args *LookupJobArgs, opts ...pulumi.InvokeOption) (*LookupJobResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupJobResult
	err := ctx.Invoke("azure-native:media/v20220701:getJob", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupJobArgs struct {
	// The Media Services account name.
	AccountName string `pulumi:"accountName"`
	// The Job name.
	JobName string `pulumi:"jobName"`
	// The name of the resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The Transform name.
	TransformName string `pulumi:"transformName"`
}

// A Job resource type. The progress and state can be obtained by polling a Job or subscribing to events using EventGrid.
type LookupJobResult struct {
	// Customer provided key, value pairs that will be returned in Job and JobOutput state events.
	CorrelationData map[string]string `pulumi:"correlationData"`
	// The UTC date and time when the customer has created the Job, in 'YYYY-MM-DDThh:mm:ssZ' format.
	Created string `pulumi:"created"`
	// Optional customer supplied description of the Job.
	Description *string `pulumi:"description"`
	// The UTC date and time at which this Job finished processing.
	EndTime string `pulumi:"endTime"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The inputs for the Job.
	Input interface{} `pulumi:"input"`
	// The UTC date and time when the customer has last updated the Job, in 'YYYY-MM-DDThh:mm:ssZ' format.
	LastModified string `pulumi:"lastModified"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The outputs for the Job.
	Outputs []JobOutputAssetResponse `pulumi:"outputs"`
	// Priority with which the job should be processed. Higher priority jobs are processed before lower priority jobs. If not set, the default is normal.
	Priority *string `pulumi:"priority"`
	// The UTC date and time at which this Job began processing.
	StartTime string `pulumi:"startTime"`
	// The current state of the job.
	State string `pulumi:"state"`
	// The system metadata relating to this resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
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
	// The Media Services account name.
	AccountName pulumi.StringInput `pulumi:"accountName"`
	// The Job name.
	JobName pulumi.StringInput `pulumi:"jobName"`
	// The name of the resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The Transform name.
	TransformName pulumi.StringInput `pulumi:"transformName"`
}

func (LookupJobOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupJobArgs)(nil)).Elem()
}

// A Job resource type. The progress and state can be obtained by polling a Job or subscribing to events using EventGrid.
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

// Customer provided key, value pairs that will be returned in Job and JobOutput state events.
func (o LookupJobResultOutput) CorrelationData() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupJobResult) map[string]string { return v.CorrelationData }).(pulumi.StringMapOutput)
}

// The UTC date and time when the customer has created the Job, in 'YYYY-MM-DDThh:mm:ssZ' format.
func (o LookupJobResultOutput) Created() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobResult) string { return v.Created }).(pulumi.StringOutput)
}

// Optional customer supplied description of the Job.
func (o LookupJobResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupJobResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The UTC date and time at which this Job finished processing.
func (o LookupJobResultOutput) EndTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobResult) string { return v.EndTime }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupJobResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobResult) string { return v.Id }).(pulumi.StringOutput)
}

// The inputs for the Job.
func (o LookupJobResultOutput) Input() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupJobResult) interface{} { return v.Input }).(pulumi.AnyOutput)
}

// The UTC date and time when the customer has last updated the Job, in 'YYYY-MM-DDThh:mm:ssZ' format.
func (o LookupJobResultOutput) LastModified() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobResult) string { return v.LastModified }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupJobResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobResult) string { return v.Name }).(pulumi.StringOutput)
}

// The outputs for the Job.
func (o LookupJobResultOutput) Outputs() JobOutputAssetResponseArrayOutput {
	return o.ApplyT(func(v LookupJobResult) []JobOutputAssetResponse { return v.Outputs }).(JobOutputAssetResponseArrayOutput)
}

// Priority with which the job should be processed. Higher priority jobs are processed before lower priority jobs. If not set, the default is normal.
func (o LookupJobResultOutput) Priority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupJobResult) *string { return v.Priority }).(pulumi.StringPtrOutput)
}

// The UTC date and time at which this Job began processing.
func (o LookupJobResultOutput) StartTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobResult) string { return v.StartTime }).(pulumi.StringOutput)
}

// The current state of the job.
func (o LookupJobResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobResult) string { return v.State }).(pulumi.StringOutput)
}

// The system metadata relating to this resource.
func (o LookupJobResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupJobResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupJobResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupJobResultOutput{})
}
