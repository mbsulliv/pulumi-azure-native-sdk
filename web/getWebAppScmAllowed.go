// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package web

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Description for Returns whether Scm basic auth is allowed on the site or not.
// Azure REST API version: 2022-09-01.
//
// Other available API versions: 2019-08-01, 2020-06-01, 2020-09-01, 2020-10-01, 2020-12-01, 2021-01-01, 2021-01-15, 2021-02-01, 2021-03-01, 2022-03-01, 2023-01-01.
func LookupWebAppScmAllowed(ctx *pulumi.Context, args *LookupWebAppScmAllowedArgs, opts ...pulumi.InvokeOption) (*LookupWebAppScmAllowedResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupWebAppScmAllowedResult
	err := ctx.Invoke("azure-native:web:getWebAppScmAllowed", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWebAppScmAllowedArgs struct {
	// Name of the app.
	Name string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Publishing Credentials Policies parameters.
type LookupWebAppScmAllowedResult struct {
	// <code>true</code> to allow access to a publishing method; otherwise, <code>false</code>.
	Allow bool `pulumi:"allow"`
	// Resource Id.
	Id string `pulumi:"id"`
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Resource Name.
	Name string `pulumi:"name"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupWebAppScmAllowedOutput(ctx *pulumi.Context, args LookupWebAppScmAllowedOutputArgs, opts ...pulumi.InvokeOption) LookupWebAppScmAllowedResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWebAppScmAllowedResult, error) {
			args := v.(LookupWebAppScmAllowedArgs)
			r, err := LookupWebAppScmAllowed(ctx, &args, opts...)
			var s LookupWebAppScmAllowedResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWebAppScmAllowedResultOutput)
}

type LookupWebAppScmAllowedOutputArgs struct {
	// Name of the app.
	Name pulumi.StringInput `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupWebAppScmAllowedOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppScmAllowedArgs)(nil)).Elem()
}

// Publishing Credentials Policies parameters.
type LookupWebAppScmAllowedResultOutput struct{ *pulumi.OutputState }

func (LookupWebAppScmAllowedResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppScmAllowedResult)(nil)).Elem()
}

func (o LookupWebAppScmAllowedResultOutput) ToLookupWebAppScmAllowedResultOutput() LookupWebAppScmAllowedResultOutput {
	return o
}

func (o LookupWebAppScmAllowedResultOutput) ToLookupWebAppScmAllowedResultOutputWithContext(ctx context.Context) LookupWebAppScmAllowedResultOutput {
	return o
}

// <code>true</code> to allow access to a publishing method; otherwise, <code>false</code>.
func (o LookupWebAppScmAllowedResultOutput) Allow() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupWebAppScmAllowedResult) bool { return v.Allow }).(pulumi.BoolOutput)
}

// Resource Id.
func (o LookupWebAppScmAllowedResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppScmAllowedResult) string { return v.Id }).(pulumi.StringOutput)
}

// Kind of resource.
func (o LookupWebAppScmAllowedResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppScmAllowedResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// Resource Name.
func (o LookupWebAppScmAllowedResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppScmAllowedResult) string { return v.Name }).(pulumi.StringOutput)
}

// Resource type.
func (o LookupWebAppScmAllowedResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppScmAllowedResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWebAppScmAllowedResultOutput{})
}
