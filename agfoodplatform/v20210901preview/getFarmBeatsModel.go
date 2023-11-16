// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get FarmBeats resource.
func LookupFarmBeatsModel(ctx *pulumi.Context, args *LookupFarmBeatsModelArgs, opts ...pulumi.InvokeOption) (*LookupFarmBeatsModelResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupFarmBeatsModelResult
	err := ctx.Invoke("azure-native:agfoodplatform/v20210901preview:getFarmBeatsModel", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFarmBeatsModelArgs struct {
	// FarmBeats resource name.
	FarmBeatsResourceName string `pulumi:"farmBeatsResourceName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// FarmBeats ARM Resource.
type LookupFarmBeatsModelResult struct {
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// Identity for the resource.
	Identity *IdentityResponse `pulumi:"identity"`
	// Uri of the FarmBeats instance.
	InstanceUri string `pulumi:"instanceUri"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The private endpoint connection resource.
	PrivateEndpointConnections PrivateEndpointConnectionResponse `pulumi:"privateEndpointConnections"`
	// FarmBeats instance provisioning state.
	ProvisioningState string `pulumi:"provisioningState"`
	// Property to allow or block public traffic for an Azure FarmBeats resource.
	PublicNetworkAccess *string `pulumi:"publicNetworkAccess"`
	// Sensor integration request model.
	SensorIntegration *SensorIntegrationResponse `pulumi:"sensorIntegration"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupFarmBeatsModelOutput(ctx *pulumi.Context, args LookupFarmBeatsModelOutputArgs, opts ...pulumi.InvokeOption) LookupFarmBeatsModelResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFarmBeatsModelResult, error) {
			args := v.(LookupFarmBeatsModelArgs)
			r, err := LookupFarmBeatsModel(ctx, &args, opts...)
			var s LookupFarmBeatsModelResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFarmBeatsModelResultOutput)
}

type LookupFarmBeatsModelOutputArgs struct {
	// FarmBeats resource name.
	FarmBeatsResourceName pulumi.StringInput `pulumi:"farmBeatsResourceName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupFarmBeatsModelOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFarmBeatsModelArgs)(nil)).Elem()
}

// FarmBeats ARM Resource.
type LookupFarmBeatsModelResultOutput struct{ *pulumi.OutputState }

func (LookupFarmBeatsModelResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFarmBeatsModelResult)(nil)).Elem()
}

func (o LookupFarmBeatsModelResultOutput) ToLookupFarmBeatsModelResultOutput() LookupFarmBeatsModelResultOutput {
	return o
}

func (o LookupFarmBeatsModelResultOutput) ToLookupFarmBeatsModelResultOutputWithContext(ctx context.Context) LookupFarmBeatsModelResultOutput {
	return o
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupFarmBeatsModelResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFarmBeatsModelResult) string { return v.Id }).(pulumi.StringOutput)
}

// Identity for the resource.
func (o LookupFarmBeatsModelResultOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupFarmBeatsModelResult) *IdentityResponse { return v.Identity }).(IdentityResponsePtrOutput)
}

// Uri of the FarmBeats instance.
func (o LookupFarmBeatsModelResultOutput) InstanceUri() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFarmBeatsModelResult) string { return v.InstanceUri }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupFarmBeatsModelResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFarmBeatsModelResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupFarmBeatsModelResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFarmBeatsModelResult) string { return v.Name }).(pulumi.StringOutput)
}

// The private endpoint connection resource.
func (o LookupFarmBeatsModelResultOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseOutput {
	return o.ApplyT(func(v LookupFarmBeatsModelResult) PrivateEndpointConnectionResponse {
		return v.PrivateEndpointConnections
	}).(PrivateEndpointConnectionResponseOutput)
}

// FarmBeats instance provisioning state.
func (o LookupFarmBeatsModelResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFarmBeatsModelResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Property to allow or block public traffic for an Azure FarmBeats resource.
func (o LookupFarmBeatsModelResultOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFarmBeatsModelResult) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

// Sensor integration request model.
func (o LookupFarmBeatsModelResultOutput) SensorIntegration() SensorIntegrationResponsePtrOutput {
	return o.ApplyT(func(v LookupFarmBeatsModelResult) *SensorIntegrationResponse { return v.SensorIntegration }).(SensorIntegrationResponsePtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupFarmBeatsModelResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupFarmBeatsModelResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupFarmBeatsModelResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupFarmBeatsModelResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupFarmBeatsModelResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFarmBeatsModelResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFarmBeatsModelResultOutput{})
}
