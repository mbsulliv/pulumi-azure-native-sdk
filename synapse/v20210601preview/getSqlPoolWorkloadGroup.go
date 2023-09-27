// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get a Sql pool's workload group.
func LookupSqlPoolWorkloadGroup(ctx *pulumi.Context, args *LookupSqlPoolWorkloadGroupArgs, opts ...pulumi.InvokeOption) (*LookupSqlPoolWorkloadGroupResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupSqlPoolWorkloadGroupResult
	err := ctx.Invoke("azure-native:synapse/v20210601preview:getSqlPoolWorkloadGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSqlPoolWorkloadGroupArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// SQL pool name
	SqlPoolName string `pulumi:"sqlPoolName"`
	// The name of the workload group.
	WorkloadGroupName string `pulumi:"workloadGroupName"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// Workload group operations for a sql pool
type LookupSqlPoolWorkloadGroupResult struct {
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The workload group importance level.
	Importance *string `pulumi:"importance"`
	// The workload group cap percentage resource.
	MaxResourcePercent int `pulumi:"maxResourcePercent"`
	// The workload group request maximum grant percentage.
	MaxResourcePercentPerRequest *float64 `pulumi:"maxResourcePercentPerRequest"`
	// The workload group minimum percentage resource.
	MinResourcePercent int `pulumi:"minResourcePercent"`
	// The workload group request minimum grant percentage.
	MinResourcePercentPerRequest float64 `pulumi:"minResourcePercentPerRequest"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The workload group query execution timeout.
	QueryExecutionTimeout *int `pulumi:"queryExecutionTimeout"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupSqlPoolWorkloadGroupOutput(ctx *pulumi.Context, args LookupSqlPoolWorkloadGroupOutputArgs, opts ...pulumi.InvokeOption) LookupSqlPoolWorkloadGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSqlPoolWorkloadGroupResult, error) {
			args := v.(LookupSqlPoolWorkloadGroupArgs)
			r, err := LookupSqlPoolWorkloadGroup(ctx, &args, opts...)
			var s LookupSqlPoolWorkloadGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSqlPoolWorkloadGroupResultOutput)
}

type LookupSqlPoolWorkloadGroupOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// SQL pool name
	SqlPoolName pulumi.StringInput `pulumi:"sqlPoolName"`
	// The name of the workload group.
	WorkloadGroupName pulumi.StringInput `pulumi:"workloadGroupName"`
	// The name of the workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupSqlPoolWorkloadGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSqlPoolWorkloadGroupArgs)(nil)).Elem()
}

// Workload group operations for a sql pool
type LookupSqlPoolWorkloadGroupResultOutput struct{ *pulumi.OutputState }

func (LookupSqlPoolWorkloadGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSqlPoolWorkloadGroupResult)(nil)).Elem()
}

func (o LookupSqlPoolWorkloadGroupResultOutput) ToLookupSqlPoolWorkloadGroupResultOutput() LookupSqlPoolWorkloadGroupResultOutput {
	return o
}

func (o LookupSqlPoolWorkloadGroupResultOutput) ToLookupSqlPoolWorkloadGroupResultOutputWithContext(ctx context.Context) LookupSqlPoolWorkloadGroupResultOutput {
	return o
}

func (o LookupSqlPoolWorkloadGroupResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupSqlPoolWorkloadGroupResult] {
	return pulumix.Output[LookupSqlPoolWorkloadGroupResult]{
		OutputState: o.OutputState,
	}
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupSqlPoolWorkloadGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlPoolWorkloadGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

// The workload group importance level.
func (o LookupSqlPoolWorkloadGroupResultOutput) Importance() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSqlPoolWorkloadGroupResult) *string { return v.Importance }).(pulumi.StringPtrOutput)
}

// The workload group cap percentage resource.
func (o LookupSqlPoolWorkloadGroupResultOutput) MaxResourcePercent() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSqlPoolWorkloadGroupResult) int { return v.MaxResourcePercent }).(pulumi.IntOutput)
}

// The workload group request maximum grant percentage.
func (o LookupSqlPoolWorkloadGroupResultOutput) MaxResourcePercentPerRequest() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupSqlPoolWorkloadGroupResult) *float64 { return v.MaxResourcePercentPerRequest }).(pulumi.Float64PtrOutput)
}

// The workload group minimum percentage resource.
func (o LookupSqlPoolWorkloadGroupResultOutput) MinResourcePercent() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSqlPoolWorkloadGroupResult) int { return v.MinResourcePercent }).(pulumi.IntOutput)
}

// The workload group request minimum grant percentage.
func (o LookupSqlPoolWorkloadGroupResultOutput) MinResourcePercentPerRequest() pulumi.Float64Output {
	return o.ApplyT(func(v LookupSqlPoolWorkloadGroupResult) float64 { return v.MinResourcePercentPerRequest }).(pulumi.Float64Output)
}

// The name of the resource
func (o LookupSqlPoolWorkloadGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlPoolWorkloadGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

// The workload group query execution timeout.
func (o LookupSqlPoolWorkloadGroupResultOutput) QueryExecutionTimeout() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupSqlPoolWorkloadGroupResult) *int { return v.QueryExecutionTimeout }).(pulumi.IntPtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupSqlPoolWorkloadGroupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlPoolWorkloadGroupResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSqlPoolWorkloadGroupResultOutput{})
}
