// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220701

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Describes an inbound endpoint for a DNS resolver.
type InboundEndpoint struct {
	pulumi.CustomResourceState

	// ETag of the inbound endpoint.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// IP configurations for the inbound endpoint.
	IpConfigurations InboundEndpointIPConfigurationResponseArrayOutput `pulumi:"ipConfigurations"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The current provisioning state of the inbound endpoint. This is a read-only property and any attempt to set this value will be ignored.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The resourceGuid property of the inbound endpoint resource.
	ResourceGuid pulumi.StringOutput `pulumi:"resourceGuid"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewInboundEndpoint registers a new resource with the given unique name, arguments, and options.
func NewInboundEndpoint(ctx *pulumi.Context,
	name string, args *InboundEndpointArgs, opts ...pulumi.ResourceOption) (*InboundEndpoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DnsResolverName == nil {
		return nil, errors.New("invalid value for required argument 'DnsResolverName'")
	}
	if args.IpConfigurations == nil {
		return nil, errors.New("invalid value for required argument 'IpConfigurations'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:InboundEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401preview:InboundEndpoint"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource InboundEndpoint
	err := ctx.RegisterResource("azure-native:network/v20220701:InboundEndpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInboundEndpoint gets an existing InboundEndpoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInboundEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InboundEndpointState, opts ...pulumi.ResourceOption) (*InboundEndpoint, error) {
	var resource InboundEndpoint
	err := ctx.ReadResource("azure-native:network/v20220701:InboundEndpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InboundEndpoint resources.
type inboundEndpointState struct {
}

type InboundEndpointState struct {
}

func (InboundEndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*inboundEndpointState)(nil)).Elem()
}

type inboundEndpointArgs struct {
	// The name of the DNS resolver.
	DnsResolverName string `pulumi:"dnsResolverName"`
	// The name of the inbound endpoint for the DNS resolver.
	InboundEndpointName *string `pulumi:"inboundEndpointName"`
	// IP configurations for the inbound endpoint.
	IpConfigurations []InboundEndpointIPConfiguration `pulumi:"ipConfigurations"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a InboundEndpoint resource.
type InboundEndpointArgs struct {
	// The name of the DNS resolver.
	DnsResolverName pulumi.StringInput
	// The name of the inbound endpoint for the DNS resolver.
	InboundEndpointName pulumi.StringPtrInput
	// IP configurations for the inbound endpoint.
	IpConfigurations InboundEndpointIPConfigurationArrayInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (InboundEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*inboundEndpointArgs)(nil)).Elem()
}

type InboundEndpointInput interface {
	pulumi.Input

	ToInboundEndpointOutput() InboundEndpointOutput
	ToInboundEndpointOutputWithContext(ctx context.Context) InboundEndpointOutput
}

func (*InboundEndpoint) ElementType() reflect.Type {
	return reflect.TypeOf((**InboundEndpoint)(nil)).Elem()
}

func (i *InboundEndpoint) ToInboundEndpointOutput() InboundEndpointOutput {
	return i.ToInboundEndpointOutputWithContext(context.Background())
}

func (i *InboundEndpoint) ToInboundEndpointOutputWithContext(ctx context.Context) InboundEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InboundEndpointOutput)
}

type InboundEndpointOutput struct{ *pulumi.OutputState }

func (InboundEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InboundEndpoint)(nil)).Elem()
}

func (o InboundEndpointOutput) ToInboundEndpointOutput() InboundEndpointOutput {
	return o
}

func (o InboundEndpointOutput) ToInboundEndpointOutputWithContext(ctx context.Context) InboundEndpointOutput {
	return o
}

// ETag of the inbound endpoint.
func (o InboundEndpointOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *InboundEndpoint) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// IP configurations for the inbound endpoint.
func (o InboundEndpointOutput) IpConfigurations() InboundEndpointIPConfigurationResponseArrayOutput {
	return o.ApplyT(func(v *InboundEndpoint) InboundEndpointIPConfigurationResponseArrayOutput { return v.IpConfigurations }).(InboundEndpointIPConfigurationResponseArrayOutput)
}

// The geo-location where the resource lives
func (o InboundEndpointOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *InboundEndpoint) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o InboundEndpointOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *InboundEndpoint) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The current provisioning state of the inbound endpoint. This is a read-only property and any attempt to set this value will be ignored.
func (o InboundEndpointOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *InboundEndpoint) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The resourceGuid property of the inbound endpoint resource.
func (o InboundEndpointOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v *InboundEndpoint) pulumi.StringOutput { return v.ResourceGuid }).(pulumi.StringOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o InboundEndpointOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *InboundEndpoint) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o InboundEndpointOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *InboundEndpoint) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o InboundEndpointOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *InboundEndpoint) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(InboundEndpointOutput{})
}
