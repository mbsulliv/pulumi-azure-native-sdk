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

// Gets Kubernetes Cluster Extension.
func LookupExtension(ctx *pulumi.Context, args *LookupExtensionArgs, opts ...pulumi.InvokeOption) (*LookupExtensionResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupExtensionResult
	err := ctx.Invoke("azure-native:kubernetesconfiguration/v20230501:getExtension", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupExtensionArgs struct {
	// The name of the kubernetes cluster.
	ClusterName string `pulumi:"clusterName"`
	// The Kubernetes cluster resource name - i.e. managedClusters, connectedClusters, provisionedClusters.
	ClusterResourceName string `pulumi:"clusterResourceName"`
	// The Kubernetes cluster RP - i.e. Microsoft.ContainerService, Microsoft.Kubernetes, Microsoft.HybridContainerService.
	ClusterRp string `pulumi:"clusterRp"`
	// Name of the Extension.
	ExtensionName string `pulumi:"extensionName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The Extension object.
type LookupExtensionResult struct {
	// Identity of the Extension resource in an AKS cluster
	AksAssignedIdentity *ExtensionResponseAksAssignedIdentity `pulumi:"aksAssignedIdentity"`
	// Flag to note if this extension participates in auto upgrade of minor version, or not.
	AutoUpgradeMinorVersion *bool `pulumi:"autoUpgradeMinorVersion"`
	// Configuration settings that are sensitive, as name-value pairs for configuring this extension.
	ConfigurationProtectedSettings map[string]string `pulumi:"configurationProtectedSettings"`
	// Configuration settings, as name-value pairs for configuring this extension.
	ConfigurationSettings map[string]string `pulumi:"configurationSettings"`
	// Currently installed version of the extension.
	CurrentVersion string `pulumi:"currentVersion"`
	// Custom Location settings properties.
	CustomLocationSettings map[string]string `pulumi:"customLocationSettings"`
	// Error information from the Agent - e.g. errors during installation.
	ErrorInfo ErrorDetailResponse `pulumi:"errorInfo"`
	// Type of the Extension, of which this resource is an instance of.  It must be one of the Extension Types registered with Microsoft.KubernetesConfiguration by the Extension publisher.
	ExtensionType *string `pulumi:"extensionType"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// Identity of the Extension resource
	Identity *IdentityResponse `pulumi:"identity"`
	// Flag to note if this extension is a system extension
	IsSystemExtension bool `pulumi:"isSystemExtension"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Uri of the Helm package
	PackageUri string `pulumi:"packageUri"`
	// The plan information.
	Plan *PlanResponse `pulumi:"plan"`
	// Status of installation of this extension.
	ProvisioningState string `pulumi:"provisioningState"`
	// ReleaseTrain this extension participates in for auto-upgrade (e.g. Stable, Preview, etc.) - only if autoUpgradeMinorVersion is 'true'.
	ReleaseTrain *string `pulumi:"releaseTrain"`
	// Scope at which the extension is installed.
	Scope *ScopeResponse `pulumi:"scope"`
	// Status from this extension.
	Statuses []ExtensionStatusResponse `pulumi:"statuses"`
	// Top level metadata https://github.com/Azure/azure-resource-manager-rpc/blob/master/v1.0/common-api-contracts.md#system-metadata-for-all-azure-resources
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// User-specified version of the extension for this extension to 'pin'. To use 'version', autoUpgradeMinorVersion must be 'false'.
	Version *string `pulumi:"version"`
}

// Defaults sets the appropriate defaults for LookupExtensionResult
func (val *LookupExtensionResult) Defaults() *LookupExtensionResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.AutoUpgradeMinorVersion == nil {
		autoUpgradeMinorVersion_ := true
		tmp.AutoUpgradeMinorVersion = &autoUpgradeMinorVersion_
	}
	if utilities.IsZero(tmp.IsSystemExtension) {
		tmp.IsSystemExtension = false
	}
	if tmp.ReleaseTrain == nil {
		releaseTrain_ := "Stable"
		tmp.ReleaseTrain = &releaseTrain_
	}
	return &tmp
}

func LookupExtensionOutput(ctx *pulumi.Context, args LookupExtensionOutputArgs, opts ...pulumi.InvokeOption) LookupExtensionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupExtensionResult, error) {
			args := v.(LookupExtensionArgs)
			r, err := LookupExtension(ctx, &args, opts...)
			var s LookupExtensionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupExtensionResultOutput)
}

type LookupExtensionOutputArgs struct {
	// The name of the kubernetes cluster.
	ClusterName pulumi.StringInput `pulumi:"clusterName"`
	// The Kubernetes cluster resource name - i.e. managedClusters, connectedClusters, provisionedClusters.
	ClusterResourceName pulumi.StringInput `pulumi:"clusterResourceName"`
	// The Kubernetes cluster RP - i.e. Microsoft.ContainerService, Microsoft.Kubernetes, Microsoft.HybridContainerService.
	ClusterRp pulumi.StringInput `pulumi:"clusterRp"`
	// Name of the Extension.
	ExtensionName pulumi.StringInput `pulumi:"extensionName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupExtensionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupExtensionArgs)(nil)).Elem()
}

// The Extension object.
type LookupExtensionResultOutput struct{ *pulumi.OutputState }

func (LookupExtensionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupExtensionResult)(nil)).Elem()
}

func (o LookupExtensionResultOutput) ToLookupExtensionResultOutput() LookupExtensionResultOutput {
	return o
}

func (o LookupExtensionResultOutput) ToLookupExtensionResultOutputWithContext(ctx context.Context) LookupExtensionResultOutput {
	return o
}

func (o LookupExtensionResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupExtensionResult] {
	return pulumix.Output[LookupExtensionResult]{
		OutputState: o.OutputState,
	}
}

// Identity of the Extension resource in an AKS cluster
func (o LookupExtensionResultOutput) AksAssignedIdentity() ExtensionResponseAksAssignedIdentityPtrOutput {
	return o.ApplyT(func(v LookupExtensionResult) *ExtensionResponseAksAssignedIdentity { return v.AksAssignedIdentity }).(ExtensionResponseAksAssignedIdentityPtrOutput)
}

// Flag to note if this extension participates in auto upgrade of minor version, or not.
func (o LookupExtensionResultOutput) AutoUpgradeMinorVersion() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupExtensionResult) *bool { return v.AutoUpgradeMinorVersion }).(pulumi.BoolPtrOutput)
}

// Configuration settings that are sensitive, as name-value pairs for configuring this extension.
func (o LookupExtensionResultOutput) ConfigurationProtectedSettings() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupExtensionResult) map[string]string { return v.ConfigurationProtectedSettings }).(pulumi.StringMapOutput)
}

// Configuration settings, as name-value pairs for configuring this extension.
func (o LookupExtensionResultOutput) ConfigurationSettings() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupExtensionResult) map[string]string { return v.ConfigurationSettings }).(pulumi.StringMapOutput)
}

// Currently installed version of the extension.
func (o LookupExtensionResultOutput) CurrentVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExtensionResult) string { return v.CurrentVersion }).(pulumi.StringOutput)
}

// Custom Location settings properties.
func (o LookupExtensionResultOutput) CustomLocationSettings() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupExtensionResult) map[string]string { return v.CustomLocationSettings }).(pulumi.StringMapOutput)
}

// Error information from the Agent - e.g. errors during installation.
func (o LookupExtensionResultOutput) ErrorInfo() ErrorDetailResponseOutput {
	return o.ApplyT(func(v LookupExtensionResult) ErrorDetailResponse { return v.ErrorInfo }).(ErrorDetailResponseOutput)
}

// Type of the Extension, of which this resource is an instance of.  It must be one of the Extension Types registered with Microsoft.KubernetesConfiguration by the Extension publisher.
func (o LookupExtensionResultOutput) ExtensionType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExtensionResult) *string { return v.ExtensionType }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupExtensionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExtensionResult) string { return v.Id }).(pulumi.StringOutput)
}

// Identity of the Extension resource
func (o LookupExtensionResultOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupExtensionResult) *IdentityResponse { return v.Identity }).(IdentityResponsePtrOutput)
}

// Flag to note if this extension is a system extension
func (o LookupExtensionResultOutput) IsSystemExtension() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupExtensionResult) bool { return v.IsSystemExtension }).(pulumi.BoolOutput)
}

// The name of the resource
func (o LookupExtensionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExtensionResult) string { return v.Name }).(pulumi.StringOutput)
}

// Uri of the Helm package
func (o LookupExtensionResultOutput) PackageUri() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExtensionResult) string { return v.PackageUri }).(pulumi.StringOutput)
}

// The plan information.
func (o LookupExtensionResultOutput) Plan() PlanResponsePtrOutput {
	return o.ApplyT(func(v LookupExtensionResult) *PlanResponse { return v.Plan }).(PlanResponsePtrOutput)
}

// Status of installation of this extension.
func (o LookupExtensionResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExtensionResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// ReleaseTrain this extension participates in for auto-upgrade (e.g. Stable, Preview, etc.) - only if autoUpgradeMinorVersion is 'true'.
func (o LookupExtensionResultOutput) ReleaseTrain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExtensionResult) *string { return v.ReleaseTrain }).(pulumi.StringPtrOutput)
}

// Scope at which the extension is installed.
func (o LookupExtensionResultOutput) Scope() ScopeResponsePtrOutput {
	return o.ApplyT(func(v LookupExtensionResult) *ScopeResponse { return v.Scope }).(ScopeResponsePtrOutput)
}

// Status from this extension.
func (o LookupExtensionResultOutput) Statuses() ExtensionStatusResponseArrayOutput {
	return o.ApplyT(func(v LookupExtensionResult) []ExtensionStatusResponse { return v.Statuses }).(ExtensionStatusResponseArrayOutput)
}

// Top level metadata https://github.com/Azure/azure-resource-manager-rpc/blob/master/v1.0/common-api-contracts.md#system-metadata-for-all-azure-resources
func (o LookupExtensionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupExtensionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupExtensionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExtensionResult) string { return v.Type }).(pulumi.StringOutput)
}

// User-specified version of the extension for this extension to 'pin'. To use 'version', autoUpgradeMinorVersion must be 'false'.
func (o LookupExtensionResultOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExtensionResult) *string { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupExtensionResultOutput{})
}
