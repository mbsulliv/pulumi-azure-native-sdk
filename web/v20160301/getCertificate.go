// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20160301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get a certificate.
func LookupCertificate(ctx *pulumi.Context, args *LookupCertificateArgs, opts ...pulumi.InvokeOption) (*LookupCertificateResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupCertificateResult
	err := ctx.Invoke("azure-native:web/v20160301:getCertificate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCertificateArgs struct {
	// Name of the certificate.
	Name string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// SSL certificate for an app.
type LookupCertificateResult struct {
	// Raw bytes of .cer file
	CerBlob string `pulumi:"cerBlob"`
	// Certificate expiration date.
	ExpirationDate string `pulumi:"expirationDate"`
	// Friendly name of the certificate.
	FriendlyName string `pulumi:"friendlyName"`
	// Region of the certificate.
	GeoRegion string `pulumi:"geoRegion"`
	// Host names the certificate applies to.
	HostNames []string `pulumi:"hostNames"`
	// Specification for the App Service Environment to use for the certificate.
	HostingEnvironmentProfile HostingEnvironmentProfileResponse `pulumi:"hostingEnvironmentProfile"`
	// Resource Id.
	Id string `pulumi:"id"`
	// Certificate issue Date.
	IssueDate string `pulumi:"issueDate"`
	// Certificate issuer.
	Issuer string `pulumi:"issuer"`
	// Key Vault Csm resource Id.
	KeyVaultId *string `pulumi:"keyVaultId"`
	// Key Vault secret name.
	KeyVaultSecretName *string `pulumi:"keyVaultSecretName"`
	// Status of the Key Vault secret.
	KeyVaultSecretStatus string `pulumi:"keyVaultSecretStatus"`
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Resource Location.
	Location string `pulumi:"location"`
	// Resource Name.
	Name string `pulumi:"name"`
	// Pfx blob.
	PfxBlob *string `pulumi:"pfxBlob"`
	// Public key hash.
	PublicKeyHash string `pulumi:"publicKeyHash"`
	// Self link.
	SelfLink string `pulumi:"selfLink"`
	// Resource ID of the associated App Service plan, formatted as: "/subscriptions/{subscriptionID}/resourceGroups/{groupName}/providers/Microsoft.Web/serverfarms/{appServicePlanName}".
	ServerFarmId *string `pulumi:"serverFarmId"`
	// App name.
	SiteName string `pulumi:"siteName"`
	// Subject name of the certificate.
	SubjectName string `pulumi:"subjectName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Certificate thumbprint.
	Thumbprint string `pulumi:"thumbprint"`
	// Resource type.
	Type string `pulumi:"type"`
	// Is the certificate valid?.
	Valid bool `pulumi:"valid"`
}

func LookupCertificateOutput(ctx *pulumi.Context, args LookupCertificateOutputArgs, opts ...pulumi.InvokeOption) LookupCertificateResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCertificateResult, error) {
			args := v.(LookupCertificateArgs)
			r, err := LookupCertificate(ctx, &args, opts...)
			var s LookupCertificateResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCertificateResultOutput)
}

type LookupCertificateOutputArgs struct {
	// Name of the certificate.
	Name pulumi.StringInput `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupCertificateOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCertificateArgs)(nil)).Elem()
}

// SSL certificate for an app.
type LookupCertificateResultOutput struct{ *pulumi.OutputState }

func (LookupCertificateResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCertificateResult)(nil)).Elem()
}

func (o LookupCertificateResultOutput) ToLookupCertificateResultOutput() LookupCertificateResultOutput {
	return o
}

func (o LookupCertificateResultOutput) ToLookupCertificateResultOutputWithContext(ctx context.Context) LookupCertificateResultOutput {
	return o
}

func (o LookupCertificateResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupCertificateResult] {
	return pulumix.Output[LookupCertificateResult]{
		OutputState: o.OutputState,
	}
}

// Raw bytes of .cer file
func (o LookupCertificateResultOutput) CerBlob() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.CerBlob }).(pulumi.StringOutput)
}

// Certificate expiration date.
func (o LookupCertificateResultOutput) ExpirationDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.ExpirationDate }).(pulumi.StringOutput)
}

// Friendly name of the certificate.
func (o LookupCertificateResultOutput) FriendlyName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.FriendlyName }).(pulumi.StringOutput)
}

// Region of the certificate.
func (o LookupCertificateResultOutput) GeoRegion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.GeoRegion }).(pulumi.StringOutput)
}

// Host names the certificate applies to.
func (o LookupCertificateResultOutput) HostNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupCertificateResult) []string { return v.HostNames }).(pulumi.StringArrayOutput)
}

// Specification for the App Service Environment to use for the certificate.
func (o LookupCertificateResultOutput) HostingEnvironmentProfile() HostingEnvironmentProfileResponseOutput {
	return o.ApplyT(func(v LookupCertificateResult) HostingEnvironmentProfileResponse { return v.HostingEnvironmentProfile }).(HostingEnvironmentProfileResponseOutput)
}

// Resource Id.
func (o LookupCertificateResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.Id }).(pulumi.StringOutput)
}

// Certificate issue Date.
func (o LookupCertificateResultOutput) IssueDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.IssueDate }).(pulumi.StringOutput)
}

// Certificate issuer.
func (o LookupCertificateResultOutput) Issuer() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.Issuer }).(pulumi.StringOutput)
}

// Key Vault Csm resource Id.
func (o LookupCertificateResultOutput) KeyVaultId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCertificateResult) *string { return v.KeyVaultId }).(pulumi.StringPtrOutput)
}

// Key Vault secret name.
func (o LookupCertificateResultOutput) KeyVaultSecretName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCertificateResult) *string { return v.KeyVaultSecretName }).(pulumi.StringPtrOutput)
}

// Status of the Key Vault secret.
func (o LookupCertificateResultOutput) KeyVaultSecretStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.KeyVaultSecretStatus }).(pulumi.StringOutput)
}

// Kind of resource.
func (o LookupCertificateResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCertificateResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// Resource Location.
func (o LookupCertificateResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.Location }).(pulumi.StringOutput)
}

// Resource Name.
func (o LookupCertificateResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.Name }).(pulumi.StringOutput)
}

// Pfx blob.
func (o LookupCertificateResultOutput) PfxBlob() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCertificateResult) *string { return v.PfxBlob }).(pulumi.StringPtrOutput)
}

// Public key hash.
func (o LookupCertificateResultOutput) PublicKeyHash() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.PublicKeyHash }).(pulumi.StringOutput)
}

// Self link.
func (o LookupCertificateResultOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.SelfLink }).(pulumi.StringOutput)
}

// Resource ID of the associated App Service plan, formatted as: "/subscriptions/{subscriptionID}/resourceGroups/{groupName}/providers/Microsoft.Web/serverfarms/{appServicePlanName}".
func (o LookupCertificateResultOutput) ServerFarmId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCertificateResult) *string { return v.ServerFarmId }).(pulumi.StringPtrOutput)
}

// App name.
func (o LookupCertificateResultOutput) SiteName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.SiteName }).(pulumi.StringOutput)
}

// Subject name of the certificate.
func (o LookupCertificateResultOutput) SubjectName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.SubjectName }).(pulumi.StringOutput)
}

// Resource tags.
func (o LookupCertificateResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupCertificateResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Certificate thumbprint.
func (o LookupCertificateResultOutput) Thumbprint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.Thumbprint }).(pulumi.StringOutput)
}

// Resource type.
func (o LookupCertificateResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.Type }).(pulumi.StringOutput)
}

// Is the certificate valid?.
func (o LookupCertificateResultOutput) Valid() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupCertificateResult) bool { return v.Valid }).(pulumi.BoolOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCertificateResultOutput{})
}
