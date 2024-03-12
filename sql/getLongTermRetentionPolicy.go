// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sql

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a database's long term retention policy.
// Azure REST API version: 2021-11-01.
//
// Other available API versions: 2022-11-01-preview, 2023-02-01-preview, 2023-05-01-preview, 2023-08-01-preview.
func LookupLongTermRetentionPolicy(ctx *pulumi.Context, args *LookupLongTermRetentionPolicyArgs, opts ...pulumi.InvokeOption) (*LookupLongTermRetentionPolicyResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupLongTermRetentionPolicyResult
	err := ctx.Invoke("azure-native:sql:getLongTermRetentionPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLongTermRetentionPolicyArgs struct {
	// The name of the database.
	DatabaseName string `pulumi:"databaseName"`
	// The policy name. Should always be Default.
	PolicyName string `pulumi:"policyName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
}

// A long term retention policy.
type LookupLongTermRetentionPolicyResult struct {
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

func LookupLongTermRetentionPolicyOutput(ctx *pulumi.Context, args LookupLongTermRetentionPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupLongTermRetentionPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupLongTermRetentionPolicyResult, error) {
			args := v.(LookupLongTermRetentionPolicyArgs)
			r, err := LookupLongTermRetentionPolicy(ctx, &args, opts...)
			var s LookupLongTermRetentionPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupLongTermRetentionPolicyResultOutput)
}

type LookupLongTermRetentionPolicyOutputArgs struct {
	// The name of the database.
	DatabaseName pulumi.StringInput `pulumi:"databaseName"`
	// The policy name. Should always be Default.
	PolicyName pulumi.StringInput `pulumi:"policyName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName pulumi.StringInput `pulumi:"serverName"`
}

func (LookupLongTermRetentionPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLongTermRetentionPolicyArgs)(nil)).Elem()
}

// A long term retention policy.
type LookupLongTermRetentionPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupLongTermRetentionPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLongTermRetentionPolicyResult)(nil)).Elem()
}

func (o LookupLongTermRetentionPolicyResultOutput) ToLookupLongTermRetentionPolicyResultOutput() LookupLongTermRetentionPolicyResultOutput {
	return o
}

func (o LookupLongTermRetentionPolicyResultOutput) ToLookupLongTermRetentionPolicyResultOutputWithContext(ctx context.Context) LookupLongTermRetentionPolicyResultOutput {
	return o
}

// Resource ID.
func (o LookupLongTermRetentionPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLongTermRetentionPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

// The monthly retention policy for an LTR backup in an ISO 8601 format.
func (o LookupLongTermRetentionPolicyResultOutput) MonthlyRetention() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLongTermRetentionPolicyResult) *string { return v.MonthlyRetention }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o LookupLongTermRetentionPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLongTermRetentionPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

// Resource type.
func (o LookupLongTermRetentionPolicyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLongTermRetentionPolicyResult) string { return v.Type }).(pulumi.StringOutput)
}

// The week of year to take the yearly backup in an ISO 8601 format.
func (o LookupLongTermRetentionPolicyResultOutput) WeekOfYear() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupLongTermRetentionPolicyResult) *int { return v.WeekOfYear }).(pulumi.IntPtrOutput)
}

// The weekly retention policy for an LTR backup in an ISO 8601 format.
func (o LookupLongTermRetentionPolicyResultOutput) WeeklyRetention() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLongTermRetentionPolicyResult) *string { return v.WeeklyRetention }).(pulumi.StringPtrOutput)
}

// The yearly retention policy for an LTR backup in an ISO 8601 format.
func (o LookupLongTermRetentionPolicyResultOutput) YearlyRetention() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLongTermRetentionPolicyResult) *string { return v.YearlyRetention }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLongTermRetentionPolicyResultOutput{})
}
