// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apimanagement

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets the details of the authorization provider specified by its identifier.
// Azure REST API version: 2022-08-01.
func LookupAuthorizationProvider(ctx *pulumi.Context, args *LookupAuthorizationProviderArgs, opts ...pulumi.InvokeOption) (*LookupAuthorizationProviderResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupAuthorizationProviderResult
	err := ctx.Invoke("azure-native:apimanagement:getAuthorizationProvider", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAuthorizationProviderArgs struct {
	// Identifier of the authorization provider.
	AuthorizationProviderId string `pulumi:"authorizationProviderId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
}

// Authorization Provider contract.
type LookupAuthorizationProviderResult struct {
	// Authorization Provider name. Must be 1 to 300 characters long.
	DisplayName *string `pulumi:"displayName"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// Identity provider name. Must be 1 to 300 characters long.
	IdentityProvider *string `pulumi:"identityProvider"`
	// The name of the resource
	Name string `pulumi:"name"`
	// OAuth2 settings
	Oauth2 *AuthorizationProviderOAuth2SettingsResponse `pulumi:"oauth2"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupAuthorizationProviderOutput(ctx *pulumi.Context, args LookupAuthorizationProviderOutputArgs, opts ...pulumi.InvokeOption) LookupAuthorizationProviderResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAuthorizationProviderResult, error) {
			args := v.(LookupAuthorizationProviderArgs)
			r, err := LookupAuthorizationProvider(ctx, &args, opts...)
			var s LookupAuthorizationProviderResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAuthorizationProviderResultOutput)
}

type LookupAuthorizationProviderOutputArgs struct {
	// Identifier of the authorization provider.
	AuthorizationProviderId pulumi.StringInput `pulumi:"authorizationProviderId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupAuthorizationProviderOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAuthorizationProviderArgs)(nil)).Elem()
}

// Authorization Provider contract.
type LookupAuthorizationProviderResultOutput struct{ *pulumi.OutputState }

func (LookupAuthorizationProviderResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAuthorizationProviderResult)(nil)).Elem()
}

func (o LookupAuthorizationProviderResultOutput) ToLookupAuthorizationProviderResultOutput() LookupAuthorizationProviderResultOutput {
	return o
}

func (o LookupAuthorizationProviderResultOutput) ToLookupAuthorizationProviderResultOutputWithContext(ctx context.Context) LookupAuthorizationProviderResultOutput {
	return o
}

func (o LookupAuthorizationProviderResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupAuthorizationProviderResult] {
	return pulumix.Output[LookupAuthorizationProviderResult]{
		OutputState: o.OutputState,
	}
}

// Authorization Provider name. Must be 1 to 300 characters long.
func (o LookupAuthorizationProviderResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAuthorizationProviderResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupAuthorizationProviderResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizationProviderResult) string { return v.Id }).(pulumi.StringOutput)
}

// Identity provider name. Must be 1 to 300 characters long.
func (o LookupAuthorizationProviderResultOutput) IdentityProvider() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAuthorizationProviderResult) *string { return v.IdentityProvider }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o LookupAuthorizationProviderResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizationProviderResult) string { return v.Name }).(pulumi.StringOutput)
}

// OAuth2 settings
func (o LookupAuthorizationProviderResultOutput) Oauth2() AuthorizationProviderOAuth2SettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupAuthorizationProviderResult) *AuthorizationProviderOAuth2SettingsResponse {
		return v.Oauth2
	}).(AuthorizationProviderOAuth2SettingsResponsePtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupAuthorizationProviderResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizationProviderResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAuthorizationProviderResultOutput{})
}
