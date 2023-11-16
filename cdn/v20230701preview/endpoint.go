// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230701preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// CDN endpoint is the entity within a CDN profile containing configuration information such as origin, protocol, content caching and delivery behavior. The CDN endpoint uses the URL format <endpointname>.azureedge.net.
type Endpoint struct {
	pulumi.CustomResourceState

	// List of content types on which compression applies. The value should be a valid MIME type.
	ContentTypesToCompress pulumi.StringArrayOutput `pulumi:"contentTypesToCompress"`
	// The custom domains under the endpoint.
	CustomDomains DeepCreatedCustomDomainResponseArrayOutput `pulumi:"customDomains"`
	// A reference to the origin group.
	DefaultOriginGroup ResourceReferenceResponsePtrOutput `pulumi:"defaultOriginGroup"`
	// A policy that specifies the delivery rules to be used for an endpoint.
	DeliveryPolicy EndpointPropertiesUpdateParametersResponseDeliveryPolicyPtrOutput `pulumi:"deliveryPolicy"`
	// List of rules defining the user's geo access within a CDN endpoint. Each geo filter defines an access rule to a specified path or content, e.g. block APAC for path /pictures/
	GeoFilters GeoFilterResponseArrayOutput `pulumi:"geoFilters"`
	// The host name of the endpoint structured as {endpointName}.{DNSZone}, e.g. contoso.azureedge.net
	HostName pulumi.StringOutput `pulumi:"hostName"`
	// Indicates whether content compression is enabled on CDN. Default value is false. If compression is enabled, content will be served as compressed if user requests for a compressed version. Content won't be compressed on CDN when requested content is smaller than 1 byte or larger than 1 MB.
	IsCompressionEnabled pulumi.BoolPtrOutput `pulumi:"isCompressionEnabled"`
	// Indicates whether HTTP traffic is allowed on the endpoint. Default value is true. At least one protocol (HTTP or HTTPS) must be allowed.
	IsHttpAllowed pulumi.BoolPtrOutput `pulumi:"isHttpAllowed"`
	// Indicates whether HTTPS traffic is allowed on the endpoint. Default value is true. At least one protocol (HTTP or HTTPS) must be allowed.
	IsHttpsAllowed pulumi.BoolPtrOutput `pulumi:"isHttpsAllowed"`
	// Resource location.
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies what scenario the customer wants this CDN endpoint to optimize for, e.g. Download, Media services. With this information, CDN can apply scenario driven optimization.
	OptimizationType pulumi.StringPtrOutput `pulumi:"optimizationType"`
	// The origin groups comprising of origins that are used for load balancing the traffic based on availability.
	OriginGroups DeepCreatedOriginGroupResponseArrayOutput `pulumi:"originGroups"`
	// The host header value sent to the origin with each request. This property at Endpoint is only allowed when endpoint uses single origin and can be overridden by the same property specified at origin.If you leave this blank, the request hostname determines this value. Azure CDN origins, such as Web Apps, Blob Storage, and Cloud Services require this host header value to match the origin hostname by default.
	OriginHostHeader pulumi.StringPtrOutput `pulumi:"originHostHeader"`
	// A directory path on the origin that CDN can use to retrieve content from, e.g. contoso.cloudapp.net/originpath.
	OriginPath pulumi.StringPtrOutput `pulumi:"originPath"`
	// The source of the content being delivered via CDN.
	Origins DeepCreatedOriginResponseArrayOutput `pulumi:"origins"`
	// Path to a file hosted on the origin which helps accelerate delivery of the dynamic content and calculate the most optimal routes for the CDN. This is relative to the origin path. This property is only relevant when using a single origin.
	ProbePath pulumi.StringPtrOutput `pulumi:"probePath"`
	// Provisioning status of the endpoint.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Defines how CDN caches requests that include query strings. You can ignore any query strings when caching, bypass caching to prevent requests that contain query strings from being cached, or cache every request with a unique URL.
	QueryStringCachingBehavior pulumi.StringPtrOutput `pulumi:"queryStringCachingBehavior"`
	// Resource status of the endpoint.
	ResourceState pulumi.StringOutput `pulumi:"resourceState"`
	// Read only system data
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// List of keys used to validate the signed URL hashes.
	UrlSigningKeys UrlSigningKeyResponseArrayOutput `pulumi:"urlSigningKeys"`
	// Defines the Web Application Firewall policy for the endpoint (if applicable)
	WebApplicationFirewallPolicyLink EndpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLinkPtrOutput `pulumi:"webApplicationFirewallPolicyLink"`
}

// NewEndpoint registers a new resource with the given unique name, arguments, and options.
func NewEndpoint(ctx *pulumi.Context,
	name string, args *EndpointArgs, opts ...pulumi.ResourceOption) (*Endpoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Origins == nil {
		return nil, errors.New("invalid value for required argument 'Origins'")
	}
	if args.ProfileName == nil {
		return nil, errors.New("invalid value for required argument 'ProfileName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.IsHttpAllowed == nil {
		args.IsHttpAllowed = pulumi.BoolPtr(true)
	}
	if args.IsHttpsAllowed == nil {
		args.IsHttpsAllowed = pulumi.BoolPtr(true)
	}
	if args.QueryStringCachingBehavior == nil {
		args.QueryStringCachingBehavior = QueryStringCachingBehavior("NotSet")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:cdn:Endpoint"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20150601:Endpoint"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20160402:Endpoint"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20161002:Endpoint"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20170402:Endpoint"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20171012:Endpoint"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20190415:Endpoint"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20190615:Endpoint"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20190615preview:Endpoint"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20191231:Endpoint"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20200331:Endpoint"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20200415:Endpoint"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20200901:Endpoint"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20210601:Endpoint"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20220501preview:Endpoint"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20221101preview:Endpoint"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20230501:Endpoint"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Endpoint
	err := ctx.RegisterResource("azure-native:cdn/v20230701preview:Endpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEndpoint gets an existing Endpoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EndpointState, opts ...pulumi.ResourceOption) (*Endpoint, error) {
	var resource Endpoint
	err := ctx.ReadResource("azure-native:cdn/v20230701preview:Endpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Endpoint resources.
type endpointState struct {
}

type EndpointState struct {
}

func (EndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointState)(nil)).Elem()
}

type endpointArgs struct {
	// List of content types on which compression applies. The value should be a valid MIME type.
	ContentTypesToCompress []string `pulumi:"contentTypesToCompress"`
	// A reference to the origin group.
	DefaultOriginGroup *ResourceReference `pulumi:"defaultOriginGroup"`
	// A policy that specifies the delivery rules to be used for an endpoint.
	DeliveryPolicy *EndpointPropertiesUpdateParametersDeliveryPolicy `pulumi:"deliveryPolicy"`
	// Name of the endpoint under the profile which is unique globally.
	EndpointName *string `pulumi:"endpointName"`
	// List of rules defining the user's geo access within a CDN endpoint. Each geo filter defines an access rule to a specified path or content, e.g. block APAC for path /pictures/
	GeoFilters []GeoFilter `pulumi:"geoFilters"`
	// Indicates whether content compression is enabled on CDN. Default value is false. If compression is enabled, content will be served as compressed if user requests for a compressed version. Content won't be compressed on CDN when requested content is smaller than 1 byte or larger than 1 MB.
	IsCompressionEnabled *bool `pulumi:"isCompressionEnabled"`
	// Indicates whether HTTP traffic is allowed on the endpoint. Default value is true. At least one protocol (HTTP or HTTPS) must be allowed.
	IsHttpAllowed *bool `pulumi:"isHttpAllowed"`
	// Indicates whether HTTPS traffic is allowed on the endpoint. Default value is true. At least one protocol (HTTP or HTTPS) must be allowed.
	IsHttpsAllowed *bool `pulumi:"isHttpsAllowed"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Specifies what scenario the customer wants this CDN endpoint to optimize for, e.g. Download, Media services. With this information, CDN can apply scenario driven optimization.
	OptimizationType *string `pulumi:"optimizationType"`
	// The origin groups comprising of origins that are used for load balancing the traffic based on availability.
	OriginGroups []DeepCreatedOriginGroup `pulumi:"originGroups"`
	// The host header value sent to the origin with each request. This property at Endpoint is only allowed when endpoint uses single origin and can be overridden by the same property specified at origin.If you leave this blank, the request hostname determines this value. Azure CDN origins, such as Web Apps, Blob Storage, and Cloud Services require this host header value to match the origin hostname by default.
	OriginHostHeader *string `pulumi:"originHostHeader"`
	// A directory path on the origin that CDN can use to retrieve content from, e.g. contoso.cloudapp.net/originpath.
	OriginPath *string `pulumi:"originPath"`
	// The source of the content being delivered via CDN.
	Origins []DeepCreatedOrigin `pulumi:"origins"`
	// Path to a file hosted on the origin which helps accelerate delivery of the dynamic content and calculate the most optimal routes for the CDN. This is relative to the origin path. This property is only relevant when using a single origin.
	ProbePath *string `pulumi:"probePath"`
	// Name of the CDN profile which is unique within the resource group.
	ProfileName string `pulumi:"profileName"`
	// Defines how CDN caches requests that include query strings. You can ignore any query strings when caching, bypass caching to prevent requests that contain query strings from being cached, or cache every request with a unique URL.
	QueryStringCachingBehavior *QueryStringCachingBehavior `pulumi:"queryStringCachingBehavior"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// List of keys used to validate the signed URL hashes.
	UrlSigningKeys []UrlSigningKey `pulumi:"urlSigningKeys"`
	// Defines the Web Application Firewall policy for the endpoint (if applicable)
	WebApplicationFirewallPolicyLink *EndpointPropertiesUpdateParametersWebApplicationFirewallPolicyLink `pulumi:"webApplicationFirewallPolicyLink"`
}

// The set of arguments for constructing a Endpoint resource.
type EndpointArgs struct {
	// List of content types on which compression applies. The value should be a valid MIME type.
	ContentTypesToCompress pulumi.StringArrayInput
	// A reference to the origin group.
	DefaultOriginGroup ResourceReferencePtrInput
	// A policy that specifies the delivery rules to be used for an endpoint.
	DeliveryPolicy EndpointPropertiesUpdateParametersDeliveryPolicyPtrInput
	// Name of the endpoint under the profile which is unique globally.
	EndpointName pulumi.StringPtrInput
	// List of rules defining the user's geo access within a CDN endpoint. Each geo filter defines an access rule to a specified path or content, e.g. block APAC for path /pictures/
	GeoFilters GeoFilterArrayInput
	// Indicates whether content compression is enabled on CDN. Default value is false. If compression is enabled, content will be served as compressed if user requests for a compressed version. Content won't be compressed on CDN when requested content is smaller than 1 byte or larger than 1 MB.
	IsCompressionEnabled pulumi.BoolPtrInput
	// Indicates whether HTTP traffic is allowed on the endpoint. Default value is true. At least one protocol (HTTP or HTTPS) must be allowed.
	IsHttpAllowed pulumi.BoolPtrInput
	// Indicates whether HTTPS traffic is allowed on the endpoint. Default value is true. At least one protocol (HTTP or HTTPS) must be allowed.
	IsHttpsAllowed pulumi.BoolPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// Specifies what scenario the customer wants this CDN endpoint to optimize for, e.g. Download, Media services. With this information, CDN can apply scenario driven optimization.
	OptimizationType pulumi.StringPtrInput
	// The origin groups comprising of origins that are used for load balancing the traffic based on availability.
	OriginGroups DeepCreatedOriginGroupArrayInput
	// The host header value sent to the origin with each request. This property at Endpoint is only allowed when endpoint uses single origin and can be overridden by the same property specified at origin.If you leave this blank, the request hostname determines this value. Azure CDN origins, such as Web Apps, Blob Storage, and Cloud Services require this host header value to match the origin hostname by default.
	OriginHostHeader pulumi.StringPtrInput
	// A directory path on the origin that CDN can use to retrieve content from, e.g. contoso.cloudapp.net/originpath.
	OriginPath pulumi.StringPtrInput
	// The source of the content being delivered via CDN.
	Origins DeepCreatedOriginArrayInput
	// Path to a file hosted on the origin which helps accelerate delivery of the dynamic content and calculate the most optimal routes for the CDN. This is relative to the origin path. This property is only relevant when using a single origin.
	ProbePath pulumi.StringPtrInput
	// Name of the CDN profile which is unique within the resource group.
	ProfileName pulumi.StringInput
	// Defines how CDN caches requests that include query strings. You can ignore any query strings when caching, bypass caching to prevent requests that contain query strings from being cached, or cache every request with a unique URL.
	QueryStringCachingBehavior QueryStringCachingBehaviorPtrInput
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// List of keys used to validate the signed URL hashes.
	UrlSigningKeys UrlSigningKeyArrayInput
	// Defines the Web Application Firewall policy for the endpoint (if applicable)
	WebApplicationFirewallPolicyLink EndpointPropertiesUpdateParametersWebApplicationFirewallPolicyLinkPtrInput
}

func (EndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointArgs)(nil)).Elem()
}

type EndpointInput interface {
	pulumi.Input

	ToEndpointOutput() EndpointOutput
	ToEndpointOutputWithContext(ctx context.Context) EndpointOutput
}

func (*Endpoint) ElementType() reflect.Type {
	return reflect.TypeOf((**Endpoint)(nil)).Elem()
}

func (i *Endpoint) ToEndpointOutput() EndpointOutput {
	return i.ToEndpointOutputWithContext(context.Background())
}

func (i *Endpoint) ToEndpointOutputWithContext(ctx context.Context) EndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointOutput)
}

type EndpointOutput struct{ *pulumi.OutputState }

func (EndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Endpoint)(nil)).Elem()
}

func (o EndpointOutput) ToEndpointOutput() EndpointOutput {
	return o
}

func (o EndpointOutput) ToEndpointOutputWithContext(ctx context.Context) EndpointOutput {
	return o
}

// List of content types on which compression applies. The value should be a valid MIME type.
func (o EndpointOutput) ContentTypesToCompress() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringArrayOutput { return v.ContentTypesToCompress }).(pulumi.StringArrayOutput)
}

// The custom domains under the endpoint.
func (o EndpointOutput) CustomDomains() DeepCreatedCustomDomainResponseArrayOutput {
	return o.ApplyT(func(v *Endpoint) DeepCreatedCustomDomainResponseArrayOutput { return v.CustomDomains }).(DeepCreatedCustomDomainResponseArrayOutput)
}

// A reference to the origin group.
func (o EndpointOutput) DefaultOriginGroup() ResourceReferenceResponsePtrOutput {
	return o.ApplyT(func(v *Endpoint) ResourceReferenceResponsePtrOutput { return v.DefaultOriginGroup }).(ResourceReferenceResponsePtrOutput)
}

// A policy that specifies the delivery rules to be used for an endpoint.
func (o EndpointOutput) DeliveryPolicy() EndpointPropertiesUpdateParametersResponseDeliveryPolicyPtrOutput {
	return o.ApplyT(func(v *Endpoint) EndpointPropertiesUpdateParametersResponseDeliveryPolicyPtrOutput {
		return v.DeliveryPolicy
	}).(EndpointPropertiesUpdateParametersResponseDeliveryPolicyPtrOutput)
}

// List of rules defining the user's geo access within a CDN endpoint. Each geo filter defines an access rule to a specified path or content, e.g. block APAC for path /pictures/
func (o EndpointOutput) GeoFilters() GeoFilterResponseArrayOutput {
	return o.ApplyT(func(v *Endpoint) GeoFilterResponseArrayOutput { return v.GeoFilters }).(GeoFilterResponseArrayOutput)
}

// The host name of the endpoint structured as {endpointName}.{DNSZone}, e.g. contoso.azureedge.net
func (o EndpointOutput) HostName() pulumi.StringOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringOutput { return v.HostName }).(pulumi.StringOutput)
}

// Indicates whether content compression is enabled on CDN. Default value is false. If compression is enabled, content will be served as compressed if user requests for a compressed version. Content won't be compressed on CDN when requested content is smaller than 1 byte or larger than 1 MB.
func (o EndpointOutput) IsCompressionEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.BoolPtrOutput { return v.IsCompressionEnabled }).(pulumi.BoolPtrOutput)
}

// Indicates whether HTTP traffic is allowed on the endpoint. Default value is true. At least one protocol (HTTP or HTTPS) must be allowed.
func (o EndpointOutput) IsHttpAllowed() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.BoolPtrOutput { return v.IsHttpAllowed }).(pulumi.BoolPtrOutput)
}

// Indicates whether HTTPS traffic is allowed on the endpoint. Default value is true. At least one protocol (HTTP or HTTPS) must be allowed.
func (o EndpointOutput) IsHttpsAllowed() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.BoolPtrOutput { return v.IsHttpsAllowed }).(pulumi.BoolPtrOutput)
}

// Resource location.
func (o EndpointOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Resource name.
func (o EndpointOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies what scenario the customer wants this CDN endpoint to optimize for, e.g. Download, Media services. With this information, CDN can apply scenario driven optimization.
func (o EndpointOutput) OptimizationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringPtrOutput { return v.OptimizationType }).(pulumi.StringPtrOutput)
}

// The origin groups comprising of origins that are used for load balancing the traffic based on availability.
func (o EndpointOutput) OriginGroups() DeepCreatedOriginGroupResponseArrayOutput {
	return o.ApplyT(func(v *Endpoint) DeepCreatedOriginGroupResponseArrayOutput { return v.OriginGroups }).(DeepCreatedOriginGroupResponseArrayOutput)
}

// The host header value sent to the origin with each request. This property at Endpoint is only allowed when endpoint uses single origin and can be overridden by the same property specified at origin.If you leave this blank, the request hostname determines this value. Azure CDN origins, such as Web Apps, Blob Storage, and Cloud Services require this host header value to match the origin hostname by default.
func (o EndpointOutput) OriginHostHeader() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringPtrOutput { return v.OriginHostHeader }).(pulumi.StringPtrOutput)
}

// A directory path on the origin that CDN can use to retrieve content from, e.g. contoso.cloudapp.net/originpath.
func (o EndpointOutput) OriginPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringPtrOutput { return v.OriginPath }).(pulumi.StringPtrOutput)
}

// The source of the content being delivered via CDN.
func (o EndpointOutput) Origins() DeepCreatedOriginResponseArrayOutput {
	return o.ApplyT(func(v *Endpoint) DeepCreatedOriginResponseArrayOutput { return v.Origins }).(DeepCreatedOriginResponseArrayOutput)
}

// Path to a file hosted on the origin which helps accelerate delivery of the dynamic content and calculate the most optimal routes for the CDN. This is relative to the origin path. This property is only relevant when using a single origin.
func (o EndpointOutput) ProbePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringPtrOutput { return v.ProbePath }).(pulumi.StringPtrOutput)
}

// Provisioning status of the endpoint.
func (o EndpointOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Defines how CDN caches requests that include query strings. You can ignore any query strings when caching, bypass caching to prevent requests that contain query strings from being cached, or cache every request with a unique URL.
func (o EndpointOutput) QueryStringCachingBehavior() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringPtrOutput { return v.QueryStringCachingBehavior }).(pulumi.StringPtrOutput)
}

// Resource status of the endpoint.
func (o EndpointOutput) ResourceState() pulumi.StringOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringOutput { return v.ResourceState }).(pulumi.StringOutput)
}

// Read only system data
func (o EndpointOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Endpoint) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o EndpointOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o EndpointOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// List of keys used to validate the signed URL hashes.
func (o EndpointOutput) UrlSigningKeys() UrlSigningKeyResponseArrayOutput {
	return o.ApplyT(func(v *Endpoint) UrlSigningKeyResponseArrayOutput { return v.UrlSigningKeys }).(UrlSigningKeyResponseArrayOutput)
}

// Defines the Web Application Firewall policy for the endpoint (if applicable)
func (o EndpointOutput) WebApplicationFirewallPolicyLink() EndpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLinkPtrOutput {
	return o.ApplyT(func(v *Endpoint) EndpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLinkPtrOutput {
		return v.WebApplicationFirewallPolicyLink
	}).(EndpointPropertiesUpdateParametersResponseWebApplicationFirewallPolicyLinkPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(EndpointOutput{})
}
