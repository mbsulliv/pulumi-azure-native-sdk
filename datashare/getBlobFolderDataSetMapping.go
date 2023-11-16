// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package datashare

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get a DataSetMapping in a shareSubscription
// Azure REST API version: 2021-08-01.
func LookupBlobFolderDataSetMapping(ctx *pulumi.Context, args *LookupBlobFolderDataSetMappingArgs, opts ...pulumi.InvokeOption) (*LookupBlobFolderDataSetMappingResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupBlobFolderDataSetMappingResult
	err := ctx.Invoke("azure-native:datashare:getBlobFolderDataSetMapping", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBlobFolderDataSetMappingArgs struct {
	// The name of the share account.
	AccountName string `pulumi:"accountName"`
	// The name of the dataSetMapping.
	DataSetMappingName string `pulumi:"dataSetMappingName"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the shareSubscription.
	ShareSubscriptionName string `pulumi:"shareSubscriptionName"`
}

// A Blob folder data set mapping.
type LookupBlobFolderDataSetMappingResult struct {
	// Container that has the file path.
	ContainerName string `pulumi:"containerName"`
	// The id of the source data set.
	DataSetId string `pulumi:"dataSetId"`
	// Gets the status of the data set mapping.
	DataSetMappingStatus string `pulumi:"dataSetMappingStatus"`
	// The resource id of the azure resource
	Id string `pulumi:"id"`
	// Kind of data set mapping.
	// Expected value is 'BlobFolder'.
	Kind string `pulumi:"kind"`
	// Name of the azure resource
	Name string `pulumi:"name"`
	// Prefix for blob folder
	Prefix string `pulumi:"prefix"`
	// Provisioning state of the data set mapping.
	ProvisioningState string `pulumi:"provisioningState"`
	// Resource group of storage account.
	ResourceGroup string `pulumi:"resourceGroup"`
	// Storage account name of the source data set.
	StorageAccountName string `pulumi:"storageAccountName"`
	// Subscription id of storage account.
	SubscriptionId string `pulumi:"subscriptionId"`
	// System Data of the Azure resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Type of the azure resource
	Type string `pulumi:"type"`
}

func LookupBlobFolderDataSetMappingOutput(ctx *pulumi.Context, args LookupBlobFolderDataSetMappingOutputArgs, opts ...pulumi.InvokeOption) LookupBlobFolderDataSetMappingResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBlobFolderDataSetMappingResult, error) {
			args := v.(LookupBlobFolderDataSetMappingArgs)
			r, err := LookupBlobFolderDataSetMapping(ctx, &args, opts...)
			var s LookupBlobFolderDataSetMappingResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBlobFolderDataSetMappingResultOutput)
}

type LookupBlobFolderDataSetMappingOutputArgs struct {
	// The name of the share account.
	AccountName pulumi.StringInput `pulumi:"accountName"`
	// The name of the dataSetMapping.
	DataSetMappingName pulumi.StringInput `pulumi:"dataSetMappingName"`
	// The resource group name.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the shareSubscription.
	ShareSubscriptionName pulumi.StringInput `pulumi:"shareSubscriptionName"`
}

func (LookupBlobFolderDataSetMappingOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBlobFolderDataSetMappingArgs)(nil)).Elem()
}

// A Blob folder data set mapping.
type LookupBlobFolderDataSetMappingResultOutput struct{ *pulumi.OutputState }

func (LookupBlobFolderDataSetMappingResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBlobFolderDataSetMappingResult)(nil)).Elem()
}

func (o LookupBlobFolderDataSetMappingResultOutput) ToLookupBlobFolderDataSetMappingResultOutput() LookupBlobFolderDataSetMappingResultOutput {
	return o
}

func (o LookupBlobFolderDataSetMappingResultOutput) ToLookupBlobFolderDataSetMappingResultOutputWithContext(ctx context.Context) LookupBlobFolderDataSetMappingResultOutput {
	return o
}

// Container that has the file path.
func (o LookupBlobFolderDataSetMappingResultOutput) ContainerName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobFolderDataSetMappingResult) string { return v.ContainerName }).(pulumi.StringOutput)
}

// The id of the source data set.
func (o LookupBlobFolderDataSetMappingResultOutput) DataSetId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobFolderDataSetMappingResult) string { return v.DataSetId }).(pulumi.StringOutput)
}

// Gets the status of the data set mapping.
func (o LookupBlobFolderDataSetMappingResultOutput) DataSetMappingStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobFolderDataSetMappingResult) string { return v.DataSetMappingStatus }).(pulumi.StringOutput)
}

// The resource id of the azure resource
func (o LookupBlobFolderDataSetMappingResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobFolderDataSetMappingResult) string { return v.Id }).(pulumi.StringOutput)
}

// Kind of data set mapping.
// Expected value is 'BlobFolder'.
func (o LookupBlobFolderDataSetMappingResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobFolderDataSetMappingResult) string { return v.Kind }).(pulumi.StringOutput)
}

// Name of the azure resource
func (o LookupBlobFolderDataSetMappingResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobFolderDataSetMappingResult) string { return v.Name }).(pulumi.StringOutput)
}

// Prefix for blob folder
func (o LookupBlobFolderDataSetMappingResultOutput) Prefix() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobFolderDataSetMappingResult) string { return v.Prefix }).(pulumi.StringOutput)
}

// Provisioning state of the data set mapping.
func (o LookupBlobFolderDataSetMappingResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobFolderDataSetMappingResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Resource group of storage account.
func (o LookupBlobFolderDataSetMappingResultOutput) ResourceGroup() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobFolderDataSetMappingResult) string { return v.ResourceGroup }).(pulumi.StringOutput)
}

// Storage account name of the source data set.
func (o LookupBlobFolderDataSetMappingResultOutput) StorageAccountName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobFolderDataSetMappingResult) string { return v.StorageAccountName }).(pulumi.StringOutput)
}

// Subscription id of storage account.
func (o LookupBlobFolderDataSetMappingResultOutput) SubscriptionId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobFolderDataSetMappingResult) string { return v.SubscriptionId }).(pulumi.StringOutput)
}

// System Data of the Azure resource.
func (o LookupBlobFolderDataSetMappingResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupBlobFolderDataSetMappingResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Type of the azure resource
func (o LookupBlobFolderDataSetMappingResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobFolderDataSetMappingResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBlobFolderDataSetMappingResultOutput{})
}
