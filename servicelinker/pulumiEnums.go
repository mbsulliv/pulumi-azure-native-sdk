// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package servicelinker

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AccessKeyPermissions string

const (
	AccessKeyPermissionsRead   = AccessKeyPermissions("Read")
	AccessKeyPermissionsWrite  = AccessKeyPermissions("Write")
	AccessKeyPermissionsListen = AccessKeyPermissions("Listen")
	AccessKeyPermissionsSend   = AccessKeyPermissions("Send")
	AccessKeyPermissionsManage = AccessKeyPermissions("Manage")
)

func (AccessKeyPermissions) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessKeyPermissions)(nil)).Elem()
}

func (e AccessKeyPermissions) ToAccessKeyPermissionsOutput() AccessKeyPermissionsOutput {
	return pulumi.ToOutput(e).(AccessKeyPermissionsOutput)
}

func (e AccessKeyPermissions) ToAccessKeyPermissionsOutputWithContext(ctx context.Context) AccessKeyPermissionsOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AccessKeyPermissionsOutput)
}

func (e AccessKeyPermissions) ToAccessKeyPermissionsPtrOutput() AccessKeyPermissionsPtrOutput {
	return e.ToAccessKeyPermissionsPtrOutputWithContext(context.Background())
}

func (e AccessKeyPermissions) ToAccessKeyPermissionsPtrOutputWithContext(ctx context.Context) AccessKeyPermissionsPtrOutput {
	return AccessKeyPermissions(e).ToAccessKeyPermissionsOutputWithContext(ctx).ToAccessKeyPermissionsPtrOutputWithContext(ctx)
}

func (e AccessKeyPermissions) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AccessKeyPermissions) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AccessKeyPermissions) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AccessKeyPermissions) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AccessKeyPermissionsOutput struct{ *pulumi.OutputState }

func (AccessKeyPermissionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessKeyPermissions)(nil)).Elem()
}

func (o AccessKeyPermissionsOutput) ToAccessKeyPermissionsOutput() AccessKeyPermissionsOutput {
	return o
}

func (o AccessKeyPermissionsOutput) ToAccessKeyPermissionsOutputWithContext(ctx context.Context) AccessKeyPermissionsOutput {
	return o
}

func (o AccessKeyPermissionsOutput) ToAccessKeyPermissionsPtrOutput() AccessKeyPermissionsPtrOutput {
	return o.ToAccessKeyPermissionsPtrOutputWithContext(context.Background())
}

func (o AccessKeyPermissionsOutput) ToAccessKeyPermissionsPtrOutputWithContext(ctx context.Context) AccessKeyPermissionsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AccessKeyPermissions) *AccessKeyPermissions {
		return &v
	}).(AccessKeyPermissionsPtrOutput)
}

func (o AccessKeyPermissionsOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AccessKeyPermissionsOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AccessKeyPermissions) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AccessKeyPermissionsOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AccessKeyPermissionsOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AccessKeyPermissions) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AccessKeyPermissionsPtrOutput struct{ *pulumi.OutputState }

func (AccessKeyPermissionsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessKeyPermissions)(nil)).Elem()
}

func (o AccessKeyPermissionsPtrOutput) ToAccessKeyPermissionsPtrOutput() AccessKeyPermissionsPtrOutput {
	return o
}

func (o AccessKeyPermissionsPtrOutput) ToAccessKeyPermissionsPtrOutputWithContext(ctx context.Context) AccessKeyPermissionsPtrOutput {
	return o
}

func (o AccessKeyPermissionsPtrOutput) Elem() AccessKeyPermissionsOutput {
	return o.ApplyT(func(v *AccessKeyPermissions) AccessKeyPermissions {
		if v != nil {
			return *v
		}
		var ret AccessKeyPermissions
		return ret
	}).(AccessKeyPermissionsOutput)
}

func (o AccessKeyPermissionsPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AccessKeyPermissionsPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AccessKeyPermissions) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// AccessKeyPermissionsInput is an input type that accepts values of the AccessKeyPermissions enum
// A concrete instance of `AccessKeyPermissionsInput` can be one of the following:
//
//	AccessKeyPermissionsRead
//	AccessKeyPermissionsWrite
//	AccessKeyPermissionsListen
//	AccessKeyPermissionsSend
//	AccessKeyPermissionsManage
type AccessKeyPermissionsInput interface {
	pulumi.Input

	ToAccessKeyPermissionsOutput() AccessKeyPermissionsOutput
	ToAccessKeyPermissionsOutputWithContext(context.Context) AccessKeyPermissionsOutput
}

var accessKeyPermissionsPtrType = reflect.TypeOf((**AccessKeyPermissions)(nil)).Elem()

type AccessKeyPermissionsPtrInput interface {
	pulumi.Input

	ToAccessKeyPermissionsPtrOutput() AccessKeyPermissionsPtrOutput
	ToAccessKeyPermissionsPtrOutputWithContext(context.Context) AccessKeyPermissionsPtrOutput
}

type accessKeyPermissionsPtr string

func AccessKeyPermissionsPtr(v string) AccessKeyPermissionsPtrInput {
	return (*accessKeyPermissionsPtr)(&v)
}

func (*accessKeyPermissionsPtr) ElementType() reflect.Type {
	return accessKeyPermissionsPtrType
}

func (in *accessKeyPermissionsPtr) ToAccessKeyPermissionsPtrOutput() AccessKeyPermissionsPtrOutput {
	return pulumi.ToOutput(in).(AccessKeyPermissionsPtrOutput)
}

func (in *accessKeyPermissionsPtr) ToAccessKeyPermissionsPtrOutputWithContext(ctx context.Context) AccessKeyPermissionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AccessKeyPermissionsPtrOutput)
}

// Optional. Indicates public network solution. If enable, enable public network access of target service with best try. Default is enable. If optOut, opt out public network access configuration.
type ActionType string

const (
	ActionTypeEnable = ActionType("enable")
	ActionTypeOptOut = ActionType("optOut")
)

func (ActionType) ElementType() reflect.Type {
	return reflect.TypeOf((*ActionType)(nil)).Elem()
}

func (e ActionType) ToActionTypeOutput() ActionTypeOutput {
	return pulumi.ToOutput(e).(ActionTypeOutput)
}

func (e ActionType) ToActionTypeOutputWithContext(ctx context.Context) ActionTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ActionTypeOutput)
}

func (e ActionType) ToActionTypePtrOutput() ActionTypePtrOutput {
	return e.ToActionTypePtrOutputWithContext(context.Background())
}

func (e ActionType) ToActionTypePtrOutputWithContext(ctx context.Context) ActionTypePtrOutput {
	return ActionType(e).ToActionTypeOutputWithContext(ctx).ToActionTypePtrOutputWithContext(ctx)
}

func (e ActionType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ActionType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ActionType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ActionType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ActionTypeOutput struct{ *pulumi.OutputState }

func (ActionTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ActionType)(nil)).Elem()
}

func (o ActionTypeOutput) ToActionTypeOutput() ActionTypeOutput {
	return o
}

func (o ActionTypeOutput) ToActionTypeOutputWithContext(ctx context.Context) ActionTypeOutput {
	return o
}

func (o ActionTypeOutput) ToActionTypePtrOutput() ActionTypePtrOutput {
	return o.ToActionTypePtrOutputWithContext(context.Background())
}

func (o ActionTypeOutput) ToActionTypePtrOutputWithContext(ctx context.Context) ActionTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ActionType) *ActionType {
		return &v
	}).(ActionTypePtrOutput)
}

func (o ActionTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ActionTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ActionType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ActionTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ActionTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ActionType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ActionTypePtrOutput struct{ *pulumi.OutputState }

func (ActionTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ActionType)(nil)).Elem()
}

func (o ActionTypePtrOutput) ToActionTypePtrOutput() ActionTypePtrOutput {
	return o
}

func (o ActionTypePtrOutput) ToActionTypePtrOutputWithContext(ctx context.Context) ActionTypePtrOutput {
	return o
}

func (o ActionTypePtrOutput) Elem() ActionTypeOutput {
	return o.ApplyT(func(v *ActionType) ActionType {
		if v != nil {
			return *v
		}
		var ret ActionType
		return ret
	}).(ActionTypeOutput)
}

func (o ActionTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ActionTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ActionType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// ActionTypeInput is an input type that accepts values of the ActionType enum
// A concrete instance of `ActionTypeInput` can be one of the following:
//
//	ActionTypeEnable
//	ActionTypeOptOut
type ActionTypeInput interface {
	pulumi.Input

	ToActionTypeOutput() ActionTypeOutput
	ToActionTypeOutputWithContext(context.Context) ActionTypeOutput
}

var actionTypePtrType = reflect.TypeOf((**ActionType)(nil)).Elem()

type ActionTypePtrInput interface {
	pulumi.Input

	ToActionTypePtrOutput() ActionTypePtrOutput
	ToActionTypePtrOutputWithContext(context.Context) ActionTypePtrOutput
}

type actionTypePtr string

func ActionTypePtr(v string) ActionTypePtrInput {
	return (*actionTypePtr)(&v)
}

func (*actionTypePtr) ElementType() reflect.Type {
	return actionTypePtrType
}

func (in *actionTypePtr) ToActionTypePtrOutput() ActionTypePtrOutput {
	return pulumi.ToOutput(in).(ActionTypePtrOutput)
}

func (in *actionTypePtr) ToActionTypePtrOutputWithContext(ctx context.Context) ActionTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ActionTypePtrOutput)
}

// Allow caller client IP to access the target service if true. the property is used when connecting local application to target service.
type AllowType string

const (
	AllowTypeTrue  = AllowType("true")
	AllowTypeFalse = AllowType("false")
)

func (AllowType) ElementType() reflect.Type {
	return reflect.TypeOf((*AllowType)(nil)).Elem()
}

func (e AllowType) ToAllowTypeOutput() AllowTypeOutput {
	return pulumi.ToOutput(e).(AllowTypeOutput)
}

func (e AllowType) ToAllowTypeOutputWithContext(ctx context.Context) AllowTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AllowTypeOutput)
}

func (e AllowType) ToAllowTypePtrOutput() AllowTypePtrOutput {
	return e.ToAllowTypePtrOutputWithContext(context.Background())
}

func (e AllowType) ToAllowTypePtrOutputWithContext(ctx context.Context) AllowTypePtrOutput {
	return AllowType(e).ToAllowTypeOutputWithContext(ctx).ToAllowTypePtrOutputWithContext(ctx)
}

func (e AllowType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AllowType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AllowType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AllowType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AllowTypeOutput struct{ *pulumi.OutputState }

func (AllowTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AllowType)(nil)).Elem()
}

func (o AllowTypeOutput) ToAllowTypeOutput() AllowTypeOutput {
	return o
}

func (o AllowTypeOutput) ToAllowTypeOutputWithContext(ctx context.Context) AllowTypeOutput {
	return o
}

func (o AllowTypeOutput) ToAllowTypePtrOutput() AllowTypePtrOutput {
	return o.ToAllowTypePtrOutputWithContext(context.Background())
}

func (o AllowTypeOutput) ToAllowTypePtrOutputWithContext(ctx context.Context) AllowTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AllowType) *AllowType {
		return &v
	}).(AllowTypePtrOutput)
}

func (o AllowTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AllowTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AllowType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AllowTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AllowTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AllowType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AllowTypePtrOutput struct{ *pulumi.OutputState }

func (AllowTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AllowType)(nil)).Elem()
}

func (o AllowTypePtrOutput) ToAllowTypePtrOutput() AllowTypePtrOutput {
	return o
}

func (o AllowTypePtrOutput) ToAllowTypePtrOutputWithContext(ctx context.Context) AllowTypePtrOutput {
	return o
}

func (o AllowTypePtrOutput) Elem() AllowTypeOutput {
	return o.ApplyT(func(v *AllowType) AllowType {
		if v != nil {
			return *v
		}
		var ret AllowType
		return ret
	}).(AllowTypeOutput)
}

func (o AllowTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AllowTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AllowType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// AllowTypeInput is an input type that accepts values of the AllowType enum
// A concrete instance of `AllowTypeInput` can be one of the following:
//
//	AllowTypeTrue
//	AllowTypeFalse
type AllowTypeInput interface {
	pulumi.Input

	ToAllowTypeOutput() AllowTypeOutput
	ToAllowTypeOutputWithContext(context.Context) AllowTypeOutput
}

var allowTypePtrType = reflect.TypeOf((**AllowType)(nil)).Elem()

type AllowTypePtrInput interface {
	pulumi.Input

	ToAllowTypePtrOutput() AllowTypePtrOutput
	ToAllowTypePtrOutputWithContext(context.Context) AllowTypePtrOutput
}

type allowTypePtr string

func AllowTypePtr(v string) AllowTypePtrInput {
	return (*allowTypePtr)(&v)
}

func (*allowTypePtr) ElementType() reflect.Type {
	return allowTypePtrType
}

func (in *allowTypePtr) ToAllowTypePtrOutput() AllowTypePtrOutput {
	return pulumi.ToOutput(in).(AllowTypePtrOutput)
}

func (in *allowTypePtr) ToAllowTypePtrOutputWithContext(ctx context.Context) AllowTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AllowTypePtrOutput)
}

// The authentication type.
type AuthType string

const (
	AuthTypeSystemAssignedIdentity      = AuthType("systemAssignedIdentity")
	AuthTypeUserAssignedIdentity        = AuthType("userAssignedIdentity")
	AuthTypeServicePrincipalSecret      = AuthType("servicePrincipalSecret")
	AuthTypeServicePrincipalCertificate = AuthType("servicePrincipalCertificate")
	AuthTypeSecret                      = AuthType("secret")
	AuthTypeAccessKey                   = AuthType("accessKey")
	AuthTypeUserAccount                 = AuthType("userAccount")
)

// The azure resource type.
type AzureResourceType string

const (
	AzureResourceTypeKeyVault = AzureResourceType("KeyVault")
)

// The application client type
type ClientType string

const (
	ClientTypeNone              = ClientType("none")
	ClientTypeDotnet            = ClientType("dotnet")
	ClientTypeJava              = ClientType("java")
	ClientTypePython            = ClientType("python")
	ClientTypeGo                = ClientType("go")
	ClientTypePhp               = ClientType("php")
	ClientTypeRuby              = ClientType("ruby")
	ClientTypeDjango            = ClientType("django")
	ClientTypeNodejs            = ClientType("nodejs")
	ClientTypeSpringBoot        = ClientType("springBoot")
	ClientType_Kafka_SpringBoot = ClientType("kafka-springBoot")
	ClientTypeDapr              = ClientType("dapr")
)

func (ClientType) ElementType() reflect.Type {
	return reflect.TypeOf((*ClientType)(nil)).Elem()
}

func (e ClientType) ToClientTypeOutput() ClientTypeOutput {
	return pulumi.ToOutput(e).(ClientTypeOutput)
}

func (e ClientType) ToClientTypeOutputWithContext(ctx context.Context) ClientTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ClientTypeOutput)
}

func (e ClientType) ToClientTypePtrOutput() ClientTypePtrOutput {
	return e.ToClientTypePtrOutputWithContext(context.Background())
}

func (e ClientType) ToClientTypePtrOutputWithContext(ctx context.Context) ClientTypePtrOutput {
	return ClientType(e).ToClientTypeOutputWithContext(ctx).ToClientTypePtrOutputWithContext(ctx)
}

func (e ClientType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ClientType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ClientType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ClientType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ClientTypeOutput struct{ *pulumi.OutputState }

func (ClientTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClientType)(nil)).Elem()
}

func (o ClientTypeOutput) ToClientTypeOutput() ClientTypeOutput {
	return o
}

func (o ClientTypeOutput) ToClientTypeOutputWithContext(ctx context.Context) ClientTypeOutput {
	return o
}

func (o ClientTypeOutput) ToClientTypePtrOutput() ClientTypePtrOutput {
	return o.ToClientTypePtrOutputWithContext(context.Background())
}

func (o ClientTypeOutput) ToClientTypePtrOutputWithContext(ctx context.Context) ClientTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ClientType) *ClientType {
		return &v
	}).(ClientTypePtrOutput)
}

func (o ClientTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ClientTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ClientType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ClientTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ClientTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ClientType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ClientTypePtrOutput struct{ *pulumi.OutputState }

func (ClientTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClientType)(nil)).Elem()
}

func (o ClientTypePtrOutput) ToClientTypePtrOutput() ClientTypePtrOutput {
	return o
}

func (o ClientTypePtrOutput) ToClientTypePtrOutputWithContext(ctx context.Context) ClientTypePtrOutput {
	return o
}

func (o ClientTypePtrOutput) Elem() ClientTypeOutput {
	return o.ApplyT(func(v *ClientType) ClientType {
		if v != nil {
			return *v
		}
		var ret ClientType
		return ret
	}).(ClientTypeOutput)
}

func (o ClientTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ClientTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ClientType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// ClientTypeInput is an input type that accepts values of the ClientType enum
// A concrete instance of `ClientTypeInput` can be one of the following:
//
//	ClientTypeNone
//	ClientTypeDotnet
//	ClientTypeJava
//	ClientTypePython
//	ClientTypeGo
//	ClientTypePhp
//	ClientTypeRuby
//	ClientTypeDjango
//	ClientTypeNodejs
//	ClientTypeSpringBoot
//	ClientType_Kafka_SpringBoot
//	ClientTypeDapr
type ClientTypeInput interface {
	pulumi.Input

	ToClientTypeOutput() ClientTypeOutput
	ToClientTypeOutputWithContext(context.Context) ClientTypeOutput
}

var clientTypePtrType = reflect.TypeOf((**ClientType)(nil)).Elem()

type ClientTypePtrInput interface {
	pulumi.Input

	ToClientTypePtrOutput() ClientTypePtrOutput
	ToClientTypePtrOutputWithContext(context.Context) ClientTypePtrOutput
}

type clientTypePtr string

func ClientTypePtr(v string) ClientTypePtrInput {
	return (*clientTypePtr)(&v)
}

func (*clientTypePtr) ElementType() reflect.Type {
	return clientTypePtrType
}

func (in *clientTypePtr) ToClientTypePtrOutput() ClientTypePtrOutput {
	return pulumi.ToOutput(in).(ClientTypePtrOutput)
}

func (in *clientTypePtr) ToClientTypePtrOutputWithContext(ctx context.Context) ClientTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ClientTypePtrOutput)
}

// Indicates whether to clean up previous operation when Linker is updating or deleting
type DeleteOrUpdateBehavior string

const (
	DeleteOrUpdateBehaviorDefault       = DeleteOrUpdateBehavior("Default")
	DeleteOrUpdateBehaviorForcedCleanup = DeleteOrUpdateBehavior("ForcedCleanup")
)

func (DeleteOrUpdateBehavior) ElementType() reflect.Type {
	return reflect.TypeOf((*DeleteOrUpdateBehavior)(nil)).Elem()
}

func (e DeleteOrUpdateBehavior) ToDeleteOrUpdateBehaviorOutput() DeleteOrUpdateBehaviorOutput {
	return pulumi.ToOutput(e).(DeleteOrUpdateBehaviorOutput)
}

func (e DeleteOrUpdateBehavior) ToDeleteOrUpdateBehaviorOutputWithContext(ctx context.Context) DeleteOrUpdateBehaviorOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DeleteOrUpdateBehaviorOutput)
}

func (e DeleteOrUpdateBehavior) ToDeleteOrUpdateBehaviorPtrOutput() DeleteOrUpdateBehaviorPtrOutput {
	return e.ToDeleteOrUpdateBehaviorPtrOutputWithContext(context.Background())
}

func (e DeleteOrUpdateBehavior) ToDeleteOrUpdateBehaviorPtrOutputWithContext(ctx context.Context) DeleteOrUpdateBehaviorPtrOutput {
	return DeleteOrUpdateBehavior(e).ToDeleteOrUpdateBehaviorOutputWithContext(ctx).ToDeleteOrUpdateBehaviorPtrOutputWithContext(ctx)
}

func (e DeleteOrUpdateBehavior) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DeleteOrUpdateBehavior) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DeleteOrUpdateBehavior) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DeleteOrUpdateBehavior) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DeleteOrUpdateBehaviorOutput struct{ *pulumi.OutputState }

func (DeleteOrUpdateBehaviorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeleteOrUpdateBehavior)(nil)).Elem()
}

func (o DeleteOrUpdateBehaviorOutput) ToDeleteOrUpdateBehaviorOutput() DeleteOrUpdateBehaviorOutput {
	return o
}

func (o DeleteOrUpdateBehaviorOutput) ToDeleteOrUpdateBehaviorOutputWithContext(ctx context.Context) DeleteOrUpdateBehaviorOutput {
	return o
}

func (o DeleteOrUpdateBehaviorOutput) ToDeleteOrUpdateBehaviorPtrOutput() DeleteOrUpdateBehaviorPtrOutput {
	return o.ToDeleteOrUpdateBehaviorPtrOutputWithContext(context.Background())
}

func (o DeleteOrUpdateBehaviorOutput) ToDeleteOrUpdateBehaviorPtrOutputWithContext(ctx context.Context) DeleteOrUpdateBehaviorPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DeleteOrUpdateBehavior) *DeleteOrUpdateBehavior {
		return &v
	}).(DeleteOrUpdateBehaviorPtrOutput)
}

func (o DeleteOrUpdateBehaviorOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DeleteOrUpdateBehaviorOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DeleteOrUpdateBehavior) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DeleteOrUpdateBehaviorOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DeleteOrUpdateBehaviorOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DeleteOrUpdateBehavior) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DeleteOrUpdateBehaviorPtrOutput struct{ *pulumi.OutputState }

func (DeleteOrUpdateBehaviorPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DeleteOrUpdateBehavior)(nil)).Elem()
}

func (o DeleteOrUpdateBehaviorPtrOutput) ToDeleteOrUpdateBehaviorPtrOutput() DeleteOrUpdateBehaviorPtrOutput {
	return o
}

func (o DeleteOrUpdateBehaviorPtrOutput) ToDeleteOrUpdateBehaviorPtrOutputWithContext(ctx context.Context) DeleteOrUpdateBehaviorPtrOutput {
	return o
}

func (o DeleteOrUpdateBehaviorPtrOutput) Elem() DeleteOrUpdateBehaviorOutput {
	return o.ApplyT(func(v *DeleteOrUpdateBehavior) DeleteOrUpdateBehavior {
		if v != nil {
			return *v
		}
		var ret DeleteOrUpdateBehavior
		return ret
	}).(DeleteOrUpdateBehaviorOutput)
}

func (o DeleteOrUpdateBehaviorPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DeleteOrUpdateBehaviorPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DeleteOrUpdateBehavior) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// DeleteOrUpdateBehaviorInput is an input type that accepts values of the DeleteOrUpdateBehavior enum
// A concrete instance of `DeleteOrUpdateBehaviorInput` can be one of the following:
//
//	DeleteOrUpdateBehaviorDefault
//	DeleteOrUpdateBehaviorForcedCleanup
type DeleteOrUpdateBehaviorInput interface {
	pulumi.Input

	ToDeleteOrUpdateBehaviorOutput() DeleteOrUpdateBehaviorOutput
	ToDeleteOrUpdateBehaviorOutputWithContext(context.Context) DeleteOrUpdateBehaviorOutput
}

var deleteOrUpdateBehaviorPtrType = reflect.TypeOf((**DeleteOrUpdateBehavior)(nil)).Elem()

type DeleteOrUpdateBehaviorPtrInput interface {
	pulumi.Input

	ToDeleteOrUpdateBehaviorPtrOutput() DeleteOrUpdateBehaviorPtrOutput
	ToDeleteOrUpdateBehaviorPtrOutputWithContext(context.Context) DeleteOrUpdateBehaviorPtrOutput
}

type deleteOrUpdateBehaviorPtr string

func DeleteOrUpdateBehaviorPtr(v string) DeleteOrUpdateBehaviorPtrInput {
	return (*deleteOrUpdateBehaviorPtr)(&v)
}

func (*deleteOrUpdateBehaviorPtr) ElementType() reflect.Type {
	return deleteOrUpdateBehaviorPtrType
}

func (in *deleteOrUpdateBehaviorPtr) ToDeleteOrUpdateBehaviorPtrOutput() DeleteOrUpdateBehaviorPtrOutput {
	return pulumi.ToOutput(in).(DeleteOrUpdateBehaviorPtrOutput)
}

func (in *deleteOrUpdateBehaviorPtr) ToDeleteOrUpdateBehaviorPtrOutputWithContext(ctx context.Context) DeleteOrUpdateBehaviorPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DeleteOrUpdateBehaviorPtrOutput)
}

// The name of action for you dryrun job.
type DryrunActionName string

const (
	DryrunActionNameCreateOrUpdate = DryrunActionName("createOrUpdate")
)

// The secret type.
type SecretType string

const (
	SecretTypeRawValue                = SecretType("rawValue")
	SecretTypeKeyVaultSecretUri       = SecretType("keyVaultSecretUri")
	SecretTypeKeyVaultSecretReference = SecretType("keyVaultSecretReference")
)

// The target service type.
type TargetServiceType string

const (
	TargetServiceTypeAzureResource            = TargetServiceType("AzureResource")
	TargetServiceTypeConfluentBootstrapServer = TargetServiceType("ConfluentBootstrapServer")
	TargetServiceTypeConfluentSchemaRegistry  = TargetServiceType("ConfluentSchemaRegistry")
	TargetServiceTypeSelfHostedServer         = TargetServiceType("SelfHostedServer")
)

// Type of VNet solution.
type VNetSolutionType string

const (
	VNetSolutionTypeServiceEndpoint = VNetSolutionType("serviceEndpoint")
	VNetSolutionTypePrivateLink     = VNetSolutionType("privateLink")
)

func (VNetSolutionType) ElementType() reflect.Type {
	return reflect.TypeOf((*VNetSolutionType)(nil)).Elem()
}

func (e VNetSolutionType) ToVNetSolutionTypeOutput() VNetSolutionTypeOutput {
	return pulumi.ToOutput(e).(VNetSolutionTypeOutput)
}

func (e VNetSolutionType) ToVNetSolutionTypeOutputWithContext(ctx context.Context) VNetSolutionTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(VNetSolutionTypeOutput)
}

func (e VNetSolutionType) ToVNetSolutionTypePtrOutput() VNetSolutionTypePtrOutput {
	return e.ToVNetSolutionTypePtrOutputWithContext(context.Background())
}

func (e VNetSolutionType) ToVNetSolutionTypePtrOutputWithContext(ctx context.Context) VNetSolutionTypePtrOutput {
	return VNetSolutionType(e).ToVNetSolutionTypeOutputWithContext(ctx).ToVNetSolutionTypePtrOutputWithContext(ctx)
}

func (e VNetSolutionType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e VNetSolutionType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e VNetSolutionType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e VNetSolutionType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type VNetSolutionTypeOutput struct{ *pulumi.OutputState }

func (VNetSolutionTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VNetSolutionType)(nil)).Elem()
}

func (o VNetSolutionTypeOutput) ToVNetSolutionTypeOutput() VNetSolutionTypeOutput {
	return o
}

func (o VNetSolutionTypeOutput) ToVNetSolutionTypeOutputWithContext(ctx context.Context) VNetSolutionTypeOutput {
	return o
}

func (o VNetSolutionTypeOutput) ToVNetSolutionTypePtrOutput() VNetSolutionTypePtrOutput {
	return o.ToVNetSolutionTypePtrOutputWithContext(context.Background())
}

func (o VNetSolutionTypeOutput) ToVNetSolutionTypePtrOutputWithContext(ctx context.Context) VNetSolutionTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VNetSolutionType) *VNetSolutionType {
		return &v
	}).(VNetSolutionTypePtrOutput)
}

func (o VNetSolutionTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o VNetSolutionTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e VNetSolutionType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o VNetSolutionTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o VNetSolutionTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e VNetSolutionType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type VNetSolutionTypePtrOutput struct{ *pulumi.OutputState }

func (VNetSolutionTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VNetSolutionType)(nil)).Elem()
}

func (o VNetSolutionTypePtrOutput) ToVNetSolutionTypePtrOutput() VNetSolutionTypePtrOutput {
	return o
}

func (o VNetSolutionTypePtrOutput) ToVNetSolutionTypePtrOutputWithContext(ctx context.Context) VNetSolutionTypePtrOutput {
	return o
}

func (o VNetSolutionTypePtrOutput) Elem() VNetSolutionTypeOutput {
	return o.ApplyT(func(v *VNetSolutionType) VNetSolutionType {
		if v != nil {
			return *v
		}
		var ret VNetSolutionType
		return ret
	}).(VNetSolutionTypeOutput)
}

func (o VNetSolutionTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o VNetSolutionTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *VNetSolutionType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// VNetSolutionTypeInput is an input type that accepts values of the VNetSolutionType enum
// A concrete instance of `VNetSolutionTypeInput` can be one of the following:
//
//	VNetSolutionTypeServiceEndpoint
//	VNetSolutionTypePrivateLink
type VNetSolutionTypeInput interface {
	pulumi.Input

	ToVNetSolutionTypeOutput() VNetSolutionTypeOutput
	ToVNetSolutionTypeOutputWithContext(context.Context) VNetSolutionTypeOutput
}

var vnetSolutionTypePtrType = reflect.TypeOf((**VNetSolutionType)(nil)).Elem()

type VNetSolutionTypePtrInput interface {
	pulumi.Input

	ToVNetSolutionTypePtrOutput() VNetSolutionTypePtrOutput
	ToVNetSolutionTypePtrOutputWithContext(context.Context) VNetSolutionTypePtrOutput
}

type vnetSolutionTypePtr string

func VNetSolutionTypePtr(v string) VNetSolutionTypePtrInput {
	return (*vnetSolutionTypePtr)(&v)
}

func (*vnetSolutionTypePtr) ElementType() reflect.Type {
	return vnetSolutionTypePtrType
}

func (in *vnetSolutionTypePtr) ToVNetSolutionTypePtrOutput() VNetSolutionTypePtrOutput {
	return pulumi.ToOutput(in).(VNetSolutionTypePtrOutput)
}

func (in *vnetSolutionTypePtr) ToVNetSolutionTypePtrOutputWithContext(ctx context.Context) VNetSolutionTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(VNetSolutionTypePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AccessKeyPermissionsOutput{})
	pulumi.RegisterOutputType(AccessKeyPermissionsPtrOutput{})
	pulumi.RegisterOutputType(ActionTypeOutput{})
	pulumi.RegisterOutputType(ActionTypePtrOutput{})
	pulumi.RegisterOutputType(AllowTypeOutput{})
	pulumi.RegisterOutputType(AllowTypePtrOutput{})
	pulumi.RegisterOutputType(ClientTypeOutput{})
	pulumi.RegisterOutputType(ClientTypePtrOutput{})
	pulumi.RegisterOutputType(DeleteOrUpdateBehaviorOutput{})
	pulumi.RegisterOutputType(DeleteOrUpdateBehaviorPtrOutput{})
	pulumi.RegisterOutputType(VNetSolutionTypeOutput{})
	pulumi.RegisterOutputType(VNetSolutionTypePtrOutput{})
}
