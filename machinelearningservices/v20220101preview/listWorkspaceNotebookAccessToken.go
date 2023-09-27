// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// return notebook access token and refresh token
func ListWorkspaceNotebookAccessToken(ctx *pulumi.Context, args *ListWorkspaceNotebookAccessTokenArgs, opts ...pulumi.InvokeOption) (*ListWorkspaceNotebookAccessTokenResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListWorkspaceNotebookAccessTokenResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20220101preview:listWorkspaceNotebookAccessToken", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWorkspaceNotebookAccessTokenArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

type ListWorkspaceNotebookAccessTokenResult struct {
	AccessToken        string `pulumi:"accessToken"`
	ExpiresIn          int    `pulumi:"expiresIn"`
	HostName           string `pulumi:"hostName"`
	NotebookResourceId string `pulumi:"notebookResourceId"`
	PublicDns          string `pulumi:"publicDns"`
	RefreshToken       string `pulumi:"refreshToken"`
	Scope              string `pulumi:"scope"`
	TokenType          string `pulumi:"tokenType"`
}

func ListWorkspaceNotebookAccessTokenOutput(ctx *pulumi.Context, args ListWorkspaceNotebookAccessTokenOutputArgs, opts ...pulumi.InvokeOption) ListWorkspaceNotebookAccessTokenResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListWorkspaceNotebookAccessTokenResult, error) {
			args := v.(ListWorkspaceNotebookAccessTokenArgs)
			r, err := ListWorkspaceNotebookAccessToken(ctx, &args, opts...)
			var s ListWorkspaceNotebookAccessTokenResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListWorkspaceNotebookAccessTokenResultOutput)
}

type ListWorkspaceNotebookAccessTokenOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (ListWorkspaceNotebookAccessTokenOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWorkspaceNotebookAccessTokenArgs)(nil)).Elem()
}

type ListWorkspaceNotebookAccessTokenResultOutput struct{ *pulumi.OutputState }

func (ListWorkspaceNotebookAccessTokenResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWorkspaceNotebookAccessTokenResult)(nil)).Elem()
}

func (o ListWorkspaceNotebookAccessTokenResultOutput) ToListWorkspaceNotebookAccessTokenResultOutput() ListWorkspaceNotebookAccessTokenResultOutput {
	return o
}

func (o ListWorkspaceNotebookAccessTokenResultOutput) ToListWorkspaceNotebookAccessTokenResultOutputWithContext(ctx context.Context) ListWorkspaceNotebookAccessTokenResultOutput {
	return o
}

func (o ListWorkspaceNotebookAccessTokenResultOutput) ToOutput(ctx context.Context) pulumix.Output[ListWorkspaceNotebookAccessTokenResult] {
	return pulumix.Output[ListWorkspaceNotebookAccessTokenResult]{
		OutputState: o.OutputState,
	}
}

func (o ListWorkspaceNotebookAccessTokenResultOutput) AccessToken() pulumi.StringOutput {
	return o.ApplyT(func(v ListWorkspaceNotebookAccessTokenResult) string { return v.AccessToken }).(pulumi.StringOutput)
}

func (o ListWorkspaceNotebookAccessTokenResultOutput) ExpiresIn() pulumi.IntOutput {
	return o.ApplyT(func(v ListWorkspaceNotebookAccessTokenResult) int { return v.ExpiresIn }).(pulumi.IntOutput)
}

func (o ListWorkspaceNotebookAccessTokenResultOutput) HostName() pulumi.StringOutput {
	return o.ApplyT(func(v ListWorkspaceNotebookAccessTokenResult) string { return v.HostName }).(pulumi.StringOutput)
}

func (o ListWorkspaceNotebookAccessTokenResultOutput) NotebookResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v ListWorkspaceNotebookAccessTokenResult) string { return v.NotebookResourceId }).(pulumi.StringOutput)
}

func (o ListWorkspaceNotebookAccessTokenResultOutput) PublicDns() pulumi.StringOutput {
	return o.ApplyT(func(v ListWorkspaceNotebookAccessTokenResult) string { return v.PublicDns }).(pulumi.StringOutput)
}

func (o ListWorkspaceNotebookAccessTokenResultOutput) RefreshToken() pulumi.StringOutput {
	return o.ApplyT(func(v ListWorkspaceNotebookAccessTokenResult) string { return v.RefreshToken }).(pulumi.StringOutput)
}

func (o ListWorkspaceNotebookAccessTokenResultOutput) Scope() pulumi.StringOutput {
	return o.ApplyT(func(v ListWorkspaceNotebookAccessTokenResult) string { return v.Scope }).(pulumi.StringOutput)
}

func (o ListWorkspaceNotebookAccessTokenResultOutput) TokenType() pulumi.StringOutput {
	return o.ApplyT(func(v ListWorkspaceNotebookAccessTokenResult) string { return v.TokenType }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListWorkspaceNotebookAccessTokenResultOutput{})
}
