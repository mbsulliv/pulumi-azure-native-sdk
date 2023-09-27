// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package videoanalyzer

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Pipeline topology describes the processing steps to be applied when processing content for a particular outcome. The topology should be defined according to the scenario to be achieved and can be reused across many pipeline instances which share the same processing characteristics. For instance, a pipeline topology which captures content from a RTSP camera and archives the content can be reused across many different cameras, as long as the same processing is to be applied across all the cameras. Individual instance properties can be defined through the use of user-defined parameters, which allow for a topology to be parameterized. This allows  individual pipelines refer to different values, such as individual cameras' RTSP endpoints and credentials. Overall a topology is composed of the following:
//
//   - Parameters: list of user defined parameters that can be references across the topology nodes.
//   - Sources: list of one or more data sources nodes such as an RTSP source which allows for content to be ingested from cameras.
//   - Processors: list of nodes which perform data analysis or transformations.
//   - Sinks: list of one or more data sinks which allow for data to be stored or exported to other destinations.
//     Azure REST API version: 2021-11-01-preview. Prior API version in Azure Native 1.x: 2021-11-01-preview
type PipelineTopology struct {
	pulumi.CustomResourceState

	// An optional description of the pipeline topology. It is recommended that the expected use of the topology to be described here.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Topology kind.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// List of the topology parameter declarations. Parameters declared here can be referenced throughout the topology nodes through the use of "${PARAMETER_NAME}" string pattern. Parameters can have optional default values and can later be defined in individual instances of the pipeline.
	Parameters ParameterDeclarationResponseArrayOutput `pulumi:"parameters"`
	// List of the topology processor nodes. Processor nodes enable pipeline data to be analyzed, processed or transformed.
	Processors EncoderProcessorResponseArrayOutput `pulumi:"processors"`
	// List of the topology sink nodes. Sink nodes allow pipeline data to be stored or exported.
	Sinks VideoSinkResponseArrayOutput `pulumi:"sinks"`
	// Describes the properties of a SKU.
	Sku SkuResponseOutput `pulumi:"sku"`
	// List of the topology source nodes. Source nodes enable external data to be ingested by the pipeline.
	Sources pulumi.ArrayOutput `pulumi:"sources"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewPipelineTopology registers a new resource with the given unique name, arguments, and options.
func NewPipelineTopology(ctx *pulumi.Context,
	name string, args *PipelineTopologyArgs, opts ...pulumi.ResourceOption) (*PipelineTopology, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Sinks == nil {
		return nil, errors.New("invalid value for required argument 'Sinks'")
	}
	if args.Sku == nil {
		return nil, errors.New("invalid value for required argument 'Sku'")
	}
	if args.Sources == nil {
		return nil, errors.New("invalid value for required argument 'Sources'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:videoanalyzer/v20211101preview:PipelineTopology"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource PipelineTopology
	err := ctx.RegisterResource("azure-native:videoanalyzer:PipelineTopology", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPipelineTopology gets an existing PipelineTopology resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPipelineTopology(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PipelineTopologyState, opts ...pulumi.ResourceOption) (*PipelineTopology, error) {
	var resource PipelineTopology
	err := ctx.ReadResource("azure-native:videoanalyzer:PipelineTopology", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PipelineTopology resources.
type pipelineTopologyState struct {
}

type PipelineTopologyState struct {
}

func (PipelineTopologyState) ElementType() reflect.Type {
	return reflect.TypeOf((*pipelineTopologyState)(nil)).Elem()
}

type pipelineTopologyArgs struct {
	// The Azure Video Analyzer account name.
	AccountName string `pulumi:"accountName"`
	// An optional description of the pipeline topology. It is recommended that the expected use of the topology to be described here.
	Description *string `pulumi:"description"`
	// Topology kind.
	Kind string `pulumi:"kind"`
	// List of the topology parameter declarations. Parameters declared here can be referenced throughout the topology nodes through the use of "${PARAMETER_NAME}" string pattern. Parameters can have optional default values and can later be defined in individual instances of the pipeline.
	Parameters []ParameterDeclaration `pulumi:"parameters"`
	// Pipeline topology unique identifier.
	PipelineTopologyName *string `pulumi:"pipelineTopologyName"`
	// List of the topology processor nodes. Processor nodes enable pipeline data to be analyzed, processed or transformed.
	Processors []EncoderProcessor `pulumi:"processors"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// List of the topology sink nodes. Sink nodes allow pipeline data to be stored or exported.
	Sinks []VideoSink `pulumi:"sinks"`
	// Describes the properties of a SKU.
	Sku Sku `pulumi:"sku"`
	// List of the topology source nodes. Source nodes enable external data to be ingested by the pipeline.
	Sources []interface{} `pulumi:"sources"`
}

// The set of arguments for constructing a PipelineTopology resource.
type PipelineTopologyArgs struct {
	// The Azure Video Analyzer account name.
	AccountName pulumi.StringInput
	// An optional description of the pipeline topology. It is recommended that the expected use of the topology to be described here.
	Description pulumi.StringPtrInput
	// Topology kind.
	Kind pulumi.StringInput
	// List of the topology parameter declarations. Parameters declared here can be referenced throughout the topology nodes through the use of "${PARAMETER_NAME}" string pattern. Parameters can have optional default values and can later be defined in individual instances of the pipeline.
	Parameters ParameterDeclarationArrayInput
	// Pipeline topology unique identifier.
	PipelineTopologyName pulumi.StringPtrInput
	// List of the topology processor nodes. Processor nodes enable pipeline data to be analyzed, processed or transformed.
	Processors EncoderProcessorArrayInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// List of the topology sink nodes. Sink nodes allow pipeline data to be stored or exported.
	Sinks VideoSinkArrayInput
	// Describes the properties of a SKU.
	Sku SkuInput
	// List of the topology source nodes. Source nodes enable external data to be ingested by the pipeline.
	Sources pulumi.ArrayInput
}

func (PipelineTopologyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*pipelineTopologyArgs)(nil)).Elem()
}

type PipelineTopologyInput interface {
	pulumi.Input

	ToPipelineTopologyOutput() PipelineTopologyOutput
	ToPipelineTopologyOutputWithContext(ctx context.Context) PipelineTopologyOutput
}

func (*PipelineTopology) ElementType() reflect.Type {
	return reflect.TypeOf((**PipelineTopology)(nil)).Elem()
}

func (i *PipelineTopology) ToPipelineTopologyOutput() PipelineTopologyOutput {
	return i.ToPipelineTopologyOutputWithContext(context.Background())
}

func (i *PipelineTopology) ToPipelineTopologyOutputWithContext(ctx context.Context) PipelineTopologyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineTopologyOutput)
}

func (i *PipelineTopology) ToOutput(ctx context.Context) pulumix.Output[*PipelineTopology] {
	return pulumix.Output[*PipelineTopology]{
		OutputState: i.ToPipelineTopologyOutputWithContext(ctx).OutputState,
	}
}

type PipelineTopologyOutput struct{ *pulumi.OutputState }

func (PipelineTopologyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PipelineTopology)(nil)).Elem()
}

func (o PipelineTopologyOutput) ToPipelineTopologyOutput() PipelineTopologyOutput {
	return o
}

func (o PipelineTopologyOutput) ToPipelineTopologyOutputWithContext(ctx context.Context) PipelineTopologyOutput {
	return o
}

func (o PipelineTopologyOutput) ToOutput(ctx context.Context) pulumix.Output[*PipelineTopology] {
	return pulumix.Output[*PipelineTopology]{
		OutputState: o.OutputState,
	}
}

// An optional description of the pipeline topology. It is recommended that the expected use of the topology to be described here.
func (o PipelineTopologyOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PipelineTopology) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Topology kind.
func (o PipelineTopologyOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *PipelineTopology) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// The name of the resource
func (o PipelineTopologyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PipelineTopology) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// List of the topology parameter declarations. Parameters declared here can be referenced throughout the topology nodes through the use of "${PARAMETER_NAME}" string pattern. Parameters can have optional default values and can later be defined in individual instances of the pipeline.
func (o PipelineTopologyOutput) Parameters() ParameterDeclarationResponseArrayOutput {
	return o.ApplyT(func(v *PipelineTopology) ParameterDeclarationResponseArrayOutput { return v.Parameters }).(ParameterDeclarationResponseArrayOutput)
}

// List of the topology processor nodes. Processor nodes enable pipeline data to be analyzed, processed or transformed.
func (o PipelineTopologyOutput) Processors() EncoderProcessorResponseArrayOutput {
	return o.ApplyT(func(v *PipelineTopology) EncoderProcessorResponseArrayOutput { return v.Processors }).(EncoderProcessorResponseArrayOutput)
}

// List of the topology sink nodes. Sink nodes allow pipeline data to be stored or exported.
func (o PipelineTopologyOutput) Sinks() VideoSinkResponseArrayOutput {
	return o.ApplyT(func(v *PipelineTopology) VideoSinkResponseArrayOutput { return v.Sinks }).(VideoSinkResponseArrayOutput)
}

// Describes the properties of a SKU.
func (o PipelineTopologyOutput) Sku() SkuResponseOutput {
	return o.ApplyT(func(v *PipelineTopology) SkuResponseOutput { return v.Sku }).(SkuResponseOutput)
}

// List of the topology source nodes. Source nodes enable external data to be ingested by the pipeline.
func (o PipelineTopologyOutput) Sources() pulumi.ArrayOutput {
	return o.ApplyT(func(v *PipelineTopology) pulumi.ArrayOutput { return v.Sources }).(pulumi.ArrayOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o PipelineTopologyOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *PipelineTopology) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o PipelineTopologyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PipelineTopology) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PipelineTopologyOutput{})
}
