// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220901preview

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
	err := ctx.Invoke("azure-native:apimanagement/v20220901preview:getApiDiagnostic", args, &rv, opts...)
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
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
}

// Diagnostic details.
type LookupApiDiagnosticResult struct {
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
	// The name of the resource group. The name is case insensitive.
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

// Specifies for what type of messages sampling settings should not apply.
func (o LookupApiDiagnosticResultOutput) AlwaysLog() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiDiagnosticResult) *string { return v.AlwaysLog }).(pulumi.StringPtrOutput)
}

// Diagnostic settings for incoming/outgoing HTTP messages to the Backend
func (o LookupApiDiagnosticResultOutput) Backend() PipelineDiagnosticSettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupApiDiagnosticResult) *PipelineDiagnosticSettingsResponse { return v.Backend }).(PipelineDiagnosticSettingsResponsePtrOutput)
}

// Diagnostic settings for incoming/outgoing HTTP messages to the Gateway.
func (o LookupApiDiagnosticResultOutput) Frontend() PipelineDiagnosticSettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupApiDiagnosticResult) *PipelineDiagnosticSettingsResponse { return v.Frontend }).(PipelineDiagnosticSettingsResponsePtrOutput)
}

// Sets correlation protocol to use for Application Insights diagnostics.
func (o LookupApiDiagnosticResultOutput) HttpCorrelationProtocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiDiagnosticResult) *string { return v.HttpCorrelationProtocol }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupApiDiagnosticResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiDiagnosticResult) string { return v.Id }).(pulumi.StringOutput)
}

// Log the ClientIP. Default is false.
func (o LookupApiDiagnosticResultOutput) LogClientIp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupApiDiagnosticResult) *bool { return v.LogClientIp }).(pulumi.BoolPtrOutput)
}

// Resource Id of a target logger.
func (o LookupApiDiagnosticResultOutput) LoggerId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiDiagnosticResult) string { return v.LoggerId }).(pulumi.StringOutput)
}

// Emit custom metrics via emit-metric policy. Applicable only to Application Insights diagnostic settings.
func (o LookupApiDiagnosticResultOutput) Metrics() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupApiDiagnosticResult) *bool { return v.Metrics }).(pulumi.BoolPtrOutput)
}

// The name of the resource
func (o LookupApiDiagnosticResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiDiagnosticResult) string { return v.Name }).(pulumi.StringOutput)
}

// The format of the Operation Name for Application Insights telemetries. Default is Name.
func (o LookupApiDiagnosticResultOutput) OperationNameFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiDiagnosticResult) *string { return v.OperationNameFormat }).(pulumi.StringPtrOutput)
}

// Sampling settings for Diagnostic.
func (o LookupApiDiagnosticResultOutput) Sampling() SamplingSettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupApiDiagnosticResult) *SamplingSettingsResponse { return v.Sampling }).(SamplingSettingsResponsePtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupApiDiagnosticResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiDiagnosticResult) string { return v.Type }).(pulumi.StringOutput)
}

// The verbosity level applied to traces emitted by trace policies.
func (o LookupApiDiagnosticResultOutput) Verbosity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiDiagnosticResult) *string { return v.Verbosity }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApiDiagnosticResultOutput{})
}
