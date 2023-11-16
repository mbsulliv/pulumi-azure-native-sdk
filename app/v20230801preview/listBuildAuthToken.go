// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230801preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the token used to connect to the endpoint where source code can be uploaded for a build.
func ListBuildAuthToken(ctx *pulumi.Context, args *ListBuildAuthTokenArgs, opts ...pulumi.InvokeOption) (*ListBuildAuthTokenResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListBuildAuthTokenResult
	err := ctx.Invoke("azure-native:app/v20230801preview:listBuildAuthToken", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListBuildAuthTokenArgs struct {
	// The name of a build.
	BuildName string `pulumi:"buildName"`
	// The name of the builder.
	BuilderName string `pulumi:"builderName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Build Auth Token.
type ListBuildAuthTokenResult struct {
	// Token expiration date.
	Expires string `pulumi:"expires"`
	// Authentication token.
	Token string `pulumi:"token"`
}

func ListBuildAuthTokenOutput(ctx *pulumi.Context, args ListBuildAuthTokenOutputArgs, opts ...pulumi.InvokeOption) ListBuildAuthTokenResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListBuildAuthTokenResult, error) {
			args := v.(ListBuildAuthTokenArgs)
			r, err := ListBuildAuthToken(ctx, &args, opts...)
			var s ListBuildAuthTokenResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListBuildAuthTokenResultOutput)
}

type ListBuildAuthTokenOutputArgs struct {
	// The name of a build.
	BuildName pulumi.StringInput `pulumi:"buildName"`
	// The name of the builder.
	BuilderName pulumi.StringInput `pulumi:"builderName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListBuildAuthTokenOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListBuildAuthTokenArgs)(nil)).Elem()
}

// Build Auth Token.
type ListBuildAuthTokenResultOutput struct{ *pulumi.OutputState }

func (ListBuildAuthTokenResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListBuildAuthTokenResult)(nil)).Elem()
}

func (o ListBuildAuthTokenResultOutput) ToListBuildAuthTokenResultOutput() ListBuildAuthTokenResultOutput {
	return o
}

func (o ListBuildAuthTokenResultOutput) ToListBuildAuthTokenResultOutputWithContext(ctx context.Context) ListBuildAuthTokenResultOutput {
	return o
}

// Token expiration date.
func (o ListBuildAuthTokenResultOutput) Expires() pulumi.StringOutput {
	return o.ApplyT(func(v ListBuildAuthTokenResult) string { return v.Expires }).(pulumi.StringOutput)
}

// Authentication token.
func (o ListBuildAuthTokenResultOutput) Token() pulumi.StringOutput {
	return o.ApplyT(func(v ListBuildAuthTokenResult) string { return v.Token }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListBuildAuthTokenResultOutput{})
}
