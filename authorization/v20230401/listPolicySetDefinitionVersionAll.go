// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This operation lists all the policy set definition versions for all policy set definitions within a subscription.
func ListPolicySetDefinitionVersionAll(ctx *pulumi.Context, args *ListPolicySetDefinitionVersionAllArgs, opts ...pulumi.InvokeOption) (*ListPolicySetDefinitionVersionAllResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListPolicySetDefinitionVersionAllResult
	err := ctx.Invoke("azure-native:authorization/v20230401:listPolicySetDefinitionVersionAll", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListPolicySetDefinitionVersionAllArgs struct {
}

// List of policy set definition versions.
type ListPolicySetDefinitionVersionAllResult struct {
	// The URL to use for getting the next set of results.
	NextLink *string `pulumi:"nextLink"`
	// An array of policy set definition versions.
	Value []PolicySetDefinitionVersionResponse `pulumi:"value"`
}

func ListPolicySetDefinitionVersionAllOutput(ctx *pulumi.Context, args ListPolicySetDefinitionVersionAllOutputArgs, opts ...pulumi.InvokeOption) ListPolicySetDefinitionVersionAllResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListPolicySetDefinitionVersionAllResult, error) {
			args := v.(ListPolicySetDefinitionVersionAllArgs)
			r, err := ListPolicySetDefinitionVersionAll(ctx, &args, opts...)
			var s ListPolicySetDefinitionVersionAllResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListPolicySetDefinitionVersionAllResultOutput)
}

type ListPolicySetDefinitionVersionAllOutputArgs struct {
}

func (ListPolicySetDefinitionVersionAllOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListPolicySetDefinitionVersionAllArgs)(nil)).Elem()
}

// List of policy set definition versions.
type ListPolicySetDefinitionVersionAllResultOutput struct{ *pulumi.OutputState }

func (ListPolicySetDefinitionVersionAllResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListPolicySetDefinitionVersionAllResult)(nil)).Elem()
}

func (o ListPolicySetDefinitionVersionAllResultOutput) ToListPolicySetDefinitionVersionAllResultOutput() ListPolicySetDefinitionVersionAllResultOutput {
	return o
}

func (o ListPolicySetDefinitionVersionAllResultOutput) ToListPolicySetDefinitionVersionAllResultOutputWithContext(ctx context.Context) ListPolicySetDefinitionVersionAllResultOutput {
	return o
}

// The URL to use for getting the next set of results.
func (o ListPolicySetDefinitionVersionAllResultOutput) NextLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListPolicySetDefinitionVersionAllResult) *string { return v.NextLink }).(pulumi.StringPtrOutput)
}

// An array of policy set definition versions.
func (o ListPolicySetDefinitionVersionAllResultOutput) Value() PolicySetDefinitionVersionResponseArrayOutput {
	return o.ApplyT(func(v ListPolicySetDefinitionVersionAllResult) []PolicySetDefinitionVersionResponse { return v.Value }).(PolicySetDefinitionVersionResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListPolicySetDefinitionVersionAllResultOutput{})
}
