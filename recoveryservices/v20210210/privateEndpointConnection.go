// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210210

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Private Endpoint Connection Response Properties
type PrivateEndpointConnection struct {
	pulumi.CustomResourceState

	// Optional ETag.
	ETag pulumi.StringPtrOutput `pulumi:"eTag"`
	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource name associated with the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// PrivateEndpointConnectionResource properties
	Properties PrivateEndpointConnectionResponseOutput `pulumi:"properties"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type represents the complete path of the form Namespace/ResourceType/ResourceType/...
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewPrivateEndpointConnection registers a new resource with the given unique name, arguments, and options.
func NewPrivateEndpointConnection(ctx *pulumi.Context,
	name string, args *PrivateEndpointConnectionArgs, opts ...pulumi.ResourceOption) (*PrivateEndpointConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.VaultName == nil {
		return nil, errors.New("invalid value for required argument 'VaultName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:recoveryservices:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20200202:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20201001:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20201201:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210101:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210201:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210201preview:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210301:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210401:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210601:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210701:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20210801:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20211001:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20211201:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220101:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220201:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220301:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220401:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220601preview:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20220901preview:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:recoveryservices/v20221001:PrivateEndpointConnection"),
		},
	})
	opts = append(opts, aliases)
	var resource PrivateEndpointConnection
	err := ctx.RegisterResource("azure-native:recoveryservices/v20210210:PrivateEndpointConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPrivateEndpointConnection gets an existing PrivateEndpointConnection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPrivateEndpointConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateEndpointConnectionState, opts ...pulumi.ResourceOption) (*PrivateEndpointConnection, error) {
	var resource PrivateEndpointConnection
	err := ctx.ReadResource("azure-native:recoveryservices/v20210210:PrivateEndpointConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PrivateEndpointConnection resources.
type privateEndpointConnectionState struct {
}

type PrivateEndpointConnectionState struct {
}

func (PrivateEndpointConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*privateEndpointConnectionState)(nil)).Elem()
}

type privateEndpointConnectionArgs struct {
	// Optional ETag.
	ETag *string `pulumi:"eTag"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The name of the private endpoint connection.
	PrivateEndpointConnectionName *string `pulumi:"privateEndpointConnectionName"`
	// PrivateEndpointConnectionResource properties
	Properties *PrivateEndpointConnectionType `pulumi:"properties"`
	// The name of the resource group where the recovery services vault is present.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The name of the recovery services vault.
	VaultName string `pulumi:"vaultName"`
}

// The set of arguments for constructing a PrivateEndpointConnection resource.
type PrivateEndpointConnectionArgs struct {
	// Optional ETag.
	ETag pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// The name of the private endpoint connection.
	PrivateEndpointConnectionName pulumi.StringPtrInput
	// PrivateEndpointConnectionResource properties
	Properties PrivateEndpointConnectionTypePtrInput
	// The name of the resource group where the recovery services vault is present.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The name of the recovery services vault.
	VaultName pulumi.StringInput
}

func (PrivateEndpointConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privateEndpointConnectionArgs)(nil)).Elem()
}

type PrivateEndpointConnectionInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionOutput() PrivateEndpointConnectionOutput
	ToPrivateEndpointConnectionOutputWithContext(ctx context.Context) PrivateEndpointConnectionOutput
}

func (*PrivateEndpointConnection) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointConnection)(nil)).Elem()
}

func (i *PrivateEndpointConnection) ToPrivateEndpointConnectionOutput() PrivateEndpointConnectionOutput {
	return i.ToPrivateEndpointConnectionOutputWithContext(context.Background())
}

func (i *PrivateEndpointConnection) ToPrivateEndpointConnectionOutputWithContext(ctx context.Context) PrivateEndpointConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionOutput)
}

type PrivateEndpointConnectionOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointConnection)(nil)).Elem()
}

func (o PrivateEndpointConnectionOutput) ToPrivateEndpointConnectionOutput() PrivateEndpointConnectionOutput {
	return o
}

func (o PrivateEndpointConnectionOutput) ToPrivateEndpointConnectionOutputWithContext(ctx context.Context) PrivateEndpointConnectionOutput {
	return o
}

// Optional ETag.
func (o PrivateEndpointConnectionOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnection) pulumi.StringPtrOutput { return v.ETag }).(pulumi.StringPtrOutput)
}

// Resource location.
func (o PrivateEndpointConnectionOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnection) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource name associated with the resource.
func (o PrivateEndpointConnectionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateEndpointConnection) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// PrivateEndpointConnectionResource properties
func (o PrivateEndpointConnectionOutput) Properties() PrivateEndpointConnectionResponseOutput {
	return o.ApplyT(func(v *PrivateEndpointConnection) PrivateEndpointConnectionResponseOutput { return v.Properties }).(PrivateEndpointConnectionResponseOutput)
}

// Resource tags.
func (o PrivateEndpointConnectionOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *PrivateEndpointConnection) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type represents the complete path of the form Namespace/ResourceType/ResourceType/...
func (o PrivateEndpointConnectionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateEndpointConnection) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PrivateEndpointConnectionOutput{})
}
