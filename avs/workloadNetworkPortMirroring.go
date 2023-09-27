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

// NSX Port Mirroring
// Azure REST API version: 2022-05-01. Prior API version in Azure Native 1.x: 2020-07-17-preview
type WorkloadNetworkPortMirroring struct {
	pulumi.CustomResourceState

	// Destination VM Group.
	Destination pulumi.StringPtrOutput `pulumi:"destination"`
	// Direction of port mirroring profile.
	Direction pulumi.StringPtrOutput `pulumi:"direction"`
	// Display name of the port mirroring profile.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioning state
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// NSX revision number.
	Revision pulumi.Float64PtrOutput `pulumi:"revision"`
	// Source VM Group.
	Source pulumi.StringPtrOutput `pulumi:"source"`
	// Port Mirroring Status.
	Status pulumi.StringOutput `pulumi:"status"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewWorkloadNetworkPortMirroring registers a new resource with the given unique name, arguments, and options.
func NewWorkloadNetworkPortMirroring(ctx *pulumi.Context,
	name string, args *WorkloadNetworkPortMirroringArgs, opts ...pulumi.ResourceOption) (*WorkloadNetworkPortMirroring, error) {
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
			Type: pulumi.String("azure-native:avs/v20200717preview:WorkloadNetworkPortMirroring"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20210101preview:WorkloadNetworkPortMirroring"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20210601:WorkloadNetworkPortMirroring"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20211201:WorkloadNetworkPortMirroring"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20220501:WorkloadNetworkPortMirroring"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20230301:WorkloadNetworkPortMirroring"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource WorkloadNetworkPortMirroring
	err := ctx.RegisterResource("azure-native:avs:WorkloadNetworkPortMirroring", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkloadNetworkPortMirroring gets an existing WorkloadNetworkPortMirroring resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkloadNetworkPortMirroring(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkloadNetworkPortMirroringState, opts ...pulumi.ResourceOption) (*WorkloadNetworkPortMirroring, error) {
	var resource WorkloadNetworkPortMirroring
	err := ctx.ReadResource("azure-native:avs:WorkloadNetworkPortMirroring", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WorkloadNetworkPortMirroring resources.
type workloadNetworkPortMirroringState struct {
}

type WorkloadNetworkPortMirroringState struct {
}

func (WorkloadNetworkPortMirroringState) ElementType() reflect.Type {
	return reflect.TypeOf((*workloadNetworkPortMirroringState)(nil)).Elem()
}

type workloadNetworkPortMirroringArgs struct {
	// Destination VM Group.
	Destination *string `pulumi:"destination"`
	// Direction of port mirroring profile.
	Direction *string `pulumi:"direction"`
	// Display name of the port mirroring profile.
	DisplayName *string `pulumi:"displayName"`
	// NSX Port Mirroring identifier. Generally the same as the Port Mirroring display name
	PortMirroringId *string `pulumi:"portMirroringId"`
	// Name of the private cloud
	PrivateCloudName string `pulumi:"privateCloudName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// NSX revision number.
	Revision *float64 `pulumi:"revision"`
	// Source VM Group.
	Source *string `pulumi:"source"`
}

// The set of arguments for constructing a WorkloadNetworkPortMirroring resource.
type WorkloadNetworkPortMirroringArgs struct {
	// Destination VM Group.
	Destination pulumi.StringPtrInput
	// Direction of port mirroring profile.
	Direction pulumi.StringPtrInput
	// Display name of the port mirroring profile.
	DisplayName pulumi.StringPtrInput
	// NSX Port Mirroring identifier. Generally the same as the Port Mirroring display name
	PortMirroringId pulumi.StringPtrInput
	// Name of the private cloud
	PrivateCloudName pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// NSX revision number.
	Revision pulumi.Float64PtrInput
	// Source VM Group.
	Source pulumi.StringPtrInput
}

func (WorkloadNetworkPortMirroringArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workloadNetworkPortMirroringArgs)(nil)).Elem()
}

type WorkloadNetworkPortMirroringInput interface {
	pulumi.Input

	ToWorkloadNetworkPortMirroringOutput() WorkloadNetworkPortMirroringOutput
	ToWorkloadNetworkPortMirroringOutputWithContext(ctx context.Context) WorkloadNetworkPortMirroringOutput
}

func (*WorkloadNetworkPortMirroring) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkloadNetworkPortMirroring)(nil)).Elem()
}

func (i *WorkloadNetworkPortMirroring) ToWorkloadNetworkPortMirroringOutput() WorkloadNetworkPortMirroringOutput {
	return i.ToWorkloadNetworkPortMirroringOutputWithContext(context.Background())
}

func (i *WorkloadNetworkPortMirroring) ToWorkloadNetworkPortMirroringOutputWithContext(ctx context.Context) WorkloadNetworkPortMirroringOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkloadNetworkPortMirroringOutput)
}

func (i *WorkloadNetworkPortMirroring) ToOutput(ctx context.Context) pulumix.Output[*WorkloadNetworkPortMirroring] {
	return pulumix.Output[*WorkloadNetworkPortMirroring]{
		OutputState: i.ToWorkloadNetworkPortMirroringOutputWithContext(ctx).OutputState,
	}
}

type WorkloadNetworkPortMirroringOutput struct{ *pulumi.OutputState }

func (WorkloadNetworkPortMirroringOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkloadNetworkPortMirroring)(nil)).Elem()
}

func (o WorkloadNetworkPortMirroringOutput) ToWorkloadNetworkPortMirroringOutput() WorkloadNetworkPortMirroringOutput {
	return o
}

func (o WorkloadNetworkPortMirroringOutput) ToWorkloadNetworkPortMirroringOutputWithContext(ctx context.Context) WorkloadNetworkPortMirroringOutput {
	return o
}

func (o WorkloadNetworkPortMirroringOutput) ToOutput(ctx context.Context) pulumix.Output[*WorkloadNetworkPortMirroring] {
	return pulumix.Output[*WorkloadNetworkPortMirroring]{
		OutputState: o.OutputState,
	}
}

// Destination VM Group.
func (o WorkloadNetworkPortMirroringOutput) Destination() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkloadNetworkPortMirroring) pulumi.StringPtrOutput { return v.Destination }).(pulumi.StringPtrOutput)
}

// Direction of port mirroring profile.
func (o WorkloadNetworkPortMirroringOutput) Direction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkloadNetworkPortMirroring) pulumi.StringPtrOutput { return v.Direction }).(pulumi.StringPtrOutput)
}

// Display name of the port mirroring profile.
func (o WorkloadNetworkPortMirroringOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkloadNetworkPortMirroring) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o WorkloadNetworkPortMirroringOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkloadNetworkPortMirroring) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state
func (o WorkloadNetworkPortMirroringOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkloadNetworkPortMirroring) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// NSX revision number.
func (o WorkloadNetworkPortMirroringOutput) Revision() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *WorkloadNetworkPortMirroring) pulumi.Float64PtrOutput { return v.Revision }).(pulumi.Float64PtrOutput)
}

// Source VM Group.
func (o WorkloadNetworkPortMirroringOutput) Source() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkloadNetworkPortMirroring) pulumi.StringPtrOutput { return v.Source }).(pulumi.StringPtrOutput)
}

// Port Mirroring Status.
func (o WorkloadNetworkPortMirroringOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkloadNetworkPortMirroring) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Resource type.
func (o WorkloadNetworkPortMirroringOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkloadNetworkPortMirroring) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WorkloadNetworkPortMirroringOutput{})
}
