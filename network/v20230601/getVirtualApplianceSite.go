// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the specified Virtual Appliance Site.
func LookupVirtualApplianceSite(ctx *pulumi.Context, args *LookupVirtualApplianceSiteArgs, opts ...pulumi.InvokeOption) (*LookupVirtualApplianceSiteResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupVirtualApplianceSiteResult
	err := ctx.Invoke("azure-native:network/v20230601:getVirtualApplianceSite", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVirtualApplianceSiteArgs struct {
	// The name of the Network Virtual Appliance.
	NetworkVirtualApplianceName string `pulumi:"networkVirtualApplianceName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the site.
	SiteName string `pulumi:"siteName"`
}

// Virtual Appliance Site resource.
type LookupVirtualApplianceSiteResult struct {
	// Address Prefix.
	AddressPrefix *string `pulumi:"addressPrefix"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// Name of the virtual appliance site.
	Name *string `pulumi:"name"`
	// Office 365 Policy.
	O365Policy *Office365PolicyPropertiesResponse `pulumi:"o365Policy"`
	// The provisioning state of the resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Site type.
	Type string `pulumi:"type"`
}

func LookupVirtualApplianceSiteOutput(ctx *pulumi.Context, args LookupVirtualApplianceSiteOutputArgs, opts ...pulumi.InvokeOption) LookupVirtualApplianceSiteResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVirtualApplianceSiteResult, error) {
			args := v.(LookupVirtualApplianceSiteArgs)
			r, err := LookupVirtualApplianceSite(ctx, &args, opts...)
			var s LookupVirtualApplianceSiteResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVirtualApplianceSiteResultOutput)
}

type LookupVirtualApplianceSiteOutputArgs struct {
	// The name of the Network Virtual Appliance.
	NetworkVirtualApplianceName pulumi.StringInput `pulumi:"networkVirtualApplianceName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the site.
	SiteName pulumi.StringInput `pulumi:"siteName"`
}

func (LookupVirtualApplianceSiteOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualApplianceSiteArgs)(nil)).Elem()
}

// Virtual Appliance Site resource.
type LookupVirtualApplianceSiteResultOutput struct{ *pulumi.OutputState }

func (LookupVirtualApplianceSiteResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualApplianceSiteResult)(nil)).Elem()
}

func (o LookupVirtualApplianceSiteResultOutput) ToLookupVirtualApplianceSiteResultOutput() LookupVirtualApplianceSiteResultOutput {
	return o
}

func (o LookupVirtualApplianceSiteResultOutput) ToLookupVirtualApplianceSiteResultOutputWithContext(ctx context.Context) LookupVirtualApplianceSiteResultOutput {
	return o
}

// Address Prefix.
func (o LookupVirtualApplianceSiteResultOutput) AddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualApplianceSiteResult) *string { return v.AddressPrefix }).(pulumi.StringPtrOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o LookupVirtualApplianceSiteResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualApplianceSiteResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Resource ID.
func (o LookupVirtualApplianceSiteResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualApplianceSiteResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// Name of the virtual appliance site.
func (o LookupVirtualApplianceSiteResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualApplianceSiteResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// Office 365 Policy.
func (o LookupVirtualApplianceSiteResultOutput) O365Policy() Office365PolicyPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualApplianceSiteResult) *Office365PolicyPropertiesResponse { return v.O365Policy }).(Office365PolicyPropertiesResponsePtrOutput)
}

// The provisioning state of the resource.
func (o LookupVirtualApplianceSiteResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualApplianceSiteResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Site type.
func (o LookupVirtualApplianceSiteResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualApplianceSiteResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVirtualApplianceSiteResultOutput{})
}
