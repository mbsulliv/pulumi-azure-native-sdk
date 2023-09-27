// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230331

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

var _ = utilities.GetEnvOrDefault

// A class that represents a VerificationStatus record.
type DnsRecordResponse struct {
	// Name of the DNS record.
	Name string `pulumi:"name"`
	// Represents an expiry time in seconds to represent how long this entry can be cached by the resolver, default = 3600sec.
	Ttl int `pulumi:"ttl"`
	// Type of the DNS record. Example: TXT
	Type string `pulumi:"type"`
	// Value of the DNS record.
	Value string `pulumi:"value"`
}

// A class that represents a VerificationStatus record.
type DnsRecordResponseOutput struct{ *pulumi.OutputState }

func (DnsRecordResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DnsRecordResponse)(nil)).Elem()
}

func (o DnsRecordResponseOutput) ToDnsRecordResponseOutput() DnsRecordResponseOutput {
	return o
}

func (o DnsRecordResponseOutput) ToDnsRecordResponseOutputWithContext(ctx context.Context) DnsRecordResponseOutput {
	return o
}

func (o DnsRecordResponseOutput) ToOutput(ctx context.Context) pulumix.Output[DnsRecordResponse] {
	return pulumix.Output[DnsRecordResponse]{
		OutputState: o.OutputState,
	}
}

// Name of the DNS record.
func (o DnsRecordResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v DnsRecordResponse) string { return v.Name }).(pulumi.StringOutput)
}

// Represents an expiry time in seconds to represent how long this entry can be cached by the resolver, default = 3600sec.
func (o DnsRecordResponseOutput) Ttl() pulumi.IntOutput {
	return o.ApplyT(func(v DnsRecordResponse) int { return v.Ttl }).(pulumi.IntOutput)
}

// Type of the DNS record. Example: TXT
func (o DnsRecordResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v DnsRecordResponse) string { return v.Type }).(pulumi.StringOutput)
}

// Value of the DNS record.
func (o DnsRecordResponseOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v DnsRecordResponse) string { return v.Value }).(pulumi.StringOutput)
}

type DnsRecordResponsePtrOutput struct{ *pulumi.OutputState }

func (DnsRecordResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DnsRecordResponse)(nil)).Elem()
}

func (o DnsRecordResponsePtrOutput) ToDnsRecordResponsePtrOutput() DnsRecordResponsePtrOutput {
	return o
}

func (o DnsRecordResponsePtrOutput) ToDnsRecordResponsePtrOutputWithContext(ctx context.Context) DnsRecordResponsePtrOutput {
	return o
}

func (o DnsRecordResponsePtrOutput) ToOutput(ctx context.Context) pulumix.Output[*DnsRecordResponse] {
	return pulumix.Output[*DnsRecordResponse]{
		OutputState: o.OutputState,
	}
}

func (o DnsRecordResponsePtrOutput) Elem() DnsRecordResponseOutput {
	return o.ApplyT(func(v *DnsRecordResponse) DnsRecordResponse {
		if v != nil {
			return *v
		}
		var ret DnsRecordResponse
		return ret
	}).(DnsRecordResponseOutput)
}

// Name of the DNS record.
func (o DnsRecordResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DnsRecordResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

// Represents an expiry time in seconds to represent how long this entry can be cached by the resolver, default = 3600sec.
func (o DnsRecordResponsePtrOutput) Ttl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DnsRecordResponse) *int {
		if v == nil {
			return nil
		}
		return &v.Ttl
	}).(pulumi.IntPtrOutput)
}

// Type of the DNS record. Example: TXT
func (o DnsRecordResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DnsRecordResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

// Value of the DNS record.
func (o DnsRecordResponsePtrOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DnsRecordResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Value
	}).(pulumi.StringPtrOutput)
}

// List of DnsRecord
type DomainPropertiesResponseVerificationRecords struct {
	// A class that represents a VerificationStatus record.
	DKIM *DnsRecordResponse `pulumi:"dKIM"`
	// A class that represents a VerificationStatus record.
	DKIM2 *DnsRecordResponse `pulumi:"dKIM2"`
	// A class that represents a VerificationStatus record.
	DMARC *DnsRecordResponse `pulumi:"dMARC"`
	// A class that represents a VerificationStatus record.
	Domain *DnsRecordResponse `pulumi:"domain"`
	// A class that represents a VerificationStatus record.
	SPF *DnsRecordResponse `pulumi:"sPF"`
}

// List of DnsRecord
type DomainPropertiesResponseVerificationRecordsOutput struct{ *pulumi.OutputState }

func (DomainPropertiesResponseVerificationRecordsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DomainPropertiesResponseVerificationRecords)(nil)).Elem()
}

func (o DomainPropertiesResponseVerificationRecordsOutput) ToDomainPropertiesResponseVerificationRecordsOutput() DomainPropertiesResponseVerificationRecordsOutput {
	return o
}

func (o DomainPropertiesResponseVerificationRecordsOutput) ToDomainPropertiesResponseVerificationRecordsOutputWithContext(ctx context.Context) DomainPropertiesResponseVerificationRecordsOutput {
	return o
}

func (o DomainPropertiesResponseVerificationRecordsOutput) ToOutput(ctx context.Context) pulumix.Output[DomainPropertiesResponseVerificationRecords] {
	return pulumix.Output[DomainPropertiesResponseVerificationRecords]{
		OutputState: o.OutputState,
	}
}

// A class that represents a VerificationStatus record.
func (o DomainPropertiesResponseVerificationRecordsOutput) DKIM() DnsRecordResponsePtrOutput {
	return o.ApplyT(func(v DomainPropertiesResponseVerificationRecords) *DnsRecordResponse { return v.DKIM }).(DnsRecordResponsePtrOutput)
}

// A class that represents a VerificationStatus record.
func (o DomainPropertiesResponseVerificationRecordsOutput) DKIM2() DnsRecordResponsePtrOutput {
	return o.ApplyT(func(v DomainPropertiesResponseVerificationRecords) *DnsRecordResponse { return v.DKIM2 }).(DnsRecordResponsePtrOutput)
}

// A class that represents a VerificationStatus record.
func (o DomainPropertiesResponseVerificationRecordsOutput) DMARC() DnsRecordResponsePtrOutput {
	return o.ApplyT(func(v DomainPropertiesResponseVerificationRecords) *DnsRecordResponse { return v.DMARC }).(DnsRecordResponsePtrOutput)
}

// A class that represents a VerificationStatus record.
func (o DomainPropertiesResponseVerificationRecordsOutput) Domain() DnsRecordResponsePtrOutput {
	return o.ApplyT(func(v DomainPropertiesResponseVerificationRecords) *DnsRecordResponse { return v.Domain }).(DnsRecordResponsePtrOutput)
}

// A class that represents a VerificationStatus record.
func (o DomainPropertiesResponseVerificationRecordsOutput) SPF() DnsRecordResponsePtrOutput {
	return o.ApplyT(func(v DomainPropertiesResponseVerificationRecords) *DnsRecordResponse { return v.SPF }).(DnsRecordResponsePtrOutput)
}

// List of VerificationStatusRecord
type DomainPropertiesResponseVerificationStates struct {
	// A class that represents a VerificationStatus record.
	DKIM *VerificationStatusRecordResponse `pulumi:"dKIM"`
	// A class that represents a VerificationStatus record.
	DKIM2 *VerificationStatusRecordResponse `pulumi:"dKIM2"`
	// A class that represents a VerificationStatus record.
	DMARC *VerificationStatusRecordResponse `pulumi:"dMARC"`
	// A class that represents a VerificationStatus record.
	Domain *VerificationStatusRecordResponse `pulumi:"domain"`
	// A class that represents a VerificationStatus record.
	SPF *VerificationStatusRecordResponse `pulumi:"sPF"`
}

// List of VerificationStatusRecord
type DomainPropertiesResponseVerificationStatesOutput struct{ *pulumi.OutputState }

func (DomainPropertiesResponseVerificationStatesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DomainPropertiesResponseVerificationStates)(nil)).Elem()
}

func (o DomainPropertiesResponseVerificationStatesOutput) ToDomainPropertiesResponseVerificationStatesOutput() DomainPropertiesResponseVerificationStatesOutput {
	return o
}

func (o DomainPropertiesResponseVerificationStatesOutput) ToDomainPropertiesResponseVerificationStatesOutputWithContext(ctx context.Context) DomainPropertiesResponseVerificationStatesOutput {
	return o
}

func (o DomainPropertiesResponseVerificationStatesOutput) ToOutput(ctx context.Context) pulumix.Output[DomainPropertiesResponseVerificationStates] {
	return pulumix.Output[DomainPropertiesResponseVerificationStates]{
		OutputState: o.OutputState,
	}
}

// A class that represents a VerificationStatus record.
func (o DomainPropertiesResponseVerificationStatesOutput) DKIM() VerificationStatusRecordResponsePtrOutput {
	return o.ApplyT(func(v DomainPropertiesResponseVerificationStates) *VerificationStatusRecordResponse { return v.DKIM }).(VerificationStatusRecordResponsePtrOutput)
}

// A class that represents a VerificationStatus record.
func (o DomainPropertiesResponseVerificationStatesOutput) DKIM2() VerificationStatusRecordResponsePtrOutput {
	return o.ApplyT(func(v DomainPropertiesResponseVerificationStates) *VerificationStatusRecordResponse { return v.DKIM2 }).(VerificationStatusRecordResponsePtrOutput)
}

// A class that represents a VerificationStatus record.
func (o DomainPropertiesResponseVerificationStatesOutput) DMARC() VerificationStatusRecordResponsePtrOutput {
	return o.ApplyT(func(v DomainPropertiesResponseVerificationStates) *VerificationStatusRecordResponse { return v.DMARC }).(VerificationStatusRecordResponsePtrOutput)
}

// A class that represents a VerificationStatus record.
func (o DomainPropertiesResponseVerificationStatesOutput) Domain() VerificationStatusRecordResponsePtrOutput {
	return o.ApplyT(func(v DomainPropertiesResponseVerificationStates) *VerificationStatusRecordResponse { return v.Domain }).(VerificationStatusRecordResponsePtrOutput)
}

// A class that represents a VerificationStatus record.
func (o DomainPropertiesResponseVerificationStatesOutput) SPF() VerificationStatusRecordResponsePtrOutput {
	return o.ApplyT(func(v DomainPropertiesResponseVerificationStates) *VerificationStatusRecordResponse { return v.SPF }).(VerificationStatusRecordResponsePtrOutput)
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

func (o SystemDataResponseOutput) ToOutput(ctx context.Context) pulumix.Output[SystemDataResponse] {
	return pulumix.Output[SystemDataResponse]{
		OutputState: o.OutputState,
	}
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

// A class that represents a VerificationStatus record.
type VerificationStatusRecordResponse struct {
	// Error code. This property will only be present if the status is UnableToVerify.
	ErrorCode string `pulumi:"errorCode"`
	// Status of the verification operation.
	Status string `pulumi:"status"`
}

// A class that represents a VerificationStatus record.
type VerificationStatusRecordResponseOutput struct{ *pulumi.OutputState }

func (VerificationStatusRecordResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VerificationStatusRecordResponse)(nil)).Elem()
}

func (o VerificationStatusRecordResponseOutput) ToVerificationStatusRecordResponseOutput() VerificationStatusRecordResponseOutput {
	return o
}

func (o VerificationStatusRecordResponseOutput) ToVerificationStatusRecordResponseOutputWithContext(ctx context.Context) VerificationStatusRecordResponseOutput {
	return o
}

func (o VerificationStatusRecordResponseOutput) ToOutput(ctx context.Context) pulumix.Output[VerificationStatusRecordResponse] {
	return pulumix.Output[VerificationStatusRecordResponse]{
		OutputState: o.OutputState,
	}
}

// Error code. This property will only be present if the status is UnableToVerify.
func (o VerificationStatusRecordResponseOutput) ErrorCode() pulumi.StringOutput {
	return o.ApplyT(func(v VerificationStatusRecordResponse) string { return v.ErrorCode }).(pulumi.StringOutput)
}

// Status of the verification operation.
func (o VerificationStatusRecordResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v VerificationStatusRecordResponse) string { return v.Status }).(pulumi.StringOutput)
}

type VerificationStatusRecordResponsePtrOutput struct{ *pulumi.OutputState }

func (VerificationStatusRecordResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VerificationStatusRecordResponse)(nil)).Elem()
}

func (o VerificationStatusRecordResponsePtrOutput) ToVerificationStatusRecordResponsePtrOutput() VerificationStatusRecordResponsePtrOutput {
	return o
}

func (o VerificationStatusRecordResponsePtrOutput) ToVerificationStatusRecordResponsePtrOutputWithContext(ctx context.Context) VerificationStatusRecordResponsePtrOutput {
	return o
}

func (o VerificationStatusRecordResponsePtrOutput) ToOutput(ctx context.Context) pulumix.Output[*VerificationStatusRecordResponse] {
	return pulumix.Output[*VerificationStatusRecordResponse]{
		OutputState: o.OutputState,
	}
}

func (o VerificationStatusRecordResponsePtrOutput) Elem() VerificationStatusRecordResponseOutput {
	return o.ApplyT(func(v *VerificationStatusRecordResponse) VerificationStatusRecordResponse {
		if v != nil {
			return *v
		}
		var ret VerificationStatusRecordResponse
		return ret
	}).(VerificationStatusRecordResponseOutput)
}

// Error code. This property will only be present if the status is UnableToVerify.
func (o VerificationStatusRecordResponsePtrOutput) ErrorCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VerificationStatusRecordResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ErrorCode
	}).(pulumi.StringPtrOutput)
}

// Status of the verification operation.
func (o VerificationStatusRecordResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VerificationStatusRecordResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Status
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(DnsRecordResponseOutput{})
	pulumi.RegisterOutputType(DnsRecordResponsePtrOutput{})
	pulumi.RegisterOutputType(DomainPropertiesResponseVerificationRecordsOutput{})
	pulumi.RegisterOutputType(DomainPropertiesResponseVerificationStatesOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(VerificationStatusRecordResponseOutput{})
	pulumi.RegisterOutputType(VerificationStatusRecordResponsePtrOutput{})
}
