// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get properties of a namespace topic.
func LookupNamespaceTopic(ctx *pulumi.Context, args *LookupNamespaceTopicArgs, opts ...pulumi.InvokeOption) (*LookupNamespaceTopicResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupNamespaceTopicResult
	err := ctx.Invoke("azure-native:eventgrid/v20230601preview:getNamespaceTopic", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupNamespaceTopicArgs struct {
	// Name of the namespace.
	NamespaceName string `pulumi:"namespaceName"`
	// The name of the resource group within the user's subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the namespace topic.
	TopicName string `pulumi:"topicName"`
}

// Namespace topic details.
type LookupNamespaceTopicResult struct {
	// Event retention for the namespace topic expressed in days. The property default value is 1 day.
	// Min event retention duration value is 1 day and max event retention duration value is 1 day.
	EventRetentionInDays *int `pulumi:"eventRetentionInDays"`
	// Fully qualified identifier of the resource.
	Id string `pulumi:"id"`
	// This determines the format that is expected for incoming events published to the topic.
	InputSchema *string `pulumi:"inputSchema"`
	// Name of the resource.
	Name string `pulumi:"name"`
	// Provisioning state of the namespace topic.
	ProvisioningState string `pulumi:"provisioningState"`
	// Publisher type of the namespace topic.
	PublisherType *string `pulumi:"publisherType"`
	// The system metadata relating to namespace topic resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Type of the resource.
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupNamespaceTopicResult
func (val *LookupNamespaceTopicResult) Defaults() *LookupNamespaceTopicResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.InputSchema == nil {
		inputSchema_ := "CloudEventSchemaV1_0"
		tmp.InputSchema = &inputSchema_
	}
	return &tmp
}

func LookupNamespaceTopicOutput(ctx *pulumi.Context, args LookupNamespaceTopicOutputArgs, opts ...pulumi.InvokeOption) LookupNamespaceTopicResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNamespaceTopicResult, error) {
			args := v.(LookupNamespaceTopicArgs)
			r, err := LookupNamespaceTopic(ctx, &args, opts...)
			var s LookupNamespaceTopicResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNamespaceTopicResultOutput)
}

type LookupNamespaceTopicOutputArgs struct {
	// Name of the namespace.
	NamespaceName pulumi.StringInput `pulumi:"namespaceName"`
	// The name of the resource group within the user's subscription.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name of the namespace topic.
	TopicName pulumi.StringInput `pulumi:"topicName"`
}

func (LookupNamespaceTopicOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNamespaceTopicArgs)(nil)).Elem()
}

// Namespace topic details.
type LookupNamespaceTopicResultOutput struct{ *pulumi.OutputState }

func (LookupNamespaceTopicResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNamespaceTopicResult)(nil)).Elem()
}

func (o LookupNamespaceTopicResultOutput) ToLookupNamespaceTopicResultOutput() LookupNamespaceTopicResultOutput {
	return o
}

func (o LookupNamespaceTopicResultOutput) ToLookupNamespaceTopicResultOutputWithContext(ctx context.Context) LookupNamespaceTopicResultOutput {
	return o
}

func (o LookupNamespaceTopicResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupNamespaceTopicResult] {
	return pulumix.Output[LookupNamespaceTopicResult]{
		OutputState: o.OutputState,
	}
}

// Event retention for the namespace topic expressed in days. The property default value is 1 day.
// Min event retention duration value is 1 day and max event retention duration value is 1 day.
func (o LookupNamespaceTopicResultOutput) EventRetentionInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupNamespaceTopicResult) *int { return v.EventRetentionInDays }).(pulumi.IntPtrOutput)
}

// Fully qualified identifier of the resource.
func (o LookupNamespaceTopicResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNamespaceTopicResult) string { return v.Id }).(pulumi.StringOutput)
}

// This determines the format that is expected for incoming events published to the topic.
func (o LookupNamespaceTopicResultOutput) InputSchema() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNamespaceTopicResult) *string { return v.InputSchema }).(pulumi.StringPtrOutput)
}

// Name of the resource.
func (o LookupNamespaceTopicResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNamespaceTopicResult) string { return v.Name }).(pulumi.StringOutput)
}

// Provisioning state of the namespace topic.
func (o LookupNamespaceTopicResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNamespaceTopicResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Publisher type of the namespace topic.
func (o LookupNamespaceTopicResultOutput) PublisherType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNamespaceTopicResult) *string { return v.PublisherType }).(pulumi.StringPtrOutput)
}

// The system metadata relating to namespace topic resource.
func (o LookupNamespaceTopicResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupNamespaceTopicResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Type of the resource.
func (o LookupNamespaceTopicResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNamespaceTopicResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNamespaceTopicResultOutput{})
}
