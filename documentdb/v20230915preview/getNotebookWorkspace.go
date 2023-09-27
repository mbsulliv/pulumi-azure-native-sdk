// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230915preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets the notebook workspace for a Cosmos DB account.
func LookupNotebookWorkspace(ctx *pulumi.Context, args *LookupNotebookWorkspaceArgs, opts ...pulumi.InvokeOption) (*LookupNotebookWorkspaceResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupNotebookWorkspaceResult
	err := ctx.Invoke("azure-native:documentdb/v20230915preview:getNotebookWorkspace", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNotebookWorkspaceArgs struct {
	// Cosmos DB database account name.
	AccountName string `pulumi:"accountName"`
	// The name of the notebook workspace resource.
	NotebookWorkspaceName string `pulumi:"notebookWorkspaceName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A notebook workspace resource
type LookupNotebookWorkspaceResult struct {
	// The unique resource identifier of the database account.
	Id string `pulumi:"id"`
	// The name of the database account.
	Name string `pulumi:"name"`
	// Specifies the endpoint of Notebook server.
	NotebookServerEndpoint string `pulumi:"notebookServerEndpoint"`
	// Status of the notebook workspace. Possible values are: Creating, Online, Deleting, Failed, Updating.
	Status string `pulumi:"status"`
	// The type of Azure resource.
	Type string `pulumi:"type"`
}

func LookupNotebookWorkspaceOutput(ctx *pulumi.Context, args LookupNotebookWorkspaceOutputArgs, opts ...pulumi.InvokeOption) LookupNotebookWorkspaceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNotebookWorkspaceResult, error) {
			args := v.(LookupNotebookWorkspaceArgs)
			r, err := LookupNotebookWorkspace(ctx, &args, opts...)
			var s LookupNotebookWorkspaceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNotebookWorkspaceResultOutput)
}

type LookupNotebookWorkspaceOutputArgs struct {
	// Cosmos DB database account name.
	AccountName pulumi.StringInput `pulumi:"accountName"`
	// The name of the notebook workspace resource.
	NotebookWorkspaceName pulumi.StringInput `pulumi:"notebookWorkspaceName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupNotebookWorkspaceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNotebookWorkspaceArgs)(nil)).Elem()
}

// A notebook workspace resource
type LookupNotebookWorkspaceResultOutput struct{ *pulumi.OutputState }

func (LookupNotebookWorkspaceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNotebookWorkspaceResult)(nil)).Elem()
}

func (o LookupNotebookWorkspaceResultOutput) ToLookupNotebookWorkspaceResultOutput() LookupNotebookWorkspaceResultOutput {
	return o
}

func (o LookupNotebookWorkspaceResultOutput) ToLookupNotebookWorkspaceResultOutputWithContext(ctx context.Context) LookupNotebookWorkspaceResultOutput {
	return o
}

func (o LookupNotebookWorkspaceResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupNotebookWorkspaceResult] {
	return pulumix.Output[LookupNotebookWorkspaceResult]{
		OutputState: o.OutputState,
	}
}

// The unique resource identifier of the database account.
func (o LookupNotebookWorkspaceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNotebookWorkspaceResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the database account.
func (o LookupNotebookWorkspaceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNotebookWorkspaceResult) string { return v.Name }).(pulumi.StringOutput)
}

// Specifies the endpoint of Notebook server.
func (o LookupNotebookWorkspaceResultOutput) NotebookServerEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNotebookWorkspaceResult) string { return v.NotebookServerEndpoint }).(pulumi.StringOutput)
}

// Status of the notebook workspace. Possible values are: Creating, Online, Deleting, Failed, Updating.
func (o LookupNotebookWorkspaceResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNotebookWorkspaceResult) string { return v.Status }).(pulumi.StringOutput)
}

// The type of Azure resource.
func (o LookupNotebookWorkspaceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNotebookWorkspaceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNotebookWorkspaceResultOutput{})
}
