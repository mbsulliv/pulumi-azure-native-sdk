// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the specified IpAllocation by resource group.
func LookupIpAllocation(ctx *pulumi.Context, args *LookupIpAllocationArgs, opts ...pulumi.InvokeOption) (*LookupIpAllocationResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupIpAllocationResult
	err := ctx.Invoke("azure-native:network/v20230201:getIpAllocation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupIpAllocationArgs struct {
	// Expands referenced resources.
	Expand *string `pulumi:"expand"`
	// The name of the IpAllocation.
	IpAllocationName string `pulumi:"ipAllocationName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// IpAllocation resource.
type LookupIpAllocationResult struct {
	// IpAllocation tags.
	AllocationTags map[string]string `pulumi:"allocationTags"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// The IPAM allocation ID.
	IpamAllocationId *string `pulumi:"ipamAllocationId"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Resource name.
	Name string `pulumi:"name"`
	// The address prefix for the IpAllocation.
	Prefix *string `pulumi:"prefix"`
	// The address prefix length for the IpAllocation.
	PrefixLength *int `pulumi:"prefixLength"`
	// The address prefix Type for the IpAllocation.
	PrefixType *string `pulumi:"prefixType"`
	// The Subnet that using the prefix of this IpAllocation resource.
	Subnet SubResourceResponse `pulumi:"subnet"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type string `pulumi:"type"`
	// The VirtualNetwork that using the prefix of this IpAllocation resource.
	VirtualNetwork SubResourceResponse `pulumi:"virtualNetwork"`
}

// Defaults sets the appropriate defaults for LookupIpAllocationResult
func (val *LookupIpAllocationResult) Defaults() *LookupIpAllocationResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.PrefixLength == nil {
		prefixLength_ := 0
		tmp.PrefixLength = &prefixLength_
	}
	return &tmp
}

func LookupIpAllocationOutput(ctx *pulumi.Context, args LookupIpAllocationOutputArgs, opts ...pulumi.InvokeOption) LookupIpAllocationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIpAllocationResult, error) {
			args := v.(LookupIpAllocationArgs)
			r, err := LookupIpAllocation(ctx, &args, opts...)
			var s LookupIpAllocationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIpAllocationResultOutput)
}

type LookupIpAllocationOutputArgs struct {
	// Expands referenced resources.
	Expand pulumi.StringPtrInput `pulumi:"expand"`
	// The name of the IpAllocation.
	IpAllocationName pulumi.StringInput `pulumi:"ipAllocationName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupIpAllocationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIpAllocationArgs)(nil)).Elem()
}

// IpAllocation resource.
type LookupIpAllocationResultOutput struct{ *pulumi.OutputState }

func (LookupIpAllocationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIpAllocationResult)(nil)).Elem()
}

func (o LookupIpAllocationResultOutput) ToLookupIpAllocationResultOutput() LookupIpAllocationResultOutput {
	return o
}

func (o LookupIpAllocationResultOutput) ToLookupIpAllocationResultOutputWithContext(ctx context.Context) LookupIpAllocationResultOutput {
	return o
}

// IpAllocation tags.
func (o LookupIpAllocationResultOutput) AllocationTags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupIpAllocationResult) map[string]string { return v.AllocationTags }).(pulumi.StringMapOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o LookupIpAllocationResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpAllocationResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Resource ID.
func (o LookupIpAllocationResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIpAllocationResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// The IPAM allocation ID.
func (o LookupIpAllocationResultOutput) IpamAllocationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIpAllocationResult) *string { return v.IpamAllocationId }).(pulumi.StringPtrOutput)
}

// Resource location.
func (o LookupIpAllocationResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIpAllocationResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o LookupIpAllocationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpAllocationResult) string { return v.Name }).(pulumi.StringOutput)
}

// The address prefix for the IpAllocation.
func (o LookupIpAllocationResultOutput) Prefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIpAllocationResult) *string { return v.Prefix }).(pulumi.StringPtrOutput)
}

// The address prefix length for the IpAllocation.
func (o LookupIpAllocationResultOutput) PrefixLength() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupIpAllocationResult) *int { return v.PrefixLength }).(pulumi.IntPtrOutput)
}

// The address prefix Type for the IpAllocation.
func (o LookupIpAllocationResultOutput) PrefixType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIpAllocationResult) *string { return v.PrefixType }).(pulumi.StringPtrOutput)
}

// The Subnet that using the prefix of this IpAllocation resource.
func (o LookupIpAllocationResultOutput) Subnet() SubResourceResponseOutput {
	return o.ApplyT(func(v LookupIpAllocationResult) SubResourceResponse { return v.Subnet }).(SubResourceResponseOutput)
}

// Resource tags.
func (o LookupIpAllocationResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupIpAllocationResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o LookupIpAllocationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpAllocationResult) string { return v.Type }).(pulumi.StringOutput)
}

// The VirtualNetwork that using the prefix of this IpAllocation resource.
func (o LookupIpAllocationResultOutput) VirtualNetwork() SubResourceResponseOutput {
	return o.ApplyT(func(v LookupIpAllocationResult) SubResourceResponse { return v.VirtualNetwork }).(SubResourceResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIpAllocationResultOutput{})
}
