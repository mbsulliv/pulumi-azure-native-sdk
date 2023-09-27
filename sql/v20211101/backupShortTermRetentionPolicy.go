// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20211101

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// A short term retention policy.
type BackupShortTermRetentionPolicy struct {
	pulumi.CustomResourceState

	// The differential backup interval in hours. This is how many interval hours between each differential backup will be supported. This is only applicable to live databases but not dropped databases.
	DiffBackupIntervalInHours pulumi.IntPtrOutput `pulumi:"diffBackupIntervalInHours"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The backup retention period in days. This is how many days Point-in-Time Restore will be supported.
	RetentionDays pulumi.IntPtrOutput `pulumi:"retentionDays"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewBackupShortTermRetentionPolicy registers a new resource with the given unique name, arguments, and options.
func NewBackupShortTermRetentionPolicy(ctx *pulumi.Context,
	name string, args *BackupShortTermRetentionPolicyArgs, opts ...pulumi.ResourceOption) (*BackupShortTermRetentionPolicy, error) {
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
			Type: pulumi.String("azure-native:sql:BackupShortTermRetentionPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20171001preview:BackupShortTermRetentionPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200202preview:BackupShortTermRetentionPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200801preview:BackupShortTermRetentionPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20201101preview:BackupShortTermRetentionPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210201preview:BackupShortTermRetentionPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210501preview:BackupShortTermRetentionPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210801preview:BackupShortTermRetentionPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101preview:BackupShortTermRetentionPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220201preview:BackupShortTermRetentionPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220501preview:BackupShortTermRetentionPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220801preview:BackupShortTermRetentionPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20221101preview:BackupShortTermRetentionPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20230201preview:BackupShortTermRetentionPolicy"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource BackupShortTermRetentionPolicy
	err := ctx.RegisterResource("azure-native:sql/v20211101:BackupShortTermRetentionPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBackupShortTermRetentionPolicy gets an existing BackupShortTermRetentionPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBackupShortTermRetentionPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BackupShortTermRetentionPolicyState, opts ...pulumi.ResourceOption) (*BackupShortTermRetentionPolicy, error) {
	var resource BackupShortTermRetentionPolicy
	err := ctx.ReadResource("azure-native:sql/v20211101:BackupShortTermRetentionPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BackupShortTermRetentionPolicy resources.
type backupShortTermRetentionPolicyState struct {
}

type BackupShortTermRetentionPolicyState struct {
}

func (BackupShortTermRetentionPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*backupShortTermRetentionPolicyState)(nil)).Elem()
}

type backupShortTermRetentionPolicyArgs struct {
	// The name of the database.
	DatabaseName string `pulumi:"databaseName"`
	// The differential backup interval in hours. This is how many interval hours between each differential backup will be supported. This is only applicable to live databases but not dropped databases.
	DiffBackupIntervalInHours *int `pulumi:"diffBackupIntervalInHours"`
	// The policy name. Should always be "default".
	PolicyName *string `pulumi:"policyName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The backup retention period in days. This is how many days Point-in-Time Restore will be supported.
	RetentionDays *int `pulumi:"retentionDays"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
}

// The set of arguments for constructing a BackupShortTermRetentionPolicy resource.
type BackupShortTermRetentionPolicyArgs struct {
	// The name of the database.
	DatabaseName pulumi.StringInput
	// The differential backup interval in hours. This is how many interval hours between each differential backup will be supported. This is only applicable to live databases but not dropped databases.
	DiffBackupIntervalInHours pulumi.IntPtrInput
	// The policy name. Should always be "default".
	PolicyName pulumi.StringPtrInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// The backup retention period in days. This is how many days Point-in-Time Restore will be supported.
	RetentionDays pulumi.IntPtrInput
	// The name of the server.
	ServerName pulumi.StringInput
}

func (BackupShortTermRetentionPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*backupShortTermRetentionPolicyArgs)(nil)).Elem()
}

type BackupShortTermRetentionPolicyInput interface {
	pulumi.Input

	ToBackupShortTermRetentionPolicyOutput() BackupShortTermRetentionPolicyOutput
	ToBackupShortTermRetentionPolicyOutputWithContext(ctx context.Context) BackupShortTermRetentionPolicyOutput
}

func (*BackupShortTermRetentionPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**BackupShortTermRetentionPolicy)(nil)).Elem()
}

func (i *BackupShortTermRetentionPolicy) ToBackupShortTermRetentionPolicyOutput() BackupShortTermRetentionPolicyOutput {
	return i.ToBackupShortTermRetentionPolicyOutputWithContext(context.Background())
}

func (i *BackupShortTermRetentionPolicy) ToBackupShortTermRetentionPolicyOutputWithContext(ctx context.Context) BackupShortTermRetentionPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupShortTermRetentionPolicyOutput)
}

func (i *BackupShortTermRetentionPolicy) ToOutput(ctx context.Context) pulumix.Output[*BackupShortTermRetentionPolicy] {
	return pulumix.Output[*BackupShortTermRetentionPolicy]{
		OutputState: i.ToBackupShortTermRetentionPolicyOutputWithContext(ctx).OutputState,
	}
}

type BackupShortTermRetentionPolicyOutput struct{ *pulumi.OutputState }

func (BackupShortTermRetentionPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BackupShortTermRetentionPolicy)(nil)).Elem()
}

func (o BackupShortTermRetentionPolicyOutput) ToBackupShortTermRetentionPolicyOutput() BackupShortTermRetentionPolicyOutput {
	return o
}

func (o BackupShortTermRetentionPolicyOutput) ToBackupShortTermRetentionPolicyOutputWithContext(ctx context.Context) BackupShortTermRetentionPolicyOutput {
	return o
}

func (o BackupShortTermRetentionPolicyOutput) ToOutput(ctx context.Context) pulumix.Output[*BackupShortTermRetentionPolicy] {
	return pulumix.Output[*BackupShortTermRetentionPolicy]{
		OutputState: o.OutputState,
	}
}

// The differential backup interval in hours. This is how many interval hours between each differential backup will be supported. This is only applicable to live databases but not dropped databases.
func (o BackupShortTermRetentionPolicyOutput) DiffBackupIntervalInHours() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *BackupShortTermRetentionPolicy) pulumi.IntPtrOutput { return v.DiffBackupIntervalInHours }).(pulumi.IntPtrOutput)
}

// Resource name.
func (o BackupShortTermRetentionPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupShortTermRetentionPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The backup retention period in days. This is how many days Point-in-Time Restore will be supported.
func (o BackupShortTermRetentionPolicyOutput) RetentionDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *BackupShortTermRetentionPolicy) pulumi.IntPtrOutput { return v.RetentionDays }).(pulumi.IntPtrOutput)
}

// Resource type.
func (o BackupShortTermRetentionPolicyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupShortTermRetentionPolicy) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(BackupShortTermRetentionPolicyOutput{})
}
