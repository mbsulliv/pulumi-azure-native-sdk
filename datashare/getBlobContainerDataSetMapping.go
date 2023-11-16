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
func LookupBlobContainerDataSetMapping(ctx *pulumi.Context, args *LookupBlobContainerDataSetMappingArgs, opts ...pulumi.InvokeOption) (*LookupBlobContainerDataSetMappingResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupBlobContainerDataSetMappingResult
	err := ctx.Invoke("azure-native:datashare:getBlobContainerDataSetMapping", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBlobContainerDataSetMappingArgs struct {
	// The name of the share account.
	AccountName string `pulumi:"accountName"`
	// The name of the dataSetMapping.
	DataSetMappingName string `pulumi:"dataSetMappingName"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the shareSubscription.
	ShareSubscriptionName string `pulumi:"shareSubscriptionName"`
}

// A Blob container data set mapping.
type LookupBlobContainerDataSetMappingResult struct {
	// BLOB Container name.
	ContainerName string `pulumi:"containerName"`
	// The id of the source data set.
	DataSetId string `pulumi:"dataSetId"`
	// Gets the status of the data set mapping.
	DataSetMappingStatus string `pulumi:"dataSetMappingStatus"`
	// The resource id of the azure resource
	Id string `pulumi:"id"`
	// Kind of data set mapping.
	// Expected value is 'Container'.
	Kind string `pulumi:"kind"`
	// Name of the azure resource
	Name string `pulumi:"name"`
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

func LookupBlobContainerDataSetMappingOutput(ctx *pulumi.Context, args LookupBlobContainerDataSetMappingOutputArgs, opts ...pulumi.InvokeOption) LookupBlobContainerDataSetMappingResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBlobContainerDataSetMappingResult, error) {
			args := v.(LookupBlobContainerDataSetMappingArgs)
			r, err := LookupBlobContainerDataSetMapping(ctx, &args, opts...)
			var s LookupBlobContainerDataSetMappingResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBlobContainerDataSetMappingResultOutput)
}

type LookupBlobContainerDataSetMappingOutputArgs struct {
	// The name of the share account.
	AccountName pulumi.StringInput `pulumi:"accountName"`
	// The name of the dataSetMapping.
	DataSetMappingName pulumi.StringInput `pulumi:"dataSetMappingName"`
	// The resource group name.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the shareSubscription.
	ShareSubscriptionName pulumi.StringInput `pulumi:"shareSubscriptionName"`
}

func (LookupBlobContainerDataSetMappingOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBlobContainerDataSetMappingArgs)(nil)).Elem()
}

// A Blob container data set mapping.
type LookupBlobContainerDataSetMappingResultOutput struct{ *pulumi.OutputState }

func (LookupBlobContainerDataSetMappingResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBlobContainerDataSetMappingResult)(nil)).Elem()
}

func (o LookupBlobContainerDataSetMappingResultOutput) ToLookupBlobContainerDataSetMappingResultOutput() LookupBlobContainerDataSetMappingResultOutput {
	return o
}

func (o LookupBlobContainerDataSetMappingResultOutput) ToLookupBlobContainerDataSetMappingResultOutputWithContext(ctx context.Context) LookupBlobContainerDataSetMappingResultOutput {
	return o
}

// BLOB Container name.
func (o LookupBlobContainerDataSetMappingResultOutput) ContainerName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobContainerDataSetMappingResult) string { return v.ContainerName }).(pulumi.StringOutput)
}

// The id of the source data set.
func (o LookupBlobContainerDataSetMappingResultOutput) DataSetId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobContainerDataSetMappingResult) string { return v.DataSetId }).(pulumi.StringOutput)
}

// Gets the status of the data set mapping.
func (o LookupBlobContainerDataSetMappingResultOutput) DataSetMappingStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobContainerDataSetMappingResult) string { return v.DataSetMappingStatus }).(pulumi.StringOutput)
}

// The resource id of the azure resource
func (o LookupBlobContainerDataSetMappingResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobContainerDataSetMappingResult) string { return v.Id }).(pulumi.StringOutput)
}

// Kind of data set mapping.
// Expected value is 'Container'.
func (o LookupBlobContainerDataSetMappingResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobContainerDataSetMappingResult) string { return v.Kind }).(pulumi.StringOutput)
}

// Name of the azure resource
func (o LookupBlobContainerDataSetMappingResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobContainerDataSetMappingResult) string { return v.Name }).(pulumi.StringOutput)
}

// Provisioning state of the data set mapping.
func (o LookupBlobContainerDataSetMappingResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobContainerDataSetMappingResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Resource group of storage account.
func (o LookupBlobContainerDataSetMappingResultOutput) ResourceGroup() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobContainerDataSetMappingResult) string { return v.ResourceGroup }).(pulumi.StringOutput)
}

// Storage account name of the source data set.
func (o LookupBlobContainerDataSetMappingResultOutput) StorageAccountName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobContainerDataSetMappingResult) string { return v.StorageAccountName }).(pulumi.StringOutput)
}

// Subscription id of storage account.
func (o LookupBlobContainerDataSetMappingResultOutput) SubscriptionId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobContainerDataSetMappingResult) string { return v.SubscriptionId }).(pulumi.StringOutput)
}

// System Data of the Azure resource.
func (o LookupBlobContainerDataSetMappingResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupBlobContainerDataSetMappingResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Type of the azure resource
func (o LookupBlobContainerDataSetMappingResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobContainerDataSetMappingResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBlobContainerDataSetMappingResultOutput{})
}
