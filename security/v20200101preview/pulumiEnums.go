// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20200101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Connect to your cloud account, for AWS use either account credentials or role-based authentication. For GCP use account organization credentials.
type AuthenticationType string

const (
	// AWS cloud account connector user credentials authentication
	AuthenticationTypeAwsCreds = AuthenticationType("awsCreds")
	// AWS account connector assume role authentication
	AuthenticationTypeAwsAssumeRole = AuthenticationType("awsAssumeRole")
	// GCP account connector service to service authentication
	AuthenticationTypeGcpCredentials = AuthenticationType("gcpCredentials")
)

// Whether or not to automatically install Azure Arc (hybrid compute) agents on machines
type AutoProvision string

const (
	// Install missing Azure Arc agents on machines automatically
	AutoProvisionOn = AutoProvision("On")
	// Do not install Azure Arc agent on the machines automatically
	AutoProvisionOff = AutoProvision("Off")
)

func (AutoProvision) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoProvision)(nil)).Elem()
}

func (e AutoProvision) ToAutoProvisionOutput() AutoProvisionOutput {
	return pulumi.ToOutput(e).(AutoProvisionOutput)
}

func (e AutoProvision) ToAutoProvisionOutputWithContext(ctx context.Context) AutoProvisionOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AutoProvisionOutput)
}

func (e AutoProvision) ToAutoProvisionPtrOutput() AutoProvisionPtrOutput {
	return e.ToAutoProvisionPtrOutputWithContext(context.Background())
}

func (e AutoProvision) ToAutoProvisionPtrOutputWithContext(ctx context.Context) AutoProvisionPtrOutput {
	return AutoProvision(e).ToAutoProvisionOutputWithContext(ctx).ToAutoProvisionPtrOutputWithContext(ctx)
}

func (e AutoProvision) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AutoProvision) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AutoProvision) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AutoProvision) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AutoProvisionOutput struct{ *pulumi.OutputState }

func (AutoProvisionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoProvision)(nil)).Elem()
}

func (o AutoProvisionOutput) ToAutoProvisionOutput() AutoProvisionOutput {
	return o
}

func (o AutoProvisionOutput) ToAutoProvisionOutputWithContext(ctx context.Context) AutoProvisionOutput {
	return o
}

func (o AutoProvisionOutput) ToAutoProvisionPtrOutput() AutoProvisionPtrOutput {
	return o.ToAutoProvisionPtrOutputWithContext(context.Background())
}

func (o AutoProvisionOutput) ToAutoProvisionPtrOutputWithContext(ctx context.Context) AutoProvisionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AutoProvision) *AutoProvision {
		return &v
	}).(AutoProvisionPtrOutput)
}

func (o AutoProvisionOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AutoProvisionOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AutoProvision) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AutoProvisionOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AutoProvisionOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AutoProvision) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AutoProvisionPtrOutput struct{ *pulumi.OutputState }

func (AutoProvisionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoProvision)(nil)).Elem()
}

func (o AutoProvisionPtrOutput) ToAutoProvisionPtrOutput() AutoProvisionPtrOutput {
	return o
}

func (o AutoProvisionPtrOutput) ToAutoProvisionPtrOutputWithContext(ctx context.Context) AutoProvisionPtrOutput {
	return o
}

func (o AutoProvisionPtrOutput) Elem() AutoProvisionOutput {
	return o.ApplyT(func(v *AutoProvision) AutoProvision {
		if v != nil {
			return *v
		}
		var ret AutoProvision
		return ret
	}).(AutoProvisionOutput)
}

func (o AutoProvisionPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AutoProvisionPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AutoProvision) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// AutoProvisionInput is an input type that accepts values of the AutoProvision enum
// A concrete instance of `AutoProvisionInput` can be one of the following:
//
//	AutoProvisionOn
//	AutoProvisionOff
type AutoProvisionInput interface {
	pulumi.Input

	ToAutoProvisionOutput() AutoProvisionOutput
	ToAutoProvisionOutputWithContext(context.Context) AutoProvisionOutput
}

var autoProvisionPtrType = reflect.TypeOf((**AutoProvision)(nil)).Elem()

type AutoProvisionPtrInput interface {
	pulumi.Input

	ToAutoProvisionPtrOutput() AutoProvisionPtrOutput
	ToAutoProvisionPtrOutputWithContext(context.Context) AutoProvisionPtrOutput
}

type autoProvisionPtr string

func AutoProvisionPtr(v string) AutoProvisionPtrInput {
	return (*autoProvisionPtr)(&v)
}

func (*autoProvisionPtr) ElementType() reflect.Type {
	return autoProvisionPtrType
}

func (in *autoProvisionPtr) ToAutoProvisionPtrOutput() AutoProvisionPtrOutput {
	return pulumi.ToOutput(in).(AutoProvisionPtrOutput)
}

func (in *autoProvisionPtr) ToAutoProvisionPtrOutputWithContext(ctx context.Context) AutoProvisionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AutoProvisionPtrOutput)
}

// Defines the minimal alert severity which will be sent as email notifications
type MinimalSeverity string

const (
	// Get notifications on new alerts with High severity
	MinimalSeverityHigh = MinimalSeverity("High")
	// Get notifications on new alerts with medium or high severity
	MinimalSeverityMedium = MinimalSeverity("Medium")
	// Don't get notifications on new alerts with low, medium or high severity
	MinimalSeverityLow = MinimalSeverity("Low")
)

func (MinimalSeverity) ElementType() reflect.Type {
	return reflect.TypeOf((*MinimalSeverity)(nil)).Elem()
}

func (e MinimalSeverity) ToMinimalSeverityOutput() MinimalSeverityOutput {
	return pulumi.ToOutput(e).(MinimalSeverityOutput)
}

func (e MinimalSeverity) ToMinimalSeverityOutputWithContext(ctx context.Context) MinimalSeverityOutput {
	return pulumi.ToOutputWithContext(ctx, e).(MinimalSeverityOutput)
}

func (e MinimalSeverity) ToMinimalSeverityPtrOutput() MinimalSeverityPtrOutput {
	return e.ToMinimalSeverityPtrOutputWithContext(context.Background())
}

func (e MinimalSeverity) ToMinimalSeverityPtrOutputWithContext(ctx context.Context) MinimalSeverityPtrOutput {
	return MinimalSeverity(e).ToMinimalSeverityOutputWithContext(ctx).ToMinimalSeverityPtrOutputWithContext(ctx)
}

func (e MinimalSeverity) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e MinimalSeverity) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e MinimalSeverity) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e MinimalSeverity) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type MinimalSeverityOutput struct{ *pulumi.OutputState }

func (MinimalSeverityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MinimalSeverity)(nil)).Elem()
}

func (o MinimalSeverityOutput) ToMinimalSeverityOutput() MinimalSeverityOutput {
	return o
}

func (o MinimalSeverityOutput) ToMinimalSeverityOutputWithContext(ctx context.Context) MinimalSeverityOutput {
	return o
}

func (o MinimalSeverityOutput) ToMinimalSeverityPtrOutput() MinimalSeverityPtrOutput {
	return o.ToMinimalSeverityPtrOutputWithContext(context.Background())
}

func (o MinimalSeverityOutput) ToMinimalSeverityPtrOutputWithContext(ctx context.Context) MinimalSeverityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MinimalSeverity) *MinimalSeverity {
		return &v
	}).(MinimalSeverityPtrOutput)
}

func (o MinimalSeverityOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o MinimalSeverityOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e MinimalSeverity) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o MinimalSeverityOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o MinimalSeverityOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e MinimalSeverity) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type MinimalSeverityPtrOutput struct{ *pulumi.OutputState }

func (MinimalSeverityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MinimalSeverity)(nil)).Elem()
}

func (o MinimalSeverityPtrOutput) ToMinimalSeverityPtrOutput() MinimalSeverityPtrOutput {
	return o
}

func (o MinimalSeverityPtrOutput) ToMinimalSeverityPtrOutputWithContext(ctx context.Context) MinimalSeverityPtrOutput {
	return o
}

func (o MinimalSeverityPtrOutput) Elem() MinimalSeverityOutput {
	return o.ApplyT(func(v *MinimalSeverity) MinimalSeverity {
		if v != nil {
			return *v
		}
		var ret MinimalSeverity
		return ret
	}).(MinimalSeverityOutput)
}

func (o MinimalSeverityPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o MinimalSeverityPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *MinimalSeverity) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// MinimalSeverityInput is an input type that accepts values of the MinimalSeverity enum
// A concrete instance of `MinimalSeverityInput` can be one of the following:
//
//	MinimalSeverityHigh
//	MinimalSeverityMedium
//	MinimalSeverityLow
type MinimalSeverityInput interface {
	pulumi.Input

	ToMinimalSeverityOutput() MinimalSeverityOutput
	ToMinimalSeverityOutputWithContext(context.Context) MinimalSeverityOutput
}

var minimalSeverityPtrType = reflect.TypeOf((**MinimalSeverity)(nil)).Elem()

type MinimalSeverityPtrInput interface {
	pulumi.Input

	ToMinimalSeverityPtrOutput() MinimalSeverityPtrOutput
	ToMinimalSeverityPtrOutputWithContext(context.Context) MinimalSeverityPtrOutput
}

type minimalSeverityPtr string

func MinimalSeverityPtr(v string) MinimalSeverityPtrInput {
	return (*minimalSeverityPtr)(&v)
}

func (*minimalSeverityPtr) ElementType() reflect.Type {
	return minimalSeverityPtrType
}

func (in *minimalSeverityPtr) ToMinimalSeverityPtrOutput() MinimalSeverityPtrOutput {
	return pulumi.ToOutput(in).(MinimalSeverityPtrOutput)
}

func (in *minimalSeverityPtr) ToMinimalSeverityPtrOutputWithContext(ctx context.Context) MinimalSeverityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(MinimalSeverityPtrOutput)
}

// A possible role to configure sending security notification alerts to
type Roles string

const (
	// If enabled, send notification on new alerts to the account admins
	RolesAccountAdmin = Roles("AccountAdmin")
	// If enabled, send notification on new alerts to the service admins
	RolesServiceAdmin = Roles("ServiceAdmin")
	// If enabled, send notification on new alerts to the subscription owners
	RolesOwner = Roles("Owner")
	// If enabled, send notification on new alerts to the subscription contributors
	RolesContributor = Roles("Contributor")
)

func (Roles) ElementType() reflect.Type {
	return reflect.TypeOf((*Roles)(nil)).Elem()
}

func (e Roles) ToRolesOutput() RolesOutput {
	return pulumi.ToOutput(e).(RolesOutput)
}

func (e Roles) ToRolesOutputWithContext(ctx context.Context) RolesOutput {
	return pulumi.ToOutputWithContext(ctx, e).(RolesOutput)
}

func (e Roles) ToRolesPtrOutput() RolesPtrOutput {
	return e.ToRolesPtrOutputWithContext(context.Background())
}

func (e Roles) ToRolesPtrOutputWithContext(ctx context.Context) RolesPtrOutput {
	return Roles(e).ToRolesOutputWithContext(ctx).ToRolesPtrOutputWithContext(ctx)
}

func (e Roles) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Roles) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Roles) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Roles) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type RolesOutput struct{ *pulumi.OutputState }

func (RolesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Roles)(nil)).Elem()
}

func (o RolesOutput) ToRolesOutput() RolesOutput {
	return o
}

func (o RolesOutput) ToRolesOutputWithContext(ctx context.Context) RolesOutput {
	return o
}

func (o RolesOutput) ToRolesPtrOutput() RolesPtrOutput {
	return o.ToRolesPtrOutputWithContext(context.Background())
}

func (o RolesOutput) ToRolesPtrOutputWithContext(ctx context.Context) RolesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Roles) *Roles {
		return &v
	}).(RolesPtrOutput)
}

func (o RolesOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o RolesOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Roles) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o RolesOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RolesOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Roles) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type RolesPtrOutput struct{ *pulumi.OutputState }

func (RolesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Roles)(nil)).Elem()
}

func (o RolesPtrOutput) ToRolesPtrOutput() RolesPtrOutput {
	return o
}

func (o RolesPtrOutput) ToRolesPtrOutputWithContext(ctx context.Context) RolesPtrOutput {
	return o
}

func (o RolesPtrOutput) Elem() RolesOutput {
	return o.ApplyT(func(v *Roles) Roles {
		if v != nil {
			return *v
		}
		var ret Roles
		return ret
	}).(RolesOutput)
}

func (o RolesPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RolesPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *Roles) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// RolesInput is an input type that accepts values of the Roles enum
// A concrete instance of `RolesInput` can be one of the following:
//
//	RolesAccountAdmin
//	RolesServiceAdmin
//	RolesOwner
//	RolesContributor
type RolesInput interface {
	pulumi.Input

	ToRolesOutput() RolesOutput
	ToRolesOutputWithContext(context.Context) RolesOutput
}

var rolesPtrType = reflect.TypeOf((**Roles)(nil)).Elem()

type RolesPtrInput interface {
	pulumi.Input

	ToRolesPtrOutput() RolesPtrOutput
	ToRolesPtrOutputWithContext(context.Context) RolesPtrOutput
}

type rolesPtr string

func RolesPtr(v string) RolesPtrInput {
	return (*rolesPtr)(&v)
}

func (*rolesPtr) ElementType() reflect.Type {
	return rolesPtrType
}

func (in *rolesPtr) ToRolesPtrOutput() RolesPtrOutput {
	return pulumi.ToOutput(in).(RolesPtrOutput)
}

func (in *rolesPtr) ToRolesPtrOutputWithContext(ctx context.Context) RolesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(RolesPtrOutput)
}

// Defines whether to send email notifications from AMicrosoft Defender for Cloud to persons with specific RBAC roles on the subscription.
type State string

const (
	// Send notification on new alerts to the subscription's admins
	StateOn = State("On")
	// Don't send notification on new alerts to the subscription's admins
	StateOff = State("Off")
)

func (State) ElementType() reflect.Type {
	return reflect.TypeOf((*State)(nil)).Elem()
}

func (e State) ToStateOutput() StateOutput {
	return pulumi.ToOutput(e).(StateOutput)
}

func (e State) ToStateOutputWithContext(ctx context.Context) StateOutput {
	return pulumi.ToOutputWithContext(ctx, e).(StateOutput)
}

func (e State) ToStatePtrOutput() StatePtrOutput {
	return e.ToStatePtrOutputWithContext(context.Background())
}

func (e State) ToStatePtrOutputWithContext(ctx context.Context) StatePtrOutput {
	return State(e).ToStateOutputWithContext(ctx).ToStatePtrOutputWithContext(ctx)
}

func (e State) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e State) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e State) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e State) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type StateOutput struct{ *pulumi.OutputState }

func (StateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*State)(nil)).Elem()
}

func (o StateOutput) ToStateOutput() StateOutput {
	return o
}

func (o StateOutput) ToStateOutputWithContext(ctx context.Context) StateOutput {
	return o
}

func (o StateOutput) ToStatePtrOutput() StatePtrOutput {
	return o.ToStatePtrOutputWithContext(context.Background())
}

func (o StateOutput) ToStatePtrOutputWithContext(ctx context.Context) StatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v State) *State {
		return &v
	}).(StatePtrOutput)
}

func (o StateOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o StateOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e State) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o StateOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o StateOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e State) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type StatePtrOutput struct{ *pulumi.OutputState }

func (StatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**State)(nil)).Elem()
}

func (o StatePtrOutput) ToStatePtrOutput() StatePtrOutput {
	return o
}

func (o StatePtrOutput) ToStatePtrOutputWithContext(ctx context.Context) StatePtrOutput {
	return o
}

func (o StatePtrOutput) Elem() StateOutput {
	return o.ApplyT(func(v *State) State {
		if v != nil {
			return *v
		}
		var ret State
		return ret
	}).(StateOutput)
}

func (o StatePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o StatePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *State) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// StateInput is an input type that accepts values of the State enum
// A concrete instance of `StateInput` can be one of the following:
//
//	StateOn
//	StateOff
type StateInput interface {
	pulumi.Input

	ToStateOutput() StateOutput
	ToStateOutputWithContext(context.Context) StateOutput
}

var statePtrType = reflect.TypeOf((**State)(nil)).Elem()

type StatePtrInput interface {
	pulumi.Input

	ToStatePtrOutput() StatePtrOutput
	ToStatePtrOutputWithContext(context.Context) StatePtrOutput
}

type statePtr string

func StatePtr(v string) StatePtrInput {
	return (*statePtr)(&v)
}

func (*statePtr) ElementType() reflect.Type {
	return statePtrType
}

func (in *statePtr) ToStatePtrOutput() StatePtrOutput {
	return pulumi.ToOutput(in).(StatePtrOutput)
}

func (in *statePtr) ToStatePtrOutputWithContext(ctx context.Context) StatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(StatePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AutoProvisionOutput{})
	pulumi.RegisterOutputType(AutoProvisionPtrOutput{})
	pulumi.RegisterOutputType(MinimalSeverityOutput{})
	pulumi.RegisterOutputType(MinimalSeverityPtrOutput{})
	pulumi.RegisterOutputType(RolesOutput{})
	pulumi.RegisterOutputType(RolesPtrOutput{})
	pulumi.RegisterOutputType(StateOutput{})
	pulumi.RegisterOutputType(StatePtrOutput{})
}
