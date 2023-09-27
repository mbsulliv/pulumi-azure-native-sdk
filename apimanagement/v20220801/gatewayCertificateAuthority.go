// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220801

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gateway certificate authority details.
type GatewayCertificateAuthority struct {
	pulumi.CustomResourceState

	// Determines whether certificate authority is trusted.
	IsTrusted pulumi.BoolPtrOutput `pulumi:"isTrusted"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewGatewayCertificateAuthority registers a new resource with the given unique name, arguments, and options.
func NewGatewayCertificateAuthority(ctx *pulumi.Context,
	name string, args *GatewayCertificateAuthorityArgs, opts ...pulumi.ResourceOption) (*GatewayCertificateAuthority, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GatewayId == nil {
		return nil, errors.New("invalid value for required argument 'GatewayId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:apimanagement:GatewayCertificateAuthority"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20200601preview:GatewayCertificateAuthority"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20201201:GatewayCertificateAuthority"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210101preview:GatewayCertificateAuthority"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210401preview:GatewayCertificateAuthority"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210801:GatewayCertificateAuthority"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20211201preview:GatewayCertificateAuthority"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220401preview:GatewayCertificateAuthority"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220901preview:GatewayCertificateAuthority"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20230301preview:GatewayCertificateAuthority"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource GatewayCertificateAuthority
	err := ctx.RegisterResource("azure-native:apimanagement/v20220801:GatewayCertificateAuthority", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGatewayCertificateAuthority gets an existing GatewayCertificateAuthority resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGatewayCertificateAuthority(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GatewayCertificateAuthorityState, opts ...pulumi.ResourceOption) (*GatewayCertificateAuthority, error) {
	var resource GatewayCertificateAuthority
	err := ctx.ReadResource("azure-native:apimanagement/v20220801:GatewayCertificateAuthority", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GatewayCertificateAuthority resources.
type gatewayCertificateAuthorityState struct {
}

type GatewayCertificateAuthorityState struct {
}

func (GatewayCertificateAuthorityState) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayCertificateAuthorityState)(nil)).Elem()
}

type gatewayCertificateAuthorityArgs struct {
	// Identifier of the certificate entity. Must be unique in the current API Management service instance.
	CertificateId *string `pulumi:"certificateId"`
	// Gateway entity identifier. Must be unique in the current API Management service instance. Must not have value 'managed'
	GatewayId string `pulumi:"gatewayId"`
	// Determines whether certificate authority is trusted.
	IsTrusted *bool `pulumi:"isTrusted"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a GatewayCertificateAuthority resource.
type GatewayCertificateAuthorityArgs struct {
	// Identifier of the certificate entity. Must be unique in the current API Management service instance.
	CertificateId pulumi.StringPtrInput
	// Gateway entity identifier. Must be unique in the current API Management service instance. Must not have value 'managed'
	GatewayId pulumi.StringInput
	// Determines whether certificate authority is trusted.
	IsTrusted pulumi.BoolPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the API Management service.
	ServiceName pulumi.StringInput
}

func (GatewayCertificateAuthorityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayCertificateAuthorityArgs)(nil)).Elem()
}

type GatewayCertificateAuthorityInput interface {
	pulumi.Input

	ToGatewayCertificateAuthorityOutput() GatewayCertificateAuthorityOutput
	ToGatewayCertificateAuthorityOutputWithContext(ctx context.Context) GatewayCertificateAuthorityOutput
}

func (*GatewayCertificateAuthority) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewayCertificateAuthority)(nil)).Elem()
}

func (i *GatewayCertificateAuthority) ToGatewayCertificateAuthorityOutput() GatewayCertificateAuthorityOutput {
	return i.ToGatewayCertificateAuthorityOutputWithContext(context.Background())
}

func (i *GatewayCertificateAuthority) ToGatewayCertificateAuthorityOutputWithContext(ctx context.Context) GatewayCertificateAuthorityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayCertificateAuthorityOutput)
}

func (i *GatewayCertificateAuthority) ToOutput(ctx context.Context) pulumix.Output[*GatewayCertificateAuthority] {
	return pulumix.Output[*GatewayCertificateAuthority]{
		OutputState: i.ToGatewayCertificateAuthorityOutputWithContext(ctx).OutputState,
	}
}

type GatewayCertificateAuthorityOutput struct{ *pulumi.OutputState }

func (GatewayCertificateAuthorityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewayCertificateAuthority)(nil)).Elem()
}

func (o GatewayCertificateAuthorityOutput) ToGatewayCertificateAuthorityOutput() GatewayCertificateAuthorityOutput {
	return o
}

func (o GatewayCertificateAuthorityOutput) ToGatewayCertificateAuthorityOutputWithContext(ctx context.Context) GatewayCertificateAuthorityOutput {
	return o
}

func (o GatewayCertificateAuthorityOutput) ToOutput(ctx context.Context) pulumix.Output[*GatewayCertificateAuthority] {
	return pulumix.Output[*GatewayCertificateAuthority]{
		OutputState: o.OutputState,
	}
}

// Determines whether certificate authority is trusted.
func (o GatewayCertificateAuthorityOutput) IsTrusted() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *GatewayCertificateAuthority) pulumi.BoolPtrOutput { return v.IsTrusted }).(pulumi.BoolPtrOutput)
}

// The name of the resource
func (o GatewayCertificateAuthorityOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayCertificateAuthority) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o GatewayCertificateAuthorityOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayCertificateAuthority) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GatewayCertificateAuthorityOutput{})
}
