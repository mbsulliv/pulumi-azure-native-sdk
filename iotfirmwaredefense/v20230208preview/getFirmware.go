// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230208preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get firmware.
func LookupFirmware(ctx *pulumi.Context, args *LookupFirmwareArgs, opts ...pulumi.InvokeOption) (*LookupFirmwareResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupFirmwareResult
	err := ctx.Invoke("azure-native:iotfirmwaredefense/v20230208preview:getFirmware", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupFirmwareArgs struct {
	// The id of the firmware.
	FirmwareId string `pulumi:"firmwareId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the firmware analysis workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// Firmware definition
type LookupFirmwareResult struct {
	// User-specified description of the firmware.
	Description *string `pulumi:"description"`
	// File name for a firmware that user uploaded.
	FileName *string `pulumi:"fileName"`
	// File size of the uploaded firmware image.
	FileSize *float64 `pulumi:"fileSize"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// Firmware model.
	Model *string `pulumi:"model"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Provisioning state of the resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// The status of firmware scan.
	Status *string `pulumi:"status"`
	// A list of errors or other messages generated during firmware analysis
	StatusMessages []interface{} `pulumi:"statusMessages"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// Firmware vendor.
	Vendor *string `pulumi:"vendor"`
	// Firmware version.
	Version *string `pulumi:"version"`
}

// Defaults sets the appropriate defaults for LookupFirmwareResult
func (val *LookupFirmwareResult) Defaults() *LookupFirmwareResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.Status == nil {
		status_ := "Pending"
		tmp.Status = &status_
	}
	return &tmp
}

func LookupFirmwareOutput(ctx *pulumi.Context, args LookupFirmwareOutputArgs, opts ...pulumi.InvokeOption) LookupFirmwareResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFirmwareResult, error) {
			args := v.(LookupFirmwareArgs)
			r, err := LookupFirmware(ctx, &args, opts...)
			var s LookupFirmwareResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFirmwareResultOutput)
}

type LookupFirmwareOutputArgs struct {
	// The id of the firmware.
	FirmwareId pulumi.StringInput `pulumi:"firmwareId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the firmware analysis workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupFirmwareOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFirmwareArgs)(nil)).Elem()
}

// Firmware definition
type LookupFirmwareResultOutput struct{ *pulumi.OutputState }

func (LookupFirmwareResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFirmwareResult)(nil)).Elem()
}

func (o LookupFirmwareResultOutput) ToLookupFirmwareResultOutput() LookupFirmwareResultOutput {
	return o
}

func (o LookupFirmwareResultOutput) ToLookupFirmwareResultOutputWithContext(ctx context.Context) LookupFirmwareResultOutput {
	return o
}

func (o LookupFirmwareResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupFirmwareResult] {
	return pulumix.Output[LookupFirmwareResult]{
		OutputState: o.OutputState,
	}
}

// User-specified description of the firmware.
func (o LookupFirmwareResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFirmwareResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// File name for a firmware that user uploaded.
func (o LookupFirmwareResultOutput) FileName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFirmwareResult) *string { return v.FileName }).(pulumi.StringPtrOutput)
}

// File size of the uploaded firmware image.
func (o LookupFirmwareResultOutput) FileSize() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupFirmwareResult) *float64 { return v.FileSize }).(pulumi.Float64PtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupFirmwareResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirmwareResult) string { return v.Id }).(pulumi.StringOutput)
}

// Firmware model.
func (o LookupFirmwareResultOutput) Model() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFirmwareResult) *string { return v.Model }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o LookupFirmwareResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirmwareResult) string { return v.Name }).(pulumi.StringOutput)
}

// Provisioning state of the resource.
func (o LookupFirmwareResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirmwareResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The status of firmware scan.
func (o LookupFirmwareResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFirmwareResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

// A list of errors or other messages generated during firmware analysis
func (o LookupFirmwareResultOutput) StatusMessages() pulumi.ArrayOutput {
	return o.ApplyT(func(v LookupFirmwareResult) []interface{} { return v.StatusMessages }).(pulumi.ArrayOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupFirmwareResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupFirmwareResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupFirmwareResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirmwareResult) string { return v.Type }).(pulumi.StringOutput)
}

// Firmware vendor.
func (o LookupFirmwareResultOutput) Vendor() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFirmwareResult) *string { return v.Vendor }).(pulumi.StringPtrOutput)
}

// Firmware version.
func (o LookupFirmwareResultOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFirmwareResult) *string { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFirmwareResultOutput{})
}
