// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package avs

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An instance of a script executed by a user - custom or AVS
// Azure REST API version: 2022-05-01. Prior API version in Azure Native 1.x: 2021-06-01.
//
// Other available API versions: 2023-03-01.
type ScriptExecution struct {
	pulumi.CustomResourceState

	// Standard error output stream from the powershell execution
	Errors pulumi.StringArrayOutput `pulumi:"errors"`
	// Error message if the script was able to run, but if the script itself had errors or powershell threw an exception
	FailureReason pulumi.StringPtrOutput `pulumi:"failureReason"`
	// Time the script execution was finished
	FinishedAt pulumi.StringOutput `pulumi:"finishedAt"`
	// Parameters that will be hidden/not visible to ARM, such as passwords and credentials
	HiddenParameters pulumi.ArrayOutput `pulumi:"hiddenParameters"`
	// Standard information out stream from the powershell execution
	Information pulumi.StringArrayOutput `pulumi:"information"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// User-defined dictionary.
	NamedOutputs pulumi.MapOutput `pulumi:"namedOutputs"`
	// Standard output stream from the powershell execution
	Output pulumi.StringArrayOutput `pulumi:"output"`
	// Parameters the script will accept
	Parameters pulumi.ArrayOutput `pulumi:"parameters"`
	// The state of the script execution resource
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Time to live for the resource. If not provided, will be available for 60 days
	Retention pulumi.StringPtrOutput `pulumi:"retention"`
	// A reference to the script cmdlet resource if user is running a AVS script
	ScriptCmdletId pulumi.StringPtrOutput `pulumi:"scriptCmdletId"`
	// Time the script execution was started
	StartedAt pulumi.StringOutput `pulumi:"startedAt"`
	// Time the script execution was submitted
	SubmittedAt pulumi.StringOutput `pulumi:"submittedAt"`
	// Time limit for execution
	Timeout pulumi.StringOutput `pulumi:"timeout"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// Standard warning out stream from the powershell execution
	Warnings pulumi.StringArrayOutput `pulumi:"warnings"`
}

// NewScriptExecution registers a new resource with the given unique name, arguments, and options.
func NewScriptExecution(ctx *pulumi.Context,
	name string, args *ScriptExecutionArgs, opts ...pulumi.ResourceOption) (*ScriptExecution, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PrivateCloudName == nil {
		return nil, errors.New("invalid value for required argument 'PrivateCloudName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Timeout == nil {
		return nil, errors.New("invalid value for required argument 'Timeout'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:avs/v20210601:ScriptExecution"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20211201:ScriptExecution"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20220501:ScriptExecution"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20230301:ScriptExecution"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ScriptExecution
	err := ctx.RegisterResource("azure-native:avs:ScriptExecution", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetScriptExecution gets an existing ScriptExecution resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetScriptExecution(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ScriptExecutionState, opts ...pulumi.ResourceOption) (*ScriptExecution, error) {
	var resource ScriptExecution
	err := ctx.ReadResource("azure-native:avs:ScriptExecution", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ScriptExecution resources.
type scriptExecutionState struct {
}

type ScriptExecutionState struct {
}

func (ScriptExecutionState) ElementType() reflect.Type {
	return reflect.TypeOf((*scriptExecutionState)(nil)).Elem()
}

type scriptExecutionArgs struct {
	// Error message if the script was able to run, but if the script itself had errors or powershell threw an exception
	FailureReason *string `pulumi:"failureReason"`
	// Parameters that will be hidden/not visible to ARM, such as passwords and credentials
	HiddenParameters []interface{} `pulumi:"hiddenParameters"`
	// User-defined dictionary.
	NamedOutputs map[string]interface{} `pulumi:"namedOutputs"`
	// Standard output stream from the powershell execution
	Output []string `pulumi:"output"`
	// Parameters the script will accept
	Parameters []interface{} `pulumi:"parameters"`
	// The name of the private cloud.
	PrivateCloudName string `pulumi:"privateCloudName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Time to live for the resource. If not provided, will be available for 60 days
	Retention *string `pulumi:"retention"`
	// A reference to the script cmdlet resource if user is running a AVS script
	ScriptCmdletId *string `pulumi:"scriptCmdletId"`
	// Name of the user-invoked script execution resource
	ScriptExecutionName *string `pulumi:"scriptExecutionName"`
	// Time limit for execution
	Timeout string `pulumi:"timeout"`
}

// The set of arguments for constructing a ScriptExecution resource.
type ScriptExecutionArgs struct {
	// Error message if the script was able to run, but if the script itself had errors or powershell threw an exception
	FailureReason pulumi.StringPtrInput
	// Parameters that will be hidden/not visible to ARM, such as passwords and credentials
	HiddenParameters pulumi.ArrayInput
	// User-defined dictionary.
	NamedOutputs pulumi.MapInput
	// Standard output stream from the powershell execution
	Output pulumi.StringArrayInput
	// Parameters the script will accept
	Parameters pulumi.ArrayInput
	// The name of the private cloud.
	PrivateCloudName pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Time to live for the resource. If not provided, will be available for 60 days
	Retention pulumi.StringPtrInput
	// A reference to the script cmdlet resource if user is running a AVS script
	ScriptCmdletId pulumi.StringPtrInput
	// Name of the user-invoked script execution resource
	ScriptExecutionName pulumi.StringPtrInput
	// Time limit for execution
	Timeout pulumi.StringInput
}

func (ScriptExecutionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*scriptExecutionArgs)(nil)).Elem()
}

type ScriptExecutionInput interface {
	pulumi.Input

	ToScriptExecutionOutput() ScriptExecutionOutput
	ToScriptExecutionOutputWithContext(ctx context.Context) ScriptExecutionOutput
}

func (*ScriptExecution) ElementType() reflect.Type {
	return reflect.TypeOf((**ScriptExecution)(nil)).Elem()
}

func (i *ScriptExecution) ToScriptExecutionOutput() ScriptExecutionOutput {
	return i.ToScriptExecutionOutputWithContext(context.Background())
}

func (i *ScriptExecution) ToScriptExecutionOutputWithContext(ctx context.Context) ScriptExecutionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScriptExecutionOutput)
}

type ScriptExecutionOutput struct{ *pulumi.OutputState }

func (ScriptExecutionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScriptExecution)(nil)).Elem()
}

func (o ScriptExecutionOutput) ToScriptExecutionOutput() ScriptExecutionOutput {
	return o
}

func (o ScriptExecutionOutput) ToScriptExecutionOutputWithContext(ctx context.Context) ScriptExecutionOutput {
	return o
}

// Standard error output stream from the powershell execution
func (o ScriptExecutionOutput) Errors() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ScriptExecution) pulumi.StringArrayOutput { return v.Errors }).(pulumi.StringArrayOutput)
}

// Error message if the script was able to run, but if the script itself had errors or powershell threw an exception
func (o ScriptExecutionOutput) FailureReason() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScriptExecution) pulumi.StringPtrOutput { return v.FailureReason }).(pulumi.StringPtrOutput)
}

// Time the script execution was finished
func (o ScriptExecutionOutput) FinishedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *ScriptExecution) pulumi.StringOutput { return v.FinishedAt }).(pulumi.StringOutput)
}

// Parameters that will be hidden/not visible to ARM, such as passwords and credentials
func (o ScriptExecutionOutput) HiddenParameters() pulumi.ArrayOutput {
	return o.ApplyT(func(v *ScriptExecution) pulumi.ArrayOutput { return v.HiddenParameters }).(pulumi.ArrayOutput)
}

// Standard information out stream from the powershell execution
func (o ScriptExecutionOutput) Information() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ScriptExecution) pulumi.StringArrayOutput { return v.Information }).(pulumi.StringArrayOutput)
}

// Resource name.
func (o ScriptExecutionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ScriptExecution) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// User-defined dictionary.
func (o ScriptExecutionOutput) NamedOutputs() pulumi.MapOutput {
	return o.ApplyT(func(v *ScriptExecution) pulumi.MapOutput { return v.NamedOutputs }).(pulumi.MapOutput)
}

// Standard output stream from the powershell execution
func (o ScriptExecutionOutput) Output() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ScriptExecution) pulumi.StringArrayOutput { return v.Output }).(pulumi.StringArrayOutput)
}

// Parameters the script will accept
func (o ScriptExecutionOutput) Parameters() pulumi.ArrayOutput {
	return o.ApplyT(func(v *ScriptExecution) pulumi.ArrayOutput { return v.Parameters }).(pulumi.ArrayOutput)
}

// The state of the script execution resource
func (o ScriptExecutionOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ScriptExecution) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Time to live for the resource. If not provided, will be available for 60 days
func (o ScriptExecutionOutput) Retention() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScriptExecution) pulumi.StringPtrOutput { return v.Retention }).(pulumi.StringPtrOutput)
}

// A reference to the script cmdlet resource if user is running a AVS script
func (o ScriptExecutionOutput) ScriptCmdletId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScriptExecution) pulumi.StringPtrOutput { return v.ScriptCmdletId }).(pulumi.StringPtrOutput)
}

// Time the script execution was started
func (o ScriptExecutionOutput) StartedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *ScriptExecution) pulumi.StringOutput { return v.StartedAt }).(pulumi.StringOutput)
}

// Time the script execution was submitted
func (o ScriptExecutionOutput) SubmittedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *ScriptExecution) pulumi.StringOutput { return v.SubmittedAt }).(pulumi.StringOutput)
}

// Time limit for execution
func (o ScriptExecutionOutput) Timeout() pulumi.StringOutput {
	return o.ApplyT(func(v *ScriptExecution) pulumi.StringOutput { return v.Timeout }).(pulumi.StringOutput)
}

// Resource type.
func (o ScriptExecutionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ScriptExecution) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Standard warning out stream from the powershell execution
func (o ScriptExecutionOutput) Warnings() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ScriptExecution) pulumi.StringArrayOutput { return v.Warnings }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ScriptExecutionOutput{})
}
