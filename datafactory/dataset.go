// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package datafactory

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Dataset resource type.
// Azure REST API version: 2018-06-01. Prior API version in Azure Native 1.x: 2018-06-01
type Dataset struct {
	pulumi.CustomResourceState

	// Etag identifies change in the resource.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Dataset properties.
	Properties pulumi.AnyOutput `pulumi:"properties"`
	// The resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDataset registers a new resource with the given unique name, arguments, and options.
func NewDataset(ctx *pulumi.Context,
	name string, args *DatasetArgs, opts ...pulumi.ResourceOption) (*Dataset, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FactoryName == nil {
		return nil, errors.New("invalid value for required argument 'FactoryName'")
	}
	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:datafactory/v20170901preview:Dataset"),
		},
		{
			Type: pulumi.String("azure-native:datafactory/v20180601:Dataset"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Dataset
	err := ctx.RegisterResource("azure-native:datafactory:Dataset", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataset gets an existing Dataset resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataset(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatasetState, opts ...pulumi.ResourceOption) (*Dataset, error) {
	var resource Dataset
	err := ctx.ReadResource("azure-native:datafactory:Dataset", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Dataset resources.
type datasetState struct {
}

type DatasetState struct {
}

func (DatasetState) ElementType() reflect.Type {
	return reflect.TypeOf((*datasetState)(nil)).Elem()
}

type datasetArgs struct {
	// The dataset name.
	DatasetName *string `pulumi:"datasetName"`
	// The factory name.
	FactoryName string `pulumi:"factoryName"`
	// Dataset properties.
	Properties interface{} `pulumi:"properties"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a Dataset resource.
type DatasetArgs struct {
	// The dataset name.
	DatasetName pulumi.StringPtrInput
	// The factory name.
	FactoryName pulumi.StringInput
	// Dataset properties.
	Properties pulumi.Input
	// The resource group name.
	ResourceGroupName pulumi.StringInput
}

func (DatasetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*datasetArgs)(nil)).Elem()
}

type DatasetInput interface {
	pulumi.Input

	ToDatasetOutput() DatasetOutput
	ToDatasetOutputWithContext(ctx context.Context) DatasetOutput
}

func (*Dataset) ElementType() reflect.Type {
	return reflect.TypeOf((**Dataset)(nil)).Elem()
}

func (i *Dataset) ToDatasetOutput() DatasetOutput {
	return i.ToDatasetOutputWithContext(context.Background())
}

func (i *Dataset) ToDatasetOutputWithContext(ctx context.Context) DatasetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasetOutput)
}

func (i *Dataset) ToOutput(ctx context.Context) pulumix.Output[*Dataset] {
	return pulumix.Output[*Dataset]{
		OutputState: i.ToDatasetOutputWithContext(ctx).OutputState,
	}
}

type DatasetOutput struct{ *pulumi.OutputState }

func (DatasetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Dataset)(nil)).Elem()
}

func (o DatasetOutput) ToDatasetOutput() DatasetOutput {
	return o
}

func (o DatasetOutput) ToDatasetOutputWithContext(ctx context.Context) DatasetOutput {
	return o
}

func (o DatasetOutput) ToOutput(ctx context.Context) pulumix.Output[*Dataset] {
	return pulumix.Output[*Dataset]{
		OutputState: o.OutputState,
	}
}

// Etag identifies change in the resource.
func (o DatasetOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *Dataset) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The resource name.
func (o DatasetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Dataset) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Dataset properties.
func (o DatasetOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v *Dataset) pulumi.AnyOutput { return v.Properties }).(pulumi.AnyOutput)
}

// The resource type.
func (o DatasetOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Dataset) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DatasetOutput{})
}
