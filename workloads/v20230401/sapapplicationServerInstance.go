// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230401

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Define the SAP Application Server Instance resource.
type SAPApplicationServerInstance struct {
	pulumi.CustomResourceState

	// Defines the Application Instance errors.
	Errors SAPVirtualInstanceErrorResponseOutput `pulumi:"errors"`
	// Application server instance gateway Port.
	GatewayPort pulumi.Float64Output `pulumi:"gatewayPort"`
	// Defines the health of SAP Instances.
	Health pulumi.StringOutput `pulumi:"health"`
	// Application server instance SAP hostname.
	Hostname pulumi.StringOutput `pulumi:"hostname"`
	// Application server instance ICM HTTP Port.
	IcmHttpPort pulumi.Float64Output `pulumi:"icmHttpPort"`
	// Application server instance ICM HTTPS Port.
	IcmHttpsPort pulumi.Float64Output `pulumi:"icmHttpsPort"`
	// Application server Instance Number.
	InstanceNo pulumi.StringOutput `pulumi:"instanceNo"`
	//  Application server instance SAP IP Address.
	IpAddress pulumi.StringOutput `pulumi:"ipAddress"`
	// Application server instance SAP Kernel Patch level.
	KernelPatch pulumi.StringOutput `pulumi:"kernelPatch"`
	//  Application server instance SAP Kernel Version.
	KernelVersion pulumi.StringOutput `pulumi:"kernelVersion"`
	// The Load Balancer details such as LoadBalancer ID attached to Application Server Virtual Machines
	LoadBalancerDetails LoadBalancerDetailsResponseOutput `pulumi:"loadBalancerDetails"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Defines the provisioning states.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Defines the SAP Instance status.
	Status pulumi.StringOutput `pulumi:"status"`
	// Application server Subnet.
	Subnet pulumi.StringOutput `pulumi:"subnet"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// The list of virtual machines.
	VmDetails ApplicationServerVmDetailsResponseArrayOutput `pulumi:"vmDetails"`
}

// NewSAPApplicationServerInstance registers a new resource with the given unique name, arguments, and options.
func NewSAPApplicationServerInstance(ctx *pulumi.Context,
	name string, args *SAPApplicationServerInstanceArgs, opts ...pulumi.ResourceOption) (*SAPApplicationServerInstance, error) {
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
			Type: pulumi.String("azure-native:workloads:SAPApplicationServerInstance"),
		},
		{
			Type: pulumi.String("azure-native:workloads/v20211201preview:SAPApplicationServerInstance"),
		},
		{
			Type: pulumi.String("azure-native:workloads/v20221101preview:SAPApplicationServerInstance"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource SAPApplicationServerInstance
	err := ctx.RegisterResource("azure-native:workloads/v20230401:SAPApplicationServerInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSAPApplicationServerInstance gets an existing SAPApplicationServerInstance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSAPApplicationServerInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SAPApplicationServerInstanceState, opts ...pulumi.ResourceOption) (*SAPApplicationServerInstance, error) {
	var resource SAPApplicationServerInstance
	err := ctx.ReadResource("azure-native:workloads/v20230401:SAPApplicationServerInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SAPApplicationServerInstance resources.
type sapapplicationServerInstanceState struct {
}

type SAPApplicationServerInstanceState struct {
}

func (SAPApplicationServerInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*sapapplicationServerInstanceState)(nil)).Elem()
}

type sapapplicationServerInstanceArgs struct {
	// The name of SAP Application Server instance resource.
	ApplicationInstanceName *string `pulumi:"applicationInstanceName"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Virtual Instances for SAP solutions resource
	SapVirtualInstanceName string `pulumi:"sapVirtualInstanceName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a SAPApplicationServerInstance resource.
type SAPApplicationServerInstanceArgs struct {
	// The name of SAP Application Server instance resource.
	ApplicationInstanceName pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the Virtual Instances for SAP solutions resource
	SapVirtualInstanceName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (SAPApplicationServerInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sapapplicationServerInstanceArgs)(nil)).Elem()
}

type SAPApplicationServerInstanceInput interface {
	pulumi.Input

	ToSAPApplicationServerInstanceOutput() SAPApplicationServerInstanceOutput
	ToSAPApplicationServerInstanceOutputWithContext(ctx context.Context) SAPApplicationServerInstanceOutput
}

func (*SAPApplicationServerInstance) ElementType() reflect.Type {
	return reflect.TypeOf((**SAPApplicationServerInstance)(nil)).Elem()
}

func (i *SAPApplicationServerInstance) ToSAPApplicationServerInstanceOutput() SAPApplicationServerInstanceOutput {
	return i.ToSAPApplicationServerInstanceOutputWithContext(context.Background())
}

func (i *SAPApplicationServerInstance) ToSAPApplicationServerInstanceOutputWithContext(ctx context.Context) SAPApplicationServerInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SAPApplicationServerInstanceOutput)
}

func (i *SAPApplicationServerInstance) ToOutput(ctx context.Context) pulumix.Output[*SAPApplicationServerInstance] {
	return pulumix.Output[*SAPApplicationServerInstance]{
		OutputState: i.ToSAPApplicationServerInstanceOutputWithContext(ctx).OutputState,
	}
}

type SAPApplicationServerInstanceOutput struct{ *pulumi.OutputState }

func (SAPApplicationServerInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SAPApplicationServerInstance)(nil)).Elem()
}

func (o SAPApplicationServerInstanceOutput) ToSAPApplicationServerInstanceOutput() SAPApplicationServerInstanceOutput {
	return o
}

func (o SAPApplicationServerInstanceOutput) ToSAPApplicationServerInstanceOutputWithContext(ctx context.Context) SAPApplicationServerInstanceOutput {
	return o
}

func (o SAPApplicationServerInstanceOutput) ToOutput(ctx context.Context) pulumix.Output[*SAPApplicationServerInstance] {
	return pulumix.Output[*SAPApplicationServerInstance]{
		OutputState: o.OutputState,
	}
}

// Defines the Application Instance errors.
func (o SAPApplicationServerInstanceOutput) Errors() SAPVirtualInstanceErrorResponseOutput {
	return o.ApplyT(func(v *SAPApplicationServerInstance) SAPVirtualInstanceErrorResponseOutput { return v.Errors }).(SAPVirtualInstanceErrorResponseOutput)
}

// Application server instance gateway Port.
func (o SAPApplicationServerInstanceOutput) GatewayPort() pulumi.Float64Output {
	return o.ApplyT(func(v *SAPApplicationServerInstance) pulumi.Float64Output { return v.GatewayPort }).(pulumi.Float64Output)
}

// Defines the health of SAP Instances.
func (o SAPApplicationServerInstanceOutput) Health() pulumi.StringOutput {
	return o.ApplyT(func(v *SAPApplicationServerInstance) pulumi.StringOutput { return v.Health }).(pulumi.StringOutput)
}

// Application server instance SAP hostname.
func (o SAPApplicationServerInstanceOutput) Hostname() pulumi.StringOutput {
	return o.ApplyT(func(v *SAPApplicationServerInstance) pulumi.StringOutput { return v.Hostname }).(pulumi.StringOutput)
}

// Application server instance ICM HTTP Port.
func (o SAPApplicationServerInstanceOutput) IcmHttpPort() pulumi.Float64Output {
	return o.ApplyT(func(v *SAPApplicationServerInstance) pulumi.Float64Output { return v.IcmHttpPort }).(pulumi.Float64Output)
}

// Application server instance ICM HTTPS Port.
func (o SAPApplicationServerInstanceOutput) IcmHttpsPort() pulumi.Float64Output {
	return o.ApplyT(func(v *SAPApplicationServerInstance) pulumi.Float64Output { return v.IcmHttpsPort }).(pulumi.Float64Output)
}

// Application server Instance Number.
func (o SAPApplicationServerInstanceOutput) InstanceNo() pulumi.StringOutput {
	return o.ApplyT(func(v *SAPApplicationServerInstance) pulumi.StringOutput { return v.InstanceNo }).(pulumi.StringOutput)
}

// Application server instance SAP IP Address.
func (o SAPApplicationServerInstanceOutput) IpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *SAPApplicationServerInstance) pulumi.StringOutput { return v.IpAddress }).(pulumi.StringOutput)
}

// Application server instance SAP Kernel Patch level.
func (o SAPApplicationServerInstanceOutput) KernelPatch() pulumi.StringOutput {
	return o.ApplyT(func(v *SAPApplicationServerInstance) pulumi.StringOutput { return v.KernelPatch }).(pulumi.StringOutput)
}

// Application server instance SAP Kernel Version.
func (o SAPApplicationServerInstanceOutput) KernelVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *SAPApplicationServerInstance) pulumi.StringOutput { return v.KernelVersion }).(pulumi.StringOutput)
}

// The Load Balancer details such as LoadBalancer ID attached to Application Server Virtual Machines
func (o SAPApplicationServerInstanceOutput) LoadBalancerDetails() LoadBalancerDetailsResponseOutput {
	return o.ApplyT(func(v *SAPApplicationServerInstance) LoadBalancerDetailsResponseOutput { return v.LoadBalancerDetails }).(LoadBalancerDetailsResponseOutput)
}

// The geo-location where the resource lives
func (o SAPApplicationServerInstanceOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *SAPApplicationServerInstance) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o SAPApplicationServerInstanceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SAPApplicationServerInstance) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Defines the provisioning states.
func (o SAPApplicationServerInstanceOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *SAPApplicationServerInstance) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Defines the SAP Instance status.
func (o SAPApplicationServerInstanceOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *SAPApplicationServerInstance) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Application server Subnet.
func (o SAPApplicationServerInstanceOutput) Subnet() pulumi.StringOutput {
	return o.ApplyT(func(v *SAPApplicationServerInstance) pulumi.StringOutput { return v.Subnet }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o SAPApplicationServerInstanceOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *SAPApplicationServerInstance) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o SAPApplicationServerInstanceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SAPApplicationServerInstance) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o SAPApplicationServerInstanceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SAPApplicationServerInstance) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The list of virtual machines.
func (o SAPApplicationServerInstanceOutput) VmDetails() ApplicationServerVmDetailsResponseArrayOutput {
	return o.ApplyT(func(v *SAPApplicationServerInstance) ApplicationServerVmDetailsResponseArrayOutput {
		return v.VmDetails
	}).(ApplicationServerVmDetailsResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(SAPApplicationServerInstanceOutput{})
}
