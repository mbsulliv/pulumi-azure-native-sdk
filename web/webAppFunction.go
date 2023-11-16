// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package web

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Function information.
// Azure REST API version: 2022-09-01. Prior API version in Azure Native 1.x: 2020-12-01.
//
// Other available API versions: 2016-08-01, 2020-10-01, 2023-01-01.
type WebAppFunction struct {
	pulumi.CustomResourceState

	// Config information.
	Config pulumi.AnyOutput `pulumi:"config"`
	// Config URI.
	ConfigHref pulumi.StringPtrOutput `pulumi:"configHref"`
	// File list.
	Files pulumi.StringMapOutput `pulumi:"files"`
	// Function App ID.
	FunctionAppId pulumi.StringPtrOutput `pulumi:"functionAppId"`
	// Function URI.
	Href pulumi.StringPtrOutput `pulumi:"href"`
	// The invocation URL
	InvokeUrlTemplate pulumi.StringPtrOutput `pulumi:"invokeUrlTemplate"`
	// Gets or sets a value indicating whether the function is disabled
	IsDisabled pulumi.BoolPtrOutput `pulumi:"isDisabled"`
	// Kind of resource.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// The function language
	Language pulumi.StringPtrOutput `pulumi:"language"`
	// Resource Name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Script URI.
	ScriptHref pulumi.StringPtrOutput `pulumi:"scriptHref"`
	// Script root path URI.
	ScriptRootPathHref pulumi.StringPtrOutput `pulumi:"scriptRootPathHref"`
	// Secrets file URI.
	SecretsFileHref pulumi.StringPtrOutput `pulumi:"secretsFileHref"`
	// Test data used when testing via the Azure Portal.
	TestData pulumi.StringPtrOutput `pulumi:"testData"`
	// Test data URI.
	TestDataHref pulumi.StringPtrOutput `pulumi:"testDataHref"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewWebAppFunction registers a new resource with the given unique name, arguments, and options.
func NewWebAppFunction(ctx *pulumi.Context,
	name string, args *WebAppFunctionArgs, opts ...pulumi.ResourceOption) (*WebAppFunction, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:web/v20160801:WebAppFunction"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:WebAppFunction"),
		},
		{
			Type: pulumi.String("azure-native:web/v20181101:WebAppFunction"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:WebAppFunction"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:WebAppFunction"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:WebAppFunction"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:WebAppFunction"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:WebAppFunction"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:WebAppFunction"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:WebAppFunction"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:WebAppFunction"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210301:WebAppFunction"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220301:WebAppFunction"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220901:WebAppFunction"),
		},
		{
			Type: pulumi.String("azure-native:web/v20230101:WebAppFunction"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource WebAppFunction
	err := ctx.RegisterResource("azure-native:web:WebAppFunction", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebAppFunction gets an existing WebAppFunction resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebAppFunction(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppFunctionState, opts ...pulumi.ResourceOption) (*WebAppFunction, error) {
	var resource WebAppFunction
	err := ctx.ReadResource("azure-native:web:WebAppFunction", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebAppFunction resources.
type webAppFunctionState struct {
}

type WebAppFunctionState struct {
}

func (WebAppFunctionState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppFunctionState)(nil)).Elem()
}

type webAppFunctionArgs struct {
	// Config information.
	Config interface{} `pulumi:"config"`
	// Config URI.
	ConfigHref *string `pulumi:"configHref"`
	// File list.
	Files map[string]string `pulumi:"files"`
	// Function App ID.
	FunctionAppId *string `pulumi:"functionAppId"`
	// Function name.
	FunctionName *string `pulumi:"functionName"`
	// Function URI.
	Href *string `pulumi:"href"`
	// The invocation URL
	InvokeUrlTemplate *string `pulumi:"invokeUrlTemplate"`
	// Gets or sets a value indicating whether the function is disabled
	IsDisabled *bool `pulumi:"isDisabled"`
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// The function language
	Language *string `pulumi:"language"`
	// Site name.
	Name string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
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
}

// The set of arguments for constructing a WebAppFunction resource.
type WebAppFunctionArgs struct {
	// Config information.
	Config pulumi.Input
	// Config URI.
	ConfigHref pulumi.StringPtrInput
	// File list.
	Files pulumi.StringMapInput
	// Function App ID.
	FunctionAppId pulumi.StringPtrInput
	// Function name.
	FunctionName pulumi.StringPtrInput
	// Function URI.
	Href pulumi.StringPtrInput
	// The invocation URL
	InvokeUrlTemplate pulumi.StringPtrInput
	// Gets or sets a value indicating whether the function is disabled
	IsDisabled pulumi.BoolPtrInput
	// Kind of resource.
	Kind pulumi.StringPtrInput
	// The function language
	Language pulumi.StringPtrInput
	// Site name.
	Name pulumi.StringInput
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput
	// Script URI.
	ScriptHref pulumi.StringPtrInput
	// Script root path URI.
	ScriptRootPathHref pulumi.StringPtrInput
	// Secrets file URI.
	SecretsFileHref pulumi.StringPtrInput
	// Test data used when testing via the Azure Portal.
	TestData pulumi.StringPtrInput
	// Test data URI.
	TestDataHref pulumi.StringPtrInput
}

func (WebAppFunctionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppFunctionArgs)(nil)).Elem()
}

type WebAppFunctionInput interface {
	pulumi.Input

	ToWebAppFunctionOutput() WebAppFunctionOutput
	ToWebAppFunctionOutputWithContext(ctx context.Context) WebAppFunctionOutput
}

func (*WebAppFunction) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppFunction)(nil)).Elem()
}

func (i *WebAppFunction) ToWebAppFunctionOutput() WebAppFunctionOutput {
	return i.ToWebAppFunctionOutputWithContext(context.Background())
}

func (i *WebAppFunction) ToWebAppFunctionOutputWithContext(ctx context.Context) WebAppFunctionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAppFunctionOutput)
}

type WebAppFunctionOutput struct{ *pulumi.OutputState }

func (WebAppFunctionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppFunction)(nil)).Elem()
}

func (o WebAppFunctionOutput) ToWebAppFunctionOutput() WebAppFunctionOutput {
	return o
}

func (o WebAppFunctionOutput) ToWebAppFunctionOutputWithContext(ctx context.Context) WebAppFunctionOutput {
	return o
}

// Config information.
func (o WebAppFunctionOutput) Config() pulumi.AnyOutput {
	return o.ApplyT(func(v *WebAppFunction) pulumi.AnyOutput { return v.Config }).(pulumi.AnyOutput)
}

// Config URI.
func (o WebAppFunctionOutput) ConfigHref() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppFunction) pulumi.StringPtrOutput { return v.ConfigHref }).(pulumi.StringPtrOutput)
}

// File list.
func (o WebAppFunctionOutput) Files() pulumi.StringMapOutput {
	return o.ApplyT(func(v *WebAppFunction) pulumi.StringMapOutput { return v.Files }).(pulumi.StringMapOutput)
}

// Function App ID.
func (o WebAppFunctionOutput) FunctionAppId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppFunction) pulumi.StringPtrOutput { return v.FunctionAppId }).(pulumi.StringPtrOutput)
}

// Function URI.
func (o WebAppFunctionOutput) Href() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppFunction) pulumi.StringPtrOutput { return v.Href }).(pulumi.StringPtrOutput)
}

// The invocation URL
func (o WebAppFunctionOutput) InvokeUrlTemplate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppFunction) pulumi.StringPtrOutput { return v.InvokeUrlTemplate }).(pulumi.StringPtrOutput)
}

// Gets or sets a value indicating whether the function is disabled
func (o WebAppFunctionOutput) IsDisabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WebAppFunction) pulumi.BoolPtrOutput { return v.IsDisabled }).(pulumi.BoolPtrOutput)
}

// Kind of resource.
func (o WebAppFunctionOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppFunction) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// The function language
func (o WebAppFunctionOutput) Language() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppFunction) pulumi.StringPtrOutput { return v.Language }).(pulumi.StringPtrOutput)
}

// Resource Name.
func (o WebAppFunctionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppFunction) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Script URI.
func (o WebAppFunctionOutput) ScriptHref() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppFunction) pulumi.StringPtrOutput { return v.ScriptHref }).(pulumi.StringPtrOutput)
}

// Script root path URI.
func (o WebAppFunctionOutput) ScriptRootPathHref() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppFunction) pulumi.StringPtrOutput { return v.ScriptRootPathHref }).(pulumi.StringPtrOutput)
}

// Secrets file URI.
func (o WebAppFunctionOutput) SecretsFileHref() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppFunction) pulumi.StringPtrOutput { return v.SecretsFileHref }).(pulumi.StringPtrOutput)
}

// Test data used when testing via the Azure Portal.
func (o WebAppFunctionOutput) TestData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppFunction) pulumi.StringPtrOutput { return v.TestData }).(pulumi.StringPtrOutput)
}

// Test data URI.
func (o WebAppFunctionOutput) TestDataHref() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppFunction) pulumi.StringPtrOutput { return v.TestDataHref }).(pulumi.StringPtrOutput)
}

// Resource type.
func (o WebAppFunctionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppFunction) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WebAppFunctionOutput{})
}
