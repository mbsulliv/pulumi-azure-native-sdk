// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package avs

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// NSX DHCP
// Azure REST API version: 2022-05-01. Prior API version in Azure Native 1.x: 2020-07-17-preview
type WorkloadNetworkDhcp struct {
	pulumi.CustomResourceState

	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// DHCP properties.
	Properties pulumi.AnyOutput `pulumi:"properties"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewWorkloadNetworkDhcp registers a new resource with the given unique name, arguments, and options.
func NewWorkloadNetworkDhcp(ctx *pulumi.Context,
	name string, args *WorkloadNetworkDhcpArgs, opts ...pulumi.ResourceOption) (*WorkloadNetworkDhcp, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PrivateCloudName == nil {
		return nil, errors.New("invalid value for required argument 'PrivateCloudName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:avs/v20200717preview:WorkloadNetworkDhcp"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20210101preview:WorkloadNetworkDhcp"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20210601:WorkloadNetworkDhcp"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20211201:WorkloadNetworkDhcp"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20220501:WorkloadNetworkDhcp"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20230301:WorkloadNetworkDhcp"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource WorkloadNetworkDhcp
	err := ctx.RegisterResource("azure-native:avs:WorkloadNetworkDhcp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkloadNetworkDhcp gets an existing WorkloadNetworkDhcp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkloadNetworkDhcp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkloadNetworkDhcpState, opts ...pulumi.ResourceOption) (*WorkloadNetworkDhcp, error) {
	var resource WorkloadNetworkDhcp
	err := ctx.ReadResource("azure-native:avs:WorkloadNetworkDhcp", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WorkloadNetworkDhcp resources.
type workloadNetworkDhcpState struct {
}

type WorkloadNetworkDhcpState struct {
}

func (WorkloadNetworkDhcpState) ElementType() reflect.Type {
	return reflect.TypeOf((*workloadNetworkDhcpState)(nil)).Elem()
}

type workloadNetworkDhcpArgs struct {
	// NSX DHCP identifier. Generally the same as the DHCP display name
	DhcpId *string `pulumi:"dhcpId"`
	// Name of the private cloud
	PrivateCloudName string `pulumi:"privateCloudName"`
	// DHCP properties.
	Properties interface{} `pulumi:"properties"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a WorkloadNetworkDhcp resource.
type WorkloadNetworkDhcpArgs struct {
	// NSX DHCP identifier. Generally the same as the DHCP display name
	DhcpId pulumi.StringPtrInput
	// Name of the private cloud
	PrivateCloudName pulumi.StringInput
	// DHCP properties.
	Properties pulumi.Input
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
}

func (WorkloadNetworkDhcpArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workloadNetworkDhcpArgs)(nil)).Elem()
}

type WorkloadNetworkDhcpInput interface {
	pulumi.Input

	ToWorkloadNetworkDhcpOutput() WorkloadNetworkDhcpOutput
	ToWorkloadNetworkDhcpOutputWithContext(ctx context.Context) WorkloadNetworkDhcpOutput
}

func (*WorkloadNetworkDhcp) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkloadNetworkDhcp)(nil)).Elem()
}

func (i *WorkloadNetworkDhcp) ToWorkloadNetworkDhcpOutput() WorkloadNetworkDhcpOutput {
	return i.ToWorkloadNetworkDhcpOutputWithContext(context.Background())
}

func (i *WorkloadNetworkDhcp) ToWorkloadNetworkDhcpOutputWithContext(ctx context.Context) WorkloadNetworkDhcpOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkloadNetworkDhcpOutput)
}

func (i *WorkloadNetworkDhcp) ToOutput(ctx context.Context) pulumix.Output[*WorkloadNetworkDhcp] {
	return pulumix.Output[*WorkloadNetworkDhcp]{
		OutputState: i.ToWorkloadNetworkDhcpOutputWithContext(ctx).OutputState,
	}
}

type WorkloadNetworkDhcpOutput struct{ *pulumi.OutputState }

func (WorkloadNetworkDhcpOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkloadNetworkDhcp)(nil)).Elem()
}

func (o WorkloadNetworkDhcpOutput) ToWorkloadNetworkDhcpOutput() WorkloadNetworkDhcpOutput {
	return o
}

func (o WorkloadNetworkDhcpOutput) ToWorkloadNetworkDhcpOutputWithContext(ctx context.Context) WorkloadNetworkDhcpOutput {
	return o
}

func (o WorkloadNetworkDhcpOutput) ToOutput(ctx context.Context) pulumix.Output[*WorkloadNetworkDhcp] {
	return pulumix.Output[*WorkloadNetworkDhcp]{
		OutputState: o.OutputState,
	}
}

// Resource name.
func (o WorkloadNetworkDhcpOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkloadNetworkDhcp) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// DHCP properties.
func (o WorkloadNetworkDhcpOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v *WorkloadNetworkDhcp) pulumi.AnyOutput { return v.Properties }).(pulumi.AnyOutput)
}

// Resource type.
func (o WorkloadNetworkDhcpOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkloadNetworkDhcp) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WorkloadNetworkDhcpOutput{})
}
