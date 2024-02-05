// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package appplatform

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get the API portal and its properties.
// Azure REST API version: 2023-05-01-preview.
//
// Other available API versions: 2023-07-01-preview, 2023-09-01-preview, 2023-11-01-preview, 2023-12-01, 2024-01-01-preview.
func LookupApiPortal(ctx *pulumi.Context, args *LookupApiPortalArgs, opts ...pulumi.InvokeOption) (*LookupApiPortalResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupApiPortalResult
	err := ctx.Invoke("azure-native:appplatform:getApiPortal", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupApiPortalArgs struct {
	// The name of API portal.
	ApiPortalName string `pulumi:"apiPortalName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Service resource.
	ServiceName string `pulumi:"serviceName"`
}

// API portal resource
type LookupApiPortalResult struct {
	// Fully qualified resource Id for the resource.
	Id string `pulumi:"id"`
	// The name of the resource.
	Name string `pulumi:"name"`
	// API portal properties payload
	Properties ApiPortalPropertiesResponse `pulumi:"properties"`
	// Sku of the API portal resource
	Sku *SkuResponse `pulumi:"sku"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource.
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupApiPortalResult
func (val *LookupApiPortalResult) Defaults() *LookupApiPortalResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Properties = *tmp.Properties.Defaults()

	tmp.Sku = tmp.Sku.Defaults()

	return &tmp
}

func LookupApiPortalOutput(ctx *pulumi.Context, args LookupApiPortalOutputArgs, opts ...pulumi.InvokeOption) LookupApiPortalResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupApiPortalResult, error) {
			args := v.(LookupApiPortalArgs)
			r, err := LookupApiPortal(ctx, &args, opts...)
			var s LookupApiPortalResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupApiPortalResultOutput)
}

type LookupApiPortalOutputArgs struct {
	// The name of API portal.
	ApiPortalName pulumi.StringInput `pulumi:"apiPortalName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the Service resource.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupApiPortalOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApiPortalArgs)(nil)).Elem()
}

// API portal resource
type LookupApiPortalResultOutput struct{ *pulumi.OutputState }

func (LookupApiPortalResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApiPortalResult)(nil)).Elem()
}

func (o LookupApiPortalResultOutput) ToLookupApiPortalResultOutput() LookupApiPortalResultOutput {
	return o
}

func (o LookupApiPortalResultOutput) ToLookupApiPortalResultOutputWithContext(ctx context.Context) LookupApiPortalResultOutput {
	return o
}

// Fully qualified resource Id for the resource.
func (o LookupApiPortalResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiPortalResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource.
func (o LookupApiPortalResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiPortalResult) string { return v.Name }).(pulumi.StringOutput)
}

// API portal properties payload
func (o LookupApiPortalResultOutput) Properties() ApiPortalPropertiesResponseOutput {
	return o.ApplyT(func(v LookupApiPortalResult) ApiPortalPropertiesResponse { return v.Properties }).(ApiPortalPropertiesResponseOutput)
}

// Sku of the API portal resource
func (o LookupApiPortalResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupApiPortalResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupApiPortalResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupApiPortalResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource.
func (o LookupApiPortalResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiPortalResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApiPortalResultOutput{})
}
