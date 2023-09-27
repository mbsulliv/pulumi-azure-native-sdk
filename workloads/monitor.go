// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package workloads

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// SAP monitor info on Azure (ARM properties and SAP monitor properties)
// Azure REST API version: 2023-04-01. Prior API version in Azure Native 1.x: 2021-12-01-preview
type Monitor struct {
	pulumi.CustomResourceState

	// The SAP monitor resources will be deployed in the SAP monitoring region. The subnet region should be same as the SAP monitoring region.
	AppLocation pulumi.StringPtrOutput `pulumi:"appLocation"`
	// Defines the SAP monitor errors.
	Errors MonitorPropertiesResponseErrorsOutput `pulumi:"errors"`
	// [currently not in use] Managed service identity(user assigned identities)
	Identity UserAssignedServiceIdentityResponsePtrOutput `pulumi:"identity"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The ARM ID of the Log Analytics Workspace that is used for SAP monitoring.
	LogAnalyticsWorkspaceArmId pulumi.StringPtrOutput `pulumi:"logAnalyticsWorkspaceArmId"`
	// Managed resource group configuration
	ManagedResourceGroupConfiguration ManagedRGConfigurationResponsePtrOutput `pulumi:"managedResourceGroupConfiguration"`
	// The subnet which the SAP monitor will be deployed in
	MonitorSubnet pulumi.StringPtrOutput `pulumi:"monitorSubnet"`
	// The ARM ID of the MSI used for SAP monitoring.
	MsiArmId pulumi.StringOutput `pulumi:"msiArmId"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// State of provisioning of the SAP monitor.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Sets the routing preference of the SAP monitor. By default only RFC1918 traffic is routed to the customer VNET.
	RoutingPreference pulumi.StringPtrOutput `pulumi:"routingPreference"`
	// The ARM ID of the Storage account used for SAP monitoring.
	StorageAccountArmId pulumi.StringOutput `pulumi:"storageAccountArmId"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// Sets the preference for zone redundancy on resources created for the SAP monitor. By default resources will be created which do not support zone redundancy.
	ZoneRedundancyPreference pulumi.StringPtrOutput `pulumi:"zoneRedundancyPreference"`
}

// NewMonitor registers a new resource with the given unique name, arguments, and options.
func NewMonitor(ctx *pulumi.Context,
	name string, args *MonitorArgs, opts ...pulumi.ResourceOption) (*Monitor, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:workloads:monitor"),
		},
		{
			Type: pulumi.String("azure-native:workloads/v20211201preview:Monitor"),
		},
		{
			Type: pulumi.String("azure-native:workloads/v20211201preview:monitor"),
		},
		{
			Type: pulumi.String("azure-native:workloads/v20221101preview:Monitor"),
		},
		{
			Type: pulumi.String("azure-native:workloads/v20221101preview:monitor"),
		},
		{
			Type: pulumi.String("azure-native:workloads/v20230401:Monitor"),
		},
		{
			Type: pulumi.String("azure-native:workloads/v20230401:monitor"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Monitor
	err := ctx.RegisterResource("azure-native:workloads:Monitor", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMonitor gets an existing Monitor resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMonitor(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MonitorState, opts ...pulumi.ResourceOption) (*Monitor, error) {
	var resource Monitor
	err := ctx.ReadResource("azure-native:workloads:Monitor", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Monitor resources.
type monitorState struct {
}

type MonitorState struct {
}

func (MonitorState) ElementType() reflect.Type {
	return reflect.TypeOf((*monitorState)(nil)).Elem()
}

type monitorArgs struct {
	// The SAP monitor resources will be deployed in the SAP monitoring region. The subnet region should be same as the SAP monitoring region.
	AppLocation *string `pulumi:"appLocation"`
	// [currently not in use] Managed service identity(user assigned identities)
	Identity *UserAssignedServiceIdentity `pulumi:"identity"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The ARM ID of the Log Analytics Workspace that is used for SAP monitoring.
	LogAnalyticsWorkspaceArmId *string `pulumi:"logAnalyticsWorkspaceArmId"`
	// Managed resource group configuration
	ManagedResourceGroupConfiguration *ManagedRGConfiguration `pulumi:"managedResourceGroupConfiguration"`
	// Name of the SAP monitor resource.
	MonitorName *string `pulumi:"monitorName"`
	// The subnet which the SAP monitor will be deployed in
	MonitorSubnet *string `pulumi:"monitorSubnet"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Sets the routing preference of the SAP monitor. By default only RFC1918 traffic is routed to the customer VNET.
	RoutingPreference *string `pulumi:"routingPreference"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Sets the preference for zone redundancy on resources created for the SAP monitor. By default resources will be created which do not support zone redundancy.
	ZoneRedundancyPreference *string `pulumi:"zoneRedundancyPreference"`
}

// The set of arguments for constructing a Monitor resource.
type MonitorArgs struct {
	// The SAP monitor resources will be deployed in the SAP monitoring region. The subnet region should be same as the SAP monitoring region.
	AppLocation pulumi.StringPtrInput
	// [currently not in use] Managed service identity(user assigned identities)
	Identity UserAssignedServiceIdentityPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The ARM ID of the Log Analytics Workspace that is used for SAP monitoring.
	LogAnalyticsWorkspaceArmId pulumi.StringPtrInput
	// Managed resource group configuration
	ManagedResourceGroupConfiguration ManagedRGConfigurationPtrInput
	// Name of the SAP monitor resource.
	MonitorName pulumi.StringPtrInput
	// The subnet which the SAP monitor will be deployed in
	MonitorSubnet pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Sets the routing preference of the SAP monitor. By default only RFC1918 traffic is routed to the customer VNET.
	RoutingPreference pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Sets the preference for zone redundancy on resources created for the SAP monitor. By default resources will be created which do not support zone redundancy.
	ZoneRedundancyPreference pulumi.StringPtrInput
}

func (MonitorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*monitorArgs)(nil)).Elem()
}

type MonitorInput interface {
	pulumi.Input

	ToMonitorOutput() MonitorOutput
	ToMonitorOutputWithContext(ctx context.Context) MonitorOutput
}

func (*Monitor) ElementType() reflect.Type {
	return reflect.TypeOf((**Monitor)(nil)).Elem()
}

func (i *Monitor) ToMonitorOutput() MonitorOutput {
	return i.ToMonitorOutputWithContext(context.Background())
}

func (i *Monitor) ToMonitorOutputWithContext(ctx context.Context) MonitorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitorOutput)
}

func (i *Monitor) ToOutput(ctx context.Context) pulumix.Output[*Monitor] {
	return pulumix.Output[*Monitor]{
		OutputState: i.ToMonitorOutputWithContext(ctx).OutputState,
	}
}

type MonitorOutput struct{ *pulumi.OutputState }

func (MonitorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Monitor)(nil)).Elem()
}

func (o MonitorOutput) ToMonitorOutput() MonitorOutput {
	return o
}

func (o MonitorOutput) ToMonitorOutputWithContext(ctx context.Context) MonitorOutput {
	return o
}

func (o MonitorOutput) ToOutput(ctx context.Context) pulumix.Output[*Monitor] {
	return pulumix.Output[*Monitor]{
		OutputState: o.OutputState,
	}
}

// The SAP monitor resources will be deployed in the SAP monitoring region. The subnet region should be same as the SAP monitoring region.
func (o MonitorOutput) AppLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Monitor) pulumi.StringPtrOutput { return v.AppLocation }).(pulumi.StringPtrOutput)
}

// Defines the SAP monitor errors.
func (o MonitorOutput) Errors() MonitorPropertiesResponseErrorsOutput {
	return o.ApplyT(func(v *Monitor) MonitorPropertiesResponseErrorsOutput { return v.Errors }).(MonitorPropertiesResponseErrorsOutput)
}

// [currently not in use] Managed service identity(user assigned identities)
func (o MonitorOutput) Identity() UserAssignedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *Monitor) UserAssignedServiceIdentityResponsePtrOutput { return v.Identity }).(UserAssignedServiceIdentityResponsePtrOutput)
}

// The geo-location where the resource lives
func (o MonitorOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Monitor) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The ARM ID of the Log Analytics Workspace that is used for SAP monitoring.
func (o MonitorOutput) LogAnalyticsWorkspaceArmId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Monitor) pulumi.StringPtrOutput { return v.LogAnalyticsWorkspaceArmId }).(pulumi.StringPtrOutput)
}

// Managed resource group configuration
func (o MonitorOutput) ManagedResourceGroupConfiguration() ManagedRGConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *Monitor) ManagedRGConfigurationResponsePtrOutput { return v.ManagedResourceGroupConfiguration }).(ManagedRGConfigurationResponsePtrOutput)
}

// The subnet which the SAP monitor will be deployed in
func (o MonitorOutput) MonitorSubnet() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Monitor) pulumi.StringPtrOutput { return v.MonitorSubnet }).(pulumi.StringPtrOutput)
}

// The ARM ID of the MSI used for SAP monitoring.
func (o MonitorOutput) MsiArmId() pulumi.StringOutput {
	return o.ApplyT(func(v *Monitor) pulumi.StringOutput { return v.MsiArmId }).(pulumi.StringOutput)
}

// The name of the resource
func (o MonitorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Monitor) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// State of provisioning of the SAP monitor.
func (o MonitorOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Monitor) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Sets the routing preference of the SAP monitor. By default only RFC1918 traffic is routed to the customer VNET.
func (o MonitorOutput) RoutingPreference() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Monitor) pulumi.StringPtrOutput { return v.RoutingPreference }).(pulumi.StringPtrOutput)
}

// The ARM ID of the Storage account used for SAP monitoring.
func (o MonitorOutput) StorageAccountArmId() pulumi.StringOutput {
	return o.ApplyT(func(v *Monitor) pulumi.StringOutput { return v.StorageAccountArmId }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o MonitorOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Monitor) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o MonitorOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Monitor) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o MonitorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Monitor) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Sets the preference for zone redundancy on resources created for the SAP monitor. By default resources will be created which do not support zone redundancy.
func (o MonitorOutput) ZoneRedundancyPreference() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Monitor) pulumi.StringPtrOutput { return v.ZoneRedundancyPreference }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(MonitorOutput{})
}
