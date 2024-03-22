// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230301

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// NSX Segment
type WorkloadNetworkSegment struct {
	pulumi.CustomResourceState

	// Gateway which to connect segment to.
	ConnectedGateway pulumi.StringPtrOutput `pulumi:"connectedGateway"`
	// Display name of the segment.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Port Vif which segment is associated with.
	PortVif WorkloadNetworkSegmentPortVifResponseArrayOutput `pulumi:"portVif"`
	// The provisioning state
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// NSX revision number.
	Revision pulumi.Float64PtrOutput `pulumi:"revision"`
	// Segment status.
	Status pulumi.StringOutput `pulumi:"status"`
	// Subnet which to connect segment to.
	Subnet WorkloadNetworkSegmentSubnetResponsePtrOutput `pulumi:"subnet"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewWorkloadNetworkSegment registers a new resource with the given unique name, arguments, and options.
func NewWorkloadNetworkSegment(ctx *pulumi.Context,
	name string, args *WorkloadNetworkSegmentArgs, opts ...pulumi.ResourceOption) (*WorkloadNetworkSegment, error) {
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
			Type: pulumi.String("azure-native:avs:WorkloadNetworkSegment"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20200717preview:WorkloadNetworkSegment"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20210101preview:WorkloadNetworkSegment"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20210601:WorkloadNetworkSegment"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20211201:WorkloadNetworkSegment"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20220501:WorkloadNetworkSegment"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20230901:WorkloadNetworkSegment"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource WorkloadNetworkSegment
	err := ctx.RegisterResource("azure-native:avs/v20230301:WorkloadNetworkSegment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkloadNetworkSegment gets an existing WorkloadNetworkSegment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkloadNetworkSegment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkloadNetworkSegmentState, opts ...pulumi.ResourceOption) (*WorkloadNetworkSegment, error) {
	var resource WorkloadNetworkSegment
	err := ctx.ReadResource("azure-native:avs/v20230301:WorkloadNetworkSegment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WorkloadNetworkSegment resources.
type workloadNetworkSegmentState struct {
}

type WorkloadNetworkSegmentState struct {
}

func (WorkloadNetworkSegmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*workloadNetworkSegmentState)(nil)).Elem()
}

type workloadNetworkSegmentArgs struct {
	// Gateway which to connect segment to.
	ConnectedGateway *string `pulumi:"connectedGateway"`
	// Display name of the segment.
	DisplayName *string `pulumi:"displayName"`
	// Name of the private cloud
	PrivateCloudName string `pulumi:"privateCloudName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// NSX revision number.
	Revision *float64 `pulumi:"revision"`
	// NSX Segment identifier. Generally the same as the Segment's display name
	SegmentId *string `pulumi:"segmentId"`
	// Subnet which to connect segment to.
	Subnet *WorkloadNetworkSegmentSubnet `pulumi:"subnet"`
}

// The set of arguments for constructing a WorkloadNetworkSegment resource.
type WorkloadNetworkSegmentArgs struct {
	// Gateway which to connect segment to.
	ConnectedGateway pulumi.StringPtrInput
	// Display name of the segment.
	DisplayName pulumi.StringPtrInput
	// Name of the private cloud
	PrivateCloudName pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// NSX revision number.
	Revision pulumi.Float64PtrInput
	// NSX Segment identifier. Generally the same as the Segment's display name
	SegmentId pulumi.StringPtrInput
	// Subnet which to connect segment to.
	Subnet WorkloadNetworkSegmentSubnetPtrInput
}

func (WorkloadNetworkSegmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workloadNetworkSegmentArgs)(nil)).Elem()
}

type WorkloadNetworkSegmentInput interface {
	pulumi.Input

	ToWorkloadNetworkSegmentOutput() WorkloadNetworkSegmentOutput
	ToWorkloadNetworkSegmentOutputWithContext(ctx context.Context) WorkloadNetworkSegmentOutput
}

func (*WorkloadNetworkSegment) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkloadNetworkSegment)(nil)).Elem()
}

func (i *WorkloadNetworkSegment) ToWorkloadNetworkSegmentOutput() WorkloadNetworkSegmentOutput {
	return i.ToWorkloadNetworkSegmentOutputWithContext(context.Background())
}

func (i *WorkloadNetworkSegment) ToWorkloadNetworkSegmentOutputWithContext(ctx context.Context) WorkloadNetworkSegmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkloadNetworkSegmentOutput)
}

type WorkloadNetworkSegmentOutput struct{ *pulumi.OutputState }

func (WorkloadNetworkSegmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkloadNetworkSegment)(nil)).Elem()
}

func (o WorkloadNetworkSegmentOutput) ToWorkloadNetworkSegmentOutput() WorkloadNetworkSegmentOutput {
	return o
}

func (o WorkloadNetworkSegmentOutput) ToWorkloadNetworkSegmentOutputWithContext(ctx context.Context) WorkloadNetworkSegmentOutput {
	return o
}

// Gateway which to connect segment to.
func (o WorkloadNetworkSegmentOutput) ConnectedGateway() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkloadNetworkSegment) pulumi.StringPtrOutput { return v.ConnectedGateway }).(pulumi.StringPtrOutput)
}

// Display name of the segment.
func (o WorkloadNetworkSegmentOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkloadNetworkSegment) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o WorkloadNetworkSegmentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkloadNetworkSegment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Port Vif which segment is associated with.
func (o WorkloadNetworkSegmentOutput) PortVif() WorkloadNetworkSegmentPortVifResponseArrayOutput {
	return o.ApplyT(func(v *WorkloadNetworkSegment) WorkloadNetworkSegmentPortVifResponseArrayOutput { return v.PortVif }).(WorkloadNetworkSegmentPortVifResponseArrayOutput)
}

// The provisioning state
func (o WorkloadNetworkSegmentOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkloadNetworkSegment) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// NSX revision number.
func (o WorkloadNetworkSegmentOutput) Revision() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *WorkloadNetworkSegment) pulumi.Float64PtrOutput { return v.Revision }).(pulumi.Float64PtrOutput)
}

// Segment status.
func (o WorkloadNetworkSegmentOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkloadNetworkSegment) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Subnet which to connect segment to.
func (o WorkloadNetworkSegmentOutput) Subnet() WorkloadNetworkSegmentSubnetResponsePtrOutput {
	return o.ApplyT(func(v *WorkloadNetworkSegment) WorkloadNetworkSegmentSubnetResponsePtrOutput { return v.Subnet }).(WorkloadNetworkSegmentSubnetResponsePtrOutput)
}

// Resource type.
func (o WorkloadNetworkSegmentOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkloadNetworkSegment) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WorkloadNetworkSegmentOutput{})
}
