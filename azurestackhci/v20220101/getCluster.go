// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get HCI cluster.
func LookupCluster(ctx *pulumi.Context, args *LookupClusterArgs, opts ...pulumi.InvokeOption) (*LookupClusterResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupClusterResult
	err := ctx.Invoke("azure-native:azurestackhci/v20220101:getCluster", args, &rv, opts...)
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
	// App id of cluster AAD identity.
	AadClientId string `pulumi:"aadClientId"`
	// Tenant id of cluster AAD identity.
	AadTenantId string `pulumi:"aadTenantId"`
	// Type of billing applied to the resource.
	BillingModel string `pulumi:"billingModel"`
	// Unique, immutable resource id.
	CloudId string `pulumi:"cloudId"`
	// Endpoint configured for management from the Azure portal.
	CloudManagementEndpoint *string `pulumi:"cloudManagementEndpoint"`
	// The timestamp of resource creation (UTC).
	CreatedAt *string `pulumi:"createdAt"`
	// The identity that created the resource.
	CreatedBy *string `pulumi:"createdBy"`
	// The type of identity that created the resource.
	CreatedByType *string `pulumi:"createdByType"`
	// Desired properties of the cluster.
	DesiredProperties *ClusterDesiredPropertiesResponse `pulumi:"desiredProperties"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// Most recent billing meter timestamp.
	LastBillingTimestamp string `pulumi:"lastBillingTimestamp"`
	// The timestamp of resource last modification (UTC)
	LastModifiedAt *string `pulumi:"lastModifiedAt"`
	// The identity that last modified the resource.
	LastModifiedBy *string `pulumi:"lastModifiedBy"`
	// The type of identity that last modified the resource.
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
	// Most recent cluster sync timestamp.
	LastSyncTimestamp string `pulumi:"lastSyncTimestamp"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Provisioning state.
	ProvisioningState string `pulumi:"provisioningState"`
	// First cluster sync timestamp.
	RegistrationTimestamp string `pulumi:"registrationTimestamp"`
	// Properties reported by cluster agent.
	ReportedProperties ClusterReportedPropertiesResponse `pulumi:"reportedProperties"`
	// Status of the cluster agent.
	Status string `pulumi:"status"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Number of days remaining in the trial period.
	TrialDaysRemaining float64 `pulumi:"trialDaysRemaining"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
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

func (o LookupClusterResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupClusterResult] {
	return pulumix.Output[LookupClusterResult]{
		OutputState: o.OutputState,
	}
}

// App id of cluster AAD identity.
func (o LookupClusterResultOutput) AadClientId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.AadClientId }).(pulumi.StringOutput)
}

// Tenant id of cluster AAD identity.
func (o LookupClusterResultOutput) AadTenantId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.AadTenantId }).(pulumi.StringOutput)
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

// The timestamp of resource creation (UTC).
func (o LookupClusterResultOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

// The identity that created the resource.
func (o LookupClusterResultOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *string { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

// The type of identity that created the resource.
func (o LookupClusterResultOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *string { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

// Desired properties of the cluster.
func (o LookupClusterResultOutput) DesiredProperties() ClusterDesiredPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *ClusterDesiredPropertiesResponse { return v.DesiredProperties }).(ClusterDesiredPropertiesResponsePtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupClusterResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.Id }).(pulumi.StringOutput)
}

// Most recent billing meter timestamp.
func (o LookupClusterResultOutput) LastBillingTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.LastBillingTimestamp }).(pulumi.StringOutput)
}

// The timestamp of resource last modification (UTC)
func (o LookupClusterResultOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *string { return v.LastModifiedAt }).(pulumi.StringPtrOutput)
}

// The identity that last modified the resource.
func (o LookupClusterResultOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *string { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

// The type of identity that last modified the resource.
func (o LookupClusterResultOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *string { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
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

// Status of the cluster agent.
func (o LookupClusterResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.Status }).(pulumi.StringOutput)
}

// Resource tags.
func (o LookupClusterResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupClusterResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Number of days remaining in the trial period.
func (o LookupClusterResultOutput) TrialDaysRemaining() pulumi.Float64Output {
	return o.ApplyT(func(v LookupClusterResult) float64 { return v.TrialDaysRemaining }).(pulumi.Float64Output)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupClusterResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupClusterResultOutput{})
}
