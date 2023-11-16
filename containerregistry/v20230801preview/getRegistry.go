// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230801preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the properties of the specified container registry.
func LookupRegistry(ctx *pulumi.Context, args *LookupRegistryArgs, opts ...pulumi.InvokeOption) (*LookupRegistryResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupRegistryResult
	err := ctx.Invoke("azure-native:containerregistry/v20230801preview:getRegistry", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupRegistryArgs struct {
	// The name of the container registry.
	RegistryName string `pulumi:"registryName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// An object that represents a container registry.
type LookupRegistryResult struct {
	// The value that indicates whether the admin user is enabled.
	AdminUserEnabled *bool `pulumi:"adminUserEnabled"`
	// Enables registry-wide pull from unauthenticated clients.
	AnonymousPullEnabled *bool `pulumi:"anonymousPullEnabled"`
	// The creation date of the container registry in ISO8601 format.
	CreationDate string `pulumi:"creationDate"`
	// Enable a single data endpoint per region for serving data.
	DataEndpointEnabled *bool `pulumi:"dataEndpointEnabled"`
	// List of host names that will serve data when dataEndpointEnabled is true.
	DataEndpointHostNames []string `pulumi:"dataEndpointHostNames"`
	// The encryption settings of container registry.
	Encryption *EncryptionPropertyResponse `pulumi:"encryption"`
	// The resource ID.
	Id string `pulumi:"id"`
	// The identity of the container registry.
	Identity *IdentityPropertiesResponse `pulumi:"identity"`
	// The location of the resource. This cannot be changed after the resource is created.
	Location string `pulumi:"location"`
	// The URL that can be used to log into the container registry.
	LoginServer string `pulumi:"loginServer"`
	// The name of the resource.
	Name string `pulumi:"name"`
	// Whether to allow trusted Azure services to access a network restricted registry.
	NetworkRuleBypassOptions *string `pulumi:"networkRuleBypassOptions"`
	// The network rule set for a container registry.
	NetworkRuleSet *NetworkRuleSetResponse `pulumi:"networkRuleSet"`
	// The policies for a container registry.
	Policies *PoliciesResponse `pulumi:"policies"`
	// List of private endpoint connections for a container registry.
	PrivateEndpointConnections []PrivateEndpointConnectionResponse `pulumi:"privateEndpointConnections"`
	// The provisioning state of the container registry at the time the operation was called.
	ProvisioningState string `pulumi:"provisioningState"`
	// Whether or not public network access is allowed for the container registry.
	PublicNetworkAccess *string `pulumi:"publicNetworkAccess"`
	// The SKU of the container registry.
	Sku SkuResponse `pulumi:"sku"`
	// The status of the container registry at the time the operation was called.
	Status StatusResponse `pulumi:"status"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The tags of the resource.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource.
	Type string `pulumi:"type"`
	// Whether or not zone redundancy is enabled for this container registry
	ZoneRedundancy *string `pulumi:"zoneRedundancy"`
}

// Defaults sets the appropriate defaults for LookupRegistryResult
func (val *LookupRegistryResult) Defaults() *LookupRegistryResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.AdminUserEnabled == nil {
		adminUserEnabled_ := false
		tmp.AdminUserEnabled = &adminUserEnabled_
	}
	if tmp.AnonymousPullEnabled == nil {
		anonymousPullEnabled_ := false
		tmp.AnonymousPullEnabled = &anonymousPullEnabled_
	}
	if tmp.NetworkRuleBypassOptions == nil {
		networkRuleBypassOptions_ := "AzureServices"
		tmp.NetworkRuleBypassOptions = &networkRuleBypassOptions_
	}
	tmp.NetworkRuleSet = tmp.NetworkRuleSet.Defaults()

	tmp.Policies = tmp.Policies.Defaults()

	if tmp.PublicNetworkAccess == nil {
		publicNetworkAccess_ := "Enabled"
		tmp.PublicNetworkAccess = &publicNetworkAccess_
	}
	if tmp.ZoneRedundancy == nil {
		zoneRedundancy_ := "Disabled"
		tmp.ZoneRedundancy = &zoneRedundancy_
	}
	return &tmp
}

func LookupRegistryOutput(ctx *pulumi.Context, args LookupRegistryOutputArgs, opts ...pulumi.InvokeOption) LookupRegistryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRegistryResult, error) {
			args := v.(LookupRegistryArgs)
			r, err := LookupRegistry(ctx, &args, opts...)
			var s LookupRegistryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRegistryResultOutput)
}

type LookupRegistryOutputArgs struct {
	// The name of the container registry.
	RegistryName pulumi.StringInput `pulumi:"registryName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupRegistryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRegistryArgs)(nil)).Elem()
}

// An object that represents a container registry.
type LookupRegistryResultOutput struct{ *pulumi.OutputState }

func (LookupRegistryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRegistryResult)(nil)).Elem()
}

func (o LookupRegistryResultOutput) ToLookupRegistryResultOutput() LookupRegistryResultOutput {
	return o
}

func (o LookupRegistryResultOutput) ToLookupRegistryResultOutputWithContext(ctx context.Context) LookupRegistryResultOutput {
	return o
}

// The value that indicates whether the admin user is enabled.
func (o LookupRegistryResultOutput) AdminUserEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRegistryResult) *bool { return v.AdminUserEnabled }).(pulumi.BoolPtrOutput)
}

// Enables registry-wide pull from unauthenticated clients.
func (o LookupRegistryResultOutput) AnonymousPullEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRegistryResult) *bool { return v.AnonymousPullEnabled }).(pulumi.BoolPtrOutput)
}

// The creation date of the container registry in ISO8601 format.
func (o LookupRegistryResultOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryResult) string { return v.CreationDate }).(pulumi.StringOutput)
}

// Enable a single data endpoint per region for serving data.
func (o LookupRegistryResultOutput) DataEndpointEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRegistryResult) *bool { return v.DataEndpointEnabled }).(pulumi.BoolPtrOutput)
}

// List of host names that will serve data when dataEndpointEnabled is true.
func (o LookupRegistryResultOutput) DataEndpointHostNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupRegistryResult) []string { return v.DataEndpointHostNames }).(pulumi.StringArrayOutput)
}

// The encryption settings of container registry.
func (o LookupRegistryResultOutput) Encryption() EncryptionPropertyResponsePtrOutput {
	return o.ApplyT(func(v LookupRegistryResult) *EncryptionPropertyResponse { return v.Encryption }).(EncryptionPropertyResponsePtrOutput)
}

// The resource ID.
func (o LookupRegistryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryResult) string { return v.Id }).(pulumi.StringOutput)
}

// The identity of the container registry.
func (o LookupRegistryResultOutput) Identity() IdentityPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupRegistryResult) *IdentityPropertiesResponse { return v.Identity }).(IdentityPropertiesResponsePtrOutput)
}

// The location of the resource. This cannot be changed after the resource is created.
func (o LookupRegistryResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryResult) string { return v.Location }).(pulumi.StringOutput)
}

// The URL that can be used to log into the container registry.
func (o LookupRegistryResultOutput) LoginServer() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryResult) string { return v.LoginServer }).(pulumi.StringOutput)
}

// The name of the resource.
func (o LookupRegistryResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryResult) string { return v.Name }).(pulumi.StringOutput)
}

// Whether to allow trusted Azure services to access a network restricted registry.
func (o LookupRegistryResultOutput) NetworkRuleBypassOptions() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRegistryResult) *string { return v.NetworkRuleBypassOptions }).(pulumi.StringPtrOutput)
}

// The network rule set for a container registry.
func (o LookupRegistryResultOutput) NetworkRuleSet() NetworkRuleSetResponsePtrOutput {
	return o.ApplyT(func(v LookupRegistryResult) *NetworkRuleSetResponse { return v.NetworkRuleSet }).(NetworkRuleSetResponsePtrOutput)
}

// The policies for a container registry.
func (o LookupRegistryResultOutput) Policies() PoliciesResponsePtrOutput {
	return o.ApplyT(func(v LookupRegistryResult) *PoliciesResponse { return v.Policies }).(PoliciesResponsePtrOutput)
}

// List of private endpoint connections for a container registry.
func (o LookupRegistryResultOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v LookupRegistryResult) []PrivateEndpointConnectionResponse { return v.PrivateEndpointConnections }).(PrivateEndpointConnectionResponseArrayOutput)
}

// The provisioning state of the container registry at the time the operation was called.
func (o LookupRegistryResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Whether or not public network access is allowed for the container registry.
func (o LookupRegistryResultOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRegistryResult) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

// The SKU of the container registry.
func (o LookupRegistryResultOutput) Sku() SkuResponseOutput {
	return o.ApplyT(func(v LookupRegistryResult) SkuResponse { return v.Sku }).(SkuResponseOutput)
}

// The status of the container registry at the time the operation was called.
func (o LookupRegistryResultOutput) Status() StatusResponseOutput {
	return o.ApplyT(func(v LookupRegistryResult) StatusResponse { return v.Status }).(StatusResponseOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupRegistryResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupRegistryResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The tags of the resource.
func (o LookupRegistryResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupRegistryResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource.
func (o LookupRegistryResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistryResult) string { return v.Type }).(pulumi.StringOutput)
}

// Whether or not zone redundancy is enabled for this container registry
func (o LookupRegistryResultOutput) ZoneRedundancy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRegistryResult) *string { return v.ZoneRedundancy }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRegistryResultOutput{})
}
