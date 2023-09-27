// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230301preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets the Single-Sign-On token for the API Management Service which is valid for 5 Minutes.
func GetApiManagementServiceSsoToken(ctx *pulumi.Context, args *GetApiManagementServiceSsoTokenArgs, opts ...pulumi.InvokeOption) (*GetApiManagementServiceSsoTokenResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv GetApiManagementServiceSsoTokenResult
	err := ctx.Invoke("azure-native:apimanagement/v20230301preview:getApiManagementServiceSsoToken", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetApiManagementServiceSsoTokenArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
}

// The response of the GetSsoToken operation.
type GetApiManagementServiceSsoTokenResult struct {
	// Redirect URL to the Publisher Portal containing the SSO token.
	RedirectUri *string `pulumi:"redirectUri"`
}

func GetApiManagementServiceSsoTokenOutput(ctx *pulumi.Context, args GetApiManagementServiceSsoTokenOutputArgs, opts ...pulumi.InvokeOption) GetApiManagementServiceSsoTokenResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetApiManagementServiceSsoTokenResult, error) {
			args := v.(GetApiManagementServiceSsoTokenArgs)
			r, err := GetApiManagementServiceSsoToken(ctx, &args, opts...)
			var s GetApiManagementServiceSsoTokenResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetApiManagementServiceSsoTokenResultOutput)
}

type GetApiManagementServiceSsoTokenOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (GetApiManagementServiceSsoTokenOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetApiManagementServiceSsoTokenArgs)(nil)).Elem()
}

// The response of the GetSsoToken operation.
type GetApiManagementServiceSsoTokenResultOutput struct{ *pulumi.OutputState }

func (GetApiManagementServiceSsoTokenResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetApiManagementServiceSsoTokenResult)(nil)).Elem()
}

func (o GetApiManagementServiceSsoTokenResultOutput) ToGetApiManagementServiceSsoTokenResultOutput() GetApiManagementServiceSsoTokenResultOutput {
	return o
}

func (o GetApiManagementServiceSsoTokenResultOutput) ToGetApiManagementServiceSsoTokenResultOutputWithContext(ctx context.Context) GetApiManagementServiceSsoTokenResultOutput {
	return o
}

func (o GetApiManagementServiceSsoTokenResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetApiManagementServiceSsoTokenResult] {
	return pulumix.Output[GetApiManagementServiceSsoTokenResult]{
		OutputState: o.OutputState,
	}
}

// Redirect URL to the Publisher Portal containing the SSO token.
func (o GetApiManagementServiceSsoTokenResultOutput) RedirectUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetApiManagementServiceSsoTokenResult) *string { return v.RedirectUri }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetApiManagementServiceSsoTokenResultOutput{})
}
