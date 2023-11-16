// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230701preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets an existing CDN endpoint with the specified endpoint name under the specified subscription, resource group and profile.
func LookupEndpoint(ctx *pulumi.Context, args *LookupEndpointArgs, opts ...pulumi.InvokeOption) (*LookupEndpointResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupEndpointResult
	err := ctx.Invoke("azure-native:cdn/v20230701preview:getEndpoint", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupEndpointArgs struct {
	// Name of the endpoint under the profile which is unique globally.
	EndpointName string `pulumi:"endpointName"`
	// Name of the CDN profile which is unique within the resource group.
	ProfileName string `pulumi:"profileName"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// CDN endpoint is the entity within a CDN profile containing configuration information such as origin, protocol, content caching and delivery behavior. The CDN endpoint uses the URL format <endpointname>.azureedge.net.
type LookupEndpointResult struct {
	// List of content types on which compression applies. The value should be a valid MIME type.
	ContentTypesToCompress []string `pulumi:"contentTypesToCompress"`
	// The custom domains under the endpoint.
	CustomDomains []DeepCreatedCustomDomainResponse `pulumi:"customDomains"`
	// A reference to the origin group.
	DefaultOriginGroup *ResourceReferenceResponse `pulumi:"defaultOriginGroup"`
	// A policy that specifies the delivery rules to be used for an endpoint.
	DeliveryPolicy *EndpointPropertiesUpdateParametersResponseDeliveryPolicy `pulumi:"deliveryPolicy"`
	// List of rules defining the user's geo access within a CDN endpoint. Each geo filter defines an access rule to a specified path or content, e.g. block APAC for path /pictures/
	GeoFilters []GeoFilterResponse `pulumi:"geoFilters"`
	// The host name of the endpoint structured as {endpointName}.{DNSZone}, e.g. contoso.azureedge.net
	HostName string `pulumi:"hostName"`
	// Resource ID.
	Id string `pulumi:"id"`
	// Indicates whether content compression is enabled on CDN. Default value is false. If compression is enabled, content will be served as compressed if user requests for a compressed version. Content won't be compressed on CDN when requested content is smaller than 1 byte or larger than 1 MB.
	IsCompressionEnabled *bool `pulumi:"isCompressionEnabled"`
	// Indicates whether HTTP traffic is allowed on the endpoint. Default value is true. At least one protocol (HTTP or HTTPS) must be allowed.
	IsHttpAllowed *bool `pulumi:"isHttpAllowed"`
	// Indicates whether HTTPS traffic is allowed on the endpoint. Default value is true. At least one protocol (HTTP or HTTPS) must be allowed.
	IsHttpsAllowed *bool `pulumi:"isHttpsAllowed"`
	// Resource location.
	Location string `pulumi:"location"`
	// Resource name.
	Name string `pulumi:"name"`
	// Specifies what scenario the customer wants this CDN endpoint to optimize for, e.g. Download, Media services. With this information, CDN can apply scenario driven optimization.
	OptimizationType *string `pulumi:"optimizationType"`
	// The origin groups comprising of origins that are used for load balancing the traffic based on availability.
	OriginGroups []DeepCreatedOriginGroupResponse `pulumi:"originGroups"`
	// The host header value sent to the origin with each request. This property at Endpoint is only allowed when endpoint uses single origin and can be overridden by the same property specified at origin.If you leave this blank, the request hostname determines this value. Azure CDN origins, such as Web Apps, Blob Storage, and Cloud Services require this host header value to match the origin hostname by default.
	OriginHostHeader *string `pulumi:"originHostHeader"`
	// A directory path on the origin that CDN can use to retrieve content from, e.g. contoso.cloudapp.net/originpath.
	OriginPath *string `pulumi:"originPath"`
	// The source of the content being delivered via CDN.
	Origins []DeepCreatedOriginResponse `pulumi:"origins"`
	// Path to a file hosted on the origin which helps accelerate delivery of the dynamic content and calculate the most optimal routes for the CDN. This is relative to the origin path. This property is only relevant when using a single origin.
	ProbePath *string `pulumi:"probePath"`
	// Provisioning status of the endpoint.
	ProvisioningState string `pulumi:"provisioningState"`
	// Defines how CDN caches requests that include query strings. You can ignore any query strings when caching, bypass caching to prevent requests that contain query strings from being cached, or cache every request with a unique URL.
	QueryStringCachingBehavior *string `pulumi:"queryStringCachingBehavior"`
	// Resource status of the endpoint.
	ResourceState string `pulumi:"resourceState"`
	// Read only system data
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type string `pulumi:"type"`
	// List of keys used to validate the signed URL hashes.
	UrlSigningKeys []UrlSigningKeyResponse `pulumi:"urlSigningKeys"`
	// Defines the Web Application Firewall policy for the endpoint (if applicable)
	WebApplicationFirewallPolicyLink *EndpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLink `pulumi:"webApplicationFirewallPolicyLink"`
}

// Defaults sets the appropriate defaults for LookupEndpointResult
func (val *LookupEndpointResult) Defaults() *LookupEndpointResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.IsHttpAllowed == nil {
		isHttpAllowed_ := true
		tmp.IsHttpAllowed = &isHttpAllowed_
	}
	if tmp.IsHttpsAllowed == nil {
		isHttpsAllowed_ := true
		tmp.IsHttpsAllowed = &isHttpsAllowed_
	}
	if tmp.QueryStringCachingBehavior == nil {
		queryStringCachingBehavior_ := "NotSet"
		tmp.QueryStringCachingBehavior = &queryStringCachingBehavior_
	}
	return &tmp
}

func LookupEndpointOutput(ctx *pulumi.Context, args LookupEndpointOutputArgs, opts ...pulumi.InvokeOption) LookupEndpointResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEndpointResult, error) {
			args := v.(LookupEndpointArgs)
			r, err := LookupEndpoint(ctx, &args, opts...)
			var s LookupEndpointResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupEndpointResultOutput)
}

type LookupEndpointOutputArgs struct {
	// Name of the endpoint under the profile which is unique globally.
	EndpointName pulumi.StringInput `pulumi:"endpointName"`
	// Name of the CDN profile which is unique within the resource group.
	ProfileName pulumi.StringInput `pulumi:"profileName"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupEndpointOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEndpointArgs)(nil)).Elem()
}

// CDN endpoint is the entity within a CDN profile containing configuration information such as origin, protocol, content caching and delivery behavior. The CDN endpoint uses the URL format <endpointname>.azureedge.net.
type LookupEndpointResultOutput struct{ *pulumi.OutputState }

func (LookupEndpointResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEndpointResult)(nil)).Elem()
}

func (o LookupEndpointResultOutput) ToLookupEndpointResultOutput() LookupEndpointResultOutput {
	return o
}

func (o LookupEndpointResultOutput) ToLookupEndpointResultOutputWithContext(ctx context.Context) LookupEndpointResultOutput {
	return o
}

// List of content types on which compression applies. The value should be a valid MIME type.
func (o LookupEndpointResultOutput) ContentTypesToCompress() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupEndpointResult) []string { return v.ContentTypesToCompress }).(pulumi.StringArrayOutput)
}

// The custom domains under the endpoint.
func (o LookupEndpointResultOutput) CustomDomains() DeepCreatedCustomDomainResponseArrayOutput {
	return o.ApplyT(func(v LookupEndpointResult) []DeepCreatedCustomDomainResponse { return v.CustomDomains }).(DeepCreatedCustomDomainResponseArrayOutput)
}

// A reference to the origin group.
func (o LookupEndpointResultOutput) DefaultOriginGroup() ResourceReferenceResponsePtrOutput {
	return o.ApplyT(func(v LookupEndpointResult) *ResourceReferenceResponse { return v.DefaultOriginGroup }).(ResourceReferenceResponsePtrOutput)
}

// A policy that specifies the delivery rules to be used for an endpoint.
func (o LookupEndpointResultOutput) DeliveryPolicy() EndpointPropertiesUpdateParametersResponseDeliveryPolicyPtrOutput {
	return o.ApplyT(func(v LookupEndpointResult) *EndpointPropertiesUpdateParametersResponseDeliveryPolicy {
		return v.DeliveryPolicy
	}).(EndpointPropertiesUpdateParametersResponseDeliveryPolicyPtrOutput)
}

// List of rules defining the user's geo access within a CDN endpoint. Each geo filter defines an access rule to a specified path or content, e.g. block APAC for path /pictures/
func (o LookupEndpointResultOutput) GeoFilters() GeoFilterResponseArrayOutput {
	return o.ApplyT(func(v LookupEndpointResult) []GeoFilterResponse { return v.GeoFilters }).(GeoFilterResponseArrayOutput)
}

// The host name of the endpoint structured as {endpointName}.{DNSZone}, e.g. contoso.azureedge.net
func (o LookupEndpointResultOutput) HostName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEndpointResult) string { return v.HostName }).(pulumi.StringOutput)
}

// Resource ID.
func (o LookupEndpointResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEndpointResult) string { return v.Id }).(pulumi.StringOutput)
}

// Indicates whether content compression is enabled on CDN. Default value is false. If compression is enabled, content will be served as compressed if user requests for a compressed version. Content won't be compressed on CDN when requested content is smaller than 1 byte or larger than 1 MB.
func (o LookupEndpointResultOutput) IsCompressionEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupEndpointResult) *bool { return v.IsCompressionEnabled }).(pulumi.BoolPtrOutput)
}

// Indicates whether HTTP traffic is allowed on the endpoint. Default value is true. At least one protocol (HTTP or HTTPS) must be allowed.
func (o LookupEndpointResultOutput) IsHttpAllowed() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupEndpointResult) *bool { return v.IsHttpAllowed }).(pulumi.BoolPtrOutput)
}

// Indicates whether HTTPS traffic is allowed on the endpoint. Default value is true. At least one protocol (HTTP or HTTPS) must be allowed.
func (o LookupEndpointResultOutput) IsHttpsAllowed() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupEndpointResult) *bool { return v.IsHttpsAllowed }).(pulumi.BoolPtrOutput)
}

// Resource location.
func (o LookupEndpointResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEndpointResult) string { return v.Location }).(pulumi.StringOutput)
}

// Resource name.
func (o LookupEndpointResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEndpointResult) string { return v.Name }).(pulumi.StringOutput)
}

// Specifies what scenario the customer wants this CDN endpoint to optimize for, e.g. Download, Media services. With this information, CDN can apply scenario driven optimization.
func (o LookupEndpointResultOutput) OptimizationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEndpointResult) *string { return v.OptimizationType }).(pulumi.StringPtrOutput)
}

// The origin groups comprising of origins that are used for load balancing the traffic based on availability.
func (o LookupEndpointResultOutput) OriginGroups() DeepCreatedOriginGroupResponseArrayOutput {
	return o.ApplyT(func(v LookupEndpointResult) []DeepCreatedOriginGroupResponse { return v.OriginGroups }).(DeepCreatedOriginGroupResponseArrayOutput)
}

// The host header value sent to the origin with each request. This property at Endpoint is only allowed when endpoint uses single origin and can be overridden by the same property specified at origin.If you leave this blank, the request hostname determines this value. Azure CDN origins, such as Web Apps, Blob Storage, and Cloud Services require this host header value to match the origin hostname by default.
func (o LookupEndpointResultOutput) OriginHostHeader() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEndpointResult) *string { return v.OriginHostHeader }).(pulumi.StringPtrOutput)
}

// A directory path on the origin that CDN can use to retrieve content from, e.g. contoso.cloudapp.net/originpath.
func (o LookupEndpointResultOutput) OriginPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEndpointResult) *string { return v.OriginPath }).(pulumi.StringPtrOutput)
}

// The source of the content being delivered via CDN.
func (o LookupEndpointResultOutput) Origins() DeepCreatedOriginResponseArrayOutput {
	return o.ApplyT(func(v LookupEndpointResult) []DeepCreatedOriginResponse { return v.Origins }).(DeepCreatedOriginResponseArrayOutput)
}

// Path to a file hosted on the origin which helps accelerate delivery of the dynamic content and calculate the most optimal routes for the CDN. This is relative to the origin path. This property is only relevant when using a single origin.
func (o LookupEndpointResultOutput) ProbePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEndpointResult) *string { return v.ProbePath }).(pulumi.StringPtrOutput)
}

// Provisioning status of the endpoint.
func (o LookupEndpointResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEndpointResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Defines how CDN caches requests that include query strings. You can ignore any query strings when caching, bypass caching to prevent requests that contain query strings from being cached, or cache every request with a unique URL.
func (o LookupEndpointResultOutput) QueryStringCachingBehavior() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEndpointResult) *string { return v.QueryStringCachingBehavior }).(pulumi.StringPtrOutput)
}

// Resource status of the endpoint.
func (o LookupEndpointResultOutput) ResourceState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEndpointResult) string { return v.ResourceState }).(pulumi.StringOutput)
}

// Read only system data
func (o LookupEndpointResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupEndpointResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupEndpointResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupEndpointResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o LookupEndpointResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEndpointResult) string { return v.Type }).(pulumi.StringOutput)
}

// List of keys used to validate the signed URL hashes.
func (o LookupEndpointResultOutput) UrlSigningKeys() UrlSigningKeyResponseArrayOutput {
	return o.ApplyT(func(v LookupEndpointResult) []UrlSigningKeyResponse { return v.UrlSigningKeys }).(UrlSigningKeyResponseArrayOutput)
}

// Defines the Web Application Firewall policy for the endpoint (if applicable)
func (o LookupEndpointResultOutput) WebApplicationFirewallPolicyLink() EndpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLinkPtrOutput {
	return o.ApplyT(func(v LookupEndpointResult) *EndpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLink {
		return v.WebApplicationFirewallPolicyLink
	}).(EndpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLinkPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEndpointResultOutput{})
}
