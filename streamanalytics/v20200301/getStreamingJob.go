// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20200301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets details about the specified streaming job.
func LookupStreamingJob(ctx *pulumi.Context, args *LookupStreamingJobArgs, opts ...pulumi.InvokeOption) (*LookupStreamingJobResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupStreamingJobResult
	err := ctx.Invoke("azure-native:streamanalytics/v20200301:getStreamingJob", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupStreamingJobArgs struct {
	// The $expand OData query parameter. This is a comma-separated list of additional streaming job properties to include in the response, beyond the default set returned when this parameter is absent. The default set is all streaming job properties other than 'inputs', 'transformation', 'outputs', and 'functions'.
	Expand *string `pulumi:"expand"`
	// The name of the streaming job.
	JobName string `pulumi:"jobName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A streaming job object, containing all information associated with the named streaming job.
type LookupStreamingJobResult struct {
	// The cluster which streaming jobs will run on.
	Cluster *ClusterInfoResponse `pulumi:"cluster"`
	// Controls certain runtime behaviors of the streaming job.
	CompatibilityLevel *string `pulumi:"compatibilityLevel"`
	// Valid values are JobStorageAccount and SystemAccount. If set to JobStorageAccount, this requires the user to also specify jobStorageAccount property. .
	ContentStoragePolicy *string `pulumi:"contentStoragePolicy"`
	// Value is an ISO-8601 formatted UTC timestamp indicating when the streaming job was created.
	CreatedDate string `pulumi:"createdDate"`
	// The data locale of the stream analytics job. Value should be the name of a supported .NET Culture from the set https://msdn.microsoft.com/en-us/library/system.globalization.culturetypes(v=vs.110).aspx. Defaults to 'en-US' if none specified.
	DataLocale *string `pulumi:"dataLocale"`
	// The current entity tag for the streaming job. This is an opaque string. You can use it to detect whether the resource has changed between requests. You can also use it in the If-Match or If-None-Match headers for write operations for optimistic concurrency.
	Etag string `pulumi:"etag"`
	// The maximum tolerable delay in seconds where events arriving late could be included.  Supported range is -1 to 1814399 (20.23:59:59 days) and -1 is used to specify wait indefinitely. If the property is absent, it is interpreted to have a value of -1.
	EventsLateArrivalMaxDelayInSeconds *int `pulumi:"eventsLateArrivalMaxDelayInSeconds"`
	// The maximum tolerable delay in seconds where out-of-order events can be adjusted to be back in order.
	EventsOutOfOrderMaxDelayInSeconds *int `pulumi:"eventsOutOfOrderMaxDelayInSeconds"`
	// Indicates the policy to apply to events that arrive out of order in the input event stream.
	EventsOutOfOrderPolicy *string `pulumi:"eventsOutOfOrderPolicy"`
	// A list of one or more functions for the streaming job. The name property for each function is required when specifying this property in a PUT request. This property cannot be modify via a PATCH operation. You must use the PATCH API available for the individual transformation.
	Functions []FunctionResponse `pulumi:"functions"`
	// Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// Describes the system-assigned managed identity assigned to this job that can be used to authenticate with inputs and outputs.
	Identity *IdentityResponse `pulumi:"identity"`
	// A list of one or more inputs to the streaming job. The name property for each input is required when specifying this property in a PUT request. This property cannot be modify via a PATCH operation. You must use the PATCH API available for the individual input.
	Inputs []InputResponse `pulumi:"inputs"`
	// A GUID uniquely identifying the streaming job. This GUID is generated upon creation of the streaming job.
	JobId string `pulumi:"jobId"`
	// Describes the state of the streaming job.
	JobState string `pulumi:"jobState"`
	// The properties that are associated with an Azure Storage account with MSI
	JobStorageAccount *JobStorageAccountResponse `pulumi:"jobStorageAccount"`
	// Describes the type of the job. Valid modes are `Cloud` and 'Edge'.
	JobType *string `pulumi:"jobType"`
	// Value is either an ISO-8601 formatted timestamp indicating the last output event time of the streaming job or null indicating that output has not yet been produced. In case of multiple outputs or multiple streams, this shows the latest value in that set.
	LastOutputEventTime string `pulumi:"lastOutputEventTime"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Indicates the policy to apply to events that arrive at the output and cannot be written to the external storage due to being malformed (missing column values, column values of wrong type or size).
	OutputErrorPolicy *string `pulumi:"outputErrorPolicy"`
	// This property should only be utilized when it is desired that the job be started immediately upon creation. Value may be JobStartTime, CustomTime, or LastOutputEventTime to indicate whether the starting point of the output event stream should start whenever the job is started, start at a custom user time stamp specified via the outputStartTime property, or start from the last event output time.
	OutputStartMode *string `pulumi:"outputStartMode"`
	// Value is either an ISO-8601 formatted time stamp that indicates the starting point of the output event stream, or null to indicate that the output event stream will start whenever the streaming job is started. This property must have a value if outputStartMode is set to CustomTime.
	OutputStartTime *string `pulumi:"outputStartTime"`
	// A list of one or more outputs for the streaming job. The name property for each output is required when specifying this property in a PUT request. This property cannot be modify via a PATCH operation. You must use the PATCH API available for the individual output.
	Outputs []OutputResponse `pulumi:"outputs"`
	// Describes the provisioning status of the streaming job.
	ProvisioningState string `pulumi:"provisioningState"`
	// Describes the SKU of the streaming job. Required on PUT (CreateOrReplace) requests.
	Sku *SkuResponse `pulumi:"sku"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Indicates the query and the number of streaming units to use for the streaming job. The name property of the transformation is required when specifying this property in a PUT request. This property cannot be modify via a PATCH operation. You must use the PATCH API available for the individual transformation.
	Transformation *TransformationResponse `pulumi:"transformation"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupStreamingJobResult
func (val *LookupStreamingJobResult) Defaults() *LookupStreamingJobResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Transformation = tmp.Transformation.Defaults()

	return &tmp
}

func LookupStreamingJobOutput(ctx *pulumi.Context, args LookupStreamingJobOutputArgs, opts ...pulumi.InvokeOption) LookupStreamingJobResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupStreamingJobResult, error) {
			args := v.(LookupStreamingJobArgs)
			r, err := LookupStreamingJob(ctx, &args, opts...)
			var s LookupStreamingJobResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupStreamingJobResultOutput)
}

type LookupStreamingJobOutputArgs struct {
	// The $expand OData query parameter. This is a comma-separated list of additional streaming job properties to include in the response, beyond the default set returned when this parameter is absent. The default set is all streaming job properties other than 'inputs', 'transformation', 'outputs', and 'functions'.
	Expand pulumi.StringPtrInput `pulumi:"expand"`
	// The name of the streaming job.
	JobName pulumi.StringInput `pulumi:"jobName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupStreamingJobOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStreamingJobArgs)(nil)).Elem()
}

// A streaming job object, containing all information associated with the named streaming job.
type LookupStreamingJobResultOutput struct{ *pulumi.OutputState }

func (LookupStreamingJobResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStreamingJobResult)(nil)).Elem()
}

func (o LookupStreamingJobResultOutput) ToLookupStreamingJobResultOutput() LookupStreamingJobResultOutput {
	return o
}

func (o LookupStreamingJobResultOutput) ToLookupStreamingJobResultOutputWithContext(ctx context.Context) LookupStreamingJobResultOutput {
	return o
}

func (o LookupStreamingJobResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupStreamingJobResult] {
	return pulumix.Output[LookupStreamingJobResult]{
		OutputState: o.OutputState,
	}
}

// The cluster which streaming jobs will run on.
func (o LookupStreamingJobResultOutput) Cluster() ClusterInfoResponsePtrOutput {
	return o.ApplyT(func(v LookupStreamingJobResult) *ClusterInfoResponse { return v.Cluster }).(ClusterInfoResponsePtrOutput)
}

// Controls certain runtime behaviors of the streaming job.
func (o LookupStreamingJobResultOutput) CompatibilityLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStreamingJobResult) *string { return v.CompatibilityLevel }).(pulumi.StringPtrOutput)
}

// Valid values are JobStorageAccount and SystemAccount. If set to JobStorageAccount, this requires the user to also specify jobStorageAccount property. .
func (o LookupStreamingJobResultOutput) ContentStoragePolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStreamingJobResult) *string { return v.ContentStoragePolicy }).(pulumi.StringPtrOutput)
}

// Value is an ISO-8601 formatted UTC timestamp indicating when the streaming job was created.
func (o LookupStreamingJobResultOutput) CreatedDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStreamingJobResult) string { return v.CreatedDate }).(pulumi.StringOutput)
}

// The data locale of the stream analytics job. Value should be the name of a supported .NET Culture from the set https://msdn.microsoft.com/en-us/library/system.globalization.culturetypes(v=vs.110).aspx. Defaults to 'en-US' if none specified.
func (o LookupStreamingJobResultOutput) DataLocale() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStreamingJobResult) *string { return v.DataLocale }).(pulumi.StringPtrOutput)
}

// The current entity tag for the streaming job. This is an opaque string. You can use it to detect whether the resource has changed between requests. You can also use it in the If-Match or If-None-Match headers for write operations for optimistic concurrency.
func (o LookupStreamingJobResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStreamingJobResult) string { return v.Etag }).(pulumi.StringOutput)
}

// The maximum tolerable delay in seconds where events arriving late could be included.  Supported range is -1 to 1814399 (20.23:59:59 days) and -1 is used to specify wait indefinitely. If the property is absent, it is interpreted to have a value of -1.
func (o LookupStreamingJobResultOutput) EventsLateArrivalMaxDelayInSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupStreamingJobResult) *int { return v.EventsLateArrivalMaxDelayInSeconds }).(pulumi.IntPtrOutput)
}

// The maximum tolerable delay in seconds where out-of-order events can be adjusted to be back in order.
func (o LookupStreamingJobResultOutput) EventsOutOfOrderMaxDelayInSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupStreamingJobResult) *int { return v.EventsOutOfOrderMaxDelayInSeconds }).(pulumi.IntPtrOutput)
}

// Indicates the policy to apply to events that arrive out of order in the input event stream.
func (o LookupStreamingJobResultOutput) EventsOutOfOrderPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStreamingJobResult) *string { return v.EventsOutOfOrderPolicy }).(pulumi.StringPtrOutput)
}

// A list of one or more functions for the streaming job. The name property for each function is required when specifying this property in a PUT request. This property cannot be modify via a PATCH operation. You must use the PATCH API available for the individual transformation.
func (o LookupStreamingJobResultOutput) Functions() FunctionResponseArrayOutput {
	return o.ApplyT(func(v LookupStreamingJobResult) []FunctionResponse { return v.Functions }).(FunctionResponseArrayOutput)
}

// Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupStreamingJobResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStreamingJobResult) string { return v.Id }).(pulumi.StringOutput)
}

// Describes the system-assigned managed identity assigned to this job that can be used to authenticate with inputs and outputs.
func (o LookupStreamingJobResultOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupStreamingJobResult) *IdentityResponse { return v.Identity }).(IdentityResponsePtrOutput)
}

// A list of one or more inputs to the streaming job. The name property for each input is required when specifying this property in a PUT request. This property cannot be modify via a PATCH operation. You must use the PATCH API available for the individual input.
func (o LookupStreamingJobResultOutput) Inputs() InputResponseArrayOutput {
	return o.ApplyT(func(v LookupStreamingJobResult) []InputResponse { return v.Inputs }).(InputResponseArrayOutput)
}

// A GUID uniquely identifying the streaming job. This GUID is generated upon creation of the streaming job.
func (o LookupStreamingJobResultOutput) JobId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStreamingJobResult) string { return v.JobId }).(pulumi.StringOutput)
}

// Describes the state of the streaming job.
func (o LookupStreamingJobResultOutput) JobState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStreamingJobResult) string { return v.JobState }).(pulumi.StringOutput)
}

// The properties that are associated with an Azure Storage account with MSI
func (o LookupStreamingJobResultOutput) JobStorageAccount() JobStorageAccountResponsePtrOutput {
	return o.ApplyT(func(v LookupStreamingJobResult) *JobStorageAccountResponse { return v.JobStorageAccount }).(JobStorageAccountResponsePtrOutput)
}

// Describes the type of the job. Valid modes are `Cloud` and 'Edge'.
func (o LookupStreamingJobResultOutput) JobType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStreamingJobResult) *string { return v.JobType }).(pulumi.StringPtrOutput)
}

// Value is either an ISO-8601 formatted timestamp indicating the last output event time of the streaming job or null indicating that output has not yet been produced. In case of multiple outputs or multiple streams, this shows the latest value in that set.
func (o LookupStreamingJobResultOutput) LastOutputEventTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStreamingJobResult) string { return v.LastOutputEventTime }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupStreamingJobResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStreamingJobResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o LookupStreamingJobResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStreamingJobResult) string { return v.Name }).(pulumi.StringOutput)
}

// Indicates the policy to apply to events that arrive at the output and cannot be written to the external storage due to being malformed (missing column values, column values of wrong type or size).
func (o LookupStreamingJobResultOutput) OutputErrorPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStreamingJobResult) *string { return v.OutputErrorPolicy }).(pulumi.StringPtrOutput)
}

// This property should only be utilized when it is desired that the job be started immediately upon creation. Value may be JobStartTime, CustomTime, or LastOutputEventTime to indicate whether the starting point of the output event stream should start whenever the job is started, start at a custom user time stamp specified via the outputStartTime property, or start from the last event output time.
func (o LookupStreamingJobResultOutput) OutputStartMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStreamingJobResult) *string { return v.OutputStartMode }).(pulumi.StringPtrOutput)
}

// Value is either an ISO-8601 formatted time stamp that indicates the starting point of the output event stream, or null to indicate that the output event stream will start whenever the streaming job is started. This property must have a value if outputStartMode is set to CustomTime.
func (o LookupStreamingJobResultOutput) OutputStartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStreamingJobResult) *string { return v.OutputStartTime }).(pulumi.StringPtrOutput)
}

// A list of one or more outputs for the streaming job. The name property for each output is required when specifying this property in a PUT request. This property cannot be modify via a PATCH operation. You must use the PATCH API available for the individual output.
func (o LookupStreamingJobResultOutput) Outputs() OutputResponseArrayOutput {
	return o.ApplyT(func(v LookupStreamingJobResult) []OutputResponse { return v.Outputs }).(OutputResponseArrayOutput)
}

// Describes the provisioning status of the streaming job.
func (o LookupStreamingJobResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStreamingJobResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Describes the SKU of the streaming job. Required on PUT (CreateOrReplace) requests.
func (o LookupStreamingJobResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupStreamingJobResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

// Resource tags.
func (o LookupStreamingJobResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupStreamingJobResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Indicates the query and the number of streaming units to use for the streaming job. The name property of the transformation is required when specifying this property in a PUT request. This property cannot be modify via a PATCH operation. You must use the PATCH API available for the individual transformation.
func (o LookupStreamingJobResultOutput) Transformation() TransformationResponsePtrOutput {
	return o.ApplyT(func(v LookupStreamingJobResult) *TransformationResponse { return v.Transformation }).(TransformationResponsePtrOutput)
}

// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
func (o LookupStreamingJobResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStreamingJobResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupStreamingJobResultOutput{})
}
