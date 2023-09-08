// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230701

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the properties of the specified cache rule resource.
func LookupCacheRule(ctx *pulumi.Context, args *LookupCacheRuleArgs, opts ...pulumi.InvokeOption) (*LookupCacheRuleResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupCacheRuleResult
	err := ctx.Invoke("azure-native:containerregistry/v20230701:getCacheRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCacheRuleArgs struct {
	// The name of the cache rule.
	CacheRuleName string `pulumi:"cacheRuleName"`
	// The name of the container registry.
	RegistryName string `pulumi:"registryName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// An object that represents a cache rule for a container registry.
type LookupCacheRuleResult struct {
	// The creation date of the cache rule.
	CreationDate string `pulumi:"creationDate"`
	// The ARM resource ID of the credential store which is associated with the cache rule.
	CredentialSetResourceId *string `pulumi:"credentialSetResourceId"`
	// The resource ID.
	Id string `pulumi:"id"`
	// The name of the resource.
	Name string `pulumi:"name"`
	// Provisioning state of the resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Source repository pulled from upstream.
	SourceRepository *string `pulumi:"sourceRepository"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Target repository specified in docker pull command.
	// Eg: docker pull myregistry.azurecr.io/{targetRepository}:{tag}
	TargetRepository *string `pulumi:"targetRepository"`
	// The type of the resource.
	Type string `pulumi:"type"`
}

func LookupCacheRuleOutput(ctx *pulumi.Context, args LookupCacheRuleOutputArgs, opts ...pulumi.InvokeOption) LookupCacheRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCacheRuleResult, error) {
			args := v.(LookupCacheRuleArgs)
			r, err := LookupCacheRule(ctx, &args, opts...)
			var s LookupCacheRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCacheRuleResultOutput)
}

type LookupCacheRuleOutputArgs struct {
	// The name of the cache rule.
	CacheRuleName pulumi.StringInput `pulumi:"cacheRuleName"`
	// The name of the container registry.
	RegistryName pulumi.StringInput `pulumi:"registryName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupCacheRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCacheRuleArgs)(nil)).Elem()
}

// An object that represents a cache rule for a container registry.
type LookupCacheRuleResultOutput struct{ *pulumi.OutputState }

func (LookupCacheRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCacheRuleResult)(nil)).Elem()
}

func (o LookupCacheRuleResultOutput) ToLookupCacheRuleResultOutput() LookupCacheRuleResultOutput {
	return o
}

func (o LookupCacheRuleResultOutput) ToLookupCacheRuleResultOutputWithContext(ctx context.Context) LookupCacheRuleResultOutput {
	return o
}

// The creation date of the cache rule.
func (o LookupCacheRuleResultOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCacheRuleResult) string { return v.CreationDate }).(pulumi.StringOutput)
}

// The ARM resource ID of the credential store which is associated with the cache rule.
func (o LookupCacheRuleResultOutput) CredentialSetResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCacheRuleResult) *string { return v.CredentialSetResourceId }).(pulumi.StringPtrOutput)
}

// The resource ID.
func (o LookupCacheRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCacheRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource.
func (o LookupCacheRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCacheRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

// Provisioning state of the resource.
func (o LookupCacheRuleResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCacheRuleResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Source repository pulled from upstream.
func (o LookupCacheRuleResultOutput) SourceRepository() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCacheRuleResult) *string { return v.SourceRepository }).(pulumi.StringPtrOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupCacheRuleResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupCacheRuleResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Target repository specified in docker pull command.
// Eg: docker pull myregistry.azurecr.io/{targetRepository}:{tag}
func (o LookupCacheRuleResultOutput) TargetRepository() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCacheRuleResult) *string { return v.TargetRepository }).(pulumi.StringPtrOutput)
}

// The type of the resource.
func (o LookupCacheRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCacheRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCacheRuleResultOutput{})
}