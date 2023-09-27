// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sql

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// A long term retention policy.
// Azure REST API version: 2017-03-01-preview.
type BackupLongTermRetentionPolicy struct {
	pulumi.CustomResourceState

	// The monthly retention policy for an LTR backup in an ISO 8601 format.
	MonthlyRetention pulumi.StringPtrOutput `pulumi:"monthlyRetention"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// The week of year to take the yearly backup in an ISO 8601 format.
	WeekOfYear pulumi.IntPtrOutput `pulumi:"weekOfYear"`
	// The weekly retention policy for an LTR backup in an ISO 8601 format.
	WeeklyRetention pulumi.StringPtrOutput `pulumi:"weeklyRetention"`
	// The yearly retention policy for an LTR backup in an ISO 8601 format.
	YearlyRetention pulumi.StringPtrOutput `pulumi:"yearlyRetention"`
}

// NewBackupLongTermRetentionPolicy registers a new resource with the given unique name, arguments, and options.
func NewBackupLongTermRetentionPolicy(ctx *pulumi.Context,
	name string, args *BackupLongTermRetentionPolicyArgs, opts ...pulumi.ResourceOption) (*BackupLongTermRetentionPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DatabaseName == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServerName == nil {
		return nil, errors.New("invalid value for required argument 'ServerName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:sql/v20170301preview:BackupLongTermRetentionPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200202preview:BackupLongTermRetentionPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200801preview:BackupLongTermRetentionPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20201101preview:BackupLongTermRetentionPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210201preview:BackupLongTermRetentionPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210501preview:BackupLongTermRetentionPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210801preview:BackupLongTermRetentionPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101:BackupLongTermRetentionPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101preview:BackupLongTermRetentionPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220201preview:BackupLongTermRetentionPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220501preview:BackupLongTermRetentionPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220801preview:BackupLongTermRetentionPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20221101preview:BackupLongTermRetentionPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20230201preview:BackupLongTermRetentionPolicy"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource BackupLongTermRetentionPolicy
	err := ctx.RegisterResource("azure-native:sql:BackupLongTermRetentionPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBackupLongTermRetentionPolicy gets an existing BackupLongTermRetentionPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBackupLongTermRetentionPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BackupLongTermRetentionPolicyState, opts ...pulumi.ResourceOption) (*BackupLongTermRetentionPolicy, error) {
	var resource BackupLongTermRetentionPolicy
	err := ctx.ReadResource("azure-native:sql:BackupLongTermRetentionPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BackupLongTermRetentionPolicy resources.
type backupLongTermRetentionPolicyState struct {
}

type BackupLongTermRetentionPolicyState struct {
}

func (BackupLongTermRetentionPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*backupLongTermRetentionPolicyState)(nil)).Elem()
}

type backupLongTermRetentionPolicyArgs struct {
	// The name of the database.
	DatabaseName string `pulumi:"databaseName"`
	// The monthly retention policy for an LTR backup in an ISO 8601 format.
	MonthlyRetention *string `pulumi:"monthlyRetention"`
	// The policy name. Should always be Default.
	PolicyName *string `pulumi:"policyName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
	// The week of year to take the yearly backup in an ISO 8601 format.
	WeekOfYear *int `pulumi:"weekOfYear"`
	// The weekly retention policy for an LTR backup in an ISO 8601 format.
	WeeklyRetention *string `pulumi:"weeklyRetention"`
	// The yearly retention policy for an LTR backup in an ISO 8601 format.
	YearlyRetention *string `pulumi:"yearlyRetention"`
}

// The set of arguments for constructing a BackupLongTermRetentionPolicy resource.
type BackupLongTermRetentionPolicyArgs struct {
	// The name of the database.
	DatabaseName pulumi.StringInput
	// The monthly retention policy for an LTR backup in an ISO 8601 format.
	MonthlyRetention pulumi.StringPtrInput
	// The policy name. Should always be Default.
	PolicyName pulumi.StringPtrInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// The name of the server.
	ServerName pulumi.StringInput
	// The week of year to take the yearly backup in an ISO 8601 format.
	WeekOfYear pulumi.IntPtrInput
	// The weekly retention policy for an LTR backup in an ISO 8601 format.
	WeeklyRetention pulumi.StringPtrInput
	// The yearly retention policy for an LTR backup in an ISO 8601 format.
	YearlyRetention pulumi.StringPtrInput
}

func (BackupLongTermRetentionPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*backupLongTermRetentionPolicyArgs)(nil)).Elem()
}

type BackupLongTermRetentionPolicyInput interface {
	pulumi.Input

	ToBackupLongTermRetentionPolicyOutput() BackupLongTermRetentionPolicyOutput
	ToBackupLongTermRetentionPolicyOutputWithContext(ctx context.Context) BackupLongTermRetentionPolicyOutput
}

func (*BackupLongTermRetentionPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**BackupLongTermRetentionPolicy)(nil)).Elem()
}

func (i *BackupLongTermRetentionPolicy) ToBackupLongTermRetentionPolicyOutput() BackupLongTermRetentionPolicyOutput {
	return i.ToBackupLongTermRetentionPolicyOutputWithContext(context.Background())
}

func (i *BackupLongTermRetentionPolicy) ToBackupLongTermRetentionPolicyOutputWithContext(ctx context.Context) BackupLongTermRetentionPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupLongTermRetentionPolicyOutput)
}

func (i *BackupLongTermRetentionPolicy) ToOutput(ctx context.Context) pulumix.Output[*BackupLongTermRetentionPolicy] {
	return pulumix.Output[*BackupLongTermRetentionPolicy]{
		OutputState: i.ToBackupLongTermRetentionPolicyOutputWithContext(ctx).OutputState,
	}
}

type BackupLongTermRetentionPolicyOutput struct{ *pulumi.OutputState }

func (BackupLongTermRetentionPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BackupLongTermRetentionPolicy)(nil)).Elem()
}

func (o BackupLongTermRetentionPolicyOutput) ToBackupLongTermRetentionPolicyOutput() BackupLongTermRetentionPolicyOutput {
	return o
}

func (o BackupLongTermRetentionPolicyOutput) ToBackupLongTermRetentionPolicyOutputWithContext(ctx context.Context) BackupLongTermRetentionPolicyOutput {
	return o
}

func (o BackupLongTermRetentionPolicyOutput) ToOutput(ctx context.Context) pulumix.Output[*BackupLongTermRetentionPolicy] {
	return pulumix.Output[*BackupLongTermRetentionPolicy]{
		OutputState: o.OutputState,
	}
}

// The monthly retention policy for an LTR backup in an ISO 8601 format.
func (o BackupLongTermRetentionPolicyOutput) MonthlyRetention() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackupLongTermRetentionPolicy) pulumi.StringPtrOutput { return v.MonthlyRetention }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o BackupLongTermRetentionPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupLongTermRetentionPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Resource type.
func (o BackupLongTermRetentionPolicyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupLongTermRetentionPolicy) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The week of year to take the yearly backup in an ISO 8601 format.
func (o BackupLongTermRetentionPolicyOutput) WeekOfYear() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *BackupLongTermRetentionPolicy) pulumi.IntPtrOutput { return v.WeekOfYear }).(pulumi.IntPtrOutput)
}

// The weekly retention policy for an LTR backup in an ISO 8601 format.
func (o BackupLongTermRetentionPolicyOutput) WeeklyRetention() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackupLongTermRetentionPolicy) pulumi.StringPtrOutput { return v.WeeklyRetention }).(pulumi.StringPtrOutput)
}

// The yearly retention policy for an LTR backup in an ISO 8601 format.
func (o BackupLongTermRetentionPolicyOutput) YearlyRetention() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackupLongTermRetentionPolicy) pulumi.StringPtrOutput { return v.YearlyRetention }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(BackupLongTermRetentionPolicyOutput{})
}
