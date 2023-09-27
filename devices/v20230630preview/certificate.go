// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230630preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The X509 Certificate.
type Certificate struct {
	pulumi.CustomResourceState

	// The entity tag.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The name of the certificate.
	Name pulumi.StringOutput `pulumi:"name"`
	// The description of an X509 CA Certificate.
	Properties CertificatePropertiesResponseOutput `pulumi:"properties"`
	// The resource type.
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
	if args.ResourceName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:devices:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20170701:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20180122:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20180401:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20181201preview:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20190322:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20190322preview:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20190701preview:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20191104:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20200301:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20200401:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20200615:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20200710preview:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20200801:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20200831:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20200831preview:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20210201preview:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20210303preview:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20210331:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20210701:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20210701preview:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20210702:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20210702preview:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20220430preview:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20221115preview:Certificate"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20230630:Certificate"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Certificate
	err := ctx.RegisterResource("azure-native:devices/v20230630preview:Certificate", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure-native:devices/v20230630preview:Certificate", name, id, state, &resource, opts...)
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
	// The name of the certificate
	CertificateName *string `pulumi:"certificateName"`
	// The description of an X509 CA Certificate.
	Properties *CertificateProperties `pulumi:"properties"`
	// The name of the resource group that contains the IoT hub.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the IoT hub.
	ResourceName string `pulumi:"resourceName"`
}

// The set of arguments for constructing a Certificate resource.
type CertificateArgs struct {
	// The name of the certificate
	CertificateName pulumi.StringPtrInput
	// The description of an X509 CA Certificate.
	Properties CertificatePropertiesPtrInput
	// The name of the resource group that contains the IoT hub.
	ResourceGroupName pulumi.StringInput
	// The name of the IoT hub.
	ResourceName pulumi.StringInput
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

// The entity tag.
func (o CertificateOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The name of the certificate.
func (o CertificateOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The description of an X509 CA Certificate.
func (o CertificateOutput) Properties() CertificatePropertiesResponseOutput {
	return o.ApplyT(func(v *Certificate) CertificatePropertiesResponseOutput { return v.Properties }).(CertificatePropertiesResponseOutput)
}

// The resource type.
func (o CertificateOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Certificate) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(CertificateOutput{})
}
