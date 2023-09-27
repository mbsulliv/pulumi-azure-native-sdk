// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package machinelearningservices

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Dto object representing feature
// Azure REST API version: 2023-02-01-preview.
func GetFeaturesetVersionFeature(ctx *pulumi.Context, args *GetFeaturesetVersionFeatureArgs, opts ...pulumi.InvokeOption) (*GetFeaturesetVersionFeatureResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv GetFeaturesetVersionFeatureResult
	err := ctx.Invoke("azure-native:machinelearningservices:getFeaturesetVersionFeature", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type GetFeaturesetVersionFeatureArgs struct {
	// Specifies name of the feature.
	FeatureName *string `pulumi:"featureName"`
	// Feature set name. This is case-sensitive.
	Name string `pulumi:"name"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Feature set version identifier. This is case-sensitive.
	Version string `pulumi:"version"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// Dto object representing feature
type GetFeaturesetVersionFeatureResult struct {
	// Specifies type
	DataType *string `pulumi:"dataType"`
	// Specifies description
	Description *string `pulumi:"description"`
	// Specifies name
	FeatureName *string `pulumi:"featureName"`
	// Specifies tags
	Tags map[string]string `pulumi:"tags"`
}

// Defaults sets the appropriate defaults for GetFeaturesetVersionFeatureResult
func (val *GetFeaturesetVersionFeatureResult) Defaults() *GetFeaturesetVersionFeatureResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.DataType == nil {
		dataType_ := "String"
		tmp.DataType = &dataType_
	}
	return &tmp
}

func GetFeaturesetVersionFeatureOutput(ctx *pulumi.Context, args GetFeaturesetVersionFeatureOutputArgs, opts ...pulumi.InvokeOption) GetFeaturesetVersionFeatureResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetFeaturesetVersionFeatureResult, error) {
			args := v.(GetFeaturesetVersionFeatureArgs)
			r, err := GetFeaturesetVersionFeature(ctx, &args, opts...)
			var s GetFeaturesetVersionFeatureResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetFeaturesetVersionFeatureResultOutput)
}

type GetFeaturesetVersionFeatureOutputArgs struct {
	// Specifies name of the feature.
	FeatureName pulumi.StringPtrInput `pulumi:"featureName"`
	// Feature set name. This is case-sensitive.
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Feature set version identifier. This is case-sensitive.
	Version pulumi.StringInput `pulumi:"version"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (GetFeaturesetVersionFeatureOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFeaturesetVersionFeatureArgs)(nil)).Elem()
}

// Dto object representing feature
type GetFeaturesetVersionFeatureResultOutput struct{ *pulumi.OutputState }

func (GetFeaturesetVersionFeatureResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFeaturesetVersionFeatureResult)(nil)).Elem()
}

func (o GetFeaturesetVersionFeatureResultOutput) ToGetFeaturesetVersionFeatureResultOutput() GetFeaturesetVersionFeatureResultOutput {
	return o
}

func (o GetFeaturesetVersionFeatureResultOutput) ToGetFeaturesetVersionFeatureResultOutputWithContext(ctx context.Context) GetFeaturesetVersionFeatureResultOutput {
	return o
}

func (o GetFeaturesetVersionFeatureResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetFeaturesetVersionFeatureResult] {
	return pulumix.Output[GetFeaturesetVersionFeatureResult]{
		OutputState: o.OutputState,
	}
}

// Specifies type
func (o GetFeaturesetVersionFeatureResultOutput) DataType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFeaturesetVersionFeatureResult) *string { return v.DataType }).(pulumi.StringPtrOutput)
}

// Specifies description
func (o GetFeaturesetVersionFeatureResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFeaturesetVersionFeatureResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Specifies name
func (o GetFeaturesetVersionFeatureResultOutput) FeatureName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFeaturesetVersionFeatureResult) *string { return v.FeatureName }).(pulumi.StringPtrOutput)
}

// Specifies tags
func (o GetFeaturesetVersionFeatureResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetFeaturesetVersionFeatureResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func init() {
	pulumi.RegisterOutputType(GetFeaturesetVersionFeatureResultOutput{})
}
