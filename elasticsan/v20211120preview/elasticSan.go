// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20211120preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Response for ElasticSan request.
type ElasticSan struct {
	pulumi.CustomResourceState

	// Logical zone for Elastic San resource; example: ["1"].
	AvailabilityZones pulumi.StringArrayOutput `pulumi:"availabilityZones"`
	// Base size of the Elastic San appliance in TiB.
	BaseSizeTiB pulumi.Float64Output `pulumi:"baseSizeTiB"`
	// Extended size of the Elastic San appliance in TiB.
	ExtendedCapacitySizeTiB pulumi.Float64Output `pulumi:"extendedCapacitySizeTiB"`
	// The geo-location where the resource lives.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Azure resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// State of the operation on the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// resource sku
	Sku SkuResponseOutput `pulumi:"sku"`
	// Resource metadata required by ARM RPC
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Azure resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Total Provisioned IOPS of the Elastic San appliance.
	TotalIops pulumi.Float64Output `pulumi:"totalIops"`
	// Total Provisioned MBps Elastic San appliance.
	TotalMBps pulumi.Float64Output `pulumi:"totalMBps"`
	// Total size of the Elastic San appliance in TB.
	TotalSizeTiB pulumi.Float64Output `pulumi:"totalSizeTiB"`
	// Total size of the provisioned Volumes in GiB.
	TotalVolumeSizeGiB pulumi.Float64Output `pulumi:"totalVolumeSizeGiB"`
	// Azure resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// Total number of volume groups in this Elastic San appliance.
	VolumeGroupCount pulumi.Float64Output `pulumi:"volumeGroupCount"`
}

// NewElasticSan registers a new resource with the given unique name, arguments, and options.
func NewElasticSan(ctx *pulumi.Context,
	name string, args *ElasticSanArgs, opts ...pulumi.ResourceOption) (*ElasticSan, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BaseSizeTiB == nil {
		return nil, errors.New("invalid value for required argument 'BaseSizeTiB'")
	}
	if args.ExtendedCapacitySizeTiB == nil {
		return nil, errors.New("invalid value for required argument 'ExtendedCapacitySizeTiB'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Sku == nil {
		return nil, errors.New("invalid value for required argument 'Sku'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:elasticsan:ElasticSan"),
		},
		{
			Type: pulumi.String("azure-native:elasticsan/v20221201preview:ElasticSan"),
		},
	})
	opts = append(opts, aliases)
	var resource ElasticSan
	err := ctx.RegisterResource("azure-native:elasticsan/v20211120preview:ElasticSan", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetElasticSan gets an existing ElasticSan resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetElasticSan(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ElasticSanState, opts ...pulumi.ResourceOption) (*ElasticSan, error) {
	var resource ElasticSan
	err := ctx.ReadResource("azure-native:elasticsan/v20211120preview:ElasticSan", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ElasticSan resources.
type elasticSanState struct {
}

type ElasticSanState struct {
}

func (ElasticSanState) ElementType() reflect.Type {
	return reflect.TypeOf((*elasticSanState)(nil)).Elem()
}

type elasticSanArgs struct {
	// Logical zone for Elastic San resource; example: ["1"].
	AvailabilityZones []string `pulumi:"availabilityZones"`
	// Base size of the Elastic San appliance in TiB.
	BaseSizeTiB float64 `pulumi:"baseSizeTiB"`
	// The name of the ElasticSan.
	ElasticSanName *string `pulumi:"elasticSanName"`
	// Extended size of the Elastic San appliance in TiB.
	ExtendedCapacitySizeTiB float64 `pulumi:"extendedCapacitySizeTiB"`
	// The geo-location where the resource lives.
	Location *string `pulumi:"location"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// resource sku
	Sku Sku `pulumi:"sku"`
	// Azure resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a ElasticSan resource.
type ElasticSanArgs struct {
	// Logical zone for Elastic San resource; example: ["1"].
	AvailabilityZones pulumi.StringArrayInput
	// Base size of the Elastic San appliance in TiB.
	BaseSizeTiB pulumi.Float64Input
	// The name of the ElasticSan.
	ElasticSanName pulumi.StringPtrInput
	// Extended size of the Elastic San appliance in TiB.
	ExtendedCapacitySizeTiB pulumi.Float64Input
	// The geo-location where the resource lives.
	Location pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// resource sku
	Sku SkuInput
	// Azure resource tags.
	Tags pulumi.StringMapInput
}

func (ElasticSanArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*elasticSanArgs)(nil)).Elem()
}

type ElasticSanInput interface {
	pulumi.Input

	ToElasticSanOutput() ElasticSanOutput
	ToElasticSanOutputWithContext(ctx context.Context) ElasticSanOutput
}

func (*ElasticSan) ElementType() reflect.Type {
	return reflect.TypeOf((**ElasticSan)(nil)).Elem()
}

func (i *ElasticSan) ToElasticSanOutput() ElasticSanOutput {
	return i.ToElasticSanOutputWithContext(context.Background())
}

func (i *ElasticSan) ToElasticSanOutputWithContext(ctx context.Context) ElasticSanOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ElasticSanOutput)
}

type ElasticSanOutput struct{ *pulumi.OutputState }

func (ElasticSanOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ElasticSan)(nil)).Elem()
}

func (o ElasticSanOutput) ToElasticSanOutput() ElasticSanOutput {
	return o
}

func (o ElasticSanOutput) ToElasticSanOutputWithContext(ctx context.Context) ElasticSanOutput {
	return o
}

// Logical zone for Elastic San resource; example: ["1"].
func (o ElasticSanOutput) AvailabilityZones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ElasticSan) pulumi.StringArrayOutput { return v.AvailabilityZones }).(pulumi.StringArrayOutput)
}

// Base size of the Elastic San appliance in TiB.
func (o ElasticSanOutput) BaseSizeTiB() pulumi.Float64Output {
	return o.ApplyT(func(v *ElasticSan) pulumi.Float64Output { return v.BaseSizeTiB }).(pulumi.Float64Output)
}

// Extended size of the Elastic San appliance in TiB.
func (o ElasticSanOutput) ExtendedCapacitySizeTiB() pulumi.Float64Output {
	return o.ApplyT(func(v *ElasticSan) pulumi.Float64Output { return v.ExtendedCapacitySizeTiB }).(pulumi.Float64Output)
}

// The geo-location where the resource lives.
func (o ElasticSanOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ElasticSan) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Azure resource name.
func (o ElasticSanOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ElasticSan) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// State of the operation on the resource.
func (o ElasticSanOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ElasticSan) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// resource sku
func (o ElasticSanOutput) Sku() SkuResponseOutput {
	return o.ApplyT(func(v *ElasticSan) SkuResponseOutput { return v.Sku }).(SkuResponseOutput)
}

// Resource metadata required by ARM RPC
func (o ElasticSanOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ElasticSan) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Azure resource tags.
func (o ElasticSanOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ElasticSan) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Total Provisioned IOPS of the Elastic San appliance.
func (o ElasticSanOutput) TotalIops() pulumi.Float64Output {
	return o.ApplyT(func(v *ElasticSan) pulumi.Float64Output { return v.TotalIops }).(pulumi.Float64Output)
}

// Total Provisioned MBps Elastic San appliance.
func (o ElasticSanOutput) TotalMBps() pulumi.Float64Output {
	return o.ApplyT(func(v *ElasticSan) pulumi.Float64Output { return v.TotalMBps }).(pulumi.Float64Output)
}

// Total size of the Elastic San appliance in TB.
func (o ElasticSanOutput) TotalSizeTiB() pulumi.Float64Output {
	return o.ApplyT(func(v *ElasticSan) pulumi.Float64Output { return v.TotalSizeTiB }).(pulumi.Float64Output)
}

// Total size of the provisioned Volumes in GiB.
func (o ElasticSanOutput) TotalVolumeSizeGiB() pulumi.Float64Output {
	return o.ApplyT(func(v *ElasticSan) pulumi.Float64Output { return v.TotalVolumeSizeGiB }).(pulumi.Float64Output)
}

// Azure resource type.
func (o ElasticSanOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ElasticSan) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Total number of volume groups in this Elastic San appliance.
func (o ElasticSanOutput) VolumeGroupCount() pulumi.Float64Output {
	return o.ApplyT(func(v *ElasticSan) pulumi.Float64Output { return v.VolumeGroupCount }).(pulumi.Float64Output)
}

func init() {
	pulumi.RegisterOutputType(ElasticSanOutput{})
}
