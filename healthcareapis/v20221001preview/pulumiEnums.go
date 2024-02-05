// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20221001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Type of data destination.
type AnalyticsConnectorDataDestinationType string

const (
	AnalyticsConnectorDataDestinationTypeDatalake = AnalyticsConnectorDataDestinationType("datalake")
)

// Type of data source.
type AnalyticsConnectorDataSourceType string

const (
	AnalyticsConnectorDataSourceTypeFhirservice = AnalyticsConnectorDataSourceType("fhirservice")
)

// Type of data mapping.
type AnalyticsConnectorMappingType string

const (
	AnalyticsConnectorMappingTypeFhirToParquet = AnalyticsConnectorMappingType("fhirToParquet")
)

// The kind of FHIR Service.
type FhirServiceVersion string

const (
	FhirServiceVersionSTU3 = FhirServiceVersion("STU3")
	FhirServiceVersionR4   = FhirServiceVersion("R4")
)

func (FhirServiceVersion) ElementType() reflect.Type {
	return reflect.TypeOf((*FhirServiceVersion)(nil)).Elem()
}

func (e FhirServiceVersion) ToFhirServiceVersionOutput() FhirServiceVersionOutput {
	return pulumi.ToOutput(e).(FhirServiceVersionOutput)
}

func (e FhirServiceVersion) ToFhirServiceVersionOutputWithContext(ctx context.Context) FhirServiceVersionOutput {
	return pulumi.ToOutputWithContext(ctx, e).(FhirServiceVersionOutput)
}

func (e FhirServiceVersion) ToFhirServiceVersionPtrOutput() FhirServiceVersionPtrOutput {
	return e.ToFhirServiceVersionPtrOutputWithContext(context.Background())
}

func (e FhirServiceVersion) ToFhirServiceVersionPtrOutputWithContext(ctx context.Context) FhirServiceVersionPtrOutput {
	return FhirServiceVersion(e).ToFhirServiceVersionOutputWithContext(ctx).ToFhirServiceVersionPtrOutputWithContext(ctx)
}

func (e FhirServiceVersion) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e FhirServiceVersion) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e FhirServiceVersion) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e FhirServiceVersion) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type FhirServiceVersionOutput struct{ *pulumi.OutputState }

func (FhirServiceVersionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FhirServiceVersion)(nil)).Elem()
}

func (o FhirServiceVersionOutput) ToFhirServiceVersionOutput() FhirServiceVersionOutput {
	return o
}

func (o FhirServiceVersionOutput) ToFhirServiceVersionOutputWithContext(ctx context.Context) FhirServiceVersionOutput {
	return o
}

func (o FhirServiceVersionOutput) ToFhirServiceVersionPtrOutput() FhirServiceVersionPtrOutput {
	return o.ToFhirServiceVersionPtrOutputWithContext(context.Background())
}

func (o FhirServiceVersionOutput) ToFhirServiceVersionPtrOutputWithContext(ctx context.Context) FhirServiceVersionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FhirServiceVersion) *FhirServiceVersion {
		return &v
	}).(FhirServiceVersionPtrOutput)
}

func (o FhirServiceVersionOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o FhirServiceVersionOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e FhirServiceVersion) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o FhirServiceVersionOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o FhirServiceVersionOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e FhirServiceVersion) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type FhirServiceVersionPtrOutput struct{ *pulumi.OutputState }

func (FhirServiceVersionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FhirServiceVersion)(nil)).Elem()
}

func (o FhirServiceVersionPtrOutput) ToFhirServiceVersionPtrOutput() FhirServiceVersionPtrOutput {
	return o
}

func (o FhirServiceVersionPtrOutput) ToFhirServiceVersionPtrOutputWithContext(ctx context.Context) FhirServiceVersionPtrOutput {
	return o
}

func (o FhirServiceVersionPtrOutput) Elem() FhirServiceVersionOutput {
	return o.ApplyT(func(v *FhirServiceVersion) FhirServiceVersion {
		if v != nil {
			return *v
		}
		var ret FhirServiceVersion
		return ret
	}).(FhirServiceVersionOutput)
}

func (o FhirServiceVersionPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o FhirServiceVersionPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *FhirServiceVersion) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// FhirServiceVersionInput is an input type that accepts values of the FhirServiceVersion enum
// A concrete instance of `FhirServiceVersionInput` can be one of the following:
//
//	FhirServiceVersionSTU3
//	FhirServiceVersionR4
type FhirServiceVersionInput interface {
	pulumi.Input

	ToFhirServiceVersionOutput() FhirServiceVersionOutput
	ToFhirServiceVersionOutputWithContext(context.Context) FhirServiceVersionOutput
}

var fhirServiceVersionPtrType = reflect.TypeOf((**FhirServiceVersion)(nil)).Elem()

type FhirServiceVersionPtrInput interface {
	pulumi.Input

	ToFhirServiceVersionPtrOutput() FhirServiceVersionPtrOutput
	ToFhirServiceVersionPtrOutputWithContext(context.Context) FhirServiceVersionPtrOutput
}

type fhirServiceVersionPtr string

func FhirServiceVersionPtr(v string) FhirServiceVersionPtrInput {
	return (*fhirServiceVersionPtr)(&v)
}

func (*fhirServiceVersionPtr) ElementType() reflect.Type {
	return fhirServiceVersionPtrType
}

func (in *fhirServiceVersionPtr) ToFhirServiceVersionPtrOutput() FhirServiceVersionPtrOutput {
	return pulumi.ToOutput(in).(FhirServiceVersionPtrOutput)
}

func (in *fhirServiceVersionPtr) ToFhirServiceVersionPtrOutputWithContext(ctx context.Context) FhirServiceVersionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(FhirServiceVersionPtrOutput)
}

// Type of identity being specified, currently SystemAssigned and None are allowed.
type ServiceManagedIdentityType string

const (
	ServiceManagedIdentityTypeNone                         = ServiceManagedIdentityType("None")
	ServiceManagedIdentityTypeSystemAssigned               = ServiceManagedIdentityType("SystemAssigned")
	ServiceManagedIdentityTypeUserAssigned                 = ServiceManagedIdentityType("UserAssigned")
	ServiceManagedIdentityType_SystemAssigned_UserAssigned = ServiceManagedIdentityType("SystemAssigned,UserAssigned")
)

func (ServiceManagedIdentityType) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceManagedIdentityType)(nil)).Elem()
}

func (e ServiceManagedIdentityType) ToServiceManagedIdentityTypeOutput() ServiceManagedIdentityTypeOutput {
	return pulumi.ToOutput(e).(ServiceManagedIdentityTypeOutput)
}

func (e ServiceManagedIdentityType) ToServiceManagedIdentityTypeOutputWithContext(ctx context.Context) ServiceManagedIdentityTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ServiceManagedIdentityTypeOutput)
}

func (e ServiceManagedIdentityType) ToServiceManagedIdentityTypePtrOutput() ServiceManagedIdentityTypePtrOutput {
	return e.ToServiceManagedIdentityTypePtrOutputWithContext(context.Background())
}

func (e ServiceManagedIdentityType) ToServiceManagedIdentityTypePtrOutputWithContext(ctx context.Context) ServiceManagedIdentityTypePtrOutput {
	return ServiceManagedIdentityType(e).ToServiceManagedIdentityTypeOutputWithContext(ctx).ToServiceManagedIdentityTypePtrOutputWithContext(ctx)
}

func (e ServiceManagedIdentityType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ServiceManagedIdentityType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ServiceManagedIdentityType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ServiceManagedIdentityType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ServiceManagedIdentityTypeOutput struct{ *pulumi.OutputState }

func (ServiceManagedIdentityTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceManagedIdentityType)(nil)).Elem()
}

func (o ServiceManagedIdentityTypeOutput) ToServiceManagedIdentityTypeOutput() ServiceManagedIdentityTypeOutput {
	return o
}

func (o ServiceManagedIdentityTypeOutput) ToServiceManagedIdentityTypeOutputWithContext(ctx context.Context) ServiceManagedIdentityTypeOutput {
	return o
}

func (o ServiceManagedIdentityTypeOutput) ToServiceManagedIdentityTypePtrOutput() ServiceManagedIdentityTypePtrOutput {
	return o.ToServiceManagedIdentityTypePtrOutputWithContext(context.Background())
}

func (o ServiceManagedIdentityTypeOutput) ToServiceManagedIdentityTypePtrOutputWithContext(ctx context.Context) ServiceManagedIdentityTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServiceManagedIdentityType) *ServiceManagedIdentityType {
		return &v
	}).(ServiceManagedIdentityTypePtrOutput)
}

func (o ServiceManagedIdentityTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ServiceManagedIdentityTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ServiceManagedIdentityType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ServiceManagedIdentityTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ServiceManagedIdentityTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ServiceManagedIdentityType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ServiceManagedIdentityTypePtrOutput struct{ *pulumi.OutputState }

func (ServiceManagedIdentityTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceManagedIdentityType)(nil)).Elem()
}

func (o ServiceManagedIdentityTypePtrOutput) ToServiceManagedIdentityTypePtrOutput() ServiceManagedIdentityTypePtrOutput {
	return o
}

func (o ServiceManagedIdentityTypePtrOutput) ToServiceManagedIdentityTypePtrOutputWithContext(ctx context.Context) ServiceManagedIdentityTypePtrOutput {
	return o
}

func (o ServiceManagedIdentityTypePtrOutput) Elem() ServiceManagedIdentityTypeOutput {
	return o.ApplyT(func(v *ServiceManagedIdentityType) ServiceManagedIdentityType {
		if v != nil {
			return *v
		}
		var ret ServiceManagedIdentityType
		return ret
	}).(ServiceManagedIdentityTypeOutput)
}

func (o ServiceManagedIdentityTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ServiceManagedIdentityTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ServiceManagedIdentityType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// ServiceManagedIdentityTypeInput is an input type that accepts values of the ServiceManagedIdentityType enum
// A concrete instance of `ServiceManagedIdentityTypeInput` can be one of the following:
//
//	ServiceManagedIdentityTypeNone
//	ServiceManagedIdentityTypeSystemAssigned
//	ServiceManagedIdentityTypeUserAssigned
//	ServiceManagedIdentityType_SystemAssigned_UserAssigned
type ServiceManagedIdentityTypeInput interface {
	pulumi.Input

	ToServiceManagedIdentityTypeOutput() ServiceManagedIdentityTypeOutput
	ToServiceManagedIdentityTypeOutputWithContext(context.Context) ServiceManagedIdentityTypeOutput
}

var serviceManagedIdentityTypePtrType = reflect.TypeOf((**ServiceManagedIdentityType)(nil)).Elem()

type ServiceManagedIdentityTypePtrInput interface {
	pulumi.Input

	ToServiceManagedIdentityTypePtrOutput() ServiceManagedIdentityTypePtrOutput
	ToServiceManagedIdentityTypePtrOutputWithContext(context.Context) ServiceManagedIdentityTypePtrOutput
}

type serviceManagedIdentityTypePtr string

func ServiceManagedIdentityTypePtr(v string) ServiceManagedIdentityTypePtrInput {
	return (*serviceManagedIdentityTypePtr)(&v)
}

func (*serviceManagedIdentityTypePtr) ElementType() reflect.Type {
	return serviceManagedIdentityTypePtrType
}

func (in *serviceManagedIdentityTypePtr) ToServiceManagedIdentityTypePtrOutput() ServiceManagedIdentityTypePtrOutput {
	return pulumi.ToOutput(in).(ServiceManagedIdentityTypePtrOutput)
}

func (in *serviceManagedIdentityTypePtr) ToServiceManagedIdentityTypePtrOutputWithContext(ctx context.Context) ServiceManagedIdentityTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ServiceManagedIdentityTypePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(FhirServiceVersionOutput{})
	pulumi.RegisterOutputType(FhirServiceVersionPtrOutput{})
	pulumi.RegisterOutputType(ServiceManagedIdentityTypeOutput{})
	pulumi.RegisterOutputType(ServiceManagedIdentityTypePtrOutput{})
}
