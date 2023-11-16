// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns whether FTP is allowed on the site or not.
func LookupWebAppFtpAllowed(ctx *pulumi.Context, args *LookupWebAppFtpAllowedArgs, opts ...pulumi.InvokeOption) (*LookupWebAppFtpAllowedResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupWebAppFtpAllowedResult
	err := ctx.Invoke("azure-native:web/v20210101:getWebAppFtpAllowed", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWebAppFtpAllowedArgs struct {
	// Name of the app.
	Name string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Publishing Credentials Policies parameters.
type LookupWebAppFtpAllowedResult struct {
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

func LookupWebAppFtpAllowedOutput(ctx *pulumi.Context, args LookupWebAppFtpAllowedOutputArgs, opts ...pulumi.InvokeOption) LookupWebAppFtpAllowedResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWebAppFtpAllowedResult, error) {
			args := v.(LookupWebAppFtpAllowedArgs)
			r, err := LookupWebAppFtpAllowed(ctx, &args, opts...)
			var s LookupWebAppFtpAllowedResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWebAppFtpAllowedResultOutput)
}

type LookupWebAppFtpAllowedOutputArgs struct {
	// Name of the app.
	Name pulumi.StringInput `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupWebAppFtpAllowedOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppFtpAllowedArgs)(nil)).Elem()
}

// Publishing Credentials Policies parameters.
type LookupWebAppFtpAllowedResultOutput struct{ *pulumi.OutputState }

func (LookupWebAppFtpAllowedResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppFtpAllowedResult)(nil)).Elem()
}

func (o LookupWebAppFtpAllowedResultOutput) ToLookupWebAppFtpAllowedResultOutput() LookupWebAppFtpAllowedResultOutput {
	return o
}

func (o LookupWebAppFtpAllowedResultOutput) ToLookupWebAppFtpAllowedResultOutputWithContext(ctx context.Context) LookupWebAppFtpAllowedResultOutput {
	return o
}

// <code>true</code> to allow access to a publishing method; otherwise, <code>false</code>.
func (o LookupWebAppFtpAllowedResultOutput) Allow() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupWebAppFtpAllowedResult) bool { return v.Allow }).(pulumi.BoolOutput)
}

// Resource Id.
func (o LookupWebAppFtpAllowedResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppFtpAllowedResult) string { return v.Id }).(pulumi.StringOutput)
}

// Kind of resource.
func (o LookupWebAppFtpAllowedResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppFtpAllowedResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// Resource Name.
func (o LookupWebAppFtpAllowedResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppFtpAllowedResult) string { return v.Name }).(pulumi.StringOutput)
}

// Resource type.
func (o LookupWebAppFtpAllowedResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppFtpAllowedResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWebAppFtpAllowedResultOutput{})
}
