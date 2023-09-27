// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230701preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get the API portal custom domain.
func LookupApiPortalCustomDomain(ctx *pulumi.Context, args *LookupApiPortalCustomDomainArgs, opts ...pulumi.InvokeOption) (*LookupApiPortalCustomDomainResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupApiPortalCustomDomainResult
	err := ctx.Invoke("azure-native:appplatform/v20230701preview:getApiPortalCustomDomain", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApiPortalCustomDomainArgs struct {
	// The name of API portal.
	ApiPortalName string `pulumi:"apiPortalName"`
	// The name of the API portal custom domain.
	DomainName string `pulumi:"domainName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Service resource.
	ServiceName string `pulumi:"serviceName"`
}

// Custom domain of the API portal
type LookupApiPortalCustomDomainResult struct {
	// Fully qualified resource Id for the resource.
	Id string `pulumi:"id"`
	// The name of the resource.
	Name string `pulumi:"name"`
	// The properties of custom domain for API portal
	Properties ApiPortalCustomDomainPropertiesResponse `pulumi:"properties"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource.
	Type string `pulumi:"type"`
}

func LookupApiPortalCustomDomainOutput(ctx *pulumi.Context, args LookupApiPortalCustomDomainOutputArgs, opts ...pulumi.InvokeOption) LookupApiPortalCustomDomainResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupApiPortalCustomDomainResult, error) {
			args := v.(LookupApiPortalCustomDomainArgs)
			r, err := LookupApiPortalCustomDomain(ctx, &args, opts...)
			var s LookupApiPortalCustomDomainResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupApiPortalCustomDomainResultOutput)
}

type LookupApiPortalCustomDomainOutputArgs struct {
	// The name of API portal.
	ApiPortalName pulumi.StringInput `pulumi:"apiPortalName"`
	// The name of the API portal custom domain.
	DomainName pulumi.StringInput `pulumi:"domainName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the Service resource.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupApiPortalCustomDomainOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApiPortalCustomDomainArgs)(nil)).Elem()
}

// Custom domain of the API portal
type LookupApiPortalCustomDomainResultOutput struct{ *pulumi.OutputState }

func (LookupApiPortalCustomDomainResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApiPortalCustomDomainResult)(nil)).Elem()
}

func (o LookupApiPortalCustomDomainResultOutput) ToLookupApiPortalCustomDomainResultOutput() LookupApiPortalCustomDomainResultOutput {
	return o
}

func (o LookupApiPortalCustomDomainResultOutput) ToLookupApiPortalCustomDomainResultOutputWithContext(ctx context.Context) LookupApiPortalCustomDomainResultOutput {
	return o
}

func (o LookupApiPortalCustomDomainResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupApiPortalCustomDomainResult] {
	return pulumix.Output[LookupApiPortalCustomDomainResult]{
		OutputState: o.OutputState,
	}
}

// Fully qualified resource Id for the resource.
func (o LookupApiPortalCustomDomainResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiPortalCustomDomainResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource.
func (o LookupApiPortalCustomDomainResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiPortalCustomDomainResult) string { return v.Name }).(pulumi.StringOutput)
}

// The properties of custom domain for API portal
func (o LookupApiPortalCustomDomainResultOutput) Properties() ApiPortalCustomDomainPropertiesResponseOutput {
	return o.ApplyT(func(v LookupApiPortalCustomDomainResult) ApiPortalCustomDomainPropertiesResponse { return v.Properties }).(ApiPortalCustomDomainPropertiesResponseOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupApiPortalCustomDomainResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupApiPortalCustomDomainResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource.
func (o LookupApiPortalCustomDomainResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiPortalCustomDomainResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApiPortalCustomDomainResultOutput{})
}
