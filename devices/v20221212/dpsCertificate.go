// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20221212

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The X509 Certificate.
type DpsCertificate struct {
	pulumi.CustomResourceState

	// The entity tag.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The name of the certificate.
	Name pulumi.StringOutput `pulumi:"name"`
	// properties of a certificate
	Properties CertificatePropertiesResponseOutput `pulumi:"properties"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDpsCertificate registers a new resource with the given unique name, arguments, and options.
func NewDpsCertificate(ctx *pulumi.Context,
	name string, args *DpsCertificateArgs, opts ...pulumi.ResourceOption) (*DpsCertificate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProvisioningServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ProvisioningServiceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:devices:DpsCertificate"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20170821preview:DpsCertificate"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20171115:DpsCertificate"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20180122:DpsCertificate"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20200101:DpsCertificate"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20200301:DpsCertificate"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20200901preview:DpsCertificate"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20211015:DpsCertificate"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20220205:DpsCertificate"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20230301preview:DpsCertificate"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource DpsCertificate
	err := ctx.RegisterResource("azure-native:devices/v20221212:DpsCertificate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDpsCertificate gets an existing DpsCertificate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDpsCertificate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DpsCertificateState, opts ...pulumi.ResourceOption) (*DpsCertificate, error) {
	var resource DpsCertificate
	err := ctx.ReadResource("azure-native:devices/v20221212:DpsCertificate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DpsCertificate resources.
type dpsCertificateState struct {
}

type DpsCertificateState struct {
}

func (DpsCertificateState) ElementType() reflect.Type {
	return reflect.TypeOf((*dpsCertificateState)(nil)).Elem()
}

type dpsCertificateArgs struct {
	// The name of the certificate create or update.
	CertificateName *string `pulumi:"certificateName"`
	// properties of a certificate
	Properties *CertificateProperties `pulumi:"properties"`
	// The name of the provisioning service.
	ProvisioningServiceName string `pulumi:"provisioningServiceName"`
	// Resource group identifier.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a DpsCertificate resource.
type DpsCertificateArgs struct {
	// The name of the certificate create or update.
	CertificateName pulumi.StringPtrInput
	// properties of a certificate
	Properties CertificatePropertiesPtrInput
	// The name of the provisioning service.
	ProvisioningServiceName pulumi.StringInput
	// Resource group identifier.
	ResourceGroupName pulumi.StringInput
}

func (DpsCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dpsCertificateArgs)(nil)).Elem()
}

type DpsCertificateInput interface {
	pulumi.Input

	ToDpsCertificateOutput() DpsCertificateOutput
	ToDpsCertificateOutputWithContext(ctx context.Context) DpsCertificateOutput
}

func (*DpsCertificate) ElementType() reflect.Type {
	return reflect.TypeOf((**DpsCertificate)(nil)).Elem()
}

func (i *DpsCertificate) ToDpsCertificateOutput() DpsCertificateOutput {
	return i.ToDpsCertificateOutputWithContext(context.Background())
}

func (i *DpsCertificate) ToDpsCertificateOutputWithContext(ctx context.Context) DpsCertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DpsCertificateOutput)
}

type DpsCertificateOutput struct{ *pulumi.OutputState }

func (DpsCertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DpsCertificate)(nil)).Elem()
}

func (o DpsCertificateOutput) ToDpsCertificateOutput() DpsCertificateOutput {
	return o
}

func (o DpsCertificateOutput) ToDpsCertificateOutputWithContext(ctx context.Context) DpsCertificateOutput {
	return o
}

// The entity tag.
func (o DpsCertificateOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *DpsCertificate) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The name of the certificate.
func (o DpsCertificateOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DpsCertificate) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// properties of a certificate
func (o DpsCertificateOutput) Properties() CertificatePropertiesResponseOutput {
	return o.ApplyT(func(v *DpsCertificate) CertificatePropertiesResponseOutput { return v.Properties }).(CertificatePropertiesResponseOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o DpsCertificateOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *DpsCertificate) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The resource type.
func (o DpsCertificateOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DpsCertificate) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DpsCertificateOutput{})
}
