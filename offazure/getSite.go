// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package offazure

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Site REST Resource.
// Azure REST API version: 2020-07-07.
func LookupSite(ctx *pulumi.Context, args *LookupSiteArgs, opts ...pulumi.InvokeOption) (*LookupSiteResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupSiteResult
	err := ctx.Invoke("azure-native:offazure:getSite", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSiteArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Site name.
	SiteName string `pulumi:"siteName"`
}

// Site REST Resource.
type LookupSiteResult struct {
	// eTag for concurrency control.
	ETag *string `pulumi:"eTag"`
	// Resource Id.
	Id string `pulumi:"id"`
	// Azure location in which Sites is created.
	Location *string `pulumi:"location"`
	// Name of the VMware site.
	Name *string `pulumi:"name"`
	// Nested properties of VMWare site.
	Properties SitePropertiesResponse `pulumi:"properties"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	Tags       map[string]string  `pulumi:"tags"`
	// Type of resource. Type = Microsoft.OffAzure/VMWareSites.
	Type string `pulumi:"type"`
}

func LookupSiteOutput(ctx *pulumi.Context, args LookupSiteOutputArgs, opts ...pulumi.InvokeOption) LookupSiteResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSiteResult, error) {
			args := v.(LookupSiteArgs)
			r, err := LookupSite(ctx, &args, opts...)
			var s LookupSiteResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSiteResultOutput)
}

type LookupSiteOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Site name.
	SiteName pulumi.StringInput `pulumi:"siteName"`
}

func (LookupSiteOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSiteArgs)(nil)).Elem()
}

// Site REST Resource.
type LookupSiteResultOutput struct{ *pulumi.OutputState }

func (LookupSiteResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSiteResult)(nil)).Elem()
}

func (o LookupSiteResultOutput) ToLookupSiteResultOutput() LookupSiteResultOutput {
	return o
}

func (o LookupSiteResultOutput) ToLookupSiteResultOutputWithContext(ctx context.Context) LookupSiteResultOutput {
	return o
}

// eTag for concurrency control.
func (o LookupSiteResultOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteResult) *string { return v.ETag }).(pulumi.StringPtrOutput)
}

// Resource Id.
func (o LookupSiteResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSiteResult) string { return v.Id }).(pulumi.StringOutput)
}

// Azure location in which Sites is created.
func (o LookupSiteResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// Name of the VMware site.
func (o LookupSiteResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// Nested properties of VMWare site.
func (o LookupSiteResultOutput) Properties() SitePropertiesResponseOutput {
	return o.ApplyT(func(v LookupSiteResult) SitePropertiesResponse { return v.Properties }).(SitePropertiesResponseOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupSiteResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupSiteResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupSiteResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSiteResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Type of resource. Type = Microsoft.OffAzure/VMWareSites.
func (o LookupSiteResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSiteResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSiteResultOutput{})
}
