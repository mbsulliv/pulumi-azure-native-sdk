// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20191001

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Azure Migrate Project.
type Project struct {
	pulumi.CustomResourceState

	// For optimistic concurrency control.
	ETag pulumi.StringPtrOutput `pulumi:"eTag"`
	// Azure location in which project is created.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Name of the project.
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties of the project.
	Properties ProjectPropertiesResponseOutput `pulumi:"properties"`
	// Tags provided by Azure Tagging service.
	Tags pulumi.AnyOutput `pulumi:"tags"`
	// Type of the object = [Microsoft.Migrate/assessmentProjects].
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewProject registers a new resource with the given unique name, arguments, and options.
func NewProject(ctx *pulumi.Context,
	name string, args *ProjectArgs, opts ...pulumi.ResourceOption) (*Project, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:migrate:Project"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Project
	err := ctx.RegisterResource("azure-native:migrate/v20191001:Project", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProject gets an existing Project resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProject(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectState, opts ...pulumi.ResourceOption) (*Project, error) {
	var resource Project
	err := ctx.ReadResource("azure-native:migrate/v20191001:Project", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Project resources.
type projectState struct {
}

type ProjectState struct {
}

func (ProjectState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectState)(nil)).Elem()
}

type projectArgs struct {
	// For optimistic concurrency control.
	ETag *string `pulumi:"eTag"`
	// Azure location in which project is created.
	Location *string `pulumi:"location"`
	// Name of the Azure Migrate project.
	ProjectName *string `pulumi:"projectName"`
	// Properties of the project.
	Properties *ProjectProperties `pulumi:"properties"`
	// Name of the Azure Resource Group that project is part of.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Tags provided by Azure Tagging service.
	Tags interface{} `pulumi:"tags"`
}

// The set of arguments for constructing a Project resource.
type ProjectArgs struct {
	// For optimistic concurrency control.
	ETag pulumi.StringPtrInput
	// Azure location in which project is created.
	Location pulumi.StringPtrInput
	// Name of the Azure Migrate project.
	ProjectName pulumi.StringPtrInput
	// Properties of the project.
	Properties ProjectPropertiesPtrInput
	// Name of the Azure Resource Group that project is part of.
	ResourceGroupName pulumi.StringInput
	// Tags provided by Azure Tagging service.
	Tags pulumi.Input
}

func (ProjectArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectArgs)(nil)).Elem()
}

type ProjectInput interface {
	pulumi.Input

	ToProjectOutput() ProjectOutput
	ToProjectOutputWithContext(ctx context.Context) ProjectOutput
}

func (*Project) ElementType() reflect.Type {
	return reflect.TypeOf((**Project)(nil)).Elem()
}

func (i *Project) ToProjectOutput() ProjectOutput {
	return i.ToProjectOutputWithContext(context.Background())
}

func (i *Project) ToProjectOutputWithContext(ctx context.Context) ProjectOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectOutput)
}

func (i *Project) ToOutput(ctx context.Context) pulumix.Output[*Project] {
	return pulumix.Output[*Project]{
		OutputState: i.ToProjectOutputWithContext(ctx).OutputState,
	}
}

type ProjectOutput struct{ *pulumi.OutputState }

func (ProjectOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Project)(nil)).Elem()
}

func (o ProjectOutput) ToProjectOutput() ProjectOutput {
	return o
}

func (o ProjectOutput) ToProjectOutputWithContext(ctx context.Context) ProjectOutput {
	return o
}

func (o ProjectOutput) ToOutput(ctx context.Context) pulumix.Output[*Project] {
	return pulumix.Output[*Project]{
		OutputState: o.OutputState,
	}
}

// For optimistic concurrency control.
func (o ProjectOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Project) pulumi.StringPtrOutput { return v.ETag }).(pulumi.StringPtrOutput)
}

// Azure location in which project is created.
func (o ProjectOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Project) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Name of the project.
func (o ProjectOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Project) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Properties of the project.
func (o ProjectOutput) Properties() ProjectPropertiesResponseOutput {
	return o.ApplyT(func(v *Project) ProjectPropertiesResponseOutput { return v.Properties }).(ProjectPropertiesResponseOutput)
}

// Tags provided by Azure Tagging service.
func (o ProjectOutput) Tags() pulumi.AnyOutput {
	return o.ApplyT(func(v *Project) pulumi.AnyOutput { return v.Tags }).(pulumi.AnyOutput)
}

// Type of the object = [Microsoft.Migrate/assessmentProjects].
func (o ProjectOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Project) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ProjectOutput{})
}
