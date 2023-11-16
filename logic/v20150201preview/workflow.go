// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20150201preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Workflow struct {
	pulumi.CustomResourceState

	// Gets the access endpoint.
	AccessEndpoint pulumi.StringOutput `pulumi:"accessEndpoint"`
	// Gets the changed time.
	ChangedTime pulumi.StringOutput `pulumi:"changedTime"`
	// Gets the created time.
	CreatedTime pulumi.StringOutput `pulumi:"createdTime"`
	// Gets or sets the definition.
	Definition pulumi.AnyOutput `pulumi:"definition"`
	// Gets or sets the link to definition.
	DefinitionLink ContentLinkResponsePtrOutput `pulumi:"definitionLink"`
	// Gets or sets the resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Gets the resource name.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// Gets or sets the parameters.
	Parameters WorkflowParameterResponseMapOutput `pulumi:"parameters"`
	// Gets or sets the link to parameters.
	ParametersLink ContentLinkResponsePtrOutput `pulumi:"parametersLink"`
	// Gets the provisioning state.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Gets or sets the sku.
	Sku SkuResponsePtrOutput `pulumi:"sku"`
	// Gets or sets the state.
	State pulumi.StringPtrOutput `pulumi:"state"`
	// Gets or sets the resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Gets the resource type.
	Type pulumi.StringPtrOutput `pulumi:"type"`
	// Gets the version.
	Version pulumi.StringOutput `pulumi:"version"`
}

// NewWorkflow registers a new resource with the given unique name, arguments, and options.
func NewWorkflow(ctx *pulumi.Context,
	name string, args *WorkflowArgs, opts ...pulumi.ResourceOption) (*Workflow, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:logic:Workflow"),
		},
		{
			Type: pulumi.String("azure-native:logic/v20160601:Workflow"),
		},
		{
			Type: pulumi.String("azure-native:logic/v20180701preview:Workflow"),
		},
		{
			Type: pulumi.String("azure-native:logic/v20190501:Workflow"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Workflow
	err := ctx.RegisterResource("azure-native:logic/v20150201preview:Workflow", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkflow gets an existing Workflow resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkflow(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkflowState, opts ...pulumi.ResourceOption) (*Workflow, error) {
	var resource Workflow
	err := ctx.ReadResource("azure-native:logic/v20150201preview:Workflow", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Workflow resources.
type workflowState struct {
}

type WorkflowState struct {
}

func (WorkflowState) ElementType() reflect.Type {
	return reflect.TypeOf((*workflowState)(nil)).Elem()
}

type workflowArgs struct {
	// Gets or sets the definition.
	Definition interface{} `pulumi:"definition"`
	// Gets or sets the link to definition.
	DefinitionLink *ContentLink `pulumi:"definitionLink"`
	// Gets or sets the resource id.
	Id *string `pulumi:"id"`
	// Gets or sets the resource location.
	Location *string `pulumi:"location"`
	// Gets the resource name.
	Name *string `pulumi:"name"`
	// Gets or sets the parameters.
	Parameters map[string]WorkflowParameter `pulumi:"parameters"`
	// Gets or sets the link to parameters.
	ParametersLink *ContentLink `pulumi:"parametersLink"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Gets or sets the sku.
	Sku *Sku `pulumi:"sku"`
	// Gets or sets the state.
	State *WorkflowStateEnum `pulumi:"state"`
	// Gets or sets the resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Gets the resource type.
	Type *string `pulumi:"type"`
	// The workflow name.
	WorkflowName *string `pulumi:"workflowName"`
}

// The set of arguments for constructing a Workflow resource.
type WorkflowArgs struct {
	// Gets or sets the definition.
	Definition pulumi.Input
	// Gets or sets the link to definition.
	DefinitionLink ContentLinkPtrInput
	// Gets or sets the resource id.
	Id pulumi.StringPtrInput
	// Gets or sets the resource location.
	Location pulumi.StringPtrInput
	// Gets the resource name.
	Name pulumi.StringPtrInput
	// Gets or sets the parameters.
	Parameters WorkflowParameterMapInput
	// Gets or sets the link to parameters.
	ParametersLink ContentLinkPtrInput
	// The resource group name.
	ResourceGroupName pulumi.StringInput
	// Gets or sets the sku.
	Sku SkuPtrInput
	// Gets or sets the state.
	State WorkflowStateEnumPtrInput
	// Gets or sets the resource tags.
	Tags pulumi.StringMapInput
	// Gets the resource type.
	Type pulumi.StringPtrInput
	// The workflow name.
	WorkflowName pulumi.StringPtrInput
}

func (WorkflowArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workflowArgs)(nil)).Elem()
}

type WorkflowInput interface {
	pulumi.Input

	ToWorkflowOutput() WorkflowOutput
	ToWorkflowOutputWithContext(ctx context.Context) WorkflowOutput
}

func (*Workflow) ElementType() reflect.Type {
	return reflect.TypeOf((**Workflow)(nil)).Elem()
}

func (i *Workflow) ToWorkflowOutput() WorkflowOutput {
	return i.ToWorkflowOutputWithContext(context.Background())
}

func (i *Workflow) ToWorkflowOutputWithContext(ctx context.Context) WorkflowOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkflowOutput)
}

type WorkflowOutput struct{ *pulumi.OutputState }

func (WorkflowOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Workflow)(nil)).Elem()
}

func (o WorkflowOutput) ToWorkflowOutput() WorkflowOutput {
	return o
}

func (o WorkflowOutput) ToWorkflowOutputWithContext(ctx context.Context) WorkflowOutput {
	return o
}

// Gets the access endpoint.
func (o WorkflowOutput) AccessEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *Workflow) pulumi.StringOutput { return v.AccessEndpoint }).(pulumi.StringOutput)
}

// Gets the changed time.
func (o WorkflowOutput) ChangedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Workflow) pulumi.StringOutput { return v.ChangedTime }).(pulumi.StringOutput)
}

// Gets the created time.
func (o WorkflowOutput) CreatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Workflow) pulumi.StringOutput { return v.CreatedTime }).(pulumi.StringOutput)
}

// Gets or sets the definition.
func (o WorkflowOutput) Definition() pulumi.AnyOutput {
	return o.ApplyT(func(v *Workflow) pulumi.AnyOutput { return v.Definition }).(pulumi.AnyOutput)
}

// Gets or sets the link to definition.
func (o WorkflowOutput) DefinitionLink() ContentLinkResponsePtrOutput {
	return o.ApplyT(func(v *Workflow) ContentLinkResponsePtrOutput { return v.DefinitionLink }).(ContentLinkResponsePtrOutput)
}

// Gets or sets the resource location.
func (o WorkflowOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workflow) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Gets the resource name.
func (o WorkflowOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workflow) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

// Gets or sets the parameters.
func (o WorkflowOutput) Parameters() WorkflowParameterResponseMapOutput {
	return o.ApplyT(func(v *Workflow) WorkflowParameterResponseMapOutput { return v.Parameters }).(WorkflowParameterResponseMapOutput)
}

// Gets or sets the link to parameters.
func (o WorkflowOutput) ParametersLink() ContentLinkResponsePtrOutput {
	return o.ApplyT(func(v *Workflow) ContentLinkResponsePtrOutput { return v.ParametersLink }).(ContentLinkResponsePtrOutput)
}

// Gets the provisioning state.
func (o WorkflowOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Workflow) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Gets or sets the sku.
func (o WorkflowOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *Workflow) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

// Gets or sets the state.
func (o WorkflowOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workflow) pulumi.StringPtrOutput { return v.State }).(pulumi.StringPtrOutput)
}

// Gets or sets the resource tags.
func (o WorkflowOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Workflow) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Gets the resource type.
func (o WorkflowOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workflow) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

// Gets the version.
func (o WorkflowOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v *Workflow) pulumi.StringOutput { return v.Version }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WorkflowOutput{})
}
