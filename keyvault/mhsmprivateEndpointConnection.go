// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package keyvault

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Private endpoint connection resource.
// Azure REST API version: 2023-02-01. Prior API version in Azure Native 1.x: 2021-06-01-preview.
//
// Other available API versions: 2023-07-01.
type MHSMPrivateEndpointConnection struct {
	pulumi.CustomResourceState

	// Modified whenever there is a change in the state of private endpoint connection.
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// The supported Azure location where the managed HSM Pool should be created.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the managed HSM Pool.
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties of the private endpoint object.
	PrivateEndpoint MHSMPrivateEndpointResponsePtrOutput `pulumi:"privateEndpoint"`
	// Approval state of the private link connection.
	PrivateLinkServiceConnectionState MHSMPrivateLinkServiceConnectionStateResponsePtrOutput `pulumi:"privateLinkServiceConnectionState"`
	// Provisioning state of the private endpoint connection.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// SKU details
	Sku ManagedHsmSkuResponsePtrOutput `pulumi:"sku"`
	// Metadata pertaining to creation and last modification of the key vault resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The resource type of the managed HSM Pool.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewMHSMPrivateEndpointConnection registers a new resource with the given unique name, arguments, and options.
func NewMHSMPrivateEndpointConnection(ctx *pulumi.Context,
	name string, args *MHSMPrivateEndpointConnectionArgs, opts ...pulumi.ResourceOption) (*MHSMPrivateEndpointConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:keyvault/v20210401preview:MHSMPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:keyvault/v20210601preview:MHSMPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:keyvault/v20211001:MHSMPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:keyvault/v20211101preview:MHSMPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:keyvault/v20220201preview:MHSMPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:keyvault/v20220701:MHSMPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:keyvault/v20221101:MHSMPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:keyvault/v20230201:MHSMPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:keyvault/v20230701:MHSMPrivateEndpointConnection"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource MHSMPrivateEndpointConnection
	err := ctx.RegisterResource("azure-native:keyvault:MHSMPrivateEndpointConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMHSMPrivateEndpointConnection gets an existing MHSMPrivateEndpointConnection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMHSMPrivateEndpointConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MHSMPrivateEndpointConnectionState, opts ...pulumi.ResourceOption) (*MHSMPrivateEndpointConnection, error) {
	var resource MHSMPrivateEndpointConnection
	err := ctx.ReadResource("azure-native:keyvault:MHSMPrivateEndpointConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MHSMPrivateEndpointConnection resources.
type mhsmprivateEndpointConnectionState struct {
}

type MHSMPrivateEndpointConnectionState struct {
}

func (MHSMPrivateEndpointConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*mhsmprivateEndpointConnectionState)(nil)).Elem()
}

type mhsmprivateEndpointConnectionArgs struct {
	// The supported Azure location where the managed HSM Pool should be created.
	Location *string `pulumi:"location"`
	// Name of the managed HSM Pool
	Name string `pulumi:"name"`
	// Name of the private endpoint connection associated with the managed hsm pool.
	PrivateEndpointConnectionName *string `pulumi:"privateEndpointConnectionName"`
	// Approval state of the private link connection.
	PrivateLinkServiceConnectionState *MHSMPrivateLinkServiceConnectionState `pulumi:"privateLinkServiceConnectionState"`
	// Name of the resource group that contains the managed HSM pool.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// SKU details
	Sku *ManagedHsmSku `pulumi:"sku"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a MHSMPrivateEndpointConnection resource.
type MHSMPrivateEndpointConnectionArgs struct {
	// The supported Azure location where the managed HSM Pool should be created.
	Location pulumi.StringPtrInput
	// Name of the managed HSM Pool
	Name pulumi.StringInput
	// Name of the private endpoint connection associated with the managed hsm pool.
	PrivateEndpointConnectionName pulumi.StringPtrInput
	// Approval state of the private link connection.
	PrivateLinkServiceConnectionState MHSMPrivateLinkServiceConnectionStatePtrInput
	// Name of the resource group that contains the managed HSM pool.
	ResourceGroupName pulumi.StringInput
	// SKU details
	Sku ManagedHsmSkuPtrInput
	// Resource tags
	Tags pulumi.StringMapInput
}

func (MHSMPrivateEndpointConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mhsmprivateEndpointConnectionArgs)(nil)).Elem()
}

type MHSMPrivateEndpointConnectionInput interface {
	pulumi.Input

	ToMHSMPrivateEndpointConnectionOutput() MHSMPrivateEndpointConnectionOutput
	ToMHSMPrivateEndpointConnectionOutputWithContext(ctx context.Context) MHSMPrivateEndpointConnectionOutput
}

func (*MHSMPrivateEndpointConnection) ElementType() reflect.Type {
	return reflect.TypeOf((**MHSMPrivateEndpointConnection)(nil)).Elem()
}

func (i *MHSMPrivateEndpointConnection) ToMHSMPrivateEndpointConnectionOutput() MHSMPrivateEndpointConnectionOutput {
	return i.ToMHSMPrivateEndpointConnectionOutputWithContext(context.Background())
}

func (i *MHSMPrivateEndpointConnection) ToMHSMPrivateEndpointConnectionOutputWithContext(ctx context.Context) MHSMPrivateEndpointConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MHSMPrivateEndpointConnectionOutput)
}

func (i *MHSMPrivateEndpointConnection) ToOutput(ctx context.Context) pulumix.Output[*MHSMPrivateEndpointConnection] {
	return pulumix.Output[*MHSMPrivateEndpointConnection]{
		OutputState: i.ToMHSMPrivateEndpointConnectionOutputWithContext(ctx).OutputState,
	}
}

type MHSMPrivateEndpointConnectionOutput struct{ *pulumi.OutputState }

func (MHSMPrivateEndpointConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MHSMPrivateEndpointConnection)(nil)).Elem()
}

func (o MHSMPrivateEndpointConnectionOutput) ToMHSMPrivateEndpointConnectionOutput() MHSMPrivateEndpointConnectionOutput {
	return o
}

func (o MHSMPrivateEndpointConnectionOutput) ToMHSMPrivateEndpointConnectionOutputWithContext(ctx context.Context) MHSMPrivateEndpointConnectionOutput {
	return o
}

func (o MHSMPrivateEndpointConnectionOutput) ToOutput(ctx context.Context) pulumix.Output[*MHSMPrivateEndpointConnection] {
	return pulumix.Output[*MHSMPrivateEndpointConnection]{
		OutputState: o.OutputState,
	}
}

// Modified whenever there is a change in the state of private endpoint connection.
func (o MHSMPrivateEndpointConnectionOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MHSMPrivateEndpointConnection) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// The supported Azure location where the managed HSM Pool should be created.
func (o MHSMPrivateEndpointConnectionOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MHSMPrivateEndpointConnection) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the managed HSM Pool.
func (o MHSMPrivateEndpointConnectionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MHSMPrivateEndpointConnection) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Properties of the private endpoint object.
func (o MHSMPrivateEndpointConnectionOutput) PrivateEndpoint() MHSMPrivateEndpointResponsePtrOutput {
	return o.ApplyT(func(v *MHSMPrivateEndpointConnection) MHSMPrivateEndpointResponsePtrOutput { return v.PrivateEndpoint }).(MHSMPrivateEndpointResponsePtrOutput)
}

// Approval state of the private link connection.
func (o MHSMPrivateEndpointConnectionOutput) PrivateLinkServiceConnectionState() MHSMPrivateLinkServiceConnectionStateResponsePtrOutput {
	return o.ApplyT(func(v *MHSMPrivateEndpointConnection) MHSMPrivateLinkServiceConnectionStateResponsePtrOutput {
		return v.PrivateLinkServiceConnectionState
	}).(MHSMPrivateLinkServiceConnectionStateResponsePtrOutput)
}

// Provisioning state of the private endpoint connection.
func (o MHSMPrivateEndpointConnectionOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *MHSMPrivateEndpointConnection) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// SKU details
func (o MHSMPrivateEndpointConnectionOutput) Sku() ManagedHsmSkuResponsePtrOutput {
	return o.ApplyT(func(v *MHSMPrivateEndpointConnection) ManagedHsmSkuResponsePtrOutput { return v.Sku }).(ManagedHsmSkuResponsePtrOutput)
}

// Metadata pertaining to creation and last modification of the key vault resource.
func (o MHSMPrivateEndpointConnectionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *MHSMPrivateEndpointConnection) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags
func (o MHSMPrivateEndpointConnectionOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *MHSMPrivateEndpointConnection) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The resource type of the managed HSM Pool.
func (o MHSMPrivateEndpointConnectionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *MHSMPrivateEndpointConnection) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(MHSMPrivateEndpointConnectionOutput{})
}
