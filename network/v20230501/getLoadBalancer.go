// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets the specified load balancer.
func LookupLoadBalancer(ctx *pulumi.Context, args *LookupLoadBalancerArgs, opts ...pulumi.InvokeOption) (*LookupLoadBalancerResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupLoadBalancerResult
	err := ctx.Invoke("azure-native:network/v20230501:getLoadBalancer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLoadBalancerArgs struct {
	// Expands referenced resources.
	Expand *string `pulumi:"expand"`
	// The name of the load balancer.
	LoadBalancerName string `pulumi:"loadBalancerName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// LoadBalancer resource.
type LookupLoadBalancerResult struct {
	// Collection of backend address pools used by a load balancer.
	BackendAddressPools []BackendAddressPoolResponse `pulumi:"backendAddressPools"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// The extended location of the load balancer.
	ExtendedLocation *ExtendedLocationResponse `pulumi:"extendedLocation"`
	// Object representing the frontend IPs to be used for the load balancer.
	FrontendIPConfigurations []FrontendIPConfigurationResponse `pulumi:"frontendIPConfigurations"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// Defines an external port range for inbound NAT to a single backend port on NICs associated with a load balancer. Inbound NAT rules are created automatically for each NIC associated with the Load Balancer using an external port from this range. Defining an Inbound NAT pool on your Load Balancer is mutually exclusive with defining inbound NAT rules. Inbound NAT pools are referenced from virtual machine scale sets. NICs that are associated with individual virtual machines cannot reference an inbound NAT pool. They have to reference individual inbound NAT rules.
	InboundNatPools []InboundNatPoolResponse `pulumi:"inboundNatPools"`
	// Collection of inbound NAT Rules used by a load balancer. Defining inbound NAT rules on your load balancer is mutually exclusive with defining an inbound NAT pool. Inbound NAT pools are referenced from virtual machine scale sets. NICs that are associated with individual virtual machines cannot reference an Inbound NAT pool. They have to reference individual inbound NAT rules.
	InboundNatRules []InboundNatRuleResponse `pulumi:"inboundNatRules"`
	// Object collection representing the load balancing rules Gets the provisioning.
	LoadBalancingRules []LoadBalancingRuleResponse `pulumi:"loadBalancingRules"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Resource name.
	Name string `pulumi:"name"`
	// The outbound rules.
	OutboundRules []OutboundRuleResponse `pulumi:"outboundRules"`
	// Collection of probe objects used in the load balancer.
	Probes []ProbeResponse `pulumi:"probes"`
	// The provisioning state of the load balancer resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// The resource GUID property of the load balancer resource.
	ResourceGuid string `pulumi:"resourceGuid"`
	// The load balancer SKU.
	Sku *LoadBalancerSkuResponse `pulumi:"sku"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupLoadBalancerOutput(ctx *pulumi.Context, args LookupLoadBalancerOutputArgs, opts ...pulumi.InvokeOption) LookupLoadBalancerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupLoadBalancerResult, error) {
			args := v.(LookupLoadBalancerArgs)
			r, err := LookupLoadBalancer(ctx, &args, opts...)
			var s LookupLoadBalancerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupLoadBalancerResultOutput)
}

type LookupLoadBalancerOutputArgs struct {
	// Expands referenced resources.
	Expand pulumi.StringPtrInput `pulumi:"expand"`
	// The name of the load balancer.
	LoadBalancerName pulumi.StringInput `pulumi:"loadBalancerName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupLoadBalancerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLoadBalancerArgs)(nil)).Elem()
}

// LoadBalancer resource.
type LookupLoadBalancerResultOutput struct{ *pulumi.OutputState }

func (LookupLoadBalancerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLoadBalancerResult)(nil)).Elem()
}

func (o LookupLoadBalancerResultOutput) ToLookupLoadBalancerResultOutput() LookupLoadBalancerResultOutput {
	return o
}

func (o LookupLoadBalancerResultOutput) ToLookupLoadBalancerResultOutputWithContext(ctx context.Context) LookupLoadBalancerResultOutput {
	return o
}

func (o LookupLoadBalancerResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupLoadBalancerResult] {
	return pulumix.Output[LookupLoadBalancerResult]{
		OutputState: o.OutputState,
	}
}

// Collection of backend address pools used by a load balancer.
func (o LookupLoadBalancerResultOutput) BackendAddressPools() BackendAddressPoolResponseArrayOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) []BackendAddressPoolResponse { return v.BackendAddressPools }).(BackendAddressPoolResponseArrayOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o LookupLoadBalancerResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) string { return v.Etag }).(pulumi.StringOutput)
}

// The extended location of the load balancer.
func (o LookupLoadBalancerResultOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) *ExtendedLocationResponse { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

// Object representing the frontend IPs to be used for the load balancer.
func (o LookupLoadBalancerResultOutput) FrontendIPConfigurations() FrontendIPConfigurationResponseArrayOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) []FrontendIPConfigurationResponse { return v.FrontendIPConfigurations }).(FrontendIPConfigurationResponseArrayOutput)
}

// Resource ID.
func (o LookupLoadBalancerResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// Defines an external port range for inbound NAT to a single backend port on NICs associated with a load balancer. Inbound NAT rules are created automatically for each NIC associated with the Load Balancer using an external port from this range. Defining an Inbound NAT pool on your Load Balancer is mutually exclusive with defining inbound NAT rules. Inbound NAT pools are referenced from virtual machine scale sets. NICs that are associated with individual virtual machines cannot reference an inbound NAT pool. They have to reference individual inbound NAT rules.
func (o LookupLoadBalancerResultOutput) InboundNatPools() InboundNatPoolResponseArrayOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) []InboundNatPoolResponse { return v.InboundNatPools }).(InboundNatPoolResponseArrayOutput)
}

// Collection of inbound NAT Rules used by a load balancer. Defining inbound NAT rules on your load balancer is mutually exclusive with defining an inbound NAT pool. Inbound NAT pools are referenced from virtual machine scale sets. NICs that are associated with individual virtual machines cannot reference an Inbound NAT pool. They have to reference individual inbound NAT rules.
func (o LookupLoadBalancerResultOutput) InboundNatRules() InboundNatRuleResponseArrayOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) []InboundNatRuleResponse { return v.InboundNatRules }).(InboundNatRuleResponseArrayOutput)
}

// Object collection representing the load balancing rules Gets the provisioning.
func (o LookupLoadBalancerResultOutput) LoadBalancingRules() LoadBalancingRuleResponseArrayOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) []LoadBalancingRuleResponse { return v.LoadBalancingRules }).(LoadBalancingRuleResponseArrayOutput)
}

// Resource location.
func (o LookupLoadBalancerResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o LookupLoadBalancerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) string { return v.Name }).(pulumi.StringOutput)
}

// The outbound rules.
func (o LookupLoadBalancerResultOutput) OutboundRules() OutboundRuleResponseArrayOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) []OutboundRuleResponse { return v.OutboundRules }).(OutboundRuleResponseArrayOutput)
}

// Collection of probe objects used in the load balancer.
func (o LookupLoadBalancerResultOutput) Probes() ProbeResponseArrayOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) []ProbeResponse { return v.Probes }).(ProbeResponseArrayOutput)
}

// The provisioning state of the load balancer resource.
func (o LookupLoadBalancerResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The resource GUID property of the load balancer resource.
func (o LookupLoadBalancerResultOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) string { return v.ResourceGuid }).(pulumi.StringOutput)
}

// The load balancer SKU.
func (o LookupLoadBalancerResultOutput) Sku() LoadBalancerSkuResponsePtrOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) *LoadBalancerSkuResponse { return v.Sku }).(LoadBalancerSkuResponsePtrOutput)
}

// Resource tags.
func (o LookupLoadBalancerResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o LookupLoadBalancerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLoadBalancerResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLoadBalancerResultOutput{})
}
