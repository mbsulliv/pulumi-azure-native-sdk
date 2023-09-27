// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230401

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// FirewallPolicy Resource.
type FirewallPolicy struct {
	pulumi.CustomResourceState

	// The parent firewall policy from which rules are inherited.
	BasePolicy SubResourceResponsePtrOutput `pulumi:"basePolicy"`
	// List of references to Child Firewall Policies.
	ChildPolicies SubResourceResponseArrayOutput `pulumi:"childPolicies"`
	// DNS Proxy Settings definition.
	DnsSettings DnsSettingsResponsePtrOutput `pulumi:"dnsSettings"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Explicit Proxy Settings definition.
	ExplicitProxy ExplicitProxyResponsePtrOutput `pulumi:"explicitProxy"`
	// List of references to Azure Firewalls that this Firewall Policy is associated with.
	Firewalls SubResourceResponseArrayOutput `pulumi:"firewalls"`
	// The identity of the firewall policy.
	Identity ManagedServiceIdentityResponsePtrOutput `pulumi:"identity"`
	// Insights on Firewall Policy.
	Insights FirewallPolicyInsightsResponsePtrOutput `pulumi:"insights"`
	// The configuration for Intrusion detection.
	IntrusionDetection FirewallPolicyIntrusionDetectionResponsePtrOutput `pulumi:"intrusionDetection"`
	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioning state of the firewall policy resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// List of references to FirewallPolicyRuleCollectionGroups.
	RuleCollectionGroups SubResourceResponseArrayOutput `pulumi:"ruleCollectionGroups"`
	// The Firewall Policy SKU.
	Sku FirewallPolicySkuResponsePtrOutput `pulumi:"sku"`
	// The private IP addresses/IP ranges to which traffic will not be SNAT.
	Snat FirewallPolicySNATResponsePtrOutput `pulumi:"snat"`
	// SQL Settings definition.
	Sql FirewallPolicySQLResponsePtrOutput `pulumi:"sql"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The operation mode for Threat Intelligence.
	ThreatIntelMode pulumi.StringPtrOutput `pulumi:"threatIntelMode"`
	// ThreatIntel Whitelist for Firewall Policy.
	ThreatIntelWhitelist FirewallPolicyThreatIntelWhitelistResponsePtrOutput `pulumi:"threatIntelWhitelist"`
	// TLS Configuration definition.
	TransportSecurity FirewallPolicyTransportSecurityResponsePtrOutput `pulumi:"transportSecurity"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewFirewallPolicy registers a new resource with the given unique name, arguments, and options.
func NewFirewallPolicy(ctx *pulumi.Context,
	name string, args *FirewallPolicyArgs, opts ...pulumi.ResourceOption) (*FirewallPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:FirewallPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:FirewallPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:FirewallPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:FirewallPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:FirewallPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:FirewallPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:FirewallPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:FirewallPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:FirewallPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:FirewallPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:FirewallPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:FirewallPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:FirewallPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:FirewallPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:FirewallPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:FirewallPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:FirewallPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:FirewallPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:FirewallPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:FirewallPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220701:FirewallPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220901:FirewallPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20221101:FirewallPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230201:FirewallPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230501:FirewallPolicy"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource FirewallPolicy
	err := ctx.RegisterResource("azure-native:network/v20230401:FirewallPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallPolicy gets an existing FirewallPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallPolicyState, opts ...pulumi.ResourceOption) (*FirewallPolicy, error) {
	var resource FirewallPolicy
	err := ctx.ReadResource("azure-native:network/v20230401:FirewallPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallPolicy resources.
type firewallPolicyState struct {
}

type FirewallPolicyState struct {
}

func (FirewallPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallPolicyState)(nil)).Elem()
}

type firewallPolicyArgs struct {
	// The parent firewall policy from which rules are inherited.
	BasePolicy *SubResource `pulumi:"basePolicy"`
	// DNS Proxy Settings definition.
	DnsSettings *DnsSettings `pulumi:"dnsSettings"`
	// Explicit Proxy Settings definition.
	ExplicitProxy *ExplicitProxy `pulumi:"explicitProxy"`
	// The name of the Firewall Policy.
	FirewallPolicyName *string `pulumi:"firewallPolicyName"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// The identity of the firewall policy.
	Identity *ManagedServiceIdentity `pulumi:"identity"`
	// Insights on Firewall Policy.
	Insights *FirewallPolicyInsights `pulumi:"insights"`
	// The configuration for Intrusion detection.
	IntrusionDetection *FirewallPolicyIntrusionDetection `pulumi:"intrusionDetection"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The Firewall Policy SKU.
	Sku *FirewallPolicySku `pulumi:"sku"`
	// The private IP addresses/IP ranges to which traffic will not be SNAT.
	Snat *FirewallPolicySNAT `pulumi:"snat"`
	// SQL Settings definition.
	Sql *FirewallPolicySQL `pulumi:"sql"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The operation mode for Threat Intelligence.
	ThreatIntelMode *string `pulumi:"threatIntelMode"`
	// ThreatIntel Whitelist for Firewall Policy.
	ThreatIntelWhitelist *FirewallPolicyThreatIntelWhitelist `pulumi:"threatIntelWhitelist"`
	// TLS Configuration definition.
	TransportSecurity *FirewallPolicyTransportSecurity `pulumi:"transportSecurity"`
}

// The set of arguments for constructing a FirewallPolicy resource.
type FirewallPolicyArgs struct {
	// The parent firewall policy from which rules are inherited.
	BasePolicy SubResourcePtrInput
	// DNS Proxy Settings definition.
	DnsSettings DnsSettingsPtrInput
	// Explicit Proxy Settings definition.
	ExplicitProxy ExplicitProxyPtrInput
	// The name of the Firewall Policy.
	FirewallPolicyName pulumi.StringPtrInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// The identity of the firewall policy.
	Identity ManagedServiceIdentityPtrInput
	// Insights on Firewall Policy.
	Insights FirewallPolicyInsightsPtrInput
	// The configuration for Intrusion detection.
	IntrusionDetection FirewallPolicyIntrusionDetectionPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The Firewall Policy SKU.
	Sku FirewallPolicySkuPtrInput
	// The private IP addresses/IP ranges to which traffic will not be SNAT.
	Snat FirewallPolicySNATPtrInput
	// SQL Settings definition.
	Sql FirewallPolicySQLPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The operation mode for Threat Intelligence.
	ThreatIntelMode pulumi.StringPtrInput
	// ThreatIntel Whitelist for Firewall Policy.
	ThreatIntelWhitelist FirewallPolicyThreatIntelWhitelistPtrInput
	// TLS Configuration definition.
	TransportSecurity FirewallPolicyTransportSecurityPtrInput
}

func (FirewallPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallPolicyArgs)(nil)).Elem()
}

type FirewallPolicyInput interface {
	pulumi.Input

	ToFirewallPolicyOutput() FirewallPolicyOutput
	ToFirewallPolicyOutputWithContext(ctx context.Context) FirewallPolicyOutput
}

func (*FirewallPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallPolicy)(nil)).Elem()
}

func (i *FirewallPolicy) ToFirewallPolicyOutput() FirewallPolicyOutput {
	return i.ToFirewallPolicyOutputWithContext(context.Background())
}

func (i *FirewallPolicy) ToFirewallPolicyOutputWithContext(ctx context.Context) FirewallPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallPolicyOutput)
}

func (i *FirewallPolicy) ToOutput(ctx context.Context) pulumix.Output[*FirewallPolicy] {
	return pulumix.Output[*FirewallPolicy]{
		OutputState: i.ToFirewallPolicyOutputWithContext(ctx).OutputState,
	}
}

type FirewallPolicyOutput struct{ *pulumi.OutputState }

func (FirewallPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallPolicy)(nil)).Elem()
}

func (o FirewallPolicyOutput) ToFirewallPolicyOutput() FirewallPolicyOutput {
	return o
}

func (o FirewallPolicyOutput) ToFirewallPolicyOutputWithContext(ctx context.Context) FirewallPolicyOutput {
	return o
}

func (o FirewallPolicyOutput) ToOutput(ctx context.Context) pulumix.Output[*FirewallPolicy] {
	return pulumix.Output[*FirewallPolicy]{
		OutputState: o.OutputState,
	}
}

// The parent firewall policy from which rules are inherited.
func (o FirewallPolicyOutput) BasePolicy() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *FirewallPolicy) SubResourceResponsePtrOutput { return v.BasePolicy }).(SubResourceResponsePtrOutput)
}

// List of references to Child Firewall Policies.
func (o FirewallPolicyOutput) ChildPolicies() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v *FirewallPolicy) SubResourceResponseArrayOutput { return v.ChildPolicies }).(SubResourceResponseArrayOutput)
}

// DNS Proxy Settings definition.
func (o FirewallPolicyOutput) DnsSettings() DnsSettingsResponsePtrOutput {
	return o.ApplyT(func(v *FirewallPolicy) DnsSettingsResponsePtrOutput { return v.DnsSettings }).(DnsSettingsResponsePtrOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o FirewallPolicyOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallPolicy) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Explicit Proxy Settings definition.
func (o FirewallPolicyOutput) ExplicitProxy() ExplicitProxyResponsePtrOutput {
	return o.ApplyT(func(v *FirewallPolicy) ExplicitProxyResponsePtrOutput { return v.ExplicitProxy }).(ExplicitProxyResponsePtrOutput)
}

// List of references to Azure Firewalls that this Firewall Policy is associated with.
func (o FirewallPolicyOutput) Firewalls() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v *FirewallPolicy) SubResourceResponseArrayOutput { return v.Firewalls }).(SubResourceResponseArrayOutput)
}

// The identity of the firewall policy.
func (o FirewallPolicyOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *FirewallPolicy) ManagedServiceIdentityResponsePtrOutput { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

// Insights on Firewall Policy.
func (o FirewallPolicyOutput) Insights() FirewallPolicyInsightsResponsePtrOutput {
	return o.ApplyT(func(v *FirewallPolicy) FirewallPolicyInsightsResponsePtrOutput { return v.Insights }).(FirewallPolicyInsightsResponsePtrOutput)
}

// The configuration for Intrusion detection.
func (o FirewallPolicyOutput) IntrusionDetection() FirewallPolicyIntrusionDetectionResponsePtrOutput {
	return o.ApplyT(func(v *FirewallPolicy) FirewallPolicyIntrusionDetectionResponsePtrOutput { return v.IntrusionDetection }).(FirewallPolicyIntrusionDetectionResponsePtrOutput)
}

// Resource location.
func (o FirewallPolicyOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallPolicy) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o FirewallPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the firewall policy resource.
func (o FirewallPolicyOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallPolicy) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// List of references to FirewallPolicyRuleCollectionGroups.
func (o FirewallPolicyOutput) RuleCollectionGroups() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v *FirewallPolicy) SubResourceResponseArrayOutput { return v.RuleCollectionGroups }).(SubResourceResponseArrayOutput)
}

// The Firewall Policy SKU.
func (o FirewallPolicyOutput) Sku() FirewallPolicySkuResponsePtrOutput {
	return o.ApplyT(func(v *FirewallPolicy) FirewallPolicySkuResponsePtrOutput { return v.Sku }).(FirewallPolicySkuResponsePtrOutput)
}

// The private IP addresses/IP ranges to which traffic will not be SNAT.
func (o FirewallPolicyOutput) Snat() FirewallPolicySNATResponsePtrOutput {
	return o.ApplyT(func(v *FirewallPolicy) FirewallPolicySNATResponsePtrOutput { return v.Snat }).(FirewallPolicySNATResponsePtrOutput)
}

// SQL Settings definition.
func (o FirewallPolicyOutput) Sql() FirewallPolicySQLResponsePtrOutput {
	return o.ApplyT(func(v *FirewallPolicy) FirewallPolicySQLResponsePtrOutput { return v.Sql }).(FirewallPolicySQLResponsePtrOutput)
}

// Resource tags.
func (o FirewallPolicyOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *FirewallPolicy) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The operation mode for Threat Intelligence.
func (o FirewallPolicyOutput) ThreatIntelMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallPolicy) pulumi.StringPtrOutput { return v.ThreatIntelMode }).(pulumi.StringPtrOutput)
}

// ThreatIntel Whitelist for Firewall Policy.
func (o FirewallPolicyOutput) ThreatIntelWhitelist() FirewallPolicyThreatIntelWhitelistResponsePtrOutput {
	return o.ApplyT(func(v *FirewallPolicy) FirewallPolicyThreatIntelWhitelistResponsePtrOutput {
		return v.ThreatIntelWhitelist
	}).(FirewallPolicyThreatIntelWhitelistResponsePtrOutput)
}

// TLS Configuration definition.
func (o FirewallPolicyOutput) TransportSecurity() FirewallPolicyTransportSecurityResponsePtrOutput {
	return o.ApplyT(func(v *FirewallPolicy) FirewallPolicyTransportSecurityResponsePtrOutput { return v.TransportSecurity }).(FirewallPolicyTransportSecurityResponsePtrOutput)
}

// Resource type.
func (o FirewallPolicyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallPolicy) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(FirewallPolicyOutput{})
}
