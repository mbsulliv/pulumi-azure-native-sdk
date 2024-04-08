// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20240401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// List of RP resources which supports pagination.
func ListGitLabSubgroup(ctx *pulumi.Context, args *ListGitLabSubgroupArgs, opts ...pulumi.InvokeOption) (*ListGitLabSubgroupResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListGitLabSubgroupResult
	err := ctx.Invoke("azure-native:security/v20240401:listGitLabSubgroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListGitLabSubgroupArgs struct {
	// The GitLab group fully-qualified name.
	GroupFQName string `pulumi:"groupFQName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The security connector name.
	SecurityConnectorName string `pulumi:"securityConnectorName"`
}

// List of RP resources which supports pagination.
type ListGitLabSubgroupResult struct {
	// Gets or sets next link to scroll over the results.
	NextLink *string `pulumi:"nextLink"`
	// Gets or sets list of resources.
	Value []GitLabGroupResponse `pulumi:"value"`
}

func ListGitLabSubgroupOutput(ctx *pulumi.Context, args ListGitLabSubgroupOutputArgs, opts ...pulumi.InvokeOption) ListGitLabSubgroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListGitLabSubgroupResult, error) {
			args := v.(ListGitLabSubgroupArgs)
			r, err := ListGitLabSubgroup(ctx, &args, opts...)
			var s ListGitLabSubgroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListGitLabSubgroupResultOutput)
}

type ListGitLabSubgroupOutputArgs struct {
	// The GitLab group fully-qualified name.
	GroupFQName pulumi.StringInput `pulumi:"groupFQName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The security connector name.
	SecurityConnectorName pulumi.StringInput `pulumi:"securityConnectorName"`
}

func (ListGitLabSubgroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListGitLabSubgroupArgs)(nil)).Elem()
}

// List of RP resources which supports pagination.
type ListGitLabSubgroupResultOutput struct{ *pulumi.OutputState }

func (ListGitLabSubgroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListGitLabSubgroupResult)(nil)).Elem()
}

func (o ListGitLabSubgroupResultOutput) ToListGitLabSubgroupResultOutput() ListGitLabSubgroupResultOutput {
	return o
}

func (o ListGitLabSubgroupResultOutput) ToListGitLabSubgroupResultOutputWithContext(ctx context.Context) ListGitLabSubgroupResultOutput {
	return o
}

// Gets or sets next link to scroll over the results.
func (o ListGitLabSubgroupResultOutput) NextLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListGitLabSubgroupResult) *string { return v.NextLink }).(pulumi.StringPtrOutput)
}

// Gets or sets list of resources.
func (o ListGitLabSubgroupResultOutput) Value() GitLabGroupResponseArrayOutput {
	return o.ApplyT(func(v ListGitLabSubgroupResult) []GitLabGroupResponse { return v.Value }).(GitLabGroupResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListGitLabSubgroupResultOutput{})
}
