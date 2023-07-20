// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220430preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The permissions assigned to the shared access policy.
type AccessRights string

const (
	AccessRightsRegistryRead                                             = AccessRights("RegistryRead")
	AccessRightsRegistryWrite                                            = AccessRights("RegistryWrite")
	AccessRightsServiceConnect                                           = AccessRights("ServiceConnect")
	AccessRightsDeviceConnect                                            = AccessRights("DeviceConnect")
	AccessRights_RegistryRead_RegistryWrite                              = AccessRights("RegistryRead, RegistryWrite")
	AccessRights_RegistryRead_ServiceConnect                             = AccessRights("RegistryRead, ServiceConnect")
	AccessRights_RegistryRead_DeviceConnect                              = AccessRights("RegistryRead, DeviceConnect")
	AccessRights_RegistryWrite_ServiceConnect                            = AccessRights("RegistryWrite, ServiceConnect")
	AccessRights_RegistryWrite_DeviceConnect                             = AccessRights("RegistryWrite, DeviceConnect")
	AccessRights_ServiceConnect_DeviceConnect                            = AccessRights("ServiceConnect, DeviceConnect")
	AccessRights_RegistryRead_RegistryWrite_ServiceConnect               = AccessRights("RegistryRead, RegistryWrite, ServiceConnect")
	AccessRights_RegistryRead_RegistryWrite_DeviceConnect                = AccessRights("RegistryRead, RegistryWrite, DeviceConnect")
	AccessRights_RegistryRead_ServiceConnect_DeviceConnect               = AccessRights("RegistryRead, ServiceConnect, DeviceConnect")
	AccessRights_RegistryWrite_ServiceConnect_DeviceConnect              = AccessRights("RegistryWrite, ServiceConnect, DeviceConnect")
	AccessRights_RegistryRead_RegistryWrite_ServiceConnect_DeviceConnect = AccessRights("RegistryRead, RegistryWrite, ServiceConnect, DeviceConnect")
)

func (AccessRights) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessRights)(nil)).Elem()
}

func (e AccessRights) ToAccessRightsOutput() AccessRightsOutput {
	return pulumi.ToOutput(e).(AccessRightsOutput)
}

func (e AccessRights) ToAccessRightsOutputWithContext(ctx context.Context) AccessRightsOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AccessRightsOutput)
}

func (e AccessRights) ToAccessRightsPtrOutput() AccessRightsPtrOutput {
	return e.ToAccessRightsPtrOutputWithContext(context.Background())
}

func (e AccessRights) ToAccessRightsPtrOutputWithContext(ctx context.Context) AccessRightsPtrOutput {
	return AccessRights(e).ToAccessRightsOutputWithContext(ctx).ToAccessRightsPtrOutputWithContext(ctx)
}

func (e AccessRights) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AccessRights) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AccessRights) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AccessRights) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AccessRightsOutput struct{ *pulumi.OutputState }

func (AccessRightsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessRights)(nil)).Elem()
}

func (o AccessRightsOutput) ToAccessRightsOutput() AccessRightsOutput {
	return o
}

func (o AccessRightsOutput) ToAccessRightsOutputWithContext(ctx context.Context) AccessRightsOutput {
	return o
}

func (o AccessRightsOutput) ToAccessRightsPtrOutput() AccessRightsPtrOutput {
	return o.ToAccessRightsPtrOutputWithContext(context.Background())
}

func (o AccessRightsOutput) ToAccessRightsPtrOutputWithContext(ctx context.Context) AccessRightsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AccessRights) *AccessRights {
		return &v
	}).(AccessRightsPtrOutput)
}

func (o AccessRightsOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AccessRightsOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AccessRights) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AccessRightsOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AccessRightsOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AccessRights) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AccessRightsPtrOutput struct{ *pulumi.OutputState }

func (AccessRightsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessRights)(nil)).Elem()
}

func (o AccessRightsPtrOutput) ToAccessRightsPtrOutput() AccessRightsPtrOutput {
	return o
}

func (o AccessRightsPtrOutput) ToAccessRightsPtrOutputWithContext(ctx context.Context) AccessRightsPtrOutput {
	return o
}

func (o AccessRightsPtrOutput) Elem() AccessRightsOutput {
	return o.ApplyT(func(v *AccessRights) AccessRights {
		if v != nil {
			return *v
		}
		var ret AccessRights
		return ret
	}).(AccessRightsOutput)
}

func (o AccessRightsPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AccessRightsPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AccessRights) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// AccessRightsInput is an input type that accepts AccessRightsArgs and AccessRightsOutput values.
// You can construct a concrete instance of `AccessRightsInput` via:
//
//	AccessRightsArgs{...}
type AccessRightsInput interface {
	pulumi.Input

	ToAccessRightsOutput() AccessRightsOutput
	ToAccessRightsOutputWithContext(context.Context) AccessRightsOutput
}

var accessRightsPtrType = reflect.TypeOf((**AccessRights)(nil)).Elem()

type AccessRightsPtrInput interface {
	pulumi.Input

	ToAccessRightsPtrOutput() AccessRightsPtrOutput
	ToAccessRightsPtrOutputWithContext(context.Context) AccessRightsPtrOutput
}

type accessRightsPtr string

func AccessRightsPtr(v string) AccessRightsPtrInput {
	return (*accessRightsPtr)(&v)
}

func (*accessRightsPtr) ElementType() reflect.Type {
	return accessRightsPtrType
}

func (in *accessRightsPtr) ToAccessRightsPtrOutput() AccessRightsPtrOutput {
	return pulumi.ToOutput(in).(AccessRightsPtrOutput)
}

func (in *accessRightsPtr) ToAccessRightsPtrOutputWithContext(ctx context.Context) AccessRightsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AccessRightsPtrOutput)
}

// Specifies authentication type being used for connecting to the storage account.
type AuthenticationType string

const (
	AuthenticationTypeKeyBased      = AuthenticationType("keyBased")
	AuthenticationTypeIdentityBased = AuthenticationType("identityBased")
)

// The capabilities and features enabled for the IoT hub.
type Capabilities string

const (
	CapabilitiesNone             = Capabilities("None")
	CapabilitiesDeviceManagement = Capabilities("DeviceManagement")
)

// Default Action for Network Rule Set
type DefaultAction string

const (
	DefaultActionDeny  = DefaultAction("Deny")
	DefaultActionAllow = DefaultAction("Allow")
)

// The name of the SKU.
type IotHubSku string

const (
	IotHubSkuF1 = IotHubSku("F1")
	IotHubSkuS1 = IotHubSku("S1")
	IotHubSkuS2 = IotHubSku("S2")
	IotHubSkuS3 = IotHubSku("S3")
	IotHubSkuB1 = IotHubSku("B1")
	IotHubSkuB2 = IotHubSku("B2")
	IotHubSkuB3 = IotHubSku("B3")
)

// The desired action for requests captured by this rule.
type IpFilterActionType string

const (
	IpFilterActionTypeAccept = IpFilterActionType("Accept")
	IpFilterActionTypeReject = IpFilterActionType("Reject")
)

func (IpFilterActionType) ElementType() reflect.Type {
	return reflect.TypeOf((*IpFilterActionType)(nil)).Elem()
}

func (e IpFilterActionType) ToIpFilterActionTypeOutput() IpFilterActionTypeOutput {
	return pulumi.ToOutput(e).(IpFilterActionTypeOutput)
}

func (e IpFilterActionType) ToIpFilterActionTypeOutputWithContext(ctx context.Context) IpFilterActionTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(IpFilterActionTypeOutput)
}

func (e IpFilterActionType) ToIpFilterActionTypePtrOutput() IpFilterActionTypePtrOutput {
	return e.ToIpFilterActionTypePtrOutputWithContext(context.Background())
}

func (e IpFilterActionType) ToIpFilterActionTypePtrOutputWithContext(ctx context.Context) IpFilterActionTypePtrOutput {
	return IpFilterActionType(e).ToIpFilterActionTypeOutputWithContext(ctx).ToIpFilterActionTypePtrOutputWithContext(ctx)
}

func (e IpFilterActionType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e IpFilterActionType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e IpFilterActionType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e IpFilterActionType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type IpFilterActionTypeOutput struct{ *pulumi.OutputState }

func (IpFilterActionTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IpFilterActionType)(nil)).Elem()
}

func (o IpFilterActionTypeOutput) ToIpFilterActionTypeOutput() IpFilterActionTypeOutput {
	return o
}

func (o IpFilterActionTypeOutput) ToIpFilterActionTypeOutputWithContext(ctx context.Context) IpFilterActionTypeOutput {
	return o
}

func (o IpFilterActionTypeOutput) ToIpFilterActionTypePtrOutput() IpFilterActionTypePtrOutput {
	return o.ToIpFilterActionTypePtrOutputWithContext(context.Background())
}

func (o IpFilterActionTypeOutput) ToIpFilterActionTypePtrOutputWithContext(ctx context.Context) IpFilterActionTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IpFilterActionType) *IpFilterActionType {
		return &v
	}).(IpFilterActionTypePtrOutput)
}

func (o IpFilterActionTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o IpFilterActionTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IpFilterActionType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o IpFilterActionTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IpFilterActionTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IpFilterActionType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type IpFilterActionTypePtrOutput struct{ *pulumi.OutputState }

func (IpFilterActionTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IpFilterActionType)(nil)).Elem()
}

func (o IpFilterActionTypePtrOutput) ToIpFilterActionTypePtrOutput() IpFilterActionTypePtrOutput {
	return o
}

func (o IpFilterActionTypePtrOutput) ToIpFilterActionTypePtrOutputWithContext(ctx context.Context) IpFilterActionTypePtrOutput {
	return o
}

func (o IpFilterActionTypePtrOutput) Elem() IpFilterActionTypeOutput {
	return o.ApplyT(func(v *IpFilterActionType) IpFilterActionType {
		if v != nil {
			return *v
		}
		var ret IpFilterActionType
		return ret
	}).(IpFilterActionTypeOutput)
}

func (o IpFilterActionTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IpFilterActionTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *IpFilterActionType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// IpFilterActionTypeInput is an input type that accepts IpFilterActionTypeArgs and IpFilterActionTypeOutput values.
// You can construct a concrete instance of `IpFilterActionTypeInput` via:
//
//	IpFilterActionTypeArgs{...}
type IpFilterActionTypeInput interface {
	pulumi.Input

	ToIpFilterActionTypeOutput() IpFilterActionTypeOutput
	ToIpFilterActionTypeOutputWithContext(context.Context) IpFilterActionTypeOutput
}

var ipFilterActionTypePtrType = reflect.TypeOf((**IpFilterActionType)(nil)).Elem()

type IpFilterActionTypePtrInput interface {
	pulumi.Input

	ToIpFilterActionTypePtrOutput() IpFilterActionTypePtrOutput
	ToIpFilterActionTypePtrOutputWithContext(context.Context) IpFilterActionTypePtrOutput
}

type ipFilterActionTypePtr string

func IpFilterActionTypePtr(v string) IpFilterActionTypePtrInput {
	return (*ipFilterActionTypePtr)(&v)
}

func (*ipFilterActionTypePtr) ElementType() reflect.Type {
	return ipFilterActionTypePtrType
}

func (in *ipFilterActionTypePtr) ToIpFilterActionTypePtrOutput() IpFilterActionTypePtrOutput {
	return pulumi.ToOutput(in).(IpFilterActionTypePtrOutput)
}

func (in *ipFilterActionTypePtr) ToIpFilterActionTypePtrOutputWithContext(ctx context.Context) IpFilterActionTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(IpFilterActionTypePtrOutput)
}

// IP Filter Action
type NetworkRuleIPAction string

const (
	NetworkRuleIPActionAllow = NetworkRuleIPAction("Allow")
)

// The status of a private endpoint connection
type PrivateLinkServiceConnectionStatus string

const (
	PrivateLinkServiceConnectionStatusPending      = PrivateLinkServiceConnectionStatus("Pending")
	PrivateLinkServiceConnectionStatusApproved     = PrivateLinkServiceConnectionStatus("Approved")
	PrivateLinkServiceConnectionStatusRejected     = PrivateLinkServiceConnectionStatus("Rejected")
	PrivateLinkServiceConnectionStatusDisconnected = PrivateLinkServiceConnectionStatus("Disconnected")
)

// Whether requests from Public Network are allowed
type PublicNetworkAccess string

const (
	PublicNetworkAccessEnabled  = PublicNetworkAccess("Enabled")
	PublicNetworkAccessDisabled = PublicNetworkAccess("Disabled")
)

// The type of identity used for the resource. The type 'SystemAssigned,UserAssigned' includes both an implicitly created identity and a set of user assigned identities. The type 'None' will remove any identities from the service.
type ResourceIdentityType string

const (
	ResourceIdentityTypeSystemAssigned               = ResourceIdentityType("SystemAssigned")
	ResourceIdentityTypeUserAssigned                 = ResourceIdentityType("UserAssigned")
	ResourceIdentityType_SystemAssigned_UserAssigned = ResourceIdentityType("SystemAssigned, UserAssigned")
	ResourceIdentityTypeNone                         = ResourceIdentityType("None")
)

func (ResourceIdentityType) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdentityType)(nil)).Elem()
}

func (e ResourceIdentityType) ToResourceIdentityTypeOutput() ResourceIdentityTypeOutput {
	return pulumi.ToOutput(e).(ResourceIdentityTypeOutput)
}

func (e ResourceIdentityType) ToResourceIdentityTypeOutputWithContext(ctx context.Context) ResourceIdentityTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ResourceIdentityTypeOutput)
}

func (e ResourceIdentityType) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return e.ToResourceIdentityTypePtrOutputWithContext(context.Background())
}

func (e ResourceIdentityType) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return ResourceIdentityType(e).ToResourceIdentityTypeOutputWithContext(ctx).ToResourceIdentityTypePtrOutputWithContext(ctx)
}

func (e ResourceIdentityType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ResourceIdentityType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ResourceIdentityType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ResourceIdentityType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ResourceIdentityTypeOutput struct{ *pulumi.OutputState }

func (ResourceIdentityTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdentityType)(nil)).Elem()
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypeOutput() ResourceIdentityTypeOutput {
	return o
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypeOutputWithContext(ctx context.Context) ResourceIdentityTypeOutput {
	return o
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return o.ToResourceIdentityTypePtrOutputWithContext(context.Background())
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceIdentityType) *ResourceIdentityType {
		return &v
	}).(ResourceIdentityTypePtrOutput)
}

func (o ResourceIdentityTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ResourceIdentityTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ResourceIdentityType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ResourceIdentityTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ResourceIdentityTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ResourceIdentityType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ResourceIdentityTypePtrOutput struct{ *pulumi.OutputState }

func (ResourceIdentityTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceIdentityType)(nil)).Elem()
}

func (o ResourceIdentityTypePtrOutput) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return o
}

func (o ResourceIdentityTypePtrOutput) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return o
}

func (o ResourceIdentityTypePtrOutput) Elem() ResourceIdentityTypeOutput {
	return o.ApplyT(func(v *ResourceIdentityType) ResourceIdentityType {
		if v != nil {
			return *v
		}
		var ret ResourceIdentityType
		return ret
	}).(ResourceIdentityTypeOutput)
}

func (o ResourceIdentityTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ResourceIdentityTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ResourceIdentityType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// ResourceIdentityTypeInput is an input type that accepts ResourceIdentityTypeArgs and ResourceIdentityTypeOutput values.
// You can construct a concrete instance of `ResourceIdentityTypeInput` via:
//
//	ResourceIdentityTypeArgs{...}
type ResourceIdentityTypeInput interface {
	pulumi.Input

	ToResourceIdentityTypeOutput() ResourceIdentityTypeOutput
	ToResourceIdentityTypeOutputWithContext(context.Context) ResourceIdentityTypeOutput
}

var resourceIdentityTypePtrType = reflect.TypeOf((**ResourceIdentityType)(nil)).Elem()

type ResourceIdentityTypePtrInput interface {
	pulumi.Input

	ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput
	ToResourceIdentityTypePtrOutputWithContext(context.Context) ResourceIdentityTypePtrOutput
}

type resourceIdentityTypePtr string

func ResourceIdentityTypePtr(v string) ResourceIdentityTypePtrInput {
	return (*resourceIdentityTypePtr)(&v)
}

func (*resourceIdentityTypePtr) ElementType() reflect.Type {
	return resourceIdentityTypePtrType
}

func (in *resourceIdentityTypePtr) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return pulumi.ToOutput(in).(ResourceIdentityTypePtrOutput)
}

func (in *resourceIdentityTypePtr) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ResourceIdentityTypePtrOutput)
}

// The source that the routing rule is to be applied to, such as DeviceMessages.
type RoutingSource string

const (
	RoutingSourceInvalid                     = RoutingSource("Invalid")
	RoutingSourceDeviceMessages              = RoutingSource("DeviceMessages")
	RoutingSourceTwinChangeEvents            = RoutingSource("TwinChangeEvents")
	RoutingSourceDeviceLifecycleEvents       = RoutingSource("DeviceLifecycleEvents")
	RoutingSourceDeviceJobLifecycleEvents    = RoutingSource("DeviceJobLifecycleEvents")
	RoutingSourceDigitalTwinChangeEvents     = RoutingSource("DigitalTwinChangeEvents")
	RoutingSourceDeviceConnectionStateEvents = RoutingSource("DeviceConnectionStateEvents")
	RoutingSourceMqttBrokerMessages          = RoutingSource("MqttBrokerMessages")
)

func init() {
	pulumi.RegisterOutputType(AccessRightsOutput{})
	pulumi.RegisterOutputType(AccessRightsPtrOutput{})
	pulumi.RegisterOutputType(IpFilterActionTypeOutput{})
	pulumi.RegisterOutputType(IpFilterActionTypePtrOutput{})
	pulumi.RegisterOutputType(ResourceIdentityTypeOutput{})
	pulumi.RegisterOutputType(ResourceIdentityTypePtrOutput{})
}
