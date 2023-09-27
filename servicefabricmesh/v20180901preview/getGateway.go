// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20180901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets the information about the gateway resource with the given name. The information include the description and other properties of the gateway.
func LookupGateway(ctx *pulumi.Context, args *LookupGatewayArgs, opts ...pulumi.InvokeOption) (*LookupGatewayResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupGatewayResult
	err := ctx.Invoke("azure-native:servicefabricmesh/v20180901preview:getGateway", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGatewayArgs struct {
	// The identity of the gateway.
	GatewayResourceName string `pulumi:"gatewayResourceName"`
	// Azure resource group name
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// This type describes a gateway resource.
type LookupGatewayResult struct {
	// User readable description of the gateway.
	Description *string `pulumi:"description"`
	// Network that the Application is using.
	DestinationNetwork NetworkRefResponse `pulumi:"destinationNetwork"`
	// Configuration for http connectivity for this gateway.
	Http []HttpConfigResponse `pulumi:"http"`
	// Fully qualified identifier for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// IP address of the gateway. This is populated in the response and is ignored for incoming requests.
	IpAddress string `pulumi:"ipAddress"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// State of the resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Network the gateway should listen on for requests.
	SourceNetwork NetworkRefResponse `pulumi:"sourceNetwork"`
	// Status of the resource.
	Status string `pulumi:"status"`
	// Gives additional information about the current status of the gateway.
	StatusDetails string `pulumi:"statusDetails"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Configuration for tcp connectivity for this gateway.
	Tcp []TcpConfigResponse `pulumi:"tcp"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type string `pulumi:"type"`
}

func LookupGatewayOutput(ctx *pulumi.Context, args LookupGatewayOutputArgs, opts ...pulumi.InvokeOption) LookupGatewayResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGatewayResult, error) {
			args := v.(LookupGatewayArgs)
			r, err := LookupGateway(ctx, &args, opts...)
			var s LookupGatewayResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGatewayResultOutput)
}

type LookupGatewayOutputArgs struct {
	// The identity of the gateway.
	GatewayResourceName pulumi.StringInput `pulumi:"gatewayResourceName"`
	// Azure resource group name
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupGatewayOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGatewayArgs)(nil)).Elem()
}

// This type describes a gateway resource.
type LookupGatewayResultOutput struct{ *pulumi.OutputState }

func (LookupGatewayResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGatewayResult)(nil)).Elem()
}

func (o LookupGatewayResultOutput) ToLookupGatewayResultOutput() LookupGatewayResultOutput {
	return o
}

func (o LookupGatewayResultOutput) ToLookupGatewayResultOutputWithContext(ctx context.Context) LookupGatewayResultOutput {
	return o
}

func (o LookupGatewayResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupGatewayResult] {
	return pulumix.Output[LookupGatewayResult]{
		OutputState: o.OutputState,
	}
}

// User readable description of the gateway.
func (o LookupGatewayResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGatewayResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Network that the Application is using.
func (o LookupGatewayResultOutput) DestinationNetwork() NetworkRefResponseOutput {
	return o.ApplyT(func(v LookupGatewayResult) NetworkRefResponse { return v.DestinationNetwork }).(NetworkRefResponseOutput)
}

// Configuration for http connectivity for this gateway.
func (o LookupGatewayResultOutput) Http() HttpConfigResponseArrayOutput {
	return o.ApplyT(func(v LookupGatewayResult) []HttpConfigResponse { return v.Http }).(HttpConfigResponseArrayOutput)
}

// Fully qualified identifier for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupGatewayResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayResult) string { return v.Id }).(pulumi.StringOutput)
}

// IP address of the gateway. This is populated in the response and is ignored for incoming requests.
func (o LookupGatewayResultOutput) IpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayResult) string { return v.IpAddress }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupGatewayResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupGatewayResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayResult) string { return v.Name }).(pulumi.StringOutput)
}

// State of the resource.
func (o LookupGatewayResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Network the gateway should listen on for requests.
func (o LookupGatewayResultOutput) SourceNetwork() NetworkRefResponseOutput {
	return o.ApplyT(func(v LookupGatewayResult) NetworkRefResponse { return v.SourceNetwork }).(NetworkRefResponseOutput)
}

// Status of the resource.
func (o LookupGatewayResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayResult) string { return v.Status }).(pulumi.StringOutput)
}

// Gives additional information about the current status of the gateway.
func (o LookupGatewayResultOutput) StatusDetails() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayResult) string { return v.StatusDetails }).(pulumi.StringOutput)
}

// Resource tags.
func (o LookupGatewayResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupGatewayResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Configuration for tcp connectivity for this gateway.
func (o LookupGatewayResultOutput) Tcp() TcpConfigResponseArrayOutput {
	return o.ApplyT(func(v LookupGatewayResult) []TcpConfigResponse { return v.Tcp }).(TcpConfigResponseArrayOutput)
}

// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
func (o LookupGatewayResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGatewayResultOutput{})
}
