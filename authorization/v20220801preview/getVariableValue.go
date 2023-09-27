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

// This operation retrieves a single variable value; given its name, subscription it was created at and the variable it's created for.
func LookupVariableValue(ctx *pulumi.Context, args *LookupVariableValueArgs, opts ...pulumi.InvokeOption) (*LookupVariableValueResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupVariableValueResult
	err := ctx.Invoke("azure-native:authorization/v20220801preview:getVariableValue", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVariableValueArgs struct {
	// The name of the variable to operate on.
	VariableName string `pulumi:"variableName"`
	// The name of the variable value to operate on.
	VariableValueName string `pulumi:"variableValueName"`
}

// The variable value.
type LookupVariableValueResult struct {
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

func LookupVariableValueOutput(ctx *pulumi.Context, args LookupVariableValueOutputArgs, opts ...pulumi.InvokeOption) LookupVariableValueResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVariableValueResult, error) {
			args := v.(LookupVariableValueArgs)
			r, err := LookupVariableValue(ctx, &args, opts...)
			var s LookupVariableValueResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVariableValueResultOutput)
}

type LookupVariableValueOutputArgs struct {
	// The name of the variable to operate on.
	VariableName pulumi.StringInput `pulumi:"variableName"`
	// The name of the variable value to operate on.
	VariableValueName pulumi.StringInput `pulumi:"variableValueName"`
}

func (LookupVariableValueOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVariableValueArgs)(nil)).Elem()
}

// The variable value.
type LookupVariableValueResultOutput struct{ *pulumi.OutputState }

func (LookupVariableValueResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVariableValueResult)(nil)).Elem()
}

func (o LookupVariableValueResultOutput) ToLookupVariableValueResultOutput() LookupVariableValueResultOutput {
	return o
}

func (o LookupVariableValueResultOutput) ToLookupVariableValueResultOutputWithContext(ctx context.Context) LookupVariableValueResultOutput {
	return o
}

func (o LookupVariableValueResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupVariableValueResult] {
	return pulumix.Output[LookupVariableValueResult]{
		OutputState: o.OutputState,
	}
}

// The ID of the variable.
func (o LookupVariableValueResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVariableValueResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the variable.
func (o LookupVariableValueResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVariableValueResult) string { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupVariableValueResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupVariableValueResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource (Microsoft.Authorization/variables/values).
func (o LookupVariableValueResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVariableValueResult) string { return v.Type }).(pulumi.StringOutput)
}

// Variable value column value array.
func (o LookupVariableValueResultOutput) Values() PolicyVariableValueColumnValueResponseArrayOutput {
	return o.ApplyT(func(v LookupVariableValueResult) []PolicyVariableValueColumnValueResponse { return v.Values }).(PolicyVariableValueColumnValueResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVariableValueResultOutput{})
}
