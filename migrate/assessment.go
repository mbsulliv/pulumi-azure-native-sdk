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

// An assessment created for a group in the Migration project.
// Azure REST API version: 2019-10-01. Prior API version in Azure Native 1.x: 2019-10-01
type Assessment struct {
	pulumi.CustomResourceState

	// For optimistic concurrency control.
	ETag pulumi.StringPtrOutput `pulumi:"eTag"`
	// Unique name of an assessment.
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties of the assessment.
	Properties AssessmentPropertiesResponseOutput `pulumi:"properties"`
	// Type of the object = [Microsoft.Migrate/assessmentProjects/groups/assessments].
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewAssessment registers a new resource with the given unique name, arguments, and options.
func NewAssessment(ctx *pulumi.Context,
	name string, args *AssessmentArgs, opts ...pulumi.ResourceOption) (*Assessment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GroupName == nil {
		return nil, errors.New("invalid value for required argument 'GroupName'")
	}
	if args.ProjectName == nil {
		return nil, errors.New("invalid value for required argument 'ProjectName'")
	}
	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:migrate/v20191001:Assessment"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Assessment
	err := ctx.RegisterResource("azure-native:migrate:Assessment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAssessment gets an existing Assessment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAssessment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AssessmentState, opts ...pulumi.ResourceOption) (*Assessment, error) {
	var resource Assessment
	err := ctx.ReadResource("azure-native:migrate:Assessment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Assessment resources.
type assessmentState struct {
}

type AssessmentState struct {
}

func (AssessmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*assessmentState)(nil)).Elem()
}

type assessmentArgs struct {
	// Unique name of an assessment within a project.
	AssessmentName *string `pulumi:"assessmentName"`
	// For optimistic concurrency control.
	ETag *string `pulumi:"eTag"`
	// Unique name of a group within a project.
	GroupName string `pulumi:"groupName"`
	// Name of the Azure Migrate project.
	ProjectName string `pulumi:"projectName"`
	// Properties of the assessment.
	Properties AssessmentProperties `pulumi:"properties"`
	// Name of the Azure Resource Group that project is part of.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a Assessment resource.
type AssessmentArgs struct {
	// Unique name of an assessment within a project.
	AssessmentName pulumi.StringPtrInput
	// For optimistic concurrency control.
	ETag pulumi.StringPtrInput
	// Unique name of a group within a project.
	GroupName pulumi.StringInput
	// Name of the Azure Migrate project.
	ProjectName pulumi.StringInput
	// Properties of the assessment.
	Properties AssessmentPropertiesInput
	// Name of the Azure Resource Group that project is part of.
	ResourceGroupName pulumi.StringInput
}

func (AssessmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*assessmentArgs)(nil)).Elem()
}

type AssessmentInput interface {
	pulumi.Input

	ToAssessmentOutput() AssessmentOutput
	ToAssessmentOutputWithContext(ctx context.Context) AssessmentOutput
}

func (*Assessment) ElementType() reflect.Type {
	return reflect.TypeOf((**Assessment)(nil)).Elem()
}

func (i *Assessment) ToAssessmentOutput() AssessmentOutput {
	return i.ToAssessmentOutputWithContext(context.Background())
}

func (i *Assessment) ToAssessmentOutputWithContext(ctx context.Context) AssessmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssessmentOutput)
}

func (i *Assessment) ToOutput(ctx context.Context) pulumix.Output[*Assessment] {
	return pulumix.Output[*Assessment]{
		OutputState: i.ToAssessmentOutputWithContext(ctx).OutputState,
	}
}

type AssessmentOutput struct{ *pulumi.OutputState }

func (AssessmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Assessment)(nil)).Elem()
}

func (o AssessmentOutput) ToAssessmentOutput() AssessmentOutput {
	return o
}

func (o AssessmentOutput) ToAssessmentOutputWithContext(ctx context.Context) AssessmentOutput {
	return o
}

func (o AssessmentOutput) ToOutput(ctx context.Context) pulumix.Output[*Assessment] {
	return pulumix.Output[*Assessment]{
		OutputState: o.OutputState,
	}
}

// For optimistic concurrency control.
func (o AssessmentOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Assessment) pulumi.StringPtrOutput { return v.ETag }).(pulumi.StringPtrOutput)
}

// Unique name of an assessment.
func (o AssessmentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Assessment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Properties of the assessment.
func (o AssessmentOutput) Properties() AssessmentPropertiesResponseOutput {
	return o.ApplyT(func(v *Assessment) AssessmentPropertiesResponseOutput { return v.Properties }).(AssessmentPropertiesResponseOutput)
}

// Type of the object = [Microsoft.Migrate/assessmentProjects/groups/assessments].
func (o AssessmentOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Assessment) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AssessmentOutput{})
}
