// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230615

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The L2 Isolation Domain resource definition.
type L2IsolationDomain struct {
	pulumi.CustomResourceState

	// Administrative state of the resource.
	AdministrativeState pulumi.StringOutput `pulumi:"administrativeState"`
	// Switch configuration description.
	Annotation pulumi.StringPtrOutput `pulumi:"annotation"`
	// Configuration state of the resource.
	ConfigurationState pulumi.StringOutput `pulumi:"configurationState"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// Maximum transmission unit. Default value is 1500.
	Mtu pulumi.IntPtrOutput `pulumi:"mtu"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// ARM Resource ID of the Network Fabric.
	NetworkFabricId pulumi.StringOutput `pulumi:"networkFabricId"`
	// Provisioning state of the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// Vlan Identifier of the Network Fabric. Example: 501.
	VlanId pulumi.IntOutput `pulumi:"vlanId"`
}

// NewL2IsolationDomain registers a new resource with the given unique name, arguments, and options.
func NewL2IsolationDomain(ctx *pulumi.Context,
	name string, args *L2IsolationDomainArgs, opts ...pulumi.ResourceOption) (*L2IsolationDomain, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NetworkFabricId == nil {
		return nil, errors.New("invalid value for required argument 'NetworkFabricId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.VlanId == nil {
		return nil, errors.New("invalid value for required argument 'VlanId'")
	}
	if args.Mtu == nil {
		args.Mtu = pulumi.IntPtr(1500)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:managednetworkfabric:L2IsolationDomain"),
		},
		{
			Type: pulumi.String("azure-native:managednetworkfabric/v20230201preview:L2IsolationDomain"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource L2IsolationDomain
	err := ctx.RegisterResource("azure-native:managednetworkfabric/v20230615:L2IsolationDomain", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetL2IsolationDomain gets an existing L2IsolationDomain resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetL2IsolationDomain(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *L2IsolationDomainState, opts ...pulumi.ResourceOption) (*L2IsolationDomain, error) {
	var resource L2IsolationDomain
	err := ctx.ReadResource("azure-native:managednetworkfabric/v20230615:L2IsolationDomain", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering L2IsolationDomain resources.
type l2isolationDomainState struct {
}

type L2IsolationDomainState struct {
}

func (L2IsolationDomainState) ElementType() reflect.Type {
	return reflect.TypeOf((*l2isolationDomainState)(nil)).Elem()
}

type l2isolationDomainArgs struct {
	// Switch configuration description.
	Annotation *string `pulumi:"annotation"`
	// Name of the L2 Isolation Domain.
	L2IsolationDomainName *string `pulumi:"l2IsolationDomainName"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// Maximum transmission unit. Default value is 1500.
	Mtu *int `pulumi:"mtu"`
	// ARM Resource ID of the Network Fabric.
	NetworkFabricId string `pulumi:"networkFabricId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Vlan Identifier of the Network Fabric. Example: 501.
	VlanId int `pulumi:"vlanId"`
}

// The set of arguments for constructing a L2IsolationDomain resource.
type L2IsolationDomainArgs struct {
	// Switch configuration description.
	Annotation pulumi.StringPtrInput
	// Name of the L2 Isolation Domain.
	L2IsolationDomainName pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// Maximum transmission unit. Default value is 1500.
	Mtu pulumi.IntPtrInput
	// ARM Resource ID of the Network Fabric.
	NetworkFabricId pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Vlan Identifier of the Network Fabric. Example: 501.
	VlanId pulumi.IntInput
}

func (L2IsolationDomainArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*l2isolationDomainArgs)(nil)).Elem()
}

type L2IsolationDomainInput interface {
	pulumi.Input

	ToL2IsolationDomainOutput() L2IsolationDomainOutput
	ToL2IsolationDomainOutputWithContext(ctx context.Context) L2IsolationDomainOutput
}

func (*L2IsolationDomain) ElementType() reflect.Type {
	return reflect.TypeOf((**L2IsolationDomain)(nil)).Elem()
}

func (i *L2IsolationDomain) ToL2IsolationDomainOutput() L2IsolationDomainOutput {
	return i.ToL2IsolationDomainOutputWithContext(context.Background())
}

func (i *L2IsolationDomain) ToL2IsolationDomainOutputWithContext(ctx context.Context) L2IsolationDomainOutput {
	return pulumi.ToOutputWithContext(ctx, i).(L2IsolationDomainOutput)
}

func (i *L2IsolationDomain) ToOutput(ctx context.Context) pulumix.Output[*L2IsolationDomain] {
	return pulumix.Output[*L2IsolationDomain]{
		OutputState: i.ToL2IsolationDomainOutputWithContext(ctx).OutputState,
	}
}

type L2IsolationDomainOutput struct{ *pulumi.OutputState }

func (L2IsolationDomainOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**L2IsolationDomain)(nil)).Elem()
}

func (o L2IsolationDomainOutput) ToL2IsolationDomainOutput() L2IsolationDomainOutput {
	return o
}

func (o L2IsolationDomainOutput) ToL2IsolationDomainOutputWithContext(ctx context.Context) L2IsolationDomainOutput {
	return o
}

func (o L2IsolationDomainOutput) ToOutput(ctx context.Context) pulumix.Output[*L2IsolationDomain] {
	return pulumix.Output[*L2IsolationDomain]{
		OutputState: o.OutputState,
	}
}

// Administrative state of the resource.
func (o L2IsolationDomainOutput) AdministrativeState() pulumi.StringOutput {
	return o.ApplyT(func(v *L2IsolationDomain) pulumi.StringOutput { return v.AdministrativeState }).(pulumi.StringOutput)
}

// Switch configuration description.
func (o L2IsolationDomainOutput) Annotation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *L2IsolationDomain) pulumi.StringPtrOutput { return v.Annotation }).(pulumi.StringPtrOutput)
}

// Configuration state of the resource.
func (o L2IsolationDomainOutput) ConfigurationState() pulumi.StringOutput {
	return o.ApplyT(func(v *L2IsolationDomain) pulumi.StringOutput { return v.ConfigurationState }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o L2IsolationDomainOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *L2IsolationDomain) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Maximum transmission unit. Default value is 1500.
func (o L2IsolationDomainOutput) Mtu() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *L2IsolationDomain) pulumi.IntPtrOutput { return v.Mtu }).(pulumi.IntPtrOutput)
}

// The name of the resource
func (o L2IsolationDomainOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *L2IsolationDomain) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// ARM Resource ID of the Network Fabric.
func (o L2IsolationDomainOutput) NetworkFabricId() pulumi.StringOutput {
	return o.ApplyT(func(v *L2IsolationDomain) pulumi.StringOutput { return v.NetworkFabricId }).(pulumi.StringOutput)
}

// Provisioning state of the resource.
func (o L2IsolationDomainOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *L2IsolationDomain) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o L2IsolationDomainOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *L2IsolationDomain) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o L2IsolationDomainOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *L2IsolationDomain) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o L2IsolationDomainOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *L2IsolationDomain) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Vlan Identifier of the Network Fabric. Example: 501.
func (o L2IsolationDomainOutput) VlanId() pulumi.IntOutput {
	return o.ApplyT(func(v *L2IsolationDomain) pulumi.IntOutput { return v.VlanId }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(L2IsolationDomainOutput{})
}
