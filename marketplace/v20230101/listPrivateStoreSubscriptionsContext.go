// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// List all the subscriptions in the private store context
func ListPrivateStoreSubscriptionsContext(ctx *pulumi.Context, args *ListPrivateStoreSubscriptionsContextArgs, opts ...pulumi.InvokeOption) (*ListPrivateStoreSubscriptionsContextResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListPrivateStoreSubscriptionsContextResult
	err := ctx.Invoke("azure-native:marketplace/v20230101:listPrivateStoreSubscriptionsContext", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListPrivateStoreSubscriptionsContextArgs struct {
	// The store ID - must use the tenant ID
	PrivateStoreId string `pulumi:"privateStoreId"`
}

// List of subscription Ids in the private store
type ListPrivateStoreSubscriptionsContextResult struct {
	SubscriptionsIds []string `pulumi:"subscriptionsIds"`
}

func ListPrivateStoreSubscriptionsContextOutput(ctx *pulumi.Context, args ListPrivateStoreSubscriptionsContextOutputArgs, opts ...pulumi.InvokeOption) ListPrivateStoreSubscriptionsContextResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListPrivateStoreSubscriptionsContextResult, error) {
			args := v.(ListPrivateStoreSubscriptionsContextArgs)
			r, err := ListPrivateStoreSubscriptionsContext(ctx, &args, opts...)
			var s ListPrivateStoreSubscriptionsContextResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListPrivateStoreSubscriptionsContextResultOutput)
}

type ListPrivateStoreSubscriptionsContextOutputArgs struct {
	// The store ID - must use the tenant ID
	PrivateStoreId pulumi.StringInput `pulumi:"privateStoreId"`
}

func (ListPrivateStoreSubscriptionsContextOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListPrivateStoreSubscriptionsContextArgs)(nil)).Elem()
}

// List of subscription Ids in the private store
type ListPrivateStoreSubscriptionsContextResultOutput struct{ *pulumi.OutputState }

func (ListPrivateStoreSubscriptionsContextResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListPrivateStoreSubscriptionsContextResult)(nil)).Elem()
}

func (o ListPrivateStoreSubscriptionsContextResultOutput) ToListPrivateStoreSubscriptionsContextResultOutput() ListPrivateStoreSubscriptionsContextResultOutput {
	return o
}

func (o ListPrivateStoreSubscriptionsContextResultOutput) ToListPrivateStoreSubscriptionsContextResultOutputWithContext(ctx context.Context) ListPrivateStoreSubscriptionsContextResultOutput {
	return o
}

func (o ListPrivateStoreSubscriptionsContextResultOutput) ToOutput(ctx context.Context) pulumix.Output[ListPrivateStoreSubscriptionsContextResult] {
	return pulumix.Output[ListPrivateStoreSubscriptionsContextResult]{
		OutputState: o.OutputState,
	}
}

func (o ListPrivateStoreSubscriptionsContextResultOutput) SubscriptionsIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ListPrivateStoreSubscriptionsContextResult) []string { return v.SubscriptionsIds }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListPrivateStoreSubscriptionsContextResultOutput{})
}
