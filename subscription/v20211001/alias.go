// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20211001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Subscription Information with the alias.
type Alias struct {
	pulumi.CustomResourceState

	// Alias ID.
	Name pulumi.StringOutput `pulumi:"name"`
	// Subscription Alias response properties.
	Properties SubscriptionAliasResponsePropertiesResponseOutput `pulumi:"properties"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource type, Microsoft.Subscription/aliases.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewAlias registers a new resource with the given unique name, arguments, and options.
func NewAlias(ctx *pulumi.Context,
	name string, args *AliasArgs, opts ...pulumi.ResourceOption) (*Alias, error) {
	if args == nil {
		args = &AliasArgs{}
	}

	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:subscription:Alias"),
		},
		{
			Type: pulumi.String("azure-native:subscription/v20191001preview:Alias"),
		},
		{
			Type: pulumi.String("azure-native:subscription/v20200901:Alias"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Alias
	err := ctx.RegisterResource("azure-native:subscription/v20211001:Alias", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAlias gets an existing Alias resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAlias(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AliasState, opts ...pulumi.ResourceOption) (*Alias, error) {
	var resource Alias
	err := ctx.ReadResource("azure-native:subscription/v20211001:Alias", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Alias resources.
type aliasState struct {
}

type AliasState struct {
}

func (AliasState) ElementType() reflect.Type {
	return reflect.TypeOf((*aliasState)(nil)).Elem()
}

type aliasArgs struct {
	// AliasName is the name for the subscription creation request. Note that this is not the same as subscription name and this doesn’t have any other lifecycle need beyond the request for subscription creation.
	AliasName *string `pulumi:"aliasName"`
	// Put alias request properties.
	Properties *PutAliasRequestProperties `pulumi:"properties"`
}

// The set of arguments for constructing a Alias resource.
type AliasArgs struct {
	// AliasName is the name for the subscription creation request. Note that this is not the same as subscription name and this doesn’t have any other lifecycle need beyond the request for subscription creation.
	AliasName pulumi.StringPtrInput
	// Put alias request properties.
	Properties PutAliasRequestPropertiesPtrInput
}

func (AliasArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*aliasArgs)(nil)).Elem()
}

type AliasInput interface {
	pulumi.Input

	ToAliasOutput() AliasOutput
	ToAliasOutputWithContext(ctx context.Context) AliasOutput
}

func (*Alias) ElementType() reflect.Type {
	return reflect.TypeOf((**Alias)(nil)).Elem()
}

func (i *Alias) ToAliasOutput() AliasOutput {
	return i.ToAliasOutputWithContext(context.Background())
}

func (i *Alias) ToAliasOutputWithContext(ctx context.Context) AliasOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AliasOutput)
}

func (i *Alias) ToOutput(ctx context.Context) pulumix.Output[*Alias] {
	return pulumix.Output[*Alias]{
		OutputState: i.ToAliasOutputWithContext(ctx).OutputState,
	}
}

type AliasOutput struct{ *pulumi.OutputState }

func (AliasOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Alias)(nil)).Elem()
}

func (o AliasOutput) ToAliasOutput() AliasOutput {
	return o
}

func (o AliasOutput) ToAliasOutputWithContext(ctx context.Context) AliasOutput {
	return o
}

func (o AliasOutput) ToOutput(ctx context.Context) pulumix.Output[*Alias] {
	return pulumix.Output[*Alias]{
		OutputState: o.OutputState,
	}
}

// Alias ID.
func (o AliasOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Alias) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Subscription Alias response properties.
func (o AliasOutput) Properties() SubscriptionAliasResponsePropertiesResponseOutput {
	return o.ApplyT(func(v *Alias) SubscriptionAliasResponsePropertiesResponseOutput { return v.Properties }).(SubscriptionAliasResponsePropertiesResponseOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o AliasOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Alias) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource type, Microsoft.Subscription/aliases.
func (o AliasOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Alias) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AliasOutput{})
}
