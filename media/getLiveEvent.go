// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package media

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets properties of a live event.
// Azure REST API version: 2022-11-01.
//
// Other available API versions: 2018-06-01-preview, 2019-05-01-preview.
func LookupLiveEvent(ctx *pulumi.Context, args *LookupLiveEventArgs, opts ...pulumi.InvokeOption) (*LookupLiveEventResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupLiveEventResult
	err := ctx.Invoke("azure-native:media:getLiveEvent", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLiveEventArgs struct {
	// The Media Services account name.
	AccountName string `pulumi:"accountName"`
	// The name of the live event, maximum length is 32.
	LiveEventName string `pulumi:"liveEventName"`
	// The name of the resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The live event.
type LookupLiveEventResult struct {
	// The creation time for the live event
	Created string `pulumi:"created"`
	// Live event cross site access policies.
	CrossSiteAccessPolicies *CrossSiteAccessPoliciesResponse `pulumi:"crossSiteAccessPolicies"`
	// A description for the live event.
	Description *string `pulumi:"description"`
	// Encoding settings for the live event. It configures whether a live encoder is used for the live event and settings for the live encoder if it is used.
	Encoding *LiveEventEncodingResponse `pulumi:"encoding"`
	// When useStaticHostname is set to true, the hostnamePrefix specifies the first part of the hostname assigned to the live event preview and ingest endpoints. The final hostname would be a combination of this prefix, the media service account name and a short code for the Azure Media Services data center.
	HostnamePrefix *string `pulumi:"hostnamePrefix"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// Live event input settings. It defines how the live event receives input from a contribution encoder.
	Input LiveEventInputResponse `pulumi:"input"`
	// The last modified time of the live event.
	LastModified string `pulumi:"lastModified"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Live event preview settings. Preview allows live event producers to preview the live streaming content without creating any live output.
	Preview *LiveEventPreviewResponse `pulumi:"preview"`
	// The provisioning state of the live event.
	ProvisioningState string `pulumi:"provisioningState"`
	// The resource state of the live event. See https://go.microsoft.com/fwlink/?linkid=2139012 for more information.
	ResourceState string `pulumi:"resourceState"`
	// The options to use for the LiveEvent. This value is specified at creation time and cannot be updated. The valid values for the array entry values are 'Default' and 'LowLatency'.
	StreamOptions []string `pulumi:"streamOptions"`
	// The system metadata relating to this resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Live transcription settings for the live event. See https://go.microsoft.com/fwlink/?linkid=2133742 for more information about the live transcription feature.
	Transcriptions []LiveEventTranscriptionResponse `pulumi:"transcriptions"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// Specifies whether a static hostname would be assigned to the live event preview and ingest endpoints. This value can only be updated if the live event is in Standby state
	UseStaticHostname *bool `pulumi:"useStaticHostname"`
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
	// The name of the live event, maximum length is 32.
	LiveEventName pulumi.StringInput `pulumi:"liveEventName"`
	// The name of the resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupLiveEventOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLiveEventArgs)(nil)).Elem()
}

// The live event.
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

// The creation time for the live event
func (o LookupLiveEventResultOutput) Created() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLiveEventResult) string { return v.Created }).(pulumi.StringOutput)
}

// Live event cross site access policies.
func (o LookupLiveEventResultOutput) CrossSiteAccessPolicies() CrossSiteAccessPoliciesResponsePtrOutput {
	return o.ApplyT(func(v LookupLiveEventResult) *CrossSiteAccessPoliciesResponse { return v.CrossSiteAccessPolicies }).(CrossSiteAccessPoliciesResponsePtrOutput)
}

// A description for the live event.
func (o LookupLiveEventResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLiveEventResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Encoding settings for the live event. It configures whether a live encoder is used for the live event and settings for the live encoder if it is used.
func (o LookupLiveEventResultOutput) Encoding() LiveEventEncodingResponsePtrOutput {
	return o.ApplyT(func(v LookupLiveEventResult) *LiveEventEncodingResponse { return v.Encoding }).(LiveEventEncodingResponsePtrOutput)
}

// When useStaticHostname is set to true, the hostnamePrefix specifies the first part of the hostname assigned to the live event preview and ingest endpoints. The final hostname would be a combination of this prefix, the media service account name and a short code for the Azure Media Services data center.
func (o LookupLiveEventResultOutput) HostnamePrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLiveEventResult) *string { return v.HostnamePrefix }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupLiveEventResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLiveEventResult) string { return v.Id }).(pulumi.StringOutput)
}

// Live event input settings. It defines how the live event receives input from a contribution encoder.
func (o LookupLiveEventResultOutput) Input() LiveEventInputResponseOutput {
	return o.ApplyT(func(v LookupLiveEventResult) LiveEventInputResponse { return v.Input }).(LiveEventInputResponseOutput)
}

// The last modified time of the live event.
func (o LookupLiveEventResultOutput) LastModified() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLiveEventResult) string { return v.LastModified }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupLiveEventResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLiveEventResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupLiveEventResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLiveEventResult) string { return v.Name }).(pulumi.StringOutput)
}

// Live event preview settings. Preview allows live event producers to preview the live streaming content without creating any live output.
func (o LookupLiveEventResultOutput) Preview() LiveEventPreviewResponsePtrOutput {
	return o.ApplyT(func(v LookupLiveEventResult) *LiveEventPreviewResponse { return v.Preview }).(LiveEventPreviewResponsePtrOutput)
}

// The provisioning state of the live event.
func (o LookupLiveEventResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLiveEventResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The resource state of the live event. See https://go.microsoft.com/fwlink/?linkid=2139012 for more information.
func (o LookupLiveEventResultOutput) ResourceState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLiveEventResult) string { return v.ResourceState }).(pulumi.StringOutput)
}

// The options to use for the LiveEvent. This value is specified at creation time and cannot be updated. The valid values for the array entry values are 'Default' and 'LowLatency'.
func (o LookupLiveEventResultOutput) StreamOptions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupLiveEventResult) []string { return v.StreamOptions }).(pulumi.StringArrayOutput)
}

// The system metadata relating to this resource.
func (o LookupLiveEventResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupLiveEventResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupLiveEventResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupLiveEventResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Live transcription settings for the live event. See https://go.microsoft.com/fwlink/?linkid=2133742 for more information about the live transcription feature.
func (o LookupLiveEventResultOutput) Transcriptions() LiveEventTranscriptionResponseArrayOutput {
	return o.ApplyT(func(v LookupLiveEventResult) []LiveEventTranscriptionResponse { return v.Transcriptions }).(LiveEventTranscriptionResponseArrayOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupLiveEventResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLiveEventResult) string { return v.Type }).(pulumi.StringOutput)
}

// Specifies whether a static hostname would be assigned to the live event preview and ingest endpoints. This value can only be updated if the live event is in Standby state
func (o LookupLiveEventResultOutput) UseStaticHostname() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLiveEventResult) *bool { return v.UseStaticHostname }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLiveEventResultOutput{})
}
