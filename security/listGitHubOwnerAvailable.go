// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package security

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// List of RP resources which supports pagination.
// Azure REST API version: 2023-09-01-preview.
func ListGitHubOwnerAvailable(ctx *pulumi.Context, args *ListGitHubOwnerAvailableArgs, opts ...pulumi.InvokeOption) (*ListGitHubOwnerAvailableResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListGitHubOwnerAvailableResult
	err := ctx.Invoke("azure-native:security:listGitHubOwnerAvailable", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListGitHubOwnerAvailableArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The security connector name.
	SecurityConnectorName string `pulumi:"securityConnectorName"`
}

// List of RP resources which supports pagination.
type ListGitHubOwnerAvailableResult struct {
	// Gets or sets next link to scroll over the results.
	NextLink *string `pulumi:"nextLink"`
	// Gets or sets list of resources.
	Value []GitHubOwnerResponse `pulumi:"value"`
}

func ListGitHubOwnerAvailableOutput(ctx *pulumi.Context, args ListGitHubOwnerAvailableOutputArgs, opts ...pulumi.InvokeOption) ListGitHubOwnerAvailableResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListGitHubOwnerAvailableResult, error) {
			args := v.(ListGitHubOwnerAvailableArgs)
			r, err := ListGitHubOwnerAvailable(ctx, &args, opts...)
			var s ListGitHubOwnerAvailableResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListGitHubOwnerAvailableResultOutput)
}

type ListGitHubOwnerAvailableOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The security connector name.
	SecurityConnectorName pulumi.StringInput `pulumi:"securityConnectorName"`
}

func (ListGitHubOwnerAvailableOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListGitHubOwnerAvailableArgs)(nil)).Elem()
}

// List of RP resources which supports pagination.
type ListGitHubOwnerAvailableResultOutput struct{ *pulumi.OutputState }

func (ListGitHubOwnerAvailableResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListGitHubOwnerAvailableResult)(nil)).Elem()
}

func (o ListGitHubOwnerAvailableResultOutput) ToListGitHubOwnerAvailableResultOutput() ListGitHubOwnerAvailableResultOutput {
	return o
}

func (o ListGitHubOwnerAvailableResultOutput) ToListGitHubOwnerAvailableResultOutputWithContext(ctx context.Context) ListGitHubOwnerAvailableResultOutput {
	return o
}

// Gets or sets next link to scroll over the results.
func (o ListGitHubOwnerAvailableResultOutput) NextLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListGitHubOwnerAvailableResult) *string { return v.NextLink }).(pulumi.StringPtrOutput)
}

// Gets or sets list of resources.
func (o ListGitHubOwnerAvailableResultOutput) Value() GitHubOwnerResponseArrayOutput {
	return o.ApplyT(func(v ListGitHubOwnerAvailableResult) []GitHubOwnerResponse { return v.Value }).(GitHubOwnerResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListGitHubOwnerAvailableResultOutput{})
}