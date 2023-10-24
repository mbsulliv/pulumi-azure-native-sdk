// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package web

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Description for This is to allow calling via powershell and ARM template.
// Azure REST API version: 2022-09-01.
//
// Other available API versions: 2016-08-01, 2018-11-01, 2020-10-01.
func ListWebAppSyncFunctionTriggers(ctx *pulumi.Context, args *ListWebAppSyncFunctionTriggersArgs, opts ...pulumi.InvokeOption) (*ListWebAppSyncFunctionTriggersResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListWebAppSyncFunctionTriggersResult
	err := ctx.Invoke("azure-native:web:listWebAppSyncFunctionTriggers", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWebAppSyncFunctionTriggersArgs struct {
	// Name of the app.
	Name string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Function secrets.
type ListWebAppSyncFunctionTriggersResult struct {
	// Secret key.
	Key *string `pulumi:"key"`
	// Trigger URL.
	TriggerUrl *string `pulumi:"triggerUrl"`
}

func ListWebAppSyncFunctionTriggersOutput(ctx *pulumi.Context, args ListWebAppSyncFunctionTriggersOutputArgs, opts ...pulumi.InvokeOption) ListWebAppSyncFunctionTriggersResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListWebAppSyncFunctionTriggersResult, error) {
			args := v.(ListWebAppSyncFunctionTriggersArgs)
			r, err := ListWebAppSyncFunctionTriggers(ctx, &args, opts...)
			var s ListWebAppSyncFunctionTriggersResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListWebAppSyncFunctionTriggersResultOutput)
}

type ListWebAppSyncFunctionTriggersOutputArgs struct {
	// Name of the app.
	Name pulumi.StringInput `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListWebAppSyncFunctionTriggersOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppSyncFunctionTriggersArgs)(nil)).Elem()
}

// Function secrets.
type ListWebAppSyncFunctionTriggersResultOutput struct{ *pulumi.OutputState }

func (ListWebAppSyncFunctionTriggersResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppSyncFunctionTriggersResult)(nil)).Elem()
}

func (o ListWebAppSyncFunctionTriggersResultOutput) ToListWebAppSyncFunctionTriggersResultOutput() ListWebAppSyncFunctionTriggersResultOutput {
	return o
}

func (o ListWebAppSyncFunctionTriggersResultOutput) ToListWebAppSyncFunctionTriggersResultOutputWithContext(ctx context.Context) ListWebAppSyncFunctionTriggersResultOutput {
	return o
}

func (o ListWebAppSyncFunctionTriggersResultOutput) ToOutput(ctx context.Context) pulumix.Output[ListWebAppSyncFunctionTriggersResult] {
	return pulumix.Output[ListWebAppSyncFunctionTriggersResult]{
		OutputState: o.OutputState,
	}
}

// Secret key.
func (o ListWebAppSyncFunctionTriggersResultOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppSyncFunctionTriggersResult) *string { return v.Key }).(pulumi.StringPtrOutput)
}

// Trigger URL.
func (o ListWebAppSyncFunctionTriggersResultOutput) TriggerUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppSyncFunctionTriggersResult) *string { return v.TriggerUrl }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListWebAppSyncFunctionTriggersResultOutput{})
}
