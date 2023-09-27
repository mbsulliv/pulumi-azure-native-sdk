// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20180202

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

	// Time when this project was created. Date-Time represented in ISO-8601 format.
	CreatedTimestamp pulumi.StringOutput `pulumi:"createdTimestamp"`
	// ARM ID of the Service Map workspace created by user.
	CustomerWorkspaceId pulumi.StringPtrOutput `pulumi:"customerWorkspaceId"`
	// Location of the Service Map workspace created by user.
	CustomerWorkspaceLocation pulumi.StringPtrOutput `pulumi:"customerWorkspaceLocation"`
	// Reports whether project is under discovery.
	DiscoveryStatus pulumi.StringOutput `pulumi:"discoveryStatus"`
	// For optimistic concurrency control.
	ETag pulumi.StringPtrOutput `pulumi:"eTag"`
	// Time when last assessment was created. Date-Time represented in ISO-8601 format. This value will be null until assessment is created.
	LastAssessmentTimestamp pulumi.StringOutput `pulumi:"lastAssessmentTimestamp"`
	// Session id of the last discovery.
	LastDiscoverySessionId pulumi.StringOutput `pulumi:"lastDiscoverySessionId"`
	// Time when this project was created. Date-Time represented in ISO-8601 format. This value will be null until discovery is complete.
	LastDiscoveryTimestamp pulumi.StringOutput `pulumi:"lastDiscoveryTimestamp"`
	// Azure location in which project is created.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Name of the project.
	Name pulumi.StringOutput `pulumi:"name"`
	// Number of assessments created in the project.
	NumberOfAssessments pulumi.IntOutput `pulumi:"numberOfAssessments"`
	// Number of groups created in the project.
	NumberOfGroups pulumi.IntOutput `pulumi:"numberOfGroups"`
	// Number of machines in the project.
	NumberOfMachines pulumi.IntOutput `pulumi:"numberOfMachines"`
	// Provisioning state of the project.
	ProvisioningState pulumi.StringPtrOutput `pulumi:"provisioningState"`
	// Tags provided by Azure Tagging service.
	Tags pulumi.AnyOutput `pulumi:"tags"`
	// Type of the object = [Microsoft.Migrate/projects].
	Type pulumi.StringOutput `pulumi:"type"`
	// Time when this project was last updated. Date-Time represented in ISO-8601 format.
	UpdatedTimestamp pulumi.StringOutput `pulumi:"updatedTimestamp"`
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
			Type: pulumi.String("azure-native:migrate/v20171111preview:Project"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Project
	err := ctx.RegisterResource("azure-native:migrate/v20180202:Project", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure-native:migrate/v20180202:Project", name, id, state, &resource, opts...)
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
	// ARM ID of the Service Map workspace created by user.
	CustomerWorkspaceId *string `pulumi:"customerWorkspaceId"`
	// Location of the Service Map workspace created by user.
	CustomerWorkspaceLocation *string `pulumi:"customerWorkspaceLocation"`
	// For optimistic concurrency control.
	ETag *string `pulumi:"eTag"`
	// Azure location in which project is created.
	Location *string `pulumi:"location"`
	// Name of the Azure Migrate project.
	ProjectName *string `pulumi:"projectName"`
	// Provisioning state of the project.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Name of the Azure Resource Group that project is part of.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Tags provided by Azure Tagging service.
	Tags interface{} `pulumi:"tags"`
}

// The set of arguments for constructing a Project resource.
type ProjectArgs struct {
	// ARM ID of the Service Map workspace created by user.
	CustomerWorkspaceId pulumi.StringPtrInput
	// Location of the Service Map workspace created by user.
	CustomerWorkspaceLocation pulumi.StringPtrInput
	// For optimistic concurrency control.
	ETag pulumi.StringPtrInput
	// Azure location in which project is created.
	Location pulumi.StringPtrInput
	// Name of the Azure Migrate project.
	ProjectName pulumi.StringPtrInput
	// Provisioning state of the project.
	ProvisioningState pulumi.StringPtrInput
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

// Time when this project was created. Date-Time represented in ISO-8601 format.
func (o ProjectOutput) CreatedTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v *Project) pulumi.StringOutput { return v.CreatedTimestamp }).(pulumi.StringOutput)
}

// ARM ID of the Service Map workspace created by user.
func (o ProjectOutput) CustomerWorkspaceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Project) pulumi.StringPtrOutput { return v.CustomerWorkspaceId }).(pulumi.StringPtrOutput)
}

// Location of the Service Map workspace created by user.
func (o ProjectOutput) CustomerWorkspaceLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Project) pulumi.StringPtrOutput { return v.CustomerWorkspaceLocation }).(pulumi.StringPtrOutput)
}

// Reports whether project is under discovery.
func (o ProjectOutput) DiscoveryStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *Project) pulumi.StringOutput { return v.DiscoveryStatus }).(pulumi.StringOutput)
}

// For optimistic concurrency control.
func (o ProjectOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Project) pulumi.StringPtrOutput { return v.ETag }).(pulumi.StringPtrOutput)
}

// Time when last assessment was created. Date-Time represented in ISO-8601 format. This value will be null until assessment is created.
func (o ProjectOutput) LastAssessmentTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v *Project) pulumi.StringOutput { return v.LastAssessmentTimestamp }).(pulumi.StringOutput)
}

// Session id of the last discovery.
func (o ProjectOutput) LastDiscoverySessionId() pulumi.StringOutput {
	return o.ApplyT(func(v *Project) pulumi.StringOutput { return v.LastDiscoverySessionId }).(pulumi.StringOutput)
}

// Time when this project was created. Date-Time represented in ISO-8601 format. This value will be null until discovery is complete.
func (o ProjectOutput) LastDiscoveryTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v *Project) pulumi.StringOutput { return v.LastDiscoveryTimestamp }).(pulumi.StringOutput)
}

// Azure location in which project is created.
func (o ProjectOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Project) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Name of the project.
func (o ProjectOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Project) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Number of assessments created in the project.
func (o ProjectOutput) NumberOfAssessments() pulumi.IntOutput {
	return o.ApplyT(func(v *Project) pulumi.IntOutput { return v.NumberOfAssessments }).(pulumi.IntOutput)
}

// Number of groups created in the project.
func (o ProjectOutput) NumberOfGroups() pulumi.IntOutput {
	return o.ApplyT(func(v *Project) pulumi.IntOutput { return v.NumberOfGroups }).(pulumi.IntOutput)
}

// Number of machines in the project.
func (o ProjectOutput) NumberOfMachines() pulumi.IntOutput {
	return o.ApplyT(func(v *Project) pulumi.IntOutput { return v.NumberOfMachines }).(pulumi.IntOutput)
}

// Provisioning state of the project.
func (o ProjectOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Project) pulumi.StringPtrOutput { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

// Tags provided by Azure Tagging service.
func (o ProjectOutput) Tags() pulumi.AnyOutput {
	return o.ApplyT(func(v *Project) pulumi.AnyOutput { return v.Tags }).(pulumi.AnyOutput)
}

// Type of the object = [Microsoft.Migrate/projects].
func (o ProjectOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Project) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Time when this project was last updated. Date-Time represented in ISO-8601 format.
func (o ProjectOutput) UpdatedTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v *Project) pulumi.StringOutput { return v.UpdatedTimestamp }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ProjectOutput{})
}
