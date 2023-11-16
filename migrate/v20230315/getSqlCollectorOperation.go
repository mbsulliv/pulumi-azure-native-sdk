// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230315

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get a SqlCollector
func LookupSqlCollectorOperation(ctx *pulumi.Context, args *LookupSqlCollectorOperationArgs, opts ...pulumi.InvokeOption) (*LookupSqlCollectorOperationResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupSqlCollectorOperationResult
	err := ctx.Invoke("azure-native:migrate/v20230315:getSqlCollectorOperation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSqlCollectorOperationArgs struct {
	// Sql collector ARM name.
	CollectorName string `pulumi:"collectorName"`
	// Assessment Project Name
	ProjectName string `pulumi:"projectName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The SQL collector REST object.
type LookupSqlCollectorOperationResult struct {
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

func LookupSqlCollectorOperationOutput(ctx *pulumi.Context, args LookupSqlCollectorOperationOutputArgs, opts ...pulumi.InvokeOption) LookupSqlCollectorOperationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSqlCollectorOperationResult, error) {
			args := v.(LookupSqlCollectorOperationArgs)
			r, err := LookupSqlCollectorOperation(ctx, &args, opts...)
			var s LookupSqlCollectorOperationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSqlCollectorOperationResultOutput)
}

type LookupSqlCollectorOperationOutputArgs struct {
	// Sql collector ARM name.
	CollectorName pulumi.StringInput `pulumi:"collectorName"`
	// Assessment Project Name
	ProjectName pulumi.StringInput `pulumi:"projectName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupSqlCollectorOperationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSqlCollectorOperationArgs)(nil)).Elem()
}

// The SQL collector REST object.
type LookupSqlCollectorOperationResultOutput struct{ *pulumi.OutputState }

func (LookupSqlCollectorOperationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSqlCollectorOperationResult)(nil)).Elem()
}

func (o LookupSqlCollectorOperationResultOutput) ToLookupSqlCollectorOperationResultOutput() LookupSqlCollectorOperationResultOutput {
	return o
}

func (o LookupSqlCollectorOperationResultOutput) ToLookupSqlCollectorOperationResultOutputWithContext(ctx context.Context) LookupSqlCollectorOperationResultOutput {
	return o
}

// Gets or sets the collector agent properties.
func (o LookupSqlCollectorOperationResultOutput) AgentProperties() CollectorAgentPropertiesBaseResponsePtrOutput {
	return o.ApplyT(func(v LookupSqlCollectorOperationResult) *CollectorAgentPropertiesBaseResponse {
		return v.AgentProperties
	}).(CollectorAgentPropertiesBaseResponsePtrOutput)
}

// Gets the Timestamp when collector was created.
func (o LookupSqlCollectorOperationResultOutput) CreatedTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlCollectorOperationResult) string { return v.CreatedTimestamp }).(pulumi.StringOutput)
}

// Gets the discovery site id.
func (o LookupSqlCollectorOperationResultOutput) DiscoverySiteId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSqlCollectorOperationResult) *string { return v.DiscoverySiteId }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupSqlCollectorOperationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlCollectorOperationResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupSqlCollectorOperationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlCollectorOperationResult) string { return v.Name }).(pulumi.StringOutput)
}

// The status of the last operation.
func (o LookupSqlCollectorOperationResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSqlCollectorOperationResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupSqlCollectorOperationResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupSqlCollectorOperationResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupSqlCollectorOperationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlCollectorOperationResult) string { return v.Type }).(pulumi.StringOutput)
}

// Timestamp when collector was last updated.
func (o LookupSqlCollectorOperationResultOutput) UpdatedTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlCollectorOperationResult) string { return v.UpdatedTimestamp }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSqlCollectorOperationResultOutput{})
}
