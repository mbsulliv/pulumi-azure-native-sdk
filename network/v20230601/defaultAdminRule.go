// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230601

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Network default admin rule.
type DefaultAdminRule struct {
	pulumi.CustomResourceState

	// Indicates the access allowed for this particular rule
	Access pulumi.StringOutput `pulumi:"access"`
	// A description for this rule. Restricted to 140 chars.
	Description pulumi.StringOutput `pulumi:"description"`
	// The destination port ranges.
	DestinationPortRanges pulumi.StringArrayOutput `pulumi:"destinationPortRanges"`
	// The destination address prefixes. CIDR or destination IP ranges.
	Destinations AddressPrefixItemResponseArrayOutput `pulumi:"destinations"`
	// Indicates if the traffic matched against the rule in inbound or outbound.
	Direction pulumi.StringOutput `pulumi:"direction"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Default rule flag.
	Flag pulumi.StringPtrOutput `pulumi:"flag"`
	// Whether the rule is custom or default.
	// Expected value is 'Default'.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The priority of the rule. The value can be between 1 and 4096. The priority number must be unique for each rule in the collection. The lower the priority number, the higher the priority of the rule.
	Priority pulumi.IntOutput `pulumi:"priority"`
	// Network protocol this rule applies to.
	Protocol pulumi.StringOutput `pulumi:"protocol"`
	// The provisioning state of the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Unique identifier for this resource.
	ResourceGuid pulumi.StringOutput `pulumi:"resourceGuid"`
	// The source port ranges.
	SourcePortRanges pulumi.StringArrayOutput `pulumi:"sourcePortRanges"`
	// The CIDR or source IP ranges.
	Sources AddressPrefixItemResponseArrayOutput `pulumi:"sources"`
	// The system metadata related to this resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDefaultAdminRule registers a new resource with the given unique name, arguments, and options.
func NewDefaultAdminRule(ctx *pulumi.Context,
	name string, args *DefaultAdminRuleArgs, opts ...pulumi.ResourceOption) (*DefaultAdminRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConfigurationName == nil {
		return nil, errors.New("invalid value for required argument 'ConfigurationName'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.NetworkManagerName == nil {
		return nil, errors.New("invalid value for required argument 'NetworkManagerName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.RuleCollectionName == nil {
		return nil, errors.New("invalid value for required argument 'RuleCollectionName'")
	}
	args.Kind = pulumi.String("Default")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:DefaultAdminRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201preview:DefaultAdminRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501preview:DefaultAdminRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:DefaultAdminRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220201preview:DefaultAdminRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220401preview:DefaultAdminRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:DefaultAdminRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220701:DefaultAdminRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220901:DefaultAdminRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20221101:DefaultAdminRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230201:DefaultAdminRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230401:DefaultAdminRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230501:DefaultAdminRule"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource DefaultAdminRule
	err := ctx.RegisterResource("azure-native:network/v20230601:DefaultAdminRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDefaultAdminRule gets an existing DefaultAdminRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDefaultAdminRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DefaultAdminRuleState, opts ...pulumi.ResourceOption) (*DefaultAdminRule, error) {
	var resource DefaultAdminRule
	err := ctx.ReadResource("azure-native:network/v20230601:DefaultAdminRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DefaultAdminRule resources.
type defaultAdminRuleState struct {
}

type DefaultAdminRuleState struct {
}

func (DefaultAdminRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*defaultAdminRuleState)(nil)).Elem()
}

type defaultAdminRuleArgs struct {
	// The name of the network manager Security Configuration.
	ConfigurationName string `pulumi:"configurationName"`
	// Default rule flag.
	Flag *string `pulumi:"flag"`
	// Whether the rule is custom or default.
	// Expected value is 'Default'.
	Kind string `pulumi:"kind"`
	// The name of the network manager.
	NetworkManagerName string `pulumi:"networkManagerName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the network manager security Configuration rule collection.
	RuleCollectionName string `pulumi:"ruleCollectionName"`
	// The name of the rule.
	RuleName *string `pulumi:"ruleName"`
}

// The set of arguments for constructing a DefaultAdminRule resource.
type DefaultAdminRuleArgs struct {
	// The name of the network manager Security Configuration.
	ConfigurationName pulumi.StringInput
	// Default rule flag.
	Flag pulumi.StringPtrInput
	// Whether the rule is custom or default.
	// Expected value is 'Default'.
	Kind pulumi.StringInput
	// The name of the network manager.
	NetworkManagerName pulumi.StringInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The name of the network manager security Configuration rule collection.
	RuleCollectionName pulumi.StringInput
	// The name of the rule.
	RuleName pulumi.StringPtrInput
}

func (DefaultAdminRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*defaultAdminRuleArgs)(nil)).Elem()
}

type DefaultAdminRuleInput interface {
	pulumi.Input

	ToDefaultAdminRuleOutput() DefaultAdminRuleOutput
	ToDefaultAdminRuleOutputWithContext(ctx context.Context) DefaultAdminRuleOutput
}

func (*DefaultAdminRule) ElementType() reflect.Type {
	return reflect.TypeOf((**DefaultAdminRule)(nil)).Elem()
}

func (i *DefaultAdminRule) ToDefaultAdminRuleOutput() DefaultAdminRuleOutput {
	return i.ToDefaultAdminRuleOutputWithContext(context.Background())
}

func (i *DefaultAdminRule) ToDefaultAdminRuleOutputWithContext(ctx context.Context) DefaultAdminRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefaultAdminRuleOutput)
}

type DefaultAdminRuleOutput struct{ *pulumi.OutputState }

func (DefaultAdminRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DefaultAdminRule)(nil)).Elem()
}

func (o DefaultAdminRuleOutput) ToDefaultAdminRuleOutput() DefaultAdminRuleOutput {
	return o
}

func (o DefaultAdminRuleOutput) ToDefaultAdminRuleOutputWithContext(ctx context.Context) DefaultAdminRuleOutput {
	return o
}

// Indicates the access allowed for this particular rule
func (o DefaultAdminRuleOutput) Access() pulumi.StringOutput {
	return o.ApplyT(func(v *DefaultAdminRule) pulumi.StringOutput { return v.Access }).(pulumi.StringOutput)
}

// A description for this rule. Restricted to 140 chars.
func (o DefaultAdminRuleOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *DefaultAdminRule) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// The destination port ranges.
func (o DefaultAdminRuleOutput) DestinationPortRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DefaultAdminRule) pulumi.StringArrayOutput { return v.DestinationPortRanges }).(pulumi.StringArrayOutput)
}

// The destination address prefixes. CIDR or destination IP ranges.
func (o DefaultAdminRuleOutput) Destinations() AddressPrefixItemResponseArrayOutput {
	return o.ApplyT(func(v *DefaultAdminRule) AddressPrefixItemResponseArrayOutput { return v.Destinations }).(AddressPrefixItemResponseArrayOutput)
}

// Indicates if the traffic matched against the rule in inbound or outbound.
func (o DefaultAdminRuleOutput) Direction() pulumi.StringOutput {
	return o.ApplyT(func(v *DefaultAdminRule) pulumi.StringOutput { return v.Direction }).(pulumi.StringOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o DefaultAdminRuleOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *DefaultAdminRule) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Default rule flag.
func (o DefaultAdminRuleOutput) Flag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DefaultAdminRule) pulumi.StringPtrOutput { return v.Flag }).(pulumi.StringPtrOutput)
}

// Whether the rule is custom or default.
// Expected value is 'Default'.
func (o DefaultAdminRuleOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *DefaultAdminRule) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Resource name.
func (o DefaultAdminRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DefaultAdminRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The priority of the rule. The value can be between 1 and 4096. The priority number must be unique for each rule in the collection. The lower the priority number, the higher the priority of the rule.
func (o DefaultAdminRuleOutput) Priority() pulumi.IntOutput {
	return o.ApplyT(func(v *DefaultAdminRule) pulumi.IntOutput { return v.Priority }).(pulumi.IntOutput)
}

// Network protocol this rule applies to.
func (o DefaultAdminRuleOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v *DefaultAdminRule) pulumi.StringOutput { return v.Protocol }).(pulumi.StringOutput)
}

// The provisioning state of the resource.
func (o DefaultAdminRuleOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *DefaultAdminRule) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Unique identifier for this resource.
func (o DefaultAdminRuleOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v *DefaultAdminRule) pulumi.StringOutput { return v.ResourceGuid }).(pulumi.StringOutput)
}

// The source port ranges.
func (o DefaultAdminRuleOutput) SourcePortRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DefaultAdminRule) pulumi.StringArrayOutput { return v.SourcePortRanges }).(pulumi.StringArrayOutput)
}

// The CIDR or source IP ranges.
func (o DefaultAdminRuleOutput) Sources() AddressPrefixItemResponseArrayOutput {
	return o.ApplyT(func(v *DefaultAdminRule) AddressPrefixItemResponseArrayOutput { return v.Sources }).(AddressPrefixItemResponseArrayOutput)
}

// The system metadata related to this resource.
func (o DefaultAdminRuleOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *DefaultAdminRule) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource type.
func (o DefaultAdminRuleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DefaultAdminRule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DefaultAdminRuleOutput{})
}
