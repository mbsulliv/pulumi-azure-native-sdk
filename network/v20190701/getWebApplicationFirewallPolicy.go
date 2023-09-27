// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20190701

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Retrieve protection policy with specified name within a resource group.
func LookupWebApplicationFirewallPolicy(ctx *pulumi.Context, args *LookupWebApplicationFirewallPolicyArgs, opts ...pulumi.InvokeOption) (*LookupWebApplicationFirewallPolicyResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupWebApplicationFirewallPolicyResult
	err := ctx.Invoke("azure-native:network/v20190701:getWebApplicationFirewallPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWebApplicationFirewallPolicyArgs struct {
	// The name of the policy.
	PolicyName string `pulumi:"policyName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Defines web application firewall policy.
type LookupWebApplicationFirewallPolicyResult struct {
	// A collection of references to application gateways.
	ApplicationGateways []ApplicationGatewayResponse `pulumi:"applicationGateways"`
	// Describes custom rules inside the policy.
	CustomRules []WebApplicationFirewallCustomRuleResponse `pulumi:"customRules"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Resource name.
	Name string `pulumi:"name"`
	// Describes policySettings for policy.
	PolicySettings *PolicySettingsResponse `pulumi:"policySettings"`
	// The provisioning state of the web application firewall policy resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Resource status of the policy.
	ResourceState string `pulumi:"resourceState"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupWebApplicationFirewallPolicyOutput(ctx *pulumi.Context, args LookupWebApplicationFirewallPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupWebApplicationFirewallPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWebApplicationFirewallPolicyResult, error) {
			args := v.(LookupWebApplicationFirewallPolicyArgs)
			r, err := LookupWebApplicationFirewallPolicy(ctx, &args, opts...)
			var s LookupWebApplicationFirewallPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWebApplicationFirewallPolicyResultOutput)
}

type LookupWebApplicationFirewallPolicyOutputArgs struct {
	// The name of the policy.
	PolicyName pulumi.StringInput `pulumi:"policyName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupWebApplicationFirewallPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebApplicationFirewallPolicyArgs)(nil)).Elem()
}

// Defines web application firewall policy.
type LookupWebApplicationFirewallPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupWebApplicationFirewallPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebApplicationFirewallPolicyResult)(nil)).Elem()
}

func (o LookupWebApplicationFirewallPolicyResultOutput) ToLookupWebApplicationFirewallPolicyResultOutput() LookupWebApplicationFirewallPolicyResultOutput {
	return o
}

func (o LookupWebApplicationFirewallPolicyResultOutput) ToLookupWebApplicationFirewallPolicyResultOutputWithContext(ctx context.Context) LookupWebApplicationFirewallPolicyResultOutput {
	return o
}

func (o LookupWebApplicationFirewallPolicyResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupWebApplicationFirewallPolicyResult] {
	return pulumix.Output[LookupWebApplicationFirewallPolicyResult]{
		OutputState: o.OutputState,
	}
}

// A collection of references to application gateways.
func (o LookupWebApplicationFirewallPolicyResultOutput) ApplicationGateways() ApplicationGatewayResponseArrayOutput {
	return o.ApplyT(func(v LookupWebApplicationFirewallPolicyResult) []ApplicationGatewayResponse {
		return v.ApplicationGateways
	}).(ApplicationGatewayResponseArrayOutput)
}

// Describes custom rules inside the policy.
func (o LookupWebApplicationFirewallPolicyResultOutput) CustomRules() WebApplicationFirewallCustomRuleResponseArrayOutput {
	return o.ApplyT(func(v LookupWebApplicationFirewallPolicyResult) []WebApplicationFirewallCustomRuleResponse {
		return v.CustomRules
	}).(WebApplicationFirewallCustomRuleResponseArrayOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o LookupWebApplicationFirewallPolicyResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebApplicationFirewallPolicyResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// Resource ID.
func (o LookupWebApplicationFirewallPolicyResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebApplicationFirewallPolicyResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// Resource location.
func (o LookupWebApplicationFirewallPolicyResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebApplicationFirewallPolicyResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o LookupWebApplicationFirewallPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebApplicationFirewallPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

// Describes policySettings for policy.
func (o LookupWebApplicationFirewallPolicyResultOutput) PolicySettings() PolicySettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupWebApplicationFirewallPolicyResult) *PolicySettingsResponse { return v.PolicySettings }).(PolicySettingsResponsePtrOutput)
}

// The provisioning state of the web application firewall policy resource.
func (o LookupWebApplicationFirewallPolicyResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebApplicationFirewallPolicyResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Resource status of the policy.
func (o LookupWebApplicationFirewallPolicyResultOutput) ResourceState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebApplicationFirewallPolicyResult) string { return v.ResourceState }).(pulumi.StringOutput)
}

// Resource tags.
func (o LookupWebApplicationFirewallPolicyResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupWebApplicationFirewallPolicyResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o LookupWebApplicationFirewallPolicyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebApplicationFirewallPolicyResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWebApplicationFirewallPolicyResultOutput{})
}
