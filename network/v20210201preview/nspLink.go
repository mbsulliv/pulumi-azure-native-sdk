// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210201preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The network security perimeter link resource
type NspLink struct {
	pulumi.CustomResourceState

	// Perimeter ARM Id for the remote NSP with which the link gets created in Auto-approval mode. It should be used when the NSP admin have Microsoft.Network/networkSecurityPerimeters/linkPerimeter/action permission on the remote NSP resource.
	AutoApprovedRemotePerimeterResourceId pulumi.StringPtrOutput `pulumi:"autoApprovedRemotePerimeterResourceId"`
	// A message passed to the owner of the remote NSP link resource with this connection request. In case of Auto-approved flow, it is default to 'Auto Approved'. Restricted to 140 chars.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Local Inbound profile names to which Inbound is allowed. Use ['*'] to allow inbound to all profiles. It's default value is ['*'].
	LocalInboundProfiles pulumi.StringArrayOutput `pulumi:"localInboundProfiles"`
	// Local Outbound profile names from which Outbound is allowed. In current version, it is readonly property and it's value is set to ['*'] to allow outbound from all profiles. In later version, user will be able to modify it.
	LocalOutboundProfiles pulumi.StringArrayOutput `pulumi:"localOutboundProfiles"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioning state of the NSP Link resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Remote Inbound profile names to which Inbound is allowed. Use ['*'] to allow inbound to all profiles. This property can only be updated in auto-approval mode. It's default value is ['*'].
	RemoteInboundProfiles pulumi.StringArrayOutput `pulumi:"remoteInboundProfiles"`
	// Remote Outbound profile names from which Outbound is allowed. In current version, it is readonly property and it's value is set to ['*'] to allow outbound from all profiles. In later version, user will be able to modify it.
	RemoteOutboundProfiles pulumi.StringArrayOutput `pulumi:"remoteOutboundProfiles"`
	// Remote NSP Guid with which the link gets created.
	RemotePerimeterGuid pulumi.StringOutput `pulumi:"remotePerimeterGuid"`
	// Remote NSP location with which the link gets created.
	RemotePerimeterLocation pulumi.StringOutput `pulumi:"remotePerimeterLocation"`
	// The NSP link state.
	Status pulumi.StringOutput `pulumi:"status"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewNspLink registers a new resource with the given unique name, arguments, and options.
func NewNspLink(ctx *pulumi.Context,
	name string, args *NspLinkArgs, opts ...pulumi.ResourceOption) (*NspLink, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NetworkSecurityPerimeterName == nil {
		return nil, errors.New("invalid value for required argument 'NetworkSecurityPerimeterName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:NspLink"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource NspLink
	err := ctx.RegisterResource("azure-native:network/v20210201preview:NspLink", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNspLink gets an existing NspLink resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNspLink(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NspLinkState, opts ...pulumi.ResourceOption) (*NspLink, error) {
	var resource NspLink
	err := ctx.ReadResource("azure-native:network/v20210201preview:NspLink", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NspLink resources.
type nspLinkState struct {
}

type NspLinkState struct {
}

func (NspLinkState) ElementType() reflect.Type {
	return reflect.TypeOf((*nspLinkState)(nil)).Elem()
}

type nspLinkArgs struct {
	// Perimeter ARM Id for the remote NSP with which the link gets created in Auto-approval mode. It should be used when the NSP admin have Microsoft.Network/networkSecurityPerimeters/linkPerimeter/action permission on the remote NSP resource.
	AutoApprovedRemotePerimeterResourceId *string `pulumi:"autoApprovedRemotePerimeterResourceId"`
	// A message passed to the owner of the remote NSP link resource with this connection request. In case of Auto-approved flow, it is default to 'Auto Approved'. Restricted to 140 chars.
	Description *string `pulumi:"description"`
	// The name of the NSP link.
	LinkName *string `pulumi:"linkName"`
	// Local Inbound profile names to which Inbound is allowed. Use ['*'] to allow inbound to all profiles. It's default value is ['*'].
	LocalInboundProfiles []string `pulumi:"localInboundProfiles"`
	// The name of the network security perimeter.
	NetworkSecurityPerimeterName string `pulumi:"networkSecurityPerimeterName"`
	// Remote Inbound profile names to which Inbound is allowed. Use ['*'] to allow inbound to all profiles. This property can only be updated in auto-approval mode. It's default value is ['*'].
	RemoteInboundProfiles []string `pulumi:"remoteInboundProfiles"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a NspLink resource.
type NspLinkArgs struct {
	// Perimeter ARM Id for the remote NSP with which the link gets created in Auto-approval mode. It should be used when the NSP admin have Microsoft.Network/networkSecurityPerimeters/linkPerimeter/action permission on the remote NSP resource.
	AutoApprovedRemotePerimeterResourceId pulumi.StringPtrInput
	// A message passed to the owner of the remote NSP link resource with this connection request. In case of Auto-approved flow, it is default to 'Auto Approved'. Restricted to 140 chars.
	Description pulumi.StringPtrInput
	// The name of the NSP link.
	LinkName pulumi.StringPtrInput
	// Local Inbound profile names to which Inbound is allowed. Use ['*'] to allow inbound to all profiles. It's default value is ['*'].
	LocalInboundProfiles pulumi.StringArrayInput
	// The name of the network security perimeter.
	NetworkSecurityPerimeterName pulumi.StringInput
	// Remote Inbound profile names to which Inbound is allowed. Use ['*'] to allow inbound to all profiles. This property can only be updated in auto-approval mode. It's default value is ['*'].
	RemoteInboundProfiles pulumi.StringArrayInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
}

func (NspLinkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*nspLinkArgs)(nil)).Elem()
}

type NspLinkInput interface {
	pulumi.Input

	ToNspLinkOutput() NspLinkOutput
	ToNspLinkOutputWithContext(ctx context.Context) NspLinkOutput
}

func (*NspLink) ElementType() reflect.Type {
	return reflect.TypeOf((**NspLink)(nil)).Elem()
}

func (i *NspLink) ToNspLinkOutput() NspLinkOutput {
	return i.ToNspLinkOutputWithContext(context.Background())
}

func (i *NspLink) ToNspLinkOutputWithContext(ctx context.Context) NspLinkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NspLinkOutput)
}

func (i *NspLink) ToOutput(ctx context.Context) pulumix.Output[*NspLink] {
	return pulumix.Output[*NspLink]{
		OutputState: i.ToNspLinkOutputWithContext(ctx).OutputState,
	}
}

type NspLinkOutput struct{ *pulumi.OutputState }

func (NspLinkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NspLink)(nil)).Elem()
}

func (o NspLinkOutput) ToNspLinkOutput() NspLinkOutput {
	return o
}

func (o NspLinkOutput) ToNspLinkOutputWithContext(ctx context.Context) NspLinkOutput {
	return o
}

func (o NspLinkOutput) ToOutput(ctx context.Context) pulumix.Output[*NspLink] {
	return pulumix.Output[*NspLink]{
		OutputState: o.OutputState,
	}
}

// Perimeter ARM Id for the remote NSP with which the link gets created in Auto-approval mode. It should be used when the NSP admin have Microsoft.Network/networkSecurityPerimeters/linkPerimeter/action permission on the remote NSP resource.
func (o NspLinkOutput) AutoApprovedRemotePerimeterResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NspLink) pulumi.StringPtrOutput { return v.AutoApprovedRemotePerimeterResourceId }).(pulumi.StringPtrOutput)
}

// A message passed to the owner of the remote NSP link resource with this connection request. In case of Auto-approved flow, it is default to 'Auto Approved'. Restricted to 140 chars.
func (o NspLinkOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NspLink) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o NspLinkOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *NspLink) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Local Inbound profile names to which Inbound is allowed. Use ['*'] to allow inbound to all profiles. It's default value is ['*'].
func (o NspLinkOutput) LocalInboundProfiles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NspLink) pulumi.StringArrayOutput { return v.LocalInboundProfiles }).(pulumi.StringArrayOutput)
}

// Local Outbound profile names from which Outbound is allowed. In current version, it is readonly property and it's value is set to ['*'] to allow outbound from all profiles. In later version, user will be able to modify it.
func (o NspLinkOutput) LocalOutboundProfiles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NspLink) pulumi.StringArrayOutput { return v.LocalOutboundProfiles }).(pulumi.StringArrayOutput)
}

// Resource name.
func (o NspLinkOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NspLink) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the NSP Link resource.
func (o NspLinkOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *NspLink) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Remote Inbound profile names to which Inbound is allowed. Use ['*'] to allow inbound to all profiles. This property can only be updated in auto-approval mode. It's default value is ['*'].
func (o NspLinkOutput) RemoteInboundProfiles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NspLink) pulumi.StringArrayOutput { return v.RemoteInboundProfiles }).(pulumi.StringArrayOutput)
}

// Remote Outbound profile names from which Outbound is allowed. In current version, it is readonly property and it's value is set to ['*'] to allow outbound from all profiles. In later version, user will be able to modify it.
func (o NspLinkOutput) RemoteOutboundProfiles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NspLink) pulumi.StringArrayOutput { return v.RemoteOutboundProfiles }).(pulumi.StringArrayOutput)
}

// Remote NSP Guid with which the link gets created.
func (o NspLinkOutput) RemotePerimeterGuid() pulumi.StringOutput {
	return o.ApplyT(func(v *NspLink) pulumi.StringOutput { return v.RemotePerimeterGuid }).(pulumi.StringOutput)
}

// Remote NSP location with which the link gets created.
func (o NspLinkOutput) RemotePerimeterLocation() pulumi.StringOutput {
	return o.ApplyT(func(v *NspLink) pulumi.StringOutput { return v.RemotePerimeterLocation }).(pulumi.StringOutput)
}

// The NSP link state.
func (o NspLinkOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *NspLink) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Resource type.
func (o NspLinkOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *NspLink) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(NspLinkOutput{})
}
