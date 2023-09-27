// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20190401

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Dedicated cloud node model
type DedicatedCloudNode struct {
	pulumi.CustomResourceState

	// Availability Zone id, e.g. "az1"
	AvailabilityZoneId pulumi.StringOutput `pulumi:"availabilityZoneId"`
	// Availability Zone name, e.g. "Availability Zone 1"
	AvailabilityZoneName pulumi.StringOutput `pulumi:"availabilityZoneName"`
	// VMWare Cloud Rack Name
	CloudRackName pulumi.StringOutput `pulumi:"cloudRackName"`
	// date time the resource was created
	Created pulumi.StringOutput `pulumi:"created"`
	// Azure region
	Location pulumi.StringOutput `pulumi:"location"`
	// SKU's name
	Name pulumi.StringOutput `pulumi:"name"`
	// count of nodes to create
	NodesCount pulumi.IntOutput `pulumi:"nodesCount"`
	// Placement Group id, e.g. "n1"
	PlacementGroupId pulumi.StringOutput `pulumi:"placementGroupId"`
	// Placement Name, e.g. "Placement Group 1"
	PlacementGroupName pulumi.StringOutput `pulumi:"placementGroupName"`
	// Private Cloud Id
	PrivateCloudId pulumi.StringOutput `pulumi:"privateCloudId"`
	// Resource Pool Name
	PrivateCloudName pulumi.StringOutput `pulumi:"privateCloudName"`
	// The provisioning status of the resource
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// purchase id
	PurchaseId pulumi.StringOutput `pulumi:"purchaseId"`
	// Dedicated Cloud Nodes SKU
	Sku SkuResponsePtrOutput `pulumi:"sku"`
	// Node status, indicates is private cloud set up on this node or not
	Status pulumi.StringOutput `pulumi:"status"`
	// Dedicated Cloud Nodes tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// {resourceProviderNamespace}/{resourceType}
	Type pulumi.StringOutput `pulumi:"type"`
	// VMWare Cluster Name
	VmwareClusterName pulumi.StringOutput `pulumi:"vmwareClusterName"`
}

// NewDedicatedCloudNode registers a new resource with the given unique name, arguments, and options.
func NewDedicatedCloudNode(ctx *pulumi.Context,
	name string, args *DedicatedCloudNodeArgs, opts ...pulumi.ResourceOption) (*DedicatedCloudNode, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AvailabilityZoneId == nil {
		return nil, errors.New("invalid value for required argument 'AvailabilityZoneId'")
	}
	if args.Id == nil {
		return nil, errors.New("invalid value for required argument 'Id'")
	}
	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.NodesCount == nil {
		return nil, errors.New("invalid value for required argument 'NodesCount'")
	}
	if args.PlacementGroupId == nil {
		return nil, errors.New("invalid value for required argument 'PlacementGroupId'")
	}
	if args.PurchaseId == nil {
		return nil, errors.New("invalid value for required argument 'PurchaseId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:vmwarecloudsimple:DedicatedCloudNode"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource DedicatedCloudNode
	err := ctx.RegisterResource("azure-native:vmwarecloudsimple/v20190401:DedicatedCloudNode", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDedicatedCloudNode gets an existing DedicatedCloudNode resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDedicatedCloudNode(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DedicatedCloudNodeState, opts ...pulumi.ResourceOption) (*DedicatedCloudNode, error) {
	var resource DedicatedCloudNode
	err := ctx.ReadResource("azure-native:vmwarecloudsimple/v20190401:DedicatedCloudNode", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DedicatedCloudNode resources.
type dedicatedCloudNodeState struct {
}

type DedicatedCloudNodeState struct {
}

func (DedicatedCloudNodeState) ElementType() reflect.Type {
	return reflect.TypeOf((*dedicatedCloudNodeState)(nil)).Elem()
}

type dedicatedCloudNodeArgs struct {
	// Availability Zone id, e.g. "az1"
	AvailabilityZoneId string `pulumi:"availabilityZoneId"`
	// dedicated cloud node name
	DedicatedCloudNodeName *string `pulumi:"dedicatedCloudNodeName"`
	// SKU's id
	Id string `pulumi:"id"`
	// Azure region
	Location *string `pulumi:"location"`
	// SKU's name
	Name string `pulumi:"name"`
	// count of nodes to create
	NodesCount int `pulumi:"nodesCount"`
	// Placement Group id, e.g. "n1"
	PlacementGroupId string `pulumi:"placementGroupId"`
	// purchase id
	PurchaseId string `pulumi:"purchaseId"`
	// The name of the resource group
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Dedicated Cloud Nodes SKU
	Sku *Sku `pulumi:"sku"`
	// Dedicated Cloud Nodes tags
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a DedicatedCloudNode resource.
type DedicatedCloudNodeArgs struct {
	// Availability Zone id, e.g. "az1"
	AvailabilityZoneId pulumi.StringInput
	// dedicated cloud node name
	DedicatedCloudNodeName pulumi.StringPtrInput
	// SKU's id
	Id pulumi.StringInput
	// Azure region
	Location pulumi.StringPtrInput
	// SKU's name
	Name pulumi.StringInput
	// count of nodes to create
	NodesCount pulumi.IntInput
	// Placement Group id, e.g. "n1"
	PlacementGroupId pulumi.StringInput
	// purchase id
	PurchaseId pulumi.StringInput
	// The name of the resource group
	ResourceGroupName pulumi.StringInput
	// Dedicated Cloud Nodes SKU
	Sku SkuPtrInput
	// Dedicated Cloud Nodes tags
	Tags pulumi.StringMapInput
}

func (DedicatedCloudNodeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dedicatedCloudNodeArgs)(nil)).Elem()
}

type DedicatedCloudNodeInput interface {
	pulumi.Input

	ToDedicatedCloudNodeOutput() DedicatedCloudNodeOutput
	ToDedicatedCloudNodeOutputWithContext(ctx context.Context) DedicatedCloudNodeOutput
}

func (*DedicatedCloudNode) ElementType() reflect.Type {
	return reflect.TypeOf((**DedicatedCloudNode)(nil)).Elem()
}

func (i *DedicatedCloudNode) ToDedicatedCloudNodeOutput() DedicatedCloudNodeOutput {
	return i.ToDedicatedCloudNodeOutputWithContext(context.Background())
}

func (i *DedicatedCloudNode) ToDedicatedCloudNodeOutputWithContext(ctx context.Context) DedicatedCloudNodeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DedicatedCloudNodeOutput)
}

func (i *DedicatedCloudNode) ToOutput(ctx context.Context) pulumix.Output[*DedicatedCloudNode] {
	return pulumix.Output[*DedicatedCloudNode]{
		OutputState: i.ToDedicatedCloudNodeOutputWithContext(ctx).OutputState,
	}
}

type DedicatedCloudNodeOutput struct{ *pulumi.OutputState }

func (DedicatedCloudNodeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DedicatedCloudNode)(nil)).Elem()
}

func (o DedicatedCloudNodeOutput) ToDedicatedCloudNodeOutput() DedicatedCloudNodeOutput {
	return o
}

func (o DedicatedCloudNodeOutput) ToDedicatedCloudNodeOutputWithContext(ctx context.Context) DedicatedCloudNodeOutput {
	return o
}

func (o DedicatedCloudNodeOutput) ToOutput(ctx context.Context) pulumix.Output[*DedicatedCloudNode] {
	return pulumix.Output[*DedicatedCloudNode]{
		OutputState: o.OutputState,
	}
}

// Availability Zone id, e.g. "az1"
func (o DedicatedCloudNodeOutput) AvailabilityZoneId() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedCloudNode) pulumi.StringOutput { return v.AvailabilityZoneId }).(pulumi.StringOutput)
}

// Availability Zone name, e.g. "Availability Zone 1"
func (o DedicatedCloudNodeOutput) AvailabilityZoneName() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedCloudNode) pulumi.StringOutput { return v.AvailabilityZoneName }).(pulumi.StringOutput)
}

// VMWare Cloud Rack Name
func (o DedicatedCloudNodeOutput) CloudRackName() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedCloudNode) pulumi.StringOutput { return v.CloudRackName }).(pulumi.StringOutput)
}

// date time the resource was created
func (o DedicatedCloudNodeOutput) Created() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedCloudNode) pulumi.StringOutput { return v.Created }).(pulumi.StringOutput)
}

// Azure region
func (o DedicatedCloudNodeOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedCloudNode) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// SKU's name
func (o DedicatedCloudNodeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedCloudNode) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// count of nodes to create
func (o DedicatedCloudNodeOutput) NodesCount() pulumi.IntOutput {
	return o.ApplyT(func(v *DedicatedCloudNode) pulumi.IntOutput { return v.NodesCount }).(pulumi.IntOutput)
}

// Placement Group id, e.g. "n1"
func (o DedicatedCloudNodeOutput) PlacementGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedCloudNode) pulumi.StringOutput { return v.PlacementGroupId }).(pulumi.StringOutput)
}

// Placement Name, e.g. "Placement Group 1"
func (o DedicatedCloudNodeOutput) PlacementGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedCloudNode) pulumi.StringOutput { return v.PlacementGroupName }).(pulumi.StringOutput)
}

// Private Cloud Id
func (o DedicatedCloudNodeOutput) PrivateCloudId() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedCloudNode) pulumi.StringOutput { return v.PrivateCloudId }).(pulumi.StringOutput)
}

// Resource Pool Name
func (o DedicatedCloudNodeOutput) PrivateCloudName() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedCloudNode) pulumi.StringOutput { return v.PrivateCloudName }).(pulumi.StringOutput)
}

// The provisioning status of the resource
func (o DedicatedCloudNodeOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedCloudNode) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// purchase id
func (o DedicatedCloudNodeOutput) PurchaseId() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedCloudNode) pulumi.StringOutput { return v.PurchaseId }).(pulumi.StringOutput)
}

// Dedicated Cloud Nodes SKU
func (o DedicatedCloudNodeOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *DedicatedCloudNode) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

// Node status, indicates is private cloud set up on this node or not
func (o DedicatedCloudNodeOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedCloudNode) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Dedicated Cloud Nodes tags
func (o DedicatedCloudNodeOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DedicatedCloudNode) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// {resourceProviderNamespace}/{resourceType}
func (o DedicatedCloudNodeOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedCloudNode) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// VMWare Cluster Name
func (o DedicatedCloudNodeOutput) VmwareClusterName() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedCloudNode) pulumi.StringOutput { return v.VmwareClusterName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DedicatedCloudNodeOutput{})
}
