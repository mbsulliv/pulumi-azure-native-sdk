// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package relay

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Single item in a List or Get AuthorizationRule operation
// Azure REST API version: 2021-11-01. Prior API version in Azure Native 1.x: 2017-04-01.
//
// Other available API versions: 2017-04-01.
type WCFRelayAuthorizationRule struct {
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

// NewWCFRelayAuthorizationRule registers a new resource with the given unique name, arguments, and options.
func NewWCFRelayAuthorizationRule(ctx *pulumi.Context,
	name string, args *WCFRelayAuthorizationRuleArgs, opts ...pulumi.ResourceOption) (*WCFRelayAuthorizationRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NamespaceName == nil {
		return nil, errors.New("invalid value for required argument 'NamespaceName'")
	}
	if args.RelayName == nil {
		return nil, errors.New("invalid value for required argument 'RelayName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Rights == nil {
		return nil, errors.New("invalid value for required argument 'Rights'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:relay/v20160701:WCFRelayAuthorizationRule"),
		},
		{
			Type: pulumi.String("azure-native:relay/v20170401:WCFRelayAuthorizationRule"),
		},
		{
			Type: pulumi.String("azure-native:relay/v20211101:WCFRelayAuthorizationRule"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource WCFRelayAuthorizationRule
	err := ctx.RegisterResource("azure-native:relay:WCFRelayAuthorizationRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWCFRelayAuthorizationRule gets an existing WCFRelayAuthorizationRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWCFRelayAuthorizationRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WCFRelayAuthorizationRuleState, opts ...pulumi.ResourceOption) (*WCFRelayAuthorizationRule, error) {
	var resource WCFRelayAuthorizationRule
	err := ctx.ReadResource("azure-native:relay:WCFRelayAuthorizationRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WCFRelayAuthorizationRule resources.
type wcfrelayAuthorizationRuleState struct {
}

type WCFRelayAuthorizationRuleState struct {
}

func (WCFRelayAuthorizationRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*wcfrelayAuthorizationRuleState)(nil)).Elem()
}

type wcfrelayAuthorizationRuleArgs struct {
	// The authorization rule name.
	AuthorizationRuleName *string `pulumi:"authorizationRuleName"`
	// The namespace name
	NamespaceName string `pulumi:"namespaceName"`
	// The relay name.
	RelayName string `pulumi:"relayName"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The rights associated with the rule.
	Rights []string `pulumi:"rights"`
}

// The set of arguments for constructing a WCFRelayAuthorizationRule resource.
type WCFRelayAuthorizationRuleArgs struct {
	// The authorization rule name.
	AuthorizationRuleName pulumi.StringPtrInput
	// The namespace name
	NamespaceName pulumi.StringInput
	// The relay name.
	RelayName pulumi.StringInput
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput
	// The rights associated with the rule.
	Rights pulumi.StringArrayInput
}

func (WCFRelayAuthorizationRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*wcfrelayAuthorizationRuleArgs)(nil)).Elem()
}

type WCFRelayAuthorizationRuleInput interface {
	pulumi.Input

	ToWCFRelayAuthorizationRuleOutput() WCFRelayAuthorizationRuleOutput
	ToWCFRelayAuthorizationRuleOutputWithContext(ctx context.Context) WCFRelayAuthorizationRuleOutput
}

func (*WCFRelayAuthorizationRule) ElementType() reflect.Type {
	return reflect.TypeOf((**WCFRelayAuthorizationRule)(nil)).Elem()
}

func (i *WCFRelayAuthorizationRule) ToWCFRelayAuthorizationRuleOutput() WCFRelayAuthorizationRuleOutput {
	return i.ToWCFRelayAuthorizationRuleOutputWithContext(context.Background())
}

func (i *WCFRelayAuthorizationRule) ToWCFRelayAuthorizationRuleOutputWithContext(ctx context.Context) WCFRelayAuthorizationRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WCFRelayAuthorizationRuleOutput)
}

func (i *WCFRelayAuthorizationRule) ToOutput(ctx context.Context) pulumix.Output[*WCFRelayAuthorizationRule] {
	return pulumix.Output[*WCFRelayAuthorizationRule]{
		OutputState: i.ToWCFRelayAuthorizationRuleOutputWithContext(ctx).OutputState,
	}
}

type WCFRelayAuthorizationRuleOutput struct{ *pulumi.OutputState }

func (WCFRelayAuthorizationRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WCFRelayAuthorizationRule)(nil)).Elem()
}

func (o WCFRelayAuthorizationRuleOutput) ToWCFRelayAuthorizationRuleOutput() WCFRelayAuthorizationRuleOutput {
	return o
}

func (o WCFRelayAuthorizationRuleOutput) ToWCFRelayAuthorizationRuleOutputWithContext(ctx context.Context) WCFRelayAuthorizationRuleOutput {
	return o
}

func (o WCFRelayAuthorizationRuleOutput) ToOutput(ctx context.Context) pulumix.Output[*WCFRelayAuthorizationRule] {
	return pulumix.Output[*WCFRelayAuthorizationRule]{
		OutputState: o.OutputState,
	}
}

// The geo-location where the resource lives
func (o WCFRelayAuthorizationRuleOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *WCFRelayAuthorizationRule) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o WCFRelayAuthorizationRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WCFRelayAuthorizationRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The rights associated with the rule.
func (o WCFRelayAuthorizationRuleOutput) Rights() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *WCFRelayAuthorizationRule) pulumi.StringArrayOutput { return v.Rights }).(pulumi.StringArrayOutput)
}

// The system meta data relating to this resource.
func (o WCFRelayAuthorizationRuleOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *WCFRelayAuthorizationRule) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.EventHub/Namespaces" or "Microsoft.EventHub/Namespaces/EventHubs"
func (o WCFRelayAuthorizationRuleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WCFRelayAuthorizationRule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WCFRelayAuthorizationRuleOutput{})
}
