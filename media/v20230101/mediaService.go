// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230101

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A Media Services account.
type MediaService struct {
	pulumi.CustomResourceState

	// The account encryption properties.
	Encryption AccountEncryptionResponsePtrOutput `pulumi:"encryption"`
	// The Managed Identity for the Media Services account.
	Identity MediaServiceIdentityResponsePtrOutput `pulumi:"identity"`
	// The Key Delivery properties for Media Services account.
	KeyDelivery KeyDeliveryResponsePtrOutput `pulumi:"keyDelivery"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The Media Services account ID.
	MediaServiceId pulumi.StringOutput `pulumi:"mediaServiceId"`
	// The minimum TLS version allowed for this account's requests. This is an optional property. If unspecified, a secure default value will be used.
	MinimumTlsVersion pulumi.StringPtrOutput `pulumi:"minimumTlsVersion"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The Private Endpoint Connections created for the Media Service account.
	PrivateEndpointConnections PrivateEndpointConnectionResponseArrayOutput `pulumi:"privateEndpointConnections"`
	// Provisioning state of the Media Services account.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Whether or not public network access is allowed for resources under the Media Services account.
	PublicNetworkAccess pulumi.StringPtrOutput `pulumi:"publicNetworkAccess"`
	// The storage accounts for this resource.
	StorageAccounts       StorageAccountResponseArrayOutput `pulumi:"storageAccounts"`
	StorageAuthentication pulumi.StringPtrOutput            `pulumi:"storageAuthentication"`
	// The system metadata relating to this resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewMediaService registers a new resource with the given unique name, arguments, and options.
func NewMediaService(ctx *pulumi.Context,
	name string, args *MediaServiceArgs, opts ...pulumi.ResourceOption) (*MediaService, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.MinimumTlsVersion == nil {
		args.MinimumTlsVersion = pulumi.StringPtr("Tls12")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:media:MediaService"),
		},
		{
			Type: pulumi.String("azure-native:media/v20151001:MediaService"),
		},
		{
			Type: pulumi.String("azure-native:media/v20180330preview:MediaService"),
		},
		{
			Type: pulumi.String("azure-native:media/v20180601preview:MediaService"),
		},
		{
			Type: pulumi.String("azure-native:media/v20180701:MediaService"),
		},
		{
			Type: pulumi.String("azure-native:media/v20200501:MediaService"),
		},
		{
			Type: pulumi.String("azure-native:media/v20210501:MediaService"),
		},
		{
			Type: pulumi.String("azure-native:media/v20210601:MediaService"),
		},
		{
			Type: pulumi.String("azure-native:media/v20211101:MediaService"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource MediaService
	err := ctx.RegisterResource("azure-native:media/v20230101:MediaService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMediaService gets an existing MediaService resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMediaService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MediaServiceState, opts ...pulumi.ResourceOption) (*MediaService, error) {
	var resource MediaService
	err := ctx.ReadResource("azure-native:media/v20230101:MediaService", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MediaService resources.
type mediaServiceState struct {
}

type MediaServiceState struct {
}

func (MediaServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*mediaServiceState)(nil)).Elem()
}

type mediaServiceArgs struct {
	// The Media Services account name.
	AccountName *string `pulumi:"accountName"`
	// The account encryption properties.
	Encryption *AccountEncryption `pulumi:"encryption"`
	// The Managed Identity for the Media Services account.
	Identity *MediaServiceIdentity `pulumi:"identity"`
	// The Key Delivery properties for Media Services account.
	KeyDelivery *KeyDelivery `pulumi:"keyDelivery"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The minimum TLS version allowed for this account's requests. This is an optional property. If unspecified, a secure default value will be used.
	MinimumTlsVersion *string `pulumi:"minimumTlsVersion"`
	// Whether or not public network access is allowed for resources under the Media Services account.
	PublicNetworkAccess *string `pulumi:"publicNetworkAccess"`
	// The name of the resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The storage accounts for this resource.
	StorageAccounts       []StorageAccount `pulumi:"storageAccounts"`
	StorageAuthentication *string          `pulumi:"storageAuthentication"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a MediaService resource.
type MediaServiceArgs struct {
	// The Media Services account name.
	AccountName pulumi.StringPtrInput
	// The account encryption properties.
	Encryption AccountEncryptionPtrInput
	// The Managed Identity for the Media Services account.
	Identity MediaServiceIdentityPtrInput
	// The Key Delivery properties for Media Services account.
	KeyDelivery KeyDeliveryPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The minimum TLS version allowed for this account's requests. This is an optional property. If unspecified, a secure default value will be used.
	MinimumTlsVersion pulumi.StringPtrInput
	// Whether or not public network access is allowed for resources under the Media Services account.
	PublicNetworkAccess pulumi.StringPtrInput
	// The name of the resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput
	// The storage accounts for this resource.
	StorageAccounts       StorageAccountArrayInput
	StorageAuthentication pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (MediaServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mediaServiceArgs)(nil)).Elem()
}

type MediaServiceInput interface {
	pulumi.Input

	ToMediaServiceOutput() MediaServiceOutput
	ToMediaServiceOutputWithContext(ctx context.Context) MediaServiceOutput
}

func (*MediaService) ElementType() reflect.Type {
	return reflect.TypeOf((**MediaService)(nil)).Elem()
}

func (i *MediaService) ToMediaServiceOutput() MediaServiceOutput {
	return i.ToMediaServiceOutputWithContext(context.Background())
}

func (i *MediaService) ToMediaServiceOutputWithContext(ctx context.Context) MediaServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MediaServiceOutput)
}

type MediaServiceOutput struct{ *pulumi.OutputState }

func (MediaServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MediaService)(nil)).Elem()
}

func (o MediaServiceOutput) ToMediaServiceOutput() MediaServiceOutput {
	return o
}

func (o MediaServiceOutput) ToMediaServiceOutputWithContext(ctx context.Context) MediaServiceOutput {
	return o
}

// The account encryption properties.
func (o MediaServiceOutput) Encryption() AccountEncryptionResponsePtrOutput {
	return o.ApplyT(func(v *MediaService) AccountEncryptionResponsePtrOutput { return v.Encryption }).(AccountEncryptionResponsePtrOutput)
}

// The Managed Identity for the Media Services account.
func (o MediaServiceOutput) Identity() MediaServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *MediaService) MediaServiceIdentityResponsePtrOutput { return v.Identity }).(MediaServiceIdentityResponsePtrOutput)
}

// The Key Delivery properties for Media Services account.
func (o MediaServiceOutput) KeyDelivery() KeyDeliveryResponsePtrOutput {
	return o.ApplyT(func(v *MediaService) KeyDeliveryResponsePtrOutput { return v.KeyDelivery }).(KeyDeliveryResponsePtrOutput)
}

// The geo-location where the resource lives
func (o MediaServiceOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *MediaService) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The Media Services account ID.
func (o MediaServiceOutput) MediaServiceId() pulumi.StringOutput {
	return o.ApplyT(func(v *MediaService) pulumi.StringOutput { return v.MediaServiceId }).(pulumi.StringOutput)
}

// The minimum TLS version allowed for this account's requests. This is an optional property. If unspecified, a secure default value will be used.
func (o MediaServiceOutput) MinimumTlsVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MediaService) pulumi.StringPtrOutput { return v.MinimumTlsVersion }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o MediaServiceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MediaService) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The Private Endpoint Connections created for the Media Service account.
func (o MediaServiceOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v *MediaService) PrivateEndpointConnectionResponseArrayOutput {
		return v.PrivateEndpointConnections
	}).(PrivateEndpointConnectionResponseArrayOutput)
}

// Provisioning state of the Media Services account.
func (o MediaServiceOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *MediaService) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Whether or not public network access is allowed for resources under the Media Services account.
func (o MediaServiceOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MediaService) pulumi.StringPtrOutput { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

// The storage accounts for this resource.
func (o MediaServiceOutput) StorageAccounts() StorageAccountResponseArrayOutput {
	return o.ApplyT(func(v *MediaService) StorageAccountResponseArrayOutput { return v.StorageAccounts }).(StorageAccountResponseArrayOutput)
}

func (o MediaServiceOutput) StorageAuthentication() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MediaService) pulumi.StringPtrOutput { return v.StorageAuthentication }).(pulumi.StringPtrOutput)
}

// The system metadata relating to this resource.
func (o MediaServiceOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *MediaService) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o MediaServiceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *MediaService) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o MediaServiceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *MediaService) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(MediaServiceOutput{})
}
