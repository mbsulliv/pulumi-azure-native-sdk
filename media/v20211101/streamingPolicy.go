// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20211101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A Streaming Policy resource
type StreamingPolicy struct {
	pulumi.CustomResourceState

	// Configuration of CommonEncryptionCbcs
	CommonEncryptionCbcs CommonEncryptionCbcsResponsePtrOutput `pulumi:"commonEncryptionCbcs"`
	// Configuration of CommonEncryptionCenc
	CommonEncryptionCenc CommonEncryptionCencResponsePtrOutput `pulumi:"commonEncryptionCenc"`
	// Creation time of Streaming Policy
	Created pulumi.StringOutput `pulumi:"created"`
	// Default ContentKey used by current Streaming Policy
	DefaultContentKeyPolicyName pulumi.StringPtrOutput `pulumi:"defaultContentKeyPolicyName"`
	// Configuration of EnvelopeEncryption
	EnvelopeEncryption EnvelopeEncryptionResponsePtrOutput `pulumi:"envelopeEncryption"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Configurations of NoEncryption
	NoEncryption NoEncryptionResponsePtrOutput `pulumi:"noEncryption"`
	// The system metadata relating to this resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewStreamingPolicy registers a new resource with the given unique name, arguments, and options.
func NewStreamingPolicy(ctx *pulumi.Context,
	name string, args *StreamingPolicyArgs, opts ...pulumi.ResourceOption) (*StreamingPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:media:StreamingPolicy"),
		},
		{
			Type: pulumi.String("azure-native:media/v20180330preview:StreamingPolicy"),
		},
		{
			Type: pulumi.String("azure-native:media/v20180601preview:StreamingPolicy"),
		},
		{
			Type: pulumi.String("azure-native:media/v20180701:StreamingPolicy"),
		},
		{
			Type: pulumi.String("azure-native:media/v20200501:StreamingPolicy"),
		},
		{
			Type: pulumi.String("azure-native:media/v20210601:StreamingPolicy"),
		},
		{
			Type: pulumi.String("azure-native:media/v20220801:StreamingPolicy"),
		},
		{
			Type: pulumi.String("azure-native:media/v20230101:StreamingPolicy"),
		},
	})
	opts = append(opts, aliases)
	var resource StreamingPolicy
	err := ctx.RegisterResource("azure-native:media/v20211101:StreamingPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStreamingPolicy gets an existing StreamingPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStreamingPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StreamingPolicyState, opts ...pulumi.ResourceOption) (*StreamingPolicy, error) {
	var resource StreamingPolicy
	err := ctx.ReadResource("azure-native:media/v20211101:StreamingPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StreamingPolicy resources.
type streamingPolicyState struct {
}

type StreamingPolicyState struct {
}

func (StreamingPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*streamingPolicyState)(nil)).Elem()
}

type streamingPolicyArgs struct {
	// The Media Services account name.
	AccountName string `pulumi:"accountName"`
	// Configuration of CommonEncryptionCbcs
	CommonEncryptionCbcs *CommonEncryptionCbcs `pulumi:"commonEncryptionCbcs"`
	// Configuration of CommonEncryptionCenc
	CommonEncryptionCenc *CommonEncryptionCenc `pulumi:"commonEncryptionCenc"`
	// Default ContentKey used by current Streaming Policy
	DefaultContentKeyPolicyName *string `pulumi:"defaultContentKeyPolicyName"`
	// Configuration of EnvelopeEncryption
	EnvelopeEncryption *EnvelopeEncryption `pulumi:"envelopeEncryption"`
	// Configurations of NoEncryption
	NoEncryption *NoEncryption `pulumi:"noEncryption"`
	// The name of the resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The Streaming Policy name.
	StreamingPolicyName *string `pulumi:"streamingPolicyName"`
}

// The set of arguments for constructing a StreamingPolicy resource.
type StreamingPolicyArgs struct {
	// The Media Services account name.
	AccountName pulumi.StringInput
	// Configuration of CommonEncryptionCbcs
	CommonEncryptionCbcs CommonEncryptionCbcsPtrInput
	// Configuration of CommonEncryptionCenc
	CommonEncryptionCenc CommonEncryptionCencPtrInput
	// Default ContentKey used by current Streaming Policy
	DefaultContentKeyPolicyName pulumi.StringPtrInput
	// Configuration of EnvelopeEncryption
	EnvelopeEncryption EnvelopeEncryptionPtrInput
	// Configurations of NoEncryption
	NoEncryption NoEncryptionPtrInput
	// The name of the resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput
	// The Streaming Policy name.
	StreamingPolicyName pulumi.StringPtrInput
}

func (StreamingPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*streamingPolicyArgs)(nil)).Elem()
}

type StreamingPolicyInput interface {
	pulumi.Input

	ToStreamingPolicyOutput() StreamingPolicyOutput
	ToStreamingPolicyOutputWithContext(ctx context.Context) StreamingPolicyOutput
}

func (*StreamingPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**StreamingPolicy)(nil)).Elem()
}

func (i *StreamingPolicy) ToStreamingPolicyOutput() StreamingPolicyOutput {
	return i.ToStreamingPolicyOutputWithContext(context.Background())
}

func (i *StreamingPolicy) ToStreamingPolicyOutputWithContext(ctx context.Context) StreamingPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamingPolicyOutput)
}

type StreamingPolicyOutput struct{ *pulumi.OutputState }

func (StreamingPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StreamingPolicy)(nil)).Elem()
}

func (o StreamingPolicyOutput) ToStreamingPolicyOutput() StreamingPolicyOutput {
	return o
}

func (o StreamingPolicyOutput) ToStreamingPolicyOutputWithContext(ctx context.Context) StreamingPolicyOutput {
	return o
}

// Configuration of CommonEncryptionCbcs
func (o StreamingPolicyOutput) CommonEncryptionCbcs() CommonEncryptionCbcsResponsePtrOutput {
	return o.ApplyT(func(v *StreamingPolicy) CommonEncryptionCbcsResponsePtrOutput { return v.CommonEncryptionCbcs }).(CommonEncryptionCbcsResponsePtrOutput)
}

// Configuration of CommonEncryptionCenc
func (o StreamingPolicyOutput) CommonEncryptionCenc() CommonEncryptionCencResponsePtrOutput {
	return o.ApplyT(func(v *StreamingPolicy) CommonEncryptionCencResponsePtrOutput { return v.CommonEncryptionCenc }).(CommonEncryptionCencResponsePtrOutput)
}

// Creation time of Streaming Policy
func (o StreamingPolicyOutput) Created() pulumi.StringOutput {
	return o.ApplyT(func(v *StreamingPolicy) pulumi.StringOutput { return v.Created }).(pulumi.StringOutput)
}

// Default ContentKey used by current Streaming Policy
func (o StreamingPolicyOutput) DefaultContentKeyPolicyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StreamingPolicy) pulumi.StringPtrOutput { return v.DefaultContentKeyPolicyName }).(pulumi.StringPtrOutput)
}

// Configuration of EnvelopeEncryption
func (o StreamingPolicyOutput) EnvelopeEncryption() EnvelopeEncryptionResponsePtrOutput {
	return o.ApplyT(func(v *StreamingPolicy) EnvelopeEncryptionResponsePtrOutput { return v.EnvelopeEncryption }).(EnvelopeEncryptionResponsePtrOutput)
}

// The name of the resource
func (o StreamingPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *StreamingPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Configurations of NoEncryption
func (o StreamingPolicyOutput) NoEncryption() NoEncryptionResponsePtrOutput {
	return o.ApplyT(func(v *StreamingPolicy) NoEncryptionResponsePtrOutput { return v.NoEncryption }).(NoEncryptionResponsePtrOutput)
}

// The system metadata relating to this resource.
func (o StreamingPolicyOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *StreamingPolicy) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o StreamingPolicyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *StreamingPolicy) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(StreamingPolicyOutput{})
}
