// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230601preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Namespace resource.
type Namespace struct {
	pulumi.CustomResourceState

	// Identity information for the Namespace resource.
	Identity IdentityInfoResponsePtrOutput `pulumi:"identity"`
	// This can be used to restrict traffic from specific IPs instead of all IPs. Note: These are considered only if PublicNetworkAccess is enabled.
	InboundIpRules InboundIpRuleResponseArrayOutput `pulumi:"inboundIpRules"`
	// Allows the user to specify if the service is zone-redundant. This is a required property and user needs to specify this value explicitly.
	// Once specified, this property cannot be updated.
	IsZoneRedundant pulumi.BoolPtrOutput `pulumi:"isZoneRedundant"`
	// Location of the resource.
	Location pulumi.StringOutput `pulumi:"location"`
	// Minimum TLS version of the publisher allowed to publish to this namespace. Only TLS version 1.2 is supported.
	MinimumTlsVersionAllowed pulumi.StringPtrOutput `pulumi:"minimumTlsVersionAllowed"`
	// Name of the resource.
	Name                       pulumi.StringOutput                          `pulumi:"name"`
	PrivateEndpointConnections PrivateEndpointConnectionResponseArrayOutput `pulumi:"privateEndpointConnections"`
	// Provisioning state of the namespace resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// This determines if traffic is allowed over public network. By default it is enabled.
	// You can further restrict to specific IPs by configuring <seealso cref="P:Microsoft.Azure.Events.ResourceProvider.Common.Contracts.PubSub.NamespaceProperties.InboundIpRules" />
	PublicNetworkAccess pulumi.StringPtrOutput `pulumi:"publicNetworkAccess"`
	// Represents available Sku pricing tiers.
	Sku NamespaceSkuResponsePtrOutput `pulumi:"sku"`
	// The system metadata relating to the namespace resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Tags of the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Topic spaces configuration information for the namespace resource
	TopicSpacesConfiguration TopicSpacesConfigurationResponsePtrOutput `pulumi:"topicSpacesConfiguration"`
	// Topics configuration information for the namespace resource
	TopicsConfiguration TopicsConfigurationResponsePtrOutput `pulumi:"topicsConfiguration"`
	// Type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewNamespace registers a new resource with the given unique name, arguments, and options.
func NewNamespace(ctx *pulumi.Context,
	name string, args *NamespaceArgs, opts ...pulumi.ResourceOption) (*Namespace, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.TopicSpacesConfiguration != nil {
		args.TopicSpacesConfiguration = args.TopicSpacesConfiguration.ToTopicSpacesConfigurationPtrOutput().ApplyT(func(v *TopicSpacesConfiguration) *TopicSpacesConfiguration { return v.Defaults() }).(TopicSpacesConfigurationPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:eventgrid:Namespace"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20231215preview:Namespace"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Namespace
	err := ctx.RegisterResource("azure-native:eventgrid/v20230601preview:Namespace", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNamespace gets an existing Namespace resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNamespace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NamespaceState, opts ...pulumi.ResourceOption) (*Namespace, error) {
	var resource Namespace
	err := ctx.ReadResource("azure-native:eventgrid/v20230601preview:Namespace", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Namespace resources.
type namespaceState struct {
}

type NamespaceState struct {
}

func (NamespaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*namespaceState)(nil)).Elem()
}

type namespaceArgs struct {
	// Identity information for the Namespace resource.
	Identity *IdentityInfo `pulumi:"identity"`
	// This can be used to restrict traffic from specific IPs instead of all IPs. Note: These are considered only if PublicNetworkAccess is enabled.
	InboundIpRules []InboundIpRule `pulumi:"inboundIpRules"`
	// Allows the user to specify if the service is zone-redundant. This is a required property and user needs to specify this value explicitly.
	// Once specified, this property cannot be updated.
	IsZoneRedundant *bool `pulumi:"isZoneRedundant"`
	// Location of the resource.
	Location *string `pulumi:"location"`
	// Minimum TLS version of the publisher allowed to publish to this namespace. Only TLS version 1.2 is supported.
	MinimumTlsVersionAllowed *string `pulumi:"minimumTlsVersionAllowed"`
	// Name of the namespace.
	NamespaceName              *string                         `pulumi:"namespaceName"`
	PrivateEndpointConnections []PrivateEndpointConnectionType `pulumi:"privateEndpointConnections"`
	// This determines if traffic is allowed over public network. By default it is enabled.
	// You can further restrict to specific IPs by configuring <seealso cref="P:Microsoft.Azure.Events.ResourceProvider.Common.Contracts.PubSub.NamespaceProperties.InboundIpRules" />
	PublicNetworkAccess *string `pulumi:"publicNetworkAccess"`
	// The name of the resource group within the user's subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Represents available Sku pricing tiers.
	Sku *NamespaceSku `pulumi:"sku"`
	// Tags of the resource.
	Tags map[string]string `pulumi:"tags"`
	// Topic spaces configuration information for the namespace resource
	TopicSpacesConfiguration *TopicSpacesConfiguration `pulumi:"topicSpacesConfiguration"`
}

// The set of arguments for constructing a Namespace resource.
type NamespaceArgs struct {
	// Identity information for the Namespace resource.
	Identity IdentityInfoPtrInput
	// This can be used to restrict traffic from specific IPs instead of all IPs. Note: These are considered only if PublicNetworkAccess is enabled.
	InboundIpRules InboundIpRuleArrayInput
	// Allows the user to specify if the service is zone-redundant. This is a required property and user needs to specify this value explicitly.
	// Once specified, this property cannot be updated.
	IsZoneRedundant pulumi.BoolPtrInput
	// Location of the resource.
	Location pulumi.StringPtrInput
	// Minimum TLS version of the publisher allowed to publish to this namespace. Only TLS version 1.2 is supported.
	MinimumTlsVersionAllowed pulumi.StringPtrInput
	// Name of the namespace.
	NamespaceName              pulumi.StringPtrInput
	PrivateEndpointConnections PrivateEndpointConnectionTypeArrayInput
	// This determines if traffic is allowed over public network. By default it is enabled.
	// You can further restrict to specific IPs by configuring <seealso cref="P:Microsoft.Azure.Events.ResourceProvider.Common.Contracts.PubSub.NamespaceProperties.InboundIpRules" />
	PublicNetworkAccess pulumi.StringPtrInput
	// The name of the resource group within the user's subscription.
	ResourceGroupName pulumi.StringInput
	// Represents available Sku pricing tiers.
	Sku NamespaceSkuPtrInput
	// Tags of the resource.
	Tags pulumi.StringMapInput
	// Topic spaces configuration information for the namespace resource
	TopicSpacesConfiguration TopicSpacesConfigurationPtrInput
}

func (NamespaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*namespaceArgs)(nil)).Elem()
}

type NamespaceInput interface {
	pulumi.Input

	ToNamespaceOutput() NamespaceOutput
	ToNamespaceOutputWithContext(ctx context.Context) NamespaceOutput
}

func (*Namespace) ElementType() reflect.Type {
	return reflect.TypeOf((**Namespace)(nil)).Elem()
}

func (i *Namespace) ToNamespaceOutput() NamespaceOutput {
	return i.ToNamespaceOutputWithContext(context.Background())
}

func (i *Namespace) ToNamespaceOutputWithContext(ctx context.Context) NamespaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamespaceOutput)
}

type NamespaceOutput struct{ *pulumi.OutputState }

func (NamespaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Namespace)(nil)).Elem()
}

func (o NamespaceOutput) ToNamespaceOutput() NamespaceOutput {
	return o
}

func (o NamespaceOutput) ToNamespaceOutputWithContext(ctx context.Context) NamespaceOutput {
	return o
}

// Identity information for the Namespace resource.
func (o NamespaceOutput) Identity() IdentityInfoResponsePtrOutput {
	return o.ApplyT(func(v *Namespace) IdentityInfoResponsePtrOutput { return v.Identity }).(IdentityInfoResponsePtrOutput)
}

// This can be used to restrict traffic from specific IPs instead of all IPs. Note: These are considered only if PublicNetworkAccess is enabled.
func (o NamespaceOutput) InboundIpRules() InboundIpRuleResponseArrayOutput {
	return o.ApplyT(func(v *Namespace) InboundIpRuleResponseArrayOutput { return v.InboundIpRules }).(InboundIpRuleResponseArrayOutput)
}

// Allows the user to specify if the service is zone-redundant. This is a required property and user needs to specify this value explicitly.
// Once specified, this property cannot be updated.
func (o NamespaceOutput) IsZoneRedundant() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Namespace) pulumi.BoolPtrOutput { return v.IsZoneRedundant }).(pulumi.BoolPtrOutput)
}

// Location of the resource.
func (o NamespaceOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Namespace) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Minimum TLS version of the publisher allowed to publish to this namespace. Only TLS version 1.2 is supported.
func (o NamespaceOutput) MinimumTlsVersionAllowed() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Namespace) pulumi.StringPtrOutput { return v.MinimumTlsVersionAllowed }).(pulumi.StringPtrOutput)
}

// Name of the resource.
func (o NamespaceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Namespace) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o NamespaceOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v *Namespace) PrivateEndpointConnectionResponseArrayOutput { return v.PrivateEndpointConnections }).(PrivateEndpointConnectionResponseArrayOutput)
}

// Provisioning state of the namespace resource.
func (o NamespaceOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Namespace) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// This determines if traffic is allowed over public network. By default it is enabled.
// You can further restrict to specific IPs by configuring <seealso cref="P:Microsoft.Azure.Events.ResourceProvider.Common.Contracts.PubSub.NamespaceProperties.InboundIpRules" />
func (o NamespaceOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Namespace) pulumi.StringPtrOutput { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

// Represents available Sku pricing tiers.
func (o NamespaceOutput) Sku() NamespaceSkuResponsePtrOutput {
	return o.ApplyT(func(v *Namespace) NamespaceSkuResponsePtrOutput { return v.Sku }).(NamespaceSkuResponsePtrOutput)
}

// The system metadata relating to the namespace resource.
func (o NamespaceOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Namespace) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Tags of the resource.
func (o NamespaceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Namespace) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Topic spaces configuration information for the namespace resource
func (o NamespaceOutput) TopicSpacesConfiguration() TopicSpacesConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *Namespace) TopicSpacesConfigurationResponsePtrOutput { return v.TopicSpacesConfiguration }).(TopicSpacesConfigurationResponsePtrOutput)
}

// Topics configuration information for the namespace resource
func (o NamespaceOutput) TopicsConfiguration() TopicsConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *Namespace) TopicsConfigurationResponsePtrOutput { return v.TopicsConfiguration }).(TopicsConfigurationResponsePtrOutput)
}

// Type of the resource.
func (o NamespaceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Namespace) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(NamespaceOutput{})
}
