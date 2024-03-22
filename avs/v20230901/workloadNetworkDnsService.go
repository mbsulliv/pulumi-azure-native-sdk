// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230901

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// NSX DNS Service
type WorkloadNetworkDnsService struct {
	pulumi.CustomResourceState

	// Default DNS zone of the DNS Service.
	DefaultDnsZone pulumi.StringPtrOutput `pulumi:"defaultDnsZone"`
	// Display name of the DNS Service.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// DNS service IP of the DNS Service.
	DnsServiceIp pulumi.StringPtrOutput `pulumi:"dnsServiceIp"`
	// FQDN zones of the DNS Service.
	FqdnZones pulumi.StringArrayOutput `pulumi:"fqdnZones"`
	// DNS Service log level.
	LogLevel pulumi.StringPtrOutput `pulumi:"logLevel"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioning state
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// NSX revision number.
	Revision pulumi.Float64PtrOutput `pulumi:"revision"`
	// DNS Service status.
	Status pulumi.StringOutput `pulumi:"status"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewWorkloadNetworkDnsService registers a new resource with the given unique name, arguments, and options.
func NewWorkloadNetworkDnsService(ctx *pulumi.Context,
	name string, args *WorkloadNetworkDnsServiceArgs, opts ...pulumi.ResourceOption) (*WorkloadNetworkDnsService, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PrivateCloudName == nil {
		return nil, errors.New("invalid value for required argument 'PrivateCloudName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:avs:WorkloadNetworkDnsService"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20200717preview:WorkloadNetworkDnsService"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20210101preview:WorkloadNetworkDnsService"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20210601:WorkloadNetworkDnsService"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20211201:WorkloadNetworkDnsService"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20220501:WorkloadNetworkDnsService"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20230301:WorkloadNetworkDnsService"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource WorkloadNetworkDnsService
	err := ctx.RegisterResource("azure-native:avs/v20230901:WorkloadNetworkDnsService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkloadNetworkDnsService gets an existing WorkloadNetworkDnsService resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkloadNetworkDnsService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkloadNetworkDnsServiceState, opts ...pulumi.ResourceOption) (*WorkloadNetworkDnsService, error) {
	var resource WorkloadNetworkDnsService
	err := ctx.ReadResource("azure-native:avs/v20230901:WorkloadNetworkDnsService", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WorkloadNetworkDnsService resources.
type workloadNetworkDnsServiceState struct {
}

type WorkloadNetworkDnsServiceState struct {
}

func (WorkloadNetworkDnsServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*workloadNetworkDnsServiceState)(nil)).Elem()
}

type workloadNetworkDnsServiceArgs struct {
	// Default DNS zone of the DNS Service.
	DefaultDnsZone *string `pulumi:"defaultDnsZone"`
	// Display name of the DNS Service.
	DisplayName *string `pulumi:"displayName"`
	// ID of the DNS service.
	DnsServiceId *string `pulumi:"dnsServiceId"`
	// DNS service IP of the DNS Service.
	DnsServiceIp *string `pulumi:"dnsServiceIp"`
	// FQDN zones of the DNS Service.
	FqdnZones []string `pulumi:"fqdnZones"`
	// DNS Service log level.
	LogLevel *string `pulumi:"logLevel"`
	// Name of the private cloud
	PrivateCloudName string `pulumi:"privateCloudName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// NSX revision number.
	Revision *float64 `pulumi:"revision"`
}

// The set of arguments for constructing a WorkloadNetworkDnsService resource.
type WorkloadNetworkDnsServiceArgs struct {
	// Default DNS zone of the DNS Service.
	DefaultDnsZone pulumi.StringPtrInput
	// Display name of the DNS Service.
	DisplayName pulumi.StringPtrInput
	// ID of the DNS service.
	DnsServiceId pulumi.StringPtrInput
	// DNS service IP of the DNS Service.
	DnsServiceIp pulumi.StringPtrInput
	// FQDN zones of the DNS Service.
	FqdnZones pulumi.StringArrayInput
	// DNS Service log level.
	LogLevel pulumi.StringPtrInput
	// Name of the private cloud
	PrivateCloudName pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// NSX revision number.
	Revision pulumi.Float64PtrInput
}

func (WorkloadNetworkDnsServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workloadNetworkDnsServiceArgs)(nil)).Elem()
}

type WorkloadNetworkDnsServiceInput interface {
	pulumi.Input

	ToWorkloadNetworkDnsServiceOutput() WorkloadNetworkDnsServiceOutput
	ToWorkloadNetworkDnsServiceOutputWithContext(ctx context.Context) WorkloadNetworkDnsServiceOutput
}

func (*WorkloadNetworkDnsService) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkloadNetworkDnsService)(nil)).Elem()
}

func (i *WorkloadNetworkDnsService) ToWorkloadNetworkDnsServiceOutput() WorkloadNetworkDnsServiceOutput {
	return i.ToWorkloadNetworkDnsServiceOutputWithContext(context.Background())
}

func (i *WorkloadNetworkDnsService) ToWorkloadNetworkDnsServiceOutputWithContext(ctx context.Context) WorkloadNetworkDnsServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkloadNetworkDnsServiceOutput)
}

type WorkloadNetworkDnsServiceOutput struct{ *pulumi.OutputState }

func (WorkloadNetworkDnsServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkloadNetworkDnsService)(nil)).Elem()
}

func (o WorkloadNetworkDnsServiceOutput) ToWorkloadNetworkDnsServiceOutput() WorkloadNetworkDnsServiceOutput {
	return o
}

func (o WorkloadNetworkDnsServiceOutput) ToWorkloadNetworkDnsServiceOutputWithContext(ctx context.Context) WorkloadNetworkDnsServiceOutput {
	return o
}

// Default DNS zone of the DNS Service.
func (o WorkloadNetworkDnsServiceOutput) DefaultDnsZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkloadNetworkDnsService) pulumi.StringPtrOutput { return v.DefaultDnsZone }).(pulumi.StringPtrOutput)
}

// Display name of the DNS Service.
func (o WorkloadNetworkDnsServiceOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkloadNetworkDnsService) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// DNS service IP of the DNS Service.
func (o WorkloadNetworkDnsServiceOutput) DnsServiceIp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkloadNetworkDnsService) pulumi.StringPtrOutput { return v.DnsServiceIp }).(pulumi.StringPtrOutput)
}

// FQDN zones of the DNS Service.
func (o WorkloadNetworkDnsServiceOutput) FqdnZones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *WorkloadNetworkDnsService) pulumi.StringArrayOutput { return v.FqdnZones }).(pulumi.StringArrayOutput)
}

// DNS Service log level.
func (o WorkloadNetworkDnsServiceOutput) LogLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkloadNetworkDnsService) pulumi.StringPtrOutput { return v.LogLevel }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o WorkloadNetworkDnsServiceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkloadNetworkDnsService) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state
func (o WorkloadNetworkDnsServiceOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkloadNetworkDnsService) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// NSX revision number.
func (o WorkloadNetworkDnsServiceOutput) Revision() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *WorkloadNetworkDnsService) pulumi.Float64PtrOutput { return v.Revision }).(pulumi.Float64PtrOutput)
}

// DNS Service status.
func (o WorkloadNetworkDnsServiceOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkloadNetworkDnsService) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o WorkloadNetworkDnsServiceOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *WorkloadNetworkDnsService) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o WorkloadNetworkDnsServiceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkloadNetworkDnsService) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WorkloadNetworkDnsServiceOutput{})
}
