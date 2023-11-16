// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20231215preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get properties of a domain topic.
func LookupDomainTopic(ctx *pulumi.Context, args *LookupDomainTopicArgs, opts ...pulumi.InvokeOption) (*LookupDomainTopicResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupDomainTopicResult
	err := ctx.Invoke("azure-native:eventgrid/v20231215preview:getDomainTopic", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDomainTopicArgs struct {
	// Name of the domain.
	DomainName string `pulumi:"domainName"`
	// Name of the topic.
	DomainTopicName string `pulumi:"domainTopicName"`
	// The name of the resource group within the user's subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Domain Topic.
type LookupDomainTopicResult struct {
	// Fully qualified identifier of the resource.
	Id string `pulumi:"id"`
	// Name of the resource.
	Name string `pulumi:"name"`
	// Provisioning state of the domain topic.
	ProvisioningState string `pulumi:"provisioningState"`
	// The system metadata relating to Domain Topic resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Type of the resource.
	Type string `pulumi:"type"`
}

func LookupDomainTopicOutput(ctx *pulumi.Context, args LookupDomainTopicOutputArgs, opts ...pulumi.InvokeOption) LookupDomainTopicResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDomainTopicResult, error) {
			args := v.(LookupDomainTopicArgs)
			r, err := LookupDomainTopic(ctx, &args, opts...)
			var s LookupDomainTopicResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDomainTopicResultOutput)
}

type LookupDomainTopicOutputArgs struct {
	// Name of the domain.
	DomainName pulumi.StringInput `pulumi:"domainName"`
	// Name of the topic.
	DomainTopicName pulumi.StringInput `pulumi:"domainTopicName"`
	// The name of the resource group within the user's subscription.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupDomainTopicOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDomainTopicArgs)(nil)).Elem()
}

// Domain Topic.
type LookupDomainTopicResultOutput struct{ *pulumi.OutputState }

func (LookupDomainTopicResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDomainTopicResult)(nil)).Elem()
}

func (o LookupDomainTopicResultOutput) ToLookupDomainTopicResultOutput() LookupDomainTopicResultOutput {
	return o
}

func (o LookupDomainTopicResultOutput) ToLookupDomainTopicResultOutputWithContext(ctx context.Context) LookupDomainTopicResultOutput {
	return o
}

// Fully qualified identifier of the resource.
func (o LookupDomainTopicResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainTopicResult) string { return v.Id }).(pulumi.StringOutput)
}

// Name of the resource.
func (o LookupDomainTopicResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainTopicResult) string { return v.Name }).(pulumi.StringOutput)
}

// Provisioning state of the domain topic.
func (o LookupDomainTopicResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainTopicResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The system metadata relating to Domain Topic resource.
func (o LookupDomainTopicResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupDomainTopicResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Type of the resource.
func (o LookupDomainTopicResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainTopicResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDomainTopicResultOutput{})
}
