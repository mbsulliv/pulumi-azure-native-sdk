// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230315

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the endpoint access credentials to the resource.
func ListEndpointCredentials(ctx *pulumi.Context, args *ListEndpointCredentialsArgs, opts ...pulumi.InvokeOption) (*ListEndpointCredentialsResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListEndpointCredentialsResult
	err := ctx.Invoke("azure-native:hybridconnectivity/v20230315:listEndpointCredentials", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListEndpointCredentialsArgs struct {
	// The endpoint name.
	EndpointName string `pulumi:"endpointName"`
	// The is how long the endpoint access token is valid (in seconds).
	Expiresin *int `pulumi:"expiresin"`
	// The fully qualified Azure Resource manager identifier of the resource to be connected.
	ResourceUri string `pulumi:"resourceUri"`
	// The name of the service. If not provided, the request will by pass the generation of service configuration token
	ServiceName *string `pulumi:"serviceName"`
}

// The endpoint access for the target resource.
type ListEndpointCredentialsResult struct {
	// Access key for hybrid connection.
	AccessKey string `pulumi:"accessKey"`
	// The expiration of access key in unix time.
	ExpiresOn *float64 `pulumi:"expiresOn"`
	// Azure Relay hybrid connection name for the resource.
	HybridConnectionName string `pulumi:"hybridConnectionName"`
	// The namespace name.
	NamespaceName string `pulumi:"namespaceName"`
	// The suffix domain name of relay namespace.
	NamespaceNameSuffix string `pulumi:"namespaceNameSuffix"`
	// The token to access the enabled service.
	ServiceConfigurationToken *string `pulumi:"serviceConfigurationToken"`
}

func ListEndpointCredentialsOutput(ctx *pulumi.Context, args ListEndpointCredentialsOutputArgs, opts ...pulumi.InvokeOption) ListEndpointCredentialsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListEndpointCredentialsResult, error) {
			args := v.(ListEndpointCredentialsArgs)
			r, err := ListEndpointCredentials(ctx, &args, opts...)
			var s ListEndpointCredentialsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListEndpointCredentialsResultOutput)
}

type ListEndpointCredentialsOutputArgs struct {
	// The endpoint name.
	EndpointName pulumi.StringInput `pulumi:"endpointName"`
	// The is how long the endpoint access token is valid (in seconds).
	Expiresin pulumi.IntPtrInput `pulumi:"expiresin"`
	// The fully qualified Azure Resource manager identifier of the resource to be connected.
	ResourceUri pulumi.StringInput `pulumi:"resourceUri"`
	// The name of the service. If not provided, the request will by pass the generation of service configuration token
	ServiceName pulumi.StringPtrInput `pulumi:"serviceName"`
}

func (ListEndpointCredentialsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListEndpointCredentialsArgs)(nil)).Elem()
}

// The endpoint access for the target resource.
type ListEndpointCredentialsResultOutput struct{ *pulumi.OutputState }

func (ListEndpointCredentialsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListEndpointCredentialsResult)(nil)).Elem()
}

func (o ListEndpointCredentialsResultOutput) ToListEndpointCredentialsResultOutput() ListEndpointCredentialsResultOutput {
	return o
}

func (o ListEndpointCredentialsResultOutput) ToListEndpointCredentialsResultOutputWithContext(ctx context.Context) ListEndpointCredentialsResultOutput {
	return o
}

// Access key for hybrid connection.
func (o ListEndpointCredentialsResultOutput) AccessKey() pulumi.StringOutput {
	return o.ApplyT(func(v ListEndpointCredentialsResult) string { return v.AccessKey }).(pulumi.StringOutput)
}

// The expiration of access key in unix time.
func (o ListEndpointCredentialsResultOutput) ExpiresOn() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ListEndpointCredentialsResult) *float64 { return v.ExpiresOn }).(pulumi.Float64PtrOutput)
}

// Azure Relay hybrid connection name for the resource.
func (o ListEndpointCredentialsResultOutput) HybridConnectionName() pulumi.StringOutput {
	return o.ApplyT(func(v ListEndpointCredentialsResult) string { return v.HybridConnectionName }).(pulumi.StringOutput)
}

// The namespace name.
func (o ListEndpointCredentialsResultOutput) NamespaceName() pulumi.StringOutput {
	return o.ApplyT(func(v ListEndpointCredentialsResult) string { return v.NamespaceName }).(pulumi.StringOutput)
}

// The suffix domain name of relay namespace.
func (o ListEndpointCredentialsResultOutput) NamespaceNameSuffix() pulumi.StringOutput {
	return o.ApplyT(func(v ListEndpointCredentialsResult) string { return v.NamespaceNameSuffix }).(pulumi.StringOutput)
}

// The token to access the enabled service.
func (o ListEndpointCredentialsResultOutput) ServiceConfigurationToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListEndpointCredentialsResult) *string { return v.ServiceConfigurationToken }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListEndpointCredentialsResultOutput{})
}