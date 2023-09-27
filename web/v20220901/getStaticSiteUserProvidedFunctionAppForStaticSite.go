// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220901

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Description for Gets the details of the user provided function app registered with a static site
func LookupStaticSiteUserProvidedFunctionAppForStaticSite(ctx *pulumi.Context, args *LookupStaticSiteUserProvidedFunctionAppForStaticSiteArgs, opts ...pulumi.InvokeOption) (*LookupStaticSiteUserProvidedFunctionAppForStaticSiteResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupStaticSiteUserProvidedFunctionAppForStaticSiteResult
	err := ctx.Invoke("azure-native:web/v20220901:getStaticSiteUserProvidedFunctionAppForStaticSite", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStaticSiteUserProvidedFunctionAppForStaticSiteArgs struct {
	// Name of the function app registered with the static site.
	FunctionAppName string `pulumi:"functionAppName"`
	// Name of the static site.
	Name string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Static Site User Provided Function App ARM resource.
type LookupStaticSiteUserProvidedFunctionAppForStaticSiteResult struct {
	// The date and time on which the function app was registered with the static site.
	CreatedOn string `pulumi:"createdOn"`
	// The region of the function app registered with the static site
	FunctionAppRegion *string `pulumi:"functionAppRegion"`
	// The resource id of the function app registered with the static site
	FunctionAppResourceId *string `pulumi:"functionAppResourceId"`
	// Resource Id.
	Id string `pulumi:"id"`
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Resource Name.
	Name string `pulumi:"name"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupStaticSiteUserProvidedFunctionAppForStaticSiteOutput(ctx *pulumi.Context, args LookupStaticSiteUserProvidedFunctionAppForStaticSiteOutputArgs, opts ...pulumi.InvokeOption) LookupStaticSiteUserProvidedFunctionAppForStaticSiteResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupStaticSiteUserProvidedFunctionAppForStaticSiteResult, error) {
			args := v.(LookupStaticSiteUserProvidedFunctionAppForStaticSiteArgs)
			r, err := LookupStaticSiteUserProvidedFunctionAppForStaticSite(ctx, &args, opts...)
			var s LookupStaticSiteUserProvidedFunctionAppForStaticSiteResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupStaticSiteUserProvidedFunctionAppForStaticSiteResultOutput)
}

type LookupStaticSiteUserProvidedFunctionAppForStaticSiteOutputArgs struct {
	// Name of the function app registered with the static site.
	FunctionAppName pulumi.StringInput `pulumi:"functionAppName"`
	// Name of the static site.
	Name pulumi.StringInput `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupStaticSiteUserProvidedFunctionAppForStaticSiteOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStaticSiteUserProvidedFunctionAppForStaticSiteArgs)(nil)).Elem()
}

// Static Site User Provided Function App ARM resource.
type LookupStaticSiteUserProvidedFunctionAppForStaticSiteResultOutput struct{ *pulumi.OutputState }

func (LookupStaticSiteUserProvidedFunctionAppForStaticSiteResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStaticSiteUserProvidedFunctionAppForStaticSiteResult)(nil)).Elem()
}

func (o LookupStaticSiteUserProvidedFunctionAppForStaticSiteResultOutput) ToLookupStaticSiteUserProvidedFunctionAppForStaticSiteResultOutput() LookupStaticSiteUserProvidedFunctionAppForStaticSiteResultOutput {
	return o
}

func (o LookupStaticSiteUserProvidedFunctionAppForStaticSiteResultOutput) ToLookupStaticSiteUserProvidedFunctionAppForStaticSiteResultOutputWithContext(ctx context.Context) LookupStaticSiteUserProvidedFunctionAppForStaticSiteResultOutput {
	return o
}

func (o LookupStaticSiteUserProvidedFunctionAppForStaticSiteResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupStaticSiteUserProvidedFunctionAppForStaticSiteResult] {
	return pulumix.Output[LookupStaticSiteUserProvidedFunctionAppForStaticSiteResult]{
		OutputState: o.OutputState,
	}
}

// The date and time on which the function app was registered with the static site.
func (o LookupStaticSiteUserProvidedFunctionAppForStaticSiteResultOutput) CreatedOn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticSiteUserProvidedFunctionAppForStaticSiteResult) string { return v.CreatedOn }).(pulumi.StringOutput)
}

// The region of the function app registered with the static site
func (o LookupStaticSiteUserProvidedFunctionAppForStaticSiteResultOutput) FunctionAppRegion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStaticSiteUserProvidedFunctionAppForStaticSiteResult) *string { return v.FunctionAppRegion }).(pulumi.StringPtrOutput)
}

// The resource id of the function app registered with the static site
func (o LookupStaticSiteUserProvidedFunctionAppForStaticSiteResultOutput) FunctionAppResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStaticSiteUserProvidedFunctionAppForStaticSiteResult) *string {
		return v.FunctionAppResourceId
	}).(pulumi.StringPtrOutput)
}

// Resource Id.
func (o LookupStaticSiteUserProvidedFunctionAppForStaticSiteResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticSiteUserProvidedFunctionAppForStaticSiteResult) string { return v.Id }).(pulumi.StringOutput)
}

// Kind of resource.
func (o LookupStaticSiteUserProvidedFunctionAppForStaticSiteResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStaticSiteUserProvidedFunctionAppForStaticSiteResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// Resource Name.
func (o LookupStaticSiteUserProvidedFunctionAppForStaticSiteResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticSiteUserProvidedFunctionAppForStaticSiteResult) string { return v.Name }).(pulumi.StringOutput)
}

// Resource type.
func (o LookupStaticSiteUserProvidedFunctionAppForStaticSiteResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticSiteUserProvidedFunctionAppForStaticSiteResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupStaticSiteUserProvidedFunctionAppForStaticSiteResultOutput{})
}
