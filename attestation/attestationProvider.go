// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package attestation

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Attestation service response message.
// Azure REST API version: 2021-06-01. Prior API version in Azure Native 1.x: 2020-10-01
type AttestationProvider struct {
	pulumi.CustomResourceState

	// Gets the uri of attestation service
	AttestUri pulumi.StringPtrOutput `pulumi:"attestUri"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// List of private endpoint connections associated with the attestation provider.
	PrivateEndpointConnections PrivateEndpointConnectionResponseArrayOutput `pulumi:"privateEndpointConnections"`
	// Controls whether traffic from the public network is allowed to access the Attestation Provider APIs.
	PublicNetworkAccess pulumi.StringPtrOutput `pulumi:"publicNetworkAccess"`
	// Status of attestation service.
	Status pulumi.StringPtrOutput `pulumi:"status"`
	// The system metadata relating to this resource
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The setting that controls whether authentication is enabled or disabled for TPM Attestation REST APIs.
	TpmAttestationAuthentication pulumi.StringPtrOutput `pulumi:"tpmAttestationAuthentication"`
	// Trust model for the attestation provider.
	TrustModel pulumi.StringPtrOutput `pulumi:"trustModel"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewAttestationProvider registers a new resource with the given unique name, arguments, and options.
func NewAttestationProvider(ctx *pulumi.Context,
	name string, args *AttestationProviderArgs, opts ...pulumi.ResourceOption) (*AttestationProvider, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:attestation/v20180901preview:AttestationProvider"),
		},
		{
			Type: pulumi.String("azure-native:attestation/v20201001:AttestationProvider"),
		},
		{
			Type: pulumi.String("azure-native:attestation/v20210601:AttestationProvider"),
		},
		{
			Type: pulumi.String("azure-native:attestation/v20210601preview:AttestationProvider"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource AttestationProvider
	err := ctx.RegisterResource("azure-native:attestation:AttestationProvider", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAttestationProvider gets an existing AttestationProvider resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAttestationProvider(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AttestationProviderState, opts ...pulumi.ResourceOption) (*AttestationProvider, error) {
	var resource AttestationProvider
	err := ctx.ReadResource("azure-native:attestation:AttestationProvider", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AttestationProvider resources.
type attestationProviderState struct {
}

type AttestationProviderState struct {
}

func (AttestationProviderState) ElementType() reflect.Type {
	return reflect.TypeOf((*attestationProviderState)(nil)).Elem()
}

type attestationProviderArgs struct {
	// The supported Azure location where the attestation provider should be created.
	Location *string `pulumi:"location"`
	// Properties of the attestation provider
	Properties AttestationServiceCreationSpecificParams `pulumi:"properties"`
	// Name of the attestation provider.
	ProviderName *string `pulumi:"providerName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The tags that will be assigned to the attestation provider.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a AttestationProvider resource.
type AttestationProviderArgs struct {
	// The supported Azure location where the attestation provider should be created.
	Location pulumi.StringPtrInput
	// Properties of the attestation provider
	Properties AttestationServiceCreationSpecificParamsInput
	// Name of the attestation provider.
	ProviderName pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The tags that will be assigned to the attestation provider.
	Tags pulumi.StringMapInput
}

func (AttestationProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*attestationProviderArgs)(nil)).Elem()
}

type AttestationProviderInput interface {
	pulumi.Input

	ToAttestationProviderOutput() AttestationProviderOutput
	ToAttestationProviderOutputWithContext(ctx context.Context) AttestationProviderOutput
}

func (*AttestationProvider) ElementType() reflect.Type {
	return reflect.TypeOf((**AttestationProvider)(nil)).Elem()
}

func (i *AttestationProvider) ToAttestationProviderOutput() AttestationProviderOutput {
	return i.ToAttestationProviderOutputWithContext(context.Background())
}

func (i *AttestationProvider) ToAttestationProviderOutputWithContext(ctx context.Context) AttestationProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AttestationProviderOutput)
}

func (i *AttestationProvider) ToOutput(ctx context.Context) pulumix.Output[*AttestationProvider] {
	return pulumix.Output[*AttestationProvider]{
		OutputState: i.ToAttestationProviderOutputWithContext(ctx).OutputState,
	}
}

type AttestationProviderOutput struct{ *pulumi.OutputState }

func (AttestationProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AttestationProvider)(nil)).Elem()
}

func (o AttestationProviderOutput) ToAttestationProviderOutput() AttestationProviderOutput {
	return o
}

func (o AttestationProviderOutput) ToAttestationProviderOutputWithContext(ctx context.Context) AttestationProviderOutput {
	return o
}

func (o AttestationProviderOutput) ToOutput(ctx context.Context) pulumix.Output[*AttestationProvider] {
	return pulumix.Output[*AttestationProvider]{
		OutputState: o.OutputState,
	}
}

// Gets the uri of attestation service
func (o AttestationProviderOutput) AttestUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AttestationProvider) pulumi.StringPtrOutput { return v.AttestUri }).(pulumi.StringPtrOutput)
}

// The geo-location where the resource lives
func (o AttestationProviderOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *AttestationProvider) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o AttestationProviderOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AttestationProvider) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// List of private endpoint connections associated with the attestation provider.
func (o AttestationProviderOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v *AttestationProvider) PrivateEndpointConnectionResponseArrayOutput {
		return v.PrivateEndpointConnections
	}).(PrivateEndpointConnectionResponseArrayOutput)
}

// Controls whether traffic from the public network is allowed to access the Attestation Provider APIs.
func (o AttestationProviderOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AttestationProvider) pulumi.StringPtrOutput { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

// Status of attestation service.
func (o AttestationProviderOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AttestationProvider) pulumi.StringPtrOutput { return v.Status }).(pulumi.StringPtrOutput)
}

// The system metadata relating to this resource
func (o AttestationProviderOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *AttestationProvider) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o AttestationProviderOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AttestationProvider) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The setting that controls whether authentication is enabled or disabled for TPM Attestation REST APIs.
func (o AttestationProviderOutput) TpmAttestationAuthentication() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AttestationProvider) pulumi.StringPtrOutput { return v.TpmAttestationAuthentication }).(pulumi.StringPtrOutput)
}

// Trust model for the attestation provider.
func (o AttestationProviderOutput) TrustModel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AttestationProvider) pulumi.StringPtrOutput { return v.TrustModel }).(pulumi.StringPtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o AttestationProviderOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AttestationProvider) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AttestationProviderOutput{})
}
