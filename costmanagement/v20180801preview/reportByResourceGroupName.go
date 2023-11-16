// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20180801preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A report resource.
type ReportByResourceGroupName struct {
	pulumi.CustomResourceState

	// Has definition for the report.
	Definition ReportDefinitionResponseOutput `pulumi:"definition"`
	// Has delivery information for the report.
	DeliveryInfo ReportDeliveryInfoResponseOutput `pulumi:"deliveryInfo"`
	// The format of the report being delivered.
	Format pulumi.StringPtrOutput `pulumi:"format"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Has schedule information for the report.
	Schedule ReportScheduleResponsePtrOutput `pulumi:"schedule"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewReportByResourceGroupName registers a new resource with the given unique name, arguments, and options.
func NewReportByResourceGroupName(ctx *pulumi.Context,
	name string, args *ReportByResourceGroupNameArgs, opts ...pulumi.ResourceOption) (*ReportByResourceGroupName, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Definition == nil {
		return nil, errors.New("invalid value for required argument 'Definition'")
	}
	if args.DeliveryInfo == nil {
		return nil, errors.New("invalid value for required argument 'DeliveryInfo'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:costmanagement:ReportByResourceGroupName"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ReportByResourceGroupName
	err := ctx.RegisterResource("azure-native:costmanagement/v20180801preview:ReportByResourceGroupName", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReportByResourceGroupName gets an existing ReportByResourceGroupName resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReportByResourceGroupName(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReportByResourceGroupNameState, opts ...pulumi.ResourceOption) (*ReportByResourceGroupName, error) {
	var resource ReportByResourceGroupName
	err := ctx.ReadResource("azure-native:costmanagement/v20180801preview:ReportByResourceGroupName", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ReportByResourceGroupName resources.
type reportByResourceGroupNameState struct {
}

type ReportByResourceGroupNameState struct {
}

func (ReportByResourceGroupNameState) ElementType() reflect.Type {
	return reflect.TypeOf((*reportByResourceGroupNameState)(nil)).Elem()
}

type reportByResourceGroupNameArgs struct {
	// Has definition for the report.
	Definition ReportDefinition `pulumi:"definition"`
	// Has delivery information for the report.
	DeliveryInfo ReportDeliveryInfo `pulumi:"deliveryInfo"`
	// The format of the report being delivered.
	Format *string `pulumi:"format"`
	// Report Name.
	ReportName *string `pulumi:"reportName"`
	// Azure Resource Group Name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Has schedule information for the report.
	Schedule *ReportSchedule `pulumi:"schedule"`
}

// The set of arguments for constructing a ReportByResourceGroupName resource.
type ReportByResourceGroupNameArgs struct {
	// Has definition for the report.
	Definition ReportDefinitionInput
	// Has delivery information for the report.
	DeliveryInfo ReportDeliveryInfoInput
	// The format of the report being delivered.
	Format pulumi.StringPtrInput
	// Report Name.
	ReportName pulumi.StringPtrInput
	// Azure Resource Group Name.
	ResourceGroupName pulumi.StringInput
	// Has schedule information for the report.
	Schedule ReportSchedulePtrInput
}

func (ReportByResourceGroupNameArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*reportByResourceGroupNameArgs)(nil)).Elem()
}

type ReportByResourceGroupNameInput interface {
	pulumi.Input

	ToReportByResourceGroupNameOutput() ReportByResourceGroupNameOutput
	ToReportByResourceGroupNameOutputWithContext(ctx context.Context) ReportByResourceGroupNameOutput
}

func (*ReportByResourceGroupName) ElementType() reflect.Type {
	return reflect.TypeOf((**ReportByResourceGroupName)(nil)).Elem()
}

func (i *ReportByResourceGroupName) ToReportByResourceGroupNameOutput() ReportByResourceGroupNameOutput {
	return i.ToReportByResourceGroupNameOutputWithContext(context.Background())
}

func (i *ReportByResourceGroupName) ToReportByResourceGroupNameOutputWithContext(ctx context.Context) ReportByResourceGroupNameOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportByResourceGroupNameOutput)
}

type ReportByResourceGroupNameOutput struct{ *pulumi.OutputState }

func (ReportByResourceGroupNameOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReportByResourceGroupName)(nil)).Elem()
}

func (o ReportByResourceGroupNameOutput) ToReportByResourceGroupNameOutput() ReportByResourceGroupNameOutput {
	return o
}

func (o ReportByResourceGroupNameOutput) ToReportByResourceGroupNameOutputWithContext(ctx context.Context) ReportByResourceGroupNameOutput {
	return o
}

// Has definition for the report.
func (o ReportByResourceGroupNameOutput) Definition() ReportDefinitionResponseOutput {
	return o.ApplyT(func(v *ReportByResourceGroupName) ReportDefinitionResponseOutput { return v.Definition }).(ReportDefinitionResponseOutput)
}

// Has delivery information for the report.
func (o ReportByResourceGroupNameOutput) DeliveryInfo() ReportDeliveryInfoResponseOutput {
	return o.ApplyT(func(v *ReportByResourceGroupName) ReportDeliveryInfoResponseOutput { return v.DeliveryInfo }).(ReportDeliveryInfoResponseOutput)
}

// The format of the report being delivered.
func (o ReportByResourceGroupNameOutput) Format() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReportByResourceGroupName) pulumi.StringPtrOutput { return v.Format }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o ReportByResourceGroupNameOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ReportByResourceGroupName) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Has schedule information for the report.
func (o ReportByResourceGroupNameOutput) Schedule() ReportScheduleResponsePtrOutput {
	return o.ApplyT(func(v *ReportByResourceGroupName) ReportScheduleResponsePtrOutput { return v.Schedule }).(ReportScheduleResponsePtrOutput)
}

// Resource tags.
func (o ReportByResourceGroupNameOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ReportByResourceGroupName) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o ReportByResourceGroupNameOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ReportByResourceGroupName) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ReportByResourceGroupNameOutput{})
}
