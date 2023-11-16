// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets information about the specified attached data network.
func LookupAttachedDataNetwork(ctx *pulumi.Context, args *LookupAttachedDataNetworkArgs, opts ...pulumi.InvokeOption) (*LookupAttachedDataNetworkResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupAttachedDataNetworkResult
	err := ctx.Invoke("azure-native:mobilenetwork/v20230601:getAttachedDataNetwork", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupAttachedDataNetworkArgs struct {
	// The name of the attached data network.
	AttachedDataNetworkName string `pulumi:"attachedDataNetworkName"`
	// The name of the packet core control plane.
	PacketCoreControlPlaneName string `pulumi:"packetCoreControlPlaneName"`
	// The name of the packet core data plane.
	PacketCoreDataPlaneName string `pulumi:"packetCoreDataPlaneName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Attached data network resource. Must be created in the same location as its parent packet core data plane.
type LookupAttachedDataNetworkResult struct {
	// The DNS servers to signal to UEs to use for this attached data network. This configuration is mandatory - if you don't want DNS servers, you must provide an empty array.
	DnsAddresses []string `pulumi:"dnsAddresses"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The network address and port translation (NAPT) configuration.
	// If this is not specified, the attached data network will use a default NAPT configuration with NAPT enabled.
	NaptConfiguration *NaptConfigurationResponse `pulumi:"naptConfiguration"`
	// The provisioning state of the attached data network resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// The user equipment (UE) address pool prefixes for the attached data network from which the packet core instance will dynamically assign IP addresses to UEs.
	// The packet core instance assigns an IP address to a UE when the UE sets up a PDU session.
	//  You must define at least one of userEquipmentAddressPoolPrefix and userEquipmentStaticAddressPoolPrefix. If you define both, they must be of the same size.
	UserEquipmentAddressPoolPrefix []string `pulumi:"userEquipmentAddressPoolPrefix"`
	// The user equipment (UE) address pool prefixes for the attached data network from which the packet core instance will assign static IP addresses to UEs.
	// The packet core instance assigns an IP address to a UE when the UE sets up a PDU session. The static IP address for a specific UE is set in StaticIPConfiguration on the corresponding SIM resource.
	// At least one of userEquipmentAddressPoolPrefix and userEquipmentStaticAddressPoolPrefix must be defined. If both are defined, they must be of the same size.
	UserEquipmentStaticAddressPoolPrefix []string `pulumi:"userEquipmentStaticAddressPoolPrefix"`
	// The user plane interface on the data network. For 5G networks, this is the N6 interface. For 4G networks, this is the SGi interface.
	UserPlaneDataInterface InterfacePropertiesResponse `pulumi:"userPlaneDataInterface"`
}

// Defaults sets the appropriate defaults for LookupAttachedDataNetworkResult
func (val *LookupAttachedDataNetworkResult) Defaults() *LookupAttachedDataNetworkResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.NaptConfiguration = tmp.NaptConfiguration.Defaults()

	return &tmp
}

func LookupAttachedDataNetworkOutput(ctx *pulumi.Context, args LookupAttachedDataNetworkOutputArgs, opts ...pulumi.InvokeOption) LookupAttachedDataNetworkResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAttachedDataNetworkResult, error) {
			args := v.(LookupAttachedDataNetworkArgs)
			r, err := LookupAttachedDataNetwork(ctx, &args, opts...)
			var s LookupAttachedDataNetworkResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAttachedDataNetworkResultOutput)
}

type LookupAttachedDataNetworkOutputArgs struct {
	// The name of the attached data network.
	AttachedDataNetworkName pulumi.StringInput `pulumi:"attachedDataNetworkName"`
	// The name of the packet core control plane.
	PacketCoreControlPlaneName pulumi.StringInput `pulumi:"packetCoreControlPlaneName"`
	// The name of the packet core data plane.
	PacketCoreDataPlaneName pulumi.StringInput `pulumi:"packetCoreDataPlaneName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupAttachedDataNetworkOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAttachedDataNetworkArgs)(nil)).Elem()
}

// Attached data network resource. Must be created in the same location as its parent packet core data plane.
type LookupAttachedDataNetworkResultOutput struct{ *pulumi.OutputState }

func (LookupAttachedDataNetworkResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAttachedDataNetworkResult)(nil)).Elem()
}

func (o LookupAttachedDataNetworkResultOutput) ToLookupAttachedDataNetworkResultOutput() LookupAttachedDataNetworkResultOutput {
	return o
}

func (o LookupAttachedDataNetworkResultOutput) ToLookupAttachedDataNetworkResultOutputWithContext(ctx context.Context) LookupAttachedDataNetworkResultOutput {
	return o
}

// The DNS servers to signal to UEs to use for this attached data network. This configuration is mandatory - if you don't want DNS servers, you must provide an empty array.
func (o LookupAttachedDataNetworkResultOutput) DnsAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAttachedDataNetworkResult) []string { return v.DnsAddresses }).(pulumi.StringArrayOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupAttachedDataNetworkResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAttachedDataNetworkResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupAttachedDataNetworkResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAttachedDataNetworkResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupAttachedDataNetworkResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAttachedDataNetworkResult) string { return v.Name }).(pulumi.StringOutput)
}

// The network address and port translation (NAPT) configuration.
// If this is not specified, the attached data network will use a default NAPT configuration with NAPT enabled.
func (o LookupAttachedDataNetworkResultOutput) NaptConfiguration() NaptConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupAttachedDataNetworkResult) *NaptConfigurationResponse { return v.NaptConfiguration }).(NaptConfigurationResponsePtrOutput)
}

// The provisioning state of the attached data network resource.
func (o LookupAttachedDataNetworkResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAttachedDataNetworkResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupAttachedDataNetworkResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupAttachedDataNetworkResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupAttachedDataNetworkResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupAttachedDataNetworkResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupAttachedDataNetworkResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAttachedDataNetworkResult) string { return v.Type }).(pulumi.StringOutput)
}

// The user equipment (UE) address pool prefixes for the attached data network from which the packet core instance will dynamically assign IP addresses to UEs.
// The packet core instance assigns an IP address to a UE when the UE sets up a PDU session.
//
//	You must define at least one of userEquipmentAddressPoolPrefix and userEquipmentStaticAddressPoolPrefix. If you define both, they must be of the same size.
func (o LookupAttachedDataNetworkResultOutput) UserEquipmentAddressPoolPrefix() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAttachedDataNetworkResult) []string { return v.UserEquipmentAddressPoolPrefix }).(pulumi.StringArrayOutput)
}

// The user equipment (UE) address pool prefixes for the attached data network from which the packet core instance will assign static IP addresses to UEs.
// The packet core instance assigns an IP address to a UE when the UE sets up a PDU session. The static IP address for a specific UE is set in StaticIPConfiguration on the corresponding SIM resource.
// At least one of userEquipmentAddressPoolPrefix and userEquipmentStaticAddressPoolPrefix must be defined. If both are defined, they must be of the same size.
func (o LookupAttachedDataNetworkResultOutput) UserEquipmentStaticAddressPoolPrefix() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAttachedDataNetworkResult) []string { return v.UserEquipmentStaticAddressPoolPrefix }).(pulumi.StringArrayOutput)
}

// The user plane interface on the data network. For 5G networks, this is the N6 interface. For 4G networks, this is the SGi interface.
func (o LookupAttachedDataNetworkResultOutput) UserPlaneDataInterface() InterfacePropertiesResponseOutput {
	return o.ApplyT(func(v LookupAttachedDataNetworkResult) InterfacePropertiesResponse { return v.UserPlaneDataInterface }).(InterfacePropertiesResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAttachedDataNetworkResultOutput{})
}
