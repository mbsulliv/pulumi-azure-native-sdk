// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210325preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get the metadata of a privateLinkServicesForEDMUpload resource.
func LookupPrivateLinkServicesForEDMUpload(ctx *pulumi.Context, args *LookupPrivateLinkServicesForEDMUploadArgs, opts ...pulumi.InvokeOption) (*LookupPrivateLinkServicesForEDMUploadResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupPrivateLinkServicesForEDMUploadResult
	err := ctx.Invoke("azure-native:m365securityandcompliance/v20210325preview:getPrivateLinkServicesForEDMUpload", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPrivateLinkServicesForEDMUploadArgs struct {
	// The name of the resource group that contains the service instance.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the service instance.
	ResourceName string `pulumi:"resourceName"`
}

// The description of the service.
type LookupPrivateLinkServicesForEDMUploadResult struct {
	// An etag associated with the resource, used for optimistic concurrency when editing it.
	Etag *string `pulumi:"etag"`
	// The resource identifier.
	Id string `pulumi:"id"`
	// Setting indicating whether the service has a managed identity associated with it.
	Identity *ServicesResourceResponseIdentity `pulumi:"identity"`
	// The kind of the service.
	Kind string `pulumi:"kind"`
	// The resource location.
	Location string `pulumi:"location"`
	// The resource name.
	Name string `pulumi:"name"`
	// The common properties of a service.
	Properties ServicesPropertiesResponse `pulumi:"properties"`
	// Required property for system data
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The resource type.
	Type string `pulumi:"type"`
}

func LookupPrivateLinkServicesForEDMUploadOutput(ctx *pulumi.Context, args LookupPrivateLinkServicesForEDMUploadOutputArgs, opts ...pulumi.InvokeOption) LookupPrivateLinkServicesForEDMUploadResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPrivateLinkServicesForEDMUploadResult, error) {
			args := v.(LookupPrivateLinkServicesForEDMUploadArgs)
			r, err := LookupPrivateLinkServicesForEDMUpload(ctx, &args, opts...)
			var s LookupPrivateLinkServicesForEDMUploadResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPrivateLinkServicesForEDMUploadResultOutput)
}

type LookupPrivateLinkServicesForEDMUploadOutputArgs struct {
	// The name of the resource group that contains the service instance.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the service instance.
	ResourceName pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupPrivateLinkServicesForEDMUploadOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateLinkServicesForEDMUploadArgs)(nil)).Elem()
}

// The description of the service.
type LookupPrivateLinkServicesForEDMUploadResultOutput struct{ *pulumi.OutputState }

func (LookupPrivateLinkServicesForEDMUploadResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateLinkServicesForEDMUploadResult)(nil)).Elem()
}

func (o LookupPrivateLinkServicesForEDMUploadResultOutput) ToLookupPrivateLinkServicesForEDMUploadResultOutput() LookupPrivateLinkServicesForEDMUploadResultOutput {
	return o
}

func (o LookupPrivateLinkServicesForEDMUploadResultOutput) ToLookupPrivateLinkServicesForEDMUploadResultOutputWithContext(ctx context.Context) LookupPrivateLinkServicesForEDMUploadResultOutput {
	return o
}

func (o LookupPrivateLinkServicesForEDMUploadResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupPrivateLinkServicesForEDMUploadResult] {
	return pulumix.Output[LookupPrivateLinkServicesForEDMUploadResult]{
		OutputState: o.OutputState,
	}
}

// An etag associated with the resource, used for optimistic concurrency when editing it.
func (o LookupPrivateLinkServicesForEDMUploadResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPrivateLinkServicesForEDMUploadResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// The resource identifier.
func (o LookupPrivateLinkServicesForEDMUploadResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateLinkServicesForEDMUploadResult) string { return v.Id }).(pulumi.StringOutput)
}

// Setting indicating whether the service has a managed identity associated with it.
func (o LookupPrivateLinkServicesForEDMUploadResultOutput) Identity() ServicesResourceResponseIdentityPtrOutput {
	return o.ApplyT(func(v LookupPrivateLinkServicesForEDMUploadResult) *ServicesResourceResponseIdentity {
		return v.Identity
	}).(ServicesResourceResponseIdentityPtrOutput)
}

// The kind of the service.
func (o LookupPrivateLinkServicesForEDMUploadResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateLinkServicesForEDMUploadResult) string { return v.Kind }).(pulumi.StringOutput)
}

// The resource location.
func (o LookupPrivateLinkServicesForEDMUploadResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateLinkServicesForEDMUploadResult) string { return v.Location }).(pulumi.StringOutput)
}

// The resource name.
func (o LookupPrivateLinkServicesForEDMUploadResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateLinkServicesForEDMUploadResult) string { return v.Name }).(pulumi.StringOutput)
}

// The common properties of a service.
func (o LookupPrivateLinkServicesForEDMUploadResultOutput) Properties() ServicesPropertiesResponseOutput {
	return o.ApplyT(func(v LookupPrivateLinkServicesForEDMUploadResult) ServicesPropertiesResponse { return v.Properties }).(ServicesPropertiesResponseOutput)
}

// Required property for system data
func (o LookupPrivateLinkServicesForEDMUploadResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupPrivateLinkServicesForEDMUploadResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The resource tags.
func (o LookupPrivateLinkServicesForEDMUploadResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupPrivateLinkServicesForEDMUploadResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The resource type.
func (o LookupPrivateLinkServicesForEDMUploadResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateLinkServicesForEDMUploadResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPrivateLinkServicesForEDMUploadResultOutput{})
}
