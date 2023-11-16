// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azurestackhci

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a network interface
// Azure REST API version: 2022-12-15-preview.
//
// Other available API versions: 2023-07-01-preview, 2023-09-01-preview.
func LookupNetworkInterface(ctx *pulumi.Context, args *LookupNetworkInterfaceArgs, opts ...pulumi.InvokeOption) (*LookupNetworkInterfaceResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupNetworkInterfaceResult
	err := ctx.Invoke("azure-native:azurestackhci:getNetworkInterface", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNetworkInterfaceArgs struct {
	// Name of the network interface
	NetworkInterfaceName string `pulumi:"networkInterfaceName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The network interface resource definition.
type LookupNetworkInterfaceResult struct {
	// DNS Settings for the interface
	DnsSettings *InterfaceDNSSettingsResponse `pulumi:"dnsSettings"`
	// The extendedLocation of the resource.
	ExtendedLocation *ExtendedLocationResponse `pulumi:"extendedLocation"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// IPConfigurations - A list of IPConfigurations of the network interface.
	IpConfigurations []IPConfigurationResponse `pulumi:"ipConfigurations"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// MacAddress - The MAC address of the network interface.
	MacAddress *string `pulumi:"macAddress"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Provisioning state of the network interface.
	ProvisioningState string `pulumi:"provisioningState"`
	// The observed state of network interfaces
	Status NetworkInterfaceStatusResponse `pulumi:"status"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupNetworkInterfaceOutput(ctx *pulumi.Context, args LookupNetworkInterfaceOutputArgs, opts ...pulumi.InvokeOption) LookupNetworkInterfaceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNetworkInterfaceResult, error) {
			args := v.(LookupNetworkInterfaceArgs)
			r, err := LookupNetworkInterface(ctx, &args, opts...)
			var s LookupNetworkInterfaceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNetworkInterfaceResultOutput)
}

type LookupNetworkInterfaceOutputArgs struct {
	// Name of the network interface
	NetworkInterfaceName pulumi.StringInput `pulumi:"networkInterfaceName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupNetworkInterfaceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkInterfaceArgs)(nil)).Elem()
}

// The network interface resource definition.
type LookupNetworkInterfaceResultOutput struct{ *pulumi.OutputState }

func (LookupNetworkInterfaceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkInterfaceResult)(nil)).Elem()
}

func (o LookupNetworkInterfaceResultOutput) ToLookupNetworkInterfaceResultOutput() LookupNetworkInterfaceResultOutput {
	return o
}

func (o LookupNetworkInterfaceResultOutput) ToLookupNetworkInterfaceResultOutputWithContext(ctx context.Context) LookupNetworkInterfaceResultOutput {
	return o
}

// DNS Settings for the interface
func (o LookupNetworkInterfaceResultOutput) DnsSettings() InterfaceDNSSettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) *InterfaceDNSSettingsResponse { return v.DnsSettings }).(InterfaceDNSSettingsResponsePtrOutput)
}

// The extendedLocation of the resource.
func (o LookupNetworkInterfaceResultOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) *ExtendedLocationResponse { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupNetworkInterfaceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) string { return v.Id }).(pulumi.StringOutput)
}

// IPConfigurations - A list of IPConfigurations of the network interface.
func (o LookupNetworkInterfaceResultOutput) IpConfigurations() IPConfigurationResponseArrayOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) []IPConfigurationResponse { return v.IpConfigurations }).(IPConfigurationResponseArrayOutput)
}

// The geo-location where the resource lives
func (o LookupNetworkInterfaceResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) string { return v.Location }).(pulumi.StringOutput)
}

// MacAddress - The MAC address of the network interface.
func (o LookupNetworkInterfaceResultOutput) MacAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) *string { return v.MacAddress }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o LookupNetworkInterfaceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) string { return v.Name }).(pulumi.StringOutput)
}

// Provisioning state of the network interface.
func (o LookupNetworkInterfaceResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The observed state of network interfaces
func (o LookupNetworkInterfaceResultOutput) Status() NetworkInterfaceStatusResponseOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) NetworkInterfaceStatusResponse { return v.Status }).(NetworkInterfaceStatusResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupNetworkInterfaceResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupNetworkInterfaceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupNetworkInterfaceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNetworkInterfaceResultOutput{})
}
