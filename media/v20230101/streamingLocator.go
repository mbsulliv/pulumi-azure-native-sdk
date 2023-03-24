// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A Streaming Locator resource
type StreamingLocator struct {
	pulumi.CustomResourceState

	// Alternative Media ID of this Streaming Locator
	AlternativeMediaId pulumi.StringPtrOutput `pulumi:"alternativeMediaId"`
	// Asset Name
	AssetName pulumi.StringOutput `pulumi:"assetName"`
	// The ContentKeys used by this Streaming Locator.
	ContentKeys StreamingLocatorContentKeyResponseArrayOutput `pulumi:"contentKeys"`
	// The creation time of the Streaming Locator.
	Created pulumi.StringOutput `pulumi:"created"`
	// Name of the default ContentKeyPolicy used by this Streaming Locator.
	DefaultContentKeyPolicyName pulumi.StringPtrOutput `pulumi:"defaultContentKeyPolicyName"`
	// The end time of the Streaming Locator.
	EndTime pulumi.StringPtrOutput `pulumi:"endTime"`
	// A list of asset or account filters which apply to this streaming locator
	Filters pulumi.StringArrayOutput `pulumi:"filters"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The start time of the Streaming Locator.
	StartTime pulumi.StringPtrOutput `pulumi:"startTime"`
	// The StreamingLocatorId of the Streaming Locator.
	StreamingLocatorId pulumi.StringPtrOutput `pulumi:"streamingLocatorId"`
	// Name of the Streaming Policy used by this Streaming Locator. Either specify the name of Streaming Policy you created or use one of the predefined Streaming Policies. The predefined Streaming Policies available are: 'Predefined_DownloadOnly', 'Predefined_ClearStreamingOnly', 'Predefined_DownloadAndClearStreaming', 'Predefined_ClearKey', 'Predefined_MultiDrmCencStreaming' and 'Predefined_MultiDrmStreaming'
	StreamingPolicyName pulumi.StringOutput `pulumi:"streamingPolicyName"`
	// The system metadata relating to this resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewStreamingLocator registers a new resource with the given unique name, arguments, and options.
func NewStreamingLocator(ctx *pulumi.Context,
	name string, args *StreamingLocatorArgs, opts ...pulumi.ResourceOption) (*StreamingLocator, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.AssetName == nil {
		return nil, errors.New("invalid value for required argument 'AssetName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.StreamingPolicyName == nil {
		return nil, errors.New("invalid value for required argument 'StreamingPolicyName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:media:StreamingLocator"),
		},
		{
			Type: pulumi.String("azure-native:media/v20180330preview:StreamingLocator"),
		},
		{
			Type: pulumi.String("azure-native:media/v20180601preview:StreamingLocator"),
		},
		{
			Type: pulumi.String("azure-native:media/v20180701:StreamingLocator"),
		},
		{
			Type: pulumi.String("azure-native:media/v20200501:StreamingLocator"),
		},
		{
			Type: pulumi.String("azure-native:media/v20210601:StreamingLocator"),
		},
		{
			Type: pulumi.String("azure-native:media/v20211101:StreamingLocator"),
		},
		{
			Type: pulumi.String("azure-native:media/v20220801:StreamingLocator"),
		},
	})
	opts = append(opts, aliases)
	var resource StreamingLocator
	err := ctx.RegisterResource("azure-native:media/v20230101:StreamingLocator", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStreamingLocator gets an existing StreamingLocator resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStreamingLocator(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StreamingLocatorState, opts ...pulumi.ResourceOption) (*StreamingLocator, error) {
	var resource StreamingLocator
	err := ctx.ReadResource("azure-native:media/v20230101:StreamingLocator", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StreamingLocator resources.
type streamingLocatorState struct {
}

type StreamingLocatorState struct {
}

func (StreamingLocatorState) ElementType() reflect.Type {
	return reflect.TypeOf((*streamingLocatorState)(nil)).Elem()
}

type streamingLocatorArgs struct {
	// The Media Services account name.
	AccountName string `pulumi:"accountName"`
	// Alternative Media ID of this Streaming Locator
	AlternativeMediaId *string `pulumi:"alternativeMediaId"`
	// Asset Name
	AssetName string `pulumi:"assetName"`
	// The ContentKeys used by this Streaming Locator.
	ContentKeys []StreamingLocatorContentKey `pulumi:"contentKeys"`
	// Name of the default ContentKeyPolicy used by this Streaming Locator.
	DefaultContentKeyPolicyName *string `pulumi:"defaultContentKeyPolicyName"`
	// The end time of the Streaming Locator.
	EndTime *string `pulumi:"endTime"`
	// A list of asset or account filters which apply to this streaming locator
	Filters []string `pulumi:"filters"`
	// The name of the resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The start time of the Streaming Locator.
	StartTime *string `pulumi:"startTime"`
	// The StreamingLocatorId of the Streaming Locator.
	StreamingLocatorId *string `pulumi:"streamingLocatorId"`
	// The Streaming Locator name.
	StreamingLocatorName *string `pulumi:"streamingLocatorName"`
	// Name of the Streaming Policy used by this Streaming Locator. Either specify the name of Streaming Policy you created or use one of the predefined Streaming Policies. The predefined Streaming Policies available are: 'Predefined_DownloadOnly', 'Predefined_ClearStreamingOnly', 'Predefined_DownloadAndClearStreaming', 'Predefined_ClearKey', 'Predefined_MultiDrmCencStreaming' and 'Predefined_MultiDrmStreaming'
	StreamingPolicyName string `pulumi:"streamingPolicyName"`
}

// The set of arguments for constructing a StreamingLocator resource.
type StreamingLocatorArgs struct {
	// The Media Services account name.
	AccountName pulumi.StringInput
	// Alternative Media ID of this Streaming Locator
	AlternativeMediaId pulumi.StringPtrInput
	// Asset Name
	AssetName pulumi.StringInput
	// The ContentKeys used by this Streaming Locator.
	ContentKeys StreamingLocatorContentKeyArrayInput
	// Name of the default ContentKeyPolicy used by this Streaming Locator.
	DefaultContentKeyPolicyName pulumi.StringPtrInput
	// The end time of the Streaming Locator.
	EndTime pulumi.StringPtrInput
	// A list of asset or account filters which apply to this streaming locator
	Filters pulumi.StringArrayInput
	// The name of the resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput
	// The start time of the Streaming Locator.
	StartTime pulumi.StringPtrInput
	// The StreamingLocatorId of the Streaming Locator.
	StreamingLocatorId pulumi.StringPtrInput
	// The Streaming Locator name.
	StreamingLocatorName pulumi.StringPtrInput
	// Name of the Streaming Policy used by this Streaming Locator. Either specify the name of Streaming Policy you created or use one of the predefined Streaming Policies. The predefined Streaming Policies available are: 'Predefined_DownloadOnly', 'Predefined_ClearStreamingOnly', 'Predefined_DownloadAndClearStreaming', 'Predefined_ClearKey', 'Predefined_MultiDrmCencStreaming' and 'Predefined_MultiDrmStreaming'
	StreamingPolicyName pulumi.StringInput
}

func (StreamingLocatorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*streamingLocatorArgs)(nil)).Elem()
}

type StreamingLocatorInput interface {
	pulumi.Input

	ToStreamingLocatorOutput() StreamingLocatorOutput
	ToStreamingLocatorOutputWithContext(ctx context.Context) StreamingLocatorOutput
}

func (*StreamingLocator) ElementType() reflect.Type {
	return reflect.TypeOf((**StreamingLocator)(nil)).Elem()
}

func (i *StreamingLocator) ToStreamingLocatorOutput() StreamingLocatorOutput {
	return i.ToStreamingLocatorOutputWithContext(context.Background())
}

func (i *StreamingLocator) ToStreamingLocatorOutputWithContext(ctx context.Context) StreamingLocatorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamingLocatorOutput)
}

type StreamingLocatorOutput struct{ *pulumi.OutputState }

func (StreamingLocatorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StreamingLocator)(nil)).Elem()
}

func (o StreamingLocatorOutput) ToStreamingLocatorOutput() StreamingLocatorOutput {
	return o
}

func (o StreamingLocatorOutput) ToStreamingLocatorOutputWithContext(ctx context.Context) StreamingLocatorOutput {
	return o
}

// Alternative Media ID of this Streaming Locator
func (o StreamingLocatorOutput) AlternativeMediaId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StreamingLocator) pulumi.StringPtrOutput { return v.AlternativeMediaId }).(pulumi.StringPtrOutput)
}

// Asset Name
func (o StreamingLocatorOutput) AssetName() pulumi.StringOutput {
	return o.ApplyT(func(v *StreamingLocator) pulumi.StringOutput { return v.AssetName }).(pulumi.StringOutput)
}

// The ContentKeys used by this Streaming Locator.
func (o StreamingLocatorOutput) ContentKeys() StreamingLocatorContentKeyResponseArrayOutput {
	return o.ApplyT(func(v *StreamingLocator) StreamingLocatorContentKeyResponseArrayOutput { return v.ContentKeys }).(StreamingLocatorContentKeyResponseArrayOutput)
}

// The creation time of the Streaming Locator.
func (o StreamingLocatorOutput) Created() pulumi.StringOutput {
	return o.ApplyT(func(v *StreamingLocator) pulumi.StringOutput { return v.Created }).(pulumi.StringOutput)
}

// Name of the default ContentKeyPolicy used by this Streaming Locator.
func (o StreamingLocatorOutput) DefaultContentKeyPolicyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StreamingLocator) pulumi.StringPtrOutput { return v.DefaultContentKeyPolicyName }).(pulumi.StringPtrOutput)
}

// The end time of the Streaming Locator.
func (o StreamingLocatorOutput) EndTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StreamingLocator) pulumi.StringPtrOutput { return v.EndTime }).(pulumi.StringPtrOutput)
}

// A list of asset or account filters which apply to this streaming locator
func (o StreamingLocatorOutput) Filters() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *StreamingLocator) pulumi.StringArrayOutput { return v.Filters }).(pulumi.StringArrayOutput)
}

// The name of the resource
func (o StreamingLocatorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *StreamingLocator) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The start time of the Streaming Locator.
func (o StreamingLocatorOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StreamingLocator) pulumi.StringPtrOutput { return v.StartTime }).(pulumi.StringPtrOutput)
}

// The StreamingLocatorId of the Streaming Locator.
func (o StreamingLocatorOutput) StreamingLocatorId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StreamingLocator) pulumi.StringPtrOutput { return v.StreamingLocatorId }).(pulumi.StringPtrOutput)
}

// Name of the Streaming Policy used by this Streaming Locator. Either specify the name of Streaming Policy you created or use one of the predefined Streaming Policies. The predefined Streaming Policies available are: 'Predefined_DownloadOnly', 'Predefined_ClearStreamingOnly', 'Predefined_DownloadAndClearStreaming', 'Predefined_ClearKey', 'Predefined_MultiDrmCencStreaming' and 'Predefined_MultiDrmStreaming'
func (o StreamingLocatorOutput) StreamingPolicyName() pulumi.StringOutput {
	return o.ApplyT(func(v *StreamingLocator) pulumi.StringOutput { return v.StreamingPolicyName }).(pulumi.StringOutput)
}

// The system metadata relating to this resource.
func (o StreamingLocatorOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *StreamingLocator) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o StreamingLocatorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *StreamingLocator) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(StreamingLocatorOutput{})
}