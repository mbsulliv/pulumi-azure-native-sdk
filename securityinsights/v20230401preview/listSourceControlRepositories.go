// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a list of repositories metadata.
func ListSourceControlRepositories(ctx *pulumi.Context, args *ListSourceControlRepositoriesArgs, opts ...pulumi.InvokeOption) (*ListSourceControlRepositoriesResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListSourceControlRepositoriesResult
	err := ctx.Invoke("azure-native:securityinsights/v20230401preview:listSourceControlRepositories", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListSourceControlRepositoriesArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// List all the source controls.
type ListSourceControlRepositoriesResult struct {
	// URL to fetch the next set of repositories.
	NextLink string `pulumi:"nextLink"`
	// Array of repositories.
	Value []RepoResponse `pulumi:"value"`
}

func ListSourceControlRepositoriesOutput(ctx *pulumi.Context, args ListSourceControlRepositoriesOutputArgs, opts ...pulumi.InvokeOption) ListSourceControlRepositoriesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListSourceControlRepositoriesResult, error) {
			args := v.(ListSourceControlRepositoriesArgs)
			r, err := ListSourceControlRepositories(ctx, &args, opts...)
			var s ListSourceControlRepositoriesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListSourceControlRepositoriesResultOutput)
}

type ListSourceControlRepositoriesOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (ListSourceControlRepositoriesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListSourceControlRepositoriesArgs)(nil)).Elem()
}

// List all the source controls.
type ListSourceControlRepositoriesResultOutput struct{ *pulumi.OutputState }

func (ListSourceControlRepositoriesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListSourceControlRepositoriesResult)(nil)).Elem()
}

func (o ListSourceControlRepositoriesResultOutput) ToListSourceControlRepositoriesResultOutput() ListSourceControlRepositoriesResultOutput {
	return o
}

func (o ListSourceControlRepositoriesResultOutput) ToListSourceControlRepositoriesResultOutputWithContext(ctx context.Context) ListSourceControlRepositoriesResultOutput {
	return o
}

// URL to fetch the next set of repositories.
func (o ListSourceControlRepositoriesResultOutput) NextLink() pulumi.StringOutput {
	return o.ApplyT(func(v ListSourceControlRepositoriesResult) string { return v.NextLink }).(pulumi.StringOutput)
}

// Array of repositories.
func (o ListSourceControlRepositoriesResultOutput) Value() RepoResponseArrayOutput {
	return o.ApplyT(func(v ListSourceControlRepositoriesResult) []RepoResponse { return v.Value }).(RepoResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListSourceControlRepositoriesResultOutput{})
}
