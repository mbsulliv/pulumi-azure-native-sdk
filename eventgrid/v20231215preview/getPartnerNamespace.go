// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20231215preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get properties of a partner namespace.
func LookupPartnerNamespace(ctx *pulumi.Context, args *LookupPartnerNamespaceArgs, opts ...pulumi.InvokeOption) (*LookupPartnerNamespaceResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupPartnerNamespaceResult
	err := ctx.Invoke("azure-native:eventgrid/v20231215preview:getPartnerNamespace", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupPartnerNamespaceArgs struct {
	// Name of the partner namespace.
	PartnerNamespaceName string `pulumi:"partnerNamespaceName"`
	// The name of the resource group within the user's subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// EventGrid Partner Namespace.
type LookupPartnerNamespaceResult struct {
	// This boolean is used to enable or disable local auth. Default value is false. When the property is set to true, only AAD token will be used to authenticate if user is allowed to publish to the partner namespace.
	DisableLocalAuth *bool `pulumi:"disableLocalAuth"`
	// Endpoint for the partner namespace.
	Endpoint string `pulumi:"endpoint"`
	// Fully qualified identifier of the resource.
	Id string `pulumi:"id"`
	// This can be used to restrict traffic from specific IPs instead of all IPs. Note: These are considered only if PublicNetworkAccess is enabled.
	InboundIpRules []InboundIpRuleResponse `pulumi:"inboundIpRules"`
	// Location of the resource.
	Location string `pulumi:"location"`
	// Minimum TLS version of the publisher allowed to publish to this partner namespace
	MinimumTlsVersionAllowed *string `pulumi:"minimumTlsVersionAllowed"`
	// Name of the resource.
	Name string `pulumi:"name"`
	// The fully qualified ARM Id of the partner registration that should be associated with this partner namespace. This takes the following format:
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/partnerRegistrations/{partnerRegistrationName}.
	PartnerRegistrationFullyQualifiedId *string `pulumi:"partnerRegistrationFullyQualifiedId"`
	// This determines if events published to this partner namespace should use the source attribute in the event payload
	// or use the channel name in the header when matching to the partner topic. If none is specified, source attribute routing will be used to match the partner topic.
	PartnerTopicRoutingMode    *string                             `pulumi:"partnerTopicRoutingMode"`
	PrivateEndpointConnections []PrivateEndpointConnectionResponse `pulumi:"privateEndpointConnections"`
	// Provisioning state of the partner namespace.
	ProvisioningState string `pulumi:"provisioningState"`
	// This determines if traffic is allowed over public network. By default it is enabled.
	// You can further restrict to specific IPs by configuring <seealso cref="P:Microsoft.Azure.Events.ResourceProvider.Common.Contracts.PartnerNamespaceProperties.InboundIpRules" />
	PublicNetworkAccess *string `pulumi:"publicNetworkAccess"`
	// The system metadata relating to Partner Namespace resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Tags of the resource.
	Tags map[string]string `pulumi:"tags"`
	// Type of the resource.
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupPartnerNamespaceResult
func (val *LookupPartnerNamespaceResult) Defaults() *LookupPartnerNamespaceResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.DisableLocalAuth == nil {
		disableLocalAuth_ := false
		tmp.DisableLocalAuth = &disableLocalAuth_
	}
	if tmp.PartnerTopicRoutingMode == nil {
		partnerTopicRoutingMode_ := "SourceEventAttribute"
		tmp.PartnerTopicRoutingMode = &partnerTopicRoutingMode_
	}
	if tmp.PublicNetworkAccess == nil {
		publicNetworkAccess_ := "Enabled"
		tmp.PublicNetworkAccess = &publicNetworkAccess_
	}
	return &tmp
}

func LookupPartnerNamespaceOutput(ctx *pulumi.Context, args LookupPartnerNamespaceOutputArgs, opts ...pulumi.InvokeOption) LookupPartnerNamespaceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPartnerNamespaceResult, error) {
			args := v.(LookupPartnerNamespaceArgs)
			r, err := LookupPartnerNamespace(ctx, &args, opts...)
			var s LookupPartnerNamespaceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPartnerNamespaceResultOutput)
}

type LookupPartnerNamespaceOutputArgs struct {
	// Name of the partner namespace.
	PartnerNamespaceName pulumi.StringInput `pulumi:"partnerNamespaceName"`
	// The name of the resource group within the user's subscription.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupPartnerNamespaceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPartnerNamespaceArgs)(nil)).Elem()
}

// EventGrid Partner Namespace.
type LookupPartnerNamespaceResultOutput struct{ *pulumi.OutputState }

func (LookupPartnerNamespaceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPartnerNamespaceResult)(nil)).Elem()
}

func (o LookupPartnerNamespaceResultOutput) ToLookupPartnerNamespaceResultOutput() LookupPartnerNamespaceResultOutput {
	return o
}

func (o LookupPartnerNamespaceResultOutput) ToLookupPartnerNamespaceResultOutputWithContext(ctx context.Context) LookupPartnerNamespaceResultOutput {
	return o
}

// This boolean is used to enable or disable local auth. Default value is false. When the property is set to true, only AAD token will be used to authenticate if user is allowed to publish to the partner namespace.
func (o LookupPartnerNamespaceResultOutput) DisableLocalAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupPartnerNamespaceResult) *bool { return v.DisableLocalAuth }).(pulumi.BoolPtrOutput)
}

// Endpoint for the partner namespace.
func (o LookupPartnerNamespaceResultOutput) Endpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPartnerNamespaceResult) string { return v.Endpoint }).(pulumi.StringOutput)
}

// Fully qualified identifier of the resource.
func (o LookupPartnerNamespaceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPartnerNamespaceResult) string { return v.Id }).(pulumi.StringOutput)
}

// This can be used to restrict traffic from specific IPs instead of all IPs. Note: These are considered only if PublicNetworkAccess is enabled.
func (o LookupPartnerNamespaceResultOutput) InboundIpRules() InboundIpRuleResponseArrayOutput {
	return o.ApplyT(func(v LookupPartnerNamespaceResult) []InboundIpRuleResponse { return v.InboundIpRules }).(InboundIpRuleResponseArrayOutput)
}

// Location of the resource.
func (o LookupPartnerNamespaceResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPartnerNamespaceResult) string { return v.Location }).(pulumi.StringOutput)
}

// Minimum TLS version of the publisher allowed to publish to this partner namespace
func (o LookupPartnerNamespaceResultOutput) MinimumTlsVersionAllowed() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPartnerNamespaceResult) *string { return v.MinimumTlsVersionAllowed }).(pulumi.StringPtrOutput)
}

// Name of the resource.
func (o LookupPartnerNamespaceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPartnerNamespaceResult) string { return v.Name }).(pulumi.StringOutput)
}

// The fully qualified ARM Id of the partner registration that should be associated with this partner namespace. This takes the following format:
// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/partnerRegistrations/{partnerRegistrationName}.
func (o LookupPartnerNamespaceResultOutput) PartnerRegistrationFullyQualifiedId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPartnerNamespaceResult) *string { return v.PartnerRegistrationFullyQualifiedId }).(pulumi.StringPtrOutput)
}

// This determines if events published to this partner namespace should use the source attribute in the event payload
// or use the channel name in the header when matching to the partner topic. If none is specified, source attribute routing will be used to match the partner topic.
func (o LookupPartnerNamespaceResultOutput) PartnerTopicRoutingMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPartnerNamespaceResult) *string { return v.PartnerTopicRoutingMode }).(pulumi.StringPtrOutput)
}

func (o LookupPartnerNamespaceResultOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v LookupPartnerNamespaceResult) []PrivateEndpointConnectionResponse {
		return v.PrivateEndpointConnections
	}).(PrivateEndpointConnectionResponseArrayOutput)
}

// Provisioning state of the partner namespace.
func (o LookupPartnerNamespaceResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPartnerNamespaceResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// This determines if traffic is allowed over public network. By default it is enabled.
// You can further restrict to specific IPs by configuring <seealso cref="P:Microsoft.Azure.Events.ResourceProvider.Common.Contracts.PartnerNamespaceProperties.InboundIpRules" />
func (o LookupPartnerNamespaceResultOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPartnerNamespaceResult) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

// The system metadata relating to Partner Namespace resource.
func (o LookupPartnerNamespaceResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupPartnerNamespaceResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Tags of the resource.
func (o LookupPartnerNamespaceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupPartnerNamespaceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Type of the resource.
func (o LookupPartnerNamespaceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPartnerNamespaceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPartnerNamespaceResultOutput{})
}
