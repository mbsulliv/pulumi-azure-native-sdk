// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220801preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// This operation retrieves a single variable value; given its name,  management group it was created at and the variable it's created for.
func LookupVariableValueAtManagementGroup(ctx *pulumi.Context, args *LookupVariableValueAtManagementGroupArgs, opts ...pulumi.InvokeOption) (*LookupVariableValueAtManagementGroupResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupVariableValueAtManagementGroupResult
	err := ctx.Invoke("azure-native:authorization/v20220801preview:getVariableValueAtManagementGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVariableValueAtManagementGroupArgs struct {
	// The ID of the management group.
	ManagementGroupId string `pulumi:"managementGroupId"`
	// The name of the variable to operate on.
	VariableName string `pulumi:"variableName"`
	// The name of the variable value to operate on.
	VariableValueName string `pulumi:"variableValueName"`
}

// The variable value.
type LookupVariableValueAtManagementGroupResult struct {
	// The ID of the variable.
	Id string `pulumi:"id"`
	// The name of the variable.
	Name string `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource (Microsoft.Authorization/variables/values).
	Type string `pulumi:"type"`
	// Variable value column value array.
	Values []PolicyVariableValueColumnValueResponse `pulumi:"values"`
}

func LookupVariableValueAtManagementGroupOutput(ctx *pulumi.Context, args LookupVariableValueAtManagementGroupOutputArgs, opts ...pulumi.InvokeOption) LookupVariableValueAtManagementGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVariableValueAtManagementGroupResult, error) {
			args := v.(LookupVariableValueAtManagementGroupArgs)
			r, err := LookupVariableValueAtManagementGroup(ctx, &args, opts...)
			var s LookupVariableValueAtManagementGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVariableValueAtManagementGroupResultOutput)
}

type LookupVariableValueAtManagementGroupOutputArgs struct {
	// The ID of the management group.
	ManagementGroupId pulumi.StringInput `pulumi:"managementGroupId"`
	// The name of the variable to operate on.
	VariableName pulumi.StringInput `pulumi:"variableName"`
	// The name of the variable value to operate on.
	VariableValueName pulumi.StringInput `pulumi:"variableValueName"`
}

func (LookupVariableValueAtManagementGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVariableValueAtManagementGroupArgs)(nil)).Elem()
}

// The variable value.
type LookupVariableValueAtManagementGroupResultOutput struct{ *pulumi.OutputState }

func (LookupVariableValueAtManagementGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVariableValueAtManagementGroupResult)(nil)).Elem()
}

func (o LookupVariableValueAtManagementGroupResultOutput) ToLookupVariableValueAtManagementGroupResultOutput() LookupVariableValueAtManagementGroupResultOutput {
	return o
}

func (o LookupVariableValueAtManagementGroupResultOutput) ToLookupVariableValueAtManagementGroupResultOutputWithContext(ctx context.Context) LookupVariableValueAtManagementGroupResultOutput {
	return o
}

func (o LookupVariableValueAtManagementGroupResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupVariableValueAtManagementGroupResult] {
	return pulumix.Output[LookupVariableValueAtManagementGroupResult]{
		OutputState: o.OutputState,
	}
}

// The ID of the variable.
func (o LookupVariableValueAtManagementGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVariableValueAtManagementGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the variable.
func (o LookupVariableValueAtManagementGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVariableValueAtManagementGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupVariableValueAtManagementGroupResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupVariableValueAtManagementGroupResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource (Microsoft.Authorization/variables/values).
func (o LookupVariableValueAtManagementGroupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVariableValueAtManagementGroupResult) string { return v.Type }).(pulumi.StringOutput)
}

// Variable value column value array.
func (o LookupVariableValueAtManagementGroupResultOutput) Values() PolicyVariableValueColumnValueResponseArrayOutput {
	return o.ApplyT(func(v LookupVariableValueAtManagementGroupResult) []PolicyVariableValueColumnValueResponse {
		return v.Values
	}).(PolicyVariableValueColumnValueResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVariableValueAtManagementGroupResultOutput{})
}
