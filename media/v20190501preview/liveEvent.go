// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20190501preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The Live Event.
type LiveEvent struct {
	pulumi.CustomResourceState

	// The exact time the Live Event was created.
	Created pulumi.StringOutput `pulumi:"created"`
	// The Live Event access policies.
	CrossSiteAccessPolicies CrossSiteAccessPoliciesResponsePtrOutput `pulumi:"crossSiteAccessPolicies"`
	// The Live Event description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The Live Event encoding.
	Encoding LiveEventEncodingResponsePtrOutput `pulumi:"encoding"`
	// The Live Event input.
	Input LiveEventInputResponseOutput `pulumi:"input"`
	// The exact time the Live Event was last modified.
	LastModified pulumi.StringOutput `pulumi:"lastModified"`
	// The Azure Region of the resource.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The Live Event preview.
	Preview LiveEventPreviewResponsePtrOutput `pulumi:"preview"`
	// The provisioning state of the Live Event.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The resource state of the Live Event.
	ResourceState pulumi.StringOutput `pulumi:"resourceState"`
	// The options to use for the LiveEvent.  This value is specified at creation time and cannot be updated.
	StreamOptions pulumi.StringArrayOutput `pulumi:"streamOptions"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The Live Event transcription.
	Transcriptions LiveEventTranscriptionResponseArrayOutput `pulumi:"transcriptions"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
	// Specifies whether to use a vanity url with the Live Event.  This value is specified at creation time and cannot be updated.
	VanityUrl pulumi.BoolPtrOutput `pulumi:"vanityUrl"`
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
			Type: pulumi.String("azure-native:media:LiveEvent"),
		},
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
	err := ctx.RegisterResource("azure-native:media/v20190501preview:LiveEvent", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure-native:media/v20190501preview:LiveEvent", name, id, state, &resource, opts...)
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
	// The Live Event access policies.
	CrossSiteAccessPolicies *CrossSiteAccessPolicies `pulumi:"crossSiteAccessPolicies"`
	// The Live Event description.
	Description *string `pulumi:"description"`
	// The Live Event encoding.
	Encoding *LiveEventEncoding `pulumi:"encoding"`
	// The Live Event input.
	Input LiveEventInputType `pulumi:"input"`
	// The name of the Live Event.
	LiveEventName *string `pulumi:"liveEventName"`
	// The Azure Region of the resource.
	Location *string `pulumi:"location"`
	// The Live Event preview.
	Preview *LiveEventPreview `pulumi:"preview"`
	// The name of the resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The options to use for the LiveEvent.  This value is specified at creation time and cannot be updated.
	StreamOptions []string `pulumi:"streamOptions"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The Live Event transcription.
	Transcriptions []LiveEventTranscription `pulumi:"transcriptions"`
	// Specifies whether to use a vanity url with the Live Event.  This value is specified at creation time and cannot be updated.
	VanityUrl *bool `pulumi:"vanityUrl"`
}

// The set of arguments for constructing a LiveEvent resource.
type LiveEventArgs struct {
	// The Media Services account name.
	AccountName pulumi.StringInput
	// The flag indicates if the resource should be automatically started on creation.
	AutoStart pulumi.BoolPtrInput
	// The Live Event access policies.
	CrossSiteAccessPolicies CrossSiteAccessPoliciesPtrInput
	// The Live Event description.
	Description pulumi.StringPtrInput
	// The Live Event encoding.
	Encoding LiveEventEncodingPtrInput
	// The Live Event input.
	Input LiveEventInputTypeInput
	// The name of the Live Event.
	LiveEventName pulumi.StringPtrInput
	// The Azure Region of the resource.
	Location pulumi.StringPtrInput
	// The Live Event preview.
	Preview LiveEventPreviewPtrInput
	// The name of the resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput
	// The options to use for the LiveEvent.  This value is specified at creation time and cannot be updated.
	StreamOptions pulumi.StringArrayInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The Live Event transcription.
	Transcriptions LiveEventTranscriptionArrayInput
	// Specifies whether to use a vanity url with the Live Event.  This value is specified at creation time and cannot be updated.
	VanityUrl pulumi.BoolPtrInput
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

// The exact time the Live Event was created.
func (o LiveEventOutput) Created() pulumi.StringOutput {
	return o.ApplyT(func(v *LiveEvent) pulumi.StringOutput { return v.Created }).(pulumi.StringOutput)
}

// The Live Event access policies.
func (o LiveEventOutput) CrossSiteAccessPolicies() CrossSiteAccessPoliciesResponsePtrOutput {
	return o.ApplyT(func(v *LiveEvent) CrossSiteAccessPoliciesResponsePtrOutput { return v.CrossSiteAccessPolicies }).(CrossSiteAccessPoliciesResponsePtrOutput)
}

// The Live Event description.
func (o LiveEventOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LiveEvent) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The Live Event encoding.
func (o LiveEventOutput) Encoding() LiveEventEncodingResponsePtrOutput {
	return o.ApplyT(func(v *LiveEvent) LiveEventEncodingResponsePtrOutput { return v.Encoding }).(LiveEventEncodingResponsePtrOutput)
}

// The Live Event input.
func (o LiveEventOutput) Input() LiveEventInputResponseOutput {
	return o.ApplyT(func(v *LiveEvent) LiveEventInputResponseOutput { return v.Input }).(LiveEventInputResponseOutput)
}

// The exact time the Live Event was last modified.
func (o LiveEventOutput) LastModified() pulumi.StringOutput {
	return o.ApplyT(func(v *LiveEvent) pulumi.StringOutput { return v.LastModified }).(pulumi.StringOutput)
}

// The Azure Region of the resource.
func (o LiveEventOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LiveEvent) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the resource.
func (o LiveEventOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *LiveEvent) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The Live Event preview.
func (o LiveEventOutput) Preview() LiveEventPreviewResponsePtrOutput {
	return o.ApplyT(func(v *LiveEvent) LiveEventPreviewResponsePtrOutput { return v.Preview }).(LiveEventPreviewResponsePtrOutput)
}

// The provisioning state of the Live Event.
func (o LiveEventOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *LiveEvent) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The resource state of the Live Event.
func (o LiveEventOutput) ResourceState() pulumi.StringOutput {
	return o.ApplyT(func(v *LiveEvent) pulumi.StringOutput { return v.ResourceState }).(pulumi.StringOutput)
}

// The options to use for the LiveEvent.  This value is specified at creation time and cannot be updated.
func (o LiveEventOutput) StreamOptions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LiveEvent) pulumi.StringArrayOutput { return v.StreamOptions }).(pulumi.StringArrayOutput)
}

// Resource tags.
func (o LiveEventOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *LiveEvent) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The Live Event transcription.
func (o LiveEventOutput) Transcriptions() LiveEventTranscriptionResponseArrayOutput {
	return o.ApplyT(func(v *LiveEvent) LiveEventTranscriptionResponseArrayOutput { return v.Transcriptions }).(LiveEventTranscriptionResponseArrayOutput)
}

// The type of the resource.
func (o LiveEventOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *LiveEvent) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Specifies whether to use a vanity url with the Live Event.  This value is specified at creation time and cannot be updated.
func (o LiveEventOutput) VanityUrl() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LiveEvent) pulumi.BoolPtrOutput { return v.VanityUrl }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LiveEventOutput{})
}
