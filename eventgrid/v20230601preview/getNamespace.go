// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get properties of a namespace.
func LookupNamespace(ctx *pulumi.Context, args *LookupNamespaceArgs, opts ...pulumi.InvokeOption) (*LookupNamespaceResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupNamespaceResult
	err := ctx.Invoke("azure-native:eventgrid/v20230601preview:getNamespace", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupNamespaceArgs struct {
	// Name of the namespace.
	NamespaceName string `pulumi:"namespaceName"`
	// The name of the resource group within the user's subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Namespace resource.
type LookupNamespaceResult struct {
	// Fully qualified identifier of the resource.
	Id string `pulumi:"id"`
	// Identity information for the Namespace resource.
	Identity *IdentityInfoResponse `pulumi:"identity"`
	// This can be used to restrict traffic from specific IPs instead of all IPs. Note: These are considered only if PublicNetworkAccess is enabled.
	InboundIpRules []InboundIpRuleResponse `pulumi:"inboundIpRules"`
	// Allows the user to specify if the service is zone-redundant. This is a required property and user needs to specify this value explicitly.
	// Once specified, this property cannot be updated.
	IsZoneRedundant *bool `pulumi:"isZoneRedundant"`
	// Location of the resource.
	Location string `pulumi:"location"`
	// Minimum TLS version of the publisher allowed to publish to this namespace. Only TLS version 1.2 is supported.
	MinimumTlsVersionAllowed *string `pulumi:"minimumTlsVersionAllowed"`
	// Name of the resource.
	Name                       string                              `pulumi:"name"`
	PrivateEndpointConnections []PrivateEndpointConnectionResponse `pulumi:"privateEndpointConnections"`
	// Provisioning state of the namespace resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// This determines if traffic is allowed over public network. By default it is enabled.
	// You can further restrict to specific IPs by configuring <seealso cref="P:Microsoft.Azure.Events.ResourceProvider.Common.Contracts.PubSub.NamespaceProperties.InboundIpRules" />
	PublicNetworkAccess *string `pulumi:"publicNetworkAccess"`
	// Represents available Sku pricing tiers.
	Sku *NamespaceSkuResponse `pulumi:"sku"`
	// The system metadata relating to the namespace resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Tags of the resource.
	Tags map[string]string `pulumi:"tags"`
	// Topic spaces configuration information for the namespace resource
	TopicSpacesConfiguration *TopicSpacesConfigurationResponse `pulumi:"topicSpacesConfiguration"`
	// Topics configuration information for the namespace resource
	TopicsConfiguration *TopicsConfigurationResponse `pulumi:"topicsConfiguration"`
	// Type of the resource.
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupNamespaceResult
func (val *LookupNamespaceResult) Defaults() *LookupNamespaceResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.TopicSpacesConfiguration = tmp.TopicSpacesConfiguration.Defaults()

	return &tmp
}

func LookupNamespaceOutput(ctx *pulumi.Context, args LookupNamespaceOutputArgs, opts ...pulumi.InvokeOption) LookupNamespaceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNamespaceResult, error) {
			args := v.(LookupNamespaceArgs)
			r, err := LookupNamespace(ctx, &args, opts...)
			var s LookupNamespaceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNamespaceResultOutput)
}

type LookupNamespaceOutputArgs struct {
	// Name of the namespace.
	NamespaceName pulumi.StringInput `pulumi:"namespaceName"`
	// The name of the resource group within the user's subscription.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupNamespaceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNamespaceArgs)(nil)).Elem()
}

// Namespace resource.
type LookupNamespaceResultOutput struct{ *pulumi.OutputState }

func (LookupNamespaceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNamespaceResult)(nil)).Elem()
}

func (o LookupNamespaceResultOutput) ToLookupNamespaceResultOutput() LookupNamespaceResultOutput {
	return o
}

func (o LookupNamespaceResultOutput) ToLookupNamespaceResultOutputWithContext(ctx context.Context) LookupNamespaceResultOutput {
	return o
}

func (o LookupNamespaceResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupNamespaceResult] {
	return pulumix.Output[LookupNamespaceResult]{
		OutputState: o.OutputState,
	}
}

// Fully qualified identifier of the resource.
func (o LookupNamespaceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNamespaceResult) string { return v.Id }).(pulumi.StringOutput)
}

// Identity information for the Namespace resource.
func (o LookupNamespaceResultOutput) Identity() IdentityInfoResponsePtrOutput {
	return o.ApplyT(func(v LookupNamespaceResult) *IdentityInfoResponse { return v.Identity }).(IdentityInfoResponsePtrOutput)
}

// This can be used to restrict traffic from specific IPs instead of all IPs. Note: These are considered only if PublicNetworkAccess is enabled.
func (o LookupNamespaceResultOutput) InboundIpRules() InboundIpRuleResponseArrayOutput {
	return o.ApplyT(func(v LookupNamespaceResult) []InboundIpRuleResponse { return v.InboundIpRules }).(InboundIpRuleResponseArrayOutput)
}

// Allows the user to specify if the service is zone-redundant. This is a required property and user needs to specify this value explicitly.
// Once specified, this property cannot be updated.
func (o LookupNamespaceResultOutput) IsZoneRedundant() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupNamespaceResult) *bool { return v.IsZoneRedundant }).(pulumi.BoolPtrOutput)
}

// Location of the resource.
func (o LookupNamespaceResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNamespaceResult) string { return v.Location }).(pulumi.StringOutput)
}

// Minimum TLS version of the publisher allowed to publish to this namespace. Only TLS version 1.2 is supported.
func (o LookupNamespaceResultOutput) MinimumTlsVersionAllowed() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNamespaceResult) *string { return v.MinimumTlsVersionAllowed }).(pulumi.StringPtrOutput)
}

// Name of the resource.
func (o LookupNamespaceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNamespaceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupNamespaceResultOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v LookupNamespaceResult) []PrivateEndpointConnectionResponse { return v.PrivateEndpointConnections }).(PrivateEndpointConnectionResponseArrayOutput)
}

// Provisioning state of the namespace resource.
func (o LookupNamespaceResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNamespaceResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// This determines if traffic is allowed over public network. By default it is enabled.
// You can further restrict to specific IPs by configuring <seealso cref="P:Microsoft.Azure.Events.ResourceProvider.Common.Contracts.PubSub.NamespaceProperties.InboundIpRules" />
func (o LookupNamespaceResultOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNamespaceResult) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

// Represents available Sku pricing tiers.
func (o LookupNamespaceResultOutput) Sku() NamespaceSkuResponsePtrOutput {
	return o.ApplyT(func(v LookupNamespaceResult) *NamespaceSkuResponse { return v.Sku }).(NamespaceSkuResponsePtrOutput)
}

// The system metadata relating to the namespace resource.
func (o LookupNamespaceResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupNamespaceResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Tags of the resource.
func (o LookupNamespaceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupNamespaceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Topic spaces configuration information for the namespace resource
func (o LookupNamespaceResultOutput) TopicSpacesConfiguration() TopicSpacesConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupNamespaceResult) *TopicSpacesConfigurationResponse { return v.TopicSpacesConfiguration }).(TopicSpacesConfigurationResponsePtrOutput)
}

// Topics configuration information for the namespace resource
func (o LookupNamespaceResultOutput) TopicsConfiguration() TopicsConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupNamespaceResult) *TopicsConfigurationResponse { return v.TopicsConfiguration }).(TopicsConfigurationResponsePtrOutput)
}

// Type of the resource.
func (o LookupNamespaceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNamespaceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNamespaceResultOutput{})
}
