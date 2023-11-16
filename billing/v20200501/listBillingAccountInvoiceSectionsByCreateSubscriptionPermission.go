// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20200501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Lists the invoice sections for which the user has permission to create Azure subscriptions. The operation is supported only for billing accounts with agreement type Microsoft Customer Agreement.
func ListBillingAccountInvoiceSectionsByCreateSubscriptionPermission(ctx *pulumi.Context, args *ListBillingAccountInvoiceSectionsByCreateSubscriptionPermissionArgs, opts ...pulumi.InvokeOption) (*ListBillingAccountInvoiceSectionsByCreateSubscriptionPermissionResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListBillingAccountInvoiceSectionsByCreateSubscriptionPermissionResult
	err := ctx.Invoke("azure-native:billing/v20200501:listBillingAccountInvoiceSectionsByCreateSubscriptionPermission", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListBillingAccountInvoiceSectionsByCreateSubscriptionPermissionArgs struct {
	// The ID that uniquely identifies a billing account.
	BillingAccountName string `pulumi:"billingAccountName"`
}

// The list of invoice section properties with create subscription permission.
type ListBillingAccountInvoiceSectionsByCreateSubscriptionPermissionResult struct {
	// The link (url) to the next page of results.
	NextLink string `pulumi:"nextLink"`
	// The list of invoice section properties with create subscription permission.
	Value []InvoiceSectionWithCreateSubPermissionResponse `pulumi:"value"`
}

func ListBillingAccountInvoiceSectionsByCreateSubscriptionPermissionOutput(ctx *pulumi.Context, args ListBillingAccountInvoiceSectionsByCreateSubscriptionPermissionOutputArgs, opts ...pulumi.InvokeOption) ListBillingAccountInvoiceSectionsByCreateSubscriptionPermissionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListBillingAccountInvoiceSectionsByCreateSubscriptionPermissionResult, error) {
			args := v.(ListBillingAccountInvoiceSectionsByCreateSubscriptionPermissionArgs)
			r, err := ListBillingAccountInvoiceSectionsByCreateSubscriptionPermission(ctx, &args, opts...)
			var s ListBillingAccountInvoiceSectionsByCreateSubscriptionPermissionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListBillingAccountInvoiceSectionsByCreateSubscriptionPermissionResultOutput)
}

type ListBillingAccountInvoiceSectionsByCreateSubscriptionPermissionOutputArgs struct {
	// The ID that uniquely identifies a billing account.
	BillingAccountName pulumi.StringInput `pulumi:"billingAccountName"`
}

func (ListBillingAccountInvoiceSectionsByCreateSubscriptionPermissionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListBillingAccountInvoiceSectionsByCreateSubscriptionPermissionArgs)(nil)).Elem()
}

// The list of invoice section properties with create subscription permission.
type ListBillingAccountInvoiceSectionsByCreateSubscriptionPermissionResultOutput struct{ *pulumi.OutputState }

func (ListBillingAccountInvoiceSectionsByCreateSubscriptionPermissionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListBillingAccountInvoiceSectionsByCreateSubscriptionPermissionResult)(nil)).Elem()
}

func (o ListBillingAccountInvoiceSectionsByCreateSubscriptionPermissionResultOutput) ToListBillingAccountInvoiceSectionsByCreateSubscriptionPermissionResultOutput() ListBillingAccountInvoiceSectionsByCreateSubscriptionPermissionResultOutput {
	return o
}

func (o ListBillingAccountInvoiceSectionsByCreateSubscriptionPermissionResultOutput) ToListBillingAccountInvoiceSectionsByCreateSubscriptionPermissionResultOutputWithContext(ctx context.Context) ListBillingAccountInvoiceSectionsByCreateSubscriptionPermissionResultOutput {
	return o
}

// The link (url) to the next page of results.
func (o ListBillingAccountInvoiceSectionsByCreateSubscriptionPermissionResultOutput) NextLink() pulumi.StringOutput {
	return o.ApplyT(func(v ListBillingAccountInvoiceSectionsByCreateSubscriptionPermissionResult) string {
		return v.NextLink
	}).(pulumi.StringOutput)
}

// The list of invoice section properties with create subscription permission.
func (o ListBillingAccountInvoiceSectionsByCreateSubscriptionPermissionResultOutput) Value() InvoiceSectionWithCreateSubPermissionResponseArrayOutput {
	return o.ApplyT(func(v ListBillingAccountInvoiceSectionsByCreateSubscriptionPermissionResult) []InvoiceSectionWithCreateSubPermissionResponse {
		return v.Value
	}).(InvoiceSectionWithCreateSubPermissionResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListBillingAccountInvoiceSectionsByCreateSubscriptionPermissionResultOutput{})
}
