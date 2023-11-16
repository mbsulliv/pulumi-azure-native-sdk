// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20181015

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get batch operation status
func GetGlobalUserOperationBatchStatus(ctx *pulumi.Context, args *GetGlobalUserOperationBatchStatusArgs, opts ...pulumi.InvokeOption) (*GetGlobalUserOperationBatchStatusResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv GetGlobalUserOperationBatchStatusResult
	err := ctx.Invoke("azure-native:labservices/v20181015:getGlobalUserOperationBatchStatus", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetGlobalUserOperationBatchStatusArgs struct {
	// The operation url of long running operation
	Urls []string `pulumi:"urls"`
	// The name of the user.
	UserName string `pulumi:"userName"`
}

// Status Details of the long running operation for an environment
type GetGlobalUserOperationBatchStatusResult struct {
	// Gets a collection of items that contain the operation url and status.
	Items []OperationBatchStatusResponseItemResponse `pulumi:"items"`
}

func GetGlobalUserOperationBatchStatusOutput(ctx *pulumi.Context, args GetGlobalUserOperationBatchStatusOutputArgs, opts ...pulumi.InvokeOption) GetGlobalUserOperationBatchStatusResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetGlobalUserOperationBatchStatusResult, error) {
			args := v.(GetGlobalUserOperationBatchStatusArgs)
			r, err := GetGlobalUserOperationBatchStatus(ctx, &args, opts...)
			var s GetGlobalUserOperationBatchStatusResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetGlobalUserOperationBatchStatusResultOutput)
}

type GetGlobalUserOperationBatchStatusOutputArgs struct {
	// The operation url of long running operation
	Urls pulumi.StringArrayInput `pulumi:"urls"`
	// The name of the user.
	UserName pulumi.StringInput `pulumi:"userName"`
}

func (GetGlobalUserOperationBatchStatusOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetGlobalUserOperationBatchStatusArgs)(nil)).Elem()
}

// Status Details of the long running operation for an environment
type GetGlobalUserOperationBatchStatusResultOutput struct{ *pulumi.OutputState }

func (GetGlobalUserOperationBatchStatusResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetGlobalUserOperationBatchStatusResult)(nil)).Elem()
}

func (o GetGlobalUserOperationBatchStatusResultOutput) ToGetGlobalUserOperationBatchStatusResultOutput() GetGlobalUserOperationBatchStatusResultOutput {
	return o
}

func (o GetGlobalUserOperationBatchStatusResultOutput) ToGetGlobalUserOperationBatchStatusResultOutputWithContext(ctx context.Context) GetGlobalUserOperationBatchStatusResultOutput {
	return o
}

// Gets a collection of items that contain the operation url and status.
func (o GetGlobalUserOperationBatchStatusResultOutput) Items() OperationBatchStatusResponseItemResponseArrayOutput {
	return o.ApplyT(func(v GetGlobalUserOperationBatchStatusResult) []OperationBatchStatusResponseItemResponse {
		return v.Items
	}).(OperationBatchStatusResponseItemResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetGlobalUserOperationBatchStatusResultOutput{})
}
