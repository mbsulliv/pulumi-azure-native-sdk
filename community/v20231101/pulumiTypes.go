// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20231101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = utilities.GetEnvOrDefault

// Details of the Community CommunityTraining Identity Configuration
type IdentityConfigurationProperties struct {
	// The name of the authentication policy registered in ADB2C for the Community Training Resource
	B2cAuthenticationPolicy *string `pulumi:"b2cAuthenticationPolicy"`
	// The name of the password reset policy registered in ADB2C for the Community Training Resource
	B2cPasswordResetPolicy *string `pulumi:"b2cPasswordResetPolicy"`
	// The clientId of the application registered in the selected identity provider for the Community Training Resource
	ClientId string `pulumi:"clientId"`
	// The client secret of the application registered in the selected identity provider for the Community Training Resource
	ClientSecret string `pulumi:"clientSecret"`
	// The custom login parameters for the Community Training Resource
	CustomLoginParameters *string `pulumi:"customLoginParameters"`
	// The domain name of the selected identity provider for the Community Training Resource
	DomainName string `pulumi:"domainName"`
	// The identity type of the Community Training Resource
	IdentityType string `pulumi:"identityType"`
	// To indicate whether the Community Training Resource has Teams enabled
	TeamsEnabled *bool `pulumi:"teamsEnabled"`
	// The tenantId of the selected identity provider for the Community Training Resource
	TenantId string `pulumi:"tenantId"`
}

// Defaults sets the appropriate defaults for IdentityConfigurationProperties
func (val *IdentityConfigurationProperties) Defaults() *IdentityConfigurationProperties {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.TeamsEnabled == nil {
		teamsEnabled_ := false
		tmp.TeamsEnabled = &teamsEnabled_
	}
	return &tmp
}

// IdentityConfigurationPropertiesInput is an input type that accepts IdentityConfigurationPropertiesArgs and IdentityConfigurationPropertiesOutput values.
// You can construct a concrete instance of `IdentityConfigurationPropertiesInput` via:
//
//	IdentityConfigurationPropertiesArgs{...}
type IdentityConfigurationPropertiesInput interface {
	pulumi.Input

	ToIdentityConfigurationPropertiesOutput() IdentityConfigurationPropertiesOutput
	ToIdentityConfigurationPropertiesOutputWithContext(context.Context) IdentityConfigurationPropertiesOutput
}

// Details of the Community CommunityTraining Identity Configuration
type IdentityConfigurationPropertiesArgs struct {
	// The name of the authentication policy registered in ADB2C for the Community Training Resource
	B2cAuthenticationPolicy pulumi.StringPtrInput `pulumi:"b2cAuthenticationPolicy"`
	// The name of the password reset policy registered in ADB2C for the Community Training Resource
	B2cPasswordResetPolicy pulumi.StringPtrInput `pulumi:"b2cPasswordResetPolicy"`
	// The clientId of the application registered in the selected identity provider for the Community Training Resource
	ClientId pulumi.StringInput `pulumi:"clientId"`
	// The client secret of the application registered in the selected identity provider for the Community Training Resource
	ClientSecret pulumi.StringInput `pulumi:"clientSecret"`
	// The custom login parameters for the Community Training Resource
	CustomLoginParameters pulumi.StringPtrInput `pulumi:"customLoginParameters"`
	// The domain name of the selected identity provider for the Community Training Resource
	DomainName pulumi.StringInput `pulumi:"domainName"`
	// The identity type of the Community Training Resource
	IdentityType pulumi.StringInput `pulumi:"identityType"`
	// To indicate whether the Community Training Resource has Teams enabled
	TeamsEnabled pulumi.BoolPtrInput `pulumi:"teamsEnabled"`
	// The tenantId of the selected identity provider for the Community Training Resource
	TenantId pulumi.StringInput `pulumi:"tenantId"`
}

// Defaults sets the appropriate defaults for IdentityConfigurationPropertiesArgs
func (val *IdentityConfigurationPropertiesArgs) Defaults() *IdentityConfigurationPropertiesArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.TeamsEnabled == nil {
		tmp.TeamsEnabled = pulumi.BoolPtr(false)
	}
	return &tmp
}
func (IdentityConfigurationPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityConfigurationProperties)(nil)).Elem()
}

func (i IdentityConfigurationPropertiesArgs) ToIdentityConfigurationPropertiesOutput() IdentityConfigurationPropertiesOutput {
	return i.ToIdentityConfigurationPropertiesOutputWithContext(context.Background())
}

func (i IdentityConfigurationPropertiesArgs) ToIdentityConfigurationPropertiesOutputWithContext(ctx context.Context) IdentityConfigurationPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityConfigurationPropertiesOutput)
}

// Details of the Community CommunityTraining Identity Configuration
type IdentityConfigurationPropertiesOutput struct{ *pulumi.OutputState }

func (IdentityConfigurationPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityConfigurationProperties)(nil)).Elem()
}

func (o IdentityConfigurationPropertiesOutput) ToIdentityConfigurationPropertiesOutput() IdentityConfigurationPropertiesOutput {
	return o
}

func (o IdentityConfigurationPropertiesOutput) ToIdentityConfigurationPropertiesOutputWithContext(ctx context.Context) IdentityConfigurationPropertiesOutput {
	return o
}

// The name of the authentication policy registered in ADB2C for the Community Training Resource
func (o IdentityConfigurationPropertiesOutput) B2cAuthenticationPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentityConfigurationProperties) *string { return v.B2cAuthenticationPolicy }).(pulumi.StringPtrOutput)
}

// The name of the password reset policy registered in ADB2C for the Community Training Resource
func (o IdentityConfigurationPropertiesOutput) B2cPasswordResetPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentityConfigurationProperties) *string { return v.B2cPasswordResetPolicy }).(pulumi.StringPtrOutput)
}

// The clientId of the application registered in the selected identity provider for the Community Training Resource
func (o IdentityConfigurationPropertiesOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityConfigurationProperties) string { return v.ClientId }).(pulumi.StringOutput)
}

// The client secret of the application registered in the selected identity provider for the Community Training Resource
func (o IdentityConfigurationPropertiesOutput) ClientSecret() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityConfigurationProperties) string { return v.ClientSecret }).(pulumi.StringOutput)
}

// The custom login parameters for the Community Training Resource
func (o IdentityConfigurationPropertiesOutput) CustomLoginParameters() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentityConfigurationProperties) *string { return v.CustomLoginParameters }).(pulumi.StringPtrOutput)
}

// The domain name of the selected identity provider for the Community Training Resource
func (o IdentityConfigurationPropertiesOutput) DomainName() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityConfigurationProperties) string { return v.DomainName }).(pulumi.StringOutput)
}

// The identity type of the Community Training Resource
func (o IdentityConfigurationPropertiesOutput) IdentityType() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityConfigurationProperties) string { return v.IdentityType }).(pulumi.StringOutput)
}

// To indicate whether the Community Training Resource has Teams enabled
func (o IdentityConfigurationPropertiesOutput) TeamsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v IdentityConfigurationProperties) *bool { return v.TeamsEnabled }).(pulumi.BoolPtrOutput)
}

// The tenantId of the selected identity provider for the Community Training Resource
func (o IdentityConfigurationPropertiesOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityConfigurationProperties) string { return v.TenantId }).(pulumi.StringOutput)
}

// Details of the Community CommunityTraining Identity Configuration
type IdentityConfigurationPropertiesResponse struct {
	// The name of the authentication policy registered in ADB2C for the Community Training Resource
	B2cAuthenticationPolicy *string `pulumi:"b2cAuthenticationPolicy"`
	// The name of the password reset policy registered in ADB2C for the Community Training Resource
	B2cPasswordResetPolicy *string `pulumi:"b2cPasswordResetPolicy"`
	// The clientId of the application registered in the selected identity provider for the Community Training Resource
	ClientId string `pulumi:"clientId"`
	// The client secret of the application registered in the selected identity provider for the Community Training Resource
	ClientSecret string `pulumi:"clientSecret"`
	// The custom login parameters for the Community Training Resource
	CustomLoginParameters *string `pulumi:"customLoginParameters"`
	// The domain name of the selected identity provider for the Community Training Resource
	DomainName string `pulumi:"domainName"`
	// The identity type of the Community Training Resource
	IdentityType string `pulumi:"identityType"`
	// To indicate whether the Community Training Resource has Teams enabled
	TeamsEnabled *bool `pulumi:"teamsEnabled"`
	// The tenantId of the selected identity provider for the Community Training Resource
	TenantId string `pulumi:"tenantId"`
}

// Defaults sets the appropriate defaults for IdentityConfigurationPropertiesResponse
func (val *IdentityConfigurationPropertiesResponse) Defaults() *IdentityConfigurationPropertiesResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.TeamsEnabled == nil {
		teamsEnabled_ := false
		tmp.TeamsEnabled = &teamsEnabled_
	}
	return &tmp
}

// Details of the Community CommunityTraining Identity Configuration
type IdentityConfigurationPropertiesResponseOutput struct{ *pulumi.OutputState }

func (IdentityConfigurationPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityConfigurationPropertiesResponse)(nil)).Elem()
}

func (o IdentityConfigurationPropertiesResponseOutput) ToIdentityConfigurationPropertiesResponseOutput() IdentityConfigurationPropertiesResponseOutput {
	return o
}

func (o IdentityConfigurationPropertiesResponseOutput) ToIdentityConfigurationPropertiesResponseOutputWithContext(ctx context.Context) IdentityConfigurationPropertiesResponseOutput {
	return o
}

// The name of the authentication policy registered in ADB2C for the Community Training Resource
func (o IdentityConfigurationPropertiesResponseOutput) B2cAuthenticationPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentityConfigurationPropertiesResponse) *string { return v.B2cAuthenticationPolicy }).(pulumi.StringPtrOutput)
}

// The name of the password reset policy registered in ADB2C for the Community Training Resource
func (o IdentityConfigurationPropertiesResponseOutput) B2cPasswordResetPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentityConfigurationPropertiesResponse) *string { return v.B2cPasswordResetPolicy }).(pulumi.StringPtrOutput)
}

// The clientId of the application registered in the selected identity provider for the Community Training Resource
func (o IdentityConfigurationPropertiesResponseOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityConfigurationPropertiesResponse) string { return v.ClientId }).(pulumi.StringOutput)
}

// The client secret of the application registered in the selected identity provider for the Community Training Resource
func (o IdentityConfigurationPropertiesResponseOutput) ClientSecret() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityConfigurationPropertiesResponse) string { return v.ClientSecret }).(pulumi.StringOutput)
}

// The custom login parameters for the Community Training Resource
func (o IdentityConfigurationPropertiesResponseOutput) CustomLoginParameters() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentityConfigurationPropertiesResponse) *string { return v.CustomLoginParameters }).(pulumi.StringPtrOutput)
}

// The domain name of the selected identity provider for the Community Training Resource
func (o IdentityConfigurationPropertiesResponseOutput) DomainName() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityConfigurationPropertiesResponse) string { return v.DomainName }).(pulumi.StringOutput)
}

// The identity type of the Community Training Resource
func (o IdentityConfigurationPropertiesResponseOutput) IdentityType() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityConfigurationPropertiesResponse) string { return v.IdentityType }).(pulumi.StringOutput)
}

// To indicate whether the Community Training Resource has Teams enabled
func (o IdentityConfigurationPropertiesResponseOutput) TeamsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v IdentityConfigurationPropertiesResponse) *bool { return v.TeamsEnabled }).(pulumi.BoolPtrOutput)
}

// The tenantId of the selected identity provider for the Community Training Resource
func (o IdentityConfigurationPropertiesResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityConfigurationPropertiesResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

// The resource model definition representing SKU
type Sku struct {
	// If the SKU supports scale out/in then the capacity integer should be included. If scale out/in is not possible for the resource this may be omitted.
	Capacity *int `pulumi:"capacity"`
	// If the service has different generations of hardware, for the same SKU, then that can be captured here.
	Family *string `pulumi:"family"`
	// The name of the SKU. Ex - P3. It is typically a letter+number code
	Name string `pulumi:"name"`
	// The SKU size. When the name field is the combination of tier and some other value, this would be the standalone code.
	Size *string `pulumi:"size"`
	// This field is required to be implemented by the Resource Provider if the service has more than one tier, but is not required on a PUT.
	Tier *SkuTier `pulumi:"tier"`
}

// SkuInput is an input type that accepts SkuArgs and SkuOutput values.
// You can construct a concrete instance of `SkuInput` via:
//
//	SkuArgs{...}
type SkuInput interface {
	pulumi.Input

	ToSkuOutput() SkuOutput
	ToSkuOutputWithContext(context.Context) SkuOutput
}

// The resource model definition representing SKU
type SkuArgs struct {
	// If the SKU supports scale out/in then the capacity integer should be included. If scale out/in is not possible for the resource this may be omitted.
	Capacity pulumi.IntPtrInput `pulumi:"capacity"`
	// If the service has different generations of hardware, for the same SKU, then that can be captured here.
	Family pulumi.StringPtrInput `pulumi:"family"`
	// The name of the SKU. Ex - P3. It is typically a letter+number code
	Name pulumi.StringInput `pulumi:"name"`
	// The SKU size. When the name field is the combination of tier and some other value, this would be the standalone code.
	Size pulumi.StringPtrInput `pulumi:"size"`
	// This field is required to be implemented by the Resource Provider if the service has more than one tier, but is not required on a PUT.
	Tier SkuTierPtrInput `pulumi:"tier"`
}

func (SkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Sku)(nil)).Elem()
}

func (i SkuArgs) ToSkuOutput() SkuOutput {
	return i.ToSkuOutputWithContext(context.Background())
}

func (i SkuArgs) ToSkuOutputWithContext(ctx context.Context) SkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuOutput)
}

func (i SkuArgs) ToSkuPtrOutput() SkuPtrOutput {
	return i.ToSkuPtrOutputWithContext(context.Background())
}

func (i SkuArgs) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuOutput).ToSkuPtrOutputWithContext(ctx)
}

// SkuPtrInput is an input type that accepts SkuArgs, SkuPtr and SkuPtrOutput values.
// You can construct a concrete instance of `SkuPtrInput` via:
//
//	        SkuArgs{...}
//
//	or:
//
//	        nil
type SkuPtrInput interface {
	pulumi.Input

	ToSkuPtrOutput() SkuPtrOutput
	ToSkuPtrOutputWithContext(context.Context) SkuPtrOutput
}

type skuPtrType SkuArgs

func SkuPtr(v *SkuArgs) SkuPtrInput {
	return (*skuPtrType)(v)
}

func (*skuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Sku)(nil)).Elem()
}

func (i *skuPtrType) ToSkuPtrOutput() SkuPtrOutput {
	return i.ToSkuPtrOutputWithContext(context.Background())
}

func (i *skuPtrType) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuPtrOutput)
}

// The resource model definition representing SKU
type SkuOutput struct{ *pulumi.OutputState }

func (SkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Sku)(nil)).Elem()
}

func (o SkuOutput) ToSkuOutput() SkuOutput {
	return o
}

func (o SkuOutput) ToSkuOutputWithContext(ctx context.Context) SkuOutput {
	return o
}

func (o SkuOutput) ToSkuPtrOutput() SkuPtrOutput {
	return o.ToSkuPtrOutputWithContext(context.Background())
}

func (o SkuOutput) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Sku) *Sku {
		return &v
	}).(SkuPtrOutput)
}

// If the SKU supports scale out/in then the capacity integer should be included. If scale out/in is not possible for the resource this may be omitted.
func (o SkuOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Sku) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

// If the service has different generations of hardware, for the same SKU, then that can be captured here.
func (o SkuOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Sku) *string { return v.Family }).(pulumi.StringPtrOutput)
}

// The name of the SKU. Ex - P3. It is typically a letter+number code
func (o SkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v Sku) string { return v.Name }).(pulumi.StringOutput)
}

// The SKU size. When the name field is the combination of tier and some other value, this would be the standalone code.
func (o SkuOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Sku) *string { return v.Size }).(pulumi.StringPtrOutput)
}

// This field is required to be implemented by the Resource Provider if the service has more than one tier, but is not required on a PUT.
func (o SkuOutput) Tier() SkuTierPtrOutput {
	return o.ApplyT(func(v Sku) *SkuTier { return v.Tier }).(SkuTierPtrOutput)
}

type SkuPtrOutput struct{ *pulumi.OutputState }

func (SkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Sku)(nil)).Elem()
}

func (o SkuPtrOutput) ToSkuPtrOutput() SkuPtrOutput {
	return o
}

func (o SkuPtrOutput) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return o
}

func (o SkuPtrOutput) Elem() SkuOutput {
	return o.ApplyT(func(v *Sku) Sku {
		if v != nil {
			return *v
		}
		var ret Sku
		return ret
	}).(SkuOutput)
}

// If the SKU supports scale out/in then the capacity integer should be included. If scale out/in is not possible for the resource this may be omitted.
func (o SkuPtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Sku) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

// If the service has different generations of hardware, for the same SKU, then that can be captured here.
func (o SkuPtrOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return v.Family
	}).(pulumi.StringPtrOutput)
}

// The name of the SKU. Ex - P3. It is typically a letter+number code
func (o SkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

// The SKU size. When the name field is the combination of tier and some other value, this would be the standalone code.
func (o SkuPtrOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return v.Size
	}).(pulumi.StringPtrOutput)
}

// This field is required to be implemented by the Resource Provider if the service has more than one tier, but is not required on a PUT.
func (o SkuPtrOutput) Tier() SkuTierPtrOutput {
	return o.ApplyT(func(v *Sku) *SkuTier {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(SkuTierPtrOutput)
}

// The resource model definition representing SKU
type SkuResponse struct {
	// If the SKU supports scale out/in then the capacity integer should be included. If scale out/in is not possible for the resource this may be omitted.
	Capacity *int `pulumi:"capacity"`
	// If the service has different generations of hardware, for the same SKU, then that can be captured here.
	Family *string `pulumi:"family"`
	// The name of the SKU. Ex - P3. It is typically a letter+number code
	Name string `pulumi:"name"`
	// The SKU size. When the name field is the combination of tier and some other value, this would be the standalone code.
	Size *string `pulumi:"size"`
	// This field is required to be implemented by the Resource Provider if the service has more than one tier, but is not required on a PUT.
	Tier *string `pulumi:"tier"`
}

// The resource model definition representing SKU
type SkuResponseOutput struct{ *pulumi.OutputState }

func (SkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuResponse)(nil)).Elem()
}

func (o SkuResponseOutput) ToSkuResponseOutput() SkuResponseOutput {
	return o
}

func (o SkuResponseOutput) ToSkuResponseOutputWithContext(ctx context.Context) SkuResponseOutput {
	return o
}

// If the SKU supports scale out/in then the capacity integer should be included. If scale out/in is not possible for the resource this may be omitted.
func (o SkuResponseOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SkuResponse) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

// If the service has different generations of hardware, for the same SKU, then that can be captured here.
func (o SkuResponseOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuResponse) *string { return v.Family }).(pulumi.StringPtrOutput)
}

// The name of the SKU. Ex - P3. It is typically a letter+number code
func (o SkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

// The SKU size. When the name field is the combination of tier and some other value, this would be the standalone code.
func (o SkuResponseOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuResponse) *string { return v.Size }).(pulumi.StringPtrOutput)
}

// This field is required to be implemented by the Resource Provider if the service has more than one tier, but is not required on a PUT.
func (o SkuResponseOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuResponse) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type SkuResponsePtrOutput struct{ *pulumi.OutputState }

func (SkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuResponse)(nil)).Elem()
}

func (o SkuResponsePtrOutput) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return o
}

func (o SkuResponsePtrOutput) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return o
}

func (o SkuResponsePtrOutput) Elem() SkuResponseOutput {
	return o.ApplyT(func(v *SkuResponse) SkuResponse {
		if v != nil {
			return *v
		}
		var ret SkuResponse
		return ret
	}).(SkuResponseOutput)
}

// If the SKU supports scale out/in then the capacity integer should be included. If scale out/in is not possible for the resource this may be omitted.
func (o SkuResponsePtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

// If the service has different generations of hardware, for the same SKU, then that can be captured here.
func (o SkuResponsePtrOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Family
	}).(pulumi.StringPtrOutput)
}

// The name of the SKU. Ex - P3. It is typically a letter+number code
func (o SkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

// The SKU size. When the name field is the combination of tier and some other value, this would be the standalone code.
func (o SkuResponsePtrOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Size
	}).(pulumi.StringPtrOutput)
}

// This field is required to be implemented by the Resource Provider if the service has more than one tier, but is not required on a PUT.
func (o SkuResponsePtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Tier
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
	pulumi.RegisterOutputType(IdentityConfigurationPropertiesOutput{})
	pulumi.RegisterOutputType(IdentityConfigurationPropertiesResponseOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuPtrOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
}
