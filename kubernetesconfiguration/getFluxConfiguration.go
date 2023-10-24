// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package kubernetesconfiguration

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets details of the Flux Configuration.
// Azure REST API version: 2023-05-01.
//
// Other available API versions: 2021-11-01-preview, 2022-01-01-preview.
func LookupFluxConfiguration(ctx *pulumi.Context, args *LookupFluxConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupFluxConfigurationResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupFluxConfigurationResult
	err := ctx.Invoke("azure-native:kubernetesconfiguration:getFluxConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupFluxConfigurationArgs struct {
	// The name of the kubernetes cluster.
	ClusterName string `pulumi:"clusterName"`
	// The Kubernetes cluster resource name - i.e. managedClusters, connectedClusters, provisionedClusters.
	ClusterResourceName string `pulumi:"clusterResourceName"`
	// The Kubernetes cluster RP - i.e. Microsoft.ContainerService, Microsoft.Kubernetes, Microsoft.HybridContainerService.
	ClusterRp string `pulumi:"clusterRp"`
	// Name of the Flux Configuration.
	FluxConfigurationName string `pulumi:"fluxConfigurationName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The Flux Configuration object returned in Get & Put response.
type LookupFluxConfigurationResult struct {
	// Parameters to reconcile to the AzureBlob source kind type.
	AzureBlob *AzureBlobDefinitionResponse `pulumi:"azureBlob"`
	// Parameters to reconcile to the Bucket source kind type.
	Bucket *BucketDefinitionResponse `pulumi:"bucket"`
	// Combined status of the Flux Kubernetes resources created by the fluxConfiguration or created by the managed objects.
	ComplianceState string `pulumi:"complianceState"`
	// Key-value pairs of protected configuration settings for the configuration
	ConfigurationProtectedSettings map[string]string `pulumi:"configurationProtectedSettings"`
	// Error message returned to the user in the case of provisioning failure.
	ErrorMessage string `pulumi:"errorMessage"`
	// Parameters to reconcile to the GitRepository source kind type.
	GitRepository *GitRepositoryDefinitionResponse `pulumi:"gitRepository"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// Array of kustomizations used to reconcile the artifact pulled by the source type on the cluster.
	Kustomizations map[string]KustomizationDefinitionResponse `pulumi:"kustomizations"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The namespace to which this configuration is installed to. Maximum of 253 lower case alphanumeric characters, hyphen and period only.
	Namespace *string `pulumi:"namespace"`
	// Status of the creation of the fluxConfiguration.
	ProvisioningState string `pulumi:"provisioningState"`
	// Maximum duration to wait for flux configuration reconciliation. E.g PT1H, PT5M, P1D
	ReconciliationWaitDuration *string `pulumi:"reconciliationWaitDuration"`
	// Public Key associated with this fluxConfiguration (either generated within the cluster or provided by the user).
	RepositoryPublicKey string `pulumi:"repositoryPublicKey"`
	// Scope at which the operator will be installed.
	Scope *string `pulumi:"scope"`
	// Source Kind to pull the configuration data from.
	SourceKind *string `pulumi:"sourceKind"`
	// Branch and/or SHA of the source commit synced with the cluster.
	SourceSyncedCommitId string `pulumi:"sourceSyncedCommitId"`
	// Datetime the fluxConfiguration synced its source on the cluster.
	SourceUpdatedAt string `pulumi:"sourceUpdatedAt"`
	// Datetime the fluxConfiguration synced its status on the cluster with Azure.
	StatusUpdatedAt string `pulumi:"statusUpdatedAt"`
	// Statuses of the Flux Kubernetes resources created by the fluxConfiguration or created by the managed objects provisioned by the fluxConfiguration.
	Statuses []ObjectStatusDefinitionResponse `pulumi:"statuses"`
	// Whether this configuration should suspend its reconciliation of its kustomizations and sources.
	Suspend *bool `pulumi:"suspend"`
	// Top level metadata https://github.com/Azure/azure-resource-manager-rpc/blob/master/v1.0/common-api-contracts.md#system-metadata-for-all-azure-resources
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// Whether flux configuration deployment should wait for cluster to reconcile the kustomizations.
	WaitForReconciliation *bool `pulumi:"waitForReconciliation"`
}

// Defaults sets the appropriate defaults for LookupFluxConfigurationResult
func (val *LookupFluxConfigurationResult) Defaults() *LookupFluxConfigurationResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.AzureBlob = tmp.AzureBlob.Defaults()

	tmp.Bucket = tmp.Bucket.Defaults()

	tmp.GitRepository = tmp.GitRepository.Defaults()

	if tmp.Namespace == nil {
		namespace_ := "default"
		tmp.Namespace = &namespace_
	}
	if tmp.SourceKind == nil {
		sourceKind_ := "GitRepository"
		tmp.SourceKind = &sourceKind_
	}
	if tmp.Suspend == nil {
		suspend_ := false
		tmp.Suspend = &suspend_
	}
	return &tmp
}

func LookupFluxConfigurationOutput(ctx *pulumi.Context, args LookupFluxConfigurationOutputArgs, opts ...pulumi.InvokeOption) LookupFluxConfigurationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFluxConfigurationResult, error) {
			args := v.(LookupFluxConfigurationArgs)
			r, err := LookupFluxConfiguration(ctx, &args, opts...)
			var s LookupFluxConfigurationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFluxConfigurationResultOutput)
}

type LookupFluxConfigurationOutputArgs struct {
	// The name of the kubernetes cluster.
	ClusterName pulumi.StringInput `pulumi:"clusterName"`
	// The Kubernetes cluster resource name - i.e. managedClusters, connectedClusters, provisionedClusters.
	ClusterResourceName pulumi.StringInput `pulumi:"clusterResourceName"`
	// The Kubernetes cluster RP - i.e. Microsoft.ContainerService, Microsoft.Kubernetes, Microsoft.HybridContainerService.
	ClusterRp pulumi.StringInput `pulumi:"clusterRp"`
	// Name of the Flux Configuration.
	FluxConfigurationName pulumi.StringInput `pulumi:"fluxConfigurationName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupFluxConfigurationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFluxConfigurationArgs)(nil)).Elem()
}

// The Flux Configuration object returned in Get & Put response.
type LookupFluxConfigurationResultOutput struct{ *pulumi.OutputState }

func (LookupFluxConfigurationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFluxConfigurationResult)(nil)).Elem()
}

func (o LookupFluxConfigurationResultOutput) ToLookupFluxConfigurationResultOutput() LookupFluxConfigurationResultOutput {
	return o
}

func (o LookupFluxConfigurationResultOutput) ToLookupFluxConfigurationResultOutputWithContext(ctx context.Context) LookupFluxConfigurationResultOutput {
	return o
}

func (o LookupFluxConfigurationResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupFluxConfigurationResult] {
	return pulumix.Output[LookupFluxConfigurationResult]{
		OutputState: o.OutputState,
	}
}

// Parameters to reconcile to the AzureBlob source kind type.
func (o LookupFluxConfigurationResultOutput) AzureBlob() AzureBlobDefinitionResponsePtrOutput {
	return o.ApplyT(func(v LookupFluxConfigurationResult) *AzureBlobDefinitionResponse { return v.AzureBlob }).(AzureBlobDefinitionResponsePtrOutput)
}

// Parameters to reconcile to the Bucket source kind type.
func (o LookupFluxConfigurationResultOutput) Bucket() BucketDefinitionResponsePtrOutput {
	return o.ApplyT(func(v LookupFluxConfigurationResult) *BucketDefinitionResponse { return v.Bucket }).(BucketDefinitionResponsePtrOutput)
}

// Combined status of the Flux Kubernetes resources created by the fluxConfiguration or created by the managed objects.
func (o LookupFluxConfigurationResultOutput) ComplianceState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFluxConfigurationResult) string { return v.ComplianceState }).(pulumi.StringOutput)
}

// Key-value pairs of protected configuration settings for the configuration
func (o LookupFluxConfigurationResultOutput) ConfigurationProtectedSettings() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupFluxConfigurationResult) map[string]string { return v.ConfigurationProtectedSettings }).(pulumi.StringMapOutput)
}

// Error message returned to the user in the case of provisioning failure.
func (o LookupFluxConfigurationResultOutput) ErrorMessage() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFluxConfigurationResult) string { return v.ErrorMessage }).(pulumi.StringOutput)
}

// Parameters to reconcile to the GitRepository source kind type.
func (o LookupFluxConfigurationResultOutput) GitRepository() GitRepositoryDefinitionResponsePtrOutput {
	return o.ApplyT(func(v LookupFluxConfigurationResult) *GitRepositoryDefinitionResponse { return v.GitRepository }).(GitRepositoryDefinitionResponsePtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupFluxConfigurationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFluxConfigurationResult) string { return v.Id }).(pulumi.StringOutput)
}

// Array of kustomizations used to reconcile the artifact pulled by the source type on the cluster.
func (o LookupFluxConfigurationResultOutput) Kustomizations() KustomizationDefinitionResponseMapOutput {
	return o.ApplyT(func(v LookupFluxConfigurationResult) map[string]KustomizationDefinitionResponse {
		return v.Kustomizations
	}).(KustomizationDefinitionResponseMapOutput)
}

// The name of the resource
func (o LookupFluxConfigurationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFluxConfigurationResult) string { return v.Name }).(pulumi.StringOutput)
}

// The namespace to which this configuration is installed to. Maximum of 253 lower case alphanumeric characters, hyphen and period only.
func (o LookupFluxConfigurationResultOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFluxConfigurationResult) *string { return v.Namespace }).(pulumi.StringPtrOutput)
}

// Status of the creation of the fluxConfiguration.
func (o LookupFluxConfigurationResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFluxConfigurationResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Maximum duration to wait for flux configuration reconciliation. E.g PT1H, PT5M, P1D
func (o LookupFluxConfigurationResultOutput) ReconciliationWaitDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFluxConfigurationResult) *string { return v.ReconciliationWaitDuration }).(pulumi.StringPtrOutput)
}

// Public Key associated with this fluxConfiguration (either generated within the cluster or provided by the user).
func (o LookupFluxConfigurationResultOutput) RepositoryPublicKey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFluxConfigurationResult) string { return v.RepositoryPublicKey }).(pulumi.StringOutput)
}

// Scope at which the operator will be installed.
func (o LookupFluxConfigurationResultOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFluxConfigurationResult) *string { return v.Scope }).(pulumi.StringPtrOutput)
}

// Source Kind to pull the configuration data from.
func (o LookupFluxConfigurationResultOutput) SourceKind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFluxConfigurationResult) *string { return v.SourceKind }).(pulumi.StringPtrOutput)
}

// Branch and/or SHA of the source commit synced with the cluster.
func (o LookupFluxConfigurationResultOutput) SourceSyncedCommitId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFluxConfigurationResult) string { return v.SourceSyncedCommitId }).(pulumi.StringOutput)
}

// Datetime the fluxConfiguration synced its source on the cluster.
func (o LookupFluxConfigurationResultOutput) SourceUpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFluxConfigurationResult) string { return v.SourceUpdatedAt }).(pulumi.StringOutput)
}

// Datetime the fluxConfiguration synced its status on the cluster with Azure.
func (o LookupFluxConfigurationResultOutput) StatusUpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFluxConfigurationResult) string { return v.StatusUpdatedAt }).(pulumi.StringOutput)
}

// Statuses of the Flux Kubernetes resources created by the fluxConfiguration or created by the managed objects provisioned by the fluxConfiguration.
func (o LookupFluxConfigurationResultOutput) Statuses() ObjectStatusDefinitionResponseArrayOutput {
	return o.ApplyT(func(v LookupFluxConfigurationResult) []ObjectStatusDefinitionResponse { return v.Statuses }).(ObjectStatusDefinitionResponseArrayOutput)
}

// Whether this configuration should suspend its reconciliation of its kustomizations and sources.
func (o LookupFluxConfigurationResultOutput) Suspend() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFluxConfigurationResult) *bool { return v.Suspend }).(pulumi.BoolPtrOutput)
}

// Top level metadata https://github.com/Azure/azure-resource-manager-rpc/blob/master/v1.0/common-api-contracts.md#system-metadata-for-all-azure-resources
func (o LookupFluxConfigurationResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupFluxConfigurationResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupFluxConfigurationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFluxConfigurationResult) string { return v.Type }).(pulumi.StringOutput)
}

// Whether flux configuration deployment should wait for cluster to reconcile the kustomizations.
func (o LookupFluxConfigurationResultOutput) WaitForReconciliation() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFluxConfigurationResult) *bool { return v.WaitForReconciliation }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFluxConfigurationResultOutput{})
}
