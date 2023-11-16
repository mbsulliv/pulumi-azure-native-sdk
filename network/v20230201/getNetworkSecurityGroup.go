// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the specified network security group.
func LookupNetworkSecurityGroup(ctx *pulumi.Context, args *LookupNetworkSecurityGroupArgs, opts ...pulumi.InvokeOption) (*LookupNetworkSecurityGroupResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupNetworkSecurityGroupResult
	err := ctx.Invoke("azure-native:network/v20230201:getNetworkSecurityGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNetworkSecurityGroupArgs struct {
	// Expands referenced resources.
	Expand *string `pulumi:"expand"`
	// The name of the network security group.
	NetworkSecurityGroupName string `pulumi:"networkSecurityGroupName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// NetworkSecurityGroup resource.
type LookupNetworkSecurityGroupResult struct {
	// The default security rules of network security group.
	DefaultSecurityRules []SecurityRuleResponse `pulumi:"defaultSecurityRules"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// A collection of references to flow log resources.
	FlowLogs []FlowLogResponse `pulumi:"flowLogs"`
	// When enabled, flows created from Network Security Group connections will be re-evaluated when rules are updates. Initial enablement will trigger re-evaluation.
	FlushConnection *bool `pulumi:"flushConnection"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Resource name.
	Name string `pulumi:"name"`
	// A collection of references to network interfaces.
	NetworkInterfaces []NetworkInterfaceResponse `pulumi:"networkInterfaces"`
	// The provisioning state of the network security group resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// The resource GUID property of the network security group resource.
	ResourceGuid string `pulumi:"resourceGuid"`
	// A collection of security rules of the network security group.
	SecurityRules []SecurityRuleResponse `pulumi:"securityRules"`
	// A collection of references to subnets.
	Subnets []SubnetResponse `pulumi:"subnets"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupNetworkSecurityGroupOutput(ctx *pulumi.Context, args LookupNetworkSecurityGroupOutputArgs, opts ...pulumi.InvokeOption) LookupNetworkSecurityGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNetworkSecurityGroupResult, error) {
			args := v.(LookupNetworkSecurityGroupArgs)
			r, err := LookupNetworkSecurityGroup(ctx, &args, opts...)
			var s LookupNetworkSecurityGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNetworkSecurityGroupResultOutput)
}

type LookupNetworkSecurityGroupOutputArgs struct {
	// Expands referenced resources.
	Expand pulumi.StringPtrInput `pulumi:"expand"`
	// The name of the network security group.
	NetworkSecurityGroupName pulumi.StringInput `pulumi:"networkSecurityGroupName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupNetworkSecurityGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkSecurityGroupArgs)(nil)).Elem()
}

// NetworkSecurityGroup resource.
type LookupNetworkSecurityGroupResultOutput struct{ *pulumi.OutputState }

func (LookupNetworkSecurityGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkSecurityGroupResult)(nil)).Elem()
}

func (o LookupNetworkSecurityGroupResultOutput) ToLookupNetworkSecurityGroupResultOutput() LookupNetworkSecurityGroupResultOutput {
	return o
}

func (o LookupNetworkSecurityGroupResultOutput) ToLookupNetworkSecurityGroupResultOutputWithContext(ctx context.Context) LookupNetworkSecurityGroupResultOutput {
	return o
}

// The default security rules of network security group.
func (o LookupNetworkSecurityGroupResultOutput) DefaultSecurityRules() SecurityRuleResponseArrayOutput {
	return o.ApplyT(func(v LookupNetworkSecurityGroupResult) []SecurityRuleResponse { return v.DefaultSecurityRules }).(SecurityRuleResponseArrayOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o LookupNetworkSecurityGroupResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkSecurityGroupResult) string { return v.Etag }).(pulumi.StringOutput)
}

// A collection of references to flow log resources.
func (o LookupNetworkSecurityGroupResultOutput) FlowLogs() FlowLogResponseArrayOutput {
	return o.ApplyT(func(v LookupNetworkSecurityGroupResult) []FlowLogResponse { return v.FlowLogs }).(FlowLogResponseArrayOutput)
}

// When enabled, flows created from Network Security Group connections will be re-evaluated when rules are updates. Initial enablement will trigger re-evaluation.
func (o LookupNetworkSecurityGroupResultOutput) FlushConnection() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupNetworkSecurityGroupResult) *bool { return v.FlushConnection }).(pulumi.BoolPtrOutput)
}

// Resource ID.
func (o LookupNetworkSecurityGroupResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkSecurityGroupResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// Resource location.
func (o LookupNetworkSecurityGroupResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkSecurityGroupResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o LookupNetworkSecurityGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkSecurityGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

// A collection of references to network interfaces.
func (o LookupNetworkSecurityGroupResultOutput) NetworkInterfaces() NetworkInterfaceResponseArrayOutput {
	return o.ApplyT(func(v LookupNetworkSecurityGroupResult) []NetworkInterfaceResponse { return v.NetworkInterfaces }).(NetworkInterfaceResponseArrayOutput)
}

// The provisioning state of the network security group resource.
func (o LookupNetworkSecurityGroupResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkSecurityGroupResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The resource GUID property of the network security group resource.
func (o LookupNetworkSecurityGroupResultOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkSecurityGroupResult) string { return v.ResourceGuid }).(pulumi.StringOutput)
}

// A collection of security rules of the network security group.
func (o LookupNetworkSecurityGroupResultOutput) SecurityRules() SecurityRuleResponseArrayOutput {
	return o.ApplyT(func(v LookupNetworkSecurityGroupResult) []SecurityRuleResponse { return v.SecurityRules }).(SecurityRuleResponseArrayOutput)
}

// A collection of references to subnets.
func (o LookupNetworkSecurityGroupResultOutput) Subnets() SubnetResponseArrayOutput {
	return o.ApplyT(func(v LookupNetworkSecurityGroupResult) []SubnetResponse { return v.Subnets }).(SubnetResponseArrayOutput)
}

// Resource tags.
func (o LookupNetworkSecurityGroupResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupNetworkSecurityGroupResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o LookupNetworkSecurityGroupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkSecurityGroupResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNetworkSecurityGroupResultOutput{})
}
