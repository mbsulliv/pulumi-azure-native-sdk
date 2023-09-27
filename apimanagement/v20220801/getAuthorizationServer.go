// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets the details of the authorization server specified by its identifier.
func LookupAuthorizationServer(ctx *pulumi.Context, args *LookupAuthorizationServerArgs, opts ...pulumi.InvokeOption) (*LookupAuthorizationServerResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupAuthorizationServerResult
	err := ctx.Invoke("azure-native:apimanagement/v20220801:getAuthorizationServer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAuthorizationServerArgs struct {
	// Identifier of the authorization server.
	Authsid string `pulumi:"authsid"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
}

// External OAuth authorization server settings.
type LookupAuthorizationServerResult struct {
	// OAuth authorization endpoint. See http://tools.ietf.org/html/rfc6749#section-3.2.
	AuthorizationEndpoint string `pulumi:"authorizationEndpoint"`
	// HTTP verbs supported by the authorization endpoint. GET must be always present. POST is optional.
	AuthorizationMethods []string `pulumi:"authorizationMethods"`
	// Specifies the mechanism by which access token is passed to the API.
	BearerTokenSendingMethods []string `pulumi:"bearerTokenSendingMethods"`
	// Method of authentication supported by the token endpoint of this authorization server. Possible values are Basic and/or Body. When Body is specified, client credentials and other parameters are passed within the request body in the application/x-www-form-urlencoded format.
	ClientAuthenticationMethod []string `pulumi:"clientAuthenticationMethod"`
	// Client or app id registered with this authorization server.
	ClientId string `pulumi:"clientId"`
	// Optional reference to a page where client or app registration for this authorization server is performed. Contains absolute URL to entity being referenced.
	ClientRegistrationEndpoint string `pulumi:"clientRegistrationEndpoint"`
	// Client or app secret registered with this authorization server. This property will not be filled on 'GET' operations! Use '/listSecrets' POST request to get the value.
	ClientSecret *string `pulumi:"clientSecret"`
	// Access token scope that is going to be requested by default. Can be overridden at the API level. Should be provided in the form of a string containing space-delimited values.
	DefaultScope *string `pulumi:"defaultScope"`
	// Description of the authorization server. Can contain HTML formatting tags.
	Description *string `pulumi:"description"`
	// User-friendly authorization server name.
	DisplayName string `pulumi:"displayName"`
	// Form of an authorization grant, which the client uses to request the access token.
	GrantTypes []string `pulumi:"grantTypes"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Can be optionally specified when resource owner password grant type is supported by this authorization server. Default resource owner password.
	ResourceOwnerPassword *string `pulumi:"resourceOwnerPassword"`
	// Can be optionally specified when resource owner password grant type is supported by this authorization server. Default resource owner username.
	ResourceOwnerUsername *string `pulumi:"resourceOwnerUsername"`
	// If true, authorization server will include state parameter from the authorization request to its response. Client may use state parameter to raise protocol security.
	SupportState *bool `pulumi:"supportState"`
	// Additional parameters required by the token endpoint of this authorization server represented as an array of JSON objects with name and value string properties, i.e. {"name" : "name value", "value": "a value"}.
	TokenBodyParameters []TokenBodyParameterContractResponse `pulumi:"tokenBodyParameters"`
	// OAuth token endpoint. Contains absolute URI to entity being referenced.
	TokenEndpoint *string `pulumi:"tokenEndpoint"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// If true, the authorization server will be used in the API documentation in the developer portal. False by default if no value is provided.
	UseInApiDocumentation *bool `pulumi:"useInApiDocumentation"`
	// If true, the authorization server may be used in the developer portal test console. True by default if no value is provided.
	UseInTestConsole *bool `pulumi:"useInTestConsole"`
}

func LookupAuthorizationServerOutput(ctx *pulumi.Context, args LookupAuthorizationServerOutputArgs, opts ...pulumi.InvokeOption) LookupAuthorizationServerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAuthorizationServerResult, error) {
			args := v.(LookupAuthorizationServerArgs)
			r, err := LookupAuthorizationServer(ctx, &args, opts...)
			var s LookupAuthorizationServerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAuthorizationServerResultOutput)
}

type LookupAuthorizationServerOutputArgs struct {
	// Identifier of the authorization server.
	Authsid pulumi.StringInput `pulumi:"authsid"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupAuthorizationServerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAuthorizationServerArgs)(nil)).Elem()
}

// External OAuth authorization server settings.
type LookupAuthorizationServerResultOutput struct{ *pulumi.OutputState }

func (LookupAuthorizationServerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAuthorizationServerResult)(nil)).Elem()
}

func (o LookupAuthorizationServerResultOutput) ToLookupAuthorizationServerResultOutput() LookupAuthorizationServerResultOutput {
	return o
}

func (o LookupAuthorizationServerResultOutput) ToLookupAuthorizationServerResultOutputWithContext(ctx context.Context) LookupAuthorizationServerResultOutput {
	return o
}

func (o LookupAuthorizationServerResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupAuthorizationServerResult] {
	return pulumix.Output[LookupAuthorizationServerResult]{
		OutputState: o.OutputState,
	}
}

// OAuth authorization endpoint. See http://tools.ietf.org/html/rfc6749#section-3.2.
func (o LookupAuthorizationServerResultOutput) AuthorizationEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizationServerResult) string { return v.AuthorizationEndpoint }).(pulumi.StringOutput)
}

// HTTP verbs supported by the authorization endpoint. GET must be always present. POST is optional.
func (o LookupAuthorizationServerResultOutput) AuthorizationMethods() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAuthorizationServerResult) []string { return v.AuthorizationMethods }).(pulumi.StringArrayOutput)
}

// Specifies the mechanism by which access token is passed to the API.
func (o LookupAuthorizationServerResultOutput) BearerTokenSendingMethods() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAuthorizationServerResult) []string { return v.BearerTokenSendingMethods }).(pulumi.StringArrayOutput)
}

// Method of authentication supported by the token endpoint of this authorization server. Possible values are Basic and/or Body. When Body is specified, client credentials and other parameters are passed within the request body in the application/x-www-form-urlencoded format.
func (o LookupAuthorizationServerResultOutput) ClientAuthenticationMethod() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAuthorizationServerResult) []string { return v.ClientAuthenticationMethod }).(pulumi.StringArrayOutput)
}

// Client or app id registered with this authorization server.
func (o LookupAuthorizationServerResultOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizationServerResult) string { return v.ClientId }).(pulumi.StringOutput)
}

// Optional reference to a page where client or app registration for this authorization server is performed. Contains absolute URL to entity being referenced.
func (o LookupAuthorizationServerResultOutput) ClientRegistrationEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizationServerResult) string { return v.ClientRegistrationEndpoint }).(pulumi.StringOutput)
}

// Client or app secret registered with this authorization server. This property will not be filled on 'GET' operations! Use '/listSecrets' POST request to get the value.
func (o LookupAuthorizationServerResultOutput) ClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAuthorizationServerResult) *string { return v.ClientSecret }).(pulumi.StringPtrOutput)
}

// Access token scope that is going to be requested by default. Can be overridden at the API level. Should be provided in the form of a string containing space-delimited values.
func (o LookupAuthorizationServerResultOutput) DefaultScope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAuthorizationServerResult) *string { return v.DefaultScope }).(pulumi.StringPtrOutput)
}

// Description of the authorization server. Can contain HTML formatting tags.
func (o LookupAuthorizationServerResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAuthorizationServerResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// User-friendly authorization server name.
func (o LookupAuthorizationServerResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizationServerResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// Form of an authorization grant, which the client uses to request the access token.
func (o LookupAuthorizationServerResultOutput) GrantTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAuthorizationServerResult) []string { return v.GrantTypes }).(pulumi.StringArrayOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupAuthorizationServerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizationServerResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupAuthorizationServerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizationServerResult) string { return v.Name }).(pulumi.StringOutput)
}

// Can be optionally specified when resource owner password grant type is supported by this authorization server. Default resource owner password.
func (o LookupAuthorizationServerResultOutput) ResourceOwnerPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAuthorizationServerResult) *string { return v.ResourceOwnerPassword }).(pulumi.StringPtrOutput)
}

// Can be optionally specified when resource owner password grant type is supported by this authorization server. Default resource owner username.
func (o LookupAuthorizationServerResultOutput) ResourceOwnerUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAuthorizationServerResult) *string { return v.ResourceOwnerUsername }).(pulumi.StringPtrOutput)
}

// If true, authorization server will include state parameter from the authorization request to its response. Client may use state parameter to raise protocol security.
func (o LookupAuthorizationServerResultOutput) SupportState() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupAuthorizationServerResult) *bool { return v.SupportState }).(pulumi.BoolPtrOutput)
}

// Additional parameters required by the token endpoint of this authorization server represented as an array of JSON objects with name and value string properties, i.e. {"name" : "name value", "value": "a value"}.
func (o LookupAuthorizationServerResultOutput) TokenBodyParameters() TokenBodyParameterContractResponseArrayOutput {
	return o.ApplyT(func(v LookupAuthorizationServerResult) []TokenBodyParameterContractResponse {
		return v.TokenBodyParameters
	}).(TokenBodyParameterContractResponseArrayOutput)
}

// OAuth token endpoint. Contains absolute URI to entity being referenced.
func (o LookupAuthorizationServerResultOutput) TokenEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAuthorizationServerResult) *string { return v.TokenEndpoint }).(pulumi.StringPtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupAuthorizationServerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizationServerResult) string { return v.Type }).(pulumi.StringOutput)
}

// If true, the authorization server will be used in the API documentation in the developer portal. False by default if no value is provided.
func (o LookupAuthorizationServerResultOutput) UseInApiDocumentation() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupAuthorizationServerResult) *bool { return v.UseInApiDocumentation }).(pulumi.BoolPtrOutput)
}

// If true, the authorization server may be used in the developer portal test console. True by default if no value is provided.
func (o LookupAuthorizationServerResultOutput) UseInTestConsole() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupAuthorizationServerResult) *bool { return v.UseInTestConsole }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAuthorizationServerResultOutput{})
}
