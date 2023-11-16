// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220801

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Diagnostic details.
type ApiDiagnostic struct {
	pulumi.CustomResourceState

	// Specifies for what type of messages sampling settings should not apply.
	AlwaysLog pulumi.StringPtrOutput `pulumi:"alwaysLog"`
	// Diagnostic settings for incoming/outgoing HTTP messages to the Backend
	Backend PipelineDiagnosticSettingsResponsePtrOutput `pulumi:"backend"`
	// Diagnostic settings for incoming/outgoing HTTP messages to the Gateway.
	Frontend PipelineDiagnosticSettingsResponsePtrOutput `pulumi:"frontend"`
	// Sets correlation protocol to use for Application Insights diagnostics.
	HttpCorrelationProtocol pulumi.StringPtrOutput `pulumi:"httpCorrelationProtocol"`
	// Log the ClientIP. Default is false.
	LogClientIp pulumi.BoolPtrOutput `pulumi:"logClientIp"`
	// Resource Id of a target logger.
	LoggerId pulumi.StringOutput `pulumi:"loggerId"`
	// Emit custom metrics via emit-metric policy. Applicable only to Application Insights diagnostic settings.
	Metrics pulumi.BoolPtrOutput `pulumi:"metrics"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The format of the Operation Name for Application Insights telemetries. Default is Name.
	OperationNameFormat pulumi.StringPtrOutput `pulumi:"operationNameFormat"`
	// Sampling settings for Diagnostic.
	Sampling SamplingSettingsResponsePtrOutput `pulumi:"sampling"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// The verbosity level applied to traces emitted by trace policies.
	Verbosity pulumi.StringPtrOutput `pulumi:"verbosity"`
}

// NewApiDiagnostic registers a new resource with the given unique name, arguments, and options.
func NewApiDiagnostic(ctx *pulumi.Context,
	name string, args *ApiDiagnosticArgs, opts ...pulumi.ResourceOption) (*ApiDiagnostic, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiId == nil {
		return nil, errors.New("invalid value for required argument 'ApiId'")
	}
	if args.LoggerId == nil {
		return nil, errors.New("invalid value for required argument 'LoggerId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:apimanagement:ApiDiagnostic"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20170301:ApiDiagnostic"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180101:ApiDiagnostic"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180601preview:ApiDiagnostic"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20190101:ApiDiagnostic"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201:ApiDiagnostic"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201preview:ApiDiagnostic"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20200601preview:ApiDiagnostic"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20201201:ApiDiagnostic"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210101preview:ApiDiagnostic"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210401preview:ApiDiagnostic"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210801:ApiDiagnostic"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20211201preview:ApiDiagnostic"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220401preview:ApiDiagnostic"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220901preview:ApiDiagnostic"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20230301preview:ApiDiagnostic"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ApiDiagnostic
	err := ctx.RegisterResource("azure-native:apimanagement/v20220801:ApiDiagnostic", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApiDiagnostic gets an existing ApiDiagnostic resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApiDiagnostic(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApiDiagnosticState, opts ...pulumi.ResourceOption) (*ApiDiagnostic, error) {
	var resource ApiDiagnostic
	err := ctx.ReadResource("azure-native:apimanagement/v20220801:ApiDiagnostic", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApiDiagnostic resources.
type apiDiagnosticState struct {
}

type ApiDiagnosticState struct {
}

func (ApiDiagnosticState) ElementType() reflect.Type {
	return reflect.TypeOf((*apiDiagnosticState)(nil)).Elem()
}

type apiDiagnosticArgs struct {
	// Specifies for what type of messages sampling settings should not apply.
	AlwaysLog *string `pulumi:"alwaysLog"`
	// API identifier. Must be unique in the current API Management service instance.
	ApiId string `pulumi:"apiId"`
	// Diagnostic settings for incoming/outgoing HTTP messages to the Backend
	Backend *PipelineDiagnosticSettings `pulumi:"backend"`
	// Diagnostic identifier. Must be unique in the current API Management service instance.
	DiagnosticId *string `pulumi:"diagnosticId"`
	// Diagnostic settings for incoming/outgoing HTTP messages to the Gateway.
	Frontend *PipelineDiagnosticSettings `pulumi:"frontend"`
	// Sets correlation protocol to use for Application Insights diagnostics.
	HttpCorrelationProtocol *string `pulumi:"httpCorrelationProtocol"`
	// Log the ClientIP. Default is false.
	LogClientIp *bool `pulumi:"logClientIp"`
	// Resource Id of a target logger.
	LoggerId string `pulumi:"loggerId"`
	// Emit custom metrics via emit-metric policy. Applicable only to Application Insights diagnostic settings.
	Metrics *bool `pulumi:"metrics"`
	// The format of the Operation Name for Application Insights telemetries. Default is Name.
	OperationNameFormat *string `pulumi:"operationNameFormat"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Sampling settings for Diagnostic.
	Sampling *SamplingSettings `pulumi:"sampling"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
	// The verbosity level applied to traces emitted by trace policies.
	Verbosity *string `pulumi:"verbosity"`
}

// The set of arguments for constructing a ApiDiagnostic resource.
type ApiDiagnosticArgs struct {
	// Specifies for what type of messages sampling settings should not apply.
	AlwaysLog pulumi.StringPtrInput
	// API identifier. Must be unique in the current API Management service instance.
	ApiId pulumi.StringInput
	// Diagnostic settings for incoming/outgoing HTTP messages to the Backend
	Backend PipelineDiagnosticSettingsPtrInput
	// Diagnostic identifier. Must be unique in the current API Management service instance.
	DiagnosticId pulumi.StringPtrInput
	// Diagnostic settings for incoming/outgoing HTTP messages to the Gateway.
	Frontend PipelineDiagnosticSettingsPtrInput
	// Sets correlation protocol to use for Application Insights diagnostics.
	HttpCorrelationProtocol pulumi.StringPtrInput
	// Log the ClientIP. Default is false.
	LogClientIp pulumi.BoolPtrInput
	// Resource Id of a target logger.
	LoggerId pulumi.StringInput
	// Emit custom metrics via emit-metric policy. Applicable only to Application Insights diagnostic settings.
	Metrics pulumi.BoolPtrInput
	// The format of the Operation Name for Application Insights telemetries. Default is Name.
	OperationNameFormat pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Sampling settings for Diagnostic.
	Sampling SamplingSettingsPtrInput
	// The name of the API Management service.
	ServiceName pulumi.StringInput
	// The verbosity level applied to traces emitted by trace policies.
	Verbosity pulumi.StringPtrInput
}

func (ApiDiagnosticArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apiDiagnosticArgs)(nil)).Elem()
}

type ApiDiagnosticInput interface {
	pulumi.Input

	ToApiDiagnosticOutput() ApiDiagnosticOutput
	ToApiDiagnosticOutputWithContext(ctx context.Context) ApiDiagnosticOutput
}

func (*ApiDiagnostic) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiDiagnostic)(nil)).Elem()
}

func (i *ApiDiagnostic) ToApiDiagnosticOutput() ApiDiagnosticOutput {
	return i.ToApiDiagnosticOutputWithContext(context.Background())
}

func (i *ApiDiagnostic) ToApiDiagnosticOutputWithContext(ctx context.Context) ApiDiagnosticOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiDiagnosticOutput)
}

type ApiDiagnosticOutput struct{ *pulumi.OutputState }

func (ApiDiagnosticOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiDiagnostic)(nil)).Elem()
}

func (o ApiDiagnosticOutput) ToApiDiagnosticOutput() ApiDiagnosticOutput {
	return o
}

func (o ApiDiagnosticOutput) ToApiDiagnosticOutputWithContext(ctx context.Context) ApiDiagnosticOutput {
	return o
}

// Specifies for what type of messages sampling settings should not apply.
func (o ApiDiagnosticOutput) AlwaysLog() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiDiagnostic) pulumi.StringPtrOutput { return v.AlwaysLog }).(pulumi.StringPtrOutput)
}

// Diagnostic settings for incoming/outgoing HTTP messages to the Backend
func (o ApiDiagnosticOutput) Backend() PipelineDiagnosticSettingsResponsePtrOutput {
	return o.ApplyT(func(v *ApiDiagnostic) PipelineDiagnosticSettingsResponsePtrOutput { return v.Backend }).(PipelineDiagnosticSettingsResponsePtrOutput)
}

// Diagnostic settings for incoming/outgoing HTTP messages to the Gateway.
func (o ApiDiagnosticOutput) Frontend() PipelineDiagnosticSettingsResponsePtrOutput {
	return o.ApplyT(func(v *ApiDiagnostic) PipelineDiagnosticSettingsResponsePtrOutput { return v.Frontend }).(PipelineDiagnosticSettingsResponsePtrOutput)
}

// Sets correlation protocol to use for Application Insights diagnostics.
func (o ApiDiagnosticOutput) HttpCorrelationProtocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiDiagnostic) pulumi.StringPtrOutput { return v.HttpCorrelationProtocol }).(pulumi.StringPtrOutput)
}

// Log the ClientIP. Default is false.
func (o ApiDiagnosticOutput) LogClientIp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApiDiagnostic) pulumi.BoolPtrOutput { return v.LogClientIp }).(pulumi.BoolPtrOutput)
}

// Resource Id of a target logger.
func (o ApiDiagnosticOutput) LoggerId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiDiagnostic) pulumi.StringOutput { return v.LoggerId }).(pulumi.StringOutput)
}

// Emit custom metrics via emit-metric policy. Applicable only to Application Insights diagnostic settings.
func (o ApiDiagnosticOutput) Metrics() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApiDiagnostic) pulumi.BoolPtrOutput { return v.Metrics }).(pulumi.BoolPtrOutput)
}

// The name of the resource
func (o ApiDiagnosticOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiDiagnostic) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The format of the Operation Name for Application Insights telemetries. Default is Name.
func (o ApiDiagnosticOutput) OperationNameFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiDiagnostic) pulumi.StringPtrOutput { return v.OperationNameFormat }).(pulumi.StringPtrOutput)
}

// Sampling settings for Diagnostic.
func (o ApiDiagnosticOutput) Sampling() SamplingSettingsResponsePtrOutput {
	return o.ApplyT(func(v *ApiDiagnostic) SamplingSettingsResponsePtrOutput { return v.Sampling }).(SamplingSettingsResponsePtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ApiDiagnosticOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiDiagnostic) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The verbosity level applied to traces emitted by trace policies.
func (o ApiDiagnosticOutput) Verbosity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiDiagnostic) pulumi.StringPtrOutput { return v.Verbosity }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ApiDiagnosticOutput{})
}
