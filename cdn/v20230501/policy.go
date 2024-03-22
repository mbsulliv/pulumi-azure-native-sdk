// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230501

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Defines web application firewall policy for Azure CDN.
type Policy struct {
	pulumi.CustomResourceState

	// Describes custom rules inside the policy.
	CustomRules CustomRuleListResponsePtrOutput `pulumi:"customRules"`
	// Describes Azure CDN endpoints associated with this Web Application Firewall policy.
	EndpointLinks CdnEndpointResponseArrayOutput `pulumi:"endpointLinks"`
	// Gets a unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// Key-Value pair representing additional properties for Web Application Firewall policy.
	ExtendedProperties pulumi.StringMapOutput `pulumi:"extendedProperties"`
	// Resource location.
	Location pulumi.StringOutput `pulumi:"location"`
	// Describes managed rules inside the policy.
	ManagedRules ManagedRuleSetListResponsePtrOutput `pulumi:"managedRules"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Describes  policySettings for policy
	PolicySettings PolicySettingsResponsePtrOutput `pulumi:"policySettings"`
	// Provisioning state of the WebApplicationFirewallPolicy.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Describes rate limit rules inside the policy.
	RateLimitRules RateLimitRuleListResponsePtrOutput `pulumi:"rateLimitRules"`
	ResourceState  pulumi.StringOutput                `pulumi:"resourceState"`
	// The pricing tier (defines a CDN provider, feature list and rate) of the CdnWebApplicationFirewallPolicy.
	Sku SkuResponseOutput `pulumi:"sku"`
	// Read only system data
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewPolicy registers a new resource with the given unique name, arguments, and options.
func NewPolicy(ctx *pulumi.Context,
	name string, args *PolicyArgs, opts ...pulumi.ResourceOption) (*Policy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Sku == nil {
		return nil, errors.New("invalid value for required argument 'Sku'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:cdn:Policy"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20190615:Policy"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20190615preview:Policy"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20200331:Policy"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20200415:Policy"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20200901:Policy"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20210601:Policy"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20220501preview:Policy"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20221101preview:Policy"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20230701preview:Policy"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20240201:Policy"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Policy
	err := ctx.RegisterResource("azure-native:cdn/v20230501:Policy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPolicy gets an existing Policy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PolicyState, opts ...pulumi.ResourceOption) (*Policy, error) {
	var resource Policy
	err := ctx.ReadResource("azure-native:cdn/v20230501:Policy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Policy resources.
type policyState struct {
}

type PolicyState struct {
}

func (PolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*policyState)(nil)).Elem()
}

type policyArgs struct {
	// Describes custom rules inside the policy.
	CustomRules *CustomRuleList `pulumi:"customRules"`
	// Key-Value pair representing additional properties for Web Application Firewall policy.
	ExtendedProperties map[string]string `pulumi:"extendedProperties"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Describes managed rules inside the policy.
	ManagedRules *ManagedRuleSetList `pulumi:"managedRules"`
	// The name of the CdnWebApplicationFirewallPolicy.
	PolicyName *string `pulumi:"policyName"`
	// Describes  policySettings for policy
	PolicySettings *PolicySettings `pulumi:"policySettings"`
	// Describes rate limit rules inside the policy.
	RateLimitRules *RateLimitRuleList `pulumi:"rateLimitRules"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The pricing tier (defines a CDN provider, feature list and rate) of the CdnWebApplicationFirewallPolicy.
	Sku Sku `pulumi:"sku"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Policy resource.
type PolicyArgs struct {
	// Describes custom rules inside the policy.
	CustomRules CustomRuleListPtrInput
	// Key-Value pair representing additional properties for Web Application Firewall policy.
	ExtendedProperties pulumi.StringMapInput
	// Resource location.
	Location pulumi.StringPtrInput
	// Describes managed rules inside the policy.
	ManagedRules ManagedRuleSetListPtrInput
	// The name of the CdnWebApplicationFirewallPolicy.
	PolicyName pulumi.StringPtrInput
	// Describes  policySettings for policy
	PolicySettings PolicySettingsPtrInput
	// Describes rate limit rules inside the policy.
	RateLimitRules RateLimitRuleListPtrInput
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput
	// The pricing tier (defines a CDN provider, feature list and rate) of the CdnWebApplicationFirewallPolicy.
	Sku SkuInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (PolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*policyArgs)(nil)).Elem()
}

type PolicyInput interface {
	pulumi.Input

	ToPolicyOutput() PolicyOutput
	ToPolicyOutputWithContext(ctx context.Context) PolicyOutput
}

func (*Policy) ElementType() reflect.Type {
	return reflect.TypeOf((**Policy)(nil)).Elem()
}

func (i *Policy) ToPolicyOutput() PolicyOutput {
	return i.ToPolicyOutputWithContext(context.Background())
}

func (i *Policy) ToPolicyOutputWithContext(ctx context.Context) PolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyOutput)
}

type PolicyOutput struct{ *pulumi.OutputState }

func (PolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Policy)(nil)).Elem()
}

func (o PolicyOutput) ToPolicyOutput() PolicyOutput {
	return o
}

func (o PolicyOutput) ToPolicyOutputWithContext(ctx context.Context) PolicyOutput {
	return o
}

// Describes custom rules inside the policy.
func (o PolicyOutput) CustomRules() CustomRuleListResponsePtrOutput {
	return o.ApplyT(func(v *Policy) CustomRuleListResponsePtrOutput { return v.CustomRules }).(CustomRuleListResponsePtrOutput)
}

// Describes Azure CDN endpoints associated with this Web Application Firewall policy.
func (o PolicyOutput) EndpointLinks() CdnEndpointResponseArrayOutput {
	return o.ApplyT(func(v *Policy) CdnEndpointResponseArrayOutput { return v.EndpointLinks }).(CdnEndpointResponseArrayOutput)
}

// Gets a unique read-only string that changes whenever the resource is updated.
func (o PolicyOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Policy) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// Key-Value pair representing additional properties for Web Application Firewall policy.
func (o PolicyOutput) ExtendedProperties() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Policy) pulumi.StringMapOutput { return v.ExtendedProperties }).(pulumi.StringMapOutput)
}

// Resource location.
func (o PolicyOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Policy) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Describes managed rules inside the policy.
func (o PolicyOutput) ManagedRules() ManagedRuleSetListResponsePtrOutput {
	return o.ApplyT(func(v *Policy) ManagedRuleSetListResponsePtrOutput { return v.ManagedRules }).(ManagedRuleSetListResponsePtrOutput)
}

// Resource name.
func (o PolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Policy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Describes  policySettings for policy
func (o PolicyOutput) PolicySettings() PolicySettingsResponsePtrOutput {
	return o.ApplyT(func(v *Policy) PolicySettingsResponsePtrOutput { return v.PolicySettings }).(PolicySettingsResponsePtrOutput)
}

// Provisioning state of the WebApplicationFirewallPolicy.
func (o PolicyOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Policy) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Describes rate limit rules inside the policy.
func (o PolicyOutput) RateLimitRules() RateLimitRuleListResponsePtrOutput {
	return o.ApplyT(func(v *Policy) RateLimitRuleListResponsePtrOutput { return v.RateLimitRules }).(RateLimitRuleListResponsePtrOutput)
}

func (o PolicyOutput) ResourceState() pulumi.StringOutput {
	return o.ApplyT(func(v *Policy) pulumi.StringOutput { return v.ResourceState }).(pulumi.StringOutput)
}

// The pricing tier (defines a CDN provider, feature list and rate) of the CdnWebApplicationFirewallPolicy.
func (o PolicyOutput) Sku() SkuResponseOutput {
	return o.ApplyT(func(v *Policy) SkuResponseOutput { return v.Sku }).(SkuResponseOutput)
}

// Read only system data
func (o PolicyOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Policy) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o PolicyOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Policy) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o PolicyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Policy) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PolicyOutput{})
}
