// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220110preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = utilities.GetEnvOrDefault

// Information about a Provider. A Provider is an entity that offers Targets to run Azure Quantum Jobs.
type Provider struct {
	// The provider's marketplace application display name.
	ApplicationName *string `pulumi:"applicationName"`
	// A Uri identifying the specific instance of this provider.
	InstanceUri *string `pulumi:"instanceUri"`
	// Unique id of this provider.
	ProviderId *string `pulumi:"providerId"`
	// The sku associated with pricing information for this provider.
	ProviderSku *string `pulumi:"providerSku"`
	// Provisioning status field
	ProvisioningState *string `pulumi:"provisioningState"`
	// Id to track resource usage for the provider.
	ResourceUsageId *string `pulumi:"resourceUsageId"`
}

// ProviderInput is an input type that accepts ProviderArgs and ProviderOutput values.
// You can construct a concrete instance of `ProviderInput` via:
//
//	ProviderArgs{...}
type ProviderInput interface {
	pulumi.Input

	ToProviderOutput() ProviderOutput
	ToProviderOutputWithContext(context.Context) ProviderOutput
}

// Information about a Provider. A Provider is an entity that offers Targets to run Azure Quantum Jobs.
type ProviderArgs struct {
	// The provider's marketplace application display name.
	ApplicationName pulumi.StringPtrInput `pulumi:"applicationName"`
	// A Uri identifying the specific instance of this provider.
	InstanceUri pulumi.StringPtrInput `pulumi:"instanceUri"`
	// Unique id of this provider.
	ProviderId pulumi.StringPtrInput `pulumi:"providerId"`
	// The sku associated with pricing information for this provider.
	ProviderSku pulumi.StringPtrInput `pulumi:"providerSku"`
	// Provisioning status field
	ProvisioningState pulumi.StringPtrInput `pulumi:"provisioningState"`
	// Id to track resource usage for the provider.
	ResourceUsageId pulumi.StringPtrInput `pulumi:"resourceUsageId"`
}

func (ProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Provider)(nil)).Elem()
}

func (i ProviderArgs) ToProviderOutput() ProviderOutput {
	return i.ToProviderOutputWithContext(context.Background())
}

func (i ProviderArgs) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderOutput)
}

// ProviderArrayInput is an input type that accepts ProviderArray and ProviderArrayOutput values.
// You can construct a concrete instance of `ProviderArrayInput` via:
//
//	ProviderArray{ ProviderArgs{...} }
type ProviderArrayInput interface {
	pulumi.Input

	ToProviderArrayOutput() ProviderArrayOutput
	ToProviderArrayOutputWithContext(context.Context) ProviderArrayOutput
}

type ProviderArray []ProviderInput

func (ProviderArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Provider)(nil)).Elem()
}

func (i ProviderArray) ToProviderArrayOutput() ProviderArrayOutput {
	return i.ToProviderArrayOutputWithContext(context.Background())
}

func (i ProviderArray) ToProviderArrayOutputWithContext(ctx context.Context) ProviderArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderArrayOutput)
}

// Information about a Provider. A Provider is an entity that offers Targets to run Azure Quantum Jobs.
type ProviderOutput struct{ *pulumi.OutputState }

func (ProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Provider)(nil)).Elem()
}

func (o ProviderOutput) ToProviderOutput() ProviderOutput {
	return o
}

func (o ProviderOutput) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return o
}

// The provider's marketplace application display name.
func (o ProviderOutput) ApplicationName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Provider) *string { return v.ApplicationName }).(pulumi.StringPtrOutput)
}

// A Uri identifying the specific instance of this provider.
func (o ProviderOutput) InstanceUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Provider) *string { return v.InstanceUri }).(pulumi.StringPtrOutput)
}

// Unique id of this provider.
func (o ProviderOutput) ProviderId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Provider) *string { return v.ProviderId }).(pulumi.StringPtrOutput)
}

// The sku associated with pricing information for this provider.
func (o ProviderOutput) ProviderSku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Provider) *string { return v.ProviderSku }).(pulumi.StringPtrOutput)
}

// Provisioning status field
func (o ProviderOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Provider) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

// Id to track resource usage for the provider.
func (o ProviderOutput) ResourceUsageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Provider) *string { return v.ResourceUsageId }).(pulumi.StringPtrOutput)
}

type ProviderArrayOutput struct{ *pulumi.OutputState }

func (ProviderArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Provider)(nil)).Elem()
}

func (o ProviderArrayOutput) ToProviderArrayOutput() ProviderArrayOutput {
	return o
}

func (o ProviderArrayOutput) ToProviderArrayOutputWithContext(ctx context.Context) ProviderArrayOutput {
	return o
}

func (o ProviderArrayOutput) Index(i pulumi.IntInput) ProviderOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Provider {
		return vs[0].([]Provider)[vs[1].(int)]
	}).(ProviderOutput)
}

// Information about a Provider. A Provider is an entity that offers Targets to run Azure Quantum Jobs.
type ProviderResponse struct {
	// The provider's marketplace application display name.
	ApplicationName *string `pulumi:"applicationName"`
	// A Uri identifying the specific instance of this provider.
	InstanceUri *string `pulumi:"instanceUri"`
	// Unique id of this provider.
	ProviderId *string `pulumi:"providerId"`
	// The sku associated with pricing information for this provider.
	ProviderSku *string `pulumi:"providerSku"`
	// Provisioning status field
	ProvisioningState *string `pulumi:"provisioningState"`
	// Id to track resource usage for the provider.
	ResourceUsageId *string `pulumi:"resourceUsageId"`
}

// Information about a Provider. A Provider is an entity that offers Targets to run Azure Quantum Jobs.
type ProviderResponseOutput struct{ *pulumi.OutputState }

func (ProviderResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProviderResponse)(nil)).Elem()
}

func (o ProviderResponseOutput) ToProviderResponseOutput() ProviderResponseOutput {
	return o
}

func (o ProviderResponseOutput) ToProviderResponseOutputWithContext(ctx context.Context) ProviderResponseOutput {
	return o
}

// The provider's marketplace application display name.
func (o ProviderResponseOutput) ApplicationName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProviderResponse) *string { return v.ApplicationName }).(pulumi.StringPtrOutput)
}

// A Uri identifying the specific instance of this provider.
func (o ProviderResponseOutput) InstanceUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProviderResponse) *string { return v.InstanceUri }).(pulumi.StringPtrOutput)
}

// Unique id of this provider.
func (o ProviderResponseOutput) ProviderId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProviderResponse) *string { return v.ProviderId }).(pulumi.StringPtrOutput)
}

// The sku associated with pricing information for this provider.
func (o ProviderResponseOutput) ProviderSku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProviderResponse) *string { return v.ProviderSku }).(pulumi.StringPtrOutput)
}

// Provisioning status field
func (o ProviderResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProviderResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

// Id to track resource usage for the provider.
func (o ProviderResponseOutput) ResourceUsageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProviderResponse) *string { return v.ResourceUsageId }).(pulumi.StringPtrOutput)
}

type ProviderResponseArrayOutput struct{ *pulumi.OutputState }

func (ProviderResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ProviderResponse)(nil)).Elem()
}

func (o ProviderResponseArrayOutput) ToProviderResponseArrayOutput() ProviderResponseArrayOutput {
	return o
}

func (o ProviderResponseArrayOutput) ToProviderResponseArrayOutputWithContext(ctx context.Context) ProviderResponseArrayOutput {
	return o
}

func (o ProviderResponseArrayOutput) Index(i pulumi.IntInput) ProviderResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ProviderResponse {
		return vs[0].([]ProviderResponse)[vs[1].(int)]
	}).(ProviderResponseOutput)
}

// Managed Identity information.
type QuantumWorkspaceIdentity struct {
	// The identity type.
	Type *string `pulumi:"type"`
}

// QuantumWorkspaceIdentityInput is an input type that accepts QuantumWorkspaceIdentityArgs and QuantumWorkspaceIdentityOutput values.
// You can construct a concrete instance of `QuantumWorkspaceIdentityInput` via:
//
//	QuantumWorkspaceIdentityArgs{...}
type QuantumWorkspaceIdentityInput interface {
	pulumi.Input

	ToQuantumWorkspaceIdentityOutput() QuantumWorkspaceIdentityOutput
	ToQuantumWorkspaceIdentityOutputWithContext(context.Context) QuantumWorkspaceIdentityOutput
}

// Managed Identity information.
type QuantumWorkspaceIdentityArgs struct {
	// The identity type.
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (QuantumWorkspaceIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*QuantumWorkspaceIdentity)(nil)).Elem()
}

func (i QuantumWorkspaceIdentityArgs) ToQuantumWorkspaceIdentityOutput() QuantumWorkspaceIdentityOutput {
	return i.ToQuantumWorkspaceIdentityOutputWithContext(context.Background())
}

func (i QuantumWorkspaceIdentityArgs) ToQuantumWorkspaceIdentityOutputWithContext(ctx context.Context) QuantumWorkspaceIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QuantumWorkspaceIdentityOutput)
}

func (i QuantumWorkspaceIdentityArgs) ToQuantumWorkspaceIdentityPtrOutput() QuantumWorkspaceIdentityPtrOutput {
	return i.ToQuantumWorkspaceIdentityPtrOutputWithContext(context.Background())
}

func (i QuantumWorkspaceIdentityArgs) ToQuantumWorkspaceIdentityPtrOutputWithContext(ctx context.Context) QuantumWorkspaceIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QuantumWorkspaceIdentityOutput).ToQuantumWorkspaceIdentityPtrOutputWithContext(ctx)
}

// QuantumWorkspaceIdentityPtrInput is an input type that accepts QuantumWorkspaceIdentityArgs, QuantumWorkspaceIdentityPtr and QuantumWorkspaceIdentityPtrOutput values.
// You can construct a concrete instance of `QuantumWorkspaceIdentityPtrInput` via:
//
//	        QuantumWorkspaceIdentityArgs{...}
//
//	or:
//
//	        nil
type QuantumWorkspaceIdentityPtrInput interface {
	pulumi.Input

	ToQuantumWorkspaceIdentityPtrOutput() QuantumWorkspaceIdentityPtrOutput
	ToQuantumWorkspaceIdentityPtrOutputWithContext(context.Context) QuantumWorkspaceIdentityPtrOutput
}

type quantumWorkspaceIdentityPtrType QuantumWorkspaceIdentityArgs

func QuantumWorkspaceIdentityPtr(v *QuantumWorkspaceIdentityArgs) QuantumWorkspaceIdentityPtrInput {
	return (*quantumWorkspaceIdentityPtrType)(v)
}

func (*quantumWorkspaceIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**QuantumWorkspaceIdentity)(nil)).Elem()
}

func (i *quantumWorkspaceIdentityPtrType) ToQuantumWorkspaceIdentityPtrOutput() QuantumWorkspaceIdentityPtrOutput {
	return i.ToQuantumWorkspaceIdentityPtrOutputWithContext(context.Background())
}

func (i *quantumWorkspaceIdentityPtrType) ToQuantumWorkspaceIdentityPtrOutputWithContext(ctx context.Context) QuantumWorkspaceIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QuantumWorkspaceIdentityPtrOutput)
}

// Managed Identity information.
type QuantumWorkspaceIdentityOutput struct{ *pulumi.OutputState }

func (QuantumWorkspaceIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*QuantumWorkspaceIdentity)(nil)).Elem()
}

func (o QuantumWorkspaceIdentityOutput) ToQuantumWorkspaceIdentityOutput() QuantumWorkspaceIdentityOutput {
	return o
}

func (o QuantumWorkspaceIdentityOutput) ToQuantumWorkspaceIdentityOutputWithContext(ctx context.Context) QuantumWorkspaceIdentityOutput {
	return o
}

func (o QuantumWorkspaceIdentityOutput) ToQuantumWorkspaceIdentityPtrOutput() QuantumWorkspaceIdentityPtrOutput {
	return o.ToQuantumWorkspaceIdentityPtrOutputWithContext(context.Background())
}

func (o QuantumWorkspaceIdentityOutput) ToQuantumWorkspaceIdentityPtrOutputWithContext(ctx context.Context) QuantumWorkspaceIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v QuantumWorkspaceIdentity) *QuantumWorkspaceIdentity {
		return &v
	}).(QuantumWorkspaceIdentityPtrOutput)
}

// The identity type.
func (o QuantumWorkspaceIdentityOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v QuantumWorkspaceIdentity) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type QuantumWorkspaceIdentityPtrOutput struct{ *pulumi.OutputState }

func (QuantumWorkspaceIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**QuantumWorkspaceIdentity)(nil)).Elem()
}

func (o QuantumWorkspaceIdentityPtrOutput) ToQuantumWorkspaceIdentityPtrOutput() QuantumWorkspaceIdentityPtrOutput {
	return o
}

func (o QuantumWorkspaceIdentityPtrOutput) ToQuantumWorkspaceIdentityPtrOutputWithContext(ctx context.Context) QuantumWorkspaceIdentityPtrOutput {
	return o
}

func (o QuantumWorkspaceIdentityPtrOutput) Elem() QuantumWorkspaceIdentityOutput {
	return o.ApplyT(func(v *QuantumWorkspaceIdentity) QuantumWorkspaceIdentity {
		if v != nil {
			return *v
		}
		var ret QuantumWorkspaceIdentity
		return ret
	}).(QuantumWorkspaceIdentityOutput)
}

// The identity type.
func (o QuantumWorkspaceIdentityPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *QuantumWorkspaceIdentity) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

// Managed Identity information.
type QuantumWorkspaceResponseIdentity struct {
	// The principal ID of resource identity.
	PrincipalId string `pulumi:"principalId"`
	// The tenant ID of resource.
	TenantId string `pulumi:"tenantId"`
	// The identity type.
	Type *string `pulumi:"type"`
}

// Managed Identity information.
type QuantumWorkspaceResponseIdentityOutput struct{ *pulumi.OutputState }

func (QuantumWorkspaceResponseIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*QuantumWorkspaceResponseIdentity)(nil)).Elem()
}

func (o QuantumWorkspaceResponseIdentityOutput) ToQuantumWorkspaceResponseIdentityOutput() QuantumWorkspaceResponseIdentityOutput {
	return o
}

func (o QuantumWorkspaceResponseIdentityOutput) ToQuantumWorkspaceResponseIdentityOutputWithContext(ctx context.Context) QuantumWorkspaceResponseIdentityOutput {
	return o
}

// The principal ID of resource identity.
func (o QuantumWorkspaceResponseIdentityOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v QuantumWorkspaceResponseIdentity) string { return v.PrincipalId }).(pulumi.StringOutput)
}

// The tenant ID of resource.
func (o QuantumWorkspaceResponseIdentityOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v QuantumWorkspaceResponseIdentity) string { return v.TenantId }).(pulumi.StringOutput)
}

// The identity type.
func (o QuantumWorkspaceResponseIdentityOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v QuantumWorkspaceResponseIdentity) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type QuantumWorkspaceResponseIdentityPtrOutput struct{ *pulumi.OutputState }

func (QuantumWorkspaceResponseIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**QuantumWorkspaceResponseIdentity)(nil)).Elem()
}

func (o QuantumWorkspaceResponseIdentityPtrOutput) ToQuantumWorkspaceResponseIdentityPtrOutput() QuantumWorkspaceResponseIdentityPtrOutput {
	return o
}

func (o QuantumWorkspaceResponseIdentityPtrOutput) ToQuantumWorkspaceResponseIdentityPtrOutputWithContext(ctx context.Context) QuantumWorkspaceResponseIdentityPtrOutput {
	return o
}

func (o QuantumWorkspaceResponseIdentityPtrOutput) Elem() QuantumWorkspaceResponseIdentityOutput {
	return o.ApplyT(func(v *QuantumWorkspaceResponseIdentity) QuantumWorkspaceResponseIdentity {
		if v != nil {
			return *v
		}
		var ret QuantumWorkspaceResponseIdentity
		return ret
	}).(QuantumWorkspaceResponseIdentityOutput)
}

// The principal ID of resource identity.
func (o QuantumWorkspaceResponseIdentityPtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *QuantumWorkspaceResponseIdentity) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

// The tenant ID of resource.
func (o QuantumWorkspaceResponseIdentityPtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *QuantumWorkspaceResponseIdentity) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

// The identity type.
func (o QuantumWorkspaceResponseIdentityPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *QuantumWorkspaceResponseIdentity) *string {
		if v == nil {
			return nil
		}
		return v.Type
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

func init() {
	pulumi.RegisterOutputType(ProviderOutput{})
	pulumi.RegisterOutputType(ProviderArrayOutput{})
	pulumi.RegisterOutputType(ProviderResponseOutput{})
	pulumi.RegisterOutputType(ProviderResponseArrayOutput{})
	pulumi.RegisterOutputType(QuantumWorkspaceIdentityOutput{})
	pulumi.RegisterOutputType(QuantumWorkspaceIdentityPtrOutput{})
	pulumi.RegisterOutputType(QuantumWorkspaceResponseIdentityOutput{})
	pulumi.RegisterOutputType(QuantumWorkspaceResponseIdentityPtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
}
