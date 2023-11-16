// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20221201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get a ElasticSan.
func LookupElasticSan(ctx *pulumi.Context, args *LookupElasticSanArgs, opts ...pulumi.InvokeOption) (*LookupElasticSanResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupElasticSanResult
	err := ctx.Invoke("azure-native:elasticsan/v20221201preview:getElasticSan", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupElasticSanArgs struct {
	// The name of the ElasticSan.
	ElasticSanName string `pulumi:"elasticSanName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Response for ElasticSan request.
type LookupElasticSanResult struct {
	// Logical zone for Elastic San resource; example: ["1"].
	AvailabilityZones []string `pulumi:"availabilityZones"`
	// Base size of the Elastic San appliance in TiB.
	BaseSizeTiB float64 `pulumi:"baseSizeTiB"`
	// Extended size of the Elastic San appliance in TiB.
	ExtendedCapacitySizeTiB float64 `pulumi:"extendedCapacitySizeTiB"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The list of Private Endpoint Connections.
	PrivateEndpointConnections []PrivateEndpointConnectionResponse `pulumi:"privateEndpointConnections"`
	// State of the operation on the resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// resource sku
	Sku SkuResponse `pulumi:"sku"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Total Provisioned IOPS of the Elastic San appliance.
	TotalIops float64 `pulumi:"totalIops"`
	// Total Provisioned MBps Elastic San appliance.
	TotalMBps float64 `pulumi:"totalMBps"`
	// Total size of the Elastic San appliance in TB.
	TotalSizeTiB float64 `pulumi:"totalSizeTiB"`
	// Total size of the provisioned Volumes in GiB.
	TotalVolumeSizeGiB float64 `pulumi:"totalVolumeSizeGiB"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// Total number of volume groups in this Elastic San appliance.
	VolumeGroupCount float64 `pulumi:"volumeGroupCount"`
}

func LookupElasticSanOutput(ctx *pulumi.Context, args LookupElasticSanOutputArgs, opts ...pulumi.InvokeOption) LookupElasticSanResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupElasticSanResult, error) {
			args := v.(LookupElasticSanArgs)
			r, err := LookupElasticSan(ctx, &args, opts...)
			var s LookupElasticSanResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupElasticSanResultOutput)
}

type LookupElasticSanOutputArgs struct {
	// The name of the ElasticSan.
	ElasticSanName pulumi.StringInput `pulumi:"elasticSanName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupElasticSanOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupElasticSanArgs)(nil)).Elem()
}

// Response for ElasticSan request.
type LookupElasticSanResultOutput struct{ *pulumi.OutputState }

func (LookupElasticSanResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupElasticSanResult)(nil)).Elem()
}

func (o LookupElasticSanResultOutput) ToLookupElasticSanResultOutput() LookupElasticSanResultOutput {
	return o
}

func (o LookupElasticSanResultOutput) ToLookupElasticSanResultOutputWithContext(ctx context.Context) LookupElasticSanResultOutput {
	return o
}

// Logical zone for Elastic San resource; example: ["1"].
func (o LookupElasticSanResultOutput) AvailabilityZones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupElasticSanResult) []string { return v.AvailabilityZones }).(pulumi.StringArrayOutput)
}

// Base size of the Elastic San appliance in TiB.
func (o LookupElasticSanResultOutput) BaseSizeTiB() pulumi.Float64Output {
	return o.ApplyT(func(v LookupElasticSanResult) float64 { return v.BaseSizeTiB }).(pulumi.Float64Output)
}

// Extended size of the Elastic San appliance in TiB.
func (o LookupElasticSanResultOutput) ExtendedCapacitySizeTiB() pulumi.Float64Output {
	return o.ApplyT(func(v LookupElasticSanResult) float64 { return v.ExtendedCapacitySizeTiB }).(pulumi.Float64Output)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupElasticSanResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupElasticSanResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupElasticSanResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupElasticSanResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupElasticSanResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupElasticSanResult) string { return v.Name }).(pulumi.StringOutput)
}

// The list of Private Endpoint Connections.
func (o LookupElasticSanResultOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v LookupElasticSanResult) []PrivateEndpointConnectionResponse {
		return v.PrivateEndpointConnections
	}).(PrivateEndpointConnectionResponseArrayOutput)
}

// State of the operation on the resource.
func (o LookupElasticSanResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupElasticSanResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// resource sku
func (o LookupElasticSanResultOutput) Sku() SkuResponseOutput {
	return o.ApplyT(func(v LookupElasticSanResult) SkuResponse { return v.Sku }).(SkuResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupElasticSanResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupElasticSanResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupElasticSanResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupElasticSanResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Total Provisioned IOPS of the Elastic San appliance.
func (o LookupElasticSanResultOutput) TotalIops() pulumi.Float64Output {
	return o.ApplyT(func(v LookupElasticSanResult) float64 { return v.TotalIops }).(pulumi.Float64Output)
}

// Total Provisioned MBps Elastic San appliance.
func (o LookupElasticSanResultOutput) TotalMBps() pulumi.Float64Output {
	return o.ApplyT(func(v LookupElasticSanResult) float64 { return v.TotalMBps }).(pulumi.Float64Output)
}

// Total size of the Elastic San appliance in TB.
func (o LookupElasticSanResultOutput) TotalSizeTiB() pulumi.Float64Output {
	return o.ApplyT(func(v LookupElasticSanResult) float64 { return v.TotalSizeTiB }).(pulumi.Float64Output)
}

// Total size of the provisioned Volumes in GiB.
func (o LookupElasticSanResultOutput) TotalVolumeSizeGiB() pulumi.Float64Output {
	return o.ApplyT(func(v LookupElasticSanResult) float64 { return v.TotalVolumeSizeGiB }).(pulumi.Float64Output)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupElasticSanResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupElasticSanResult) string { return v.Type }).(pulumi.StringOutput)
}

// Total number of volume groups in this Elastic San appliance.
func (o LookupElasticSanResultOutput) VolumeGroupCount() pulumi.Float64Output {
	return o.ApplyT(func(v LookupElasticSanResult) float64 { return v.VolumeGroupCount }).(pulumi.Float64Output)
}

func init() {
	pulumi.RegisterOutputType(LookupElasticSanResultOutput{})
}
