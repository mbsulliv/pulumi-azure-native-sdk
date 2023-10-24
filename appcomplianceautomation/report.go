// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package appcomplianceautomation

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// A class represent an AppComplianceAutomation report resource.
// Azure REST API version: 2022-11-16-preview. Prior API version in Azure Native 1.x: 2022-11-16-preview.
type Report struct {
	pulumi.CustomResourceState

	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Report property.
	Properties ReportPropertiesResponseOutput `pulumi:"properties"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewReport registers a new resource with the given unique name, arguments, and options.
func NewReport(ctx *pulumi.Context,
	name string, args *ReportArgs, opts ...pulumi.ResourceOption) (*Report, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:appcomplianceautomation/v20221116preview:Report"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Report
	err := ctx.RegisterResource("azure-native:appcomplianceautomation:Report", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReport gets an existing Report resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReport(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReportState, opts ...pulumi.ResourceOption) (*Report, error) {
	var resource Report
	err := ctx.ReadResource("azure-native:appcomplianceautomation:Report", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Report resources.
type reportState struct {
}

type ReportState struct {
}

func (ReportState) ElementType() reflect.Type {
	return reflect.TypeOf((*reportState)(nil)).Elem()
}

type reportArgs struct {
	// Report property.
	Properties ReportProperties `pulumi:"properties"`
	// Report Name.
	ReportName *string `pulumi:"reportName"`
}

// The set of arguments for constructing a Report resource.
type ReportArgs struct {
	// Report property.
	Properties ReportPropertiesInput
	// Report Name.
	ReportName pulumi.StringPtrInput
}

func (ReportArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*reportArgs)(nil)).Elem()
}

type ReportInput interface {
	pulumi.Input

	ToReportOutput() ReportOutput
	ToReportOutputWithContext(ctx context.Context) ReportOutput
}

func (*Report) ElementType() reflect.Type {
	return reflect.TypeOf((**Report)(nil)).Elem()
}

func (i *Report) ToReportOutput() ReportOutput {
	return i.ToReportOutputWithContext(context.Background())
}

func (i *Report) ToReportOutputWithContext(ctx context.Context) ReportOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportOutput)
}

func (i *Report) ToOutput(ctx context.Context) pulumix.Output[*Report] {
	return pulumix.Output[*Report]{
		OutputState: i.ToReportOutputWithContext(ctx).OutputState,
	}
}

type ReportOutput struct{ *pulumi.OutputState }

func (ReportOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Report)(nil)).Elem()
}

func (o ReportOutput) ToReportOutput() ReportOutput {
	return o
}

func (o ReportOutput) ToReportOutputWithContext(ctx context.Context) ReportOutput {
	return o
}

func (o ReportOutput) ToOutput(ctx context.Context) pulumix.Output[*Report] {
	return pulumix.Output[*Report]{
		OutputState: o.OutputState,
	}
}

// The name of the resource
func (o ReportOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Report) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Report property.
func (o ReportOutput) Properties() ReportPropertiesResponseOutput {
	return o.ApplyT(func(v *Report) ReportPropertiesResponseOutput { return v.Properties }).(ReportPropertiesResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o ReportOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Report) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ReportOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Report) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ReportOutput{})
}
