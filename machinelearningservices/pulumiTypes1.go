// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package machinelearningservices

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

var _ = utilities.GetEnvOrDefault

// Settings for user account that gets created on each on the nodes of a compute.
type UserAccountCredentials struct {
	// Name of the administrator user account which can be used to SSH to nodes.
	AdminUserName string `pulumi:"adminUserName"`
	// Password of the administrator user account.
	AdminUserPassword *string `pulumi:"adminUserPassword"`
	// SSH public key of the administrator user account.
	AdminUserSshPublicKey *string `pulumi:"adminUserSshPublicKey"`
}

// Settings for user account that gets created on each on the nodes of a compute.
type UserAccountCredentialsResponse struct {
	// Name of the administrator user account which can be used to SSH to nodes.
	AdminUserName string `pulumi:"adminUserName"`
	// Password of the administrator user account.
	AdminUserPassword *string `pulumi:"adminUserPassword"`
	// SSH public key of the administrator user account.
	AdminUserSshPublicKey *string `pulumi:"adminUserSshPublicKey"`
}

// User assigned identity properties
type UserAssignedIdentityResponse struct {
	// The client ID of the assigned identity.
	ClientId string `pulumi:"clientId"`
	// The principal ID of the assigned identity.
	PrincipalId string `pulumi:"principalId"`
	// The tenant ID of the user assigned identity.
	TenantId *string `pulumi:"tenantId"`
}

// User assigned identity properties
type UserAssignedIdentityResponseOutput struct{ *pulumi.OutputState }

func (UserAssignedIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserAssignedIdentityResponse)(nil)).Elem()
}

func (o UserAssignedIdentityResponseOutput) ToUserAssignedIdentityResponseOutput() UserAssignedIdentityResponseOutput {
	return o
}

func (o UserAssignedIdentityResponseOutput) ToUserAssignedIdentityResponseOutputWithContext(ctx context.Context) UserAssignedIdentityResponseOutput {
	return o
}

func (o UserAssignedIdentityResponseOutput) ToOutput(ctx context.Context) pulumix.Output[UserAssignedIdentityResponse] {
	return pulumix.Output[UserAssignedIdentityResponse]{
		OutputState: o.OutputState,
	}
}

// The client ID of the assigned identity.
func (o UserAssignedIdentityResponseOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v UserAssignedIdentityResponse) string { return v.ClientId }).(pulumi.StringOutput)
}

// The principal ID of the assigned identity.
func (o UserAssignedIdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v UserAssignedIdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

// The tenant ID of the user assigned identity.
func (o UserAssignedIdentityResponseOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserAssignedIdentityResponse) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

type UserAssignedIdentityResponseMapOutput struct{ *pulumi.OutputState }

func (UserAssignedIdentityResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]UserAssignedIdentityResponse)(nil)).Elem()
}

func (o UserAssignedIdentityResponseMapOutput) ToUserAssignedIdentityResponseMapOutput() UserAssignedIdentityResponseMapOutput {
	return o
}

func (o UserAssignedIdentityResponseMapOutput) ToUserAssignedIdentityResponseMapOutputWithContext(ctx context.Context) UserAssignedIdentityResponseMapOutput {
	return o
}

func (o UserAssignedIdentityResponseMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]UserAssignedIdentityResponse] {
	return pulumix.Output[map[string]UserAssignedIdentityResponse]{
		OutputState: o.OutputState,
	}
}

func (o UserAssignedIdentityResponseMapOutput) MapIndex(k pulumi.StringInput) UserAssignedIdentityResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) UserAssignedIdentityResponse {
		return vs[0].(map[string]UserAssignedIdentityResponse)[vs[1].(string)]
	}).(UserAssignedIdentityResponseOutput)
}

type UserCreatedAcrAccount struct {
	// ARM ResourceId of a resource
	ArmResourceId *ArmResourceId `pulumi:"armResourceId"`
}

// UserCreatedAcrAccountInput is an input type that accepts UserCreatedAcrAccountArgs and UserCreatedAcrAccountOutput values.
// You can construct a concrete instance of `UserCreatedAcrAccountInput` via:
//
//	UserCreatedAcrAccountArgs{...}
type UserCreatedAcrAccountInput interface {
	pulumi.Input

	ToUserCreatedAcrAccountOutput() UserCreatedAcrAccountOutput
	ToUserCreatedAcrAccountOutputWithContext(context.Context) UserCreatedAcrAccountOutput
}

type UserCreatedAcrAccountArgs struct {
	// ARM ResourceId of a resource
	ArmResourceId ArmResourceIdPtrInput `pulumi:"armResourceId"`
}

func (UserCreatedAcrAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UserCreatedAcrAccount)(nil)).Elem()
}

func (i UserCreatedAcrAccountArgs) ToUserCreatedAcrAccountOutput() UserCreatedAcrAccountOutput {
	return i.ToUserCreatedAcrAccountOutputWithContext(context.Background())
}

func (i UserCreatedAcrAccountArgs) ToUserCreatedAcrAccountOutputWithContext(ctx context.Context) UserCreatedAcrAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserCreatedAcrAccountOutput)
}

func (i UserCreatedAcrAccountArgs) ToOutput(ctx context.Context) pulumix.Output[UserCreatedAcrAccount] {
	return pulumix.Output[UserCreatedAcrAccount]{
		OutputState: i.ToUserCreatedAcrAccountOutputWithContext(ctx).OutputState,
	}
}

func (i UserCreatedAcrAccountArgs) ToUserCreatedAcrAccountPtrOutput() UserCreatedAcrAccountPtrOutput {
	return i.ToUserCreatedAcrAccountPtrOutputWithContext(context.Background())
}

func (i UserCreatedAcrAccountArgs) ToUserCreatedAcrAccountPtrOutputWithContext(ctx context.Context) UserCreatedAcrAccountPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserCreatedAcrAccountOutput).ToUserCreatedAcrAccountPtrOutputWithContext(ctx)
}

// UserCreatedAcrAccountPtrInput is an input type that accepts UserCreatedAcrAccountArgs, UserCreatedAcrAccountPtr and UserCreatedAcrAccountPtrOutput values.
// You can construct a concrete instance of `UserCreatedAcrAccountPtrInput` via:
//
//	        UserCreatedAcrAccountArgs{...}
//
//	or:
//
//	        nil
type UserCreatedAcrAccountPtrInput interface {
	pulumi.Input

	ToUserCreatedAcrAccountPtrOutput() UserCreatedAcrAccountPtrOutput
	ToUserCreatedAcrAccountPtrOutputWithContext(context.Context) UserCreatedAcrAccountPtrOutput
}

type userCreatedAcrAccountPtrType UserCreatedAcrAccountArgs

func UserCreatedAcrAccountPtr(v *UserCreatedAcrAccountArgs) UserCreatedAcrAccountPtrInput {
	return (*userCreatedAcrAccountPtrType)(v)
}

func (*userCreatedAcrAccountPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**UserCreatedAcrAccount)(nil)).Elem()
}

func (i *userCreatedAcrAccountPtrType) ToUserCreatedAcrAccountPtrOutput() UserCreatedAcrAccountPtrOutput {
	return i.ToUserCreatedAcrAccountPtrOutputWithContext(context.Background())
}

func (i *userCreatedAcrAccountPtrType) ToUserCreatedAcrAccountPtrOutputWithContext(ctx context.Context) UserCreatedAcrAccountPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserCreatedAcrAccountPtrOutput)
}

func (i *userCreatedAcrAccountPtrType) ToOutput(ctx context.Context) pulumix.Output[*UserCreatedAcrAccount] {
	return pulumix.Output[*UserCreatedAcrAccount]{
		OutputState: i.ToUserCreatedAcrAccountPtrOutputWithContext(ctx).OutputState,
	}
}

type UserCreatedAcrAccountOutput struct{ *pulumi.OutputState }

func (UserCreatedAcrAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserCreatedAcrAccount)(nil)).Elem()
}

func (o UserCreatedAcrAccountOutput) ToUserCreatedAcrAccountOutput() UserCreatedAcrAccountOutput {
	return o
}

func (o UserCreatedAcrAccountOutput) ToUserCreatedAcrAccountOutputWithContext(ctx context.Context) UserCreatedAcrAccountOutput {
	return o
}

func (o UserCreatedAcrAccountOutput) ToUserCreatedAcrAccountPtrOutput() UserCreatedAcrAccountPtrOutput {
	return o.ToUserCreatedAcrAccountPtrOutputWithContext(context.Background())
}

func (o UserCreatedAcrAccountOutput) ToUserCreatedAcrAccountPtrOutputWithContext(ctx context.Context) UserCreatedAcrAccountPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v UserCreatedAcrAccount) *UserCreatedAcrAccount {
		return &v
	}).(UserCreatedAcrAccountPtrOutput)
}

func (o UserCreatedAcrAccountOutput) ToOutput(ctx context.Context) pulumix.Output[UserCreatedAcrAccount] {
	return pulumix.Output[UserCreatedAcrAccount]{
		OutputState: o.OutputState,
	}
}

// ARM ResourceId of a resource
func (o UserCreatedAcrAccountOutput) ArmResourceId() ArmResourceIdPtrOutput {
	return o.ApplyT(func(v UserCreatedAcrAccount) *ArmResourceId { return v.ArmResourceId }).(ArmResourceIdPtrOutput)
}

type UserCreatedAcrAccountPtrOutput struct{ *pulumi.OutputState }

func (UserCreatedAcrAccountPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserCreatedAcrAccount)(nil)).Elem()
}

func (o UserCreatedAcrAccountPtrOutput) ToUserCreatedAcrAccountPtrOutput() UserCreatedAcrAccountPtrOutput {
	return o
}

func (o UserCreatedAcrAccountPtrOutput) ToUserCreatedAcrAccountPtrOutputWithContext(ctx context.Context) UserCreatedAcrAccountPtrOutput {
	return o
}

func (o UserCreatedAcrAccountPtrOutput) ToOutput(ctx context.Context) pulumix.Output[*UserCreatedAcrAccount] {
	return pulumix.Output[*UserCreatedAcrAccount]{
		OutputState: o.OutputState,
	}
}

func (o UserCreatedAcrAccountPtrOutput) Elem() UserCreatedAcrAccountOutput {
	return o.ApplyT(func(v *UserCreatedAcrAccount) UserCreatedAcrAccount {
		if v != nil {
			return *v
		}
		var ret UserCreatedAcrAccount
		return ret
	}).(UserCreatedAcrAccountOutput)
}

// ARM ResourceId of a resource
func (o UserCreatedAcrAccountPtrOutput) ArmResourceId() ArmResourceIdPtrOutput {
	return o.ApplyT(func(v *UserCreatedAcrAccount) *ArmResourceId {
		if v == nil {
			return nil
		}
		return v.ArmResourceId
	}).(ArmResourceIdPtrOutput)
}

type UserCreatedAcrAccountResponse struct {
	// ARM ResourceId of a resource
	ArmResourceId *ArmResourceIdResponse `pulumi:"armResourceId"`
}

type UserCreatedAcrAccountResponseOutput struct{ *pulumi.OutputState }

func (UserCreatedAcrAccountResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserCreatedAcrAccountResponse)(nil)).Elem()
}

func (o UserCreatedAcrAccountResponseOutput) ToUserCreatedAcrAccountResponseOutput() UserCreatedAcrAccountResponseOutput {
	return o
}

func (o UserCreatedAcrAccountResponseOutput) ToUserCreatedAcrAccountResponseOutputWithContext(ctx context.Context) UserCreatedAcrAccountResponseOutput {
	return o
}

func (o UserCreatedAcrAccountResponseOutput) ToOutput(ctx context.Context) pulumix.Output[UserCreatedAcrAccountResponse] {
	return pulumix.Output[UserCreatedAcrAccountResponse]{
		OutputState: o.OutputState,
	}
}

// ARM ResourceId of a resource
func (o UserCreatedAcrAccountResponseOutput) ArmResourceId() ArmResourceIdResponsePtrOutput {
	return o.ApplyT(func(v UserCreatedAcrAccountResponse) *ArmResourceIdResponse { return v.ArmResourceId }).(ArmResourceIdResponsePtrOutput)
}

type UserCreatedAcrAccountResponsePtrOutput struct{ *pulumi.OutputState }

func (UserCreatedAcrAccountResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserCreatedAcrAccountResponse)(nil)).Elem()
}

func (o UserCreatedAcrAccountResponsePtrOutput) ToUserCreatedAcrAccountResponsePtrOutput() UserCreatedAcrAccountResponsePtrOutput {
	return o
}

func (o UserCreatedAcrAccountResponsePtrOutput) ToUserCreatedAcrAccountResponsePtrOutputWithContext(ctx context.Context) UserCreatedAcrAccountResponsePtrOutput {
	return o
}

func (o UserCreatedAcrAccountResponsePtrOutput) ToOutput(ctx context.Context) pulumix.Output[*UserCreatedAcrAccountResponse] {
	return pulumix.Output[*UserCreatedAcrAccountResponse]{
		OutputState: o.OutputState,
	}
}

func (o UserCreatedAcrAccountResponsePtrOutput) Elem() UserCreatedAcrAccountResponseOutput {
	return o.ApplyT(func(v *UserCreatedAcrAccountResponse) UserCreatedAcrAccountResponse {
		if v != nil {
			return *v
		}
		var ret UserCreatedAcrAccountResponse
		return ret
	}).(UserCreatedAcrAccountResponseOutput)
}

// ARM ResourceId of a resource
func (o UserCreatedAcrAccountResponsePtrOutput) ArmResourceId() ArmResourceIdResponsePtrOutput {
	return o.ApplyT(func(v *UserCreatedAcrAccountResponse) *ArmResourceIdResponse {
		if v == nil {
			return nil
		}
		return v.ArmResourceId
	}).(ArmResourceIdResponsePtrOutput)
}

type UserCreatedStorageAccount struct {
	// ARM ResourceId of a resource
	ArmResourceId *ArmResourceId `pulumi:"armResourceId"`
}

// UserCreatedStorageAccountInput is an input type that accepts UserCreatedStorageAccountArgs and UserCreatedStorageAccountOutput values.
// You can construct a concrete instance of `UserCreatedStorageAccountInput` via:
//
//	UserCreatedStorageAccountArgs{...}
type UserCreatedStorageAccountInput interface {
	pulumi.Input

	ToUserCreatedStorageAccountOutput() UserCreatedStorageAccountOutput
	ToUserCreatedStorageAccountOutputWithContext(context.Context) UserCreatedStorageAccountOutput
}

type UserCreatedStorageAccountArgs struct {
	// ARM ResourceId of a resource
	ArmResourceId ArmResourceIdPtrInput `pulumi:"armResourceId"`
}

func (UserCreatedStorageAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UserCreatedStorageAccount)(nil)).Elem()
}

func (i UserCreatedStorageAccountArgs) ToUserCreatedStorageAccountOutput() UserCreatedStorageAccountOutput {
	return i.ToUserCreatedStorageAccountOutputWithContext(context.Background())
}

func (i UserCreatedStorageAccountArgs) ToUserCreatedStorageAccountOutputWithContext(ctx context.Context) UserCreatedStorageAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserCreatedStorageAccountOutput)
}

func (i UserCreatedStorageAccountArgs) ToOutput(ctx context.Context) pulumix.Output[UserCreatedStorageAccount] {
	return pulumix.Output[UserCreatedStorageAccount]{
		OutputState: i.ToUserCreatedStorageAccountOutputWithContext(ctx).OutputState,
	}
}

func (i UserCreatedStorageAccountArgs) ToUserCreatedStorageAccountPtrOutput() UserCreatedStorageAccountPtrOutput {
	return i.ToUserCreatedStorageAccountPtrOutputWithContext(context.Background())
}

func (i UserCreatedStorageAccountArgs) ToUserCreatedStorageAccountPtrOutputWithContext(ctx context.Context) UserCreatedStorageAccountPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserCreatedStorageAccountOutput).ToUserCreatedStorageAccountPtrOutputWithContext(ctx)
}

// UserCreatedStorageAccountPtrInput is an input type that accepts UserCreatedStorageAccountArgs, UserCreatedStorageAccountPtr and UserCreatedStorageAccountPtrOutput values.
// You can construct a concrete instance of `UserCreatedStorageAccountPtrInput` via:
//
//	        UserCreatedStorageAccountArgs{...}
//
//	or:
//
//	        nil
type UserCreatedStorageAccountPtrInput interface {
	pulumi.Input

	ToUserCreatedStorageAccountPtrOutput() UserCreatedStorageAccountPtrOutput
	ToUserCreatedStorageAccountPtrOutputWithContext(context.Context) UserCreatedStorageAccountPtrOutput
}

type userCreatedStorageAccountPtrType UserCreatedStorageAccountArgs

func UserCreatedStorageAccountPtr(v *UserCreatedStorageAccountArgs) UserCreatedStorageAccountPtrInput {
	return (*userCreatedStorageAccountPtrType)(v)
}

func (*userCreatedStorageAccountPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**UserCreatedStorageAccount)(nil)).Elem()
}

func (i *userCreatedStorageAccountPtrType) ToUserCreatedStorageAccountPtrOutput() UserCreatedStorageAccountPtrOutput {
	return i.ToUserCreatedStorageAccountPtrOutputWithContext(context.Background())
}

func (i *userCreatedStorageAccountPtrType) ToUserCreatedStorageAccountPtrOutputWithContext(ctx context.Context) UserCreatedStorageAccountPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserCreatedStorageAccountPtrOutput)
}

func (i *userCreatedStorageAccountPtrType) ToOutput(ctx context.Context) pulumix.Output[*UserCreatedStorageAccount] {
	return pulumix.Output[*UserCreatedStorageAccount]{
		OutputState: i.ToUserCreatedStorageAccountPtrOutputWithContext(ctx).OutputState,
	}
}

type UserCreatedStorageAccountOutput struct{ *pulumi.OutputState }

func (UserCreatedStorageAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserCreatedStorageAccount)(nil)).Elem()
}

func (o UserCreatedStorageAccountOutput) ToUserCreatedStorageAccountOutput() UserCreatedStorageAccountOutput {
	return o
}

func (o UserCreatedStorageAccountOutput) ToUserCreatedStorageAccountOutputWithContext(ctx context.Context) UserCreatedStorageAccountOutput {
	return o
}

func (o UserCreatedStorageAccountOutput) ToUserCreatedStorageAccountPtrOutput() UserCreatedStorageAccountPtrOutput {
	return o.ToUserCreatedStorageAccountPtrOutputWithContext(context.Background())
}

func (o UserCreatedStorageAccountOutput) ToUserCreatedStorageAccountPtrOutputWithContext(ctx context.Context) UserCreatedStorageAccountPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v UserCreatedStorageAccount) *UserCreatedStorageAccount {
		return &v
	}).(UserCreatedStorageAccountPtrOutput)
}

func (o UserCreatedStorageAccountOutput) ToOutput(ctx context.Context) pulumix.Output[UserCreatedStorageAccount] {
	return pulumix.Output[UserCreatedStorageAccount]{
		OutputState: o.OutputState,
	}
}

// ARM ResourceId of a resource
func (o UserCreatedStorageAccountOutput) ArmResourceId() ArmResourceIdPtrOutput {
	return o.ApplyT(func(v UserCreatedStorageAccount) *ArmResourceId { return v.ArmResourceId }).(ArmResourceIdPtrOutput)
}

type UserCreatedStorageAccountPtrOutput struct{ *pulumi.OutputState }

func (UserCreatedStorageAccountPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserCreatedStorageAccount)(nil)).Elem()
}

func (o UserCreatedStorageAccountPtrOutput) ToUserCreatedStorageAccountPtrOutput() UserCreatedStorageAccountPtrOutput {
	return o
}

func (o UserCreatedStorageAccountPtrOutput) ToUserCreatedStorageAccountPtrOutputWithContext(ctx context.Context) UserCreatedStorageAccountPtrOutput {
	return o
}

func (o UserCreatedStorageAccountPtrOutput) ToOutput(ctx context.Context) pulumix.Output[*UserCreatedStorageAccount] {
	return pulumix.Output[*UserCreatedStorageAccount]{
		OutputState: o.OutputState,
	}
}

func (o UserCreatedStorageAccountPtrOutput) Elem() UserCreatedStorageAccountOutput {
	return o.ApplyT(func(v *UserCreatedStorageAccount) UserCreatedStorageAccount {
		if v != nil {
			return *v
		}
		var ret UserCreatedStorageAccount
		return ret
	}).(UserCreatedStorageAccountOutput)
}

// ARM ResourceId of a resource
func (o UserCreatedStorageAccountPtrOutput) ArmResourceId() ArmResourceIdPtrOutput {
	return o.ApplyT(func(v *UserCreatedStorageAccount) *ArmResourceId {
		if v == nil {
			return nil
		}
		return v.ArmResourceId
	}).(ArmResourceIdPtrOutput)
}

type UserCreatedStorageAccountResponse struct {
	// ARM ResourceId of a resource
	ArmResourceId *ArmResourceIdResponse `pulumi:"armResourceId"`
}

type UserCreatedStorageAccountResponseOutput struct{ *pulumi.OutputState }

func (UserCreatedStorageAccountResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserCreatedStorageAccountResponse)(nil)).Elem()
}

func (o UserCreatedStorageAccountResponseOutput) ToUserCreatedStorageAccountResponseOutput() UserCreatedStorageAccountResponseOutput {
	return o
}

func (o UserCreatedStorageAccountResponseOutput) ToUserCreatedStorageAccountResponseOutputWithContext(ctx context.Context) UserCreatedStorageAccountResponseOutput {
	return o
}

func (o UserCreatedStorageAccountResponseOutput) ToOutput(ctx context.Context) pulumix.Output[UserCreatedStorageAccountResponse] {
	return pulumix.Output[UserCreatedStorageAccountResponse]{
		OutputState: o.OutputState,
	}
}

// ARM ResourceId of a resource
func (o UserCreatedStorageAccountResponseOutput) ArmResourceId() ArmResourceIdResponsePtrOutput {
	return o.ApplyT(func(v UserCreatedStorageAccountResponse) *ArmResourceIdResponse { return v.ArmResourceId }).(ArmResourceIdResponsePtrOutput)
}

type UserCreatedStorageAccountResponsePtrOutput struct{ *pulumi.OutputState }

func (UserCreatedStorageAccountResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserCreatedStorageAccountResponse)(nil)).Elem()
}

func (o UserCreatedStorageAccountResponsePtrOutput) ToUserCreatedStorageAccountResponsePtrOutput() UserCreatedStorageAccountResponsePtrOutput {
	return o
}

func (o UserCreatedStorageAccountResponsePtrOutput) ToUserCreatedStorageAccountResponsePtrOutputWithContext(ctx context.Context) UserCreatedStorageAccountResponsePtrOutput {
	return o
}

func (o UserCreatedStorageAccountResponsePtrOutput) ToOutput(ctx context.Context) pulumix.Output[*UserCreatedStorageAccountResponse] {
	return pulumix.Output[*UserCreatedStorageAccountResponse]{
		OutputState: o.OutputState,
	}
}

func (o UserCreatedStorageAccountResponsePtrOutput) Elem() UserCreatedStorageAccountResponseOutput {
	return o.ApplyT(func(v *UserCreatedStorageAccountResponse) UserCreatedStorageAccountResponse {
		if v != nil {
			return *v
		}
		var ret UserCreatedStorageAccountResponse
		return ret
	}).(UserCreatedStorageAccountResponseOutput)
}

// ARM ResourceId of a resource
func (o UserCreatedStorageAccountResponsePtrOutput) ArmResourceId() ArmResourceIdResponsePtrOutput {
	return o.ApplyT(func(v *UserCreatedStorageAccountResponse) *ArmResourceIdResponse {
		if v == nil {
			return nil
		}
		return v.ArmResourceId
	}).(ArmResourceIdResponsePtrOutput)
}

// User identity configuration.
type UserIdentity struct {
	// Enum to determine identity framework.
	// Expected value is 'UserIdentity'.
	IdentityType string `pulumi:"identityType"`
}

// User identity configuration.
type UserIdentityResponse struct {
	// Enum to determine identity framework.
	// Expected value is 'UserIdentity'.
	IdentityType string `pulumi:"identityType"`
}

// User who created.
type UserInfoResponse struct {
	// A user alternate sec id. This represents the user in a different identity provider system Eg.1:live.com:puid
	UserAltSecId *string `pulumi:"userAltSecId"`
	// A user identity provider. Eg live.com
	UserIdp *string `pulumi:"userIdp"`
	// The issuer which issued the token for this user.
	UserIss *string `pulumi:"userIss"`
	//  A user's full name or a service principal's app ID.
	UserName *string `pulumi:"userName"`
	// A user or service principal's object ID..
	UserObjectId *string `pulumi:"userObjectId"`
	// A user or service principal's PuID.
	UserPuId *string `pulumi:"userPuId"`
	// A user or service principal's tenant ID.
	UserTenantId *string `pulumi:"userTenantId"`
}

// User who created.
type UserInfoResponseOutput struct{ *pulumi.OutputState }

func (UserInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserInfoResponse)(nil)).Elem()
}

func (o UserInfoResponseOutput) ToUserInfoResponseOutput() UserInfoResponseOutput {
	return o
}

func (o UserInfoResponseOutput) ToUserInfoResponseOutputWithContext(ctx context.Context) UserInfoResponseOutput {
	return o
}

func (o UserInfoResponseOutput) ToOutput(ctx context.Context) pulumix.Output[UserInfoResponse] {
	return pulumix.Output[UserInfoResponse]{
		OutputState: o.OutputState,
	}
}

// A user alternate sec id. This represents the user in a different identity provider system Eg.1:live.com:puid
func (o UserInfoResponseOutput) UserAltSecId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserInfoResponse) *string { return v.UserAltSecId }).(pulumi.StringPtrOutput)
}

// A user identity provider. Eg live.com
func (o UserInfoResponseOutput) UserIdp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserInfoResponse) *string { return v.UserIdp }).(pulumi.StringPtrOutput)
}

// The issuer which issued the token for this user.
func (o UserInfoResponseOutput) UserIss() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserInfoResponse) *string { return v.UserIss }).(pulumi.StringPtrOutput)
}

// A user's full name or a service principal's app ID.
func (o UserInfoResponseOutput) UserName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserInfoResponse) *string { return v.UserName }).(pulumi.StringPtrOutput)
}

// A user or service principal's object ID..
func (o UserInfoResponseOutput) UserObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserInfoResponse) *string { return v.UserObjectId }).(pulumi.StringPtrOutput)
}

// A user or service principal's PuID.
func (o UserInfoResponseOutput) UserPuId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserInfoResponse) *string { return v.UserPuId }).(pulumi.StringPtrOutput)
}

// A user or service principal's tenant ID.
func (o UserInfoResponseOutput) UserTenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserInfoResponse) *string { return v.UserTenantId }).(pulumi.StringPtrOutput)
}

type UserInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (UserInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserInfoResponse)(nil)).Elem()
}

func (o UserInfoResponsePtrOutput) ToUserInfoResponsePtrOutput() UserInfoResponsePtrOutput {
	return o
}

func (o UserInfoResponsePtrOutput) ToUserInfoResponsePtrOutputWithContext(ctx context.Context) UserInfoResponsePtrOutput {
	return o
}

func (o UserInfoResponsePtrOutput) ToOutput(ctx context.Context) pulumix.Output[*UserInfoResponse] {
	return pulumix.Output[*UserInfoResponse]{
		OutputState: o.OutputState,
	}
}

func (o UserInfoResponsePtrOutput) Elem() UserInfoResponseOutput {
	return o.ApplyT(func(v *UserInfoResponse) UserInfoResponse {
		if v != nil {
			return *v
		}
		var ret UserInfoResponse
		return ret
	}).(UserInfoResponseOutput)
}

// A user alternate sec id. This represents the user in a different identity provider system Eg.1:live.com:puid
func (o UserInfoResponsePtrOutput) UserAltSecId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.UserAltSecId
	}).(pulumi.StringPtrOutput)
}

// A user identity provider. Eg live.com
func (o UserInfoResponsePtrOutput) UserIdp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.UserIdp
	}).(pulumi.StringPtrOutput)
}

// The issuer which issued the token for this user.
func (o UserInfoResponsePtrOutput) UserIss() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.UserIss
	}).(pulumi.StringPtrOutput)
}

// A user's full name or a service principal's app ID.
func (o UserInfoResponsePtrOutput) UserName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.UserName
	}).(pulumi.StringPtrOutput)
}

// A user or service principal's object ID..
func (o UserInfoResponsePtrOutput) UserObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.UserObjectId
	}).(pulumi.StringPtrOutput)
}

// A user or service principal's PuID.
func (o UserInfoResponsePtrOutput) UserPuId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.UserPuId
	}).(pulumi.StringPtrOutput)
}

// A user or service principal's tenant ID.
func (o UserInfoResponsePtrOutput) UserTenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.UserTenantId
	}).(pulumi.StringPtrOutput)
}

type UsernamePasswordAuthTypeWorkspaceConnectionProperties struct {
	// Authentication type of the connection target
	// Expected value is 'UsernamePassword'.
	AuthType string `pulumi:"authType"`
	// Category of the connection
	Category    *string                              `pulumi:"category"`
	Credentials *WorkspaceConnectionUsernamePassword `pulumi:"credentials"`
	Target      *string                              `pulumi:"target"`
	// Value details of the workspace connection.
	Value *string `pulumi:"value"`
	// format for the workspace connection value
	ValueFormat *string `pulumi:"valueFormat"`
}

type UsernamePasswordAuthTypeWorkspaceConnectionPropertiesResponse struct {
	// Authentication type of the connection target
	// Expected value is 'UsernamePassword'.
	AuthType string `pulumi:"authType"`
	// Category of the connection
	Category    *string                                      `pulumi:"category"`
	Credentials *WorkspaceConnectionUsernamePasswordResponse `pulumi:"credentials"`
	ExpiryTime  *string                                      `pulumi:"expiryTime"`
	Metadata    interface{}                                  `pulumi:"metadata"`
	Target      *string                                      `pulumi:"target"`
	// Value details of the workspace connection.
	Value *string `pulumi:"value"`
	// format for the workspace connection value
	ValueFormat *string `pulumi:"valueFormat"`
}

// A Machine Learning compute based on Azure Virtual Machines.
type VirtualMachine struct {
	// Location for the underlying compute
	ComputeLocation *string `pulumi:"computeLocation"`
	// The type of compute
	// Expected value is 'VirtualMachine'.
	ComputeType string `pulumi:"computeType"`
	// The description of the Machine Learning compute.
	Description *string `pulumi:"description"`
	// Opt-out of local authentication and ensure customers can use only MSI and AAD exclusively for authentication.
	DisableLocalAuth *bool                           `pulumi:"disableLocalAuth"`
	Properties       *VirtualMachineSchemaProperties `pulumi:"properties"`
	// ARM resource id of the underlying compute
	ResourceId *string `pulumi:"resourceId"`
}

// Virtual Machine image for Windows AML Compute
type VirtualMachineImage struct {
	// Virtual Machine image path
	Id string `pulumi:"id"`
}

// Virtual Machine image for Windows AML Compute
type VirtualMachineImageResponse struct {
	// Virtual Machine image path
	Id string `pulumi:"id"`
}

// A Machine Learning compute based on Azure Virtual Machines.
type VirtualMachineResponse struct {
	// Location for the underlying compute
	ComputeLocation *string `pulumi:"computeLocation"`
	// The type of compute
	// Expected value is 'VirtualMachine'.
	ComputeType string `pulumi:"computeType"`
	// The time at which the compute was created.
	CreatedOn string `pulumi:"createdOn"`
	// The description of the Machine Learning compute.
	Description *string `pulumi:"description"`
	// Opt-out of local authentication and ensure customers can use only MSI and AAD exclusively for authentication.
	DisableLocalAuth *bool `pulumi:"disableLocalAuth"`
	// Indicating whether the compute was provisioned by user and brought from outside if true, or machine learning service provisioned it if false.
	IsAttachedCompute bool `pulumi:"isAttachedCompute"`
	// The time at which the compute was last modified.
	ModifiedOn string                                  `pulumi:"modifiedOn"`
	Properties *VirtualMachineSchemaResponseProperties `pulumi:"properties"`
	// Errors during provisioning
	ProvisioningErrors []ErrorResponseResponse `pulumi:"provisioningErrors"`
	// The provision state of the cluster. Valid values are Unknown, Updating, Provisioning, Succeeded, and Failed.
	ProvisioningState string `pulumi:"provisioningState"`
	// ARM resource id of the underlying compute
	ResourceId *string `pulumi:"resourceId"`
}

type VirtualMachineSchemaProperties struct {
	// Public IP address of the virtual machine.
	Address *string `pulumi:"address"`
	// Admin credentials for virtual machine
	AdministratorAccount *VirtualMachineSshCredentials `pulumi:"administratorAccount"`
	// Indicates whether this compute will be used for running notebooks.
	IsNotebookInstanceCompute *bool `pulumi:"isNotebookInstanceCompute"`
	// Notebook server port open for ssh connections.
	NotebookServerPort *int `pulumi:"notebookServerPort"`
	// Port open for ssh connections.
	SshPort *int `pulumi:"sshPort"`
	// Virtual Machine size
	VirtualMachineSize *string `pulumi:"virtualMachineSize"`
}

type VirtualMachineSchemaResponseProperties struct {
	// Public IP address of the virtual machine.
	Address *string `pulumi:"address"`
	// Admin credentials for virtual machine
	AdministratorAccount *VirtualMachineSshCredentialsResponse `pulumi:"administratorAccount"`
	// Indicates whether this compute will be used for running notebooks.
	IsNotebookInstanceCompute *bool `pulumi:"isNotebookInstanceCompute"`
	// Notebook server port open for ssh connections.
	NotebookServerPort *int `pulumi:"notebookServerPort"`
	// Port open for ssh connections.
	SshPort *int `pulumi:"sshPort"`
	// Virtual Machine size
	VirtualMachineSize *string `pulumi:"virtualMachineSize"`
}

// Admin credentials for virtual machine
type VirtualMachineSshCredentials struct {
	// Password of admin account
	Password *string `pulumi:"password"`
	// Private key data
	PrivateKeyData *string `pulumi:"privateKeyData"`
	// Public key data
	PublicKeyData *string `pulumi:"publicKeyData"`
	// Username of admin account
	Username *string `pulumi:"username"`
}

// Admin credentials for virtual machine
type VirtualMachineSshCredentialsResponse struct {
	// Password of admin account
	Password *string `pulumi:"password"`
	// Private key data
	PrivateKeyData *string `pulumi:"privateKeyData"`
	// Public key data
	PublicKeyData *string `pulumi:"publicKeyData"`
	// Username of admin account
	Username *string `pulumi:"username"`
}

// Describes the volume configuration for the container
type VolumeDefinition struct {
	// Bind Options of the mount
	Bind *BindOptions `pulumi:"bind"`
	// Consistency of the volume
	Consistency *string `pulumi:"consistency"`
	// Indicate whether to mount volume as readOnly. Default value for this is false.
	ReadOnly *bool `pulumi:"readOnly"`
	// Source of the mount. For bind mounts this is the host path.
	Source *string `pulumi:"source"`
	// Target of the mount. For bind mounts this is the path in the container.
	Target *string `pulumi:"target"`
	// tmpfs option of the mount
	Tmpfs *TmpfsOptions `pulumi:"tmpfs"`
	// Type of Volume Definition. Possible Values: bind,volume,tmpfs,npipe
	Type *string `pulumi:"type"`
	// Volume Options of the mount
	Volume *VolumeOptions `pulumi:"volume"`
}

// Defaults sets the appropriate defaults for VolumeDefinition
func (val *VolumeDefinition) Defaults() *VolumeDefinition {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.Type == nil {
		type_ := "bind"
		tmp.Type = &type_
	}
	return &tmp
}

// Describes the volume configuration for the container
type VolumeDefinitionResponse struct {
	// Bind Options of the mount
	Bind *BindOptionsResponse `pulumi:"bind"`
	// Consistency of the volume
	Consistency *string `pulumi:"consistency"`
	// Indicate whether to mount volume as readOnly. Default value for this is false.
	ReadOnly *bool `pulumi:"readOnly"`
	// Source of the mount. For bind mounts this is the host path.
	Source *string `pulumi:"source"`
	// Target of the mount. For bind mounts this is the path in the container.
	Target *string `pulumi:"target"`
	// tmpfs option of the mount
	Tmpfs *TmpfsOptionsResponse `pulumi:"tmpfs"`
	// Type of Volume Definition. Possible Values: bind,volume,tmpfs,npipe
	Type *string `pulumi:"type"`
	// Volume Options of the mount
	Volume *VolumeOptionsResponse `pulumi:"volume"`
}

// Defaults sets the appropriate defaults for VolumeDefinitionResponse
func (val *VolumeDefinitionResponse) Defaults() *VolumeDefinitionResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.Type == nil {
		type_ := "bind"
		tmp.Type = &type_
	}
	return &tmp
}

// Describes the volume options for the container
type VolumeOptions struct {
	// Indicate whether volume is nocopy
	Nocopy *bool `pulumi:"nocopy"`
}

// Describes the volume options for the container
type VolumeOptionsResponse struct {
	// Indicate whether volume is nocopy
	Nocopy *bool `pulumi:"nocopy"`
}

type WorkspaceConnectionAccessKeyResponse struct {
	AccessKeyId     *string `pulumi:"accessKeyId"`
	SecretAccessKey *string `pulumi:"secretAccessKey"`
}

// Api key object for workspace connection credential.
type WorkspaceConnectionApiKeyResponse struct {
	Key *string `pulumi:"key"`
}

type WorkspaceConnectionManagedIdentity struct {
	ClientId   *string `pulumi:"clientId"`
	ResourceId *string `pulumi:"resourceId"`
}

type WorkspaceConnectionManagedIdentityResponse struct {
	ClientId   *string `pulumi:"clientId"`
	ResourceId *string `pulumi:"resourceId"`
}

type WorkspaceConnectionPersonalAccessToken struct {
	Pat *string `pulumi:"pat"`
}

type WorkspaceConnectionPersonalAccessTokenResponse struct {
	Pat *string `pulumi:"pat"`
}

type WorkspaceConnectionServicePrincipalResponse struct {
	ClientId     *string `pulumi:"clientId"`
	ClientSecret *string `pulumi:"clientSecret"`
	TenantId     *string `pulumi:"tenantId"`
}

type WorkspaceConnectionSharedAccessSignature struct {
	Sas *string `pulumi:"sas"`
}

type WorkspaceConnectionSharedAccessSignatureResponse struct {
	Sas *string `pulumi:"sas"`
}

type WorkspaceConnectionUsernamePassword struct {
	Password *string `pulumi:"password"`
	Username *string `pulumi:"username"`
}

type WorkspaceConnectionUsernamePasswordResponse struct {
	Password *string `pulumi:"password"`
	Username *string `pulumi:"username"`
}

func init() {
	pulumi.RegisterOutputType(UserAssignedIdentityResponseOutput{})
	pulumi.RegisterOutputType(UserAssignedIdentityResponseMapOutput{})
	pulumi.RegisterOutputType(UserCreatedAcrAccountOutput{})
	pulumi.RegisterOutputType(UserCreatedAcrAccountPtrOutput{})
	pulumi.RegisterOutputType(UserCreatedAcrAccountResponseOutput{})
	pulumi.RegisterOutputType(UserCreatedAcrAccountResponsePtrOutput{})
	pulumi.RegisterOutputType(UserCreatedStorageAccountOutput{})
	pulumi.RegisterOutputType(UserCreatedStorageAccountPtrOutput{})
	pulumi.RegisterOutputType(UserCreatedStorageAccountResponseOutput{})
	pulumi.RegisterOutputType(UserCreatedStorageAccountResponsePtrOutput{})
	pulumi.RegisterOutputType(UserInfoResponseOutput{})
	pulumi.RegisterOutputType(UserInfoResponsePtrOutput{})
}
