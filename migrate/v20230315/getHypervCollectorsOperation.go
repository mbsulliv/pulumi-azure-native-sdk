// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230315

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get a HypervCollector
func LookupHypervCollectorsOperation(ctx *pulumi.Context, args *LookupHypervCollectorsOperationArgs, opts ...pulumi.InvokeOption) (*LookupHypervCollectorsOperationResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupHypervCollectorsOperationResult
	err := ctx.Invoke("azure-native:migrate/v20230315:getHypervCollectorsOperation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupHypervCollectorsOperationArgs struct {
	// Hyper-V collector ARM name
	HypervCollectorName string `pulumi:"hypervCollectorName"`
	// Assessment Project Name
	ProjectName string `pulumi:"projectName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Hyper-V collector resource.
type LookupHypervCollectorsOperationResult struct {
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

func LookupHypervCollectorsOperationOutput(ctx *pulumi.Context, args LookupHypervCollectorsOperationOutputArgs, opts ...pulumi.InvokeOption) LookupHypervCollectorsOperationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupHypervCollectorsOperationResult, error) {
			args := v.(LookupHypervCollectorsOperationArgs)
			r, err := LookupHypervCollectorsOperation(ctx, &args, opts...)
			var s LookupHypervCollectorsOperationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupHypervCollectorsOperationResultOutput)
}

type LookupHypervCollectorsOperationOutputArgs struct {
	// Hyper-V collector ARM name
	HypervCollectorName pulumi.StringInput `pulumi:"hypervCollectorName"`
	// Assessment Project Name
	ProjectName pulumi.StringInput `pulumi:"projectName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupHypervCollectorsOperationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHypervCollectorsOperationArgs)(nil)).Elem()
}

// Hyper-V collector resource.
type LookupHypervCollectorsOperationResultOutput struct{ *pulumi.OutputState }

func (LookupHypervCollectorsOperationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHypervCollectorsOperationResult)(nil)).Elem()
}

func (o LookupHypervCollectorsOperationResultOutput) ToLookupHypervCollectorsOperationResultOutput() LookupHypervCollectorsOperationResultOutput {
	return o
}

func (o LookupHypervCollectorsOperationResultOutput) ToLookupHypervCollectorsOperationResultOutputWithContext(ctx context.Context) LookupHypervCollectorsOperationResultOutput {
	return o
}

func (o LookupHypervCollectorsOperationResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupHypervCollectorsOperationResult] {
	return pulumix.Output[LookupHypervCollectorsOperationResult]{
		OutputState: o.OutputState,
	}
}

// Gets or sets the collector agent properties.
func (o LookupHypervCollectorsOperationResultOutput) AgentProperties() CollectorAgentPropertiesBaseResponsePtrOutput {
	return o.ApplyT(func(v LookupHypervCollectorsOperationResult) *CollectorAgentPropertiesBaseResponse {
		return v.AgentProperties
	}).(CollectorAgentPropertiesBaseResponsePtrOutput)
}

// Gets the Timestamp when collector was created.
func (o LookupHypervCollectorsOperationResultOutput) CreatedTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHypervCollectorsOperationResult) string { return v.CreatedTimestamp }).(pulumi.StringOutput)
}

// Gets the discovery site id.
func (o LookupHypervCollectorsOperationResultOutput) DiscoverySiteId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHypervCollectorsOperationResult) *string { return v.DiscoverySiteId }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupHypervCollectorsOperationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHypervCollectorsOperationResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupHypervCollectorsOperationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHypervCollectorsOperationResult) string { return v.Name }).(pulumi.StringOutput)
}

// The status of the last operation.
func (o LookupHypervCollectorsOperationResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHypervCollectorsOperationResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupHypervCollectorsOperationResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupHypervCollectorsOperationResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupHypervCollectorsOperationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHypervCollectorsOperationResult) string { return v.Type }).(pulumi.StringOutput)
}

// Timestamp when collector was last updated.
func (o LookupHypervCollectorsOperationResultOutput) UpdatedTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHypervCollectorsOperationResult) string { return v.UpdatedTimestamp }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupHypervCollectorsOperationResultOutput{})
}
