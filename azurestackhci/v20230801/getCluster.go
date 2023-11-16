// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get HCI cluster.
func LookupCluster(ctx *pulumi.Context, args *LookupClusterArgs, opts ...pulumi.InvokeOption) (*LookupClusterResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupClusterResult
	err := ctx.Invoke("azure-native:azurestackhci/v20230801:getCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupClusterArgs struct {
	// The name of the cluster.
	ClusterName string `pulumi:"clusterName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Cluster details.
type LookupClusterResult struct {
	// Object id of cluster AAD identity.
	AadApplicationObjectId *string `pulumi:"aadApplicationObjectId"`
	// App id of cluster AAD identity.
	AadClientId *string `pulumi:"aadClientId"`
	// Id of cluster identity service principal.
	AadServicePrincipalObjectId *string `pulumi:"aadServicePrincipalObjectId"`
	// Tenant id of cluster AAD identity.
	AadTenantId *string `pulumi:"aadTenantId"`
	// Type of billing applied to the resource.
	BillingModel string `pulumi:"billingModel"`
	// Unique, immutable resource id.
	CloudId string `pulumi:"cloudId"`
	// Endpoint configured for management from the Azure portal.
	CloudManagementEndpoint *string `pulumi:"cloudManagementEndpoint"`
	// Overall connectivity status for the cluster resource.
	ConnectivityStatus string `pulumi:"connectivityStatus"`
	// Desired properties of the cluster.
	DesiredProperties *ClusterDesiredPropertiesResponse `pulumi:"desiredProperties"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// Attestation configurations for isolated VM (e.g. TVM, CVM) of the cluster.
	IsolatedVmAttestationConfiguration IsolatedVmAttestationConfigurationResponse `pulumi:"isolatedVmAttestationConfiguration"`
	// Most recent billing meter timestamp.
	LastBillingTimestamp string `pulumi:"lastBillingTimestamp"`
	// Most recent cluster sync timestamp.
	LastSyncTimestamp string `pulumi:"lastSyncTimestamp"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The service principal ID of the system assigned identity. This property will only be provided for a system assigned identity.
	PrincipalId string `pulumi:"principalId"`
	// Provisioning state.
	ProvisioningState string `pulumi:"provisioningState"`
	// First cluster sync timestamp.
	RegistrationTimestamp string `pulumi:"registrationTimestamp"`
	// Properties reported by cluster agent.
	ReportedProperties ClusterReportedPropertiesResponse `pulumi:"reportedProperties"`
	// Object id of RP Service Principal
	ResourceProviderObjectId string `pulumi:"resourceProviderObjectId"`
	// Region specific DataPath Endpoint of the cluster.
	ServiceEndpoint string `pulumi:"serviceEndpoint"`
	// Software Assurance properties of the cluster.
	SoftwareAssuranceProperties *SoftwareAssurancePropertiesResponse `pulumi:"softwareAssuranceProperties"`
	// Status of the cluster agent.
	Status string `pulumi:"status"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The tenant ID of the system assigned identity. This property will only be provided for a system assigned identity.
	TenantId string `pulumi:"tenantId"`
	// Number of days remaining in the trial period.
	TrialDaysRemaining float64 `pulumi:"trialDaysRemaining"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// The set of user assigned identities associated with the resource. The userAssignedIdentities dictionary keys will be ARM resource ids in the form: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}. The dictionary values can be empty objects ({}) in requests.
	UserAssignedIdentities map[string]UserAssignedIdentityResponse `pulumi:"userAssignedIdentities"`
}

func LookupClusterOutput(ctx *pulumi.Context, args LookupClusterOutputArgs, opts ...pulumi.InvokeOption) LookupClusterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupClusterResult, error) {
			args := v.(LookupClusterArgs)
			r, err := LookupCluster(ctx, &args, opts...)
			var s LookupClusterResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupClusterResultOutput)
}

type LookupClusterOutputArgs struct {
	// The name of the cluster.
	ClusterName pulumi.StringInput `pulumi:"clusterName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupClusterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupClusterArgs)(nil)).Elem()
}

// Cluster details.
type LookupClusterResultOutput struct{ *pulumi.OutputState }

func (LookupClusterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupClusterResult)(nil)).Elem()
}

func (o LookupClusterResultOutput) ToLookupClusterResultOutput() LookupClusterResultOutput {
	return o
}

func (o LookupClusterResultOutput) ToLookupClusterResultOutputWithContext(ctx context.Context) LookupClusterResultOutput {
	return o
}

// Object id of cluster AAD identity.
func (o LookupClusterResultOutput) AadApplicationObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *string { return v.AadApplicationObjectId }).(pulumi.StringPtrOutput)
}

// App id of cluster AAD identity.
func (o LookupClusterResultOutput) AadClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *string { return v.AadClientId }).(pulumi.StringPtrOutput)
}

// Id of cluster identity service principal.
func (o LookupClusterResultOutput) AadServicePrincipalObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *string { return v.AadServicePrincipalObjectId }).(pulumi.StringPtrOutput)
}

// Tenant id of cluster AAD identity.
func (o LookupClusterResultOutput) AadTenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *string { return v.AadTenantId }).(pulumi.StringPtrOutput)
}

// Type of billing applied to the resource.
func (o LookupClusterResultOutput) BillingModel() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.BillingModel }).(pulumi.StringOutput)
}

// Unique, immutable resource id.
func (o LookupClusterResultOutput) CloudId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.CloudId }).(pulumi.StringOutput)
}

// Endpoint configured for management from the Azure portal.
func (o LookupClusterResultOutput) CloudManagementEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *string { return v.CloudManagementEndpoint }).(pulumi.StringPtrOutput)
}

// Overall connectivity status for the cluster resource.
func (o LookupClusterResultOutput) ConnectivityStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.ConnectivityStatus }).(pulumi.StringOutput)
}

// Desired properties of the cluster.
func (o LookupClusterResultOutput) DesiredProperties() ClusterDesiredPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *ClusterDesiredPropertiesResponse { return v.DesiredProperties }).(ClusterDesiredPropertiesResponsePtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupClusterResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.Id }).(pulumi.StringOutput)
}

// Attestation configurations for isolated VM (e.g. TVM, CVM) of the cluster.
func (o LookupClusterResultOutput) IsolatedVmAttestationConfiguration() IsolatedVmAttestationConfigurationResponseOutput {
	return o.ApplyT(func(v LookupClusterResult) IsolatedVmAttestationConfigurationResponse {
		return v.IsolatedVmAttestationConfiguration
	}).(IsolatedVmAttestationConfigurationResponseOutput)
}

// Most recent billing meter timestamp.
func (o LookupClusterResultOutput) LastBillingTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.LastBillingTimestamp }).(pulumi.StringOutput)
}

// Most recent cluster sync timestamp.
func (o LookupClusterResultOutput) LastSyncTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.LastSyncTimestamp }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupClusterResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupClusterResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.Name }).(pulumi.StringOutput)
}

// The service principal ID of the system assigned identity. This property will only be provided for a system assigned identity.
func (o LookupClusterResultOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.PrincipalId }).(pulumi.StringOutput)
}

// Provisioning state.
func (o LookupClusterResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// First cluster sync timestamp.
func (o LookupClusterResultOutput) RegistrationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.RegistrationTimestamp }).(pulumi.StringOutput)
}

// Properties reported by cluster agent.
func (o LookupClusterResultOutput) ReportedProperties() ClusterReportedPropertiesResponseOutput {
	return o.ApplyT(func(v LookupClusterResult) ClusterReportedPropertiesResponse { return v.ReportedProperties }).(ClusterReportedPropertiesResponseOutput)
}

// Object id of RP Service Principal
func (o LookupClusterResultOutput) ResourceProviderObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.ResourceProviderObjectId }).(pulumi.StringOutput)
}

// Region specific DataPath Endpoint of the cluster.
func (o LookupClusterResultOutput) ServiceEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.ServiceEndpoint }).(pulumi.StringOutput)
}

// Software Assurance properties of the cluster.
func (o LookupClusterResultOutput) SoftwareAssuranceProperties() SoftwareAssurancePropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *SoftwareAssurancePropertiesResponse { return v.SoftwareAssuranceProperties }).(SoftwareAssurancePropertiesResponsePtrOutput)
}

// Status of the cluster agent.
func (o LookupClusterResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.Status }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupClusterResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupClusterResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupClusterResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupClusterResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The tenant ID of the system assigned identity. This property will only be provided for a system assigned identity.
func (o LookupClusterResultOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.TenantId }).(pulumi.StringOutput)
}

// Number of days remaining in the trial period.
func (o LookupClusterResultOutput) TrialDaysRemaining() pulumi.Float64Output {
	return o.ApplyT(func(v LookupClusterResult) float64 { return v.TrialDaysRemaining }).(pulumi.Float64Output)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupClusterResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.Type }).(pulumi.StringOutput)
}

// The set of user assigned identities associated with the resource. The userAssignedIdentities dictionary keys will be ARM resource ids in the form: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}. The dictionary values can be empty objects ({}) in requests.
func (o LookupClusterResultOutput) UserAssignedIdentities() UserAssignedIdentityResponseMapOutput {
	return o.ApplyT(func(v LookupClusterResult) map[string]UserAssignedIdentityResponse { return v.UserAssignedIdentities }).(UserAssignedIdentityResponseMapOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupClusterResultOutput{})
}
