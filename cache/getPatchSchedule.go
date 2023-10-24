// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cache

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets the patching schedule of a redis cache.
// Azure REST API version: 2023-04-01.
//
// Other available API versions: 2023-05-01-preview, 2023-08-01.
func LookupPatchSchedule(ctx *pulumi.Context, args *LookupPatchScheduleArgs, opts ...pulumi.InvokeOption) (*LookupPatchScheduleResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupPatchScheduleResult
	err := ctx.Invoke("azure-native:cache:getPatchSchedule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPatchScheduleArgs struct {
	// Default string modeled as parameter for auto generation to work correctly.
	Default string `pulumi:"default"`
	// The name of the redis cache.
	Name string `pulumi:"name"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Response to put/get patch schedules for Redis cache.
type LookupPatchScheduleResult struct {
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// List of patch schedules for a Redis cache.
	ScheduleEntries []ScheduleEntryResponse `pulumi:"scheduleEntries"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupPatchScheduleOutput(ctx *pulumi.Context, args LookupPatchScheduleOutputArgs, opts ...pulumi.InvokeOption) LookupPatchScheduleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPatchScheduleResult, error) {
			args := v.(LookupPatchScheduleArgs)
			r, err := LookupPatchSchedule(ctx, &args, opts...)
			var s LookupPatchScheduleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPatchScheduleResultOutput)
}

type LookupPatchScheduleOutputArgs struct {
	// Default string modeled as parameter for auto generation to work correctly.
	Default pulumi.StringInput `pulumi:"default"`
	// The name of the redis cache.
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupPatchScheduleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPatchScheduleArgs)(nil)).Elem()
}

// Response to put/get patch schedules for Redis cache.
type LookupPatchScheduleResultOutput struct{ *pulumi.OutputState }

func (LookupPatchScheduleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPatchScheduleResult)(nil)).Elem()
}

func (o LookupPatchScheduleResultOutput) ToLookupPatchScheduleResultOutput() LookupPatchScheduleResultOutput {
	return o
}

func (o LookupPatchScheduleResultOutput) ToLookupPatchScheduleResultOutputWithContext(ctx context.Context) LookupPatchScheduleResultOutput {
	return o
}

func (o LookupPatchScheduleResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupPatchScheduleResult] {
	return pulumix.Output[LookupPatchScheduleResult]{
		OutputState: o.OutputState,
	}
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupPatchScheduleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPatchScheduleResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupPatchScheduleResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPatchScheduleResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupPatchScheduleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPatchScheduleResult) string { return v.Name }).(pulumi.StringOutput)
}

// List of patch schedules for a Redis cache.
func (o LookupPatchScheduleResultOutput) ScheduleEntries() ScheduleEntryResponseArrayOutput {
	return o.ApplyT(func(v LookupPatchScheduleResult) []ScheduleEntryResponse { return v.ScheduleEntries }).(ScheduleEntryResponseArrayOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupPatchScheduleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPatchScheduleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPatchScheduleResultOutput{})
}
