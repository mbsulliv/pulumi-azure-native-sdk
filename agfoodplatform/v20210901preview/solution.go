// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210901preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Solution resource.
type Solution struct {
	pulumi.CustomResourceState

	// The ETag value to implement optimistic concurrency.
	ETag pulumi.StringOutput `pulumi:"eTag"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Solution resource properties.
	Properties SolutionPropertiesResponseOutput `pulumi:"properties"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewSolution registers a new resource with the given unique name, arguments, and options.
func NewSolution(ctx *pulumi.Context,
	name string, args *SolutionArgs, opts ...pulumi.ResourceOption) (*Solution, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FarmBeatsResourceName == nil {
		return nil, errors.New("invalid value for required argument 'FarmBeatsResourceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:agfoodplatform:Solution"),
		},
		{
			Type: pulumi.String("azure-native:agfoodplatform/v20230601preview:Solution"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Solution
	err := ctx.RegisterResource("azure-native:agfoodplatform/v20210901preview:Solution", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure-native:agfoodplatform/v20210901preview:Solution", name, id, state, &resource, opts...)
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
	// FarmBeats resource name.
	FarmBeatsResourceName string `pulumi:"farmBeatsResourceName"`
	// Solution resource properties.
	Properties *SolutionProperties `pulumi:"properties"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Solution Id of the solution.
	SolutionId *string `pulumi:"solutionId"`
}

// The set of arguments for constructing a Solution resource.
type SolutionArgs struct {
	// FarmBeats resource name.
	FarmBeatsResourceName pulumi.StringInput
	// Solution resource properties.
	Properties SolutionPropertiesPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Solution Id of the solution.
	SolutionId pulumi.StringPtrInput
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

// The ETag value to implement optimistic concurrency.
func (o SolutionOutput) ETag() pulumi.StringOutput {
	return o.ApplyT(func(v *Solution) pulumi.StringOutput { return v.ETag }).(pulumi.StringOutput)
}

// The name of the resource
func (o SolutionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Solution) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Solution resource properties.
func (o SolutionOutput) Properties() SolutionPropertiesResponseOutput {
	return o.ApplyT(func(v *Solution) SolutionPropertiesResponseOutput { return v.Properties }).(SolutionPropertiesResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o SolutionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Solution) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o SolutionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Solution) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SolutionOutput{})
}
