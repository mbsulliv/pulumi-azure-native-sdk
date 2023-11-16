// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package insights

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// list date to migrate to new pricing model.
// Azure REST API version: 2017-10-01.
func ListEASubscriptionListMigrationDatePost(ctx *pulumi.Context, args *ListEASubscriptionListMigrationDatePostArgs, opts ...pulumi.InvokeOption) (*ListEASubscriptionListMigrationDatePostResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListEASubscriptionListMigrationDatePostResult
	err := ctx.Invoke("azure-native:insights:listEASubscriptionListMigrationDatePost", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListEASubscriptionListMigrationDatePostArgs struct {
}

// Subscription migrate date information properties
type ListEASubscriptionListMigrationDatePostResult struct {
	// Is subscription in the grand fatherable subscription list.
	IsGrandFatherableSubscription *bool `pulumi:"isGrandFatherableSubscription"`
	// Time to start using new pricing model.
	OptedInDate *string `pulumi:"optedInDate"`
}

func ListEASubscriptionListMigrationDatePostOutput(ctx *pulumi.Context, args ListEASubscriptionListMigrationDatePostOutputArgs, opts ...pulumi.InvokeOption) ListEASubscriptionListMigrationDatePostResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListEASubscriptionListMigrationDatePostResult, error) {
			args := v.(ListEASubscriptionListMigrationDatePostArgs)
			r, err := ListEASubscriptionListMigrationDatePost(ctx, &args, opts...)
			var s ListEASubscriptionListMigrationDatePostResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListEASubscriptionListMigrationDatePostResultOutput)
}

type ListEASubscriptionListMigrationDatePostOutputArgs struct {
}

func (ListEASubscriptionListMigrationDatePostOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListEASubscriptionListMigrationDatePostArgs)(nil)).Elem()
}

// Subscription migrate date information properties
type ListEASubscriptionListMigrationDatePostResultOutput struct{ *pulumi.OutputState }

func (ListEASubscriptionListMigrationDatePostResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListEASubscriptionListMigrationDatePostResult)(nil)).Elem()
}

func (o ListEASubscriptionListMigrationDatePostResultOutput) ToListEASubscriptionListMigrationDatePostResultOutput() ListEASubscriptionListMigrationDatePostResultOutput {
	return o
}

func (o ListEASubscriptionListMigrationDatePostResultOutput) ToListEASubscriptionListMigrationDatePostResultOutputWithContext(ctx context.Context) ListEASubscriptionListMigrationDatePostResultOutput {
	return o
}

// Is subscription in the grand fatherable subscription list.
func (o ListEASubscriptionListMigrationDatePostResultOutput) IsGrandFatherableSubscription() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ListEASubscriptionListMigrationDatePostResult) *bool { return v.IsGrandFatherableSubscription }).(pulumi.BoolPtrOutput)
}

// Time to start using new pricing model.
func (o ListEASubscriptionListMigrationDatePostResultOutput) OptedInDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListEASubscriptionListMigrationDatePostResult) *string { return v.OptedInDate }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListEASubscriptionListMigrationDatePostResultOutput{})
}
