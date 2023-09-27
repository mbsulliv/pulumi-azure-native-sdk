// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets information about the specified job.
func LookupJob(ctx *pulumi.Context, args *LookupJobArgs, opts ...pulumi.InvokeOption) (*LookupJobResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupJobResult
	err := ctx.Invoke("azure-native:databox/v20230301:getJob", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupJobArgs struct {
	// $expand is supported on details parameter for job, which provides details on the job stages.
	Expand *string `pulumi:"expand"`
	// The name of the job Resource within the specified resource group. job names must be between 3 and 24 characters in length and use any alphanumeric and underscore only
	JobName string `pulumi:"jobName"`
	// The Resource Group Name
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Job Resource.
type LookupJobResult struct {
	// Reason for cancellation.
	CancellationReason string `pulumi:"cancellationReason"`
	// Name of the stage where delay might be present.
	DelayedStage string `pulumi:"delayedStage"`
	// Delivery Info of Job.
	DeliveryInfo *JobDeliveryInfoResponse `pulumi:"deliveryInfo"`
	// Delivery type of Job.
	DeliveryType *string `pulumi:"deliveryType"`
	// Details of a job run. This field will only be sent for expand details filter.
	Details interface{} `pulumi:"details"`
	// Top level error for the job.
	Error CloudErrorResponse `pulumi:"error"`
	// Id of the object.
	Id string `pulumi:"id"`
	// Msi identity of the resource
	Identity *ResourceIdentityResponse `pulumi:"identity"`
	// Describes whether the job is cancellable or not.
	IsCancellable bool `pulumi:"isCancellable"`
	// Flag to indicate cancellation of scheduled job.
	IsCancellableWithoutFee bool `pulumi:"isCancellableWithoutFee"`
	// Describes whether the job is deletable or not.
	IsDeletable bool `pulumi:"isDeletable"`
	// Is Prepare To Ship Enabled on this job
	IsPrepareToShipEnabled bool `pulumi:"isPrepareToShipEnabled"`
	// Describes whether the shipping address is editable or not.
	IsShippingAddressEditable bool `pulumi:"isShippingAddressEditable"`
	// The location of the resource. This will be one of the supported and registered Azure Regions (e.g. West US, East US, Southeast Asia, etc.). The region of a resource cannot be changed once it is created, but if an identical region is specified on update the request will succeed.
	Location string `pulumi:"location"`
	// Name of the object.
	Name string `pulumi:"name"`
	// The Editable status for Reverse Shipping Address and Contact Info
	ReverseShippingDetailsUpdate string `pulumi:"reverseShippingDetailsUpdate"`
	// The Editable status for Reverse Transport preferences
	ReverseTransportPreferenceUpdate string `pulumi:"reverseTransportPreferenceUpdate"`
	// The sku type.
	Sku SkuResponse `pulumi:"sku"`
	// Time at which the job was started in UTC ISO 8601 format.
	StartTime string `pulumi:"startTime"`
	// Name of the stage which is in progress.
	Status string `pulumi:"status"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The list of key value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups).
	Tags map[string]string `pulumi:"tags"`
	// Type of the data transfer.
	TransferType string `pulumi:"transferType"`
	// Type of the object.
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupJobResult
func (val *LookupJobResult) Defaults() *LookupJobResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.DeliveryType == nil {
		deliveryType_ := "NonScheduled"
		tmp.DeliveryType = &deliveryType_
	}
	tmp.Identity = tmp.Identity.Defaults()

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
	// $expand is supported on details parameter for job, which provides details on the job stages.
	Expand pulumi.StringPtrInput `pulumi:"expand"`
	// The name of the job Resource within the specified resource group. job names must be between 3 and 24 characters in length and use any alphanumeric and underscore only
	JobName pulumi.StringInput `pulumi:"jobName"`
	// The Resource Group Name
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupJobOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupJobArgs)(nil)).Elem()
}

// Job Resource.
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

// Reason for cancellation.
func (o LookupJobResultOutput) CancellationReason() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobResult) string { return v.CancellationReason }).(pulumi.StringOutput)
}

// Name of the stage where delay might be present.
func (o LookupJobResultOutput) DelayedStage() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobResult) string { return v.DelayedStage }).(pulumi.StringOutput)
}

// Delivery Info of Job.
func (o LookupJobResultOutput) DeliveryInfo() JobDeliveryInfoResponsePtrOutput {
	return o.ApplyT(func(v LookupJobResult) *JobDeliveryInfoResponse { return v.DeliveryInfo }).(JobDeliveryInfoResponsePtrOutput)
}

// Delivery type of Job.
func (o LookupJobResultOutput) DeliveryType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupJobResult) *string { return v.DeliveryType }).(pulumi.StringPtrOutput)
}

// Details of a job run. This field will only be sent for expand details filter.
func (o LookupJobResultOutput) Details() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupJobResult) interface{} { return v.Details }).(pulumi.AnyOutput)
}

// Top level error for the job.
func (o LookupJobResultOutput) Error() CloudErrorResponseOutput {
	return o.ApplyT(func(v LookupJobResult) CloudErrorResponse { return v.Error }).(CloudErrorResponseOutput)
}

// Id of the object.
func (o LookupJobResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobResult) string { return v.Id }).(pulumi.StringOutput)
}

// Msi identity of the resource
func (o LookupJobResultOutput) Identity() ResourceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupJobResult) *ResourceIdentityResponse { return v.Identity }).(ResourceIdentityResponsePtrOutput)
}

// Describes whether the job is cancellable or not.
func (o LookupJobResultOutput) IsCancellable() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupJobResult) bool { return v.IsCancellable }).(pulumi.BoolOutput)
}

// Flag to indicate cancellation of scheduled job.
func (o LookupJobResultOutput) IsCancellableWithoutFee() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupJobResult) bool { return v.IsCancellableWithoutFee }).(pulumi.BoolOutput)
}

// Describes whether the job is deletable or not.
func (o LookupJobResultOutput) IsDeletable() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupJobResult) bool { return v.IsDeletable }).(pulumi.BoolOutput)
}

// Is Prepare To Ship Enabled on this job
func (o LookupJobResultOutput) IsPrepareToShipEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupJobResult) bool { return v.IsPrepareToShipEnabled }).(pulumi.BoolOutput)
}

// Describes whether the shipping address is editable or not.
func (o LookupJobResultOutput) IsShippingAddressEditable() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupJobResult) bool { return v.IsShippingAddressEditable }).(pulumi.BoolOutput)
}

// The location of the resource. This will be one of the supported and registered Azure Regions (e.g. West US, East US, Southeast Asia, etc.). The region of a resource cannot be changed once it is created, but if an identical region is specified on update the request will succeed.
func (o LookupJobResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobResult) string { return v.Location }).(pulumi.StringOutput)
}

// Name of the object.
func (o LookupJobResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobResult) string { return v.Name }).(pulumi.StringOutput)
}

// The Editable status for Reverse Shipping Address and Contact Info
func (o LookupJobResultOutput) ReverseShippingDetailsUpdate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobResult) string { return v.ReverseShippingDetailsUpdate }).(pulumi.StringOutput)
}

// The Editable status for Reverse Transport preferences
func (o LookupJobResultOutput) ReverseTransportPreferenceUpdate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobResult) string { return v.ReverseTransportPreferenceUpdate }).(pulumi.StringOutput)
}

// The sku type.
func (o LookupJobResultOutput) Sku() SkuResponseOutput {
	return o.ApplyT(func(v LookupJobResult) SkuResponse { return v.Sku }).(SkuResponseOutput)
}

// Time at which the job was started in UTC ISO 8601 format.
func (o LookupJobResultOutput) StartTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobResult) string { return v.StartTime }).(pulumi.StringOutput)
}

// Name of the stage which is in progress.
func (o LookupJobResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobResult) string { return v.Status }).(pulumi.StringOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupJobResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupJobResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The list of key value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups).
func (o LookupJobResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupJobResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Type of the data transfer.
func (o LookupJobResultOutput) TransferType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobResult) string { return v.TransferType }).(pulumi.StringOutput)
}

// Type of the object.
func (o LookupJobResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupJobResultOutput{})
}
