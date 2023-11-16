// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package automation

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieve the deleted runbooks for an automation account.
// Azure REST API version: 2023-05-15-preview.
func ListAutomationAccountDeletedRunbooks(ctx *pulumi.Context, args *ListAutomationAccountDeletedRunbooksArgs, opts ...pulumi.InvokeOption) (*ListAutomationAccountDeletedRunbooksResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListAutomationAccountDeletedRunbooksResult
	err := ctx.Invoke("azure-native:automation:listAutomationAccountDeletedRunbooks", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListAutomationAccountDeletedRunbooksArgs struct {
	// The name of the automation account.
	AutomationAccountName string `pulumi:"automationAccountName"`
	// Name of an Azure Resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The response model for the list deleted runbook.
type ListAutomationAccountDeletedRunbooksResult struct {
	// Gets or sets the next link.
	NextLink *string `pulumi:"nextLink"`
	// List of deleted runbooks in automation account.
	Value []DeletedRunbookResponse `pulumi:"value"`
}

func ListAutomationAccountDeletedRunbooksOutput(ctx *pulumi.Context, args ListAutomationAccountDeletedRunbooksOutputArgs, opts ...pulumi.InvokeOption) ListAutomationAccountDeletedRunbooksResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListAutomationAccountDeletedRunbooksResult, error) {
			args := v.(ListAutomationAccountDeletedRunbooksArgs)
			r, err := ListAutomationAccountDeletedRunbooks(ctx, &args, opts...)
			var s ListAutomationAccountDeletedRunbooksResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListAutomationAccountDeletedRunbooksResultOutput)
}

type ListAutomationAccountDeletedRunbooksOutputArgs struct {
	// The name of the automation account.
	AutomationAccountName pulumi.StringInput `pulumi:"automationAccountName"`
	// Name of an Azure Resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListAutomationAccountDeletedRunbooksOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListAutomationAccountDeletedRunbooksArgs)(nil)).Elem()
}

// The response model for the list deleted runbook.
type ListAutomationAccountDeletedRunbooksResultOutput struct{ *pulumi.OutputState }

func (ListAutomationAccountDeletedRunbooksResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListAutomationAccountDeletedRunbooksResult)(nil)).Elem()
}

func (o ListAutomationAccountDeletedRunbooksResultOutput) ToListAutomationAccountDeletedRunbooksResultOutput() ListAutomationAccountDeletedRunbooksResultOutput {
	return o
}

func (o ListAutomationAccountDeletedRunbooksResultOutput) ToListAutomationAccountDeletedRunbooksResultOutputWithContext(ctx context.Context) ListAutomationAccountDeletedRunbooksResultOutput {
	return o
}

// Gets or sets the next link.
func (o ListAutomationAccountDeletedRunbooksResultOutput) NextLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListAutomationAccountDeletedRunbooksResult) *string { return v.NextLink }).(pulumi.StringPtrOutput)
}

// List of deleted runbooks in automation account.
func (o ListAutomationAccountDeletedRunbooksResultOutput) Value() DeletedRunbookResponseArrayOutput {
	return o.ApplyT(func(v ListAutomationAccountDeletedRunbooksResult) []DeletedRunbookResponse { return v.Value }).(DeletedRunbookResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListAutomationAccountDeletedRunbooksResultOutput{})
}
