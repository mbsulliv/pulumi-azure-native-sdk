// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package automation

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieve the source control identified by source control name.
// Azure REST API version: 2022-08-08.
//
// Other available API versions: 2023-05-15-preview, 2023-11-01.
func LookupSourceControl(ctx *pulumi.Context, args *LookupSourceControlArgs, opts ...pulumi.InvokeOption) (*LookupSourceControlResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupSourceControlResult
	err := ctx.Invoke("azure-native:automation:getSourceControl", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSourceControlArgs struct {
	// The name of the automation account.
	AutomationAccountName string `pulumi:"automationAccountName"`
	// Name of an Azure Resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of source control.
	SourceControlName string `pulumi:"sourceControlName"`
}

// Definition of the source control.
type LookupSourceControlResult struct {
	// The auto sync of the source control. Default is false.
	AutoSync *bool `pulumi:"autoSync"`
	// The repo branch of the source control. Include branch as empty string for VsoTfvc.
	Branch *string `pulumi:"branch"`
	// The creation time.
	CreationTime *string `pulumi:"creationTime"`
	// The description.
	Description *string `pulumi:"description"`
	// The folder path of the source control.
	FolderPath *string `pulumi:"folderPath"`
	// Fully qualified resource Id for the resource
	Id string `pulumi:"id"`
	// The last modified time.
	LastModifiedTime *string `pulumi:"lastModifiedTime"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The auto publish of the source control. Default is true.
	PublishRunbook *bool `pulumi:"publishRunbook"`
	// The repo url of the source control.
	RepoUrl *string `pulumi:"repoUrl"`
	// The source type. Must be one of VsoGit, VsoTfvc, GitHub.
	SourceType *string `pulumi:"sourceType"`
	// The type of the resource.
	Type string `pulumi:"type"`
}

func LookupSourceControlOutput(ctx *pulumi.Context, args LookupSourceControlOutputArgs, opts ...pulumi.InvokeOption) LookupSourceControlResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceControlResult, error) {
			args := v.(LookupSourceControlArgs)
			r, err := LookupSourceControl(ctx, &args, opts...)
			var s LookupSourceControlResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceControlResultOutput)
}

type LookupSourceControlOutputArgs struct {
	// The name of the automation account.
	AutomationAccountName pulumi.StringInput `pulumi:"automationAccountName"`
	// Name of an Azure Resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of source control.
	SourceControlName pulumi.StringInput `pulumi:"sourceControlName"`
}

func (LookupSourceControlOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceControlArgs)(nil)).Elem()
}

// Definition of the source control.
type LookupSourceControlResultOutput struct{ *pulumi.OutputState }

func (LookupSourceControlResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceControlResult)(nil)).Elem()
}

func (o LookupSourceControlResultOutput) ToLookupSourceControlResultOutput() LookupSourceControlResultOutput {
	return o
}

func (o LookupSourceControlResultOutput) ToLookupSourceControlResultOutputWithContext(ctx context.Context) LookupSourceControlResultOutput {
	return o
}

// The auto sync of the source control. Default is false.
func (o LookupSourceControlResultOutput) AutoSync() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupSourceControlResult) *bool { return v.AutoSync }).(pulumi.BoolPtrOutput)
}

// The repo branch of the source control. Include branch as empty string for VsoTfvc.
func (o LookupSourceControlResultOutput) Branch() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceControlResult) *string { return v.Branch }).(pulumi.StringPtrOutput)
}

// The creation time.
func (o LookupSourceControlResultOutput) CreationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceControlResult) *string { return v.CreationTime }).(pulumi.StringPtrOutput)
}

// The description.
func (o LookupSourceControlResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceControlResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The folder path of the source control.
func (o LookupSourceControlResultOutput) FolderPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceControlResult) *string { return v.FolderPath }).(pulumi.StringPtrOutput)
}

// Fully qualified resource Id for the resource
func (o LookupSourceControlResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceControlResult) string { return v.Id }).(pulumi.StringOutput)
}

// The last modified time.
func (o LookupSourceControlResultOutput) LastModifiedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceControlResult) *string { return v.LastModifiedTime }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o LookupSourceControlResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceControlResult) string { return v.Name }).(pulumi.StringOutput)
}

// The auto publish of the source control. Default is true.
func (o LookupSourceControlResultOutput) PublishRunbook() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupSourceControlResult) *bool { return v.PublishRunbook }).(pulumi.BoolPtrOutput)
}

// The repo url of the source control.
func (o LookupSourceControlResultOutput) RepoUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceControlResult) *string { return v.RepoUrl }).(pulumi.StringPtrOutput)
}

// The source type. Must be one of VsoGit, VsoTfvc, GitHub.
func (o LookupSourceControlResultOutput) SourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceControlResult) *string { return v.SourceType }).(pulumi.StringPtrOutput)
}

// The type of the resource.
func (o LookupSourceControlResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceControlResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceControlResultOutput{})
}
