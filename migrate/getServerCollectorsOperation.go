// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package migrate

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get a ServerCollector
// Azure REST API version: 2023-03-15.
func LookupServerCollectorsOperation(ctx *pulumi.Context, args *LookupServerCollectorsOperationArgs, opts ...pulumi.InvokeOption) (*LookupServerCollectorsOperationResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupServerCollectorsOperationResult
	err := ctx.Invoke("azure-native:migrate:getServerCollectorsOperation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServerCollectorsOperationArgs struct {
	// Assessment Project Name
	ProjectName string `pulumi:"projectName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Physical server collector ARM name
	ServerCollectorName string `pulumi:"serverCollectorName"`
}

// Physical server collector resource.
type LookupServerCollectorsOperationResult struct {
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

func LookupServerCollectorsOperationOutput(ctx *pulumi.Context, args LookupServerCollectorsOperationOutputArgs, opts ...pulumi.InvokeOption) LookupServerCollectorsOperationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupServerCollectorsOperationResult, error) {
			args := v.(LookupServerCollectorsOperationArgs)
			r, err := LookupServerCollectorsOperation(ctx, &args, opts...)
			var s LookupServerCollectorsOperationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupServerCollectorsOperationResultOutput)
}

type LookupServerCollectorsOperationOutputArgs struct {
	// Assessment Project Name
	ProjectName pulumi.StringInput `pulumi:"projectName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Physical server collector ARM name
	ServerCollectorName pulumi.StringInput `pulumi:"serverCollectorName"`
}

func (LookupServerCollectorsOperationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServerCollectorsOperationArgs)(nil)).Elem()
}

// Physical server collector resource.
type LookupServerCollectorsOperationResultOutput struct{ *pulumi.OutputState }

func (LookupServerCollectorsOperationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServerCollectorsOperationResult)(nil)).Elem()
}

func (o LookupServerCollectorsOperationResultOutput) ToLookupServerCollectorsOperationResultOutput() LookupServerCollectorsOperationResultOutput {
	return o
}

func (o LookupServerCollectorsOperationResultOutput) ToLookupServerCollectorsOperationResultOutputWithContext(ctx context.Context) LookupServerCollectorsOperationResultOutput {
	return o
}

func (o LookupServerCollectorsOperationResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupServerCollectorsOperationResult] {
	return pulumix.Output[LookupServerCollectorsOperationResult]{
		OutputState: o.OutputState,
	}
}

// Gets or sets the collector agent properties.
func (o LookupServerCollectorsOperationResultOutput) AgentProperties() CollectorAgentPropertiesBaseResponsePtrOutput {
	return o.ApplyT(func(v LookupServerCollectorsOperationResult) *CollectorAgentPropertiesBaseResponse {
		return v.AgentProperties
	}).(CollectorAgentPropertiesBaseResponsePtrOutput)
}

// Gets the Timestamp when collector was created.
func (o LookupServerCollectorsOperationResultOutput) CreatedTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerCollectorsOperationResult) string { return v.CreatedTimestamp }).(pulumi.StringOutput)
}

// Gets the discovery site id.
func (o LookupServerCollectorsOperationResultOutput) DiscoverySiteId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerCollectorsOperationResult) *string { return v.DiscoverySiteId }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupServerCollectorsOperationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerCollectorsOperationResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupServerCollectorsOperationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerCollectorsOperationResult) string { return v.Name }).(pulumi.StringOutput)
}

// The status of the last operation.
func (o LookupServerCollectorsOperationResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServerCollectorsOperationResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupServerCollectorsOperationResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupServerCollectorsOperationResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupServerCollectorsOperationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerCollectorsOperationResult) string { return v.Type }).(pulumi.StringOutput)
}

// Timestamp when collector was last updated.
func (o LookupServerCollectorsOperationResultOutput) UpdatedTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerCollectorsOperationResult) string { return v.UpdatedTimestamp }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupServerCollectorsOperationResultOutput{})
}
