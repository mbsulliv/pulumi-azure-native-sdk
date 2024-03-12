// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package orbital

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the specified L2 connection in a specified resource group.
// Azure REST API version: 2024-03-01-preview.
//
// Other available API versions: 2024-03-01.
func LookupL2Connection(ctx *pulumi.Context, args *LookupL2ConnectionArgs, opts ...pulumi.InvokeOption) (*LookupL2ConnectionResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupL2ConnectionResult
	err := ctx.Invoke("azure-native:orbital:getL2Connection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupL2ConnectionArgs struct {
	// L2 Connection name.
	L2ConnectionName string `pulumi:"l2ConnectionName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Connects an edge site to an orbital gateway and describes what layer 2 traffic to forward between them.
type LookupL2ConnectionResult struct {
	// Globally-unique identifier for this connection that is to be used as a circuit ID.
	CircuitId string `pulumi:"circuitId"`
	// A reference to an Microsoft.Orbital/edgeSites resource to route traffic for.
	EdgeSite L2ConnectionsPropertiesResponseEdgeSite `pulumi:"edgeSite"`
	// A reference to an Microsoft.Orbital/groundStations resource to route traffic for.
	GroundStation L2ConnectionsPropertiesResponseGroundStation `pulumi:"groundStation"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// The VLAN ID for the L2 connection.
	VlanId int `pulumi:"vlanId"`
}

func LookupL2ConnectionOutput(ctx *pulumi.Context, args LookupL2ConnectionOutputArgs, opts ...pulumi.InvokeOption) LookupL2ConnectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupL2ConnectionResult, error) {
			args := v.(LookupL2ConnectionArgs)
			r, err := LookupL2Connection(ctx, &args, opts...)
			var s LookupL2ConnectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupL2ConnectionResultOutput)
}

type LookupL2ConnectionOutputArgs struct {
	// L2 Connection name.
	L2ConnectionName pulumi.StringInput `pulumi:"l2ConnectionName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupL2ConnectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupL2ConnectionArgs)(nil)).Elem()
}

// Connects an edge site to an orbital gateway and describes what layer 2 traffic to forward between them.
type LookupL2ConnectionResultOutput struct{ *pulumi.OutputState }

func (LookupL2ConnectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupL2ConnectionResult)(nil)).Elem()
}

func (o LookupL2ConnectionResultOutput) ToLookupL2ConnectionResultOutput() LookupL2ConnectionResultOutput {
	return o
}

func (o LookupL2ConnectionResultOutput) ToLookupL2ConnectionResultOutputWithContext(ctx context.Context) LookupL2ConnectionResultOutput {
	return o
}

// Globally-unique identifier for this connection that is to be used as a circuit ID.
func (o LookupL2ConnectionResultOutput) CircuitId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupL2ConnectionResult) string { return v.CircuitId }).(pulumi.StringOutput)
}

// A reference to an Microsoft.Orbital/edgeSites resource to route traffic for.
func (o LookupL2ConnectionResultOutput) EdgeSite() L2ConnectionsPropertiesResponseEdgeSiteOutput {
	return o.ApplyT(func(v LookupL2ConnectionResult) L2ConnectionsPropertiesResponseEdgeSite { return v.EdgeSite }).(L2ConnectionsPropertiesResponseEdgeSiteOutput)
}

// A reference to an Microsoft.Orbital/groundStations resource to route traffic for.
func (o LookupL2ConnectionResultOutput) GroundStation() L2ConnectionsPropertiesResponseGroundStationOutput {
	return o.ApplyT(func(v LookupL2ConnectionResult) L2ConnectionsPropertiesResponseGroundStation { return v.GroundStation }).(L2ConnectionsPropertiesResponseGroundStationOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupL2ConnectionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupL2ConnectionResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupL2ConnectionResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupL2ConnectionResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupL2ConnectionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupL2ConnectionResult) string { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupL2ConnectionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupL2ConnectionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupL2ConnectionResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupL2ConnectionResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupL2ConnectionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupL2ConnectionResult) string { return v.Type }).(pulumi.StringOutput)
}

// The VLAN ID for the L2 connection.
func (o LookupL2ConnectionResultOutput) VlanId() pulumi.IntOutput {
	return o.ApplyT(func(v LookupL2ConnectionResult) int { return v.VlanId }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupL2ConnectionResultOutput{})
}
