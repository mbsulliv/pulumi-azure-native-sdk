// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package m365securityandcompliance

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The Private Endpoint Connection resource.
// Azure REST API version: 2021-03-25-preview. Prior API version in Azure Native 1.x: 2021-03-25-preview.
type PrivateEndpointConnectionsForSCCPowershell struct {
	pulumi.CustomResourceState

	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The resource of private end point.
	PrivateEndpoint PrivateEndpointResponsePtrOutput `pulumi:"privateEndpoint"`
	// A collection of information about the state of the connection between service consumer and provider.
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponseOutput `pulumi:"privateLinkServiceConnectionState"`
	// The provisioning state of the private endpoint connection resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Required property for system data
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewPrivateEndpointConnectionsForSCCPowershell registers a new resource with the given unique name, arguments, and options.
func NewPrivateEndpointConnectionsForSCCPowershell(ctx *pulumi.Context,
	name string, args *PrivateEndpointConnectionsForSCCPowershellArgs, opts ...pulumi.ResourceOption) (*PrivateEndpointConnectionsForSCCPowershell, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PrivateLinkServiceConnectionState == nil {
		return nil, errors.New("invalid value for required argument 'PrivateLinkServiceConnectionState'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ResourceName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:m365securityandcompliance/v20210325preview:PrivateEndpointConnectionsForSCCPowershell"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource PrivateEndpointConnectionsForSCCPowershell
	err := ctx.RegisterResource("azure-native:m365securityandcompliance:PrivateEndpointConnectionsForSCCPowershell", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPrivateEndpointConnectionsForSCCPowershell gets an existing PrivateEndpointConnectionsForSCCPowershell resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPrivateEndpointConnectionsForSCCPowershell(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateEndpointConnectionsForSCCPowershellState, opts ...pulumi.ResourceOption) (*PrivateEndpointConnectionsForSCCPowershell, error) {
	var resource PrivateEndpointConnectionsForSCCPowershell
	err := ctx.ReadResource("azure-native:m365securityandcompliance:PrivateEndpointConnectionsForSCCPowershell", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PrivateEndpointConnectionsForSCCPowershell resources.
type privateEndpointConnectionsForSCCPowershellState struct {
}

type PrivateEndpointConnectionsForSCCPowershellState struct {
}

func (PrivateEndpointConnectionsForSCCPowershellState) ElementType() reflect.Type {
	return reflect.TypeOf((*privateEndpointConnectionsForSCCPowershellState)(nil)).Elem()
}

type privateEndpointConnectionsForSCCPowershellArgs struct {
	// The name of the private endpoint connection associated with the Azure resource
	PrivateEndpointConnectionName *string `pulumi:"privateEndpointConnectionName"`
	// A collection of information about the state of the connection between service consumer and provider.
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionState `pulumi:"privateLinkServiceConnectionState"`
	// The name of the resource group that contains the service instance.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the service instance.
	ResourceName string `pulumi:"resourceName"`
}

// The set of arguments for constructing a PrivateEndpointConnectionsForSCCPowershell resource.
type PrivateEndpointConnectionsForSCCPowershellArgs struct {
	// The name of the private endpoint connection associated with the Azure resource
	PrivateEndpointConnectionName pulumi.StringPtrInput
	// A collection of information about the state of the connection between service consumer and provider.
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateInput
	// The name of the resource group that contains the service instance.
	ResourceGroupName pulumi.StringInput
	// The name of the service instance.
	ResourceName pulumi.StringInput
}

func (PrivateEndpointConnectionsForSCCPowershellArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privateEndpointConnectionsForSCCPowershellArgs)(nil)).Elem()
}

type PrivateEndpointConnectionsForSCCPowershellInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionsForSCCPowershellOutput() PrivateEndpointConnectionsForSCCPowershellOutput
	ToPrivateEndpointConnectionsForSCCPowershellOutputWithContext(ctx context.Context) PrivateEndpointConnectionsForSCCPowershellOutput
}

func (*PrivateEndpointConnectionsForSCCPowershell) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointConnectionsForSCCPowershell)(nil)).Elem()
}

func (i *PrivateEndpointConnectionsForSCCPowershell) ToPrivateEndpointConnectionsForSCCPowershellOutput() PrivateEndpointConnectionsForSCCPowershellOutput {
	return i.ToPrivateEndpointConnectionsForSCCPowershellOutputWithContext(context.Background())
}

func (i *PrivateEndpointConnectionsForSCCPowershell) ToPrivateEndpointConnectionsForSCCPowershellOutputWithContext(ctx context.Context) PrivateEndpointConnectionsForSCCPowershellOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionsForSCCPowershellOutput)
}

type PrivateEndpointConnectionsForSCCPowershellOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionsForSCCPowershellOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointConnectionsForSCCPowershell)(nil)).Elem()
}

func (o PrivateEndpointConnectionsForSCCPowershellOutput) ToPrivateEndpointConnectionsForSCCPowershellOutput() PrivateEndpointConnectionsForSCCPowershellOutput {
	return o
}

func (o PrivateEndpointConnectionsForSCCPowershellOutput) ToPrivateEndpointConnectionsForSCCPowershellOutputWithContext(ctx context.Context) PrivateEndpointConnectionsForSCCPowershellOutput {
	return o
}

// The name of the resource
func (o PrivateEndpointConnectionsForSCCPowershellOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionsForSCCPowershell) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The resource of private end point.
func (o PrivateEndpointConnectionsForSCCPowershellOutput) PrivateEndpoint() PrivateEndpointResponsePtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionsForSCCPowershell) PrivateEndpointResponsePtrOutput {
		return v.PrivateEndpoint
	}).(PrivateEndpointResponsePtrOutput)
}

// A collection of information about the state of the connection between service consumer and provider.
func (o PrivateEndpointConnectionsForSCCPowershellOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateResponseOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionsForSCCPowershell) PrivateLinkServiceConnectionStateResponseOutput {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateResponseOutput)
}

// The provisioning state of the private endpoint connection resource.
func (o PrivateEndpointConnectionsForSCCPowershellOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionsForSCCPowershell) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Required property for system data
func (o PrivateEndpointConnectionsForSCCPowershellOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionsForSCCPowershell) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o PrivateEndpointConnectionsForSCCPowershellOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionsForSCCPowershell) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PrivateEndpointConnectionsForSCCPowershellOutput{})
}
