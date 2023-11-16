// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package redhatopenshift

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The operation returns properties of a OpenShift cluster.
// Azure REST API version: 2022-09-04.
//
// Other available API versions: 2023-04-01, 2023-07-01-preview, 2023-09-04.
func LookupOpenShiftCluster(ctx *pulumi.Context, args *LookupOpenShiftClusterArgs, opts ...pulumi.InvokeOption) (*LookupOpenShiftClusterResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupOpenShiftClusterResult
	err := ctx.Invoke("azure-native:redhatopenshift:getOpenShiftCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupOpenShiftClusterArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the OpenShift cluster resource.
	ResourceName string `pulumi:"resourceName"`
}

// OpenShiftCluster represents an Azure Red Hat OpenShift cluster.
type LookupOpenShiftClusterResult struct {
	// The cluster API server profile.
	ApiserverProfile *APIServerProfileResponse `pulumi:"apiserverProfile"`
	// The cluster profile.
	ClusterProfile *ClusterProfileResponse `pulumi:"clusterProfile"`
	// The console profile.
	ConsoleProfile *ConsoleProfileResponse `pulumi:"consoleProfile"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The cluster ingress profiles.
	IngressProfiles []IngressProfileResponse `pulumi:"ingressProfiles"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The cluster master profile.
	MasterProfile *MasterProfileResponse `pulumi:"masterProfile"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The cluster network profile.
	NetworkProfile *NetworkProfileResponse `pulumi:"networkProfile"`
	// The cluster provisioning state.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The cluster service principal profile.
	ServicePrincipalProfile *ServicePrincipalProfileResponse `pulumi:"servicePrincipalProfile"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// The cluster worker profiles.
	WorkerProfiles []WorkerProfileResponse `pulumi:"workerProfiles"`
}

func LookupOpenShiftClusterOutput(ctx *pulumi.Context, args LookupOpenShiftClusterOutputArgs, opts ...pulumi.InvokeOption) LookupOpenShiftClusterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupOpenShiftClusterResult, error) {
			args := v.(LookupOpenShiftClusterArgs)
			r, err := LookupOpenShiftCluster(ctx, &args, opts...)
			var s LookupOpenShiftClusterResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupOpenShiftClusterResultOutput)
}

type LookupOpenShiftClusterOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the OpenShift cluster resource.
	ResourceName pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupOpenShiftClusterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOpenShiftClusterArgs)(nil)).Elem()
}

// OpenShiftCluster represents an Azure Red Hat OpenShift cluster.
type LookupOpenShiftClusterResultOutput struct{ *pulumi.OutputState }

func (LookupOpenShiftClusterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOpenShiftClusterResult)(nil)).Elem()
}

func (o LookupOpenShiftClusterResultOutput) ToLookupOpenShiftClusterResultOutput() LookupOpenShiftClusterResultOutput {
	return o
}

func (o LookupOpenShiftClusterResultOutput) ToLookupOpenShiftClusterResultOutputWithContext(ctx context.Context) LookupOpenShiftClusterResultOutput {
	return o
}

// The cluster API server profile.
func (o LookupOpenShiftClusterResultOutput) ApiserverProfile() APIServerProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupOpenShiftClusterResult) *APIServerProfileResponse { return v.ApiserverProfile }).(APIServerProfileResponsePtrOutput)
}

// The cluster profile.
func (o LookupOpenShiftClusterResultOutput) ClusterProfile() ClusterProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupOpenShiftClusterResult) *ClusterProfileResponse { return v.ClusterProfile }).(ClusterProfileResponsePtrOutput)
}

// The console profile.
func (o LookupOpenShiftClusterResultOutput) ConsoleProfile() ConsoleProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupOpenShiftClusterResult) *ConsoleProfileResponse { return v.ConsoleProfile }).(ConsoleProfileResponsePtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupOpenShiftClusterResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOpenShiftClusterResult) string { return v.Id }).(pulumi.StringOutput)
}

// The cluster ingress profiles.
func (o LookupOpenShiftClusterResultOutput) IngressProfiles() IngressProfileResponseArrayOutput {
	return o.ApplyT(func(v LookupOpenShiftClusterResult) []IngressProfileResponse { return v.IngressProfiles }).(IngressProfileResponseArrayOutput)
}

// The geo-location where the resource lives
func (o LookupOpenShiftClusterResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOpenShiftClusterResult) string { return v.Location }).(pulumi.StringOutput)
}

// The cluster master profile.
func (o LookupOpenShiftClusterResultOutput) MasterProfile() MasterProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupOpenShiftClusterResult) *MasterProfileResponse { return v.MasterProfile }).(MasterProfileResponsePtrOutput)
}

// The name of the resource
func (o LookupOpenShiftClusterResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOpenShiftClusterResult) string { return v.Name }).(pulumi.StringOutput)
}

// The cluster network profile.
func (o LookupOpenShiftClusterResultOutput) NetworkProfile() NetworkProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupOpenShiftClusterResult) *NetworkProfileResponse { return v.NetworkProfile }).(NetworkProfileResponsePtrOutput)
}

// The cluster provisioning state.
func (o LookupOpenShiftClusterResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupOpenShiftClusterResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

// The cluster service principal profile.
func (o LookupOpenShiftClusterResultOutput) ServicePrincipalProfile() ServicePrincipalProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupOpenShiftClusterResult) *ServicePrincipalProfileResponse {
		return v.ServicePrincipalProfile
	}).(ServicePrincipalProfileResponsePtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupOpenShiftClusterResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupOpenShiftClusterResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupOpenShiftClusterResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupOpenShiftClusterResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupOpenShiftClusterResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOpenShiftClusterResult) string { return v.Type }).(pulumi.StringOutput)
}

// The cluster worker profiles.
func (o LookupOpenShiftClusterResultOutput) WorkerProfiles() WorkerProfileResponseArrayOutput {
	return o.ApplyT(func(v LookupOpenShiftClusterResult) []WorkerProfileResponse { return v.WorkerProfiles }).(WorkerProfileResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupOpenShiftClusterResultOutput{})
}
