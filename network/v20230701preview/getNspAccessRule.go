// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230701preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets the specified NSP access rule by name.
func LookupNspAccessRule(ctx *pulumi.Context, args *LookupNspAccessRuleArgs, opts ...pulumi.InvokeOption) (*LookupNspAccessRuleResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupNspAccessRuleResult
	err := ctx.Invoke("azure-native:network/v20230701preview:getNspAccessRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNspAccessRuleArgs struct {
	// The name of the NSP access rule.
	AccessRuleName string `pulumi:"accessRuleName"`
	// The name of the network security perimeter.
	NetworkSecurityPerimeterName string `pulumi:"networkSecurityPerimeterName"`
	// The name of the NSP profile.
	ProfileName string `pulumi:"profileName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The NSP access rule resource
type LookupNspAccessRuleResult struct {
	// Inbound address prefixes (IPv4/IPv6)
	AddressPrefixes []string `pulumi:"addressPrefixes"`
	// Direction that specifies whether the access rules is inbound/outbound.
	Direction *string `pulumi:"direction"`
	// Outbound rules email address format.
	EmailAddresses []string `pulumi:"emailAddresses"`
	// Outbound rules fully qualified domain name format.
	FullyQualifiedDomainNames []string `pulumi:"fullyQualifiedDomainNames"`
	// Resource ID.
	Id string `pulumi:"id"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Resource name.
	Name string `pulumi:"name"`
	// Rule specified by the perimeter id.
	NetworkSecurityPerimeters []PerimeterBasedAccessRuleResponse `pulumi:"networkSecurityPerimeters"`
	// Outbound rules phone number format.
	PhoneNumbers []string `pulumi:"phoneNumbers"`
	// The provisioning state of the scope assignment resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// List of subscription ids
	Subscriptions []SubscriptionIdResponse `pulumi:"subscriptions"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupNspAccessRuleOutput(ctx *pulumi.Context, args LookupNspAccessRuleOutputArgs, opts ...pulumi.InvokeOption) LookupNspAccessRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNspAccessRuleResult, error) {
			args := v.(LookupNspAccessRuleArgs)
			r, err := LookupNspAccessRule(ctx, &args, opts...)
			var s LookupNspAccessRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNspAccessRuleResultOutput)
}

type LookupNspAccessRuleOutputArgs struct {
	// The name of the NSP access rule.
	AccessRuleName pulumi.StringInput `pulumi:"accessRuleName"`
	// The name of the network security perimeter.
	NetworkSecurityPerimeterName pulumi.StringInput `pulumi:"networkSecurityPerimeterName"`
	// The name of the NSP profile.
	ProfileName pulumi.StringInput `pulumi:"profileName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupNspAccessRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNspAccessRuleArgs)(nil)).Elem()
}

// The NSP access rule resource
type LookupNspAccessRuleResultOutput struct{ *pulumi.OutputState }

func (LookupNspAccessRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNspAccessRuleResult)(nil)).Elem()
}

func (o LookupNspAccessRuleResultOutput) ToLookupNspAccessRuleResultOutput() LookupNspAccessRuleResultOutput {
	return o
}

func (o LookupNspAccessRuleResultOutput) ToLookupNspAccessRuleResultOutputWithContext(ctx context.Context) LookupNspAccessRuleResultOutput {
	return o
}

func (o LookupNspAccessRuleResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupNspAccessRuleResult] {
	return pulumix.Output[LookupNspAccessRuleResult]{
		OutputState: o.OutputState,
	}
}

// Inbound address prefixes (IPv4/IPv6)
func (o LookupNspAccessRuleResultOutput) AddressPrefixes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupNspAccessRuleResult) []string { return v.AddressPrefixes }).(pulumi.StringArrayOutput)
}

// Direction that specifies whether the access rules is inbound/outbound.
func (o LookupNspAccessRuleResultOutput) Direction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNspAccessRuleResult) *string { return v.Direction }).(pulumi.StringPtrOutput)
}

// Outbound rules email address format.
func (o LookupNspAccessRuleResultOutput) EmailAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupNspAccessRuleResult) []string { return v.EmailAddresses }).(pulumi.StringArrayOutput)
}

// Outbound rules fully qualified domain name format.
func (o LookupNspAccessRuleResultOutput) FullyQualifiedDomainNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupNspAccessRuleResult) []string { return v.FullyQualifiedDomainNames }).(pulumi.StringArrayOutput)
}

// Resource ID.
func (o LookupNspAccessRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNspAccessRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource location.
func (o LookupNspAccessRuleResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNspAccessRuleResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o LookupNspAccessRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNspAccessRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

// Rule specified by the perimeter id.
func (o LookupNspAccessRuleResultOutput) NetworkSecurityPerimeters() PerimeterBasedAccessRuleResponseArrayOutput {
	return o.ApplyT(func(v LookupNspAccessRuleResult) []PerimeterBasedAccessRuleResponse {
		return v.NetworkSecurityPerimeters
	}).(PerimeterBasedAccessRuleResponseArrayOutput)
}

// Outbound rules phone number format.
func (o LookupNspAccessRuleResultOutput) PhoneNumbers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupNspAccessRuleResult) []string { return v.PhoneNumbers }).(pulumi.StringArrayOutput)
}

// The provisioning state of the scope assignment resource.
func (o LookupNspAccessRuleResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNspAccessRuleResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// List of subscription ids
func (o LookupNspAccessRuleResultOutput) Subscriptions() SubscriptionIdResponseArrayOutput {
	return o.ApplyT(func(v LookupNspAccessRuleResult) []SubscriptionIdResponse { return v.Subscriptions }).(SubscriptionIdResponseArrayOutput)
}

// Resource tags.
func (o LookupNspAccessRuleResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupNspAccessRuleResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o LookupNspAccessRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNspAccessRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNspAccessRuleResultOutput{})
}
