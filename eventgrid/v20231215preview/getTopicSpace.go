// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20231215preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get properties of a topic space.
func LookupTopicSpace(ctx *pulumi.Context, args *LookupTopicSpaceArgs, opts ...pulumi.InvokeOption) (*LookupTopicSpaceResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupTopicSpaceResult
	err := ctx.Invoke("azure-native:eventgrid/v20231215preview:getTopicSpace", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTopicSpaceArgs struct {
	// Name of the namespace.
	NamespaceName string `pulumi:"namespaceName"`
	// The name of the resource group within the user's subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the Topic space.
	TopicSpaceName string `pulumi:"topicSpaceName"`
}

// The Topic space resource.
type LookupTopicSpaceResult struct {
	// Description for the Topic Space resource.
	Description *string `pulumi:"description"`
	// Fully qualified identifier of the resource.
	Id string `pulumi:"id"`
	// Name of the resource.
	Name string `pulumi:"name"`
	// Provisioning state of the TopicSpace resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// The system metadata relating to the TopicSpace resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The topic filters in the topic space.
	// Example: "topicTemplates": [
	//               "devices/foo/bar",
	//               "devices/topic1/+",
	//               "devices/${principal.name}/${principal.attributes.keyName}" ].
	TopicTemplates []string `pulumi:"topicTemplates"`
	// Type of the resource.
	Type string `pulumi:"type"`
}

func LookupTopicSpaceOutput(ctx *pulumi.Context, args LookupTopicSpaceOutputArgs, opts ...pulumi.InvokeOption) LookupTopicSpaceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTopicSpaceResult, error) {
			args := v.(LookupTopicSpaceArgs)
			r, err := LookupTopicSpace(ctx, &args, opts...)
			var s LookupTopicSpaceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTopicSpaceResultOutput)
}

type LookupTopicSpaceOutputArgs struct {
	// Name of the namespace.
	NamespaceName pulumi.StringInput `pulumi:"namespaceName"`
	// The name of the resource group within the user's subscription.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name of the Topic space.
	TopicSpaceName pulumi.StringInput `pulumi:"topicSpaceName"`
}

func (LookupTopicSpaceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTopicSpaceArgs)(nil)).Elem()
}

// The Topic space resource.
type LookupTopicSpaceResultOutput struct{ *pulumi.OutputState }

func (LookupTopicSpaceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTopicSpaceResult)(nil)).Elem()
}

func (o LookupTopicSpaceResultOutput) ToLookupTopicSpaceResultOutput() LookupTopicSpaceResultOutput {
	return o
}

func (o LookupTopicSpaceResultOutput) ToLookupTopicSpaceResultOutputWithContext(ctx context.Context) LookupTopicSpaceResultOutput {
	return o
}

// Description for the Topic Space resource.
func (o LookupTopicSpaceResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTopicSpaceResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Fully qualified identifier of the resource.
func (o LookupTopicSpaceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTopicSpaceResult) string { return v.Id }).(pulumi.StringOutput)
}

// Name of the resource.
func (o LookupTopicSpaceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTopicSpaceResult) string { return v.Name }).(pulumi.StringOutput)
}

// Provisioning state of the TopicSpace resource.
func (o LookupTopicSpaceResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTopicSpaceResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The system metadata relating to the TopicSpace resource.
func (o LookupTopicSpaceResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupTopicSpaceResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The topic filters in the topic space.
// Example: "topicTemplates": [
//
//	"devices/foo/bar",
//	"devices/topic1/+",
//	"devices/${principal.name}/${principal.attributes.keyName}" ].
func (o LookupTopicSpaceResultOutput) TopicTemplates() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupTopicSpaceResult) []string { return v.TopicTemplates }).(pulumi.StringArrayOutput)
}

// Type of the resource.
func (o LookupTopicSpaceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTopicSpaceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTopicSpaceResultOutput{})
}