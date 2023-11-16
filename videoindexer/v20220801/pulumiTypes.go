// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = utilities.GetEnvOrDefault

// Managed service identity (system assigned and/or user assigned identities)
type ManagedServiceIdentity struct {
	// Type of managed service identity (where both SystemAssigned and UserAssigned types are allowed).
	Type string `pulumi:"type"`
	// The set of user assigned identities associated with the resource. The userAssignedIdentities dictionary keys will be ARM resource ids in the form: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}. The dictionary values can be empty objects ({}) in requests.
	UserAssignedIdentities []string `pulumi:"userAssignedIdentities"`
}

// ManagedServiceIdentityInput is an input type that accepts ManagedServiceIdentityArgs and ManagedServiceIdentityOutput values.
// You can construct a concrete instance of `ManagedServiceIdentityInput` via:
//
//	ManagedServiceIdentityArgs{...}
type ManagedServiceIdentityInput interface {
	pulumi.Input

	ToManagedServiceIdentityOutput() ManagedServiceIdentityOutput
	ToManagedServiceIdentityOutputWithContext(context.Context) ManagedServiceIdentityOutput
}

// Managed service identity (system assigned and/or user assigned identities)
type ManagedServiceIdentityArgs struct {
	// Type of managed service identity (where both SystemAssigned and UserAssigned types are allowed).
	Type pulumi.StringInput `pulumi:"type"`
	// The set of user assigned identities associated with the resource. The userAssignedIdentities dictionary keys will be ARM resource ids in the form: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}. The dictionary values can be empty objects ({}) in requests.
	UserAssignedIdentities pulumi.StringArrayInput `pulumi:"userAssignedIdentities"`
}

func (ManagedServiceIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedServiceIdentity)(nil)).Elem()
}

func (i ManagedServiceIdentityArgs) ToManagedServiceIdentityOutput() ManagedServiceIdentityOutput {
	return i.ToManagedServiceIdentityOutputWithContext(context.Background())
}

func (i ManagedServiceIdentityArgs) ToManagedServiceIdentityOutputWithContext(ctx context.Context) ManagedServiceIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedServiceIdentityOutput)
}

func (i ManagedServiceIdentityArgs) ToManagedServiceIdentityPtrOutput() ManagedServiceIdentityPtrOutput {
	return i.ToManagedServiceIdentityPtrOutputWithContext(context.Background())
}

func (i ManagedServiceIdentityArgs) ToManagedServiceIdentityPtrOutputWithContext(ctx context.Context) ManagedServiceIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedServiceIdentityOutput).ToManagedServiceIdentityPtrOutputWithContext(ctx)
}

// ManagedServiceIdentityPtrInput is an input type that accepts ManagedServiceIdentityArgs, ManagedServiceIdentityPtr and ManagedServiceIdentityPtrOutput values.
// You can construct a concrete instance of `ManagedServiceIdentityPtrInput` via:
//
//	        ManagedServiceIdentityArgs{...}
//
//	or:
//
//	        nil
type ManagedServiceIdentityPtrInput interface {
	pulumi.Input

	ToManagedServiceIdentityPtrOutput() ManagedServiceIdentityPtrOutput
	ToManagedServiceIdentityPtrOutputWithContext(context.Context) ManagedServiceIdentityPtrOutput
}

type managedServiceIdentityPtrType ManagedServiceIdentityArgs

func ManagedServiceIdentityPtr(v *ManagedServiceIdentityArgs) ManagedServiceIdentityPtrInput {
	return (*managedServiceIdentityPtrType)(v)
}

func (*managedServiceIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedServiceIdentity)(nil)).Elem()
}

func (i *managedServiceIdentityPtrType) ToManagedServiceIdentityPtrOutput() ManagedServiceIdentityPtrOutput {
	return i.ToManagedServiceIdentityPtrOutputWithContext(context.Background())
}

func (i *managedServiceIdentityPtrType) ToManagedServiceIdentityPtrOutputWithContext(ctx context.Context) ManagedServiceIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedServiceIdentityPtrOutput)
}

// Managed service identity (system assigned and/or user assigned identities)
type ManagedServiceIdentityOutput struct{ *pulumi.OutputState }

func (ManagedServiceIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedServiceIdentity)(nil)).Elem()
}

func (o ManagedServiceIdentityOutput) ToManagedServiceIdentityOutput() ManagedServiceIdentityOutput {
	return o
}

func (o ManagedServiceIdentityOutput) ToManagedServiceIdentityOutputWithContext(ctx context.Context) ManagedServiceIdentityOutput {
	return o
}

func (o ManagedServiceIdentityOutput) ToManagedServiceIdentityPtrOutput() ManagedServiceIdentityPtrOutput {
	return o.ToManagedServiceIdentityPtrOutputWithContext(context.Background())
}

func (o ManagedServiceIdentityOutput) ToManagedServiceIdentityPtrOutputWithContext(ctx context.Context) ManagedServiceIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedServiceIdentity) *ManagedServiceIdentity {
		return &v
	}).(ManagedServiceIdentityPtrOutput)
}

// Type of managed service identity (where both SystemAssigned and UserAssigned types are allowed).
func (o ManagedServiceIdentityOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedServiceIdentity) string { return v.Type }).(pulumi.StringOutput)
}

// The set of user assigned identities associated with the resource. The userAssignedIdentities dictionary keys will be ARM resource ids in the form: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}. The dictionary values can be empty objects ({}) in requests.
func (o ManagedServiceIdentityOutput) UserAssignedIdentities() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ManagedServiceIdentity) []string { return v.UserAssignedIdentities }).(pulumi.StringArrayOutput)
}

type ManagedServiceIdentityPtrOutput struct{ *pulumi.OutputState }

func (ManagedServiceIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedServiceIdentity)(nil)).Elem()
}

func (o ManagedServiceIdentityPtrOutput) ToManagedServiceIdentityPtrOutput() ManagedServiceIdentityPtrOutput {
	return o
}

func (o ManagedServiceIdentityPtrOutput) ToManagedServiceIdentityPtrOutputWithContext(ctx context.Context) ManagedServiceIdentityPtrOutput {
	return o
}

func (o ManagedServiceIdentityPtrOutput) Elem() ManagedServiceIdentityOutput {
	return o.ApplyT(func(v *ManagedServiceIdentity) ManagedServiceIdentity {
		if v != nil {
			return *v
		}
		var ret ManagedServiceIdentity
		return ret
	}).(ManagedServiceIdentityOutput)
}

// Type of managed service identity (where both SystemAssigned and UserAssigned types are allowed).
func (o ManagedServiceIdentityPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedServiceIdentity) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

// The set of user assigned identities associated with the resource. The userAssignedIdentities dictionary keys will be ARM resource ids in the form: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}. The dictionary values can be empty objects ({}) in requests.
func (o ManagedServiceIdentityPtrOutput) UserAssignedIdentities() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ManagedServiceIdentity) []string {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(pulumi.StringArrayOutput)
}

// Managed service identity (system assigned and/or user assigned identities)
type ManagedServiceIdentityResponse struct {
	// The service principal ID of the system assigned identity. This property will only be provided for a system assigned identity.
	PrincipalId string `pulumi:"principalId"`
	// The tenant ID of the system assigned identity. This property will only be provided for a system assigned identity.
	TenantId string `pulumi:"tenantId"`
	// Type of managed service identity (where both SystemAssigned and UserAssigned types are allowed).
	Type string `pulumi:"type"`
	// The set of user assigned identities associated with the resource. The userAssignedIdentities dictionary keys will be ARM resource ids in the form: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}. The dictionary values can be empty objects ({}) in requests.
	UserAssignedIdentities map[string]UserAssignedIdentityResponse `pulumi:"userAssignedIdentities"`
}

// Managed service identity (system assigned and/or user assigned identities)
type ManagedServiceIdentityResponseOutput struct{ *pulumi.OutputState }

func (ManagedServiceIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedServiceIdentityResponse)(nil)).Elem()
}

func (o ManagedServiceIdentityResponseOutput) ToManagedServiceIdentityResponseOutput() ManagedServiceIdentityResponseOutput {
	return o
}

func (o ManagedServiceIdentityResponseOutput) ToManagedServiceIdentityResponseOutputWithContext(ctx context.Context) ManagedServiceIdentityResponseOutput {
	return o
}

// The service principal ID of the system assigned identity. This property will only be provided for a system assigned identity.
func (o ManagedServiceIdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedServiceIdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

// The tenant ID of the system assigned identity. This property will only be provided for a system assigned identity.
func (o ManagedServiceIdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedServiceIdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

// Type of managed service identity (where both SystemAssigned and UserAssigned types are allowed).
func (o ManagedServiceIdentityResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedServiceIdentityResponse) string { return v.Type }).(pulumi.StringOutput)
}

// The set of user assigned identities associated with the resource. The userAssignedIdentities dictionary keys will be ARM resource ids in the form: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}. The dictionary values can be empty objects ({}) in requests.
func (o ManagedServiceIdentityResponseOutput) UserAssignedIdentities() UserAssignedIdentityResponseMapOutput {
	return o.ApplyT(func(v ManagedServiceIdentityResponse) map[string]UserAssignedIdentityResponse {
		return v.UserAssignedIdentities
	}).(UserAssignedIdentityResponseMapOutput)
}

type ManagedServiceIdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (ManagedServiceIdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedServiceIdentityResponse)(nil)).Elem()
}

func (o ManagedServiceIdentityResponsePtrOutput) ToManagedServiceIdentityResponsePtrOutput() ManagedServiceIdentityResponsePtrOutput {
	return o
}

func (o ManagedServiceIdentityResponsePtrOutput) ToManagedServiceIdentityResponsePtrOutputWithContext(ctx context.Context) ManagedServiceIdentityResponsePtrOutput {
	return o
}

func (o ManagedServiceIdentityResponsePtrOutput) Elem() ManagedServiceIdentityResponseOutput {
	return o.ApplyT(func(v *ManagedServiceIdentityResponse) ManagedServiceIdentityResponse {
		if v != nil {
			return *v
		}
		var ret ManagedServiceIdentityResponse
		return ret
	}).(ManagedServiceIdentityResponseOutput)
}

// The service principal ID of the system assigned identity. This property will only be provided for a system assigned identity.
func (o ManagedServiceIdentityResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedServiceIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

// The tenant ID of the system assigned identity. This property will only be provided for a system assigned identity.
func (o ManagedServiceIdentityResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedServiceIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

// Type of managed service identity (where both SystemAssigned and UserAssigned types are allowed).
func (o ManagedServiceIdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedServiceIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

// The set of user assigned identities associated with the resource. The userAssignedIdentities dictionary keys will be ARM resource ids in the form: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}. The dictionary values can be empty objects ({}) in requests.
func (o ManagedServiceIdentityResponsePtrOutput) UserAssignedIdentities() UserAssignedIdentityResponseMapOutput {
	return o.ApplyT(func(v *ManagedServiceIdentityResponse) map[string]UserAssignedIdentityResponse {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(UserAssignedIdentityResponseMapOutput)
}

// The media services details
type MediaServicesForPutRequest struct {
	// The media services resource id
	ResourceId *string `pulumi:"resourceId"`
	// The user assigned identity to be used to grant permissions
	UserAssignedIdentity *string `pulumi:"userAssignedIdentity"`
}

// MediaServicesForPutRequestInput is an input type that accepts MediaServicesForPutRequestArgs and MediaServicesForPutRequestOutput values.
// You can construct a concrete instance of `MediaServicesForPutRequestInput` via:
//
//	MediaServicesForPutRequestArgs{...}
type MediaServicesForPutRequestInput interface {
	pulumi.Input

	ToMediaServicesForPutRequestOutput() MediaServicesForPutRequestOutput
	ToMediaServicesForPutRequestOutputWithContext(context.Context) MediaServicesForPutRequestOutput
}

// The media services details
type MediaServicesForPutRequestArgs struct {
	// The media services resource id
	ResourceId pulumi.StringPtrInput `pulumi:"resourceId"`
	// The user assigned identity to be used to grant permissions
	UserAssignedIdentity pulumi.StringPtrInput `pulumi:"userAssignedIdentity"`
}

func (MediaServicesForPutRequestArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MediaServicesForPutRequest)(nil)).Elem()
}

func (i MediaServicesForPutRequestArgs) ToMediaServicesForPutRequestOutput() MediaServicesForPutRequestOutput {
	return i.ToMediaServicesForPutRequestOutputWithContext(context.Background())
}

func (i MediaServicesForPutRequestArgs) ToMediaServicesForPutRequestOutputWithContext(ctx context.Context) MediaServicesForPutRequestOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MediaServicesForPutRequestOutput)
}

func (i MediaServicesForPutRequestArgs) ToMediaServicesForPutRequestPtrOutput() MediaServicesForPutRequestPtrOutput {
	return i.ToMediaServicesForPutRequestPtrOutputWithContext(context.Background())
}

func (i MediaServicesForPutRequestArgs) ToMediaServicesForPutRequestPtrOutputWithContext(ctx context.Context) MediaServicesForPutRequestPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MediaServicesForPutRequestOutput).ToMediaServicesForPutRequestPtrOutputWithContext(ctx)
}

// MediaServicesForPutRequestPtrInput is an input type that accepts MediaServicesForPutRequestArgs, MediaServicesForPutRequestPtr and MediaServicesForPutRequestPtrOutput values.
// You can construct a concrete instance of `MediaServicesForPutRequestPtrInput` via:
//
//	        MediaServicesForPutRequestArgs{...}
//
//	or:
//
//	        nil
type MediaServicesForPutRequestPtrInput interface {
	pulumi.Input

	ToMediaServicesForPutRequestPtrOutput() MediaServicesForPutRequestPtrOutput
	ToMediaServicesForPutRequestPtrOutputWithContext(context.Context) MediaServicesForPutRequestPtrOutput
}

type mediaServicesForPutRequestPtrType MediaServicesForPutRequestArgs

func MediaServicesForPutRequestPtr(v *MediaServicesForPutRequestArgs) MediaServicesForPutRequestPtrInput {
	return (*mediaServicesForPutRequestPtrType)(v)
}

func (*mediaServicesForPutRequestPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MediaServicesForPutRequest)(nil)).Elem()
}

func (i *mediaServicesForPutRequestPtrType) ToMediaServicesForPutRequestPtrOutput() MediaServicesForPutRequestPtrOutput {
	return i.ToMediaServicesForPutRequestPtrOutputWithContext(context.Background())
}

func (i *mediaServicesForPutRequestPtrType) ToMediaServicesForPutRequestPtrOutputWithContext(ctx context.Context) MediaServicesForPutRequestPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MediaServicesForPutRequestPtrOutput)
}

// The media services details
type MediaServicesForPutRequestOutput struct{ *pulumi.OutputState }

func (MediaServicesForPutRequestOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MediaServicesForPutRequest)(nil)).Elem()
}

func (o MediaServicesForPutRequestOutput) ToMediaServicesForPutRequestOutput() MediaServicesForPutRequestOutput {
	return o
}

func (o MediaServicesForPutRequestOutput) ToMediaServicesForPutRequestOutputWithContext(ctx context.Context) MediaServicesForPutRequestOutput {
	return o
}

func (o MediaServicesForPutRequestOutput) ToMediaServicesForPutRequestPtrOutput() MediaServicesForPutRequestPtrOutput {
	return o.ToMediaServicesForPutRequestPtrOutputWithContext(context.Background())
}

func (o MediaServicesForPutRequestOutput) ToMediaServicesForPutRequestPtrOutputWithContext(ctx context.Context) MediaServicesForPutRequestPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MediaServicesForPutRequest) *MediaServicesForPutRequest {
		return &v
	}).(MediaServicesForPutRequestPtrOutput)
}

// The media services resource id
func (o MediaServicesForPutRequestOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MediaServicesForPutRequest) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

// The user assigned identity to be used to grant permissions
func (o MediaServicesForPutRequestOutput) UserAssignedIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MediaServicesForPutRequest) *string { return v.UserAssignedIdentity }).(pulumi.StringPtrOutput)
}

type MediaServicesForPutRequestPtrOutput struct{ *pulumi.OutputState }

func (MediaServicesForPutRequestPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MediaServicesForPutRequest)(nil)).Elem()
}

func (o MediaServicesForPutRequestPtrOutput) ToMediaServicesForPutRequestPtrOutput() MediaServicesForPutRequestPtrOutput {
	return o
}

func (o MediaServicesForPutRequestPtrOutput) ToMediaServicesForPutRequestPtrOutputWithContext(ctx context.Context) MediaServicesForPutRequestPtrOutput {
	return o
}

func (o MediaServicesForPutRequestPtrOutput) Elem() MediaServicesForPutRequestOutput {
	return o.ApplyT(func(v *MediaServicesForPutRequest) MediaServicesForPutRequest {
		if v != nil {
			return *v
		}
		var ret MediaServicesForPutRequest
		return ret
	}).(MediaServicesForPutRequestOutput)
}

// The media services resource id
func (o MediaServicesForPutRequestPtrOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MediaServicesForPutRequest) *string {
		if v == nil {
			return nil
		}
		return v.ResourceId
	}).(pulumi.StringPtrOutput)
}

// The user assigned identity to be used to grant permissions
func (o MediaServicesForPutRequestPtrOutput) UserAssignedIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MediaServicesForPutRequest) *string {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentity
	}).(pulumi.StringPtrOutput)
}

// The media services details
type MediaServicesForPutRequestResponse struct {
	// The media services resource id
	ResourceId *string `pulumi:"resourceId"`
	// The user assigned identity to be used to grant permissions
	UserAssignedIdentity *string `pulumi:"userAssignedIdentity"`
}

// The media services details
type MediaServicesForPutRequestResponseOutput struct{ *pulumi.OutputState }

func (MediaServicesForPutRequestResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MediaServicesForPutRequestResponse)(nil)).Elem()
}

func (o MediaServicesForPutRequestResponseOutput) ToMediaServicesForPutRequestResponseOutput() MediaServicesForPutRequestResponseOutput {
	return o
}

func (o MediaServicesForPutRequestResponseOutput) ToMediaServicesForPutRequestResponseOutputWithContext(ctx context.Context) MediaServicesForPutRequestResponseOutput {
	return o
}

// The media services resource id
func (o MediaServicesForPutRequestResponseOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MediaServicesForPutRequestResponse) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

// The user assigned identity to be used to grant permissions
func (o MediaServicesForPutRequestResponseOutput) UserAssignedIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MediaServicesForPutRequestResponse) *string { return v.UserAssignedIdentity }).(pulumi.StringPtrOutput)
}

type MediaServicesForPutRequestResponsePtrOutput struct{ *pulumi.OutputState }

func (MediaServicesForPutRequestResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MediaServicesForPutRequestResponse)(nil)).Elem()
}

func (o MediaServicesForPutRequestResponsePtrOutput) ToMediaServicesForPutRequestResponsePtrOutput() MediaServicesForPutRequestResponsePtrOutput {
	return o
}

func (o MediaServicesForPutRequestResponsePtrOutput) ToMediaServicesForPutRequestResponsePtrOutputWithContext(ctx context.Context) MediaServicesForPutRequestResponsePtrOutput {
	return o
}

func (o MediaServicesForPutRequestResponsePtrOutput) Elem() MediaServicesForPutRequestResponseOutput {
	return o.ApplyT(func(v *MediaServicesForPutRequestResponse) MediaServicesForPutRequestResponse {
		if v != nil {
			return *v
		}
		var ret MediaServicesForPutRequestResponse
		return ret
	}).(MediaServicesForPutRequestResponseOutput)
}

// The media services resource id
func (o MediaServicesForPutRequestResponsePtrOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MediaServicesForPutRequestResponse) *string {
		if v == nil {
			return nil
		}
		return v.ResourceId
	}).(pulumi.StringPtrOutput)
}

// The user assigned identity to be used to grant permissions
func (o MediaServicesForPutRequestResponsePtrOutput) UserAssignedIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MediaServicesForPutRequestResponse) *string {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentity
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

func init() {
	pulumi.RegisterOutputType(ManagedServiceIdentityOutput{})
	pulumi.RegisterOutputType(ManagedServiceIdentityPtrOutput{})
	pulumi.RegisterOutputType(ManagedServiceIdentityResponseOutput{})
	pulumi.RegisterOutputType(ManagedServiceIdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(MediaServicesForPutRequestOutput{})
	pulumi.RegisterOutputType(MediaServicesForPutRequestPtrOutput{})
	pulumi.RegisterOutputType(MediaServicesForPutRequestResponseOutput{})
	pulumi.RegisterOutputType(MediaServicesForPutRequestResponsePtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(UserAssignedIdentityResponseOutput{})
	pulumi.RegisterOutputType(UserAssignedIdentityResponseMapOutput{})
}
