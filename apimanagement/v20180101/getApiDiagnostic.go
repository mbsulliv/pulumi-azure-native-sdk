// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20180101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the details of the Diagnostic for an API specified by its identifier.
func LookupApiDiagnostic(ctx *pulumi.Context, args *LookupApiDiagnosticArgs, opts ...pulumi.InvokeOption) (*LookupApiDiagnosticResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupApiDiagnosticResult
	err := ctx.Invoke("azure-native:apimanagement/v20180101:getApiDiagnostic", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApiDiagnosticArgs struct {
	// API identifier. Must be unique in the current API Management service instance.
	ApiId string `pulumi:"apiId"`
	// Diagnostic identifier. Must be unique in the current API Management service instance.
	DiagnosticId string `pulumi:"diagnosticId"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
}

// Diagnostic details.
type LookupApiDiagnosticResult struct {
	// Indicates whether a diagnostic should receive data or not.
	Enabled bool `pulumi:"enabled"`
	// Resource ID.
	Id string `pulumi:"id"`
	// Resource name.
	Name string `pulumi:"name"`
	// Resource type for API Management resource.
	Type string `pulumi:"type"`
}

func LookupApiDiagnosticOutput(ctx *pulumi.Context, args LookupApiDiagnosticOutputArgs, opts ...pulumi.InvokeOption) LookupApiDiagnosticResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupApiDiagnosticResult, error) {
			args := v.(LookupApiDiagnosticArgs)
			r, err := LookupApiDiagnostic(ctx, &args, opts...)
			var s LookupApiDiagnosticResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupApiDiagnosticResultOutput)
}

type LookupApiDiagnosticOutputArgs struct {
	// API identifier. Must be unique in the current API Management service instance.
	ApiId pulumi.StringInput `pulumi:"apiId"`
	// Diagnostic identifier. Must be unique in the current API Management service instance.
	DiagnosticId pulumi.StringInput `pulumi:"diagnosticId"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupApiDiagnosticOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApiDiagnosticArgs)(nil)).Elem()
}

// Diagnostic details.
type LookupApiDiagnosticResultOutput struct{ *pulumi.OutputState }

func (LookupApiDiagnosticResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApiDiagnosticResult)(nil)).Elem()
}

func (o LookupApiDiagnosticResultOutput) ToLookupApiDiagnosticResultOutput() LookupApiDiagnosticResultOutput {
	return o
}

func (o LookupApiDiagnosticResultOutput) ToLookupApiDiagnosticResultOutputWithContext(ctx context.Context) LookupApiDiagnosticResultOutput {
	return o
}

// Indicates whether a diagnostic should receive data or not.
func (o LookupApiDiagnosticResultOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupApiDiagnosticResult) bool { return v.Enabled }).(pulumi.BoolOutput)
}

// Resource ID.
func (o LookupApiDiagnosticResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiDiagnosticResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource name.
func (o LookupApiDiagnosticResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiDiagnosticResult) string { return v.Name }).(pulumi.StringOutput)
}

// Resource type for API Management resource.
func (o LookupApiDiagnosticResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiDiagnosticResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApiDiagnosticResultOutput{})
}
