// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the specified application gateway.
func LookupApplicationGateway(ctx *pulumi.Context, args *LookupApplicationGatewayArgs, opts ...pulumi.InvokeOption) (*LookupApplicationGatewayResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupApplicationGatewayResult
	err := ctx.Invoke("azure-native:network/v20230601:getApplicationGateway", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApplicationGatewayArgs struct {
	// The name of the application gateway.
	ApplicationGatewayName string `pulumi:"applicationGatewayName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Application gateway resource.
type LookupApplicationGatewayResult struct {
	// Authentication certificates of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
	AuthenticationCertificates []ApplicationGatewayAuthenticationCertificateResponse `pulumi:"authenticationCertificates"`
	// Autoscale Configuration.
	AutoscaleConfiguration *ApplicationGatewayAutoscaleConfigurationResponse `pulumi:"autoscaleConfiguration"`
	// Backend address pool of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
	BackendAddressPools []ApplicationGatewayBackendAddressPoolResponse `pulumi:"backendAddressPools"`
	// Backend http settings of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
	BackendHttpSettingsCollection []ApplicationGatewayBackendHttpSettingsResponse `pulumi:"backendHttpSettingsCollection"`
	// Backend settings of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
	BackendSettingsCollection []ApplicationGatewayBackendSettingsResponse `pulumi:"backendSettingsCollection"`
	// Custom error configurations of the application gateway resource.
	CustomErrorConfigurations []ApplicationGatewayCustomErrorResponse `pulumi:"customErrorConfigurations"`
	// The default predefined SSL Policy applied on the application gateway resource.
	DefaultPredefinedSslPolicy string `pulumi:"defaultPredefinedSslPolicy"`
	// Whether FIPS is enabled on the application gateway resource.
	EnableFips *bool `pulumi:"enableFips"`
	// Whether HTTP2 is enabled on the application gateway resource.
	EnableHttp2 *bool `pulumi:"enableHttp2"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// Reference to the FirewallPolicy resource.
	FirewallPolicy *SubResourceResponse `pulumi:"firewallPolicy"`
	// If true, associates a firewall policy with an application gateway regardless whether the policy differs from the WAF Config.
	ForceFirewallPolicyAssociation *bool `pulumi:"forceFirewallPolicyAssociation"`
	// Frontend IP addresses of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
	FrontendIPConfigurations []ApplicationGatewayFrontendIPConfigurationResponse `pulumi:"frontendIPConfigurations"`
	// Frontend ports of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
	FrontendPorts []ApplicationGatewayFrontendPortResponse `pulumi:"frontendPorts"`
	// Subnets of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
	GatewayIPConfigurations []ApplicationGatewayIPConfigurationResponse `pulumi:"gatewayIPConfigurations"`
	// Global Configuration.
	GlobalConfiguration *ApplicationGatewayGlobalConfigurationResponse `pulumi:"globalConfiguration"`
	// Http listeners of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
	HttpListeners []ApplicationGatewayHttpListenerResponse `pulumi:"httpListeners"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// The identity of the application gateway, if configured.
	Identity *ManagedServiceIdentityResponse `pulumi:"identity"`
	// Listeners of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
	Listeners []ApplicationGatewayListenerResponse `pulumi:"listeners"`
	// Load distribution policies of the application gateway resource.
	LoadDistributionPolicies []ApplicationGatewayLoadDistributionPolicyResponse `pulumi:"loadDistributionPolicies"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Resource name.
	Name string `pulumi:"name"`
	// Operational state of the application gateway resource.
	OperationalState string `pulumi:"operationalState"`
	// Private Endpoint connections on application gateway.
	PrivateEndpointConnections []ApplicationGatewayPrivateEndpointConnectionResponse `pulumi:"privateEndpointConnections"`
	// PrivateLink configurations on application gateway.
	PrivateLinkConfigurations []ApplicationGatewayPrivateLinkConfigurationResponse `pulumi:"privateLinkConfigurations"`
	// Probes of the application gateway resource.
	Probes []ApplicationGatewayProbeResponse `pulumi:"probes"`
	// The provisioning state of the application gateway resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Redirect configurations of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
	RedirectConfigurations []ApplicationGatewayRedirectConfigurationResponse `pulumi:"redirectConfigurations"`
	// Request routing rules of the application gateway resource.
	RequestRoutingRules []ApplicationGatewayRequestRoutingRuleResponse `pulumi:"requestRoutingRules"`
	// The resource GUID property of the application gateway resource.
	ResourceGuid string `pulumi:"resourceGuid"`
	// Rewrite rules for the application gateway resource.
	RewriteRuleSets []ApplicationGatewayRewriteRuleSetResponse `pulumi:"rewriteRuleSets"`
	// Routing rules of the application gateway resource.
	RoutingRules []ApplicationGatewayRoutingRuleResponse `pulumi:"routingRules"`
	// SKU of the application gateway resource.
	Sku *ApplicationGatewaySkuResponse `pulumi:"sku"`
	// SSL certificates of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
	SslCertificates []ApplicationGatewaySslCertificateResponse `pulumi:"sslCertificates"`
	// SSL policy of the application gateway resource.
	SslPolicy *ApplicationGatewaySslPolicyResponse `pulumi:"sslPolicy"`
	// SSL profiles of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
	SslProfiles []ApplicationGatewaySslProfileResponse `pulumi:"sslProfiles"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Trusted client certificates of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
	TrustedClientCertificates []ApplicationGatewayTrustedClientCertificateResponse `pulumi:"trustedClientCertificates"`
	// Trusted Root certificates of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
	TrustedRootCertificates []ApplicationGatewayTrustedRootCertificateResponse `pulumi:"trustedRootCertificates"`
	// Resource type.
	Type string `pulumi:"type"`
	// URL path map of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
	UrlPathMaps []ApplicationGatewayUrlPathMapResponse `pulumi:"urlPathMaps"`
	// Web application firewall configuration.
	WebApplicationFirewallConfiguration *ApplicationGatewayWebApplicationFirewallConfigurationResponse `pulumi:"webApplicationFirewallConfiguration"`
	// A list of availability zones denoting where the resource needs to come from.
	Zones []string `pulumi:"zones"`
}

func LookupApplicationGatewayOutput(ctx *pulumi.Context, args LookupApplicationGatewayOutputArgs, opts ...pulumi.InvokeOption) LookupApplicationGatewayResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupApplicationGatewayResult, error) {
			args := v.(LookupApplicationGatewayArgs)
			r, err := LookupApplicationGateway(ctx, &args, opts...)
			var s LookupApplicationGatewayResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupApplicationGatewayResultOutput)
}

type LookupApplicationGatewayOutputArgs struct {
	// The name of the application gateway.
	ApplicationGatewayName pulumi.StringInput `pulumi:"applicationGatewayName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupApplicationGatewayOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplicationGatewayArgs)(nil)).Elem()
}

// Application gateway resource.
type LookupApplicationGatewayResultOutput struct{ *pulumi.OutputState }

func (LookupApplicationGatewayResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplicationGatewayResult)(nil)).Elem()
}

func (o LookupApplicationGatewayResultOutput) ToLookupApplicationGatewayResultOutput() LookupApplicationGatewayResultOutput {
	return o
}

func (o LookupApplicationGatewayResultOutput) ToLookupApplicationGatewayResultOutputWithContext(ctx context.Context) LookupApplicationGatewayResultOutput {
	return o
}

// Authentication certificates of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
func (o LookupApplicationGatewayResultOutput) AuthenticationCertificates() ApplicationGatewayAuthenticationCertificateResponseArrayOutput {
	return o.ApplyT(func(v LookupApplicationGatewayResult) []ApplicationGatewayAuthenticationCertificateResponse {
		return v.AuthenticationCertificates
	}).(ApplicationGatewayAuthenticationCertificateResponseArrayOutput)
}

// Autoscale Configuration.
func (o LookupApplicationGatewayResultOutput) AutoscaleConfiguration() ApplicationGatewayAutoscaleConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupApplicationGatewayResult) *ApplicationGatewayAutoscaleConfigurationResponse {
		return v.AutoscaleConfiguration
	}).(ApplicationGatewayAutoscaleConfigurationResponsePtrOutput)
}

// Backend address pool of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
func (o LookupApplicationGatewayResultOutput) BackendAddressPools() ApplicationGatewayBackendAddressPoolResponseArrayOutput {
	return o.ApplyT(func(v LookupApplicationGatewayResult) []ApplicationGatewayBackendAddressPoolResponse {
		return v.BackendAddressPools
	}).(ApplicationGatewayBackendAddressPoolResponseArrayOutput)
}

// Backend http settings of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
func (o LookupApplicationGatewayResultOutput) BackendHttpSettingsCollection() ApplicationGatewayBackendHttpSettingsResponseArrayOutput {
	return o.ApplyT(func(v LookupApplicationGatewayResult) []ApplicationGatewayBackendHttpSettingsResponse {
		return v.BackendHttpSettingsCollection
	}).(ApplicationGatewayBackendHttpSettingsResponseArrayOutput)
}

// Backend settings of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
func (o LookupApplicationGatewayResultOutput) BackendSettingsCollection() ApplicationGatewayBackendSettingsResponseArrayOutput {
	return o.ApplyT(func(v LookupApplicationGatewayResult) []ApplicationGatewayBackendSettingsResponse {
		return v.BackendSettingsCollection
	}).(ApplicationGatewayBackendSettingsResponseArrayOutput)
}

// Custom error configurations of the application gateway resource.
func (o LookupApplicationGatewayResultOutput) CustomErrorConfigurations() ApplicationGatewayCustomErrorResponseArrayOutput {
	return o.ApplyT(func(v LookupApplicationGatewayResult) []ApplicationGatewayCustomErrorResponse {
		return v.CustomErrorConfigurations
	}).(ApplicationGatewayCustomErrorResponseArrayOutput)
}

// The default predefined SSL Policy applied on the application gateway resource.
func (o LookupApplicationGatewayResultOutput) DefaultPredefinedSslPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationGatewayResult) string { return v.DefaultPredefinedSslPolicy }).(pulumi.StringOutput)
}

// Whether FIPS is enabled on the application gateway resource.
func (o LookupApplicationGatewayResultOutput) EnableFips() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupApplicationGatewayResult) *bool { return v.EnableFips }).(pulumi.BoolPtrOutput)
}

// Whether HTTP2 is enabled on the application gateway resource.
func (o LookupApplicationGatewayResultOutput) EnableHttp2() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupApplicationGatewayResult) *bool { return v.EnableHttp2 }).(pulumi.BoolPtrOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o LookupApplicationGatewayResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationGatewayResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Reference to the FirewallPolicy resource.
func (o LookupApplicationGatewayResultOutput) FirewallPolicy() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupApplicationGatewayResult) *SubResourceResponse { return v.FirewallPolicy }).(SubResourceResponsePtrOutput)
}

// If true, associates a firewall policy with an application gateway regardless whether the policy differs from the WAF Config.
func (o LookupApplicationGatewayResultOutput) ForceFirewallPolicyAssociation() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupApplicationGatewayResult) *bool { return v.ForceFirewallPolicyAssociation }).(pulumi.BoolPtrOutput)
}

// Frontend IP addresses of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
func (o LookupApplicationGatewayResultOutput) FrontendIPConfigurations() ApplicationGatewayFrontendIPConfigurationResponseArrayOutput {
	return o.ApplyT(func(v LookupApplicationGatewayResult) []ApplicationGatewayFrontendIPConfigurationResponse {
		return v.FrontendIPConfigurations
	}).(ApplicationGatewayFrontendIPConfigurationResponseArrayOutput)
}

// Frontend ports of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
func (o LookupApplicationGatewayResultOutput) FrontendPorts() ApplicationGatewayFrontendPortResponseArrayOutput {
	return o.ApplyT(func(v LookupApplicationGatewayResult) []ApplicationGatewayFrontendPortResponse {
		return v.FrontendPorts
	}).(ApplicationGatewayFrontendPortResponseArrayOutput)
}

// Subnets of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
func (o LookupApplicationGatewayResultOutput) GatewayIPConfigurations() ApplicationGatewayIPConfigurationResponseArrayOutput {
	return o.ApplyT(func(v LookupApplicationGatewayResult) []ApplicationGatewayIPConfigurationResponse {
		return v.GatewayIPConfigurations
	}).(ApplicationGatewayIPConfigurationResponseArrayOutput)
}

// Global Configuration.
func (o LookupApplicationGatewayResultOutput) GlobalConfiguration() ApplicationGatewayGlobalConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupApplicationGatewayResult) *ApplicationGatewayGlobalConfigurationResponse {
		return v.GlobalConfiguration
	}).(ApplicationGatewayGlobalConfigurationResponsePtrOutput)
}

// Http listeners of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
func (o LookupApplicationGatewayResultOutput) HttpListeners() ApplicationGatewayHttpListenerResponseArrayOutput {
	return o.ApplyT(func(v LookupApplicationGatewayResult) []ApplicationGatewayHttpListenerResponse {
		return v.HttpListeners
	}).(ApplicationGatewayHttpListenerResponseArrayOutput)
}

// Resource ID.
func (o LookupApplicationGatewayResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationGatewayResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// The identity of the application gateway, if configured.
func (o LookupApplicationGatewayResultOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupApplicationGatewayResult) *ManagedServiceIdentityResponse { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

// Listeners of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
func (o LookupApplicationGatewayResultOutput) Listeners() ApplicationGatewayListenerResponseArrayOutput {
	return o.ApplyT(func(v LookupApplicationGatewayResult) []ApplicationGatewayListenerResponse { return v.Listeners }).(ApplicationGatewayListenerResponseArrayOutput)
}

// Load distribution policies of the application gateway resource.
func (o LookupApplicationGatewayResultOutput) LoadDistributionPolicies() ApplicationGatewayLoadDistributionPolicyResponseArrayOutput {
	return o.ApplyT(func(v LookupApplicationGatewayResult) []ApplicationGatewayLoadDistributionPolicyResponse {
		return v.LoadDistributionPolicies
	}).(ApplicationGatewayLoadDistributionPolicyResponseArrayOutput)
}

// Resource location.
func (o LookupApplicationGatewayResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationGatewayResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o LookupApplicationGatewayResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationGatewayResult) string { return v.Name }).(pulumi.StringOutput)
}

// Operational state of the application gateway resource.
func (o LookupApplicationGatewayResultOutput) OperationalState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationGatewayResult) string { return v.OperationalState }).(pulumi.StringOutput)
}

// Private Endpoint connections on application gateway.
func (o LookupApplicationGatewayResultOutput) PrivateEndpointConnections() ApplicationGatewayPrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v LookupApplicationGatewayResult) []ApplicationGatewayPrivateEndpointConnectionResponse {
		return v.PrivateEndpointConnections
	}).(ApplicationGatewayPrivateEndpointConnectionResponseArrayOutput)
}

// PrivateLink configurations on application gateway.
func (o LookupApplicationGatewayResultOutput) PrivateLinkConfigurations() ApplicationGatewayPrivateLinkConfigurationResponseArrayOutput {
	return o.ApplyT(func(v LookupApplicationGatewayResult) []ApplicationGatewayPrivateLinkConfigurationResponse {
		return v.PrivateLinkConfigurations
	}).(ApplicationGatewayPrivateLinkConfigurationResponseArrayOutput)
}

// Probes of the application gateway resource.
func (o LookupApplicationGatewayResultOutput) Probes() ApplicationGatewayProbeResponseArrayOutput {
	return o.ApplyT(func(v LookupApplicationGatewayResult) []ApplicationGatewayProbeResponse { return v.Probes }).(ApplicationGatewayProbeResponseArrayOutput)
}

// The provisioning state of the application gateway resource.
func (o LookupApplicationGatewayResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationGatewayResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Redirect configurations of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
func (o LookupApplicationGatewayResultOutput) RedirectConfigurations() ApplicationGatewayRedirectConfigurationResponseArrayOutput {
	return o.ApplyT(func(v LookupApplicationGatewayResult) []ApplicationGatewayRedirectConfigurationResponse {
		return v.RedirectConfigurations
	}).(ApplicationGatewayRedirectConfigurationResponseArrayOutput)
}

// Request routing rules of the application gateway resource.
func (o LookupApplicationGatewayResultOutput) RequestRoutingRules() ApplicationGatewayRequestRoutingRuleResponseArrayOutput {
	return o.ApplyT(func(v LookupApplicationGatewayResult) []ApplicationGatewayRequestRoutingRuleResponse {
		return v.RequestRoutingRules
	}).(ApplicationGatewayRequestRoutingRuleResponseArrayOutput)
}

// The resource GUID property of the application gateway resource.
func (o LookupApplicationGatewayResultOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationGatewayResult) string { return v.ResourceGuid }).(pulumi.StringOutput)
}

// Rewrite rules for the application gateway resource.
func (o LookupApplicationGatewayResultOutput) RewriteRuleSets() ApplicationGatewayRewriteRuleSetResponseArrayOutput {
	return o.ApplyT(func(v LookupApplicationGatewayResult) []ApplicationGatewayRewriteRuleSetResponse {
		return v.RewriteRuleSets
	}).(ApplicationGatewayRewriteRuleSetResponseArrayOutput)
}

// Routing rules of the application gateway resource.
func (o LookupApplicationGatewayResultOutput) RoutingRules() ApplicationGatewayRoutingRuleResponseArrayOutput {
	return o.ApplyT(func(v LookupApplicationGatewayResult) []ApplicationGatewayRoutingRuleResponse { return v.RoutingRules }).(ApplicationGatewayRoutingRuleResponseArrayOutput)
}

// SKU of the application gateway resource.
func (o LookupApplicationGatewayResultOutput) Sku() ApplicationGatewaySkuResponsePtrOutput {
	return o.ApplyT(func(v LookupApplicationGatewayResult) *ApplicationGatewaySkuResponse { return v.Sku }).(ApplicationGatewaySkuResponsePtrOutput)
}

// SSL certificates of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
func (o LookupApplicationGatewayResultOutput) SslCertificates() ApplicationGatewaySslCertificateResponseArrayOutput {
	return o.ApplyT(func(v LookupApplicationGatewayResult) []ApplicationGatewaySslCertificateResponse {
		return v.SslCertificates
	}).(ApplicationGatewaySslCertificateResponseArrayOutput)
}

// SSL policy of the application gateway resource.
func (o LookupApplicationGatewayResultOutput) SslPolicy() ApplicationGatewaySslPolicyResponsePtrOutput {
	return o.ApplyT(func(v LookupApplicationGatewayResult) *ApplicationGatewaySslPolicyResponse { return v.SslPolicy }).(ApplicationGatewaySslPolicyResponsePtrOutput)
}

// SSL profiles of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
func (o LookupApplicationGatewayResultOutput) SslProfiles() ApplicationGatewaySslProfileResponseArrayOutput {
	return o.ApplyT(func(v LookupApplicationGatewayResult) []ApplicationGatewaySslProfileResponse { return v.SslProfiles }).(ApplicationGatewaySslProfileResponseArrayOutput)
}

// Resource tags.
func (o LookupApplicationGatewayResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupApplicationGatewayResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Trusted client certificates of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
func (o LookupApplicationGatewayResultOutput) TrustedClientCertificates() ApplicationGatewayTrustedClientCertificateResponseArrayOutput {
	return o.ApplyT(func(v LookupApplicationGatewayResult) []ApplicationGatewayTrustedClientCertificateResponse {
		return v.TrustedClientCertificates
	}).(ApplicationGatewayTrustedClientCertificateResponseArrayOutput)
}

// Trusted Root certificates of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
func (o LookupApplicationGatewayResultOutput) TrustedRootCertificates() ApplicationGatewayTrustedRootCertificateResponseArrayOutput {
	return o.ApplyT(func(v LookupApplicationGatewayResult) []ApplicationGatewayTrustedRootCertificateResponse {
		return v.TrustedRootCertificates
	}).(ApplicationGatewayTrustedRootCertificateResponseArrayOutput)
}

// Resource type.
func (o LookupApplicationGatewayResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationGatewayResult) string { return v.Type }).(pulumi.StringOutput)
}

// URL path map of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
func (o LookupApplicationGatewayResultOutput) UrlPathMaps() ApplicationGatewayUrlPathMapResponseArrayOutput {
	return o.ApplyT(func(v LookupApplicationGatewayResult) []ApplicationGatewayUrlPathMapResponse { return v.UrlPathMaps }).(ApplicationGatewayUrlPathMapResponseArrayOutput)
}

// Web application firewall configuration.
func (o LookupApplicationGatewayResultOutput) WebApplicationFirewallConfiguration() ApplicationGatewayWebApplicationFirewallConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupApplicationGatewayResult) *ApplicationGatewayWebApplicationFirewallConfigurationResponse {
		return v.WebApplicationFirewallConfiguration
	}).(ApplicationGatewayWebApplicationFirewallConfigurationResponsePtrOutput)
}

// A list of availability zones denoting where the resource needs to come from.
func (o LookupApplicationGatewayResultOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupApplicationGatewayResult) []string { return v.Zones }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApplicationGatewayResultOutput{})
}
