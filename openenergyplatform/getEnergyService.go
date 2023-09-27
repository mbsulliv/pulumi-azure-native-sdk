// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package openenergyplatform

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Returns oep resource for a given name.
// Azure REST API version: 2022-04-04-preview.
func LookupEnergyService(ctx *pulumi.Context, args *LookupEnergyServiceArgs, opts ...pulumi.InvokeOption) (*LookupEnergyServiceResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupEnergyServiceResult
	err := ctx.Invoke("azure-native:openenergyplatform:getEnergyService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEnergyServiceArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The resource name.
	ResourceName string `pulumi:"resourceName"`
}

type LookupEnergyServiceResult struct {
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// Geo-location where the resource lives.
	Location string `pulumi:"location"`
	// The name of the resource
	Name       string                          `pulumi:"name"`
	Properties EnergyServicePropertiesResponse `pulumi:"properties"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupEnergyServiceOutput(ctx *pulumi.Context, args LookupEnergyServiceOutputArgs, opts ...pulumi.InvokeOption) LookupEnergyServiceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEnergyServiceResult, error) {
			args := v.(LookupEnergyServiceArgs)
			r, err := LookupEnergyService(ctx, &args, opts...)
			var s LookupEnergyServiceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupEnergyServiceResultOutput)
}

type LookupEnergyServiceOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The resource name.
	ResourceName pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupEnergyServiceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEnergyServiceArgs)(nil)).Elem()
}

type LookupEnergyServiceResultOutput struct{ *pulumi.OutputState }

func (LookupEnergyServiceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEnergyServiceResult)(nil)).Elem()
}

func (o LookupEnergyServiceResultOutput) ToLookupEnergyServiceResultOutput() LookupEnergyServiceResultOutput {
	return o
}

func (o LookupEnergyServiceResultOutput) ToLookupEnergyServiceResultOutputWithContext(ctx context.Context) LookupEnergyServiceResultOutput {
	return o
}

func (o LookupEnergyServiceResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupEnergyServiceResult] {
	return pulumix.Output[LookupEnergyServiceResult]{
		OutputState: o.OutputState,
	}
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupEnergyServiceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnergyServiceResult) string { return v.Id }).(pulumi.StringOutput)
}

// Geo-location where the resource lives.
func (o LookupEnergyServiceResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnergyServiceResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupEnergyServiceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnergyServiceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupEnergyServiceResultOutput) Properties() EnergyServicePropertiesResponseOutput {
	return o.ApplyT(func(v LookupEnergyServiceResult) EnergyServicePropertiesResponse { return v.Properties }).(EnergyServicePropertiesResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupEnergyServiceResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupEnergyServiceResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupEnergyServiceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupEnergyServiceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupEnergyServiceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnergyServiceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEnergyServiceResultOutput{})
}
