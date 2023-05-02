// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220901preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An device group resource belonging to a product resource.
type DeviceGroup struct {
	pulumi.CustomResourceState

	// Flag to define if the user allows for crash dump collection.
	AllowCrashDumpsCollection pulumi.StringPtrOutput `pulumi:"allowCrashDumpsCollection"`
	// Description of the device group.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Deployment status for the device group.
	HasDeployment pulumi.BoolOutput `pulumi:"hasDeployment"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Operating system feed type of the device group.
	OsFeedType pulumi.StringPtrOutput `pulumi:"osFeedType"`
	// The status of the last operation.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Regional data boundary for the device group.
	RegionalDataBoundary pulumi.StringPtrOutput `pulumi:"regionalDataBoundary"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// Update policy of the device group.
	UpdatePolicy pulumi.StringPtrOutput `pulumi:"updatePolicy"`
}

// NewDeviceGroup registers a new resource with the given unique name, arguments, and options.
func NewDeviceGroup(ctx *pulumi.Context,
	name string, args *DeviceGroupArgs, opts ...pulumi.ResourceOption) (*DeviceGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CatalogName == nil {
		return nil, errors.New("invalid value for required argument 'CatalogName'")
	}
	if args.ProductName == nil {
		return nil, errors.New("invalid value for required argument 'ProductName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:azuresphere:DeviceGroup"),
		},
	})
	opts = append(opts, aliases)
	var resource DeviceGroup
	err := ctx.RegisterResource("azure-native:azuresphere/v20220901preview:DeviceGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDeviceGroup gets an existing DeviceGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDeviceGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DeviceGroupState, opts ...pulumi.ResourceOption) (*DeviceGroup, error) {
	var resource DeviceGroup
	err := ctx.ReadResource("azure-native:azuresphere/v20220901preview:DeviceGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DeviceGroup resources.
type deviceGroupState struct {
}

type DeviceGroupState struct {
}

func (DeviceGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*deviceGroupState)(nil)).Elem()
}

type deviceGroupArgs struct {
	// Flag to define if the user allows for crash dump collection.
	AllowCrashDumpsCollection *string `pulumi:"allowCrashDumpsCollection"`
	// Name of catalog
	CatalogName string `pulumi:"catalogName"`
	// Description of the device group.
	Description *string `pulumi:"description"`
	// Name of device group.
	DeviceGroupName *string `pulumi:"deviceGroupName"`
	// Operating system feed type of the device group.
	OsFeedType *string `pulumi:"osFeedType"`
	// Name of product.
	ProductName string `pulumi:"productName"`
	// Regional data boundary for the device group.
	RegionalDataBoundary *string `pulumi:"regionalDataBoundary"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Update policy of the device group.
	UpdatePolicy *string `pulumi:"updatePolicy"`
}

// The set of arguments for constructing a DeviceGroup resource.
type DeviceGroupArgs struct {
	// Flag to define if the user allows for crash dump collection.
	AllowCrashDumpsCollection pulumi.StringPtrInput
	// Name of catalog
	CatalogName pulumi.StringInput
	// Description of the device group.
	Description pulumi.StringPtrInput
	// Name of device group.
	DeviceGroupName pulumi.StringPtrInput
	// Operating system feed type of the device group.
	OsFeedType pulumi.StringPtrInput
	// Name of product.
	ProductName pulumi.StringInput
	// Regional data boundary for the device group.
	RegionalDataBoundary pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Update policy of the device group.
	UpdatePolicy pulumi.StringPtrInput
}

func (DeviceGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*deviceGroupArgs)(nil)).Elem()
}

type DeviceGroupInput interface {
	pulumi.Input

	ToDeviceGroupOutput() DeviceGroupOutput
	ToDeviceGroupOutputWithContext(ctx context.Context) DeviceGroupOutput
}

func (*DeviceGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**DeviceGroup)(nil)).Elem()
}

func (i *DeviceGroup) ToDeviceGroupOutput() DeviceGroupOutput {
	return i.ToDeviceGroupOutputWithContext(context.Background())
}

func (i *DeviceGroup) ToDeviceGroupOutputWithContext(ctx context.Context) DeviceGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeviceGroupOutput)
}

type DeviceGroupOutput struct{ *pulumi.OutputState }

func (DeviceGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DeviceGroup)(nil)).Elem()
}

func (o DeviceGroupOutput) ToDeviceGroupOutput() DeviceGroupOutput {
	return o
}

func (o DeviceGroupOutput) ToDeviceGroupOutputWithContext(ctx context.Context) DeviceGroupOutput {
	return o
}

// Flag to define if the user allows for crash dump collection.
func (o DeviceGroupOutput) AllowCrashDumpsCollection() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeviceGroup) pulumi.StringPtrOutput { return v.AllowCrashDumpsCollection }).(pulumi.StringPtrOutput)
}

// Description of the device group.
func (o DeviceGroupOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeviceGroup) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Deployment status for the device group.
func (o DeviceGroupOutput) HasDeployment() pulumi.BoolOutput {
	return o.ApplyT(func(v *DeviceGroup) pulumi.BoolOutput { return v.HasDeployment }).(pulumi.BoolOutput)
}

// The name of the resource
func (o DeviceGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DeviceGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Operating system feed type of the device group.
func (o DeviceGroupOutput) OsFeedType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeviceGroup) pulumi.StringPtrOutput { return v.OsFeedType }).(pulumi.StringPtrOutput)
}

// The status of the last operation.
func (o DeviceGroupOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *DeviceGroup) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Regional data boundary for the device group.
func (o DeviceGroupOutput) RegionalDataBoundary() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeviceGroup) pulumi.StringPtrOutput { return v.RegionalDataBoundary }).(pulumi.StringPtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o DeviceGroupOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *DeviceGroup) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o DeviceGroupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DeviceGroup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Update policy of the device group.
func (o DeviceGroupOutput) UpdatePolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeviceGroup) pulumi.StringPtrOutput { return v.UpdatePolicy }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(DeviceGroupOutput{})
}