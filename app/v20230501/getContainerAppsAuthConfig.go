// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configuration settings for the Azure ContainerApp Service Authentication / Authorization feature.
func LookupContainerAppsAuthConfig(ctx *pulumi.Context, args *LookupContainerAppsAuthConfigArgs, opts ...pulumi.InvokeOption) (*LookupContainerAppsAuthConfigResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupContainerAppsAuthConfigResult
	err := ctx.Invoke("azure-native:app/v20230501:getContainerAppsAuthConfig", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupContainerAppsAuthConfigArgs struct {
	// Name of the Container App AuthConfig.
	AuthConfigName string `pulumi:"authConfigName"`
	// Name of the Container App.
	ContainerAppName string `pulumi:"containerAppName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Configuration settings for the Azure ContainerApp Service Authentication / Authorization feature.
type LookupContainerAppsAuthConfigResult struct {
	// The configuration settings that determines the validation flow of users using  Service Authentication/Authorization.
	GlobalValidation *GlobalValidationResponse `pulumi:"globalValidation"`
	// The configuration settings of the HTTP requests for authentication and authorization requests made against ContainerApp Service Authentication/Authorization.
	HttpSettings *HttpSettingsResponse `pulumi:"httpSettings"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The configuration settings of each of the identity providers used to configure ContainerApp Service Authentication/Authorization.
	IdentityProviders *IdentityProvidersResponse `pulumi:"identityProviders"`
	// The configuration settings of the login flow of users using ContainerApp Service Authentication/Authorization.
	Login *LoginResponse `pulumi:"login"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The configuration settings of the platform of ContainerApp Service Authentication/Authorization.
	Platform *AuthPlatformResponse `pulumi:"platform"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupContainerAppsAuthConfigOutput(ctx *pulumi.Context, args LookupContainerAppsAuthConfigOutputArgs, opts ...pulumi.InvokeOption) LookupContainerAppsAuthConfigResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupContainerAppsAuthConfigResult, error) {
			args := v.(LookupContainerAppsAuthConfigArgs)
			r, err := LookupContainerAppsAuthConfig(ctx, &args, opts...)
			var s LookupContainerAppsAuthConfigResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupContainerAppsAuthConfigResultOutput)
}

type LookupContainerAppsAuthConfigOutputArgs struct {
	// Name of the Container App AuthConfig.
	AuthConfigName pulumi.StringInput `pulumi:"authConfigName"`
	// Name of the Container App.
	ContainerAppName pulumi.StringInput `pulumi:"containerAppName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupContainerAppsAuthConfigOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupContainerAppsAuthConfigArgs)(nil)).Elem()
}

// Configuration settings for the Azure ContainerApp Service Authentication / Authorization feature.
type LookupContainerAppsAuthConfigResultOutput struct{ *pulumi.OutputState }

func (LookupContainerAppsAuthConfigResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupContainerAppsAuthConfigResult)(nil)).Elem()
}

func (o LookupContainerAppsAuthConfigResultOutput) ToLookupContainerAppsAuthConfigResultOutput() LookupContainerAppsAuthConfigResultOutput {
	return o
}

func (o LookupContainerAppsAuthConfigResultOutput) ToLookupContainerAppsAuthConfigResultOutputWithContext(ctx context.Context) LookupContainerAppsAuthConfigResultOutput {
	return o
}

// The configuration settings that determines the validation flow of users using  Service Authentication/Authorization.
func (o LookupContainerAppsAuthConfigResultOutput) GlobalValidation() GlobalValidationResponsePtrOutput {
	return o.ApplyT(func(v LookupContainerAppsAuthConfigResult) *GlobalValidationResponse { return v.GlobalValidation }).(GlobalValidationResponsePtrOutput)
}

// The configuration settings of the HTTP requests for authentication and authorization requests made against ContainerApp Service Authentication/Authorization.
func (o LookupContainerAppsAuthConfigResultOutput) HttpSettings() HttpSettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupContainerAppsAuthConfigResult) *HttpSettingsResponse { return v.HttpSettings }).(HttpSettingsResponsePtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupContainerAppsAuthConfigResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerAppsAuthConfigResult) string { return v.Id }).(pulumi.StringOutput)
}

// The configuration settings of each of the identity providers used to configure ContainerApp Service Authentication/Authorization.
func (o LookupContainerAppsAuthConfigResultOutput) IdentityProviders() IdentityProvidersResponsePtrOutput {
	return o.ApplyT(func(v LookupContainerAppsAuthConfigResult) *IdentityProvidersResponse { return v.IdentityProviders }).(IdentityProvidersResponsePtrOutput)
}

// The configuration settings of the login flow of users using ContainerApp Service Authentication/Authorization.
func (o LookupContainerAppsAuthConfigResultOutput) Login() LoginResponsePtrOutput {
	return o.ApplyT(func(v LookupContainerAppsAuthConfigResult) *LoginResponse { return v.Login }).(LoginResponsePtrOutput)
}

// The name of the resource
func (o LookupContainerAppsAuthConfigResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerAppsAuthConfigResult) string { return v.Name }).(pulumi.StringOutput)
}

// The configuration settings of the platform of ContainerApp Service Authentication/Authorization.
func (o LookupContainerAppsAuthConfigResultOutput) Platform() AuthPlatformResponsePtrOutput {
	return o.ApplyT(func(v LookupContainerAppsAuthConfigResult) *AuthPlatformResponse { return v.Platform }).(AuthPlatformResponsePtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupContainerAppsAuthConfigResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupContainerAppsAuthConfigResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupContainerAppsAuthConfigResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerAppsAuthConfigResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupContainerAppsAuthConfigResultOutput{})
}
