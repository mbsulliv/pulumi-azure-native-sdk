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

// Gets the client secret details of the authorization server.
// Azure REST API version: 2022-08-01.
//
// Other available API versions: 2022-09-01-preview, 2023-03-01-preview.
func ListAuthorizationServerSecrets(ctx *pulumi.Context, args *ListAuthorizationServerSecretsArgs, opts ...pulumi.InvokeOption) (*ListAuthorizationServerSecretsResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListAuthorizationServerSecretsResult
	err := ctx.Invoke("azure-native:apimanagement:listAuthorizationServerSecrets", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListAuthorizationServerSecretsArgs struct {
	// Identifier of the authorization server.
	Authsid string `pulumi:"authsid"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
}

// OAuth Server Secrets Contract.
type ListAuthorizationServerSecretsResult struct {
	// oAuth Authorization Server Secrets.
	ClientSecret *string `pulumi:"clientSecret"`
	// Can be optionally specified when resource owner password grant type is supported by this authorization server. Default resource owner password.
	ResourceOwnerPassword *string `pulumi:"resourceOwnerPassword"`
	// Can be optionally specified when resource owner password grant type is supported by this authorization server. Default resource owner username.
	ResourceOwnerUsername *string `pulumi:"resourceOwnerUsername"`
}

func ListAuthorizationServerSecretsOutput(ctx *pulumi.Context, args ListAuthorizationServerSecretsOutputArgs, opts ...pulumi.InvokeOption) ListAuthorizationServerSecretsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListAuthorizationServerSecretsResult, error) {
			args := v.(ListAuthorizationServerSecretsArgs)
			r, err := ListAuthorizationServerSecrets(ctx, &args, opts...)
			var s ListAuthorizationServerSecretsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListAuthorizationServerSecretsResultOutput)
}

type ListAuthorizationServerSecretsOutputArgs struct {
	// Identifier of the authorization server.
	Authsid pulumi.StringInput `pulumi:"authsid"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (ListAuthorizationServerSecretsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListAuthorizationServerSecretsArgs)(nil)).Elem()
}

// OAuth Server Secrets Contract.
type ListAuthorizationServerSecretsResultOutput struct{ *pulumi.OutputState }

func (ListAuthorizationServerSecretsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListAuthorizationServerSecretsResult)(nil)).Elem()
}

func (o ListAuthorizationServerSecretsResultOutput) ToListAuthorizationServerSecretsResultOutput() ListAuthorizationServerSecretsResultOutput {
	return o
}

func (o ListAuthorizationServerSecretsResultOutput) ToListAuthorizationServerSecretsResultOutputWithContext(ctx context.Context) ListAuthorizationServerSecretsResultOutput {
	return o
}

func (o ListAuthorizationServerSecretsResultOutput) ToOutput(ctx context.Context) pulumix.Output[ListAuthorizationServerSecretsResult] {
	return pulumix.Output[ListAuthorizationServerSecretsResult]{
		OutputState: o.OutputState,
	}
}

// oAuth Authorization Server Secrets.
func (o ListAuthorizationServerSecretsResultOutput) ClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListAuthorizationServerSecretsResult) *string { return v.ClientSecret }).(pulumi.StringPtrOutput)
}

// Can be optionally specified when resource owner password grant type is supported by this authorization server. Default resource owner password.
func (o ListAuthorizationServerSecretsResultOutput) ResourceOwnerPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListAuthorizationServerSecretsResult) *string { return v.ResourceOwnerPassword }).(pulumi.StringPtrOutput)
}

// Can be optionally specified when resource owner password grant type is supported by this authorization server. Default resource owner username.
func (o ListAuthorizationServerSecretsResultOutput) ResourceOwnerUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListAuthorizationServerSecretsResult) *string { return v.ResourceOwnerUsername }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListAuthorizationServerSecretsResultOutput{})
}
