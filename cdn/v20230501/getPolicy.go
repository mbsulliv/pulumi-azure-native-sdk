// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Retrieve protection policy with specified name within a resource group.
func LookupPolicy(ctx *pulumi.Context, args *LookupPolicyArgs, opts ...pulumi.InvokeOption) (*LookupPolicyResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupPolicyResult
	err := ctx.Invoke("azure-native:cdn/v20230501:getPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPolicyArgs struct {
	// The name of the CdnWebApplicationFirewallPolicy.
	PolicyName string `pulumi:"policyName"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Defines web application firewall policy for Azure CDN.
type LookupPolicyResult struct {
	// Describes custom rules inside the policy.
	CustomRules *CustomRuleListResponse `pulumi:"customRules"`
	// Describes Azure CDN endpoints associated with this Web Application Firewall policy.
	EndpointLinks []CdnEndpointResponse `pulumi:"endpointLinks"`
	// Gets a unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// Key-Value pair representing additional properties for Web Application Firewall policy.
	ExtendedProperties map[string]string `pulumi:"extendedProperties"`
	// Resource ID.
	Id string `pulumi:"id"`
	// Resource location.
	Location string `pulumi:"location"`
	// Describes managed rules inside the policy.
	ManagedRules *ManagedRuleSetListResponse `pulumi:"managedRules"`
	// Resource name.
	Name string `pulumi:"name"`
	// Describes  policySettings for policy
	PolicySettings *PolicySettingsResponse `pulumi:"policySettings"`
	// Provisioning state of the WebApplicationFirewallPolicy.
	ProvisioningState string `pulumi:"provisioningState"`
	// Describes rate limit rules inside the policy.
	RateLimitRules *RateLimitRuleListResponse `pulumi:"rateLimitRules"`
	ResourceState  string                     `pulumi:"resourceState"`
	// The pricing tier (defines a CDN provider, feature list and rate) of the CdnWebApplicationFirewallPolicy.
	Sku SkuResponse `pulumi:"sku"`
	// Read only system data
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupPolicyOutput(ctx *pulumi.Context, args LookupPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPolicyResult, error) {
			args := v.(LookupPolicyArgs)
			r, err := LookupPolicy(ctx, &args, opts...)
			var s LookupPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPolicyResultOutput)
}

type LookupPolicyOutputArgs struct {
	// The name of the CdnWebApplicationFirewallPolicy.
	PolicyName pulumi.StringInput `pulumi:"policyName"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPolicyArgs)(nil)).Elem()
}

// Defines web application firewall policy for Azure CDN.
type LookupPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPolicyResult)(nil)).Elem()
}

func (o LookupPolicyResultOutput) ToLookupPolicyResultOutput() LookupPolicyResultOutput {
	return o
}

func (o LookupPolicyResultOutput) ToLookupPolicyResultOutputWithContext(ctx context.Context) LookupPolicyResultOutput {
	return o
}

func (o LookupPolicyResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupPolicyResult] {
	return pulumix.Output[LookupPolicyResult]{
		OutputState: o.OutputState,
	}
}

// Describes custom rules inside the policy.
func (o LookupPolicyResultOutput) CustomRules() CustomRuleListResponsePtrOutput {
	return o.ApplyT(func(v LookupPolicyResult) *CustomRuleListResponse { return v.CustomRules }).(CustomRuleListResponsePtrOutput)
}

// Describes Azure CDN endpoints associated with this Web Application Firewall policy.
func (o LookupPolicyResultOutput) EndpointLinks() CdnEndpointResponseArrayOutput {
	return o.ApplyT(func(v LookupPolicyResult) []CdnEndpointResponse { return v.EndpointLinks }).(CdnEndpointResponseArrayOutput)
}

// Gets a unique read-only string that changes whenever the resource is updated.
func (o LookupPolicyResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicyResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// Key-Value pair representing additional properties for Web Application Firewall policy.
func (o LookupPolicyResultOutput) ExtendedProperties() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupPolicyResult) map[string]string { return v.ExtendedProperties }).(pulumi.StringMapOutput)
}

// Resource ID.
func (o LookupPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource location.
func (o LookupPolicyResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyResult) string { return v.Location }).(pulumi.StringOutput)
}

// Describes managed rules inside the policy.
func (o LookupPolicyResultOutput) ManagedRules() ManagedRuleSetListResponsePtrOutput {
	return o.ApplyT(func(v LookupPolicyResult) *ManagedRuleSetListResponse { return v.ManagedRules }).(ManagedRuleSetListResponsePtrOutput)
}

// Resource name.
func (o LookupPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

// Describes  policySettings for policy
func (o LookupPolicyResultOutput) PolicySettings() PolicySettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupPolicyResult) *PolicySettingsResponse { return v.PolicySettings }).(PolicySettingsResponsePtrOutput)
}

// Provisioning state of the WebApplicationFirewallPolicy.
func (o LookupPolicyResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Describes rate limit rules inside the policy.
func (o LookupPolicyResultOutput) RateLimitRules() RateLimitRuleListResponsePtrOutput {
	return o.ApplyT(func(v LookupPolicyResult) *RateLimitRuleListResponse { return v.RateLimitRules }).(RateLimitRuleListResponsePtrOutput)
}

func (o LookupPolicyResultOutput) ResourceState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyResult) string { return v.ResourceState }).(pulumi.StringOutput)
}

// The pricing tier (defines a CDN provider, feature list and rate) of the CdnWebApplicationFirewallPolicy.
func (o LookupPolicyResultOutput) Sku() SkuResponseOutput {
	return o.ApplyT(func(v LookupPolicyResult) SkuResponse { return v.Sku }).(SkuResponseOutput)
}

// Read only system data
func (o LookupPolicyResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupPolicyResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupPolicyResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupPolicyResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o LookupPolicyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPolicyResultOutput{})
}
