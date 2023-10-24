// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Defines web application firewall policy.
// Azure REST API version: 2023-02-01. Prior API version in Azure Native 1.x: 2020-11-01.
//
// Other available API versions: 2019-07-01, 2023-04-01, 2023-05-01.
type WebApplicationFirewallPolicy struct {
	pulumi.CustomResourceState

	// A collection of references to application gateways.
	ApplicationGateways ApplicationGatewayResponseArrayOutput `pulumi:"applicationGateways"`
	// The custom rules inside the policy.
	CustomRules WebApplicationFirewallCustomRuleResponseArrayOutput `pulumi:"customRules"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// A collection of references to application gateway http listeners.
	HttpListeners SubResourceResponseArrayOutput `pulumi:"httpListeners"`
	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Describes the managedRules structure.
	ManagedRules ManagedRulesDefinitionResponseOutput `pulumi:"managedRules"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// A collection of references to application gateway path rules.
	PathBasedRules SubResourceResponseArrayOutput `pulumi:"pathBasedRules"`
	// The PolicySettings for policy.
	PolicySettings PolicySettingsResponsePtrOutput `pulumi:"policySettings"`
	// The provisioning state of the web application firewall policy resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Resource status of the policy.
	ResourceState pulumi.StringOutput `pulumi:"resourceState"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewWebApplicationFirewallPolicy registers a new resource with the given unique name, arguments, and options.
func NewWebApplicationFirewallPolicy(ctx *pulumi.Context,
	name string, args *WebApplicationFirewallPolicyArgs, opts ...pulumi.ResourceOption) (*WebApplicationFirewallPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ManagedRules == nil {
		return nil, errors.New("invalid value for required argument 'ManagedRules'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.PolicySettings != nil {
		args.PolicySettings = args.PolicySettings.ToPolicySettingsPtrOutput().ApplyT(func(v *PolicySettings) *PolicySettings { return v.Defaults() }).(PolicySettingsPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network/v20181201:WebApplicationFirewallPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190201:WebApplicationFirewallPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190401:WebApplicationFirewallPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:WebApplicationFirewallPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:WebApplicationFirewallPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:WebApplicationFirewallPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:WebApplicationFirewallPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:WebApplicationFirewallPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:WebApplicationFirewallPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:WebApplicationFirewallPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:WebApplicationFirewallPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:WebApplicationFirewallPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:WebApplicationFirewallPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:WebApplicationFirewallPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:WebApplicationFirewallPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:WebApplicationFirewallPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:WebApplicationFirewallPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:WebApplicationFirewallPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:WebApplicationFirewallPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:WebApplicationFirewallPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:WebApplicationFirewallPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:WebApplicationFirewallPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220701:WebApplicationFirewallPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220901:WebApplicationFirewallPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20221101:WebApplicationFirewallPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230201:WebApplicationFirewallPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230401:WebApplicationFirewallPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230501:WebApplicationFirewallPolicy"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource WebApplicationFirewallPolicy
	err := ctx.RegisterResource("azure-native:network:WebApplicationFirewallPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebApplicationFirewallPolicy gets an existing WebApplicationFirewallPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebApplicationFirewallPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebApplicationFirewallPolicyState, opts ...pulumi.ResourceOption) (*WebApplicationFirewallPolicy, error) {
	var resource WebApplicationFirewallPolicy
	err := ctx.ReadResource("azure-native:network:WebApplicationFirewallPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebApplicationFirewallPolicy resources.
type webApplicationFirewallPolicyState struct {
}

type WebApplicationFirewallPolicyState struct {
}

func (WebApplicationFirewallPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*webApplicationFirewallPolicyState)(nil)).Elem()
}

type webApplicationFirewallPolicyArgs struct {
	// The custom rules inside the policy.
	CustomRules []WebApplicationFirewallCustomRule `pulumi:"customRules"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Describes the managedRules structure.
	ManagedRules ManagedRulesDefinition `pulumi:"managedRules"`
	// The name of the policy.
	PolicyName *string `pulumi:"policyName"`
	// The PolicySettings for policy.
	PolicySettings *PolicySettings `pulumi:"policySettings"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a WebApplicationFirewallPolicy resource.
type WebApplicationFirewallPolicyArgs struct {
	// The custom rules inside the policy.
	CustomRules WebApplicationFirewallCustomRuleArrayInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// Describes the managedRules structure.
	ManagedRules ManagedRulesDefinitionInput
	// The name of the policy.
	PolicyName pulumi.StringPtrInput
	// The PolicySettings for policy.
	PolicySettings PolicySettingsPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (WebApplicationFirewallPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webApplicationFirewallPolicyArgs)(nil)).Elem()
}

type WebApplicationFirewallPolicyInput interface {
	pulumi.Input

	ToWebApplicationFirewallPolicyOutput() WebApplicationFirewallPolicyOutput
	ToWebApplicationFirewallPolicyOutputWithContext(ctx context.Context) WebApplicationFirewallPolicyOutput
}

func (*WebApplicationFirewallPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**WebApplicationFirewallPolicy)(nil)).Elem()
}

func (i *WebApplicationFirewallPolicy) ToWebApplicationFirewallPolicyOutput() WebApplicationFirewallPolicyOutput {
	return i.ToWebApplicationFirewallPolicyOutputWithContext(context.Background())
}

func (i *WebApplicationFirewallPolicy) ToWebApplicationFirewallPolicyOutputWithContext(ctx context.Context) WebApplicationFirewallPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebApplicationFirewallPolicyOutput)
}

func (i *WebApplicationFirewallPolicy) ToOutput(ctx context.Context) pulumix.Output[*WebApplicationFirewallPolicy] {
	return pulumix.Output[*WebApplicationFirewallPolicy]{
		OutputState: i.ToWebApplicationFirewallPolicyOutputWithContext(ctx).OutputState,
	}
}

type WebApplicationFirewallPolicyOutput struct{ *pulumi.OutputState }

func (WebApplicationFirewallPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebApplicationFirewallPolicy)(nil)).Elem()
}

func (o WebApplicationFirewallPolicyOutput) ToWebApplicationFirewallPolicyOutput() WebApplicationFirewallPolicyOutput {
	return o
}

func (o WebApplicationFirewallPolicyOutput) ToWebApplicationFirewallPolicyOutputWithContext(ctx context.Context) WebApplicationFirewallPolicyOutput {
	return o
}

func (o WebApplicationFirewallPolicyOutput) ToOutput(ctx context.Context) pulumix.Output[*WebApplicationFirewallPolicy] {
	return pulumix.Output[*WebApplicationFirewallPolicy]{
		OutputState: o.OutputState,
	}
}

// A collection of references to application gateways.
func (o WebApplicationFirewallPolicyOutput) ApplicationGateways() ApplicationGatewayResponseArrayOutput {
	return o.ApplyT(func(v *WebApplicationFirewallPolicy) ApplicationGatewayResponseArrayOutput {
		return v.ApplicationGateways
	}).(ApplicationGatewayResponseArrayOutput)
}

// The custom rules inside the policy.
func (o WebApplicationFirewallPolicyOutput) CustomRules() WebApplicationFirewallCustomRuleResponseArrayOutput {
	return o.ApplyT(func(v *WebApplicationFirewallPolicy) WebApplicationFirewallCustomRuleResponseArrayOutput {
		return v.CustomRules
	}).(WebApplicationFirewallCustomRuleResponseArrayOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o WebApplicationFirewallPolicyOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *WebApplicationFirewallPolicy) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// A collection of references to application gateway http listeners.
func (o WebApplicationFirewallPolicyOutput) HttpListeners() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v *WebApplicationFirewallPolicy) SubResourceResponseArrayOutput { return v.HttpListeners }).(SubResourceResponseArrayOutput)
}

// Resource location.
func (o WebApplicationFirewallPolicyOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebApplicationFirewallPolicy) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Describes the managedRules structure.
func (o WebApplicationFirewallPolicyOutput) ManagedRules() ManagedRulesDefinitionResponseOutput {
	return o.ApplyT(func(v *WebApplicationFirewallPolicy) ManagedRulesDefinitionResponseOutput { return v.ManagedRules }).(ManagedRulesDefinitionResponseOutput)
}

// Resource name.
func (o WebApplicationFirewallPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WebApplicationFirewallPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// A collection of references to application gateway path rules.
func (o WebApplicationFirewallPolicyOutput) PathBasedRules() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v *WebApplicationFirewallPolicy) SubResourceResponseArrayOutput { return v.PathBasedRules }).(SubResourceResponseArrayOutput)
}

// The PolicySettings for policy.
func (o WebApplicationFirewallPolicyOutput) PolicySettings() PolicySettingsResponsePtrOutput {
	return o.ApplyT(func(v *WebApplicationFirewallPolicy) PolicySettingsResponsePtrOutput { return v.PolicySettings }).(PolicySettingsResponsePtrOutput)
}

// The provisioning state of the web application firewall policy resource.
func (o WebApplicationFirewallPolicyOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *WebApplicationFirewallPolicy) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Resource status of the policy.
func (o WebApplicationFirewallPolicyOutput) ResourceState() pulumi.StringOutput {
	return o.ApplyT(func(v *WebApplicationFirewallPolicy) pulumi.StringOutput { return v.ResourceState }).(pulumi.StringOutput)
}

// Resource tags.
func (o WebApplicationFirewallPolicyOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *WebApplicationFirewallPolicy) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o WebApplicationFirewallPolicyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WebApplicationFirewallPolicy) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WebApplicationFirewallPolicyOutput{})
}
