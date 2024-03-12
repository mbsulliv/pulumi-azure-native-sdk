// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package search

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Describes an existing Private Endpoint connection to the Azure Cognitive Search service.
// Azure REST API version: 2022-09-01. Prior API version in Azure Native 1.x: 2020-08-01.
//
// Other available API versions: 2019-10-01-preview, 2023-11-01, 2024-03-01-preview.
type PrivateEndpointConnection struct {
	pulumi.CustomResourceState

	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Describes the properties of an existing Private Endpoint connection to the Azure Cognitive Search service.
	Properties PrivateEndpointConnectionPropertiesResponseOutput `pulumi:"properties"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
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
	if args.SearchServiceName == nil {
		return nil, errors.New("invalid value for required argument 'SearchServiceName'")
	}
	if args.Properties != nil {
		args.Properties = args.Properties.ToPrivateEndpointConnectionPropertiesPtrOutput().ApplyT(func(v *PrivateEndpointConnectionProperties) *PrivateEndpointConnectionProperties { return v.Defaults() }).(PrivateEndpointConnectionPropertiesPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:search/v20191001preview:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:search/v20200313:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:search/v20200801:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:search/v20200801preview:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:search/v20210401preview:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:search/v20220901:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:search/v20231101:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:search/v20240301preview:PrivateEndpointConnection"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource PrivateEndpointConnection
	err := ctx.RegisterResource("azure-native:search:PrivateEndpointConnection", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure-native:search:PrivateEndpointConnection", name, id, state, &resource, opts...)
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
	// The name of the private endpoint connection to the Azure Cognitive Search service with the specified resource group.
	PrivateEndpointConnectionName *string `pulumi:"privateEndpointConnectionName"`
	// Describes the properties of an existing Private Endpoint connection to the Azure Cognitive Search service.
	Properties *PrivateEndpointConnectionProperties `pulumi:"properties"`
	// The name of the resource group within the current subscription. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Azure Cognitive Search service associated with the specified resource group.
	SearchServiceName string `pulumi:"searchServiceName"`
}

// The set of arguments for constructing a PrivateEndpointConnection resource.
type PrivateEndpointConnectionArgs struct {
	// The name of the private endpoint connection to the Azure Cognitive Search service with the specified resource group.
	PrivateEndpointConnectionName pulumi.StringPtrInput
	// Describes the properties of an existing Private Endpoint connection to the Azure Cognitive Search service.
	Properties PrivateEndpointConnectionPropertiesPtrInput
	// The name of the resource group within the current subscription. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// The name of the Azure Cognitive Search service associated with the specified resource group.
	SearchServiceName pulumi.StringInput
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

// The name of the resource
func (o PrivateEndpointConnectionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateEndpointConnection) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Describes the properties of an existing Private Endpoint connection to the Azure Cognitive Search service.
func (o PrivateEndpointConnectionOutput) Properties() PrivateEndpointConnectionPropertiesResponseOutput {
	return o.ApplyT(func(v *PrivateEndpointConnection) PrivateEndpointConnectionPropertiesResponseOutput {
		return v.Properties
	}).(PrivateEndpointConnectionPropertiesResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o PrivateEndpointConnectionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateEndpointConnection) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PrivateEndpointConnectionOutput{})
}
