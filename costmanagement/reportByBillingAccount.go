// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package costmanagement

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// A report resource.
// Azure REST API version: 2018-08-01-preview. Prior API version in Azure Native 1.x: 2018-08-01-preview
type ReportByBillingAccount struct {
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

// NewReportByBillingAccount registers a new resource with the given unique name, arguments, and options.
func NewReportByBillingAccount(ctx *pulumi.Context,
	name string, args *ReportByBillingAccountArgs, opts ...pulumi.ResourceOption) (*ReportByBillingAccount, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BillingAccountId == nil {
		return nil, errors.New("invalid value for required argument 'BillingAccountId'")
	}
	if args.Definition == nil {
		return nil, errors.New("invalid value for required argument 'Definition'")
	}
	if args.DeliveryInfo == nil {
		return nil, errors.New("invalid value for required argument 'DeliveryInfo'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:costmanagement/v20180801preview:ReportByBillingAccount"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ReportByBillingAccount
	err := ctx.RegisterResource("azure-native:costmanagement:ReportByBillingAccount", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReportByBillingAccount gets an existing ReportByBillingAccount resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReportByBillingAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReportByBillingAccountState, opts ...pulumi.ResourceOption) (*ReportByBillingAccount, error) {
	var resource ReportByBillingAccount
	err := ctx.ReadResource("azure-native:costmanagement:ReportByBillingAccount", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ReportByBillingAccount resources.
type reportByBillingAccountState struct {
}

type ReportByBillingAccountState struct {
}

func (ReportByBillingAccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*reportByBillingAccountState)(nil)).Elem()
}

type reportByBillingAccountArgs struct {
	// BillingAccount ID
	BillingAccountId string `pulumi:"billingAccountId"`
	// Has definition for the report.
	Definition ReportDefinition `pulumi:"definition"`
	// Has delivery information for the report.
	DeliveryInfo ReportDeliveryInfo `pulumi:"deliveryInfo"`
	// The format of the report being delivered.
	Format *string `pulumi:"format"`
	// Report Name.
	ReportName *string `pulumi:"reportName"`
	// Has schedule information for the report.
	Schedule *ReportSchedule `pulumi:"schedule"`
}

// The set of arguments for constructing a ReportByBillingAccount resource.
type ReportByBillingAccountArgs struct {
	// BillingAccount ID
	BillingAccountId pulumi.StringInput
	// Has definition for the report.
	Definition ReportDefinitionInput
	// Has delivery information for the report.
	DeliveryInfo ReportDeliveryInfoInput
	// The format of the report being delivered.
	Format pulumi.StringPtrInput
	// Report Name.
	ReportName pulumi.StringPtrInput
	// Has schedule information for the report.
	Schedule ReportSchedulePtrInput
}

func (ReportByBillingAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*reportByBillingAccountArgs)(nil)).Elem()
}

type ReportByBillingAccountInput interface {
	pulumi.Input

	ToReportByBillingAccountOutput() ReportByBillingAccountOutput
	ToReportByBillingAccountOutputWithContext(ctx context.Context) ReportByBillingAccountOutput
}

func (*ReportByBillingAccount) ElementType() reflect.Type {
	return reflect.TypeOf((**ReportByBillingAccount)(nil)).Elem()
}

func (i *ReportByBillingAccount) ToReportByBillingAccountOutput() ReportByBillingAccountOutput {
	return i.ToReportByBillingAccountOutputWithContext(context.Background())
}

func (i *ReportByBillingAccount) ToReportByBillingAccountOutputWithContext(ctx context.Context) ReportByBillingAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportByBillingAccountOutput)
}

func (i *ReportByBillingAccount) ToOutput(ctx context.Context) pulumix.Output[*ReportByBillingAccount] {
	return pulumix.Output[*ReportByBillingAccount]{
		OutputState: i.ToReportByBillingAccountOutputWithContext(ctx).OutputState,
	}
}

type ReportByBillingAccountOutput struct{ *pulumi.OutputState }

func (ReportByBillingAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReportByBillingAccount)(nil)).Elem()
}

func (o ReportByBillingAccountOutput) ToReportByBillingAccountOutput() ReportByBillingAccountOutput {
	return o
}

func (o ReportByBillingAccountOutput) ToReportByBillingAccountOutputWithContext(ctx context.Context) ReportByBillingAccountOutput {
	return o
}

func (o ReportByBillingAccountOutput) ToOutput(ctx context.Context) pulumix.Output[*ReportByBillingAccount] {
	return pulumix.Output[*ReportByBillingAccount]{
		OutputState: o.OutputState,
	}
}

// Has definition for the report.
func (o ReportByBillingAccountOutput) Definition() ReportDefinitionResponseOutput {
	return o.ApplyT(func(v *ReportByBillingAccount) ReportDefinitionResponseOutput { return v.Definition }).(ReportDefinitionResponseOutput)
}

// Has delivery information for the report.
func (o ReportByBillingAccountOutput) DeliveryInfo() ReportDeliveryInfoResponseOutput {
	return o.ApplyT(func(v *ReportByBillingAccount) ReportDeliveryInfoResponseOutput { return v.DeliveryInfo }).(ReportDeliveryInfoResponseOutput)
}

// The format of the report being delivered.
func (o ReportByBillingAccountOutput) Format() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReportByBillingAccount) pulumi.StringPtrOutput { return v.Format }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o ReportByBillingAccountOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ReportByBillingAccount) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Has schedule information for the report.
func (o ReportByBillingAccountOutput) Schedule() ReportScheduleResponsePtrOutput {
	return o.ApplyT(func(v *ReportByBillingAccount) ReportScheduleResponsePtrOutput { return v.Schedule }).(ReportScheduleResponsePtrOutput)
}

// Resource tags.
func (o ReportByBillingAccountOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ReportByBillingAccount) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o ReportByBillingAccountOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ReportByBillingAccount) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ReportByBillingAccountOutput{})
}
