// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230606

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// REST model used to encapsulate Private Link properties for tracked resources.
type PrivateEndpointConnectionController struct {
	pulumi.CustomResourceState

	// array of group ids
	GroupIds pulumi.StringArrayOutput `pulumi:"groupIds"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// private endpoints
	PrivateEndpoint ResourceIdResponseOutput `pulumi:"privateEndpoint"`
	// private endpoints connection state
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponsePtrOutput `pulumi:"privateLinkServiceConnectionState"`
	// provisioning state enum
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewPrivateEndpointConnectionController registers a new resource with the given unique name, arguments, and options.
func NewPrivateEndpointConnectionController(ctx *pulumi.Context,
	name string, args *PrivateEndpointConnectionControllerArgs, opts ...pulumi.ResourceOption) (*PrivateEndpointConnectionController, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SiteName == nil {
		return nil, errors.New("invalid value for required argument 'SiteName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:offazure:PrivateEndpointConnectionController"),
		},
		{
			Type: pulumi.String("azure-native:offazure/v20200707:PrivateEndpointConnectionController"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource PrivateEndpointConnectionController
	err := ctx.RegisterResource("azure-native:offazure/v20230606:PrivateEndpointConnectionController", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPrivateEndpointConnectionController gets an existing PrivateEndpointConnectionController resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPrivateEndpointConnectionController(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateEndpointConnectionControllerState, opts ...pulumi.ResourceOption) (*PrivateEndpointConnectionController, error) {
	var resource PrivateEndpointConnectionController
	err := ctx.ReadResource("azure-native:offazure/v20230606:PrivateEndpointConnectionController", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PrivateEndpointConnectionController resources.
type privateEndpointConnectionControllerState struct {
}

type PrivateEndpointConnectionControllerState struct {
}

func (PrivateEndpointConnectionControllerState) ElementType() reflect.Type {
	return reflect.TypeOf((*privateEndpointConnectionControllerState)(nil)).Elem()
}

type privateEndpointConnectionControllerArgs struct {
	//  Private link resource name.
	PeConnectionName *string `pulumi:"peConnectionName"`
	// private endpoints connection state
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionState `pulumi:"privateLinkServiceConnectionState"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Site name
	SiteName string `pulumi:"siteName"`
}

// The set of arguments for constructing a PrivateEndpointConnectionController resource.
type PrivateEndpointConnectionControllerArgs struct {
	//  Private link resource name.
	PeConnectionName pulumi.StringPtrInput
	// private endpoints connection state
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStatePtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Site name
	SiteName pulumi.StringInput
}

func (PrivateEndpointConnectionControllerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privateEndpointConnectionControllerArgs)(nil)).Elem()
}

type PrivateEndpointConnectionControllerInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionControllerOutput() PrivateEndpointConnectionControllerOutput
	ToPrivateEndpointConnectionControllerOutputWithContext(ctx context.Context) PrivateEndpointConnectionControllerOutput
}

func (*PrivateEndpointConnectionController) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointConnectionController)(nil)).Elem()
}

func (i *PrivateEndpointConnectionController) ToPrivateEndpointConnectionControllerOutput() PrivateEndpointConnectionControllerOutput {
	return i.ToPrivateEndpointConnectionControllerOutputWithContext(context.Background())
}

func (i *PrivateEndpointConnectionController) ToPrivateEndpointConnectionControllerOutputWithContext(ctx context.Context) PrivateEndpointConnectionControllerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionControllerOutput)
}

func (i *PrivateEndpointConnectionController) ToOutput(ctx context.Context) pulumix.Output[*PrivateEndpointConnectionController] {
	return pulumix.Output[*PrivateEndpointConnectionController]{
		OutputState: i.ToPrivateEndpointConnectionControllerOutputWithContext(ctx).OutputState,
	}
}

type PrivateEndpointConnectionControllerOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionControllerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointConnectionController)(nil)).Elem()
}

func (o PrivateEndpointConnectionControllerOutput) ToPrivateEndpointConnectionControllerOutput() PrivateEndpointConnectionControllerOutput {
	return o
}

func (o PrivateEndpointConnectionControllerOutput) ToPrivateEndpointConnectionControllerOutputWithContext(ctx context.Context) PrivateEndpointConnectionControllerOutput {
	return o
}

func (o PrivateEndpointConnectionControllerOutput) ToOutput(ctx context.Context) pulumix.Output[*PrivateEndpointConnectionController] {
	return pulumix.Output[*PrivateEndpointConnectionController]{
		OutputState: o.OutputState,
	}
}

// array of group ids
func (o PrivateEndpointConnectionControllerOutput) GroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionController) pulumi.StringArrayOutput { return v.GroupIds }).(pulumi.StringArrayOutput)
}

// The name of the resource
func (o PrivateEndpointConnectionControllerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionController) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// private endpoints
func (o PrivateEndpointConnectionControllerOutput) PrivateEndpoint() ResourceIdResponseOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionController) ResourceIdResponseOutput { return v.PrivateEndpoint }).(ResourceIdResponseOutput)
}

// private endpoints connection state
func (o PrivateEndpointConnectionControllerOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionController) PrivateLinkServiceConnectionStateResponsePtrOutput {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateResponsePtrOutput)
}

// provisioning state enum
func (o PrivateEndpointConnectionControllerOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionController) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o PrivateEndpointConnectionControllerOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionController) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o PrivateEndpointConnectionControllerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionController) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PrivateEndpointConnectionControllerOutput{})
}
