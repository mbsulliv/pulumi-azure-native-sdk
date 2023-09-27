// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cdn

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Friendly Routes name mapping to the any Routes or secret related information.
// Azure REST API version: 2023-05-01. Prior API version in Azure Native 1.x: 2020-09-01
type Route struct {
	pulumi.CustomResourceState

	// The caching configuration for this route. To disable caching, do not provide a cacheConfiguration object.
	CacheConfiguration AfdRouteCacheConfigurationResponsePtrOutput `pulumi:"cacheConfiguration"`
	// Domains referenced by this endpoint.
	CustomDomains    ActivatedResourceReferenceResponseArrayOutput `pulumi:"customDomains"`
	DeploymentStatus pulumi.StringOutput                           `pulumi:"deploymentStatus"`
	// Whether to enable use of this rule. Permitted values are 'Enabled' or 'Disabled'
	EnabledState pulumi.StringPtrOutput `pulumi:"enabledState"`
	// The name of the endpoint which holds the route.
	EndpointName pulumi.StringOutput `pulumi:"endpointName"`
	// Protocol this rule will use when forwarding traffic to backends.
	ForwardingProtocol pulumi.StringPtrOutput `pulumi:"forwardingProtocol"`
	// Whether to automatically redirect HTTP traffic to HTTPS traffic. Note that this is a easy way to set up this rule and it will be the first rule that gets executed.
	HttpsRedirect pulumi.StringPtrOutput `pulumi:"httpsRedirect"`
	// whether this route will be linked to the default endpoint domain.
	LinkToDefaultDomain pulumi.StringPtrOutput `pulumi:"linkToDefaultDomain"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// A reference to the origin group.
	OriginGroup ResourceReferenceResponseOutput `pulumi:"originGroup"`
	// A directory path on the origin that AzureFrontDoor can use to retrieve content from, e.g. contoso.cloudapp.net/originpath.
	OriginPath pulumi.StringPtrOutput `pulumi:"originPath"`
	// The route patterns of the rule.
	PatternsToMatch pulumi.StringArrayOutput `pulumi:"patternsToMatch"`
	// Provisioning status
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// rule sets referenced by this endpoint.
	RuleSets ResourceReferenceResponseArrayOutput `pulumi:"ruleSets"`
	// List of supported protocols for this route.
	SupportedProtocols pulumi.StringArrayOutput `pulumi:"supportedProtocols"`
	// Read only system data
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewRoute registers a new resource with the given unique name, arguments, and options.
func NewRoute(ctx *pulumi.Context,
	name string, args *RouteArgs, opts ...pulumi.ResourceOption) (*Route, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EndpointName == nil {
		return nil, errors.New("invalid value for required argument 'EndpointName'")
	}
	if args.OriginGroup == nil {
		return nil, errors.New("invalid value for required argument 'OriginGroup'")
	}
	if args.ProfileName == nil {
		return nil, errors.New("invalid value for required argument 'ProfileName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ForwardingProtocol == nil {
		args.ForwardingProtocol = pulumi.StringPtr("MatchRequest")
	}
	if args.HttpsRedirect == nil {
		args.HttpsRedirect = pulumi.StringPtr("Disabled")
	}
	if args.LinkToDefaultDomain == nil {
		args.LinkToDefaultDomain = pulumi.StringPtr("Disabled")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:cdn/v20200901:Route"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20210601:Route"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20220501preview:Route"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20221101preview:Route"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20230501:Route"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20230701preview:Route"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Route
	err := ctx.RegisterResource("azure-native:cdn:Route", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRoute gets an existing Route resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRoute(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RouteState, opts ...pulumi.ResourceOption) (*Route, error) {
	var resource Route
	err := ctx.ReadResource("azure-native:cdn:Route", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Route resources.
type routeState struct {
}

type RouteState struct {
}

func (RouteState) ElementType() reflect.Type {
	return reflect.TypeOf((*routeState)(nil)).Elem()
}

type routeArgs struct {
	// The caching configuration for this route. To disable caching, do not provide a cacheConfiguration object.
	CacheConfiguration *AfdRouteCacheConfiguration `pulumi:"cacheConfiguration"`
	// Domains referenced by this endpoint.
	CustomDomains []ActivatedResourceReference `pulumi:"customDomains"`
	// Whether to enable use of this rule. Permitted values are 'Enabled' or 'Disabled'
	EnabledState *string `pulumi:"enabledState"`
	// Name of the endpoint under the profile which is unique globally.
	EndpointName string `pulumi:"endpointName"`
	// Protocol this rule will use when forwarding traffic to backends.
	ForwardingProtocol *string `pulumi:"forwardingProtocol"`
	// Whether to automatically redirect HTTP traffic to HTTPS traffic. Note that this is a easy way to set up this rule and it will be the first rule that gets executed.
	HttpsRedirect *string `pulumi:"httpsRedirect"`
	// whether this route will be linked to the default endpoint domain.
	LinkToDefaultDomain *string `pulumi:"linkToDefaultDomain"`
	// A reference to the origin group.
	OriginGroup ResourceReference `pulumi:"originGroup"`
	// A directory path on the origin that AzureFrontDoor can use to retrieve content from, e.g. contoso.cloudapp.net/originpath.
	OriginPath *string `pulumi:"originPath"`
	// The route patterns of the rule.
	PatternsToMatch []string `pulumi:"patternsToMatch"`
	// Name of the Azure Front Door Standard or Azure Front Door Premium profile which is unique within the resource group.
	ProfileName string `pulumi:"profileName"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the routing rule.
	RouteName *string `pulumi:"routeName"`
	// rule sets referenced by this endpoint.
	RuleSets []ResourceReference `pulumi:"ruleSets"`
	// List of supported protocols for this route.
	SupportedProtocols []string `pulumi:"supportedProtocols"`
}

// The set of arguments for constructing a Route resource.
type RouteArgs struct {
	// The caching configuration for this route. To disable caching, do not provide a cacheConfiguration object.
	CacheConfiguration AfdRouteCacheConfigurationPtrInput
	// Domains referenced by this endpoint.
	CustomDomains ActivatedResourceReferenceArrayInput
	// Whether to enable use of this rule. Permitted values are 'Enabled' or 'Disabled'
	EnabledState pulumi.StringPtrInput
	// Name of the endpoint under the profile which is unique globally.
	EndpointName pulumi.StringInput
	// Protocol this rule will use when forwarding traffic to backends.
	ForwardingProtocol pulumi.StringPtrInput
	// Whether to automatically redirect HTTP traffic to HTTPS traffic. Note that this is a easy way to set up this rule and it will be the first rule that gets executed.
	HttpsRedirect pulumi.StringPtrInput
	// whether this route will be linked to the default endpoint domain.
	LinkToDefaultDomain pulumi.StringPtrInput
	// A reference to the origin group.
	OriginGroup ResourceReferenceInput
	// A directory path on the origin that AzureFrontDoor can use to retrieve content from, e.g. contoso.cloudapp.net/originpath.
	OriginPath pulumi.StringPtrInput
	// The route patterns of the rule.
	PatternsToMatch pulumi.StringArrayInput
	// Name of the Azure Front Door Standard or Azure Front Door Premium profile which is unique within the resource group.
	ProfileName pulumi.StringInput
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput
	// Name of the routing rule.
	RouteName pulumi.StringPtrInput
	// rule sets referenced by this endpoint.
	RuleSets ResourceReferenceArrayInput
	// List of supported protocols for this route.
	SupportedProtocols pulumi.StringArrayInput
}

func (RouteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*routeArgs)(nil)).Elem()
}

type RouteInput interface {
	pulumi.Input

	ToRouteOutput() RouteOutput
	ToRouteOutputWithContext(ctx context.Context) RouteOutput
}

func (*Route) ElementType() reflect.Type {
	return reflect.TypeOf((**Route)(nil)).Elem()
}

func (i *Route) ToRouteOutput() RouteOutput {
	return i.ToRouteOutputWithContext(context.Background())
}

func (i *Route) ToRouteOutputWithContext(ctx context.Context) RouteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteOutput)
}

func (i *Route) ToOutput(ctx context.Context) pulumix.Output[*Route] {
	return pulumix.Output[*Route]{
		OutputState: i.ToRouteOutputWithContext(ctx).OutputState,
	}
}

type RouteOutput struct{ *pulumi.OutputState }

func (RouteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Route)(nil)).Elem()
}

func (o RouteOutput) ToRouteOutput() RouteOutput {
	return o
}

func (o RouteOutput) ToRouteOutputWithContext(ctx context.Context) RouteOutput {
	return o
}

func (o RouteOutput) ToOutput(ctx context.Context) pulumix.Output[*Route] {
	return pulumix.Output[*Route]{
		OutputState: o.OutputState,
	}
}

// The caching configuration for this route. To disable caching, do not provide a cacheConfiguration object.
func (o RouteOutput) CacheConfiguration() AfdRouteCacheConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *Route) AfdRouteCacheConfigurationResponsePtrOutput { return v.CacheConfiguration }).(AfdRouteCacheConfigurationResponsePtrOutput)
}

// Domains referenced by this endpoint.
func (o RouteOutput) CustomDomains() ActivatedResourceReferenceResponseArrayOutput {
	return o.ApplyT(func(v *Route) ActivatedResourceReferenceResponseArrayOutput { return v.CustomDomains }).(ActivatedResourceReferenceResponseArrayOutput)
}

func (o RouteOutput) DeploymentStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *Route) pulumi.StringOutput { return v.DeploymentStatus }).(pulumi.StringOutput)
}

// Whether to enable use of this rule. Permitted values are 'Enabled' or 'Disabled'
func (o RouteOutput) EnabledState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Route) pulumi.StringPtrOutput { return v.EnabledState }).(pulumi.StringPtrOutput)
}

// The name of the endpoint which holds the route.
func (o RouteOutput) EndpointName() pulumi.StringOutput {
	return o.ApplyT(func(v *Route) pulumi.StringOutput { return v.EndpointName }).(pulumi.StringOutput)
}

// Protocol this rule will use when forwarding traffic to backends.
func (o RouteOutput) ForwardingProtocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Route) pulumi.StringPtrOutput { return v.ForwardingProtocol }).(pulumi.StringPtrOutput)
}

// Whether to automatically redirect HTTP traffic to HTTPS traffic. Note that this is a easy way to set up this rule and it will be the first rule that gets executed.
func (o RouteOutput) HttpsRedirect() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Route) pulumi.StringPtrOutput { return v.HttpsRedirect }).(pulumi.StringPtrOutput)
}

// whether this route will be linked to the default endpoint domain.
func (o RouteOutput) LinkToDefaultDomain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Route) pulumi.StringPtrOutput { return v.LinkToDefaultDomain }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o RouteOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Route) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// A reference to the origin group.
func (o RouteOutput) OriginGroup() ResourceReferenceResponseOutput {
	return o.ApplyT(func(v *Route) ResourceReferenceResponseOutput { return v.OriginGroup }).(ResourceReferenceResponseOutput)
}

// A directory path on the origin that AzureFrontDoor can use to retrieve content from, e.g. contoso.cloudapp.net/originpath.
func (o RouteOutput) OriginPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Route) pulumi.StringPtrOutput { return v.OriginPath }).(pulumi.StringPtrOutput)
}

// The route patterns of the rule.
func (o RouteOutput) PatternsToMatch() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Route) pulumi.StringArrayOutput { return v.PatternsToMatch }).(pulumi.StringArrayOutput)
}

// Provisioning status
func (o RouteOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Route) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// rule sets referenced by this endpoint.
func (o RouteOutput) RuleSets() ResourceReferenceResponseArrayOutput {
	return o.ApplyT(func(v *Route) ResourceReferenceResponseArrayOutput { return v.RuleSets }).(ResourceReferenceResponseArrayOutput)
}

// List of supported protocols for this route.
func (o RouteOutput) SupportedProtocols() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Route) pulumi.StringArrayOutput { return v.SupportedProtocols }).(pulumi.StringArrayOutput)
}

// Read only system data
func (o RouteOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Route) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource type.
func (o RouteOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Route) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(RouteOutput{})
}
