// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package hybridcontainerservice

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The agentPool resource definition
// Azure REST API version: 2022-09-01-preview. Prior API version in Azure Native 1.x: 2022-05-01-preview.
type AgentPool struct {
	pulumi.CustomResourceState

	// AvailabilityZones - The list of Availability zones to use for nodes. Datacenter racks modelled as zones
	AvailabilityZones pulumi.StringArrayOutput `pulumi:"availabilityZones"`
	// The underlying cloud infra provider properties.
	CloudProviderProfile CloudProviderProfileResponsePtrOutput `pulumi:"cloudProviderProfile"`
	// Count - Number of agents to host docker containers. Allowed values must be in the range of 1 to 100 (inclusive). The default value is 1.
	Count            pulumi.IntPtrOutput                        `pulumi:"count"`
	ExtendedLocation AgentPoolResponseExtendedLocationPtrOutput `pulumi:"extendedLocation"`
	// The resource location
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The maximum number of nodes for auto-scaling
	MaxCount pulumi.IntPtrOutput `pulumi:"maxCount"`
	// The maximum number of pods that can run on a node.
	MaxPods pulumi.IntPtrOutput `pulumi:"maxPods"`
	// The minimum number of nodes for auto-scaling
	MinCount pulumi.IntPtrOutput `pulumi:"minCount"`
	// Mode - AgentPoolMode represents mode of an agent pool. Possible values include: 'System', 'LB', 'User'. Default is 'User'
	Mode pulumi.StringPtrOutput `pulumi:"mode"`
	// Resource Name
	Name pulumi.StringOutput `pulumi:"name"`
	// The version of node image
	NodeImageVersion pulumi.StringPtrOutput `pulumi:"nodeImageVersion"`
	// NodeLabels - Agent pool node labels to be persisted across all nodes in agent pool.
	NodeLabels pulumi.StringMapOutput `pulumi:"nodeLabels"`
	// NodeTaints - Taints added to new nodes during node pool create and scale. For example, key=value:NoSchedule.
	NodeTaints pulumi.StringArrayOutput `pulumi:"nodeTaints"`
	// OsType - OsType to be used to specify os type. Choose from Linux and Windows. Default to Linux. Possible values include: 'Linux', 'Windows'
	OsType            pulumi.StringPtrOutput `pulumi:"osType"`
	ProvisioningState pulumi.StringOutput    `pulumi:"provisioningState"`
	// HybridAKSNodePoolStatus defines the observed state of HybridAKSNodePool
	Status AgentPoolProvisioningStatusResponseStatusPtrOutput `pulumi:"status"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource Type
	Type pulumi.StringOutput `pulumi:"type"`
	// VmSize - The size of the agent pool VMs.
	VmSize pulumi.StringPtrOutput `pulumi:"vmSize"`
}

// NewAgentPool registers a new resource with the given unique name, arguments, and options.
func NewAgentPool(ctx *pulumi.Context,
	name string, args *AgentPoolArgs, opts ...pulumi.ResourceOption) (*AgentPool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ResourceName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceName'")
	}
	if args.Count == nil {
		args.Count = pulumi.IntPtr(1)
	}
	if args.Mode == nil {
		args.Mode = pulumi.StringPtr("User")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:hybridcontainerservice:agentPool"),
		},
		{
			Type: pulumi.String("azure-native:hybridcontainerservice/v20220501preview:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:hybridcontainerservice/v20220501preview:agentPool"),
		},
		{
			Type: pulumi.String("azure-native:hybridcontainerservice/v20220901preview:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:hybridcontainerservice/v20220901preview:agentPool"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource AgentPool
	err := ctx.RegisterResource("azure-native:hybridcontainerservice:AgentPool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAgentPool gets an existing AgentPool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAgentPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AgentPoolState, opts ...pulumi.ResourceOption) (*AgentPool, error) {
	var resource AgentPool
	err := ctx.ReadResource("azure-native:hybridcontainerservice:AgentPool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AgentPool resources.
type agentPoolState struct {
}

type AgentPoolState struct {
}

func (AgentPoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*agentPoolState)(nil)).Elem()
}

type agentPoolArgs struct {
	// Parameter for the name of the agent pool in the provisioned cluster
	AgentPoolName *string `pulumi:"agentPoolName"`
	// AvailabilityZones - The list of Availability zones to use for nodes. Datacenter racks modelled as zones
	AvailabilityZones []string `pulumi:"availabilityZones"`
	// The underlying cloud infra provider properties.
	CloudProviderProfile *CloudProviderProfile `pulumi:"cloudProviderProfile"`
	// Count - Number of agents to host docker containers. Allowed values must be in the range of 1 to 100 (inclusive). The default value is 1.
	Count            *int                       `pulumi:"count"`
	ExtendedLocation *AgentPoolExtendedLocation `pulumi:"extendedLocation"`
	// The resource location
	Location *string `pulumi:"location"`
	// The maximum number of nodes for auto-scaling
	MaxCount *int `pulumi:"maxCount"`
	// The maximum number of pods that can run on a node.
	MaxPods *int `pulumi:"maxPods"`
	// The minimum number of nodes for auto-scaling
	MinCount *int `pulumi:"minCount"`
	// Mode - AgentPoolMode represents mode of an agent pool. Possible values include: 'System', 'LB', 'User'. Default is 'User'
	Mode *string `pulumi:"mode"`
	// The version of node image
	NodeImageVersion *string `pulumi:"nodeImageVersion"`
	// NodeLabels - Agent pool node labels to be persisted across all nodes in agent pool.
	NodeLabels map[string]string `pulumi:"nodeLabels"`
	// NodeTaints - Taints added to new nodes during node pool create and scale. For example, key=value:NoSchedule.
	NodeTaints []string `pulumi:"nodeTaints"`
	// OsType - OsType to be used to specify os type. Choose from Linux and Windows. Default to Linux. Possible values include: 'Linux', 'Windows'
	OsType *string `pulumi:"osType"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Parameter for the name of the provisioned cluster
	ResourceName string `pulumi:"resourceName"`
	// HybridAKSNodePoolStatus defines the observed state of HybridAKSNodePool
	Status *AgentPoolProvisioningStatusStatus `pulumi:"status"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// VmSize - The size of the agent pool VMs.
	VmSize *string `pulumi:"vmSize"`
}

// The set of arguments for constructing a AgentPool resource.
type AgentPoolArgs struct {
	// Parameter for the name of the agent pool in the provisioned cluster
	AgentPoolName pulumi.StringPtrInput
	// AvailabilityZones - The list of Availability zones to use for nodes. Datacenter racks modelled as zones
	AvailabilityZones pulumi.StringArrayInput
	// The underlying cloud infra provider properties.
	CloudProviderProfile CloudProviderProfilePtrInput
	// Count - Number of agents to host docker containers. Allowed values must be in the range of 1 to 100 (inclusive). The default value is 1.
	Count            pulumi.IntPtrInput
	ExtendedLocation AgentPoolExtendedLocationPtrInput
	// The resource location
	Location pulumi.StringPtrInput
	// The maximum number of nodes for auto-scaling
	MaxCount pulumi.IntPtrInput
	// The maximum number of pods that can run on a node.
	MaxPods pulumi.IntPtrInput
	// The minimum number of nodes for auto-scaling
	MinCount pulumi.IntPtrInput
	// Mode - AgentPoolMode represents mode of an agent pool. Possible values include: 'System', 'LB', 'User'. Default is 'User'
	Mode pulumi.StringPtrInput
	// The version of node image
	NodeImageVersion pulumi.StringPtrInput
	// NodeLabels - Agent pool node labels to be persisted across all nodes in agent pool.
	NodeLabels pulumi.StringMapInput
	// NodeTaints - Taints added to new nodes during node pool create and scale. For example, key=value:NoSchedule.
	NodeTaints pulumi.StringArrayInput
	// OsType - OsType to be used to specify os type. Choose from Linux and Windows. Default to Linux. Possible values include: 'Linux', 'Windows'
	OsType pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Parameter for the name of the provisioned cluster
	ResourceName pulumi.StringInput
	// HybridAKSNodePoolStatus defines the observed state of HybridAKSNodePool
	Status AgentPoolProvisioningStatusStatusPtrInput
	// Resource tags
	Tags pulumi.StringMapInput
	// VmSize - The size of the agent pool VMs.
	VmSize pulumi.StringPtrInput
}

func (AgentPoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*agentPoolArgs)(nil)).Elem()
}

type AgentPoolInput interface {
	pulumi.Input

	ToAgentPoolOutput() AgentPoolOutput
	ToAgentPoolOutputWithContext(ctx context.Context) AgentPoolOutput
}

func (*AgentPool) ElementType() reflect.Type {
	return reflect.TypeOf((**AgentPool)(nil)).Elem()
}

func (i *AgentPool) ToAgentPoolOutput() AgentPoolOutput {
	return i.ToAgentPoolOutputWithContext(context.Background())
}

func (i *AgentPool) ToAgentPoolOutputWithContext(ctx context.Context) AgentPoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AgentPoolOutput)
}

func (i *AgentPool) ToOutput(ctx context.Context) pulumix.Output[*AgentPool] {
	return pulumix.Output[*AgentPool]{
		OutputState: i.ToAgentPoolOutputWithContext(ctx).OutputState,
	}
}

type AgentPoolOutput struct{ *pulumi.OutputState }

func (AgentPoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AgentPool)(nil)).Elem()
}

func (o AgentPoolOutput) ToAgentPoolOutput() AgentPoolOutput {
	return o
}

func (o AgentPoolOutput) ToAgentPoolOutputWithContext(ctx context.Context) AgentPoolOutput {
	return o
}

func (o AgentPoolOutput) ToOutput(ctx context.Context) pulumix.Output[*AgentPool] {
	return pulumix.Output[*AgentPool]{
		OutputState: o.OutputState,
	}
}

// AvailabilityZones - The list of Availability zones to use for nodes. Datacenter racks modelled as zones
func (o AgentPoolOutput) AvailabilityZones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringArrayOutput { return v.AvailabilityZones }).(pulumi.StringArrayOutput)
}

// The underlying cloud infra provider properties.
func (o AgentPoolOutput) CloudProviderProfile() CloudProviderProfileResponsePtrOutput {
	return o.ApplyT(func(v *AgentPool) CloudProviderProfileResponsePtrOutput { return v.CloudProviderProfile }).(CloudProviderProfileResponsePtrOutput)
}

// Count - Number of agents to host docker containers. Allowed values must be in the range of 1 to 100 (inclusive). The default value is 1.
func (o AgentPoolOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.IntPtrOutput { return v.Count }).(pulumi.IntPtrOutput)
}

func (o AgentPoolOutput) ExtendedLocation() AgentPoolResponseExtendedLocationPtrOutput {
	return o.ApplyT(func(v *AgentPool) AgentPoolResponseExtendedLocationPtrOutput { return v.ExtendedLocation }).(AgentPoolResponseExtendedLocationPtrOutput)
}

// The resource location
func (o AgentPoolOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The maximum number of nodes for auto-scaling
func (o AgentPoolOutput) MaxCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.IntPtrOutput { return v.MaxCount }).(pulumi.IntPtrOutput)
}

// The maximum number of pods that can run on a node.
func (o AgentPoolOutput) MaxPods() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.IntPtrOutput { return v.MaxPods }).(pulumi.IntPtrOutput)
}

// The minimum number of nodes for auto-scaling
func (o AgentPoolOutput) MinCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.IntPtrOutput { return v.MinCount }).(pulumi.IntPtrOutput)
}

// Mode - AgentPoolMode represents mode of an agent pool. Possible values include: 'System', 'LB', 'User'. Default is 'User'
func (o AgentPoolOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringPtrOutput { return v.Mode }).(pulumi.StringPtrOutput)
}

// Resource Name
func (o AgentPoolOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The version of node image
func (o AgentPoolOutput) NodeImageVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringPtrOutput { return v.NodeImageVersion }).(pulumi.StringPtrOutput)
}

// NodeLabels - Agent pool node labels to be persisted across all nodes in agent pool.
func (o AgentPoolOutput) NodeLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringMapOutput { return v.NodeLabels }).(pulumi.StringMapOutput)
}

// NodeTaints - Taints added to new nodes during node pool create and scale. For example, key=value:NoSchedule.
func (o AgentPoolOutput) NodeTaints() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringArrayOutput { return v.NodeTaints }).(pulumi.StringArrayOutput)
}

// OsType - OsType to be used to specify os type. Choose from Linux and Windows. Default to Linux. Possible values include: 'Linux', 'Windows'
func (o AgentPoolOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringPtrOutput { return v.OsType }).(pulumi.StringPtrOutput)
}

func (o AgentPoolOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// HybridAKSNodePoolStatus defines the observed state of HybridAKSNodePool
func (o AgentPoolOutput) Status() AgentPoolProvisioningStatusResponseStatusPtrOutput {
	return o.ApplyT(func(v *AgentPool) AgentPoolProvisioningStatusResponseStatusPtrOutput { return v.Status }).(AgentPoolProvisioningStatusResponseStatusPtrOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o AgentPoolOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *AgentPool) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags
func (o AgentPoolOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource Type
func (o AgentPoolOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// VmSize - The size of the agent pool VMs.
func (o AgentPoolOutput) VmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringPtrOutput { return v.VmSize }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AgentPoolOutput{})
}
