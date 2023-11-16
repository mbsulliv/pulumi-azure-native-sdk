// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220808

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of the runbook type.
type Runbook struct {
	pulumi.CustomResourceState

	// Gets or sets the creation time.
	CreationTime pulumi.StringPtrOutput `pulumi:"creationTime"`
	// Gets or sets the description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Gets or sets the draft runbook properties.
	Draft RunbookDraftResponsePtrOutput `pulumi:"draft"`
	// Gets or sets the etag of the resource.
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// Gets or sets the job count of the runbook.
	JobCount pulumi.IntPtrOutput `pulumi:"jobCount"`
	// Gets or sets the last modified by.
	LastModifiedBy pulumi.StringPtrOutput `pulumi:"lastModifiedBy"`
	// Gets or sets the last modified time.
	LastModifiedTime pulumi.StringPtrOutput `pulumi:"lastModifiedTime"`
	// The Azure Region where the resource lives
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Gets or sets the option to log activity trace of the runbook.
	LogActivityTrace pulumi.IntPtrOutput `pulumi:"logActivityTrace"`
	// Gets or sets progress log option.
	LogProgress pulumi.BoolPtrOutput `pulumi:"logProgress"`
	// Gets or sets verbose log option.
	LogVerbose pulumi.BoolPtrOutput `pulumi:"logVerbose"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Gets or sets the runbook output types.
	OutputTypes pulumi.StringArrayOutput `pulumi:"outputTypes"`
	// Gets or sets the runbook parameters.
	Parameters RunbookParameterResponseMapOutput `pulumi:"parameters"`
	// Gets or sets the provisioning state of the runbook.
	ProvisioningState pulumi.StringPtrOutput `pulumi:"provisioningState"`
	// Gets or sets the published runbook content link.
	PublishContentLink ContentLinkResponsePtrOutput `pulumi:"publishContentLink"`
	// Gets or sets the type of the runbook.
	RunbookType pulumi.StringPtrOutput `pulumi:"runbookType"`
	// Gets or sets the state of the runbook.
	State pulumi.StringPtrOutput `pulumi:"state"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewRunbook registers a new resource with the given unique name, arguments, and options.
func NewRunbook(ctx *pulumi.Context,
	name string, args *RunbookArgs, opts ...pulumi.ResourceOption) (*Runbook, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AutomationAccountName == nil {
		return nil, errors.New("invalid value for required argument 'AutomationAccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.RunbookType == nil {
		return nil, errors.New("invalid value for required argument 'RunbookType'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:automation:Runbook"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20151031:Runbook"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20180630:Runbook"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20190601:Runbook"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20230515preview:Runbook"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20231101:Runbook"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Runbook
	err := ctx.RegisterResource("azure-native:automation/v20220808:Runbook", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRunbook gets an existing Runbook resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRunbook(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RunbookState, opts ...pulumi.ResourceOption) (*Runbook, error) {
	var resource Runbook
	err := ctx.ReadResource("azure-native:automation/v20220808:Runbook", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Runbook resources.
type runbookState struct {
}

type RunbookState struct {
}

func (RunbookState) ElementType() reflect.Type {
	return reflect.TypeOf((*runbookState)(nil)).Elem()
}

type runbookArgs struct {
	// The name of the automation account.
	AutomationAccountName string `pulumi:"automationAccountName"`
	// Gets or sets the description of the runbook.
	Description *string `pulumi:"description"`
	// Gets or sets the draft runbook properties.
	Draft *RunbookDraft `pulumi:"draft"`
	// Gets or sets the location of the resource.
	Location *string `pulumi:"location"`
	// Gets or sets the activity-level tracing options of the runbook.
	LogActivityTrace *int `pulumi:"logActivityTrace"`
	// Gets or sets progress log option.
	LogProgress *bool `pulumi:"logProgress"`
	// Gets or sets verbose log option.
	LogVerbose *bool `pulumi:"logVerbose"`
	// Gets or sets the name of the resource.
	Name *string `pulumi:"name"`
	// Gets or sets the published runbook content link.
	PublishContentLink *ContentLink `pulumi:"publishContentLink"`
	// Name of an Azure Resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The runbook name.
	RunbookName *string `pulumi:"runbookName"`
	// Gets or sets the type of the runbook.
	RunbookType string `pulumi:"runbookType"`
	// Gets or sets the tags attached to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Runbook resource.
type RunbookArgs struct {
	// The name of the automation account.
	AutomationAccountName pulumi.StringInput
	// Gets or sets the description of the runbook.
	Description pulumi.StringPtrInput
	// Gets or sets the draft runbook properties.
	Draft RunbookDraftPtrInput
	// Gets or sets the location of the resource.
	Location pulumi.StringPtrInput
	// Gets or sets the activity-level tracing options of the runbook.
	LogActivityTrace pulumi.IntPtrInput
	// Gets or sets progress log option.
	LogProgress pulumi.BoolPtrInput
	// Gets or sets verbose log option.
	LogVerbose pulumi.BoolPtrInput
	// Gets or sets the name of the resource.
	Name pulumi.StringPtrInput
	// Gets or sets the published runbook content link.
	PublishContentLink ContentLinkPtrInput
	// Name of an Azure Resource group.
	ResourceGroupName pulumi.StringInput
	// The runbook name.
	RunbookName pulumi.StringPtrInput
	// Gets or sets the type of the runbook.
	RunbookType pulumi.StringInput
	// Gets or sets the tags attached to the resource.
	Tags pulumi.StringMapInput
}

func (RunbookArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*runbookArgs)(nil)).Elem()
}

type RunbookInput interface {
	pulumi.Input

	ToRunbookOutput() RunbookOutput
	ToRunbookOutputWithContext(ctx context.Context) RunbookOutput
}

func (*Runbook) ElementType() reflect.Type {
	return reflect.TypeOf((**Runbook)(nil)).Elem()
}

func (i *Runbook) ToRunbookOutput() RunbookOutput {
	return i.ToRunbookOutputWithContext(context.Background())
}

func (i *Runbook) ToRunbookOutputWithContext(ctx context.Context) RunbookOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RunbookOutput)
}

type RunbookOutput struct{ *pulumi.OutputState }

func (RunbookOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Runbook)(nil)).Elem()
}

func (o RunbookOutput) ToRunbookOutput() RunbookOutput {
	return o
}

func (o RunbookOutput) ToRunbookOutputWithContext(ctx context.Context) RunbookOutput {
	return o
}

// Gets or sets the creation time.
func (o RunbookOutput) CreationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Runbook) pulumi.StringPtrOutput { return v.CreationTime }).(pulumi.StringPtrOutput)
}

// Gets or sets the description.
func (o RunbookOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Runbook) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Gets or sets the draft runbook properties.
func (o RunbookOutput) Draft() RunbookDraftResponsePtrOutput {
	return o.ApplyT(func(v *Runbook) RunbookDraftResponsePtrOutput { return v.Draft }).(RunbookDraftResponsePtrOutput)
}

// Gets or sets the etag of the resource.
func (o RunbookOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Runbook) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// Gets or sets the job count of the runbook.
func (o RunbookOutput) JobCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Runbook) pulumi.IntPtrOutput { return v.JobCount }).(pulumi.IntPtrOutput)
}

// Gets or sets the last modified by.
func (o RunbookOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Runbook) pulumi.StringPtrOutput { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

// Gets or sets the last modified time.
func (o RunbookOutput) LastModifiedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Runbook) pulumi.StringPtrOutput { return v.LastModifiedTime }).(pulumi.StringPtrOutput)
}

// The Azure Region where the resource lives
func (o RunbookOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Runbook) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Gets or sets the option to log activity trace of the runbook.
func (o RunbookOutput) LogActivityTrace() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Runbook) pulumi.IntPtrOutput { return v.LogActivityTrace }).(pulumi.IntPtrOutput)
}

// Gets or sets progress log option.
func (o RunbookOutput) LogProgress() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Runbook) pulumi.BoolPtrOutput { return v.LogProgress }).(pulumi.BoolPtrOutput)
}

// Gets or sets verbose log option.
func (o RunbookOutput) LogVerbose() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Runbook) pulumi.BoolPtrOutput { return v.LogVerbose }).(pulumi.BoolPtrOutput)
}

// The name of the resource
func (o RunbookOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Runbook) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Gets or sets the runbook output types.
func (o RunbookOutput) OutputTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Runbook) pulumi.StringArrayOutput { return v.OutputTypes }).(pulumi.StringArrayOutput)
}

// Gets or sets the runbook parameters.
func (o RunbookOutput) Parameters() RunbookParameterResponseMapOutput {
	return o.ApplyT(func(v *Runbook) RunbookParameterResponseMapOutput { return v.Parameters }).(RunbookParameterResponseMapOutput)
}

// Gets or sets the provisioning state of the runbook.
func (o RunbookOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Runbook) pulumi.StringPtrOutput { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

// Gets or sets the published runbook content link.
func (o RunbookOutput) PublishContentLink() ContentLinkResponsePtrOutput {
	return o.ApplyT(func(v *Runbook) ContentLinkResponsePtrOutput { return v.PublishContentLink }).(ContentLinkResponsePtrOutput)
}

// Gets or sets the type of the runbook.
func (o RunbookOutput) RunbookType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Runbook) pulumi.StringPtrOutput { return v.RunbookType }).(pulumi.StringPtrOutput)
}

// Gets or sets the state of the runbook.
func (o RunbookOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Runbook) pulumi.StringPtrOutput { return v.State }).(pulumi.StringPtrOutput)
}

// Resource tags.
func (o RunbookOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Runbook) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource.
func (o RunbookOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Runbook) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(RunbookOutput{})
}
