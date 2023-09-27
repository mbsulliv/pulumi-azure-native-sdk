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

// Get a workload classifier of Sql pool's workload group.
func LookupSqlPoolWorkloadClassifier(ctx *pulumi.Context, args *LookupSqlPoolWorkloadClassifierArgs, opts ...pulumi.InvokeOption) (*LookupSqlPoolWorkloadClassifierResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupSqlPoolWorkloadClassifierResult
	err := ctx.Invoke("azure-native:synapse/v20210601preview:getSqlPoolWorkloadClassifier", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSqlPoolWorkloadClassifierArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// SQL pool name
	SqlPoolName string `pulumi:"sqlPoolName"`
	// The name of the workload classifier.
	WorkloadClassifierName string `pulumi:"workloadClassifierName"`
	// The name of the workload group.
	WorkloadGroupName string `pulumi:"workloadGroupName"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// Workload classifier operations for a data warehouse
type LookupSqlPoolWorkloadClassifierResult struct {
	// The workload classifier context.
	Context *string `pulumi:"context"`
	// The workload classifier end time for classification.
	EndTime *string `pulumi:"endTime"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The workload classifier importance.
	Importance *string `pulumi:"importance"`
	// The workload classifier label.
	Label *string `pulumi:"label"`
	// The workload classifier member name.
	MemberName string `pulumi:"memberName"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The workload classifier start time for classification.
	StartTime *string `pulumi:"startTime"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupSqlPoolWorkloadClassifierOutput(ctx *pulumi.Context, args LookupSqlPoolWorkloadClassifierOutputArgs, opts ...pulumi.InvokeOption) LookupSqlPoolWorkloadClassifierResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSqlPoolWorkloadClassifierResult, error) {
			args := v.(LookupSqlPoolWorkloadClassifierArgs)
			r, err := LookupSqlPoolWorkloadClassifier(ctx, &args, opts...)
			var s LookupSqlPoolWorkloadClassifierResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSqlPoolWorkloadClassifierResultOutput)
}

type LookupSqlPoolWorkloadClassifierOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// SQL pool name
	SqlPoolName pulumi.StringInput `pulumi:"sqlPoolName"`
	// The name of the workload classifier.
	WorkloadClassifierName pulumi.StringInput `pulumi:"workloadClassifierName"`
	// The name of the workload group.
	WorkloadGroupName pulumi.StringInput `pulumi:"workloadGroupName"`
	// The name of the workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupSqlPoolWorkloadClassifierOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSqlPoolWorkloadClassifierArgs)(nil)).Elem()
}

// Workload classifier operations for a data warehouse
type LookupSqlPoolWorkloadClassifierResultOutput struct{ *pulumi.OutputState }

func (LookupSqlPoolWorkloadClassifierResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSqlPoolWorkloadClassifierResult)(nil)).Elem()
}

func (o LookupSqlPoolWorkloadClassifierResultOutput) ToLookupSqlPoolWorkloadClassifierResultOutput() LookupSqlPoolWorkloadClassifierResultOutput {
	return o
}

func (o LookupSqlPoolWorkloadClassifierResultOutput) ToLookupSqlPoolWorkloadClassifierResultOutputWithContext(ctx context.Context) LookupSqlPoolWorkloadClassifierResultOutput {
	return o
}

func (o LookupSqlPoolWorkloadClassifierResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupSqlPoolWorkloadClassifierResult] {
	return pulumix.Output[LookupSqlPoolWorkloadClassifierResult]{
		OutputState: o.OutputState,
	}
}

// The workload classifier context.
func (o LookupSqlPoolWorkloadClassifierResultOutput) Context() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSqlPoolWorkloadClassifierResult) *string { return v.Context }).(pulumi.StringPtrOutput)
}

// The workload classifier end time for classification.
func (o LookupSqlPoolWorkloadClassifierResultOutput) EndTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSqlPoolWorkloadClassifierResult) *string { return v.EndTime }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupSqlPoolWorkloadClassifierResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlPoolWorkloadClassifierResult) string { return v.Id }).(pulumi.StringOutput)
}

// The workload classifier importance.
func (o LookupSqlPoolWorkloadClassifierResultOutput) Importance() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSqlPoolWorkloadClassifierResult) *string { return v.Importance }).(pulumi.StringPtrOutput)
}

// The workload classifier label.
func (o LookupSqlPoolWorkloadClassifierResultOutput) Label() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSqlPoolWorkloadClassifierResult) *string { return v.Label }).(pulumi.StringPtrOutput)
}

// The workload classifier member name.
func (o LookupSqlPoolWorkloadClassifierResultOutput) MemberName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlPoolWorkloadClassifierResult) string { return v.MemberName }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupSqlPoolWorkloadClassifierResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlPoolWorkloadClassifierResult) string { return v.Name }).(pulumi.StringOutput)
}

// The workload classifier start time for classification.
func (o LookupSqlPoolWorkloadClassifierResultOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSqlPoolWorkloadClassifierResult) *string { return v.StartTime }).(pulumi.StringPtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupSqlPoolWorkloadClassifierResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlPoolWorkloadClassifierResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSqlPoolWorkloadClassifierResultOutput{})
}
