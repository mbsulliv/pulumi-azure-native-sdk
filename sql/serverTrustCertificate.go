// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sql

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Server trust certificate imported from box to enable connection between box and Sql Managed Instance.
// Azure REST API version: 2021-11-01. Prior API version in Azure Native 1.x: 2021-05-01-preview.
//
// Other available API versions: 2022-11-01-preview, 2023-02-01-preview, 2023-05-01-preview.
type ServerTrustCertificate struct {
	pulumi.CustomResourceState

	// The certificate name
	CertificateName pulumi.StringOutput `pulumi:"certificateName"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The certificate public blob
	PublicBlob pulumi.StringPtrOutput `pulumi:"publicBlob"`
	// The certificate thumbprint
	Thumbprint pulumi.StringOutput `pulumi:"thumbprint"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewServerTrustCertificate registers a new resource with the given unique name, arguments, and options.
func NewServerTrustCertificate(ctx *pulumi.Context,
	name string, args *ServerTrustCertificateArgs, opts ...pulumi.ResourceOption) (*ServerTrustCertificate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ManagedInstanceName == nil {
		return nil, errors.New("invalid value for required argument 'ManagedInstanceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:sql/v20210501preview:ServerTrustCertificate"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210801preview:ServerTrustCertificate"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101:ServerTrustCertificate"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101preview:ServerTrustCertificate"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220201preview:ServerTrustCertificate"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220501preview:ServerTrustCertificate"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220801preview:ServerTrustCertificate"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20221101preview:ServerTrustCertificate"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20230201preview:ServerTrustCertificate"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20230501preview:ServerTrustCertificate"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ServerTrustCertificate
	err := ctx.RegisterResource("azure-native:sql:ServerTrustCertificate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServerTrustCertificate gets an existing ServerTrustCertificate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServerTrustCertificate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServerTrustCertificateState, opts ...pulumi.ResourceOption) (*ServerTrustCertificate, error) {
	var resource ServerTrustCertificate
	err := ctx.ReadResource("azure-native:sql:ServerTrustCertificate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServerTrustCertificate resources.
type serverTrustCertificateState struct {
}

type ServerTrustCertificateState struct {
}

func (ServerTrustCertificateState) ElementType() reflect.Type {
	return reflect.TypeOf((*serverTrustCertificateState)(nil)).Elem()
}

type serverTrustCertificateArgs struct {
	// Name of of the certificate to upload.
	CertificateName *string `pulumi:"certificateName"`
	// The name of the managed instance.
	ManagedInstanceName string `pulumi:"managedInstanceName"`
	// The certificate public blob
	PublicBlob *string `pulumi:"publicBlob"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a ServerTrustCertificate resource.
type ServerTrustCertificateArgs struct {
	// Name of of the certificate to upload.
	CertificateName pulumi.StringPtrInput
	// The name of the managed instance.
	ManagedInstanceName pulumi.StringInput
	// The certificate public blob
	PublicBlob pulumi.StringPtrInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
}

func (ServerTrustCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serverTrustCertificateArgs)(nil)).Elem()
}

type ServerTrustCertificateInput interface {
	pulumi.Input

	ToServerTrustCertificateOutput() ServerTrustCertificateOutput
	ToServerTrustCertificateOutputWithContext(ctx context.Context) ServerTrustCertificateOutput
}

func (*ServerTrustCertificate) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerTrustCertificate)(nil)).Elem()
}

func (i *ServerTrustCertificate) ToServerTrustCertificateOutput() ServerTrustCertificateOutput {
	return i.ToServerTrustCertificateOutputWithContext(context.Background())
}

func (i *ServerTrustCertificate) ToServerTrustCertificateOutputWithContext(ctx context.Context) ServerTrustCertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerTrustCertificateOutput)
}

type ServerTrustCertificateOutput struct{ *pulumi.OutputState }

func (ServerTrustCertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerTrustCertificate)(nil)).Elem()
}

func (o ServerTrustCertificateOutput) ToServerTrustCertificateOutput() ServerTrustCertificateOutput {
	return o
}

func (o ServerTrustCertificateOutput) ToServerTrustCertificateOutputWithContext(ctx context.Context) ServerTrustCertificateOutput {
	return o
}

// The certificate name
func (o ServerTrustCertificateOutput) CertificateName() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerTrustCertificate) pulumi.StringOutput { return v.CertificateName }).(pulumi.StringOutput)
}

// Resource name.
func (o ServerTrustCertificateOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerTrustCertificate) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The certificate public blob
func (o ServerTrustCertificateOutput) PublicBlob() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerTrustCertificate) pulumi.StringPtrOutput { return v.PublicBlob }).(pulumi.StringPtrOutput)
}

// The certificate thumbprint
func (o ServerTrustCertificateOutput) Thumbprint() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerTrustCertificate) pulumi.StringOutput { return v.Thumbprint }).(pulumi.StringOutput)
}

// Resource type.
func (o ServerTrustCertificateOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerTrustCertificate) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ServerTrustCertificateOutput{})
}
