// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20151101preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The container for solution.
type Solution struct {
	pulumi.CustomResourceState

	// Resource location
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Plan for solution object supported by the OperationsManagement resource provider.
	Plan SolutionPlanResponsePtrOutput `pulumi:"plan"`
	// Properties for solution object supported by the OperationsManagement resource provider.
	Properties SolutionPropertiesResponseOutput `pulumi:"properties"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewSolution registers a new resource with the given unique name, arguments, and options.
func NewSolution(ctx *pulumi.Context,
	name string, args *SolutionArgs, opts ...pulumi.ResourceOption) (*Solution, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:operationsmanagement:Solution"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Solution
	err := ctx.RegisterResource("azure-native:operationsmanagement/v20151101preview:Solution", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSolution gets an existing Solution resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSolution(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SolutionState, opts ...pulumi.ResourceOption) (*Solution, error) {
	var resource Solution
	err := ctx.ReadResource("azure-native:operationsmanagement/v20151101preview:Solution", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Solution resources.
type solutionState struct {
}

type SolutionState struct {
}

func (SolutionState) ElementType() reflect.Type {
	return reflect.TypeOf((*solutionState)(nil)).Elem()
}

type solutionArgs struct {
	// Resource location
	Location *string `pulumi:"location"`
	// Plan for solution object supported by the OperationsManagement resource provider.
	Plan *SolutionPlan `pulumi:"plan"`
	// Properties for solution object supported by the OperationsManagement resource provider.
	Properties *SolutionProperties `pulumi:"properties"`
	// The name of the resource group to get. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// User Solution Name.
	SolutionName *string `pulumi:"solutionName"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Solution resource.
type SolutionArgs struct {
	// Resource location
	Location pulumi.StringPtrInput
	// Plan for solution object supported by the OperationsManagement resource provider.
	Plan SolutionPlanPtrInput
	// Properties for solution object supported by the OperationsManagement resource provider.
	Properties SolutionPropertiesPtrInput
	// The name of the resource group to get. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// User Solution Name.
	SolutionName pulumi.StringPtrInput
	// Resource tags
	Tags pulumi.StringMapInput
}

func (SolutionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*solutionArgs)(nil)).Elem()
}

type SolutionInput interface {
	pulumi.Input

	ToSolutionOutput() SolutionOutput
	ToSolutionOutputWithContext(ctx context.Context) SolutionOutput
}

func (*Solution) ElementType() reflect.Type {
	return reflect.TypeOf((**Solution)(nil)).Elem()
}

func (i *Solution) ToSolutionOutput() SolutionOutput {
	return i.ToSolutionOutputWithContext(context.Background())
}

func (i *Solution) ToSolutionOutputWithContext(ctx context.Context) SolutionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SolutionOutput)
}

func (i *Solution) ToOutput(ctx context.Context) pulumix.Output[*Solution] {
	return pulumix.Output[*Solution]{
		OutputState: i.ToSolutionOutputWithContext(ctx).OutputState,
	}
}

type SolutionOutput struct{ *pulumi.OutputState }

func (SolutionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Solution)(nil)).Elem()
}

func (o SolutionOutput) ToSolutionOutput() SolutionOutput {
	return o
}

func (o SolutionOutput) ToSolutionOutputWithContext(ctx context.Context) SolutionOutput {
	return o
}

func (o SolutionOutput) ToOutput(ctx context.Context) pulumix.Output[*Solution] {
	return pulumix.Output[*Solution]{
		OutputState: o.OutputState,
	}
}

// Resource location
func (o SolutionOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Solution) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o SolutionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Solution) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Plan for solution object supported by the OperationsManagement resource provider.
func (o SolutionOutput) Plan() SolutionPlanResponsePtrOutput {
	return o.ApplyT(func(v *Solution) SolutionPlanResponsePtrOutput { return v.Plan }).(SolutionPlanResponsePtrOutput)
}

// Properties for solution object supported by the OperationsManagement resource provider.
func (o SolutionOutput) Properties() SolutionPropertiesResponseOutput {
	return o.ApplyT(func(v *Solution) SolutionPropertiesResponseOutput { return v.Properties }).(SolutionPropertiesResponseOutput)
}

// Resource tags
func (o SolutionOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Solution) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o SolutionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Solution) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SolutionOutput{})
}
