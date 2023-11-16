// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20191001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The operation to get the export for the defined scope by export name.
func LookupExport(ctx *pulumi.Context, args *LookupExportArgs, opts ...pulumi.InvokeOption) (*LookupExportResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupExportResult
	err := ctx.Invoke("azure-native:costmanagement/v20191001:getExport", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupExportArgs struct {
	// Export Name.
	ExportName string `pulumi:"exportName"`
	// The scope associated with query and export operations. This includes '/subscriptions/{subscriptionId}/' for subscription scope, '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for resourceGroup scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope and '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/departments/{departmentId}' for Department scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/enrollmentAccounts/{enrollmentAccountId}' for EnrollmentAccount scope, '/providers/Microsoft.Management/managementGroups/{managementGroupId} for Management Group scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}' for billingProfile scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}/invoiceSections/{invoiceSectionId}' for invoiceSection scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/customers/{customerId}' specific for partners, 'providers/Microsoft.CostManagement/ExternalSubscriptions/{externalSubscriptionId}' for linked account and 'providers/Microsoft.CostManagement/externalBillingAccounts/{externalBillingAccountId}' for consolidated account
	Scope string `pulumi:"scope"`
}

// A export resource.
type LookupExportResult struct {
	// Has definition for the export.
	Definition QueryDefinitionResponse `pulumi:"definition"`
	// Has delivery information for the export.
	DeliveryInfo ExportDeliveryInfoResponse `pulumi:"deliveryInfo"`
	// The format of the export being delivered.
	Format *string `pulumi:"format"`
	// Resource Id.
	Id string `pulumi:"id"`
	// Resource name.
	Name string `pulumi:"name"`
	// Has schedule information for the export.
	Schedule *ExportScheduleResponse `pulumi:"schedule"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupExportOutput(ctx *pulumi.Context, args LookupExportOutputArgs, opts ...pulumi.InvokeOption) LookupExportResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupExportResult, error) {
			args := v.(LookupExportArgs)
			r, err := LookupExport(ctx, &args, opts...)
			var s LookupExportResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupExportResultOutput)
}

type LookupExportOutputArgs struct {
	// Export Name.
	ExportName pulumi.StringInput `pulumi:"exportName"`
	// The scope associated with query and export operations. This includes '/subscriptions/{subscriptionId}/' for subscription scope, '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for resourceGroup scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope and '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/departments/{departmentId}' for Department scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/enrollmentAccounts/{enrollmentAccountId}' for EnrollmentAccount scope, '/providers/Microsoft.Management/managementGroups/{managementGroupId} for Management Group scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}' for billingProfile scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}/invoiceSections/{invoiceSectionId}' for invoiceSection scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/customers/{customerId}' specific for partners, 'providers/Microsoft.CostManagement/ExternalSubscriptions/{externalSubscriptionId}' for linked account and 'providers/Microsoft.CostManagement/externalBillingAccounts/{externalBillingAccountId}' for consolidated account
	Scope pulumi.StringInput `pulumi:"scope"`
}

func (LookupExportOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupExportArgs)(nil)).Elem()
}

// A export resource.
type LookupExportResultOutput struct{ *pulumi.OutputState }

func (LookupExportResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupExportResult)(nil)).Elem()
}

func (o LookupExportResultOutput) ToLookupExportResultOutput() LookupExportResultOutput {
	return o
}

func (o LookupExportResultOutput) ToLookupExportResultOutputWithContext(ctx context.Context) LookupExportResultOutput {
	return o
}

// Has definition for the export.
func (o LookupExportResultOutput) Definition() QueryDefinitionResponseOutput {
	return o.ApplyT(func(v LookupExportResult) QueryDefinitionResponse { return v.Definition }).(QueryDefinitionResponseOutput)
}

// Has delivery information for the export.
func (o LookupExportResultOutput) DeliveryInfo() ExportDeliveryInfoResponseOutput {
	return o.ApplyT(func(v LookupExportResult) ExportDeliveryInfoResponse { return v.DeliveryInfo }).(ExportDeliveryInfoResponseOutput)
}

// The format of the export being delivered.
func (o LookupExportResultOutput) Format() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExportResult) *string { return v.Format }).(pulumi.StringPtrOutput)
}

// Resource Id.
func (o LookupExportResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExportResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource name.
func (o LookupExportResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExportResult) string { return v.Name }).(pulumi.StringOutput)
}

// Has schedule information for the export.
func (o LookupExportResultOutput) Schedule() ExportScheduleResponsePtrOutput {
	return o.ApplyT(func(v LookupExportResult) *ExportScheduleResponse { return v.Schedule }).(ExportScheduleResponsePtrOutput)
}

// Resource tags.
func (o LookupExportResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupExportResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o LookupExportResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExportResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupExportResultOutput{})
}
