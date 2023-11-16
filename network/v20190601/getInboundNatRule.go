// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20190601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the specified load balancer inbound nat rule.
func LookupInboundNatRule(ctx *pulumi.Context, args *LookupInboundNatRuleArgs, opts ...pulumi.InvokeOption) (*LookupInboundNatRuleResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupInboundNatRuleResult
	err := ctx.Invoke("azure-native:network/v20190601:getInboundNatRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupInboundNatRuleArgs struct {
	// Expands referenced resources.
	Expand *string `pulumi:"expand"`
	// The name of the inbound nat rule.
	InboundNatRuleName string `pulumi:"inboundNatRuleName"`
	// The name of the load balancer.
	LoadBalancerName string `pulumi:"loadBalancerName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Inbound NAT rule of the load balancer.
type LookupInboundNatRuleResult struct {
	// A reference to a private IP address defined on a network interface of a VM. Traffic sent to the frontend port of each of the frontend IP configurations is forwarded to the backend IP.
	BackendIPConfiguration NetworkInterfaceIPConfigurationResponse `pulumi:"backendIPConfiguration"`
	// The port used for the internal endpoint. Acceptable values range from 1 to 65535.
	BackendPort *int `pulumi:"backendPort"`
	// Configures a virtual machine's endpoint for the floating IP capability required to configure a SQL AlwaysOn Availability Group. This setting is required when using the SQL AlwaysOn Availability Groups in SQL server. This setting can't be changed after you create the endpoint.
	EnableFloatingIP *bool `pulumi:"enableFloatingIP"`
	// Receive bidirectional TCP Reset on TCP flow idle timeout or unexpected connection termination. This element is only used when the protocol is set to TCP.
	EnableTcpReset *bool `pulumi:"enableTcpReset"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// A reference to frontend IP addresses.
	FrontendIPConfiguration *SubResourceResponse `pulumi:"frontendIPConfiguration"`
	// The port for the external endpoint. Port numbers for each rule must be unique within the Load Balancer. Acceptable values range from 1 to 65534.
	FrontendPort *int `pulumi:"frontendPort"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// The timeout for the TCP idle connection. The value can be set between 4 and 30 minutes. The default value is 4 minutes. This element is only used when the protocol is set to TCP.
	IdleTimeoutInMinutes *int `pulumi:"idleTimeoutInMinutes"`
	// Gets name of the resource that is unique within the set of inbound NAT rules used by the load balancer. This name can be used to access the resource.
	Name *string `pulumi:"name"`
	// The reference to the transport protocol used by the load balancing rule.
	Protocol *string `pulumi:"protocol"`
	// Gets the provisioning state of the public IP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Type of the resource.
	Type string `pulumi:"type"`
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
	// The name of the inbound nat rule.
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
func (o LookupInboundNatRuleResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInboundNatRuleResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// A reference to frontend IP addresses.
func (o LookupInboundNatRuleResultOutput) FrontendIPConfiguration() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupInboundNatRuleResult) *SubResourceResponse { return v.FrontendIPConfiguration }).(SubResourceResponsePtrOutput)
}

// The port for the external endpoint. Port numbers for each rule must be unique within the Load Balancer. Acceptable values range from 1 to 65534.
func (o LookupInboundNatRuleResultOutput) FrontendPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupInboundNatRuleResult) *int { return v.FrontendPort }).(pulumi.IntPtrOutput)
}

// Resource ID.
func (o LookupInboundNatRuleResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInboundNatRuleResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// The timeout for the TCP idle connection. The value can be set between 4 and 30 minutes. The default value is 4 minutes. This element is only used when the protocol is set to TCP.
func (o LookupInboundNatRuleResultOutput) IdleTimeoutInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupInboundNatRuleResult) *int { return v.IdleTimeoutInMinutes }).(pulumi.IntPtrOutput)
}

// Gets name of the resource that is unique within the set of inbound NAT rules used by the load balancer. This name can be used to access the resource.
func (o LookupInboundNatRuleResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInboundNatRuleResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The reference to the transport protocol used by the load balancing rule.
func (o LookupInboundNatRuleResultOutput) Protocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInboundNatRuleResult) *string { return v.Protocol }).(pulumi.StringPtrOutput)
}

// Gets the provisioning state of the public IP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
func (o LookupInboundNatRuleResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInboundNatRuleResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

// Type of the resource.
func (o LookupInboundNatRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInboundNatRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupInboundNatRuleResultOutput{})
}
