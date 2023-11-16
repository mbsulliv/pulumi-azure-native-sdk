// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230615

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Implements L2 Isolation Domain GET method.
func LookupL2IsolationDomain(ctx *pulumi.Context, args *LookupL2IsolationDomainArgs, opts ...pulumi.InvokeOption) (*LookupL2IsolationDomainResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupL2IsolationDomainResult
	err := ctx.Invoke("azure-native:managednetworkfabric/v20230615:getL2IsolationDomain", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupL2IsolationDomainArgs struct {
	// Name of the L2 Isolation Domain.
	L2IsolationDomainName string `pulumi:"l2IsolationDomainName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The L2 Isolation Domain resource definition.
type LookupL2IsolationDomainResult struct {
	// Administrative state of the resource.
	AdministrativeState string `pulumi:"administrativeState"`
	// Switch configuration description.
	Annotation *string `pulumi:"annotation"`
	// Configuration state of the resource.
	ConfigurationState string `pulumi:"configurationState"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// Maximum transmission unit. Default value is 1500.
	Mtu *int `pulumi:"mtu"`
	// The name of the resource
	Name string `pulumi:"name"`
	// ARM Resource ID of the Network Fabric.
	NetworkFabricId string `pulumi:"networkFabricId"`
	// Provisioning state of the resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// Vlan Identifier of the Network Fabric. Example: 501.
	VlanId int `pulumi:"vlanId"`
}

// Defaults sets the appropriate defaults for LookupL2IsolationDomainResult
func (val *LookupL2IsolationDomainResult) Defaults() *LookupL2IsolationDomainResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.Mtu == nil {
		mtu_ := 1500
		tmp.Mtu = &mtu_
	}
	return &tmp
}

func LookupL2IsolationDomainOutput(ctx *pulumi.Context, args LookupL2IsolationDomainOutputArgs, opts ...pulumi.InvokeOption) LookupL2IsolationDomainResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupL2IsolationDomainResult, error) {
			args := v.(LookupL2IsolationDomainArgs)
			r, err := LookupL2IsolationDomain(ctx, &args, opts...)
			var s LookupL2IsolationDomainResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupL2IsolationDomainResultOutput)
}

type LookupL2IsolationDomainOutputArgs struct {
	// Name of the L2 Isolation Domain.
	L2IsolationDomainName pulumi.StringInput `pulumi:"l2IsolationDomainName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupL2IsolationDomainOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupL2IsolationDomainArgs)(nil)).Elem()
}

// The L2 Isolation Domain resource definition.
type LookupL2IsolationDomainResultOutput struct{ *pulumi.OutputState }

func (LookupL2IsolationDomainResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupL2IsolationDomainResult)(nil)).Elem()
}

func (o LookupL2IsolationDomainResultOutput) ToLookupL2IsolationDomainResultOutput() LookupL2IsolationDomainResultOutput {
	return o
}

func (o LookupL2IsolationDomainResultOutput) ToLookupL2IsolationDomainResultOutputWithContext(ctx context.Context) LookupL2IsolationDomainResultOutput {
	return o
}

// Administrative state of the resource.
func (o LookupL2IsolationDomainResultOutput) AdministrativeState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupL2IsolationDomainResult) string { return v.AdministrativeState }).(pulumi.StringOutput)
}

// Switch configuration description.
func (o LookupL2IsolationDomainResultOutput) Annotation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupL2IsolationDomainResult) *string { return v.Annotation }).(pulumi.StringPtrOutput)
}

// Configuration state of the resource.
func (o LookupL2IsolationDomainResultOutput) ConfigurationState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupL2IsolationDomainResult) string { return v.ConfigurationState }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupL2IsolationDomainResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupL2IsolationDomainResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupL2IsolationDomainResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupL2IsolationDomainResult) string { return v.Location }).(pulumi.StringOutput)
}

// Maximum transmission unit. Default value is 1500.
func (o LookupL2IsolationDomainResultOutput) Mtu() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupL2IsolationDomainResult) *int { return v.Mtu }).(pulumi.IntPtrOutput)
}

// The name of the resource
func (o LookupL2IsolationDomainResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupL2IsolationDomainResult) string { return v.Name }).(pulumi.StringOutput)
}

// ARM Resource ID of the Network Fabric.
func (o LookupL2IsolationDomainResultOutput) NetworkFabricId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupL2IsolationDomainResult) string { return v.NetworkFabricId }).(pulumi.StringOutput)
}

// Provisioning state of the resource.
func (o LookupL2IsolationDomainResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupL2IsolationDomainResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupL2IsolationDomainResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupL2IsolationDomainResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupL2IsolationDomainResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupL2IsolationDomainResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupL2IsolationDomainResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupL2IsolationDomainResult) string { return v.Type }).(pulumi.StringOutput)
}

// Vlan Identifier of the Network Fabric. Example: 501.
func (o LookupL2IsolationDomainResultOutput) VlanId() pulumi.IntOutput {
	return o.ApplyT(func(v LookupL2IsolationDomainResult) int { return v.VlanId }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupL2IsolationDomainResultOutput{})
}
