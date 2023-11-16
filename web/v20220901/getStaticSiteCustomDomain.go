// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220901

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Description for Gets an existing custom domain for a particular static site.
func LookupStaticSiteCustomDomain(ctx *pulumi.Context, args *LookupStaticSiteCustomDomainArgs, opts ...pulumi.InvokeOption) (*LookupStaticSiteCustomDomainResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupStaticSiteCustomDomainResult
	err := ctx.Invoke("azure-native:web/v20220901:getStaticSiteCustomDomain", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStaticSiteCustomDomainArgs struct {
	// The custom domain name.
	DomainName string `pulumi:"domainName"`
	// Name of the static site resource to search in.
	Name string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Static Site Custom Domain Overview ARM resource.
type LookupStaticSiteCustomDomainResult struct {
	// The date and time on which the custom domain was created for the static site.
	CreatedOn string `pulumi:"createdOn"`
	// The domain name for the static site custom domain.
	DomainName   string `pulumi:"domainName"`
	ErrorMessage string `pulumi:"errorMessage"`
	// Resource Id.
	Id string `pulumi:"id"`
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Resource Name.
	Name string `pulumi:"name"`
	// The status of the custom domain
	Status string `pulumi:"status"`
	// Resource type.
	Type string `pulumi:"type"`
	// The TXT record validation token
	ValidationToken string `pulumi:"validationToken"`
}

func LookupStaticSiteCustomDomainOutput(ctx *pulumi.Context, args LookupStaticSiteCustomDomainOutputArgs, opts ...pulumi.InvokeOption) LookupStaticSiteCustomDomainResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupStaticSiteCustomDomainResult, error) {
			args := v.(LookupStaticSiteCustomDomainArgs)
			r, err := LookupStaticSiteCustomDomain(ctx, &args, opts...)
			var s LookupStaticSiteCustomDomainResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupStaticSiteCustomDomainResultOutput)
}

type LookupStaticSiteCustomDomainOutputArgs struct {
	// The custom domain name.
	DomainName pulumi.StringInput `pulumi:"domainName"`
	// Name of the static site resource to search in.
	Name pulumi.StringInput `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupStaticSiteCustomDomainOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStaticSiteCustomDomainArgs)(nil)).Elem()
}

// Static Site Custom Domain Overview ARM resource.
type LookupStaticSiteCustomDomainResultOutput struct{ *pulumi.OutputState }

func (LookupStaticSiteCustomDomainResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStaticSiteCustomDomainResult)(nil)).Elem()
}

func (o LookupStaticSiteCustomDomainResultOutput) ToLookupStaticSiteCustomDomainResultOutput() LookupStaticSiteCustomDomainResultOutput {
	return o
}

func (o LookupStaticSiteCustomDomainResultOutput) ToLookupStaticSiteCustomDomainResultOutputWithContext(ctx context.Context) LookupStaticSiteCustomDomainResultOutput {
	return o
}

// The date and time on which the custom domain was created for the static site.
func (o LookupStaticSiteCustomDomainResultOutput) CreatedOn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticSiteCustomDomainResult) string { return v.CreatedOn }).(pulumi.StringOutput)
}

// The domain name for the static site custom domain.
func (o LookupStaticSiteCustomDomainResultOutput) DomainName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticSiteCustomDomainResult) string { return v.DomainName }).(pulumi.StringOutput)
}

func (o LookupStaticSiteCustomDomainResultOutput) ErrorMessage() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticSiteCustomDomainResult) string { return v.ErrorMessage }).(pulumi.StringOutput)
}

// Resource Id.
func (o LookupStaticSiteCustomDomainResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticSiteCustomDomainResult) string { return v.Id }).(pulumi.StringOutput)
}

// Kind of resource.
func (o LookupStaticSiteCustomDomainResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStaticSiteCustomDomainResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// Resource Name.
func (o LookupStaticSiteCustomDomainResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticSiteCustomDomainResult) string { return v.Name }).(pulumi.StringOutput)
}

// The status of the custom domain
func (o LookupStaticSiteCustomDomainResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticSiteCustomDomainResult) string { return v.Status }).(pulumi.StringOutput)
}

// Resource type.
func (o LookupStaticSiteCustomDomainResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticSiteCustomDomainResult) string { return v.Type }).(pulumi.StringOutput)
}

// The TXT record validation token
func (o LookupStaticSiteCustomDomainResultOutput) ValidationToken() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticSiteCustomDomainResult) string { return v.ValidationToken }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupStaticSiteCustomDomainResultOutput{})
}
