// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get the metadata of an IoT Central application.
func LookupApp(ctx *pulumi.Context, args *LookupAppArgs, opts ...pulumi.InvokeOption) (*LookupAppResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupAppResult
	err := ctx.Invoke("azure-native:iotcentral/v20210601:getApp", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAppArgs struct {
	// The name of the resource group that contains the IoT Central application.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The ARM resource name of the IoT Central application.
	ResourceName string `pulumi:"resourceName"`
}

// The IoT Central application.
type LookupAppResult struct {
	// The ID of the application.
	ApplicationId string `pulumi:"applicationId"`
	// The display name of the application.
	DisplayName *string `pulumi:"displayName"`
	// The ARM resource identifier.
	Id string `pulumi:"id"`
	// The managed identities for the IoT Central application.
	Identity *SystemAssignedServiceIdentityResponse `pulumi:"identity"`
	// The resource location.
	Location string `pulumi:"location"`
	// The ARM resource name.
	Name string `pulumi:"name"`
	// A valid instance SKU.
	Sku AppSkuInfoResponse `pulumi:"sku"`
	// The current state of the application.
	State string `pulumi:"state"`
	// The subdomain of the application.
	Subdomain *string `pulumi:"subdomain"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The ID of the application template, which is a blueprint that defines the characteristics and behaviors of an application. Optional; if not specified, defaults to a blank blueprint and allows the application to be defined from scratch.
	Template *string `pulumi:"template"`
	// The resource type.
	Type string `pulumi:"type"`
}

func LookupAppOutput(ctx *pulumi.Context, args LookupAppOutputArgs, opts ...pulumi.InvokeOption) LookupAppResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAppResult, error) {
			args := v.(LookupAppArgs)
			r, err := LookupApp(ctx, &args, opts...)
			var s LookupAppResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAppResultOutput)
}

type LookupAppOutputArgs struct {
	// The name of the resource group that contains the IoT Central application.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The ARM resource name of the IoT Central application.
	ResourceName pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupAppOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAppArgs)(nil)).Elem()
}

// The IoT Central application.
type LookupAppResultOutput struct{ *pulumi.OutputState }

func (LookupAppResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAppResult)(nil)).Elem()
}

func (o LookupAppResultOutput) ToLookupAppResultOutput() LookupAppResultOutput {
	return o
}

func (o LookupAppResultOutput) ToLookupAppResultOutputWithContext(ctx context.Context) LookupAppResultOutput {
	return o
}

// The ID of the application.
func (o LookupAppResultOutput) ApplicationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppResult) string { return v.ApplicationId }).(pulumi.StringOutput)
}

// The display name of the application.
func (o LookupAppResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAppResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// The ARM resource identifier.
func (o LookupAppResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppResult) string { return v.Id }).(pulumi.StringOutput)
}

// The managed identities for the IoT Central application.
func (o LookupAppResultOutput) Identity() SystemAssignedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupAppResult) *SystemAssignedServiceIdentityResponse { return v.Identity }).(SystemAssignedServiceIdentityResponsePtrOutput)
}

// The resource location.
func (o LookupAppResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppResult) string { return v.Location }).(pulumi.StringOutput)
}

// The ARM resource name.
func (o LookupAppResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppResult) string { return v.Name }).(pulumi.StringOutput)
}

// A valid instance SKU.
func (o LookupAppResultOutput) Sku() AppSkuInfoResponseOutput {
	return o.ApplyT(func(v LookupAppResult) AppSkuInfoResponse { return v.Sku }).(AppSkuInfoResponseOutput)
}

// The current state of the application.
func (o LookupAppResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppResult) string { return v.State }).(pulumi.StringOutput)
}

// The subdomain of the application.
func (o LookupAppResultOutput) Subdomain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAppResult) *string { return v.Subdomain }).(pulumi.StringPtrOutput)
}

// The resource tags.
func (o LookupAppResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupAppResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The ID of the application template, which is a blueprint that defines the characteristics and behaviors of an application. Optional; if not specified, defaults to a blank blueprint and allows the application to be defined from scratch.
func (o LookupAppResultOutput) Template() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAppResult) *string { return v.Template }).(pulumi.StringPtrOutput)
}

// The resource type.
func (o LookupAppResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAppResultOutput{})
}
