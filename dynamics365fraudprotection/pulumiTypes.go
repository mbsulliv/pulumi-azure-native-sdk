// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dynamics365fraudprotection

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

var _ = utilities.GetEnvOrDefault

// An array of administrator user identities
type DFPInstanceAdministrators struct {
	// An array of administrator user identities.
	Members []string `pulumi:"members"`
}

// DFPInstanceAdministratorsInput is an input type that accepts DFPInstanceAdministratorsArgs and DFPInstanceAdministratorsOutput values.
// You can construct a concrete instance of `DFPInstanceAdministratorsInput` via:
//
//	DFPInstanceAdministratorsArgs{...}
type DFPInstanceAdministratorsInput interface {
	pulumi.Input

	ToDFPInstanceAdministratorsOutput() DFPInstanceAdministratorsOutput
	ToDFPInstanceAdministratorsOutputWithContext(context.Context) DFPInstanceAdministratorsOutput
}

// An array of administrator user identities
type DFPInstanceAdministratorsArgs struct {
	// An array of administrator user identities.
	Members pulumi.StringArrayInput `pulumi:"members"`
}

func (DFPInstanceAdministratorsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DFPInstanceAdministrators)(nil)).Elem()
}

func (i DFPInstanceAdministratorsArgs) ToDFPInstanceAdministratorsOutput() DFPInstanceAdministratorsOutput {
	return i.ToDFPInstanceAdministratorsOutputWithContext(context.Background())
}

func (i DFPInstanceAdministratorsArgs) ToDFPInstanceAdministratorsOutputWithContext(ctx context.Context) DFPInstanceAdministratorsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DFPInstanceAdministratorsOutput)
}

func (i DFPInstanceAdministratorsArgs) ToOutput(ctx context.Context) pulumix.Output[DFPInstanceAdministrators] {
	return pulumix.Output[DFPInstanceAdministrators]{
		OutputState: i.ToDFPInstanceAdministratorsOutputWithContext(ctx).OutputState,
	}
}

func (i DFPInstanceAdministratorsArgs) ToDFPInstanceAdministratorsPtrOutput() DFPInstanceAdministratorsPtrOutput {
	return i.ToDFPInstanceAdministratorsPtrOutputWithContext(context.Background())
}

func (i DFPInstanceAdministratorsArgs) ToDFPInstanceAdministratorsPtrOutputWithContext(ctx context.Context) DFPInstanceAdministratorsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DFPInstanceAdministratorsOutput).ToDFPInstanceAdministratorsPtrOutputWithContext(ctx)
}

// DFPInstanceAdministratorsPtrInput is an input type that accepts DFPInstanceAdministratorsArgs, DFPInstanceAdministratorsPtr and DFPInstanceAdministratorsPtrOutput values.
// You can construct a concrete instance of `DFPInstanceAdministratorsPtrInput` via:
//
//	        DFPInstanceAdministratorsArgs{...}
//
//	or:
//
//	        nil
type DFPInstanceAdministratorsPtrInput interface {
	pulumi.Input

	ToDFPInstanceAdministratorsPtrOutput() DFPInstanceAdministratorsPtrOutput
	ToDFPInstanceAdministratorsPtrOutputWithContext(context.Context) DFPInstanceAdministratorsPtrOutput
}

type dfpinstanceAdministratorsPtrType DFPInstanceAdministratorsArgs

func DFPInstanceAdministratorsPtr(v *DFPInstanceAdministratorsArgs) DFPInstanceAdministratorsPtrInput {
	return (*dfpinstanceAdministratorsPtrType)(v)
}

func (*dfpinstanceAdministratorsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DFPInstanceAdministrators)(nil)).Elem()
}

func (i *dfpinstanceAdministratorsPtrType) ToDFPInstanceAdministratorsPtrOutput() DFPInstanceAdministratorsPtrOutput {
	return i.ToDFPInstanceAdministratorsPtrOutputWithContext(context.Background())
}

func (i *dfpinstanceAdministratorsPtrType) ToDFPInstanceAdministratorsPtrOutputWithContext(ctx context.Context) DFPInstanceAdministratorsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DFPInstanceAdministratorsPtrOutput)
}

func (i *dfpinstanceAdministratorsPtrType) ToOutput(ctx context.Context) pulumix.Output[*DFPInstanceAdministrators] {
	return pulumix.Output[*DFPInstanceAdministrators]{
		OutputState: i.ToDFPInstanceAdministratorsPtrOutputWithContext(ctx).OutputState,
	}
}

// An array of administrator user identities
type DFPInstanceAdministratorsOutput struct{ *pulumi.OutputState }

func (DFPInstanceAdministratorsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DFPInstanceAdministrators)(nil)).Elem()
}

func (o DFPInstanceAdministratorsOutput) ToDFPInstanceAdministratorsOutput() DFPInstanceAdministratorsOutput {
	return o
}

func (o DFPInstanceAdministratorsOutput) ToDFPInstanceAdministratorsOutputWithContext(ctx context.Context) DFPInstanceAdministratorsOutput {
	return o
}

func (o DFPInstanceAdministratorsOutput) ToDFPInstanceAdministratorsPtrOutput() DFPInstanceAdministratorsPtrOutput {
	return o.ToDFPInstanceAdministratorsPtrOutputWithContext(context.Background())
}

func (o DFPInstanceAdministratorsOutput) ToDFPInstanceAdministratorsPtrOutputWithContext(ctx context.Context) DFPInstanceAdministratorsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DFPInstanceAdministrators) *DFPInstanceAdministrators {
		return &v
	}).(DFPInstanceAdministratorsPtrOutput)
}

func (o DFPInstanceAdministratorsOutput) ToOutput(ctx context.Context) pulumix.Output[DFPInstanceAdministrators] {
	return pulumix.Output[DFPInstanceAdministrators]{
		OutputState: o.OutputState,
	}
}

// An array of administrator user identities.
func (o DFPInstanceAdministratorsOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DFPInstanceAdministrators) []string { return v.Members }).(pulumi.StringArrayOutput)
}

type DFPInstanceAdministratorsPtrOutput struct{ *pulumi.OutputState }

func (DFPInstanceAdministratorsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DFPInstanceAdministrators)(nil)).Elem()
}

func (o DFPInstanceAdministratorsPtrOutput) ToDFPInstanceAdministratorsPtrOutput() DFPInstanceAdministratorsPtrOutput {
	return o
}

func (o DFPInstanceAdministratorsPtrOutput) ToDFPInstanceAdministratorsPtrOutputWithContext(ctx context.Context) DFPInstanceAdministratorsPtrOutput {
	return o
}

func (o DFPInstanceAdministratorsPtrOutput) ToOutput(ctx context.Context) pulumix.Output[*DFPInstanceAdministrators] {
	return pulumix.Output[*DFPInstanceAdministrators]{
		OutputState: o.OutputState,
	}
}

func (o DFPInstanceAdministratorsPtrOutput) Elem() DFPInstanceAdministratorsOutput {
	return o.ApplyT(func(v *DFPInstanceAdministrators) DFPInstanceAdministrators {
		if v != nil {
			return *v
		}
		var ret DFPInstanceAdministrators
		return ret
	}).(DFPInstanceAdministratorsOutput)
}

// An array of administrator user identities.
func (o DFPInstanceAdministratorsPtrOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DFPInstanceAdministrators) []string {
		if v == nil {
			return nil
		}
		return v.Members
	}).(pulumi.StringArrayOutput)
}

// An array of administrator user identities
type DFPInstanceAdministratorsResponse struct {
	// An array of administrator user identities.
	Members []string `pulumi:"members"`
}

// An array of administrator user identities
type DFPInstanceAdministratorsResponseOutput struct{ *pulumi.OutputState }

func (DFPInstanceAdministratorsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DFPInstanceAdministratorsResponse)(nil)).Elem()
}

func (o DFPInstanceAdministratorsResponseOutput) ToDFPInstanceAdministratorsResponseOutput() DFPInstanceAdministratorsResponseOutput {
	return o
}

func (o DFPInstanceAdministratorsResponseOutput) ToDFPInstanceAdministratorsResponseOutputWithContext(ctx context.Context) DFPInstanceAdministratorsResponseOutput {
	return o
}

func (o DFPInstanceAdministratorsResponseOutput) ToOutput(ctx context.Context) pulumix.Output[DFPInstanceAdministratorsResponse] {
	return pulumix.Output[DFPInstanceAdministratorsResponse]{
		OutputState: o.OutputState,
	}
}

// An array of administrator user identities.
func (o DFPInstanceAdministratorsResponseOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DFPInstanceAdministratorsResponse) []string { return v.Members }).(pulumi.StringArrayOutput)
}

type DFPInstanceAdministratorsResponsePtrOutput struct{ *pulumi.OutputState }

func (DFPInstanceAdministratorsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DFPInstanceAdministratorsResponse)(nil)).Elem()
}

func (o DFPInstanceAdministratorsResponsePtrOutput) ToDFPInstanceAdministratorsResponsePtrOutput() DFPInstanceAdministratorsResponsePtrOutput {
	return o
}

func (o DFPInstanceAdministratorsResponsePtrOutput) ToDFPInstanceAdministratorsResponsePtrOutputWithContext(ctx context.Context) DFPInstanceAdministratorsResponsePtrOutput {
	return o
}

func (o DFPInstanceAdministratorsResponsePtrOutput) ToOutput(ctx context.Context) pulumix.Output[*DFPInstanceAdministratorsResponse] {
	return pulumix.Output[*DFPInstanceAdministratorsResponse]{
		OutputState: o.OutputState,
	}
}

func (o DFPInstanceAdministratorsResponsePtrOutput) Elem() DFPInstanceAdministratorsResponseOutput {
	return o.ApplyT(func(v *DFPInstanceAdministratorsResponse) DFPInstanceAdministratorsResponse {
		if v != nil {
			return *v
		}
		var ret DFPInstanceAdministratorsResponse
		return ret
	}).(DFPInstanceAdministratorsResponseOutput)
}

// An array of administrator user identities.
func (o DFPInstanceAdministratorsResponsePtrOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DFPInstanceAdministratorsResponse) []string {
		if v == nil {
			return nil
		}
		return v.Members
	}).(pulumi.StringArrayOutput)
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

func init() {
	pulumi.RegisterOutputType(DFPInstanceAdministratorsOutput{})
	pulumi.RegisterOutputType(DFPInstanceAdministratorsPtrOutput{})
	pulumi.RegisterOutputType(DFPInstanceAdministratorsResponseOutput{})
	pulumi.RegisterOutputType(DFPInstanceAdministratorsResponsePtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
}
