// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20180330preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A Streaming Locator resource
type StreamingLocator struct {
	pulumi.CustomResourceState

	// Asset Name
	AssetName pulumi.StringOutput `pulumi:"assetName"`
	// ContentKeys used by this Streaming Locator
	ContentKeys StreamingLocatorUserDefinedContentKeyResponseArrayOutput `pulumi:"contentKeys"`
	// Creation time of Streaming Locator
	Created pulumi.StringOutput `pulumi:"created"`
	// Default ContentKeyPolicy used by this Streaming Locator
	DefaultContentKeyPolicyName pulumi.StringPtrOutput `pulumi:"defaultContentKeyPolicyName"`
	// EndTime of Streaming Locator
	EndTime pulumi.StringPtrOutput `pulumi:"endTime"`
	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// StartTime of Streaming Locator
	StartTime pulumi.StringPtrOutput `pulumi:"startTime"`
	// StreamingLocatorId of Streaming Locator
	StreamingLocatorId pulumi.StringPtrOutput `pulumi:"streamingLocatorId"`
	// Streaming policy name used by this streaming locator. Either specify the name of streaming policy you created or use one of the predefined streaming polices. The predefined streaming policies available are: 'Predefined_DownloadOnly', 'Predefined_ClearStreamingOnly', 'Predefined_DownloadAndClearStreaming', 'Predefined_ClearKey', 'Predefined_SecureStreaming' and 'Predefined_SecureStreamingWithFairPlay'
	StreamingPolicyName pulumi.StringOutput `pulumi:"streamingPolicyName"`
	// The type of the resource.
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
		{
			Type: pulumi.String("azure-native:media/v20230101:StreamingLocator"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource StreamingLocator
	err := ctx.RegisterResource("azure-native:media/v20180330preview:StreamingLocator", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure-native:media/v20180330preview:StreamingLocator", name, id, state, &resource, opts...)
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
	// Asset Name
	AssetName string `pulumi:"assetName"`
	// ContentKeys used by this Streaming Locator
	ContentKeys []StreamingLocatorUserDefinedContentKey `pulumi:"contentKeys"`
	// Default ContentKeyPolicy used by this Streaming Locator
	DefaultContentKeyPolicyName *string `pulumi:"defaultContentKeyPolicyName"`
	// EndTime of Streaming Locator
	EndTime *string `pulumi:"endTime"`
	// The name of the resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// StartTime of Streaming Locator
	StartTime *string `pulumi:"startTime"`
	// StreamingLocatorId of Streaming Locator
	StreamingLocatorId *string `pulumi:"streamingLocatorId"`
	// The Streaming Locator name.
	StreamingLocatorName *string `pulumi:"streamingLocatorName"`
	// Streaming policy name used by this streaming locator. Either specify the name of streaming policy you created or use one of the predefined streaming polices. The predefined streaming policies available are: 'Predefined_DownloadOnly', 'Predefined_ClearStreamingOnly', 'Predefined_DownloadAndClearStreaming', 'Predefined_ClearKey', 'Predefined_SecureStreaming' and 'Predefined_SecureStreamingWithFairPlay'
	StreamingPolicyName string `pulumi:"streamingPolicyName"`
}

// The set of arguments for constructing a StreamingLocator resource.
type StreamingLocatorArgs struct {
	// The Media Services account name.
	AccountName pulumi.StringInput
	// Asset Name
	AssetName pulumi.StringInput
	// ContentKeys used by this Streaming Locator
	ContentKeys StreamingLocatorUserDefinedContentKeyArrayInput
	// Default ContentKeyPolicy used by this Streaming Locator
	DefaultContentKeyPolicyName pulumi.StringPtrInput
	// EndTime of Streaming Locator
	EndTime pulumi.StringPtrInput
	// The name of the resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput
	// StartTime of Streaming Locator
	StartTime pulumi.StringPtrInput
	// StreamingLocatorId of Streaming Locator
	StreamingLocatorId pulumi.StringPtrInput
	// The Streaming Locator name.
	StreamingLocatorName pulumi.StringPtrInput
	// Streaming policy name used by this streaming locator. Either specify the name of streaming policy you created or use one of the predefined streaming polices. The predefined streaming policies available are: 'Predefined_DownloadOnly', 'Predefined_ClearStreamingOnly', 'Predefined_DownloadAndClearStreaming', 'Predefined_ClearKey', 'Predefined_SecureStreaming' and 'Predefined_SecureStreamingWithFairPlay'
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

// Asset Name
func (o StreamingLocatorOutput) AssetName() pulumi.StringOutput {
	return o.ApplyT(func(v *StreamingLocator) pulumi.StringOutput { return v.AssetName }).(pulumi.StringOutput)
}

// ContentKeys used by this Streaming Locator
func (o StreamingLocatorOutput) ContentKeys() StreamingLocatorUserDefinedContentKeyResponseArrayOutput {
	return o.ApplyT(func(v *StreamingLocator) StreamingLocatorUserDefinedContentKeyResponseArrayOutput {
		return v.ContentKeys
	}).(StreamingLocatorUserDefinedContentKeyResponseArrayOutput)
}

// Creation time of Streaming Locator
func (o StreamingLocatorOutput) Created() pulumi.StringOutput {
	return o.ApplyT(func(v *StreamingLocator) pulumi.StringOutput { return v.Created }).(pulumi.StringOutput)
}

// Default ContentKeyPolicy used by this Streaming Locator
func (o StreamingLocatorOutput) DefaultContentKeyPolicyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StreamingLocator) pulumi.StringPtrOutput { return v.DefaultContentKeyPolicyName }).(pulumi.StringPtrOutput)
}

// EndTime of Streaming Locator
func (o StreamingLocatorOutput) EndTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StreamingLocator) pulumi.StringPtrOutput { return v.EndTime }).(pulumi.StringPtrOutput)
}

// The name of the resource.
func (o StreamingLocatorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *StreamingLocator) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// StartTime of Streaming Locator
func (o StreamingLocatorOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StreamingLocator) pulumi.StringPtrOutput { return v.StartTime }).(pulumi.StringPtrOutput)
}

// StreamingLocatorId of Streaming Locator
func (o StreamingLocatorOutput) StreamingLocatorId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StreamingLocator) pulumi.StringPtrOutput { return v.StreamingLocatorId }).(pulumi.StringPtrOutput)
}

// Streaming policy name used by this streaming locator. Either specify the name of streaming policy you created or use one of the predefined streaming polices. The predefined streaming policies available are: 'Predefined_DownloadOnly', 'Predefined_ClearStreamingOnly', 'Predefined_DownloadAndClearStreaming', 'Predefined_ClearKey', 'Predefined_SecureStreaming' and 'Predefined_SecureStreamingWithFairPlay'
func (o StreamingLocatorOutput) StreamingPolicyName() pulumi.StringOutput {
	return o.ApplyT(func(v *StreamingLocator) pulumi.StringOutput { return v.StreamingPolicyName }).(pulumi.StringOutput)
}

// The type of the resource.
func (o StreamingLocatorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *StreamingLocator) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(StreamingLocatorOutput{})
}
