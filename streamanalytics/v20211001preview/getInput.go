// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20211001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets details about the specified input.
func LookupInput(ctx *pulumi.Context, args *LookupInputArgs, opts ...pulumi.InvokeOption) (*LookupInputResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupInputResult
	err := ctx.Invoke("azure-native:streamanalytics/v20211001preview:getInput", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupInputArgs struct {
	// The name of the input.
	InputName string `pulumi:"inputName"`
	// The name of the streaming job.
	JobName string `pulumi:"jobName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// An input object, containing all information associated with the named input. All inputs are contained under a streaming job.
type LookupInputResult struct {
	// Resource Id
	Id string `pulumi:"id"`
	// Resource name
	Name *string `pulumi:"name"`
	// The properties that are associated with an input. Required on PUT (CreateOrReplace) requests.
	Properties interface{} `pulumi:"properties"`
	// Resource type
	Type string `pulumi:"type"`
}

func LookupInputOutput(ctx *pulumi.Context, args LookupInputOutputArgs, opts ...pulumi.InvokeOption) LookupInputResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupInputResult, error) {
			args := v.(LookupInputArgs)
			r, err := LookupInput(ctx, &args, opts...)
			var s LookupInputResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupInputResultOutput)
}

type LookupInputOutputArgs struct {
	// The name of the input.
	InputName pulumi.StringInput `pulumi:"inputName"`
	// The name of the streaming job.
	JobName pulumi.StringInput `pulumi:"jobName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupInputOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInputArgs)(nil)).Elem()
}

// An input object, containing all information associated with the named input. All inputs are contained under a streaming job.
type LookupInputResultOutput struct{ *pulumi.OutputState }

func (LookupInputResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInputResult)(nil)).Elem()
}

func (o LookupInputResultOutput) ToLookupInputResultOutput() LookupInputResultOutput {
	return o
}

func (o LookupInputResultOutput) ToLookupInputResultOutputWithContext(ctx context.Context) LookupInputResultOutput {
	return o
}

func (o LookupInputResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupInputResult] {
	return pulumix.Output[LookupInputResult]{
		OutputState: o.OutputState,
	}
}

// Resource Id
func (o LookupInputResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInputResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource name
func (o LookupInputResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInputResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The properties that are associated with an input. Required on PUT (CreateOrReplace) requests.
func (o LookupInputResultOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupInputResult) interface{} { return v.Properties }).(pulumi.AnyOutput)
}

// Resource type
func (o LookupInputResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInputResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupInputResultOutput{})
}
