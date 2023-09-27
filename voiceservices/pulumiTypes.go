// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package voiceservices

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
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

func (i ManagedServiceIdentityArgs) ToOutput(ctx context.Context) pulumix.Output[ManagedServiceIdentity] {
	return pulumix.Output[ManagedServiceIdentity]{
		OutputState: i.ToManagedServiceIdentityOutputWithContext(ctx).OutputState,
	}
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

func (i *managedServiceIdentityPtrType) ToOutput(ctx context.Context) pulumix.Output[*ManagedServiceIdentity] {
	return pulumix.Output[*ManagedServiceIdentity]{
		OutputState: i.ToManagedServiceIdentityPtrOutputWithContext(ctx).OutputState,
	}
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

func (o ManagedServiceIdentityOutput) ToOutput(ctx context.Context) pulumix.Output[ManagedServiceIdentity] {
	return pulumix.Output[ManagedServiceIdentity]{
		OutputState: o.OutputState,
	}
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

func (o ManagedServiceIdentityPtrOutput) ToOutput(ctx context.Context) pulumix.Output[*ManagedServiceIdentity] {
	return pulumix.Output[*ManagedServiceIdentity]{
		OutputState: o.OutputState,
	}
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

func (o ManagedServiceIdentityResponseOutput) ToOutput(ctx context.Context) pulumix.Output[ManagedServiceIdentityResponse] {
	return pulumix.Output[ManagedServiceIdentityResponse]{
		OutputState: o.OutputState,
	}
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

func (o ManagedServiceIdentityResponsePtrOutput) ToOutput(ctx context.Context) pulumix.Output[*ManagedServiceIdentityResponse] {
	return pulumix.Output[*ManagedServiceIdentityResponse]{
		OutputState: o.OutputState,
	}
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

// The configuration used in this region as primary, and other regions as backup.
type PrimaryRegionProperties struct {
	// The allowed source IP address or CIDR ranges for media
	AllowedMediaSourceAddressPrefixes []string `pulumi:"allowedMediaSourceAddressPrefixes"`
	// The allowed source IP address or CIDR ranges for signaling
	AllowedSignalingSourceAddressPrefixes []string `pulumi:"allowedSignalingSourceAddressPrefixes"`
	// IP address to use to contact the ESRP from this region
	EsrpAddresses []string `pulumi:"esrpAddresses"`
	// IP address to use to contact the operator network from this region
	OperatorAddresses []string `pulumi:"operatorAddresses"`
}

// PrimaryRegionPropertiesInput is an input type that accepts PrimaryRegionPropertiesArgs and PrimaryRegionPropertiesOutput values.
// You can construct a concrete instance of `PrimaryRegionPropertiesInput` via:
//
//	PrimaryRegionPropertiesArgs{...}
type PrimaryRegionPropertiesInput interface {
	pulumi.Input

	ToPrimaryRegionPropertiesOutput() PrimaryRegionPropertiesOutput
	ToPrimaryRegionPropertiesOutputWithContext(context.Context) PrimaryRegionPropertiesOutput
}

// The configuration used in this region as primary, and other regions as backup.
type PrimaryRegionPropertiesArgs struct {
	// The allowed source IP address or CIDR ranges for media
	AllowedMediaSourceAddressPrefixes pulumi.StringArrayInput `pulumi:"allowedMediaSourceAddressPrefixes"`
	// The allowed source IP address or CIDR ranges for signaling
	AllowedSignalingSourceAddressPrefixes pulumi.StringArrayInput `pulumi:"allowedSignalingSourceAddressPrefixes"`
	// IP address to use to contact the ESRP from this region
	EsrpAddresses pulumi.StringArrayInput `pulumi:"esrpAddresses"`
	// IP address to use to contact the operator network from this region
	OperatorAddresses pulumi.StringArrayInput `pulumi:"operatorAddresses"`
}

func (PrimaryRegionPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrimaryRegionProperties)(nil)).Elem()
}

func (i PrimaryRegionPropertiesArgs) ToPrimaryRegionPropertiesOutput() PrimaryRegionPropertiesOutput {
	return i.ToPrimaryRegionPropertiesOutputWithContext(context.Background())
}

func (i PrimaryRegionPropertiesArgs) ToPrimaryRegionPropertiesOutputWithContext(ctx context.Context) PrimaryRegionPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrimaryRegionPropertiesOutput)
}

func (i PrimaryRegionPropertiesArgs) ToOutput(ctx context.Context) pulumix.Output[PrimaryRegionProperties] {
	return pulumix.Output[PrimaryRegionProperties]{
		OutputState: i.ToPrimaryRegionPropertiesOutputWithContext(ctx).OutputState,
	}
}

// The configuration used in this region as primary, and other regions as backup.
type PrimaryRegionPropertiesOutput struct{ *pulumi.OutputState }

func (PrimaryRegionPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrimaryRegionProperties)(nil)).Elem()
}

func (o PrimaryRegionPropertiesOutput) ToPrimaryRegionPropertiesOutput() PrimaryRegionPropertiesOutput {
	return o
}

func (o PrimaryRegionPropertiesOutput) ToPrimaryRegionPropertiesOutputWithContext(ctx context.Context) PrimaryRegionPropertiesOutput {
	return o
}

func (o PrimaryRegionPropertiesOutput) ToOutput(ctx context.Context) pulumix.Output[PrimaryRegionProperties] {
	return pulumix.Output[PrimaryRegionProperties]{
		OutputState: o.OutputState,
	}
}

// The allowed source IP address or CIDR ranges for media
func (o PrimaryRegionPropertiesOutput) AllowedMediaSourceAddressPrefixes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PrimaryRegionProperties) []string { return v.AllowedMediaSourceAddressPrefixes }).(pulumi.StringArrayOutput)
}

// The allowed source IP address or CIDR ranges for signaling
func (o PrimaryRegionPropertiesOutput) AllowedSignalingSourceAddressPrefixes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PrimaryRegionProperties) []string { return v.AllowedSignalingSourceAddressPrefixes }).(pulumi.StringArrayOutput)
}

// IP address to use to contact the ESRP from this region
func (o PrimaryRegionPropertiesOutput) EsrpAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PrimaryRegionProperties) []string { return v.EsrpAddresses }).(pulumi.StringArrayOutput)
}

// IP address to use to contact the operator network from this region
func (o PrimaryRegionPropertiesOutput) OperatorAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PrimaryRegionProperties) []string { return v.OperatorAddresses }).(pulumi.StringArrayOutput)
}

// The configuration used in this region as primary, and other regions as backup.
type PrimaryRegionPropertiesResponse struct {
	// The allowed source IP address or CIDR ranges for media
	AllowedMediaSourceAddressPrefixes []string `pulumi:"allowedMediaSourceAddressPrefixes"`
	// The allowed source IP address or CIDR ranges for signaling
	AllowedSignalingSourceAddressPrefixes []string `pulumi:"allowedSignalingSourceAddressPrefixes"`
	// IP address to use to contact the ESRP from this region
	EsrpAddresses []string `pulumi:"esrpAddresses"`
	// IP address to use to contact the operator network from this region
	OperatorAddresses []string `pulumi:"operatorAddresses"`
}

// The configuration used in this region as primary, and other regions as backup.
type PrimaryRegionPropertiesResponseOutput struct{ *pulumi.OutputState }

func (PrimaryRegionPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrimaryRegionPropertiesResponse)(nil)).Elem()
}

func (o PrimaryRegionPropertiesResponseOutput) ToPrimaryRegionPropertiesResponseOutput() PrimaryRegionPropertiesResponseOutput {
	return o
}

func (o PrimaryRegionPropertiesResponseOutput) ToPrimaryRegionPropertiesResponseOutputWithContext(ctx context.Context) PrimaryRegionPropertiesResponseOutput {
	return o
}

func (o PrimaryRegionPropertiesResponseOutput) ToOutput(ctx context.Context) pulumix.Output[PrimaryRegionPropertiesResponse] {
	return pulumix.Output[PrimaryRegionPropertiesResponse]{
		OutputState: o.OutputState,
	}
}

// The allowed source IP address or CIDR ranges for media
func (o PrimaryRegionPropertiesResponseOutput) AllowedMediaSourceAddressPrefixes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PrimaryRegionPropertiesResponse) []string { return v.AllowedMediaSourceAddressPrefixes }).(pulumi.StringArrayOutput)
}

// The allowed source IP address or CIDR ranges for signaling
func (o PrimaryRegionPropertiesResponseOutput) AllowedSignalingSourceAddressPrefixes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PrimaryRegionPropertiesResponse) []string { return v.AllowedSignalingSourceAddressPrefixes }).(pulumi.StringArrayOutput)
}

// IP address to use to contact the ESRP from this region
func (o PrimaryRegionPropertiesResponseOutput) EsrpAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PrimaryRegionPropertiesResponse) []string { return v.EsrpAddresses }).(pulumi.StringArrayOutput)
}

// IP address to use to contact the operator network from this region
func (o PrimaryRegionPropertiesResponseOutput) OperatorAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PrimaryRegionPropertiesResponse) []string { return v.OperatorAddresses }).(pulumi.StringArrayOutput)
}

// The service region configuration needed for Teams Callings.
type ServiceRegionProperties struct {
	// The name of the region in which the resources needed for Teams Calling will be deployed.
	Name string `pulumi:"name"`
	// The configuration used in this region as primary, and other regions as backup.
	PrimaryRegionProperties PrimaryRegionProperties `pulumi:"primaryRegionProperties"`
}

// ServiceRegionPropertiesInput is an input type that accepts ServiceRegionPropertiesArgs and ServiceRegionPropertiesOutput values.
// You can construct a concrete instance of `ServiceRegionPropertiesInput` via:
//
//	ServiceRegionPropertiesArgs{...}
type ServiceRegionPropertiesInput interface {
	pulumi.Input

	ToServiceRegionPropertiesOutput() ServiceRegionPropertiesOutput
	ToServiceRegionPropertiesOutputWithContext(context.Context) ServiceRegionPropertiesOutput
}

// The service region configuration needed for Teams Callings.
type ServiceRegionPropertiesArgs struct {
	// The name of the region in which the resources needed for Teams Calling will be deployed.
	Name pulumi.StringInput `pulumi:"name"`
	// The configuration used in this region as primary, and other regions as backup.
	PrimaryRegionProperties PrimaryRegionPropertiesInput `pulumi:"primaryRegionProperties"`
}

func (ServiceRegionPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceRegionProperties)(nil)).Elem()
}

func (i ServiceRegionPropertiesArgs) ToServiceRegionPropertiesOutput() ServiceRegionPropertiesOutput {
	return i.ToServiceRegionPropertiesOutputWithContext(context.Background())
}

func (i ServiceRegionPropertiesArgs) ToServiceRegionPropertiesOutputWithContext(ctx context.Context) ServiceRegionPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceRegionPropertiesOutput)
}

func (i ServiceRegionPropertiesArgs) ToOutput(ctx context.Context) pulumix.Output[ServiceRegionProperties] {
	return pulumix.Output[ServiceRegionProperties]{
		OutputState: i.ToServiceRegionPropertiesOutputWithContext(ctx).OutputState,
	}
}

// ServiceRegionPropertiesArrayInput is an input type that accepts ServiceRegionPropertiesArray and ServiceRegionPropertiesArrayOutput values.
// You can construct a concrete instance of `ServiceRegionPropertiesArrayInput` via:
//
//	ServiceRegionPropertiesArray{ ServiceRegionPropertiesArgs{...} }
type ServiceRegionPropertiesArrayInput interface {
	pulumi.Input

	ToServiceRegionPropertiesArrayOutput() ServiceRegionPropertiesArrayOutput
	ToServiceRegionPropertiesArrayOutputWithContext(context.Context) ServiceRegionPropertiesArrayOutput
}

type ServiceRegionPropertiesArray []ServiceRegionPropertiesInput

func (ServiceRegionPropertiesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceRegionProperties)(nil)).Elem()
}

func (i ServiceRegionPropertiesArray) ToServiceRegionPropertiesArrayOutput() ServiceRegionPropertiesArrayOutput {
	return i.ToServiceRegionPropertiesArrayOutputWithContext(context.Background())
}

func (i ServiceRegionPropertiesArray) ToServiceRegionPropertiesArrayOutputWithContext(ctx context.Context) ServiceRegionPropertiesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceRegionPropertiesArrayOutput)
}

func (i ServiceRegionPropertiesArray) ToOutput(ctx context.Context) pulumix.Output[[]ServiceRegionProperties] {
	return pulumix.Output[[]ServiceRegionProperties]{
		OutputState: i.ToServiceRegionPropertiesArrayOutputWithContext(ctx).OutputState,
	}
}

// The service region configuration needed for Teams Callings.
type ServiceRegionPropertiesOutput struct{ *pulumi.OutputState }

func (ServiceRegionPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceRegionProperties)(nil)).Elem()
}

func (o ServiceRegionPropertiesOutput) ToServiceRegionPropertiesOutput() ServiceRegionPropertiesOutput {
	return o
}

func (o ServiceRegionPropertiesOutput) ToServiceRegionPropertiesOutputWithContext(ctx context.Context) ServiceRegionPropertiesOutput {
	return o
}

func (o ServiceRegionPropertiesOutput) ToOutput(ctx context.Context) pulumix.Output[ServiceRegionProperties] {
	return pulumix.Output[ServiceRegionProperties]{
		OutputState: o.OutputState,
	}
}

// The name of the region in which the resources needed for Teams Calling will be deployed.
func (o ServiceRegionPropertiesOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceRegionProperties) string { return v.Name }).(pulumi.StringOutput)
}

// The configuration used in this region as primary, and other regions as backup.
func (o ServiceRegionPropertiesOutput) PrimaryRegionProperties() PrimaryRegionPropertiesOutput {
	return o.ApplyT(func(v ServiceRegionProperties) PrimaryRegionProperties { return v.PrimaryRegionProperties }).(PrimaryRegionPropertiesOutput)
}

type ServiceRegionPropertiesArrayOutput struct{ *pulumi.OutputState }

func (ServiceRegionPropertiesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceRegionProperties)(nil)).Elem()
}

func (o ServiceRegionPropertiesArrayOutput) ToServiceRegionPropertiesArrayOutput() ServiceRegionPropertiesArrayOutput {
	return o
}

func (o ServiceRegionPropertiesArrayOutput) ToServiceRegionPropertiesArrayOutputWithContext(ctx context.Context) ServiceRegionPropertiesArrayOutput {
	return o
}

func (o ServiceRegionPropertiesArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]ServiceRegionProperties] {
	return pulumix.Output[[]ServiceRegionProperties]{
		OutputState: o.OutputState,
	}
}

func (o ServiceRegionPropertiesArrayOutput) Index(i pulumi.IntInput) ServiceRegionPropertiesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ServiceRegionProperties {
		return vs[0].([]ServiceRegionProperties)[vs[1].(int)]
	}).(ServiceRegionPropertiesOutput)
}

// The service region configuration needed for Teams Callings.
type ServiceRegionPropertiesResponse struct {
	// The name of the region in which the resources needed for Teams Calling will be deployed.
	Name string `pulumi:"name"`
	// The configuration used in this region as primary, and other regions as backup.
	PrimaryRegionProperties PrimaryRegionPropertiesResponse `pulumi:"primaryRegionProperties"`
}

// The service region configuration needed for Teams Callings.
type ServiceRegionPropertiesResponseOutput struct{ *pulumi.OutputState }

func (ServiceRegionPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceRegionPropertiesResponse)(nil)).Elem()
}

func (o ServiceRegionPropertiesResponseOutput) ToServiceRegionPropertiesResponseOutput() ServiceRegionPropertiesResponseOutput {
	return o
}

func (o ServiceRegionPropertiesResponseOutput) ToServiceRegionPropertiesResponseOutputWithContext(ctx context.Context) ServiceRegionPropertiesResponseOutput {
	return o
}

func (o ServiceRegionPropertiesResponseOutput) ToOutput(ctx context.Context) pulumix.Output[ServiceRegionPropertiesResponse] {
	return pulumix.Output[ServiceRegionPropertiesResponse]{
		OutputState: o.OutputState,
	}
}

// The name of the region in which the resources needed for Teams Calling will be deployed.
func (o ServiceRegionPropertiesResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceRegionPropertiesResponse) string { return v.Name }).(pulumi.StringOutput)
}

// The configuration used in this region as primary, and other regions as backup.
func (o ServiceRegionPropertiesResponseOutput) PrimaryRegionProperties() PrimaryRegionPropertiesResponseOutput {
	return o.ApplyT(func(v ServiceRegionPropertiesResponse) PrimaryRegionPropertiesResponse {
		return v.PrimaryRegionProperties
	}).(PrimaryRegionPropertiesResponseOutput)
}

type ServiceRegionPropertiesResponseArrayOutput struct{ *pulumi.OutputState }

func (ServiceRegionPropertiesResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceRegionPropertiesResponse)(nil)).Elem()
}

func (o ServiceRegionPropertiesResponseArrayOutput) ToServiceRegionPropertiesResponseArrayOutput() ServiceRegionPropertiesResponseArrayOutput {
	return o
}

func (o ServiceRegionPropertiesResponseArrayOutput) ToServiceRegionPropertiesResponseArrayOutputWithContext(ctx context.Context) ServiceRegionPropertiesResponseArrayOutput {
	return o
}

func (o ServiceRegionPropertiesResponseArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]ServiceRegionPropertiesResponse] {
	return pulumix.Output[[]ServiceRegionPropertiesResponse]{
		OutputState: o.OutputState,
	}
}

func (o ServiceRegionPropertiesResponseArrayOutput) Index(i pulumi.IntInput) ServiceRegionPropertiesResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ServiceRegionPropertiesResponse {
		return vs[0].([]ServiceRegionPropertiesResponse)[vs[1].(int)]
	}).(ServiceRegionPropertiesResponseOutput)
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

func init() {
	pulumi.RegisterOutputType(ManagedServiceIdentityOutput{})
	pulumi.RegisterOutputType(ManagedServiceIdentityPtrOutput{})
	pulumi.RegisterOutputType(ManagedServiceIdentityResponseOutput{})
	pulumi.RegisterOutputType(ManagedServiceIdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(PrimaryRegionPropertiesOutput{})
	pulumi.RegisterOutputType(PrimaryRegionPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ServiceRegionPropertiesOutput{})
	pulumi.RegisterOutputType(ServiceRegionPropertiesArrayOutput{})
	pulumi.RegisterOutputType(ServiceRegionPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ServiceRegionPropertiesResponseArrayOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(UserAssignedIdentityResponseOutput{})
	pulumi.RegisterOutputType(UserAssignedIdentityResponseMapOutput{})
}
