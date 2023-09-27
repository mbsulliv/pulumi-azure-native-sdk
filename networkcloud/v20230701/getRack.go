// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230701

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get properties of the provided rack.
func LookupRack(ctx *pulumi.Context, args *LookupRackArgs, opts ...pulumi.InvokeOption) (*LookupRackResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupRackResult
	err := ctx.Invoke("azure-native:networkcloud/v20230701:getRack", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRackArgs struct {
	// The name of the rack.
	RackName string `pulumi:"rackName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

type LookupRackResult struct {
	// The value that will be used for machines in this rack to represent the availability zones that can be referenced by Hybrid AKS Clusters for node arrangement.
	AvailabilityZone string `pulumi:"availabilityZone"`
	// The resource ID of the cluster the rack is created for. This value is set when the rack is created by the cluster.
	ClusterId string `pulumi:"clusterId"`
	// The more detailed status of the rack.
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
	// The provisioning state of the rack resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// The free-form description of the rack location. (e.g. “DTN Datacenter, Floor 3, Isle 9, Rack 2B”)
	RackLocation string `pulumi:"rackLocation"`
	// The unique identifier for the rack within Network Cloud cluster. An alternate unique alphanumeric value other than a serial number may be provided if desired.
	RackSerialNumber string `pulumi:"rackSerialNumber"`
	// The SKU for the rack.
	RackSkuId string `pulumi:"rackSkuId"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupRackOutput(ctx *pulumi.Context, args LookupRackOutputArgs, opts ...pulumi.InvokeOption) LookupRackResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRackResult, error) {
			args := v.(LookupRackArgs)
			r, err := LookupRack(ctx, &args, opts...)
			var s LookupRackResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRackResultOutput)
}

type LookupRackOutputArgs struct {
	// The name of the rack.
	RackName pulumi.StringInput `pulumi:"rackName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupRackOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRackArgs)(nil)).Elem()
}

type LookupRackResultOutput struct{ *pulumi.OutputState }

func (LookupRackResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRackResult)(nil)).Elem()
}

func (o LookupRackResultOutput) ToLookupRackResultOutput() LookupRackResultOutput {
	return o
}

func (o LookupRackResultOutput) ToLookupRackResultOutputWithContext(ctx context.Context) LookupRackResultOutput {
	return o
}

func (o LookupRackResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupRackResult] {
	return pulumix.Output[LookupRackResult]{
		OutputState: o.OutputState,
	}
}

// The value that will be used for machines in this rack to represent the availability zones that can be referenced by Hybrid AKS Clusters for node arrangement.
func (o LookupRackResultOutput) AvailabilityZone() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRackResult) string { return v.AvailabilityZone }).(pulumi.StringOutput)
}

// The resource ID of the cluster the rack is created for. This value is set when the rack is created by the cluster.
func (o LookupRackResultOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRackResult) string { return v.ClusterId }).(pulumi.StringOutput)
}

// The more detailed status of the rack.
func (o LookupRackResultOutput) DetailedStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRackResult) string { return v.DetailedStatus }).(pulumi.StringOutput)
}

// The descriptive message about the current detailed status.
func (o LookupRackResultOutput) DetailedStatusMessage() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRackResult) string { return v.DetailedStatusMessage }).(pulumi.StringOutput)
}

// The extended location of the cluster associated with the resource.
func (o LookupRackResultOutput) ExtendedLocation() ExtendedLocationResponseOutput {
	return o.ApplyT(func(v LookupRackResult) ExtendedLocationResponse { return v.ExtendedLocation }).(ExtendedLocationResponseOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupRackResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRackResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupRackResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRackResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupRackResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRackResult) string { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the rack resource.
func (o LookupRackResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRackResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The free-form description of the rack location. (e.g. “DTN Datacenter, Floor 3, Isle 9, Rack 2B”)
func (o LookupRackResultOutput) RackLocation() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRackResult) string { return v.RackLocation }).(pulumi.StringOutput)
}

// The unique identifier for the rack within Network Cloud cluster. An alternate unique alphanumeric value other than a serial number may be provided if desired.
func (o LookupRackResultOutput) RackSerialNumber() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRackResult) string { return v.RackSerialNumber }).(pulumi.StringOutput)
}

// The SKU for the rack.
func (o LookupRackResultOutput) RackSkuId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRackResult) string { return v.RackSkuId }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupRackResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupRackResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupRackResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupRackResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupRackResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRackResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRackResultOutput{})
}
