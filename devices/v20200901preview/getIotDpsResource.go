// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20200901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get the metadata of the provisioning service without SAS keys.
func LookupIotDpsResource(ctx *pulumi.Context, args *LookupIotDpsResourceArgs, opts ...pulumi.InvokeOption) (*LookupIotDpsResourceResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupIotDpsResourceResult
	err := ctx.Invoke("azure-native:devices/v20200901preview:getIotDpsResource", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIotDpsResourceArgs struct {
	// Name of the provisioning service to retrieve.
	ProvisioningServiceName string `pulumi:"provisioningServiceName"`
	// Resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The description of the provisioning service.
type LookupIotDpsResourceResult struct {
	// The Etag field is *not* required. If it is provided in the response body, it must also be provided as a header per the normal ETag convention.
	Etag *string `pulumi:"etag"`
	// The resource identifier.
	Id string `pulumi:"id"`
	// The managed identities for the IotDps instance.
	Identity *ArmIdentityResponse `pulumi:"identity"`
	// The resource location.
	Location string `pulumi:"location"`
	// The resource name.
	Name string `pulumi:"name"`
	// Service specific properties for a provisioning service
	Properties IotDpsPropertiesDescriptionResponse `pulumi:"properties"`
	// Sku info for a provisioning Service.
	Sku IotDpsSkuInfoResponse `pulumi:"sku"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The resource type.
	Type string `pulumi:"type"`
}

func LookupIotDpsResourceOutput(ctx *pulumi.Context, args LookupIotDpsResourceOutputArgs, opts ...pulumi.InvokeOption) LookupIotDpsResourceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIotDpsResourceResult, error) {
			args := v.(LookupIotDpsResourceArgs)
			r, err := LookupIotDpsResource(ctx, &args, opts...)
			var s LookupIotDpsResourceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIotDpsResourceResultOutput)
}

type LookupIotDpsResourceOutputArgs struct {
	// Name of the provisioning service to retrieve.
	ProvisioningServiceName pulumi.StringInput `pulumi:"provisioningServiceName"`
	// Resource group name.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupIotDpsResourceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIotDpsResourceArgs)(nil)).Elem()
}

// The description of the provisioning service.
type LookupIotDpsResourceResultOutput struct{ *pulumi.OutputState }

func (LookupIotDpsResourceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIotDpsResourceResult)(nil)).Elem()
}

func (o LookupIotDpsResourceResultOutput) ToLookupIotDpsResourceResultOutput() LookupIotDpsResourceResultOutput {
	return o
}

func (o LookupIotDpsResourceResultOutput) ToLookupIotDpsResourceResultOutputWithContext(ctx context.Context) LookupIotDpsResourceResultOutput {
	return o
}

// The Etag field is *not* required. If it is provided in the response body, it must also be provided as a header per the normal ETag convention.
func (o LookupIotDpsResourceResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIotDpsResourceResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// The resource identifier.
func (o LookupIotDpsResourceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIotDpsResourceResult) string { return v.Id }).(pulumi.StringOutput)
}

// The managed identities for the IotDps instance.
func (o LookupIotDpsResourceResultOutput) Identity() ArmIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupIotDpsResourceResult) *ArmIdentityResponse { return v.Identity }).(ArmIdentityResponsePtrOutput)
}

// The resource location.
func (o LookupIotDpsResourceResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIotDpsResourceResult) string { return v.Location }).(pulumi.StringOutput)
}

// The resource name.
func (o LookupIotDpsResourceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIotDpsResourceResult) string { return v.Name }).(pulumi.StringOutput)
}

// Service specific properties for a provisioning service
func (o LookupIotDpsResourceResultOutput) Properties() IotDpsPropertiesDescriptionResponseOutput {
	return o.ApplyT(func(v LookupIotDpsResourceResult) IotDpsPropertiesDescriptionResponse { return v.Properties }).(IotDpsPropertiesDescriptionResponseOutput)
}

// Sku info for a provisioning Service.
func (o LookupIotDpsResourceResultOutput) Sku() IotDpsSkuInfoResponseOutput {
	return o.ApplyT(func(v LookupIotDpsResourceResult) IotDpsSkuInfoResponse { return v.Sku }).(IotDpsSkuInfoResponseOutput)
}

// The resource tags.
func (o LookupIotDpsResourceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupIotDpsResourceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The resource type.
func (o LookupIotDpsResourceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIotDpsResourceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIotDpsResourceResultOutput{})
}
