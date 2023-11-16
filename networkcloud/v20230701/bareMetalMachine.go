// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230701

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type BareMetalMachine struct {
	pulumi.CustomResourceState

	// The list of resource IDs for the other Microsoft.NetworkCloud resources that have attached this network.
	AssociatedResourceIds pulumi.StringArrayOutput `pulumi:"associatedResourceIds"`
	// The connection string for the baseboard management controller including IP address and protocol.
	BmcConnectionString pulumi.StringOutput `pulumi:"bmcConnectionString"`
	// The credentials of the baseboard management controller on this bare metal machine.
	BmcCredentials AdministrativeCredentialsResponseOutput `pulumi:"bmcCredentials"`
	// The MAC address of the BMC device.
	BmcMacAddress pulumi.StringOutput `pulumi:"bmcMacAddress"`
	// The MAC address of a NIC connected to the PXE network.
	BootMacAddress pulumi.StringOutput `pulumi:"bootMacAddress"`
	// The resource ID of the cluster this bare metal machine is associated with.
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// The cordon status of the bare metal machine.
	CordonStatus pulumi.StringOutput `pulumi:"cordonStatus"`
	// The more detailed status of the bare metal machine.
	DetailedStatus pulumi.StringOutput `pulumi:"detailedStatus"`
	// The descriptive message about the current detailed status.
	DetailedStatusMessage pulumi.StringOutput `pulumi:"detailedStatusMessage"`
	// The extended location of the cluster associated with the resource.
	ExtendedLocation ExtendedLocationResponseOutput `pulumi:"extendedLocation"`
	// The hardware inventory, including information acquired from the model/sku information and from the ironic inspector.
	HardwareInventory HardwareInventoryResponseOutput `pulumi:"hardwareInventory"`
	// The details of the latest hardware validation performed for this bare metal machine.
	HardwareValidationStatus HardwareValidationStatusResponseOutput `pulumi:"hardwareValidationStatus"`
	// Field Deprecated. These fields will be empty/omitted. The list of the resource IDs for the HybridAksClusters that have nodes hosted on this bare metal machine.
	HybridAksClustersAssociatedIds pulumi.StringArrayOutput `pulumi:"hybridAksClustersAssociatedIds"`
	// The name of this machine represented by the host object in the Cluster's Kubernetes control plane.
	KubernetesNodeName pulumi.StringOutput `pulumi:"kubernetesNodeName"`
	// The version of Kubernetes running on this machine.
	KubernetesVersion pulumi.StringOutput `pulumi:"kubernetesVersion"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The custom details provided by the customer.
	MachineDetails pulumi.StringOutput `pulumi:"machineDetails"`
	// The OS-level hostname assigned to this machine.
	MachineName pulumi.StringOutput `pulumi:"machineName"`
	// The unique internal identifier of the bare metal machine SKU.
	MachineSkuId pulumi.StringOutput `pulumi:"machineSkuId"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The IPv4 address that is assigned to the bare metal machine during the cluster deployment.
	OamIpv4Address pulumi.StringOutput `pulumi:"oamIpv4Address"`
	// The IPv6 address that is assigned to the bare metal machine during the cluster deployment.
	OamIpv6Address pulumi.StringOutput `pulumi:"oamIpv6Address"`
	// The image that is currently provisioned to the OS disk.
	OsImage pulumi.StringOutput `pulumi:"osImage"`
	// The power state derived from the baseboard management controller.
	PowerState pulumi.StringOutput `pulumi:"powerState"`
	// The provisioning state of the bare metal machine.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The resource ID of the rack where this bare metal machine resides.
	RackId pulumi.StringOutput `pulumi:"rackId"`
	// The rack slot in which this bare metal machine is located, ordered from the bottom up i.e. the lowest slot is 1.
	RackSlot pulumi.Float64Output `pulumi:"rackSlot"`
	// The indicator of whether the bare metal machine is ready to receive workloads.
	ReadyState pulumi.StringOutput `pulumi:"readyState"`
	// The serial number of the bare metal machine.
	SerialNumber pulumi.StringOutput `pulumi:"serialNumber"`
	// The discovered value of the machine's service tag.
	ServiceTag pulumi.StringOutput `pulumi:"serviceTag"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// Field Deprecated. These fields will be empty/omitted. The list of the resource IDs for the VirtualMachines that are hosted on this bare metal machine.
	VirtualMachinesAssociatedIds pulumi.StringArrayOutput `pulumi:"virtualMachinesAssociatedIds"`
}

// NewBareMetalMachine registers a new resource with the given unique name, arguments, and options.
func NewBareMetalMachine(ctx *pulumi.Context,
	name string, args *BareMetalMachineArgs, opts ...pulumi.ResourceOption) (*BareMetalMachine, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BmcConnectionString == nil {
		return nil, errors.New("invalid value for required argument 'BmcConnectionString'")
	}
	if args.BmcCredentials == nil {
		return nil, errors.New("invalid value for required argument 'BmcCredentials'")
	}
	if args.BmcMacAddress == nil {
		return nil, errors.New("invalid value for required argument 'BmcMacAddress'")
	}
	if args.BootMacAddress == nil {
		return nil, errors.New("invalid value for required argument 'BootMacAddress'")
	}
	if args.ExtendedLocation == nil {
		return nil, errors.New("invalid value for required argument 'ExtendedLocation'")
	}
	if args.MachineDetails == nil {
		return nil, errors.New("invalid value for required argument 'MachineDetails'")
	}
	if args.MachineName == nil {
		return nil, errors.New("invalid value for required argument 'MachineName'")
	}
	if args.MachineSkuId == nil {
		return nil, errors.New("invalid value for required argument 'MachineSkuId'")
	}
	if args.RackId == nil {
		return nil, errors.New("invalid value for required argument 'RackId'")
	}
	if args.RackSlot == nil {
		return nil, errors.New("invalid value for required argument 'RackSlot'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SerialNumber == nil {
		return nil, errors.New("invalid value for required argument 'SerialNumber'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:networkcloud:BareMetalMachine"),
		},
		{
			Type: pulumi.String("azure-native:networkcloud/v20221212preview:BareMetalMachine"),
		},
		{
			Type: pulumi.String("azure-native:networkcloud/v20230501preview:BareMetalMachine"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource BareMetalMachine
	err := ctx.RegisterResource("azure-native:networkcloud/v20230701:BareMetalMachine", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBareMetalMachine gets an existing BareMetalMachine resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBareMetalMachine(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BareMetalMachineState, opts ...pulumi.ResourceOption) (*BareMetalMachine, error) {
	var resource BareMetalMachine
	err := ctx.ReadResource("azure-native:networkcloud/v20230701:BareMetalMachine", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BareMetalMachine resources.
type bareMetalMachineState struct {
}

type BareMetalMachineState struct {
}

func (BareMetalMachineState) ElementType() reflect.Type {
	return reflect.TypeOf((*bareMetalMachineState)(nil)).Elem()
}

type bareMetalMachineArgs struct {
	// The name of the bare metal machine.
	BareMetalMachineName *string `pulumi:"bareMetalMachineName"`
	// The connection string for the baseboard management controller including IP address and protocol.
	BmcConnectionString string `pulumi:"bmcConnectionString"`
	// The credentials of the baseboard management controller on this bare metal machine.
	BmcCredentials AdministrativeCredentials `pulumi:"bmcCredentials"`
	// The MAC address of the BMC device.
	BmcMacAddress string `pulumi:"bmcMacAddress"`
	// The MAC address of a NIC connected to the PXE network.
	BootMacAddress string `pulumi:"bootMacAddress"`
	// The extended location of the cluster associated with the resource.
	ExtendedLocation ExtendedLocation `pulumi:"extendedLocation"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The custom details provided by the customer.
	MachineDetails string `pulumi:"machineDetails"`
	// The OS-level hostname assigned to this machine.
	MachineName string `pulumi:"machineName"`
	// The unique internal identifier of the bare metal machine SKU.
	MachineSkuId string `pulumi:"machineSkuId"`
	// The resource ID of the rack where this bare metal machine resides.
	RackId string `pulumi:"rackId"`
	// The rack slot in which this bare metal machine is located, ordered from the bottom up i.e. the lowest slot is 1.
	RackSlot float64 `pulumi:"rackSlot"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The serial number of the bare metal machine.
	SerialNumber string `pulumi:"serialNumber"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a BareMetalMachine resource.
type BareMetalMachineArgs struct {
	// The name of the bare metal machine.
	BareMetalMachineName pulumi.StringPtrInput
	// The connection string for the baseboard management controller including IP address and protocol.
	BmcConnectionString pulumi.StringInput
	// The credentials of the baseboard management controller on this bare metal machine.
	BmcCredentials AdministrativeCredentialsInput
	// The MAC address of the BMC device.
	BmcMacAddress pulumi.StringInput
	// The MAC address of a NIC connected to the PXE network.
	BootMacAddress pulumi.StringInput
	// The extended location of the cluster associated with the resource.
	ExtendedLocation ExtendedLocationInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The custom details provided by the customer.
	MachineDetails pulumi.StringInput
	// The OS-level hostname assigned to this machine.
	MachineName pulumi.StringInput
	// The unique internal identifier of the bare metal machine SKU.
	MachineSkuId pulumi.StringInput
	// The resource ID of the rack where this bare metal machine resides.
	RackId pulumi.StringInput
	// The rack slot in which this bare metal machine is located, ordered from the bottom up i.e. the lowest slot is 1.
	RackSlot pulumi.Float64Input
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The serial number of the bare metal machine.
	SerialNumber pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (BareMetalMachineArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bareMetalMachineArgs)(nil)).Elem()
}

type BareMetalMachineInput interface {
	pulumi.Input

	ToBareMetalMachineOutput() BareMetalMachineOutput
	ToBareMetalMachineOutputWithContext(ctx context.Context) BareMetalMachineOutput
}

func (*BareMetalMachine) ElementType() reflect.Type {
	return reflect.TypeOf((**BareMetalMachine)(nil)).Elem()
}

func (i *BareMetalMachine) ToBareMetalMachineOutput() BareMetalMachineOutput {
	return i.ToBareMetalMachineOutputWithContext(context.Background())
}

func (i *BareMetalMachine) ToBareMetalMachineOutputWithContext(ctx context.Context) BareMetalMachineOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BareMetalMachineOutput)
}

type BareMetalMachineOutput struct{ *pulumi.OutputState }

func (BareMetalMachineOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BareMetalMachine)(nil)).Elem()
}

func (o BareMetalMachineOutput) ToBareMetalMachineOutput() BareMetalMachineOutput {
	return o
}

func (o BareMetalMachineOutput) ToBareMetalMachineOutputWithContext(ctx context.Context) BareMetalMachineOutput {
	return o
}

// The list of resource IDs for the other Microsoft.NetworkCloud resources that have attached this network.
func (o BareMetalMachineOutput) AssociatedResourceIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *BareMetalMachine) pulumi.StringArrayOutput { return v.AssociatedResourceIds }).(pulumi.StringArrayOutput)
}

// The connection string for the baseboard management controller including IP address and protocol.
func (o BareMetalMachineOutput) BmcConnectionString() pulumi.StringOutput {
	return o.ApplyT(func(v *BareMetalMachine) pulumi.StringOutput { return v.BmcConnectionString }).(pulumi.StringOutput)
}

// The credentials of the baseboard management controller on this bare metal machine.
func (o BareMetalMachineOutput) BmcCredentials() AdministrativeCredentialsResponseOutput {
	return o.ApplyT(func(v *BareMetalMachine) AdministrativeCredentialsResponseOutput { return v.BmcCredentials }).(AdministrativeCredentialsResponseOutput)
}

// The MAC address of the BMC device.
func (o BareMetalMachineOutput) BmcMacAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *BareMetalMachine) pulumi.StringOutput { return v.BmcMacAddress }).(pulumi.StringOutput)
}

// The MAC address of a NIC connected to the PXE network.
func (o BareMetalMachineOutput) BootMacAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *BareMetalMachine) pulumi.StringOutput { return v.BootMacAddress }).(pulumi.StringOutput)
}

// The resource ID of the cluster this bare metal machine is associated with.
func (o BareMetalMachineOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *BareMetalMachine) pulumi.StringOutput { return v.ClusterId }).(pulumi.StringOutput)
}

// The cordon status of the bare metal machine.
func (o BareMetalMachineOutput) CordonStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *BareMetalMachine) pulumi.StringOutput { return v.CordonStatus }).(pulumi.StringOutput)
}

// The more detailed status of the bare metal machine.
func (o BareMetalMachineOutput) DetailedStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *BareMetalMachine) pulumi.StringOutput { return v.DetailedStatus }).(pulumi.StringOutput)
}

// The descriptive message about the current detailed status.
func (o BareMetalMachineOutput) DetailedStatusMessage() pulumi.StringOutput {
	return o.ApplyT(func(v *BareMetalMachine) pulumi.StringOutput { return v.DetailedStatusMessage }).(pulumi.StringOutput)
}

// The extended location of the cluster associated with the resource.
func (o BareMetalMachineOutput) ExtendedLocation() ExtendedLocationResponseOutput {
	return o.ApplyT(func(v *BareMetalMachine) ExtendedLocationResponseOutput { return v.ExtendedLocation }).(ExtendedLocationResponseOutput)
}

// The hardware inventory, including information acquired from the model/sku information and from the ironic inspector.
func (o BareMetalMachineOutput) HardwareInventory() HardwareInventoryResponseOutput {
	return o.ApplyT(func(v *BareMetalMachine) HardwareInventoryResponseOutput { return v.HardwareInventory }).(HardwareInventoryResponseOutput)
}

// The details of the latest hardware validation performed for this bare metal machine.
func (o BareMetalMachineOutput) HardwareValidationStatus() HardwareValidationStatusResponseOutput {
	return o.ApplyT(func(v *BareMetalMachine) HardwareValidationStatusResponseOutput { return v.HardwareValidationStatus }).(HardwareValidationStatusResponseOutput)
}

// Field Deprecated. These fields will be empty/omitted. The list of the resource IDs for the HybridAksClusters that have nodes hosted on this bare metal machine.
func (o BareMetalMachineOutput) HybridAksClustersAssociatedIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *BareMetalMachine) pulumi.StringArrayOutput { return v.HybridAksClustersAssociatedIds }).(pulumi.StringArrayOutput)
}

// The name of this machine represented by the host object in the Cluster's Kubernetes control plane.
func (o BareMetalMachineOutput) KubernetesNodeName() pulumi.StringOutput {
	return o.ApplyT(func(v *BareMetalMachine) pulumi.StringOutput { return v.KubernetesNodeName }).(pulumi.StringOutput)
}

// The version of Kubernetes running on this machine.
func (o BareMetalMachineOutput) KubernetesVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *BareMetalMachine) pulumi.StringOutput { return v.KubernetesVersion }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o BareMetalMachineOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *BareMetalMachine) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The custom details provided by the customer.
func (o BareMetalMachineOutput) MachineDetails() pulumi.StringOutput {
	return o.ApplyT(func(v *BareMetalMachine) pulumi.StringOutput { return v.MachineDetails }).(pulumi.StringOutput)
}

// The OS-level hostname assigned to this machine.
func (o BareMetalMachineOutput) MachineName() pulumi.StringOutput {
	return o.ApplyT(func(v *BareMetalMachine) pulumi.StringOutput { return v.MachineName }).(pulumi.StringOutput)
}

// The unique internal identifier of the bare metal machine SKU.
func (o BareMetalMachineOutput) MachineSkuId() pulumi.StringOutput {
	return o.ApplyT(func(v *BareMetalMachine) pulumi.StringOutput { return v.MachineSkuId }).(pulumi.StringOutput)
}

// The name of the resource
func (o BareMetalMachineOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *BareMetalMachine) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The IPv4 address that is assigned to the bare metal machine during the cluster deployment.
func (o BareMetalMachineOutput) OamIpv4Address() pulumi.StringOutput {
	return o.ApplyT(func(v *BareMetalMachine) pulumi.StringOutput { return v.OamIpv4Address }).(pulumi.StringOutput)
}

// The IPv6 address that is assigned to the bare metal machine during the cluster deployment.
func (o BareMetalMachineOutput) OamIpv6Address() pulumi.StringOutput {
	return o.ApplyT(func(v *BareMetalMachine) pulumi.StringOutput { return v.OamIpv6Address }).(pulumi.StringOutput)
}

// The image that is currently provisioned to the OS disk.
func (o BareMetalMachineOutput) OsImage() pulumi.StringOutput {
	return o.ApplyT(func(v *BareMetalMachine) pulumi.StringOutput { return v.OsImage }).(pulumi.StringOutput)
}

// The power state derived from the baseboard management controller.
func (o BareMetalMachineOutput) PowerState() pulumi.StringOutput {
	return o.ApplyT(func(v *BareMetalMachine) pulumi.StringOutput { return v.PowerState }).(pulumi.StringOutput)
}

// The provisioning state of the bare metal machine.
func (o BareMetalMachineOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *BareMetalMachine) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The resource ID of the rack where this bare metal machine resides.
func (o BareMetalMachineOutput) RackId() pulumi.StringOutput {
	return o.ApplyT(func(v *BareMetalMachine) pulumi.StringOutput { return v.RackId }).(pulumi.StringOutput)
}

// The rack slot in which this bare metal machine is located, ordered from the bottom up i.e. the lowest slot is 1.
func (o BareMetalMachineOutput) RackSlot() pulumi.Float64Output {
	return o.ApplyT(func(v *BareMetalMachine) pulumi.Float64Output { return v.RackSlot }).(pulumi.Float64Output)
}

// The indicator of whether the bare metal machine is ready to receive workloads.
func (o BareMetalMachineOutput) ReadyState() pulumi.StringOutput {
	return o.ApplyT(func(v *BareMetalMachine) pulumi.StringOutput { return v.ReadyState }).(pulumi.StringOutput)
}

// The serial number of the bare metal machine.
func (o BareMetalMachineOutput) SerialNumber() pulumi.StringOutput {
	return o.ApplyT(func(v *BareMetalMachine) pulumi.StringOutput { return v.SerialNumber }).(pulumi.StringOutput)
}

// The discovered value of the machine's service tag.
func (o BareMetalMachineOutput) ServiceTag() pulumi.StringOutput {
	return o.ApplyT(func(v *BareMetalMachine) pulumi.StringOutput { return v.ServiceTag }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o BareMetalMachineOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *BareMetalMachine) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o BareMetalMachineOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *BareMetalMachine) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o BareMetalMachineOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *BareMetalMachine) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Field Deprecated. These fields will be empty/omitted. The list of the resource IDs for the VirtualMachines that are hosted on this bare metal machine.
func (o BareMetalMachineOutput) VirtualMachinesAssociatedIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *BareMetalMachine) pulumi.StringArrayOutput { return v.VirtualMachinesAssociatedIds }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(BareMetalMachineOutput{})
}
