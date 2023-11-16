// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230915

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves the connection info for the notebook workspace
func ListNotebookWorkspaceConnectionInfo(ctx *pulumi.Context, args *ListNotebookWorkspaceConnectionInfoArgs, opts ...pulumi.InvokeOption) (*ListNotebookWorkspaceConnectionInfoResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListNotebookWorkspaceConnectionInfoResult
	err := ctx.Invoke("azure-native:documentdb/v20230915:listNotebookWorkspaceConnectionInfo", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListNotebookWorkspaceConnectionInfoArgs struct {
	// Cosmos DB database account name.
	AccountName string `pulumi:"accountName"`
	// The name of the notebook workspace resource.
	NotebookWorkspaceName string `pulumi:"notebookWorkspaceName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The connection info for the given notebook workspace
type ListNotebookWorkspaceConnectionInfoResult struct {
	// Specifies auth token used for connecting to Notebook server (uses token-based auth).
	AuthToken string `pulumi:"authToken"`
	// Specifies the endpoint of Notebook server.
	NotebookServerEndpoint string `pulumi:"notebookServerEndpoint"`
}

func ListNotebookWorkspaceConnectionInfoOutput(ctx *pulumi.Context, args ListNotebookWorkspaceConnectionInfoOutputArgs, opts ...pulumi.InvokeOption) ListNotebookWorkspaceConnectionInfoResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListNotebookWorkspaceConnectionInfoResult, error) {
			args := v.(ListNotebookWorkspaceConnectionInfoArgs)
			r, err := ListNotebookWorkspaceConnectionInfo(ctx, &args, opts...)
			var s ListNotebookWorkspaceConnectionInfoResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListNotebookWorkspaceConnectionInfoResultOutput)
}

type ListNotebookWorkspaceConnectionInfoOutputArgs struct {
	// Cosmos DB database account name.
	AccountName pulumi.StringInput `pulumi:"accountName"`
	// The name of the notebook workspace resource.
	NotebookWorkspaceName pulumi.StringInput `pulumi:"notebookWorkspaceName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListNotebookWorkspaceConnectionInfoOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListNotebookWorkspaceConnectionInfoArgs)(nil)).Elem()
}

// The connection info for the given notebook workspace
type ListNotebookWorkspaceConnectionInfoResultOutput struct{ *pulumi.OutputState }

func (ListNotebookWorkspaceConnectionInfoResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListNotebookWorkspaceConnectionInfoResult)(nil)).Elem()
}

func (o ListNotebookWorkspaceConnectionInfoResultOutput) ToListNotebookWorkspaceConnectionInfoResultOutput() ListNotebookWorkspaceConnectionInfoResultOutput {
	return o
}

func (o ListNotebookWorkspaceConnectionInfoResultOutput) ToListNotebookWorkspaceConnectionInfoResultOutputWithContext(ctx context.Context) ListNotebookWorkspaceConnectionInfoResultOutput {
	return o
}

// Specifies auth token used for connecting to Notebook server (uses token-based auth).
func (o ListNotebookWorkspaceConnectionInfoResultOutput) AuthToken() pulumi.StringOutput {
	return o.ApplyT(func(v ListNotebookWorkspaceConnectionInfoResult) string { return v.AuthToken }).(pulumi.StringOutput)
}

// Specifies the endpoint of Notebook server.
func (o ListNotebookWorkspaceConnectionInfoResultOutput) NotebookServerEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v ListNotebookWorkspaceConnectionInfoResult) string { return v.NotebookServerEndpoint }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListNotebookWorkspaceConnectionInfoResultOutput{})
}
