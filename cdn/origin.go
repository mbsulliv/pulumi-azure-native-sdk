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

// CDN origin is the source of the content being delivered via CDN. When the edge nodes represented by an endpoint do not have the requested content cached, they attempt to fetch it from one or more of the configured origins.
// Azure REST API version: 2023-05-01. Prior API version in Azure Native 1.x: 2020-09-01
type Origin struct {
	pulumi.CustomResourceState

	// Origin is enabled for load balancing or not
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// The address of the origin. Domain names, IPv4 addresses, and IPv6 addresses are supported.This should be unique across all origins in an endpoint.
	HostName pulumi.StringOutput `pulumi:"hostName"`
	// The value of the HTTP port. Must be between 1 and 65535.
	HttpPort pulumi.IntPtrOutput `pulumi:"httpPort"`
	// The value of the HTTPS port. Must be between 1 and 65535.
	HttpsPort pulumi.IntPtrOutput `pulumi:"httpsPort"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The host header value sent to the origin with each request. If you leave this blank, the request hostname determines this value. Azure CDN origins, such as Web Apps, Blob Storage, and Cloud Services require this host header value to match the origin hostname by default. This overrides the host header defined at Endpoint
	OriginHostHeader pulumi.StringPtrOutput `pulumi:"originHostHeader"`
	// Priority of origin in given origin group for load balancing. Higher priorities will not be used for load balancing if any lower priority origin is healthy.Must be between 1 and 5
	Priority pulumi.IntPtrOutput `pulumi:"priority"`
	// The approval status for the connection to the Private Link
	PrivateEndpointStatus pulumi.StringOutput `pulumi:"privateEndpointStatus"`
	// The Alias of the Private Link resource. Populating this optional field indicates that this origin is 'Private'
	PrivateLinkAlias pulumi.StringPtrOutput `pulumi:"privateLinkAlias"`
	// A custom message to be included in the approval request to connect to the Private Link.
	PrivateLinkApprovalMessage pulumi.StringPtrOutput `pulumi:"privateLinkApprovalMessage"`
	// The location of the Private Link resource. Required only if 'privateLinkResourceId' is populated
	PrivateLinkLocation pulumi.StringPtrOutput `pulumi:"privateLinkLocation"`
	// The Resource Id of the Private Link resource. Populating this optional field indicates that this backend is 'Private'
	PrivateLinkResourceId pulumi.StringPtrOutput `pulumi:"privateLinkResourceId"`
	// Provisioning status of the origin.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Resource status of the origin.
	ResourceState pulumi.StringOutput `pulumi:"resourceState"`
	// Read only system data
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// Weight of the origin in given origin group for load balancing. Must be between 1 and 1000
	Weight pulumi.IntPtrOutput `pulumi:"weight"`
}

// NewOrigin registers a new resource with the given unique name, arguments, and options.
func NewOrigin(ctx *pulumi.Context,
	name string, args *OriginArgs, opts ...pulumi.ResourceOption) (*Origin, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EndpointName == nil {
		return nil, errors.New("invalid value for required argument 'EndpointName'")
	}
	if args.HostName == nil {
		return nil, errors.New("invalid value for required argument 'HostName'")
	}
	if args.ProfileName == nil {
		return nil, errors.New("invalid value for required argument 'ProfileName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:cdn/v20150601:Origin"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20160402:Origin"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20191231:Origin"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20200331:Origin"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20200415:Origin"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20200901:Origin"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20210601:Origin"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20220501preview:Origin"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20221101preview:Origin"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20230501:Origin"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20230701preview:Origin"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Origin
	err := ctx.RegisterResource("azure-native:cdn:Origin", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOrigin gets an existing Origin resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOrigin(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OriginState, opts ...pulumi.ResourceOption) (*Origin, error) {
	var resource Origin
	err := ctx.ReadResource("azure-native:cdn:Origin", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Origin resources.
type originState struct {
}

type OriginState struct {
}

func (OriginState) ElementType() reflect.Type {
	return reflect.TypeOf((*originState)(nil)).Elem()
}

type originArgs struct {
	// Origin is enabled for load balancing or not
	Enabled *bool `pulumi:"enabled"`
	// Name of the endpoint under the profile which is unique globally.
	EndpointName string `pulumi:"endpointName"`
	// The address of the origin. Domain names, IPv4 addresses, and IPv6 addresses are supported.This should be unique across all origins in an endpoint.
	HostName string `pulumi:"hostName"`
	// The value of the HTTP port. Must be between 1 and 65535.
	HttpPort *int `pulumi:"httpPort"`
	// The value of the HTTPS port. Must be between 1 and 65535.
	HttpsPort *int `pulumi:"httpsPort"`
	// The host header value sent to the origin with each request. If you leave this blank, the request hostname determines this value. Azure CDN origins, such as Web Apps, Blob Storage, and Cloud Services require this host header value to match the origin hostname by default. This overrides the host header defined at Endpoint
	OriginHostHeader *string `pulumi:"originHostHeader"`
	// Name of the origin that is unique within the endpoint.
	OriginName *string `pulumi:"originName"`
	// Priority of origin in given origin group for load balancing. Higher priorities will not be used for load balancing if any lower priority origin is healthy.Must be between 1 and 5
	Priority *int `pulumi:"priority"`
	// The Alias of the Private Link resource. Populating this optional field indicates that this origin is 'Private'
	PrivateLinkAlias *string `pulumi:"privateLinkAlias"`
	// A custom message to be included in the approval request to connect to the Private Link.
	PrivateLinkApprovalMessage *string `pulumi:"privateLinkApprovalMessage"`
	// The location of the Private Link resource. Required only if 'privateLinkResourceId' is populated
	PrivateLinkLocation *string `pulumi:"privateLinkLocation"`
	// The Resource Id of the Private Link resource. Populating this optional field indicates that this backend is 'Private'
	PrivateLinkResourceId *string `pulumi:"privateLinkResourceId"`
	// Name of the CDN profile which is unique within the resource group.
	ProfileName string `pulumi:"profileName"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Weight of the origin in given origin group for load balancing. Must be between 1 and 1000
	Weight *int `pulumi:"weight"`
}

// The set of arguments for constructing a Origin resource.
type OriginArgs struct {
	// Origin is enabled for load balancing or not
	Enabled pulumi.BoolPtrInput
	// Name of the endpoint under the profile which is unique globally.
	EndpointName pulumi.StringInput
	// The address of the origin. Domain names, IPv4 addresses, and IPv6 addresses are supported.This should be unique across all origins in an endpoint.
	HostName pulumi.StringInput
	// The value of the HTTP port. Must be between 1 and 65535.
	HttpPort pulumi.IntPtrInput
	// The value of the HTTPS port. Must be between 1 and 65535.
	HttpsPort pulumi.IntPtrInput
	// The host header value sent to the origin with each request. If you leave this blank, the request hostname determines this value. Azure CDN origins, such as Web Apps, Blob Storage, and Cloud Services require this host header value to match the origin hostname by default. This overrides the host header defined at Endpoint
	OriginHostHeader pulumi.StringPtrInput
	// Name of the origin that is unique within the endpoint.
	OriginName pulumi.StringPtrInput
	// Priority of origin in given origin group for load balancing. Higher priorities will not be used for load balancing if any lower priority origin is healthy.Must be between 1 and 5
	Priority pulumi.IntPtrInput
	// The Alias of the Private Link resource. Populating this optional field indicates that this origin is 'Private'
	PrivateLinkAlias pulumi.StringPtrInput
	// A custom message to be included in the approval request to connect to the Private Link.
	PrivateLinkApprovalMessage pulumi.StringPtrInput
	// The location of the Private Link resource. Required only if 'privateLinkResourceId' is populated
	PrivateLinkLocation pulumi.StringPtrInput
	// The Resource Id of the Private Link resource. Populating this optional field indicates that this backend is 'Private'
	PrivateLinkResourceId pulumi.StringPtrInput
	// Name of the CDN profile which is unique within the resource group.
	ProfileName pulumi.StringInput
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput
	// Weight of the origin in given origin group for load balancing. Must be between 1 and 1000
	Weight pulumi.IntPtrInput
}

func (OriginArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*originArgs)(nil)).Elem()
}

type OriginInput interface {
	pulumi.Input

	ToOriginOutput() OriginOutput
	ToOriginOutputWithContext(ctx context.Context) OriginOutput
}

func (*Origin) ElementType() reflect.Type {
	return reflect.TypeOf((**Origin)(nil)).Elem()
}

func (i *Origin) ToOriginOutput() OriginOutput {
	return i.ToOriginOutputWithContext(context.Background())
}

func (i *Origin) ToOriginOutputWithContext(ctx context.Context) OriginOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OriginOutput)
}

func (i *Origin) ToOutput(ctx context.Context) pulumix.Output[*Origin] {
	return pulumix.Output[*Origin]{
		OutputState: i.ToOriginOutputWithContext(ctx).OutputState,
	}
}

type OriginOutput struct{ *pulumi.OutputState }

func (OriginOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Origin)(nil)).Elem()
}

func (o OriginOutput) ToOriginOutput() OriginOutput {
	return o
}

func (o OriginOutput) ToOriginOutputWithContext(ctx context.Context) OriginOutput {
	return o
}

func (o OriginOutput) ToOutput(ctx context.Context) pulumix.Output[*Origin] {
	return pulumix.Output[*Origin]{
		OutputState: o.OutputState,
	}
}

// Origin is enabled for load balancing or not
func (o OriginOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Origin) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// The address of the origin. Domain names, IPv4 addresses, and IPv6 addresses are supported.This should be unique across all origins in an endpoint.
func (o OriginOutput) HostName() pulumi.StringOutput {
	return o.ApplyT(func(v *Origin) pulumi.StringOutput { return v.HostName }).(pulumi.StringOutput)
}

// The value of the HTTP port. Must be between 1 and 65535.
func (o OriginOutput) HttpPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Origin) pulumi.IntPtrOutput { return v.HttpPort }).(pulumi.IntPtrOutput)
}

// The value of the HTTPS port. Must be between 1 and 65535.
func (o OriginOutput) HttpsPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Origin) pulumi.IntPtrOutput { return v.HttpsPort }).(pulumi.IntPtrOutput)
}

// Resource name.
func (o OriginOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Origin) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The host header value sent to the origin with each request. If you leave this blank, the request hostname determines this value. Azure CDN origins, such as Web Apps, Blob Storage, and Cloud Services require this host header value to match the origin hostname by default. This overrides the host header defined at Endpoint
func (o OriginOutput) OriginHostHeader() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Origin) pulumi.StringPtrOutput { return v.OriginHostHeader }).(pulumi.StringPtrOutput)
}

// Priority of origin in given origin group for load balancing. Higher priorities will not be used for load balancing if any lower priority origin is healthy.Must be between 1 and 5
func (o OriginOutput) Priority() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Origin) pulumi.IntPtrOutput { return v.Priority }).(pulumi.IntPtrOutput)
}

// The approval status for the connection to the Private Link
func (o OriginOutput) PrivateEndpointStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *Origin) pulumi.StringOutput { return v.PrivateEndpointStatus }).(pulumi.StringOutput)
}

// The Alias of the Private Link resource. Populating this optional field indicates that this origin is 'Private'
func (o OriginOutput) PrivateLinkAlias() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Origin) pulumi.StringPtrOutput { return v.PrivateLinkAlias }).(pulumi.StringPtrOutput)
}

// A custom message to be included in the approval request to connect to the Private Link.
func (o OriginOutput) PrivateLinkApprovalMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Origin) pulumi.StringPtrOutput { return v.PrivateLinkApprovalMessage }).(pulumi.StringPtrOutput)
}

// The location of the Private Link resource. Required only if 'privateLinkResourceId' is populated
func (o OriginOutput) PrivateLinkLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Origin) pulumi.StringPtrOutput { return v.PrivateLinkLocation }).(pulumi.StringPtrOutput)
}

// The Resource Id of the Private Link resource. Populating this optional field indicates that this backend is 'Private'
func (o OriginOutput) PrivateLinkResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Origin) pulumi.StringPtrOutput { return v.PrivateLinkResourceId }).(pulumi.StringPtrOutput)
}

// Provisioning status of the origin.
func (o OriginOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Origin) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Resource status of the origin.
func (o OriginOutput) ResourceState() pulumi.StringOutput {
	return o.ApplyT(func(v *Origin) pulumi.StringOutput { return v.ResourceState }).(pulumi.StringOutput)
}

// Read only system data
func (o OriginOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Origin) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource type.
func (o OriginOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Origin) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Weight of the origin in given origin group for load balancing. Must be between 1 and 1000
func (o OriginOutput) Weight() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Origin) pulumi.IntPtrOutput { return v.Weight }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(OriginOutput{})
}
