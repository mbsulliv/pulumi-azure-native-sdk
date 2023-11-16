// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the specified Security Partner Provider.
func LookupSecurityPartnerProvider(ctx *pulumi.Context, args *LookupSecurityPartnerProviderArgs, opts ...pulumi.InvokeOption) (*LookupSecurityPartnerProviderResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupSecurityPartnerProviderResult
	err := ctx.Invoke("azure-native:network/v20230601:getSecurityPartnerProvider", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSecurityPartnerProviderArgs struct {
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Security Partner Provider.
	SecurityPartnerProviderName string `pulumi:"securityPartnerProviderName"`
}

// Security Partner Provider resource.
type LookupSecurityPartnerProviderResult struct {
	// The connection status with the Security Partner Provider.
	ConnectionStatus string `pulumi:"connectionStatus"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Resource name.
	Name string `pulumi:"name"`
	// The provisioning state of the Security Partner Provider resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// The security provider name.
	SecurityProviderName *string `pulumi:"securityProviderName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type string `pulumi:"type"`
	// The virtualHub to which the Security Partner Provider belongs.
	VirtualHub *SubResourceResponse `pulumi:"virtualHub"`
}

func LookupSecurityPartnerProviderOutput(ctx *pulumi.Context, args LookupSecurityPartnerProviderOutputArgs, opts ...pulumi.InvokeOption) LookupSecurityPartnerProviderResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSecurityPartnerProviderResult, error) {
			args := v.(LookupSecurityPartnerProviderArgs)
			r, err := LookupSecurityPartnerProvider(ctx, &args, opts...)
			var s LookupSecurityPartnerProviderResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSecurityPartnerProviderResultOutput)
}

type LookupSecurityPartnerProviderOutputArgs struct {
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the Security Partner Provider.
	SecurityPartnerProviderName pulumi.StringInput `pulumi:"securityPartnerProviderName"`
}

func (LookupSecurityPartnerProviderOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSecurityPartnerProviderArgs)(nil)).Elem()
}

// Security Partner Provider resource.
type LookupSecurityPartnerProviderResultOutput struct{ *pulumi.OutputState }

func (LookupSecurityPartnerProviderResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSecurityPartnerProviderResult)(nil)).Elem()
}

func (o LookupSecurityPartnerProviderResultOutput) ToLookupSecurityPartnerProviderResultOutput() LookupSecurityPartnerProviderResultOutput {
	return o
}

func (o LookupSecurityPartnerProviderResultOutput) ToLookupSecurityPartnerProviderResultOutputWithContext(ctx context.Context) LookupSecurityPartnerProviderResultOutput {
	return o
}

// The connection status with the Security Partner Provider.
func (o LookupSecurityPartnerProviderResultOutput) ConnectionStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecurityPartnerProviderResult) string { return v.ConnectionStatus }).(pulumi.StringOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o LookupSecurityPartnerProviderResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecurityPartnerProviderResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Resource ID.
func (o LookupSecurityPartnerProviderResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSecurityPartnerProviderResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// Resource location.
func (o LookupSecurityPartnerProviderResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSecurityPartnerProviderResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o LookupSecurityPartnerProviderResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecurityPartnerProviderResult) string { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the Security Partner Provider resource.
func (o LookupSecurityPartnerProviderResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecurityPartnerProviderResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The security provider name.
func (o LookupSecurityPartnerProviderResultOutput) SecurityProviderName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSecurityPartnerProviderResult) *string { return v.SecurityProviderName }).(pulumi.StringPtrOutput)
}

// Resource tags.
func (o LookupSecurityPartnerProviderResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSecurityPartnerProviderResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o LookupSecurityPartnerProviderResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecurityPartnerProviderResult) string { return v.Type }).(pulumi.StringOutput)
}

// The virtualHub to which the Security Partner Provider belongs.
func (o LookupSecurityPartnerProviderResultOutput) VirtualHub() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupSecurityPartnerProviderResult) *SubResourceResponse { return v.VirtualHub }).(SubResourceResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSecurityPartnerProviderResultOutput{})
}
