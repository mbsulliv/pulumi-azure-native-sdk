// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230301preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets specific OpenID Connect Provider without secrets.
func LookupOpenIdConnectProvider(ctx *pulumi.Context, args *LookupOpenIdConnectProviderArgs, opts ...pulumi.InvokeOption) (*LookupOpenIdConnectProviderResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupOpenIdConnectProviderResult
	err := ctx.Invoke("azure-native:apimanagement/v20230301preview:getOpenIdConnectProvider", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupOpenIdConnectProviderArgs struct {
	// Identifier of the OpenID Connect Provider.
	Opid string `pulumi:"opid"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
}

// OpenId Connect Provider details.
type LookupOpenIdConnectProviderResult struct {
	// Client ID of developer console which is the client application.
	ClientId string `pulumi:"clientId"`
	// Client Secret of developer console which is the client application.
	ClientSecret *string `pulumi:"clientSecret"`
	// User-friendly description of OpenID Connect Provider.
	Description *string `pulumi:"description"`
	// User-friendly OpenID Connect Provider name.
	DisplayName string `pulumi:"displayName"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// Metadata endpoint URI.
	MetadataEndpoint string `pulumi:"metadataEndpoint"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// If true, the Open ID Connect provider will be used in the API documentation in the developer portal. False by default if no value is provided.
	UseInApiDocumentation *bool `pulumi:"useInApiDocumentation"`
	// If true, the Open ID Connect provider may be used in the developer portal test console. True by default if no value is provided.
	UseInTestConsole *bool `pulumi:"useInTestConsole"`
}

func LookupOpenIdConnectProviderOutput(ctx *pulumi.Context, args LookupOpenIdConnectProviderOutputArgs, opts ...pulumi.InvokeOption) LookupOpenIdConnectProviderResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupOpenIdConnectProviderResult, error) {
			args := v.(LookupOpenIdConnectProviderArgs)
			r, err := LookupOpenIdConnectProvider(ctx, &args, opts...)
			var s LookupOpenIdConnectProviderResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupOpenIdConnectProviderResultOutput)
}

type LookupOpenIdConnectProviderOutputArgs struct {
	// Identifier of the OpenID Connect Provider.
	Opid pulumi.StringInput `pulumi:"opid"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupOpenIdConnectProviderOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOpenIdConnectProviderArgs)(nil)).Elem()
}

// OpenId Connect Provider details.
type LookupOpenIdConnectProviderResultOutput struct{ *pulumi.OutputState }

func (LookupOpenIdConnectProviderResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOpenIdConnectProviderResult)(nil)).Elem()
}

func (o LookupOpenIdConnectProviderResultOutput) ToLookupOpenIdConnectProviderResultOutput() LookupOpenIdConnectProviderResultOutput {
	return o
}

func (o LookupOpenIdConnectProviderResultOutput) ToLookupOpenIdConnectProviderResultOutputWithContext(ctx context.Context) LookupOpenIdConnectProviderResultOutput {
	return o
}

// Client ID of developer console which is the client application.
func (o LookupOpenIdConnectProviderResultOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOpenIdConnectProviderResult) string { return v.ClientId }).(pulumi.StringOutput)
}

// Client Secret of developer console which is the client application.
func (o LookupOpenIdConnectProviderResultOutput) ClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupOpenIdConnectProviderResult) *string { return v.ClientSecret }).(pulumi.StringPtrOutput)
}

// User-friendly description of OpenID Connect Provider.
func (o LookupOpenIdConnectProviderResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupOpenIdConnectProviderResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// User-friendly OpenID Connect Provider name.
func (o LookupOpenIdConnectProviderResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOpenIdConnectProviderResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupOpenIdConnectProviderResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOpenIdConnectProviderResult) string { return v.Id }).(pulumi.StringOutput)
}

// Metadata endpoint URI.
func (o LookupOpenIdConnectProviderResultOutput) MetadataEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOpenIdConnectProviderResult) string { return v.MetadataEndpoint }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupOpenIdConnectProviderResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOpenIdConnectProviderResult) string { return v.Name }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupOpenIdConnectProviderResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOpenIdConnectProviderResult) string { return v.Type }).(pulumi.StringOutput)
}

// If true, the Open ID Connect provider will be used in the API documentation in the developer portal. False by default if no value is provided.
func (o LookupOpenIdConnectProviderResultOutput) UseInApiDocumentation() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupOpenIdConnectProviderResult) *bool { return v.UseInApiDocumentation }).(pulumi.BoolPtrOutput)
}

// If true, the Open ID Connect provider may be used in the developer portal test console. True by default if no value is provided.
func (o LookupOpenIdConnectProviderResultOutput) UseInTestConsole() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupOpenIdConnectProviderResult) *bool { return v.UseInTestConsole }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupOpenIdConnectProviderResultOutput{})
}
