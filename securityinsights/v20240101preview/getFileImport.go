// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20240101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a file import.
func LookupFileImport(ctx *pulumi.Context, args *LookupFileImportArgs, opts ...pulumi.InvokeOption) (*LookupFileImportResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupFileImportResult
	err := ctx.Invoke("azure-native:securityinsights/v20240101preview:getFileImport", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFileImportArgs struct {
	// File import ID
	FileImportId string `pulumi:"fileImportId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// Represents a file import in Azure Security Insights.
type LookupFileImportResult struct {
	// The content type of this file.
	ContentType string `pulumi:"contentType"`
	// The time the file was imported.
	CreatedTimeUTC string `pulumi:"createdTimeUTC"`
	// Represents the error file (if the import was ingested with errors or failed the validation).
	ErrorFile FileMetadataResponse `pulumi:"errorFile"`
	// An ordered list of some of the errors that were encountered during validation.
	ErrorsPreview []ValidationErrorResponse `pulumi:"errorsPreview"`
	// The time the files associated with this import are deleted from the storage account.
	FilesValidUntilTimeUTC string `pulumi:"filesValidUntilTimeUTC"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// Represents the imported file.
	ImportFile FileMetadataResponse `pulumi:"importFile"`
	// The time the file import record is soft deleted from the database and history.
	ImportValidUntilTimeUTC string `pulumi:"importValidUntilTimeUTC"`
	// The number of records that have been successfully ingested.
	IngestedRecordCount int `pulumi:"ingestedRecordCount"`
	// Describes how to ingest the records in the file.
	IngestionMode string `pulumi:"ingestionMode"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The source for the data in the file.
	Source string `pulumi:"source"`
	// The state of the file import.
	State string `pulumi:"state"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The number of records in the file.
	TotalRecordCount int `pulumi:"totalRecordCount"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// The number of records that have passed validation.
	ValidRecordCount int `pulumi:"validRecordCount"`
}

func LookupFileImportOutput(ctx *pulumi.Context, args LookupFileImportOutputArgs, opts ...pulumi.InvokeOption) LookupFileImportResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFileImportResult, error) {
			args := v.(LookupFileImportArgs)
			r, err := LookupFileImport(ctx, &args, opts...)
			var s LookupFileImportResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFileImportResultOutput)
}

type LookupFileImportOutputArgs struct {
	// File import ID
	FileImportId pulumi.StringInput `pulumi:"fileImportId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupFileImportOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFileImportArgs)(nil)).Elem()
}

// Represents a file import in Azure Security Insights.
type LookupFileImportResultOutput struct{ *pulumi.OutputState }

func (LookupFileImportResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFileImportResult)(nil)).Elem()
}

func (o LookupFileImportResultOutput) ToLookupFileImportResultOutput() LookupFileImportResultOutput {
	return o
}

func (o LookupFileImportResultOutput) ToLookupFileImportResultOutputWithContext(ctx context.Context) LookupFileImportResultOutput {
	return o
}

// The content type of this file.
func (o LookupFileImportResultOutput) ContentType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFileImportResult) string { return v.ContentType }).(pulumi.StringOutput)
}

// The time the file was imported.
func (o LookupFileImportResultOutput) CreatedTimeUTC() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFileImportResult) string { return v.CreatedTimeUTC }).(pulumi.StringOutput)
}

// Represents the error file (if the import was ingested with errors or failed the validation).
func (o LookupFileImportResultOutput) ErrorFile() FileMetadataResponseOutput {
	return o.ApplyT(func(v LookupFileImportResult) FileMetadataResponse { return v.ErrorFile }).(FileMetadataResponseOutput)
}

// An ordered list of some of the errors that were encountered during validation.
func (o LookupFileImportResultOutput) ErrorsPreview() ValidationErrorResponseArrayOutput {
	return o.ApplyT(func(v LookupFileImportResult) []ValidationErrorResponse { return v.ErrorsPreview }).(ValidationErrorResponseArrayOutput)
}

// The time the files associated with this import are deleted from the storage account.
func (o LookupFileImportResultOutput) FilesValidUntilTimeUTC() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFileImportResult) string { return v.FilesValidUntilTimeUTC }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupFileImportResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFileImportResult) string { return v.Id }).(pulumi.StringOutput)
}

// Represents the imported file.
func (o LookupFileImportResultOutput) ImportFile() FileMetadataResponseOutput {
	return o.ApplyT(func(v LookupFileImportResult) FileMetadataResponse { return v.ImportFile }).(FileMetadataResponseOutput)
}

// The time the file import record is soft deleted from the database and history.
func (o LookupFileImportResultOutput) ImportValidUntilTimeUTC() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFileImportResult) string { return v.ImportValidUntilTimeUTC }).(pulumi.StringOutput)
}

// The number of records that have been successfully ingested.
func (o LookupFileImportResultOutput) IngestedRecordCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupFileImportResult) int { return v.IngestedRecordCount }).(pulumi.IntOutput)
}

// Describes how to ingest the records in the file.
func (o LookupFileImportResultOutput) IngestionMode() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFileImportResult) string { return v.IngestionMode }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupFileImportResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFileImportResult) string { return v.Name }).(pulumi.StringOutput)
}

// The source for the data in the file.
func (o LookupFileImportResultOutput) Source() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFileImportResult) string { return v.Source }).(pulumi.StringOutput)
}

// The state of the file import.
func (o LookupFileImportResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFileImportResult) string { return v.State }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupFileImportResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupFileImportResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The number of records in the file.
func (o LookupFileImportResultOutput) TotalRecordCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupFileImportResult) int { return v.TotalRecordCount }).(pulumi.IntOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupFileImportResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFileImportResult) string { return v.Type }).(pulumi.StringOutput)
}

// The number of records that have passed validation.
func (o LookupFileImportResultOutput) ValidRecordCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupFileImportResult) int { return v.ValidRecordCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFileImportResultOutput{})
}
