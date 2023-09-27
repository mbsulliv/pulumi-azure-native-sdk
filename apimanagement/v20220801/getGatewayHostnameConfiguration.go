// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get details of a hostname configuration
func LookupGatewayHostnameConfiguration(ctx *pulumi.Context, args *LookupGatewayHostnameConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupGatewayHostnameConfigurationResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupGatewayHostnameConfigurationResult
	err := ctx.Invoke("azure-native:apimanagement/v20220801:getGatewayHostnameConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGatewayHostnameConfigurationArgs struct {
	// Gateway entity identifier. Must be unique in the current API Management service instance. Must not have value 'managed'
	GatewayId string `pulumi:"gatewayId"`
	// Gateway hostname configuration identifier. Must be unique in the scope of parent Gateway entity.
	HcId string `pulumi:"hcId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
}

// Gateway hostname configuration details.
type LookupGatewayHostnameConfigurationResult struct {
	// Identifier of Certificate entity that will be used for TLS connection establishment
	CertificateId *string `pulumi:"certificateId"`
	// Hostname value. Supports valid domain name, partial or full wildcard
	Hostname *string `pulumi:"hostname"`
	// Specifies if HTTP/2.0 is supported
	Http2Enabled *bool `pulumi:"http2Enabled"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Determines whether gateway requests client certificate
	NegotiateClientCertificate *bool `pulumi:"negotiateClientCertificate"`
	// Specifies if TLS 1.0 is supported
	Tls10Enabled *bool `pulumi:"tls10Enabled"`
	// Specifies if TLS 1.1 is supported
	Tls11Enabled *bool `pulumi:"tls11Enabled"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupGatewayHostnameConfigurationOutput(ctx *pulumi.Context, args LookupGatewayHostnameConfigurationOutputArgs, opts ...pulumi.InvokeOption) LookupGatewayHostnameConfigurationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGatewayHostnameConfigurationResult, error) {
			args := v.(LookupGatewayHostnameConfigurationArgs)
			r, err := LookupGatewayHostnameConfiguration(ctx, &args, opts...)
			var s LookupGatewayHostnameConfigurationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGatewayHostnameConfigurationResultOutput)
}

type LookupGatewayHostnameConfigurationOutputArgs struct {
	// Gateway entity identifier. Must be unique in the current API Management service instance. Must not have value 'managed'
	GatewayId pulumi.StringInput `pulumi:"gatewayId"`
	// Gateway hostname configuration identifier. Must be unique in the scope of parent Gateway entity.
	HcId pulumi.StringInput `pulumi:"hcId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupGatewayHostnameConfigurationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGatewayHostnameConfigurationArgs)(nil)).Elem()
}

// Gateway hostname configuration details.
type LookupGatewayHostnameConfigurationResultOutput struct{ *pulumi.OutputState }

func (LookupGatewayHostnameConfigurationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGatewayHostnameConfigurationResult)(nil)).Elem()
}

func (o LookupGatewayHostnameConfigurationResultOutput) ToLookupGatewayHostnameConfigurationResultOutput() LookupGatewayHostnameConfigurationResultOutput {
	return o
}

func (o LookupGatewayHostnameConfigurationResultOutput) ToLookupGatewayHostnameConfigurationResultOutputWithContext(ctx context.Context) LookupGatewayHostnameConfigurationResultOutput {
	return o
}

func (o LookupGatewayHostnameConfigurationResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupGatewayHostnameConfigurationResult] {
	return pulumix.Output[LookupGatewayHostnameConfigurationResult]{
		OutputState: o.OutputState,
	}
}

// Identifier of Certificate entity that will be used for TLS connection establishment
func (o LookupGatewayHostnameConfigurationResultOutput) CertificateId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGatewayHostnameConfigurationResult) *string { return v.CertificateId }).(pulumi.StringPtrOutput)
}

// Hostname value. Supports valid domain name, partial or full wildcard
func (o LookupGatewayHostnameConfigurationResultOutput) Hostname() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGatewayHostnameConfigurationResult) *string { return v.Hostname }).(pulumi.StringPtrOutput)
}

// Specifies if HTTP/2.0 is supported
func (o LookupGatewayHostnameConfigurationResultOutput) Http2Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupGatewayHostnameConfigurationResult) *bool { return v.Http2Enabled }).(pulumi.BoolPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupGatewayHostnameConfigurationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayHostnameConfigurationResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupGatewayHostnameConfigurationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayHostnameConfigurationResult) string { return v.Name }).(pulumi.StringOutput)
}

// Determines whether gateway requests client certificate
func (o LookupGatewayHostnameConfigurationResultOutput) NegotiateClientCertificate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupGatewayHostnameConfigurationResult) *bool { return v.NegotiateClientCertificate }).(pulumi.BoolPtrOutput)
}

// Specifies if TLS 1.0 is supported
func (o LookupGatewayHostnameConfigurationResultOutput) Tls10Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupGatewayHostnameConfigurationResult) *bool { return v.Tls10Enabled }).(pulumi.BoolPtrOutput)
}

// Specifies if TLS 1.1 is supported
func (o LookupGatewayHostnameConfigurationResultOutput) Tls11Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupGatewayHostnameConfigurationResult) *bool { return v.Tls11Enabled }).(pulumi.BoolPtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupGatewayHostnameConfigurationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayHostnameConfigurationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGatewayHostnameConfigurationResultOutput{})
}
