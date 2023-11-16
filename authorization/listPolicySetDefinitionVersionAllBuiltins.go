// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package authorization

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This operation lists all the built-in policy set definition versions for all built-in policy set definitions.
// Azure REST API version: 2023-04-01.
func ListPolicySetDefinitionVersionAllBuiltins(ctx *pulumi.Context, args *ListPolicySetDefinitionVersionAllBuiltinsArgs, opts ...pulumi.InvokeOption) (*ListPolicySetDefinitionVersionAllBuiltinsResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListPolicySetDefinitionVersionAllBuiltinsResult
	err := ctx.Invoke("azure-native:authorization:listPolicySetDefinitionVersionAllBuiltins", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListPolicySetDefinitionVersionAllBuiltinsArgs struct {
}

// List of policy set definition versions.
type ListPolicySetDefinitionVersionAllBuiltinsResult struct {
	// The URL to use for getting the next set of results.
	NextLink *string `pulumi:"nextLink"`
	// An array of policy set definition versions.
	Value []PolicySetDefinitionVersionResponse `pulumi:"value"`
}

func ListPolicySetDefinitionVersionAllBuiltinsOutput(ctx *pulumi.Context, args ListPolicySetDefinitionVersionAllBuiltinsOutputArgs, opts ...pulumi.InvokeOption) ListPolicySetDefinitionVersionAllBuiltinsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListPolicySetDefinitionVersionAllBuiltinsResult, error) {
			args := v.(ListPolicySetDefinitionVersionAllBuiltinsArgs)
			r, err := ListPolicySetDefinitionVersionAllBuiltins(ctx, &args, opts...)
			var s ListPolicySetDefinitionVersionAllBuiltinsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListPolicySetDefinitionVersionAllBuiltinsResultOutput)
}

type ListPolicySetDefinitionVersionAllBuiltinsOutputArgs struct {
}

func (ListPolicySetDefinitionVersionAllBuiltinsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListPolicySetDefinitionVersionAllBuiltinsArgs)(nil)).Elem()
}

// List of policy set definition versions.
type ListPolicySetDefinitionVersionAllBuiltinsResultOutput struct{ *pulumi.OutputState }

func (ListPolicySetDefinitionVersionAllBuiltinsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListPolicySetDefinitionVersionAllBuiltinsResult)(nil)).Elem()
}

func (o ListPolicySetDefinitionVersionAllBuiltinsResultOutput) ToListPolicySetDefinitionVersionAllBuiltinsResultOutput() ListPolicySetDefinitionVersionAllBuiltinsResultOutput {
	return o
}

func (o ListPolicySetDefinitionVersionAllBuiltinsResultOutput) ToListPolicySetDefinitionVersionAllBuiltinsResultOutputWithContext(ctx context.Context) ListPolicySetDefinitionVersionAllBuiltinsResultOutput {
	return o
}

// The URL to use for getting the next set of results.
func (o ListPolicySetDefinitionVersionAllBuiltinsResultOutput) NextLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListPolicySetDefinitionVersionAllBuiltinsResult) *string { return v.NextLink }).(pulumi.StringPtrOutput)
}

// An array of policy set definition versions.
func (o ListPolicySetDefinitionVersionAllBuiltinsResultOutput) Value() PolicySetDefinitionVersionResponseArrayOutput {
	return o.ApplyT(func(v ListPolicySetDefinitionVersionAllBuiltinsResult) []PolicySetDefinitionVersionResponse {
		return v.Value
	}).(PolicySetDefinitionVersionResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListPolicySetDefinitionVersionAllBuiltinsResultOutput{})
}
