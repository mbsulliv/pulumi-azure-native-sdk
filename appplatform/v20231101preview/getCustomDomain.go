// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20231101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get the custom domain of one lifecycle application.
func LookupCustomDomain(ctx *pulumi.Context, args *LookupCustomDomainArgs, opts ...pulumi.InvokeOption) (*LookupCustomDomainResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupCustomDomainResult
	err := ctx.Invoke("azure-native:appplatform/v20231101preview:getCustomDomain", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCustomDomainArgs struct {
	// The name of the App resource.
	AppName string `pulumi:"appName"`
	// The name of the custom domain resource.
	DomainName string `pulumi:"domainName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Service resource.
	ServiceName string `pulumi:"serviceName"`
}

// Custom domain resource payload.
type LookupCustomDomainResult struct {
	// Fully qualified resource Id for the resource.
	Id string `pulumi:"id"`
	// The name of the resource.
	Name string `pulumi:"name"`
	// Properties of the custom domain resource.
	Properties CustomDomainPropertiesResponse `pulumi:"properties"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource.
	Type string `pulumi:"type"`
}

func LookupCustomDomainOutput(ctx *pulumi.Context, args LookupCustomDomainOutputArgs, opts ...pulumi.InvokeOption) LookupCustomDomainResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCustomDomainResult, error) {
			args := v.(LookupCustomDomainArgs)
			r, err := LookupCustomDomain(ctx, &args, opts...)
			var s LookupCustomDomainResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCustomDomainResultOutput)
}

type LookupCustomDomainOutputArgs struct {
	// The name of the App resource.
	AppName pulumi.StringInput `pulumi:"appName"`
	// The name of the custom domain resource.
	DomainName pulumi.StringInput `pulumi:"domainName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the Service resource.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupCustomDomainOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCustomDomainArgs)(nil)).Elem()
}

// Custom domain resource payload.
type LookupCustomDomainResultOutput struct{ *pulumi.OutputState }

func (LookupCustomDomainResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCustomDomainResult)(nil)).Elem()
}

func (o LookupCustomDomainResultOutput) ToLookupCustomDomainResultOutput() LookupCustomDomainResultOutput {
	return o
}

func (o LookupCustomDomainResultOutput) ToLookupCustomDomainResultOutputWithContext(ctx context.Context) LookupCustomDomainResultOutput {
	return o
}

func (o LookupCustomDomainResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupCustomDomainResult] {
	return pulumix.Output[LookupCustomDomainResult]{
		OutputState: o.OutputState,
	}
}

// Fully qualified resource Id for the resource.
func (o LookupCustomDomainResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomDomainResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource.
func (o LookupCustomDomainResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomDomainResult) string { return v.Name }).(pulumi.StringOutput)
}

// Properties of the custom domain resource.
func (o LookupCustomDomainResultOutput) Properties() CustomDomainPropertiesResponseOutput {
	return o.ApplyT(func(v LookupCustomDomainResult) CustomDomainPropertiesResponse { return v.Properties }).(CustomDomainPropertiesResponseOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupCustomDomainResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupCustomDomainResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource.
func (o LookupCustomDomainResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomDomainResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCustomDomainResultOutput{})
}
