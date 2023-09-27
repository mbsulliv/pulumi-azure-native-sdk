// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230801

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Tag Inheritance Setting definition.
type TagInheritanceSetting struct {
	pulumi.CustomResourceState

	// Specifies the kind of settings.
	// Expected value is 'taginheritance'.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The properties of the tag inheritance setting.
	Properties TagInheritancePropertiesResponseOutput `pulumi:"properties"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewTagInheritanceSetting registers a new resource with the given unique name, arguments, and options.
func NewTagInheritanceSetting(ctx *pulumi.Context,
	name string, args *TagInheritanceSettingArgs, opts ...pulumi.ResourceOption) (*TagInheritanceSetting, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.Scope == nil {
		return nil, errors.New("invalid value for required argument 'Scope'")
	}
	args.Kind = pulumi.String("taginheritance")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:costmanagement:TagInheritanceSetting"),
		},
		{
			Type: pulumi.String("azure-native:costmanagement/v20221001preview:TagInheritanceSetting"),
		},
		{
			Type: pulumi.String("azure-native:costmanagement/v20221005preview:TagInheritanceSetting"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource TagInheritanceSetting
	err := ctx.RegisterResource("azure-native:costmanagement/v20230801:TagInheritanceSetting", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTagInheritanceSetting gets an existing TagInheritanceSetting resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTagInheritanceSetting(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TagInheritanceSettingState, opts ...pulumi.ResourceOption) (*TagInheritanceSetting, error) {
	var resource TagInheritanceSetting
	err := ctx.ReadResource("azure-native:costmanagement/v20230801:TagInheritanceSetting", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TagInheritanceSetting resources.
type tagInheritanceSettingState struct {
}

type TagInheritanceSettingState struct {
}

func (TagInheritanceSettingState) ElementType() reflect.Type {
	return reflect.TypeOf((*tagInheritanceSettingState)(nil)).Elem()
}

type tagInheritanceSettingArgs struct {
	// Specifies the kind of settings.
	// Expected value is 'taginheritance'.
	Kind string `pulumi:"kind"`
	// The properties of the tag inheritance setting.
	Properties *TagInheritanceProperties `pulumi:"properties"`
	// The scope associated with this setting. This includes 'subscriptions/{subscriptionId}' for subscription scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}' for billing profile scope.
	Scope string `pulumi:"scope"`
	// Setting type.
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a TagInheritanceSetting resource.
type TagInheritanceSettingArgs struct {
	// Specifies the kind of settings.
	// Expected value is 'taginheritance'.
	Kind pulumi.StringInput
	// The properties of the tag inheritance setting.
	Properties TagInheritancePropertiesPtrInput
	// The scope associated with this setting. This includes 'subscriptions/{subscriptionId}' for subscription scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}' for billing profile scope.
	Scope pulumi.StringInput
	// Setting type.
	Type pulumi.StringPtrInput
}

func (TagInheritanceSettingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tagInheritanceSettingArgs)(nil)).Elem()
}

type TagInheritanceSettingInput interface {
	pulumi.Input

	ToTagInheritanceSettingOutput() TagInheritanceSettingOutput
	ToTagInheritanceSettingOutputWithContext(ctx context.Context) TagInheritanceSettingOutput
}

func (*TagInheritanceSetting) ElementType() reflect.Type {
	return reflect.TypeOf((**TagInheritanceSetting)(nil)).Elem()
}

func (i *TagInheritanceSetting) ToTagInheritanceSettingOutput() TagInheritanceSettingOutput {
	return i.ToTagInheritanceSettingOutputWithContext(context.Background())
}

func (i *TagInheritanceSetting) ToTagInheritanceSettingOutputWithContext(ctx context.Context) TagInheritanceSettingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagInheritanceSettingOutput)
}

func (i *TagInheritanceSetting) ToOutput(ctx context.Context) pulumix.Output[*TagInheritanceSetting] {
	return pulumix.Output[*TagInheritanceSetting]{
		OutputState: i.ToTagInheritanceSettingOutputWithContext(ctx).OutputState,
	}
}

type TagInheritanceSettingOutput struct{ *pulumi.OutputState }

func (TagInheritanceSettingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TagInheritanceSetting)(nil)).Elem()
}

func (o TagInheritanceSettingOutput) ToTagInheritanceSettingOutput() TagInheritanceSettingOutput {
	return o
}

func (o TagInheritanceSettingOutput) ToTagInheritanceSettingOutputWithContext(ctx context.Context) TagInheritanceSettingOutput {
	return o
}

func (o TagInheritanceSettingOutput) ToOutput(ctx context.Context) pulumix.Output[*TagInheritanceSetting] {
	return pulumix.Output[*TagInheritanceSetting]{
		OutputState: o.OutputState,
	}
}

// Specifies the kind of settings.
// Expected value is 'taginheritance'.
func (o TagInheritanceSettingOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *TagInheritanceSetting) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// The name of the resource
func (o TagInheritanceSettingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *TagInheritanceSetting) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The properties of the tag inheritance setting.
func (o TagInheritanceSettingOutput) Properties() TagInheritancePropertiesResponseOutput {
	return o.ApplyT(func(v *TagInheritanceSetting) TagInheritancePropertiesResponseOutput { return v.Properties }).(TagInheritancePropertiesResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o TagInheritanceSettingOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *TagInheritanceSetting) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(TagInheritanceSettingOutput{})
}
