// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iotoperationsmq

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get a DiagnosticServiceResource
// Azure REST API version: 2023-10-04-preview.
func LookupDiagnosticService(ctx *pulumi.Context, args *LookupDiagnosticServiceArgs, opts ...pulumi.InvokeOption) (*LookupDiagnosticServiceResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupDiagnosticServiceResult
	err := ctx.Invoke("azure-native:iotoperationsmq:getDiagnosticService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupDiagnosticServiceArgs struct {
	// Name of MQ diagnostic resource
	DiagnosticServiceName string `pulumi:"diagnosticServiceName"`
	// Name of MQ resource
	MqName string `pulumi:"mqName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// MQ diagnostic services resource
type LookupDiagnosticServiceResult struct {
	// The frequency at which the data will be exported.
	DataExportFrequencySeconds *int `pulumi:"dataExportFrequencySeconds"`
	// Extended Location
	ExtendedLocation ExtendedLocationPropertyResponse `pulumi:"extendedLocation"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The details of Diagnostic Service Docker Image.
	Image ContainerImageResponse `pulumi:"image"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The format for the logs generated.
	LogFormat *string `pulumi:"logFormat"`
	// The format for the logs generated.
	LogLevel *string `pulumi:"logLevel"`
	// The maximum data stored in MiB.
	MaxDataStorageSize *float64 `pulumi:"maxDataStorageSize"`
	// The port at which metrics is exposed.
	MetricsPort *int `pulumi:"metricsPort"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The destination to collect traces. Diagnostic service will push traces to this endpoint
	OpenTelemetryTracesCollectorAddr *string `pulumi:"openTelemetryTracesCollectorAddr"`
	// The status of the last operation.
	ProvisioningState string `pulumi:"provisioningState"`
	// Metric inactivity timeout.
	StaleDataTimeoutSeconds *int `pulumi:"staleDataTimeoutSeconds"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupDiagnosticServiceResult
func (val *LookupDiagnosticServiceResult) Defaults() *LookupDiagnosticServiceResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.DataExportFrequencySeconds == nil {
		dataExportFrequencySeconds_ := 10
		tmp.DataExportFrequencySeconds = &dataExportFrequencySeconds_
	}
	if tmp.LogFormat == nil {
		logFormat_ := "json"
		tmp.LogFormat = &logFormat_
	}
	if tmp.LogLevel == nil {
		logLevel_ := "info"
		tmp.LogLevel = &logLevel_
	}
	if tmp.MaxDataStorageSize == nil {
		maxDataStorageSize_ := 16.0
		tmp.MaxDataStorageSize = &maxDataStorageSize_
	}
	if tmp.MetricsPort == nil {
		metricsPort_ := 9600
		tmp.MetricsPort = &metricsPort_
	}
	if tmp.StaleDataTimeoutSeconds == nil {
		staleDataTimeoutSeconds_ := 600
		tmp.StaleDataTimeoutSeconds = &staleDataTimeoutSeconds_
	}
	return &tmp
}

func LookupDiagnosticServiceOutput(ctx *pulumi.Context, args LookupDiagnosticServiceOutputArgs, opts ...pulumi.InvokeOption) LookupDiagnosticServiceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDiagnosticServiceResult, error) {
			args := v.(LookupDiagnosticServiceArgs)
			r, err := LookupDiagnosticService(ctx, &args, opts...)
			var s LookupDiagnosticServiceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDiagnosticServiceResultOutput)
}

type LookupDiagnosticServiceOutputArgs struct {
	// Name of MQ diagnostic resource
	DiagnosticServiceName pulumi.StringInput `pulumi:"diagnosticServiceName"`
	// Name of MQ resource
	MqName pulumi.StringInput `pulumi:"mqName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupDiagnosticServiceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDiagnosticServiceArgs)(nil)).Elem()
}

// MQ diagnostic services resource
type LookupDiagnosticServiceResultOutput struct{ *pulumi.OutputState }

func (LookupDiagnosticServiceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDiagnosticServiceResult)(nil)).Elem()
}

func (o LookupDiagnosticServiceResultOutput) ToLookupDiagnosticServiceResultOutput() LookupDiagnosticServiceResultOutput {
	return o
}

func (o LookupDiagnosticServiceResultOutput) ToLookupDiagnosticServiceResultOutputWithContext(ctx context.Context) LookupDiagnosticServiceResultOutput {
	return o
}

// The frequency at which the data will be exported.
func (o LookupDiagnosticServiceResultOutput) DataExportFrequencySeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupDiagnosticServiceResult) *int { return v.DataExportFrequencySeconds }).(pulumi.IntPtrOutput)
}

// Extended Location
func (o LookupDiagnosticServiceResultOutput) ExtendedLocation() ExtendedLocationPropertyResponseOutput {
	return o.ApplyT(func(v LookupDiagnosticServiceResult) ExtendedLocationPropertyResponse { return v.ExtendedLocation }).(ExtendedLocationPropertyResponseOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupDiagnosticServiceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiagnosticServiceResult) string { return v.Id }).(pulumi.StringOutput)
}

// The details of Diagnostic Service Docker Image.
func (o LookupDiagnosticServiceResultOutput) Image() ContainerImageResponseOutput {
	return o.ApplyT(func(v LookupDiagnosticServiceResult) ContainerImageResponse { return v.Image }).(ContainerImageResponseOutput)
}

// The geo-location where the resource lives
func (o LookupDiagnosticServiceResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiagnosticServiceResult) string { return v.Location }).(pulumi.StringOutput)
}

// The format for the logs generated.
func (o LookupDiagnosticServiceResultOutput) LogFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDiagnosticServiceResult) *string { return v.LogFormat }).(pulumi.StringPtrOutput)
}

// The format for the logs generated.
func (o LookupDiagnosticServiceResultOutput) LogLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDiagnosticServiceResult) *string { return v.LogLevel }).(pulumi.StringPtrOutput)
}

// The maximum data stored in MiB.
func (o LookupDiagnosticServiceResultOutput) MaxDataStorageSize() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupDiagnosticServiceResult) *float64 { return v.MaxDataStorageSize }).(pulumi.Float64PtrOutput)
}

// The port at which metrics is exposed.
func (o LookupDiagnosticServiceResultOutput) MetricsPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupDiagnosticServiceResult) *int { return v.MetricsPort }).(pulumi.IntPtrOutput)
}

// The name of the resource
func (o LookupDiagnosticServiceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiagnosticServiceResult) string { return v.Name }).(pulumi.StringOutput)
}

// The destination to collect traces. Diagnostic service will push traces to this endpoint
func (o LookupDiagnosticServiceResultOutput) OpenTelemetryTracesCollectorAddr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDiagnosticServiceResult) *string { return v.OpenTelemetryTracesCollectorAddr }).(pulumi.StringPtrOutput)
}

// The status of the last operation.
func (o LookupDiagnosticServiceResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiagnosticServiceResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Metric inactivity timeout.
func (o LookupDiagnosticServiceResultOutput) StaleDataTimeoutSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupDiagnosticServiceResult) *int { return v.StaleDataTimeoutSeconds }).(pulumi.IntPtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupDiagnosticServiceResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupDiagnosticServiceResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupDiagnosticServiceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDiagnosticServiceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupDiagnosticServiceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiagnosticServiceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDiagnosticServiceResultOutput{})
}
