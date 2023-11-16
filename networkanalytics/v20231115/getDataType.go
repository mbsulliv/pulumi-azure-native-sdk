// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20231115

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieve data type resource.
func LookupDataType(ctx *pulumi.Context, args *LookupDataTypeArgs, opts ...pulumi.InvokeOption) (*LookupDataTypeResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupDataTypeResult
	err := ctx.Invoke("azure-native:networkanalytics/v20231115:getDataType", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDataTypeArgs struct {
	// The data product resource name
	DataProductName string `pulumi:"dataProductName"`
	// The data type name.
	DataTypeName string `pulumi:"dataTypeName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The data type resource.
type LookupDataTypeResult struct {
	// Field for database cache retention in days.
	DatabaseCacheRetention *int `pulumi:"databaseCacheRetention"`
	// Field for database data retention in days.
	DatabaseRetention *int `pulumi:"databaseRetention"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Latest provisioning state  of data product.
	ProvisioningState string `pulumi:"provisioningState"`
	// State of data type.
	State *string `pulumi:"state"`
	// Reason for the state of data type.
	StateReason string `pulumi:"stateReason"`
	// Field for storage output retention in days.
	StorageOutputRetention *int `pulumi:"storageOutputRetention"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// Url for data visualization.
	VisualizationUrl string `pulumi:"visualizationUrl"`
}

func LookupDataTypeOutput(ctx *pulumi.Context, args LookupDataTypeOutputArgs, opts ...pulumi.InvokeOption) LookupDataTypeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDataTypeResult, error) {
			args := v.(LookupDataTypeArgs)
			r, err := LookupDataType(ctx, &args, opts...)
			var s LookupDataTypeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDataTypeResultOutput)
}

type LookupDataTypeOutputArgs struct {
	// The data product resource name
	DataProductName pulumi.StringInput `pulumi:"dataProductName"`
	// The data type name.
	DataTypeName pulumi.StringInput `pulumi:"dataTypeName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupDataTypeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDataTypeArgs)(nil)).Elem()
}

// The data type resource.
type LookupDataTypeResultOutput struct{ *pulumi.OutputState }

func (LookupDataTypeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDataTypeResult)(nil)).Elem()
}

func (o LookupDataTypeResultOutput) ToLookupDataTypeResultOutput() LookupDataTypeResultOutput {
	return o
}

func (o LookupDataTypeResultOutput) ToLookupDataTypeResultOutputWithContext(ctx context.Context) LookupDataTypeResultOutput {
	return o
}

// Field for database cache retention in days.
func (o LookupDataTypeResultOutput) DatabaseCacheRetention() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupDataTypeResult) *int { return v.DatabaseCacheRetention }).(pulumi.IntPtrOutput)
}

// Field for database data retention in days.
func (o LookupDataTypeResultOutput) DatabaseRetention() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupDataTypeResult) *int { return v.DatabaseRetention }).(pulumi.IntPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupDataTypeResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataTypeResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupDataTypeResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataTypeResult) string { return v.Name }).(pulumi.StringOutput)
}

// Latest provisioning state  of data product.
func (o LookupDataTypeResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataTypeResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// State of data type.
func (o LookupDataTypeResultOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDataTypeResult) *string { return v.State }).(pulumi.StringPtrOutput)
}

// Reason for the state of data type.
func (o LookupDataTypeResultOutput) StateReason() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataTypeResult) string { return v.StateReason }).(pulumi.StringOutput)
}

// Field for storage output retention in days.
func (o LookupDataTypeResultOutput) StorageOutputRetention() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupDataTypeResult) *int { return v.StorageOutputRetention }).(pulumi.IntPtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupDataTypeResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupDataTypeResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupDataTypeResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataTypeResult) string { return v.Type }).(pulumi.StringOutput)
}

// Url for data visualization.
func (o LookupDataTypeResultOutput) VisualizationUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataTypeResult) string { return v.VisualizationUrl }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDataTypeResultOutput{})
}
