// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get private DNS zone suffix in the cloud.
func GetGetPrivateDnsZoneSuffixExecute(ctx *pulumi.Context, args *GetGetPrivateDnsZoneSuffixExecuteArgs, opts ...pulumi.InvokeOption) (*GetGetPrivateDnsZoneSuffixExecuteResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv GetGetPrivateDnsZoneSuffixExecuteResult
	err := ctx.Invoke("azure-native:dbformysql/v20220101:getGetPrivateDnsZoneSuffixExecute", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetGetPrivateDnsZoneSuffixExecuteArgs struct {
}

// The response of get private dns zone suffix.
type GetGetPrivateDnsZoneSuffixExecuteResult struct {
	// Represents the private DNS zone suffix.
	PrivateDnsZoneSuffix *string `pulumi:"privateDnsZoneSuffix"`
}

func GetGetPrivateDnsZoneSuffixExecuteOutput(ctx *pulumi.Context, args GetGetPrivateDnsZoneSuffixExecuteOutputArgs, opts ...pulumi.InvokeOption) GetGetPrivateDnsZoneSuffixExecuteResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetGetPrivateDnsZoneSuffixExecuteResult, error) {
			args := v.(GetGetPrivateDnsZoneSuffixExecuteArgs)
			r, err := GetGetPrivateDnsZoneSuffixExecute(ctx, &args, opts...)
			var s GetGetPrivateDnsZoneSuffixExecuteResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetGetPrivateDnsZoneSuffixExecuteResultOutput)
}

type GetGetPrivateDnsZoneSuffixExecuteOutputArgs struct {
}

func (GetGetPrivateDnsZoneSuffixExecuteOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetGetPrivateDnsZoneSuffixExecuteArgs)(nil)).Elem()
}

// The response of get private dns zone suffix.
type GetGetPrivateDnsZoneSuffixExecuteResultOutput struct{ *pulumi.OutputState }

func (GetGetPrivateDnsZoneSuffixExecuteResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetGetPrivateDnsZoneSuffixExecuteResult)(nil)).Elem()
}

func (o GetGetPrivateDnsZoneSuffixExecuteResultOutput) ToGetGetPrivateDnsZoneSuffixExecuteResultOutput() GetGetPrivateDnsZoneSuffixExecuteResultOutput {
	return o
}

func (o GetGetPrivateDnsZoneSuffixExecuteResultOutput) ToGetGetPrivateDnsZoneSuffixExecuteResultOutputWithContext(ctx context.Context) GetGetPrivateDnsZoneSuffixExecuteResultOutput {
	return o
}

// Represents the private DNS zone suffix.
func (o GetGetPrivateDnsZoneSuffixExecuteResultOutput) PrivateDnsZoneSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetGetPrivateDnsZoneSuffixExecuteResult) *string { return v.PrivateDnsZoneSuffix }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetGetPrivateDnsZoneSuffixExecuteResultOutput{})
}
