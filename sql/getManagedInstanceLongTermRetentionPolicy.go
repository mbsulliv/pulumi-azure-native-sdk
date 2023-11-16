// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sql

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a managed database's long term retention policy.
// Azure REST API version: 2022-11-01-preview.
//
// Other available API versions: 2023-02-01-preview, 2023-05-01-preview.
func LookupManagedInstanceLongTermRetentionPolicy(ctx *pulumi.Context, args *LookupManagedInstanceLongTermRetentionPolicyArgs, opts ...pulumi.InvokeOption) (*LookupManagedInstanceLongTermRetentionPolicyResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupManagedInstanceLongTermRetentionPolicyResult
	err := ctx.Invoke("azure-native:sql:getManagedInstanceLongTermRetentionPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupManagedInstanceLongTermRetentionPolicyArgs struct {
	// The name of the database.
	DatabaseName string `pulumi:"databaseName"`
	// The name of the managed instance.
	ManagedInstanceName string `pulumi:"managedInstanceName"`
	// The policy name. Should always be Default.
	PolicyName string `pulumi:"policyName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A long term retention policy.
type LookupManagedInstanceLongTermRetentionPolicyResult struct {
	// Resource ID.
	Id string `pulumi:"id"`
	// The monthly retention policy for an LTR backup in an ISO 8601 format.
	MonthlyRetention *string `pulumi:"monthlyRetention"`
	// Resource name.
	Name string `pulumi:"name"`
	// Resource type.
	Type string `pulumi:"type"`
	// The week of year to take the yearly backup in an ISO 8601 format.
	WeekOfYear *int `pulumi:"weekOfYear"`
	// The weekly retention policy for an LTR backup in an ISO 8601 format.
	WeeklyRetention *string `pulumi:"weeklyRetention"`
	// The yearly retention policy for an LTR backup in an ISO 8601 format.
	YearlyRetention *string `pulumi:"yearlyRetention"`
}

func LookupManagedInstanceLongTermRetentionPolicyOutput(ctx *pulumi.Context, args LookupManagedInstanceLongTermRetentionPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupManagedInstanceLongTermRetentionPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupManagedInstanceLongTermRetentionPolicyResult, error) {
			args := v.(LookupManagedInstanceLongTermRetentionPolicyArgs)
			r, err := LookupManagedInstanceLongTermRetentionPolicy(ctx, &args, opts...)
			var s LookupManagedInstanceLongTermRetentionPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupManagedInstanceLongTermRetentionPolicyResultOutput)
}

type LookupManagedInstanceLongTermRetentionPolicyOutputArgs struct {
	// The name of the database.
	DatabaseName pulumi.StringInput `pulumi:"databaseName"`
	// The name of the managed instance.
	ManagedInstanceName pulumi.StringInput `pulumi:"managedInstanceName"`
	// The policy name. Should always be Default.
	PolicyName pulumi.StringInput `pulumi:"policyName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupManagedInstanceLongTermRetentionPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagedInstanceLongTermRetentionPolicyArgs)(nil)).Elem()
}

// A long term retention policy.
type LookupManagedInstanceLongTermRetentionPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupManagedInstanceLongTermRetentionPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagedInstanceLongTermRetentionPolicyResult)(nil)).Elem()
}

func (o LookupManagedInstanceLongTermRetentionPolicyResultOutput) ToLookupManagedInstanceLongTermRetentionPolicyResultOutput() LookupManagedInstanceLongTermRetentionPolicyResultOutput {
	return o
}

func (o LookupManagedInstanceLongTermRetentionPolicyResultOutput) ToLookupManagedInstanceLongTermRetentionPolicyResultOutputWithContext(ctx context.Context) LookupManagedInstanceLongTermRetentionPolicyResultOutput {
	return o
}

// Resource ID.
func (o LookupManagedInstanceLongTermRetentionPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedInstanceLongTermRetentionPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

// The monthly retention policy for an LTR backup in an ISO 8601 format.
func (o LookupManagedInstanceLongTermRetentionPolicyResultOutput) MonthlyRetention() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedInstanceLongTermRetentionPolicyResult) *string { return v.MonthlyRetention }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o LookupManagedInstanceLongTermRetentionPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedInstanceLongTermRetentionPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

// Resource type.
func (o LookupManagedInstanceLongTermRetentionPolicyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedInstanceLongTermRetentionPolicyResult) string { return v.Type }).(pulumi.StringOutput)
}

// The week of year to take the yearly backup in an ISO 8601 format.
func (o LookupManagedInstanceLongTermRetentionPolicyResultOutput) WeekOfYear() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupManagedInstanceLongTermRetentionPolicyResult) *int { return v.WeekOfYear }).(pulumi.IntPtrOutput)
}

// The weekly retention policy for an LTR backup in an ISO 8601 format.
func (o LookupManagedInstanceLongTermRetentionPolicyResultOutput) WeeklyRetention() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedInstanceLongTermRetentionPolicyResult) *string { return v.WeeklyRetention }).(pulumi.StringPtrOutput)
}

// The yearly retention policy for an LTR backup in an ISO 8601 format.
func (o LookupManagedInstanceLongTermRetentionPolicyResultOutput) YearlyRetention() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedInstanceLongTermRetentionPolicyResult) *string { return v.YearlyRetention }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupManagedInstanceLongTermRetentionPolicyResultOutput{})
}
