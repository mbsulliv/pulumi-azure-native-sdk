// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get a DataSetMapping in a shareSubscription
func LookupKustoTableDataSetMapping(ctx *pulumi.Context, args *LookupKustoTableDataSetMappingArgs, opts ...pulumi.InvokeOption) (*LookupKustoTableDataSetMappingResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupKustoTableDataSetMappingResult
	err := ctx.Invoke("azure-native:datashare/v20210801:getKustoTableDataSetMapping", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupKustoTableDataSetMappingArgs struct {
	// The name of the share account.
	AccountName string `pulumi:"accountName"`
	// The name of the dataSetMapping.
	DataSetMappingName string `pulumi:"dataSetMappingName"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the shareSubscription.
	ShareSubscriptionName string `pulumi:"shareSubscriptionName"`
}

// A Kusto database data set mapping
type LookupKustoTableDataSetMappingResult struct {
	// The id of the source data set.
	DataSetId string `pulumi:"dataSetId"`
	// Gets the status of the data set mapping.
	DataSetMappingStatus string `pulumi:"dataSetMappingStatus"`
	// The resource id of the azure resource
	Id string `pulumi:"id"`
	// Kind of data set mapping.
	// Expected value is 'KustoTable'.
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

func LookupKustoTableDataSetMappingOutput(ctx *pulumi.Context, args LookupKustoTableDataSetMappingOutputArgs, opts ...pulumi.InvokeOption) LookupKustoTableDataSetMappingResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupKustoTableDataSetMappingResult, error) {
			args := v.(LookupKustoTableDataSetMappingArgs)
			r, err := LookupKustoTableDataSetMapping(ctx, &args, opts...)
			var s LookupKustoTableDataSetMappingResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupKustoTableDataSetMappingResultOutput)
}

type LookupKustoTableDataSetMappingOutputArgs struct {
	// The name of the share account.
	AccountName pulumi.StringInput `pulumi:"accountName"`
	// The name of the dataSetMapping.
	DataSetMappingName pulumi.StringInput `pulumi:"dataSetMappingName"`
	// The resource group name.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the shareSubscription.
	ShareSubscriptionName pulumi.StringInput `pulumi:"shareSubscriptionName"`
}

func (LookupKustoTableDataSetMappingOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupKustoTableDataSetMappingArgs)(nil)).Elem()
}

// A Kusto database data set mapping
type LookupKustoTableDataSetMappingResultOutput struct{ *pulumi.OutputState }

func (LookupKustoTableDataSetMappingResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupKustoTableDataSetMappingResult)(nil)).Elem()
}

func (o LookupKustoTableDataSetMappingResultOutput) ToLookupKustoTableDataSetMappingResultOutput() LookupKustoTableDataSetMappingResultOutput {
	return o
}

func (o LookupKustoTableDataSetMappingResultOutput) ToLookupKustoTableDataSetMappingResultOutputWithContext(ctx context.Context) LookupKustoTableDataSetMappingResultOutput {
	return o
}

func (o LookupKustoTableDataSetMappingResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupKustoTableDataSetMappingResult] {
	return pulumix.Output[LookupKustoTableDataSetMappingResult]{
		OutputState: o.OutputState,
	}
}

// The id of the source data set.
func (o LookupKustoTableDataSetMappingResultOutput) DataSetId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoTableDataSetMappingResult) string { return v.DataSetId }).(pulumi.StringOutput)
}

// Gets the status of the data set mapping.
func (o LookupKustoTableDataSetMappingResultOutput) DataSetMappingStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoTableDataSetMappingResult) string { return v.DataSetMappingStatus }).(pulumi.StringOutput)
}

// The resource id of the azure resource
func (o LookupKustoTableDataSetMappingResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoTableDataSetMappingResult) string { return v.Id }).(pulumi.StringOutput)
}

// Kind of data set mapping.
// Expected value is 'KustoTable'.
func (o LookupKustoTableDataSetMappingResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoTableDataSetMappingResult) string { return v.Kind }).(pulumi.StringOutput)
}

// Resource id of the sink kusto cluster.
func (o LookupKustoTableDataSetMappingResultOutput) KustoClusterResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoTableDataSetMappingResult) string { return v.KustoClusterResourceId }).(pulumi.StringOutput)
}

// Location of the sink kusto cluster.
func (o LookupKustoTableDataSetMappingResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoTableDataSetMappingResult) string { return v.Location }).(pulumi.StringOutput)
}

// Name of the azure resource
func (o LookupKustoTableDataSetMappingResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoTableDataSetMappingResult) string { return v.Name }).(pulumi.StringOutput)
}

// Provisioning state of the data set mapping.
func (o LookupKustoTableDataSetMappingResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoTableDataSetMappingResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// System Data of the Azure resource.
func (o LookupKustoTableDataSetMappingResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupKustoTableDataSetMappingResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Type of the azure resource
func (o LookupKustoTableDataSetMappingResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKustoTableDataSetMappingResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupKustoTableDataSetMappingResultOutput{})
}
