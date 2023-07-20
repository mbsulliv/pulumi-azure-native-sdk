// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230201preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The IpCommunity resource definition.
type IpCommunity struct {
	pulumi.CustomResourceState

	// Action to be taken on the configuration. Example: Permit | Deny.
	Action pulumi.StringOutput `pulumi:"action"`
	// Switch configuration description.
	Annotation pulumi.StringPtrOutput `pulumi:"annotation"`
	// List the communityMembers of IP Community .
	CommunityMembers pulumi.StringArrayOutput `pulumi:"communityMembers"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Gets the provisioning state of the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// Supported well known Community List.
	WellKnownCommunities pulumi.StringArrayOutput `pulumi:"wellKnownCommunities"`
}

// NewIpCommunity registers a new resource with the given unique name, arguments, and options.
func NewIpCommunity(ctx *pulumi.Context,
	name string, args *IpCommunityArgs, opts ...pulumi.ResourceOption) (*IpCommunity, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Action == nil {
		return nil, errors.New("invalid value for required argument 'Action'")
	}
	if args.CommunityMembers == nil {
		return nil, errors.New("invalid value for required argument 'CommunityMembers'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:managednetworkfabric:IpCommunity"),
		},
		{
			Type: pulumi.String("azure-native:managednetworkfabric/v20230615:IpCommunity"),
		},
	})
	opts = append(opts, aliases)
	var resource IpCommunity
	err := ctx.RegisterResource("azure-native:managednetworkfabric/v20230201preview:IpCommunity", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIpCommunity gets an existing IpCommunity resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIpCommunity(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IpCommunityState, opts ...pulumi.ResourceOption) (*IpCommunity, error) {
	var resource IpCommunity
	err := ctx.ReadResource("azure-native:managednetworkfabric/v20230201preview:IpCommunity", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IpCommunity resources.
type ipCommunityState struct {
}

type IpCommunityState struct {
}

func (IpCommunityState) ElementType() reflect.Type {
	return reflect.TypeOf((*ipCommunityState)(nil)).Elem()
}

type ipCommunityArgs struct {
	// Action to be taken on the configuration. Example: Permit | Deny.
	Action string `pulumi:"action"`
	// Switch configuration description.
	Annotation *string `pulumi:"annotation"`
	// List the communityMembers of IP Community .
	CommunityMembers []string `pulumi:"communityMembers"`
	// Name of the IP Community
	IpCommunityName *string `pulumi:"ipCommunityName"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Supported well known Community List.
	WellKnownCommunities []string `pulumi:"wellKnownCommunities"`
}

// The set of arguments for constructing a IpCommunity resource.
type IpCommunityArgs struct {
	// Action to be taken on the configuration. Example: Permit | Deny.
	Action pulumi.StringInput
	// Switch configuration description.
	Annotation pulumi.StringPtrInput
	// List the communityMembers of IP Community .
	CommunityMembers pulumi.StringArrayInput
	// Name of the IP Community
	IpCommunityName pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Supported well known Community List.
	WellKnownCommunities pulumi.StringArrayInput
}

func (IpCommunityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ipCommunityArgs)(nil)).Elem()
}

type IpCommunityInput interface {
	pulumi.Input

	ToIpCommunityOutput() IpCommunityOutput
	ToIpCommunityOutputWithContext(ctx context.Context) IpCommunityOutput
}

func (*IpCommunity) ElementType() reflect.Type {
	return reflect.TypeOf((**IpCommunity)(nil)).Elem()
}

func (i *IpCommunity) ToIpCommunityOutput() IpCommunityOutput {
	return i.ToIpCommunityOutputWithContext(context.Background())
}

func (i *IpCommunity) ToIpCommunityOutputWithContext(ctx context.Context) IpCommunityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpCommunityOutput)
}

type IpCommunityOutput struct{ *pulumi.OutputState }

func (IpCommunityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IpCommunity)(nil)).Elem()
}

func (o IpCommunityOutput) ToIpCommunityOutput() IpCommunityOutput {
	return o
}

func (o IpCommunityOutput) ToIpCommunityOutputWithContext(ctx context.Context) IpCommunityOutput {
	return o
}

// Action to be taken on the configuration. Example: Permit | Deny.
func (o IpCommunityOutput) Action() pulumi.StringOutput {
	return o.ApplyT(func(v *IpCommunity) pulumi.StringOutput { return v.Action }).(pulumi.StringOutput)
}

// Switch configuration description.
func (o IpCommunityOutput) Annotation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IpCommunity) pulumi.StringPtrOutput { return v.Annotation }).(pulumi.StringPtrOutput)
}

// List the communityMembers of IP Community .
func (o IpCommunityOutput) CommunityMembers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *IpCommunity) pulumi.StringArrayOutput { return v.CommunityMembers }).(pulumi.StringArrayOutput)
}

// The geo-location where the resource lives
func (o IpCommunityOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *IpCommunity) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o IpCommunityOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *IpCommunity) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Gets the provisioning state of the resource.
func (o IpCommunityOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *IpCommunity) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o IpCommunityOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *IpCommunity) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o IpCommunityOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *IpCommunity) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o IpCommunityOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *IpCommunity) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Supported well known Community List.
func (o IpCommunityOutput) WellKnownCommunities() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *IpCommunity) pulumi.StringArrayOutput { return v.WellKnownCommunities }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(IpCommunityOutput{})
}
