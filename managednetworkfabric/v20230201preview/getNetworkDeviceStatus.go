// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get the running status of the Network Device.
func GetNetworkDeviceStatus(ctx *pulumi.Context, args *GetNetworkDeviceStatusArgs, opts ...pulumi.InvokeOption) (*GetNetworkDeviceStatusResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv GetNetworkDeviceStatusResult
	err := ctx.Invoke("azure-native:managednetworkfabric/v20230201preview:getNetworkDeviceStatus", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetNetworkDeviceStatusArgs struct {
	// Name of the NetworkDevice.
	NetworkDeviceName string `pulumi:"networkDeviceName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Get Device status response properties.
type GetNetworkDeviceStatusResult struct {
	// Primary or Secondary power end.
	OperationalStatus string `pulumi:"operationalStatus"`
	// On or Off power cycle state.
	PowerCycleState string `pulumi:"powerCycleState"`
	// The serial number of the device
	SerialNumber string `pulumi:"serialNumber"`
}

func GetNetworkDeviceStatusOutput(ctx *pulumi.Context, args GetNetworkDeviceStatusOutputArgs, opts ...pulumi.InvokeOption) GetNetworkDeviceStatusResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetNetworkDeviceStatusResult, error) {
			args := v.(GetNetworkDeviceStatusArgs)
			r, err := GetNetworkDeviceStatus(ctx, &args, opts...)
			var s GetNetworkDeviceStatusResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetNetworkDeviceStatusResultOutput)
}

type GetNetworkDeviceStatusOutputArgs struct {
	// Name of the NetworkDevice.
	NetworkDeviceName pulumi.StringInput `pulumi:"networkDeviceName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (GetNetworkDeviceStatusOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetNetworkDeviceStatusArgs)(nil)).Elem()
}

// Get Device status response properties.
type GetNetworkDeviceStatusResultOutput struct{ *pulumi.OutputState }

func (GetNetworkDeviceStatusResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetNetworkDeviceStatusResult)(nil)).Elem()
}

func (o GetNetworkDeviceStatusResultOutput) ToGetNetworkDeviceStatusResultOutput() GetNetworkDeviceStatusResultOutput {
	return o
}

func (o GetNetworkDeviceStatusResultOutput) ToGetNetworkDeviceStatusResultOutputWithContext(ctx context.Context) GetNetworkDeviceStatusResultOutput {
	return o
}

func (o GetNetworkDeviceStatusResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetNetworkDeviceStatusResult] {
	return pulumix.Output[GetNetworkDeviceStatusResult]{
		OutputState: o.OutputState,
	}
}

// Primary or Secondary power end.
func (o GetNetworkDeviceStatusResultOutput) OperationalStatus() pulumi.StringOutput {
	return o.ApplyT(func(v GetNetworkDeviceStatusResult) string { return v.OperationalStatus }).(pulumi.StringOutput)
}

// On or Off power cycle state.
func (o GetNetworkDeviceStatusResultOutput) PowerCycleState() pulumi.StringOutput {
	return o.ApplyT(func(v GetNetworkDeviceStatusResult) string { return v.PowerCycleState }).(pulumi.StringOutput)
}

// The serial number of the device
func (o GetNetworkDeviceStatusResultOutput) SerialNumber() pulumi.StringOutput {
	return o.ApplyT(func(v GetNetworkDeviceStatusResult) string { return v.SerialNumber }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetNetworkDeviceStatusResultOutput{})
}
