// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20211101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = utilities.GetEnvOrDefault

// The authentication info when authType is secret
type SecretAuthInfo struct {
	// The authentication type.
	// Expected value is 'secret'.
	AuthType string `pulumi:"authType"`
	// Username or account name for secret auth.
	Name *string `pulumi:"name"`
	// Password or account key for secret auth.
	Secret *string `pulumi:"secret"`
}

// The authentication info when authType is secret
type SecretAuthInfoResponse struct {
	// The authentication type.
	// Expected value is 'secret'.
	AuthType string `pulumi:"authType"`
	// Username or account name for secret auth.
	Name *string `pulumi:"name"`
	// Password or account key for secret auth.
	Secret *string `pulumi:"secret"`
}

// An option to store secret value in secure place
type SecretStore struct {
	// The key vault id to store secret
	KeyVaultId *string `pulumi:"keyVaultId"`
}

// SecretStoreInput is an input type that accepts SecretStoreArgs and SecretStoreOutput values.
// You can construct a concrete instance of `SecretStoreInput` via:
//
//	SecretStoreArgs{...}
type SecretStoreInput interface {
	pulumi.Input

	ToSecretStoreOutput() SecretStoreOutput
	ToSecretStoreOutputWithContext(context.Context) SecretStoreOutput
}

// An option to store secret value in secure place
type SecretStoreArgs struct {
	// The key vault id to store secret
	KeyVaultId pulumi.StringPtrInput `pulumi:"keyVaultId"`
}

func (SecretStoreArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretStore)(nil)).Elem()
}

func (i SecretStoreArgs) ToSecretStoreOutput() SecretStoreOutput {
	return i.ToSecretStoreOutputWithContext(context.Background())
}

func (i SecretStoreArgs) ToSecretStoreOutputWithContext(ctx context.Context) SecretStoreOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretStoreOutput)
}

func (i SecretStoreArgs) ToSecretStorePtrOutput() SecretStorePtrOutput {
	return i.ToSecretStorePtrOutputWithContext(context.Background())
}

func (i SecretStoreArgs) ToSecretStorePtrOutputWithContext(ctx context.Context) SecretStorePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretStoreOutput).ToSecretStorePtrOutputWithContext(ctx)
}

// SecretStorePtrInput is an input type that accepts SecretStoreArgs, SecretStorePtr and SecretStorePtrOutput values.
// You can construct a concrete instance of `SecretStorePtrInput` via:
//
//	        SecretStoreArgs{...}
//
//	or:
//
//	        nil
type SecretStorePtrInput interface {
	pulumi.Input

	ToSecretStorePtrOutput() SecretStorePtrOutput
	ToSecretStorePtrOutputWithContext(context.Context) SecretStorePtrOutput
}

type secretStorePtrType SecretStoreArgs

func SecretStorePtr(v *SecretStoreArgs) SecretStorePtrInput {
	return (*secretStorePtrType)(v)
}

func (*secretStorePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretStore)(nil)).Elem()
}

func (i *secretStorePtrType) ToSecretStorePtrOutput() SecretStorePtrOutput {
	return i.ToSecretStorePtrOutputWithContext(context.Background())
}

func (i *secretStorePtrType) ToSecretStorePtrOutputWithContext(ctx context.Context) SecretStorePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretStorePtrOutput)
}

// An option to store secret value in secure place
type SecretStoreOutput struct{ *pulumi.OutputState }

func (SecretStoreOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretStore)(nil)).Elem()
}

func (o SecretStoreOutput) ToSecretStoreOutput() SecretStoreOutput {
	return o
}

func (o SecretStoreOutput) ToSecretStoreOutputWithContext(ctx context.Context) SecretStoreOutput {
	return o
}

func (o SecretStoreOutput) ToSecretStorePtrOutput() SecretStorePtrOutput {
	return o.ToSecretStorePtrOutputWithContext(context.Background())
}

func (o SecretStoreOutput) ToSecretStorePtrOutputWithContext(ctx context.Context) SecretStorePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SecretStore) *SecretStore {
		return &v
	}).(SecretStorePtrOutput)
}

// The key vault id to store secret
func (o SecretStoreOutput) KeyVaultId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecretStore) *string { return v.KeyVaultId }).(pulumi.StringPtrOutput)
}

type SecretStorePtrOutput struct{ *pulumi.OutputState }

func (SecretStorePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretStore)(nil)).Elem()
}

func (o SecretStorePtrOutput) ToSecretStorePtrOutput() SecretStorePtrOutput {
	return o
}

func (o SecretStorePtrOutput) ToSecretStorePtrOutputWithContext(ctx context.Context) SecretStorePtrOutput {
	return o
}

func (o SecretStorePtrOutput) Elem() SecretStoreOutput {
	return o.ApplyT(func(v *SecretStore) SecretStore {
		if v != nil {
			return *v
		}
		var ret SecretStore
		return ret
	}).(SecretStoreOutput)
}

// The key vault id to store secret
func (o SecretStorePtrOutput) KeyVaultId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretStore) *string {
		if v == nil {
			return nil
		}
		return v.KeyVaultId
	}).(pulumi.StringPtrOutput)
}

// An option to store secret value in secure place
type SecretStoreResponse struct {
	// The key vault id to store secret
	KeyVaultId *string `pulumi:"keyVaultId"`
}

// An option to store secret value in secure place
type SecretStoreResponseOutput struct{ *pulumi.OutputState }

func (SecretStoreResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretStoreResponse)(nil)).Elem()
}

func (o SecretStoreResponseOutput) ToSecretStoreResponseOutput() SecretStoreResponseOutput {
	return o
}

func (o SecretStoreResponseOutput) ToSecretStoreResponseOutputWithContext(ctx context.Context) SecretStoreResponseOutput {
	return o
}

// The key vault id to store secret
func (o SecretStoreResponseOutput) KeyVaultId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecretStoreResponse) *string { return v.KeyVaultId }).(pulumi.StringPtrOutput)
}

type SecretStoreResponsePtrOutput struct{ *pulumi.OutputState }

func (SecretStoreResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretStoreResponse)(nil)).Elem()
}

func (o SecretStoreResponsePtrOutput) ToSecretStoreResponsePtrOutput() SecretStoreResponsePtrOutput {
	return o
}

func (o SecretStoreResponsePtrOutput) ToSecretStoreResponsePtrOutputWithContext(ctx context.Context) SecretStoreResponsePtrOutput {
	return o
}

func (o SecretStoreResponsePtrOutput) Elem() SecretStoreResponseOutput {
	return o.ApplyT(func(v *SecretStoreResponse) SecretStoreResponse {
		if v != nil {
			return *v
		}
		var ret SecretStoreResponse
		return ret
	}).(SecretStoreResponseOutput)
}

// The key vault id to store secret
func (o SecretStoreResponsePtrOutput) KeyVaultId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretStoreResponse) *string {
		if v == nil {
			return nil
		}
		return v.KeyVaultId
	}).(pulumi.StringPtrOutput)
}

// The authentication info when authType is servicePrincipal certificate
type ServicePrincipalCertificateAuthInfo struct {
	// The authentication type.
	// Expected value is 'servicePrincipalCertificate'.
	AuthType string `pulumi:"authType"`
	// ServicePrincipal certificate for servicePrincipal auth.
	Certificate string `pulumi:"certificate"`
	// Application clientId for servicePrincipal auth.
	ClientId string `pulumi:"clientId"`
	// Principal Id for servicePrincipal auth.
	PrincipalId string `pulumi:"principalId"`
}

// The authentication info when authType is servicePrincipal certificate
type ServicePrincipalCertificateAuthInfoResponse struct {
	// The authentication type.
	// Expected value is 'servicePrincipalCertificate'.
	AuthType string `pulumi:"authType"`
	// ServicePrincipal certificate for servicePrincipal auth.
	Certificate string `pulumi:"certificate"`
	// Application clientId for servicePrincipal auth.
	ClientId string `pulumi:"clientId"`
	// Principal Id for servicePrincipal auth.
	PrincipalId string `pulumi:"principalId"`
}

// The authentication info when authType is servicePrincipal secret
type ServicePrincipalSecretAuthInfo struct {
	// The authentication type.
	// Expected value is 'servicePrincipalSecret'.
	AuthType string `pulumi:"authType"`
	// ServicePrincipal application clientId for servicePrincipal auth.
	ClientId string `pulumi:"clientId"`
	// Principal Id for servicePrincipal auth.
	PrincipalId string `pulumi:"principalId"`
	// Secret for servicePrincipal auth.
	Secret string `pulumi:"secret"`
}

// The authentication info when authType is servicePrincipal secret
type ServicePrincipalSecretAuthInfoResponse struct {
	// The authentication type.
	// Expected value is 'servicePrincipalSecret'.
	AuthType string `pulumi:"authType"`
	// ServicePrincipal application clientId for servicePrincipal auth.
	ClientId string `pulumi:"clientId"`
	// Principal Id for servicePrincipal auth.
	PrincipalId string `pulumi:"principalId"`
	// Secret for servicePrincipal auth.
	Secret string `pulumi:"secret"`
}

// A configuration item for source resource
type SourceConfigurationResponse struct {
	// The name of setting.
	Name *string `pulumi:"name"`
	// The value of setting
	Value *string `pulumi:"value"`
}

// A configuration item for source resource
type SourceConfigurationResponseOutput struct{ *pulumi.OutputState }

func (SourceConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SourceConfigurationResponse)(nil)).Elem()
}

func (o SourceConfigurationResponseOutput) ToSourceConfigurationResponseOutput() SourceConfigurationResponseOutput {
	return o
}

func (o SourceConfigurationResponseOutput) ToSourceConfigurationResponseOutputWithContext(ctx context.Context) SourceConfigurationResponseOutput {
	return o
}

// The name of setting.
func (o SourceConfigurationResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SourceConfigurationResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The value of setting
func (o SourceConfigurationResponseOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SourceConfigurationResponse) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type SourceConfigurationResponseArrayOutput struct{ *pulumi.OutputState }

func (SourceConfigurationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SourceConfigurationResponse)(nil)).Elem()
}

func (o SourceConfigurationResponseArrayOutput) ToSourceConfigurationResponseArrayOutput() SourceConfigurationResponseArrayOutput {
	return o
}

func (o SourceConfigurationResponseArrayOutput) ToSourceConfigurationResponseArrayOutputWithContext(ctx context.Context) SourceConfigurationResponseArrayOutput {
	return o
}

func (o SourceConfigurationResponseArrayOutput) Index(i pulumi.IntInput) SourceConfigurationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SourceConfigurationResponse {
		return vs[0].([]SourceConfigurationResponse)[vs[1].(int)]
	}).(SourceConfigurationResponseOutput)
}

// The authentication info when authType is systemAssignedIdentity
type SystemAssignedIdentityAuthInfo struct {
	// The authentication type.
	// Expected value is 'systemAssignedIdentity'.
	AuthType string `pulumi:"authType"`
}

// The authentication info when authType is systemAssignedIdentity
type SystemAssignedIdentityAuthInfoResponse struct {
	// The authentication type.
	// Expected value is 'systemAssignedIdentity'.
	AuthType string `pulumi:"authType"`
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

// The authentication info when authType is userAssignedIdentity
type UserAssignedIdentityAuthInfo struct {
	// The authentication type.
	// Expected value is 'userAssignedIdentity'.
	AuthType string `pulumi:"authType"`
	// Client Id for userAssignedIdentity.
	ClientId string `pulumi:"clientId"`
	// Subscription id for userAssignedIdentity.
	SubscriptionId string `pulumi:"subscriptionId"`
}

// The authentication info when authType is userAssignedIdentity
type UserAssignedIdentityAuthInfoResponse struct {
	// The authentication type.
	// Expected value is 'userAssignedIdentity'.
	AuthType string `pulumi:"authType"`
	// Client Id for userAssignedIdentity.
	ClientId string `pulumi:"clientId"`
	// Subscription id for userAssignedIdentity.
	SubscriptionId string `pulumi:"subscriptionId"`
}

// The VNet solution for linker
type VNetSolution struct {
	// Type of VNet solution.
	Type *string `pulumi:"type"`
}

// VNetSolutionInput is an input type that accepts VNetSolutionArgs and VNetSolutionOutput values.
// You can construct a concrete instance of `VNetSolutionInput` via:
//
//	VNetSolutionArgs{...}
type VNetSolutionInput interface {
	pulumi.Input

	ToVNetSolutionOutput() VNetSolutionOutput
	ToVNetSolutionOutputWithContext(context.Context) VNetSolutionOutput
}

// The VNet solution for linker
type VNetSolutionArgs struct {
	// Type of VNet solution.
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (VNetSolutionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VNetSolution)(nil)).Elem()
}

func (i VNetSolutionArgs) ToVNetSolutionOutput() VNetSolutionOutput {
	return i.ToVNetSolutionOutputWithContext(context.Background())
}

func (i VNetSolutionArgs) ToVNetSolutionOutputWithContext(ctx context.Context) VNetSolutionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VNetSolutionOutput)
}

func (i VNetSolutionArgs) ToVNetSolutionPtrOutput() VNetSolutionPtrOutput {
	return i.ToVNetSolutionPtrOutputWithContext(context.Background())
}

func (i VNetSolutionArgs) ToVNetSolutionPtrOutputWithContext(ctx context.Context) VNetSolutionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VNetSolutionOutput).ToVNetSolutionPtrOutputWithContext(ctx)
}

// VNetSolutionPtrInput is an input type that accepts VNetSolutionArgs, VNetSolutionPtr and VNetSolutionPtrOutput values.
// You can construct a concrete instance of `VNetSolutionPtrInput` via:
//
//	        VNetSolutionArgs{...}
//
//	or:
//
//	        nil
type VNetSolutionPtrInput interface {
	pulumi.Input

	ToVNetSolutionPtrOutput() VNetSolutionPtrOutput
	ToVNetSolutionPtrOutputWithContext(context.Context) VNetSolutionPtrOutput
}

type vnetSolutionPtrType VNetSolutionArgs

func VNetSolutionPtr(v *VNetSolutionArgs) VNetSolutionPtrInput {
	return (*vnetSolutionPtrType)(v)
}

func (*vnetSolutionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VNetSolution)(nil)).Elem()
}

func (i *vnetSolutionPtrType) ToVNetSolutionPtrOutput() VNetSolutionPtrOutput {
	return i.ToVNetSolutionPtrOutputWithContext(context.Background())
}

func (i *vnetSolutionPtrType) ToVNetSolutionPtrOutputWithContext(ctx context.Context) VNetSolutionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VNetSolutionPtrOutput)
}

// The VNet solution for linker
type VNetSolutionOutput struct{ *pulumi.OutputState }

func (VNetSolutionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VNetSolution)(nil)).Elem()
}

func (o VNetSolutionOutput) ToVNetSolutionOutput() VNetSolutionOutput {
	return o
}

func (o VNetSolutionOutput) ToVNetSolutionOutputWithContext(ctx context.Context) VNetSolutionOutput {
	return o
}

func (o VNetSolutionOutput) ToVNetSolutionPtrOutput() VNetSolutionPtrOutput {
	return o.ToVNetSolutionPtrOutputWithContext(context.Background())
}

func (o VNetSolutionOutput) ToVNetSolutionPtrOutputWithContext(ctx context.Context) VNetSolutionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VNetSolution) *VNetSolution {
		return &v
	}).(VNetSolutionPtrOutput)
}

// Type of VNet solution.
func (o VNetSolutionOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VNetSolution) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type VNetSolutionPtrOutput struct{ *pulumi.OutputState }

func (VNetSolutionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VNetSolution)(nil)).Elem()
}

func (o VNetSolutionPtrOutput) ToVNetSolutionPtrOutput() VNetSolutionPtrOutput {
	return o
}

func (o VNetSolutionPtrOutput) ToVNetSolutionPtrOutputWithContext(ctx context.Context) VNetSolutionPtrOutput {
	return o
}

func (o VNetSolutionPtrOutput) Elem() VNetSolutionOutput {
	return o.ApplyT(func(v *VNetSolution) VNetSolution {
		if v != nil {
			return *v
		}
		var ret VNetSolution
		return ret
	}).(VNetSolutionOutput)
}

// Type of VNet solution.
func (o VNetSolutionPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VNetSolution) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

// The VNet solution for linker
type VNetSolutionResponse struct {
	// Type of VNet solution.
	Type *string `pulumi:"type"`
}

// The VNet solution for linker
type VNetSolutionResponseOutput struct{ *pulumi.OutputState }

func (VNetSolutionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VNetSolutionResponse)(nil)).Elem()
}

func (o VNetSolutionResponseOutput) ToVNetSolutionResponseOutput() VNetSolutionResponseOutput {
	return o
}

func (o VNetSolutionResponseOutput) ToVNetSolutionResponseOutputWithContext(ctx context.Context) VNetSolutionResponseOutput {
	return o
}

// Type of VNet solution.
func (o VNetSolutionResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VNetSolutionResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type VNetSolutionResponsePtrOutput struct{ *pulumi.OutputState }

func (VNetSolutionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VNetSolutionResponse)(nil)).Elem()
}

func (o VNetSolutionResponsePtrOutput) ToVNetSolutionResponsePtrOutput() VNetSolutionResponsePtrOutput {
	return o
}

func (o VNetSolutionResponsePtrOutput) ToVNetSolutionResponsePtrOutputWithContext(ctx context.Context) VNetSolutionResponsePtrOutput {
	return o
}

func (o VNetSolutionResponsePtrOutput) Elem() VNetSolutionResponseOutput {
	return o.ApplyT(func(v *VNetSolutionResponse) VNetSolutionResponse {
		if v != nil {
			return *v
		}
		var ret VNetSolutionResponse
		return ret
	}).(VNetSolutionResponseOutput)
}

// Type of VNet solution.
func (o VNetSolutionResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VNetSolutionResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(SecretStoreOutput{})
	pulumi.RegisterOutputType(SecretStorePtrOutput{})
	pulumi.RegisterOutputType(SecretStoreResponseOutput{})
	pulumi.RegisterOutputType(SecretStoreResponsePtrOutput{})
	pulumi.RegisterOutputType(SourceConfigurationResponseOutput{})
	pulumi.RegisterOutputType(SourceConfigurationResponseArrayOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(VNetSolutionOutput{})
	pulumi.RegisterOutputType(VNetSolutionPtrOutput{})
	pulumi.RegisterOutputType(VNetSolutionResponseOutput{})
	pulumi.RegisterOutputType(VNetSolutionResponsePtrOutput{})
}
