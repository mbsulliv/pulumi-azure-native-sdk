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

// Get a DataSet in a share
func LookupSynapseWorkspaceSqlPoolTableDataSet(ctx *pulumi.Context, args *LookupSynapseWorkspaceSqlPoolTableDataSetArgs, opts ...pulumi.InvokeOption) (*LookupSynapseWorkspaceSqlPoolTableDataSetResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupSynapseWorkspaceSqlPoolTableDataSetResult
	err := ctx.Invoke("azure-native:datashare/v20210801:getSynapseWorkspaceSqlPoolTableDataSet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSynapseWorkspaceSqlPoolTableDataSetArgs struct {
	// The name of the share account.
	AccountName string `pulumi:"accountName"`
	// The name of the dataSet.
	DataSetName string `pulumi:"dataSetName"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the share.
	ShareName string `pulumi:"shareName"`
}

// A Synapse Workspace Sql Pool Table data set.
type LookupSynapseWorkspaceSqlPoolTableDataSetResult struct {
	// Unique id for identifying a data set resource
	DataSetId string `pulumi:"dataSetId"`
	// The resource id of the azure resource
	Id string `pulumi:"id"`
	// Kind of data set.
	// Expected value is 'SynapseWorkspaceSqlPoolTable'.
	Kind string `pulumi:"kind"`
	// Name of the azure resource
	Name string `pulumi:"name"`
	// Resource id of the Synapse Workspace SQL Pool Table
	SynapseWorkspaceSqlPoolTableResourceId string `pulumi:"synapseWorkspaceSqlPoolTableResourceId"`
	// System Data of the Azure resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Type of the azure resource
	Type string `pulumi:"type"`
}

func LookupSynapseWorkspaceSqlPoolTableDataSetOutput(ctx *pulumi.Context, args LookupSynapseWorkspaceSqlPoolTableDataSetOutputArgs, opts ...pulumi.InvokeOption) LookupSynapseWorkspaceSqlPoolTableDataSetResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSynapseWorkspaceSqlPoolTableDataSetResult, error) {
			args := v.(LookupSynapseWorkspaceSqlPoolTableDataSetArgs)
			r, err := LookupSynapseWorkspaceSqlPoolTableDataSet(ctx, &args, opts...)
			var s LookupSynapseWorkspaceSqlPoolTableDataSetResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSynapseWorkspaceSqlPoolTableDataSetResultOutput)
}

type LookupSynapseWorkspaceSqlPoolTableDataSetOutputArgs struct {
	// The name of the share account.
	AccountName pulumi.StringInput `pulumi:"accountName"`
	// The name of the dataSet.
	DataSetName pulumi.StringInput `pulumi:"dataSetName"`
	// The resource group name.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the share.
	ShareName pulumi.StringInput `pulumi:"shareName"`
}

func (LookupSynapseWorkspaceSqlPoolTableDataSetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSynapseWorkspaceSqlPoolTableDataSetArgs)(nil)).Elem()
}

// A Synapse Workspace Sql Pool Table data set.
type LookupSynapseWorkspaceSqlPoolTableDataSetResultOutput struct{ *pulumi.OutputState }

func (LookupSynapseWorkspaceSqlPoolTableDataSetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSynapseWorkspaceSqlPoolTableDataSetResult)(nil)).Elem()
}

func (o LookupSynapseWorkspaceSqlPoolTableDataSetResultOutput) ToLookupSynapseWorkspaceSqlPoolTableDataSetResultOutput() LookupSynapseWorkspaceSqlPoolTableDataSetResultOutput {
	return o
}

func (o LookupSynapseWorkspaceSqlPoolTableDataSetResultOutput) ToLookupSynapseWorkspaceSqlPoolTableDataSetResultOutputWithContext(ctx context.Context) LookupSynapseWorkspaceSqlPoolTableDataSetResultOutput {
	return o
}

func (o LookupSynapseWorkspaceSqlPoolTableDataSetResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupSynapseWorkspaceSqlPoolTableDataSetResult] {
	return pulumix.Output[LookupSynapseWorkspaceSqlPoolTableDataSetResult]{
		OutputState: o.OutputState,
	}
}

// Unique id for identifying a data set resource
func (o LookupSynapseWorkspaceSqlPoolTableDataSetResultOutput) DataSetId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSynapseWorkspaceSqlPoolTableDataSetResult) string { return v.DataSetId }).(pulumi.StringOutput)
}

// The resource id of the azure resource
func (o LookupSynapseWorkspaceSqlPoolTableDataSetResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSynapseWorkspaceSqlPoolTableDataSetResult) string { return v.Id }).(pulumi.StringOutput)
}

// Kind of data set.
// Expected value is 'SynapseWorkspaceSqlPoolTable'.
func (o LookupSynapseWorkspaceSqlPoolTableDataSetResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSynapseWorkspaceSqlPoolTableDataSetResult) string { return v.Kind }).(pulumi.StringOutput)
}

// Name of the azure resource
func (o LookupSynapseWorkspaceSqlPoolTableDataSetResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSynapseWorkspaceSqlPoolTableDataSetResult) string { return v.Name }).(pulumi.StringOutput)
}

// Resource id of the Synapse Workspace SQL Pool Table
func (o LookupSynapseWorkspaceSqlPoolTableDataSetResultOutput) SynapseWorkspaceSqlPoolTableResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSynapseWorkspaceSqlPoolTableDataSetResult) string {
		return v.SynapseWorkspaceSqlPoolTableResourceId
	}).(pulumi.StringOutput)
}

// System Data of the Azure resource.
func (o LookupSynapseWorkspaceSqlPoolTableDataSetResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupSynapseWorkspaceSqlPoolTableDataSetResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Type of the azure resource
func (o LookupSynapseWorkspaceSqlPoolTableDataSetResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSynapseWorkspaceSqlPoolTableDataSetResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSynapseWorkspaceSqlPoolTableDataSetResultOutput{})
}
