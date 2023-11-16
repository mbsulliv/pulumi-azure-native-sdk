// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20221101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a database's short term retention policy.
func LookupBackupShortTermRetentionPolicy(ctx *pulumi.Context, args *LookupBackupShortTermRetentionPolicyArgs, opts ...pulumi.InvokeOption) (*LookupBackupShortTermRetentionPolicyResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupBackupShortTermRetentionPolicyResult
	err := ctx.Invoke("azure-native:sql/v20221101preview:getBackupShortTermRetentionPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBackupShortTermRetentionPolicyArgs struct {
	// The name of the database.
	DatabaseName string `pulumi:"databaseName"`
	// The policy name. Should always be "default".
	PolicyName string `pulumi:"policyName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
}

// A short term retention policy.
type LookupBackupShortTermRetentionPolicyResult struct {
	// The differential backup interval in hours. This is how many interval hours between each differential backup will be supported. This is only applicable to live databases but not dropped databases.
	DiffBackupIntervalInHours *int `pulumi:"diffBackupIntervalInHours"`
	// Resource ID.
	Id string `pulumi:"id"`
	// Resource name.
	Name string `pulumi:"name"`
	// The backup retention period in days. This is how many days Point-in-Time Restore will be supported.
	RetentionDays *int `pulumi:"retentionDays"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupBackupShortTermRetentionPolicyOutput(ctx *pulumi.Context, args LookupBackupShortTermRetentionPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupBackupShortTermRetentionPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBackupShortTermRetentionPolicyResult, error) {
			args := v.(LookupBackupShortTermRetentionPolicyArgs)
			r, err := LookupBackupShortTermRetentionPolicy(ctx, &args, opts...)
			var s LookupBackupShortTermRetentionPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBackupShortTermRetentionPolicyResultOutput)
}

type LookupBackupShortTermRetentionPolicyOutputArgs struct {
	// The name of the database.
	DatabaseName pulumi.StringInput `pulumi:"databaseName"`
	// The policy name. Should always be "default".
	PolicyName pulumi.StringInput `pulumi:"policyName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName pulumi.StringInput `pulumi:"serverName"`
}

func (LookupBackupShortTermRetentionPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBackupShortTermRetentionPolicyArgs)(nil)).Elem()
}

// A short term retention policy.
type LookupBackupShortTermRetentionPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupBackupShortTermRetentionPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBackupShortTermRetentionPolicyResult)(nil)).Elem()
}

func (o LookupBackupShortTermRetentionPolicyResultOutput) ToLookupBackupShortTermRetentionPolicyResultOutput() LookupBackupShortTermRetentionPolicyResultOutput {
	return o
}

func (o LookupBackupShortTermRetentionPolicyResultOutput) ToLookupBackupShortTermRetentionPolicyResultOutputWithContext(ctx context.Context) LookupBackupShortTermRetentionPolicyResultOutput {
	return o
}

// The differential backup interval in hours. This is how many interval hours between each differential backup will be supported. This is only applicable to live databases but not dropped databases.
func (o LookupBackupShortTermRetentionPolicyResultOutput) DiffBackupIntervalInHours() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupBackupShortTermRetentionPolicyResult) *int { return v.DiffBackupIntervalInHours }).(pulumi.IntPtrOutput)
}

// Resource ID.
func (o LookupBackupShortTermRetentionPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackupShortTermRetentionPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource name.
func (o LookupBackupShortTermRetentionPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackupShortTermRetentionPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

// The backup retention period in days. This is how many days Point-in-Time Restore will be supported.
func (o LookupBackupShortTermRetentionPolicyResultOutput) RetentionDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupBackupShortTermRetentionPolicyResult) *int { return v.RetentionDays }).(pulumi.IntPtrOutput)
}

// Resource type.
func (o LookupBackupShortTermRetentionPolicyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackupShortTermRetentionPolicyResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBackupShortTermRetentionPolicyResultOutput{})
}
