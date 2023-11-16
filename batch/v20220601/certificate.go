// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220601

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Contains information about a certificate.
type Certificate struct {
	pulumi.CustomResourceState

	// This is only returned when the certificate provisioningState is 'Failed'.
	DeleteCertificateError DeleteCertificateErrorResponseOutput `pulumi:"deleteCertificateError"`
	// The ETag of the resource, used for concurrency statements.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The format of the certificate - either Pfx or Cer. If omitted, the default is Pfx.
	Format pulumi.StringPtrOutput `pulumi:"format"`
	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The previous provisioned state of the resource
	PreviousProvisioningState               pulumi.StringOutput `pulumi:"previousProvisioningState"`
	PreviousProvisioningStateTransitionTime pulumi.StringOutput `pulumi:"previousProvisioningStateTransitionTime"`
	ProvisioningState                       pulumi.StringOutput `pulumi:"provisioningState"`
	ProvisioningStateTransitionTime         pulumi.StringOutput `pulumi:"provisioningStateTransitionTime"`
	// The public key of the certificate.
	PublicData pulumi.StringOutput `pulumi:"publicData"`
	// This must match the thumbprint from the name.
	Thumbprint pulumi.StringPtrOutput `pulumi:"thumbprint"`
	// This must match the first portion of the certificate name. Currently required to be 'SHA1'.
	ThumbprintAlgorithm pulumi.StringPtrOutput `pulumi:"thumbprintAlgorithm"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewCertificate registers a new resource with the given unique name, arguments, and options.
func NewCertificate(ctx *pulumi.Context,
	name string, args *CertificateArgs, opts ...pulumi.ResourceOption) (*Certificate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.Data == nil {
		return nil, errors.New("invalid value for required argument 'Data'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:batch/v20170901:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20181201:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20190401:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20190801:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20200301:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20200501:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20200901:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20210101:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20210601:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:batch/v20220101:Certificate"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Certificate
	err := ctx.RegisterResource("azure-native:batch/v20220601:Certificate", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure-native:batch/v20220601:Certificate", name, id, state, &resource, opts...)
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
	// The name of the Batch account.
	AccountName string `pulumi:"accountName"`
	// The identifier for the certificate. This must be made up of algorithm and thumbprint separated by a dash, and must match the certificate data in the request. For example SHA1-a3d1c5.
	CertificateName *string `pulumi:"certificateName"`
	// The maximum size is 10KB.
	Data string `pulumi:"data"`
	// The format of the certificate - either Pfx or Cer. If omitted, the default is Pfx.
	Format *CertificateFormat `pulumi:"format"`
	// This must not be specified if the certificate format is Cer.
	Password *string `pulumi:"password"`
	// The name of the resource group that contains the Batch account.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// This must match the thumbprint from the name.
	Thumbprint *string `pulumi:"thumbprint"`
	// This must match the first portion of the certificate name. Currently required to be 'SHA1'.
	ThumbprintAlgorithm *string `pulumi:"thumbprintAlgorithm"`
}

// The set of arguments for constructing a Certificate resource.
type CertificateArgs struct {
	// The name of the Batch account.
	AccountName pulumi.StringInput
	// The identifier for the certificate. This must be made up of algorithm and thumbprint separated by a dash, and must match the certificate data in the request. For example SHA1-a3d1c5.
	CertificateName pulumi.StringPtrInput
	// The maximum size is 10KB.
	Data pulumi.StringInput
	// The format of the certificate - either Pfx or Cer. If omitted, the default is Pfx.
	Format CertificateFormatPtrInput
	// This must not be specified if the certificate format is Cer.
	Password pulumi.StringPtrInput
	// The name of the resource group that contains the Batch account.
	ResourceGroupName pulumi.StringInput
	// This must match the thumbprint from the name.
	Thumbprint pulumi.StringPtrInput
	// This must match the first portion of the certificate name. Currently required to be 'SHA1'.
	ThumbprintAlgorithm pulumi.StringPtrInput
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

// This is only returned when the certificate provisioningState is 'Failed'.
func (o CertificateOutput) DeleteCertificateError() DeleteCertificateErrorResponseOutput {
	return o.ApplyT(func(v *Certificate) DeleteCertificateErrorResponseOutput { return v.DeleteCertificateError }).(DeleteCertificateErrorResponseOutput)
}

// The ETag of the resource, used for concurrency statements.
func (o CertificateOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The format of the certificate - either Pfx or Cer. If omitted, the default is Pfx.
func (o CertificateOutput) Format() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringPtrOutput { return v.Format }).(pulumi.StringPtrOutput)
}

// The name of the resource.
func (o CertificateOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The previous provisioned state of the resource
func (o CertificateOutput) PreviousProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.PreviousProvisioningState }).(pulumi.StringOutput)
}

func (o CertificateOutput) PreviousProvisioningStateTransitionTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.PreviousProvisioningStateTransitionTime }).(pulumi.StringOutput)
}

func (o CertificateOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o CertificateOutput) ProvisioningStateTransitionTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.ProvisioningStateTransitionTime }).(pulumi.StringOutput)
}

// The public key of the certificate.
func (o CertificateOutput) PublicData() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.PublicData }).(pulumi.StringOutput)
}

// This must match the thumbprint from the name.
func (o CertificateOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringPtrOutput { return v.Thumbprint }).(pulumi.StringPtrOutput)
}

// This must match the first portion of the certificate name. Currently required to be 'SHA1'.
func (o CertificateOutput) ThumbprintAlgorithm() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringPtrOutput { return v.ThumbprintAlgorithm }).(pulumi.StringPtrOutput)
}

// The type of the resource.
func (o CertificateOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(CertificateOutput{})
}
