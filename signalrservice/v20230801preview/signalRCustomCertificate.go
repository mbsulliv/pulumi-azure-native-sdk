// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230801preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// A custom certificate.
type SignalRCustomCertificate struct {
	pulumi.CustomResourceState

	// Base uri of the KeyVault that stores certificate.
	KeyVaultBaseUri pulumi.StringOutput `pulumi:"keyVaultBaseUri"`
	// Certificate secret name.
	KeyVaultSecretName pulumi.StringOutput `pulumi:"keyVaultSecretName"`
	// Certificate secret version.
	KeyVaultSecretVersion pulumi.StringPtrOutput `pulumi:"keyVaultSecretVersion"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Provisioning state of the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewSignalRCustomCertificate registers a new resource with the given unique name, arguments, and options.
func NewSignalRCustomCertificate(ctx *pulumi.Context,
	name string, args *SignalRCustomCertificateArgs, opts ...pulumi.ResourceOption) (*SignalRCustomCertificate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.KeyVaultBaseUri == nil {
		return nil, errors.New("invalid value for required argument 'KeyVaultBaseUri'")
	}
	if args.KeyVaultSecretName == nil {
		return nil, errors.New("invalid value for required argument 'KeyVaultSecretName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ResourceName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:signalrservice:SignalRCustomCertificate"),
		},
		{
			Type: pulumi.String("azure-native:signalrservice/v20220201:SignalRCustomCertificate"),
		},
		{
			Type: pulumi.String("azure-native:signalrservice/v20220801preview:SignalRCustomCertificate"),
		},
		{
			Type: pulumi.String("azure-native:signalrservice/v20230201:SignalRCustomCertificate"),
		},
		{
			Type: pulumi.String("azure-native:signalrservice/v20230301preview:SignalRCustomCertificate"),
		},
		{
			Type: pulumi.String("azure-native:signalrservice/v20230601preview:SignalRCustomCertificate"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource SignalRCustomCertificate
	err := ctx.RegisterResource("azure-native:signalrservice/v20230801preview:SignalRCustomCertificate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSignalRCustomCertificate gets an existing SignalRCustomCertificate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSignalRCustomCertificate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SignalRCustomCertificateState, opts ...pulumi.ResourceOption) (*SignalRCustomCertificate, error) {
	var resource SignalRCustomCertificate
	err := ctx.ReadResource("azure-native:signalrservice/v20230801preview:SignalRCustomCertificate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SignalRCustomCertificate resources.
type signalRCustomCertificateState struct {
}

type SignalRCustomCertificateState struct {
}

func (SignalRCustomCertificateState) ElementType() reflect.Type {
	return reflect.TypeOf((*signalRCustomCertificateState)(nil)).Elem()
}

type signalRCustomCertificateArgs struct {
	// Custom certificate name
	CertificateName *string `pulumi:"certificateName"`
	// Base uri of the KeyVault that stores certificate.
	KeyVaultBaseUri string `pulumi:"keyVaultBaseUri"`
	// Certificate secret name.
	KeyVaultSecretName string `pulumi:"keyVaultSecretName"`
	// Certificate secret version.
	KeyVaultSecretVersion *string `pulumi:"keyVaultSecretVersion"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the resource.
	ResourceName string `pulumi:"resourceName"`
}

// The set of arguments for constructing a SignalRCustomCertificate resource.
type SignalRCustomCertificateArgs struct {
	// Custom certificate name
	CertificateName pulumi.StringPtrInput
	// Base uri of the KeyVault that stores certificate.
	KeyVaultBaseUri pulumi.StringInput
	// Certificate secret name.
	KeyVaultSecretName pulumi.StringInput
	// Certificate secret version.
	KeyVaultSecretVersion pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the resource.
	ResourceName pulumi.StringInput
}

func (SignalRCustomCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*signalRCustomCertificateArgs)(nil)).Elem()
}

type SignalRCustomCertificateInput interface {
	pulumi.Input

	ToSignalRCustomCertificateOutput() SignalRCustomCertificateOutput
	ToSignalRCustomCertificateOutputWithContext(ctx context.Context) SignalRCustomCertificateOutput
}

func (*SignalRCustomCertificate) ElementType() reflect.Type {
	return reflect.TypeOf((**SignalRCustomCertificate)(nil)).Elem()
}

func (i *SignalRCustomCertificate) ToSignalRCustomCertificateOutput() SignalRCustomCertificateOutput {
	return i.ToSignalRCustomCertificateOutputWithContext(context.Background())
}

func (i *SignalRCustomCertificate) ToSignalRCustomCertificateOutputWithContext(ctx context.Context) SignalRCustomCertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SignalRCustomCertificateOutput)
}

func (i *SignalRCustomCertificate) ToOutput(ctx context.Context) pulumix.Output[*SignalRCustomCertificate] {
	return pulumix.Output[*SignalRCustomCertificate]{
		OutputState: i.ToSignalRCustomCertificateOutputWithContext(ctx).OutputState,
	}
}

type SignalRCustomCertificateOutput struct{ *pulumi.OutputState }

func (SignalRCustomCertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SignalRCustomCertificate)(nil)).Elem()
}

func (o SignalRCustomCertificateOutput) ToSignalRCustomCertificateOutput() SignalRCustomCertificateOutput {
	return o
}

func (o SignalRCustomCertificateOutput) ToSignalRCustomCertificateOutputWithContext(ctx context.Context) SignalRCustomCertificateOutput {
	return o
}

func (o SignalRCustomCertificateOutput) ToOutput(ctx context.Context) pulumix.Output[*SignalRCustomCertificate] {
	return pulumix.Output[*SignalRCustomCertificate]{
		OutputState: o.OutputState,
	}
}

// Base uri of the KeyVault that stores certificate.
func (o SignalRCustomCertificateOutput) KeyVaultBaseUri() pulumi.StringOutput {
	return o.ApplyT(func(v *SignalRCustomCertificate) pulumi.StringOutput { return v.KeyVaultBaseUri }).(pulumi.StringOutput)
}

// Certificate secret name.
func (o SignalRCustomCertificateOutput) KeyVaultSecretName() pulumi.StringOutput {
	return o.ApplyT(func(v *SignalRCustomCertificate) pulumi.StringOutput { return v.KeyVaultSecretName }).(pulumi.StringOutput)
}

// Certificate secret version.
func (o SignalRCustomCertificateOutput) KeyVaultSecretVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SignalRCustomCertificate) pulumi.StringPtrOutput { return v.KeyVaultSecretVersion }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o SignalRCustomCertificateOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SignalRCustomCertificate) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Provisioning state of the resource.
func (o SignalRCustomCertificateOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *SignalRCustomCertificate) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o SignalRCustomCertificateOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *SignalRCustomCertificate) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o SignalRCustomCertificateOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SignalRCustomCertificate) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SignalRCustomCertificateOutput{})
}
