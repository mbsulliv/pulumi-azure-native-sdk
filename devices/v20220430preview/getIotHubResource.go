// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220430preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get the non-security related metadata of an IoT hub.
func LookupIotHubResource(ctx *pulumi.Context, args *LookupIotHubResourceArgs, opts ...pulumi.InvokeOption) (*LookupIotHubResourceResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupIotHubResourceResult
	err := ctx.Invoke("azure-native:devices/v20220430preview:getIotHubResource", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupIotHubResourceArgs struct {
	// The name of the resource group that contains the IoT hub.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the IoT hub.
	ResourceName string `pulumi:"resourceName"`
}

// The description of the IoT hub.
type LookupIotHubResourceResult struct {
	// The Etag field is *not* required. If it is provided in the response body, it must also be provided as a header per the normal ETag convention.
	Etag *string `pulumi:"etag"`
	// The resource identifier.
	Id string `pulumi:"id"`
	// The managed identities for the IotHub.
	Identity *ArmIdentityResponse `pulumi:"identity"`
	// The resource location.
	Location string `pulumi:"location"`
	// The resource name.
	Name string `pulumi:"name"`
	// IotHub properties
	Properties IotHubPropertiesResponse `pulumi:"properties"`
	// IotHub SKU info
	Sku IotHubSkuInfoResponse `pulumi:"sku"`
	// The system meta data relating to this resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The resource type.
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupIotHubResourceResult
func (val *LookupIotHubResourceResult) Defaults() *LookupIotHubResourceResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Properties = *tmp.Properties.Defaults()

	return &tmp
}

func LookupIotHubResourceOutput(ctx *pulumi.Context, args LookupIotHubResourceOutputArgs, opts ...pulumi.InvokeOption) LookupIotHubResourceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIotHubResourceResult, error) {
			args := v.(LookupIotHubResourceArgs)
			r, err := LookupIotHubResource(ctx, &args, opts...)
			var s LookupIotHubResourceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIotHubResourceResultOutput)
}

type LookupIotHubResourceOutputArgs struct {
	// The name of the resource group that contains the IoT hub.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the IoT hub.
	ResourceName pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupIotHubResourceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIotHubResourceArgs)(nil)).Elem()
}

// The description of the IoT hub.
type LookupIotHubResourceResultOutput struct{ *pulumi.OutputState }

func (LookupIotHubResourceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIotHubResourceResult)(nil)).Elem()
}

func (o LookupIotHubResourceResultOutput) ToLookupIotHubResourceResultOutput() LookupIotHubResourceResultOutput {
	return o
}

func (o LookupIotHubResourceResultOutput) ToLookupIotHubResourceResultOutputWithContext(ctx context.Context) LookupIotHubResourceResultOutput {
	return o
}

// The Etag field is *not* required. If it is provided in the response body, it must also be provided as a header per the normal ETag convention.
func (o LookupIotHubResourceResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIotHubResourceResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// The resource identifier.
func (o LookupIotHubResourceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIotHubResourceResult) string { return v.Id }).(pulumi.StringOutput)
}

// The managed identities for the IotHub.
func (o LookupIotHubResourceResultOutput) Identity() ArmIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupIotHubResourceResult) *ArmIdentityResponse { return v.Identity }).(ArmIdentityResponsePtrOutput)
}

// The resource location.
func (o LookupIotHubResourceResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIotHubResourceResult) string { return v.Location }).(pulumi.StringOutput)
}

// The resource name.
func (o LookupIotHubResourceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIotHubResourceResult) string { return v.Name }).(pulumi.StringOutput)
}

// IotHub properties
func (o LookupIotHubResourceResultOutput) Properties() IotHubPropertiesResponseOutput {
	return o.ApplyT(func(v LookupIotHubResourceResult) IotHubPropertiesResponse { return v.Properties }).(IotHubPropertiesResponseOutput)
}

// IotHub SKU info
func (o LookupIotHubResourceResultOutput) Sku() IotHubSkuInfoResponseOutput {
	return o.ApplyT(func(v LookupIotHubResourceResult) IotHubSkuInfoResponse { return v.Sku }).(IotHubSkuInfoResponseOutput)
}

// The system meta data relating to this resource.
func (o LookupIotHubResourceResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupIotHubResourceResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The resource tags.
func (o LookupIotHubResourceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupIotHubResourceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The resource type.
func (o LookupIotHubResourceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIotHubResourceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIotHubResourceResultOutput{})
}
