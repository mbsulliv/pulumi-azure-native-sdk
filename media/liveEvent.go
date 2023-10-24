// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package media

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The live event.
// Azure REST API version: 2022-11-01. Prior API version in Azure Native 1.x: 2020-05-01.
//
// Other available API versions: 2018-06-01-preview, 2019-05-01-preview.
type LiveEvent struct {
	pulumi.CustomResourceState

	// The creation time for the live event
	Created pulumi.StringOutput `pulumi:"created"`
	// Live event cross site access policies.
	CrossSiteAccessPolicies CrossSiteAccessPoliciesResponsePtrOutput `pulumi:"crossSiteAccessPolicies"`
	// A description for the live event.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Encoding settings for the live event. It configures whether a live encoder is used for the live event and settings for the live encoder if it is used.
	Encoding LiveEventEncodingResponsePtrOutput `pulumi:"encoding"`
	// When useStaticHostname is set to true, the hostnamePrefix specifies the first part of the hostname assigned to the live event preview and ingest endpoints. The final hostname would be a combination of this prefix, the media service account name and a short code for the Azure Media Services data center.
	HostnamePrefix pulumi.StringPtrOutput `pulumi:"hostnamePrefix"`
	// Live event input settings. It defines how the live event receives input from a contribution encoder.
	Input LiveEventInputResponseOutput `pulumi:"input"`
	// The last modified time of the live event.
	LastModified pulumi.StringOutput `pulumi:"lastModified"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Live event preview settings. Preview allows live event producers to preview the live streaming content without creating any live output.
	Preview LiveEventPreviewResponsePtrOutput `pulumi:"preview"`
	// The provisioning state of the live event.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The resource state of the live event. See https://go.microsoft.com/fwlink/?linkid=2139012 for more information.
	ResourceState pulumi.StringOutput `pulumi:"resourceState"`
	// The options to use for the LiveEvent. This value is specified at creation time and cannot be updated. The valid values for the array entry values are 'Default' and 'LowLatency'.
	StreamOptions pulumi.StringArrayOutput `pulumi:"streamOptions"`
	// The system metadata relating to this resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Live transcription settings for the live event. See https://go.microsoft.com/fwlink/?linkid=2133742 for more information about the live transcription feature.
	Transcriptions LiveEventTranscriptionResponseArrayOutput `pulumi:"transcriptions"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// Specifies whether a static hostname would be assigned to the live event preview and ingest endpoints. This value can only be updated if the live event is in Standby state
	UseStaticHostname pulumi.BoolPtrOutput `pulumi:"useStaticHostname"`
}

// NewLiveEvent registers a new resource with the given unique name, arguments, and options.
func NewLiveEvent(ctx *pulumi.Context,
	name string, args *LiveEventArgs, opts ...pulumi.ResourceOption) (*LiveEvent, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.Input == nil {
		return nil, errors.New("invalid value for required argument 'Input'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:media/v20180330preview:LiveEvent"),
		},
		{
			Type: pulumi.String("azure-native:media/v20180601preview:LiveEvent"),
		},
		{
			Type: pulumi.String("azure-native:media/v20180701:LiveEvent"),
		},
		{
			Type: pulumi.String("azure-native:media/v20190501preview:LiveEvent"),
		},
		{
			Type: pulumi.String("azure-native:media/v20200501:LiveEvent"),
		},
		{
			Type: pulumi.String("azure-native:media/v20210601:LiveEvent"),
		},
		{
			Type: pulumi.String("azure-native:media/v20211101:LiveEvent"),
		},
		{
			Type: pulumi.String("azure-native:media/v20220801:LiveEvent"),
		},
		{
			Type: pulumi.String("azure-native:media/v20221101:LiveEvent"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource LiveEvent
	err := ctx.RegisterResource("azure-native:media:LiveEvent", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLiveEvent gets an existing LiveEvent resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLiveEvent(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LiveEventState, opts ...pulumi.ResourceOption) (*LiveEvent, error) {
	var resource LiveEvent
	err := ctx.ReadResource("azure-native:media:LiveEvent", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LiveEvent resources.
type liveEventState struct {
}

type LiveEventState struct {
}

func (LiveEventState) ElementType() reflect.Type {
	return reflect.TypeOf((*liveEventState)(nil)).Elem()
}

type liveEventArgs struct {
	// The Media Services account name.
	AccountName string `pulumi:"accountName"`
	// The flag indicates if the resource should be automatically started on creation.
	AutoStart *bool `pulumi:"autoStart"`
	// Live event cross site access policies.
	CrossSiteAccessPolicies *CrossSiteAccessPolicies `pulumi:"crossSiteAccessPolicies"`
	// A description for the live event.
	Description *string `pulumi:"description"`
	// Encoding settings for the live event. It configures whether a live encoder is used for the live event and settings for the live encoder if it is used.
	Encoding *LiveEventEncoding `pulumi:"encoding"`
	// When useStaticHostname is set to true, the hostnamePrefix specifies the first part of the hostname assigned to the live event preview and ingest endpoints. The final hostname would be a combination of this prefix, the media service account name and a short code for the Azure Media Services data center.
	HostnamePrefix *string `pulumi:"hostnamePrefix"`
	// Live event input settings. It defines how the live event receives input from a contribution encoder.
	Input LiveEventInputType `pulumi:"input"`
	// The name of the live event, maximum length is 32.
	LiveEventName *string `pulumi:"liveEventName"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// Live event preview settings. Preview allows live event producers to preview the live streaming content without creating any live output.
	Preview *LiveEventPreview `pulumi:"preview"`
	// The name of the resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The options to use for the LiveEvent. This value is specified at creation time and cannot be updated. The valid values for the array entry values are 'Default' and 'LowLatency'.
	StreamOptions []string `pulumi:"streamOptions"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Live transcription settings for the live event. See https://go.microsoft.com/fwlink/?linkid=2133742 for more information about the live transcription feature.
	Transcriptions []LiveEventTranscription `pulumi:"transcriptions"`
	// Specifies whether a static hostname would be assigned to the live event preview and ingest endpoints. This value can only be updated if the live event is in Standby state
	UseStaticHostname *bool `pulumi:"useStaticHostname"`
}

// The set of arguments for constructing a LiveEvent resource.
type LiveEventArgs struct {
	// The Media Services account name.
	AccountName pulumi.StringInput
	// The flag indicates if the resource should be automatically started on creation.
	AutoStart pulumi.BoolPtrInput
	// Live event cross site access policies.
	CrossSiteAccessPolicies CrossSiteAccessPoliciesPtrInput
	// A description for the live event.
	Description pulumi.StringPtrInput
	// Encoding settings for the live event. It configures whether a live encoder is used for the live event and settings for the live encoder if it is used.
	Encoding LiveEventEncodingPtrInput
	// When useStaticHostname is set to true, the hostnamePrefix specifies the first part of the hostname assigned to the live event preview and ingest endpoints. The final hostname would be a combination of this prefix, the media service account name and a short code for the Azure Media Services data center.
	HostnamePrefix pulumi.StringPtrInput
	// Live event input settings. It defines how the live event receives input from a contribution encoder.
	Input LiveEventInputTypeInput
	// The name of the live event, maximum length is 32.
	LiveEventName pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// Live event preview settings. Preview allows live event producers to preview the live streaming content without creating any live output.
	Preview LiveEventPreviewPtrInput
	// The name of the resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput
	// The options to use for the LiveEvent. This value is specified at creation time and cannot be updated. The valid values for the array entry values are 'Default' and 'LowLatency'.
	StreamOptions pulumi.StringArrayInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Live transcription settings for the live event. See https://go.microsoft.com/fwlink/?linkid=2133742 for more information about the live transcription feature.
	Transcriptions LiveEventTranscriptionArrayInput
	// Specifies whether a static hostname would be assigned to the live event preview and ingest endpoints. This value can only be updated if the live event is in Standby state
	UseStaticHostname pulumi.BoolPtrInput
}

func (LiveEventArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*liveEventArgs)(nil)).Elem()
}

type LiveEventInput interface {
	pulumi.Input

	ToLiveEventOutput() LiveEventOutput
	ToLiveEventOutputWithContext(ctx context.Context) LiveEventOutput
}

func (*LiveEvent) ElementType() reflect.Type {
	return reflect.TypeOf((**LiveEvent)(nil)).Elem()
}

func (i *LiveEvent) ToLiveEventOutput() LiveEventOutput {
	return i.ToLiveEventOutputWithContext(context.Background())
}

func (i *LiveEvent) ToLiveEventOutputWithContext(ctx context.Context) LiveEventOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiveEventOutput)
}

func (i *LiveEvent) ToOutput(ctx context.Context) pulumix.Output[*LiveEvent] {
	return pulumix.Output[*LiveEvent]{
		OutputState: i.ToLiveEventOutputWithContext(ctx).OutputState,
	}
}

type LiveEventOutput struct{ *pulumi.OutputState }

func (LiveEventOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LiveEvent)(nil)).Elem()
}

func (o LiveEventOutput) ToLiveEventOutput() LiveEventOutput {
	return o
}

func (o LiveEventOutput) ToLiveEventOutputWithContext(ctx context.Context) LiveEventOutput {
	return o
}

func (o LiveEventOutput) ToOutput(ctx context.Context) pulumix.Output[*LiveEvent] {
	return pulumix.Output[*LiveEvent]{
		OutputState: o.OutputState,
	}
}

// The creation time for the live event
func (o LiveEventOutput) Created() pulumi.StringOutput {
	return o.ApplyT(func(v *LiveEvent) pulumi.StringOutput { return v.Created }).(pulumi.StringOutput)
}

// Live event cross site access policies.
func (o LiveEventOutput) CrossSiteAccessPolicies() CrossSiteAccessPoliciesResponsePtrOutput {
	return o.ApplyT(func(v *LiveEvent) CrossSiteAccessPoliciesResponsePtrOutput { return v.CrossSiteAccessPolicies }).(CrossSiteAccessPoliciesResponsePtrOutput)
}

// A description for the live event.
func (o LiveEventOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LiveEvent) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Encoding settings for the live event. It configures whether a live encoder is used for the live event and settings for the live encoder if it is used.
func (o LiveEventOutput) Encoding() LiveEventEncodingResponsePtrOutput {
	return o.ApplyT(func(v *LiveEvent) LiveEventEncodingResponsePtrOutput { return v.Encoding }).(LiveEventEncodingResponsePtrOutput)
}

// When useStaticHostname is set to true, the hostnamePrefix specifies the first part of the hostname assigned to the live event preview and ingest endpoints. The final hostname would be a combination of this prefix, the media service account name and a short code for the Azure Media Services data center.
func (o LiveEventOutput) HostnamePrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LiveEvent) pulumi.StringPtrOutput { return v.HostnamePrefix }).(pulumi.StringPtrOutput)
}

// Live event input settings. It defines how the live event receives input from a contribution encoder.
func (o LiveEventOutput) Input() LiveEventInputResponseOutput {
	return o.ApplyT(func(v *LiveEvent) LiveEventInputResponseOutput { return v.Input }).(LiveEventInputResponseOutput)
}

// The last modified time of the live event.
func (o LiveEventOutput) LastModified() pulumi.StringOutput {
	return o.ApplyT(func(v *LiveEvent) pulumi.StringOutput { return v.LastModified }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LiveEventOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *LiveEvent) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LiveEventOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *LiveEvent) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Live event preview settings. Preview allows live event producers to preview the live streaming content without creating any live output.
func (o LiveEventOutput) Preview() LiveEventPreviewResponsePtrOutput {
	return o.ApplyT(func(v *LiveEvent) LiveEventPreviewResponsePtrOutput { return v.Preview }).(LiveEventPreviewResponsePtrOutput)
}

// The provisioning state of the live event.
func (o LiveEventOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *LiveEvent) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The resource state of the live event. See https://go.microsoft.com/fwlink/?linkid=2139012 for more information.
func (o LiveEventOutput) ResourceState() pulumi.StringOutput {
	return o.ApplyT(func(v *LiveEvent) pulumi.StringOutput { return v.ResourceState }).(pulumi.StringOutput)
}

// The options to use for the LiveEvent. This value is specified at creation time and cannot be updated. The valid values for the array entry values are 'Default' and 'LowLatency'.
func (o LiveEventOutput) StreamOptions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LiveEvent) pulumi.StringArrayOutput { return v.StreamOptions }).(pulumi.StringArrayOutput)
}

// The system metadata relating to this resource.
func (o LiveEventOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *LiveEvent) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LiveEventOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *LiveEvent) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Live transcription settings for the live event. See https://go.microsoft.com/fwlink/?linkid=2133742 for more information about the live transcription feature.
func (o LiveEventOutput) Transcriptions() LiveEventTranscriptionResponseArrayOutput {
	return o.ApplyT(func(v *LiveEvent) LiveEventTranscriptionResponseArrayOutput { return v.Transcriptions }).(LiveEventTranscriptionResponseArrayOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LiveEventOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *LiveEvent) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Specifies whether a static hostname would be assigned to the live event preview and ingest endpoints. This value can only be updated if the live event is in Standby state
func (o LiveEventOutput) UseStaticHostname() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LiveEvent) pulumi.BoolPtrOutput { return v.UseStaticHostname }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LiveEventOutput{})
}
