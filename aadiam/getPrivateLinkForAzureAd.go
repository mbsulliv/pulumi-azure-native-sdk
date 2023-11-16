// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package aadiam

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a private link policy with a given name.
// Azure REST API version: 2020-03-01.
func LookupPrivateLinkForAzureAd(ctx *pulumi.Context, args *LookupPrivateLinkForAzureAdArgs, opts ...pulumi.InvokeOption) (*LookupPrivateLinkForAzureAdResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupPrivateLinkForAzureAdResult
	err := ctx.Invoke("azure-native:aadiam:getPrivateLinkForAzureAd", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPrivateLinkForAzureAdArgs struct {
	// The name of the private link policy in Azure AD.
	PolicyName string `pulumi:"policyName"`
	// Name of an Azure resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// PrivateLink Policy configuration object.
type LookupPrivateLinkForAzureAdResult struct {
	// Flag indicating whether all tenants are allowed
	AllTenants *bool `pulumi:"allTenants"`
	// String Id used to locate any resource on Azure.
	Id string `pulumi:"id"`
	// Name of this resource.
	Name *string `pulumi:"name"`
	// Guid of the owner tenant
	OwnerTenantId *string `pulumi:"ownerTenantId"`
	// Name of the resource group
	ResourceGroup *string `pulumi:"resourceGroup"`
	// Name of the private link policy resource
	ResourceName *string `pulumi:"resourceName"`
	// Subscription Identifier
	SubscriptionId *string `pulumi:"subscriptionId"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The list of tenantIds.
	Tenants []string `pulumi:"tenants"`
	// Type of this resource.
	Type string `pulumi:"type"`
}

func LookupPrivateLinkForAzureAdOutput(ctx *pulumi.Context, args LookupPrivateLinkForAzureAdOutputArgs, opts ...pulumi.InvokeOption) LookupPrivateLinkForAzureAdResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPrivateLinkForAzureAdResult, error) {
			args := v.(LookupPrivateLinkForAzureAdArgs)
			r, err := LookupPrivateLinkForAzureAd(ctx, &args, opts...)
			var s LookupPrivateLinkForAzureAdResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPrivateLinkForAzureAdResultOutput)
}

type LookupPrivateLinkForAzureAdOutputArgs struct {
	// The name of the private link policy in Azure AD.
	PolicyName pulumi.StringInput `pulumi:"policyName"`
	// Name of an Azure resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupPrivateLinkForAzureAdOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateLinkForAzureAdArgs)(nil)).Elem()
}

// PrivateLink Policy configuration object.
type LookupPrivateLinkForAzureAdResultOutput struct{ *pulumi.OutputState }

func (LookupPrivateLinkForAzureAdResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateLinkForAzureAdResult)(nil)).Elem()
}

func (o LookupPrivateLinkForAzureAdResultOutput) ToLookupPrivateLinkForAzureAdResultOutput() LookupPrivateLinkForAzureAdResultOutput {
	return o
}

func (o LookupPrivateLinkForAzureAdResultOutput) ToLookupPrivateLinkForAzureAdResultOutputWithContext(ctx context.Context) LookupPrivateLinkForAzureAdResultOutput {
	return o
}

// Flag indicating whether all tenants are allowed
func (o LookupPrivateLinkForAzureAdResultOutput) AllTenants() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupPrivateLinkForAzureAdResult) *bool { return v.AllTenants }).(pulumi.BoolPtrOutput)
}

// String Id used to locate any resource on Azure.
func (o LookupPrivateLinkForAzureAdResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateLinkForAzureAdResult) string { return v.Id }).(pulumi.StringOutput)
}

// Name of this resource.
func (o LookupPrivateLinkForAzureAdResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPrivateLinkForAzureAdResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// Guid of the owner tenant
func (o LookupPrivateLinkForAzureAdResultOutput) OwnerTenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPrivateLinkForAzureAdResult) *string { return v.OwnerTenantId }).(pulumi.StringPtrOutput)
}

// Name of the resource group
func (o LookupPrivateLinkForAzureAdResultOutput) ResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPrivateLinkForAzureAdResult) *string { return v.ResourceGroup }).(pulumi.StringPtrOutput)
}

// Name of the private link policy resource
func (o LookupPrivateLinkForAzureAdResultOutput) ResourceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPrivateLinkForAzureAdResult) *string { return v.ResourceName }).(pulumi.StringPtrOutput)
}

// Subscription Identifier
func (o LookupPrivateLinkForAzureAdResultOutput) SubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPrivateLinkForAzureAdResult) *string { return v.SubscriptionId }).(pulumi.StringPtrOutput)
}

// Resource tags.
func (o LookupPrivateLinkForAzureAdResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupPrivateLinkForAzureAdResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The list of tenantIds.
func (o LookupPrivateLinkForAzureAdResultOutput) Tenants() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupPrivateLinkForAzureAdResult) []string { return v.Tenants }).(pulumi.StringArrayOutput)
}

// Type of this resource.
func (o LookupPrivateLinkForAzureAdResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateLinkForAzureAdResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPrivateLinkForAzureAdResultOutput{})
}
