// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210915preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Model that represents a Experiment resource.
type Experiment struct {
	pulumi.CustomResourceState

	// The identity of the experiment resource.
	Identity ResourceIdentityResponsePtrOutput `pulumi:"identity"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The properties of the experiment resource.
	Properties ExperimentPropertiesResponseOutput `pulumi:"properties"`
	// The system metadata of the experiment resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewExperiment registers a new resource with the given unique name, arguments, and options.
func NewExperiment(ctx *pulumi.Context,
	name string, args *ExperimentArgs, opts ...pulumi.ResourceOption) (*Experiment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:chaos:Experiment"),
		},
		{
			Type: pulumi.String("azure-native:chaos/v20220701preview:Experiment"),
		},
		{
			Type: pulumi.String("azure-native:chaos/v20221001preview:Experiment"),
		},
		{
			Type: pulumi.String("azure-native:chaos/v20230401preview:Experiment"),
		},
	})
	opts = append(opts, aliases)
	var resource Experiment
	err := ctx.RegisterResource("azure-native:chaos/v20210915preview:Experiment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetExperiment gets an existing Experiment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetExperiment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExperimentState, opts ...pulumi.ResourceOption) (*Experiment, error) {
	var resource Experiment
	err := ctx.ReadResource("azure-native:chaos/v20210915preview:Experiment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Experiment resources.
type experimentState struct {
}

type ExperimentState struct {
}

func (ExperimentState) ElementType() reflect.Type {
	return reflect.TypeOf((*experimentState)(nil)).Elem()
}

type experimentArgs struct {
	// String that represents a Experiment resource name.
	ExperimentName *string `pulumi:"experimentName"`
	// The identity of the experiment resource.
	Identity *ResourceIdentity `pulumi:"identity"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The properties of the experiment resource.
	Properties ExperimentProperties `pulumi:"properties"`
	// String that represents an Azure resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Experiment resource.
type ExperimentArgs struct {
	// String that represents a Experiment resource name.
	ExperimentName pulumi.StringPtrInput
	// The identity of the experiment resource.
	Identity ResourceIdentityPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The properties of the experiment resource.
	Properties ExperimentPropertiesInput
	// String that represents an Azure resource group.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (ExperimentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*experimentArgs)(nil)).Elem()
}

type ExperimentInput interface {
	pulumi.Input

	ToExperimentOutput() ExperimentOutput
	ToExperimentOutputWithContext(ctx context.Context) ExperimentOutput
}

func (*Experiment) ElementType() reflect.Type {
	return reflect.TypeOf((**Experiment)(nil)).Elem()
}

func (i *Experiment) ToExperimentOutput() ExperimentOutput {
	return i.ToExperimentOutputWithContext(context.Background())
}

func (i *Experiment) ToExperimentOutputWithContext(ctx context.Context) ExperimentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExperimentOutput)
}

type ExperimentOutput struct{ *pulumi.OutputState }

func (ExperimentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Experiment)(nil)).Elem()
}

func (o ExperimentOutput) ToExperimentOutput() ExperimentOutput {
	return o
}

func (o ExperimentOutput) ToExperimentOutputWithContext(ctx context.Context) ExperimentOutput {
	return o
}

// The identity of the experiment resource.
func (o ExperimentOutput) Identity() ResourceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *Experiment) ResourceIdentityResponsePtrOutput { return v.Identity }).(ResourceIdentityResponsePtrOutput)
}

// The geo-location where the resource lives
func (o ExperimentOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Experiment) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o ExperimentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Experiment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The properties of the experiment resource.
func (o ExperimentOutput) Properties() ExperimentPropertiesResponseOutput {
	return o.ApplyT(func(v *Experiment) ExperimentPropertiesResponseOutput { return v.Properties }).(ExperimentPropertiesResponseOutput)
}

// The system metadata of the experiment resource.
func (o ExperimentOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Experiment) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o ExperimentOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Experiment) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ExperimentOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Experiment) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ExperimentOutput{})
}
