// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Describes an outbound endpoint for a DNS resolver.
// Azure REST API version: 2022-07-01. Prior API version in Azure Native 1.x: 2020-04-01-preview
type OutboundEndpoint struct {
	pulumi.CustomResourceState

	// ETag of the outbound endpoint.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The current provisioning state of the outbound endpoint. This is a read-only property and any attempt to set this value will be ignored.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The resourceGuid property of the outbound endpoint resource.
	ResourceGuid pulumi.StringOutput `pulumi:"resourceGuid"`
	// The reference to the subnet used for the outbound endpoint.
	Subnet SubResourceResponseOutput `pulumi:"subnet"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewOutboundEndpoint registers a new resource with the given unique name, arguments, and options.
func NewOutboundEndpoint(ctx *pulumi.Context,
	name string, args *OutboundEndpointArgs, opts ...pulumi.ResourceOption) (*OutboundEndpoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DnsResolverName == nil {
		return nil, errors.New("invalid value for required argument 'DnsResolverName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Subnet == nil {
		return nil, errors.New("invalid value for required argument 'Subnet'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network/v20200401preview:OutboundEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220701:OutboundEndpoint"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource OutboundEndpoint
	err := ctx.RegisterResource("azure-native:network:OutboundEndpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOutboundEndpoint gets an existing OutboundEndpoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOutboundEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OutboundEndpointState, opts ...pulumi.ResourceOption) (*OutboundEndpoint, error) {
	var resource OutboundEndpoint
	err := ctx.ReadResource("azure-native:network:OutboundEndpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OutboundEndpoint resources.
type outboundEndpointState struct {
}

type OutboundEndpointState struct {
}

func (OutboundEndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*outboundEndpointState)(nil)).Elem()
}

type outboundEndpointArgs struct {
	// The name of the DNS resolver.
	DnsResolverName string `pulumi:"dnsResolverName"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the outbound endpoint for the DNS resolver.
	OutboundEndpointName *string `pulumi:"outboundEndpointName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The reference to the subnet used for the outbound endpoint.
	Subnet SubResource `pulumi:"subnet"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a OutboundEndpoint resource.
type OutboundEndpointArgs struct {
	// The name of the DNS resolver.
	DnsResolverName pulumi.StringInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the outbound endpoint for the DNS resolver.
	OutboundEndpointName pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The reference to the subnet used for the outbound endpoint.
	Subnet SubResourceInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (OutboundEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*outboundEndpointArgs)(nil)).Elem()
}

type OutboundEndpointInput interface {
	pulumi.Input

	ToOutboundEndpointOutput() OutboundEndpointOutput
	ToOutboundEndpointOutputWithContext(ctx context.Context) OutboundEndpointOutput
}

func (*OutboundEndpoint) ElementType() reflect.Type {
	return reflect.TypeOf((**OutboundEndpoint)(nil)).Elem()
}

func (i *OutboundEndpoint) ToOutboundEndpointOutput() OutboundEndpointOutput {
	return i.ToOutboundEndpointOutputWithContext(context.Background())
}

func (i *OutboundEndpoint) ToOutboundEndpointOutputWithContext(ctx context.Context) OutboundEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OutboundEndpointOutput)
}

func (i *OutboundEndpoint) ToOutput(ctx context.Context) pulumix.Output[*OutboundEndpoint] {
	return pulumix.Output[*OutboundEndpoint]{
		OutputState: i.ToOutboundEndpointOutputWithContext(ctx).OutputState,
	}
}

type OutboundEndpointOutput struct{ *pulumi.OutputState }

func (OutboundEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OutboundEndpoint)(nil)).Elem()
}

func (o OutboundEndpointOutput) ToOutboundEndpointOutput() OutboundEndpointOutput {
	return o
}

func (o OutboundEndpointOutput) ToOutboundEndpointOutputWithContext(ctx context.Context) OutboundEndpointOutput {
	return o
}

func (o OutboundEndpointOutput) ToOutput(ctx context.Context) pulumix.Output[*OutboundEndpoint] {
	return pulumix.Output[*OutboundEndpoint]{
		OutputState: o.OutputState,
	}
}

// ETag of the outbound endpoint.
func (o OutboundEndpointOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *OutboundEndpoint) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o OutboundEndpointOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *OutboundEndpoint) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o OutboundEndpointOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *OutboundEndpoint) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The current provisioning state of the outbound endpoint. This is a read-only property and any attempt to set this value will be ignored.
func (o OutboundEndpointOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *OutboundEndpoint) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The resourceGuid property of the outbound endpoint resource.
func (o OutboundEndpointOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v *OutboundEndpoint) pulumi.StringOutput { return v.ResourceGuid }).(pulumi.StringOutput)
}

// The reference to the subnet used for the outbound endpoint.
func (o OutboundEndpointOutput) Subnet() SubResourceResponseOutput {
	return o.ApplyT(func(v *OutboundEndpoint) SubResourceResponseOutput { return v.Subnet }).(SubResourceResponseOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o OutboundEndpointOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *OutboundEndpoint) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o OutboundEndpointOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *OutboundEndpoint) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o OutboundEndpointOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *OutboundEndpoint) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(OutboundEndpointOutput{})
}
