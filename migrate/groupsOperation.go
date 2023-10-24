// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package migrate

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Group resource.
// Azure REST API version: 2023-03-15.
type GroupsOperation struct {
	pulumi.CustomResourceState

	// If the assessments are in running state.
	AreAssessmentsRunning pulumi.BoolOutput `pulumi:"areAssessmentsRunning"`
	// List of References to Assessments created on this group.
	Assessments pulumi.StringArrayOutput `pulumi:"assessments"`
	// Time when this group was created. Date-Time represented in ISO-8601 format.
	CreatedTimestamp pulumi.StringOutput `pulumi:"createdTimestamp"`
	// Whether the group has been created and is valid.
	GroupStatus pulumi.StringOutput `pulumi:"groupStatus"`
	// The type of group.
	GroupType pulumi.StringPtrOutput `pulumi:"groupType"`
	// Number of machines part of this group.
	MachineCount pulumi.IntOutput `pulumi:"machineCount"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The status of the last operation.
	ProvisioningState pulumi.StringPtrOutput `pulumi:"provisioningState"`
	// List of assessment types supported on this group.
	SupportedAssessmentTypes pulumi.StringArrayOutput `pulumi:"supportedAssessmentTypes"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// Time when this group was last updated. Date-Time represented in ISO-8601 format.
	UpdatedTimestamp pulumi.StringOutput `pulumi:"updatedTimestamp"`
}

// NewGroupsOperation registers a new resource with the given unique name, arguments, and options.
func NewGroupsOperation(ctx *pulumi.Context,
	name string, args *GroupsOperationArgs, opts ...pulumi.ResourceOption) (*GroupsOperation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProjectName == nil {
		return nil, errors.New("invalid value for required argument 'ProjectName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:migrate/v20191001:GroupsOperation"),
		},
		{
			Type: pulumi.String("azure-native:migrate/v20230315:GroupsOperation"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource GroupsOperation
	err := ctx.RegisterResource("azure-native:migrate:GroupsOperation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGroupsOperation gets an existing GroupsOperation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGroupsOperation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GroupsOperationState, opts ...pulumi.ResourceOption) (*GroupsOperation, error) {
	var resource GroupsOperation
	err := ctx.ReadResource("azure-native:migrate:GroupsOperation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GroupsOperation resources.
type groupsOperationState struct {
}

type GroupsOperationState struct {
}

func (GroupsOperationState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupsOperationState)(nil)).Elem()
}

type groupsOperationArgs struct {
	// Group ARM name
	GroupName *string `pulumi:"groupName"`
	// The type of group.
	GroupType *string `pulumi:"groupType"`
	// Assessment Project Name
	ProjectName string `pulumi:"projectName"`
	// The status of the last operation.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// List of assessment types supported on this group.
	SupportedAssessmentTypes []string `pulumi:"supportedAssessmentTypes"`
}

// The set of arguments for constructing a GroupsOperation resource.
type GroupsOperationArgs struct {
	// Group ARM name
	GroupName pulumi.StringPtrInput
	// The type of group.
	GroupType pulumi.StringPtrInput
	// Assessment Project Name
	ProjectName pulumi.StringInput
	// The status of the last operation.
	ProvisioningState pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// List of assessment types supported on this group.
	SupportedAssessmentTypes pulumi.StringArrayInput
}

func (GroupsOperationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*groupsOperationArgs)(nil)).Elem()
}

type GroupsOperationInput interface {
	pulumi.Input

	ToGroupsOperationOutput() GroupsOperationOutput
	ToGroupsOperationOutputWithContext(ctx context.Context) GroupsOperationOutput
}

func (*GroupsOperation) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupsOperation)(nil)).Elem()
}

func (i *GroupsOperation) ToGroupsOperationOutput() GroupsOperationOutput {
	return i.ToGroupsOperationOutputWithContext(context.Background())
}

func (i *GroupsOperation) ToGroupsOperationOutputWithContext(ctx context.Context) GroupsOperationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupsOperationOutput)
}

func (i *GroupsOperation) ToOutput(ctx context.Context) pulumix.Output[*GroupsOperation] {
	return pulumix.Output[*GroupsOperation]{
		OutputState: i.ToGroupsOperationOutputWithContext(ctx).OutputState,
	}
}

type GroupsOperationOutput struct{ *pulumi.OutputState }

func (GroupsOperationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupsOperation)(nil)).Elem()
}

func (o GroupsOperationOutput) ToGroupsOperationOutput() GroupsOperationOutput {
	return o
}

func (o GroupsOperationOutput) ToGroupsOperationOutputWithContext(ctx context.Context) GroupsOperationOutput {
	return o
}

func (o GroupsOperationOutput) ToOutput(ctx context.Context) pulumix.Output[*GroupsOperation] {
	return pulumix.Output[*GroupsOperation]{
		OutputState: o.OutputState,
	}
}

// If the assessments are in running state.
func (o GroupsOperationOutput) AreAssessmentsRunning() pulumi.BoolOutput {
	return o.ApplyT(func(v *GroupsOperation) pulumi.BoolOutput { return v.AreAssessmentsRunning }).(pulumi.BoolOutput)
}

// List of References to Assessments created on this group.
func (o GroupsOperationOutput) Assessments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *GroupsOperation) pulumi.StringArrayOutput { return v.Assessments }).(pulumi.StringArrayOutput)
}

// Time when this group was created. Date-Time represented in ISO-8601 format.
func (o GroupsOperationOutput) CreatedTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v *GroupsOperation) pulumi.StringOutput { return v.CreatedTimestamp }).(pulumi.StringOutput)
}

// Whether the group has been created and is valid.
func (o GroupsOperationOutput) GroupStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *GroupsOperation) pulumi.StringOutput { return v.GroupStatus }).(pulumi.StringOutput)
}

// The type of group.
func (o GroupsOperationOutput) GroupType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GroupsOperation) pulumi.StringPtrOutput { return v.GroupType }).(pulumi.StringPtrOutput)
}

// Number of machines part of this group.
func (o GroupsOperationOutput) MachineCount() pulumi.IntOutput {
	return o.ApplyT(func(v *GroupsOperation) pulumi.IntOutput { return v.MachineCount }).(pulumi.IntOutput)
}

// The name of the resource
func (o GroupsOperationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *GroupsOperation) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The status of the last operation.
func (o GroupsOperationOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GroupsOperation) pulumi.StringPtrOutput { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

// List of assessment types supported on this group.
func (o GroupsOperationOutput) SupportedAssessmentTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *GroupsOperation) pulumi.StringArrayOutput { return v.SupportedAssessmentTypes }).(pulumi.StringArrayOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o GroupsOperationOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *GroupsOperation) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o GroupsOperationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *GroupsOperation) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Time when this group was last updated. Date-Time represented in ISO-8601 format.
func (o GroupsOperationOutput) UpdatedTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v *GroupsOperation) pulumi.StringOutput { return v.UpdatedTimestamp }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GroupsOperationOutput{})
}
