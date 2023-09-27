// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230501preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// A firewall rule on a redis cache has a name, and describes a contiguous range of IP addresses permitted to connect
type FirewallRule struct {
	pulumi.CustomResourceState

	// highest IP address included in the range
	EndIP pulumi.StringOutput `pulumi:"endIP"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// lowest IP address included in the range
	StartIP pulumi.StringOutput `pulumi:"startIP"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewFirewallRule registers a new resource with the given unique name, arguments, and options.
func NewFirewallRule(ctx *pulumi.Context,
	name string, args *FirewallRuleArgs, opts ...pulumi.ResourceOption) (*FirewallRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CacheName == nil {
		return nil, errors.New("invalid value for required argument 'CacheName'")
	}
	if args.EndIP == nil {
		return nil, errors.New("invalid value for required argument 'EndIP'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.StartIP == nil {
		return nil, errors.New("invalid value for required argument 'StartIP'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:cache:FirewallRule"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20160401:FirewallRule"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20170201:FirewallRule"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20171001:FirewallRule"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20180301:FirewallRule"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20190701:FirewallRule"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20200601:FirewallRule"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20201201:FirewallRule"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20210601:FirewallRule"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20220501:FirewallRule"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20220601:FirewallRule"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20230401:FirewallRule"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20230801:FirewallRule"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource FirewallRule
	err := ctx.RegisterResource("azure-native:cache/v20230501preview:FirewallRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallRule gets an existing FirewallRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallRuleState, opts ...pulumi.ResourceOption) (*FirewallRule, error) {
	var resource FirewallRule
	err := ctx.ReadResource("azure-native:cache/v20230501preview:FirewallRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallRule resources.
type firewallRuleState struct {
}

type FirewallRuleState struct {
}

func (FirewallRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallRuleState)(nil)).Elem()
}

type firewallRuleArgs struct {
	// The name of the Redis cache.
	CacheName string `pulumi:"cacheName"`
	// highest IP address included in the range
	EndIP string `pulumi:"endIP"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the firewall rule.
	RuleName *string `pulumi:"ruleName"`
	// lowest IP address included in the range
	StartIP string `pulumi:"startIP"`
}

// The set of arguments for constructing a FirewallRule resource.
type FirewallRuleArgs struct {
	// The name of the Redis cache.
	CacheName pulumi.StringInput
	// highest IP address included in the range
	EndIP pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the firewall rule.
	RuleName pulumi.StringPtrInput
	// lowest IP address included in the range
	StartIP pulumi.StringInput
}

func (FirewallRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallRuleArgs)(nil)).Elem()
}

type FirewallRuleInput interface {
	pulumi.Input

	ToFirewallRuleOutput() FirewallRuleOutput
	ToFirewallRuleOutputWithContext(ctx context.Context) FirewallRuleOutput
}

func (*FirewallRule) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallRule)(nil)).Elem()
}

func (i *FirewallRule) ToFirewallRuleOutput() FirewallRuleOutput {
	return i.ToFirewallRuleOutputWithContext(context.Background())
}

func (i *FirewallRule) ToFirewallRuleOutputWithContext(ctx context.Context) FirewallRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallRuleOutput)
}

func (i *FirewallRule) ToOutput(ctx context.Context) pulumix.Output[*FirewallRule] {
	return pulumix.Output[*FirewallRule]{
		OutputState: i.ToFirewallRuleOutputWithContext(ctx).OutputState,
	}
}

type FirewallRuleOutput struct{ *pulumi.OutputState }

func (FirewallRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallRule)(nil)).Elem()
}

func (o FirewallRuleOutput) ToFirewallRuleOutput() FirewallRuleOutput {
	return o
}

func (o FirewallRuleOutput) ToFirewallRuleOutputWithContext(ctx context.Context) FirewallRuleOutput {
	return o
}

func (o FirewallRuleOutput) ToOutput(ctx context.Context) pulumix.Output[*FirewallRule] {
	return pulumix.Output[*FirewallRule]{
		OutputState: o.OutputState,
	}
}

// highest IP address included in the range
func (o FirewallRuleOutput) EndIP() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallRule) pulumi.StringOutput { return v.EndIP }).(pulumi.StringOutput)
}

// The name of the resource
func (o FirewallRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// lowest IP address included in the range
func (o FirewallRuleOutput) StartIP() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallRule) pulumi.StringOutput { return v.StartIP }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o FirewallRuleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallRule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(FirewallRuleOutput{})
}
