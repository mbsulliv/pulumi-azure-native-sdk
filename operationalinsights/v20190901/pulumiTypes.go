// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20190901

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = utilities.GetEnvOrDefault

// The related metadata items for the function.
type LogAnalyticsQueryPackQueryPropertiesRelated struct {
	// The related categories for the function.
	Categories []string `pulumi:"categories"`
	// The related resource types for the function.
	ResourceTypes []string `pulumi:"resourceTypes"`
	// The related Log Analytics solutions for the function.
	Solutions []string `pulumi:"solutions"`
}

// LogAnalyticsQueryPackQueryPropertiesRelatedInput is an input type that accepts LogAnalyticsQueryPackQueryPropertiesRelatedArgs and LogAnalyticsQueryPackQueryPropertiesRelatedOutput values.
// You can construct a concrete instance of `LogAnalyticsQueryPackQueryPropertiesRelatedInput` via:
//
//	LogAnalyticsQueryPackQueryPropertiesRelatedArgs{...}
type LogAnalyticsQueryPackQueryPropertiesRelatedInput interface {
	pulumi.Input

	ToLogAnalyticsQueryPackQueryPropertiesRelatedOutput() LogAnalyticsQueryPackQueryPropertiesRelatedOutput
	ToLogAnalyticsQueryPackQueryPropertiesRelatedOutputWithContext(context.Context) LogAnalyticsQueryPackQueryPropertiesRelatedOutput
}

// The related metadata items for the function.
type LogAnalyticsQueryPackQueryPropertiesRelatedArgs struct {
	// The related categories for the function.
	Categories pulumi.StringArrayInput `pulumi:"categories"`
	// The related resource types for the function.
	ResourceTypes pulumi.StringArrayInput `pulumi:"resourceTypes"`
	// The related Log Analytics solutions for the function.
	Solutions pulumi.StringArrayInput `pulumi:"solutions"`
}

func (LogAnalyticsQueryPackQueryPropertiesRelatedArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LogAnalyticsQueryPackQueryPropertiesRelated)(nil)).Elem()
}

func (i LogAnalyticsQueryPackQueryPropertiesRelatedArgs) ToLogAnalyticsQueryPackQueryPropertiesRelatedOutput() LogAnalyticsQueryPackQueryPropertiesRelatedOutput {
	return i.ToLogAnalyticsQueryPackQueryPropertiesRelatedOutputWithContext(context.Background())
}

func (i LogAnalyticsQueryPackQueryPropertiesRelatedArgs) ToLogAnalyticsQueryPackQueryPropertiesRelatedOutputWithContext(ctx context.Context) LogAnalyticsQueryPackQueryPropertiesRelatedOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogAnalyticsQueryPackQueryPropertiesRelatedOutput)
}

func (i LogAnalyticsQueryPackQueryPropertiesRelatedArgs) ToLogAnalyticsQueryPackQueryPropertiesRelatedPtrOutput() LogAnalyticsQueryPackQueryPropertiesRelatedPtrOutput {
	return i.ToLogAnalyticsQueryPackQueryPropertiesRelatedPtrOutputWithContext(context.Background())
}

func (i LogAnalyticsQueryPackQueryPropertiesRelatedArgs) ToLogAnalyticsQueryPackQueryPropertiesRelatedPtrOutputWithContext(ctx context.Context) LogAnalyticsQueryPackQueryPropertiesRelatedPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogAnalyticsQueryPackQueryPropertiesRelatedOutput).ToLogAnalyticsQueryPackQueryPropertiesRelatedPtrOutputWithContext(ctx)
}

// LogAnalyticsQueryPackQueryPropertiesRelatedPtrInput is an input type that accepts LogAnalyticsQueryPackQueryPropertiesRelatedArgs, LogAnalyticsQueryPackQueryPropertiesRelatedPtr and LogAnalyticsQueryPackQueryPropertiesRelatedPtrOutput values.
// You can construct a concrete instance of `LogAnalyticsQueryPackQueryPropertiesRelatedPtrInput` via:
//
//	        LogAnalyticsQueryPackQueryPropertiesRelatedArgs{...}
//
//	or:
//
//	        nil
type LogAnalyticsQueryPackQueryPropertiesRelatedPtrInput interface {
	pulumi.Input

	ToLogAnalyticsQueryPackQueryPropertiesRelatedPtrOutput() LogAnalyticsQueryPackQueryPropertiesRelatedPtrOutput
	ToLogAnalyticsQueryPackQueryPropertiesRelatedPtrOutputWithContext(context.Context) LogAnalyticsQueryPackQueryPropertiesRelatedPtrOutput
}

type logAnalyticsQueryPackQueryPropertiesRelatedPtrType LogAnalyticsQueryPackQueryPropertiesRelatedArgs

func LogAnalyticsQueryPackQueryPropertiesRelatedPtr(v *LogAnalyticsQueryPackQueryPropertiesRelatedArgs) LogAnalyticsQueryPackQueryPropertiesRelatedPtrInput {
	return (*logAnalyticsQueryPackQueryPropertiesRelatedPtrType)(v)
}

func (*logAnalyticsQueryPackQueryPropertiesRelatedPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LogAnalyticsQueryPackQueryPropertiesRelated)(nil)).Elem()
}

func (i *logAnalyticsQueryPackQueryPropertiesRelatedPtrType) ToLogAnalyticsQueryPackQueryPropertiesRelatedPtrOutput() LogAnalyticsQueryPackQueryPropertiesRelatedPtrOutput {
	return i.ToLogAnalyticsQueryPackQueryPropertiesRelatedPtrOutputWithContext(context.Background())
}

func (i *logAnalyticsQueryPackQueryPropertiesRelatedPtrType) ToLogAnalyticsQueryPackQueryPropertiesRelatedPtrOutputWithContext(ctx context.Context) LogAnalyticsQueryPackQueryPropertiesRelatedPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogAnalyticsQueryPackQueryPropertiesRelatedPtrOutput)
}

// The related metadata items for the function.
type LogAnalyticsQueryPackQueryPropertiesRelatedOutput struct{ *pulumi.OutputState }

func (LogAnalyticsQueryPackQueryPropertiesRelatedOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LogAnalyticsQueryPackQueryPropertiesRelated)(nil)).Elem()
}

func (o LogAnalyticsQueryPackQueryPropertiesRelatedOutput) ToLogAnalyticsQueryPackQueryPropertiesRelatedOutput() LogAnalyticsQueryPackQueryPropertiesRelatedOutput {
	return o
}

func (o LogAnalyticsQueryPackQueryPropertiesRelatedOutput) ToLogAnalyticsQueryPackQueryPropertiesRelatedOutputWithContext(ctx context.Context) LogAnalyticsQueryPackQueryPropertiesRelatedOutput {
	return o
}

func (o LogAnalyticsQueryPackQueryPropertiesRelatedOutput) ToLogAnalyticsQueryPackQueryPropertiesRelatedPtrOutput() LogAnalyticsQueryPackQueryPropertiesRelatedPtrOutput {
	return o.ToLogAnalyticsQueryPackQueryPropertiesRelatedPtrOutputWithContext(context.Background())
}

func (o LogAnalyticsQueryPackQueryPropertiesRelatedOutput) ToLogAnalyticsQueryPackQueryPropertiesRelatedPtrOutputWithContext(ctx context.Context) LogAnalyticsQueryPackQueryPropertiesRelatedPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LogAnalyticsQueryPackQueryPropertiesRelated) *LogAnalyticsQueryPackQueryPropertiesRelated {
		return &v
	}).(LogAnalyticsQueryPackQueryPropertiesRelatedPtrOutput)
}

// The related categories for the function.
func (o LogAnalyticsQueryPackQueryPropertiesRelatedOutput) Categories() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LogAnalyticsQueryPackQueryPropertiesRelated) []string { return v.Categories }).(pulumi.StringArrayOutput)
}

// The related resource types for the function.
func (o LogAnalyticsQueryPackQueryPropertiesRelatedOutput) ResourceTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LogAnalyticsQueryPackQueryPropertiesRelated) []string { return v.ResourceTypes }).(pulumi.StringArrayOutput)
}

// The related Log Analytics solutions for the function.
func (o LogAnalyticsQueryPackQueryPropertiesRelatedOutput) Solutions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LogAnalyticsQueryPackQueryPropertiesRelated) []string { return v.Solutions }).(pulumi.StringArrayOutput)
}

type LogAnalyticsQueryPackQueryPropertiesRelatedPtrOutput struct{ *pulumi.OutputState }

func (LogAnalyticsQueryPackQueryPropertiesRelatedPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogAnalyticsQueryPackQueryPropertiesRelated)(nil)).Elem()
}

func (o LogAnalyticsQueryPackQueryPropertiesRelatedPtrOutput) ToLogAnalyticsQueryPackQueryPropertiesRelatedPtrOutput() LogAnalyticsQueryPackQueryPropertiesRelatedPtrOutput {
	return o
}

func (o LogAnalyticsQueryPackQueryPropertiesRelatedPtrOutput) ToLogAnalyticsQueryPackQueryPropertiesRelatedPtrOutputWithContext(ctx context.Context) LogAnalyticsQueryPackQueryPropertiesRelatedPtrOutput {
	return o
}

func (o LogAnalyticsQueryPackQueryPropertiesRelatedPtrOutput) Elem() LogAnalyticsQueryPackQueryPropertiesRelatedOutput {
	return o.ApplyT(func(v *LogAnalyticsQueryPackQueryPropertiesRelated) LogAnalyticsQueryPackQueryPropertiesRelated {
		if v != nil {
			return *v
		}
		var ret LogAnalyticsQueryPackQueryPropertiesRelated
		return ret
	}).(LogAnalyticsQueryPackQueryPropertiesRelatedOutput)
}

// The related categories for the function.
func (o LogAnalyticsQueryPackQueryPropertiesRelatedPtrOutput) Categories() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LogAnalyticsQueryPackQueryPropertiesRelated) []string {
		if v == nil {
			return nil
		}
		return v.Categories
	}).(pulumi.StringArrayOutput)
}

// The related resource types for the function.
func (o LogAnalyticsQueryPackQueryPropertiesRelatedPtrOutput) ResourceTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LogAnalyticsQueryPackQueryPropertiesRelated) []string {
		if v == nil {
			return nil
		}
		return v.ResourceTypes
	}).(pulumi.StringArrayOutput)
}

// The related Log Analytics solutions for the function.
func (o LogAnalyticsQueryPackQueryPropertiesRelatedPtrOutput) Solutions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LogAnalyticsQueryPackQueryPropertiesRelated) []string {
		if v == nil {
			return nil
		}
		return v.Solutions
	}).(pulumi.StringArrayOutput)
}

// The related metadata items for the function.
type LogAnalyticsQueryPackQueryPropertiesResponseRelated struct {
	// The related categories for the function.
	Categories []string `pulumi:"categories"`
	// The related resource types for the function.
	ResourceTypes []string `pulumi:"resourceTypes"`
	// The related Log Analytics solutions for the function.
	Solutions []string `pulumi:"solutions"`
}

// The related metadata items for the function.
type LogAnalyticsQueryPackQueryPropertiesResponseRelatedOutput struct{ *pulumi.OutputState }

func (LogAnalyticsQueryPackQueryPropertiesResponseRelatedOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LogAnalyticsQueryPackQueryPropertiesResponseRelated)(nil)).Elem()
}

func (o LogAnalyticsQueryPackQueryPropertiesResponseRelatedOutput) ToLogAnalyticsQueryPackQueryPropertiesResponseRelatedOutput() LogAnalyticsQueryPackQueryPropertiesResponseRelatedOutput {
	return o
}

func (o LogAnalyticsQueryPackQueryPropertiesResponseRelatedOutput) ToLogAnalyticsQueryPackQueryPropertiesResponseRelatedOutputWithContext(ctx context.Context) LogAnalyticsQueryPackQueryPropertiesResponseRelatedOutput {
	return o
}

// The related categories for the function.
func (o LogAnalyticsQueryPackQueryPropertiesResponseRelatedOutput) Categories() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LogAnalyticsQueryPackQueryPropertiesResponseRelated) []string { return v.Categories }).(pulumi.StringArrayOutput)
}

// The related resource types for the function.
func (o LogAnalyticsQueryPackQueryPropertiesResponseRelatedOutput) ResourceTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LogAnalyticsQueryPackQueryPropertiesResponseRelated) []string { return v.ResourceTypes }).(pulumi.StringArrayOutput)
}

// The related Log Analytics solutions for the function.
func (o LogAnalyticsQueryPackQueryPropertiesResponseRelatedOutput) Solutions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LogAnalyticsQueryPackQueryPropertiesResponseRelated) []string { return v.Solutions }).(pulumi.StringArrayOutput)
}

type LogAnalyticsQueryPackQueryPropertiesResponseRelatedPtrOutput struct{ *pulumi.OutputState }

func (LogAnalyticsQueryPackQueryPropertiesResponseRelatedPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogAnalyticsQueryPackQueryPropertiesResponseRelated)(nil)).Elem()
}

func (o LogAnalyticsQueryPackQueryPropertiesResponseRelatedPtrOutput) ToLogAnalyticsQueryPackQueryPropertiesResponseRelatedPtrOutput() LogAnalyticsQueryPackQueryPropertiesResponseRelatedPtrOutput {
	return o
}

func (o LogAnalyticsQueryPackQueryPropertiesResponseRelatedPtrOutput) ToLogAnalyticsQueryPackQueryPropertiesResponseRelatedPtrOutputWithContext(ctx context.Context) LogAnalyticsQueryPackQueryPropertiesResponseRelatedPtrOutput {
	return o
}

func (o LogAnalyticsQueryPackQueryPropertiesResponseRelatedPtrOutput) Elem() LogAnalyticsQueryPackQueryPropertiesResponseRelatedOutput {
	return o.ApplyT(func(v *LogAnalyticsQueryPackQueryPropertiesResponseRelated) LogAnalyticsQueryPackQueryPropertiesResponseRelated {
		if v != nil {
			return *v
		}
		var ret LogAnalyticsQueryPackQueryPropertiesResponseRelated
		return ret
	}).(LogAnalyticsQueryPackQueryPropertiesResponseRelatedOutput)
}

// The related categories for the function.
func (o LogAnalyticsQueryPackQueryPropertiesResponseRelatedPtrOutput) Categories() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LogAnalyticsQueryPackQueryPropertiesResponseRelated) []string {
		if v == nil {
			return nil
		}
		return v.Categories
	}).(pulumi.StringArrayOutput)
}

// The related resource types for the function.
func (o LogAnalyticsQueryPackQueryPropertiesResponseRelatedPtrOutput) ResourceTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LogAnalyticsQueryPackQueryPropertiesResponseRelated) []string {
		if v == nil {
			return nil
		}
		return v.ResourceTypes
	}).(pulumi.StringArrayOutput)
}

// The related Log Analytics solutions for the function.
func (o LogAnalyticsQueryPackQueryPropertiesResponseRelatedPtrOutput) Solutions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LogAnalyticsQueryPackQueryPropertiesResponseRelated) []string {
		if v == nil {
			return nil
		}
		return v.Solutions
	}).(pulumi.StringArrayOutput)
}

// Read only system data
type SystemDataResponse struct {
	// The timestamp of resource creation (UTC)
	CreatedAt *string `pulumi:"createdAt"`
	// An identifier for the identity that created the resource
	CreatedBy *string `pulumi:"createdBy"`
	// The type of identity that created the resource
	CreatedByType *string `pulumi:"createdByType"`
	// The timestamp of resource last modification (UTC)
	LastModifiedAt *string `pulumi:"lastModifiedAt"`
	// An identifier for the identity that last modified the resource
	LastModifiedBy *string `pulumi:"lastModifiedBy"`
	// The type of identity that last modified the resource
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
}

// Read only system data
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

// The timestamp of resource creation (UTC)
func (o SystemDataResponseOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

// An identifier for the identity that created the resource
func (o SystemDataResponseOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

// The type of identity that created the resource
func (o SystemDataResponseOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

// The timestamp of resource last modification (UTC)
func (o SystemDataResponseOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedAt }).(pulumi.StringPtrOutput)
}

// An identifier for the identity that last modified the resource
func (o SystemDataResponseOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

// The type of identity that last modified the resource
func (o SystemDataResponseOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LogAnalyticsQueryPackQueryPropertiesRelatedOutput{})
	pulumi.RegisterOutputType(LogAnalyticsQueryPackQueryPropertiesRelatedPtrOutput{})
	pulumi.RegisterOutputType(LogAnalyticsQueryPackQueryPropertiesResponseRelatedOutput{})
	pulumi.RegisterOutputType(LogAnalyticsQueryPackQueryPropertiesResponseRelatedPtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
}
