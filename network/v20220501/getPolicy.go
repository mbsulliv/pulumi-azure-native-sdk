// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieve protection policy with specified name within a resource group.
func LookupPolicy(ctx *pulumi.Context, args *LookupPolicyArgs, opts ...pulumi.InvokeOption) (*LookupPolicyResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupPolicyResult
	err := ctx.Invoke("azure-native:network/v20220501:getPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPolicyArgs struct {
	// The name of the Web Application Firewall Policy.
	PolicyName string `pulumi:"policyName"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Defines web application firewall policy.
type LookupPolicyResult struct {
	// Describes custom rules inside the policy.
	CustomRules *CustomRuleListResponse `pulumi:"customRules"`
	// Gets a unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// Describes Frontend Endpoints associated with this Web Application Firewall policy.
	FrontendEndpointLinks []FrontendEndpointLinkResponse `pulumi:"frontendEndpointLinks"`
	// Resource ID.
	Id string `pulumi:"id"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Describes managed rules inside the policy.
	ManagedRules *ManagedRuleSetListResponse `pulumi:"managedRules"`
	// Resource name.
	Name string `pulumi:"name"`
	// Describes settings for the policy.
	PolicySettings *FrontDoorPolicySettingsResponse `pulumi:"policySettings"`
	// Provisioning state of the policy.
	ProvisioningState string `pulumi:"provisioningState"`
	ResourceState     string `pulumi:"resourceState"`
	// Describes Routing Rules associated with this Web Application Firewall policy.
	RoutingRuleLinks []RoutingRuleLinkResponse `pulumi:"routingRuleLinks"`
	// Describes Security Policy associated with this Web Application Firewall policy.
	SecurityPolicyLinks []SecurityPolicyLinkResponse `pulumi:"securityPolicyLinks"`
	// The pricing tier of web application firewall policy. Defaults to Classic_AzureFrontDoor if not specified.
	Sku *SkuResponse `pulumi:"sku"`
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
	// The name of the Web Application Firewall Policy.
	PolicyName pulumi.StringInput `pulumi:"policyName"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPolicyArgs)(nil)).Elem()
}

// Defines web application firewall policy.
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

// Describes custom rules inside the policy.
func (o LookupPolicyResultOutput) CustomRules() CustomRuleListResponsePtrOutput {
	return o.ApplyT(func(v LookupPolicyResult) *CustomRuleListResponse { return v.CustomRules }).(CustomRuleListResponsePtrOutput)
}

// Gets a unique read-only string that changes whenever the resource is updated.
func (o LookupPolicyResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicyResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// Describes Frontend Endpoints associated with this Web Application Firewall policy.
func (o LookupPolicyResultOutput) FrontendEndpointLinks() FrontendEndpointLinkResponseArrayOutput {
	return o.ApplyT(func(v LookupPolicyResult) []FrontendEndpointLinkResponse { return v.FrontendEndpointLinks }).(FrontendEndpointLinkResponseArrayOutput)
}

// Resource ID.
func (o LookupPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource location.
func (o LookupPolicyResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicyResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// Describes managed rules inside the policy.
func (o LookupPolicyResultOutput) ManagedRules() ManagedRuleSetListResponsePtrOutput {
	return o.ApplyT(func(v LookupPolicyResult) *ManagedRuleSetListResponse { return v.ManagedRules }).(ManagedRuleSetListResponsePtrOutput)
}

// Resource name.
func (o LookupPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

// Describes settings for the policy.
func (o LookupPolicyResultOutput) PolicySettings() FrontDoorPolicySettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupPolicyResult) *FrontDoorPolicySettingsResponse { return v.PolicySettings }).(FrontDoorPolicySettingsResponsePtrOutput)
}

// Provisioning state of the policy.
func (o LookupPolicyResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupPolicyResultOutput) ResourceState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyResult) string { return v.ResourceState }).(pulumi.StringOutput)
}

// Describes Routing Rules associated with this Web Application Firewall policy.
func (o LookupPolicyResultOutput) RoutingRuleLinks() RoutingRuleLinkResponseArrayOutput {
	return o.ApplyT(func(v LookupPolicyResult) []RoutingRuleLinkResponse { return v.RoutingRuleLinks }).(RoutingRuleLinkResponseArrayOutput)
}

// Describes Security Policy associated with this Web Application Firewall policy.
func (o LookupPolicyResultOutput) SecurityPolicyLinks() SecurityPolicyLinkResponseArrayOutput {
	return o.ApplyT(func(v LookupPolicyResult) []SecurityPolicyLinkResponse { return v.SecurityPolicyLinks }).(SecurityPolicyLinkResponseArrayOutput)
}

// The pricing tier of web application firewall policy. Defaults to Classic_AzureFrontDoor if not specified.
func (o LookupPolicyResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupPolicyResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
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
