// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Description for Get function information by its ID for web site, or a deployment slot.
func LookupWebAppInstanceFunctionSlot(ctx *pulumi.Context, args *LookupWebAppInstanceFunctionSlotArgs, opts ...pulumi.InvokeOption) (*LookupWebAppInstanceFunctionSlotResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupWebAppInstanceFunctionSlotResult
	err := ctx.Invoke("azure-native:web/v20230101:getWebAppInstanceFunctionSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWebAppInstanceFunctionSlotArgs struct {
	// Function name.
	FunctionName string `pulumi:"functionName"`
	// Site name.
	Name string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the deployment slot.
	Slot string `pulumi:"slot"`
}

// Function information.
type LookupWebAppInstanceFunctionSlotResult struct {
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

func LookupWebAppInstanceFunctionSlotOutput(ctx *pulumi.Context, args LookupWebAppInstanceFunctionSlotOutputArgs, opts ...pulumi.InvokeOption) LookupWebAppInstanceFunctionSlotResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWebAppInstanceFunctionSlotResult, error) {
			args := v.(LookupWebAppInstanceFunctionSlotArgs)
			r, err := LookupWebAppInstanceFunctionSlot(ctx, &args, opts...)
			var s LookupWebAppInstanceFunctionSlotResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWebAppInstanceFunctionSlotResultOutput)
}

type LookupWebAppInstanceFunctionSlotOutputArgs struct {
	// Function name.
	FunctionName pulumi.StringInput `pulumi:"functionName"`
	// Site name.
	Name pulumi.StringInput `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name of the deployment slot.
	Slot pulumi.StringInput `pulumi:"slot"`
}

func (LookupWebAppInstanceFunctionSlotOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppInstanceFunctionSlotArgs)(nil)).Elem()
}

// Function information.
type LookupWebAppInstanceFunctionSlotResultOutput struct{ *pulumi.OutputState }

func (LookupWebAppInstanceFunctionSlotResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppInstanceFunctionSlotResult)(nil)).Elem()
}

func (o LookupWebAppInstanceFunctionSlotResultOutput) ToLookupWebAppInstanceFunctionSlotResultOutput() LookupWebAppInstanceFunctionSlotResultOutput {
	return o
}

func (o LookupWebAppInstanceFunctionSlotResultOutput) ToLookupWebAppInstanceFunctionSlotResultOutputWithContext(ctx context.Context) LookupWebAppInstanceFunctionSlotResultOutput {
	return o
}

// Config information.
func (o LookupWebAppInstanceFunctionSlotResultOutput) Config() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupWebAppInstanceFunctionSlotResult) interface{} { return v.Config }).(pulumi.AnyOutput)
}

// Config URI.
func (o LookupWebAppInstanceFunctionSlotResultOutput) ConfigHref() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppInstanceFunctionSlotResult) *string { return v.ConfigHref }).(pulumi.StringPtrOutput)
}

// File list.
func (o LookupWebAppInstanceFunctionSlotResultOutput) Files() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupWebAppInstanceFunctionSlotResult) map[string]string { return v.Files }).(pulumi.StringMapOutput)
}

// Function App ID.
func (o LookupWebAppInstanceFunctionSlotResultOutput) FunctionAppId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppInstanceFunctionSlotResult) *string { return v.FunctionAppId }).(pulumi.StringPtrOutput)
}

// Function URI.
func (o LookupWebAppInstanceFunctionSlotResultOutput) Href() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppInstanceFunctionSlotResult) *string { return v.Href }).(pulumi.StringPtrOutput)
}

// Resource Id.
func (o LookupWebAppInstanceFunctionSlotResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppInstanceFunctionSlotResult) string { return v.Id }).(pulumi.StringOutput)
}

// The invocation URL
func (o LookupWebAppInstanceFunctionSlotResultOutput) InvokeUrlTemplate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppInstanceFunctionSlotResult) *string { return v.InvokeUrlTemplate }).(pulumi.StringPtrOutput)
}

// Gets or sets a value indicating whether the function is disabled
func (o LookupWebAppInstanceFunctionSlotResultOutput) IsDisabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupWebAppInstanceFunctionSlotResult) *bool { return v.IsDisabled }).(pulumi.BoolPtrOutput)
}

// Kind of resource.
func (o LookupWebAppInstanceFunctionSlotResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppInstanceFunctionSlotResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// The function language
func (o LookupWebAppInstanceFunctionSlotResultOutput) Language() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppInstanceFunctionSlotResult) *string { return v.Language }).(pulumi.StringPtrOutput)
}

// Resource Name.
func (o LookupWebAppInstanceFunctionSlotResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppInstanceFunctionSlotResult) string { return v.Name }).(pulumi.StringOutput)
}

// Script URI.
func (o LookupWebAppInstanceFunctionSlotResultOutput) ScriptHref() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppInstanceFunctionSlotResult) *string { return v.ScriptHref }).(pulumi.StringPtrOutput)
}

// Script root path URI.
func (o LookupWebAppInstanceFunctionSlotResultOutput) ScriptRootPathHref() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppInstanceFunctionSlotResult) *string { return v.ScriptRootPathHref }).(pulumi.StringPtrOutput)
}

// Secrets file URI.
func (o LookupWebAppInstanceFunctionSlotResultOutput) SecretsFileHref() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppInstanceFunctionSlotResult) *string { return v.SecretsFileHref }).(pulumi.StringPtrOutput)
}

// Test data used when testing via the Azure Portal.
func (o LookupWebAppInstanceFunctionSlotResultOutput) TestData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppInstanceFunctionSlotResult) *string { return v.TestData }).(pulumi.StringPtrOutput)
}

// Test data URI.
func (o LookupWebAppInstanceFunctionSlotResultOutput) TestDataHref() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppInstanceFunctionSlotResult) *string { return v.TestDataHref }).(pulumi.StringPtrOutput)
}

// Resource type.
func (o LookupWebAppInstanceFunctionSlotResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppInstanceFunctionSlotResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWebAppInstanceFunctionSlotResultOutput{})
}