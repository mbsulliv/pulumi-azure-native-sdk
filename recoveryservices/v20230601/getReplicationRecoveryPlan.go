// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets the details of the recovery plan.
func LookupReplicationRecoveryPlan(ctx *pulumi.Context, args *LookupReplicationRecoveryPlanArgs, opts ...pulumi.InvokeOption) (*LookupReplicationRecoveryPlanResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupReplicationRecoveryPlanResult
	err := ctx.Invoke("azure-native:recoveryservices/v20230601:getReplicationRecoveryPlan", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupReplicationRecoveryPlanArgs struct {
	// Name of the recovery plan.
	RecoveryPlanName string `pulumi:"recoveryPlanName"`
	// The name of the resource group where the recovery services vault is present.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the recovery services vault.
	ResourceName string `pulumi:"resourceName"`
}

// Recovery plan details.
type LookupReplicationRecoveryPlanResult struct {
	// Resource Id
	Id string `pulumi:"id"`
	// Resource Location
	Location *string `pulumi:"location"`
	// Resource Name
	Name string `pulumi:"name"`
	// The custom details.
	Properties RecoveryPlanPropertiesResponse `pulumi:"properties"`
	// Resource Type
	Type string `pulumi:"type"`
}

func LookupReplicationRecoveryPlanOutput(ctx *pulumi.Context, args LookupReplicationRecoveryPlanOutputArgs, opts ...pulumi.InvokeOption) LookupReplicationRecoveryPlanResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupReplicationRecoveryPlanResult, error) {
			args := v.(LookupReplicationRecoveryPlanArgs)
			r, err := LookupReplicationRecoveryPlan(ctx, &args, opts...)
			var s LookupReplicationRecoveryPlanResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupReplicationRecoveryPlanResultOutput)
}

type LookupReplicationRecoveryPlanOutputArgs struct {
	// Name of the recovery plan.
	RecoveryPlanName pulumi.StringInput `pulumi:"recoveryPlanName"`
	// The name of the resource group where the recovery services vault is present.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the recovery services vault.
	ResourceName pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupReplicationRecoveryPlanOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReplicationRecoveryPlanArgs)(nil)).Elem()
}

// Recovery plan details.
type LookupReplicationRecoveryPlanResultOutput struct{ *pulumi.OutputState }

func (LookupReplicationRecoveryPlanResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReplicationRecoveryPlanResult)(nil)).Elem()
}

func (o LookupReplicationRecoveryPlanResultOutput) ToLookupReplicationRecoveryPlanResultOutput() LookupReplicationRecoveryPlanResultOutput {
	return o
}

func (o LookupReplicationRecoveryPlanResultOutput) ToLookupReplicationRecoveryPlanResultOutputWithContext(ctx context.Context) LookupReplicationRecoveryPlanResultOutput {
	return o
}

func (o LookupReplicationRecoveryPlanResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupReplicationRecoveryPlanResult] {
	return pulumix.Output[LookupReplicationRecoveryPlanResult]{
		OutputState: o.OutputState,
	}
}

// Resource Id
func (o LookupReplicationRecoveryPlanResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReplicationRecoveryPlanResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource Location
func (o LookupReplicationRecoveryPlanResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupReplicationRecoveryPlanResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource Name
func (o LookupReplicationRecoveryPlanResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReplicationRecoveryPlanResult) string { return v.Name }).(pulumi.StringOutput)
}

// The custom details.
func (o LookupReplicationRecoveryPlanResultOutput) Properties() RecoveryPlanPropertiesResponseOutput {
	return o.ApplyT(func(v LookupReplicationRecoveryPlanResult) RecoveryPlanPropertiesResponse { return v.Properties }).(RecoveryPlanPropertiesResponseOutput)
}

// Resource Type
func (o LookupReplicationRecoveryPlanResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReplicationRecoveryPlanResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupReplicationRecoveryPlanResultOutput{})
}
