// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220615

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// EventGrid Partner Namespace.
type PartnerNamespace struct {
	pulumi.CustomResourceState

	// This boolean is used to enable or disable local auth. Default value is false. When the property is set to true, only AAD token will be used to authenticate if user is allowed to publish to the partner namespace.
	DisableLocalAuth pulumi.BoolPtrOutput `pulumi:"disableLocalAuth"`
	// Endpoint for the partner namespace.
	Endpoint pulumi.StringOutput `pulumi:"endpoint"`
	// This can be used to restrict traffic from specific IPs instead of all IPs. Note: These are considered only if PublicNetworkAccess is enabled.
	InboundIpRules InboundIpRuleResponseArrayOutput `pulumi:"inboundIpRules"`
	// Location of the resource.
	Location pulumi.StringOutput `pulumi:"location"`
	// Name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The fully qualified ARM Id of the partner registration that should be associated with this partner namespace. This takes the following format:
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/partnerRegistrations/{partnerRegistrationName}.
	PartnerRegistrationFullyQualifiedId pulumi.StringPtrOutput `pulumi:"partnerRegistrationFullyQualifiedId"`
	// This determines if events published to this partner namespace should use the source attribute in the event payload
	// or use the channel name in the header when matching to the partner topic. If none is specified, source attribute routing will be used to match the partner topic.
	PartnerTopicRoutingMode    pulumi.StringPtrOutput                       `pulumi:"partnerTopicRoutingMode"`
	PrivateEndpointConnections PrivateEndpointConnectionResponseArrayOutput `pulumi:"privateEndpointConnections"`
	// Provisioning state of the partner namespace.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// This determines if traffic is allowed over public network. By default it is enabled.
	// You can further restrict to specific IPs by configuring <seealso cref="P:Microsoft.Azure.Events.ResourceProvider.Common.Contracts.PartnerNamespaceProperties.InboundIpRules" />
	PublicNetworkAccess pulumi.StringPtrOutput `pulumi:"publicNetworkAccess"`
	// The system metadata relating to Partner Namespace resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Tags of the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewPartnerNamespace registers a new resource with the given unique name, arguments, and options.
func NewPartnerNamespace(ctx *pulumi.Context,
	name string, args *PartnerNamespaceArgs, opts ...pulumi.ResourceOption) (*PartnerNamespace, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.DisableLocalAuth == nil {
		args.DisableLocalAuth = pulumi.BoolPtr(false)
	}
	if args.PartnerTopicRoutingMode == nil {
		args.PartnerTopicRoutingMode = pulumi.StringPtr("SourceEventAttribute")
	}
	if args.PublicNetworkAccess == nil {
		args.PublicNetworkAccess = pulumi.StringPtr("Enabled")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:eventgrid:PartnerNamespace"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20200401preview:PartnerNamespace"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20201015preview:PartnerNamespace"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20210601preview:PartnerNamespace"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20211015preview:PartnerNamespace"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20230601preview:PartnerNamespace"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource PartnerNamespace
	err := ctx.RegisterResource("azure-native:eventgrid/v20220615:PartnerNamespace", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPartnerNamespace gets an existing PartnerNamespace resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPartnerNamespace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PartnerNamespaceState, opts ...pulumi.ResourceOption) (*PartnerNamespace, error) {
	var resource PartnerNamespace
	err := ctx.ReadResource("azure-native:eventgrid/v20220615:PartnerNamespace", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PartnerNamespace resources.
type partnerNamespaceState struct {
}

type PartnerNamespaceState struct {
}

func (PartnerNamespaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*partnerNamespaceState)(nil)).Elem()
}

type partnerNamespaceArgs struct {
	// This boolean is used to enable or disable local auth. Default value is false. When the property is set to true, only AAD token will be used to authenticate if user is allowed to publish to the partner namespace.
	DisableLocalAuth *bool `pulumi:"disableLocalAuth"`
	// This can be used to restrict traffic from specific IPs instead of all IPs. Note: These are considered only if PublicNetworkAccess is enabled.
	InboundIpRules []InboundIpRule `pulumi:"inboundIpRules"`
	// Location of the resource.
	Location *string `pulumi:"location"`
	// Name of the partner namespace.
	PartnerNamespaceName *string `pulumi:"partnerNamespaceName"`
	// The fully qualified ARM Id of the partner registration that should be associated with this partner namespace. This takes the following format:
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/partnerRegistrations/{partnerRegistrationName}.
	PartnerRegistrationFullyQualifiedId *string `pulumi:"partnerRegistrationFullyQualifiedId"`
	// This determines if events published to this partner namespace should use the source attribute in the event payload
	// or use the channel name in the header when matching to the partner topic. If none is specified, source attribute routing will be used to match the partner topic.
	PartnerTopicRoutingMode *string `pulumi:"partnerTopicRoutingMode"`
	// This determines if traffic is allowed over public network. By default it is enabled.
	// You can further restrict to specific IPs by configuring <seealso cref="P:Microsoft.Azure.Events.ResourceProvider.Common.Contracts.PartnerNamespaceProperties.InboundIpRules" />
	PublicNetworkAccess *string `pulumi:"publicNetworkAccess"`
	// The name of the resource group within the user's subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Tags of the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a PartnerNamespace resource.
type PartnerNamespaceArgs struct {
	// This boolean is used to enable or disable local auth. Default value is false. When the property is set to true, only AAD token will be used to authenticate if user is allowed to publish to the partner namespace.
	DisableLocalAuth pulumi.BoolPtrInput
	// This can be used to restrict traffic from specific IPs instead of all IPs. Note: These are considered only if PublicNetworkAccess is enabled.
	InboundIpRules InboundIpRuleArrayInput
	// Location of the resource.
	Location pulumi.StringPtrInput
	// Name of the partner namespace.
	PartnerNamespaceName pulumi.StringPtrInput
	// The fully qualified ARM Id of the partner registration that should be associated with this partner namespace. This takes the following format:
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/partnerRegistrations/{partnerRegistrationName}.
	PartnerRegistrationFullyQualifiedId pulumi.StringPtrInput
	// This determines if events published to this partner namespace should use the source attribute in the event payload
	// or use the channel name in the header when matching to the partner topic. If none is specified, source attribute routing will be used to match the partner topic.
	PartnerTopicRoutingMode pulumi.StringPtrInput
	// This determines if traffic is allowed over public network. By default it is enabled.
	// You can further restrict to specific IPs by configuring <seealso cref="P:Microsoft.Azure.Events.ResourceProvider.Common.Contracts.PartnerNamespaceProperties.InboundIpRules" />
	PublicNetworkAccess pulumi.StringPtrInput
	// The name of the resource group within the user's subscription.
	ResourceGroupName pulumi.StringInput
	// Tags of the resource.
	Tags pulumi.StringMapInput
}

func (PartnerNamespaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*partnerNamespaceArgs)(nil)).Elem()
}

type PartnerNamespaceInput interface {
	pulumi.Input

	ToPartnerNamespaceOutput() PartnerNamespaceOutput
	ToPartnerNamespaceOutputWithContext(ctx context.Context) PartnerNamespaceOutput
}

func (*PartnerNamespace) ElementType() reflect.Type {
	return reflect.TypeOf((**PartnerNamespace)(nil)).Elem()
}

func (i *PartnerNamespace) ToPartnerNamespaceOutput() PartnerNamespaceOutput {
	return i.ToPartnerNamespaceOutputWithContext(context.Background())
}

func (i *PartnerNamespace) ToPartnerNamespaceOutputWithContext(ctx context.Context) PartnerNamespaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PartnerNamespaceOutput)
}

func (i *PartnerNamespace) ToOutput(ctx context.Context) pulumix.Output[*PartnerNamespace] {
	return pulumix.Output[*PartnerNamespace]{
		OutputState: i.ToPartnerNamespaceOutputWithContext(ctx).OutputState,
	}
}

type PartnerNamespaceOutput struct{ *pulumi.OutputState }

func (PartnerNamespaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PartnerNamespace)(nil)).Elem()
}

func (o PartnerNamespaceOutput) ToPartnerNamespaceOutput() PartnerNamespaceOutput {
	return o
}

func (o PartnerNamespaceOutput) ToPartnerNamespaceOutputWithContext(ctx context.Context) PartnerNamespaceOutput {
	return o
}

func (o PartnerNamespaceOutput) ToOutput(ctx context.Context) pulumix.Output[*PartnerNamespace] {
	return pulumix.Output[*PartnerNamespace]{
		OutputState: o.OutputState,
	}
}

// This boolean is used to enable or disable local auth. Default value is false. When the property is set to true, only AAD token will be used to authenticate if user is allowed to publish to the partner namespace.
func (o PartnerNamespaceOutput) DisableLocalAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *PartnerNamespace) pulumi.BoolPtrOutput { return v.DisableLocalAuth }).(pulumi.BoolPtrOutput)
}

// Endpoint for the partner namespace.
func (o PartnerNamespaceOutput) Endpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *PartnerNamespace) pulumi.StringOutput { return v.Endpoint }).(pulumi.StringOutput)
}

// This can be used to restrict traffic from specific IPs instead of all IPs. Note: These are considered only if PublicNetworkAccess is enabled.
func (o PartnerNamespaceOutput) InboundIpRules() InboundIpRuleResponseArrayOutput {
	return o.ApplyT(func(v *PartnerNamespace) InboundIpRuleResponseArrayOutput { return v.InboundIpRules }).(InboundIpRuleResponseArrayOutput)
}

// Location of the resource.
func (o PartnerNamespaceOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *PartnerNamespace) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Name of the resource.
func (o PartnerNamespaceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PartnerNamespace) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The fully qualified ARM Id of the partner registration that should be associated with this partner namespace. This takes the following format:
// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/partnerRegistrations/{partnerRegistrationName}.
func (o PartnerNamespaceOutput) PartnerRegistrationFullyQualifiedId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PartnerNamespace) pulumi.StringPtrOutput { return v.PartnerRegistrationFullyQualifiedId }).(pulumi.StringPtrOutput)
}

// This determines if events published to this partner namespace should use the source attribute in the event payload
// or use the channel name in the header when matching to the partner topic. If none is specified, source attribute routing will be used to match the partner topic.
func (o PartnerNamespaceOutput) PartnerTopicRoutingMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PartnerNamespace) pulumi.StringPtrOutput { return v.PartnerTopicRoutingMode }).(pulumi.StringPtrOutput)
}

func (o PartnerNamespaceOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v *PartnerNamespace) PrivateEndpointConnectionResponseArrayOutput {
		return v.PrivateEndpointConnections
	}).(PrivateEndpointConnectionResponseArrayOutput)
}

// Provisioning state of the partner namespace.
func (o PartnerNamespaceOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *PartnerNamespace) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// This determines if traffic is allowed over public network. By default it is enabled.
// You can further restrict to specific IPs by configuring <seealso cref="P:Microsoft.Azure.Events.ResourceProvider.Common.Contracts.PartnerNamespaceProperties.InboundIpRules" />
func (o PartnerNamespaceOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PartnerNamespace) pulumi.StringPtrOutput { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

// The system metadata relating to Partner Namespace resource.
func (o PartnerNamespaceOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *PartnerNamespace) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Tags of the resource.
func (o PartnerNamespaceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *PartnerNamespace) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Type of the resource.
func (o PartnerNamespaceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PartnerNamespace) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PartnerNamespaceOutput{})
}
