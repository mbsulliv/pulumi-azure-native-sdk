// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get a DataSet in a share
func LookupADLSGen2FileDataSet(ctx *pulumi.Context, args *LookupADLSGen2FileDataSetArgs, opts ...pulumi.InvokeOption) (*LookupADLSGen2FileDataSetResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupADLSGen2FileDataSetResult
	err := ctx.Invoke("azure-native:datashare/v20210801:getADLSGen2FileDataSet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupADLSGen2FileDataSetArgs struct {
	// The name of the share account.
	AccountName string `pulumi:"accountName"`
	// The name of the dataSet.
	DataSetName string `pulumi:"dataSetName"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the share.
	ShareName string `pulumi:"shareName"`
}

// An ADLS Gen 2 file data set.
type LookupADLSGen2FileDataSetResult struct {
	// Unique id for identifying a data set resource
	DataSetId string `pulumi:"dataSetId"`
	// File path within the file system.
	FilePath string `pulumi:"filePath"`
	// File system to which the file belongs.
	FileSystem string `pulumi:"fileSystem"`
	// The resource id of the azure resource
	Id string `pulumi:"id"`
	// Kind of data set.
	// Expected value is 'AdlsGen2File'.
	Kind string `pulumi:"kind"`
	// Name of the azure resource
	Name string `pulumi:"name"`
	// Resource group of storage account
	ResourceGroup string `pulumi:"resourceGroup"`
	// Storage account name of the source data set
	StorageAccountName string `pulumi:"storageAccountName"`
	// Subscription id of storage account
	SubscriptionId string `pulumi:"subscriptionId"`
	// System Data of the Azure resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Type of the azure resource
	Type string `pulumi:"type"`
}

func LookupADLSGen2FileDataSetOutput(ctx *pulumi.Context, args LookupADLSGen2FileDataSetOutputArgs, opts ...pulumi.InvokeOption) LookupADLSGen2FileDataSetResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupADLSGen2FileDataSetResult, error) {
			args := v.(LookupADLSGen2FileDataSetArgs)
			r, err := LookupADLSGen2FileDataSet(ctx, &args, opts...)
			var s LookupADLSGen2FileDataSetResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupADLSGen2FileDataSetResultOutput)
}

type LookupADLSGen2FileDataSetOutputArgs struct {
	// The name of the share account.
	AccountName pulumi.StringInput `pulumi:"accountName"`
	// The name of the dataSet.
	DataSetName pulumi.StringInput `pulumi:"dataSetName"`
	// The resource group name.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the share.
	ShareName pulumi.StringInput `pulumi:"shareName"`
}

func (LookupADLSGen2FileDataSetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupADLSGen2FileDataSetArgs)(nil)).Elem()
}

// An ADLS Gen 2 file data set.
type LookupADLSGen2FileDataSetResultOutput struct{ *pulumi.OutputState }

func (LookupADLSGen2FileDataSetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupADLSGen2FileDataSetResult)(nil)).Elem()
}

func (o LookupADLSGen2FileDataSetResultOutput) ToLookupADLSGen2FileDataSetResultOutput() LookupADLSGen2FileDataSetResultOutput {
	return o
}

func (o LookupADLSGen2FileDataSetResultOutput) ToLookupADLSGen2FileDataSetResultOutputWithContext(ctx context.Context) LookupADLSGen2FileDataSetResultOutput {
	return o
}

// Unique id for identifying a data set resource
func (o LookupADLSGen2FileDataSetResultOutput) DataSetId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen2FileDataSetResult) string { return v.DataSetId }).(pulumi.StringOutput)
}

// File path within the file system.
func (o LookupADLSGen2FileDataSetResultOutput) FilePath() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen2FileDataSetResult) string { return v.FilePath }).(pulumi.StringOutput)
}

// File system to which the file belongs.
func (o LookupADLSGen2FileDataSetResultOutput) FileSystem() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen2FileDataSetResult) string { return v.FileSystem }).(pulumi.StringOutput)
}

// The resource id of the azure resource
func (o LookupADLSGen2FileDataSetResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen2FileDataSetResult) string { return v.Id }).(pulumi.StringOutput)
}

// Kind of data set.
// Expected value is 'AdlsGen2File'.
func (o LookupADLSGen2FileDataSetResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen2FileDataSetResult) string { return v.Kind }).(pulumi.StringOutput)
}

// Name of the azure resource
func (o LookupADLSGen2FileDataSetResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen2FileDataSetResult) string { return v.Name }).(pulumi.StringOutput)
}

// Resource group of storage account
func (o LookupADLSGen2FileDataSetResultOutput) ResourceGroup() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen2FileDataSetResult) string { return v.ResourceGroup }).(pulumi.StringOutput)
}

// Storage account name of the source data set
func (o LookupADLSGen2FileDataSetResultOutput) StorageAccountName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen2FileDataSetResult) string { return v.StorageAccountName }).(pulumi.StringOutput)
}

// Subscription id of storage account
func (o LookupADLSGen2FileDataSetResultOutput) SubscriptionId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen2FileDataSetResult) string { return v.SubscriptionId }).(pulumi.StringOutput)
}

// System Data of the Azure resource.
func (o LookupADLSGen2FileDataSetResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupADLSGen2FileDataSetResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Type of the azure resource
func (o LookupADLSGen2FileDataSetResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupADLSGen2FileDataSetResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupADLSGen2FileDataSetResultOutput{})
}
