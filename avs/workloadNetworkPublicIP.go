// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package avs

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// NSX Public IP Block
// Azure REST API version: 2022-05-01. Prior API version in Azure Native 1.x: 2021-06-01.
//
// Other available API versions: 2023-03-01, 2023-09-01.
type WorkloadNetworkPublicIP struct {
	pulumi.CustomResourceState

	// Display name of the Public IP Block.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Number of Public IPs requested.
	NumberOfPublicIPs pulumi.Float64PtrOutput `pulumi:"numberOfPublicIPs"`
	// The provisioning state
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// CIDR Block of the Public IP Block.
	PublicIPBlock pulumi.StringOutput `pulumi:"publicIPBlock"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewWorkloadNetworkPublicIP registers a new resource with the given unique name, arguments, and options.
func NewWorkloadNetworkPublicIP(ctx *pulumi.Context,
	name string, args *WorkloadNetworkPublicIPArgs, opts ...pulumi.ResourceOption) (*WorkloadNetworkPublicIP, error) {
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
			Type: pulumi.String("azure-native:avs/v20210601:WorkloadNetworkPublicIP"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20211201:WorkloadNetworkPublicIP"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20220501:WorkloadNetworkPublicIP"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20230301:WorkloadNetworkPublicIP"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20230901:WorkloadNetworkPublicIP"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource WorkloadNetworkPublicIP
	err := ctx.RegisterResource("azure-native:avs:WorkloadNetworkPublicIP", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkloadNetworkPublicIP gets an existing WorkloadNetworkPublicIP resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkloadNetworkPublicIP(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkloadNetworkPublicIPState, opts ...pulumi.ResourceOption) (*WorkloadNetworkPublicIP, error) {
	var resource WorkloadNetworkPublicIP
	err := ctx.ReadResource("azure-native:avs:WorkloadNetworkPublicIP", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WorkloadNetworkPublicIP resources.
type workloadNetworkPublicIPState struct {
}

type WorkloadNetworkPublicIPState struct {
}

func (WorkloadNetworkPublicIPState) ElementType() reflect.Type {
	return reflect.TypeOf((*workloadNetworkPublicIPState)(nil)).Elem()
}

type workloadNetworkPublicIPArgs struct {
	// Display name of the Public IP Block.
	DisplayName *string `pulumi:"displayName"`
	// Number of Public IPs requested.
	NumberOfPublicIPs *float64 `pulumi:"numberOfPublicIPs"`
	// Name of the private cloud
	PrivateCloudName string `pulumi:"privateCloudName"`
	// NSX Public IP Block identifier. Generally the same as the Public IP Block's display name
	PublicIPId *string `pulumi:"publicIPId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a WorkloadNetworkPublicIP resource.
type WorkloadNetworkPublicIPArgs struct {
	// Display name of the Public IP Block.
	DisplayName pulumi.StringPtrInput
	// Number of Public IPs requested.
	NumberOfPublicIPs pulumi.Float64PtrInput
	// Name of the private cloud
	PrivateCloudName pulumi.StringInput
	// NSX Public IP Block identifier. Generally the same as the Public IP Block's display name
	PublicIPId pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
}

func (WorkloadNetworkPublicIPArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workloadNetworkPublicIPArgs)(nil)).Elem()
}

type WorkloadNetworkPublicIPInput interface {
	pulumi.Input

	ToWorkloadNetworkPublicIPOutput() WorkloadNetworkPublicIPOutput
	ToWorkloadNetworkPublicIPOutputWithContext(ctx context.Context) WorkloadNetworkPublicIPOutput
}

func (*WorkloadNetworkPublicIP) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkloadNetworkPublicIP)(nil)).Elem()
}

func (i *WorkloadNetworkPublicIP) ToWorkloadNetworkPublicIPOutput() WorkloadNetworkPublicIPOutput {
	return i.ToWorkloadNetworkPublicIPOutputWithContext(context.Background())
}

func (i *WorkloadNetworkPublicIP) ToWorkloadNetworkPublicIPOutputWithContext(ctx context.Context) WorkloadNetworkPublicIPOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkloadNetworkPublicIPOutput)
}

type WorkloadNetworkPublicIPOutput struct{ *pulumi.OutputState }

func (WorkloadNetworkPublicIPOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkloadNetworkPublicIP)(nil)).Elem()
}

func (o WorkloadNetworkPublicIPOutput) ToWorkloadNetworkPublicIPOutput() WorkloadNetworkPublicIPOutput {
	return o
}

func (o WorkloadNetworkPublicIPOutput) ToWorkloadNetworkPublicIPOutputWithContext(ctx context.Context) WorkloadNetworkPublicIPOutput {
	return o
}

// Display name of the Public IP Block.
func (o WorkloadNetworkPublicIPOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkloadNetworkPublicIP) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o WorkloadNetworkPublicIPOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkloadNetworkPublicIP) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Number of Public IPs requested.
func (o WorkloadNetworkPublicIPOutput) NumberOfPublicIPs() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *WorkloadNetworkPublicIP) pulumi.Float64PtrOutput { return v.NumberOfPublicIPs }).(pulumi.Float64PtrOutput)
}

// The provisioning state
func (o WorkloadNetworkPublicIPOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkloadNetworkPublicIP) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// CIDR Block of the Public IP Block.
func (o WorkloadNetworkPublicIPOutput) PublicIPBlock() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkloadNetworkPublicIP) pulumi.StringOutput { return v.PublicIPBlock }).(pulumi.StringOutput)
}

// Resource type.
func (o WorkloadNetworkPublicIPOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkloadNetworkPublicIP) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WorkloadNetworkPublicIPOutput{})
}
