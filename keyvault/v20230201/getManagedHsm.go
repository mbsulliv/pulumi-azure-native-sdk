// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets the specified managed HSM Pool.
func LookupManagedHsm(ctx *pulumi.Context, args *LookupManagedHsmArgs, opts ...pulumi.InvokeOption) (*LookupManagedHsmResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupManagedHsmResult
	err := ctx.Invoke("azure-native:keyvault/v20230201:getManagedHsm", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupManagedHsmArgs struct {
	// The name of the managed HSM Pool.
	Name string `pulumi:"name"`
	// Name of the resource group that contains the managed HSM pool.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Resource information with extended details.
type LookupManagedHsmResult struct {
	// The Azure Resource Manager resource ID for the managed HSM Pool.
	Id string `pulumi:"id"`
	// The supported Azure location where the managed HSM Pool should be created.
	Location *string `pulumi:"location"`
	// The name of the managed HSM Pool.
	Name string `pulumi:"name"`
	// Properties of the managed HSM
	Properties ManagedHsmPropertiesResponse `pulumi:"properties"`
	// SKU details
	Sku *ManagedHsmSkuResponse `pulumi:"sku"`
	// Metadata pertaining to creation and last modification of the key vault resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// The resource type of the managed HSM Pool.
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupManagedHsmResult
func (val *LookupManagedHsmResult) Defaults() *LookupManagedHsmResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Properties = *tmp.Properties.Defaults()

	return &tmp
}

func LookupManagedHsmOutput(ctx *pulumi.Context, args LookupManagedHsmOutputArgs, opts ...pulumi.InvokeOption) LookupManagedHsmResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupManagedHsmResult, error) {
			args := v.(LookupManagedHsmArgs)
			r, err := LookupManagedHsm(ctx, &args, opts...)
			var s LookupManagedHsmResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupManagedHsmResultOutput)
}

type LookupManagedHsmOutputArgs struct {
	// The name of the managed HSM Pool.
	Name pulumi.StringInput `pulumi:"name"`
	// Name of the resource group that contains the managed HSM pool.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupManagedHsmOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagedHsmArgs)(nil)).Elem()
}

// Resource information with extended details.
type LookupManagedHsmResultOutput struct{ *pulumi.OutputState }

func (LookupManagedHsmResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagedHsmResult)(nil)).Elem()
}

func (o LookupManagedHsmResultOutput) ToLookupManagedHsmResultOutput() LookupManagedHsmResultOutput {
	return o
}

func (o LookupManagedHsmResultOutput) ToLookupManagedHsmResultOutputWithContext(ctx context.Context) LookupManagedHsmResultOutput {
	return o
}

func (o LookupManagedHsmResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupManagedHsmResult] {
	return pulumix.Output[LookupManagedHsmResult]{
		OutputState: o.OutputState,
	}
}

// The Azure Resource Manager resource ID for the managed HSM Pool.
func (o LookupManagedHsmResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedHsmResult) string { return v.Id }).(pulumi.StringOutput)
}

// The supported Azure location where the managed HSM Pool should be created.
func (o LookupManagedHsmResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedHsmResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the managed HSM Pool.
func (o LookupManagedHsmResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedHsmResult) string { return v.Name }).(pulumi.StringOutput)
}

// Properties of the managed HSM
func (o LookupManagedHsmResultOutput) Properties() ManagedHsmPropertiesResponseOutput {
	return o.ApplyT(func(v LookupManagedHsmResult) ManagedHsmPropertiesResponse { return v.Properties }).(ManagedHsmPropertiesResponseOutput)
}

// SKU details
func (o LookupManagedHsmResultOutput) Sku() ManagedHsmSkuResponsePtrOutput {
	return o.ApplyT(func(v LookupManagedHsmResult) *ManagedHsmSkuResponse { return v.Sku }).(ManagedHsmSkuResponsePtrOutput)
}

// Metadata pertaining to creation and last modification of the key vault resource.
func (o LookupManagedHsmResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupManagedHsmResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags
func (o LookupManagedHsmResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupManagedHsmResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The resource type of the managed HSM Pool.
func (o LookupManagedHsmResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedHsmResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupManagedHsmResultOutput{})
}
