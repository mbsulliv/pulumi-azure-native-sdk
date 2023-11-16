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
func LookupKustoClusterDataSetMapping(ctx *pulumi.Context, args *LookupKustoClusterDataSetMappingArgs, opts ...pulumi.InvokeOption) (*LookupKustoClusterDataSetMappingResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupKustoClusterDataSetMappingResult
	err := ctx.Invoke("azure-native:datashare:getKustoClusterDataSetMapping", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupKustoClusterDataSetMappingArgs struct {
	// The name of the share account.
	AccountName string `pulumi:"accountName"`
	// The name of the dataSetMapping.
	DataSetMappingName string `pulumi:"dataSetMappingName"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the shareSubscription.
	ShareSubscriptionName string `pulumi:"shareSubscriptionName"`
}

// A Kusto cluster data set mapping
type LookupKustoClusterDataSetMappingResult struct {
	// The id of the source data set.
	DataSetId string `pulumi:"dataSetId"`
	// Gets the status of the data set mapping.
	DataSetMappingStatus string `pulumi:"dataSetMappingStatus"`
	// The resource id of the azure resource
	Id string `pulumi:"id"`
	// Kind of data set mapping.
	// Expected value is 'KustoCluster'.
	Kind string `pulumi:"kind"`
	// Resource id of the sink kusto cluster.
	KustoClusterResourceId string `pulumi:"kustoClusterResourceId"`
	// Location of the sink kusto cluster.
	Location string `pulumi:"location"`
	// Name of the azure resource
	Name string `pulumi:"name"`
	// Provisioning state of the data set mapping.
	ProvisioningState string `pulumi:"provisioningState"`
	// System Data of the Azure resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Type of the azure resource
	Type string `pulumi:"type"`
}

func LookupKustoClusterDataSetMappingOutput(ctx *pulumi.Context, args LookupKustoClusterDataSetMappingOutputArgs, opts ...pulumi.InvokeOption) LookupKustoClusterDataSetMappingResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupKustoClusterDataSetMappingResult, error) {
			args := v.(LookupKustoClusterDataSetMappingArgs)
			r, err := LookupKustoClusterDataSetMapping(ctx, &args, opts...)
			var s LookupKustoClusterDataSetMappingResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupKustoClusterDataSetMappingResultOutput)
}

type LookupKustoClusterDataSetMappingOutputArgs struct {
	// The name of the share account.
	AccountName pulumi.StringInput `pulumi:"accountName"`
	// The name of the dataSetMapping.
	DataSetMappingName pulumi.StringInput `pulumi:"dataSetMappingName"`
	// The resource group name.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the shareSubscription.
	ShareSubscriptionName pulumi.StringInput `pulumi:"shareSubscriptionName"`
}

func (LookupKustoClusterDataSetMappingOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupKustoClusterDataSetMappingArgs)(nil)).Elem()
}

// A Kusto cluster data set mapping
type LookupKustoClusterDataSetMappingResultOutput struct{ *pulumi.OutputState }

func (LookupKustoClusterDataSetMappingResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupKustoClusterDataSetMappingResult)(nil)).Elem()
}

func (o LookupKustoClusterDataSetMappingResultOutput) ToLookupKustoClusterDataSetMappingResultOutput() LookupKustoClusterDataSetMappingResultOutput {
	return o
}

func (o LookupKustoClusterDataSetMappingResultOutput) ToLookupKustoClusterDataSetMappingResultOutputWithContext(ctx context.Context) LookupKustoClusterDataSetMappingResultOutput {
	return o
}

// The id of the source data set.
func (o LookupKustoClusterDataSetMappingResultOutput) DataSetId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoClusterDataSetMappingResult) string { return v.DataSetId }).(pulumi.StringOutput)
}

// Gets the status of the data set mapping.
func (o LookupKustoClusterDataSetMappingResultOutput) DataSetMappingStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoClusterDataSetMappingResult) string { return v.DataSetMappingStatus }).(pulumi.StringOutput)
}

// The resource id of the azure resource
func (o LookupKustoClusterDataSetMappingResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoClusterDataSetMappingResult) string { return v.Id }).(pulumi.StringOutput)
}

// Kind of data set mapping.
// Expected value is 'KustoCluster'.
func (o LookupKustoClusterDataSetMappingResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoClusterDataSetMappingResult) string { return v.Kind }).(pulumi.StringOutput)
}

// Resource id of the sink kusto cluster.
func (o LookupKustoClusterDataSetMappingResultOutput) KustoClusterResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoClusterDataSetMappingResult) string { return v.KustoClusterResourceId }).(pulumi.StringOutput)
}

// Location of the sink kusto cluster.
func (o LookupKustoClusterDataSetMappingResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoClusterDataSetMappingResult) string { return v.Location }).(pulumi.StringOutput)
}

// Name of the azure resource
func (o LookupKustoClusterDataSetMappingResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoClusterDataSetMappingResult) string { return v.Name }).(pulumi.StringOutput)
}

// Provisioning state of the data set mapping.
func (o LookupKustoClusterDataSetMappingResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoClusterDataSetMappingResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// System Data of the Azure resource.
func (o LookupKustoClusterDataSetMappingResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupKustoClusterDataSetMappingResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Type of the azure resource
func (o LookupKustoClusterDataSetMappingResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoClusterDataSetMappingResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupKustoClusterDataSetMappingResultOutput{})
}
