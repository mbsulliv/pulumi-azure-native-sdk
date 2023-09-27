// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package customerinsights

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets a view in the hub.
// Azure REST API version: 2017-04-26.
func LookupView(ctx *pulumi.Context, args *LookupViewArgs, opts ...pulumi.InvokeOption) (*LookupViewResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupViewResult
	err := ctx.Invoke("azure-native:customerinsights:getView", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupViewArgs struct {
	// The name of the hub.
	HubName string `pulumi:"hubName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The user ID. Use * to retrieve hub level view.
	UserId string `pulumi:"userId"`
	// The name of the view.
	ViewName string `pulumi:"viewName"`
}

// The view resource format.
type LookupViewResult struct {
	// Date time when view was last modified.
	Changed string `pulumi:"changed"`
	// Date time when view was created.
	Created string `pulumi:"created"`
	// View definition.
	Definition string `pulumi:"definition"`
	// Localized display name for the view.
	DisplayName map[string]string `pulumi:"displayName"`
	// Resource ID.
	Id string `pulumi:"id"`
	// Resource name.
	Name string `pulumi:"name"`
	// the hub name.
	TenantId string `pulumi:"tenantId"`
	// Resource type.
	Type string `pulumi:"type"`
	// the user ID.
	UserId *string `pulumi:"userId"`
	// Name of the view.
	ViewName string `pulumi:"viewName"`
}

func LookupViewOutput(ctx *pulumi.Context, args LookupViewOutputArgs, opts ...pulumi.InvokeOption) LookupViewResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupViewResult, error) {
			args := v.(LookupViewArgs)
			r, err := LookupView(ctx, &args, opts...)
			var s LookupViewResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupViewResultOutput)
}

type LookupViewOutputArgs struct {
	// The name of the hub.
	HubName pulumi.StringInput `pulumi:"hubName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The user ID. Use * to retrieve hub level view.
	UserId pulumi.StringInput `pulumi:"userId"`
	// The name of the view.
	ViewName pulumi.StringInput `pulumi:"viewName"`
}

func (LookupViewOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupViewArgs)(nil)).Elem()
}

// The view resource format.
type LookupViewResultOutput struct{ *pulumi.OutputState }

func (LookupViewResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupViewResult)(nil)).Elem()
}

func (o LookupViewResultOutput) ToLookupViewResultOutput() LookupViewResultOutput {
	return o
}

func (o LookupViewResultOutput) ToLookupViewResultOutputWithContext(ctx context.Context) LookupViewResultOutput {
	return o
}

func (o LookupViewResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupViewResult] {
	return pulumix.Output[LookupViewResult]{
		OutputState: o.OutputState,
	}
}

// Date time when view was last modified.
func (o LookupViewResultOutput) Changed() pulumi.StringOutput {
	return o.ApplyT(func(v LookupViewResult) string { return v.Changed }).(pulumi.StringOutput)
}

// Date time when view was created.
func (o LookupViewResultOutput) Created() pulumi.StringOutput {
	return o.ApplyT(func(v LookupViewResult) string { return v.Created }).(pulumi.StringOutput)
}

// View definition.
func (o LookupViewResultOutput) Definition() pulumi.StringOutput {
	return o.ApplyT(func(v LookupViewResult) string { return v.Definition }).(pulumi.StringOutput)
}

// Localized display name for the view.
func (o LookupViewResultOutput) DisplayName() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupViewResult) map[string]string { return v.DisplayName }).(pulumi.StringMapOutput)
}

// Resource ID.
func (o LookupViewResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupViewResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource name.
func (o LookupViewResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupViewResult) string { return v.Name }).(pulumi.StringOutput)
}

// the hub name.
func (o LookupViewResultOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupViewResult) string { return v.TenantId }).(pulumi.StringOutput)
}

// Resource type.
func (o LookupViewResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupViewResult) string { return v.Type }).(pulumi.StringOutput)
}

// the user ID.
func (o LookupViewResultOutput) UserId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupViewResult) *string { return v.UserId }).(pulumi.StringPtrOutput)
}

// Name of the view.
func (o LookupViewResultOutput) ViewName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupViewResult) string { return v.ViewName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupViewResultOutput{})
}
