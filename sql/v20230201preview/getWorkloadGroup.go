// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets a workload group
func LookupWorkloadGroup(ctx *pulumi.Context, args *LookupWorkloadGroupArgs, opts ...pulumi.InvokeOption) (*LookupWorkloadGroupResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupWorkloadGroupResult
	err := ctx.Invoke("azure-native:sql/v20230201preview:getWorkloadGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWorkloadGroupArgs struct {
	// The name of the database.
	DatabaseName string `pulumi:"databaseName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
	// The name of the workload group.
	WorkloadGroupName string `pulumi:"workloadGroupName"`
}

// Workload group operations for a data warehouse
type LookupWorkloadGroupResult struct {
	// Resource ID.
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
	// Resource name.
	Name string `pulumi:"name"`
	// The workload group query execution timeout.
	QueryExecutionTimeout *int `pulumi:"queryExecutionTimeout"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupWorkloadGroupOutput(ctx *pulumi.Context, args LookupWorkloadGroupOutputArgs, opts ...pulumi.InvokeOption) LookupWorkloadGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWorkloadGroupResult, error) {
			args := v.(LookupWorkloadGroupArgs)
			r, err := LookupWorkloadGroup(ctx, &args, opts...)
			var s LookupWorkloadGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWorkloadGroupResultOutput)
}

type LookupWorkloadGroupOutputArgs struct {
	// The name of the database.
	DatabaseName pulumi.StringInput `pulumi:"databaseName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName pulumi.StringInput `pulumi:"serverName"`
	// The name of the workload group.
	WorkloadGroupName pulumi.StringInput `pulumi:"workloadGroupName"`
}

func (LookupWorkloadGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkloadGroupArgs)(nil)).Elem()
}

// Workload group operations for a data warehouse
type LookupWorkloadGroupResultOutput struct{ *pulumi.OutputState }

func (LookupWorkloadGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkloadGroupResult)(nil)).Elem()
}

func (o LookupWorkloadGroupResultOutput) ToLookupWorkloadGroupResultOutput() LookupWorkloadGroupResultOutput {
	return o
}

func (o LookupWorkloadGroupResultOutput) ToLookupWorkloadGroupResultOutputWithContext(ctx context.Context) LookupWorkloadGroupResultOutput {
	return o
}

func (o LookupWorkloadGroupResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupWorkloadGroupResult] {
	return pulumix.Output[LookupWorkloadGroupResult]{
		OutputState: o.OutputState,
	}
}

// Resource ID.
func (o LookupWorkloadGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkloadGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

// The workload group importance level.
func (o LookupWorkloadGroupResultOutput) Importance() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkloadGroupResult) *string { return v.Importance }).(pulumi.StringPtrOutput)
}

// The workload group cap percentage resource.
func (o LookupWorkloadGroupResultOutput) MaxResourcePercent() pulumi.IntOutput {
	return o.ApplyT(func(v LookupWorkloadGroupResult) int { return v.MaxResourcePercent }).(pulumi.IntOutput)
}

// The workload group request maximum grant percentage.
func (o LookupWorkloadGroupResultOutput) MaxResourcePercentPerRequest() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupWorkloadGroupResult) *float64 { return v.MaxResourcePercentPerRequest }).(pulumi.Float64PtrOutput)
}

// The workload group minimum percentage resource.
func (o LookupWorkloadGroupResultOutput) MinResourcePercent() pulumi.IntOutput {
	return o.ApplyT(func(v LookupWorkloadGroupResult) int { return v.MinResourcePercent }).(pulumi.IntOutput)
}

// The workload group request minimum grant percentage.
func (o LookupWorkloadGroupResultOutput) MinResourcePercentPerRequest() pulumi.Float64Output {
	return o.ApplyT(func(v LookupWorkloadGroupResult) float64 { return v.MinResourcePercentPerRequest }).(pulumi.Float64Output)
}

// Resource name.
func (o LookupWorkloadGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkloadGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

// The workload group query execution timeout.
func (o LookupWorkloadGroupResultOutput) QueryExecutionTimeout() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupWorkloadGroupResult) *int { return v.QueryExecutionTimeout }).(pulumi.IntPtrOutput)
}

// Resource type.
func (o LookupWorkloadGroupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkloadGroupResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWorkloadGroupResultOutput{})
}
