// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = utilities.GetEnvOrDefault

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

// User assigned identity properties
type UserAssignedIdentityResponse struct {
	// The client ID of the assigned identity.
	ClientId string `pulumi:"clientId"`
	// The principal ID of the assigned identity.
	PrincipalId string `pulumi:"principalId"`
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

// The client ID of the assigned identity.
func (o UserAssignedIdentityResponseOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v UserAssignedIdentityResponse) string { return v.ClientId }).(pulumi.StringOutput)
}

// The principal ID of the assigned identity.
func (o UserAssignedIdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v UserAssignedIdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
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

func (o UserAssignedIdentityResponseMapOutput) MapIndex(k pulumi.StringInput) UserAssignedIdentityResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) UserAssignedIdentityResponse {
		return vs[0].(map[string]UserAssignedIdentityResponse)[vs[1].(string)]
	}).(UserAssignedIdentityResponseOutput)
}

// Identity used for BYOS
type WorkbookResourceIdentity struct {
	// Type of managed service identity (where both SystemAssigned and UserAssigned types are allowed).
	Type string `pulumi:"type"`
	// The set of user assigned identities associated with the resource. The userAssignedIdentities dictionary keys will be ARM resource ids in the form: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}. The dictionary values can be empty objects ({}) in requests.
	UserAssignedIdentities []string `pulumi:"userAssignedIdentities"`
}

// WorkbookResourceIdentityInput is an input type that accepts WorkbookResourceIdentityArgs and WorkbookResourceIdentityOutput values.
// You can construct a concrete instance of `WorkbookResourceIdentityInput` via:
//
//	WorkbookResourceIdentityArgs{...}
type WorkbookResourceIdentityInput interface {
	pulumi.Input

	ToWorkbookResourceIdentityOutput() WorkbookResourceIdentityOutput
	ToWorkbookResourceIdentityOutputWithContext(context.Context) WorkbookResourceIdentityOutput
}

// Identity used for BYOS
type WorkbookResourceIdentityArgs struct {
	// Type of managed service identity (where both SystemAssigned and UserAssigned types are allowed).
	Type pulumi.StringInput `pulumi:"type"`
	// The set of user assigned identities associated with the resource. The userAssignedIdentities dictionary keys will be ARM resource ids in the form: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}. The dictionary values can be empty objects ({}) in requests.
	UserAssignedIdentities pulumi.StringArrayInput `pulumi:"userAssignedIdentities"`
}

func (WorkbookResourceIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkbookResourceIdentity)(nil)).Elem()
}

func (i WorkbookResourceIdentityArgs) ToWorkbookResourceIdentityOutput() WorkbookResourceIdentityOutput {
	return i.ToWorkbookResourceIdentityOutputWithContext(context.Background())
}

func (i WorkbookResourceIdentityArgs) ToWorkbookResourceIdentityOutputWithContext(ctx context.Context) WorkbookResourceIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkbookResourceIdentityOutput)
}

func (i WorkbookResourceIdentityArgs) ToWorkbookResourceIdentityPtrOutput() WorkbookResourceIdentityPtrOutput {
	return i.ToWorkbookResourceIdentityPtrOutputWithContext(context.Background())
}

func (i WorkbookResourceIdentityArgs) ToWorkbookResourceIdentityPtrOutputWithContext(ctx context.Context) WorkbookResourceIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkbookResourceIdentityOutput).ToWorkbookResourceIdentityPtrOutputWithContext(ctx)
}

// WorkbookResourceIdentityPtrInput is an input type that accepts WorkbookResourceIdentityArgs, WorkbookResourceIdentityPtr and WorkbookResourceIdentityPtrOutput values.
// You can construct a concrete instance of `WorkbookResourceIdentityPtrInput` via:
//
//	        WorkbookResourceIdentityArgs{...}
//
//	or:
//
//	        nil
type WorkbookResourceIdentityPtrInput interface {
	pulumi.Input

	ToWorkbookResourceIdentityPtrOutput() WorkbookResourceIdentityPtrOutput
	ToWorkbookResourceIdentityPtrOutputWithContext(context.Context) WorkbookResourceIdentityPtrOutput
}

type workbookResourceIdentityPtrType WorkbookResourceIdentityArgs

func WorkbookResourceIdentityPtr(v *WorkbookResourceIdentityArgs) WorkbookResourceIdentityPtrInput {
	return (*workbookResourceIdentityPtrType)(v)
}

func (*workbookResourceIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkbookResourceIdentity)(nil)).Elem()
}

func (i *workbookResourceIdentityPtrType) ToWorkbookResourceIdentityPtrOutput() WorkbookResourceIdentityPtrOutput {
	return i.ToWorkbookResourceIdentityPtrOutputWithContext(context.Background())
}

func (i *workbookResourceIdentityPtrType) ToWorkbookResourceIdentityPtrOutputWithContext(ctx context.Context) WorkbookResourceIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkbookResourceIdentityPtrOutput)
}

// Identity used for BYOS
type WorkbookResourceIdentityOutput struct{ *pulumi.OutputState }

func (WorkbookResourceIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkbookResourceIdentity)(nil)).Elem()
}

func (o WorkbookResourceIdentityOutput) ToWorkbookResourceIdentityOutput() WorkbookResourceIdentityOutput {
	return o
}

func (o WorkbookResourceIdentityOutput) ToWorkbookResourceIdentityOutputWithContext(ctx context.Context) WorkbookResourceIdentityOutput {
	return o
}

func (o WorkbookResourceIdentityOutput) ToWorkbookResourceIdentityPtrOutput() WorkbookResourceIdentityPtrOutput {
	return o.ToWorkbookResourceIdentityPtrOutputWithContext(context.Background())
}

func (o WorkbookResourceIdentityOutput) ToWorkbookResourceIdentityPtrOutputWithContext(ctx context.Context) WorkbookResourceIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WorkbookResourceIdentity) *WorkbookResourceIdentity {
		return &v
	}).(WorkbookResourceIdentityPtrOutput)
}

// Type of managed service identity (where both SystemAssigned and UserAssigned types are allowed).
func (o WorkbookResourceIdentityOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v WorkbookResourceIdentity) string { return v.Type }).(pulumi.StringOutput)
}

// The set of user assigned identities associated with the resource. The userAssignedIdentities dictionary keys will be ARM resource ids in the form: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}. The dictionary values can be empty objects ({}) in requests.
func (o WorkbookResourceIdentityOutput) UserAssignedIdentities() pulumi.StringArrayOutput {
	return o.ApplyT(func(v WorkbookResourceIdentity) []string { return v.UserAssignedIdentities }).(pulumi.StringArrayOutput)
}

type WorkbookResourceIdentityPtrOutput struct{ *pulumi.OutputState }

func (WorkbookResourceIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkbookResourceIdentity)(nil)).Elem()
}

func (o WorkbookResourceIdentityPtrOutput) ToWorkbookResourceIdentityPtrOutput() WorkbookResourceIdentityPtrOutput {
	return o
}

func (o WorkbookResourceIdentityPtrOutput) ToWorkbookResourceIdentityPtrOutputWithContext(ctx context.Context) WorkbookResourceIdentityPtrOutput {
	return o
}

func (o WorkbookResourceIdentityPtrOutput) Elem() WorkbookResourceIdentityOutput {
	return o.ApplyT(func(v *WorkbookResourceIdentity) WorkbookResourceIdentity {
		if v != nil {
			return *v
		}
		var ret WorkbookResourceIdentity
		return ret
	}).(WorkbookResourceIdentityOutput)
}

// Type of managed service identity (where both SystemAssigned and UserAssigned types are allowed).
func (o WorkbookResourceIdentityPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkbookResourceIdentity) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

// The set of user assigned identities associated with the resource. The userAssignedIdentities dictionary keys will be ARM resource ids in the form: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}. The dictionary values can be empty objects ({}) in requests.
func (o WorkbookResourceIdentityPtrOutput) UserAssignedIdentities() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *WorkbookResourceIdentity) []string {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(pulumi.StringArrayOutput)
}

// Identity used for BYOS
type WorkbookResourceResponseIdentity struct {
	// The service principal ID of the system assigned identity. This property will only be provided for a system assigned identity.
	PrincipalId string `pulumi:"principalId"`
	// The tenant ID of the system assigned identity. This property will only be provided for a system assigned identity.
	TenantId string `pulumi:"tenantId"`
	// Type of managed service identity (where both SystemAssigned and UserAssigned types are allowed).
	Type string `pulumi:"type"`
	// The set of user assigned identities associated with the resource. The userAssignedIdentities dictionary keys will be ARM resource ids in the form: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}. The dictionary values can be empty objects ({}) in requests.
	UserAssignedIdentities map[string]UserAssignedIdentityResponse `pulumi:"userAssignedIdentities"`
}

// Identity used for BYOS
type WorkbookResourceResponseIdentityOutput struct{ *pulumi.OutputState }

func (WorkbookResourceResponseIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkbookResourceResponseIdentity)(nil)).Elem()
}

func (o WorkbookResourceResponseIdentityOutput) ToWorkbookResourceResponseIdentityOutput() WorkbookResourceResponseIdentityOutput {
	return o
}

func (o WorkbookResourceResponseIdentityOutput) ToWorkbookResourceResponseIdentityOutputWithContext(ctx context.Context) WorkbookResourceResponseIdentityOutput {
	return o
}

// The service principal ID of the system assigned identity. This property will only be provided for a system assigned identity.
func (o WorkbookResourceResponseIdentityOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v WorkbookResourceResponseIdentity) string { return v.PrincipalId }).(pulumi.StringOutput)
}

// The tenant ID of the system assigned identity. This property will only be provided for a system assigned identity.
func (o WorkbookResourceResponseIdentityOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v WorkbookResourceResponseIdentity) string { return v.TenantId }).(pulumi.StringOutput)
}

// Type of managed service identity (where both SystemAssigned and UserAssigned types are allowed).
func (o WorkbookResourceResponseIdentityOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v WorkbookResourceResponseIdentity) string { return v.Type }).(pulumi.StringOutput)
}

// The set of user assigned identities associated with the resource. The userAssignedIdentities dictionary keys will be ARM resource ids in the form: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}. The dictionary values can be empty objects ({}) in requests.
func (o WorkbookResourceResponseIdentityOutput) UserAssignedIdentities() UserAssignedIdentityResponseMapOutput {
	return o.ApplyT(func(v WorkbookResourceResponseIdentity) map[string]UserAssignedIdentityResponse {
		return v.UserAssignedIdentities
	}).(UserAssignedIdentityResponseMapOutput)
}

type WorkbookResourceResponseIdentityPtrOutput struct{ *pulumi.OutputState }

func (WorkbookResourceResponseIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkbookResourceResponseIdentity)(nil)).Elem()
}

func (o WorkbookResourceResponseIdentityPtrOutput) ToWorkbookResourceResponseIdentityPtrOutput() WorkbookResourceResponseIdentityPtrOutput {
	return o
}

func (o WorkbookResourceResponseIdentityPtrOutput) ToWorkbookResourceResponseIdentityPtrOutputWithContext(ctx context.Context) WorkbookResourceResponseIdentityPtrOutput {
	return o
}

func (o WorkbookResourceResponseIdentityPtrOutput) Elem() WorkbookResourceResponseIdentityOutput {
	return o.ApplyT(func(v *WorkbookResourceResponseIdentity) WorkbookResourceResponseIdentity {
		if v != nil {
			return *v
		}
		var ret WorkbookResourceResponseIdentity
		return ret
	}).(WorkbookResourceResponseIdentityOutput)
}

// The service principal ID of the system assigned identity. This property will only be provided for a system assigned identity.
func (o WorkbookResourceResponseIdentityPtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkbookResourceResponseIdentity) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

// The tenant ID of the system assigned identity. This property will only be provided for a system assigned identity.
func (o WorkbookResourceResponseIdentityPtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkbookResourceResponseIdentity) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

// Type of managed service identity (where both SystemAssigned and UserAssigned types are allowed).
func (o WorkbookResourceResponseIdentityPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkbookResourceResponseIdentity) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

// The set of user assigned identities associated with the resource. The userAssignedIdentities dictionary keys will be ARM resource ids in the form: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}. The dictionary values can be empty objects ({}) in requests.
func (o WorkbookResourceResponseIdentityPtrOutput) UserAssignedIdentities() UserAssignedIdentityResponseMapOutput {
	return o.ApplyT(func(v *WorkbookResourceResponseIdentity) map[string]UserAssignedIdentityResponse {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(UserAssignedIdentityResponseMapOutput)
}

func init() {
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(UserAssignedIdentityResponseOutput{})
	pulumi.RegisterOutputType(UserAssignedIdentityResponseMapOutput{})
	pulumi.RegisterOutputType(WorkbookResourceIdentityOutput{})
	pulumi.RegisterOutputType(WorkbookResourceIdentityPtrOutput{})
	pulumi.RegisterOutputType(WorkbookResourceResponseIdentityOutput{})
	pulumi.RegisterOutputType(WorkbookResourceResponseIdentityPtrOutput{})
}
