// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package networkcloud

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Azure REST API version: 2023-05-01-preview. Prior API version in Azure Native 1.x: 2022-12-12-preview
type L2Network struct {
	pulumi.CustomResourceState

	// The list of resource IDs for the other Microsoft.NetworkCloud resources that have attached this network.
	AssociatedResourceIds pulumi.StringArrayOutput `pulumi:"associatedResourceIds"`
	// The resource ID of the Network Cloud cluster this L2 network is associated with.
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// The more detailed status of the L2 network.
	DetailedStatus pulumi.StringOutput `pulumi:"detailedStatus"`
	// The descriptive message about the current detailed status.
	DetailedStatusMessage pulumi.StringOutput `pulumi:"detailedStatusMessage"`
	// The extended location of the cluster associated with the resource.
	ExtendedLocation ExtendedLocationResponseOutput `pulumi:"extendedLocation"`
	// Field Deprecated. These fields will be empty/omitted. The list of Hybrid AKS cluster resource ID(s) that are associated with this L2 network.
	HybridAksClustersAssociatedIds pulumi.StringArrayOutput `pulumi:"hybridAksClustersAssociatedIds"`
	// Field Deprecated. The field was previously optional, now it will have no defined behavior and will be ignored. The network plugin type for Hybrid AKS.
	HybridAksPluginType pulumi.StringPtrOutput `pulumi:"hybridAksPluginType"`
	// The default interface name for this L2 network in the virtual machine. This name can be overridden by the name supplied in the network attachment configuration of that virtual machine.
	InterfaceName pulumi.StringPtrOutput `pulumi:"interfaceName"`
	// The resource ID of the Network Fabric l2IsolationDomain.
	L2IsolationDomainId pulumi.StringOutput `pulumi:"l2IsolationDomainId"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioning state of the L2 network.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// Field Deprecated. These fields will be empty/omitted. The list of virtual machine resource ID(s), excluding any Hybrid AKS virtual machines, that are currently using this L2 network.
	VirtualMachinesAssociatedIds pulumi.StringArrayOutput `pulumi:"virtualMachinesAssociatedIds"`
}

// NewL2Network registers a new resource with the given unique name, arguments, and options.
func NewL2Network(ctx *pulumi.Context,
	name string, args *L2NetworkArgs, opts ...pulumi.ResourceOption) (*L2Network, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ExtendedLocation == nil {
		return nil, errors.New("invalid value for required argument 'ExtendedLocation'")
	}
	if args.L2IsolationDomainId == nil {
		return nil, errors.New("invalid value for required argument 'L2IsolationDomainId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.HybridAksPluginType == nil {
		args.HybridAksPluginType = pulumi.StringPtr("SRIOV")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:networkcloud/v20221212preview:L2Network"),
		},
		{
			Type: pulumi.String("azure-native:networkcloud/v20230501preview:L2Network"),
		},
		{
			Type: pulumi.String("azure-native:networkcloud/v20230701:L2Network"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource L2Network
	err := ctx.RegisterResource("azure-native:networkcloud:L2Network", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetL2Network gets an existing L2Network resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetL2Network(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *L2NetworkState, opts ...pulumi.ResourceOption) (*L2Network, error) {
	var resource L2Network
	err := ctx.ReadResource("azure-native:networkcloud:L2Network", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering L2Network resources.
type l2networkState struct {
}

type L2NetworkState struct {
}

func (L2NetworkState) ElementType() reflect.Type {
	return reflect.TypeOf((*l2networkState)(nil)).Elem()
}

type l2networkArgs struct {
	// The extended location of the cluster associated with the resource.
	ExtendedLocation ExtendedLocation `pulumi:"extendedLocation"`
	// Field Deprecated. The field was previously optional, now it will have no defined behavior and will be ignored. The network plugin type for Hybrid AKS.
	HybridAksPluginType *string `pulumi:"hybridAksPluginType"`
	// The default interface name for this L2 network in the virtual machine. This name can be overridden by the name supplied in the network attachment configuration of that virtual machine.
	InterfaceName *string `pulumi:"interfaceName"`
	// The resource ID of the Network Fabric l2IsolationDomain.
	L2IsolationDomainId string `pulumi:"l2IsolationDomainId"`
	// The name of the L2 network.
	L2NetworkName *string `pulumi:"l2NetworkName"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a L2Network resource.
type L2NetworkArgs struct {
	// The extended location of the cluster associated with the resource.
	ExtendedLocation ExtendedLocationInput
	// Field Deprecated. The field was previously optional, now it will have no defined behavior and will be ignored. The network plugin type for Hybrid AKS.
	HybridAksPluginType pulumi.StringPtrInput
	// The default interface name for this L2 network in the virtual machine. This name can be overridden by the name supplied in the network attachment configuration of that virtual machine.
	InterfaceName pulumi.StringPtrInput
	// The resource ID of the Network Fabric l2IsolationDomain.
	L2IsolationDomainId pulumi.StringInput
	// The name of the L2 network.
	L2NetworkName pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (L2NetworkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*l2networkArgs)(nil)).Elem()
}

type L2NetworkInput interface {
	pulumi.Input

	ToL2NetworkOutput() L2NetworkOutput
	ToL2NetworkOutputWithContext(ctx context.Context) L2NetworkOutput
}

func (*L2Network) ElementType() reflect.Type {
	return reflect.TypeOf((**L2Network)(nil)).Elem()
}

func (i *L2Network) ToL2NetworkOutput() L2NetworkOutput {
	return i.ToL2NetworkOutputWithContext(context.Background())
}

func (i *L2Network) ToL2NetworkOutputWithContext(ctx context.Context) L2NetworkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(L2NetworkOutput)
}

func (i *L2Network) ToOutput(ctx context.Context) pulumix.Output[*L2Network] {
	return pulumix.Output[*L2Network]{
		OutputState: i.ToL2NetworkOutputWithContext(ctx).OutputState,
	}
}

type L2NetworkOutput struct{ *pulumi.OutputState }

func (L2NetworkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**L2Network)(nil)).Elem()
}

func (o L2NetworkOutput) ToL2NetworkOutput() L2NetworkOutput {
	return o
}

func (o L2NetworkOutput) ToL2NetworkOutputWithContext(ctx context.Context) L2NetworkOutput {
	return o
}

func (o L2NetworkOutput) ToOutput(ctx context.Context) pulumix.Output[*L2Network] {
	return pulumix.Output[*L2Network]{
		OutputState: o.OutputState,
	}
}

// The list of resource IDs for the other Microsoft.NetworkCloud resources that have attached this network.
func (o L2NetworkOutput) AssociatedResourceIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *L2Network) pulumi.StringArrayOutput { return v.AssociatedResourceIds }).(pulumi.StringArrayOutput)
}

// The resource ID of the Network Cloud cluster this L2 network is associated with.
func (o L2NetworkOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *L2Network) pulumi.StringOutput { return v.ClusterId }).(pulumi.StringOutput)
}

// The more detailed status of the L2 network.
func (o L2NetworkOutput) DetailedStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *L2Network) pulumi.StringOutput { return v.DetailedStatus }).(pulumi.StringOutput)
}

// The descriptive message about the current detailed status.
func (o L2NetworkOutput) DetailedStatusMessage() pulumi.StringOutput {
	return o.ApplyT(func(v *L2Network) pulumi.StringOutput { return v.DetailedStatusMessage }).(pulumi.StringOutput)
}

// The extended location of the cluster associated with the resource.
func (o L2NetworkOutput) ExtendedLocation() ExtendedLocationResponseOutput {
	return o.ApplyT(func(v *L2Network) ExtendedLocationResponseOutput { return v.ExtendedLocation }).(ExtendedLocationResponseOutput)
}

// Field Deprecated. These fields will be empty/omitted. The list of Hybrid AKS cluster resource ID(s) that are associated with this L2 network.
func (o L2NetworkOutput) HybridAksClustersAssociatedIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *L2Network) pulumi.StringArrayOutput { return v.HybridAksClustersAssociatedIds }).(pulumi.StringArrayOutput)
}

// Field Deprecated. The field was previously optional, now it will have no defined behavior and will be ignored. The network plugin type for Hybrid AKS.
func (o L2NetworkOutput) HybridAksPluginType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *L2Network) pulumi.StringPtrOutput { return v.HybridAksPluginType }).(pulumi.StringPtrOutput)
}

// The default interface name for this L2 network in the virtual machine. This name can be overridden by the name supplied in the network attachment configuration of that virtual machine.
func (o L2NetworkOutput) InterfaceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *L2Network) pulumi.StringPtrOutput { return v.InterfaceName }).(pulumi.StringPtrOutput)
}

// The resource ID of the Network Fabric l2IsolationDomain.
func (o L2NetworkOutput) L2IsolationDomainId() pulumi.StringOutput {
	return o.ApplyT(func(v *L2Network) pulumi.StringOutput { return v.L2IsolationDomainId }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o L2NetworkOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *L2Network) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o L2NetworkOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *L2Network) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the L2 network.
func (o L2NetworkOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *L2Network) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o L2NetworkOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *L2Network) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o L2NetworkOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *L2Network) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o L2NetworkOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *L2Network) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Field Deprecated. These fields will be empty/omitted. The list of virtual machine resource ID(s), excluding any Hybrid AKS virtual machines, that are currently using this L2 network.
func (o L2NetworkOutput) VirtualMachinesAssociatedIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *L2Network) pulumi.StringArrayOutput { return v.VirtualMachinesAssociatedIds }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(L2NetworkOutput{})
}
