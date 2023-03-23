// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20200501

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A Content Key Policy resource.
type ContentKeyPolicy struct {
	pulumi.CustomResourceState

	// The creation date of the Policy
	Created pulumi.StringOutput `pulumi:"created"`
	// A description for the Policy.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The last modified date of the Policy
	LastModified pulumi.StringOutput `pulumi:"lastModified"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The Key Policy options.
	Options ContentKeyPolicyOptionResponseArrayOutput `pulumi:"options"`
	// The legacy Policy ID.
	PolicyId pulumi.StringOutput `pulumi:"policyId"`
	// The system metadata relating to this resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewContentKeyPolicy registers a new resource with the given unique name, arguments, and options.
func NewContentKeyPolicy(ctx *pulumi.Context,
	name string, args *ContentKeyPolicyArgs, opts ...pulumi.ResourceOption) (*ContentKeyPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.Options == nil {
		return nil, errors.New("invalid value for required argument 'Options'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:media:ContentKeyPolicy"),
		},
		{
			Type: pulumi.String("azure-native:media/v20180330preview:ContentKeyPolicy"),
		},
		{
			Type: pulumi.String("azure-native:media/v20180601preview:ContentKeyPolicy"),
		},
		{
			Type: pulumi.String("azure-native:media/v20180701:ContentKeyPolicy"),
		},
		{
			Type: pulumi.String("azure-native:media/v20210601:ContentKeyPolicy"),
		},
		{
			Type: pulumi.String("azure-native:media/v20211101:ContentKeyPolicy"),
		},
		{
			Type: pulumi.String("azure-native:media/v20220801:ContentKeyPolicy"),
		},
		{
			Type: pulumi.String("azure-native:media/v20230101:ContentKeyPolicy"),
		},
	})
	opts = append(opts, aliases)
	var resource ContentKeyPolicy
	err := ctx.RegisterResource("azure-native:media/v20200501:ContentKeyPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetContentKeyPolicy gets an existing ContentKeyPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetContentKeyPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ContentKeyPolicyState, opts ...pulumi.ResourceOption) (*ContentKeyPolicy, error) {
	var resource ContentKeyPolicy
	err := ctx.ReadResource("azure-native:media/v20200501:ContentKeyPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ContentKeyPolicy resources.
type contentKeyPolicyState struct {
}

type ContentKeyPolicyState struct {
}

func (ContentKeyPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*contentKeyPolicyState)(nil)).Elem()
}

type contentKeyPolicyArgs struct {
	// The Media Services account name.
	AccountName string `pulumi:"accountName"`
	// The Content Key Policy name.
	ContentKeyPolicyName *string `pulumi:"contentKeyPolicyName"`
	// A description for the Policy.
	Description *string `pulumi:"description"`
	// The Key Policy options.
	Options []ContentKeyPolicyOption `pulumi:"options"`
	// The name of the resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a ContentKeyPolicy resource.
type ContentKeyPolicyArgs struct {
	// The Media Services account name.
	AccountName pulumi.StringInput
	// The Content Key Policy name.
	ContentKeyPolicyName pulumi.StringPtrInput
	// A description for the Policy.
	Description pulumi.StringPtrInput
	// The Key Policy options.
	Options ContentKeyPolicyOptionArrayInput
	// The name of the resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput
}

func (ContentKeyPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*contentKeyPolicyArgs)(nil)).Elem()
}

type ContentKeyPolicyInput interface {
	pulumi.Input

	ToContentKeyPolicyOutput() ContentKeyPolicyOutput
	ToContentKeyPolicyOutputWithContext(ctx context.Context) ContentKeyPolicyOutput
}

func (*ContentKeyPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**ContentKeyPolicy)(nil)).Elem()
}

func (i *ContentKeyPolicy) ToContentKeyPolicyOutput() ContentKeyPolicyOutput {
	return i.ToContentKeyPolicyOutputWithContext(context.Background())
}

func (i *ContentKeyPolicy) ToContentKeyPolicyOutputWithContext(ctx context.Context) ContentKeyPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentKeyPolicyOutput)
}

type ContentKeyPolicyOutput struct{ *pulumi.OutputState }

func (ContentKeyPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContentKeyPolicy)(nil)).Elem()
}

func (o ContentKeyPolicyOutput) ToContentKeyPolicyOutput() ContentKeyPolicyOutput {
	return o
}

func (o ContentKeyPolicyOutput) ToContentKeyPolicyOutputWithContext(ctx context.Context) ContentKeyPolicyOutput {
	return o
}

// The creation date of the Policy
func (o ContentKeyPolicyOutput) Created() pulumi.StringOutput {
	return o.ApplyT(func(v *ContentKeyPolicy) pulumi.StringOutput { return v.Created }).(pulumi.StringOutput)
}

// A description for the Policy.
func (o ContentKeyPolicyOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContentKeyPolicy) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The last modified date of the Policy
func (o ContentKeyPolicyOutput) LastModified() pulumi.StringOutput {
	return o.ApplyT(func(v *ContentKeyPolicy) pulumi.StringOutput { return v.LastModified }).(pulumi.StringOutput)
}

// The name of the resource
func (o ContentKeyPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ContentKeyPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The Key Policy options.
func (o ContentKeyPolicyOutput) Options() ContentKeyPolicyOptionResponseArrayOutput {
	return o.ApplyT(func(v *ContentKeyPolicy) ContentKeyPolicyOptionResponseArrayOutput { return v.Options }).(ContentKeyPolicyOptionResponseArrayOutput)
}

// The legacy Policy ID.
func (o ContentKeyPolicyOutput) PolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v *ContentKeyPolicy) pulumi.StringOutput { return v.PolicyId }).(pulumi.StringOutput)
}

// The system metadata relating to this resource.
func (o ContentKeyPolicyOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ContentKeyPolicy) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ContentKeyPolicyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ContentKeyPolicy) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ContentKeyPolicyOutput{})
}
