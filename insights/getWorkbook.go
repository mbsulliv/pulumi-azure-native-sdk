// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package insights

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get a single workbook by its resourceName.
// Azure REST API version: 2022-04-01.
//
// Other available API versions: 2015-05-01, 2021-03-08, 2021-08-01, 2023-06-01.
func LookupWorkbook(ctx *pulumi.Context, args *LookupWorkbookArgs, opts ...pulumi.InvokeOption) (*LookupWorkbookResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupWorkbookResult
	err := ctx.Invoke("azure-native:insights:getWorkbook", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWorkbookArgs struct {
	// Flag indicating whether or not to return the full content for each applicable workbook. If false, only return summary content for workbooks.
	CanFetchContent *bool `pulumi:"canFetchContent"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the resource.
	ResourceName string `pulumi:"resourceName"`
}

// A workbook definition.
type LookupWorkbookResult struct {
	// Workbook category, as defined by the user at creation time.
	Category string `pulumi:"category"`
	// The description of the workbook.
	Description *string `pulumi:"description"`
	// The user-defined name (display name) of the workbook.
	DisplayName string `pulumi:"displayName"`
	// Resource etag
	Etag *string `pulumi:"etag"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// Identity used for BYOS
	Identity *WorkbookResourceResponseIdentity `pulumi:"identity"`
	// The kind of workbook. Only valid value is shared.
	Kind *string `pulumi:"kind"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The unique revision id for this workbook definition
	Revision string `pulumi:"revision"`
	// Configuration of this particular workbook. Configuration data is a string containing valid JSON
	SerializedData string `pulumi:"serializedData"`
	// ResourceId for a source resource.
	SourceId *string `pulumi:"sourceId"`
	// The resourceId to the storage account when bring your own storage is used
	StorageUri *string `pulumi:"storageUri"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Date and time in UTC of the last modification that was made to this workbook definition.
	TimeModified string `pulumi:"timeModified"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// Unique user id of the specific user that owns this workbook.
	UserId string `pulumi:"userId"`
	// Workbook schema version format, like 'Notebook/1.0', which should match the workbook in serializedData
	Version *string `pulumi:"version"`
}

func LookupWorkbookOutput(ctx *pulumi.Context, args LookupWorkbookOutputArgs, opts ...pulumi.InvokeOption) LookupWorkbookResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWorkbookResult, error) {
			args := v.(LookupWorkbookArgs)
			r, err := LookupWorkbook(ctx, &args, opts...)
			var s LookupWorkbookResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWorkbookResultOutput)
}

type LookupWorkbookOutputArgs struct {
	// Flag indicating whether or not to return the full content for each applicable workbook. If false, only return summary content for workbooks.
	CanFetchContent pulumi.BoolPtrInput `pulumi:"canFetchContent"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the resource.
	ResourceName pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupWorkbookOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkbookArgs)(nil)).Elem()
}

// A workbook definition.
type LookupWorkbookResultOutput struct{ *pulumi.OutputState }

func (LookupWorkbookResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkbookResult)(nil)).Elem()
}

func (o LookupWorkbookResultOutput) ToLookupWorkbookResultOutput() LookupWorkbookResultOutput {
	return o
}

func (o LookupWorkbookResultOutput) ToLookupWorkbookResultOutputWithContext(ctx context.Context) LookupWorkbookResultOutput {
	return o
}

func (o LookupWorkbookResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupWorkbookResult] {
	return pulumix.Output[LookupWorkbookResult]{
		OutputState: o.OutputState,
	}
}

// Workbook category, as defined by the user at creation time.
func (o LookupWorkbookResultOutput) Category() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkbookResult) string { return v.Category }).(pulumi.StringOutput)
}

// The description of the workbook.
func (o LookupWorkbookResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkbookResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The user-defined name (display name) of the workbook.
func (o LookupWorkbookResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkbookResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// Resource etag
func (o LookupWorkbookResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkbookResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupWorkbookResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkbookResult) string { return v.Id }).(pulumi.StringOutput)
}

// Identity used for BYOS
func (o LookupWorkbookResultOutput) Identity() WorkbookResourceResponseIdentityPtrOutput {
	return o.ApplyT(func(v LookupWorkbookResult) *WorkbookResourceResponseIdentity { return v.Identity }).(WorkbookResourceResponseIdentityPtrOutput)
}

// The kind of workbook. Only valid value is shared.
func (o LookupWorkbookResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkbookResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// The geo-location where the resource lives
func (o LookupWorkbookResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkbookResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupWorkbookResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkbookResult) string { return v.Name }).(pulumi.StringOutput)
}

// The unique revision id for this workbook definition
func (o LookupWorkbookResultOutput) Revision() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkbookResult) string { return v.Revision }).(pulumi.StringOutput)
}

// Configuration of this particular workbook. Configuration data is a string containing valid JSON
func (o LookupWorkbookResultOutput) SerializedData() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkbookResult) string { return v.SerializedData }).(pulumi.StringOutput)
}

// ResourceId for a source resource.
func (o LookupWorkbookResultOutput) SourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkbookResult) *string { return v.SourceId }).(pulumi.StringPtrOutput)
}

// The resourceId to the storage account when bring your own storage is used
func (o LookupWorkbookResultOutput) StorageUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkbookResult) *string { return v.StorageUri }).(pulumi.StringPtrOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupWorkbookResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupWorkbookResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupWorkbookResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupWorkbookResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Date and time in UTC of the last modification that was made to this workbook definition.
func (o LookupWorkbookResultOutput) TimeModified() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkbookResult) string { return v.TimeModified }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupWorkbookResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkbookResult) string { return v.Type }).(pulumi.StringOutput)
}

// Unique user id of the specific user that owns this workbook.
func (o LookupWorkbookResultOutput) UserId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkbookResult) string { return v.UserId }).(pulumi.StringOutput)
}

// Workbook schema version format, like 'Notebook/1.0', which should match the workbook in serializedData
func (o LookupWorkbookResultOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkbookResult) *string { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWorkbookResultOutput{})
}
