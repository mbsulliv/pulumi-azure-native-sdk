// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets information about the specified Batch account.
func LookupBatchAccount(ctx *pulumi.Context, args *LookupBatchAccountArgs, opts ...pulumi.InvokeOption) (*LookupBatchAccountResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupBatchAccountResult
	err := ctx.Invoke("azure-native:batch/v20230501:getBatchAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupBatchAccountArgs struct {
	// The name of the Batch account.
	AccountName string `pulumi:"accountName"`
	// The name of the resource group that contains the Batch account.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Contains information about an Azure Batch account.
type LookupBatchAccountResult struct {
	// The account endpoint used to interact with the Batch service.
	AccountEndpoint              string `pulumi:"accountEndpoint"`
	ActiveJobAndJobScheduleQuota int    `pulumi:"activeJobAndJobScheduleQuota"`
	// List of allowed authentication modes for the Batch account that can be used to authenticate with the data plane. This does not affect authentication with the control plane.
	AllowedAuthenticationModes []string `pulumi:"allowedAuthenticationModes"`
	// Contains information about the auto-storage account associated with a Batch account.
	AutoStorage AutoStoragePropertiesResponse `pulumi:"autoStorage"`
	// For accounts with PoolAllocationMode set to UserSubscription, quota is managed on the subscription so this value is not returned.
	DedicatedCoreQuota int `pulumi:"dedicatedCoreQuota"`
	// A list of the dedicated core quota per Virtual Machine family for the Batch account. For accounts with PoolAllocationMode set to UserSubscription, quota is managed on the subscription so this value is not returned.
	DedicatedCoreQuotaPerVMFamily []VirtualMachineFamilyCoreQuotaResponse `pulumi:"dedicatedCoreQuotaPerVMFamily"`
	// If this flag is true, dedicated core quota is enforced via both the dedicatedCoreQuotaPerVMFamily and dedicatedCoreQuota properties on the account. If this flag is false, dedicated core quota is enforced only via the dedicatedCoreQuota property on the account and does not consider Virtual Machine family.
	DedicatedCoreQuotaPerVMFamilyEnforced bool `pulumi:"dedicatedCoreQuotaPerVMFamilyEnforced"`
	// Configures how customer data is encrypted inside the Batch account. By default, accounts are encrypted using a Microsoft managed key. For additional control, a customer-managed key can be used instead.
	Encryption EncryptionPropertiesResponse `pulumi:"encryption"`
	// The ID of the resource.
	Id string `pulumi:"id"`
	// The identity of the Batch account.
	Identity *BatchAccountIdentityResponse `pulumi:"identity"`
	// Identifies the Azure key vault associated with a Batch account.
	KeyVaultReference KeyVaultReferenceResponse `pulumi:"keyVaultReference"`
	// The location of the resource.
	Location string `pulumi:"location"`
	// For accounts with PoolAllocationMode set to UserSubscription, quota is managed on the subscription so this value is not returned.
	LowPriorityCoreQuota int `pulumi:"lowPriorityCoreQuota"`
	// The name of the resource.
	Name string `pulumi:"name"`
	// The network profile only takes effect when publicNetworkAccess is enabled.
	NetworkProfile *NetworkProfileResponse `pulumi:"networkProfile"`
	// The endpoint used by compute node to connect to the Batch node management service.
	NodeManagementEndpoint string `pulumi:"nodeManagementEndpoint"`
	// The allocation mode for creating pools in the Batch account.
	PoolAllocationMode string `pulumi:"poolAllocationMode"`
	PoolQuota          int    `pulumi:"poolQuota"`
	// List of private endpoint connections associated with the Batch account
	PrivateEndpointConnections []PrivateEndpointConnectionResponse `pulumi:"privateEndpointConnections"`
	// The provisioned state of the resource
	ProvisioningState string `pulumi:"provisioningState"`
	// If not specified, the default value is 'enabled'.
	PublicNetworkAccess *string `pulumi:"publicNetworkAccess"`
	// The tags of the resource.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource.
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupBatchAccountResult
func (val *LookupBatchAccountResult) Defaults() *LookupBatchAccountResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.AutoStorage = *tmp.AutoStorage.Defaults()

	return &tmp
}

func LookupBatchAccountOutput(ctx *pulumi.Context, args LookupBatchAccountOutputArgs, opts ...pulumi.InvokeOption) LookupBatchAccountResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBatchAccountResult, error) {
			args := v.(LookupBatchAccountArgs)
			r, err := LookupBatchAccount(ctx, &args, opts...)
			var s LookupBatchAccountResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBatchAccountResultOutput)
}

type LookupBatchAccountOutputArgs struct {
	// The name of the Batch account.
	AccountName pulumi.StringInput `pulumi:"accountName"`
	// The name of the resource group that contains the Batch account.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupBatchAccountOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBatchAccountArgs)(nil)).Elem()
}

// Contains information about an Azure Batch account.
type LookupBatchAccountResultOutput struct{ *pulumi.OutputState }

func (LookupBatchAccountResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBatchAccountResult)(nil)).Elem()
}

func (o LookupBatchAccountResultOutput) ToLookupBatchAccountResultOutput() LookupBatchAccountResultOutput {
	return o
}

func (o LookupBatchAccountResultOutput) ToLookupBatchAccountResultOutputWithContext(ctx context.Context) LookupBatchAccountResultOutput {
	return o
}

func (o LookupBatchAccountResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupBatchAccountResult] {
	return pulumix.Output[LookupBatchAccountResult]{
		OutputState: o.OutputState,
	}
}

// The account endpoint used to interact with the Batch service.
func (o LookupBatchAccountResultOutput) AccountEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBatchAccountResult) string { return v.AccountEndpoint }).(pulumi.StringOutput)
}

func (o LookupBatchAccountResultOutput) ActiveJobAndJobScheduleQuota() pulumi.IntOutput {
	return o.ApplyT(func(v LookupBatchAccountResult) int { return v.ActiveJobAndJobScheduleQuota }).(pulumi.IntOutput)
}

// List of allowed authentication modes for the Batch account that can be used to authenticate with the data plane. This does not affect authentication with the control plane.
func (o LookupBatchAccountResultOutput) AllowedAuthenticationModes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupBatchAccountResult) []string { return v.AllowedAuthenticationModes }).(pulumi.StringArrayOutput)
}

// Contains information about the auto-storage account associated with a Batch account.
func (o LookupBatchAccountResultOutput) AutoStorage() AutoStoragePropertiesResponseOutput {
	return o.ApplyT(func(v LookupBatchAccountResult) AutoStoragePropertiesResponse { return v.AutoStorage }).(AutoStoragePropertiesResponseOutput)
}

// For accounts with PoolAllocationMode set to UserSubscription, quota is managed on the subscription so this value is not returned.
func (o LookupBatchAccountResultOutput) DedicatedCoreQuota() pulumi.IntOutput {
	return o.ApplyT(func(v LookupBatchAccountResult) int { return v.DedicatedCoreQuota }).(pulumi.IntOutput)
}

// A list of the dedicated core quota per Virtual Machine family for the Batch account. For accounts with PoolAllocationMode set to UserSubscription, quota is managed on the subscription so this value is not returned.
func (o LookupBatchAccountResultOutput) DedicatedCoreQuotaPerVMFamily() VirtualMachineFamilyCoreQuotaResponseArrayOutput {
	return o.ApplyT(func(v LookupBatchAccountResult) []VirtualMachineFamilyCoreQuotaResponse {
		return v.DedicatedCoreQuotaPerVMFamily
	}).(VirtualMachineFamilyCoreQuotaResponseArrayOutput)
}

// If this flag is true, dedicated core quota is enforced via both the dedicatedCoreQuotaPerVMFamily and dedicatedCoreQuota properties on the account. If this flag is false, dedicated core quota is enforced only via the dedicatedCoreQuota property on the account and does not consider Virtual Machine family.
func (o LookupBatchAccountResultOutput) DedicatedCoreQuotaPerVMFamilyEnforced() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupBatchAccountResult) bool { return v.DedicatedCoreQuotaPerVMFamilyEnforced }).(pulumi.BoolOutput)
}

// Configures how customer data is encrypted inside the Batch account. By default, accounts are encrypted using a Microsoft managed key. For additional control, a customer-managed key can be used instead.
func (o LookupBatchAccountResultOutput) Encryption() EncryptionPropertiesResponseOutput {
	return o.ApplyT(func(v LookupBatchAccountResult) EncryptionPropertiesResponse { return v.Encryption }).(EncryptionPropertiesResponseOutput)
}

// The ID of the resource.
func (o LookupBatchAccountResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBatchAccountResult) string { return v.Id }).(pulumi.StringOutput)
}

// The identity of the Batch account.
func (o LookupBatchAccountResultOutput) Identity() BatchAccountIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupBatchAccountResult) *BatchAccountIdentityResponse { return v.Identity }).(BatchAccountIdentityResponsePtrOutput)
}

// Identifies the Azure key vault associated with a Batch account.
func (o LookupBatchAccountResultOutput) KeyVaultReference() KeyVaultReferenceResponseOutput {
	return o.ApplyT(func(v LookupBatchAccountResult) KeyVaultReferenceResponse { return v.KeyVaultReference }).(KeyVaultReferenceResponseOutput)
}

// The location of the resource.
func (o LookupBatchAccountResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBatchAccountResult) string { return v.Location }).(pulumi.StringOutput)
}

// For accounts with PoolAllocationMode set to UserSubscription, quota is managed on the subscription so this value is not returned.
func (o LookupBatchAccountResultOutput) LowPriorityCoreQuota() pulumi.IntOutput {
	return o.ApplyT(func(v LookupBatchAccountResult) int { return v.LowPriorityCoreQuota }).(pulumi.IntOutput)
}

// The name of the resource.
func (o LookupBatchAccountResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBatchAccountResult) string { return v.Name }).(pulumi.StringOutput)
}

// The network profile only takes effect when publicNetworkAccess is enabled.
func (o LookupBatchAccountResultOutput) NetworkProfile() NetworkProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupBatchAccountResult) *NetworkProfileResponse { return v.NetworkProfile }).(NetworkProfileResponsePtrOutput)
}

// The endpoint used by compute node to connect to the Batch node management service.
func (o LookupBatchAccountResultOutput) NodeManagementEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBatchAccountResult) string { return v.NodeManagementEndpoint }).(pulumi.StringOutput)
}

// The allocation mode for creating pools in the Batch account.
func (o LookupBatchAccountResultOutput) PoolAllocationMode() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBatchAccountResult) string { return v.PoolAllocationMode }).(pulumi.StringOutput)
}

func (o LookupBatchAccountResultOutput) PoolQuota() pulumi.IntOutput {
	return o.ApplyT(func(v LookupBatchAccountResult) int { return v.PoolQuota }).(pulumi.IntOutput)
}

// List of private endpoint connections associated with the Batch account
func (o LookupBatchAccountResultOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v LookupBatchAccountResult) []PrivateEndpointConnectionResponse {
		return v.PrivateEndpointConnections
	}).(PrivateEndpointConnectionResponseArrayOutput)
}

// The provisioned state of the resource
func (o LookupBatchAccountResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBatchAccountResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// If not specified, the default value is 'enabled'.
func (o LookupBatchAccountResultOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBatchAccountResult) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

// The tags of the resource.
func (o LookupBatchAccountResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupBatchAccountResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource.
func (o LookupBatchAccountResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBatchAccountResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBatchAccountResultOutput{})
}
