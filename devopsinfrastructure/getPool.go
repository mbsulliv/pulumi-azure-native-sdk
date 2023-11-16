// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package devopsinfrastructure

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get a Pool
// Azure REST API version: 2023-10-30-preview.
func LookupPool(ctx *pulumi.Context, args *LookupPoolArgs, opts ...pulumi.InvokeOption) (*LookupPoolResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupPoolResult
	err := ctx.Invoke("azure-native:devopsinfrastructure:getPool", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupPoolArgs struct {
	// Name of the pool. It needs to be globally unique.
	PoolName string `pulumi:"poolName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Concrete tracked resource types can be created by aliasing this type using a specific property type.
type LookupPoolResult struct {
	// Defines how the machine will be handled once it executed a job.
	AgentProfile interface{} `pulumi:"agentProfile"`
	// The resource id of the DevCenter Project the pool belongs to.
	DevCenterProjectResourceId string `pulumi:"devCenterProjectResourceId"`
	// Defines the type of fabric the agent will run on.
	FabricProfile VmssFabricProfileResponse `pulumi:"fabricProfile"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The managed service identities assigned to this resource.
	Identity *ManagedServiceIdentityResponse `pulumi:"identity"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// Defines how many resources can there be created at any given time.
	MaximumConcurrency int `pulumi:"maximumConcurrency"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Defines the organization in which the pool will be used.
	OrganizationProfile AzureDevOpsOrganizationProfileResponse `pulumi:"organizationProfile"`
	// The status of the current operation.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupPoolResult
func (val *LookupPoolResult) Defaults() *LookupPoolResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.FabricProfile = *tmp.FabricProfile.Defaults()

	return &tmp
}

func LookupPoolOutput(ctx *pulumi.Context, args LookupPoolOutputArgs, opts ...pulumi.InvokeOption) LookupPoolResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPoolResult, error) {
			args := v.(LookupPoolArgs)
			r, err := LookupPool(ctx, &args, opts...)
			var s LookupPoolResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPoolResultOutput)
}

type LookupPoolOutputArgs struct {
	// Name of the pool. It needs to be globally unique.
	PoolName pulumi.StringInput `pulumi:"poolName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupPoolOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPoolArgs)(nil)).Elem()
}

// Concrete tracked resource types can be created by aliasing this type using a specific property type.
type LookupPoolResultOutput struct{ *pulumi.OutputState }

func (LookupPoolResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPoolResult)(nil)).Elem()
}

func (o LookupPoolResultOutput) ToLookupPoolResultOutput() LookupPoolResultOutput {
	return o
}

func (o LookupPoolResultOutput) ToLookupPoolResultOutputWithContext(ctx context.Context) LookupPoolResultOutput {
	return o
}

// Defines how the machine will be handled once it executed a job.
func (o LookupPoolResultOutput) AgentProfile() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupPoolResult) interface{} { return v.AgentProfile }).(pulumi.AnyOutput)
}

// The resource id of the DevCenter Project the pool belongs to.
func (o LookupPoolResultOutput) DevCenterProjectResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPoolResult) string { return v.DevCenterProjectResourceId }).(pulumi.StringOutput)
}

// Defines the type of fabric the agent will run on.
func (o LookupPoolResultOutput) FabricProfile() VmssFabricProfileResponseOutput {
	return o.ApplyT(func(v LookupPoolResult) VmssFabricProfileResponse { return v.FabricProfile }).(VmssFabricProfileResponseOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupPoolResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPoolResult) string { return v.Id }).(pulumi.StringOutput)
}

// The managed service identities assigned to this resource.
func (o LookupPoolResultOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupPoolResult) *ManagedServiceIdentityResponse { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

// The geo-location where the resource lives
func (o LookupPoolResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPoolResult) string { return v.Location }).(pulumi.StringOutput)
}

// Defines how many resources can there be created at any given time.
func (o LookupPoolResultOutput) MaximumConcurrency() pulumi.IntOutput {
	return o.ApplyT(func(v LookupPoolResult) int { return v.MaximumConcurrency }).(pulumi.IntOutput)
}

// The name of the resource
func (o LookupPoolResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPoolResult) string { return v.Name }).(pulumi.StringOutput)
}

// Defines the organization in which the pool will be used.
func (o LookupPoolResultOutput) OrganizationProfile() AzureDevOpsOrganizationProfileResponseOutput {
	return o.ApplyT(func(v LookupPoolResult) AzureDevOpsOrganizationProfileResponse { return v.OrganizationProfile }).(AzureDevOpsOrganizationProfileResponseOutput)
}

// The status of the current operation.
func (o LookupPoolResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPoolResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupPoolResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupPoolResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupPoolResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupPoolResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupPoolResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPoolResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPoolResultOutput{})
}
