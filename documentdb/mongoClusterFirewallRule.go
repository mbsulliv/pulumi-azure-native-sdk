// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package documentdb

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents a mongo cluster firewall rule.
// Azure REST API version: 2023-03-15-preview.
//
// Other available API versions: 2023-09-15-preview.
type MongoClusterFirewallRule struct {
	pulumi.CustomResourceState

	// The end IP address of the mongo cluster firewall rule. Must be IPv4 format.
	EndIpAddress pulumi.StringOutput `pulumi:"endIpAddress"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioning state of the firewall rule.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The start IP address of the mongo cluster firewall rule. Must be IPv4 format.
	StartIpAddress pulumi.StringOutput `pulumi:"startIpAddress"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewMongoClusterFirewallRule registers a new resource with the given unique name, arguments, and options.
func NewMongoClusterFirewallRule(ctx *pulumi.Context,
	name string, args *MongoClusterFirewallRuleArgs, opts ...pulumi.ResourceOption) (*MongoClusterFirewallRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EndIpAddress == nil {
		return nil, errors.New("invalid value for required argument 'EndIpAddress'")
	}
	if args.MongoClusterName == nil {
		return nil, errors.New("invalid value for required argument 'MongoClusterName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.StartIpAddress == nil {
		return nil, errors.New("invalid value for required argument 'StartIpAddress'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:documentdb/v20230301preview:MongoClusterFirewallRule"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20230315preview:MongoClusterFirewallRule"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20230915preview:MongoClusterFirewallRule"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource MongoClusterFirewallRule
	err := ctx.RegisterResource("azure-native:documentdb:MongoClusterFirewallRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMongoClusterFirewallRule gets an existing MongoClusterFirewallRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMongoClusterFirewallRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MongoClusterFirewallRuleState, opts ...pulumi.ResourceOption) (*MongoClusterFirewallRule, error) {
	var resource MongoClusterFirewallRule
	err := ctx.ReadResource("azure-native:documentdb:MongoClusterFirewallRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MongoClusterFirewallRule resources.
type mongoClusterFirewallRuleState struct {
}

type MongoClusterFirewallRuleState struct {
}

func (MongoClusterFirewallRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*mongoClusterFirewallRuleState)(nil)).Elem()
}

type mongoClusterFirewallRuleArgs struct {
	// The end IP address of the mongo cluster firewall rule. Must be IPv4 format.
	EndIpAddress string `pulumi:"endIpAddress"`
	// The name of the mongo cluster firewall rule.
	FirewallRuleName *string `pulumi:"firewallRuleName"`
	// The name of the mongo cluster.
	MongoClusterName string `pulumi:"mongoClusterName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The start IP address of the mongo cluster firewall rule. Must be IPv4 format.
	StartIpAddress string `pulumi:"startIpAddress"`
}

// The set of arguments for constructing a MongoClusterFirewallRule resource.
type MongoClusterFirewallRuleArgs struct {
	// The end IP address of the mongo cluster firewall rule. Must be IPv4 format.
	EndIpAddress pulumi.StringInput
	// The name of the mongo cluster firewall rule.
	FirewallRuleName pulumi.StringPtrInput
	// The name of the mongo cluster.
	MongoClusterName pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The start IP address of the mongo cluster firewall rule. Must be IPv4 format.
	StartIpAddress pulumi.StringInput
}

func (MongoClusterFirewallRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mongoClusterFirewallRuleArgs)(nil)).Elem()
}

type MongoClusterFirewallRuleInput interface {
	pulumi.Input

	ToMongoClusterFirewallRuleOutput() MongoClusterFirewallRuleOutput
	ToMongoClusterFirewallRuleOutputWithContext(ctx context.Context) MongoClusterFirewallRuleOutput
}

func (*MongoClusterFirewallRule) ElementType() reflect.Type {
	return reflect.TypeOf((**MongoClusterFirewallRule)(nil)).Elem()
}

func (i *MongoClusterFirewallRule) ToMongoClusterFirewallRuleOutput() MongoClusterFirewallRuleOutput {
	return i.ToMongoClusterFirewallRuleOutputWithContext(context.Background())
}

func (i *MongoClusterFirewallRule) ToMongoClusterFirewallRuleOutputWithContext(ctx context.Context) MongoClusterFirewallRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MongoClusterFirewallRuleOutput)
}

type MongoClusterFirewallRuleOutput struct{ *pulumi.OutputState }

func (MongoClusterFirewallRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MongoClusterFirewallRule)(nil)).Elem()
}

func (o MongoClusterFirewallRuleOutput) ToMongoClusterFirewallRuleOutput() MongoClusterFirewallRuleOutput {
	return o
}

func (o MongoClusterFirewallRuleOutput) ToMongoClusterFirewallRuleOutputWithContext(ctx context.Context) MongoClusterFirewallRuleOutput {
	return o
}

// The end IP address of the mongo cluster firewall rule. Must be IPv4 format.
func (o MongoClusterFirewallRuleOutput) EndIpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *MongoClusterFirewallRule) pulumi.StringOutput { return v.EndIpAddress }).(pulumi.StringOutput)
}

// The name of the resource
func (o MongoClusterFirewallRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MongoClusterFirewallRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the firewall rule.
func (o MongoClusterFirewallRuleOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *MongoClusterFirewallRule) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The start IP address of the mongo cluster firewall rule. Must be IPv4 format.
func (o MongoClusterFirewallRuleOutput) StartIpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *MongoClusterFirewallRule) pulumi.StringOutput { return v.StartIpAddress }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o MongoClusterFirewallRuleOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *MongoClusterFirewallRule) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o MongoClusterFirewallRuleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *MongoClusterFirewallRule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(MongoClusterFirewallRuleOutput{})
}
