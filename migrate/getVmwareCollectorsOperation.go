// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package migrate

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get a VmwareCollector
// Azure REST API version: 2023-03-15.
func LookupVmwareCollectorsOperation(ctx *pulumi.Context, args *LookupVmwareCollectorsOperationArgs, opts ...pulumi.InvokeOption) (*LookupVmwareCollectorsOperationResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupVmwareCollectorsOperationResult
	err := ctx.Invoke("azure-native:migrate:getVmwareCollectorsOperation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVmwareCollectorsOperationArgs struct {
	// Assessment Project Name
	ProjectName string `pulumi:"projectName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// VMware collector ARM name
	VmWareCollectorName string `pulumi:"vmWareCollectorName"`
}

// VMware collector resource.
type LookupVmwareCollectorsOperationResult struct {
	// Gets or sets the collector agent properties.
	AgentProperties *CollectorAgentPropertiesBaseResponse `pulumi:"agentProperties"`
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

func LookupVmwareCollectorsOperationOutput(ctx *pulumi.Context, args LookupVmwareCollectorsOperationOutputArgs, opts ...pulumi.InvokeOption) LookupVmwareCollectorsOperationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVmwareCollectorsOperationResult, error) {
			args := v.(LookupVmwareCollectorsOperationArgs)
			r, err := LookupVmwareCollectorsOperation(ctx, &args, opts...)
			var s LookupVmwareCollectorsOperationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVmwareCollectorsOperationResultOutput)
}

type LookupVmwareCollectorsOperationOutputArgs struct {
	// Assessment Project Name
	ProjectName pulumi.StringInput `pulumi:"projectName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// VMware collector ARM name
	VmWareCollectorName pulumi.StringInput `pulumi:"vmWareCollectorName"`
}

func (LookupVmwareCollectorsOperationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVmwareCollectorsOperationArgs)(nil)).Elem()
}

// VMware collector resource.
type LookupVmwareCollectorsOperationResultOutput struct{ *pulumi.OutputState }

func (LookupVmwareCollectorsOperationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVmwareCollectorsOperationResult)(nil)).Elem()
}

func (o LookupVmwareCollectorsOperationResultOutput) ToLookupVmwareCollectorsOperationResultOutput() LookupVmwareCollectorsOperationResultOutput {
	return o
}

func (o LookupVmwareCollectorsOperationResultOutput) ToLookupVmwareCollectorsOperationResultOutputWithContext(ctx context.Context) LookupVmwareCollectorsOperationResultOutput {
	return o
}

// Gets or sets the collector agent properties.
func (o LookupVmwareCollectorsOperationResultOutput) AgentProperties() CollectorAgentPropertiesBaseResponsePtrOutput {
	return o.ApplyT(func(v LookupVmwareCollectorsOperationResult) *CollectorAgentPropertiesBaseResponse {
		return v.AgentProperties
	}).(CollectorAgentPropertiesBaseResponsePtrOutput)
}

// Gets the Timestamp when collector was created.
func (o LookupVmwareCollectorsOperationResultOutput) CreatedTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVmwareCollectorsOperationResult) string { return v.CreatedTimestamp }).(pulumi.StringOutput)
}

// Gets the discovery site id.
func (o LookupVmwareCollectorsOperationResultOutput) DiscoverySiteId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVmwareCollectorsOperationResult) *string { return v.DiscoverySiteId }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupVmwareCollectorsOperationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVmwareCollectorsOperationResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupVmwareCollectorsOperationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVmwareCollectorsOperationResult) string { return v.Name }).(pulumi.StringOutput)
}

// The status of the last operation.
func (o LookupVmwareCollectorsOperationResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVmwareCollectorsOperationResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupVmwareCollectorsOperationResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupVmwareCollectorsOperationResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupVmwareCollectorsOperationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVmwareCollectorsOperationResult) string { return v.Type }).(pulumi.StringOutput)
}

// Timestamp when collector was last updated.
func (o LookupVmwareCollectorsOperationResultOutput) UpdatedTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVmwareCollectorsOperationResult) string { return v.UpdatedTimestamp }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVmwareCollectorsOperationResultOutput{})
}
