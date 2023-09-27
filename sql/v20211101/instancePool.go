// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20211101

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// An Azure SQL instance pool.
type InstancePool struct {
	pulumi.CustomResourceState

	// The license type. Possible values are 'LicenseIncluded' (price for SQL license is included) and 'BasePrice' (without SQL license price).
	LicenseType pulumi.StringOutput `pulumi:"licenseType"`
	// Resource location.
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name and tier of the SKU.
	Sku SkuResponsePtrOutput `pulumi:"sku"`
	// Resource ID of the subnet to place this instance pool in.
	SubnetId pulumi.StringOutput `pulumi:"subnetId"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// Count of vCores belonging to this instance pool.
	VCores pulumi.IntOutput `pulumi:"vCores"`
}

// NewInstancePool registers a new resource with the given unique name, arguments, and options.
func NewInstancePool(ctx *pulumi.Context,
	name string, args *InstancePoolArgs, opts ...pulumi.ResourceOption) (*InstancePool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LicenseType == nil {
		return nil, errors.New("invalid value for required argument 'LicenseType'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SubnetId == nil {
		return nil, errors.New("invalid value for required argument 'SubnetId'")
	}
	if args.VCores == nil {
		return nil, errors.New("invalid value for required argument 'VCores'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:sql:InstancePool"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20180601preview:InstancePool"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200202preview:InstancePool"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200801preview:InstancePool"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20201101preview:InstancePool"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210201preview:InstancePool"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210501preview:InstancePool"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210801preview:InstancePool"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101preview:InstancePool"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220201preview:InstancePool"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220501preview:InstancePool"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220801preview:InstancePool"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20221101preview:InstancePool"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20230201preview:InstancePool"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource InstancePool
	err := ctx.RegisterResource("azure-native:sql/v20211101:InstancePool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstancePool gets an existing InstancePool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstancePool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstancePoolState, opts ...pulumi.ResourceOption) (*InstancePool, error) {
	var resource InstancePool
	err := ctx.ReadResource("azure-native:sql/v20211101:InstancePool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InstancePool resources.
type instancePoolState struct {
}

type InstancePoolState struct {
}

func (InstancePoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*instancePoolState)(nil)).Elem()
}

type instancePoolArgs struct {
	// The name of the instance pool to be created or updated.
	InstancePoolName *string `pulumi:"instancePoolName"`
	// The license type. Possible values are 'LicenseIncluded' (price for SQL license is included) and 'BasePrice' (without SQL license price).
	LicenseType string `pulumi:"licenseType"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name and tier of the SKU.
	Sku *Sku `pulumi:"sku"`
	// Resource ID of the subnet to place this instance pool in.
	SubnetId string `pulumi:"subnetId"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Count of vCores belonging to this instance pool.
	VCores int `pulumi:"vCores"`
}

// The set of arguments for constructing a InstancePool resource.
type InstancePoolArgs struct {
	// The name of the instance pool to be created or updated.
	InstancePoolName pulumi.StringPtrInput
	// The license type. Possible values are 'LicenseIncluded' (price for SQL license is included) and 'BasePrice' (without SQL license price).
	LicenseType pulumi.StringInput
	// Resource location.
	Location pulumi.StringPtrInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// The name and tier of the SKU.
	Sku SkuPtrInput
	// Resource ID of the subnet to place this instance pool in.
	SubnetId pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Count of vCores belonging to this instance pool.
	VCores pulumi.IntInput
}

func (InstancePoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instancePoolArgs)(nil)).Elem()
}

type InstancePoolInput interface {
	pulumi.Input

	ToInstancePoolOutput() InstancePoolOutput
	ToInstancePoolOutputWithContext(ctx context.Context) InstancePoolOutput
}

func (*InstancePool) ElementType() reflect.Type {
	return reflect.TypeOf((**InstancePool)(nil)).Elem()
}

func (i *InstancePool) ToInstancePoolOutput() InstancePoolOutput {
	return i.ToInstancePoolOutputWithContext(context.Background())
}

func (i *InstancePool) ToInstancePoolOutputWithContext(ctx context.Context) InstancePoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstancePoolOutput)
}

func (i *InstancePool) ToOutput(ctx context.Context) pulumix.Output[*InstancePool] {
	return pulumix.Output[*InstancePool]{
		OutputState: i.ToInstancePoolOutputWithContext(ctx).OutputState,
	}
}

type InstancePoolOutput struct{ *pulumi.OutputState }

func (InstancePoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InstancePool)(nil)).Elem()
}

func (o InstancePoolOutput) ToInstancePoolOutput() InstancePoolOutput {
	return o
}

func (o InstancePoolOutput) ToInstancePoolOutputWithContext(ctx context.Context) InstancePoolOutput {
	return o
}

func (o InstancePoolOutput) ToOutput(ctx context.Context) pulumix.Output[*InstancePool] {
	return pulumix.Output[*InstancePool]{
		OutputState: o.OutputState,
	}
}

// The license type. Possible values are 'LicenseIncluded' (price for SQL license is included) and 'BasePrice' (without SQL license price).
func (o InstancePoolOutput) LicenseType() pulumi.StringOutput {
	return o.ApplyT(func(v *InstancePool) pulumi.StringOutput { return v.LicenseType }).(pulumi.StringOutput)
}

// Resource location.
func (o InstancePoolOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *InstancePool) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Resource name.
func (o InstancePoolOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *InstancePool) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The name and tier of the SKU.
func (o InstancePoolOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *InstancePool) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

// Resource ID of the subnet to place this instance pool in.
func (o InstancePoolOutput) SubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v *InstancePool) pulumi.StringOutput { return v.SubnetId }).(pulumi.StringOutput)
}

// Resource tags.
func (o InstancePoolOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *InstancePool) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o InstancePoolOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *InstancePool) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Count of vCores belonging to this instance pool.
func (o InstancePoolOutput) VCores() pulumi.IntOutput {
	return o.ApplyT(func(v *InstancePool) pulumi.IntOutput { return v.VCores }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(InstancePoolOutput{})
}
