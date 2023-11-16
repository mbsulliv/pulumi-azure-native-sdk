// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20231004preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A Solution resource belonging to an Instance resource.
type Solution struct {
	pulumi.CustomResourceState

	// A list of components
	Components ComponentPropertiesResponseArrayOutput `pulumi:"components"`
	// Edge location of the resource.
	ExtendedLocation ExtendedLocationResponseOutput `pulumi:"extendedLocation"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The status of the last operation.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// Version of the particular resource.
	Version pulumi.StringPtrOutput `pulumi:"version"`
}

// NewSolution registers a new resource with the given unique name, arguments, and options.
func NewSolution(ctx *pulumi.Context,
	name string, args *SolutionArgs, opts ...pulumi.ResourceOption) (*Solution, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ExtendedLocation == nil {
		return nil, errors.New("invalid value for required argument 'ExtendedLocation'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:iotoperationsorchestrator:Solution"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Solution
	err := ctx.RegisterResource("azure-native:iotoperationsorchestrator/v20231004preview:Solution", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure-native:iotoperationsorchestrator/v20231004preview:Solution", name, id, state, &resource, opts...)
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
	// A list of components
	Components []ComponentProperties `pulumi:"components"`
	// Edge location of the resource.
	ExtendedLocation ExtendedLocation `pulumi:"extendedLocation"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// Name of solution.
	Name *string `pulumi:"name"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Version of the particular resource.
	Version *string `pulumi:"version"`
}

// The set of arguments for constructing a Solution resource.
type SolutionArgs struct {
	// A list of components
	Components ComponentPropertiesArrayInput
	// Edge location of the resource.
	ExtendedLocation ExtendedLocationInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// Name of solution.
	Name pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Version of the particular resource.
	Version pulumi.StringPtrInput
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

// A list of components
func (o SolutionOutput) Components() ComponentPropertiesResponseArrayOutput {
	return o.ApplyT(func(v *Solution) ComponentPropertiesResponseArrayOutput { return v.Components }).(ComponentPropertiesResponseArrayOutput)
}

// Edge location of the resource.
func (o SolutionOutput) ExtendedLocation() ExtendedLocationResponseOutput {
	return o.ApplyT(func(v *Solution) ExtendedLocationResponseOutput { return v.ExtendedLocation }).(ExtendedLocationResponseOutput)
}

// The geo-location where the resource lives
func (o SolutionOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Solution) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o SolutionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Solution) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The status of the last operation.
func (o SolutionOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Solution) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o SolutionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Solution) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o SolutionOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Solution) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o SolutionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Solution) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Version of the particular resource.
func (o SolutionOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Solution) pulumi.StringPtrOutput { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(SolutionOutput{})
}
