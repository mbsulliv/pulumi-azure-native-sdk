// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package networkcloud

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get properties of the provided volume.
// Azure REST API version: 2023-05-01-preview.
//
// Other available API versions: 2023-07-01.
func LookupVolume(ctx *pulumi.Context, args *LookupVolumeArgs, opts ...pulumi.InvokeOption) (*LookupVolumeResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupVolumeResult
	err := ctx.Invoke("azure-native:networkcloud:getVolume", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVolumeArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the volume.
	VolumeName string `pulumi:"volumeName"`
}

type LookupVolumeResult struct {
	// The list of resource IDs that attach the volume. It may include virtual machines and Hybrid AKS clusters.
	AttachedTo []string `pulumi:"attachedTo"`
	// The more detailed status of the volume.
	DetailedStatus string `pulumi:"detailedStatus"`
	// The descriptive message about the current detailed status.
	DetailedStatusMessage string `pulumi:"detailedStatusMessage"`
	// The extended location of the cluster associated with the resource.
	ExtendedLocation ExtendedLocationResponse `pulumi:"extendedLocation"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The provisioning state of the volume.
	ProvisioningState string `pulumi:"provisioningState"`
	// The unique identifier of the volume.
	SerialNumber string `pulumi:"serialNumber"`
	// The size of the allocation for this volume in Mebibytes.
	SizeMiB float64 `pulumi:"sizeMiB"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupVolumeOutput(ctx *pulumi.Context, args LookupVolumeOutputArgs, opts ...pulumi.InvokeOption) LookupVolumeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVolumeResult, error) {
			args := v.(LookupVolumeArgs)
			r, err := LookupVolume(ctx, &args, opts...)
			var s LookupVolumeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVolumeResultOutput)
}

type LookupVolumeOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the volume.
	VolumeName pulumi.StringInput `pulumi:"volumeName"`
}

func (LookupVolumeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVolumeArgs)(nil)).Elem()
}

type LookupVolumeResultOutput struct{ *pulumi.OutputState }

func (LookupVolumeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVolumeResult)(nil)).Elem()
}

func (o LookupVolumeResultOutput) ToLookupVolumeResultOutput() LookupVolumeResultOutput {
	return o
}

func (o LookupVolumeResultOutput) ToLookupVolumeResultOutputWithContext(ctx context.Context) LookupVolumeResultOutput {
	return o
}

// The list of resource IDs that attach the volume. It may include virtual machines and Hybrid AKS clusters.
func (o LookupVolumeResultOutput) AttachedTo() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVolumeResult) []string { return v.AttachedTo }).(pulumi.StringArrayOutput)
}

// The more detailed status of the volume.
func (o LookupVolumeResultOutput) DetailedStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.DetailedStatus }).(pulumi.StringOutput)
}

// The descriptive message about the current detailed status.
func (o LookupVolumeResultOutput) DetailedStatusMessage() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.DetailedStatusMessage }).(pulumi.StringOutput)
}

// The extended location of the cluster associated with the resource.
func (o LookupVolumeResultOutput) ExtendedLocation() ExtendedLocationResponseOutput {
	return o.ApplyT(func(v LookupVolumeResult) ExtendedLocationResponse { return v.ExtendedLocation }).(ExtendedLocationResponseOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupVolumeResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupVolumeResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupVolumeResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the volume.
func (o LookupVolumeResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The unique identifier of the volume.
func (o LookupVolumeResultOutput) SerialNumber() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.SerialNumber }).(pulumi.StringOutput)
}

// The size of the allocation for this volume in Mebibytes.
func (o LookupVolumeResultOutput) SizeMiB() pulumi.Float64Output {
	return o.ApplyT(func(v LookupVolumeResult) float64 { return v.SizeMiB }).(pulumi.Float64Output)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupVolumeResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupVolumeResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupVolumeResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupVolumeResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupVolumeResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVolumeResultOutput{})
}
