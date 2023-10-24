// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mobilenetwork

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets information about the specified mobile network.
// Azure REST API version: 2023-06-01.
//
// Other available API versions: 2022-04-01-preview, 2022-11-01, 2023-09-01.
func LookupMobileNetwork(ctx *pulumi.Context, args *LookupMobileNetworkArgs, opts ...pulumi.InvokeOption) (*LookupMobileNetworkResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupMobileNetworkResult
	err := ctx.Invoke("azure-native:mobilenetwork:getMobileNetwork", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMobileNetworkArgs struct {
	// The name of the mobile network.
	MobileNetworkName string `pulumi:"mobileNetworkName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Mobile network resource.
type LookupMobileNetworkResult struct {
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The provisioning state of the mobile network resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// The unique public land mobile network identifier for the network. This is made up of the mobile country code and mobile network code, as defined in https://www.itu.int/rec/T-REC-E.212. The values 001-01 and 001-001 can be used for testing and the values 999-99 and 999-999 can be used on internal private networks.
	PublicLandMobileNetworkIdentifier PlmnIdResponse `pulumi:"publicLandMobileNetworkIdentifier"`
	// The mobile network resource identifier
	ServiceKey string `pulumi:"serviceKey"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupMobileNetworkOutput(ctx *pulumi.Context, args LookupMobileNetworkOutputArgs, opts ...pulumi.InvokeOption) LookupMobileNetworkResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMobileNetworkResult, error) {
			args := v.(LookupMobileNetworkArgs)
			r, err := LookupMobileNetwork(ctx, &args, opts...)
			var s LookupMobileNetworkResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMobileNetworkResultOutput)
}

type LookupMobileNetworkOutputArgs struct {
	// The name of the mobile network.
	MobileNetworkName pulumi.StringInput `pulumi:"mobileNetworkName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupMobileNetworkOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMobileNetworkArgs)(nil)).Elem()
}

// Mobile network resource.
type LookupMobileNetworkResultOutput struct{ *pulumi.OutputState }

func (LookupMobileNetworkResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMobileNetworkResult)(nil)).Elem()
}

func (o LookupMobileNetworkResultOutput) ToLookupMobileNetworkResultOutput() LookupMobileNetworkResultOutput {
	return o
}

func (o LookupMobileNetworkResultOutput) ToLookupMobileNetworkResultOutputWithContext(ctx context.Context) LookupMobileNetworkResultOutput {
	return o
}

func (o LookupMobileNetworkResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupMobileNetworkResult] {
	return pulumix.Output[LookupMobileNetworkResult]{
		OutputState: o.OutputState,
	}
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupMobileNetworkResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMobileNetworkResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupMobileNetworkResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMobileNetworkResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupMobileNetworkResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMobileNetworkResult) string { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the mobile network resource.
func (o LookupMobileNetworkResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMobileNetworkResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The unique public land mobile network identifier for the network. This is made up of the mobile country code and mobile network code, as defined in https://www.itu.int/rec/T-REC-E.212. The values 001-01 and 001-001 can be used for testing and the values 999-99 and 999-999 can be used on internal private networks.
func (o LookupMobileNetworkResultOutput) PublicLandMobileNetworkIdentifier() PlmnIdResponseOutput {
	return o.ApplyT(func(v LookupMobileNetworkResult) PlmnIdResponse { return v.PublicLandMobileNetworkIdentifier }).(PlmnIdResponseOutput)
}

// The mobile network resource identifier
func (o LookupMobileNetworkResultOutput) ServiceKey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMobileNetworkResult) string { return v.ServiceKey }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupMobileNetworkResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupMobileNetworkResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupMobileNetworkResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupMobileNetworkResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupMobileNetworkResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMobileNetworkResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMobileNetworkResultOutput{})
}
