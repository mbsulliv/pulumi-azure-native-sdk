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

// Gets the Shared Access Authorization Token for the User.
// Azure REST API version: 2022-08-01.
func GetUserSharedAccessToken(ctx *pulumi.Context, args *GetUserSharedAccessTokenArgs, opts ...pulumi.InvokeOption) (*GetUserSharedAccessTokenResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv GetUserSharedAccessTokenResult
	err := ctx.Invoke("azure-native:apimanagement:getUserSharedAccessToken", args.Defaults(), &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetUserSharedAccessTokenArgs struct {
	// The Expiry time of the Token. Maximum token expiry time is set to 30 days. The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
	Expiry string `pulumi:"expiry"`
	// The Key to be used to generate token for user.
	KeyType KeyType `pulumi:"keyType"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
	// User identifier. Must be unique in the current API Management service instance.
	UserId string `pulumi:"userId"`
}

// Defaults sets the appropriate defaults for GetUserSharedAccessTokenArgs
func (val *GetUserSharedAccessTokenArgs) Defaults() *GetUserSharedAccessTokenArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if utilities.IsZero(tmp.KeyType) {
		tmp.KeyType = KeyType("primary")
	}
	return &tmp
}

// Get User Token response details.
type GetUserSharedAccessTokenResult struct {
	// Shared Access Authorization token for the User.
	Value *string `pulumi:"value"`
}

func GetUserSharedAccessTokenOutput(ctx *pulumi.Context, args GetUserSharedAccessTokenOutputArgs, opts ...pulumi.InvokeOption) GetUserSharedAccessTokenResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetUserSharedAccessTokenResult, error) {
			args := v.(GetUserSharedAccessTokenArgs)
			r, err := GetUserSharedAccessToken(ctx, &args, opts...)
			var s GetUserSharedAccessTokenResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetUserSharedAccessTokenResultOutput)
}

type GetUserSharedAccessTokenOutputArgs struct {
	// The Expiry time of the Token. Maximum token expiry time is set to 30 days. The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
	Expiry pulumi.StringInput `pulumi:"expiry"`
	// The Key to be used to generate token for user.
	KeyType KeyTypeInput `pulumi:"keyType"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
	// User identifier. Must be unique in the current API Management service instance.
	UserId pulumi.StringInput `pulumi:"userId"`
}

func (GetUserSharedAccessTokenOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetUserSharedAccessTokenArgs)(nil)).Elem()
}

// Get User Token response details.
type GetUserSharedAccessTokenResultOutput struct{ *pulumi.OutputState }

func (GetUserSharedAccessTokenResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetUserSharedAccessTokenResult)(nil)).Elem()
}

func (o GetUserSharedAccessTokenResultOutput) ToGetUserSharedAccessTokenResultOutput() GetUserSharedAccessTokenResultOutput {
	return o
}

func (o GetUserSharedAccessTokenResultOutput) ToGetUserSharedAccessTokenResultOutputWithContext(ctx context.Context) GetUserSharedAccessTokenResultOutput {
	return o
}

func (o GetUserSharedAccessTokenResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetUserSharedAccessTokenResult] {
	return pulumix.Output[GetUserSharedAccessTokenResult]{
		OutputState: o.OutputState,
	}
}

// Shared Access Authorization token for the User.
func (o GetUserSharedAccessTokenResultOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetUserSharedAccessTokenResult) *string { return v.Value }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetUserSharedAccessTokenResultOutput{})
}
