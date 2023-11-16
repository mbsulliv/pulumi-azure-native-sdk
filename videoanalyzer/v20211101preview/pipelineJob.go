// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20211101preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Pipeline job represents a unique instance of a batch topology, used for offline processing of selected portions of archived content.
type PipelineJob struct {
	pulumi.CustomResourceState

	// An optional description for the pipeline.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Details about the error, in case the pipeline job fails.
	Error PipelineJobErrorResponseOutput `pulumi:"error"`
	// The date-time by when this pipeline job will be automatically deleted from your account.
	Expiration pulumi.StringOutput `pulumi:"expiration"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// List of the instance level parameter values for the user-defined topology parameters. A pipeline can only define or override parameters values for parameters which have been declared in the referenced topology. Topology parameters without a default value must be defined. Topology parameters with a default value can be optionally be overridden.
	Parameters ParameterDefinitionResponseArrayOutput `pulumi:"parameters"`
	// Current state of the pipeline (read-only).
	State pulumi.StringOutput `pulumi:"state"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Reference to an existing pipeline topology. When activated, this pipeline job will process content according to the pipeline topology definition.
	TopologyName pulumi.StringOutput `pulumi:"topologyName"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewPipelineJob registers a new resource with the given unique name, arguments, and options.
func NewPipelineJob(ctx *pulumi.Context,
	name string, args *PipelineJobArgs, opts ...pulumi.ResourceOption) (*PipelineJob, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.TopologyName == nil {
		return nil, errors.New("invalid value for required argument 'TopologyName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:videoanalyzer:PipelineJob"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource PipelineJob
	err := ctx.RegisterResource("azure-native:videoanalyzer/v20211101preview:PipelineJob", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPipelineJob gets an existing PipelineJob resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPipelineJob(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PipelineJobState, opts ...pulumi.ResourceOption) (*PipelineJob, error) {
	var resource PipelineJob
	err := ctx.ReadResource("azure-native:videoanalyzer/v20211101preview:PipelineJob", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PipelineJob resources.
type pipelineJobState struct {
}

type PipelineJobState struct {
}

func (PipelineJobState) ElementType() reflect.Type {
	return reflect.TypeOf((*pipelineJobState)(nil)).Elem()
}

type pipelineJobArgs struct {
	// The Azure Video Analyzer account name.
	AccountName string `pulumi:"accountName"`
	// An optional description for the pipeline.
	Description *string `pulumi:"description"`
	// List of the instance level parameter values for the user-defined topology parameters. A pipeline can only define or override parameters values for parameters which have been declared in the referenced topology. Topology parameters without a default value must be defined. Topology parameters with a default value can be optionally be overridden.
	Parameters []ParameterDefinition `pulumi:"parameters"`
	// The pipeline job name.
	PipelineJobName *string `pulumi:"pipelineJobName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Reference to an existing pipeline topology. When activated, this pipeline job will process content according to the pipeline topology definition.
	TopologyName string `pulumi:"topologyName"`
}

// The set of arguments for constructing a PipelineJob resource.
type PipelineJobArgs struct {
	// The Azure Video Analyzer account name.
	AccountName pulumi.StringInput
	// An optional description for the pipeline.
	Description pulumi.StringPtrInput
	// List of the instance level parameter values for the user-defined topology parameters. A pipeline can only define or override parameters values for parameters which have been declared in the referenced topology. Topology parameters without a default value must be defined. Topology parameters with a default value can be optionally be overridden.
	Parameters ParameterDefinitionArrayInput
	// The pipeline job name.
	PipelineJobName pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Reference to an existing pipeline topology. When activated, this pipeline job will process content according to the pipeline topology definition.
	TopologyName pulumi.StringInput
}

func (PipelineJobArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*pipelineJobArgs)(nil)).Elem()
}

type PipelineJobInput interface {
	pulumi.Input

	ToPipelineJobOutput() PipelineJobOutput
	ToPipelineJobOutputWithContext(ctx context.Context) PipelineJobOutput
}

func (*PipelineJob) ElementType() reflect.Type {
	return reflect.TypeOf((**PipelineJob)(nil)).Elem()
}

func (i *PipelineJob) ToPipelineJobOutput() PipelineJobOutput {
	return i.ToPipelineJobOutputWithContext(context.Background())
}

func (i *PipelineJob) ToPipelineJobOutputWithContext(ctx context.Context) PipelineJobOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineJobOutput)
}

type PipelineJobOutput struct{ *pulumi.OutputState }

func (PipelineJobOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PipelineJob)(nil)).Elem()
}

func (o PipelineJobOutput) ToPipelineJobOutput() PipelineJobOutput {
	return o
}

func (o PipelineJobOutput) ToPipelineJobOutputWithContext(ctx context.Context) PipelineJobOutput {
	return o
}

// An optional description for the pipeline.
func (o PipelineJobOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PipelineJob) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Details about the error, in case the pipeline job fails.
func (o PipelineJobOutput) Error() PipelineJobErrorResponseOutput {
	return o.ApplyT(func(v *PipelineJob) PipelineJobErrorResponseOutput { return v.Error }).(PipelineJobErrorResponseOutput)
}

// The date-time by when this pipeline job will be automatically deleted from your account.
func (o PipelineJobOutput) Expiration() pulumi.StringOutput {
	return o.ApplyT(func(v *PipelineJob) pulumi.StringOutput { return v.Expiration }).(pulumi.StringOutput)
}

// The name of the resource
func (o PipelineJobOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PipelineJob) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// List of the instance level parameter values for the user-defined topology parameters. A pipeline can only define or override parameters values for parameters which have been declared in the referenced topology. Topology parameters without a default value must be defined. Topology parameters with a default value can be optionally be overridden.
func (o PipelineJobOutput) Parameters() ParameterDefinitionResponseArrayOutput {
	return o.ApplyT(func(v *PipelineJob) ParameterDefinitionResponseArrayOutput { return v.Parameters }).(ParameterDefinitionResponseArrayOutput)
}

// Current state of the pipeline (read-only).
func (o PipelineJobOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *PipelineJob) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o PipelineJobOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *PipelineJob) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Reference to an existing pipeline topology. When activated, this pipeline job will process content according to the pipeline topology definition.
func (o PipelineJobOutput) TopologyName() pulumi.StringOutput {
	return o.ApplyT(func(v *PipelineJob) pulumi.StringOutput { return v.TopologyName }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o PipelineJobOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PipelineJob) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PipelineJobOutput{})
}
