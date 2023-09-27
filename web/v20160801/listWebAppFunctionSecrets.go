// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20160801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get function secrets for a function in a web site, or a deployment slot.
func ListWebAppFunctionSecrets(ctx *pulumi.Context, args *ListWebAppFunctionSecretsArgs, opts ...pulumi.InvokeOption) (*ListWebAppFunctionSecretsResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListWebAppFunctionSecretsResult
	err := ctx.Invoke("azure-native:web/v20160801:listWebAppFunctionSecrets", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWebAppFunctionSecretsArgs struct {
	// Function name.
	FunctionName string `pulumi:"functionName"`
	// Site name.
	Name string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Function secrets.
type ListWebAppFunctionSecretsResult struct {
	// Resource Id.
	Id string `pulumi:"id"`
	// Secret key.
	Key *string `pulumi:"key"`
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Resource Name.
	Name string `pulumi:"name"`
	// Trigger URL.
	TriggerUrl *string `pulumi:"triggerUrl"`
	// Resource type.
	Type string `pulumi:"type"`
}

func ListWebAppFunctionSecretsOutput(ctx *pulumi.Context, args ListWebAppFunctionSecretsOutputArgs, opts ...pulumi.InvokeOption) ListWebAppFunctionSecretsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListWebAppFunctionSecretsResult, error) {
			args := v.(ListWebAppFunctionSecretsArgs)
			r, err := ListWebAppFunctionSecrets(ctx, &args, opts...)
			var s ListWebAppFunctionSecretsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListWebAppFunctionSecretsResultOutput)
}

type ListWebAppFunctionSecretsOutputArgs struct {
	// Function name.
	FunctionName pulumi.StringInput `pulumi:"functionName"`
	// Site name.
	Name pulumi.StringInput `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListWebAppFunctionSecretsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppFunctionSecretsArgs)(nil)).Elem()
}

// Function secrets.
type ListWebAppFunctionSecretsResultOutput struct{ *pulumi.OutputState }

func (ListWebAppFunctionSecretsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppFunctionSecretsResult)(nil)).Elem()
}

func (o ListWebAppFunctionSecretsResultOutput) ToListWebAppFunctionSecretsResultOutput() ListWebAppFunctionSecretsResultOutput {
	return o
}

func (o ListWebAppFunctionSecretsResultOutput) ToListWebAppFunctionSecretsResultOutputWithContext(ctx context.Context) ListWebAppFunctionSecretsResultOutput {
	return o
}

func (o ListWebAppFunctionSecretsResultOutput) ToOutput(ctx context.Context) pulumix.Output[ListWebAppFunctionSecretsResult] {
	return pulumix.Output[ListWebAppFunctionSecretsResult]{
		OutputState: o.OutputState,
	}
}

// Resource Id.
func (o ListWebAppFunctionSecretsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppFunctionSecretsResult) string { return v.Id }).(pulumi.StringOutput)
}

// Secret key.
func (o ListWebAppFunctionSecretsResultOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppFunctionSecretsResult) *string { return v.Key }).(pulumi.StringPtrOutput)
}

// Kind of resource.
func (o ListWebAppFunctionSecretsResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppFunctionSecretsResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// Resource Name.
func (o ListWebAppFunctionSecretsResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppFunctionSecretsResult) string { return v.Name }).(pulumi.StringOutput)
}

// Trigger URL.
func (o ListWebAppFunctionSecretsResultOutput) TriggerUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppFunctionSecretsResult) *string { return v.TriggerUrl }).(pulumi.StringPtrOutput)
}

// Resource type.
func (o ListWebAppFunctionSecretsResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppFunctionSecretsResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListWebAppFunctionSecretsResultOutput{})
}
