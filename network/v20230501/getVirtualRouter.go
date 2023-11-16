// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the specified Virtual Router.
func LookupVirtualRouter(ctx *pulumi.Context, args *LookupVirtualRouterArgs, opts ...pulumi.InvokeOption) (*LookupVirtualRouterResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupVirtualRouterResult
	err := ctx.Invoke("azure-native:network/v20230501:getVirtualRouter", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVirtualRouterArgs struct {
	// Expands referenced resources.
	Expand *string `pulumi:"expand"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Virtual Router.
	VirtualRouterName string `pulumi:"virtualRouterName"`
}

// VirtualRouter Resource.
type LookupVirtualRouterResult struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// The Gateway on which VirtualRouter is hosted.
	HostedGateway *SubResourceResponse `pulumi:"hostedGateway"`
	// The Subnet on which VirtualRouter is hosted.
	HostedSubnet *SubResourceResponse `pulumi:"hostedSubnet"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Resource name.
	Name string `pulumi:"name"`
	// List of references to VirtualRouterPeerings.
	Peerings []SubResourceResponse `pulumi:"peerings"`
	// The provisioning state of the resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type string `pulumi:"type"`
	// VirtualRouter ASN.
	VirtualRouterAsn *float64 `pulumi:"virtualRouterAsn"`
	// VirtualRouter IPs.
	VirtualRouterIps []string `pulumi:"virtualRouterIps"`
}

func LookupVirtualRouterOutput(ctx *pulumi.Context, args LookupVirtualRouterOutputArgs, opts ...pulumi.InvokeOption) LookupVirtualRouterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVirtualRouterResult, error) {
			args := v.(LookupVirtualRouterArgs)
			r, err := LookupVirtualRouter(ctx, &args, opts...)
			var s LookupVirtualRouterResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVirtualRouterResultOutput)
}

type LookupVirtualRouterOutputArgs struct {
	// Expands referenced resources.
	Expand pulumi.StringPtrInput `pulumi:"expand"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the Virtual Router.
	VirtualRouterName pulumi.StringInput `pulumi:"virtualRouterName"`
}

func (LookupVirtualRouterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualRouterArgs)(nil)).Elem()
}

// VirtualRouter Resource.
type LookupVirtualRouterResultOutput struct{ *pulumi.OutputState }

func (LookupVirtualRouterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualRouterResult)(nil)).Elem()
}

func (o LookupVirtualRouterResultOutput) ToLookupVirtualRouterResultOutput() LookupVirtualRouterResultOutput {
	return o
}

func (o LookupVirtualRouterResultOutput) ToLookupVirtualRouterResultOutputWithContext(ctx context.Context) LookupVirtualRouterResultOutput {
	return o
}

// A unique read-only string that changes whenever the resource is updated.
func (o LookupVirtualRouterResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualRouterResult) string { return v.Etag }).(pulumi.StringOutput)
}

// The Gateway on which VirtualRouter is hosted.
func (o LookupVirtualRouterResultOutput) HostedGateway() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualRouterResult) *SubResourceResponse { return v.HostedGateway }).(SubResourceResponsePtrOutput)
}

// The Subnet on which VirtualRouter is hosted.
func (o LookupVirtualRouterResultOutput) HostedSubnet() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualRouterResult) *SubResourceResponse { return v.HostedSubnet }).(SubResourceResponsePtrOutput)
}

// Resource ID.
func (o LookupVirtualRouterResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualRouterResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// Resource location.
func (o LookupVirtualRouterResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualRouterResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o LookupVirtualRouterResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualRouterResult) string { return v.Name }).(pulumi.StringOutput)
}

// List of references to VirtualRouterPeerings.
func (o LookupVirtualRouterResultOutput) Peerings() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v LookupVirtualRouterResult) []SubResourceResponse { return v.Peerings }).(SubResourceResponseArrayOutput)
}

// The provisioning state of the resource.
func (o LookupVirtualRouterResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualRouterResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Resource tags.
func (o LookupVirtualRouterResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupVirtualRouterResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o LookupVirtualRouterResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualRouterResult) string { return v.Type }).(pulumi.StringOutput)
}

// VirtualRouter ASN.
func (o LookupVirtualRouterResultOutput) VirtualRouterAsn() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupVirtualRouterResult) *float64 { return v.VirtualRouterAsn }).(pulumi.Float64PtrOutput)
}

// VirtualRouter IPs.
func (o LookupVirtualRouterResultOutput) VirtualRouterIps() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVirtualRouterResult) []string { return v.VirtualRouterIps }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVirtualRouterResultOutput{})
}
