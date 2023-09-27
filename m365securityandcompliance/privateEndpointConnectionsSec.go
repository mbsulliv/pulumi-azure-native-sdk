// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package m365securityandcompliance

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The Private Endpoint Connection resource.
// Azure REST API version: 2021-03-25-preview. Prior API version in Azure Native 1.x: 2021-03-25-preview
type PrivateEndpointConnectionsSec struct {
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

// NewPrivateEndpointConnectionsSec registers a new resource with the given unique name, arguments, and options.
func NewPrivateEndpointConnectionsSec(ctx *pulumi.Context,
	name string, args *PrivateEndpointConnectionsSecArgs, opts ...pulumi.ResourceOption) (*PrivateEndpointConnectionsSec, error) {
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
			Type: pulumi.String("azure-native:m365securityandcompliance/v20210325preview:PrivateEndpointConnectionsSec"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource PrivateEndpointConnectionsSec
	err := ctx.RegisterResource("azure-native:m365securityandcompliance:PrivateEndpointConnectionsSec", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPrivateEndpointConnectionsSec gets an existing PrivateEndpointConnectionsSec resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPrivateEndpointConnectionsSec(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateEndpointConnectionsSecState, opts ...pulumi.ResourceOption) (*PrivateEndpointConnectionsSec, error) {
	var resource PrivateEndpointConnectionsSec
	err := ctx.ReadResource("azure-native:m365securityandcompliance:PrivateEndpointConnectionsSec", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PrivateEndpointConnectionsSec resources.
type privateEndpointConnectionsSecState struct {
}

type PrivateEndpointConnectionsSecState struct {
}

func (PrivateEndpointConnectionsSecState) ElementType() reflect.Type {
	return reflect.TypeOf((*privateEndpointConnectionsSecState)(nil)).Elem()
}

type privateEndpointConnectionsSecArgs struct {
	// The name of the private endpoint connection associated with the Azure resource
	PrivateEndpointConnectionName *string `pulumi:"privateEndpointConnectionName"`
	// A collection of information about the state of the connection between service consumer and provider.
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionState `pulumi:"privateLinkServiceConnectionState"`
	// The name of the resource group that contains the service instance.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the service instance.
	ResourceName string `pulumi:"resourceName"`
}

// The set of arguments for constructing a PrivateEndpointConnectionsSec resource.
type PrivateEndpointConnectionsSecArgs struct {
	// The name of the private endpoint connection associated with the Azure resource
	PrivateEndpointConnectionName pulumi.StringPtrInput
	// A collection of information about the state of the connection between service consumer and provider.
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateInput
	// The name of the resource group that contains the service instance.
	ResourceGroupName pulumi.StringInput
	// The name of the service instance.
	ResourceName pulumi.StringInput
}

func (PrivateEndpointConnectionsSecArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privateEndpointConnectionsSecArgs)(nil)).Elem()
}

type PrivateEndpointConnectionsSecInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionsSecOutput() PrivateEndpointConnectionsSecOutput
	ToPrivateEndpointConnectionsSecOutputWithContext(ctx context.Context) PrivateEndpointConnectionsSecOutput
}

func (*PrivateEndpointConnectionsSec) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointConnectionsSec)(nil)).Elem()
}

func (i *PrivateEndpointConnectionsSec) ToPrivateEndpointConnectionsSecOutput() PrivateEndpointConnectionsSecOutput {
	return i.ToPrivateEndpointConnectionsSecOutputWithContext(context.Background())
}

func (i *PrivateEndpointConnectionsSec) ToPrivateEndpointConnectionsSecOutputWithContext(ctx context.Context) PrivateEndpointConnectionsSecOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionsSecOutput)
}

func (i *PrivateEndpointConnectionsSec) ToOutput(ctx context.Context) pulumix.Output[*PrivateEndpointConnectionsSec] {
	return pulumix.Output[*PrivateEndpointConnectionsSec]{
		OutputState: i.ToPrivateEndpointConnectionsSecOutputWithContext(ctx).OutputState,
	}
}

type PrivateEndpointConnectionsSecOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionsSecOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointConnectionsSec)(nil)).Elem()
}

func (o PrivateEndpointConnectionsSecOutput) ToPrivateEndpointConnectionsSecOutput() PrivateEndpointConnectionsSecOutput {
	return o
}

func (o PrivateEndpointConnectionsSecOutput) ToPrivateEndpointConnectionsSecOutputWithContext(ctx context.Context) PrivateEndpointConnectionsSecOutput {
	return o
}

func (o PrivateEndpointConnectionsSecOutput) ToOutput(ctx context.Context) pulumix.Output[*PrivateEndpointConnectionsSec] {
	return pulumix.Output[*PrivateEndpointConnectionsSec]{
		OutputState: o.OutputState,
	}
}

// The name of the resource
func (o PrivateEndpointConnectionsSecOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionsSec) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The resource of private end point.
func (o PrivateEndpointConnectionsSecOutput) PrivateEndpoint() PrivateEndpointResponsePtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionsSec) PrivateEndpointResponsePtrOutput { return v.PrivateEndpoint }).(PrivateEndpointResponsePtrOutput)
}

// A collection of information about the state of the connection between service consumer and provider.
func (o PrivateEndpointConnectionsSecOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateResponseOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionsSec) PrivateLinkServiceConnectionStateResponseOutput {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateResponseOutput)
}

// The provisioning state of the private endpoint connection resource.
func (o PrivateEndpointConnectionsSecOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionsSec) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Required property for system data
func (o PrivateEndpointConnectionsSecOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionsSec) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o PrivateEndpointConnectionsSecOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionsSec) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PrivateEndpointConnectionsSecOutput{})
}
