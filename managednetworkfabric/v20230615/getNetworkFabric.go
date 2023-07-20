// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230615

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get Network Fabric resource details.
func LookupNetworkFabric(ctx *pulumi.Context, args *LookupNetworkFabricArgs, opts ...pulumi.InvokeOption) (*LookupNetworkFabricResult, error) {
	var rv LookupNetworkFabricResult
	err := ctx.Invoke("azure-native:managednetworkfabric/v20230615:getNetworkFabric", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupNetworkFabricArgs struct {
	// Name of the Network Fabric.
	NetworkFabricName string `pulumi:"networkFabricName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The Network Fabric resource definition.
type LookupNetworkFabricResult struct {
	// Administrative state of the resource.
	AdministrativeState string `pulumi:"administrativeState"`
	// Switch configuration description.
	Annotation *string `pulumi:"annotation"`
	// Configuration state of the resource.
	ConfigurationState string `pulumi:"configurationState"`
	// ASN of CE devices for CE/PE connectivity.
	FabricASN float64 `pulumi:"fabricASN"`
	// The version of Network Fabric.
	FabricVersion string `pulumi:"fabricVersion"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// IPv4Prefix for Management Network. Example: 10.1.0.0/19.
	Ipv4Prefix string `pulumi:"ipv4Prefix"`
	// IPv6Prefix for Management Network. Example: 3FFE:FFFF:0:CD40::/59
	Ipv6Prefix *string `pulumi:"ipv6Prefix"`
	// List of L2 Isolation Domain resource IDs under the Network Fabric.
	L2IsolationDomains []string `pulumi:"l2IsolationDomains"`
	// List of L3 Isolation Domain resource IDs under the Network Fabric.
	L3IsolationDomains []string `pulumi:"l3IsolationDomains"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// Configuration to be used to setup the management network.
	ManagementNetworkConfiguration ManagementNetworkConfigurationPropertiesResponse `pulumi:"managementNetworkConfiguration"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Azure resource ID for the NetworkFabricController the NetworkFabric belongs.
	NetworkFabricControllerId string `pulumi:"networkFabricControllerId"`
	// Supported Network Fabric SKU.Example: Compute / Aggregate racks. Once the user chooses a particular SKU, only supported racks can be added to the Network Fabric. The SKU determines whether it is a single / multi rack Network Fabric.
	NetworkFabricSku string `pulumi:"networkFabricSku"`
	// Provides you the latest status of the NFC service, whether it is Accepted, updating, Succeeded or Failed. During this process, the states keep changing based on the status of NFC provisioning.
	ProvisioningState string `pulumi:"provisioningState"`
	// Number of compute racks associated to Network Fabric.
	RackCount *int `pulumi:"rackCount"`
	// List of NetworkRack resource IDs under the Network Fabric. The number of racks allowed depends on the Network Fabric SKU.
	Racks []string `pulumi:"racks"`
	// Array of router IDs.
	RouterIds []string `pulumi:"routerIds"`
	// Number of servers.Possible values are from 1-16.
	ServerCountPerRack int `pulumi:"serverCountPerRack"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Network and credentials configuration currently applied to terminal server.
	TerminalServerConfiguration TerminalServerConfigurationResponse `pulumi:"terminalServerConfiguration"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupNetworkFabricResult
func (val *LookupNetworkFabricResult) Defaults() *LookupNetworkFabricResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.ManagementNetworkConfiguration = *tmp.ManagementNetworkConfiguration.Defaults()

	return &tmp
}

func LookupNetworkFabricOutput(ctx *pulumi.Context, args LookupNetworkFabricOutputArgs, opts ...pulumi.InvokeOption) LookupNetworkFabricResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNetworkFabricResult, error) {
			args := v.(LookupNetworkFabricArgs)
			r, err := LookupNetworkFabric(ctx, &args, opts...)
			var s LookupNetworkFabricResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNetworkFabricResultOutput)
}

type LookupNetworkFabricOutputArgs struct {
	// Name of the Network Fabric.
	NetworkFabricName pulumi.StringInput `pulumi:"networkFabricName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupNetworkFabricOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkFabricArgs)(nil)).Elem()
}

// The Network Fabric resource definition.
type LookupNetworkFabricResultOutput struct{ *pulumi.OutputState }

func (LookupNetworkFabricResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkFabricResult)(nil)).Elem()
}

func (o LookupNetworkFabricResultOutput) ToLookupNetworkFabricResultOutput() LookupNetworkFabricResultOutput {
	return o
}

func (o LookupNetworkFabricResultOutput) ToLookupNetworkFabricResultOutputWithContext(ctx context.Context) LookupNetworkFabricResultOutput {
	return o
}

// Administrative state of the resource.
func (o LookupNetworkFabricResultOutput) AdministrativeState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkFabricResult) string { return v.AdministrativeState }).(pulumi.StringOutput)
}

// Switch configuration description.
func (o LookupNetworkFabricResultOutput) Annotation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkFabricResult) *string { return v.Annotation }).(pulumi.StringPtrOutput)
}

// Configuration state of the resource.
func (o LookupNetworkFabricResultOutput) ConfigurationState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkFabricResult) string { return v.ConfigurationState }).(pulumi.StringOutput)
}

// ASN of CE devices for CE/PE connectivity.
func (o LookupNetworkFabricResultOutput) FabricASN() pulumi.Float64Output {
	return o.ApplyT(func(v LookupNetworkFabricResult) float64 { return v.FabricASN }).(pulumi.Float64Output)
}

// The version of Network Fabric.
func (o LookupNetworkFabricResultOutput) FabricVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkFabricResult) string { return v.FabricVersion }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupNetworkFabricResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkFabricResult) string { return v.Id }).(pulumi.StringOutput)
}

// IPv4Prefix for Management Network. Example: 10.1.0.0/19.
func (o LookupNetworkFabricResultOutput) Ipv4Prefix() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkFabricResult) string { return v.Ipv4Prefix }).(pulumi.StringOutput)
}

// IPv6Prefix for Management Network. Example: 3FFE:FFFF:0:CD40::/59
func (o LookupNetworkFabricResultOutput) Ipv6Prefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkFabricResult) *string { return v.Ipv6Prefix }).(pulumi.StringPtrOutput)
}

// List of L2 Isolation Domain resource IDs under the Network Fabric.
func (o LookupNetworkFabricResultOutput) L2IsolationDomains() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupNetworkFabricResult) []string { return v.L2IsolationDomains }).(pulumi.StringArrayOutput)
}

// List of L3 Isolation Domain resource IDs under the Network Fabric.
func (o LookupNetworkFabricResultOutput) L3IsolationDomains() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupNetworkFabricResult) []string { return v.L3IsolationDomains }).(pulumi.StringArrayOutput)
}

// The geo-location where the resource lives
func (o LookupNetworkFabricResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkFabricResult) string { return v.Location }).(pulumi.StringOutput)
}

// Configuration to be used to setup the management network.
func (o LookupNetworkFabricResultOutput) ManagementNetworkConfiguration() ManagementNetworkConfigurationPropertiesResponseOutput {
	return o.ApplyT(func(v LookupNetworkFabricResult) ManagementNetworkConfigurationPropertiesResponse {
		return v.ManagementNetworkConfiguration
	}).(ManagementNetworkConfigurationPropertiesResponseOutput)
}

// The name of the resource
func (o LookupNetworkFabricResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkFabricResult) string { return v.Name }).(pulumi.StringOutput)
}

// Azure resource ID for the NetworkFabricController the NetworkFabric belongs.
func (o LookupNetworkFabricResultOutput) NetworkFabricControllerId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkFabricResult) string { return v.NetworkFabricControllerId }).(pulumi.StringOutput)
}

// Supported Network Fabric SKU.Example: Compute / Aggregate racks. Once the user chooses a particular SKU, only supported racks can be added to the Network Fabric. The SKU determines whether it is a single / multi rack Network Fabric.
func (o LookupNetworkFabricResultOutput) NetworkFabricSku() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkFabricResult) string { return v.NetworkFabricSku }).(pulumi.StringOutput)
}

// Provides you the latest status of the NFC service, whether it is Accepted, updating, Succeeded or Failed. During this process, the states keep changing based on the status of NFC provisioning.
func (o LookupNetworkFabricResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkFabricResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Number of compute racks associated to Network Fabric.
func (o LookupNetworkFabricResultOutput) RackCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupNetworkFabricResult) *int { return v.RackCount }).(pulumi.IntPtrOutput)
}

// List of NetworkRack resource IDs under the Network Fabric. The number of racks allowed depends on the Network Fabric SKU.
func (o LookupNetworkFabricResultOutput) Racks() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupNetworkFabricResult) []string { return v.Racks }).(pulumi.StringArrayOutput)
}

// Array of router IDs.
func (o LookupNetworkFabricResultOutput) RouterIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupNetworkFabricResult) []string { return v.RouterIds }).(pulumi.StringArrayOutput)
}

// Number of servers.Possible values are from 1-16.
func (o LookupNetworkFabricResultOutput) ServerCountPerRack() pulumi.IntOutput {
	return o.ApplyT(func(v LookupNetworkFabricResult) int { return v.ServerCountPerRack }).(pulumi.IntOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupNetworkFabricResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupNetworkFabricResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupNetworkFabricResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupNetworkFabricResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Network and credentials configuration currently applied to terminal server.
func (o LookupNetworkFabricResultOutput) TerminalServerConfiguration() TerminalServerConfigurationResponseOutput {
	return o.ApplyT(func(v LookupNetworkFabricResult) TerminalServerConfigurationResponse {
		return v.TerminalServerConfiguration
	}).(TerminalServerConfigurationResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupNetworkFabricResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkFabricResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNetworkFabricResultOutput{})
}
