// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package desktopvirtualization

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The Private Endpoint Connection resource.
// Azure REST API version: 2022-10-14-preview. Prior API version in Azure Native 1.x: 2021-04-01-preview
type PrivateEndpointConnectionByWorkspace struct {
	pulumi.CustomResourceState

	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The resource of private end point.
	PrivateEndpoint PrivateEndpointResponsePtrOutput `pulumi:"privateEndpoint"`
	// A collection of information about the state of the connection between service consumer and provider.
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponseOutput `pulumi:"privateLinkServiceConnectionState"`
	// The provisioning state of the private endpoint connection resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewPrivateEndpointConnectionByWorkspace registers a new resource with the given unique name, arguments, and options.
func NewPrivateEndpointConnectionByWorkspace(ctx *pulumi.Context,
	name string, args *PrivateEndpointConnectionByWorkspaceArgs, opts ...pulumi.ResourceOption) (*PrivateEndpointConnectionByWorkspace, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PrivateLinkServiceConnectionState == nil {
		return nil, errors.New("invalid value for required argument 'PrivateLinkServiceConnectionState'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20210401preview:PrivateEndpointConnectionByWorkspace"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20210903preview:PrivateEndpointConnectionByWorkspace"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20220210preview:PrivateEndpointConnectionByWorkspace"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20220401preview:PrivateEndpointConnectionByWorkspace"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20221014preview:PrivateEndpointConnectionByWorkspace"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20230707preview:PrivateEndpointConnectionByWorkspace"),
		},
	})
	opts = append(opts, aliases)
	var resource PrivateEndpointConnectionByWorkspace
	err := ctx.RegisterResource("azure-native:desktopvirtualization:PrivateEndpointConnectionByWorkspace", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPrivateEndpointConnectionByWorkspace gets an existing PrivateEndpointConnectionByWorkspace resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPrivateEndpointConnectionByWorkspace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateEndpointConnectionByWorkspaceState, opts ...pulumi.ResourceOption) (*PrivateEndpointConnectionByWorkspace, error) {
	var resource PrivateEndpointConnectionByWorkspace
	err := ctx.ReadResource("azure-native:desktopvirtualization:PrivateEndpointConnectionByWorkspace", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PrivateEndpointConnectionByWorkspace resources.
type privateEndpointConnectionByWorkspaceState struct {
}

type PrivateEndpointConnectionByWorkspaceState struct {
}

func (PrivateEndpointConnectionByWorkspaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*privateEndpointConnectionByWorkspaceState)(nil)).Elem()
}

type privateEndpointConnectionByWorkspaceArgs struct {
	// The name of the private endpoint connection associated with the Azure resource
	PrivateEndpointConnectionName *string `pulumi:"privateEndpointConnectionName"`
	// A collection of information about the state of the connection between service consumer and provider.
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionState `pulumi:"privateLinkServiceConnectionState"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the workspace
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a PrivateEndpointConnectionByWorkspace resource.
type PrivateEndpointConnectionByWorkspaceArgs struct {
	// The name of the private endpoint connection associated with the Azure resource
	PrivateEndpointConnectionName pulumi.StringPtrInput
	// A collection of information about the state of the connection between service consumer and provider.
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the workspace
	WorkspaceName pulumi.StringInput
}

func (PrivateEndpointConnectionByWorkspaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privateEndpointConnectionByWorkspaceArgs)(nil)).Elem()
}

type PrivateEndpointConnectionByWorkspaceInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionByWorkspaceOutput() PrivateEndpointConnectionByWorkspaceOutput
	ToPrivateEndpointConnectionByWorkspaceOutputWithContext(ctx context.Context) PrivateEndpointConnectionByWorkspaceOutput
}

func (*PrivateEndpointConnectionByWorkspace) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointConnectionByWorkspace)(nil)).Elem()
}

func (i *PrivateEndpointConnectionByWorkspace) ToPrivateEndpointConnectionByWorkspaceOutput() PrivateEndpointConnectionByWorkspaceOutput {
	return i.ToPrivateEndpointConnectionByWorkspaceOutputWithContext(context.Background())
}

func (i *PrivateEndpointConnectionByWorkspace) ToPrivateEndpointConnectionByWorkspaceOutputWithContext(ctx context.Context) PrivateEndpointConnectionByWorkspaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionByWorkspaceOutput)
}

type PrivateEndpointConnectionByWorkspaceOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionByWorkspaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointConnectionByWorkspace)(nil)).Elem()
}

func (o PrivateEndpointConnectionByWorkspaceOutput) ToPrivateEndpointConnectionByWorkspaceOutput() PrivateEndpointConnectionByWorkspaceOutput {
	return o
}

func (o PrivateEndpointConnectionByWorkspaceOutput) ToPrivateEndpointConnectionByWorkspaceOutputWithContext(ctx context.Context) PrivateEndpointConnectionByWorkspaceOutput {
	return o
}

// The name of the resource
func (o PrivateEndpointConnectionByWorkspaceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionByWorkspace) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The resource of private end point.
func (o PrivateEndpointConnectionByWorkspaceOutput) PrivateEndpoint() PrivateEndpointResponsePtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionByWorkspace) PrivateEndpointResponsePtrOutput {
		return v.PrivateEndpoint
	}).(PrivateEndpointResponsePtrOutput)
}

// A collection of information about the state of the connection between service consumer and provider.
func (o PrivateEndpointConnectionByWorkspaceOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateResponseOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionByWorkspace) PrivateLinkServiceConnectionStateResponseOutput {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateResponseOutput)
}

// The provisioning state of the private endpoint connection resource.
func (o PrivateEndpointConnectionByWorkspaceOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionByWorkspace) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o PrivateEndpointConnectionByWorkspaceOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionByWorkspace) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o PrivateEndpointConnectionByWorkspaceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionByWorkspace) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PrivateEndpointConnectionByWorkspaceOutput{})
}
