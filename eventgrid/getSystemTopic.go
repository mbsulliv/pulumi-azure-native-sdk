// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package eventgrid

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get properties of a system topic.
// Azure REST API version: 2022-06-15.
func LookupSystemTopic(ctx *pulumi.Context, args *LookupSystemTopicArgs, opts ...pulumi.InvokeOption) (*LookupSystemTopicResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupSystemTopicResult
	err := ctx.Invoke("azure-native:eventgrid:getSystemTopic", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSystemTopicArgs struct {
	// The name of the resource group within the user's subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the system topic.
	SystemTopicName string `pulumi:"systemTopicName"`
}

// EventGrid System Topic.
type LookupSystemTopicResult struct {
	// Fully qualified identifier of the resource.
	Id string `pulumi:"id"`
	// Identity information for the resource.
	Identity *IdentityInfoResponse `pulumi:"identity"`
	// Location of the resource.
	Location string `pulumi:"location"`
	// Metric resource id for the system topic.
	MetricResourceId string `pulumi:"metricResourceId"`
	// Name of the resource.
	Name string `pulumi:"name"`
	// Provisioning state of the system topic.
	ProvisioningState string `pulumi:"provisioningState"`
	// Source for the system topic.
	Source *string `pulumi:"source"`
	// The system metadata relating to System Topic resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Tags of the resource.
	Tags map[string]string `pulumi:"tags"`
	// TopicType for the system topic.
	TopicType *string `pulumi:"topicType"`
	// Type of the resource.
	Type string `pulumi:"type"`
}

func LookupSystemTopicOutput(ctx *pulumi.Context, args LookupSystemTopicOutputArgs, opts ...pulumi.InvokeOption) LookupSystemTopicResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSystemTopicResult, error) {
			args := v.(LookupSystemTopicArgs)
			r, err := LookupSystemTopic(ctx, &args, opts...)
			var s LookupSystemTopicResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSystemTopicResultOutput)
}

type LookupSystemTopicOutputArgs struct {
	// The name of the resource group within the user's subscription.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name of the system topic.
	SystemTopicName pulumi.StringInput `pulumi:"systemTopicName"`
}

func (LookupSystemTopicOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemTopicArgs)(nil)).Elem()
}

// EventGrid System Topic.
type LookupSystemTopicResultOutput struct{ *pulumi.OutputState }

func (LookupSystemTopicResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemTopicResult)(nil)).Elem()
}

func (o LookupSystemTopicResultOutput) ToLookupSystemTopicResultOutput() LookupSystemTopicResultOutput {
	return o
}

func (o LookupSystemTopicResultOutput) ToLookupSystemTopicResultOutputWithContext(ctx context.Context) LookupSystemTopicResultOutput {
	return o
}

func (o LookupSystemTopicResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupSystemTopicResult] {
	return pulumix.Output[LookupSystemTopicResult]{
		OutputState: o.OutputState,
	}
}

// Fully qualified identifier of the resource.
func (o LookupSystemTopicResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemTopicResult) string { return v.Id }).(pulumi.StringOutput)
}

// Identity information for the resource.
func (o LookupSystemTopicResultOutput) Identity() IdentityInfoResponsePtrOutput {
	return o.ApplyT(func(v LookupSystemTopicResult) *IdentityInfoResponse { return v.Identity }).(IdentityInfoResponsePtrOutput)
}

// Location of the resource.
func (o LookupSystemTopicResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemTopicResult) string { return v.Location }).(pulumi.StringOutput)
}

// Metric resource id for the system topic.
func (o LookupSystemTopicResultOutput) MetricResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemTopicResult) string { return v.MetricResourceId }).(pulumi.StringOutput)
}

// Name of the resource.
func (o LookupSystemTopicResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemTopicResult) string { return v.Name }).(pulumi.StringOutput)
}

// Provisioning state of the system topic.
func (o LookupSystemTopicResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemTopicResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Source for the system topic.
func (o LookupSystemTopicResultOutput) Source() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSystemTopicResult) *string { return v.Source }).(pulumi.StringPtrOutput)
}

// The system metadata relating to System Topic resource.
func (o LookupSystemTopicResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupSystemTopicResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Tags of the resource.
func (o LookupSystemTopicResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSystemTopicResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// TopicType for the system topic.
func (o LookupSystemTopicResultOutput) TopicType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSystemTopicResult) *string { return v.TopicType }).(pulumi.StringPtrOutput)
}

// Type of the resource.
func (o LookupSystemTopicResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemTopicResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSystemTopicResultOutput{})
}
