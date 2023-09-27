// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package m365securityandcompliance

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get the metadata of a privateLinkServicesForMIPPolicySync resource.
// Azure REST API version: 2021-03-25-preview.
func LookupPrivateLinkServicesForMIPPolicySync(ctx *pulumi.Context, args *LookupPrivateLinkServicesForMIPPolicySyncArgs, opts ...pulumi.InvokeOption) (*LookupPrivateLinkServicesForMIPPolicySyncResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupPrivateLinkServicesForMIPPolicySyncResult
	err := ctx.Invoke("azure-native:m365securityandcompliance:getPrivateLinkServicesForMIPPolicySync", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPrivateLinkServicesForMIPPolicySyncArgs struct {
	// The name of the resource group that contains the service instance.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the service instance.
	ResourceName string `pulumi:"resourceName"`
}

// The description of the service.
type LookupPrivateLinkServicesForMIPPolicySyncResult struct {
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

func LookupPrivateLinkServicesForMIPPolicySyncOutput(ctx *pulumi.Context, args LookupPrivateLinkServicesForMIPPolicySyncOutputArgs, opts ...pulumi.InvokeOption) LookupPrivateLinkServicesForMIPPolicySyncResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPrivateLinkServicesForMIPPolicySyncResult, error) {
			args := v.(LookupPrivateLinkServicesForMIPPolicySyncArgs)
			r, err := LookupPrivateLinkServicesForMIPPolicySync(ctx, &args, opts...)
			var s LookupPrivateLinkServicesForMIPPolicySyncResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPrivateLinkServicesForMIPPolicySyncResultOutput)
}

type LookupPrivateLinkServicesForMIPPolicySyncOutputArgs struct {
	// The name of the resource group that contains the service instance.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the service instance.
	ResourceName pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupPrivateLinkServicesForMIPPolicySyncOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateLinkServicesForMIPPolicySyncArgs)(nil)).Elem()
}

// The description of the service.
type LookupPrivateLinkServicesForMIPPolicySyncResultOutput struct{ *pulumi.OutputState }

func (LookupPrivateLinkServicesForMIPPolicySyncResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateLinkServicesForMIPPolicySyncResult)(nil)).Elem()
}

func (o LookupPrivateLinkServicesForMIPPolicySyncResultOutput) ToLookupPrivateLinkServicesForMIPPolicySyncResultOutput() LookupPrivateLinkServicesForMIPPolicySyncResultOutput {
	return o
}

func (o LookupPrivateLinkServicesForMIPPolicySyncResultOutput) ToLookupPrivateLinkServicesForMIPPolicySyncResultOutputWithContext(ctx context.Context) LookupPrivateLinkServicesForMIPPolicySyncResultOutput {
	return o
}

func (o LookupPrivateLinkServicesForMIPPolicySyncResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupPrivateLinkServicesForMIPPolicySyncResult] {
	return pulumix.Output[LookupPrivateLinkServicesForMIPPolicySyncResult]{
		OutputState: o.OutputState,
	}
}

// An etag associated with the resource, used for optimistic concurrency when editing it.
func (o LookupPrivateLinkServicesForMIPPolicySyncResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPrivateLinkServicesForMIPPolicySyncResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// The resource identifier.
func (o LookupPrivateLinkServicesForMIPPolicySyncResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateLinkServicesForMIPPolicySyncResult) string { return v.Id }).(pulumi.StringOutput)
}

// Setting indicating whether the service has a managed identity associated with it.
func (o LookupPrivateLinkServicesForMIPPolicySyncResultOutput) Identity() ServicesResourceResponseIdentityPtrOutput {
	return o.ApplyT(func(v LookupPrivateLinkServicesForMIPPolicySyncResult) *ServicesResourceResponseIdentity {
		return v.Identity
	}).(ServicesResourceResponseIdentityPtrOutput)
}

// The kind of the service.
func (o LookupPrivateLinkServicesForMIPPolicySyncResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateLinkServicesForMIPPolicySyncResult) string { return v.Kind }).(pulumi.StringOutput)
}

// The resource location.
func (o LookupPrivateLinkServicesForMIPPolicySyncResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateLinkServicesForMIPPolicySyncResult) string { return v.Location }).(pulumi.StringOutput)
}

// The resource name.
func (o LookupPrivateLinkServicesForMIPPolicySyncResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateLinkServicesForMIPPolicySyncResult) string { return v.Name }).(pulumi.StringOutput)
}

// The common properties of a service.
func (o LookupPrivateLinkServicesForMIPPolicySyncResultOutput) Properties() ServicesPropertiesResponseOutput {
	return o.ApplyT(func(v LookupPrivateLinkServicesForMIPPolicySyncResult) ServicesPropertiesResponse {
		return v.Properties
	}).(ServicesPropertiesResponseOutput)
}

// Required property for system data
func (o LookupPrivateLinkServicesForMIPPolicySyncResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupPrivateLinkServicesForMIPPolicySyncResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The resource tags.
func (o LookupPrivateLinkServicesForMIPPolicySyncResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupPrivateLinkServicesForMIPPolicySyncResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The resource type.
func (o LookupPrivateLinkServicesForMIPPolicySyncResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateLinkServicesForMIPPolicySyncResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPrivateLinkServicesForMIPPolicySyncResultOutput{})
}
