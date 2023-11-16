// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the specified ipGroups.
func LookupIpGroup(ctx *pulumi.Context, args *LookupIpGroupArgs, opts ...pulumi.InvokeOption) (*LookupIpGroupResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupIpGroupResult
	err := ctx.Invoke("azure-native:network/v20230601:getIpGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIpGroupArgs struct {
	// Expands resourceIds (of Firewalls/Network Security Groups etc.) back referenced by the IpGroups resource.
	Expand *string `pulumi:"expand"`
	// The name of the ipGroups.
	IpGroupsName string `pulumi:"ipGroupsName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The IpGroups resource information.
type LookupIpGroupResult struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// List of references to Firewall Policies resources that this IpGroups is associated with.
	FirewallPolicies []SubResourceResponse `pulumi:"firewallPolicies"`
	// List of references to Firewall resources that this IpGroups is associated with.
	Firewalls []SubResourceResponse `pulumi:"firewalls"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// IpAddresses/IpAddressPrefixes in the IpGroups resource.
	IpAddresses []string `pulumi:"ipAddresses"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Resource name.
	Name string `pulumi:"name"`
	// The provisioning state of the IpGroups resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupIpGroupOutput(ctx *pulumi.Context, args LookupIpGroupOutputArgs, opts ...pulumi.InvokeOption) LookupIpGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIpGroupResult, error) {
			args := v.(LookupIpGroupArgs)
			r, err := LookupIpGroup(ctx, &args, opts...)
			var s LookupIpGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIpGroupResultOutput)
}

type LookupIpGroupOutputArgs struct {
	// Expands resourceIds (of Firewalls/Network Security Groups etc.) back referenced by the IpGroups resource.
	Expand pulumi.StringPtrInput `pulumi:"expand"`
	// The name of the ipGroups.
	IpGroupsName pulumi.StringInput `pulumi:"ipGroupsName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupIpGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIpGroupArgs)(nil)).Elem()
}

// The IpGroups resource information.
type LookupIpGroupResultOutput struct{ *pulumi.OutputState }

func (LookupIpGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIpGroupResult)(nil)).Elem()
}

func (o LookupIpGroupResultOutput) ToLookupIpGroupResultOutput() LookupIpGroupResultOutput {
	return o
}

func (o LookupIpGroupResultOutput) ToLookupIpGroupResultOutputWithContext(ctx context.Context) LookupIpGroupResultOutput {
	return o
}

// A unique read-only string that changes whenever the resource is updated.
func (o LookupIpGroupResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpGroupResult) string { return v.Etag }).(pulumi.StringOutput)
}

// List of references to Firewall Policies resources that this IpGroups is associated with.
func (o LookupIpGroupResultOutput) FirewallPolicies() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v LookupIpGroupResult) []SubResourceResponse { return v.FirewallPolicies }).(SubResourceResponseArrayOutput)
}

// List of references to Firewall resources that this IpGroups is associated with.
func (o LookupIpGroupResultOutput) Firewalls() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v LookupIpGroupResult) []SubResourceResponse { return v.Firewalls }).(SubResourceResponseArrayOutput)
}

// Resource ID.
func (o LookupIpGroupResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIpGroupResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// IpAddresses/IpAddressPrefixes in the IpGroups resource.
func (o LookupIpGroupResultOutput) IpAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupIpGroupResult) []string { return v.IpAddresses }).(pulumi.StringArrayOutput)
}

// Resource location.
func (o LookupIpGroupResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIpGroupResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o LookupIpGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the IpGroups resource.
func (o LookupIpGroupResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpGroupResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Resource tags.
func (o LookupIpGroupResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupIpGroupResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o LookupIpGroupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpGroupResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIpGroupResultOutput{})
}
