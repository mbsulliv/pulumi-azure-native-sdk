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

// Description for Gets the logging configuration of an app.
// Azure REST API version: 2022-09-01.
//
// Other available API versions: 2020-10-01.
func LookupWebAppDiagnosticLogsConfiguration(ctx *pulumi.Context, args *LookupWebAppDiagnosticLogsConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupWebAppDiagnosticLogsConfigurationResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupWebAppDiagnosticLogsConfigurationResult
	err := ctx.Invoke("azure-native:web:getWebAppDiagnosticLogsConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupWebAppDiagnosticLogsConfigurationArgs struct {
	// Name of the app.
	Name string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Configuration of App Service site logs.
type LookupWebAppDiagnosticLogsConfigurationResult struct {
	// Application logs configuration.
	ApplicationLogs *ApplicationLogsConfigResponse `pulumi:"applicationLogs"`
	// Detailed error messages configuration.
	DetailedErrorMessages *EnabledConfigResponse `pulumi:"detailedErrorMessages"`
	// Failed requests tracing configuration.
	FailedRequestsTracing *EnabledConfigResponse `pulumi:"failedRequestsTracing"`
	// HTTP logs configuration.
	HttpLogs *HttpLogsConfigResponse `pulumi:"httpLogs"`
	// Resource Id.
	Id string `pulumi:"id"`
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Resource Name.
	Name string `pulumi:"name"`
	// Resource type.
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupWebAppDiagnosticLogsConfigurationResult
func (val *LookupWebAppDiagnosticLogsConfigurationResult) Defaults() *LookupWebAppDiagnosticLogsConfigurationResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.ApplicationLogs = tmp.ApplicationLogs.Defaults()

	return &tmp
}

func LookupWebAppDiagnosticLogsConfigurationOutput(ctx *pulumi.Context, args LookupWebAppDiagnosticLogsConfigurationOutputArgs, opts ...pulumi.InvokeOption) LookupWebAppDiagnosticLogsConfigurationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWebAppDiagnosticLogsConfigurationResult, error) {
			args := v.(LookupWebAppDiagnosticLogsConfigurationArgs)
			r, err := LookupWebAppDiagnosticLogsConfiguration(ctx, &args, opts...)
			var s LookupWebAppDiagnosticLogsConfigurationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWebAppDiagnosticLogsConfigurationResultOutput)
}

type LookupWebAppDiagnosticLogsConfigurationOutputArgs struct {
	// Name of the app.
	Name pulumi.StringInput `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupWebAppDiagnosticLogsConfigurationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppDiagnosticLogsConfigurationArgs)(nil)).Elem()
}

// Configuration of App Service site logs.
type LookupWebAppDiagnosticLogsConfigurationResultOutput struct{ *pulumi.OutputState }

func (LookupWebAppDiagnosticLogsConfigurationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppDiagnosticLogsConfigurationResult)(nil)).Elem()
}

func (o LookupWebAppDiagnosticLogsConfigurationResultOutput) ToLookupWebAppDiagnosticLogsConfigurationResultOutput() LookupWebAppDiagnosticLogsConfigurationResultOutput {
	return o
}

func (o LookupWebAppDiagnosticLogsConfigurationResultOutput) ToLookupWebAppDiagnosticLogsConfigurationResultOutputWithContext(ctx context.Context) LookupWebAppDiagnosticLogsConfigurationResultOutput {
	return o
}

func (o LookupWebAppDiagnosticLogsConfigurationResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupWebAppDiagnosticLogsConfigurationResult] {
	return pulumix.Output[LookupWebAppDiagnosticLogsConfigurationResult]{
		OutputState: o.OutputState,
	}
}

// Application logs configuration.
func (o LookupWebAppDiagnosticLogsConfigurationResultOutput) ApplicationLogs() ApplicationLogsConfigResponsePtrOutput {
	return o.ApplyT(func(v LookupWebAppDiagnosticLogsConfigurationResult) *ApplicationLogsConfigResponse {
		return v.ApplicationLogs
	}).(ApplicationLogsConfigResponsePtrOutput)
}

// Detailed error messages configuration.
func (o LookupWebAppDiagnosticLogsConfigurationResultOutput) DetailedErrorMessages() EnabledConfigResponsePtrOutput {
	return o.ApplyT(func(v LookupWebAppDiagnosticLogsConfigurationResult) *EnabledConfigResponse {
		return v.DetailedErrorMessages
	}).(EnabledConfigResponsePtrOutput)
}

// Failed requests tracing configuration.
func (o LookupWebAppDiagnosticLogsConfigurationResultOutput) FailedRequestsTracing() EnabledConfigResponsePtrOutput {
	return o.ApplyT(func(v LookupWebAppDiagnosticLogsConfigurationResult) *EnabledConfigResponse {
		return v.FailedRequestsTracing
	}).(EnabledConfigResponsePtrOutput)
}

// HTTP logs configuration.
func (o LookupWebAppDiagnosticLogsConfigurationResultOutput) HttpLogs() HttpLogsConfigResponsePtrOutput {
	return o.ApplyT(func(v LookupWebAppDiagnosticLogsConfigurationResult) *HttpLogsConfigResponse { return v.HttpLogs }).(HttpLogsConfigResponsePtrOutput)
}

// Resource Id.
func (o LookupWebAppDiagnosticLogsConfigurationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppDiagnosticLogsConfigurationResult) string { return v.Id }).(pulumi.StringOutput)
}

// Kind of resource.
func (o LookupWebAppDiagnosticLogsConfigurationResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppDiagnosticLogsConfigurationResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// Resource Name.
func (o LookupWebAppDiagnosticLogsConfigurationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppDiagnosticLogsConfigurationResult) string { return v.Name }).(pulumi.StringOutput)
}

// Resource type.
func (o LookupWebAppDiagnosticLogsConfigurationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppDiagnosticLogsConfigurationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWebAppDiagnosticLogsConfigurationResultOutput{})
}
