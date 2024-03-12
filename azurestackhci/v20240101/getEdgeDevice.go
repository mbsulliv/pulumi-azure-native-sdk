// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20240101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get a EdgeDevice
func LookupEdgeDevice(ctx *pulumi.Context, args *LookupEdgeDeviceArgs, opts ...pulumi.InvokeOption) (*LookupEdgeDeviceResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupEdgeDeviceResult
	err := ctx.Invoke("azure-native:azurestackhci/v20240101:getEdgeDevice", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEdgeDeviceArgs struct {
	// Name of Device
	EdgeDeviceName string `pulumi:"edgeDeviceName"`
	// The fully qualified Azure Resource manager identifier of the resource.
	ResourceUri string `pulumi:"resourceUri"`
}

// Edge device resource
type LookupEdgeDeviceResult struct {
	// Device Configuration
	DeviceConfiguration *DeviceConfigurationResponse `pulumi:"deviceConfiguration"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Provisioning state of edgeDevice resource
	ProvisioningState string `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupEdgeDeviceOutput(ctx *pulumi.Context, args LookupEdgeDeviceOutputArgs, opts ...pulumi.InvokeOption) LookupEdgeDeviceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEdgeDeviceResult, error) {
			args := v.(LookupEdgeDeviceArgs)
			r, err := LookupEdgeDevice(ctx, &args, opts...)
			var s LookupEdgeDeviceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupEdgeDeviceResultOutput)
}

type LookupEdgeDeviceOutputArgs struct {
	// Name of Device
	EdgeDeviceName pulumi.StringInput `pulumi:"edgeDeviceName"`
	// The fully qualified Azure Resource manager identifier of the resource.
	ResourceUri pulumi.StringInput `pulumi:"resourceUri"`
}

func (LookupEdgeDeviceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEdgeDeviceArgs)(nil)).Elem()
}

// Edge device resource
type LookupEdgeDeviceResultOutput struct{ *pulumi.OutputState }

func (LookupEdgeDeviceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEdgeDeviceResult)(nil)).Elem()
}

func (o LookupEdgeDeviceResultOutput) ToLookupEdgeDeviceResultOutput() LookupEdgeDeviceResultOutput {
	return o
}

func (o LookupEdgeDeviceResultOutput) ToLookupEdgeDeviceResultOutputWithContext(ctx context.Context) LookupEdgeDeviceResultOutput {
	return o
}

// Device Configuration
func (o LookupEdgeDeviceResultOutput) DeviceConfiguration() DeviceConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupEdgeDeviceResult) *DeviceConfigurationResponse { return v.DeviceConfiguration }).(DeviceConfigurationResponsePtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupEdgeDeviceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEdgeDeviceResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupEdgeDeviceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEdgeDeviceResult) string { return v.Name }).(pulumi.StringOutput)
}

// Provisioning state of edgeDevice resource
func (o LookupEdgeDeviceResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEdgeDeviceResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupEdgeDeviceResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupEdgeDeviceResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupEdgeDeviceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEdgeDeviceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEdgeDeviceResultOutput{})
}
