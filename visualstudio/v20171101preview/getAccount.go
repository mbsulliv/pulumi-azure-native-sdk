// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20171101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the Visual Studio Team Services account resource details.
func LookupAccount(ctx *pulumi.Context, args *LookupAccountArgs, opts ...pulumi.InvokeOption) (*LookupAccountResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupAccountResult
	err := ctx.Invoke("azure-native:visualstudio/v20171101preview:getAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAccountArgs struct {
	// Name of the resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the resource.
	ResourceName string `pulumi:"resourceName"`
}

// The response to an account resource GET request.
type LookupAccountResult struct {
	// Unique identifier of the resource.
	Id string `pulumi:"id"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Resource name.
	Name string `pulumi:"name"`
	// Resource properties.
	Properties map[string]string `pulumi:"properties"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupAccountOutput(ctx *pulumi.Context, args LookupAccountOutputArgs, opts ...pulumi.InvokeOption) LookupAccountResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAccountResult, error) {
			args := v.(LookupAccountArgs)
			r, err := LookupAccount(ctx, &args, opts...)
			var s LookupAccountResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAccountResultOutput)
}

type LookupAccountOutputArgs struct {
	// Name of the resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name of the resource.
	ResourceName pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupAccountOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAccountArgs)(nil)).Elem()
}

// The response to an account resource GET request.
type LookupAccountResultOutput struct{ *pulumi.OutputState }

func (LookupAccountResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAccountResult)(nil)).Elem()
}

func (o LookupAccountResultOutput) ToLookupAccountResultOutput() LookupAccountResultOutput {
	return o
}

func (o LookupAccountResultOutput) ToLookupAccountResultOutputWithContext(ctx context.Context) LookupAccountResultOutput {
	return o
}

// Unique identifier of the resource.
func (o LookupAccountResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource location.
func (o LookupAccountResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAccountResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o LookupAccountResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountResult) string { return v.Name }).(pulumi.StringOutput)
}

// Resource properties.
func (o LookupAccountResultOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupAccountResult) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

// Resource tags.
func (o LookupAccountResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupAccountResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o LookupAccountResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAccountResultOutput{})
}
