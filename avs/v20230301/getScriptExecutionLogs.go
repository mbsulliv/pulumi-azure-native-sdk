// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Return the logs for a script execution resource
func GetScriptExecutionLogs(ctx *pulumi.Context, args *GetScriptExecutionLogsArgs, opts ...pulumi.InvokeOption) (*GetScriptExecutionLogsResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv GetScriptExecutionLogsResult
	err := ctx.Invoke("azure-native:avs/v20230301:getScriptExecutionLogs", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetScriptExecutionLogsArgs struct {
	// Name of the private cloud
	PrivateCloudName string `pulumi:"privateCloudName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the user-invoked script execution resource
	ScriptExecutionName string `pulumi:"scriptExecutionName"`
}

// An instance of a script executed by a user - custom or AVS
type GetScriptExecutionLogsResult struct {
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

func GetScriptExecutionLogsOutput(ctx *pulumi.Context, args GetScriptExecutionLogsOutputArgs, opts ...pulumi.InvokeOption) GetScriptExecutionLogsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetScriptExecutionLogsResult, error) {
			args := v.(GetScriptExecutionLogsArgs)
			r, err := GetScriptExecutionLogs(ctx, &args, opts...)
			var s GetScriptExecutionLogsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetScriptExecutionLogsResultOutput)
}

type GetScriptExecutionLogsOutputArgs struct {
	// Name of the private cloud
	PrivateCloudName pulumi.StringInput `pulumi:"privateCloudName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name of the user-invoked script execution resource
	ScriptExecutionName pulumi.StringInput `pulumi:"scriptExecutionName"`
}

func (GetScriptExecutionLogsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetScriptExecutionLogsArgs)(nil)).Elem()
}

// An instance of a script executed by a user - custom or AVS
type GetScriptExecutionLogsResultOutput struct{ *pulumi.OutputState }

func (GetScriptExecutionLogsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetScriptExecutionLogsResult)(nil)).Elem()
}

func (o GetScriptExecutionLogsResultOutput) ToGetScriptExecutionLogsResultOutput() GetScriptExecutionLogsResultOutput {
	return o
}

func (o GetScriptExecutionLogsResultOutput) ToGetScriptExecutionLogsResultOutputWithContext(ctx context.Context) GetScriptExecutionLogsResultOutput {
	return o
}

// Standard error output stream from the powershell execution
func (o GetScriptExecutionLogsResultOutput) Errors() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetScriptExecutionLogsResult) []string { return v.Errors }).(pulumi.StringArrayOutput)
}

// Error message if the script was able to run, but if the script itself had errors or powershell threw an exception
func (o GetScriptExecutionLogsResultOutput) FailureReason() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetScriptExecutionLogsResult) *string { return v.FailureReason }).(pulumi.StringPtrOutput)
}

// Time the script execution was finished
func (o GetScriptExecutionLogsResultOutput) FinishedAt() pulumi.StringOutput {
	return o.ApplyT(func(v GetScriptExecutionLogsResult) string { return v.FinishedAt }).(pulumi.StringOutput)
}

// Parameters that will be hidden/not visible to ARM, such as passwords and credentials
func (o GetScriptExecutionLogsResultOutput) HiddenParameters() pulumi.ArrayOutput {
	return o.ApplyT(func(v GetScriptExecutionLogsResult) []interface{} { return v.HiddenParameters }).(pulumi.ArrayOutput)
}

// Resource ID.
func (o GetScriptExecutionLogsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetScriptExecutionLogsResult) string { return v.Id }).(pulumi.StringOutput)
}

// Standard information out stream from the powershell execution
func (o GetScriptExecutionLogsResultOutput) Information() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetScriptExecutionLogsResult) []string { return v.Information }).(pulumi.StringArrayOutput)
}

// Resource name.
func (o GetScriptExecutionLogsResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetScriptExecutionLogsResult) string { return v.Name }).(pulumi.StringOutput)
}

// User-defined dictionary.
func (o GetScriptExecutionLogsResultOutput) NamedOutputs() pulumi.MapOutput {
	return o.ApplyT(func(v GetScriptExecutionLogsResult) map[string]interface{} { return v.NamedOutputs }).(pulumi.MapOutput)
}

// Standard output stream from the powershell execution
func (o GetScriptExecutionLogsResultOutput) Output() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetScriptExecutionLogsResult) []string { return v.Output }).(pulumi.StringArrayOutput)
}

// Parameters the script will accept
func (o GetScriptExecutionLogsResultOutput) Parameters() pulumi.ArrayOutput {
	return o.ApplyT(func(v GetScriptExecutionLogsResult) []interface{} { return v.Parameters }).(pulumi.ArrayOutput)
}

// The state of the script execution resource
func (o GetScriptExecutionLogsResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v GetScriptExecutionLogsResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Time to live for the resource. If not provided, will be available for 60 days
func (o GetScriptExecutionLogsResultOutput) Retention() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetScriptExecutionLogsResult) *string { return v.Retention }).(pulumi.StringPtrOutput)
}

// A reference to the script cmdlet resource if user is running a AVS script
func (o GetScriptExecutionLogsResultOutput) ScriptCmdletId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetScriptExecutionLogsResult) *string { return v.ScriptCmdletId }).(pulumi.StringPtrOutput)
}

// Time the script execution was started
func (o GetScriptExecutionLogsResultOutput) StartedAt() pulumi.StringOutput {
	return o.ApplyT(func(v GetScriptExecutionLogsResult) string { return v.StartedAt }).(pulumi.StringOutput)
}

// Time the script execution was submitted
func (o GetScriptExecutionLogsResultOutput) SubmittedAt() pulumi.StringOutput {
	return o.ApplyT(func(v GetScriptExecutionLogsResult) string { return v.SubmittedAt }).(pulumi.StringOutput)
}

// Time limit for execution
func (o GetScriptExecutionLogsResultOutput) Timeout() pulumi.StringOutput {
	return o.ApplyT(func(v GetScriptExecutionLogsResult) string { return v.Timeout }).(pulumi.StringOutput)
}

// Resource type.
func (o GetScriptExecutionLogsResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GetScriptExecutionLogsResult) string { return v.Type }).(pulumi.StringOutput)
}

// Standard warning out stream from the powershell execution
func (o GetScriptExecutionLogsResultOutput) Warnings() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetScriptExecutionLogsResult) []string { return v.Warnings }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetScriptExecutionLogsResultOutput{})
}