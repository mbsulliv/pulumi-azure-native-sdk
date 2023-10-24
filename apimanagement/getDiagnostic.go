// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apimanagement

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets the details of the Diagnostic specified by its identifier.
// Azure REST API version: 2022-08-01.
//
// Other available API versions: 2018-01-01, 2019-01-01, 2022-09-01-preview, 2023-03-01-preview.
func LookupDiagnostic(ctx *pulumi.Context, args *LookupDiagnosticArgs, opts ...pulumi.InvokeOption) (*LookupDiagnosticResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupDiagnosticResult
	err := ctx.Invoke("azure-native:apimanagement:getDiagnostic", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDiagnosticArgs struct {
	// Diagnostic identifier. Must be unique in the current API Management service instance.
	DiagnosticId string `pulumi:"diagnosticId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
}

// Diagnostic details.
type LookupDiagnosticResult struct {
	// Specifies for what type of messages sampling settings should not apply.
	AlwaysLog *string `pulumi:"alwaysLog"`
	// Diagnostic settings for incoming/outgoing HTTP messages to the Backend
	Backend *PipelineDiagnosticSettingsResponse `pulumi:"backend"`
	// Diagnostic settings for incoming/outgoing HTTP messages to the Gateway.
	Frontend *PipelineDiagnosticSettingsResponse `pulumi:"frontend"`
	// Sets correlation protocol to use for Application Insights diagnostics.
	HttpCorrelationProtocol *string `pulumi:"httpCorrelationProtocol"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// Log the ClientIP. Default is false.
	LogClientIp *bool `pulumi:"logClientIp"`
	// Resource Id of a target logger.
	LoggerId string `pulumi:"loggerId"`
	// Emit custom metrics via emit-metric policy. Applicable only to Application Insights diagnostic settings.
	Metrics *bool `pulumi:"metrics"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The format of the Operation Name for Application Insights telemetries. Default is Name.
	OperationNameFormat *string `pulumi:"operationNameFormat"`
	// Sampling settings for Diagnostic.
	Sampling *SamplingSettingsResponse `pulumi:"sampling"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// The verbosity level applied to traces emitted by trace policies.
	Verbosity *string `pulumi:"verbosity"`
}

func LookupDiagnosticOutput(ctx *pulumi.Context, args LookupDiagnosticOutputArgs, opts ...pulumi.InvokeOption) LookupDiagnosticResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDiagnosticResult, error) {
			args := v.(LookupDiagnosticArgs)
			r, err := LookupDiagnostic(ctx, &args, opts...)
			var s LookupDiagnosticResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDiagnosticResultOutput)
}

type LookupDiagnosticOutputArgs struct {
	// Diagnostic identifier. Must be unique in the current API Management service instance.
	DiagnosticId pulumi.StringInput `pulumi:"diagnosticId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupDiagnosticOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDiagnosticArgs)(nil)).Elem()
}

// Diagnostic details.
type LookupDiagnosticResultOutput struct{ *pulumi.OutputState }

func (LookupDiagnosticResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDiagnosticResult)(nil)).Elem()
}

func (o LookupDiagnosticResultOutput) ToLookupDiagnosticResultOutput() LookupDiagnosticResultOutput {
	return o
}

func (o LookupDiagnosticResultOutput) ToLookupDiagnosticResultOutputWithContext(ctx context.Context) LookupDiagnosticResultOutput {
	return o
}

func (o LookupDiagnosticResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupDiagnosticResult] {
	return pulumix.Output[LookupDiagnosticResult]{
		OutputState: o.OutputState,
	}
}

// Specifies for what type of messages sampling settings should not apply.
func (o LookupDiagnosticResultOutput) AlwaysLog() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDiagnosticResult) *string { return v.AlwaysLog }).(pulumi.StringPtrOutput)
}

// Diagnostic settings for incoming/outgoing HTTP messages to the Backend
func (o LookupDiagnosticResultOutput) Backend() PipelineDiagnosticSettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupDiagnosticResult) *PipelineDiagnosticSettingsResponse { return v.Backend }).(PipelineDiagnosticSettingsResponsePtrOutput)
}

// Diagnostic settings for incoming/outgoing HTTP messages to the Gateway.
func (o LookupDiagnosticResultOutput) Frontend() PipelineDiagnosticSettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupDiagnosticResult) *PipelineDiagnosticSettingsResponse { return v.Frontend }).(PipelineDiagnosticSettingsResponsePtrOutput)
}

// Sets correlation protocol to use for Application Insights diagnostics.
func (o LookupDiagnosticResultOutput) HttpCorrelationProtocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDiagnosticResult) *string { return v.HttpCorrelationProtocol }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupDiagnosticResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiagnosticResult) string { return v.Id }).(pulumi.StringOutput)
}

// Log the ClientIP. Default is false.
func (o LookupDiagnosticResultOutput) LogClientIp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupDiagnosticResult) *bool { return v.LogClientIp }).(pulumi.BoolPtrOutput)
}

// Resource Id of a target logger.
func (o LookupDiagnosticResultOutput) LoggerId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiagnosticResult) string { return v.LoggerId }).(pulumi.StringOutput)
}

// Emit custom metrics via emit-metric policy. Applicable only to Application Insights diagnostic settings.
func (o LookupDiagnosticResultOutput) Metrics() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupDiagnosticResult) *bool { return v.Metrics }).(pulumi.BoolPtrOutput)
}

// The name of the resource
func (o LookupDiagnosticResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiagnosticResult) string { return v.Name }).(pulumi.StringOutput)
}

// The format of the Operation Name for Application Insights telemetries. Default is Name.
func (o LookupDiagnosticResultOutput) OperationNameFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDiagnosticResult) *string { return v.OperationNameFormat }).(pulumi.StringPtrOutput)
}

// Sampling settings for Diagnostic.
func (o LookupDiagnosticResultOutput) Sampling() SamplingSettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupDiagnosticResult) *SamplingSettingsResponse { return v.Sampling }).(SamplingSettingsResponsePtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupDiagnosticResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiagnosticResult) string { return v.Type }).(pulumi.StringOutput)
}

// The verbosity level applied to traces emitted by trace policies.
func (o LookupDiagnosticResultOutput) Verbosity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDiagnosticResult) *string { return v.Verbosity }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDiagnosticResultOutput{})
}
