// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package logic

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets a workflow access key.
// Azure REST API version: 2015-02-01-preview.
func LookupWorkflowAccessKey(ctx *pulumi.Context, args *LookupWorkflowAccessKeyArgs, opts ...pulumi.InvokeOption) (*LookupWorkflowAccessKeyResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupWorkflowAccessKeyResult
	err := ctx.Invoke("azure-native:logic:getWorkflowAccessKey", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWorkflowAccessKeyArgs struct {
	// The workflow access key name.
	AccessKeyName string `pulumi:"accessKeyName"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The workflow name.
	WorkflowName string `pulumi:"workflowName"`
}

type LookupWorkflowAccessKeyResult struct {
	// Gets or sets the resource id.
	Id *string `pulumi:"id"`
	// Gets the workflow access key name.
	Name string `pulumi:"name"`
	// Gets or sets the not-after time.
	NotAfter *string `pulumi:"notAfter"`
	// Gets or sets the not-before time.
	NotBefore *string `pulumi:"notBefore"`
	// Gets the workflow access key type.
	Type string `pulumi:"type"`
}

func LookupWorkflowAccessKeyOutput(ctx *pulumi.Context, args LookupWorkflowAccessKeyOutputArgs, opts ...pulumi.InvokeOption) LookupWorkflowAccessKeyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWorkflowAccessKeyResult, error) {
			args := v.(LookupWorkflowAccessKeyArgs)
			r, err := LookupWorkflowAccessKey(ctx, &args, opts...)
			var s LookupWorkflowAccessKeyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWorkflowAccessKeyResultOutput)
}

type LookupWorkflowAccessKeyOutputArgs struct {
	// The workflow access key name.
	AccessKeyName pulumi.StringInput `pulumi:"accessKeyName"`
	// The resource group name.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The workflow name.
	WorkflowName pulumi.StringInput `pulumi:"workflowName"`
}

func (LookupWorkflowAccessKeyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkflowAccessKeyArgs)(nil)).Elem()
}

type LookupWorkflowAccessKeyResultOutput struct{ *pulumi.OutputState }

func (LookupWorkflowAccessKeyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkflowAccessKeyResult)(nil)).Elem()
}

func (o LookupWorkflowAccessKeyResultOutput) ToLookupWorkflowAccessKeyResultOutput() LookupWorkflowAccessKeyResultOutput {
	return o
}

func (o LookupWorkflowAccessKeyResultOutput) ToLookupWorkflowAccessKeyResultOutputWithContext(ctx context.Context) LookupWorkflowAccessKeyResultOutput {
	return o
}

func (o LookupWorkflowAccessKeyResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupWorkflowAccessKeyResult] {
	return pulumix.Output[LookupWorkflowAccessKeyResult]{
		OutputState: o.OutputState,
	}
}

// Gets or sets the resource id.
func (o LookupWorkflowAccessKeyResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkflowAccessKeyResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// Gets the workflow access key name.
func (o LookupWorkflowAccessKeyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkflowAccessKeyResult) string { return v.Name }).(pulumi.StringOutput)
}

// Gets or sets the not-after time.
func (o LookupWorkflowAccessKeyResultOutput) NotAfter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkflowAccessKeyResult) *string { return v.NotAfter }).(pulumi.StringPtrOutput)
}

// Gets or sets the not-before time.
func (o LookupWorkflowAccessKeyResultOutput) NotBefore() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkflowAccessKeyResult) *string { return v.NotBefore }).(pulumi.StringPtrOutput)
}

// Gets the workflow access key type.
func (o LookupWorkflowAccessKeyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkflowAccessKeyResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWorkflowAccessKeyResultOutput{})
}
