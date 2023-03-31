// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20190401

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Agent Pool.
//
// Deprecated: Version 2019-04-01 will be removed in v2 of the provider.
type AgentPool struct {
	pulumi.CustomResourceState

	// (PREVIEW) Availability zones for nodes. Must use VirtualMachineScaleSets AgentPoolType.
	AvailabilityZones pulumi.StringArrayOutput `pulumi:"availabilityZones"`
	// Number of agents (VMs) to host docker containers. Allowed values must be in the range of 1 to 100 (inclusive). The default value is 1.
	Count pulumi.IntOutput `pulumi:"count"`
	// Whether to enable auto-scaler
	EnableAutoScaling pulumi.BoolPtrOutput `pulumi:"enableAutoScaling"`
	// Maximum number of nodes for auto-scaling
	MaxCount pulumi.IntPtrOutput `pulumi:"maxCount"`
	// Maximum number of pods that can run on a node.
	MaxPods pulumi.IntPtrOutput `pulumi:"maxPods"`
	// Minimum number of nodes for auto-scaling
	MinCount pulumi.IntPtrOutput `pulumi:"minCount"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Version of orchestrator specified when creating the managed cluster.
	OrchestratorVersion pulumi.StringPtrOutput `pulumi:"orchestratorVersion"`
	// OS Disk Size in GB to be used to specify the disk size for every machine in this master/agent pool. If you specify 0, it will apply the default osDisk size according to the vmSize specified.
	OsDiskSizeGB pulumi.IntPtrOutput `pulumi:"osDiskSizeGB"`
	// OsType to be used to specify os type. Choose from Linux and Windows. Default to Linux.
	OsType pulumi.StringPtrOutput `pulumi:"osType"`
	// The current deployment or provisioning state, which only appears in the response.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// AgentPoolType represents types of an agent pool
	Type pulumi.StringOutput `pulumi:"type"`
	// Size of agent VMs.
	VmSize pulumi.StringOutput `pulumi:"vmSize"`
	// VNet SubnetID specifies the VNet's subnet identifier.
	VnetSubnetID pulumi.StringPtrOutput `pulumi:"vnetSubnetID"`
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
	if args.VmSize == nil {
		return nil, errors.New("invalid value for required argument 'VmSize'")
	}
	if isZero(args.Count) {
		args.Count = pulumi.Int(1)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:containerservice:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20190201:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20190601:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20190801:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20191001:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20191101:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20200101:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20200201:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20200301:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20200401:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20200601:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20200701:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20200901:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20201101:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20201201:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20210201:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20210301:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20210501:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20210701:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20210801:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20210901:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20211001:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20211101preview:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220101:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220102preview:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220201:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220202preview:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220301:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220302preview:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220401:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220402preview:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220502preview:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220601:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220602preview:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220701:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220702preview:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220802preview:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220803preview:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220901:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20220902preview:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20221002preview:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20221101:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20221102preview:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20230101:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20230102preview:AgentPool"),
		},
	})
	opts = append(opts, aliases)
	var resource AgentPool
	err := ctx.RegisterResource("azure-native:containerservice/v20190401:AgentPool", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure-native:containerservice/v20190401:AgentPool", name, id, state, &resource, opts...)
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
	// The name of the agent pool.
	AgentPoolName *string `pulumi:"agentPoolName"`
	// (PREVIEW) Availability zones for nodes. Must use VirtualMachineScaleSets AgentPoolType.
	AvailabilityZones []string `pulumi:"availabilityZones"`
	// Number of agents (VMs) to host docker containers. Allowed values must be in the range of 1 to 100 (inclusive). The default value is 1.
	Count int `pulumi:"count"`
	// Whether to enable auto-scaler
	EnableAutoScaling *bool `pulumi:"enableAutoScaling"`
	// Maximum number of nodes for auto-scaling
	MaxCount *int `pulumi:"maxCount"`
	// Maximum number of pods that can run on a node.
	MaxPods *int `pulumi:"maxPods"`
	// Minimum number of nodes for auto-scaling
	MinCount *int `pulumi:"minCount"`
	// Version of orchestrator specified when creating the managed cluster.
	OrchestratorVersion *string `pulumi:"orchestratorVersion"`
	// OS Disk Size in GB to be used to specify the disk size for every machine in this master/agent pool. If you specify 0, it will apply the default osDisk size according to the vmSize specified.
	OsDiskSizeGB *int `pulumi:"osDiskSizeGB"`
	// OsType to be used to specify os type. Choose from Linux and Windows. Default to Linux.
	OsType *string `pulumi:"osType"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the managed cluster resource.
	ResourceName string `pulumi:"resourceName"`
	// AgentPoolType represents types of an agent pool
	Type *string `pulumi:"type"`
	// Size of agent VMs.
	VmSize string `pulumi:"vmSize"`
	// VNet SubnetID specifies the VNet's subnet identifier.
	VnetSubnetID *string `pulumi:"vnetSubnetID"`
}

// The set of arguments for constructing a AgentPool resource.
type AgentPoolArgs struct {
	// The name of the agent pool.
	AgentPoolName pulumi.StringPtrInput
	// (PREVIEW) Availability zones for nodes. Must use VirtualMachineScaleSets AgentPoolType.
	AvailabilityZones pulumi.StringArrayInput
	// Number of agents (VMs) to host docker containers. Allowed values must be in the range of 1 to 100 (inclusive). The default value is 1.
	Count pulumi.IntInput
	// Whether to enable auto-scaler
	EnableAutoScaling pulumi.BoolPtrInput
	// Maximum number of nodes for auto-scaling
	MaxCount pulumi.IntPtrInput
	// Maximum number of pods that can run on a node.
	MaxPods pulumi.IntPtrInput
	// Minimum number of nodes for auto-scaling
	MinCount pulumi.IntPtrInput
	// Version of orchestrator specified when creating the managed cluster.
	OrchestratorVersion pulumi.StringPtrInput
	// OS Disk Size in GB to be used to specify the disk size for every machine in this master/agent pool. If you specify 0, it will apply the default osDisk size according to the vmSize specified.
	OsDiskSizeGB pulumi.IntPtrInput
	// OsType to be used to specify os type. Choose from Linux and Windows. Default to Linux.
	OsType pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The name of the managed cluster resource.
	ResourceName pulumi.StringInput
	// AgentPoolType represents types of an agent pool
	Type pulumi.StringPtrInput
	// Size of agent VMs.
	VmSize pulumi.StringInput
	// VNet SubnetID specifies the VNet's subnet identifier.
	VnetSubnetID pulumi.StringPtrInput
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

// (PREVIEW) Availability zones for nodes. Must use VirtualMachineScaleSets AgentPoolType.
func (o AgentPoolOutput) AvailabilityZones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringArrayOutput { return v.AvailabilityZones }).(pulumi.StringArrayOutput)
}

// Number of agents (VMs) to host docker containers. Allowed values must be in the range of 1 to 100 (inclusive). The default value is 1.
func (o AgentPoolOutput) Count() pulumi.IntOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.IntOutput { return v.Count }).(pulumi.IntOutput)
}

// Whether to enable auto-scaler
func (o AgentPoolOutput) EnableAutoScaling() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.BoolPtrOutput { return v.EnableAutoScaling }).(pulumi.BoolPtrOutput)
}

// Maximum number of nodes for auto-scaling
func (o AgentPoolOutput) MaxCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.IntPtrOutput { return v.MaxCount }).(pulumi.IntPtrOutput)
}

// Maximum number of pods that can run on a node.
func (o AgentPoolOutput) MaxPods() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.IntPtrOutput { return v.MaxPods }).(pulumi.IntPtrOutput)
}

// Minimum number of nodes for auto-scaling
func (o AgentPoolOutput) MinCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.IntPtrOutput { return v.MinCount }).(pulumi.IntPtrOutput)
}

// The name of the resource that is unique within a resource group. This name can be used to access the resource.
func (o AgentPoolOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Version of orchestrator specified when creating the managed cluster.
func (o AgentPoolOutput) OrchestratorVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringPtrOutput { return v.OrchestratorVersion }).(pulumi.StringPtrOutput)
}

// OS Disk Size in GB to be used to specify the disk size for every machine in this master/agent pool. If you specify 0, it will apply the default osDisk size according to the vmSize specified.
func (o AgentPoolOutput) OsDiskSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.IntPtrOutput { return v.OsDiskSizeGB }).(pulumi.IntPtrOutput)
}

// OsType to be used to specify os type. Choose from Linux and Windows. Default to Linux.
func (o AgentPoolOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringPtrOutput { return v.OsType }).(pulumi.StringPtrOutput)
}

// The current deployment or provisioning state, which only appears in the response.
func (o AgentPoolOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// AgentPoolType represents types of an agent pool
func (o AgentPoolOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Size of agent VMs.
func (o AgentPoolOutput) VmSize() pulumi.StringOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringOutput { return v.VmSize }).(pulumi.StringOutput)
}

// VNet SubnetID specifies the VNet's subnet identifier.
func (o AgentPoolOutput) VnetSubnetID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringPtrOutput { return v.VnetSubnetID }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AgentPoolOutput{})
}
