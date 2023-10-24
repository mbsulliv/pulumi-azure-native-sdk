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

// Azure REST API version: 2023-05-01-preview. Prior API version in Azure Native 1.x: 2022-12-12-preview.
//
// Other available API versions: 2023-07-01.
type TrunkedNetwork struct {
	pulumi.CustomResourceState

	// The list of resource IDs for the other Microsoft.NetworkCloud resources that have attached this network.
	AssociatedResourceIds pulumi.StringArrayOutput `pulumi:"associatedResourceIds"`
	// The resource ID of the Network Cloud cluster this trunked network is associated with.
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// The more detailed status of the trunked network.
	DetailedStatus pulumi.StringOutput `pulumi:"detailedStatus"`
	// The descriptive message about the current detailed status.
	DetailedStatusMessage pulumi.StringOutput `pulumi:"detailedStatusMessage"`
	// The extended location of the cluster associated with the resource.
	ExtendedLocation ExtendedLocationResponseOutput `pulumi:"extendedLocation"`
	// Field Deprecated. These fields will be empty/omitted. The list of Hybrid AKS cluster resource IDs that are associated with this trunked network.
	HybridAksClustersAssociatedIds pulumi.StringArrayOutput `pulumi:"hybridAksClustersAssociatedIds"`
	// Field Deprecated. The field was previously optional, now it will have no defined behavior and will be ignored. The network plugin type for Hybrid AKS.
	HybridAksPluginType pulumi.StringPtrOutput `pulumi:"hybridAksPluginType"`
	// The default interface name for this trunked network in the virtual machine. This name can be overridden by the name supplied in the network attachment configuration of that virtual machine.
	InterfaceName pulumi.StringPtrOutput `pulumi:"interfaceName"`
	// The list of resource IDs representing the Network Fabric isolation domains. It can be any combination of l2IsolationDomain and l3IsolationDomain resources.
	IsolationDomainIds pulumi.StringArrayOutput `pulumi:"isolationDomainIds"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioning state of the trunked network.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// Field Deprecated. These fields will be empty/omitted. The list of virtual machine resource IDs, excluding any Hybrid AKS virtual machines, that are currently using this trunked network.
	VirtualMachinesAssociatedIds pulumi.StringArrayOutput `pulumi:"virtualMachinesAssociatedIds"`
	// The list of vlans that are selected from the isolation domains for trunking.
	Vlans pulumi.Float64ArrayOutput `pulumi:"vlans"`
}

// NewTrunkedNetwork registers a new resource with the given unique name, arguments, and options.
func NewTrunkedNetwork(ctx *pulumi.Context,
	name string, args *TrunkedNetworkArgs, opts ...pulumi.ResourceOption) (*TrunkedNetwork, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ExtendedLocation == nil {
		return nil, errors.New("invalid value for required argument 'ExtendedLocation'")
	}
	if args.IsolationDomainIds == nil {
		return nil, errors.New("invalid value for required argument 'IsolationDomainIds'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Vlans == nil {
		return nil, errors.New("invalid value for required argument 'Vlans'")
	}
	if args.HybridAksPluginType == nil {
		args.HybridAksPluginType = pulumi.StringPtr("SRIOV")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:networkcloud/v20221212preview:TrunkedNetwork"),
		},
		{
			Type: pulumi.String("azure-native:networkcloud/v20230501preview:TrunkedNetwork"),
		},
		{
			Type: pulumi.String("azure-native:networkcloud/v20230701:TrunkedNetwork"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource TrunkedNetwork
	err := ctx.RegisterResource("azure-native:networkcloud:TrunkedNetwork", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTrunkedNetwork gets an existing TrunkedNetwork resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTrunkedNetwork(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TrunkedNetworkState, opts ...pulumi.ResourceOption) (*TrunkedNetwork, error) {
	var resource TrunkedNetwork
	err := ctx.ReadResource("azure-native:networkcloud:TrunkedNetwork", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TrunkedNetwork resources.
type trunkedNetworkState struct {
}

type TrunkedNetworkState struct {
}

func (TrunkedNetworkState) ElementType() reflect.Type {
	return reflect.TypeOf((*trunkedNetworkState)(nil)).Elem()
}

type trunkedNetworkArgs struct {
	// The extended location of the cluster associated with the resource.
	ExtendedLocation ExtendedLocation `pulumi:"extendedLocation"`
	// Field Deprecated. The field was previously optional, now it will have no defined behavior and will be ignored. The network plugin type for Hybrid AKS.
	HybridAksPluginType *string `pulumi:"hybridAksPluginType"`
	// The default interface name for this trunked network in the virtual machine. This name can be overridden by the name supplied in the network attachment configuration of that virtual machine.
	InterfaceName *string `pulumi:"interfaceName"`
	// The list of resource IDs representing the Network Fabric isolation domains. It can be any combination of l2IsolationDomain and l3IsolationDomain resources.
	IsolationDomainIds []string `pulumi:"isolationDomainIds"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The name of the trunked network.
	TrunkedNetworkName *string `pulumi:"trunkedNetworkName"`
	// The list of vlans that are selected from the isolation domains for trunking.
	Vlans []float64 `pulumi:"vlans"`
}

// The set of arguments for constructing a TrunkedNetwork resource.
type TrunkedNetworkArgs struct {
	// The extended location of the cluster associated with the resource.
	ExtendedLocation ExtendedLocationInput
	// Field Deprecated. The field was previously optional, now it will have no defined behavior and will be ignored. The network plugin type for Hybrid AKS.
	HybridAksPluginType pulumi.StringPtrInput
	// The default interface name for this trunked network in the virtual machine. This name can be overridden by the name supplied in the network attachment configuration of that virtual machine.
	InterfaceName pulumi.StringPtrInput
	// The list of resource IDs representing the Network Fabric isolation domains. It can be any combination of l2IsolationDomain and l3IsolationDomain resources.
	IsolationDomainIds pulumi.StringArrayInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The name of the trunked network.
	TrunkedNetworkName pulumi.StringPtrInput
	// The list of vlans that are selected from the isolation domains for trunking.
	Vlans pulumi.Float64ArrayInput
}

func (TrunkedNetworkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*trunkedNetworkArgs)(nil)).Elem()
}

type TrunkedNetworkInput interface {
	pulumi.Input

	ToTrunkedNetworkOutput() TrunkedNetworkOutput
	ToTrunkedNetworkOutputWithContext(ctx context.Context) TrunkedNetworkOutput
}

func (*TrunkedNetwork) ElementType() reflect.Type {
	return reflect.TypeOf((**TrunkedNetwork)(nil)).Elem()
}

func (i *TrunkedNetwork) ToTrunkedNetworkOutput() TrunkedNetworkOutput {
	return i.ToTrunkedNetworkOutputWithContext(context.Background())
}

func (i *TrunkedNetwork) ToTrunkedNetworkOutputWithContext(ctx context.Context) TrunkedNetworkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrunkedNetworkOutput)
}

func (i *TrunkedNetwork) ToOutput(ctx context.Context) pulumix.Output[*TrunkedNetwork] {
	return pulumix.Output[*TrunkedNetwork]{
		OutputState: i.ToTrunkedNetworkOutputWithContext(ctx).OutputState,
	}
}

type TrunkedNetworkOutput struct{ *pulumi.OutputState }

func (TrunkedNetworkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TrunkedNetwork)(nil)).Elem()
}

func (o TrunkedNetworkOutput) ToTrunkedNetworkOutput() TrunkedNetworkOutput {
	return o
}

func (o TrunkedNetworkOutput) ToTrunkedNetworkOutputWithContext(ctx context.Context) TrunkedNetworkOutput {
	return o
}

func (o TrunkedNetworkOutput) ToOutput(ctx context.Context) pulumix.Output[*TrunkedNetwork] {
	return pulumix.Output[*TrunkedNetwork]{
		OutputState: o.OutputState,
	}
}

// The list of resource IDs for the other Microsoft.NetworkCloud resources that have attached this network.
func (o TrunkedNetworkOutput) AssociatedResourceIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *TrunkedNetwork) pulumi.StringArrayOutput { return v.AssociatedResourceIds }).(pulumi.StringArrayOutput)
}

// The resource ID of the Network Cloud cluster this trunked network is associated with.
func (o TrunkedNetworkOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *TrunkedNetwork) pulumi.StringOutput { return v.ClusterId }).(pulumi.StringOutput)
}

// The more detailed status of the trunked network.
func (o TrunkedNetworkOutput) DetailedStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *TrunkedNetwork) pulumi.StringOutput { return v.DetailedStatus }).(pulumi.StringOutput)
}

// The descriptive message about the current detailed status.
func (o TrunkedNetworkOutput) DetailedStatusMessage() pulumi.StringOutput {
	return o.ApplyT(func(v *TrunkedNetwork) pulumi.StringOutput { return v.DetailedStatusMessage }).(pulumi.StringOutput)
}

// The extended location of the cluster associated with the resource.
func (o TrunkedNetworkOutput) ExtendedLocation() ExtendedLocationResponseOutput {
	return o.ApplyT(func(v *TrunkedNetwork) ExtendedLocationResponseOutput { return v.ExtendedLocation }).(ExtendedLocationResponseOutput)
}

// Field Deprecated. These fields will be empty/omitted. The list of Hybrid AKS cluster resource IDs that are associated with this trunked network.
func (o TrunkedNetworkOutput) HybridAksClustersAssociatedIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *TrunkedNetwork) pulumi.StringArrayOutput { return v.HybridAksClustersAssociatedIds }).(pulumi.StringArrayOutput)
}

// Field Deprecated. The field was previously optional, now it will have no defined behavior and will be ignored. The network plugin type for Hybrid AKS.
func (o TrunkedNetworkOutput) HybridAksPluginType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TrunkedNetwork) pulumi.StringPtrOutput { return v.HybridAksPluginType }).(pulumi.StringPtrOutput)
}

// The default interface name for this trunked network in the virtual machine. This name can be overridden by the name supplied in the network attachment configuration of that virtual machine.
func (o TrunkedNetworkOutput) InterfaceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TrunkedNetwork) pulumi.StringPtrOutput { return v.InterfaceName }).(pulumi.StringPtrOutput)
}

// The list of resource IDs representing the Network Fabric isolation domains. It can be any combination of l2IsolationDomain and l3IsolationDomain resources.
func (o TrunkedNetworkOutput) IsolationDomainIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *TrunkedNetwork) pulumi.StringArrayOutput { return v.IsolationDomainIds }).(pulumi.StringArrayOutput)
}

// The geo-location where the resource lives
func (o TrunkedNetworkOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *TrunkedNetwork) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o TrunkedNetworkOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *TrunkedNetwork) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the trunked network.
func (o TrunkedNetworkOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *TrunkedNetwork) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o TrunkedNetworkOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *TrunkedNetwork) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o TrunkedNetworkOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *TrunkedNetwork) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o TrunkedNetworkOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *TrunkedNetwork) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Field Deprecated. These fields will be empty/omitted. The list of virtual machine resource IDs, excluding any Hybrid AKS virtual machines, that are currently using this trunked network.
func (o TrunkedNetworkOutput) VirtualMachinesAssociatedIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *TrunkedNetwork) pulumi.StringArrayOutput { return v.VirtualMachinesAssociatedIds }).(pulumi.StringArrayOutput)
}

// The list of vlans that are selected from the isolation domains for trunking.
func (o TrunkedNetworkOutput) Vlans() pulumi.Float64ArrayOutput {
	return o.ApplyT(func(v *TrunkedNetwork) pulumi.Float64ArrayOutput { return v.Vlans }).(pulumi.Float64ArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(TrunkedNetworkOutput{})
}
