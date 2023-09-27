// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230501preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

type StorageAppliance struct {
	pulumi.CustomResourceState

	// The credentials of the administrative interface on this storage appliance.
	AdministratorCredentials AdministrativeCredentialsResponseOutput `pulumi:"administratorCredentials"`
	// The total capacity of the storage appliance.
	Capacity pulumi.Float64Output `pulumi:"capacity"`
	// The amount of storage consumed.
	CapacityUsed pulumi.Float64Output `pulumi:"capacityUsed"`
	// The resource ID of the cluster this storage appliance is associated with.
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// The detailed status of the storage appliance.
	DetailedStatus pulumi.StringOutput `pulumi:"detailedStatus"`
	// The descriptive message about the current detailed status.
	DetailedStatusMessage pulumi.StringOutput `pulumi:"detailedStatusMessage"`
	// The extended location of the cluster associated with the resource.
	ExtendedLocation ExtendedLocationResponseOutput `pulumi:"extendedLocation"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The endpoint for the management interface of the storage appliance.
	ManagementIpv4Address pulumi.StringOutput `pulumi:"managementIpv4Address"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioning state of the storage appliance.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The resource ID of the rack where this storage appliance resides.
	RackId pulumi.StringOutput `pulumi:"rackId"`
	// The slot the storage appliance is in the rack based on the BOM configuration.
	RackSlot pulumi.Float64Output `pulumi:"rackSlot"`
	// The indicator of whether the storage appliance supports remote vendor management.
	RemoteVendorManagementFeature pulumi.StringOutput `pulumi:"remoteVendorManagementFeature"`
	// The indicator of whether the remote vendor management feature is enabled or disabled, or unsupported if it is an unsupported feature.
	RemoteVendorManagementStatus pulumi.StringOutput `pulumi:"remoteVendorManagementStatus"`
	// The serial number for the storage appliance.
	SerialNumber pulumi.StringOutput `pulumi:"serialNumber"`
	// The SKU for the storage appliance.
	StorageApplianceSkuId pulumi.StringOutput `pulumi:"storageApplianceSkuId"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewStorageAppliance registers a new resource with the given unique name, arguments, and options.
func NewStorageAppliance(ctx *pulumi.Context,
	name string, args *StorageApplianceArgs, opts ...pulumi.ResourceOption) (*StorageAppliance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AdministratorCredentials == nil {
		return nil, errors.New("invalid value for required argument 'AdministratorCredentials'")
	}
	if args.ExtendedLocation == nil {
		return nil, errors.New("invalid value for required argument 'ExtendedLocation'")
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
	if args.StorageApplianceSkuId == nil {
		return nil, errors.New("invalid value for required argument 'StorageApplianceSkuId'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:networkcloud:StorageAppliance"),
		},
		{
			Type: pulumi.String("azure-native:networkcloud/v20221212preview:StorageAppliance"),
		},
		{
			Type: pulumi.String("azure-native:networkcloud/v20230701:StorageAppliance"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource StorageAppliance
	err := ctx.RegisterResource("azure-native:networkcloud/v20230501preview:StorageAppliance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStorageAppliance gets an existing StorageAppliance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStorageAppliance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StorageApplianceState, opts ...pulumi.ResourceOption) (*StorageAppliance, error) {
	var resource StorageAppliance
	err := ctx.ReadResource("azure-native:networkcloud/v20230501preview:StorageAppliance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StorageAppliance resources.
type storageApplianceState struct {
}

type StorageApplianceState struct {
}

func (StorageApplianceState) ElementType() reflect.Type {
	return reflect.TypeOf((*storageApplianceState)(nil)).Elem()
}

type storageApplianceArgs struct {
	// The credentials of the administrative interface on this storage appliance.
	AdministratorCredentials AdministrativeCredentials `pulumi:"administratorCredentials"`
	// The extended location of the cluster associated with the resource.
	ExtendedLocation ExtendedLocation `pulumi:"extendedLocation"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The resource ID of the rack where this storage appliance resides.
	RackId string `pulumi:"rackId"`
	// The slot the storage appliance is in the rack based on the BOM configuration.
	RackSlot float64 `pulumi:"rackSlot"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The serial number for the storage appliance.
	SerialNumber string `pulumi:"serialNumber"`
	// The name of the storage appliance.
	StorageApplianceName *string `pulumi:"storageApplianceName"`
	// The SKU for the storage appliance.
	StorageApplianceSkuId string `pulumi:"storageApplianceSkuId"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a StorageAppliance resource.
type StorageApplianceArgs struct {
	// The credentials of the administrative interface on this storage appliance.
	AdministratorCredentials AdministrativeCredentialsInput
	// The extended location of the cluster associated with the resource.
	ExtendedLocation ExtendedLocationInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The resource ID of the rack where this storage appliance resides.
	RackId pulumi.StringInput
	// The slot the storage appliance is in the rack based on the BOM configuration.
	RackSlot pulumi.Float64Input
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The serial number for the storage appliance.
	SerialNumber pulumi.StringInput
	// The name of the storage appliance.
	StorageApplianceName pulumi.StringPtrInput
	// The SKU for the storage appliance.
	StorageApplianceSkuId pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (StorageApplianceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*storageApplianceArgs)(nil)).Elem()
}

type StorageApplianceInput interface {
	pulumi.Input

	ToStorageApplianceOutput() StorageApplianceOutput
	ToStorageApplianceOutputWithContext(ctx context.Context) StorageApplianceOutput
}

func (*StorageAppliance) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageAppliance)(nil)).Elem()
}

func (i *StorageAppliance) ToStorageApplianceOutput() StorageApplianceOutput {
	return i.ToStorageApplianceOutputWithContext(context.Background())
}

func (i *StorageAppliance) ToStorageApplianceOutputWithContext(ctx context.Context) StorageApplianceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageApplianceOutput)
}

func (i *StorageAppliance) ToOutput(ctx context.Context) pulumix.Output[*StorageAppliance] {
	return pulumix.Output[*StorageAppliance]{
		OutputState: i.ToStorageApplianceOutputWithContext(ctx).OutputState,
	}
}

type StorageApplianceOutput struct{ *pulumi.OutputState }

func (StorageApplianceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageAppliance)(nil)).Elem()
}

func (o StorageApplianceOutput) ToStorageApplianceOutput() StorageApplianceOutput {
	return o
}

func (o StorageApplianceOutput) ToStorageApplianceOutputWithContext(ctx context.Context) StorageApplianceOutput {
	return o
}

func (o StorageApplianceOutput) ToOutput(ctx context.Context) pulumix.Output[*StorageAppliance] {
	return pulumix.Output[*StorageAppliance]{
		OutputState: o.OutputState,
	}
}

// The credentials of the administrative interface on this storage appliance.
func (o StorageApplianceOutput) AdministratorCredentials() AdministrativeCredentialsResponseOutput {
	return o.ApplyT(func(v *StorageAppliance) AdministrativeCredentialsResponseOutput { return v.AdministratorCredentials }).(AdministrativeCredentialsResponseOutput)
}

// The total capacity of the storage appliance.
func (o StorageApplianceOutput) Capacity() pulumi.Float64Output {
	return o.ApplyT(func(v *StorageAppliance) pulumi.Float64Output { return v.Capacity }).(pulumi.Float64Output)
}

// The amount of storage consumed.
func (o StorageApplianceOutput) CapacityUsed() pulumi.Float64Output {
	return o.ApplyT(func(v *StorageAppliance) pulumi.Float64Output { return v.CapacityUsed }).(pulumi.Float64Output)
}

// The resource ID of the cluster this storage appliance is associated with.
func (o StorageApplianceOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageAppliance) pulumi.StringOutput { return v.ClusterId }).(pulumi.StringOutput)
}

// The detailed status of the storage appliance.
func (o StorageApplianceOutput) DetailedStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageAppliance) pulumi.StringOutput { return v.DetailedStatus }).(pulumi.StringOutput)
}

// The descriptive message about the current detailed status.
func (o StorageApplianceOutput) DetailedStatusMessage() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageAppliance) pulumi.StringOutput { return v.DetailedStatusMessage }).(pulumi.StringOutput)
}

// The extended location of the cluster associated with the resource.
func (o StorageApplianceOutput) ExtendedLocation() ExtendedLocationResponseOutput {
	return o.ApplyT(func(v *StorageAppliance) ExtendedLocationResponseOutput { return v.ExtendedLocation }).(ExtendedLocationResponseOutput)
}

// The geo-location where the resource lives
func (o StorageApplianceOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageAppliance) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The endpoint for the management interface of the storage appliance.
func (o StorageApplianceOutput) ManagementIpv4Address() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageAppliance) pulumi.StringOutput { return v.ManagementIpv4Address }).(pulumi.StringOutput)
}

// The name of the resource
func (o StorageApplianceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageAppliance) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the storage appliance.
func (o StorageApplianceOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageAppliance) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The resource ID of the rack where this storage appliance resides.
func (o StorageApplianceOutput) RackId() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageAppliance) pulumi.StringOutput { return v.RackId }).(pulumi.StringOutput)
}

// The slot the storage appliance is in the rack based on the BOM configuration.
func (o StorageApplianceOutput) RackSlot() pulumi.Float64Output {
	return o.ApplyT(func(v *StorageAppliance) pulumi.Float64Output { return v.RackSlot }).(pulumi.Float64Output)
}

// The indicator of whether the storage appliance supports remote vendor management.
func (o StorageApplianceOutput) RemoteVendorManagementFeature() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageAppliance) pulumi.StringOutput { return v.RemoteVendorManagementFeature }).(pulumi.StringOutput)
}

// The indicator of whether the remote vendor management feature is enabled or disabled, or unsupported if it is an unsupported feature.
func (o StorageApplianceOutput) RemoteVendorManagementStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageAppliance) pulumi.StringOutput { return v.RemoteVendorManagementStatus }).(pulumi.StringOutput)
}

// The serial number for the storage appliance.
func (o StorageApplianceOutput) SerialNumber() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageAppliance) pulumi.StringOutput { return v.SerialNumber }).(pulumi.StringOutput)
}

// The SKU for the storage appliance.
func (o StorageApplianceOutput) StorageApplianceSkuId() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageAppliance) pulumi.StringOutput { return v.StorageApplianceSkuId }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o StorageApplianceOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *StorageAppliance) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o StorageApplianceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *StorageAppliance) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o StorageApplianceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageAppliance) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(StorageApplianceOutput{})
}
