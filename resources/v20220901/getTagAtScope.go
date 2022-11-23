// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220901

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Wrapper resource for tags API requests and responses.
func LookupTagAtScope(ctx *pulumi.Context, args *LookupTagAtScopeArgs, opts ...pulumi.InvokeOption) (*LookupTagAtScopeResult, error) {
	var rv LookupTagAtScopeResult
	err := ctx.Invoke("azure-native:resources/v20220901:getTagAtScope", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTagAtScopeArgs struct {
	// The resource scope.
	Scope string `pulumi:"scope"`
}

// Wrapper resource for tags API requests and responses.
type LookupTagAtScopeResult struct {
	// The ID of the tags wrapper resource.
	Id string `pulumi:"id"`
	// The name of the tags wrapper resource.
	Name string `pulumi:"name"`
	// The set of tags.
	Properties TagsResponse `pulumi:"properties"`
	// The type of the tags wrapper resource.
	Type string `pulumi:"type"`
}

func LookupTagAtScopeOutput(ctx *pulumi.Context, args LookupTagAtScopeOutputArgs, opts ...pulumi.InvokeOption) LookupTagAtScopeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTagAtScopeResult, error) {
			args := v.(LookupTagAtScopeArgs)
			r, err := LookupTagAtScope(ctx, &args, opts...)
			var s LookupTagAtScopeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTagAtScopeResultOutput)
}

type LookupTagAtScopeOutputArgs struct {
	// The resource scope.
	Scope pulumi.StringInput `pulumi:"scope"`
}

func (LookupTagAtScopeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTagAtScopeArgs)(nil)).Elem()
}

// Wrapper resource for tags API requests and responses.
type LookupTagAtScopeResultOutput struct{ *pulumi.OutputState }

func (LookupTagAtScopeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTagAtScopeResult)(nil)).Elem()
}

func (o LookupTagAtScopeResultOutput) ToLookupTagAtScopeResultOutput() LookupTagAtScopeResultOutput {
	return o
}

func (o LookupTagAtScopeResultOutput) ToLookupTagAtScopeResultOutputWithContext(ctx context.Context) LookupTagAtScopeResultOutput {
	return o
}

// The ID of the tags wrapper resource.
func (o LookupTagAtScopeResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTagAtScopeResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the tags wrapper resource.
func (o LookupTagAtScopeResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTagAtScopeResult) string { return v.Name }).(pulumi.StringOutput)
}

// The set of tags.
func (o LookupTagAtScopeResultOutput) Properties() TagsResponseOutput {
	return o.ApplyT(func(v LookupTagAtScopeResult) TagsResponse { return v.Properties }).(TagsResponseOutput)
}

// The type of the tags wrapper resource.
func (o LookupTagAtScopeResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTagAtScopeResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTagAtScopeResultOutput{})
}
