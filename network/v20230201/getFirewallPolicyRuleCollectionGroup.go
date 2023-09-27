// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets the specified FirewallPolicyRuleCollectionGroup.
func LookupFirewallPolicyRuleCollectionGroup(ctx *pulumi.Context, args *LookupFirewallPolicyRuleCollectionGroupArgs, opts ...pulumi.InvokeOption) (*LookupFirewallPolicyRuleCollectionGroupResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupFirewallPolicyRuleCollectionGroupResult
	err := ctx.Invoke("azure-native:network/v20230201:getFirewallPolicyRuleCollectionGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFirewallPolicyRuleCollectionGroupArgs struct {
	// The name of the Firewall Policy.
	FirewallPolicyName string `pulumi:"firewallPolicyName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the FirewallPolicyRuleCollectionGroup.
	RuleCollectionGroupName string `pulumi:"ruleCollectionGroupName"`
}

// Rule Collection Group resource.
type LookupFirewallPolicyRuleCollectionGroupResult struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `pulumi:"name"`
	// Priority of the Firewall Policy Rule Collection Group resource.
	Priority *int `pulumi:"priority"`
	// The provisioning state of the firewall policy rule collection group resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Group of Firewall Policy rule collections.
	RuleCollections []interface{} `pulumi:"ruleCollections"`
	// Rule Group type.
	Type string `pulumi:"type"`
}

func LookupFirewallPolicyRuleCollectionGroupOutput(ctx *pulumi.Context, args LookupFirewallPolicyRuleCollectionGroupOutputArgs, opts ...pulumi.InvokeOption) LookupFirewallPolicyRuleCollectionGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFirewallPolicyRuleCollectionGroupResult, error) {
			args := v.(LookupFirewallPolicyRuleCollectionGroupArgs)
			r, err := LookupFirewallPolicyRuleCollectionGroup(ctx, &args, opts...)
			var s LookupFirewallPolicyRuleCollectionGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFirewallPolicyRuleCollectionGroupResultOutput)
}

type LookupFirewallPolicyRuleCollectionGroupOutputArgs struct {
	// The name of the Firewall Policy.
	FirewallPolicyName pulumi.StringInput `pulumi:"firewallPolicyName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the FirewallPolicyRuleCollectionGroup.
	RuleCollectionGroupName pulumi.StringInput `pulumi:"ruleCollectionGroupName"`
}

func (LookupFirewallPolicyRuleCollectionGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFirewallPolicyRuleCollectionGroupArgs)(nil)).Elem()
}

// Rule Collection Group resource.
type LookupFirewallPolicyRuleCollectionGroupResultOutput struct{ *pulumi.OutputState }

func (LookupFirewallPolicyRuleCollectionGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFirewallPolicyRuleCollectionGroupResult)(nil)).Elem()
}

func (o LookupFirewallPolicyRuleCollectionGroupResultOutput) ToLookupFirewallPolicyRuleCollectionGroupResultOutput() LookupFirewallPolicyRuleCollectionGroupResultOutput {
	return o
}

func (o LookupFirewallPolicyRuleCollectionGroupResultOutput) ToLookupFirewallPolicyRuleCollectionGroupResultOutputWithContext(ctx context.Context) LookupFirewallPolicyRuleCollectionGroupResultOutput {
	return o
}

func (o LookupFirewallPolicyRuleCollectionGroupResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupFirewallPolicyRuleCollectionGroupResult] {
	return pulumix.Output[LookupFirewallPolicyRuleCollectionGroupResult]{
		OutputState: o.OutputState,
	}
}

// A unique read-only string that changes whenever the resource is updated.
func (o LookupFirewallPolicyRuleCollectionGroupResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicyRuleCollectionGroupResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Resource ID.
func (o LookupFirewallPolicyRuleCollectionGroupResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFirewallPolicyRuleCollectionGroupResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// The name of the resource that is unique within a resource group. This name can be used to access the resource.
func (o LookupFirewallPolicyRuleCollectionGroupResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFirewallPolicyRuleCollectionGroupResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// Priority of the Firewall Policy Rule Collection Group resource.
func (o LookupFirewallPolicyRuleCollectionGroupResultOutput) Priority() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupFirewallPolicyRuleCollectionGroupResult) *int { return v.Priority }).(pulumi.IntPtrOutput)
}

// The provisioning state of the firewall policy rule collection group resource.
func (o LookupFirewallPolicyRuleCollectionGroupResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicyRuleCollectionGroupResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Group of Firewall Policy rule collections.
func (o LookupFirewallPolicyRuleCollectionGroupResultOutput) RuleCollections() pulumi.ArrayOutput {
	return o.ApplyT(func(v LookupFirewallPolicyRuleCollectionGroupResult) []interface{} { return v.RuleCollections }).(pulumi.ArrayOutput)
}

// Rule Group type.
func (o LookupFirewallPolicyRuleCollectionGroupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicyRuleCollectionGroupResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFirewallPolicyRuleCollectionGroupResultOutput{})
}
