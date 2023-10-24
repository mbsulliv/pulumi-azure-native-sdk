// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apimanagement

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Certificate details.
// Azure REST API version: 2022-08-01. Prior API version in Azure Native 1.x: 2020-12-01.
//
// Other available API versions: 2016-07-07, 2016-10-10, 2022-09-01-preview, 2023-03-01-preview.
type Certificate struct {
	pulumi.CustomResourceState

	// Expiration date of the certificate. The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
	ExpirationDate pulumi.StringOutput `pulumi:"expirationDate"`
	// KeyVault location details of the certificate.
	KeyVault KeyVaultContractPropertiesResponsePtrOutput `pulumi:"keyVault"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Subject attribute of the certificate.
	Subject pulumi.StringOutput `pulumi:"subject"`
	// Thumbprint of the certificate.
	Thumbprint pulumi.StringOutput `pulumi:"thumbprint"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
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
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:apimanagement/v20160707:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20161010:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20170301:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180101:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180601preview:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20190101:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201preview:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20200601preview:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20201201:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210101preview:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210401preview:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210801:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20211201preview:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220401preview:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220801:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220901preview:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20230301preview:Certificate"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Certificate
	err := ctx.RegisterResource("azure-native:apimanagement:Certificate", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure-native:apimanagement:Certificate", name, id, state, &resource, opts...)
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
	// Identifier of the certificate entity. Must be unique in the current API Management service instance.
	CertificateId *string `pulumi:"certificateId"`
	// Base 64 encoded certificate using the application/x-pkcs12 representation.
	Data *string `pulumi:"data"`
	// KeyVault location details of the certificate.
	KeyVault *KeyVaultContractCreateProperties `pulumi:"keyVault"`
	// Password for the Certificate
	Password *string `pulumi:"password"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a Certificate resource.
type CertificateArgs struct {
	// Identifier of the certificate entity. Must be unique in the current API Management service instance.
	CertificateId pulumi.StringPtrInput
	// Base 64 encoded certificate using the application/x-pkcs12 representation.
	Data pulumi.StringPtrInput
	// KeyVault location details of the certificate.
	KeyVault KeyVaultContractCreatePropertiesPtrInput
	// Password for the Certificate
	Password pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the API Management service.
	ServiceName pulumi.StringInput
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

func (i *Certificate) ToOutput(ctx context.Context) pulumix.Output[*Certificate] {
	return pulumix.Output[*Certificate]{
		OutputState: i.ToCertificateOutputWithContext(ctx).OutputState,
	}
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

func (o CertificateOutput) ToOutput(ctx context.Context) pulumix.Output[*Certificate] {
	return pulumix.Output[*Certificate]{
		OutputState: o.OutputState,
	}
}

// Expiration date of the certificate. The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
func (o CertificateOutput) ExpirationDate() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.ExpirationDate }).(pulumi.StringOutput)
}

// KeyVault location details of the certificate.
func (o CertificateOutput) KeyVault() KeyVaultContractPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *Certificate) KeyVaultContractPropertiesResponsePtrOutput { return v.KeyVault }).(KeyVaultContractPropertiesResponsePtrOutput)
}

// The name of the resource
func (o CertificateOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Subject attribute of the certificate.
func (o CertificateOutput) Subject() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.Subject }).(pulumi.StringOutput)
}

// Thumbprint of the certificate.
func (o CertificateOutput) Thumbprint() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.Thumbprint }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o CertificateOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(CertificateOutput{})
}
