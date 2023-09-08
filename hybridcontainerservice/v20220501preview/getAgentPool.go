// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220501preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the agent pool in the Hybrid AKS provisioned cluster
func LookupAgentPool(ctx *pulumi.Context, args *LookupAgentPoolArgs, opts ...pulumi.InvokeOption) (*LookupAgentPoolResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupAgentPoolResult
	err := ctx.Invoke("azure-native:hybridcontainerservice/v20220501preview:getAgentPool", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupAgentPoolArgs struct {
	// Parameter for the name of the agent pool in the provisioned cluster
	AgentPoolName string `pulumi:"agentPoolName"`
	// Parameter for the name of the provisioned cluster
	ProvisionedClustersName string `pulumi:"provisionedClustersName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The agentPool resource definition
type LookupAgentPoolResult struct {
	// AvailabilityZones - The list of Availability zones to use for nodes. Datacenter racks modelled as zones
	AvailabilityZones []string `pulumi:"availabilityZones"`
	// The underlying cloud infra provider properties.
	CloudProviderProfile *CloudProviderProfileResponse `pulumi:"cloudProviderProfile"`
	// Count - Number of agents to host docker containers. Allowed values must be in the range of 1 to 100 (inclusive). The default value is 1.
	Count            *int                               `pulumi:"count"`
	ExtendedLocation *AgentPoolResponseExtendedLocation `pulumi:"extendedLocation"`
	// Resource Id
	Id string `pulumi:"id"`
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
	// Resource Name
	Name string `pulumi:"name"`
	// The version of node image
	NodeImageVersion *string `pulumi:"nodeImageVersion"`
	// NodeLabels - Agent pool node labels to be persisted across all nodes in agent pool.
	NodeLabels map[string]string `pulumi:"nodeLabels"`
	// NodeTaints - Taints added to new nodes during node pool create and scale. For example, key=value:NoSchedule.
	NodeTaints []string `pulumi:"nodeTaints"`
	// OsType - OsType to be used to specify os type. Choose from Linux and Windows. Default to Linux. Possible values include: 'Linux', 'Windows'
	OsType            *string `pulumi:"osType"`
	ProvisioningState string  `pulumi:"provisioningState"`
	// HybridAKSNodePoolStatus defines the observed state of HybridAKSNodePool
	Status *AgentPoolProvisioningStatusResponseStatus `pulumi:"status"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource Type
	Type string `pulumi:"type"`
	// VmSize - The size of the agent pool VMs.
	VmSize *string `pulumi:"vmSize"`
}

// Defaults sets the appropriate defaults for LookupAgentPoolResult
func (val *LookupAgentPoolResult) Defaults() *LookupAgentPoolResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.Count == nil {
		count_ := 1
		tmp.Count = &count_
	}
	if tmp.Mode == nil {
		mode_ := "User"
		tmp.Mode = &mode_
	}
	if tmp.OsType == nil {
		osType_ := "Linux"
		tmp.OsType = &osType_
	}
	return &tmp
}

func LookupAgentPoolOutput(ctx *pulumi.Context, args LookupAgentPoolOutputArgs, opts ...pulumi.InvokeOption) LookupAgentPoolResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAgentPoolResult, error) {
			args := v.(LookupAgentPoolArgs)
			r, err := LookupAgentPool(ctx, &args, opts...)
			var s LookupAgentPoolResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAgentPoolResultOutput)
}

type LookupAgentPoolOutputArgs struct {
	// Parameter for the name of the agent pool in the provisioned cluster
	AgentPoolName pulumi.StringInput `pulumi:"agentPoolName"`
	// Parameter for the name of the provisioned cluster
	ProvisionedClustersName pulumi.StringInput `pulumi:"provisionedClustersName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupAgentPoolOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAgentPoolArgs)(nil)).Elem()
}

// The agentPool resource definition
type LookupAgentPoolResultOutput struct{ *pulumi.OutputState }

func (LookupAgentPoolResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAgentPoolResult)(nil)).Elem()
}

func (o LookupAgentPoolResultOutput) ToLookupAgentPoolResultOutput() LookupAgentPoolResultOutput {
	return o
}

func (o LookupAgentPoolResultOutput) ToLookupAgentPoolResultOutputWithContext(ctx context.Context) LookupAgentPoolResultOutput {
	return o
}

// AvailabilityZones - The list of Availability zones to use for nodes. Datacenter racks modelled as zones
func (o LookupAgentPoolResultOutput) AvailabilityZones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) []string { return v.AvailabilityZones }).(pulumi.StringArrayOutput)
}

// The underlying cloud infra provider properties.
func (o LookupAgentPoolResultOutput) CloudProviderProfile() CloudProviderProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) *CloudProviderProfileResponse { return v.CloudProviderProfile }).(CloudProviderProfileResponsePtrOutput)
}

// Count - Number of agents to host docker containers. Allowed values must be in the range of 1 to 100 (inclusive). The default value is 1.
func (o LookupAgentPoolResultOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) *int { return v.Count }).(pulumi.IntPtrOutput)
}

func (o LookupAgentPoolResultOutput) ExtendedLocation() AgentPoolResponseExtendedLocationPtrOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) *AgentPoolResponseExtendedLocation { return v.ExtendedLocation }).(AgentPoolResponseExtendedLocationPtrOutput)
}

// Resource Id
func (o LookupAgentPoolResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) string { return v.Id }).(pulumi.StringOutput)
}

// The resource location
func (o LookupAgentPoolResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// The maximum number of nodes for auto-scaling
func (o LookupAgentPoolResultOutput) MaxCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) *int { return v.MaxCount }).(pulumi.IntPtrOutput)
}

// The maximum number of pods that can run on a node.
func (o LookupAgentPoolResultOutput) MaxPods() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) *int { return v.MaxPods }).(pulumi.IntPtrOutput)
}

// The minimum number of nodes for auto-scaling
func (o LookupAgentPoolResultOutput) MinCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) *int { return v.MinCount }).(pulumi.IntPtrOutput)
}

// Mode - AgentPoolMode represents mode of an agent pool. Possible values include: 'System', 'LB', 'User'. Default is 'User'
func (o LookupAgentPoolResultOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) *string { return v.Mode }).(pulumi.StringPtrOutput)
}

// Resource Name
func (o LookupAgentPoolResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) string { return v.Name }).(pulumi.StringOutput)
}

// The version of node image
func (o LookupAgentPoolResultOutput) NodeImageVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) *string { return v.NodeImageVersion }).(pulumi.StringPtrOutput)
}

// NodeLabels - Agent pool node labels to be persisted across all nodes in agent pool.
func (o LookupAgentPoolResultOutput) NodeLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) map[string]string { return v.NodeLabels }).(pulumi.StringMapOutput)
}

// NodeTaints - Taints added to new nodes during node pool create and scale. For example, key=value:NoSchedule.
func (o LookupAgentPoolResultOutput) NodeTaints() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) []string { return v.NodeTaints }).(pulumi.StringArrayOutput)
}

// OsType - OsType to be used to specify os type. Choose from Linux and Windows. Default to Linux. Possible values include: 'Linux', 'Windows'
func (o LookupAgentPoolResultOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) *string { return v.OsType }).(pulumi.StringPtrOutput)
}

func (o LookupAgentPoolResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// HybridAKSNodePoolStatus defines the observed state of HybridAKSNodePool
func (o LookupAgentPoolResultOutput) Status() AgentPoolProvisioningStatusResponseStatusPtrOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) *AgentPoolProvisioningStatusResponseStatus { return v.Status }).(AgentPoolProvisioningStatusResponseStatusPtrOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupAgentPoolResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags
func (o LookupAgentPoolResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource Type
func (o LookupAgentPoolResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) string { return v.Type }).(pulumi.StringOutput)
}

// VmSize - The size of the agent pool VMs.
func (o LookupAgentPoolResultOutput) VmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) *string { return v.VmSize }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAgentPoolResultOutput{})
}