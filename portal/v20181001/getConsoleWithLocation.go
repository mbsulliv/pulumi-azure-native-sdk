// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20181001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the console for the user.
func LookupConsoleWithLocation(ctx *pulumi.Context, args *LookupConsoleWithLocationArgs, opts ...pulumi.InvokeOption) (*LookupConsoleWithLocationResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupConsoleWithLocationResult
	err := ctx.Invoke("azure-native:portal/v20181001:getConsoleWithLocation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConsoleWithLocationArgs struct {
	// The name of the console
	ConsoleName string `pulumi:"consoleName"`
	// The provider location
	Location string `pulumi:"location"`
}

// Cloud shell console
type LookupConsoleWithLocationResult struct {
	// Cloud shell console properties.
	Properties ConsolePropertiesResponse `pulumi:"properties"`
}

func LookupConsoleWithLocationOutput(ctx *pulumi.Context, args LookupConsoleWithLocationOutputArgs, opts ...pulumi.InvokeOption) LookupConsoleWithLocationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupConsoleWithLocationResult, error) {
			args := v.(LookupConsoleWithLocationArgs)
			r, err := LookupConsoleWithLocation(ctx, &args, opts...)
			var s LookupConsoleWithLocationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupConsoleWithLocationResultOutput)
}

type LookupConsoleWithLocationOutputArgs struct {
	// The name of the console
	ConsoleName pulumi.StringInput `pulumi:"consoleName"`
	// The provider location
	Location pulumi.StringInput `pulumi:"location"`
}

func (LookupConsoleWithLocationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConsoleWithLocationArgs)(nil)).Elem()
}

// Cloud shell console
type LookupConsoleWithLocationResultOutput struct{ *pulumi.OutputState }

func (LookupConsoleWithLocationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConsoleWithLocationResult)(nil)).Elem()
}

func (o LookupConsoleWithLocationResultOutput) ToLookupConsoleWithLocationResultOutput() LookupConsoleWithLocationResultOutput {
	return o
}

func (o LookupConsoleWithLocationResultOutput) ToLookupConsoleWithLocationResultOutputWithContext(ctx context.Context) LookupConsoleWithLocationResultOutput {
	return o
}

// Cloud shell console properties.
func (o LookupConsoleWithLocationResultOutput) Properties() ConsolePropertiesResponseOutput {
	return o.ApplyT(func(v LookupConsoleWithLocationResult) ConsolePropertiesResponse { return v.Properties }).(ConsolePropertiesResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupConsoleWithLocationResultOutput{})
}
