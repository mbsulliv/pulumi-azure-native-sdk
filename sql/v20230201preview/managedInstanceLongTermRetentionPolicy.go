// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230201preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// A long term retention policy.
type ManagedInstanceLongTermRetentionPolicy struct {
	pulumi.CustomResourceState

	// The monthly retention policy for an LTR backup.
	MonthlyRetention pulumi.StringPtrOutput `pulumi:"monthlyRetention"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// The week of year to take the yearly backup.
	WeekOfYear pulumi.IntPtrOutput `pulumi:"weekOfYear"`
	// The weekly retention policy for an LTR backup.
	WeeklyRetention pulumi.StringPtrOutput `pulumi:"weeklyRetention"`
	// The yearly retention policy for an LTR backup.
	YearlyRetention pulumi.StringPtrOutput `pulumi:"yearlyRetention"`
}

// NewManagedInstanceLongTermRetentionPolicy registers a new resource with the given unique name, arguments, and options.
func NewManagedInstanceLongTermRetentionPolicy(ctx *pulumi.Context,
	name string, args *ManagedInstanceLongTermRetentionPolicyArgs, opts ...pulumi.ResourceOption) (*ManagedInstanceLongTermRetentionPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DatabaseName == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseName'")
	}
	if args.ManagedInstanceName == nil {
		return nil, errors.New("invalid value for required argument 'ManagedInstanceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:sql:ManagedInstanceLongTermRetentionPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220501preview:ManagedInstanceLongTermRetentionPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220801preview:ManagedInstanceLongTermRetentionPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20221101preview:ManagedInstanceLongTermRetentionPolicy"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ManagedInstanceLongTermRetentionPolicy
	err := ctx.RegisterResource("azure-native:sql/v20230201preview:ManagedInstanceLongTermRetentionPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetManagedInstanceLongTermRetentionPolicy gets an existing ManagedInstanceLongTermRetentionPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetManagedInstanceLongTermRetentionPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagedInstanceLongTermRetentionPolicyState, opts ...pulumi.ResourceOption) (*ManagedInstanceLongTermRetentionPolicy, error) {
	var resource ManagedInstanceLongTermRetentionPolicy
	err := ctx.ReadResource("azure-native:sql/v20230201preview:ManagedInstanceLongTermRetentionPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ManagedInstanceLongTermRetentionPolicy resources.
type managedInstanceLongTermRetentionPolicyState struct {
}

type ManagedInstanceLongTermRetentionPolicyState struct {
}

func (ManagedInstanceLongTermRetentionPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*managedInstanceLongTermRetentionPolicyState)(nil)).Elem()
}

type managedInstanceLongTermRetentionPolicyArgs struct {
	// The name of the database.
	DatabaseName string `pulumi:"databaseName"`
	// The name of the managed instance.
	ManagedInstanceName string `pulumi:"managedInstanceName"`
	// The monthly retention policy for an LTR backup.
	MonthlyRetention *string `pulumi:"monthlyRetention"`
	// The policy name. Should always be Default.
	PolicyName *string `pulumi:"policyName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The week of year to take the yearly backup.
	WeekOfYear *int `pulumi:"weekOfYear"`
	// The weekly retention policy for an LTR backup.
	WeeklyRetention *string `pulumi:"weeklyRetention"`
	// The yearly retention policy for an LTR backup.
	YearlyRetention *string `pulumi:"yearlyRetention"`
}

// The set of arguments for constructing a ManagedInstanceLongTermRetentionPolicy resource.
type ManagedInstanceLongTermRetentionPolicyArgs struct {
	// The name of the database.
	DatabaseName pulumi.StringInput
	// The name of the managed instance.
	ManagedInstanceName pulumi.StringInput
	// The monthly retention policy for an LTR backup.
	MonthlyRetention pulumi.StringPtrInput
	// The policy name. Should always be Default.
	PolicyName pulumi.StringPtrInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// The week of year to take the yearly backup.
	WeekOfYear pulumi.IntPtrInput
	// The weekly retention policy for an LTR backup.
	WeeklyRetention pulumi.StringPtrInput
	// The yearly retention policy for an LTR backup.
	YearlyRetention pulumi.StringPtrInput
}

func (ManagedInstanceLongTermRetentionPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managedInstanceLongTermRetentionPolicyArgs)(nil)).Elem()
}

type ManagedInstanceLongTermRetentionPolicyInput interface {
	pulumi.Input

	ToManagedInstanceLongTermRetentionPolicyOutput() ManagedInstanceLongTermRetentionPolicyOutput
	ToManagedInstanceLongTermRetentionPolicyOutputWithContext(ctx context.Context) ManagedInstanceLongTermRetentionPolicyOutput
}

func (*ManagedInstanceLongTermRetentionPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedInstanceLongTermRetentionPolicy)(nil)).Elem()
}

func (i *ManagedInstanceLongTermRetentionPolicy) ToManagedInstanceLongTermRetentionPolicyOutput() ManagedInstanceLongTermRetentionPolicyOutput {
	return i.ToManagedInstanceLongTermRetentionPolicyOutputWithContext(context.Background())
}

func (i *ManagedInstanceLongTermRetentionPolicy) ToManagedInstanceLongTermRetentionPolicyOutputWithContext(ctx context.Context) ManagedInstanceLongTermRetentionPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedInstanceLongTermRetentionPolicyOutput)
}

func (i *ManagedInstanceLongTermRetentionPolicy) ToOutput(ctx context.Context) pulumix.Output[*ManagedInstanceLongTermRetentionPolicy] {
	return pulumix.Output[*ManagedInstanceLongTermRetentionPolicy]{
		OutputState: i.ToManagedInstanceLongTermRetentionPolicyOutputWithContext(ctx).OutputState,
	}
}

type ManagedInstanceLongTermRetentionPolicyOutput struct{ *pulumi.OutputState }

func (ManagedInstanceLongTermRetentionPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedInstanceLongTermRetentionPolicy)(nil)).Elem()
}

func (o ManagedInstanceLongTermRetentionPolicyOutput) ToManagedInstanceLongTermRetentionPolicyOutput() ManagedInstanceLongTermRetentionPolicyOutput {
	return o
}

func (o ManagedInstanceLongTermRetentionPolicyOutput) ToManagedInstanceLongTermRetentionPolicyOutputWithContext(ctx context.Context) ManagedInstanceLongTermRetentionPolicyOutput {
	return o
}

func (o ManagedInstanceLongTermRetentionPolicyOutput) ToOutput(ctx context.Context) pulumix.Output[*ManagedInstanceLongTermRetentionPolicy] {
	return pulumix.Output[*ManagedInstanceLongTermRetentionPolicy]{
		OutputState: o.OutputState,
	}
}

// The monthly retention policy for an LTR backup.
func (o ManagedInstanceLongTermRetentionPolicyOutput) MonthlyRetention() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedInstanceLongTermRetentionPolicy) pulumi.StringPtrOutput { return v.MonthlyRetention }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o ManagedInstanceLongTermRetentionPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedInstanceLongTermRetentionPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Resource type.
func (o ManagedInstanceLongTermRetentionPolicyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedInstanceLongTermRetentionPolicy) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The week of year to take the yearly backup.
func (o ManagedInstanceLongTermRetentionPolicyOutput) WeekOfYear() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ManagedInstanceLongTermRetentionPolicy) pulumi.IntPtrOutput { return v.WeekOfYear }).(pulumi.IntPtrOutput)
}

// The weekly retention policy for an LTR backup.
func (o ManagedInstanceLongTermRetentionPolicyOutput) WeeklyRetention() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedInstanceLongTermRetentionPolicy) pulumi.StringPtrOutput { return v.WeeklyRetention }).(pulumi.StringPtrOutput)
}

// The yearly retention policy for an LTR backup.
func (o ManagedInstanceLongTermRetentionPolicyOutput) YearlyRetention() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedInstanceLongTermRetentionPolicy) pulumi.StringPtrOutput { return v.YearlyRetention }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ManagedInstanceLongTermRetentionPolicyOutput{})
}
