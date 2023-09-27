// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230501preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get the Spring Cloud Gateway custom domain.
func LookupGatewayCustomDomain(ctx *pulumi.Context, args *LookupGatewayCustomDomainArgs, opts ...pulumi.InvokeOption) (*LookupGatewayCustomDomainResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupGatewayCustomDomainResult
	err := ctx.Invoke("azure-native:appplatform/v20230501preview:getGatewayCustomDomain", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGatewayCustomDomainArgs struct {
	// The name of the Spring Cloud Gateway custom domain.
	DomainName string `pulumi:"domainName"`
	// The name of Spring Cloud Gateway.
	GatewayName string `pulumi:"gatewayName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Service resource.
	ServiceName string `pulumi:"serviceName"`
}

// Custom domain of the Spring Cloud Gateway
type LookupGatewayCustomDomainResult struct {
	// Fully qualified resource Id for the resource.
	Id string `pulumi:"id"`
	// The name of the resource.
	Name string `pulumi:"name"`
	// The properties of custom domain for Spring Cloud Gateway
	Properties GatewayCustomDomainPropertiesResponse `pulumi:"properties"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource.
	Type string `pulumi:"type"`
}

func LookupGatewayCustomDomainOutput(ctx *pulumi.Context, args LookupGatewayCustomDomainOutputArgs, opts ...pulumi.InvokeOption) LookupGatewayCustomDomainResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGatewayCustomDomainResult, error) {
			args := v.(LookupGatewayCustomDomainArgs)
			r, err := LookupGatewayCustomDomain(ctx, &args, opts...)
			var s LookupGatewayCustomDomainResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGatewayCustomDomainResultOutput)
}

type LookupGatewayCustomDomainOutputArgs struct {
	// The name of the Spring Cloud Gateway custom domain.
	DomainName pulumi.StringInput `pulumi:"domainName"`
	// The name of Spring Cloud Gateway.
	GatewayName pulumi.StringInput `pulumi:"gatewayName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the Service resource.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupGatewayCustomDomainOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGatewayCustomDomainArgs)(nil)).Elem()
}

// Custom domain of the Spring Cloud Gateway
type LookupGatewayCustomDomainResultOutput struct{ *pulumi.OutputState }

func (LookupGatewayCustomDomainResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGatewayCustomDomainResult)(nil)).Elem()
}

func (o LookupGatewayCustomDomainResultOutput) ToLookupGatewayCustomDomainResultOutput() LookupGatewayCustomDomainResultOutput {
	return o
}

func (o LookupGatewayCustomDomainResultOutput) ToLookupGatewayCustomDomainResultOutputWithContext(ctx context.Context) LookupGatewayCustomDomainResultOutput {
	return o
}

func (o LookupGatewayCustomDomainResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupGatewayCustomDomainResult] {
	return pulumix.Output[LookupGatewayCustomDomainResult]{
		OutputState: o.OutputState,
	}
}

// Fully qualified resource Id for the resource.
func (o LookupGatewayCustomDomainResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayCustomDomainResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource.
func (o LookupGatewayCustomDomainResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayCustomDomainResult) string { return v.Name }).(pulumi.StringOutput)
}

// The properties of custom domain for Spring Cloud Gateway
func (o LookupGatewayCustomDomainResultOutput) Properties() GatewayCustomDomainPropertiesResponseOutput {
	return o.ApplyT(func(v LookupGatewayCustomDomainResult) GatewayCustomDomainPropertiesResponse { return v.Properties }).(GatewayCustomDomainPropertiesResponseOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupGatewayCustomDomainResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupGatewayCustomDomainResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource.
func (o LookupGatewayCustomDomainResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayCustomDomainResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGatewayCustomDomainResultOutput{})
}
