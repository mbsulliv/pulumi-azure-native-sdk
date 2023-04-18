// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20221101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Azure Front Door origin is the source of the content being delivered via Azure Front Door. When the edge nodes represented by an endpoint do not have the requested content cached, they attempt to fetch it from one or more of the configured origins.
type AFDOrigin struct {
	pulumi.CustomResourceState

	// Resource reference to the Azure origin resource.
	AzureOrigin      ResourceReferenceResponsePtrOutput `pulumi:"azureOrigin"`
	DeploymentStatus pulumi.StringOutput                `pulumi:"deploymentStatus"`
	// Whether to enable health probes to be made against backends defined under backendPools. Health probes can only be disabled if there is a single enabled backend in single enabled backend pool.
	EnabledState pulumi.StringPtrOutput `pulumi:"enabledState"`
	// Whether to enable certificate name check at origin level
	EnforceCertificateNameCheck pulumi.BoolPtrOutput `pulumi:"enforceCertificateNameCheck"`
	// The address of the origin. Domain names, IPv4 addresses, and IPv6 addresses are supported.This should be unique across all origins in an endpoint.
	HostName pulumi.StringOutput `pulumi:"hostName"`
	// The value of the HTTP port. Must be between 1 and 65535.
	HttpPort pulumi.IntPtrOutput `pulumi:"httpPort"`
	// The value of the HTTPS port. Must be between 1 and 65535.
	HttpsPort pulumi.IntPtrOutput `pulumi:"httpsPort"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the origin group which contains this origin.
	OriginGroupName pulumi.StringOutput `pulumi:"originGroupName"`
	// The host header value sent to the origin with each request. If you leave this blank, the request hostname determines this value. Azure Front Door origins, such as Web Apps, Blob Storage, and Cloud Services require this host header value to match the origin hostname by default. This overrides the host header defined at Endpoint
	OriginHostHeader pulumi.StringPtrOutput `pulumi:"originHostHeader"`
	// Priority of origin in given origin group for load balancing. Higher priorities will not be used for load balancing if any lower priority origin is healthy.Must be between 1 and 5
	Priority pulumi.IntPtrOutput `pulumi:"priority"`
	// Provisioning status
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The properties of the private link resource for private origin.
	SharedPrivateLinkResource SharedPrivateLinkResourcePropertiesResponsePtrOutput `pulumi:"sharedPrivateLinkResource"`
	// Read only system data
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// Weight of the origin in given origin group for load balancing. Must be between 1 and 1000
	Weight pulumi.IntPtrOutput `pulumi:"weight"`
}

// NewAFDOrigin registers a new resource with the given unique name, arguments, and options.
func NewAFDOrigin(ctx *pulumi.Context,
	name string, args *AFDOriginArgs, opts ...pulumi.ResourceOption) (*AFDOrigin, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.HostName == nil {
		return nil, errors.New("invalid value for required argument 'HostName'")
	}
	if args.OriginGroupName == nil {
		return nil, errors.New("invalid value for required argument 'OriginGroupName'")
	}
	if args.ProfileName == nil {
		return nil, errors.New("invalid value for required argument 'ProfileName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if isZero(args.EnforceCertificateNameCheck) {
		args.EnforceCertificateNameCheck = pulumi.BoolPtr(true)
	}
	if isZero(args.HttpPort) {
		args.HttpPort = pulumi.IntPtr(80)
	}
	if isZero(args.HttpsPort) {
		args.HttpsPort = pulumi.IntPtr(443)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:cdn:AFDOrigin"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20200901:AFDOrigin"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20210601:AFDOrigin"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20220501preview:AFDOrigin"),
		},
	})
	opts = append(opts, aliases)
	var resource AFDOrigin
	err := ctx.RegisterResource("azure-native:cdn/v20221101preview:AFDOrigin", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAFDOrigin gets an existing AFDOrigin resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAFDOrigin(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AFDOriginState, opts ...pulumi.ResourceOption) (*AFDOrigin, error) {
	var resource AFDOrigin
	err := ctx.ReadResource("azure-native:cdn/v20221101preview:AFDOrigin", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AFDOrigin resources.
type afdoriginState struct {
}

type AFDOriginState struct {
}

func (AFDOriginState) ElementType() reflect.Type {
	return reflect.TypeOf((*afdoriginState)(nil)).Elem()
}

type afdoriginArgs struct {
	// Resource reference to the Azure origin resource.
	AzureOrigin *ResourceReference `pulumi:"azureOrigin"`
	// Whether to enable health probes to be made against backends defined under backendPools. Health probes can only be disabled if there is a single enabled backend in single enabled backend pool.
	EnabledState *string `pulumi:"enabledState"`
	// Whether to enable certificate name check at origin level
	EnforceCertificateNameCheck *bool `pulumi:"enforceCertificateNameCheck"`
	// The address of the origin. Domain names, IPv4 addresses, and IPv6 addresses are supported.This should be unique across all origins in an endpoint.
	HostName string `pulumi:"hostName"`
	// The value of the HTTP port. Must be between 1 and 65535.
	HttpPort *int `pulumi:"httpPort"`
	// The value of the HTTPS port. Must be between 1 and 65535.
	HttpsPort *int `pulumi:"httpsPort"`
	// Name of the origin group which is unique within the profile.
	OriginGroupName string `pulumi:"originGroupName"`
	// The host header value sent to the origin with each request. If you leave this blank, the request hostname determines this value. Azure Front Door origins, such as Web Apps, Blob Storage, and Cloud Services require this host header value to match the origin hostname by default. This overrides the host header defined at Endpoint
	OriginHostHeader *string `pulumi:"originHostHeader"`
	// Name of the origin that is unique within the profile.
	OriginName *string `pulumi:"originName"`
	// Priority of origin in given origin group for load balancing. Higher priorities will not be used for load balancing if any lower priority origin is healthy.Must be between 1 and 5
	Priority *int `pulumi:"priority"`
	// Name of the Azure Front Door Standard or Azure Front Door Premium profile which is unique within the resource group.
	ProfileName string `pulumi:"profileName"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The properties of the private link resource for private origin.
	SharedPrivateLinkResource *SharedPrivateLinkResourceProperties `pulumi:"sharedPrivateLinkResource"`
	// Weight of the origin in given origin group for load balancing. Must be between 1 and 1000
	Weight *int `pulumi:"weight"`
}

// The set of arguments for constructing a AFDOrigin resource.
type AFDOriginArgs struct {
	// Resource reference to the Azure origin resource.
	AzureOrigin ResourceReferencePtrInput
	// Whether to enable health probes to be made against backends defined under backendPools. Health probes can only be disabled if there is a single enabled backend in single enabled backend pool.
	EnabledState pulumi.StringPtrInput
	// Whether to enable certificate name check at origin level
	EnforceCertificateNameCheck pulumi.BoolPtrInput
	// The address of the origin. Domain names, IPv4 addresses, and IPv6 addresses are supported.This should be unique across all origins in an endpoint.
	HostName pulumi.StringInput
	// The value of the HTTP port. Must be between 1 and 65535.
	HttpPort pulumi.IntPtrInput
	// The value of the HTTPS port. Must be between 1 and 65535.
	HttpsPort pulumi.IntPtrInput
	// Name of the origin group which is unique within the profile.
	OriginGroupName pulumi.StringInput
	// The host header value sent to the origin with each request. If you leave this blank, the request hostname determines this value. Azure Front Door origins, such as Web Apps, Blob Storage, and Cloud Services require this host header value to match the origin hostname by default. This overrides the host header defined at Endpoint
	OriginHostHeader pulumi.StringPtrInput
	// Name of the origin that is unique within the profile.
	OriginName pulumi.StringPtrInput
	// Priority of origin in given origin group for load balancing. Higher priorities will not be used for load balancing if any lower priority origin is healthy.Must be between 1 and 5
	Priority pulumi.IntPtrInput
	// Name of the Azure Front Door Standard or Azure Front Door Premium profile which is unique within the resource group.
	ProfileName pulumi.StringInput
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput
	// The properties of the private link resource for private origin.
	SharedPrivateLinkResource SharedPrivateLinkResourcePropertiesPtrInput
	// Weight of the origin in given origin group for load balancing. Must be between 1 and 1000
	Weight pulumi.IntPtrInput
}

func (AFDOriginArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*afdoriginArgs)(nil)).Elem()
}

type AFDOriginInput interface {
	pulumi.Input

	ToAFDOriginOutput() AFDOriginOutput
	ToAFDOriginOutputWithContext(ctx context.Context) AFDOriginOutput
}

func (*AFDOrigin) ElementType() reflect.Type {
	return reflect.TypeOf((**AFDOrigin)(nil)).Elem()
}

func (i *AFDOrigin) ToAFDOriginOutput() AFDOriginOutput {
	return i.ToAFDOriginOutputWithContext(context.Background())
}

func (i *AFDOrigin) ToAFDOriginOutputWithContext(ctx context.Context) AFDOriginOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AFDOriginOutput)
}

type AFDOriginOutput struct{ *pulumi.OutputState }

func (AFDOriginOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AFDOrigin)(nil)).Elem()
}

func (o AFDOriginOutput) ToAFDOriginOutput() AFDOriginOutput {
	return o
}

func (o AFDOriginOutput) ToAFDOriginOutputWithContext(ctx context.Context) AFDOriginOutput {
	return o
}

// Resource reference to the Azure origin resource.
func (o AFDOriginOutput) AzureOrigin() ResourceReferenceResponsePtrOutput {
	return o.ApplyT(func(v *AFDOrigin) ResourceReferenceResponsePtrOutput { return v.AzureOrigin }).(ResourceReferenceResponsePtrOutput)
}

func (o AFDOriginOutput) DeploymentStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *AFDOrigin) pulumi.StringOutput { return v.DeploymentStatus }).(pulumi.StringOutput)
}

// Whether to enable health probes to be made against backends defined under backendPools. Health probes can only be disabled if there is a single enabled backend in single enabled backend pool.
func (o AFDOriginOutput) EnabledState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AFDOrigin) pulumi.StringPtrOutput { return v.EnabledState }).(pulumi.StringPtrOutput)
}

// Whether to enable certificate name check at origin level
func (o AFDOriginOutput) EnforceCertificateNameCheck() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AFDOrigin) pulumi.BoolPtrOutput { return v.EnforceCertificateNameCheck }).(pulumi.BoolPtrOutput)
}

// The address of the origin. Domain names, IPv4 addresses, and IPv6 addresses are supported.This should be unique across all origins in an endpoint.
func (o AFDOriginOutput) HostName() pulumi.StringOutput {
	return o.ApplyT(func(v *AFDOrigin) pulumi.StringOutput { return v.HostName }).(pulumi.StringOutput)
}

// The value of the HTTP port. Must be between 1 and 65535.
func (o AFDOriginOutput) HttpPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AFDOrigin) pulumi.IntPtrOutput { return v.HttpPort }).(pulumi.IntPtrOutput)
}

// The value of the HTTPS port. Must be between 1 and 65535.
func (o AFDOriginOutput) HttpsPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AFDOrigin) pulumi.IntPtrOutput { return v.HttpsPort }).(pulumi.IntPtrOutput)
}

// Resource name.
func (o AFDOriginOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AFDOrigin) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The name of the origin group which contains this origin.
func (o AFDOriginOutput) OriginGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v *AFDOrigin) pulumi.StringOutput { return v.OriginGroupName }).(pulumi.StringOutput)
}

// The host header value sent to the origin with each request. If you leave this blank, the request hostname determines this value. Azure Front Door origins, such as Web Apps, Blob Storage, and Cloud Services require this host header value to match the origin hostname by default. This overrides the host header defined at Endpoint
func (o AFDOriginOutput) OriginHostHeader() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AFDOrigin) pulumi.StringPtrOutput { return v.OriginHostHeader }).(pulumi.StringPtrOutput)
}

// Priority of origin in given origin group for load balancing. Higher priorities will not be used for load balancing if any lower priority origin is healthy.Must be between 1 and 5
func (o AFDOriginOutput) Priority() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AFDOrigin) pulumi.IntPtrOutput { return v.Priority }).(pulumi.IntPtrOutput)
}

// Provisioning status
func (o AFDOriginOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *AFDOrigin) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The properties of the private link resource for private origin.
func (o AFDOriginOutput) SharedPrivateLinkResource() SharedPrivateLinkResourcePropertiesResponsePtrOutput {
	return o.ApplyT(func(v *AFDOrigin) SharedPrivateLinkResourcePropertiesResponsePtrOutput {
		return v.SharedPrivateLinkResource
	}).(SharedPrivateLinkResourcePropertiesResponsePtrOutput)
}

// Read only system data
func (o AFDOriginOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *AFDOrigin) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource type.
func (o AFDOriginOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AFDOrigin) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Weight of the origin in given origin group for load balancing. Must be between 1 and 1000
func (o AFDOriginOutput) Weight() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AFDOrigin) pulumi.IntPtrOutput { return v.Weight }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AFDOriginOutput{})
}
