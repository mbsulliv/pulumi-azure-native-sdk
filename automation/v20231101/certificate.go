// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20231101

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of the certificate.
type Certificate struct {
	pulumi.CustomResourceState

	// Gets the creation time.
	CreationTime pulumi.StringOutput `pulumi:"creationTime"`
	// Gets or sets the description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Gets the expiry time of the certificate.
	ExpiryTime pulumi.StringOutput `pulumi:"expiryTime"`
	// Gets the is exportable flag of the certificate.
	IsExportable pulumi.BoolOutput `pulumi:"isExportable"`
	// Gets the last modified time.
	LastModifiedTime pulumi.StringOutput `pulumi:"lastModifiedTime"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Gets the thumbprint of the certificate.
	Thumbprint pulumi.StringOutput `pulumi:"thumbprint"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewCertificate registers a new resource with the given unique name, arguments, and options.
func NewCertificate(ctx *pulumi.Context,
	name string, args *CertificateArgs, opts ...pulumi.ResourceOption) (*Certificate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AutomationAccountName == nil {
		return nil, errors.New("invalid value for required argument 'AutomationAccountName'")
	}
	if args.Base64Value == nil {
		return nil, errors.New("invalid value for required argument 'Base64Value'")
	}
	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:automation:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20151031:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20190601:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20200113preview:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20220808:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20230515preview:Certificate"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Certificate
	err := ctx.RegisterResource("azure-native:automation/v20231101:Certificate", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure-native:automation/v20231101:Certificate", name, id, state, &resource, opts...)
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
	// The name of the automation account.
	AutomationAccountName string `pulumi:"automationAccountName"`
	// Gets or sets the base64 encoded value of the certificate.
	Base64Value string `pulumi:"base64Value"`
	// The parameters supplied to the create or update certificate operation.
	CertificateName *string `pulumi:"certificateName"`
	// Gets or sets the description of the certificate.
	Description *string `pulumi:"description"`
	// Gets or sets the is exportable flag of the certificate.
	IsExportable *bool `pulumi:"isExportable"`
	// Gets or sets the name of the certificate.
	Name string `pulumi:"name"`
	// Name of an Azure Resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Gets or sets the thumbprint of the certificate.
	Thumbprint *string `pulumi:"thumbprint"`
}

// The set of arguments for constructing a Certificate resource.
type CertificateArgs struct {
	// The name of the automation account.
	AutomationAccountName pulumi.StringInput
	// Gets or sets the base64 encoded value of the certificate.
	Base64Value pulumi.StringInput
	// The parameters supplied to the create or update certificate operation.
	CertificateName pulumi.StringPtrInput
	// Gets or sets the description of the certificate.
	Description pulumi.StringPtrInput
	// Gets or sets the is exportable flag of the certificate.
	IsExportable pulumi.BoolPtrInput
	// Gets or sets the name of the certificate.
	Name pulumi.StringInput
	// Name of an Azure Resource group.
	ResourceGroupName pulumi.StringInput
	// Gets or sets the thumbprint of the certificate.
	Thumbprint pulumi.StringPtrInput
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

// Gets the creation time.
func (o CertificateOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.CreationTime }).(pulumi.StringOutput)
}

// Gets or sets the description.
func (o CertificateOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Gets the expiry time of the certificate.
func (o CertificateOutput) ExpiryTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.ExpiryTime }).(pulumi.StringOutput)
}

// Gets the is exportable flag of the certificate.
func (o CertificateOutput) IsExportable() pulumi.BoolOutput {
	return o.ApplyT(func(v *Certificate) pulumi.BoolOutput { return v.IsExportable }).(pulumi.BoolOutput)
}

// Gets the last modified time.
func (o CertificateOutput) LastModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.LastModifiedTime }).(pulumi.StringOutput)
}

// The name of the resource
func (o CertificateOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Gets the thumbprint of the certificate.
func (o CertificateOutput) Thumbprint() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.Thumbprint }).(pulumi.StringOutput)
}

// The type of the resource.
func (o CertificateOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(CertificateOutput{})
}
