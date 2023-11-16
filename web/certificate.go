// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package web

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// SSL certificate for an app.
// Azure REST API version: 2022-09-01. Prior API version in Azure Native 1.x: 2020-12-01.
//
// Other available API versions: 2015-08-01, 2016-03-01, 2020-10-01, 2023-01-01.
type Certificate struct {
	pulumi.CustomResourceState

	// CNAME of the certificate to be issued via free certificate
	CanonicalName pulumi.StringPtrOutput `pulumi:"canonicalName"`
	// Raw bytes of .cer file
	CerBlob pulumi.StringOutput `pulumi:"cerBlob"`
	// Method of domain validation for free cert
	DomainValidationMethod pulumi.StringPtrOutput `pulumi:"domainValidationMethod"`
	// Certificate expiration date.
	ExpirationDate pulumi.StringOutput `pulumi:"expirationDate"`
	// Friendly name of the certificate.
	FriendlyName pulumi.StringOutput `pulumi:"friendlyName"`
	// Host names the certificate applies to.
	HostNames pulumi.StringArrayOutput `pulumi:"hostNames"`
	// Specification for the App Service Environment to use for the certificate.
	HostingEnvironmentProfile HostingEnvironmentProfileResponseOutput `pulumi:"hostingEnvironmentProfile"`
	// Certificate issue Date.
	IssueDate pulumi.StringOutput `pulumi:"issueDate"`
	// Certificate issuer.
	Issuer pulumi.StringOutput `pulumi:"issuer"`
	// Key Vault Csm resource Id.
	KeyVaultId pulumi.StringPtrOutput `pulumi:"keyVaultId"`
	// Key Vault secret name.
	KeyVaultSecretName pulumi.StringPtrOutput `pulumi:"keyVaultSecretName"`
	// Status of the Key Vault secret.
	KeyVaultSecretStatus pulumi.StringOutput `pulumi:"keyVaultSecretStatus"`
	// Kind of resource.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Resource Location.
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource Name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Pfx blob.
	PfxBlob pulumi.StringPtrOutput `pulumi:"pfxBlob"`
	// Public key hash.
	PublicKeyHash pulumi.StringOutput `pulumi:"publicKeyHash"`
	// Self link.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// Resource ID of the associated App Service plan, formatted as: "/subscriptions/{subscriptionID}/resourceGroups/{groupName}/providers/Microsoft.Web/serverfarms/{appServicePlanName}".
	ServerFarmId pulumi.StringPtrOutput `pulumi:"serverFarmId"`
	// App name.
	SiteName pulumi.StringOutput `pulumi:"siteName"`
	// Subject name of the certificate.
	SubjectName pulumi.StringOutput `pulumi:"subjectName"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Certificate thumbprint.
	Thumbprint pulumi.StringOutput `pulumi:"thumbprint"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// Is the certificate valid?.
	Valid pulumi.BoolOutput `pulumi:"valid"`
}

// NewCertificate registers a new resource with the given unique name, arguments, and options.
func NewCertificate(ctx *pulumi.Context,
	name string, args *CertificateArgs, opts ...pulumi.ResourceOption) (*Certificate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:web/v20150801:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:web/v20160301:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:web/v20181101:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210301:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220301:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220901:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:web/v20230101:Certificate"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Certificate
	err := ctx.RegisterResource("azure-native:web:Certificate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCertificate gets an existing Certificate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCertificate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CertificateState, opts ...pulumi.ResourceOption) (*Certificate, error) {
	var resource Certificate
	err := ctx.ReadResource("azure-native:web:Certificate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Certificate resources.
type certificateState struct {
}

type CertificateState struct {
}

func (CertificateState) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateState)(nil)).Elem()
}

type certificateArgs struct {
	// CNAME of the certificate to be issued via free certificate
	CanonicalName *string `pulumi:"canonicalName"`
	// Method of domain validation for free cert
	DomainValidationMethod *string `pulumi:"domainValidationMethod"`
	// Host names the certificate applies to.
	HostNames []string `pulumi:"hostNames"`
	// Key Vault Csm resource Id.
	KeyVaultId *string `pulumi:"keyVaultId"`
	// Key Vault secret name.
	KeyVaultSecretName *string `pulumi:"keyVaultSecretName"`
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Resource Location.
	Location *string `pulumi:"location"`
	// Name of the certificate.
	Name *string `pulumi:"name"`
	// Certificate password.
	Password *string `pulumi:"password"`
	// Pfx blob.
	PfxBlob *string `pulumi:"pfxBlob"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource ID of the associated App Service plan, formatted as: "/subscriptions/{subscriptionID}/resourceGroups/{groupName}/providers/Microsoft.Web/serverfarms/{appServicePlanName}".
	ServerFarmId *string `pulumi:"serverFarmId"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Certificate resource.
type CertificateArgs struct {
	// CNAME of the certificate to be issued via free certificate
	CanonicalName pulumi.StringPtrInput
	// Method of domain validation for free cert
	DomainValidationMethod pulumi.StringPtrInput
	// Host names the certificate applies to.
	HostNames pulumi.StringArrayInput
	// Key Vault Csm resource Id.
	KeyVaultId pulumi.StringPtrInput
	// Key Vault secret name.
	KeyVaultSecretName pulumi.StringPtrInput
	// Kind of resource.
	Kind pulumi.StringPtrInput
	// Resource Location.
	Location pulumi.StringPtrInput
	// Name of the certificate.
	Name pulumi.StringPtrInput
	// Certificate password.
	Password pulumi.StringPtrInput
	// Pfx blob.
	PfxBlob pulumi.StringPtrInput
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput
	// Resource ID of the associated App Service plan, formatted as: "/subscriptions/{subscriptionID}/resourceGroups/{groupName}/providers/Microsoft.Web/serverfarms/{appServicePlanName}".
	ServerFarmId pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (CertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateArgs)(nil)).Elem()
}

type CertificateInput interface {
	pulumi.Input

	ToCertificateOutput() CertificateOutput
	ToCertificateOutputWithContext(ctx context.Context) CertificateOutput
}

func (*Certificate) ElementType() reflect.Type {
	return reflect.TypeOf((**Certificate)(nil)).Elem()
}

func (i *Certificate) ToCertificateOutput() CertificateOutput {
	return i.ToCertificateOutputWithContext(context.Background())
}

func (i *Certificate) ToCertificateOutputWithContext(ctx context.Context) CertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateOutput)
}

type CertificateOutput struct{ *pulumi.OutputState }

func (CertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Certificate)(nil)).Elem()
}

func (o CertificateOutput) ToCertificateOutput() CertificateOutput {
	return o
}

func (o CertificateOutput) ToCertificateOutputWithContext(ctx context.Context) CertificateOutput {
	return o
}

// CNAME of the certificate to be issued via free certificate
func (o CertificateOutput) CanonicalName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringPtrOutput { return v.CanonicalName }).(pulumi.StringPtrOutput)
}

// Raw bytes of .cer file
func (o CertificateOutput) CerBlob() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.CerBlob }).(pulumi.StringOutput)
}

// Method of domain validation for free cert
func (o CertificateOutput) DomainValidationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringPtrOutput { return v.DomainValidationMethod }).(pulumi.StringPtrOutput)
}

// Certificate expiration date.
func (o CertificateOutput) ExpirationDate() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.ExpirationDate }).(pulumi.StringOutput)
}

// Friendly name of the certificate.
func (o CertificateOutput) FriendlyName() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.FriendlyName }).(pulumi.StringOutput)
}

// Host names the certificate applies to.
func (o CertificateOutput) HostNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringArrayOutput { return v.HostNames }).(pulumi.StringArrayOutput)
}

// Specification for the App Service Environment to use for the certificate.
func (o CertificateOutput) HostingEnvironmentProfile() HostingEnvironmentProfileResponseOutput {
	return o.ApplyT(func(v *Certificate) HostingEnvironmentProfileResponseOutput { return v.HostingEnvironmentProfile }).(HostingEnvironmentProfileResponseOutput)
}

// Certificate issue Date.
func (o CertificateOutput) IssueDate() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.IssueDate }).(pulumi.StringOutput)
}

// Certificate issuer.
func (o CertificateOutput) Issuer() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.Issuer }).(pulumi.StringOutput)
}

// Key Vault Csm resource Id.
func (o CertificateOutput) KeyVaultId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringPtrOutput { return v.KeyVaultId }).(pulumi.StringPtrOutput)
}

// Key Vault secret name.
func (o CertificateOutput) KeyVaultSecretName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringPtrOutput { return v.KeyVaultSecretName }).(pulumi.StringPtrOutput)
}

// Status of the Key Vault secret.
func (o CertificateOutput) KeyVaultSecretStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.KeyVaultSecretStatus }).(pulumi.StringOutput)
}

// Kind of resource.
func (o CertificateOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Resource Location.
func (o CertificateOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Resource Name.
func (o CertificateOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Pfx blob.
func (o CertificateOutput) PfxBlob() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringPtrOutput { return v.PfxBlob }).(pulumi.StringPtrOutput)
}

// Public key hash.
func (o CertificateOutput) PublicKeyHash() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.PublicKeyHash }).(pulumi.StringOutput)
}

// Self link.
func (o CertificateOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.SelfLink }).(pulumi.StringOutput)
}

// Resource ID of the associated App Service plan, formatted as: "/subscriptions/{subscriptionID}/resourceGroups/{groupName}/providers/Microsoft.Web/serverfarms/{appServicePlanName}".
func (o CertificateOutput) ServerFarmId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringPtrOutput { return v.ServerFarmId }).(pulumi.StringPtrOutput)
}

// App name.
func (o CertificateOutput) SiteName() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.SiteName }).(pulumi.StringOutput)
}

// Subject name of the certificate.
func (o CertificateOutput) SubjectName() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.SubjectName }).(pulumi.StringOutput)
}

// Resource tags.
func (o CertificateOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Certificate thumbprint.
func (o CertificateOutput) Thumbprint() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.Thumbprint }).(pulumi.StringOutput)
}

// Resource type.
func (o CertificateOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Is the certificate valid?.
func (o CertificateOutput) Valid() pulumi.BoolOutput {
	return o.ApplyT(func(v *Certificate) pulumi.BoolOutput { return v.Valid }).(pulumi.BoolOutput)
}

func init() {
	pulumi.RegisterOutputType(CertificateOutput{})
}
