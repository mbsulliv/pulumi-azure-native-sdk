// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets an attached NetworkConnection.
func LookupAttachedNetworkByDevCenter(ctx *pulumi.Context, args *LookupAttachedNetworkByDevCenterArgs, opts ...pulumi.InvokeOption) (*LookupAttachedNetworkByDevCenterResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupAttachedNetworkByDevCenterResult
	err := ctx.Invoke("azure-native:devcenter/v20230401:getAttachedNetworkByDevCenter", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAttachedNetworkByDevCenterArgs struct {
	// The name of the attached NetworkConnection.
	AttachedNetworkConnectionName string `pulumi:"attachedNetworkConnectionName"`
	// The name of the devcenter.
	DevCenterName string `pulumi:"devCenterName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Represents an attached NetworkConnection.
type LookupAttachedNetworkByDevCenterResult struct {
	// AAD Join type of the network. This is populated based on the referenced Network Connection.
	DomainJoinType string `pulumi:"domainJoinType"`
	// Health check status values
	HealthCheckStatus string `pulumi:"healthCheckStatus"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The resource ID of the NetworkConnection you want to attach.
	NetworkConnectionId string `pulumi:"networkConnectionId"`
	// The geo-location where the NetworkConnection resource specified in 'networkConnectionResourceId' property lives.
	NetworkConnectionLocation string `pulumi:"networkConnectionLocation"`
	// The provisioning state of the resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupAttachedNetworkByDevCenterOutput(ctx *pulumi.Context, args LookupAttachedNetworkByDevCenterOutputArgs, opts ...pulumi.InvokeOption) LookupAttachedNetworkByDevCenterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAttachedNetworkByDevCenterResult, error) {
			args := v.(LookupAttachedNetworkByDevCenterArgs)
			r, err := LookupAttachedNetworkByDevCenter(ctx, &args, opts...)
			var s LookupAttachedNetworkByDevCenterResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAttachedNetworkByDevCenterResultOutput)
}

type LookupAttachedNetworkByDevCenterOutputArgs struct {
	// The name of the attached NetworkConnection.
	AttachedNetworkConnectionName pulumi.StringInput `pulumi:"attachedNetworkConnectionName"`
	// The name of the devcenter.
	DevCenterName pulumi.StringInput `pulumi:"devCenterName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupAttachedNetworkByDevCenterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAttachedNetworkByDevCenterArgs)(nil)).Elem()
}

// Represents an attached NetworkConnection.
type LookupAttachedNetworkByDevCenterResultOutput struct{ *pulumi.OutputState }

func (LookupAttachedNetworkByDevCenterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAttachedNetworkByDevCenterResult)(nil)).Elem()
}

func (o LookupAttachedNetworkByDevCenterResultOutput) ToLookupAttachedNetworkByDevCenterResultOutput() LookupAttachedNetworkByDevCenterResultOutput {
	return o
}

func (o LookupAttachedNetworkByDevCenterResultOutput) ToLookupAttachedNetworkByDevCenterResultOutputWithContext(ctx context.Context) LookupAttachedNetworkByDevCenterResultOutput {
	return o
}

func (o LookupAttachedNetworkByDevCenterResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupAttachedNetworkByDevCenterResult] {
	return pulumix.Output[LookupAttachedNetworkByDevCenterResult]{
		OutputState: o.OutputState,
	}
}

// AAD Join type of the network. This is populated based on the referenced Network Connection.
func (o LookupAttachedNetworkByDevCenterResultOutput) DomainJoinType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAttachedNetworkByDevCenterResult) string { return v.DomainJoinType }).(pulumi.StringOutput)
}

// Health check status values
func (o LookupAttachedNetworkByDevCenterResultOutput) HealthCheckStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAttachedNetworkByDevCenterResult) string { return v.HealthCheckStatus }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupAttachedNetworkByDevCenterResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAttachedNetworkByDevCenterResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupAttachedNetworkByDevCenterResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAttachedNetworkByDevCenterResult) string { return v.Name }).(pulumi.StringOutput)
}

// The resource ID of the NetworkConnection you want to attach.
func (o LookupAttachedNetworkByDevCenterResultOutput) NetworkConnectionId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAttachedNetworkByDevCenterResult) string { return v.NetworkConnectionId }).(pulumi.StringOutput)
}

// The geo-location where the NetworkConnection resource specified in 'networkConnectionResourceId' property lives.
func (o LookupAttachedNetworkByDevCenterResultOutput) NetworkConnectionLocation() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAttachedNetworkByDevCenterResult) string { return v.NetworkConnectionLocation }).(pulumi.StringOutput)
}

// The provisioning state of the resource.
func (o LookupAttachedNetworkByDevCenterResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAttachedNetworkByDevCenterResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupAttachedNetworkByDevCenterResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupAttachedNetworkByDevCenterResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupAttachedNetworkByDevCenterResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAttachedNetworkByDevCenterResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAttachedNetworkByDevCenterResultOutput{})
}
