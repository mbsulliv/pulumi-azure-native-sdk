// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20191001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

var _ = utilities.GetEnvOrDefault

// Details of the Azure plan.
type AzurePlanResponse struct {
	// The sku description.
	SkuDescription string `pulumi:"skuDescription"`
	// The sku id.
	SkuId *string `pulumi:"skuId"`
}

// Details of the Azure plan.
type AzurePlanResponseOutput struct{ *pulumi.OutputState }

func (AzurePlanResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzurePlanResponse)(nil)).Elem()
}

func (o AzurePlanResponseOutput) ToAzurePlanResponseOutput() AzurePlanResponseOutput {
	return o
}

func (o AzurePlanResponseOutput) ToAzurePlanResponseOutputWithContext(ctx context.Context) AzurePlanResponseOutput {
	return o
}

func (o AzurePlanResponseOutput) ToOutput(ctx context.Context) pulumix.Output[AzurePlanResponse] {
	return pulumix.Output[AzurePlanResponse]{
		OutputState: o.OutputState,
	}
}

// The sku description.
func (o AzurePlanResponseOutput) SkuDescription() pulumi.StringOutput {
	return o.ApplyT(func(v AzurePlanResponse) string { return v.SkuDescription }).(pulumi.StringOutput)
}

// The sku id.
func (o AzurePlanResponseOutput) SkuId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzurePlanResponse) *string { return v.SkuId }).(pulumi.StringPtrOutput)
}

type AzurePlanResponseArrayOutput struct{ *pulumi.OutputState }

func (AzurePlanResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AzurePlanResponse)(nil)).Elem()
}

func (o AzurePlanResponseArrayOutput) ToAzurePlanResponseArrayOutput() AzurePlanResponseArrayOutput {
	return o
}

func (o AzurePlanResponseArrayOutput) ToAzurePlanResponseArrayOutputWithContext(ctx context.Context) AzurePlanResponseArrayOutput {
	return o
}

func (o AzurePlanResponseArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]AzurePlanResponse] {
	return pulumix.Output[[]AzurePlanResponse]{
		OutputState: o.OutputState,
	}
}

func (o AzurePlanResponseArrayOutput) Index(i pulumi.IntInput) AzurePlanResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AzurePlanResponse {
		return vs[0].([]AzurePlanResponse)[vs[1].(int)]
	}).(AzurePlanResponseOutput)
}

// Invoice section properties with create subscription permission.
type InvoiceSectionWithCreateSubPermissionResponse struct {
	// The name of the billing profile for the invoice section.
	BillingProfileDisplayName string `pulumi:"billingProfileDisplayName"`
	// The ID of the billing profile for the invoice section.
	BillingProfileId string `pulumi:"billingProfileId"`
	// The billing profile spending limit.
	BillingProfileSpendingLimit string `pulumi:"billingProfileSpendingLimit"`
	// The status of the billing profile.
	BillingProfileStatus string `pulumi:"billingProfileStatus"`
	// Reason for the specified billing profile status.
	BillingProfileStatusReasonCode string `pulumi:"billingProfileStatusReasonCode"`
	// Enabled azure plans for the associated billing profile.
	EnabledAzurePlans []AzurePlanResponse `pulumi:"enabledAzurePlans"`
	// The name of the invoice section.
	InvoiceSectionDisplayName string `pulumi:"invoiceSectionDisplayName"`
	// The ID of the invoice section.
	InvoiceSectionId string `pulumi:"invoiceSectionId"`
}

// Invoice section properties with create subscription permission.
type InvoiceSectionWithCreateSubPermissionResponseOutput struct{ *pulumi.OutputState }

func (InvoiceSectionWithCreateSubPermissionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InvoiceSectionWithCreateSubPermissionResponse)(nil)).Elem()
}

func (o InvoiceSectionWithCreateSubPermissionResponseOutput) ToInvoiceSectionWithCreateSubPermissionResponseOutput() InvoiceSectionWithCreateSubPermissionResponseOutput {
	return o
}

func (o InvoiceSectionWithCreateSubPermissionResponseOutput) ToInvoiceSectionWithCreateSubPermissionResponseOutputWithContext(ctx context.Context) InvoiceSectionWithCreateSubPermissionResponseOutput {
	return o
}

func (o InvoiceSectionWithCreateSubPermissionResponseOutput) ToOutput(ctx context.Context) pulumix.Output[InvoiceSectionWithCreateSubPermissionResponse] {
	return pulumix.Output[InvoiceSectionWithCreateSubPermissionResponse]{
		OutputState: o.OutputState,
	}
}

// The name of the billing profile for the invoice section.
func (o InvoiceSectionWithCreateSubPermissionResponseOutput) BillingProfileDisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v InvoiceSectionWithCreateSubPermissionResponse) string { return v.BillingProfileDisplayName }).(pulumi.StringOutput)
}

// The ID of the billing profile for the invoice section.
func (o InvoiceSectionWithCreateSubPermissionResponseOutput) BillingProfileId() pulumi.StringOutput {
	return o.ApplyT(func(v InvoiceSectionWithCreateSubPermissionResponse) string { return v.BillingProfileId }).(pulumi.StringOutput)
}

// The billing profile spending limit.
func (o InvoiceSectionWithCreateSubPermissionResponseOutput) BillingProfileSpendingLimit() pulumi.StringOutput {
	return o.ApplyT(func(v InvoiceSectionWithCreateSubPermissionResponse) string { return v.BillingProfileSpendingLimit }).(pulumi.StringOutput)
}

// The status of the billing profile.
func (o InvoiceSectionWithCreateSubPermissionResponseOutput) BillingProfileStatus() pulumi.StringOutput {
	return o.ApplyT(func(v InvoiceSectionWithCreateSubPermissionResponse) string { return v.BillingProfileStatus }).(pulumi.StringOutput)
}

// Reason for the specified billing profile status.
func (o InvoiceSectionWithCreateSubPermissionResponseOutput) BillingProfileStatusReasonCode() pulumi.StringOutput {
	return o.ApplyT(func(v InvoiceSectionWithCreateSubPermissionResponse) string { return v.BillingProfileStatusReasonCode }).(pulumi.StringOutput)
}

// Enabled azure plans for the associated billing profile.
func (o InvoiceSectionWithCreateSubPermissionResponseOutput) EnabledAzurePlans() AzurePlanResponseArrayOutput {
	return o.ApplyT(func(v InvoiceSectionWithCreateSubPermissionResponse) []AzurePlanResponse { return v.EnabledAzurePlans }).(AzurePlanResponseArrayOutput)
}

// The name of the invoice section.
func (o InvoiceSectionWithCreateSubPermissionResponseOutput) InvoiceSectionDisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v InvoiceSectionWithCreateSubPermissionResponse) string { return v.InvoiceSectionDisplayName }).(pulumi.StringOutput)
}

// The ID of the invoice section.
func (o InvoiceSectionWithCreateSubPermissionResponseOutput) InvoiceSectionId() pulumi.StringOutput {
	return o.ApplyT(func(v InvoiceSectionWithCreateSubPermissionResponse) string { return v.InvoiceSectionId }).(pulumi.StringOutput)
}

type InvoiceSectionWithCreateSubPermissionResponseArrayOutput struct{ *pulumi.OutputState }

func (InvoiceSectionWithCreateSubPermissionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InvoiceSectionWithCreateSubPermissionResponse)(nil)).Elem()
}

func (o InvoiceSectionWithCreateSubPermissionResponseArrayOutput) ToInvoiceSectionWithCreateSubPermissionResponseArrayOutput() InvoiceSectionWithCreateSubPermissionResponseArrayOutput {
	return o
}

func (o InvoiceSectionWithCreateSubPermissionResponseArrayOutput) ToInvoiceSectionWithCreateSubPermissionResponseArrayOutputWithContext(ctx context.Context) InvoiceSectionWithCreateSubPermissionResponseArrayOutput {
	return o
}

func (o InvoiceSectionWithCreateSubPermissionResponseArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]InvoiceSectionWithCreateSubPermissionResponse] {
	return pulumix.Output[[]InvoiceSectionWithCreateSubPermissionResponse]{
		OutputState: o.OutputState,
	}
}

func (o InvoiceSectionWithCreateSubPermissionResponseArrayOutput) Index(i pulumi.IntInput) InvoiceSectionWithCreateSubPermissionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) InvoiceSectionWithCreateSubPermissionResponse {
		return vs[0].([]InvoiceSectionWithCreateSubPermissionResponse)[vs[1].(int)]
	}).(InvoiceSectionWithCreateSubPermissionResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(AzurePlanResponseOutput{})
	pulumi.RegisterOutputType(AzurePlanResponseArrayOutput{})
	pulumi.RegisterOutputType(InvoiceSectionWithCreateSubPermissionResponseOutput{})
	pulumi.RegisterOutputType(InvoiceSectionWithCreateSubPermissionResponseArrayOutput{})
}
