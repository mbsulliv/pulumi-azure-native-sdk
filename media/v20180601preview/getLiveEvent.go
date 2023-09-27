// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20180601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets a Live Event.
func LookupLiveEvent(ctx *pulumi.Context, args *LookupLiveEventArgs, opts ...pulumi.InvokeOption) (*LookupLiveEventResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupLiveEventResult
	err := ctx.Invoke("azure-native:media/v20180601preview:getLiveEvent", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLiveEventArgs struct {
	// The Media Services account name.
	AccountName string `pulumi:"accountName"`
	// The name of the Live Event.
	LiveEventName string `pulumi:"liveEventName"`
	// The name of the resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The Live Event.
type LookupLiveEventResult struct {
	// The exact time the Live Event was created.
	Created string `pulumi:"created"`
	// The Live Event access policies.
	CrossSiteAccessPolicies *CrossSiteAccessPoliciesResponse `pulumi:"crossSiteAccessPolicies"`
	// The Live Event description.
	Description *string `pulumi:"description"`
	// The Live Event encoding.
	Encoding *LiveEventEncodingResponse `pulumi:"encoding"`
	// Fully qualified resource ID for the resource.
	Id string `pulumi:"id"`
	// The Live Event input.
	Input LiveEventInputResponse `pulumi:"input"`
	// The exact time the Live Event was last modified.
	LastModified string `pulumi:"lastModified"`
	// The Azure Region of the resource.
	Location *string `pulumi:"location"`
	// The name of the resource.
	Name string `pulumi:"name"`
	// The Live Event preview.
	Preview *LiveEventPreviewResponse `pulumi:"preview"`
	// The provisioning state of the Live Event.
	ProvisioningState string `pulumi:"provisioningState"`
	// The resource state of the Live Event.
	ResourceState string `pulumi:"resourceState"`
	// The stream options.
	StreamOptions []string `pulumi:"streamOptions"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource.
	Type string `pulumi:"type"`
	// The Live Event vanity URL flag.
	VanityUrl *bool `pulumi:"vanityUrl"`
}

func LookupLiveEventOutput(ctx *pulumi.Context, args LookupLiveEventOutputArgs, opts ...pulumi.InvokeOption) LookupLiveEventResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupLiveEventResult, error) {
			args := v.(LookupLiveEventArgs)
			r, err := LookupLiveEvent(ctx, &args, opts...)
			var s LookupLiveEventResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupLiveEventResultOutput)
}

type LookupLiveEventOutputArgs struct {
	// The Media Services account name.
	AccountName pulumi.StringInput `pulumi:"accountName"`
	// The name of the Live Event.
	LiveEventName pulumi.StringInput `pulumi:"liveEventName"`
	// The name of the resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupLiveEventOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLiveEventArgs)(nil)).Elem()
}

// The Live Event.
type LookupLiveEventResultOutput struct{ *pulumi.OutputState }

func (LookupLiveEventResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLiveEventResult)(nil)).Elem()
}

func (o LookupLiveEventResultOutput) ToLookupLiveEventResultOutput() LookupLiveEventResultOutput {
	return o
}

func (o LookupLiveEventResultOutput) ToLookupLiveEventResultOutputWithContext(ctx context.Context) LookupLiveEventResultOutput {
	return o
}

func (o LookupLiveEventResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupLiveEventResult] {
	return pulumix.Output[LookupLiveEventResult]{
		OutputState: o.OutputState,
	}
}

// The exact time the Live Event was created.
func (o LookupLiveEventResultOutput) Created() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLiveEventResult) string { return v.Created }).(pulumi.StringOutput)
}

// The Live Event access policies.
func (o LookupLiveEventResultOutput) CrossSiteAccessPolicies() CrossSiteAccessPoliciesResponsePtrOutput {
	return o.ApplyT(func(v LookupLiveEventResult) *CrossSiteAccessPoliciesResponse { return v.CrossSiteAccessPolicies }).(CrossSiteAccessPoliciesResponsePtrOutput)
}

// The Live Event description.
func (o LookupLiveEventResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLiveEventResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The Live Event encoding.
func (o LookupLiveEventResultOutput) Encoding() LiveEventEncodingResponsePtrOutput {
	return o.ApplyT(func(v LookupLiveEventResult) *LiveEventEncodingResponse { return v.Encoding }).(LiveEventEncodingResponsePtrOutput)
}

// Fully qualified resource ID for the resource.
func (o LookupLiveEventResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLiveEventResult) string { return v.Id }).(pulumi.StringOutput)
}

// The Live Event input.
func (o LookupLiveEventResultOutput) Input() LiveEventInputResponseOutput {
	return o.ApplyT(func(v LookupLiveEventResult) LiveEventInputResponse { return v.Input }).(LiveEventInputResponseOutput)
}

// The exact time the Live Event was last modified.
func (o LookupLiveEventResultOutput) LastModified() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLiveEventResult) string { return v.LastModified }).(pulumi.StringOutput)
}

// The Azure Region of the resource.
func (o LookupLiveEventResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLiveEventResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the resource.
func (o LookupLiveEventResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLiveEventResult) string { return v.Name }).(pulumi.StringOutput)
}

// The Live Event preview.
func (o LookupLiveEventResultOutput) Preview() LiveEventPreviewResponsePtrOutput {
	return o.ApplyT(func(v LookupLiveEventResult) *LiveEventPreviewResponse { return v.Preview }).(LiveEventPreviewResponsePtrOutput)
}

// The provisioning state of the Live Event.
func (o LookupLiveEventResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLiveEventResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The resource state of the Live Event.
func (o LookupLiveEventResultOutput) ResourceState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLiveEventResult) string { return v.ResourceState }).(pulumi.StringOutput)
}

// The stream options.
func (o LookupLiveEventResultOutput) StreamOptions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupLiveEventResult) []string { return v.StreamOptions }).(pulumi.StringArrayOutput)
}

// Resource tags.
func (o LookupLiveEventResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupLiveEventResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource.
func (o LookupLiveEventResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLiveEventResult) string { return v.Type }).(pulumi.StringOutput)
}

// The Live Event vanity URL flag.
func (o LookupLiveEventResultOutput) VanityUrl() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLiveEventResult) *bool { return v.VanityUrl }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLiveEventResultOutput{})
}
