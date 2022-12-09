// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package workloads

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Define the SAP Central Services Instance resource.
// API Version: 2021-12-01-preview.
type SAPCentralInstance struct {
	pulumi.CustomResourceState

	// Defines the SAP Enqueue Replication Server (ERS) properties.
	EnqueueReplicationServerProperties EnqueueReplicationServerPropertiesResponsePtrOutput `pulumi:"enqueueReplicationServerProperties"`
	// Defines the SAP Enqueue Server properties.
	EnqueueServerProperties EnqueueServerPropertiesResponsePtrOutput `pulumi:"enqueueServerProperties"`
	// Defines the errors related to SAP Central Services Instance resource.
	Errors SAPVirtualInstanceErrorResponseOutput `pulumi:"errors"`
	// Defines the SAP Gateway Server properties.
	GatewayServerProperties GatewayServerPropertiesResponsePtrOutput `pulumi:"gatewayServerProperties"`
	// Defines the health of SAP Instances.
	Health pulumi.StringOutput `pulumi:"health"`
	// The central services instance number.
	InstanceNo pulumi.StringOutput `pulumi:"instanceNo"`
	// The central services instance Kernel Patch level.
	KernelPatch pulumi.StringOutput `pulumi:"kernelPatch"`
	// The central services instance Kernel Version.
	KernelVersion pulumi.StringOutput `pulumi:"kernelVersion"`
	// The Load Balancer details such as LoadBalancer ID attached to ASCS Virtual Machines
	LoadBalancerDetails LoadBalancerDetailsResponseOutput `pulumi:"loadBalancerDetails"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// Defines the SAP Message Server properties.
	MessageServerProperties MessageServerPropertiesResponsePtrOutput `pulumi:"messageServerProperties"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Defines the provisioning states.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Defines the SAP Instance status.
	Status pulumi.StringOutput `pulumi:"status"`
	// The central services instance subnet.
	Subnet pulumi.StringOutput `pulumi:"subnet"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// The list of virtual machines corresponding to the Central Services instance.
	VmDetails CentralServerVmDetailsResponseArrayOutput `pulumi:"vmDetails"`
}

// NewSAPCentralInstance registers a new resource with the given unique name, arguments, and options.
func NewSAPCentralInstance(ctx *pulumi.Context,
	name string, args *SAPCentralInstanceArgs, opts ...pulumi.ResourceOption) (*SAPCentralInstance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SapVirtualInstanceName == nil {
		return nil, errors.New("invalid value for required argument 'SapVirtualInstanceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:workloads/v20211201preview:SAPCentralInstance"),
		},
	})
	opts = append(opts, aliases)
	var resource SAPCentralInstance
	err := ctx.RegisterResource("azure-native:workloads:SAPCentralInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSAPCentralInstance gets an existing SAPCentralInstance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSAPCentralInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SAPCentralInstanceState, opts ...pulumi.ResourceOption) (*SAPCentralInstance, error) {
	var resource SAPCentralInstance
	err := ctx.ReadResource("azure-native:workloads:SAPCentralInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SAPCentralInstance resources.
type sapcentralInstanceState struct {
}

type SAPCentralInstanceState struct {
}

func (SAPCentralInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*sapcentralInstanceState)(nil)).Elem()
}

type sapcentralInstanceArgs struct {
	// Central Services Instance resource name string modeled as parameter for auto generation to work correctly.
	CentralInstanceName *string `pulumi:"centralInstanceName"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Virtual Instances for SAP solutions resource
	SapVirtualInstanceName string `pulumi:"sapVirtualInstanceName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a SAPCentralInstance resource.
type SAPCentralInstanceArgs struct {
	// Central Services Instance resource name string modeled as parameter for auto generation to work correctly.
	CentralInstanceName pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the Virtual Instances for SAP solutions resource
	SapVirtualInstanceName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (SAPCentralInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sapcentralInstanceArgs)(nil)).Elem()
}

type SAPCentralInstanceInput interface {
	pulumi.Input

	ToSAPCentralInstanceOutput() SAPCentralInstanceOutput
	ToSAPCentralInstanceOutputWithContext(ctx context.Context) SAPCentralInstanceOutput
}

func (*SAPCentralInstance) ElementType() reflect.Type {
	return reflect.TypeOf((**SAPCentralInstance)(nil)).Elem()
}

func (i *SAPCentralInstance) ToSAPCentralInstanceOutput() SAPCentralInstanceOutput {
	return i.ToSAPCentralInstanceOutputWithContext(context.Background())
}

func (i *SAPCentralInstance) ToSAPCentralInstanceOutputWithContext(ctx context.Context) SAPCentralInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SAPCentralInstanceOutput)
}

type SAPCentralInstanceOutput struct{ *pulumi.OutputState }

func (SAPCentralInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SAPCentralInstance)(nil)).Elem()
}

func (o SAPCentralInstanceOutput) ToSAPCentralInstanceOutput() SAPCentralInstanceOutput {
	return o
}

func (o SAPCentralInstanceOutput) ToSAPCentralInstanceOutputWithContext(ctx context.Context) SAPCentralInstanceOutput {
	return o
}

// Defines the SAP Enqueue Replication Server (ERS) properties.
func (o SAPCentralInstanceOutput) EnqueueReplicationServerProperties() EnqueueReplicationServerPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *SAPCentralInstance) EnqueueReplicationServerPropertiesResponsePtrOutput {
		return v.EnqueueReplicationServerProperties
	}).(EnqueueReplicationServerPropertiesResponsePtrOutput)
}

// Defines the SAP Enqueue Server properties.
func (o SAPCentralInstanceOutput) EnqueueServerProperties() EnqueueServerPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *SAPCentralInstance) EnqueueServerPropertiesResponsePtrOutput { return v.EnqueueServerProperties }).(EnqueueServerPropertiesResponsePtrOutput)
}

// Defines the errors related to SAP Central Services Instance resource.
func (o SAPCentralInstanceOutput) Errors() SAPVirtualInstanceErrorResponseOutput {
	return o.ApplyT(func(v *SAPCentralInstance) SAPVirtualInstanceErrorResponseOutput { return v.Errors }).(SAPVirtualInstanceErrorResponseOutput)
}

// Defines the SAP Gateway Server properties.
func (o SAPCentralInstanceOutput) GatewayServerProperties() GatewayServerPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *SAPCentralInstance) GatewayServerPropertiesResponsePtrOutput { return v.GatewayServerProperties }).(GatewayServerPropertiesResponsePtrOutput)
}

// Defines the health of SAP Instances.
func (o SAPCentralInstanceOutput) Health() pulumi.StringOutput {
	return o.ApplyT(func(v *SAPCentralInstance) pulumi.StringOutput { return v.Health }).(pulumi.StringOutput)
}

// The central services instance number.
func (o SAPCentralInstanceOutput) InstanceNo() pulumi.StringOutput {
	return o.ApplyT(func(v *SAPCentralInstance) pulumi.StringOutput { return v.InstanceNo }).(pulumi.StringOutput)
}

// The central services instance Kernel Patch level.
func (o SAPCentralInstanceOutput) KernelPatch() pulumi.StringOutput {
	return o.ApplyT(func(v *SAPCentralInstance) pulumi.StringOutput { return v.KernelPatch }).(pulumi.StringOutput)
}

// The central services instance Kernel Version.
func (o SAPCentralInstanceOutput) KernelVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *SAPCentralInstance) pulumi.StringOutput { return v.KernelVersion }).(pulumi.StringOutput)
}

// The Load Balancer details such as LoadBalancer ID attached to ASCS Virtual Machines
func (o SAPCentralInstanceOutput) LoadBalancerDetails() LoadBalancerDetailsResponseOutput {
	return o.ApplyT(func(v *SAPCentralInstance) LoadBalancerDetailsResponseOutput { return v.LoadBalancerDetails }).(LoadBalancerDetailsResponseOutput)
}

// The geo-location where the resource lives
func (o SAPCentralInstanceOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *SAPCentralInstance) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Defines the SAP Message Server properties.
func (o SAPCentralInstanceOutput) MessageServerProperties() MessageServerPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *SAPCentralInstance) MessageServerPropertiesResponsePtrOutput { return v.MessageServerProperties }).(MessageServerPropertiesResponsePtrOutput)
}

// The name of the resource
func (o SAPCentralInstanceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SAPCentralInstance) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Defines the provisioning states.
func (o SAPCentralInstanceOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *SAPCentralInstance) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Defines the SAP Instance status.
func (o SAPCentralInstanceOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *SAPCentralInstance) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The central services instance subnet.
func (o SAPCentralInstanceOutput) Subnet() pulumi.StringOutput {
	return o.ApplyT(func(v *SAPCentralInstance) pulumi.StringOutput { return v.Subnet }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o SAPCentralInstanceOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *SAPCentralInstance) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o SAPCentralInstanceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SAPCentralInstance) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o SAPCentralInstanceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SAPCentralInstance) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The list of virtual machines corresponding to the Central Services instance.
func (o SAPCentralInstanceOutput) VmDetails() CentralServerVmDetailsResponseArrayOutput {
	return o.ApplyT(func(v *SAPCentralInstance) CentralServerVmDetailsResponseArrayOutput { return v.VmDetails }).(CentralServerVmDetailsResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(SAPCentralInstanceOutput{})
}
