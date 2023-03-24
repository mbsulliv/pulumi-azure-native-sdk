// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the client secret details of the Identity Provider.
func ListIdentityProviderSecrets(ctx *pulumi.Context, args *ListIdentityProviderSecretsArgs, opts ...pulumi.InvokeOption) (*ListIdentityProviderSecretsResult, error) {
	var rv ListIdentityProviderSecretsResult
	err := ctx.Invoke("azure-native:apimanagement/v20220801:listIdentityProviderSecrets", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListIdentityProviderSecretsArgs struct {
	// Identity Provider Type identifier.
	IdentityProviderName string `pulumi:"identityProviderName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
}

// Client or app secret used in IdentityProviders, Aad, OpenID or OAuth.
type ListIdentityProviderSecretsResult struct {
	// Client or app secret used in IdentityProviders, Aad, OpenID or OAuth.
	ClientSecret *string `pulumi:"clientSecret"`
}

func ListIdentityProviderSecretsOutput(ctx *pulumi.Context, args ListIdentityProviderSecretsOutputArgs, opts ...pulumi.InvokeOption) ListIdentityProviderSecretsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListIdentityProviderSecretsResult, error) {
			args := v.(ListIdentityProviderSecretsArgs)
			r, err := ListIdentityProviderSecrets(ctx, &args, opts...)
			var s ListIdentityProviderSecretsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListIdentityProviderSecretsResultOutput)
}

type ListIdentityProviderSecretsOutputArgs struct {
	// Identity Provider Type identifier.
	IdentityProviderName pulumi.StringInput `pulumi:"identityProviderName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (ListIdentityProviderSecretsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListIdentityProviderSecretsArgs)(nil)).Elem()
}

// Client or app secret used in IdentityProviders, Aad, OpenID or OAuth.
type ListIdentityProviderSecretsResultOutput struct{ *pulumi.OutputState }

func (ListIdentityProviderSecretsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListIdentityProviderSecretsResult)(nil)).Elem()
}

func (o ListIdentityProviderSecretsResultOutput) ToListIdentityProviderSecretsResultOutput() ListIdentityProviderSecretsResultOutput {
	return o
}

func (o ListIdentityProviderSecretsResultOutput) ToListIdentityProviderSecretsResultOutputWithContext(ctx context.Context) ListIdentityProviderSecretsResultOutput {
	return o
}

// Client or app secret used in IdentityProviders, Aad, OpenID or OAuth.
func (o ListIdentityProviderSecretsResultOutput) ClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListIdentityProviderSecretsResult) *string { return v.ClientSecret }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListIdentityProviderSecretsResultOutput{})
}