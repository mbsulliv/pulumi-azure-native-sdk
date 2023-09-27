// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets the specified load balancer inbound NAT rule.
func LookupInboundNatRule(ctx *pulumi.Context, args *LookupInboundNatRuleArgs, opts ...pulumi.InvokeOption) (*LookupInboundNatRuleResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupInboundNatRuleResult
	err := ctx.Invoke("azure-native:network/v20230201:getInboundNatRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupInboundNatRuleArgs struct {
	// Expands referenced resources.
	Expand *string `pulumi:"expand"`
	// The name of the inbound NAT rule.
	InboundNatRuleName string `pulumi:"inboundNatRuleName"`
	// The name of the load balancer.
	LoadBalancerName string `pulumi:"loadBalancerName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Inbound NAT rule of the load balancer.
type LookupInboundNatRuleResult struct {
	// A reference to backendAddressPool resource.
	BackendAddressPool *SubResourceResponse `pulumi:"backendAddressPool"`
	// A reference to a private IP address defined on a network interface of a VM. Traffic sent to the frontend port of each of the frontend IP configurations is forwarded to the backend IP.
	BackendIPConfiguration NetworkInterfaceIPConfigurationResponse `pulumi:"backendIPConfiguration"`
	// The port used for the internal endpoint. Acceptable values range from 1 to 65535.
	BackendPort *int `pulumi:"backendPort"`
	// Configures a virtual machine's endpoint for the floating IP capability required to configure a SQL AlwaysOn Availability Group. This setting is required when using the SQL AlwaysOn Availability Groups in SQL server. This setting can't be changed after you create the endpoint.
	EnableFloatingIP *bool `pulumi:"enableFloatingIP"`
	// Receive bidirectional TCP Reset on TCP flow idle timeout or unexpected connection termination. This element is only used when the protocol is set to TCP.
	EnableTcpReset *bool `pulumi:"enableTcpReset"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// A reference to frontend IP addresses.
	FrontendIPConfiguration *SubResourceResponse `pulumi:"frontendIPConfiguration"`
	// The port for the external endpoint. Port numbers for each rule must be unique within the Load Balancer. Acceptable values range from 1 to 65534.
	FrontendPort *int `pulumi:"frontendPort"`
	// The port range end for the external endpoint. This property is used together with BackendAddressPool and FrontendPortRangeStart. Individual inbound NAT rule port mappings will be created for each backend address from BackendAddressPool. Acceptable values range from 1 to 65534.
	FrontendPortRangeEnd *int `pulumi:"frontendPortRangeEnd"`
	// The port range start for the external endpoint. This property is used together with BackendAddressPool and FrontendPortRangeEnd. Individual inbound NAT rule port mappings will be created for each backend address from BackendAddressPool. Acceptable values range from 1 to 65534.
	FrontendPortRangeStart *int `pulumi:"frontendPortRangeStart"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// The timeout for the TCP idle connection. The value can be set between 4 and 30 minutes. The default value is 4 minutes. This element is only used when the protocol is set to TCP.
	IdleTimeoutInMinutes *int `pulumi:"idleTimeoutInMinutes"`
	// The name of the resource that is unique within the set of inbound NAT rules used by the load balancer. This name can be used to access the resource.
	Name *string `pulumi:"name"`
	// The reference to the transport protocol used by the load balancing rule.
	Protocol *string `pulumi:"protocol"`
	// The provisioning state of the inbound NAT rule resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Type of the resource.
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupInboundNatRuleResult
func (val *LookupInboundNatRuleResult) Defaults() *LookupInboundNatRuleResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.BackendIPConfiguration = *tmp.BackendIPConfiguration.Defaults()

	return &tmp
}

func LookupInboundNatRuleOutput(ctx *pulumi.Context, args LookupInboundNatRuleOutputArgs, opts ...pulumi.InvokeOption) LookupInboundNatRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupInboundNatRuleResult, error) {
			args := v.(LookupInboundNatRuleArgs)
			r, err := LookupInboundNatRule(ctx, &args, opts...)
			var s LookupInboundNatRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupInboundNatRuleResultOutput)
}

type LookupInboundNatRuleOutputArgs struct {
	// Expands referenced resources.
	Expand pulumi.StringPtrInput `pulumi:"expand"`
	// The name of the inbound NAT rule.
	InboundNatRuleName pulumi.StringInput `pulumi:"inboundNatRuleName"`
	// The name of the load balancer.
	LoadBalancerName pulumi.StringInput `pulumi:"loadBalancerName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupInboundNatRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInboundNatRuleArgs)(nil)).Elem()
}

// Inbound NAT rule of the load balancer.
type LookupInboundNatRuleResultOutput struct{ *pulumi.OutputState }

func (LookupInboundNatRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInboundNatRuleResult)(nil)).Elem()
}

func (o LookupInboundNatRuleResultOutput) ToLookupInboundNatRuleResultOutput() LookupInboundNatRuleResultOutput {
	return o
}

func (o LookupInboundNatRuleResultOutput) ToLookupInboundNatRuleResultOutputWithContext(ctx context.Context) LookupInboundNatRuleResultOutput {
	return o
}

func (o LookupInboundNatRuleResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupInboundNatRuleResult] {
	return pulumix.Output[LookupInboundNatRuleResult]{
		OutputState: o.OutputState,
	}
}

// A reference to backendAddressPool resource.
func (o LookupInboundNatRuleResultOutput) BackendAddressPool() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupInboundNatRuleResult) *SubResourceResponse { return v.BackendAddressPool }).(SubResourceResponsePtrOutput)
}

// A reference to a private IP address defined on a network interface of a VM. Traffic sent to the frontend port of each of the frontend IP configurations is forwarded to the backend IP.
func (o LookupInboundNatRuleResultOutput) BackendIPConfiguration() NetworkInterfaceIPConfigurationResponseOutput {
	return o.ApplyT(func(v LookupInboundNatRuleResult) NetworkInterfaceIPConfigurationResponse {
		return v.BackendIPConfiguration
	}).(NetworkInterfaceIPConfigurationResponseOutput)
}

// The port used for the internal endpoint. Acceptable values range from 1 to 65535.
func (o LookupInboundNatRuleResultOutput) BackendPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupInboundNatRuleResult) *int { return v.BackendPort }).(pulumi.IntPtrOutput)
}

// Configures a virtual machine's endpoint for the floating IP capability required to configure a SQL AlwaysOn Availability Group. This setting is required when using the SQL AlwaysOn Availability Groups in SQL server. This setting can't be changed after you create the endpoint.
func (o LookupInboundNatRuleResultOutput) EnableFloatingIP() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupInboundNatRuleResult) *bool { return v.EnableFloatingIP }).(pulumi.BoolPtrOutput)
}

// Receive bidirectional TCP Reset on TCP flow idle timeout or unexpected connection termination. This element is only used when the protocol is set to TCP.
func (o LookupInboundNatRuleResultOutput) EnableTcpReset() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupInboundNatRuleResult) *bool { return v.EnableTcpReset }).(pulumi.BoolPtrOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o LookupInboundNatRuleResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInboundNatRuleResult) string { return v.Etag }).(pulumi.StringOutput)
}

// A reference to frontend IP addresses.
func (o LookupInboundNatRuleResultOutput) FrontendIPConfiguration() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupInboundNatRuleResult) *SubResourceResponse { return v.FrontendIPConfiguration }).(SubResourceResponsePtrOutput)
}

// The port for the external endpoint. Port numbers for each rule must be unique within the Load Balancer. Acceptable values range from 1 to 65534.
func (o LookupInboundNatRuleResultOutput) FrontendPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupInboundNatRuleResult) *int { return v.FrontendPort }).(pulumi.IntPtrOutput)
}

// The port range end for the external endpoint. This property is used together with BackendAddressPool and FrontendPortRangeStart. Individual inbound NAT rule port mappings will be created for each backend address from BackendAddressPool. Acceptable values range from 1 to 65534.
func (o LookupInboundNatRuleResultOutput) FrontendPortRangeEnd() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupInboundNatRuleResult) *int { return v.FrontendPortRangeEnd }).(pulumi.IntPtrOutput)
}

// The port range start for the external endpoint. This property is used together with BackendAddressPool and FrontendPortRangeEnd. Individual inbound NAT rule port mappings will be created for each backend address from BackendAddressPool. Acceptable values range from 1 to 65534.
func (o LookupInboundNatRuleResultOutput) FrontendPortRangeStart() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupInboundNatRuleResult) *int { return v.FrontendPortRangeStart }).(pulumi.IntPtrOutput)
}

// Resource ID.
func (o LookupInboundNatRuleResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInboundNatRuleResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// The timeout for the TCP idle connection. The value can be set between 4 and 30 minutes. The default value is 4 minutes. This element is only used when the protocol is set to TCP.
func (o LookupInboundNatRuleResultOutput) IdleTimeoutInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupInboundNatRuleResult) *int { return v.IdleTimeoutInMinutes }).(pulumi.IntPtrOutput)
}

// The name of the resource that is unique within the set of inbound NAT rules used by the load balancer. This name can be used to access the resource.
func (o LookupInboundNatRuleResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInboundNatRuleResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The reference to the transport protocol used by the load balancing rule.
func (o LookupInboundNatRuleResultOutput) Protocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInboundNatRuleResult) *string { return v.Protocol }).(pulumi.StringPtrOutput)
}

// The provisioning state of the inbound NAT rule resource.
func (o LookupInboundNatRuleResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInboundNatRuleResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Type of the resource.
func (o LookupInboundNatRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInboundNatRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupInboundNatRuleResultOutput{})
}
