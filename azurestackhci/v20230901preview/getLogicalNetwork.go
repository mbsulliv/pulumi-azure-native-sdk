// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The logical network resource definition.
func LookupLogicalNetwork(ctx *pulumi.Context, args *LookupLogicalNetworkArgs, opts ...pulumi.InvokeOption) (*LookupLogicalNetworkResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupLogicalNetworkResult
	err := ctx.Invoke("azure-native:azurestackhci/v20230901preview:getLogicalNetwork", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLogicalNetworkArgs struct {
	// Name of the logical network
	LogicalNetworkName string `pulumi:"logicalNetworkName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The logical network resource definition.
type LookupLogicalNetworkResult struct {
	// DhcpOptions contains an array of DNS servers available to VMs deployed in the logical network. Standard DHCP option for a subnet overrides logical network DHCP options.
	DhcpOptions *LogicalNetworkPropertiesResponseDhcpOptions `pulumi:"dhcpOptions"`
	// The extendedLocation of the resource.
	ExtendedLocation *ExtendedLocationResponse `pulumi:"extendedLocation"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Provisioning state of the logical network.
	ProvisioningState string `pulumi:"provisioningState"`
	// The observed state of logical networks
	Status LogicalNetworkStatusResponse `pulumi:"status"`
	// Subnet - list of subnets under the logical network
	Subnets []SubnetResponse `pulumi:"subnets"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// name of the network switch to be used for VMs
	VmSwitchName *string `pulumi:"vmSwitchName"`
}

func LookupLogicalNetworkOutput(ctx *pulumi.Context, args LookupLogicalNetworkOutputArgs, opts ...pulumi.InvokeOption) LookupLogicalNetworkResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupLogicalNetworkResult, error) {
			args := v.(LookupLogicalNetworkArgs)
			r, err := LookupLogicalNetwork(ctx, &args, opts...)
			var s LookupLogicalNetworkResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupLogicalNetworkResultOutput)
}

type LookupLogicalNetworkOutputArgs struct {
	// Name of the logical network
	LogicalNetworkName pulumi.StringInput `pulumi:"logicalNetworkName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupLogicalNetworkOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLogicalNetworkArgs)(nil)).Elem()
}

// The logical network resource definition.
type LookupLogicalNetworkResultOutput struct{ *pulumi.OutputState }

func (LookupLogicalNetworkResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLogicalNetworkResult)(nil)).Elem()
}

func (o LookupLogicalNetworkResultOutput) ToLookupLogicalNetworkResultOutput() LookupLogicalNetworkResultOutput {
	return o
}

func (o LookupLogicalNetworkResultOutput) ToLookupLogicalNetworkResultOutputWithContext(ctx context.Context) LookupLogicalNetworkResultOutput {
	return o
}

// DhcpOptions contains an array of DNS servers available to VMs deployed in the logical network. Standard DHCP option for a subnet overrides logical network DHCP options.
func (o LookupLogicalNetworkResultOutput) DhcpOptions() LogicalNetworkPropertiesResponseDhcpOptionsPtrOutput {
	return o.ApplyT(func(v LookupLogicalNetworkResult) *LogicalNetworkPropertiesResponseDhcpOptions { return v.DhcpOptions }).(LogicalNetworkPropertiesResponseDhcpOptionsPtrOutput)
}

// The extendedLocation of the resource.
func (o LookupLogicalNetworkResultOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v LookupLogicalNetworkResult) *ExtendedLocationResponse { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupLogicalNetworkResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLogicalNetworkResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupLogicalNetworkResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLogicalNetworkResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupLogicalNetworkResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLogicalNetworkResult) string { return v.Name }).(pulumi.StringOutput)
}

// Provisioning state of the logical network.
func (o LookupLogicalNetworkResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLogicalNetworkResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The observed state of logical networks
func (o LookupLogicalNetworkResultOutput) Status() LogicalNetworkStatusResponseOutput {
	return o.ApplyT(func(v LookupLogicalNetworkResult) LogicalNetworkStatusResponse { return v.Status }).(LogicalNetworkStatusResponseOutput)
}

// Subnet - list of subnets under the logical network
func (o LookupLogicalNetworkResultOutput) Subnets() SubnetResponseArrayOutput {
	return o.ApplyT(func(v LookupLogicalNetworkResult) []SubnetResponse { return v.Subnets }).(SubnetResponseArrayOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupLogicalNetworkResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupLogicalNetworkResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupLogicalNetworkResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupLogicalNetworkResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupLogicalNetworkResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLogicalNetworkResult) string { return v.Type }).(pulumi.StringOutput)
}

// name of the network switch to be used for VMs
func (o LookupLogicalNetworkResultOutput) VmSwitchName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLogicalNetworkResult) *string { return v.VmSwitchName }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLogicalNetworkResultOutput{})
}
