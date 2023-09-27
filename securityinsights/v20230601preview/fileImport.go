// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230601preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Represents a file import in Azure Security Insights.
type FileImport struct {
	pulumi.CustomResourceState

	// The content type of this file.
	ContentType pulumi.StringOutput `pulumi:"contentType"`
	// The time the file was imported.
	CreatedTimeUTC pulumi.StringOutput `pulumi:"createdTimeUTC"`
	// Represents the error file (if the import was ingested with errors or failed the validation).
	ErrorFile FileMetadataResponseOutput `pulumi:"errorFile"`
	// An ordered list of some of the errors that were encountered during validation.
	ErrorsPreview ValidationErrorResponseArrayOutput `pulumi:"errorsPreview"`
	// The time the files associated with this import are deleted from the storage account.
	FilesValidUntilTimeUTC pulumi.StringOutput `pulumi:"filesValidUntilTimeUTC"`
	// Represents the imported file.
	ImportFile FileMetadataResponseOutput `pulumi:"importFile"`
	// The time the file import record is soft deleted from the database and history.
	ImportValidUntilTimeUTC pulumi.StringOutput `pulumi:"importValidUntilTimeUTC"`
	// The number of records that have been successfully ingested.
	IngestedRecordCount pulumi.IntOutput `pulumi:"ingestedRecordCount"`
	// Describes how to ingest the records in the file.
	IngestionMode pulumi.StringOutput `pulumi:"ingestionMode"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The source for the data in the file.
	Source pulumi.StringOutput `pulumi:"source"`
	// The state of the file import.
	State pulumi.StringOutput `pulumi:"state"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The number of records in the file.
	TotalRecordCount pulumi.IntOutput `pulumi:"totalRecordCount"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// The number of records that have passed validation.
	ValidRecordCount pulumi.IntOutput `pulumi:"validRecordCount"`
}

// NewFileImport registers a new resource with the given unique name, arguments, and options.
func NewFileImport(ctx *pulumi.Context,
	name string, args *FileImportArgs, opts ...pulumi.ResourceOption) (*FileImport, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ContentType == nil {
		return nil, errors.New("invalid value for required argument 'ContentType'")
	}
	if args.ImportFile == nil {
		return nil, errors.New("invalid value for required argument 'ImportFile'")
	}
	if args.IngestionMode == nil {
		return nil, errors.New("invalid value for required argument 'IngestionMode'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Source == nil {
		return nil, errors.New("invalid value for required argument 'Source'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights:FileImport"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801preview:FileImport"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220901preview:FileImport"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221001preview:FileImport"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101preview:FileImport"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221201preview:FileImport"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230201preview:FileImport"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230301preview:FileImport"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230401preview:FileImport"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230501preview:FileImport"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230701preview:FileImport"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230801preview:FileImport"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230901preview:FileImport"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource FileImport
	err := ctx.RegisterResource("azure-native:securityinsights/v20230601preview:FileImport", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFileImport gets an existing FileImport resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFileImport(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FileImportState, opts ...pulumi.ResourceOption) (*FileImport, error) {
	var resource FileImport
	err := ctx.ReadResource("azure-native:securityinsights/v20230601preview:FileImport", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FileImport resources.
type fileImportState struct {
}

type FileImportState struct {
}

func (FileImportState) ElementType() reflect.Type {
	return reflect.TypeOf((*fileImportState)(nil)).Elem()
}

type fileImportArgs struct {
	// The content type of this file.
	ContentType string `pulumi:"contentType"`
	// File import ID
	FileImportId *string `pulumi:"fileImportId"`
	// Represents the imported file.
	ImportFile FileMetadata `pulumi:"importFile"`
	// Describes how to ingest the records in the file.
	IngestionMode string `pulumi:"ingestionMode"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The source for the data in the file.
	Source string `pulumi:"source"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a FileImport resource.
type FileImportArgs struct {
	// The content type of this file.
	ContentType pulumi.StringInput
	// File import ID
	FileImportId pulumi.StringPtrInput
	// Represents the imported file.
	ImportFile FileMetadataInput
	// Describes how to ingest the records in the file.
	IngestionMode pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The source for the data in the file.
	Source pulumi.StringInput
	// The name of the workspace.
	WorkspaceName pulumi.StringInput
}

func (FileImportArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fileImportArgs)(nil)).Elem()
}

type FileImportInput interface {
	pulumi.Input

	ToFileImportOutput() FileImportOutput
	ToFileImportOutputWithContext(ctx context.Context) FileImportOutput
}

func (*FileImport) ElementType() reflect.Type {
	return reflect.TypeOf((**FileImport)(nil)).Elem()
}

func (i *FileImport) ToFileImportOutput() FileImportOutput {
	return i.ToFileImportOutputWithContext(context.Background())
}

func (i *FileImport) ToFileImportOutputWithContext(ctx context.Context) FileImportOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileImportOutput)
}

func (i *FileImport) ToOutput(ctx context.Context) pulumix.Output[*FileImport] {
	return pulumix.Output[*FileImport]{
		OutputState: i.ToFileImportOutputWithContext(ctx).OutputState,
	}
}

type FileImportOutput struct{ *pulumi.OutputState }

func (FileImportOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FileImport)(nil)).Elem()
}

func (o FileImportOutput) ToFileImportOutput() FileImportOutput {
	return o
}

func (o FileImportOutput) ToFileImportOutputWithContext(ctx context.Context) FileImportOutput {
	return o
}

func (o FileImportOutput) ToOutput(ctx context.Context) pulumix.Output[*FileImport] {
	return pulumix.Output[*FileImport]{
		OutputState: o.OutputState,
	}
}

// The content type of this file.
func (o FileImportOutput) ContentType() pulumi.StringOutput {
	return o.ApplyT(func(v *FileImport) pulumi.StringOutput { return v.ContentType }).(pulumi.StringOutput)
}

// The time the file was imported.
func (o FileImportOutput) CreatedTimeUTC() pulumi.StringOutput {
	return o.ApplyT(func(v *FileImport) pulumi.StringOutput { return v.CreatedTimeUTC }).(pulumi.StringOutput)
}

// Represents the error file (if the import was ingested with errors or failed the validation).
func (o FileImportOutput) ErrorFile() FileMetadataResponseOutput {
	return o.ApplyT(func(v *FileImport) FileMetadataResponseOutput { return v.ErrorFile }).(FileMetadataResponseOutput)
}

// An ordered list of some of the errors that were encountered during validation.
func (o FileImportOutput) ErrorsPreview() ValidationErrorResponseArrayOutput {
	return o.ApplyT(func(v *FileImport) ValidationErrorResponseArrayOutput { return v.ErrorsPreview }).(ValidationErrorResponseArrayOutput)
}

// The time the files associated with this import are deleted from the storage account.
func (o FileImportOutput) FilesValidUntilTimeUTC() pulumi.StringOutput {
	return o.ApplyT(func(v *FileImport) pulumi.StringOutput { return v.FilesValidUntilTimeUTC }).(pulumi.StringOutput)
}

// Represents the imported file.
func (o FileImportOutput) ImportFile() FileMetadataResponseOutput {
	return o.ApplyT(func(v *FileImport) FileMetadataResponseOutput { return v.ImportFile }).(FileMetadataResponseOutput)
}

// The time the file import record is soft deleted from the database and history.
func (o FileImportOutput) ImportValidUntilTimeUTC() pulumi.StringOutput {
	return o.ApplyT(func(v *FileImport) pulumi.StringOutput { return v.ImportValidUntilTimeUTC }).(pulumi.StringOutput)
}

// The number of records that have been successfully ingested.
func (o FileImportOutput) IngestedRecordCount() pulumi.IntOutput {
	return o.ApplyT(func(v *FileImport) pulumi.IntOutput { return v.IngestedRecordCount }).(pulumi.IntOutput)
}

// Describes how to ingest the records in the file.
func (o FileImportOutput) IngestionMode() pulumi.StringOutput {
	return o.ApplyT(func(v *FileImport) pulumi.StringOutput { return v.IngestionMode }).(pulumi.StringOutput)
}

// The name of the resource
func (o FileImportOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FileImport) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The source for the data in the file.
func (o FileImportOutput) Source() pulumi.StringOutput {
	return o.ApplyT(func(v *FileImport) pulumi.StringOutput { return v.Source }).(pulumi.StringOutput)
}

// The state of the file import.
func (o FileImportOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *FileImport) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o FileImportOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *FileImport) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The number of records in the file.
func (o FileImportOutput) TotalRecordCount() pulumi.IntOutput {
	return o.ApplyT(func(v *FileImport) pulumi.IntOutput { return v.TotalRecordCount }).(pulumi.IntOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o FileImportOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *FileImport) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The number of records that have passed validation.
func (o FileImportOutput) ValidRecordCount() pulumi.IntOutput {
	return o.ApplyT(func(v *FileImport) pulumi.IntOutput { return v.ValidRecordCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(FileImportOutput{})
}
