// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230101preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Single item in a List or Get AuthorizationRule operation
type NamespaceAuthorizationRule struct {
	pulumi.CustomResourceState

	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The rights associated with the rule.
	Rights pulumi.StringArrayOutput `pulumi:"rights"`
	// The system meta data relating to this resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.EventHub/Namespaces" or "Microsoft.EventHub/Namespaces/EventHubs"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewNamespaceAuthorizationRule registers a new resource with the given unique name, arguments, and options.
func NewNamespaceAuthorizationRule(ctx *pulumi.Context,
	name string, args *NamespaceAuthorizationRuleArgs, opts ...pulumi.ResourceOption) (*NamespaceAuthorizationRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NamespaceName == nil {
		return nil, errors.New("invalid value for required argument 'NamespaceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Rights == nil {
		return nil, errors.New("invalid value for required argument 'Rights'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:eventhub:NamespaceAuthorizationRule"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20140901:NamespaceAuthorizationRule"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20150801:NamespaceAuthorizationRule"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20170401:NamespaceAuthorizationRule"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20180101preview:NamespaceAuthorizationRule"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20210101preview:NamespaceAuthorizationRule"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20210601preview:NamespaceAuthorizationRule"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20211101:NamespaceAuthorizationRule"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20220101preview:NamespaceAuthorizationRule"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20221001preview:NamespaceAuthorizationRule"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource NamespaceAuthorizationRule
	err := ctx.RegisterResource("azure-native:eventhub/v20230101preview:NamespaceAuthorizationRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNamespaceAuthorizationRule gets an existing NamespaceAuthorizationRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNamespaceAuthorizationRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NamespaceAuthorizationRuleState, opts ...pulumi.ResourceOption) (*NamespaceAuthorizationRule, error) {
	var resource NamespaceAuthorizationRule
	err := ctx.ReadResource("azure-native:eventhub/v20230101preview:NamespaceAuthorizationRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NamespaceAuthorizationRule resources.
type namespaceAuthorizationRuleState struct {
}

type NamespaceAuthorizationRuleState struct {
}

func (NamespaceAuthorizationRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*namespaceAuthorizationRuleState)(nil)).Elem()
}

type namespaceAuthorizationRuleArgs struct {
	// The authorization rule name.
	AuthorizationRuleName *string `pulumi:"authorizationRuleName"`
	// The Namespace name
	NamespaceName string `pulumi:"namespaceName"`
	// Name of the resource group within the azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The rights associated with the rule.
	Rights []string `pulumi:"rights"`
}

// The set of arguments for constructing a NamespaceAuthorizationRule resource.
type NamespaceAuthorizationRuleArgs struct {
	// The authorization rule name.
	AuthorizationRuleName pulumi.StringPtrInput
	// The Namespace name
	NamespaceName pulumi.StringInput
	// Name of the resource group within the azure subscription.
	ResourceGroupName pulumi.StringInput
	// The rights associated with the rule.
	Rights pulumi.StringArrayInput
}

func (NamespaceAuthorizationRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*namespaceAuthorizationRuleArgs)(nil)).Elem()
}

type NamespaceAuthorizationRuleInput interface {
	pulumi.Input

	ToNamespaceAuthorizationRuleOutput() NamespaceAuthorizationRuleOutput
	ToNamespaceAuthorizationRuleOutputWithContext(ctx context.Context) NamespaceAuthorizationRuleOutput
}

func (*NamespaceAuthorizationRule) ElementType() reflect.Type {
	return reflect.TypeOf((**NamespaceAuthorizationRule)(nil)).Elem()
}

func (i *NamespaceAuthorizationRule) ToNamespaceAuthorizationRuleOutput() NamespaceAuthorizationRuleOutput {
	return i.ToNamespaceAuthorizationRuleOutputWithContext(context.Background())
}

func (i *NamespaceAuthorizationRule) ToNamespaceAuthorizationRuleOutputWithContext(ctx context.Context) NamespaceAuthorizationRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamespaceAuthorizationRuleOutput)
}

type NamespaceAuthorizationRuleOutput struct{ *pulumi.OutputState }

func (NamespaceAuthorizationRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NamespaceAuthorizationRule)(nil)).Elem()
}

func (o NamespaceAuthorizationRuleOutput) ToNamespaceAuthorizationRuleOutput() NamespaceAuthorizationRuleOutput {
	return o
}

func (o NamespaceAuthorizationRuleOutput) ToNamespaceAuthorizationRuleOutputWithContext(ctx context.Context) NamespaceAuthorizationRuleOutput {
	return o
}

// The geo-location where the resource lives
func (o NamespaceAuthorizationRuleOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *NamespaceAuthorizationRule) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o NamespaceAuthorizationRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NamespaceAuthorizationRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The rights associated with the rule.
func (o NamespaceAuthorizationRuleOutput) Rights() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NamespaceAuthorizationRule) pulumi.StringArrayOutput { return v.Rights }).(pulumi.StringArrayOutput)
}

// The system meta data relating to this resource.
func (o NamespaceAuthorizationRuleOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *NamespaceAuthorizationRule) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.EventHub/Namespaces" or "Microsoft.EventHub/Namespaces/EventHubs"
func (o NamespaceAuthorizationRuleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *NamespaceAuthorizationRule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(NamespaceAuthorizationRuleOutput{})
}
