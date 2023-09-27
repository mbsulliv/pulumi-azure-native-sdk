// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets the specified private endpoint by resource group.
func LookupPrivateEndpoint(ctx *pulumi.Context, args *LookupPrivateEndpointArgs, opts ...pulumi.InvokeOption) (*LookupPrivateEndpointResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupPrivateEndpointResult
	err := ctx.Invoke("azure-native:network/v20230401:getPrivateEndpoint", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupPrivateEndpointArgs struct {
	// Expands referenced resources.
	Expand *string `pulumi:"expand"`
	// The name of the private endpoint.
	PrivateEndpointName string `pulumi:"privateEndpointName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Private endpoint resource.
type LookupPrivateEndpointResult struct {
	// Application security groups in which the private endpoint IP configuration is included.
	ApplicationSecurityGroups []ApplicationSecurityGroupResponse `pulumi:"applicationSecurityGroups"`
	// An array of custom dns configurations.
	CustomDnsConfigs []CustomDnsConfigPropertiesFormatResponse `pulumi:"customDnsConfigs"`
	// The custom name of the network interface attached to the private endpoint.
	CustomNetworkInterfaceName *string `pulumi:"customNetworkInterfaceName"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// The extended location of the load balancer.
	ExtendedLocation *ExtendedLocationResponse `pulumi:"extendedLocation"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// A list of IP configurations of the private endpoint. This will be used to map to the First Party Service's endpoints.
	IpConfigurations []PrivateEndpointIPConfigurationResponse `pulumi:"ipConfigurations"`
	// Resource location.
	Location *string `pulumi:"location"`
	// A grouping of information about the connection to the remote resource. Used when the network admin does not have access to approve connections to the remote resource.
	ManualPrivateLinkServiceConnections []PrivateLinkServiceConnectionResponse `pulumi:"manualPrivateLinkServiceConnections"`
	// Resource name.
	Name string `pulumi:"name"`
	// An array of references to the network interfaces created for this private endpoint.
	NetworkInterfaces []NetworkInterfaceResponse `pulumi:"networkInterfaces"`
	// A grouping of information about the connection to the remote resource.
	PrivateLinkServiceConnections []PrivateLinkServiceConnectionResponse `pulumi:"privateLinkServiceConnections"`
	// The provisioning state of the private endpoint resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// The ID of the subnet from which the private IP will be allocated.
	Subnet *SubnetResponse `pulumi:"subnet"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupPrivateEndpointResult
func (val *LookupPrivateEndpointResult) Defaults() *LookupPrivateEndpointResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Subnet = tmp.Subnet.Defaults()

	return &tmp
}

func LookupPrivateEndpointOutput(ctx *pulumi.Context, args LookupPrivateEndpointOutputArgs, opts ...pulumi.InvokeOption) LookupPrivateEndpointResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPrivateEndpointResult, error) {
			args := v.(LookupPrivateEndpointArgs)
			r, err := LookupPrivateEndpoint(ctx, &args, opts...)
			var s LookupPrivateEndpointResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPrivateEndpointResultOutput)
}

type LookupPrivateEndpointOutputArgs struct {
	// Expands referenced resources.
	Expand pulumi.StringPtrInput `pulumi:"expand"`
	// The name of the private endpoint.
	PrivateEndpointName pulumi.StringInput `pulumi:"privateEndpointName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupPrivateEndpointOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateEndpointArgs)(nil)).Elem()
}

// Private endpoint resource.
type LookupPrivateEndpointResultOutput struct{ *pulumi.OutputState }

func (LookupPrivateEndpointResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateEndpointResult)(nil)).Elem()
}

func (o LookupPrivateEndpointResultOutput) ToLookupPrivateEndpointResultOutput() LookupPrivateEndpointResultOutput {
	return o
}

func (o LookupPrivateEndpointResultOutput) ToLookupPrivateEndpointResultOutputWithContext(ctx context.Context) LookupPrivateEndpointResultOutput {
	return o
}

func (o LookupPrivateEndpointResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupPrivateEndpointResult] {
	return pulumix.Output[LookupPrivateEndpointResult]{
		OutputState: o.OutputState,
	}
}

// Application security groups in which the private endpoint IP configuration is included.
func (o LookupPrivateEndpointResultOutput) ApplicationSecurityGroups() ApplicationSecurityGroupResponseArrayOutput {
	return o.ApplyT(func(v LookupPrivateEndpointResult) []ApplicationSecurityGroupResponse {
		return v.ApplicationSecurityGroups
	}).(ApplicationSecurityGroupResponseArrayOutput)
}

// An array of custom dns configurations.
func (o LookupPrivateEndpointResultOutput) CustomDnsConfigs() CustomDnsConfigPropertiesFormatResponseArrayOutput {
	return o.ApplyT(func(v LookupPrivateEndpointResult) []CustomDnsConfigPropertiesFormatResponse {
		return v.CustomDnsConfigs
	}).(CustomDnsConfigPropertiesFormatResponseArrayOutput)
}

// The custom name of the network interface attached to the private endpoint.
func (o LookupPrivateEndpointResultOutput) CustomNetworkInterfaceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPrivateEndpointResult) *string { return v.CustomNetworkInterfaceName }).(pulumi.StringPtrOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o LookupPrivateEndpointResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointResult) string { return v.Etag }).(pulumi.StringOutput)
}

// The extended location of the load balancer.
func (o LookupPrivateEndpointResultOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v LookupPrivateEndpointResult) *ExtendedLocationResponse { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

// Resource ID.
func (o LookupPrivateEndpointResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPrivateEndpointResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// A list of IP configurations of the private endpoint. This will be used to map to the First Party Service's endpoints.
func (o LookupPrivateEndpointResultOutput) IpConfigurations() PrivateEndpointIPConfigurationResponseArrayOutput {
	return o.ApplyT(func(v LookupPrivateEndpointResult) []PrivateEndpointIPConfigurationResponse {
		return v.IpConfigurations
	}).(PrivateEndpointIPConfigurationResponseArrayOutput)
}

// Resource location.
func (o LookupPrivateEndpointResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPrivateEndpointResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// A grouping of information about the connection to the remote resource. Used when the network admin does not have access to approve connections to the remote resource.
func (o LookupPrivateEndpointResultOutput) ManualPrivateLinkServiceConnections() PrivateLinkServiceConnectionResponseArrayOutput {
	return o.ApplyT(func(v LookupPrivateEndpointResult) []PrivateLinkServiceConnectionResponse {
		return v.ManualPrivateLinkServiceConnections
	}).(PrivateLinkServiceConnectionResponseArrayOutput)
}

// Resource name.
func (o LookupPrivateEndpointResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointResult) string { return v.Name }).(pulumi.StringOutput)
}

// An array of references to the network interfaces created for this private endpoint.
func (o LookupPrivateEndpointResultOutput) NetworkInterfaces() NetworkInterfaceResponseArrayOutput {
	return o.ApplyT(func(v LookupPrivateEndpointResult) []NetworkInterfaceResponse { return v.NetworkInterfaces }).(NetworkInterfaceResponseArrayOutput)
}

// A grouping of information about the connection to the remote resource.
func (o LookupPrivateEndpointResultOutput) PrivateLinkServiceConnections() PrivateLinkServiceConnectionResponseArrayOutput {
	return o.ApplyT(func(v LookupPrivateEndpointResult) []PrivateLinkServiceConnectionResponse {
		return v.PrivateLinkServiceConnections
	}).(PrivateLinkServiceConnectionResponseArrayOutput)
}

// The provisioning state of the private endpoint resource.
func (o LookupPrivateEndpointResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The ID of the subnet from which the private IP will be allocated.
func (o LookupPrivateEndpointResultOutput) Subnet() SubnetResponsePtrOutput {
	return o.ApplyT(func(v LookupPrivateEndpointResult) *SubnetResponse { return v.Subnet }).(SubnetResponsePtrOutput)
}

// Resource tags.
func (o LookupPrivateEndpointResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupPrivateEndpointResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o LookupPrivateEndpointResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPrivateEndpointResultOutput{})
}
