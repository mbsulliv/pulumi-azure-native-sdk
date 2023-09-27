// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230501

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Snapshot policy information
type SnapshotPolicy struct {
	pulumi.CustomResourceState

	// Schedule for daily snapshots
	DailySchedule DailyScheduleResponsePtrOutput `pulumi:"dailySchedule"`
	// The property to decide policy is enabled or not
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Schedule for hourly snapshots
	HourlySchedule HourlyScheduleResponsePtrOutput `pulumi:"hourlySchedule"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// Schedule for monthly snapshots
	MonthlySchedule MonthlyScheduleResponsePtrOutput `pulumi:"monthlySchedule"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Azure lifecycle management
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// Schedule for weekly snapshots
	WeeklySchedule WeeklyScheduleResponsePtrOutput `pulumi:"weeklySchedule"`
}

// NewSnapshotPolicy registers a new resource with the given unique name, arguments, and options.
func NewSnapshotPolicy(ctx *pulumi.Context,
	name string, args *SnapshotPolicyArgs, opts ...pulumi.ResourceOption) (*SnapshotPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:netapp:SnapshotPolicy"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20200501:SnapshotPolicy"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20200601:SnapshotPolicy"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20200701:SnapshotPolicy"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20200801:SnapshotPolicy"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20200901:SnapshotPolicy"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20201101:SnapshotPolicy"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20201201:SnapshotPolicy"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20210201:SnapshotPolicy"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20210401:SnapshotPolicy"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20210401preview:SnapshotPolicy"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20210601:SnapshotPolicy"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20210801:SnapshotPolicy"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20211001:SnapshotPolicy"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20220101:SnapshotPolicy"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20220301:SnapshotPolicy"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20220501:SnapshotPolicy"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20220901:SnapshotPolicy"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20221101:SnapshotPolicy"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20221101preview:SnapshotPolicy"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource SnapshotPolicy
	err := ctx.RegisterResource("azure-native:netapp/v20230501:SnapshotPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSnapshotPolicy gets an existing SnapshotPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSnapshotPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SnapshotPolicyState, opts ...pulumi.ResourceOption) (*SnapshotPolicy, error) {
	var resource SnapshotPolicy
	err := ctx.ReadResource("azure-native:netapp/v20230501:SnapshotPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SnapshotPolicy resources.
type snapshotPolicyState struct {
}

type SnapshotPolicyState struct {
}

func (SnapshotPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*snapshotPolicyState)(nil)).Elem()
}

type snapshotPolicyArgs struct {
	// The name of the NetApp account
	AccountName string `pulumi:"accountName"`
	// Schedule for daily snapshots
	DailySchedule *DailySchedule `pulumi:"dailySchedule"`
	// The property to decide policy is enabled or not
	Enabled *bool `pulumi:"enabled"`
	// Schedule for hourly snapshots
	HourlySchedule *HourlySchedule `pulumi:"hourlySchedule"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// Schedule for monthly snapshots
	MonthlySchedule *MonthlySchedule `pulumi:"monthlySchedule"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the snapshot policy
	SnapshotPolicyName *string `pulumi:"snapshotPolicyName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Schedule for weekly snapshots
	WeeklySchedule *WeeklySchedule `pulumi:"weeklySchedule"`
}

// The set of arguments for constructing a SnapshotPolicy resource.
type SnapshotPolicyArgs struct {
	// The name of the NetApp account
	AccountName pulumi.StringInput
	// Schedule for daily snapshots
	DailySchedule DailySchedulePtrInput
	// The property to decide policy is enabled or not
	Enabled pulumi.BoolPtrInput
	// Schedule for hourly snapshots
	HourlySchedule HourlySchedulePtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// Schedule for monthly snapshots
	MonthlySchedule MonthlySchedulePtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the snapshot policy
	SnapshotPolicyName pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Schedule for weekly snapshots
	WeeklySchedule WeeklySchedulePtrInput
}

func (SnapshotPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*snapshotPolicyArgs)(nil)).Elem()
}

type SnapshotPolicyInput interface {
	pulumi.Input

	ToSnapshotPolicyOutput() SnapshotPolicyOutput
	ToSnapshotPolicyOutputWithContext(ctx context.Context) SnapshotPolicyOutput
}

func (*SnapshotPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**SnapshotPolicy)(nil)).Elem()
}

func (i *SnapshotPolicy) ToSnapshotPolicyOutput() SnapshotPolicyOutput {
	return i.ToSnapshotPolicyOutputWithContext(context.Background())
}

func (i *SnapshotPolicy) ToSnapshotPolicyOutputWithContext(ctx context.Context) SnapshotPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotPolicyOutput)
}

func (i *SnapshotPolicy) ToOutput(ctx context.Context) pulumix.Output[*SnapshotPolicy] {
	return pulumix.Output[*SnapshotPolicy]{
		OutputState: i.ToSnapshotPolicyOutputWithContext(ctx).OutputState,
	}
}

type SnapshotPolicyOutput struct{ *pulumi.OutputState }

func (SnapshotPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SnapshotPolicy)(nil)).Elem()
}

func (o SnapshotPolicyOutput) ToSnapshotPolicyOutput() SnapshotPolicyOutput {
	return o
}

func (o SnapshotPolicyOutput) ToSnapshotPolicyOutputWithContext(ctx context.Context) SnapshotPolicyOutput {
	return o
}

func (o SnapshotPolicyOutput) ToOutput(ctx context.Context) pulumix.Output[*SnapshotPolicy] {
	return pulumix.Output[*SnapshotPolicy]{
		OutputState: o.OutputState,
	}
}

// Schedule for daily snapshots
func (o SnapshotPolicyOutput) DailySchedule() DailyScheduleResponsePtrOutput {
	return o.ApplyT(func(v *SnapshotPolicy) DailyScheduleResponsePtrOutput { return v.DailySchedule }).(DailyScheduleResponsePtrOutput)
}

// The property to decide policy is enabled or not
func (o SnapshotPolicyOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SnapshotPolicy) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o SnapshotPolicyOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *SnapshotPolicy) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Schedule for hourly snapshots
func (o SnapshotPolicyOutput) HourlySchedule() HourlyScheduleResponsePtrOutput {
	return o.ApplyT(func(v *SnapshotPolicy) HourlyScheduleResponsePtrOutput { return v.HourlySchedule }).(HourlyScheduleResponsePtrOutput)
}

// The geo-location where the resource lives
func (o SnapshotPolicyOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *SnapshotPolicy) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Schedule for monthly snapshots
func (o SnapshotPolicyOutput) MonthlySchedule() MonthlyScheduleResponsePtrOutput {
	return o.ApplyT(func(v *SnapshotPolicy) MonthlyScheduleResponsePtrOutput { return v.MonthlySchedule }).(MonthlyScheduleResponsePtrOutput)
}

// The name of the resource
func (o SnapshotPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SnapshotPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Azure lifecycle management
func (o SnapshotPolicyOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *SnapshotPolicy) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o SnapshotPolicyOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *SnapshotPolicy) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o SnapshotPolicyOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SnapshotPolicy) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o SnapshotPolicyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SnapshotPolicy) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Schedule for weekly snapshots
func (o SnapshotPolicyOutput) WeeklySchedule() WeeklyScheduleResponsePtrOutput {
	return o.ApplyT(func(v *SnapshotPolicy) WeeklyScheduleResponsePtrOutput { return v.WeeklySchedule }).(WeeklyScheduleResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(SnapshotPolicyOutput{})
}
