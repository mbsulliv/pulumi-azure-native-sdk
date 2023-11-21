// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20240101

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Single item in a List or Get AuthorizationRule operation
type EventHubAuthorizationRule struct {
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

// NewEventHubAuthorizationRule registers a new resource with the given unique name, arguments, and options.
func NewEventHubAuthorizationRule(ctx *pulumi.Context,
	name string, args *EventHubAuthorizationRuleArgs, opts ...pulumi.ResourceOption) (*EventHubAuthorizationRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EventHubName == nil {
		return nil, errors.New("invalid value for required argument 'EventHubName'")
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
			Type: pulumi.String("azure-native:eventhub:EventHubAuthorizationRule"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20140901:EventHubAuthorizationRule"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20150801:EventHubAuthorizationRule"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20170401:EventHubAuthorizationRule"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20180101preview:EventHubAuthorizationRule"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20210101preview:EventHubAuthorizationRule"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20210601preview:EventHubAuthorizationRule"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20211101:EventHubAuthorizationRule"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20220101preview:EventHubAuthorizationRule"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20221001preview:EventHubAuthorizationRule"),
		},
		{
			Type: pulumi.String("azure-native:eventhub/v20230101preview:EventHubAuthorizationRule"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource EventHubAuthorizationRule
	err := ctx.RegisterResource("azure-native:eventhub/v20240101:EventHubAuthorizationRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEventHubAuthorizationRule gets an existing EventHubAuthorizationRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEventHubAuthorizationRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EventHubAuthorizationRuleState, opts ...pulumi.ResourceOption) (*EventHubAuthorizationRule, error) {
	var resource EventHubAuthorizationRule
	err := ctx.ReadResource("azure-native:eventhub/v20240101:EventHubAuthorizationRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EventHubAuthorizationRule resources.
type eventHubAuthorizationRuleState struct {
}

type EventHubAuthorizationRuleState struct {
}

func (EventHubAuthorizationRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*eventHubAuthorizationRuleState)(nil)).Elem()
}

type eventHubAuthorizationRuleArgs struct {
	// The authorization rule name.
	AuthorizationRuleName *string `pulumi:"authorizationRuleName"`
	// The Event Hub name
	EventHubName string `pulumi:"eventHubName"`
	// The Namespace name
	NamespaceName string `pulumi:"namespaceName"`
	// Name of the resource group within the azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The rights associated with the rule.
	Rights []string `pulumi:"rights"`
}

// The set of arguments for constructing a EventHubAuthorizationRule resource.
type EventHubAuthorizationRuleArgs struct {
	// The authorization rule name.
	AuthorizationRuleName pulumi.StringPtrInput
	// The Event Hub name
	EventHubName pulumi.StringInput
	// The Namespace name
	NamespaceName pulumi.StringInput
	// Name of the resource group within the azure subscription.
	ResourceGroupName pulumi.StringInput
	// The rights associated with the rule.
	Rights pulumi.StringArrayInput
}

func (EventHubAuthorizationRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*eventHubAuthorizationRuleArgs)(nil)).Elem()
}

type EventHubAuthorizationRuleInput interface {
	pulumi.Input

	ToEventHubAuthorizationRuleOutput() EventHubAuthorizationRuleOutput
	ToEventHubAuthorizationRuleOutputWithContext(ctx context.Context) EventHubAuthorizationRuleOutput
}

func (*EventHubAuthorizationRule) ElementType() reflect.Type {
	return reflect.TypeOf((**EventHubAuthorizationRule)(nil)).Elem()
}

func (i *EventHubAuthorizationRule) ToEventHubAuthorizationRuleOutput() EventHubAuthorizationRuleOutput {
	return i.ToEventHubAuthorizationRuleOutputWithContext(context.Background())
}

func (i *EventHubAuthorizationRule) ToEventHubAuthorizationRuleOutputWithContext(ctx context.Context) EventHubAuthorizationRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventHubAuthorizationRuleOutput)
}

type EventHubAuthorizationRuleOutput struct{ *pulumi.OutputState }

func (EventHubAuthorizationRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EventHubAuthorizationRule)(nil)).Elem()
}

func (o EventHubAuthorizationRuleOutput) ToEventHubAuthorizationRuleOutput() EventHubAuthorizationRuleOutput {
	return o
}

func (o EventHubAuthorizationRuleOutput) ToEventHubAuthorizationRuleOutputWithContext(ctx context.Context) EventHubAuthorizationRuleOutput {
	return o
}

// The geo-location where the resource lives
func (o EventHubAuthorizationRuleOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *EventHubAuthorizationRule) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o EventHubAuthorizationRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *EventHubAuthorizationRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The rights associated with the rule.
func (o EventHubAuthorizationRuleOutput) Rights() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *EventHubAuthorizationRule) pulumi.StringArrayOutput { return v.Rights }).(pulumi.StringArrayOutput)
}

// The system meta data relating to this resource.
func (o EventHubAuthorizationRuleOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *EventHubAuthorizationRule) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.EventHub/Namespaces" or "Microsoft.EventHub/Namespaces/EventHubs"
func (o EventHubAuthorizationRuleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *EventHubAuthorizationRule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(EventHubAuthorizationRuleOutput{})
}
