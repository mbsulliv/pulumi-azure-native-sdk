// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20180101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets the details of the backend specified by its identifier.
func LookupBackend(ctx *pulumi.Context, args *LookupBackendArgs, opts ...pulumi.InvokeOption) (*LookupBackendResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupBackendResult
	err := ctx.Invoke("azure-native:apimanagement/v20180101:getBackend", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupBackendArgs struct {
	// Identifier of the Backend entity. Must be unique in the current API Management service instance.
	Backendid string `pulumi:"backendid"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
}

// Backend details.
type LookupBackendResult struct {
	// Backend Credentials Contract Properties
	Credentials *BackendCredentialsContractResponse `pulumi:"credentials"`
	// Backend Description.
	Description *string `pulumi:"description"`
	// Resource ID.
	Id string `pulumi:"id"`
	// Resource name.
	Name string `pulumi:"name"`
	// Backend Properties contract
	Properties BackendPropertiesResponse `pulumi:"properties"`
	// Backend communication protocol.
	Protocol string `pulumi:"protocol"`
	// Backend Proxy Contract Properties
	Proxy *BackendProxyContractResponse `pulumi:"proxy"`
	// Management Uri of the Resource in External System. This url can be the Arm Resource Id of Logic Apps, Function Apps or Api Apps.
	ResourceId *string `pulumi:"resourceId"`
	// Backend Title.
	Title *string `pulumi:"title"`
	// Backend TLS Properties
	Tls *BackendTlsPropertiesResponse `pulumi:"tls"`
	// Resource type for API Management resource.
	Type string `pulumi:"type"`
	// Runtime Url of the Backend.
	Url string `pulumi:"url"`
}

// Defaults sets the appropriate defaults for LookupBackendResult
func (val *LookupBackendResult) Defaults() *LookupBackendResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Tls = tmp.Tls.Defaults()

	return &tmp
}

func LookupBackendOutput(ctx *pulumi.Context, args LookupBackendOutputArgs, opts ...pulumi.InvokeOption) LookupBackendResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBackendResult, error) {
			args := v.(LookupBackendArgs)
			r, err := LookupBackend(ctx, &args, opts...)
			var s LookupBackendResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBackendResultOutput)
}

type LookupBackendOutputArgs struct {
	// Identifier of the Backend entity. Must be unique in the current API Management service instance.
	Backendid pulumi.StringInput `pulumi:"backendid"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupBackendOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBackendArgs)(nil)).Elem()
}

// Backend details.
type LookupBackendResultOutput struct{ *pulumi.OutputState }

func (LookupBackendResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBackendResult)(nil)).Elem()
}

func (o LookupBackendResultOutput) ToLookupBackendResultOutput() LookupBackendResultOutput {
	return o
}

func (o LookupBackendResultOutput) ToLookupBackendResultOutputWithContext(ctx context.Context) LookupBackendResultOutput {
	return o
}

func (o LookupBackendResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupBackendResult] {
	return pulumix.Output[LookupBackendResult]{
		OutputState: o.OutputState,
	}
}

// Backend Credentials Contract Properties
func (o LookupBackendResultOutput) Credentials() BackendCredentialsContractResponsePtrOutput {
	return o.ApplyT(func(v LookupBackendResult) *BackendCredentialsContractResponse { return v.Credentials }).(BackendCredentialsContractResponsePtrOutput)
}

// Backend Description.
func (o LookupBackendResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBackendResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Resource ID.
func (o LookupBackendResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackendResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource name.
func (o LookupBackendResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackendResult) string { return v.Name }).(pulumi.StringOutput)
}

// Backend Properties contract
func (o LookupBackendResultOutput) Properties() BackendPropertiesResponseOutput {
	return o.ApplyT(func(v LookupBackendResult) BackendPropertiesResponse { return v.Properties }).(BackendPropertiesResponseOutput)
}

// Backend communication protocol.
func (o LookupBackendResultOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackendResult) string { return v.Protocol }).(pulumi.StringOutput)
}

// Backend Proxy Contract Properties
func (o LookupBackendResultOutput) Proxy() BackendProxyContractResponsePtrOutput {
	return o.ApplyT(func(v LookupBackendResult) *BackendProxyContractResponse { return v.Proxy }).(BackendProxyContractResponsePtrOutput)
}

// Management Uri of the Resource in External System. This url can be the Arm Resource Id of Logic Apps, Function Apps or Api Apps.
func (o LookupBackendResultOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBackendResult) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

// Backend Title.
func (o LookupBackendResultOutput) Title() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBackendResult) *string { return v.Title }).(pulumi.StringPtrOutput)
}

// Backend TLS Properties
func (o LookupBackendResultOutput) Tls() BackendTlsPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupBackendResult) *BackendTlsPropertiesResponse { return v.Tls }).(BackendTlsPropertiesResponsePtrOutput)
}

// Resource type for API Management resource.
func (o LookupBackendResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackendResult) string { return v.Type }).(pulumi.StringOutput)
}

// Runtime Url of the Backend.
func (o LookupBackendResultOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBackendResult) string { return v.Url }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBackendResultOutput{})
}
