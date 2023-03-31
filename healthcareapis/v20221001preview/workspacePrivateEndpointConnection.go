// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20221001preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The Private Endpoint Connection resource.
type WorkspacePrivateEndpointConnection struct {
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

// NewWorkspacePrivateEndpointConnection registers a new resource with the given unique name, arguments, and options.
func NewWorkspacePrivateEndpointConnection(ctx *pulumi.Context,
	name string, args *WorkspacePrivateEndpointConnectionArgs, opts ...pulumi.ResourceOption) (*WorkspacePrivateEndpointConnection, error) {
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
			Type: pulumi.String("azure-native:healthcareapis:WorkspacePrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:healthcareapis/v20211101:WorkspacePrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:healthcareapis/v20220131preview:WorkspacePrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:healthcareapis/v20220515:WorkspacePrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:healthcareapis/v20220601:WorkspacePrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:healthcareapis/v20221201:WorkspacePrivateEndpointConnection"),
		},
	})
	opts = append(opts, aliases)
	var resource WorkspacePrivateEndpointConnection
	err := ctx.RegisterResource("azure-native:healthcareapis/v20221001preview:WorkspacePrivateEndpointConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkspacePrivateEndpointConnection gets an existing WorkspacePrivateEndpointConnection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkspacePrivateEndpointConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkspacePrivateEndpointConnectionState, opts ...pulumi.ResourceOption) (*WorkspacePrivateEndpointConnection, error) {
	var resource WorkspacePrivateEndpointConnection
	err := ctx.ReadResource("azure-native:healthcareapis/v20221001preview:WorkspacePrivateEndpointConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WorkspacePrivateEndpointConnection resources.
type workspacePrivateEndpointConnectionState struct {
}

type WorkspacePrivateEndpointConnectionState struct {
}

func (WorkspacePrivateEndpointConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*workspacePrivateEndpointConnectionState)(nil)).Elem()
}

type workspacePrivateEndpointConnectionArgs struct {
	// The name of the private endpoint connection associated with the Azure resource
	PrivateEndpointConnectionName *string `pulumi:"privateEndpointConnectionName"`
	// A collection of information about the state of the connection between service consumer and provider.
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionState `pulumi:"privateLinkServiceConnectionState"`
	// The name of the resource group that contains the service instance.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of workspace resource.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a WorkspacePrivateEndpointConnection resource.
type WorkspacePrivateEndpointConnectionArgs struct {
	// The name of the private endpoint connection associated with the Azure resource
	PrivateEndpointConnectionName pulumi.StringPtrInput
	// A collection of information about the state of the connection between service consumer and provider.
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateInput
	// The name of the resource group that contains the service instance.
	ResourceGroupName pulumi.StringInput
	// The name of workspace resource.
	WorkspaceName pulumi.StringInput
}

func (WorkspacePrivateEndpointConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workspacePrivateEndpointConnectionArgs)(nil)).Elem()
}

type WorkspacePrivateEndpointConnectionInput interface {
	pulumi.Input

	ToWorkspacePrivateEndpointConnectionOutput() WorkspacePrivateEndpointConnectionOutput
	ToWorkspacePrivateEndpointConnectionOutputWithContext(ctx context.Context) WorkspacePrivateEndpointConnectionOutput
}

func (*WorkspacePrivateEndpointConnection) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspacePrivateEndpointConnection)(nil)).Elem()
}

func (i *WorkspacePrivateEndpointConnection) ToWorkspacePrivateEndpointConnectionOutput() WorkspacePrivateEndpointConnectionOutput {
	return i.ToWorkspacePrivateEndpointConnectionOutputWithContext(context.Background())
}

func (i *WorkspacePrivateEndpointConnection) ToWorkspacePrivateEndpointConnectionOutputWithContext(ctx context.Context) WorkspacePrivateEndpointConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspacePrivateEndpointConnectionOutput)
}

type WorkspacePrivateEndpointConnectionOutput struct{ *pulumi.OutputState }

func (WorkspacePrivateEndpointConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspacePrivateEndpointConnection)(nil)).Elem()
}

func (o WorkspacePrivateEndpointConnectionOutput) ToWorkspacePrivateEndpointConnectionOutput() WorkspacePrivateEndpointConnectionOutput {
	return o
}

func (o WorkspacePrivateEndpointConnectionOutput) ToWorkspacePrivateEndpointConnectionOutputWithContext(ctx context.Context) WorkspacePrivateEndpointConnectionOutput {
	return o
}

// The name of the resource
func (o WorkspacePrivateEndpointConnectionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkspacePrivateEndpointConnection) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The resource of private end point.
func (o WorkspacePrivateEndpointConnectionOutput) PrivateEndpoint() PrivateEndpointResponsePtrOutput {
	return o.ApplyT(func(v *WorkspacePrivateEndpointConnection) PrivateEndpointResponsePtrOutput { return v.PrivateEndpoint }).(PrivateEndpointResponsePtrOutput)
}

// A collection of information about the state of the connection between service consumer and provider.
func (o WorkspacePrivateEndpointConnectionOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateResponseOutput {
	return o.ApplyT(func(v *WorkspacePrivateEndpointConnection) PrivateLinkServiceConnectionStateResponseOutput {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateResponseOutput)
}

// The provisioning state of the private endpoint connection resource.
func (o WorkspacePrivateEndpointConnectionOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkspacePrivateEndpointConnection) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o WorkspacePrivateEndpointConnectionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *WorkspacePrivateEndpointConnection) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o WorkspacePrivateEndpointConnectionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkspacePrivateEndpointConnection) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WorkspacePrivateEndpointConnectionOutput{})
}
