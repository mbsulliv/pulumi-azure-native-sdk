// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230601preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A custom certificate.
type WebPubSubCustomCertificate struct {
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

// NewWebPubSubCustomCertificate registers a new resource with the given unique name, arguments, and options.
func NewWebPubSubCustomCertificate(ctx *pulumi.Context,
	name string, args *WebPubSubCustomCertificateArgs, opts ...pulumi.ResourceOption) (*WebPubSubCustomCertificate, error) {
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
			Type: pulumi.String("azure-native:webpubsub:WebPubSubCustomCertificate"),
		},
		{
			Type: pulumi.String("azure-native:webpubsub/v20220801preview:WebPubSubCustomCertificate"),
		},
		{
			Type: pulumi.String("azure-native:webpubsub/v20230201:WebPubSubCustomCertificate"),
		},
		{
			Type: pulumi.String("azure-native:webpubsub/v20230301preview:WebPubSubCustomCertificate"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource WebPubSubCustomCertificate
	err := ctx.RegisterResource("azure-native:webpubsub/v20230601preview:WebPubSubCustomCertificate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebPubSubCustomCertificate gets an existing WebPubSubCustomCertificate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebPubSubCustomCertificate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebPubSubCustomCertificateState, opts ...pulumi.ResourceOption) (*WebPubSubCustomCertificate, error) {
	var resource WebPubSubCustomCertificate
	err := ctx.ReadResource("azure-native:webpubsub/v20230601preview:WebPubSubCustomCertificate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebPubSubCustomCertificate resources.
type webPubSubCustomCertificateState struct {
}

type WebPubSubCustomCertificateState struct {
}

func (WebPubSubCustomCertificateState) ElementType() reflect.Type {
	return reflect.TypeOf((*webPubSubCustomCertificateState)(nil)).Elem()
}

type webPubSubCustomCertificateArgs struct {
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

// The set of arguments for constructing a WebPubSubCustomCertificate resource.
type WebPubSubCustomCertificateArgs struct {
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

func (WebPubSubCustomCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webPubSubCustomCertificateArgs)(nil)).Elem()
}

type WebPubSubCustomCertificateInput interface {
	pulumi.Input

	ToWebPubSubCustomCertificateOutput() WebPubSubCustomCertificateOutput
	ToWebPubSubCustomCertificateOutputWithContext(ctx context.Context) WebPubSubCustomCertificateOutput
}

func (*WebPubSubCustomCertificate) ElementType() reflect.Type {
	return reflect.TypeOf((**WebPubSubCustomCertificate)(nil)).Elem()
}

func (i *WebPubSubCustomCertificate) ToWebPubSubCustomCertificateOutput() WebPubSubCustomCertificateOutput {
	return i.ToWebPubSubCustomCertificateOutputWithContext(context.Background())
}

func (i *WebPubSubCustomCertificate) ToWebPubSubCustomCertificateOutputWithContext(ctx context.Context) WebPubSubCustomCertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebPubSubCustomCertificateOutput)
}

type WebPubSubCustomCertificateOutput struct{ *pulumi.OutputState }

func (WebPubSubCustomCertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebPubSubCustomCertificate)(nil)).Elem()
}

func (o WebPubSubCustomCertificateOutput) ToWebPubSubCustomCertificateOutput() WebPubSubCustomCertificateOutput {
	return o
}

func (o WebPubSubCustomCertificateOutput) ToWebPubSubCustomCertificateOutputWithContext(ctx context.Context) WebPubSubCustomCertificateOutput {
	return o
}

// Base uri of the KeyVault that stores certificate.
func (o WebPubSubCustomCertificateOutput) KeyVaultBaseUri() pulumi.StringOutput {
	return o.ApplyT(func(v *WebPubSubCustomCertificate) pulumi.StringOutput { return v.KeyVaultBaseUri }).(pulumi.StringOutput)
}

// Certificate secret name.
func (o WebPubSubCustomCertificateOutput) KeyVaultSecretName() pulumi.StringOutput {
	return o.ApplyT(func(v *WebPubSubCustomCertificate) pulumi.StringOutput { return v.KeyVaultSecretName }).(pulumi.StringOutput)
}

// Certificate secret version.
func (o WebPubSubCustomCertificateOutput) KeyVaultSecretVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebPubSubCustomCertificate) pulumi.StringPtrOutput { return v.KeyVaultSecretVersion }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o WebPubSubCustomCertificateOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WebPubSubCustomCertificate) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Provisioning state of the resource.
func (o WebPubSubCustomCertificateOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *WebPubSubCustomCertificate) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o WebPubSubCustomCertificateOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *WebPubSubCustomCertificate) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o WebPubSubCustomCertificateOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WebPubSubCustomCertificate) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WebPubSubCustomCertificateOutput{})
}