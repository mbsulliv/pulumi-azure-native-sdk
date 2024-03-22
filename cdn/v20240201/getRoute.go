// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20240201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets an existing route with the specified route name under the specified subscription, resource group, profile, and AzureFrontDoor endpoint.
func LookupRoute(ctx *pulumi.Context, args *LookupRouteArgs, opts ...pulumi.InvokeOption) (*LookupRouteResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupRouteResult
	err := ctx.Invoke("azure-native:cdn/v20240201:getRoute", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupRouteArgs struct {
	// Name of the endpoint under the profile which is unique globally.
	EndpointName string `pulumi:"endpointName"`
	// Name of the Azure Front Door Standard or Azure Front Door Premium profile which is unique within the resource group.
	ProfileName string `pulumi:"profileName"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the routing rule.
	RouteName string `pulumi:"routeName"`
}

// Friendly Routes name mapping to the any Routes or secret related information.
type LookupRouteResult struct {
	// The caching configuration for this route. To disable caching, do not provide a cacheConfiguration object.
	CacheConfiguration *AfdRouteCacheConfigurationResponse `pulumi:"cacheConfiguration"`
	// Domains referenced by this endpoint.
	CustomDomains    []ActivatedResourceReferenceResponse `pulumi:"customDomains"`
	DeploymentStatus string                               `pulumi:"deploymentStatus"`
	// Whether to enable use of this rule. Permitted values are 'Enabled' or 'Disabled'
	EnabledState *string `pulumi:"enabledState"`
	// The name of the endpoint which holds the route.
	EndpointName string `pulumi:"endpointName"`
	// Protocol this rule will use when forwarding traffic to backends.
	ForwardingProtocol *string `pulumi:"forwardingProtocol"`
	// Whether to automatically redirect HTTP traffic to HTTPS traffic. Note that this is a easy way to set up this rule and it will be the first rule that gets executed.
	HttpsRedirect *string `pulumi:"httpsRedirect"`
	// Resource ID.
	Id string `pulumi:"id"`
	// whether this route will be linked to the default endpoint domain.
	LinkToDefaultDomain *string `pulumi:"linkToDefaultDomain"`
	// Resource name.
	Name string `pulumi:"name"`
	// A reference to the origin group.
	OriginGroup ResourceReferenceResponse `pulumi:"originGroup"`
	// A directory path on the origin that AzureFrontDoor can use to retrieve content from, e.g. contoso.cloudapp.net/originpath.
	OriginPath *string `pulumi:"originPath"`
	// The route patterns of the rule.
	PatternsToMatch []string `pulumi:"patternsToMatch"`
	// Provisioning status
	ProvisioningState string `pulumi:"provisioningState"`
	// rule sets referenced by this endpoint.
	RuleSets []ResourceReferenceResponse `pulumi:"ruleSets"`
	// List of supported protocols for this route.
	SupportedProtocols []string `pulumi:"supportedProtocols"`
	// Read only system data
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource type.
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupRouteResult
func (val *LookupRouteResult) Defaults() *LookupRouteResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.ForwardingProtocol == nil {
		forwardingProtocol_ := "MatchRequest"
		tmp.ForwardingProtocol = &forwardingProtocol_
	}
	if tmp.HttpsRedirect == nil {
		httpsRedirect_ := "Disabled"
		tmp.HttpsRedirect = &httpsRedirect_
	}
	if tmp.LinkToDefaultDomain == nil {
		linkToDefaultDomain_ := "Disabled"
		tmp.LinkToDefaultDomain = &linkToDefaultDomain_
	}
	return &tmp
}

func LookupRouteOutput(ctx *pulumi.Context, args LookupRouteOutputArgs, opts ...pulumi.InvokeOption) LookupRouteResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRouteResult, error) {
			args := v.(LookupRouteArgs)
			r, err := LookupRoute(ctx, &args, opts...)
			var s LookupRouteResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRouteResultOutput)
}

type LookupRouteOutputArgs struct {
	// Name of the endpoint under the profile which is unique globally.
	EndpointName pulumi.StringInput `pulumi:"endpointName"`
	// Name of the Azure Front Door Standard or Azure Front Door Premium profile which is unique within the resource group.
	ProfileName pulumi.StringInput `pulumi:"profileName"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name of the routing rule.
	RouteName pulumi.StringInput `pulumi:"routeName"`
}

func (LookupRouteOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRouteArgs)(nil)).Elem()
}

// Friendly Routes name mapping to the any Routes or secret related information.
type LookupRouteResultOutput struct{ *pulumi.OutputState }

func (LookupRouteResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRouteResult)(nil)).Elem()
}

func (o LookupRouteResultOutput) ToLookupRouteResultOutput() LookupRouteResultOutput {
	return o
}

func (o LookupRouteResultOutput) ToLookupRouteResultOutputWithContext(ctx context.Context) LookupRouteResultOutput {
	return o
}

// The caching configuration for this route. To disable caching, do not provide a cacheConfiguration object.
func (o LookupRouteResultOutput) CacheConfiguration() AfdRouteCacheConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupRouteResult) *AfdRouteCacheConfigurationResponse { return v.CacheConfiguration }).(AfdRouteCacheConfigurationResponsePtrOutput)
}

// Domains referenced by this endpoint.
func (o LookupRouteResultOutput) CustomDomains() ActivatedResourceReferenceResponseArrayOutput {
	return o.ApplyT(func(v LookupRouteResult) []ActivatedResourceReferenceResponse { return v.CustomDomains }).(ActivatedResourceReferenceResponseArrayOutput)
}

func (o LookupRouteResultOutput) DeploymentStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteResult) string { return v.DeploymentStatus }).(pulumi.StringOutput)
}

// Whether to enable use of this rule. Permitted values are 'Enabled' or 'Disabled'
func (o LookupRouteResultOutput) EnabledState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRouteResult) *string { return v.EnabledState }).(pulumi.StringPtrOutput)
}

// The name of the endpoint which holds the route.
func (o LookupRouteResultOutput) EndpointName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteResult) string { return v.EndpointName }).(pulumi.StringOutput)
}

// Protocol this rule will use when forwarding traffic to backends.
func (o LookupRouteResultOutput) ForwardingProtocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRouteResult) *string { return v.ForwardingProtocol }).(pulumi.StringPtrOutput)
}

// Whether to automatically redirect HTTP traffic to HTTPS traffic. Note that this is a easy way to set up this rule and it will be the first rule that gets executed.
func (o LookupRouteResultOutput) HttpsRedirect() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRouteResult) *string { return v.HttpsRedirect }).(pulumi.StringPtrOutput)
}

// Resource ID.
func (o LookupRouteResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteResult) string { return v.Id }).(pulumi.StringOutput)
}

// whether this route will be linked to the default endpoint domain.
func (o LookupRouteResultOutput) LinkToDefaultDomain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRouteResult) *string { return v.LinkToDefaultDomain }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o LookupRouteResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteResult) string { return v.Name }).(pulumi.StringOutput)
}

// A reference to the origin group.
func (o LookupRouteResultOutput) OriginGroup() ResourceReferenceResponseOutput {
	return o.ApplyT(func(v LookupRouteResult) ResourceReferenceResponse { return v.OriginGroup }).(ResourceReferenceResponseOutput)
}

// A directory path on the origin that AzureFrontDoor can use to retrieve content from, e.g. contoso.cloudapp.net/originpath.
func (o LookupRouteResultOutput) OriginPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRouteResult) *string { return v.OriginPath }).(pulumi.StringPtrOutput)
}

// The route patterns of the rule.
func (o LookupRouteResultOutput) PatternsToMatch() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupRouteResult) []string { return v.PatternsToMatch }).(pulumi.StringArrayOutput)
}

// Provisioning status
func (o LookupRouteResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// rule sets referenced by this endpoint.
func (o LookupRouteResultOutput) RuleSets() ResourceReferenceResponseArrayOutput {
	return o.ApplyT(func(v LookupRouteResult) []ResourceReferenceResponse { return v.RuleSets }).(ResourceReferenceResponseArrayOutput)
}

// List of supported protocols for this route.
func (o LookupRouteResultOutput) SupportedProtocols() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupRouteResult) []string { return v.SupportedProtocols }).(pulumi.StringArrayOutput)
}

// Read only system data
func (o LookupRouteResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupRouteResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource type.
func (o LookupRouteResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRouteResultOutput{})
}
