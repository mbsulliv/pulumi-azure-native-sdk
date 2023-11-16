// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scheduler

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Azure REST API version: 2016-03-01. Prior API version in Azure Native 1.x: 2016-03-01.
type JobCollection struct {
	pulumi.CustomResourceState

	// Gets or sets the storage account location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Gets or sets the job collection resource name.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// Gets or sets the job collection properties.
	Properties JobCollectionPropertiesResponseOutput `pulumi:"properties"`
	// Gets or sets the tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Gets the job collection resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewJobCollection registers a new resource with the given unique name, arguments, and options.
func NewJobCollection(ctx *pulumi.Context,
	name string, args *JobCollectionArgs, opts ...pulumi.ResourceOption) (*JobCollection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:scheduler/v20140801preview:JobCollection"),
		},
		{
			Type: pulumi.String("azure-native:scheduler/v20160101:JobCollection"),
		},
		{
			Type: pulumi.String("azure-native:scheduler/v20160301:JobCollection"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource JobCollection
	err := ctx.RegisterResource("azure-native:scheduler:JobCollection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetJobCollection gets an existing JobCollection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetJobCollection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *JobCollectionState, opts ...pulumi.ResourceOption) (*JobCollection, error) {
	var resource JobCollection
	err := ctx.ReadResource("azure-native:scheduler:JobCollection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering JobCollection resources.
type jobCollectionState struct {
}

type JobCollectionState struct {
}

func (JobCollectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*jobCollectionState)(nil)).Elem()
}

type jobCollectionArgs struct {
	// The job collection name.
	JobCollectionName *string `pulumi:"jobCollectionName"`
	// Gets or sets the storage account location.
	Location *string `pulumi:"location"`
	// Gets or sets the job collection resource name.
	Name *string `pulumi:"name"`
	// Gets or sets the job collection properties.
	Properties *JobCollectionProperties `pulumi:"properties"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Gets or sets the tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a JobCollection resource.
type JobCollectionArgs struct {
	// The job collection name.
	JobCollectionName pulumi.StringPtrInput
	// Gets or sets the storage account location.
	Location pulumi.StringPtrInput
	// Gets or sets the job collection resource name.
	Name pulumi.StringPtrInput
	// Gets or sets the job collection properties.
	Properties JobCollectionPropertiesPtrInput
	// The resource group name.
	ResourceGroupName pulumi.StringInput
	// Gets or sets the tags.
	Tags pulumi.StringMapInput
}

func (JobCollectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*jobCollectionArgs)(nil)).Elem()
}

type JobCollectionInput interface {
	pulumi.Input

	ToJobCollectionOutput() JobCollectionOutput
	ToJobCollectionOutputWithContext(ctx context.Context) JobCollectionOutput
}

func (*JobCollection) ElementType() reflect.Type {
	return reflect.TypeOf((**JobCollection)(nil)).Elem()
}

func (i *JobCollection) ToJobCollectionOutput() JobCollectionOutput {
	return i.ToJobCollectionOutputWithContext(context.Background())
}

func (i *JobCollection) ToJobCollectionOutputWithContext(ctx context.Context) JobCollectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobCollectionOutput)
}

type JobCollectionOutput struct{ *pulumi.OutputState }

func (JobCollectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JobCollection)(nil)).Elem()
}

func (o JobCollectionOutput) ToJobCollectionOutput() JobCollectionOutput {
	return o
}

func (o JobCollectionOutput) ToJobCollectionOutputWithContext(ctx context.Context) JobCollectionOutput {
	return o
}

// Gets or sets the storage account location.
func (o JobCollectionOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobCollection) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Gets or sets the job collection resource name.
func (o JobCollectionOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobCollection) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

// Gets or sets the job collection properties.
func (o JobCollectionOutput) Properties() JobCollectionPropertiesResponseOutput {
	return o.ApplyT(func(v *JobCollection) JobCollectionPropertiesResponseOutput { return v.Properties }).(JobCollectionPropertiesResponseOutput)
}

// Gets or sets the tags.
func (o JobCollectionOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *JobCollection) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Gets the job collection resource type.
func (o JobCollectionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *JobCollection) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(JobCollectionOutput{})
}
