// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220504

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = utilities.GetEnvOrDefault

// Automanage configuration profile assignment properties.
type ConfigurationProfileAssignmentProperties struct {
	// The Automanage configurationProfile ARM Resource URI.
	ConfigurationProfile *string `pulumi:"configurationProfile"`
}

// ConfigurationProfileAssignmentPropertiesInput is an input type that accepts ConfigurationProfileAssignmentPropertiesArgs and ConfigurationProfileAssignmentPropertiesOutput values.
// You can construct a concrete instance of `ConfigurationProfileAssignmentPropertiesInput` via:
//
//	ConfigurationProfileAssignmentPropertiesArgs{...}
type ConfigurationProfileAssignmentPropertiesInput interface {
	pulumi.Input

	ToConfigurationProfileAssignmentPropertiesOutput() ConfigurationProfileAssignmentPropertiesOutput
	ToConfigurationProfileAssignmentPropertiesOutputWithContext(context.Context) ConfigurationProfileAssignmentPropertiesOutput
}

// Automanage configuration profile assignment properties.
type ConfigurationProfileAssignmentPropertiesArgs struct {
	// The Automanage configurationProfile ARM Resource URI.
	ConfigurationProfile pulumi.StringPtrInput `pulumi:"configurationProfile"`
}

func (ConfigurationProfileAssignmentPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationProfileAssignmentProperties)(nil)).Elem()
}

func (i ConfigurationProfileAssignmentPropertiesArgs) ToConfigurationProfileAssignmentPropertiesOutput() ConfigurationProfileAssignmentPropertiesOutput {
	return i.ToConfigurationProfileAssignmentPropertiesOutputWithContext(context.Background())
}

func (i ConfigurationProfileAssignmentPropertiesArgs) ToConfigurationProfileAssignmentPropertiesOutputWithContext(ctx context.Context) ConfigurationProfileAssignmentPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationProfileAssignmentPropertiesOutput)
}

func (i ConfigurationProfileAssignmentPropertiesArgs) ToConfigurationProfileAssignmentPropertiesPtrOutput() ConfigurationProfileAssignmentPropertiesPtrOutput {
	return i.ToConfigurationProfileAssignmentPropertiesPtrOutputWithContext(context.Background())
}

func (i ConfigurationProfileAssignmentPropertiesArgs) ToConfigurationProfileAssignmentPropertiesPtrOutputWithContext(ctx context.Context) ConfigurationProfileAssignmentPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationProfileAssignmentPropertiesOutput).ToConfigurationProfileAssignmentPropertiesPtrOutputWithContext(ctx)
}

// ConfigurationProfileAssignmentPropertiesPtrInput is an input type that accepts ConfigurationProfileAssignmentPropertiesArgs, ConfigurationProfileAssignmentPropertiesPtr and ConfigurationProfileAssignmentPropertiesPtrOutput values.
// You can construct a concrete instance of `ConfigurationProfileAssignmentPropertiesPtrInput` via:
//
//	        ConfigurationProfileAssignmentPropertiesArgs{...}
//
//	or:
//
//	        nil
type ConfigurationProfileAssignmentPropertiesPtrInput interface {
	pulumi.Input

	ToConfigurationProfileAssignmentPropertiesPtrOutput() ConfigurationProfileAssignmentPropertiesPtrOutput
	ToConfigurationProfileAssignmentPropertiesPtrOutputWithContext(context.Context) ConfigurationProfileAssignmentPropertiesPtrOutput
}

type configurationProfileAssignmentPropertiesPtrType ConfigurationProfileAssignmentPropertiesArgs

func ConfigurationProfileAssignmentPropertiesPtr(v *ConfigurationProfileAssignmentPropertiesArgs) ConfigurationProfileAssignmentPropertiesPtrInput {
	return (*configurationProfileAssignmentPropertiesPtrType)(v)
}

func (*configurationProfileAssignmentPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationProfileAssignmentProperties)(nil)).Elem()
}

func (i *configurationProfileAssignmentPropertiesPtrType) ToConfigurationProfileAssignmentPropertiesPtrOutput() ConfigurationProfileAssignmentPropertiesPtrOutput {
	return i.ToConfigurationProfileAssignmentPropertiesPtrOutputWithContext(context.Background())
}

func (i *configurationProfileAssignmentPropertiesPtrType) ToConfigurationProfileAssignmentPropertiesPtrOutputWithContext(ctx context.Context) ConfigurationProfileAssignmentPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationProfileAssignmentPropertiesPtrOutput)
}

// Automanage configuration profile assignment properties.
type ConfigurationProfileAssignmentPropertiesOutput struct{ *pulumi.OutputState }

func (ConfigurationProfileAssignmentPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationProfileAssignmentProperties)(nil)).Elem()
}

func (o ConfigurationProfileAssignmentPropertiesOutput) ToConfigurationProfileAssignmentPropertiesOutput() ConfigurationProfileAssignmentPropertiesOutput {
	return o
}

func (o ConfigurationProfileAssignmentPropertiesOutput) ToConfigurationProfileAssignmentPropertiesOutputWithContext(ctx context.Context) ConfigurationProfileAssignmentPropertiesOutput {
	return o
}

func (o ConfigurationProfileAssignmentPropertiesOutput) ToConfigurationProfileAssignmentPropertiesPtrOutput() ConfigurationProfileAssignmentPropertiesPtrOutput {
	return o.ToConfigurationProfileAssignmentPropertiesPtrOutputWithContext(context.Background())
}

func (o ConfigurationProfileAssignmentPropertiesOutput) ToConfigurationProfileAssignmentPropertiesPtrOutputWithContext(ctx context.Context) ConfigurationProfileAssignmentPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConfigurationProfileAssignmentProperties) *ConfigurationProfileAssignmentProperties {
		return &v
	}).(ConfigurationProfileAssignmentPropertiesPtrOutput)
}

// The Automanage configurationProfile ARM Resource URI.
func (o ConfigurationProfileAssignmentPropertiesOutput) ConfigurationProfile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConfigurationProfileAssignmentProperties) *string { return v.ConfigurationProfile }).(pulumi.StringPtrOutput)
}

type ConfigurationProfileAssignmentPropertiesPtrOutput struct{ *pulumi.OutputState }

func (ConfigurationProfileAssignmentPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationProfileAssignmentProperties)(nil)).Elem()
}

func (o ConfigurationProfileAssignmentPropertiesPtrOutput) ToConfigurationProfileAssignmentPropertiesPtrOutput() ConfigurationProfileAssignmentPropertiesPtrOutput {
	return o
}

func (o ConfigurationProfileAssignmentPropertiesPtrOutput) ToConfigurationProfileAssignmentPropertiesPtrOutputWithContext(ctx context.Context) ConfigurationProfileAssignmentPropertiesPtrOutput {
	return o
}

func (o ConfigurationProfileAssignmentPropertiesPtrOutput) Elem() ConfigurationProfileAssignmentPropertiesOutput {
	return o.ApplyT(func(v *ConfigurationProfileAssignmentProperties) ConfigurationProfileAssignmentProperties {
		if v != nil {
			return *v
		}
		var ret ConfigurationProfileAssignmentProperties
		return ret
	}).(ConfigurationProfileAssignmentPropertiesOutput)
}

// The Automanage configurationProfile ARM Resource URI.
func (o ConfigurationProfileAssignmentPropertiesPtrOutput) ConfigurationProfile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigurationProfileAssignmentProperties) *string {
		if v == nil {
			return nil
		}
		return v.ConfigurationProfile
	}).(pulumi.StringPtrOutput)
}

// Automanage configuration profile assignment properties.
type ConfigurationProfileAssignmentPropertiesResponse struct {
	// The Automanage configurationProfile ARM Resource URI.
	ConfigurationProfile *string `pulumi:"configurationProfile"`
	// The status of onboarding, which only appears in the response.
	Status string `pulumi:"status"`
	// The target VM resource URI
	TargetId string `pulumi:"targetId"`
}

// Automanage configuration profile assignment properties.
type ConfigurationProfileAssignmentPropertiesResponseOutput struct{ *pulumi.OutputState }

func (ConfigurationProfileAssignmentPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationProfileAssignmentPropertiesResponse)(nil)).Elem()
}

func (o ConfigurationProfileAssignmentPropertiesResponseOutput) ToConfigurationProfileAssignmentPropertiesResponseOutput() ConfigurationProfileAssignmentPropertiesResponseOutput {
	return o
}

func (o ConfigurationProfileAssignmentPropertiesResponseOutput) ToConfigurationProfileAssignmentPropertiesResponseOutputWithContext(ctx context.Context) ConfigurationProfileAssignmentPropertiesResponseOutput {
	return o
}

// The Automanage configurationProfile ARM Resource URI.
func (o ConfigurationProfileAssignmentPropertiesResponseOutput) ConfigurationProfile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConfigurationProfileAssignmentPropertiesResponse) *string { return v.ConfigurationProfile }).(pulumi.StringPtrOutput)
}

// The status of onboarding, which only appears in the response.
func (o ConfigurationProfileAssignmentPropertiesResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v ConfigurationProfileAssignmentPropertiesResponse) string { return v.Status }).(pulumi.StringOutput)
}

// The target VM resource URI
func (o ConfigurationProfileAssignmentPropertiesResponseOutput) TargetId() pulumi.StringOutput {
	return o.ApplyT(func(v ConfigurationProfileAssignmentPropertiesResponse) string { return v.TargetId }).(pulumi.StringOutput)
}

// Automanage configuration profile properties.
type ConfigurationProfileProperties struct {
	// configuration dictionary of the configuration profile.
	Configuration interface{} `pulumi:"configuration"`
}

// ConfigurationProfilePropertiesInput is an input type that accepts ConfigurationProfilePropertiesArgs and ConfigurationProfilePropertiesOutput values.
// You can construct a concrete instance of `ConfigurationProfilePropertiesInput` via:
//
//	ConfigurationProfilePropertiesArgs{...}
type ConfigurationProfilePropertiesInput interface {
	pulumi.Input

	ToConfigurationProfilePropertiesOutput() ConfigurationProfilePropertiesOutput
	ToConfigurationProfilePropertiesOutputWithContext(context.Context) ConfigurationProfilePropertiesOutput
}

// Automanage configuration profile properties.
type ConfigurationProfilePropertiesArgs struct {
	// configuration dictionary of the configuration profile.
	Configuration pulumi.Input `pulumi:"configuration"`
}

func (ConfigurationProfilePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationProfileProperties)(nil)).Elem()
}

func (i ConfigurationProfilePropertiesArgs) ToConfigurationProfilePropertiesOutput() ConfigurationProfilePropertiesOutput {
	return i.ToConfigurationProfilePropertiesOutputWithContext(context.Background())
}

func (i ConfigurationProfilePropertiesArgs) ToConfigurationProfilePropertiesOutputWithContext(ctx context.Context) ConfigurationProfilePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationProfilePropertiesOutput)
}

func (i ConfigurationProfilePropertiesArgs) ToConfigurationProfilePropertiesPtrOutput() ConfigurationProfilePropertiesPtrOutput {
	return i.ToConfigurationProfilePropertiesPtrOutputWithContext(context.Background())
}

func (i ConfigurationProfilePropertiesArgs) ToConfigurationProfilePropertiesPtrOutputWithContext(ctx context.Context) ConfigurationProfilePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationProfilePropertiesOutput).ToConfigurationProfilePropertiesPtrOutputWithContext(ctx)
}

// ConfigurationProfilePropertiesPtrInput is an input type that accepts ConfigurationProfilePropertiesArgs, ConfigurationProfilePropertiesPtr and ConfigurationProfilePropertiesPtrOutput values.
// You can construct a concrete instance of `ConfigurationProfilePropertiesPtrInput` via:
//
//	        ConfigurationProfilePropertiesArgs{...}
//
//	or:
//
//	        nil
type ConfigurationProfilePropertiesPtrInput interface {
	pulumi.Input

	ToConfigurationProfilePropertiesPtrOutput() ConfigurationProfilePropertiesPtrOutput
	ToConfigurationProfilePropertiesPtrOutputWithContext(context.Context) ConfigurationProfilePropertiesPtrOutput
}

type configurationProfilePropertiesPtrType ConfigurationProfilePropertiesArgs

func ConfigurationProfilePropertiesPtr(v *ConfigurationProfilePropertiesArgs) ConfigurationProfilePropertiesPtrInput {
	return (*configurationProfilePropertiesPtrType)(v)
}

func (*configurationProfilePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationProfileProperties)(nil)).Elem()
}

func (i *configurationProfilePropertiesPtrType) ToConfigurationProfilePropertiesPtrOutput() ConfigurationProfilePropertiesPtrOutput {
	return i.ToConfigurationProfilePropertiesPtrOutputWithContext(context.Background())
}

func (i *configurationProfilePropertiesPtrType) ToConfigurationProfilePropertiesPtrOutputWithContext(ctx context.Context) ConfigurationProfilePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationProfilePropertiesPtrOutput)
}

// Automanage configuration profile properties.
type ConfigurationProfilePropertiesOutput struct{ *pulumi.OutputState }

func (ConfigurationProfilePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationProfileProperties)(nil)).Elem()
}

func (o ConfigurationProfilePropertiesOutput) ToConfigurationProfilePropertiesOutput() ConfigurationProfilePropertiesOutput {
	return o
}

func (o ConfigurationProfilePropertiesOutput) ToConfigurationProfilePropertiesOutputWithContext(ctx context.Context) ConfigurationProfilePropertiesOutput {
	return o
}

func (o ConfigurationProfilePropertiesOutput) ToConfigurationProfilePropertiesPtrOutput() ConfigurationProfilePropertiesPtrOutput {
	return o.ToConfigurationProfilePropertiesPtrOutputWithContext(context.Background())
}

func (o ConfigurationProfilePropertiesOutput) ToConfigurationProfilePropertiesPtrOutputWithContext(ctx context.Context) ConfigurationProfilePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConfigurationProfileProperties) *ConfigurationProfileProperties {
		return &v
	}).(ConfigurationProfilePropertiesPtrOutput)
}

// configuration dictionary of the configuration profile.
func (o ConfigurationProfilePropertiesOutput) Configuration() pulumi.AnyOutput {
	return o.ApplyT(func(v ConfigurationProfileProperties) interface{} { return v.Configuration }).(pulumi.AnyOutput)
}

type ConfigurationProfilePropertiesPtrOutput struct{ *pulumi.OutputState }

func (ConfigurationProfilePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationProfileProperties)(nil)).Elem()
}

func (o ConfigurationProfilePropertiesPtrOutput) ToConfigurationProfilePropertiesPtrOutput() ConfigurationProfilePropertiesPtrOutput {
	return o
}

func (o ConfigurationProfilePropertiesPtrOutput) ToConfigurationProfilePropertiesPtrOutputWithContext(ctx context.Context) ConfigurationProfilePropertiesPtrOutput {
	return o
}

func (o ConfigurationProfilePropertiesPtrOutput) Elem() ConfigurationProfilePropertiesOutput {
	return o.ApplyT(func(v *ConfigurationProfileProperties) ConfigurationProfileProperties {
		if v != nil {
			return *v
		}
		var ret ConfigurationProfileProperties
		return ret
	}).(ConfigurationProfilePropertiesOutput)
}

// configuration dictionary of the configuration profile.
func (o ConfigurationProfilePropertiesPtrOutput) Configuration() pulumi.AnyOutput {
	return o.ApplyT(func(v *ConfigurationProfileProperties) interface{} {
		if v == nil {
			return nil
		}
		return v.Configuration
	}).(pulumi.AnyOutput)
}

// Automanage configuration profile properties.
type ConfigurationProfilePropertiesResponse struct {
	// configuration dictionary of the configuration profile.
	Configuration interface{} `pulumi:"configuration"`
}

// Automanage configuration profile properties.
type ConfigurationProfilePropertiesResponseOutput struct{ *pulumi.OutputState }

func (ConfigurationProfilePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationProfilePropertiesResponse)(nil)).Elem()
}

func (o ConfigurationProfilePropertiesResponseOutput) ToConfigurationProfilePropertiesResponseOutput() ConfigurationProfilePropertiesResponseOutput {
	return o
}

func (o ConfigurationProfilePropertiesResponseOutput) ToConfigurationProfilePropertiesResponseOutputWithContext(ctx context.Context) ConfigurationProfilePropertiesResponseOutput {
	return o
}

// configuration dictionary of the configuration profile.
func (o ConfigurationProfilePropertiesResponseOutput) Configuration() pulumi.AnyOutput {
	return o.ApplyT(func(v ConfigurationProfilePropertiesResponse) interface{} { return v.Configuration }).(pulumi.AnyOutput)
}

// Metadata pertaining to creation and last modification of the resource.
type SystemDataResponse struct {
	// The timestamp of resource creation (UTC).
	CreatedAt *string `pulumi:"createdAt"`
	// The identity that created the resource.
	CreatedBy *string `pulumi:"createdBy"`
	// The type of identity that created the resource.
	CreatedByType *string `pulumi:"createdByType"`
	// The timestamp of resource last modification (UTC)
	LastModifiedAt *string `pulumi:"lastModifiedAt"`
	// The identity that last modified the resource.
	LastModifiedBy *string `pulumi:"lastModifiedBy"`
	// The type of identity that last modified the resource.
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
}

// Metadata pertaining to creation and last modification of the resource.
type SystemDataResponseOutput struct{ *pulumi.OutputState }

func (SystemDataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemDataResponse)(nil)).Elem()
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutput() SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutputWithContext(ctx context.Context) SystemDataResponseOutput {
	return o
}

// The timestamp of resource creation (UTC).
func (o SystemDataResponseOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

// The identity that created the resource.
func (o SystemDataResponseOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

// The type of identity that created the resource.
func (o SystemDataResponseOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

// The timestamp of resource last modification (UTC)
func (o SystemDataResponseOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedAt }).(pulumi.StringPtrOutput)
}

// The identity that last modified the resource.
func (o SystemDataResponseOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

// The type of identity that last modified the resource.
func (o SystemDataResponseOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ConfigurationProfileAssignmentPropertiesOutput{})
	pulumi.RegisterOutputType(ConfigurationProfileAssignmentPropertiesPtrOutput{})
	pulumi.RegisterOutputType(ConfigurationProfileAssignmentPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ConfigurationProfilePropertiesOutput{})
	pulumi.RegisterOutputType(ConfigurationProfilePropertiesPtrOutput{})
	pulumi.RegisterOutputType(ConfigurationProfilePropertiesResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
}
