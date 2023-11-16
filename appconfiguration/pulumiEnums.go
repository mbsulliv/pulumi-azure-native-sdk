// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package appconfiguration

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The private link service connection status.
type ConnectionStatus string

const (
	ConnectionStatusPending      = ConnectionStatus("Pending")
	ConnectionStatusApproved     = ConnectionStatus("Approved")
	ConnectionStatusRejected     = ConnectionStatus("Rejected")
	ConnectionStatusDisconnected = ConnectionStatus("Disconnected")
)

// Indicates whether the configuration store need to be recovered.
type CreateMode string

const (
	CreateModeRecover = CreateMode("Recover")
	CreateModeDefault = CreateMode("Default")
)

func (CreateMode) ElementType() reflect.Type {
	return reflect.TypeOf((*CreateMode)(nil)).Elem()
}

func (e CreateMode) ToCreateModeOutput() CreateModeOutput {
	return pulumi.ToOutput(e).(CreateModeOutput)
}

func (e CreateMode) ToCreateModeOutputWithContext(ctx context.Context) CreateModeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(CreateModeOutput)
}

func (e CreateMode) ToCreateModePtrOutput() CreateModePtrOutput {
	return e.ToCreateModePtrOutputWithContext(context.Background())
}

func (e CreateMode) ToCreateModePtrOutputWithContext(ctx context.Context) CreateModePtrOutput {
	return CreateMode(e).ToCreateModeOutputWithContext(ctx).ToCreateModePtrOutputWithContext(ctx)
}

func (e CreateMode) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e CreateMode) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e CreateMode) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e CreateMode) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type CreateModeOutput struct{ *pulumi.OutputState }

func (CreateModeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CreateMode)(nil)).Elem()
}

func (o CreateModeOutput) ToCreateModeOutput() CreateModeOutput {
	return o
}

func (o CreateModeOutput) ToCreateModeOutputWithContext(ctx context.Context) CreateModeOutput {
	return o
}

func (o CreateModeOutput) ToCreateModePtrOutput() CreateModePtrOutput {
	return o.ToCreateModePtrOutputWithContext(context.Background())
}

func (o CreateModeOutput) ToCreateModePtrOutputWithContext(ctx context.Context) CreateModePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CreateMode) *CreateMode {
		return &v
	}).(CreateModePtrOutput)
}

func (o CreateModeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o CreateModeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CreateMode) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o CreateModeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CreateModeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CreateMode) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type CreateModePtrOutput struct{ *pulumi.OutputState }

func (CreateModePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CreateMode)(nil)).Elem()
}

func (o CreateModePtrOutput) ToCreateModePtrOutput() CreateModePtrOutput {
	return o
}

func (o CreateModePtrOutput) ToCreateModePtrOutputWithContext(ctx context.Context) CreateModePtrOutput {
	return o
}

func (o CreateModePtrOutput) Elem() CreateModeOutput {
	return o.ApplyT(func(v *CreateMode) CreateMode {
		if v != nil {
			return *v
		}
		var ret CreateMode
		return ret
	}).(CreateModeOutput)
}

func (o CreateModePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CreateModePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *CreateMode) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// CreateModeInput is an input type that accepts CreateModeArgs and CreateModeOutput values.
// You can construct a concrete instance of `CreateModeInput` via:
//
//	CreateModeArgs{...}
type CreateModeInput interface {
	pulumi.Input

	ToCreateModeOutput() CreateModeOutput
	ToCreateModeOutputWithContext(context.Context) CreateModeOutput
}

var createModePtrType = reflect.TypeOf((**CreateMode)(nil)).Elem()

type CreateModePtrInput interface {
	pulumi.Input

	ToCreateModePtrOutput() CreateModePtrOutput
	ToCreateModePtrOutputWithContext(context.Context) CreateModePtrOutput
}

type createModePtr string

func CreateModePtr(v string) CreateModePtrInput {
	return (*createModePtr)(&v)
}

func (*createModePtr) ElementType() reflect.Type {
	return createModePtrType
}

func (in *createModePtr) ToCreateModePtrOutput() CreateModePtrOutput {
	return pulumi.ToOutput(in).(CreateModePtrOutput)
}

func (in *createModePtr) ToCreateModePtrOutputWithContext(ctx context.Context) CreateModePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(CreateModePtrOutput)
}

func (in *createModePtr) ToOutput(ctx context.Context) pulumix.Output[*CreateMode] {
	return pulumix.Output[*CreateMode]{
		OutputState: in.ToCreateModePtrOutputWithContext(ctx).OutputState,
	}
}

// The type of managed identity used. The type 'SystemAssigned, UserAssigned' includes both an implicitly created identity and a set of user-assigned identities. The type 'None' will remove any identities.
type IdentityType string

const (
	IdentityTypeNone                         = IdentityType("None")
	IdentityTypeSystemAssigned               = IdentityType("SystemAssigned")
	IdentityTypeUserAssigned                 = IdentityType("UserAssigned")
	IdentityType_SystemAssigned_UserAssigned = IdentityType("SystemAssigned, UserAssigned")
)

// Control permission for data plane traffic coming from public networks while private endpoint is enabled.
type PublicNetworkAccess string

const (
	PublicNetworkAccessEnabled  = PublicNetworkAccess("Enabled")
	PublicNetworkAccessDisabled = PublicNetworkAccess("Disabled")
)

func init() {
	pulumi.RegisterOutputType(CreateModeOutput{})
	pulumi.RegisterOutputType(CreateModePtrOutput{})
}
