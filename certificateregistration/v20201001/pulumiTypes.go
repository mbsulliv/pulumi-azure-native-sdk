// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20201001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

var _ = utilities.GetEnvOrDefault

// Key Vault container for a certificate that is purchased through Azure.
type AppServiceCertificate struct {
	// Key Vault resource Id.
	KeyVaultId *string `pulumi:"keyVaultId"`
	// Key Vault secret name.
	KeyVaultSecretName *string `pulumi:"keyVaultSecretName"`
}

// AppServiceCertificateInput is an input type that accepts AppServiceCertificateArgs and AppServiceCertificateOutput values.
// You can construct a concrete instance of `AppServiceCertificateInput` via:
//
//	AppServiceCertificateArgs{...}
type AppServiceCertificateInput interface {
	pulumi.Input

	ToAppServiceCertificateOutput() AppServiceCertificateOutput
	ToAppServiceCertificateOutputWithContext(context.Context) AppServiceCertificateOutput
}

// Key Vault container for a certificate that is purchased through Azure.
type AppServiceCertificateArgs struct {
	// Key Vault resource Id.
	KeyVaultId pulumi.StringPtrInput `pulumi:"keyVaultId"`
	// Key Vault secret name.
	KeyVaultSecretName pulumi.StringPtrInput `pulumi:"keyVaultSecretName"`
}

func (AppServiceCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AppServiceCertificate)(nil)).Elem()
}

func (i AppServiceCertificateArgs) ToAppServiceCertificateOutput() AppServiceCertificateOutput {
	return i.ToAppServiceCertificateOutputWithContext(context.Background())
}

func (i AppServiceCertificateArgs) ToAppServiceCertificateOutputWithContext(ctx context.Context) AppServiceCertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppServiceCertificateOutput)
}

func (i AppServiceCertificateArgs) ToOutput(ctx context.Context) pulumix.Output[AppServiceCertificate] {
	return pulumix.Output[AppServiceCertificate]{
		OutputState: i.ToAppServiceCertificateOutputWithContext(ctx).OutputState,
	}
}

// AppServiceCertificateMapInput is an input type that accepts AppServiceCertificateMap and AppServiceCertificateMapOutput values.
// You can construct a concrete instance of `AppServiceCertificateMapInput` via:
//
//	AppServiceCertificateMap{ "key": AppServiceCertificateArgs{...} }
type AppServiceCertificateMapInput interface {
	pulumi.Input

	ToAppServiceCertificateMapOutput() AppServiceCertificateMapOutput
	ToAppServiceCertificateMapOutputWithContext(context.Context) AppServiceCertificateMapOutput
}

type AppServiceCertificateMap map[string]AppServiceCertificateInput

func (AppServiceCertificateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]AppServiceCertificate)(nil)).Elem()
}

func (i AppServiceCertificateMap) ToAppServiceCertificateMapOutput() AppServiceCertificateMapOutput {
	return i.ToAppServiceCertificateMapOutputWithContext(context.Background())
}

func (i AppServiceCertificateMap) ToAppServiceCertificateMapOutputWithContext(ctx context.Context) AppServiceCertificateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppServiceCertificateMapOutput)
}

func (i AppServiceCertificateMap) ToOutput(ctx context.Context) pulumix.Output[map[string]AppServiceCertificate] {
	return pulumix.Output[map[string]AppServiceCertificate]{
		OutputState: i.ToAppServiceCertificateMapOutputWithContext(ctx).OutputState,
	}
}

// Key Vault container for a certificate that is purchased through Azure.
type AppServiceCertificateOutput struct{ *pulumi.OutputState }

func (AppServiceCertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AppServiceCertificate)(nil)).Elem()
}

func (o AppServiceCertificateOutput) ToAppServiceCertificateOutput() AppServiceCertificateOutput {
	return o
}

func (o AppServiceCertificateOutput) ToAppServiceCertificateOutputWithContext(ctx context.Context) AppServiceCertificateOutput {
	return o
}

func (o AppServiceCertificateOutput) ToOutput(ctx context.Context) pulumix.Output[AppServiceCertificate] {
	return pulumix.Output[AppServiceCertificate]{
		OutputState: o.OutputState,
	}
}

// Key Vault resource Id.
func (o AppServiceCertificateOutput) KeyVaultId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AppServiceCertificate) *string { return v.KeyVaultId }).(pulumi.StringPtrOutput)
}

// Key Vault secret name.
func (o AppServiceCertificateOutput) KeyVaultSecretName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AppServiceCertificate) *string { return v.KeyVaultSecretName }).(pulumi.StringPtrOutput)
}

type AppServiceCertificateMapOutput struct{ *pulumi.OutputState }

func (AppServiceCertificateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]AppServiceCertificate)(nil)).Elem()
}

func (o AppServiceCertificateMapOutput) ToAppServiceCertificateMapOutput() AppServiceCertificateMapOutput {
	return o
}

func (o AppServiceCertificateMapOutput) ToAppServiceCertificateMapOutputWithContext(ctx context.Context) AppServiceCertificateMapOutput {
	return o
}

func (o AppServiceCertificateMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]AppServiceCertificate] {
	return pulumix.Output[map[string]AppServiceCertificate]{
		OutputState: o.OutputState,
	}
}

func (o AppServiceCertificateMapOutput) MapIndex(k pulumi.StringInput) AppServiceCertificateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) AppServiceCertificate {
		return vs[0].(map[string]AppServiceCertificate)[vs[1].(string)]
	}).(AppServiceCertificateOutput)
}

// Key Vault container for a certificate that is purchased through Azure.
type AppServiceCertificateResponse struct {
	// Key Vault resource Id.
	KeyVaultId *string `pulumi:"keyVaultId"`
	// Key Vault secret name.
	KeyVaultSecretName *string `pulumi:"keyVaultSecretName"`
	// Status of the Key Vault secret.
	ProvisioningState string `pulumi:"provisioningState"`
}

// Key Vault container for a certificate that is purchased through Azure.
type AppServiceCertificateResponseOutput struct{ *pulumi.OutputState }

func (AppServiceCertificateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AppServiceCertificateResponse)(nil)).Elem()
}

func (o AppServiceCertificateResponseOutput) ToAppServiceCertificateResponseOutput() AppServiceCertificateResponseOutput {
	return o
}

func (o AppServiceCertificateResponseOutput) ToAppServiceCertificateResponseOutputWithContext(ctx context.Context) AppServiceCertificateResponseOutput {
	return o
}

func (o AppServiceCertificateResponseOutput) ToOutput(ctx context.Context) pulumix.Output[AppServiceCertificateResponse] {
	return pulumix.Output[AppServiceCertificateResponse]{
		OutputState: o.OutputState,
	}
}

// Key Vault resource Id.
func (o AppServiceCertificateResponseOutput) KeyVaultId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AppServiceCertificateResponse) *string { return v.KeyVaultId }).(pulumi.StringPtrOutput)
}

// Key Vault secret name.
func (o AppServiceCertificateResponseOutput) KeyVaultSecretName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AppServiceCertificateResponse) *string { return v.KeyVaultSecretName }).(pulumi.StringPtrOutput)
}

// Status of the Key Vault secret.
func (o AppServiceCertificateResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v AppServiceCertificateResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

type AppServiceCertificateResponseMapOutput struct{ *pulumi.OutputState }

func (AppServiceCertificateResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]AppServiceCertificateResponse)(nil)).Elem()
}

func (o AppServiceCertificateResponseMapOutput) ToAppServiceCertificateResponseMapOutput() AppServiceCertificateResponseMapOutput {
	return o
}

func (o AppServiceCertificateResponseMapOutput) ToAppServiceCertificateResponseMapOutputWithContext(ctx context.Context) AppServiceCertificateResponseMapOutput {
	return o
}

func (o AppServiceCertificateResponseMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]AppServiceCertificateResponse] {
	return pulumix.Output[map[string]AppServiceCertificateResponse]{
		OutputState: o.OutputState,
	}
}

func (o AppServiceCertificateResponseMapOutput) MapIndex(k pulumi.StringInput) AppServiceCertificateResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) AppServiceCertificateResponse {
		return vs[0].(map[string]AppServiceCertificateResponse)[vs[1].(string)]
	}).(AppServiceCertificateResponseOutput)
}

// SSL certificate details.
type CertificateDetailsResponse struct {
	// Certificate Issuer.
	Issuer string `pulumi:"issuer"`
	// Date Certificate is valid to.
	NotAfter string `pulumi:"notAfter"`
	// Date Certificate is valid from.
	NotBefore string `pulumi:"notBefore"`
	// Raw certificate data.
	RawData string `pulumi:"rawData"`
	// Certificate Serial Number.
	SerialNumber string `pulumi:"serialNumber"`
	// Certificate Signature algorithm.
	SignatureAlgorithm string `pulumi:"signatureAlgorithm"`
	// Certificate Subject.
	Subject string `pulumi:"subject"`
	// Certificate Thumbprint.
	Thumbprint string `pulumi:"thumbprint"`
	// Certificate Version.
	Version int `pulumi:"version"`
}

// SSL certificate details.
type CertificateDetailsResponseOutput struct{ *pulumi.OutputState }

func (CertificateDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CertificateDetailsResponse)(nil)).Elem()
}

func (o CertificateDetailsResponseOutput) ToCertificateDetailsResponseOutput() CertificateDetailsResponseOutput {
	return o
}

func (o CertificateDetailsResponseOutput) ToCertificateDetailsResponseOutputWithContext(ctx context.Context) CertificateDetailsResponseOutput {
	return o
}

func (o CertificateDetailsResponseOutput) ToOutput(ctx context.Context) pulumix.Output[CertificateDetailsResponse] {
	return pulumix.Output[CertificateDetailsResponse]{
		OutputState: o.OutputState,
	}
}

// Certificate Issuer.
func (o CertificateDetailsResponseOutput) Issuer() pulumi.StringOutput {
	return o.ApplyT(func(v CertificateDetailsResponse) string { return v.Issuer }).(pulumi.StringOutput)
}

// Date Certificate is valid to.
func (o CertificateDetailsResponseOutput) NotAfter() pulumi.StringOutput {
	return o.ApplyT(func(v CertificateDetailsResponse) string { return v.NotAfter }).(pulumi.StringOutput)
}

// Date Certificate is valid from.
func (o CertificateDetailsResponseOutput) NotBefore() pulumi.StringOutput {
	return o.ApplyT(func(v CertificateDetailsResponse) string { return v.NotBefore }).(pulumi.StringOutput)
}

// Raw certificate data.
func (o CertificateDetailsResponseOutput) RawData() pulumi.StringOutput {
	return o.ApplyT(func(v CertificateDetailsResponse) string { return v.RawData }).(pulumi.StringOutput)
}

// Certificate Serial Number.
func (o CertificateDetailsResponseOutput) SerialNumber() pulumi.StringOutput {
	return o.ApplyT(func(v CertificateDetailsResponse) string { return v.SerialNumber }).(pulumi.StringOutput)
}

// Certificate Signature algorithm.
func (o CertificateDetailsResponseOutput) SignatureAlgorithm() pulumi.StringOutput {
	return o.ApplyT(func(v CertificateDetailsResponse) string { return v.SignatureAlgorithm }).(pulumi.StringOutput)
}

// Certificate Subject.
func (o CertificateDetailsResponseOutput) Subject() pulumi.StringOutput {
	return o.ApplyT(func(v CertificateDetailsResponse) string { return v.Subject }).(pulumi.StringOutput)
}

// Certificate Thumbprint.
func (o CertificateDetailsResponseOutput) Thumbprint() pulumi.StringOutput {
	return o.ApplyT(func(v CertificateDetailsResponse) string { return v.Thumbprint }).(pulumi.StringOutput)
}

// Certificate Version.
func (o CertificateDetailsResponseOutput) Version() pulumi.IntOutput {
	return o.ApplyT(func(v CertificateDetailsResponse) int { return v.Version }).(pulumi.IntOutput)
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

func init() {
	pulumi.RegisterOutputType(AppServiceCertificateOutput{})
	pulumi.RegisterOutputType(AppServiceCertificateMapOutput{})
	pulumi.RegisterOutputType(AppServiceCertificateResponseOutput{})
	pulumi.RegisterOutputType(AppServiceCertificateResponseMapOutput{})
	pulumi.RegisterOutputType(CertificateDetailsResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
}
