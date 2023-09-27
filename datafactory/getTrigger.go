// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package datafactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets a trigger.
// Azure REST API version: 2018-06-01.
func LookupTrigger(ctx *pulumi.Context, args *LookupTriggerArgs, opts ...pulumi.InvokeOption) (*LookupTriggerResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupTriggerResult
	err := ctx.Invoke("azure-native:datafactory:getTrigger", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTriggerArgs struct {
	// The factory name.
	FactoryName string `pulumi:"factoryName"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The trigger name.
	TriggerName string `pulumi:"triggerName"`
}

// Trigger resource type.
type LookupTriggerResult struct {
	// Etag identifies change in the resource.
	Etag string `pulumi:"etag"`
	// The resource identifier.
	Id string `pulumi:"id"`
	// The resource name.
	Name string `pulumi:"name"`
	// Properties of the trigger.
	Properties interface{} `pulumi:"properties"`
	// The resource type.
	Type string `pulumi:"type"`
}

func LookupTriggerOutput(ctx *pulumi.Context, args LookupTriggerOutputArgs, opts ...pulumi.InvokeOption) LookupTriggerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTriggerResult, error) {
			args := v.(LookupTriggerArgs)
			r, err := LookupTrigger(ctx, &args, opts...)
			var s LookupTriggerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTriggerResultOutput)
}

type LookupTriggerOutputArgs struct {
	// The factory name.
	FactoryName pulumi.StringInput `pulumi:"factoryName"`
	// The resource group name.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The trigger name.
	TriggerName pulumi.StringInput `pulumi:"triggerName"`
}

func (LookupTriggerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTriggerArgs)(nil)).Elem()
}

// Trigger resource type.
type LookupTriggerResultOutput struct{ *pulumi.OutputState }

func (LookupTriggerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTriggerResult)(nil)).Elem()
}

func (o LookupTriggerResultOutput) ToLookupTriggerResultOutput() LookupTriggerResultOutput {
	return o
}

func (o LookupTriggerResultOutput) ToLookupTriggerResultOutputWithContext(ctx context.Context) LookupTriggerResultOutput {
	return o
}

func (o LookupTriggerResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupTriggerResult] {
	return pulumix.Output[LookupTriggerResult]{
		OutputState: o.OutputState,
	}
}

// Etag identifies change in the resource.
func (o LookupTriggerResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTriggerResult) string { return v.Etag }).(pulumi.StringOutput)
}

// The resource identifier.
func (o LookupTriggerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTriggerResult) string { return v.Id }).(pulumi.StringOutput)
}

// The resource name.
func (o LookupTriggerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTriggerResult) string { return v.Name }).(pulumi.StringOutput)
}

// Properties of the trigger.
func (o LookupTriggerResultOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupTriggerResult) interface{} { return v.Properties }).(pulumi.AnyOutput)
}

// The resource type.
func (o LookupTriggerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTriggerResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTriggerResultOutput{})
}
