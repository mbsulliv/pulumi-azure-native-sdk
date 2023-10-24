// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package avs

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// An instance of a script executed by a user - custom or AVS
// Azure REST API version: 2022-05-01.
//
// Other available API versions: 2023-03-01.
func LookupScriptExecution(ctx *pulumi.Context, args *LookupScriptExecutionArgs, opts ...pulumi.InvokeOption) (*LookupScriptExecutionResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupScriptExecutionResult
	err := ctx.Invoke("azure-native:avs:getScriptExecution", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupScriptExecutionArgs struct {
	// Name of the private cloud
	PrivateCloudName string `pulumi:"privateCloudName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the user-invoked script execution resource
	ScriptExecutionName string `pulumi:"scriptExecutionName"`
}

// An instance of a script executed by a user - custom or AVS
type LookupScriptExecutionResult struct {
	// Standard error output stream from the powershell execution
	Errors []string `pulumi:"errors"`
	// Error message if the script was able to run, but if the script itself had errors or powershell threw an exception
	FailureReason *string `pulumi:"failureReason"`
	// Time the script execution was finished
	FinishedAt string `pulumi:"finishedAt"`
	// Parameters that will be hidden/not visible to ARM, such as passwords and credentials
	HiddenParameters []interface{} `pulumi:"hiddenParameters"`
	// Resource ID.
	Id string `pulumi:"id"`
	// Standard information out stream from the powershell execution
	Information []string `pulumi:"information"`
	// Resource name.
	Name string `pulumi:"name"`
	// User-defined dictionary.
	NamedOutputs map[string]interface{} `pulumi:"namedOutputs"`
	// Standard output stream from the powershell execution
	Output []string `pulumi:"output"`
	// Parameters the script will accept
	Parameters []interface{} `pulumi:"parameters"`
	// The state of the script execution resource
	ProvisioningState string `pulumi:"provisioningState"`
	// Time to live for the resource. If not provided, will be available for 60 days
	Retention *string `pulumi:"retention"`
	// A reference to the script cmdlet resource if user is running a AVS script
	ScriptCmdletId *string `pulumi:"scriptCmdletId"`
	// Time the script execution was started
	StartedAt string `pulumi:"startedAt"`
	// Time the script execution was submitted
	SubmittedAt string `pulumi:"submittedAt"`
	// Time limit for execution
	Timeout string `pulumi:"timeout"`
	// Resource type.
	Type string `pulumi:"type"`
	// Standard warning out stream from the powershell execution
	Warnings []string `pulumi:"warnings"`
}

func LookupScriptExecutionOutput(ctx *pulumi.Context, args LookupScriptExecutionOutputArgs, opts ...pulumi.InvokeOption) LookupScriptExecutionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupScriptExecutionResult, error) {
			args := v.(LookupScriptExecutionArgs)
			r, err := LookupScriptExecution(ctx, &args, opts...)
			var s LookupScriptExecutionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupScriptExecutionResultOutput)
}

type LookupScriptExecutionOutputArgs struct {
	// Name of the private cloud
	PrivateCloudName pulumi.StringInput `pulumi:"privateCloudName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name of the user-invoked script execution resource
	ScriptExecutionName pulumi.StringInput `pulumi:"scriptExecutionName"`
}

func (LookupScriptExecutionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupScriptExecutionArgs)(nil)).Elem()
}

// An instance of a script executed by a user - custom or AVS
type LookupScriptExecutionResultOutput struct{ *pulumi.OutputState }

func (LookupScriptExecutionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupScriptExecutionResult)(nil)).Elem()
}

func (o LookupScriptExecutionResultOutput) ToLookupScriptExecutionResultOutput() LookupScriptExecutionResultOutput {
	return o
}

func (o LookupScriptExecutionResultOutput) ToLookupScriptExecutionResultOutputWithContext(ctx context.Context) LookupScriptExecutionResultOutput {
	return o
}

func (o LookupScriptExecutionResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupScriptExecutionResult] {
	return pulumix.Output[LookupScriptExecutionResult]{
		OutputState: o.OutputState,
	}
}

// Standard error output stream from the powershell execution
func (o LookupScriptExecutionResultOutput) Errors() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupScriptExecutionResult) []string { return v.Errors }).(pulumi.StringArrayOutput)
}

// Error message if the script was able to run, but if the script itself had errors or powershell threw an exception
func (o LookupScriptExecutionResultOutput) FailureReason() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScriptExecutionResult) *string { return v.FailureReason }).(pulumi.StringPtrOutput)
}

// Time the script execution was finished
func (o LookupScriptExecutionResultOutput) FinishedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScriptExecutionResult) string { return v.FinishedAt }).(pulumi.StringOutput)
}

// Parameters that will be hidden/not visible to ARM, such as passwords and credentials
func (o LookupScriptExecutionResultOutput) HiddenParameters() pulumi.ArrayOutput {
	return o.ApplyT(func(v LookupScriptExecutionResult) []interface{} { return v.HiddenParameters }).(pulumi.ArrayOutput)
}

// Resource ID.
func (o LookupScriptExecutionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScriptExecutionResult) string { return v.Id }).(pulumi.StringOutput)
}

// Standard information out stream from the powershell execution
func (o LookupScriptExecutionResultOutput) Information() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupScriptExecutionResult) []string { return v.Information }).(pulumi.StringArrayOutput)
}

// Resource name.
func (o LookupScriptExecutionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScriptExecutionResult) string { return v.Name }).(pulumi.StringOutput)
}

// User-defined dictionary.
func (o LookupScriptExecutionResultOutput) NamedOutputs() pulumi.MapOutput {
	return o.ApplyT(func(v LookupScriptExecutionResult) map[string]interface{} { return v.NamedOutputs }).(pulumi.MapOutput)
}

// Standard output stream from the powershell execution
func (o LookupScriptExecutionResultOutput) Output() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupScriptExecutionResult) []string { return v.Output }).(pulumi.StringArrayOutput)
}

// Parameters the script will accept
func (o LookupScriptExecutionResultOutput) Parameters() pulumi.ArrayOutput {
	return o.ApplyT(func(v LookupScriptExecutionResult) []interface{} { return v.Parameters }).(pulumi.ArrayOutput)
}

// The state of the script execution resource
func (o LookupScriptExecutionResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScriptExecutionResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Time to live for the resource. If not provided, will be available for 60 days
func (o LookupScriptExecutionResultOutput) Retention() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScriptExecutionResult) *string { return v.Retention }).(pulumi.StringPtrOutput)
}

// A reference to the script cmdlet resource if user is running a AVS script
func (o LookupScriptExecutionResultOutput) ScriptCmdletId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScriptExecutionResult) *string { return v.ScriptCmdletId }).(pulumi.StringPtrOutput)
}

// Time the script execution was started
func (o LookupScriptExecutionResultOutput) StartedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScriptExecutionResult) string { return v.StartedAt }).(pulumi.StringOutput)
}

// Time the script execution was submitted
func (o LookupScriptExecutionResultOutput) SubmittedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScriptExecutionResult) string { return v.SubmittedAt }).(pulumi.StringOutput)
}

// Time limit for execution
func (o LookupScriptExecutionResultOutput) Timeout() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScriptExecutionResult) string { return v.Timeout }).(pulumi.StringOutput)
}

// Resource type.
func (o LookupScriptExecutionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScriptExecutionResult) string { return v.Type }).(pulumi.StringOutput)
}

// Standard warning out stream from the powershell execution
func (o LookupScriptExecutionResultOutput) Warnings() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupScriptExecutionResult) []string { return v.Warnings }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupScriptExecutionResultOutput{})
}
