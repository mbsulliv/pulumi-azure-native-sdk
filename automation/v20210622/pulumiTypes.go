// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210622

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

var _ = utilities.GetEnvOrDefault

// Definition of hybrid runbook worker Legacy.
type HybridRunbookWorkerLegacyResponse struct {
	// Gets or sets the assigned machine IP address.
	Ip *string `pulumi:"ip"`
	// Last Heartbeat from the Worker
	LastSeenDateTime *string `pulumi:"lastSeenDateTime"`
	// Gets or sets the worker machine name.
	Name *string `pulumi:"name"`
	// Gets or sets the registration time of the worker machine.
	RegistrationTime *string `pulumi:"registrationTime"`
}

// Definition of hybrid runbook worker Legacy.
type HybridRunbookWorkerLegacyResponseOutput struct{ *pulumi.OutputState }

func (HybridRunbookWorkerLegacyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HybridRunbookWorkerLegacyResponse)(nil)).Elem()
}

func (o HybridRunbookWorkerLegacyResponseOutput) ToHybridRunbookWorkerLegacyResponseOutput() HybridRunbookWorkerLegacyResponseOutput {
	return o
}

func (o HybridRunbookWorkerLegacyResponseOutput) ToHybridRunbookWorkerLegacyResponseOutputWithContext(ctx context.Context) HybridRunbookWorkerLegacyResponseOutput {
	return o
}

func (o HybridRunbookWorkerLegacyResponseOutput) ToOutput(ctx context.Context) pulumix.Output[HybridRunbookWorkerLegacyResponse] {
	return pulumix.Output[HybridRunbookWorkerLegacyResponse]{
		OutputState: o.OutputState,
	}
}

// Gets or sets the assigned machine IP address.
func (o HybridRunbookWorkerLegacyResponseOutput) Ip() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HybridRunbookWorkerLegacyResponse) *string { return v.Ip }).(pulumi.StringPtrOutput)
}

// Last Heartbeat from the Worker
func (o HybridRunbookWorkerLegacyResponseOutput) LastSeenDateTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HybridRunbookWorkerLegacyResponse) *string { return v.LastSeenDateTime }).(pulumi.StringPtrOutput)
}

// Gets or sets the worker machine name.
func (o HybridRunbookWorkerLegacyResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HybridRunbookWorkerLegacyResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// Gets or sets the registration time of the worker machine.
func (o HybridRunbookWorkerLegacyResponseOutput) RegistrationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HybridRunbookWorkerLegacyResponse) *string { return v.RegistrationTime }).(pulumi.StringPtrOutput)
}

type HybridRunbookWorkerLegacyResponseArrayOutput struct{ *pulumi.OutputState }

func (HybridRunbookWorkerLegacyResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HybridRunbookWorkerLegacyResponse)(nil)).Elem()
}

func (o HybridRunbookWorkerLegacyResponseArrayOutput) ToHybridRunbookWorkerLegacyResponseArrayOutput() HybridRunbookWorkerLegacyResponseArrayOutput {
	return o
}

func (o HybridRunbookWorkerLegacyResponseArrayOutput) ToHybridRunbookWorkerLegacyResponseArrayOutputWithContext(ctx context.Context) HybridRunbookWorkerLegacyResponseArrayOutput {
	return o
}

func (o HybridRunbookWorkerLegacyResponseArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]HybridRunbookWorkerLegacyResponse] {
	return pulumix.Output[[]HybridRunbookWorkerLegacyResponse]{
		OutputState: o.OutputState,
	}
}

func (o HybridRunbookWorkerLegacyResponseArrayOutput) Index(i pulumi.IntInput) HybridRunbookWorkerLegacyResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) HybridRunbookWorkerLegacyResponse {
		return vs[0].([]HybridRunbookWorkerLegacyResponse)[vs[1].(int)]
	}).(HybridRunbookWorkerLegacyResponseOutput)
}

// Definition of RunAs credential to use for hybrid worker.
type RunAsCredentialAssociationProperty struct {
	// Gets or sets the name of the credential.
	Name *string `pulumi:"name"`
}

// RunAsCredentialAssociationPropertyInput is an input type that accepts RunAsCredentialAssociationPropertyArgs and RunAsCredentialAssociationPropertyOutput values.
// You can construct a concrete instance of `RunAsCredentialAssociationPropertyInput` via:
//
//	RunAsCredentialAssociationPropertyArgs{...}
type RunAsCredentialAssociationPropertyInput interface {
	pulumi.Input

	ToRunAsCredentialAssociationPropertyOutput() RunAsCredentialAssociationPropertyOutput
	ToRunAsCredentialAssociationPropertyOutputWithContext(context.Context) RunAsCredentialAssociationPropertyOutput
}

// Definition of RunAs credential to use for hybrid worker.
type RunAsCredentialAssociationPropertyArgs struct {
	// Gets or sets the name of the credential.
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (RunAsCredentialAssociationPropertyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RunAsCredentialAssociationProperty)(nil)).Elem()
}

func (i RunAsCredentialAssociationPropertyArgs) ToRunAsCredentialAssociationPropertyOutput() RunAsCredentialAssociationPropertyOutput {
	return i.ToRunAsCredentialAssociationPropertyOutputWithContext(context.Background())
}

func (i RunAsCredentialAssociationPropertyArgs) ToRunAsCredentialAssociationPropertyOutputWithContext(ctx context.Context) RunAsCredentialAssociationPropertyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RunAsCredentialAssociationPropertyOutput)
}

func (i RunAsCredentialAssociationPropertyArgs) ToOutput(ctx context.Context) pulumix.Output[RunAsCredentialAssociationProperty] {
	return pulumix.Output[RunAsCredentialAssociationProperty]{
		OutputState: i.ToRunAsCredentialAssociationPropertyOutputWithContext(ctx).OutputState,
	}
}

func (i RunAsCredentialAssociationPropertyArgs) ToRunAsCredentialAssociationPropertyPtrOutput() RunAsCredentialAssociationPropertyPtrOutput {
	return i.ToRunAsCredentialAssociationPropertyPtrOutputWithContext(context.Background())
}

func (i RunAsCredentialAssociationPropertyArgs) ToRunAsCredentialAssociationPropertyPtrOutputWithContext(ctx context.Context) RunAsCredentialAssociationPropertyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RunAsCredentialAssociationPropertyOutput).ToRunAsCredentialAssociationPropertyPtrOutputWithContext(ctx)
}

// RunAsCredentialAssociationPropertyPtrInput is an input type that accepts RunAsCredentialAssociationPropertyArgs, RunAsCredentialAssociationPropertyPtr and RunAsCredentialAssociationPropertyPtrOutput values.
// You can construct a concrete instance of `RunAsCredentialAssociationPropertyPtrInput` via:
//
//	        RunAsCredentialAssociationPropertyArgs{...}
//
//	or:
//
//	        nil
type RunAsCredentialAssociationPropertyPtrInput interface {
	pulumi.Input

	ToRunAsCredentialAssociationPropertyPtrOutput() RunAsCredentialAssociationPropertyPtrOutput
	ToRunAsCredentialAssociationPropertyPtrOutputWithContext(context.Context) RunAsCredentialAssociationPropertyPtrOutput
}

type runAsCredentialAssociationPropertyPtrType RunAsCredentialAssociationPropertyArgs

func RunAsCredentialAssociationPropertyPtr(v *RunAsCredentialAssociationPropertyArgs) RunAsCredentialAssociationPropertyPtrInput {
	return (*runAsCredentialAssociationPropertyPtrType)(v)
}

func (*runAsCredentialAssociationPropertyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RunAsCredentialAssociationProperty)(nil)).Elem()
}

func (i *runAsCredentialAssociationPropertyPtrType) ToRunAsCredentialAssociationPropertyPtrOutput() RunAsCredentialAssociationPropertyPtrOutput {
	return i.ToRunAsCredentialAssociationPropertyPtrOutputWithContext(context.Background())
}

func (i *runAsCredentialAssociationPropertyPtrType) ToRunAsCredentialAssociationPropertyPtrOutputWithContext(ctx context.Context) RunAsCredentialAssociationPropertyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RunAsCredentialAssociationPropertyPtrOutput)
}

func (i *runAsCredentialAssociationPropertyPtrType) ToOutput(ctx context.Context) pulumix.Output[*RunAsCredentialAssociationProperty] {
	return pulumix.Output[*RunAsCredentialAssociationProperty]{
		OutputState: i.ToRunAsCredentialAssociationPropertyPtrOutputWithContext(ctx).OutputState,
	}
}

// Definition of RunAs credential to use for hybrid worker.
type RunAsCredentialAssociationPropertyOutput struct{ *pulumi.OutputState }

func (RunAsCredentialAssociationPropertyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RunAsCredentialAssociationProperty)(nil)).Elem()
}

func (o RunAsCredentialAssociationPropertyOutput) ToRunAsCredentialAssociationPropertyOutput() RunAsCredentialAssociationPropertyOutput {
	return o
}

func (o RunAsCredentialAssociationPropertyOutput) ToRunAsCredentialAssociationPropertyOutputWithContext(ctx context.Context) RunAsCredentialAssociationPropertyOutput {
	return o
}

func (o RunAsCredentialAssociationPropertyOutput) ToRunAsCredentialAssociationPropertyPtrOutput() RunAsCredentialAssociationPropertyPtrOutput {
	return o.ToRunAsCredentialAssociationPropertyPtrOutputWithContext(context.Background())
}

func (o RunAsCredentialAssociationPropertyOutput) ToRunAsCredentialAssociationPropertyPtrOutputWithContext(ctx context.Context) RunAsCredentialAssociationPropertyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RunAsCredentialAssociationProperty) *RunAsCredentialAssociationProperty {
		return &v
	}).(RunAsCredentialAssociationPropertyPtrOutput)
}

func (o RunAsCredentialAssociationPropertyOutput) ToOutput(ctx context.Context) pulumix.Output[RunAsCredentialAssociationProperty] {
	return pulumix.Output[RunAsCredentialAssociationProperty]{
		OutputState: o.OutputState,
	}
}

// Gets or sets the name of the credential.
func (o RunAsCredentialAssociationPropertyOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RunAsCredentialAssociationProperty) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type RunAsCredentialAssociationPropertyPtrOutput struct{ *pulumi.OutputState }

func (RunAsCredentialAssociationPropertyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RunAsCredentialAssociationProperty)(nil)).Elem()
}

func (o RunAsCredentialAssociationPropertyPtrOutput) ToRunAsCredentialAssociationPropertyPtrOutput() RunAsCredentialAssociationPropertyPtrOutput {
	return o
}

func (o RunAsCredentialAssociationPropertyPtrOutput) ToRunAsCredentialAssociationPropertyPtrOutputWithContext(ctx context.Context) RunAsCredentialAssociationPropertyPtrOutput {
	return o
}

func (o RunAsCredentialAssociationPropertyPtrOutput) ToOutput(ctx context.Context) pulumix.Output[*RunAsCredentialAssociationProperty] {
	return pulumix.Output[*RunAsCredentialAssociationProperty]{
		OutputState: o.OutputState,
	}
}

func (o RunAsCredentialAssociationPropertyPtrOutput) Elem() RunAsCredentialAssociationPropertyOutput {
	return o.ApplyT(func(v *RunAsCredentialAssociationProperty) RunAsCredentialAssociationProperty {
		if v != nil {
			return *v
		}
		var ret RunAsCredentialAssociationProperty
		return ret
	}).(RunAsCredentialAssociationPropertyOutput)
}

// Gets or sets the name of the credential.
func (o RunAsCredentialAssociationPropertyPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RunAsCredentialAssociationProperty) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

// Definition of RunAs credential to use for hybrid worker.
type RunAsCredentialAssociationPropertyResponse struct {
	// Gets or sets the name of the credential.
	Name *string `pulumi:"name"`
}

// Definition of RunAs credential to use for hybrid worker.
type RunAsCredentialAssociationPropertyResponseOutput struct{ *pulumi.OutputState }

func (RunAsCredentialAssociationPropertyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RunAsCredentialAssociationPropertyResponse)(nil)).Elem()
}

func (o RunAsCredentialAssociationPropertyResponseOutput) ToRunAsCredentialAssociationPropertyResponseOutput() RunAsCredentialAssociationPropertyResponseOutput {
	return o
}

func (o RunAsCredentialAssociationPropertyResponseOutput) ToRunAsCredentialAssociationPropertyResponseOutputWithContext(ctx context.Context) RunAsCredentialAssociationPropertyResponseOutput {
	return o
}

func (o RunAsCredentialAssociationPropertyResponseOutput) ToOutput(ctx context.Context) pulumix.Output[RunAsCredentialAssociationPropertyResponse] {
	return pulumix.Output[RunAsCredentialAssociationPropertyResponse]{
		OutputState: o.OutputState,
	}
}

// Gets or sets the name of the credential.
func (o RunAsCredentialAssociationPropertyResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RunAsCredentialAssociationPropertyResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type RunAsCredentialAssociationPropertyResponsePtrOutput struct{ *pulumi.OutputState }

func (RunAsCredentialAssociationPropertyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RunAsCredentialAssociationPropertyResponse)(nil)).Elem()
}

func (o RunAsCredentialAssociationPropertyResponsePtrOutput) ToRunAsCredentialAssociationPropertyResponsePtrOutput() RunAsCredentialAssociationPropertyResponsePtrOutput {
	return o
}

func (o RunAsCredentialAssociationPropertyResponsePtrOutput) ToRunAsCredentialAssociationPropertyResponsePtrOutputWithContext(ctx context.Context) RunAsCredentialAssociationPropertyResponsePtrOutput {
	return o
}

func (o RunAsCredentialAssociationPropertyResponsePtrOutput) ToOutput(ctx context.Context) pulumix.Output[*RunAsCredentialAssociationPropertyResponse] {
	return pulumix.Output[*RunAsCredentialAssociationPropertyResponse]{
		OutputState: o.OutputState,
	}
}

func (o RunAsCredentialAssociationPropertyResponsePtrOutput) Elem() RunAsCredentialAssociationPropertyResponseOutput {
	return o.ApplyT(func(v *RunAsCredentialAssociationPropertyResponse) RunAsCredentialAssociationPropertyResponse {
		if v != nil {
			return *v
		}
		var ret RunAsCredentialAssociationPropertyResponse
		return ret
	}).(RunAsCredentialAssociationPropertyResponseOutput)
}

// Gets or sets the name of the credential.
func (o RunAsCredentialAssociationPropertyResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RunAsCredentialAssociationPropertyResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
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
	pulumi.RegisterOutputType(HybridRunbookWorkerLegacyResponseOutput{})
	pulumi.RegisterOutputType(HybridRunbookWorkerLegacyResponseArrayOutput{})
	pulumi.RegisterOutputType(RunAsCredentialAssociationPropertyOutput{})
	pulumi.RegisterOutputType(RunAsCredentialAssociationPropertyPtrOutput{})
	pulumi.RegisterOutputType(RunAsCredentialAssociationPropertyResponseOutput{})
	pulumi.RegisterOutputType(RunAsCredentialAssociationPropertyResponsePtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
}
