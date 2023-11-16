// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package automation

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieve the watcher identified by watcher name.
// Azure REST API version: 2020-01-13-preview.
//
// Other available API versions: 2023-05-15-preview.
func LookupWatcher(ctx *pulumi.Context, args *LookupWatcherArgs, opts ...pulumi.InvokeOption) (*LookupWatcherResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupWatcherResult
	err := ctx.Invoke("azure-native:automation:getWatcher", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWatcherArgs struct {
	// The name of the automation account.
	AutomationAccountName string `pulumi:"automationAccountName"`
	// Name of an Azure Resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The watcher name.
	WatcherName string `pulumi:"watcherName"`
}

// Definition of the watcher type.
type LookupWatcherResult struct {
	// Gets or sets the creation time.
	CreationTime string `pulumi:"creationTime"`
	// Gets or sets the description.
	Description *string `pulumi:"description"`
	// Gets or sets the etag of the resource.
	Etag *string `pulumi:"etag"`
	// Gets or sets the frequency at which the watcher is invoked.
	ExecutionFrequencyInSeconds *float64 `pulumi:"executionFrequencyInSeconds"`
	// Fully qualified resource Id for the resource
	Id string `pulumi:"id"`
	// Details of the user who last modified the watcher.
	LastModifiedBy string `pulumi:"lastModifiedBy"`
	// Gets or sets the last modified time.
	LastModifiedTime string `pulumi:"lastModifiedTime"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Gets or sets the name of the script the watcher is attached to, i.e. the name of an existing runbook.
	ScriptName *string `pulumi:"scriptName"`
	// Gets or sets the parameters of the script.
	ScriptParameters map[string]string `pulumi:"scriptParameters"`
	// Gets or sets the name of the hybrid worker group the watcher will run on.
	ScriptRunOn *string `pulumi:"scriptRunOn"`
	// Gets the current status of the watcher.
	Status string `pulumi:"status"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource.
	Type string `pulumi:"type"`
}

func LookupWatcherOutput(ctx *pulumi.Context, args LookupWatcherOutputArgs, opts ...pulumi.InvokeOption) LookupWatcherResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWatcherResult, error) {
			args := v.(LookupWatcherArgs)
			r, err := LookupWatcher(ctx, &args, opts...)
			var s LookupWatcherResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWatcherResultOutput)
}

type LookupWatcherOutputArgs struct {
	// The name of the automation account.
	AutomationAccountName pulumi.StringInput `pulumi:"automationAccountName"`
	// Name of an Azure Resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The watcher name.
	WatcherName pulumi.StringInput `pulumi:"watcherName"`
}

func (LookupWatcherOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWatcherArgs)(nil)).Elem()
}

// Definition of the watcher type.
type LookupWatcherResultOutput struct{ *pulumi.OutputState }

func (LookupWatcherResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWatcherResult)(nil)).Elem()
}

func (o LookupWatcherResultOutput) ToLookupWatcherResultOutput() LookupWatcherResultOutput {
	return o
}

func (o LookupWatcherResultOutput) ToLookupWatcherResultOutputWithContext(ctx context.Context) LookupWatcherResultOutput {
	return o
}

// Gets or sets the creation time.
func (o LookupWatcherResultOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWatcherResult) string { return v.CreationTime }).(pulumi.StringOutput)
}

// Gets or sets the description.
func (o LookupWatcherResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWatcherResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Gets or sets the etag of the resource.
func (o LookupWatcherResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWatcherResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// Gets or sets the frequency at which the watcher is invoked.
func (o LookupWatcherResultOutput) ExecutionFrequencyInSeconds() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupWatcherResult) *float64 { return v.ExecutionFrequencyInSeconds }).(pulumi.Float64PtrOutput)
}

// Fully qualified resource Id for the resource
func (o LookupWatcherResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWatcherResult) string { return v.Id }).(pulumi.StringOutput)
}

// Details of the user who last modified the watcher.
func (o LookupWatcherResultOutput) LastModifiedBy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWatcherResult) string { return v.LastModifiedBy }).(pulumi.StringOutput)
}

// Gets or sets the last modified time.
func (o LookupWatcherResultOutput) LastModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWatcherResult) string { return v.LastModifiedTime }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupWatcherResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWatcherResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o LookupWatcherResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWatcherResult) string { return v.Name }).(pulumi.StringOutput)
}

// Gets or sets the name of the script the watcher is attached to, i.e. the name of an existing runbook.
func (o LookupWatcherResultOutput) ScriptName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWatcherResult) *string { return v.ScriptName }).(pulumi.StringPtrOutput)
}

// Gets or sets the parameters of the script.
func (o LookupWatcherResultOutput) ScriptParameters() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupWatcherResult) map[string]string { return v.ScriptParameters }).(pulumi.StringMapOutput)
}

// Gets or sets the name of the hybrid worker group the watcher will run on.
func (o LookupWatcherResultOutput) ScriptRunOn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWatcherResult) *string { return v.ScriptRunOn }).(pulumi.StringPtrOutput)
}

// Gets the current status of the watcher.
func (o LookupWatcherResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWatcherResult) string { return v.Status }).(pulumi.StringOutput)
}

// Resource tags.
func (o LookupWatcherResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupWatcherResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource.
func (o LookupWatcherResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWatcherResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWatcherResultOutput{})
}
