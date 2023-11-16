// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package containerregistry

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An object that represents an import pipeline for a container registry.
// Azure REST API version: 2023-01-01-preview. Prior API version in Azure Native 1.x: 2020-11-01-preview.
//
// Other available API versions: 2023-06-01-preview, 2023-08-01-preview, 2023-11-01-preview.
type ImportPipeline struct {
	pulumi.CustomResourceState

	// The identity of the import pipeline.
	Identity IdentityPropertiesResponsePtrOutput `pulumi:"identity"`
	// The location of the import pipeline.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The list of all options configured for the pipeline.
	Options pulumi.StringArrayOutput `pulumi:"options"`
	// The provisioning state of the pipeline at the time the operation was called.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The source properties of the import pipeline.
	Source ImportPipelineSourcePropertiesResponseOutput `pulumi:"source"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The properties that describe the trigger of the import pipeline.
	Trigger PipelineTriggerPropertiesResponsePtrOutput `pulumi:"trigger"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewImportPipeline registers a new resource with the given unique name, arguments, and options.
func NewImportPipeline(ctx *pulumi.Context,
	name string, args *ImportPipelineArgs, opts ...pulumi.ResourceOption) (*ImportPipeline, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RegistryName == nil {
		return nil, errors.New("invalid value for required argument 'RegistryName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Source == nil {
		return nil, errors.New("invalid value for required argument 'Source'")
	}
	args.Source = args.Source.ToImportPipelineSourcePropertiesOutput().ApplyT(func(v ImportPipelineSourceProperties) ImportPipelineSourceProperties { return *v.Defaults() }).(ImportPipelineSourcePropertiesOutput)
	if args.Trigger != nil {
		args.Trigger = args.Trigger.ToPipelineTriggerPropertiesPtrOutput().ApplyT(func(v *PipelineTriggerProperties) *PipelineTriggerProperties { return v.Defaults() }).(PipelineTriggerPropertiesPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:containerregistry/v20191201preview:ImportPipeline"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20201101preview:ImportPipeline"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20210601preview:ImportPipeline"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20210801preview:ImportPipeline"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20211201preview:ImportPipeline"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20220201preview:ImportPipeline"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20230101preview:ImportPipeline"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20230601preview:ImportPipeline"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20230801preview:ImportPipeline"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20231101preview:ImportPipeline"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ImportPipeline
	err := ctx.RegisterResource("azure-native:containerregistry:ImportPipeline", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetImportPipeline gets an existing ImportPipeline resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetImportPipeline(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ImportPipelineState, opts ...pulumi.ResourceOption) (*ImportPipeline, error) {
	var resource ImportPipeline
	err := ctx.ReadResource("azure-native:containerregistry:ImportPipeline", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ImportPipeline resources.
type importPipelineState struct {
}

type ImportPipelineState struct {
}

func (ImportPipelineState) ElementType() reflect.Type {
	return reflect.TypeOf((*importPipelineState)(nil)).Elem()
}

type importPipelineArgs struct {
	// The identity of the import pipeline.
	Identity *IdentityProperties `pulumi:"identity"`
	// The name of the import pipeline.
	ImportPipelineName *string `pulumi:"importPipelineName"`
	// The location of the import pipeline.
	Location *string `pulumi:"location"`
	// The list of all options configured for the pipeline.
	Options []string `pulumi:"options"`
	// The name of the container registry.
	RegistryName string `pulumi:"registryName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The source properties of the import pipeline.
	Source ImportPipelineSourceProperties `pulumi:"source"`
	// The properties that describe the trigger of the import pipeline.
	Trigger *PipelineTriggerProperties `pulumi:"trigger"`
}

// The set of arguments for constructing a ImportPipeline resource.
type ImportPipelineArgs struct {
	// The identity of the import pipeline.
	Identity IdentityPropertiesPtrInput
	// The name of the import pipeline.
	ImportPipelineName pulumi.StringPtrInput
	// The location of the import pipeline.
	Location pulumi.StringPtrInput
	// The list of all options configured for the pipeline.
	Options pulumi.StringArrayInput
	// The name of the container registry.
	RegistryName pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The source properties of the import pipeline.
	Source ImportPipelineSourcePropertiesInput
	// The properties that describe the trigger of the import pipeline.
	Trigger PipelineTriggerPropertiesPtrInput
}

func (ImportPipelineArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*importPipelineArgs)(nil)).Elem()
}

type ImportPipelineInput interface {
	pulumi.Input

	ToImportPipelineOutput() ImportPipelineOutput
	ToImportPipelineOutputWithContext(ctx context.Context) ImportPipelineOutput
}

func (*ImportPipeline) ElementType() reflect.Type {
	return reflect.TypeOf((**ImportPipeline)(nil)).Elem()
}

func (i *ImportPipeline) ToImportPipelineOutput() ImportPipelineOutput {
	return i.ToImportPipelineOutputWithContext(context.Background())
}

func (i *ImportPipeline) ToImportPipelineOutputWithContext(ctx context.Context) ImportPipelineOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImportPipelineOutput)
}

type ImportPipelineOutput struct{ *pulumi.OutputState }

func (ImportPipelineOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ImportPipeline)(nil)).Elem()
}

func (o ImportPipelineOutput) ToImportPipelineOutput() ImportPipelineOutput {
	return o
}

func (o ImportPipelineOutput) ToImportPipelineOutputWithContext(ctx context.Context) ImportPipelineOutput {
	return o
}

// The identity of the import pipeline.
func (o ImportPipelineOutput) Identity() IdentityPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *ImportPipeline) IdentityPropertiesResponsePtrOutput { return v.Identity }).(IdentityPropertiesResponsePtrOutput)
}

// The location of the import pipeline.
func (o ImportPipelineOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImportPipeline) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the resource.
func (o ImportPipelineOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ImportPipeline) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The list of all options configured for the pipeline.
func (o ImportPipelineOutput) Options() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ImportPipeline) pulumi.StringArrayOutput { return v.Options }).(pulumi.StringArrayOutput)
}

// The provisioning state of the pipeline at the time the operation was called.
func (o ImportPipelineOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ImportPipeline) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The source properties of the import pipeline.
func (o ImportPipelineOutput) Source() ImportPipelineSourcePropertiesResponseOutput {
	return o.ApplyT(func(v *ImportPipeline) ImportPipelineSourcePropertiesResponseOutput { return v.Source }).(ImportPipelineSourcePropertiesResponseOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o ImportPipelineOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ImportPipeline) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The properties that describe the trigger of the import pipeline.
func (o ImportPipelineOutput) Trigger() PipelineTriggerPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *ImportPipeline) PipelineTriggerPropertiesResponsePtrOutput { return v.Trigger }).(PipelineTriggerPropertiesResponsePtrOutput)
}

// The type of the resource.
func (o ImportPipelineOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ImportPipeline) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ImportPipelineOutput{})
}
