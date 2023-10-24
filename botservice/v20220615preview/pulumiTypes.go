// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220615preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

var _ = utilities.GetEnvOrDefault

// Extra Parameters specific to each Service Provider
type ServiceProviderParameterResponse struct {
	// Default Name for the Service Provider
	Default string `pulumi:"default"`
	// Description of the Service Provider
	Description string `pulumi:"description"`
	// Display Name of the Service Provider
	DisplayName string `pulumi:"displayName"`
	// Help Url for the  Service Provider
	HelpUrl string `pulumi:"helpUrl"`
	// Meta data for the Service Provider
	Metadata ServiceProviderParameterResponseMetadata `pulumi:"metadata"`
	// Name of the Service Provider
	Name string `pulumi:"name"`
	// Type of the Service Provider
	Type string `pulumi:"type"`
}

// Extra Parameters specific to each Service Provider
type ServiceProviderParameterResponseOutput struct{ *pulumi.OutputState }

func (ServiceProviderParameterResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceProviderParameterResponse)(nil)).Elem()
}

func (o ServiceProviderParameterResponseOutput) ToServiceProviderParameterResponseOutput() ServiceProviderParameterResponseOutput {
	return o
}

func (o ServiceProviderParameterResponseOutput) ToServiceProviderParameterResponseOutputWithContext(ctx context.Context) ServiceProviderParameterResponseOutput {
	return o
}

func (o ServiceProviderParameterResponseOutput) ToOutput(ctx context.Context) pulumix.Output[ServiceProviderParameterResponse] {
	return pulumix.Output[ServiceProviderParameterResponse]{
		OutputState: o.OutputState,
	}
}

// Default Name for the Service Provider
func (o ServiceProviderParameterResponseOutput) Default() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceProviderParameterResponse) string { return v.Default }).(pulumi.StringOutput)
}

// Description of the Service Provider
func (o ServiceProviderParameterResponseOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceProviderParameterResponse) string { return v.Description }).(pulumi.StringOutput)
}

// Display Name of the Service Provider
func (o ServiceProviderParameterResponseOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceProviderParameterResponse) string { return v.DisplayName }).(pulumi.StringOutput)
}

// Help Url for the  Service Provider
func (o ServiceProviderParameterResponseOutput) HelpUrl() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceProviderParameterResponse) string { return v.HelpUrl }).(pulumi.StringOutput)
}

// Meta data for the Service Provider
func (o ServiceProviderParameterResponseOutput) Metadata() ServiceProviderParameterResponseMetadataOutput {
	return o.ApplyT(func(v ServiceProviderParameterResponse) ServiceProviderParameterResponseMetadata { return v.Metadata }).(ServiceProviderParameterResponseMetadataOutput)
}

// Name of the Service Provider
func (o ServiceProviderParameterResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceProviderParameterResponse) string { return v.Name }).(pulumi.StringOutput)
}

// Type of the Service Provider
func (o ServiceProviderParameterResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceProviderParameterResponse) string { return v.Type }).(pulumi.StringOutput)
}

type ServiceProviderParameterResponseArrayOutput struct{ *pulumi.OutputState }

func (ServiceProviderParameterResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceProviderParameterResponse)(nil)).Elem()
}

func (o ServiceProviderParameterResponseArrayOutput) ToServiceProviderParameterResponseArrayOutput() ServiceProviderParameterResponseArrayOutput {
	return o
}

func (o ServiceProviderParameterResponseArrayOutput) ToServiceProviderParameterResponseArrayOutputWithContext(ctx context.Context) ServiceProviderParameterResponseArrayOutput {
	return o
}

func (o ServiceProviderParameterResponseArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]ServiceProviderParameterResponse] {
	return pulumix.Output[[]ServiceProviderParameterResponse]{
		OutputState: o.OutputState,
	}
}

func (o ServiceProviderParameterResponseArrayOutput) Index(i pulumi.IntInput) ServiceProviderParameterResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ServiceProviderParameterResponse {
		return vs[0].([]ServiceProviderParameterResponse)[vs[1].(int)]
	}).(ServiceProviderParameterResponseOutput)
}

// the constraints of the bot meta data.
type ServiceProviderParameterResponseConstraints struct {
	// Whether required the constraints of the bot meta data.
	Required *bool `pulumi:"required"`
}

// the constraints of the bot meta data.
type ServiceProviderParameterResponseConstraintsOutput struct{ *pulumi.OutputState }

func (ServiceProviderParameterResponseConstraintsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceProviderParameterResponseConstraints)(nil)).Elem()
}

func (o ServiceProviderParameterResponseConstraintsOutput) ToServiceProviderParameterResponseConstraintsOutput() ServiceProviderParameterResponseConstraintsOutput {
	return o
}

func (o ServiceProviderParameterResponseConstraintsOutput) ToServiceProviderParameterResponseConstraintsOutputWithContext(ctx context.Context) ServiceProviderParameterResponseConstraintsOutput {
	return o
}

func (o ServiceProviderParameterResponseConstraintsOutput) ToOutput(ctx context.Context) pulumix.Output[ServiceProviderParameterResponseConstraints] {
	return pulumix.Output[ServiceProviderParameterResponseConstraints]{
		OutputState: o.OutputState,
	}
}

// Whether required the constraints of the bot meta data.
func (o ServiceProviderParameterResponseConstraintsOutput) Required() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ServiceProviderParameterResponseConstraints) *bool { return v.Required }).(pulumi.BoolPtrOutput)
}

type ServiceProviderParameterResponseConstraintsPtrOutput struct{ *pulumi.OutputState }

func (ServiceProviderParameterResponseConstraintsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceProviderParameterResponseConstraints)(nil)).Elem()
}

func (o ServiceProviderParameterResponseConstraintsPtrOutput) ToServiceProviderParameterResponseConstraintsPtrOutput() ServiceProviderParameterResponseConstraintsPtrOutput {
	return o
}

func (o ServiceProviderParameterResponseConstraintsPtrOutput) ToServiceProviderParameterResponseConstraintsPtrOutputWithContext(ctx context.Context) ServiceProviderParameterResponseConstraintsPtrOutput {
	return o
}

func (o ServiceProviderParameterResponseConstraintsPtrOutput) ToOutput(ctx context.Context) pulumix.Output[*ServiceProviderParameterResponseConstraints] {
	return pulumix.Output[*ServiceProviderParameterResponseConstraints]{
		OutputState: o.OutputState,
	}
}

func (o ServiceProviderParameterResponseConstraintsPtrOutput) Elem() ServiceProviderParameterResponseConstraintsOutput {
	return o.ApplyT(func(v *ServiceProviderParameterResponseConstraints) ServiceProviderParameterResponseConstraints {
		if v != nil {
			return *v
		}
		var ret ServiceProviderParameterResponseConstraints
		return ret
	}).(ServiceProviderParameterResponseConstraintsOutput)
}

// Whether required the constraints of the bot meta data.
func (o ServiceProviderParameterResponseConstraintsPtrOutput) Required() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ServiceProviderParameterResponseConstraints) *bool {
		if v == nil {
			return nil
		}
		return v.Required
	}).(pulumi.BoolPtrOutput)
}

// Meta data for the Service Provider
type ServiceProviderParameterResponseMetadata struct {
	// the constraints of the bot meta data.
	Constraints *ServiceProviderParameterResponseConstraints `pulumi:"constraints"`
}

// Meta data for the Service Provider
type ServiceProviderParameterResponseMetadataOutput struct{ *pulumi.OutputState }

func (ServiceProviderParameterResponseMetadataOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceProviderParameterResponseMetadata)(nil)).Elem()
}

func (o ServiceProviderParameterResponseMetadataOutput) ToServiceProviderParameterResponseMetadataOutput() ServiceProviderParameterResponseMetadataOutput {
	return o
}

func (o ServiceProviderParameterResponseMetadataOutput) ToServiceProviderParameterResponseMetadataOutputWithContext(ctx context.Context) ServiceProviderParameterResponseMetadataOutput {
	return o
}

func (o ServiceProviderParameterResponseMetadataOutput) ToOutput(ctx context.Context) pulumix.Output[ServiceProviderParameterResponseMetadata] {
	return pulumix.Output[ServiceProviderParameterResponseMetadata]{
		OutputState: o.OutputState,
	}
}

// the constraints of the bot meta data.
func (o ServiceProviderParameterResponseMetadataOutput) Constraints() ServiceProviderParameterResponseConstraintsPtrOutput {
	return o.ApplyT(func(v ServiceProviderParameterResponseMetadata) *ServiceProviderParameterResponseConstraints {
		return v.Constraints
	}).(ServiceProviderParameterResponseConstraintsPtrOutput)
}

// The Object used to describe a Service Provider supported by Bot Service
type ServiceProviderPropertiesResponse struct {
	// URL of Dev Portal
	DevPortalUrl string `pulumi:"devPortalUrl"`
	// Display Name of the Service Provider
	DisplayName string `pulumi:"displayName"`
	// The URL of icon
	IconUrl *string `pulumi:"iconUrl"`
	// Id for Service Provider
	Id string `pulumi:"id"`
	// The list of parameters for the Service Provider
	Parameters []ServiceProviderParameterResponse `pulumi:"parameters"`
	// Name of the Service Provider
	ServiceProviderName string `pulumi:"serviceProviderName"`
}

// Defaults sets the appropriate defaults for ServiceProviderPropertiesResponse
func (val *ServiceProviderPropertiesResponse) Defaults() *ServiceProviderPropertiesResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.IconUrl == nil {
		iconUrl_ := ""
		tmp.IconUrl = &iconUrl_
	}
	return &tmp
}

// The Object used to describe a Service Provider supported by Bot Service
type ServiceProviderPropertiesResponseOutput struct{ *pulumi.OutputState }

func (ServiceProviderPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceProviderPropertiesResponse)(nil)).Elem()
}

func (o ServiceProviderPropertiesResponseOutput) ToServiceProviderPropertiesResponseOutput() ServiceProviderPropertiesResponseOutput {
	return o
}

func (o ServiceProviderPropertiesResponseOutput) ToServiceProviderPropertiesResponseOutputWithContext(ctx context.Context) ServiceProviderPropertiesResponseOutput {
	return o
}

func (o ServiceProviderPropertiesResponseOutput) ToOutput(ctx context.Context) pulumix.Output[ServiceProviderPropertiesResponse] {
	return pulumix.Output[ServiceProviderPropertiesResponse]{
		OutputState: o.OutputState,
	}
}

// URL of Dev Portal
func (o ServiceProviderPropertiesResponseOutput) DevPortalUrl() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceProviderPropertiesResponse) string { return v.DevPortalUrl }).(pulumi.StringOutput)
}

// Display Name of the Service Provider
func (o ServiceProviderPropertiesResponseOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceProviderPropertiesResponse) string { return v.DisplayName }).(pulumi.StringOutput)
}

// The URL of icon
func (o ServiceProviderPropertiesResponseOutput) IconUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceProviderPropertiesResponse) *string { return v.IconUrl }).(pulumi.StringPtrOutput)
}

// Id for Service Provider
func (o ServiceProviderPropertiesResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceProviderPropertiesResponse) string { return v.Id }).(pulumi.StringOutput)
}

// The list of parameters for the Service Provider
func (o ServiceProviderPropertiesResponseOutput) Parameters() ServiceProviderParameterResponseArrayOutput {
	return o.ApplyT(func(v ServiceProviderPropertiesResponse) []ServiceProviderParameterResponse { return v.Parameters }).(ServiceProviderParameterResponseArrayOutput)
}

// Name of the Service Provider
func (o ServiceProviderPropertiesResponseOutput) ServiceProviderName() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceProviderPropertiesResponse) string { return v.ServiceProviderName }).(pulumi.StringOutput)
}

type ServiceProviderPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (ServiceProviderPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceProviderPropertiesResponse)(nil)).Elem()
}

func (o ServiceProviderPropertiesResponsePtrOutput) ToServiceProviderPropertiesResponsePtrOutput() ServiceProviderPropertiesResponsePtrOutput {
	return o
}

func (o ServiceProviderPropertiesResponsePtrOutput) ToServiceProviderPropertiesResponsePtrOutputWithContext(ctx context.Context) ServiceProviderPropertiesResponsePtrOutput {
	return o
}

func (o ServiceProviderPropertiesResponsePtrOutput) ToOutput(ctx context.Context) pulumix.Output[*ServiceProviderPropertiesResponse] {
	return pulumix.Output[*ServiceProviderPropertiesResponse]{
		OutputState: o.OutputState,
	}
}

func (o ServiceProviderPropertiesResponsePtrOutput) Elem() ServiceProviderPropertiesResponseOutput {
	return o.ApplyT(func(v *ServiceProviderPropertiesResponse) ServiceProviderPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret ServiceProviderPropertiesResponse
		return ret
	}).(ServiceProviderPropertiesResponseOutput)
}

// URL of Dev Portal
func (o ServiceProviderPropertiesResponsePtrOutput) DevPortalUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceProviderPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.DevPortalUrl
	}).(pulumi.StringPtrOutput)
}

// Display Name of the Service Provider
func (o ServiceProviderPropertiesResponsePtrOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceProviderPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.DisplayName
	}).(pulumi.StringPtrOutput)
}

// The URL of icon
func (o ServiceProviderPropertiesResponsePtrOutput) IconUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceProviderPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.IconUrl
	}).(pulumi.StringPtrOutput)
}

// Id for Service Provider
func (o ServiceProviderPropertiesResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceProviderPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

// The list of parameters for the Service Provider
func (o ServiceProviderPropertiesResponsePtrOutput) Parameters() ServiceProviderParameterResponseArrayOutput {
	return o.ApplyT(func(v *ServiceProviderPropertiesResponse) []ServiceProviderParameterResponse {
		if v == nil {
			return nil
		}
		return v.Parameters
	}).(ServiceProviderParameterResponseArrayOutput)
}

// Name of the Service Provider
func (o ServiceProviderPropertiesResponsePtrOutput) ServiceProviderName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceProviderPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ServiceProviderName
	}).(pulumi.StringPtrOutput)
}

// Service Provider Definition
type ServiceProviderResponse struct {
	// The Properties of a Service Provider Object
	Properties *ServiceProviderPropertiesResponse `pulumi:"properties"`
}

// Defaults sets the appropriate defaults for ServiceProviderResponse
func (val *ServiceProviderResponse) Defaults() *ServiceProviderResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Properties = tmp.Properties.Defaults()

	return &tmp
}

// Service Provider Definition
type ServiceProviderResponseOutput struct{ *pulumi.OutputState }

func (ServiceProviderResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceProviderResponse)(nil)).Elem()
}

func (o ServiceProviderResponseOutput) ToServiceProviderResponseOutput() ServiceProviderResponseOutput {
	return o
}

func (o ServiceProviderResponseOutput) ToServiceProviderResponseOutputWithContext(ctx context.Context) ServiceProviderResponseOutput {
	return o
}

func (o ServiceProviderResponseOutput) ToOutput(ctx context.Context) pulumix.Output[ServiceProviderResponse] {
	return pulumix.Output[ServiceProviderResponse]{
		OutputState: o.OutputState,
	}
}

// The Properties of a Service Provider Object
func (o ServiceProviderResponseOutput) Properties() ServiceProviderPropertiesResponsePtrOutput {
	return o.ApplyT(func(v ServiceProviderResponse) *ServiceProviderPropertiesResponse { return v.Properties }).(ServiceProviderPropertiesResponsePtrOutput)
}

type ServiceProviderResponseArrayOutput struct{ *pulumi.OutputState }

func (ServiceProviderResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceProviderResponse)(nil)).Elem()
}

func (o ServiceProviderResponseArrayOutput) ToServiceProviderResponseArrayOutput() ServiceProviderResponseArrayOutput {
	return o
}

func (o ServiceProviderResponseArrayOutput) ToServiceProviderResponseArrayOutputWithContext(ctx context.Context) ServiceProviderResponseArrayOutput {
	return o
}

func (o ServiceProviderResponseArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]ServiceProviderResponse] {
	return pulumix.Output[[]ServiceProviderResponse]{
		OutputState: o.OutputState,
	}
}

func (o ServiceProviderResponseArrayOutput) Index(i pulumi.IntInput) ServiceProviderResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ServiceProviderResponse {
		return vs[0].([]ServiceProviderResponse)[vs[1].(int)]
	}).(ServiceProviderResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(ServiceProviderParameterResponseOutput{})
	pulumi.RegisterOutputType(ServiceProviderParameterResponseArrayOutput{})
	pulumi.RegisterOutputType(ServiceProviderParameterResponseConstraintsOutput{})
	pulumi.RegisterOutputType(ServiceProviderParameterResponseConstraintsPtrOutput{})
	pulumi.RegisterOutputType(ServiceProviderParameterResponseMetadataOutput{})
	pulumi.RegisterOutputType(ServiceProviderPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ServiceProviderPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(ServiceProviderResponseOutput{})
	pulumi.RegisterOutputType(ServiceProviderResponseArrayOutput{})
}
