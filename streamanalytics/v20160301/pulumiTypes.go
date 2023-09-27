// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20160301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

var _ = utilities.GetEnvOrDefault

// The binding to an Azure Machine Learning web service.
type AzureMachineLearningWebServiceFunctionBinding struct {
	// The API key used to authenticate with Request-Response endpoint.
	ApiKey *string `pulumi:"apiKey"`
	// Number between 1 and 10000 describing maximum number of rows for every Azure ML RRS execute request. Default is 1000.
	BatchSize *int `pulumi:"batchSize"`
	// The Request-Response execute endpoint of the Azure Machine Learning web service. Find out more here: https://docs.microsoft.com/en-us/azure/machine-learning/machine-learning-consume-web-services#request-response-service-rrs
	Endpoint *string `pulumi:"endpoint"`
	// The inputs for the Azure Machine Learning web service endpoint.
	Inputs *AzureMachineLearningWebServiceInputs `pulumi:"inputs"`
	// A list of outputs from the Azure Machine Learning web service endpoint execution.
	Outputs []AzureMachineLearningWebServiceOutputColumn `pulumi:"outputs"`
	// Indicates the function binding type.
	// Expected value is 'Microsoft.MachineLearning/WebService'.
	Type string `pulumi:"type"`
}

// The binding to an Azure Machine Learning web service.
type AzureMachineLearningWebServiceFunctionBindingResponse struct {
	// The API key used to authenticate with Request-Response endpoint.
	ApiKey *string `pulumi:"apiKey"`
	// Number between 1 and 10000 describing maximum number of rows for every Azure ML RRS execute request. Default is 1000.
	BatchSize *int `pulumi:"batchSize"`
	// The Request-Response execute endpoint of the Azure Machine Learning web service. Find out more here: https://docs.microsoft.com/en-us/azure/machine-learning/machine-learning-consume-web-services#request-response-service-rrs
	Endpoint *string `pulumi:"endpoint"`
	// The inputs for the Azure Machine Learning web service endpoint.
	Inputs *AzureMachineLearningWebServiceInputsResponse `pulumi:"inputs"`
	// A list of outputs from the Azure Machine Learning web service endpoint execution.
	Outputs []AzureMachineLearningWebServiceOutputColumnResponse `pulumi:"outputs"`
	// Indicates the function binding type.
	// Expected value is 'Microsoft.MachineLearning/WebService'.
	Type string `pulumi:"type"`
}

// Describes an input column for the Azure Machine Learning web service endpoint.
type AzureMachineLearningWebServiceInputColumn struct {
	// The (Azure Machine Learning supported) data type of the input column. A list of valid  Azure Machine Learning data types are described at https://msdn.microsoft.com/en-us/library/azure/dn905923.aspx .
	DataType *string `pulumi:"dataType"`
	// The zero based index of the function parameter this input maps to.
	MapTo *int `pulumi:"mapTo"`
	// The name of the input column.
	Name *string `pulumi:"name"`
}

// Describes an input column for the Azure Machine Learning web service endpoint.
type AzureMachineLearningWebServiceInputColumnResponse struct {
	// The (Azure Machine Learning supported) data type of the input column. A list of valid  Azure Machine Learning data types are described at https://msdn.microsoft.com/en-us/library/azure/dn905923.aspx .
	DataType *string `pulumi:"dataType"`
	// The zero based index of the function parameter this input maps to.
	MapTo *int `pulumi:"mapTo"`
	// The name of the input column.
	Name *string `pulumi:"name"`
}

// The inputs for the Azure Machine Learning web service endpoint.
type AzureMachineLearningWebServiceInputs struct {
	// A list of input columns for the Azure Machine Learning web service endpoint.
	ColumnNames []AzureMachineLearningWebServiceInputColumn `pulumi:"columnNames"`
	// The name of the input. This is the name provided while authoring the endpoint.
	Name *string `pulumi:"name"`
}

// The inputs for the Azure Machine Learning web service endpoint.
type AzureMachineLearningWebServiceInputsResponse struct {
	// A list of input columns for the Azure Machine Learning web service endpoint.
	ColumnNames []AzureMachineLearningWebServiceInputColumnResponse `pulumi:"columnNames"`
	// The name of the input. This is the name provided while authoring the endpoint.
	Name *string `pulumi:"name"`
}

// Describes an output column for the Azure Machine Learning web service endpoint.
type AzureMachineLearningWebServiceOutputColumn struct {
	// The (Azure Machine Learning supported) data type of the output column. A list of valid  Azure Machine Learning data types are described at https://msdn.microsoft.com/en-us/library/azure/dn905923.aspx .
	DataType *string `pulumi:"dataType"`
	// The name of the output column.
	Name *string `pulumi:"name"`
}

// Describes an output column for the Azure Machine Learning web service endpoint.
type AzureMachineLearningWebServiceOutputColumnResponse struct {
	// The (Azure Machine Learning supported) data type of the output column. A list of valid  Azure Machine Learning data types are described at https://msdn.microsoft.com/en-us/library/azure/dn905923.aspx .
	DataType *string `pulumi:"dataType"`
	// The name of the output column.
	Name *string `pulumi:"name"`
}

// Describes one input parameter of a function.
type FunctionInputType struct {
	// The (Azure Stream Analytics supported) data type of the function input parameter. A list of valid Azure Stream Analytics data types are described at https://msdn.microsoft.com/en-us/library/azure/dn835065.aspx
	DataType *string `pulumi:"dataType"`
	// A flag indicating if the parameter is a configuration parameter. True if this input parameter is expected to be a constant. Default is false.
	IsConfigurationParameter *bool `pulumi:"isConfigurationParameter"`
}

// FunctionInputTypeInput is an input type that accepts FunctionInputTypeArgs and FunctionInputTypeOutput values.
// You can construct a concrete instance of `FunctionInputTypeInput` via:
//
//	FunctionInputTypeArgs{...}
type FunctionInputTypeInput interface {
	pulumi.Input

	ToFunctionInputTypeOutput() FunctionInputTypeOutput
	ToFunctionInputTypeOutputWithContext(context.Context) FunctionInputTypeOutput
}

// Describes one input parameter of a function.
type FunctionInputTypeArgs struct {
	// The (Azure Stream Analytics supported) data type of the function input parameter. A list of valid Azure Stream Analytics data types are described at https://msdn.microsoft.com/en-us/library/azure/dn835065.aspx
	DataType pulumi.StringPtrInput `pulumi:"dataType"`
	// A flag indicating if the parameter is a configuration parameter. True if this input parameter is expected to be a constant. Default is false.
	IsConfigurationParameter pulumi.BoolPtrInput `pulumi:"isConfigurationParameter"`
}

func (FunctionInputTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FunctionInputType)(nil)).Elem()
}

func (i FunctionInputTypeArgs) ToFunctionInputTypeOutput() FunctionInputTypeOutput {
	return i.ToFunctionInputTypeOutputWithContext(context.Background())
}

func (i FunctionInputTypeArgs) ToFunctionInputTypeOutputWithContext(ctx context.Context) FunctionInputTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionInputTypeOutput)
}

func (i FunctionInputTypeArgs) ToOutput(ctx context.Context) pulumix.Output[FunctionInputType] {
	return pulumix.Output[FunctionInputType]{
		OutputState: i.ToFunctionInputTypeOutputWithContext(ctx).OutputState,
	}
}

// FunctionInputTypeArrayInput is an input type that accepts FunctionInputTypeArray and FunctionInputTypeArrayOutput values.
// You can construct a concrete instance of `FunctionInputTypeArrayInput` via:
//
//	FunctionInputTypeArray{ FunctionInputTypeArgs{...} }
type FunctionInputTypeArrayInput interface {
	pulumi.Input

	ToFunctionInputTypeArrayOutput() FunctionInputTypeArrayOutput
	ToFunctionInputTypeArrayOutputWithContext(context.Context) FunctionInputTypeArrayOutput
}

type FunctionInputTypeArray []FunctionInputTypeInput

func (FunctionInputTypeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FunctionInputType)(nil)).Elem()
}

func (i FunctionInputTypeArray) ToFunctionInputTypeArrayOutput() FunctionInputTypeArrayOutput {
	return i.ToFunctionInputTypeArrayOutputWithContext(context.Background())
}

func (i FunctionInputTypeArray) ToFunctionInputTypeArrayOutputWithContext(ctx context.Context) FunctionInputTypeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionInputTypeArrayOutput)
}

func (i FunctionInputTypeArray) ToOutput(ctx context.Context) pulumix.Output[[]FunctionInputType] {
	return pulumix.Output[[]FunctionInputType]{
		OutputState: i.ToFunctionInputTypeArrayOutputWithContext(ctx).OutputState,
	}
}

// Describes one input parameter of a function.
type FunctionInputTypeOutput struct{ *pulumi.OutputState }

func (FunctionInputTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FunctionInputType)(nil)).Elem()
}

func (o FunctionInputTypeOutput) ToFunctionInputTypeOutput() FunctionInputTypeOutput {
	return o
}

func (o FunctionInputTypeOutput) ToFunctionInputTypeOutputWithContext(ctx context.Context) FunctionInputTypeOutput {
	return o
}

func (o FunctionInputTypeOutput) ToOutput(ctx context.Context) pulumix.Output[FunctionInputType] {
	return pulumix.Output[FunctionInputType]{
		OutputState: o.OutputState,
	}
}

// The (Azure Stream Analytics supported) data type of the function input parameter. A list of valid Azure Stream Analytics data types are described at https://msdn.microsoft.com/en-us/library/azure/dn835065.aspx
func (o FunctionInputTypeOutput) DataType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FunctionInputType) *string { return v.DataType }).(pulumi.StringPtrOutput)
}

// A flag indicating if the parameter is a configuration parameter. True if this input parameter is expected to be a constant. Default is false.
func (o FunctionInputTypeOutput) IsConfigurationParameter() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v FunctionInputType) *bool { return v.IsConfigurationParameter }).(pulumi.BoolPtrOutput)
}

type FunctionInputTypeArrayOutput struct{ *pulumi.OutputState }

func (FunctionInputTypeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FunctionInputType)(nil)).Elem()
}

func (o FunctionInputTypeArrayOutput) ToFunctionInputTypeArrayOutput() FunctionInputTypeArrayOutput {
	return o
}

func (o FunctionInputTypeArrayOutput) ToFunctionInputTypeArrayOutputWithContext(ctx context.Context) FunctionInputTypeArrayOutput {
	return o
}

func (o FunctionInputTypeArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]FunctionInputType] {
	return pulumix.Output[[]FunctionInputType]{
		OutputState: o.OutputState,
	}
}

func (o FunctionInputTypeArrayOutput) Index(i pulumi.IntInput) FunctionInputTypeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FunctionInputType {
		return vs[0].([]FunctionInputType)[vs[1].(int)]
	}).(FunctionInputTypeOutput)
}

// Describes one input parameter of a function.
type FunctionInputResponse struct {
	// The (Azure Stream Analytics supported) data type of the function input parameter. A list of valid Azure Stream Analytics data types are described at https://msdn.microsoft.com/en-us/library/azure/dn835065.aspx
	DataType *string `pulumi:"dataType"`
	// A flag indicating if the parameter is a configuration parameter. True if this input parameter is expected to be a constant. Default is false.
	IsConfigurationParameter *bool `pulumi:"isConfigurationParameter"`
}

// Describes one input parameter of a function.
type FunctionInputResponseOutput struct{ *pulumi.OutputState }

func (FunctionInputResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FunctionInputResponse)(nil)).Elem()
}

func (o FunctionInputResponseOutput) ToFunctionInputResponseOutput() FunctionInputResponseOutput {
	return o
}

func (o FunctionInputResponseOutput) ToFunctionInputResponseOutputWithContext(ctx context.Context) FunctionInputResponseOutput {
	return o
}

func (o FunctionInputResponseOutput) ToOutput(ctx context.Context) pulumix.Output[FunctionInputResponse] {
	return pulumix.Output[FunctionInputResponse]{
		OutputState: o.OutputState,
	}
}

// The (Azure Stream Analytics supported) data type of the function input parameter. A list of valid Azure Stream Analytics data types are described at https://msdn.microsoft.com/en-us/library/azure/dn835065.aspx
func (o FunctionInputResponseOutput) DataType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FunctionInputResponse) *string { return v.DataType }).(pulumi.StringPtrOutput)
}

// A flag indicating if the parameter is a configuration parameter. True if this input parameter is expected to be a constant. Default is false.
func (o FunctionInputResponseOutput) IsConfigurationParameter() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v FunctionInputResponse) *bool { return v.IsConfigurationParameter }).(pulumi.BoolPtrOutput)
}

type FunctionInputResponseArrayOutput struct{ *pulumi.OutputState }

func (FunctionInputResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FunctionInputResponse)(nil)).Elem()
}

func (o FunctionInputResponseArrayOutput) ToFunctionInputResponseArrayOutput() FunctionInputResponseArrayOutput {
	return o
}

func (o FunctionInputResponseArrayOutput) ToFunctionInputResponseArrayOutputWithContext(ctx context.Context) FunctionInputResponseArrayOutput {
	return o
}

func (o FunctionInputResponseArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]FunctionInputResponse] {
	return pulumix.Output[[]FunctionInputResponse]{
		OutputState: o.OutputState,
	}
}

func (o FunctionInputResponseArrayOutput) Index(i pulumi.IntInput) FunctionInputResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FunctionInputResponse {
		return vs[0].([]FunctionInputResponse)[vs[1].(int)]
	}).(FunctionInputResponseOutput)
}

// Describes the output of a function.
type FunctionOutputType struct {
	// The (Azure Stream Analytics supported) data type of the function output. A list of valid Azure Stream Analytics data types are described at https://msdn.microsoft.com/en-us/library/azure/dn835065.aspx
	DataType *string `pulumi:"dataType"`
}

// FunctionOutputTypeInput is an input type that accepts FunctionOutputTypeArgs and FunctionOutputTypeOutput values.
// You can construct a concrete instance of `FunctionOutputTypeInput` via:
//
//	FunctionOutputTypeArgs{...}
type FunctionOutputTypeInput interface {
	pulumi.Input

	ToFunctionOutputTypeOutput() FunctionOutputTypeOutput
	ToFunctionOutputTypeOutputWithContext(context.Context) FunctionOutputTypeOutput
}

// Describes the output of a function.
type FunctionOutputTypeArgs struct {
	// The (Azure Stream Analytics supported) data type of the function output. A list of valid Azure Stream Analytics data types are described at https://msdn.microsoft.com/en-us/library/azure/dn835065.aspx
	DataType pulumi.StringPtrInput `pulumi:"dataType"`
}

func (FunctionOutputTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FunctionOutputType)(nil)).Elem()
}

func (i FunctionOutputTypeArgs) ToFunctionOutputTypeOutput() FunctionOutputTypeOutput {
	return i.ToFunctionOutputTypeOutputWithContext(context.Background())
}

func (i FunctionOutputTypeArgs) ToFunctionOutputTypeOutputWithContext(ctx context.Context) FunctionOutputTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionOutputTypeOutput)
}

func (i FunctionOutputTypeArgs) ToOutput(ctx context.Context) pulumix.Output[FunctionOutputType] {
	return pulumix.Output[FunctionOutputType]{
		OutputState: i.ToFunctionOutputTypeOutputWithContext(ctx).OutputState,
	}
}

func (i FunctionOutputTypeArgs) ToFunctionOutputTypePtrOutput() FunctionOutputTypePtrOutput {
	return i.ToFunctionOutputTypePtrOutputWithContext(context.Background())
}

func (i FunctionOutputTypeArgs) ToFunctionOutputTypePtrOutputWithContext(ctx context.Context) FunctionOutputTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionOutputTypeOutput).ToFunctionOutputTypePtrOutputWithContext(ctx)
}

// FunctionOutputTypePtrInput is an input type that accepts FunctionOutputTypeArgs, FunctionOutputTypePtr and FunctionOutputTypePtrOutput values.
// You can construct a concrete instance of `FunctionOutputTypePtrInput` via:
//
//	        FunctionOutputTypeArgs{...}
//
//	or:
//
//	        nil
type FunctionOutputTypePtrInput interface {
	pulumi.Input

	ToFunctionOutputTypePtrOutput() FunctionOutputTypePtrOutput
	ToFunctionOutputTypePtrOutputWithContext(context.Context) FunctionOutputTypePtrOutput
}

type functionOutputTypePtrType FunctionOutputTypeArgs

func FunctionOutputTypePtr(v *FunctionOutputTypeArgs) FunctionOutputTypePtrInput {
	return (*functionOutputTypePtrType)(v)
}

func (*functionOutputTypePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FunctionOutputType)(nil)).Elem()
}

func (i *functionOutputTypePtrType) ToFunctionOutputTypePtrOutput() FunctionOutputTypePtrOutput {
	return i.ToFunctionOutputTypePtrOutputWithContext(context.Background())
}

func (i *functionOutputTypePtrType) ToFunctionOutputTypePtrOutputWithContext(ctx context.Context) FunctionOutputTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionOutputTypePtrOutput)
}

func (i *functionOutputTypePtrType) ToOutput(ctx context.Context) pulumix.Output[*FunctionOutputType] {
	return pulumix.Output[*FunctionOutputType]{
		OutputState: i.ToFunctionOutputTypePtrOutputWithContext(ctx).OutputState,
	}
}

// Describes the output of a function.
type FunctionOutputTypeOutput struct{ *pulumi.OutputState }

func (FunctionOutputTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FunctionOutputType)(nil)).Elem()
}

func (o FunctionOutputTypeOutput) ToFunctionOutputTypeOutput() FunctionOutputTypeOutput {
	return o
}

func (o FunctionOutputTypeOutput) ToFunctionOutputTypeOutputWithContext(ctx context.Context) FunctionOutputTypeOutput {
	return o
}

func (o FunctionOutputTypeOutput) ToFunctionOutputTypePtrOutput() FunctionOutputTypePtrOutput {
	return o.ToFunctionOutputTypePtrOutputWithContext(context.Background())
}

func (o FunctionOutputTypeOutput) ToFunctionOutputTypePtrOutputWithContext(ctx context.Context) FunctionOutputTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FunctionOutputType) *FunctionOutputType {
		return &v
	}).(FunctionOutputTypePtrOutput)
}

func (o FunctionOutputTypeOutput) ToOutput(ctx context.Context) pulumix.Output[FunctionOutputType] {
	return pulumix.Output[FunctionOutputType]{
		OutputState: o.OutputState,
	}
}

// The (Azure Stream Analytics supported) data type of the function output. A list of valid Azure Stream Analytics data types are described at https://msdn.microsoft.com/en-us/library/azure/dn835065.aspx
func (o FunctionOutputTypeOutput) DataType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FunctionOutputType) *string { return v.DataType }).(pulumi.StringPtrOutput)
}

type FunctionOutputTypePtrOutput struct{ *pulumi.OutputState }

func (FunctionOutputTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FunctionOutputType)(nil)).Elem()
}

func (o FunctionOutputTypePtrOutput) ToFunctionOutputTypePtrOutput() FunctionOutputTypePtrOutput {
	return o
}

func (o FunctionOutputTypePtrOutput) ToFunctionOutputTypePtrOutputWithContext(ctx context.Context) FunctionOutputTypePtrOutput {
	return o
}

func (o FunctionOutputTypePtrOutput) ToOutput(ctx context.Context) pulumix.Output[*FunctionOutputType] {
	return pulumix.Output[*FunctionOutputType]{
		OutputState: o.OutputState,
	}
}

func (o FunctionOutputTypePtrOutput) Elem() FunctionOutputTypeOutput {
	return o.ApplyT(func(v *FunctionOutputType) FunctionOutputType {
		if v != nil {
			return *v
		}
		var ret FunctionOutputType
		return ret
	}).(FunctionOutputTypeOutput)
}

// The (Azure Stream Analytics supported) data type of the function output. A list of valid Azure Stream Analytics data types are described at https://msdn.microsoft.com/en-us/library/azure/dn835065.aspx
func (o FunctionOutputTypePtrOutput) DataType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FunctionOutputType) *string {
		if v == nil {
			return nil
		}
		return v.DataType
	}).(pulumi.StringPtrOutput)
}

// Describes the output of a function.
type FunctionOutputResponse struct {
	// The (Azure Stream Analytics supported) data type of the function output. A list of valid Azure Stream Analytics data types are described at https://msdn.microsoft.com/en-us/library/azure/dn835065.aspx
	DataType *string `pulumi:"dataType"`
}

// Describes the output of a function.
type FunctionOutputResponseOutput struct{ *pulumi.OutputState }

func (FunctionOutputResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FunctionOutputResponse)(nil)).Elem()
}

func (o FunctionOutputResponseOutput) ToFunctionOutputResponseOutput() FunctionOutputResponseOutput {
	return o
}

func (o FunctionOutputResponseOutput) ToFunctionOutputResponseOutputWithContext(ctx context.Context) FunctionOutputResponseOutput {
	return o
}

func (o FunctionOutputResponseOutput) ToOutput(ctx context.Context) pulumix.Output[FunctionOutputResponse] {
	return pulumix.Output[FunctionOutputResponse]{
		OutputState: o.OutputState,
	}
}

// The (Azure Stream Analytics supported) data type of the function output. A list of valid Azure Stream Analytics data types are described at https://msdn.microsoft.com/en-us/library/azure/dn835065.aspx
func (o FunctionOutputResponseOutput) DataType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FunctionOutputResponse) *string { return v.DataType }).(pulumi.StringPtrOutput)
}

type FunctionOutputResponsePtrOutput struct{ *pulumi.OutputState }

func (FunctionOutputResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FunctionOutputResponse)(nil)).Elem()
}

func (o FunctionOutputResponsePtrOutput) ToFunctionOutputResponsePtrOutput() FunctionOutputResponsePtrOutput {
	return o
}

func (o FunctionOutputResponsePtrOutput) ToFunctionOutputResponsePtrOutputWithContext(ctx context.Context) FunctionOutputResponsePtrOutput {
	return o
}

func (o FunctionOutputResponsePtrOutput) ToOutput(ctx context.Context) pulumix.Output[*FunctionOutputResponse] {
	return pulumix.Output[*FunctionOutputResponse]{
		OutputState: o.OutputState,
	}
}

func (o FunctionOutputResponsePtrOutput) Elem() FunctionOutputResponseOutput {
	return o.ApplyT(func(v *FunctionOutputResponse) FunctionOutputResponse {
		if v != nil {
			return *v
		}
		var ret FunctionOutputResponse
		return ret
	}).(FunctionOutputResponseOutput)
}

// The (Azure Stream Analytics supported) data type of the function output. A list of valid Azure Stream Analytics data types are described at https://msdn.microsoft.com/en-us/library/azure/dn835065.aspx
func (o FunctionOutputResponsePtrOutput) DataType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FunctionOutputResponse) *string {
		if v == nil {
			return nil
		}
		return v.DataType
	}).(pulumi.StringPtrOutput)
}

// The binding to a JavaScript function.
type JavaScriptFunctionBinding struct {
	// The JavaScript code containing a single function definition. For example: 'function (x, y) { return x + y; }'
	Script *string `pulumi:"script"`
	// Indicates the function binding type.
	// Expected value is 'Microsoft.StreamAnalytics/JavascriptUdf'.
	Type string `pulumi:"type"`
}

// The binding to a JavaScript function.
type JavaScriptFunctionBindingResponse struct {
	// The JavaScript code containing a single function definition. For example: 'function (x, y) { return x + y; }'
	Script *string `pulumi:"script"`
	// Indicates the function binding type.
	// Expected value is 'Microsoft.StreamAnalytics/JavascriptUdf'.
	Type string `pulumi:"type"`
}

// The properties that are associated with a scalar function.
type ScalarFunctionProperties struct {
	// The physical binding of the function. For example, in the Azure Machine Learning web service’s case, this describes the endpoint.
	Binding interface{} `pulumi:"binding"`
	// A list of inputs describing the parameters of the function.
	Inputs []FunctionInputType `pulumi:"inputs"`
	// The output of the function.
	Output *FunctionOutputType `pulumi:"output"`
	// Indicates the type of function.
	// Expected value is 'Scalar'.
	Type string `pulumi:"type"`
}

// ScalarFunctionPropertiesInput is an input type that accepts ScalarFunctionPropertiesArgs and ScalarFunctionPropertiesOutput values.
// You can construct a concrete instance of `ScalarFunctionPropertiesInput` via:
//
//	ScalarFunctionPropertiesArgs{...}
type ScalarFunctionPropertiesInput interface {
	pulumi.Input

	ToScalarFunctionPropertiesOutput() ScalarFunctionPropertiesOutput
	ToScalarFunctionPropertiesOutputWithContext(context.Context) ScalarFunctionPropertiesOutput
}

// The properties that are associated with a scalar function.
type ScalarFunctionPropertiesArgs struct {
	// The physical binding of the function. For example, in the Azure Machine Learning web service’s case, this describes the endpoint.
	Binding pulumi.Input `pulumi:"binding"`
	// A list of inputs describing the parameters of the function.
	Inputs FunctionInputTypeArrayInput `pulumi:"inputs"`
	// The output of the function.
	Output FunctionOutputTypePtrInput `pulumi:"output"`
	// Indicates the type of function.
	// Expected value is 'Scalar'.
	Type pulumi.StringInput `pulumi:"type"`
}

func (ScalarFunctionPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ScalarFunctionProperties)(nil)).Elem()
}

func (i ScalarFunctionPropertiesArgs) ToScalarFunctionPropertiesOutput() ScalarFunctionPropertiesOutput {
	return i.ToScalarFunctionPropertiesOutputWithContext(context.Background())
}

func (i ScalarFunctionPropertiesArgs) ToScalarFunctionPropertiesOutputWithContext(ctx context.Context) ScalarFunctionPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScalarFunctionPropertiesOutput)
}

func (i ScalarFunctionPropertiesArgs) ToOutput(ctx context.Context) pulumix.Output[ScalarFunctionProperties] {
	return pulumix.Output[ScalarFunctionProperties]{
		OutputState: i.ToScalarFunctionPropertiesOutputWithContext(ctx).OutputState,
	}
}

func (i ScalarFunctionPropertiesArgs) ToScalarFunctionPropertiesPtrOutput() ScalarFunctionPropertiesPtrOutput {
	return i.ToScalarFunctionPropertiesPtrOutputWithContext(context.Background())
}

func (i ScalarFunctionPropertiesArgs) ToScalarFunctionPropertiesPtrOutputWithContext(ctx context.Context) ScalarFunctionPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScalarFunctionPropertiesOutput).ToScalarFunctionPropertiesPtrOutputWithContext(ctx)
}

// ScalarFunctionPropertiesPtrInput is an input type that accepts ScalarFunctionPropertiesArgs, ScalarFunctionPropertiesPtr and ScalarFunctionPropertiesPtrOutput values.
// You can construct a concrete instance of `ScalarFunctionPropertiesPtrInput` via:
//
//	        ScalarFunctionPropertiesArgs{...}
//
//	or:
//
//	        nil
type ScalarFunctionPropertiesPtrInput interface {
	pulumi.Input

	ToScalarFunctionPropertiesPtrOutput() ScalarFunctionPropertiesPtrOutput
	ToScalarFunctionPropertiesPtrOutputWithContext(context.Context) ScalarFunctionPropertiesPtrOutput
}

type scalarFunctionPropertiesPtrType ScalarFunctionPropertiesArgs

func ScalarFunctionPropertiesPtr(v *ScalarFunctionPropertiesArgs) ScalarFunctionPropertiesPtrInput {
	return (*scalarFunctionPropertiesPtrType)(v)
}

func (*scalarFunctionPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ScalarFunctionProperties)(nil)).Elem()
}

func (i *scalarFunctionPropertiesPtrType) ToScalarFunctionPropertiesPtrOutput() ScalarFunctionPropertiesPtrOutput {
	return i.ToScalarFunctionPropertiesPtrOutputWithContext(context.Background())
}

func (i *scalarFunctionPropertiesPtrType) ToScalarFunctionPropertiesPtrOutputWithContext(ctx context.Context) ScalarFunctionPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScalarFunctionPropertiesPtrOutput)
}

func (i *scalarFunctionPropertiesPtrType) ToOutput(ctx context.Context) pulumix.Output[*ScalarFunctionProperties] {
	return pulumix.Output[*ScalarFunctionProperties]{
		OutputState: i.ToScalarFunctionPropertiesPtrOutputWithContext(ctx).OutputState,
	}
}

// The properties that are associated with a scalar function.
type ScalarFunctionPropertiesOutput struct{ *pulumi.OutputState }

func (ScalarFunctionPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScalarFunctionProperties)(nil)).Elem()
}

func (o ScalarFunctionPropertiesOutput) ToScalarFunctionPropertiesOutput() ScalarFunctionPropertiesOutput {
	return o
}

func (o ScalarFunctionPropertiesOutput) ToScalarFunctionPropertiesOutputWithContext(ctx context.Context) ScalarFunctionPropertiesOutput {
	return o
}

func (o ScalarFunctionPropertiesOutput) ToScalarFunctionPropertiesPtrOutput() ScalarFunctionPropertiesPtrOutput {
	return o.ToScalarFunctionPropertiesPtrOutputWithContext(context.Background())
}

func (o ScalarFunctionPropertiesOutput) ToScalarFunctionPropertiesPtrOutputWithContext(ctx context.Context) ScalarFunctionPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ScalarFunctionProperties) *ScalarFunctionProperties {
		return &v
	}).(ScalarFunctionPropertiesPtrOutput)
}

func (o ScalarFunctionPropertiesOutput) ToOutput(ctx context.Context) pulumix.Output[ScalarFunctionProperties] {
	return pulumix.Output[ScalarFunctionProperties]{
		OutputState: o.OutputState,
	}
}

// The physical binding of the function. For example, in the Azure Machine Learning web service’s case, this describes the endpoint.
func (o ScalarFunctionPropertiesOutput) Binding() pulumi.AnyOutput {
	return o.ApplyT(func(v ScalarFunctionProperties) interface{} { return v.Binding }).(pulumi.AnyOutput)
}

// A list of inputs describing the parameters of the function.
func (o ScalarFunctionPropertiesOutput) Inputs() FunctionInputTypeArrayOutput {
	return o.ApplyT(func(v ScalarFunctionProperties) []FunctionInputType { return v.Inputs }).(FunctionInputTypeArrayOutput)
}

// The output of the function.
func (o ScalarFunctionPropertiesOutput) Output() FunctionOutputTypePtrOutput {
	return o.ApplyT(func(v ScalarFunctionProperties) *FunctionOutputType { return v.Output }).(FunctionOutputTypePtrOutput)
}

// Indicates the type of function.
// Expected value is 'Scalar'.
func (o ScalarFunctionPropertiesOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ScalarFunctionProperties) string { return v.Type }).(pulumi.StringOutput)
}

type ScalarFunctionPropertiesPtrOutput struct{ *pulumi.OutputState }

func (ScalarFunctionPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScalarFunctionProperties)(nil)).Elem()
}

func (o ScalarFunctionPropertiesPtrOutput) ToScalarFunctionPropertiesPtrOutput() ScalarFunctionPropertiesPtrOutput {
	return o
}

func (o ScalarFunctionPropertiesPtrOutput) ToScalarFunctionPropertiesPtrOutputWithContext(ctx context.Context) ScalarFunctionPropertiesPtrOutput {
	return o
}

func (o ScalarFunctionPropertiesPtrOutput) ToOutput(ctx context.Context) pulumix.Output[*ScalarFunctionProperties] {
	return pulumix.Output[*ScalarFunctionProperties]{
		OutputState: o.OutputState,
	}
}

func (o ScalarFunctionPropertiesPtrOutput) Elem() ScalarFunctionPropertiesOutput {
	return o.ApplyT(func(v *ScalarFunctionProperties) ScalarFunctionProperties {
		if v != nil {
			return *v
		}
		var ret ScalarFunctionProperties
		return ret
	}).(ScalarFunctionPropertiesOutput)
}

// The physical binding of the function. For example, in the Azure Machine Learning web service’s case, this describes the endpoint.
func (o ScalarFunctionPropertiesPtrOutput) Binding() pulumi.AnyOutput {
	return o.ApplyT(func(v *ScalarFunctionProperties) interface{} {
		if v == nil {
			return nil
		}
		return v.Binding
	}).(pulumi.AnyOutput)
}

// A list of inputs describing the parameters of the function.
func (o ScalarFunctionPropertiesPtrOutput) Inputs() FunctionInputTypeArrayOutput {
	return o.ApplyT(func(v *ScalarFunctionProperties) []FunctionInputType {
		if v == nil {
			return nil
		}
		return v.Inputs
	}).(FunctionInputTypeArrayOutput)
}

// The output of the function.
func (o ScalarFunctionPropertiesPtrOutput) Output() FunctionOutputTypePtrOutput {
	return o.ApplyT(func(v *ScalarFunctionProperties) *FunctionOutputType {
		if v == nil {
			return nil
		}
		return v.Output
	}).(FunctionOutputTypePtrOutput)
}

// Indicates the type of function.
// Expected value is 'Scalar'.
func (o ScalarFunctionPropertiesPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalarFunctionProperties) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

// The properties that are associated with a scalar function.
type ScalarFunctionPropertiesResponse struct {
	// The physical binding of the function. For example, in the Azure Machine Learning web service’s case, this describes the endpoint.
	Binding interface{} `pulumi:"binding"`
	// The current entity tag for the function. This is an opaque string. You can use it to detect whether the resource has changed between requests. You can also use it in the If-Match or If-None-Match headers for write operations for optimistic concurrency.
	Etag string `pulumi:"etag"`
	// A list of inputs describing the parameters of the function.
	Inputs []FunctionInputResponse `pulumi:"inputs"`
	// The output of the function.
	Output *FunctionOutputResponse `pulumi:"output"`
	// Indicates the type of function.
	// Expected value is 'Scalar'.
	Type string `pulumi:"type"`
}

// The properties that are associated with a scalar function.
type ScalarFunctionPropertiesResponseOutput struct{ *pulumi.OutputState }

func (ScalarFunctionPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScalarFunctionPropertiesResponse)(nil)).Elem()
}

func (o ScalarFunctionPropertiesResponseOutput) ToScalarFunctionPropertiesResponseOutput() ScalarFunctionPropertiesResponseOutput {
	return o
}

func (o ScalarFunctionPropertiesResponseOutput) ToScalarFunctionPropertiesResponseOutputWithContext(ctx context.Context) ScalarFunctionPropertiesResponseOutput {
	return o
}

func (o ScalarFunctionPropertiesResponseOutput) ToOutput(ctx context.Context) pulumix.Output[ScalarFunctionPropertiesResponse] {
	return pulumix.Output[ScalarFunctionPropertiesResponse]{
		OutputState: o.OutputState,
	}
}

// The physical binding of the function. For example, in the Azure Machine Learning web service’s case, this describes the endpoint.
func (o ScalarFunctionPropertiesResponseOutput) Binding() pulumi.AnyOutput {
	return o.ApplyT(func(v ScalarFunctionPropertiesResponse) interface{} { return v.Binding }).(pulumi.AnyOutput)
}

// The current entity tag for the function. This is an opaque string. You can use it to detect whether the resource has changed between requests. You can also use it in the If-Match or If-None-Match headers for write operations for optimistic concurrency.
func (o ScalarFunctionPropertiesResponseOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v ScalarFunctionPropertiesResponse) string { return v.Etag }).(pulumi.StringOutput)
}

// A list of inputs describing the parameters of the function.
func (o ScalarFunctionPropertiesResponseOutput) Inputs() FunctionInputResponseArrayOutput {
	return o.ApplyT(func(v ScalarFunctionPropertiesResponse) []FunctionInputResponse { return v.Inputs }).(FunctionInputResponseArrayOutput)
}

// The output of the function.
func (o ScalarFunctionPropertiesResponseOutput) Output() FunctionOutputResponsePtrOutput {
	return o.ApplyT(func(v ScalarFunctionPropertiesResponse) *FunctionOutputResponse { return v.Output }).(FunctionOutputResponsePtrOutput)
}

// Indicates the type of function.
// Expected value is 'Scalar'.
func (o ScalarFunctionPropertiesResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ScalarFunctionPropertiesResponse) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(FunctionInputTypeOutput{})
	pulumi.RegisterOutputType(FunctionInputTypeArrayOutput{})
	pulumi.RegisterOutputType(FunctionInputResponseOutput{})
	pulumi.RegisterOutputType(FunctionInputResponseArrayOutput{})
	pulumi.RegisterOutputType(FunctionOutputTypeOutput{})
	pulumi.RegisterOutputType(FunctionOutputTypePtrOutput{})
	pulumi.RegisterOutputType(FunctionOutputResponseOutput{})
	pulumi.RegisterOutputType(FunctionOutputResponsePtrOutput{})
	pulumi.RegisterOutputType(ScalarFunctionPropertiesOutput{})
	pulumi.RegisterOutputType(ScalarFunctionPropertiesPtrOutput{})
	pulumi.RegisterOutputType(ScalarFunctionPropertiesResponseOutput{})
}
