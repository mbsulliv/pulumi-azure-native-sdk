// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220702preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The Fleet resource which contains multiple Kubernetes clusters as its members.
func LookupFleet(ctx *pulumi.Context, args *LookupFleetArgs, opts ...pulumi.InvokeOption) (*LookupFleetResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupFleetResult
	err := ctx.Invoke("azure-native:containerservice/v20220702preview:getFleet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFleetArgs struct {
	// The name of the Fleet resource.
	FleetName string `pulumi:"fleetName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The Fleet resource which contains multiple Kubernetes clusters as its members.
type LookupFleetResult struct {
	// Resource Etag.
	Etag string `pulumi:"etag"`
	// The FleetHubProfile configures the Fleet's hub.
	HubProfile *FleetHubProfileResponse `pulumi:"hubProfile"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The provisioning state of the last accepted operation.
	ProvisioningState string `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupFleetOutput(ctx *pulumi.Context, args LookupFleetOutputArgs, opts ...pulumi.InvokeOption) LookupFleetResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFleetResult, error) {
			args := v.(LookupFleetArgs)
			r, err := LookupFleet(ctx, &args, opts...)
			var s LookupFleetResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFleetResultOutput)
}

type LookupFleetOutputArgs struct {
	// The name of the Fleet resource.
	FleetName pulumi.StringInput `pulumi:"fleetName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupFleetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFleetArgs)(nil)).Elem()
}

// The Fleet resource which contains multiple Kubernetes clusters as its members.
type LookupFleetResultOutput struct{ *pulumi.OutputState }

func (LookupFleetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFleetResult)(nil)).Elem()
}

func (o LookupFleetResultOutput) ToLookupFleetResultOutput() LookupFleetResultOutput {
	return o
}

func (o LookupFleetResultOutput) ToLookupFleetResultOutputWithContext(ctx context.Context) LookupFleetResultOutput {
	return o
}

// Resource Etag.
func (o LookupFleetResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFleetResult) string { return v.Etag }).(pulumi.StringOutput)
}

// The FleetHubProfile configures the Fleet's hub.
func (o LookupFleetResultOutput) HubProfile() FleetHubProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupFleetResult) *FleetHubProfileResponse { return v.HubProfile }).(FleetHubProfileResponsePtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupFleetResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFleetResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupFleetResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFleetResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupFleetResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFleetResult) string { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the last accepted operation.
func (o LookupFleetResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFleetResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupFleetResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupFleetResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupFleetResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupFleetResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupFleetResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFleetResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFleetResultOutput{})
}
