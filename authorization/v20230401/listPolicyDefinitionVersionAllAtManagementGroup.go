// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// This operation lists all the policy definition versions for all policy definitions at the management group scope.
func ListPolicyDefinitionVersionAllAtManagementGroup(ctx *pulumi.Context, args *ListPolicyDefinitionVersionAllAtManagementGroupArgs, opts ...pulumi.InvokeOption) (*ListPolicyDefinitionVersionAllAtManagementGroupResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListPolicyDefinitionVersionAllAtManagementGroupResult
	err := ctx.Invoke("azure-native:authorization/v20230401:listPolicyDefinitionVersionAllAtManagementGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListPolicyDefinitionVersionAllAtManagementGroupArgs struct {
	// The name of the management group. The name is case insensitive.
	ManagementGroupName string `pulumi:"managementGroupName"`
}

// List of policy definition versions.
type ListPolicyDefinitionVersionAllAtManagementGroupResult struct {
	// The URL to use for getting the next set of results.
	NextLink *string `pulumi:"nextLink"`
	// An array of policy definitions versions.
	Value []PolicyDefinitionVersionResponse `pulumi:"value"`
}

func ListPolicyDefinitionVersionAllAtManagementGroupOutput(ctx *pulumi.Context, args ListPolicyDefinitionVersionAllAtManagementGroupOutputArgs, opts ...pulumi.InvokeOption) ListPolicyDefinitionVersionAllAtManagementGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListPolicyDefinitionVersionAllAtManagementGroupResult, error) {
			args := v.(ListPolicyDefinitionVersionAllAtManagementGroupArgs)
			r, err := ListPolicyDefinitionVersionAllAtManagementGroup(ctx, &args, opts...)
			var s ListPolicyDefinitionVersionAllAtManagementGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListPolicyDefinitionVersionAllAtManagementGroupResultOutput)
}

type ListPolicyDefinitionVersionAllAtManagementGroupOutputArgs struct {
	// The name of the management group. The name is case insensitive.
	ManagementGroupName pulumi.StringInput `pulumi:"managementGroupName"`
}

func (ListPolicyDefinitionVersionAllAtManagementGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListPolicyDefinitionVersionAllAtManagementGroupArgs)(nil)).Elem()
}

// List of policy definition versions.
type ListPolicyDefinitionVersionAllAtManagementGroupResultOutput struct{ *pulumi.OutputState }

func (ListPolicyDefinitionVersionAllAtManagementGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListPolicyDefinitionVersionAllAtManagementGroupResult)(nil)).Elem()
}

func (o ListPolicyDefinitionVersionAllAtManagementGroupResultOutput) ToListPolicyDefinitionVersionAllAtManagementGroupResultOutput() ListPolicyDefinitionVersionAllAtManagementGroupResultOutput {
	return o
}

func (o ListPolicyDefinitionVersionAllAtManagementGroupResultOutput) ToListPolicyDefinitionVersionAllAtManagementGroupResultOutputWithContext(ctx context.Context) ListPolicyDefinitionVersionAllAtManagementGroupResultOutput {
	return o
}

func (o ListPolicyDefinitionVersionAllAtManagementGroupResultOutput) ToOutput(ctx context.Context) pulumix.Output[ListPolicyDefinitionVersionAllAtManagementGroupResult] {
	return pulumix.Output[ListPolicyDefinitionVersionAllAtManagementGroupResult]{
		OutputState: o.OutputState,
	}
}

// The URL to use for getting the next set of results.
func (o ListPolicyDefinitionVersionAllAtManagementGroupResultOutput) NextLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListPolicyDefinitionVersionAllAtManagementGroupResult) *string { return v.NextLink }).(pulumi.StringPtrOutput)
}

// An array of policy definitions versions.
func (o ListPolicyDefinitionVersionAllAtManagementGroupResultOutput) Value() PolicyDefinitionVersionResponseArrayOutput {
	return o.ApplyT(func(v ListPolicyDefinitionVersionAllAtManagementGroupResult) []PolicyDefinitionVersionResponse {
		return v.Value
	}).(PolicyDefinitionVersionResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListPolicyDefinitionVersionAllAtManagementGroupResultOutput{})
}
