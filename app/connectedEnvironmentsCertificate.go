// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package app

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Certificate used for Custom Domain bindings of Container Apps in a Managed Environment
// Azure REST API version: 2022-10-01.
//
// Other available API versions: 2023-04-01-preview, 2023-05-01, 2023-05-02-preview, 2023-08-01-preview.
type ConnectedEnvironmentsCertificate struct {
	pulumi.CustomResourceState

	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Certificate resource specific properties
	Properties CertificateResponsePropertiesOutput `pulumi:"properties"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewConnectedEnvironmentsCertificate registers a new resource with the given unique name, arguments, and options.
func NewConnectedEnvironmentsCertificate(ctx *pulumi.Context,
	name string, args *ConnectedEnvironmentsCertificateArgs, opts ...pulumi.ResourceOption) (*ConnectedEnvironmentsCertificate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConnectedEnvironmentName == nil {
		return nil, errors.New("invalid value for required argument 'ConnectedEnvironmentName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:app/v20220601preview:ConnectedEnvironmentsCertificate"),
		},
		{
			Type: pulumi.String("azure-native:app/v20221001:ConnectedEnvironmentsCertificate"),
		},
		{
			Type: pulumi.String("azure-native:app/v20221101preview:ConnectedEnvironmentsCertificate"),
		},
		{
			Type: pulumi.String("azure-native:app/v20230401preview:ConnectedEnvironmentsCertificate"),
		},
		{
			Type: pulumi.String("azure-native:app/v20230501:ConnectedEnvironmentsCertificate"),
		},
		{
			Type: pulumi.String("azure-native:app/v20230502preview:ConnectedEnvironmentsCertificate"),
		},
		{
			Type: pulumi.String("azure-native:app/v20230801preview:ConnectedEnvironmentsCertificate"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ConnectedEnvironmentsCertificate
	err := ctx.RegisterResource("azure-native:app:ConnectedEnvironmentsCertificate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConnectedEnvironmentsCertificate gets an existing ConnectedEnvironmentsCertificate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConnectedEnvironmentsCertificate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConnectedEnvironmentsCertificateState, opts ...pulumi.ResourceOption) (*ConnectedEnvironmentsCertificate, error) {
	var resource ConnectedEnvironmentsCertificate
	err := ctx.ReadResource("azure-native:app:ConnectedEnvironmentsCertificate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ConnectedEnvironmentsCertificate resources.
type connectedEnvironmentsCertificateState struct {
}

type ConnectedEnvironmentsCertificateState struct {
}

func (ConnectedEnvironmentsCertificateState) ElementType() reflect.Type {
	return reflect.TypeOf((*connectedEnvironmentsCertificateState)(nil)).Elem()
}

type connectedEnvironmentsCertificateArgs struct {
	// Name of the Certificate.
	CertificateName *string `pulumi:"certificateName"`
	// Name of the Connected Environment.
	ConnectedEnvironmentName string `pulumi:"connectedEnvironmentName"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// Certificate resource specific properties
	Properties *CertificateProperties `pulumi:"properties"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a ConnectedEnvironmentsCertificate resource.
type ConnectedEnvironmentsCertificateArgs struct {
	// Name of the Certificate.
	CertificateName pulumi.StringPtrInput
	// Name of the Connected Environment.
	ConnectedEnvironmentName pulumi.StringInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// Certificate resource specific properties
	Properties CertificatePropertiesPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (ConnectedEnvironmentsCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*connectedEnvironmentsCertificateArgs)(nil)).Elem()
}

type ConnectedEnvironmentsCertificateInput interface {
	pulumi.Input

	ToConnectedEnvironmentsCertificateOutput() ConnectedEnvironmentsCertificateOutput
	ToConnectedEnvironmentsCertificateOutputWithContext(ctx context.Context) ConnectedEnvironmentsCertificateOutput
}

func (*ConnectedEnvironmentsCertificate) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectedEnvironmentsCertificate)(nil)).Elem()
}

func (i *ConnectedEnvironmentsCertificate) ToConnectedEnvironmentsCertificateOutput() ConnectedEnvironmentsCertificateOutput {
	return i.ToConnectedEnvironmentsCertificateOutputWithContext(context.Background())
}

func (i *ConnectedEnvironmentsCertificate) ToConnectedEnvironmentsCertificateOutputWithContext(ctx context.Context) ConnectedEnvironmentsCertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectedEnvironmentsCertificateOutput)
}

type ConnectedEnvironmentsCertificateOutput struct{ *pulumi.OutputState }

func (ConnectedEnvironmentsCertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectedEnvironmentsCertificate)(nil)).Elem()
}

func (o ConnectedEnvironmentsCertificateOutput) ToConnectedEnvironmentsCertificateOutput() ConnectedEnvironmentsCertificateOutput {
	return o
}

func (o ConnectedEnvironmentsCertificateOutput) ToConnectedEnvironmentsCertificateOutputWithContext(ctx context.Context) ConnectedEnvironmentsCertificateOutput {
	return o
}

// The geo-location where the resource lives
func (o ConnectedEnvironmentsCertificateOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectedEnvironmentsCertificate) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o ConnectedEnvironmentsCertificateOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectedEnvironmentsCertificate) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Certificate resource specific properties
func (o ConnectedEnvironmentsCertificateOutput) Properties() CertificateResponsePropertiesOutput {
	return o.ApplyT(func(v *ConnectedEnvironmentsCertificate) CertificateResponsePropertiesOutput { return v.Properties }).(CertificateResponsePropertiesOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o ConnectedEnvironmentsCertificateOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ConnectedEnvironmentsCertificate) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o ConnectedEnvironmentsCertificateOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ConnectedEnvironmentsCertificate) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ConnectedEnvironmentsCertificateOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectedEnvironmentsCertificate) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ConnectedEnvironmentsCertificateOutput{})
}
