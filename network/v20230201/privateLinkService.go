// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230201

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Private link service resource.
type PrivateLinkService struct {
	pulumi.CustomResourceState

	// The alias of the private link service.
	Alias pulumi.StringOutput `pulumi:"alias"`
	// The auto-approval list of the private link service.
	AutoApproval PrivateLinkServicePropertiesResponseAutoApprovalPtrOutput `pulumi:"autoApproval"`
	// Whether the private link service is enabled for proxy protocol or not.
	EnableProxyProtocol pulumi.BoolPtrOutput `pulumi:"enableProxyProtocol"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The extended location of the load balancer.
	ExtendedLocation ExtendedLocationResponsePtrOutput `pulumi:"extendedLocation"`
	// The list of Fqdn.
	Fqdns pulumi.StringArrayOutput `pulumi:"fqdns"`
	// An array of private link service IP configurations.
	IpConfigurations PrivateLinkServiceIpConfigurationResponseArrayOutput `pulumi:"ipConfigurations"`
	// An array of references to the load balancer IP configurations.
	LoadBalancerFrontendIpConfigurations FrontendIPConfigurationResponseArrayOutput `pulumi:"loadBalancerFrontendIpConfigurations"`
	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// An array of references to the network interfaces created for this private link service.
	NetworkInterfaces NetworkInterfaceResponseArrayOutput `pulumi:"networkInterfaces"`
	// An array of list about connections to the private endpoint.
	PrivateEndpointConnections PrivateEndpointConnectionResponseArrayOutput `pulumi:"privateEndpointConnections"`
	// The provisioning state of the private link service resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// The visibility list of the private link service.
	Visibility PrivateLinkServicePropertiesResponseVisibilityPtrOutput `pulumi:"visibility"`
}

// NewPrivateLinkService registers a new resource with the given unique name, arguments, and options.
func NewPrivateLinkService(ctx *pulumi.Context,
	name string, args *PrivateLinkServiceArgs, opts ...pulumi.ResourceOption) (*PrivateLinkService, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:PrivateLinkService"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190401:PrivateLinkService"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:PrivateLinkService"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:PrivateLinkService"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:PrivateLinkService"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:PrivateLinkService"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:PrivateLinkService"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:PrivateLinkService"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:PrivateLinkService"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:PrivateLinkService"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:PrivateLinkService"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:PrivateLinkService"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:PrivateLinkService"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:PrivateLinkService"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:PrivateLinkService"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:PrivateLinkService"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:PrivateLinkService"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:PrivateLinkService"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:PrivateLinkService"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:PrivateLinkService"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:PrivateLinkService"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220701:PrivateLinkService"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220901:PrivateLinkService"),
		},
		{
			Type: pulumi.String("azure-native:network/v20221101:PrivateLinkService"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230401:PrivateLinkService"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230501:PrivateLinkService"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230601:PrivateLinkService"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource PrivateLinkService
	err := ctx.RegisterResource("azure-native:network/v20230201:PrivateLinkService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPrivateLinkService gets an existing PrivateLinkService resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPrivateLinkService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateLinkServiceState, opts ...pulumi.ResourceOption) (*PrivateLinkService, error) {
	var resource PrivateLinkService
	err := ctx.ReadResource("azure-native:network/v20230201:PrivateLinkService", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PrivateLinkService resources.
type privateLinkServiceState struct {
}

type PrivateLinkServiceState struct {
}

func (PrivateLinkServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*privateLinkServiceState)(nil)).Elem()
}

type privateLinkServiceArgs struct {
	// The auto-approval list of the private link service.
	AutoApproval *PrivateLinkServicePropertiesAutoApproval `pulumi:"autoApproval"`
	// Whether the private link service is enabled for proxy protocol or not.
	EnableProxyProtocol *bool `pulumi:"enableProxyProtocol"`
	// The extended location of the load balancer.
	ExtendedLocation *ExtendedLocation `pulumi:"extendedLocation"`
	// The list of Fqdn.
	Fqdns []string `pulumi:"fqdns"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// An array of private link service IP configurations.
	IpConfigurations []PrivateLinkServiceIpConfiguration `pulumi:"ipConfigurations"`
	// An array of references to the load balancer IP configurations.
	LoadBalancerFrontendIpConfigurations []FrontendIPConfiguration `pulumi:"loadBalancerFrontendIpConfigurations"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the private link service.
	ServiceName *string `pulumi:"serviceName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The visibility list of the private link service.
	Visibility *PrivateLinkServicePropertiesVisibility `pulumi:"visibility"`
}

// The set of arguments for constructing a PrivateLinkService resource.
type PrivateLinkServiceArgs struct {
	// The auto-approval list of the private link service.
	AutoApproval PrivateLinkServicePropertiesAutoApprovalPtrInput
	// Whether the private link service is enabled for proxy protocol or not.
	EnableProxyProtocol pulumi.BoolPtrInput
	// The extended location of the load balancer.
	ExtendedLocation ExtendedLocationPtrInput
	// The list of Fqdn.
	Fqdns pulumi.StringArrayInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// An array of private link service IP configurations.
	IpConfigurations PrivateLinkServiceIpConfigurationArrayInput
	// An array of references to the load balancer IP configurations.
	LoadBalancerFrontendIpConfigurations FrontendIPConfigurationArrayInput
	// Resource location.
	Location pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The name of the private link service.
	ServiceName pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The visibility list of the private link service.
	Visibility PrivateLinkServicePropertiesVisibilityPtrInput
}

func (PrivateLinkServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privateLinkServiceArgs)(nil)).Elem()
}

type PrivateLinkServiceInput interface {
	pulumi.Input

	ToPrivateLinkServiceOutput() PrivateLinkServiceOutput
	ToPrivateLinkServiceOutputWithContext(ctx context.Context) PrivateLinkServiceOutput
}

func (*PrivateLinkService) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkService)(nil)).Elem()
}

func (i *PrivateLinkService) ToPrivateLinkServiceOutput() PrivateLinkServiceOutput {
	return i.ToPrivateLinkServiceOutputWithContext(context.Background())
}

func (i *PrivateLinkService) ToPrivateLinkServiceOutputWithContext(ctx context.Context) PrivateLinkServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceOutput)
}

type PrivateLinkServiceOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkService)(nil)).Elem()
}

func (o PrivateLinkServiceOutput) ToPrivateLinkServiceOutput() PrivateLinkServiceOutput {
	return o
}

func (o PrivateLinkServiceOutput) ToPrivateLinkServiceOutputWithContext(ctx context.Context) PrivateLinkServiceOutput {
	return o
}

// The alias of the private link service.
func (o PrivateLinkServiceOutput) Alias() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateLinkService) pulumi.StringOutput { return v.Alias }).(pulumi.StringOutput)
}

// The auto-approval list of the private link service.
func (o PrivateLinkServiceOutput) AutoApproval() PrivateLinkServicePropertiesResponseAutoApprovalPtrOutput {
	return o.ApplyT(func(v *PrivateLinkService) PrivateLinkServicePropertiesResponseAutoApprovalPtrOutput {
		return v.AutoApproval
	}).(PrivateLinkServicePropertiesResponseAutoApprovalPtrOutput)
}

// Whether the private link service is enabled for proxy protocol or not.
func (o PrivateLinkServiceOutput) EnableProxyProtocol() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *PrivateLinkService) pulumi.BoolPtrOutput { return v.EnableProxyProtocol }).(pulumi.BoolPtrOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o PrivateLinkServiceOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateLinkService) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The extended location of the load balancer.
func (o PrivateLinkServiceOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v *PrivateLinkService) ExtendedLocationResponsePtrOutput { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

// The list of Fqdn.
func (o PrivateLinkServiceOutput) Fqdns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *PrivateLinkService) pulumi.StringArrayOutput { return v.Fqdns }).(pulumi.StringArrayOutput)
}

// An array of private link service IP configurations.
func (o PrivateLinkServiceOutput) IpConfigurations() PrivateLinkServiceIpConfigurationResponseArrayOutput {
	return o.ApplyT(func(v *PrivateLinkService) PrivateLinkServiceIpConfigurationResponseArrayOutput {
		return v.IpConfigurations
	}).(PrivateLinkServiceIpConfigurationResponseArrayOutput)
}

// An array of references to the load balancer IP configurations.
func (o PrivateLinkServiceOutput) LoadBalancerFrontendIpConfigurations() FrontendIPConfigurationResponseArrayOutput {
	return o.ApplyT(func(v *PrivateLinkService) FrontendIPConfigurationResponseArrayOutput {
		return v.LoadBalancerFrontendIpConfigurations
	}).(FrontendIPConfigurationResponseArrayOutput)
}

// Resource location.
func (o PrivateLinkServiceOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkService) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o PrivateLinkServiceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateLinkService) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// An array of references to the network interfaces created for this private link service.
func (o PrivateLinkServiceOutput) NetworkInterfaces() NetworkInterfaceResponseArrayOutput {
	return o.ApplyT(func(v *PrivateLinkService) NetworkInterfaceResponseArrayOutput { return v.NetworkInterfaces }).(NetworkInterfaceResponseArrayOutput)
}

// An array of list about connections to the private endpoint.
func (o PrivateLinkServiceOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v *PrivateLinkService) PrivateEndpointConnectionResponseArrayOutput {
		return v.PrivateEndpointConnections
	}).(PrivateEndpointConnectionResponseArrayOutput)
}

// The provisioning state of the private link service resource.
func (o PrivateLinkServiceOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateLinkService) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Resource tags.
func (o PrivateLinkServiceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *PrivateLinkService) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o PrivateLinkServiceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateLinkService) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The visibility list of the private link service.
func (o PrivateLinkServiceOutput) Visibility() PrivateLinkServicePropertiesResponseVisibilityPtrOutput {
	return o.ApplyT(func(v *PrivateLinkService) PrivateLinkServicePropertiesResponseVisibilityPtrOutput {
		return v.Visibility
	}).(PrivateLinkServicePropertiesResponseVisibilityPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(PrivateLinkServiceOutput{})
}
