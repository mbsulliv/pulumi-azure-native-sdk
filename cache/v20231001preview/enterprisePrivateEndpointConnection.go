// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20231001preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The Private Endpoint Connection resource.
type EnterprisePrivateEndpointConnection struct {
	pulumi.CustomResourceState

	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The resource of private end point.
	PrivateEndpoint PrivateEndpointResponsePtrOutput `pulumi:"privateEndpoint"`
	// A collection of information about the state of the connection between service consumer and provider.
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponseOutput `pulumi:"privateLinkServiceConnectionState"`
	// The provisioning state of the private endpoint connection resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewEnterprisePrivateEndpointConnection registers a new resource with the given unique name, arguments, and options.
func NewEnterprisePrivateEndpointConnection(ctx *pulumi.Context,
	name string, args *EnterprisePrivateEndpointConnectionArgs, opts ...pulumi.ResourceOption) (*EnterprisePrivateEndpointConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterName == nil {
		return nil, errors.New("invalid value for required argument 'ClusterName'")
	}
	if args.PrivateLinkServiceConnectionState == nil {
		return nil, errors.New("invalid value for required argument 'PrivateLinkServiceConnectionState'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:cache:EnterprisePrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20201001preview:EnterprisePrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20210201preview:EnterprisePrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20210301:EnterprisePrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20210801:EnterprisePrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20220101:EnterprisePrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20221101preview:EnterprisePrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20230301preview:EnterprisePrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20230701:EnterprisePrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20230801preview:EnterprisePrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:cache/v20231101:EnterprisePrivateEndpointConnection"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource EnterprisePrivateEndpointConnection
	err := ctx.RegisterResource("azure-native:cache/v20231001preview:EnterprisePrivateEndpointConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEnterprisePrivateEndpointConnection gets an existing EnterprisePrivateEndpointConnection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEnterprisePrivateEndpointConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EnterprisePrivateEndpointConnectionState, opts ...pulumi.ResourceOption) (*EnterprisePrivateEndpointConnection, error) {
	var resource EnterprisePrivateEndpointConnection
	err := ctx.ReadResource("azure-native:cache/v20231001preview:EnterprisePrivateEndpointConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EnterprisePrivateEndpointConnection resources.
type enterprisePrivateEndpointConnectionState struct {
}

type EnterprisePrivateEndpointConnectionState struct {
}

func (EnterprisePrivateEndpointConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*enterprisePrivateEndpointConnectionState)(nil)).Elem()
}

type enterprisePrivateEndpointConnectionArgs struct {
	// The name of the RedisEnterprise cluster.
	ClusterName string `pulumi:"clusterName"`
	// The name of the private endpoint connection associated with the Azure resource
	PrivateEndpointConnectionName *string `pulumi:"privateEndpointConnectionName"`
	// A collection of information about the state of the connection between service consumer and provider.
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionState `pulumi:"privateLinkServiceConnectionState"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a EnterprisePrivateEndpointConnection resource.
type EnterprisePrivateEndpointConnectionArgs struct {
	// The name of the RedisEnterprise cluster.
	ClusterName pulumi.StringInput
	// The name of the private endpoint connection associated with the Azure resource
	PrivateEndpointConnectionName pulumi.StringPtrInput
	// A collection of information about the state of the connection between service consumer and provider.
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
}

func (EnterprisePrivateEndpointConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*enterprisePrivateEndpointConnectionArgs)(nil)).Elem()
}

type EnterprisePrivateEndpointConnectionInput interface {
	pulumi.Input

	ToEnterprisePrivateEndpointConnectionOutput() EnterprisePrivateEndpointConnectionOutput
	ToEnterprisePrivateEndpointConnectionOutputWithContext(ctx context.Context) EnterprisePrivateEndpointConnectionOutput
}

func (*EnterprisePrivateEndpointConnection) ElementType() reflect.Type {
	return reflect.TypeOf((**EnterprisePrivateEndpointConnection)(nil)).Elem()
}

func (i *EnterprisePrivateEndpointConnection) ToEnterprisePrivateEndpointConnectionOutput() EnterprisePrivateEndpointConnectionOutput {
	return i.ToEnterprisePrivateEndpointConnectionOutputWithContext(context.Background())
}

func (i *EnterprisePrivateEndpointConnection) ToEnterprisePrivateEndpointConnectionOutputWithContext(ctx context.Context) EnterprisePrivateEndpointConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnterprisePrivateEndpointConnectionOutput)
}

type EnterprisePrivateEndpointConnectionOutput struct{ *pulumi.OutputState }

func (EnterprisePrivateEndpointConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EnterprisePrivateEndpointConnection)(nil)).Elem()
}

func (o EnterprisePrivateEndpointConnectionOutput) ToEnterprisePrivateEndpointConnectionOutput() EnterprisePrivateEndpointConnectionOutput {
	return o
}

func (o EnterprisePrivateEndpointConnectionOutput) ToEnterprisePrivateEndpointConnectionOutputWithContext(ctx context.Context) EnterprisePrivateEndpointConnectionOutput {
	return o
}

// The name of the resource
func (o EnterprisePrivateEndpointConnectionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *EnterprisePrivateEndpointConnection) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The resource of private end point.
func (o EnterprisePrivateEndpointConnectionOutput) PrivateEndpoint() PrivateEndpointResponsePtrOutput {
	return o.ApplyT(func(v *EnterprisePrivateEndpointConnection) PrivateEndpointResponsePtrOutput {
		return v.PrivateEndpoint
	}).(PrivateEndpointResponsePtrOutput)
}

// A collection of information about the state of the connection between service consumer and provider.
func (o EnterprisePrivateEndpointConnectionOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateResponseOutput {
	return o.ApplyT(func(v *EnterprisePrivateEndpointConnection) PrivateLinkServiceConnectionStateResponseOutput {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateResponseOutput)
}

// The provisioning state of the private endpoint connection resource.
func (o EnterprisePrivateEndpointConnectionOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *EnterprisePrivateEndpointConnection) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o EnterprisePrivateEndpointConnectionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *EnterprisePrivateEndpointConnection) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o EnterprisePrivateEndpointConnectionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *EnterprisePrivateEndpointConnection) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(EnterprisePrivateEndpointConnectionOutput{})
}
