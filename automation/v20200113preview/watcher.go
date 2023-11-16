// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20200113preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of the watcher type.
type Watcher struct {
	pulumi.CustomResourceState

	// Gets or sets the creation time.
	CreationTime pulumi.StringOutput `pulumi:"creationTime"`
	// Gets or sets the description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Gets or sets the etag of the resource.
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// Gets or sets the frequency at which the watcher is invoked.
	ExecutionFrequencyInSeconds pulumi.Float64PtrOutput `pulumi:"executionFrequencyInSeconds"`
	// Details of the user who last modified the watcher.
	LastModifiedBy pulumi.StringOutput `pulumi:"lastModifiedBy"`
	// Gets or sets the last modified time.
	LastModifiedTime pulumi.StringOutput `pulumi:"lastModifiedTime"`
	// The geo-location where the resource lives
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Gets or sets the name of the script the watcher is attached to, i.e. the name of an existing runbook.
	ScriptName pulumi.StringPtrOutput `pulumi:"scriptName"`
	// Gets or sets the parameters of the script.
	ScriptParameters pulumi.StringMapOutput `pulumi:"scriptParameters"`
	// Gets or sets the name of the hybrid worker group the watcher will run on.
	ScriptRunOn pulumi.StringPtrOutput `pulumi:"scriptRunOn"`
	// Gets the current status of the watcher.
	Status pulumi.StringOutput `pulumi:"status"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewWatcher registers a new resource with the given unique name, arguments, and options.
func NewWatcher(ctx *pulumi.Context,
	name string, args *WatcherArgs, opts ...pulumi.ResourceOption) (*Watcher, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AutomationAccountName == nil {
		return nil, errors.New("invalid value for required argument 'AutomationAccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:automation:Watcher"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20151031:Watcher"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20190601:Watcher"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20230515preview:Watcher"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Watcher
	err := ctx.RegisterResource("azure-native:automation/v20200113preview:Watcher", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWatcher gets an existing Watcher resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWatcher(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WatcherState, opts ...pulumi.ResourceOption) (*Watcher, error) {
	var resource Watcher
	err := ctx.ReadResource("azure-native:automation/v20200113preview:Watcher", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Watcher resources.
type watcherState struct {
}

type WatcherState struct {
}

func (WatcherState) ElementType() reflect.Type {
	return reflect.TypeOf((*watcherState)(nil)).Elem()
}

type watcherArgs struct {
	// The name of the automation account.
	AutomationAccountName string `pulumi:"automationAccountName"`
	// Gets or sets the description.
	Description *string `pulumi:"description"`
	// Gets or sets the frequency at which the watcher is invoked.
	ExecutionFrequencyInSeconds *float64 `pulumi:"executionFrequencyInSeconds"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// Name of an Azure Resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Gets or sets the name of the script the watcher is attached to, i.e. the name of an existing runbook.
	ScriptName *string `pulumi:"scriptName"`
	// Gets or sets the parameters of the script.
	ScriptParameters map[string]string `pulumi:"scriptParameters"`
	// Gets or sets the name of the hybrid worker group the watcher will run on.
	ScriptRunOn *string `pulumi:"scriptRunOn"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The watcher name.
	WatcherName *string `pulumi:"watcherName"`
}

// The set of arguments for constructing a Watcher resource.
type WatcherArgs struct {
	// The name of the automation account.
	AutomationAccountName pulumi.StringInput
	// Gets or sets the description.
	Description pulumi.StringPtrInput
	// Gets or sets the frequency at which the watcher is invoked.
	ExecutionFrequencyInSeconds pulumi.Float64PtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// Name of an Azure Resource group.
	ResourceGroupName pulumi.StringInput
	// Gets or sets the name of the script the watcher is attached to, i.e. the name of an existing runbook.
	ScriptName pulumi.StringPtrInput
	// Gets or sets the parameters of the script.
	ScriptParameters pulumi.StringMapInput
	// Gets or sets the name of the hybrid worker group the watcher will run on.
	ScriptRunOn pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The watcher name.
	WatcherName pulumi.StringPtrInput
}

func (WatcherArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*watcherArgs)(nil)).Elem()
}

type WatcherInput interface {
	pulumi.Input

	ToWatcherOutput() WatcherOutput
	ToWatcherOutputWithContext(ctx context.Context) WatcherOutput
}

func (*Watcher) ElementType() reflect.Type {
	return reflect.TypeOf((**Watcher)(nil)).Elem()
}

func (i *Watcher) ToWatcherOutput() WatcherOutput {
	return i.ToWatcherOutputWithContext(context.Background())
}

func (i *Watcher) ToWatcherOutputWithContext(ctx context.Context) WatcherOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WatcherOutput)
}

type WatcherOutput struct{ *pulumi.OutputState }

func (WatcherOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Watcher)(nil)).Elem()
}

func (o WatcherOutput) ToWatcherOutput() WatcherOutput {
	return o
}

func (o WatcherOutput) ToWatcherOutputWithContext(ctx context.Context) WatcherOutput {
	return o
}

// Gets or sets the creation time.
func (o WatcherOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Watcher) pulumi.StringOutput { return v.CreationTime }).(pulumi.StringOutput)
}

// Gets or sets the description.
func (o WatcherOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Watcher) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Gets or sets the etag of the resource.
func (o WatcherOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Watcher) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// Gets or sets the frequency at which the watcher is invoked.
func (o WatcherOutput) ExecutionFrequencyInSeconds() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *Watcher) pulumi.Float64PtrOutput { return v.ExecutionFrequencyInSeconds }).(pulumi.Float64PtrOutput)
}

// Details of the user who last modified the watcher.
func (o WatcherOutput) LastModifiedBy() pulumi.StringOutput {
	return o.ApplyT(func(v *Watcher) pulumi.StringOutput { return v.LastModifiedBy }).(pulumi.StringOutput)
}

// Gets or sets the last modified time.
func (o WatcherOutput) LastModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Watcher) pulumi.StringOutput { return v.LastModifiedTime }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o WatcherOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Watcher) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o WatcherOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Watcher) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Gets or sets the name of the script the watcher is attached to, i.e. the name of an existing runbook.
func (o WatcherOutput) ScriptName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Watcher) pulumi.StringPtrOutput { return v.ScriptName }).(pulumi.StringPtrOutput)
}

// Gets or sets the parameters of the script.
func (o WatcherOutput) ScriptParameters() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Watcher) pulumi.StringMapOutput { return v.ScriptParameters }).(pulumi.StringMapOutput)
}

// Gets or sets the name of the hybrid worker group the watcher will run on.
func (o WatcherOutput) ScriptRunOn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Watcher) pulumi.StringPtrOutput { return v.ScriptRunOn }).(pulumi.StringPtrOutput)
}

// Gets the current status of the watcher.
func (o WatcherOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Watcher) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Resource tags.
func (o WatcherOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Watcher) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource.
func (o WatcherOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Watcher) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WatcherOutput{})
}
