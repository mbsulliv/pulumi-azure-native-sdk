// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220901

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Description for Get function information by its ID for web site, or a deployment slot.
func LookupWebAppFunction(ctx *pulumi.Context, args *LookupWebAppFunctionArgs, opts ...pulumi.InvokeOption) (*LookupWebAppFunctionResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupWebAppFunctionResult
	err := ctx.Invoke("azure-native:web/v20220901:getWebAppFunction", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWebAppFunctionArgs struct {
	// Function name.
	FunctionName string `pulumi:"functionName"`
	// Site name.
	Name string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Function information.
type LookupWebAppFunctionResult struct {
	// Config information.
	Config interface{} `pulumi:"config"`
	// Config URI.
	ConfigHref *string `pulumi:"configHref"`
	// File list.
	Files map[string]string `pulumi:"files"`
	// Function App ID.
	FunctionAppId *string `pulumi:"functionAppId"`
	// Function URI.
	Href *string `pulumi:"href"`
	// Resource Id.
	Id string `pulumi:"id"`
	// The invocation URL
	InvokeUrlTemplate *string `pulumi:"invokeUrlTemplate"`
	// Gets or sets a value indicating whether the function is disabled
	IsDisabled *bool `pulumi:"isDisabled"`
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// The function language
	Language *string `pulumi:"language"`
	// Resource Name.
	Name string `pulumi:"name"`
	// Script URI.
	ScriptHref *string `pulumi:"scriptHref"`
	// Script root path URI.
	ScriptRootPathHref *string `pulumi:"scriptRootPathHref"`
	// Secrets file URI.
	SecretsFileHref *string `pulumi:"secretsFileHref"`
	// Test data used when testing via the Azure Portal.
	TestData *string `pulumi:"testData"`
	// Test data URI.
	TestDataHref *string `pulumi:"testDataHref"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupWebAppFunctionOutput(ctx *pulumi.Context, args LookupWebAppFunctionOutputArgs, opts ...pulumi.InvokeOption) LookupWebAppFunctionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWebAppFunctionResult, error) {
			args := v.(LookupWebAppFunctionArgs)
			r, err := LookupWebAppFunction(ctx, &args, opts...)
			var s LookupWebAppFunctionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWebAppFunctionResultOutput)
}

type LookupWebAppFunctionOutputArgs struct {
	// Function name.
	FunctionName pulumi.StringInput `pulumi:"functionName"`
	// Site name.
	Name pulumi.StringInput `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupWebAppFunctionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppFunctionArgs)(nil)).Elem()
}

// Function information.
type LookupWebAppFunctionResultOutput struct{ *pulumi.OutputState }

func (LookupWebAppFunctionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppFunctionResult)(nil)).Elem()
}

func (o LookupWebAppFunctionResultOutput) ToLookupWebAppFunctionResultOutput() LookupWebAppFunctionResultOutput {
	return o
}

func (o LookupWebAppFunctionResultOutput) ToLookupWebAppFunctionResultOutputWithContext(ctx context.Context) LookupWebAppFunctionResultOutput {
	return o
}

// Config information.
func (o LookupWebAppFunctionResultOutput) Config() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupWebAppFunctionResult) interface{} { return v.Config }).(pulumi.AnyOutput)
}

// Config URI.
func (o LookupWebAppFunctionResultOutput) ConfigHref() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppFunctionResult) *string { return v.ConfigHref }).(pulumi.StringPtrOutput)
}

// File list.
func (o LookupWebAppFunctionResultOutput) Files() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupWebAppFunctionResult) map[string]string { return v.Files }).(pulumi.StringMapOutput)
}

// Function App ID.
func (o LookupWebAppFunctionResultOutput) FunctionAppId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppFunctionResult) *string { return v.FunctionAppId }).(pulumi.StringPtrOutput)
}

// Function URI.
func (o LookupWebAppFunctionResultOutput) Href() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppFunctionResult) *string { return v.Href }).(pulumi.StringPtrOutput)
}

// Resource Id.
func (o LookupWebAppFunctionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppFunctionResult) string { return v.Id }).(pulumi.StringOutput)
}

// The invocation URL
func (o LookupWebAppFunctionResultOutput) InvokeUrlTemplate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppFunctionResult) *string { return v.InvokeUrlTemplate }).(pulumi.StringPtrOutput)
}

// Gets or sets a value indicating whether the function is disabled
func (o LookupWebAppFunctionResultOutput) IsDisabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupWebAppFunctionResult) *bool { return v.IsDisabled }).(pulumi.BoolPtrOutput)
}

// Kind of resource.
func (o LookupWebAppFunctionResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppFunctionResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// The function language
func (o LookupWebAppFunctionResultOutput) Language() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppFunctionResult) *string { return v.Language }).(pulumi.StringPtrOutput)
}

// Resource Name.
func (o LookupWebAppFunctionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppFunctionResult) string { return v.Name }).(pulumi.StringOutput)
}

// Script URI.
func (o LookupWebAppFunctionResultOutput) ScriptHref() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppFunctionResult) *string { return v.ScriptHref }).(pulumi.StringPtrOutput)
}

// Script root path URI.
func (o LookupWebAppFunctionResultOutput) ScriptRootPathHref() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppFunctionResult) *string { return v.ScriptRootPathHref }).(pulumi.StringPtrOutput)
}

// Secrets file URI.
func (o LookupWebAppFunctionResultOutput) SecretsFileHref() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppFunctionResult) *string { return v.SecretsFileHref }).(pulumi.StringPtrOutput)
}

// Test data used when testing via the Azure Portal.
func (o LookupWebAppFunctionResultOutput) TestData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppFunctionResult) *string { return v.TestData }).(pulumi.StringPtrOutput)
}

// Test data URI.
func (o LookupWebAppFunctionResultOutput) TestDataHref() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppFunctionResult) *string { return v.TestDataHref }).(pulumi.StringPtrOutput)
}

// Resource type.
func (o LookupWebAppFunctionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppFunctionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWebAppFunctionResultOutput{})
}
