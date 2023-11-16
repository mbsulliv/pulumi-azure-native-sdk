// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230315

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get a ImportCollector
func LookupImportCollectorsOperation(ctx *pulumi.Context, args *LookupImportCollectorsOperationArgs, opts ...pulumi.InvokeOption) (*LookupImportCollectorsOperationResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupImportCollectorsOperationResult
	err := ctx.Invoke("azure-native:migrate/v20230315:getImportCollectorsOperation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupImportCollectorsOperationArgs struct {
	// Import collector ARM name
	ImportCollectorName string `pulumi:"importCollectorName"`
	// Assessment Project Name
	ProjectName string `pulumi:"projectName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Import collector resource.
type LookupImportCollectorsOperationResult struct {
	// Gets the Timestamp when collector was created.
	CreatedTimestamp string `pulumi:"createdTimestamp"`
	// Gets the discovery site id.
	DiscoverySiteId *string `pulumi:"discoverySiteId"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The status of the last operation.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// Timestamp when collector was last updated.
	UpdatedTimestamp string `pulumi:"updatedTimestamp"`
}

func LookupImportCollectorsOperationOutput(ctx *pulumi.Context, args LookupImportCollectorsOperationOutputArgs, opts ...pulumi.InvokeOption) LookupImportCollectorsOperationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupImportCollectorsOperationResult, error) {
			args := v.(LookupImportCollectorsOperationArgs)
			r, err := LookupImportCollectorsOperation(ctx, &args, opts...)
			var s LookupImportCollectorsOperationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupImportCollectorsOperationResultOutput)
}

type LookupImportCollectorsOperationOutputArgs struct {
	// Import collector ARM name
	ImportCollectorName pulumi.StringInput `pulumi:"importCollectorName"`
	// Assessment Project Name
	ProjectName pulumi.StringInput `pulumi:"projectName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupImportCollectorsOperationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupImportCollectorsOperationArgs)(nil)).Elem()
}

// Import collector resource.
type LookupImportCollectorsOperationResultOutput struct{ *pulumi.OutputState }

func (LookupImportCollectorsOperationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupImportCollectorsOperationResult)(nil)).Elem()
}

func (o LookupImportCollectorsOperationResultOutput) ToLookupImportCollectorsOperationResultOutput() LookupImportCollectorsOperationResultOutput {
	return o
}

func (o LookupImportCollectorsOperationResultOutput) ToLookupImportCollectorsOperationResultOutputWithContext(ctx context.Context) LookupImportCollectorsOperationResultOutput {
	return o
}

// Gets the Timestamp when collector was created.
func (o LookupImportCollectorsOperationResultOutput) CreatedTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImportCollectorsOperationResult) string { return v.CreatedTimestamp }).(pulumi.StringOutput)
}

// Gets the discovery site id.
func (o LookupImportCollectorsOperationResultOutput) DiscoverySiteId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupImportCollectorsOperationResult) *string { return v.DiscoverySiteId }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupImportCollectorsOperationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImportCollectorsOperationResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupImportCollectorsOperationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImportCollectorsOperationResult) string { return v.Name }).(pulumi.StringOutput)
}

// The status of the last operation.
func (o LookupImportCollectorsOperationResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupImportCollectorsOperationResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupImportCollectorsOperationResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupImportCollectorsOperationResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupImportCollectorsOperationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImportCollectorsOperationResult) string { return v.Type }).(pulumi.StringOutput)
}

// Timestamp when collector was last updated.
func (o LookupImportCollectorsOperationResultOutput) UpdatedTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImportCollectorsOperationResult) string { return v.UpdatedTimestamp }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupImportCollectorsOperationResultOutput{})
}
