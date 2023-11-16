// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20231115preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the Hybrid AKS virtual network
func LookupVirtualNetworkRetrieve(ctx *pulumi.Context, args *LookupVirtualNetworkRetrieveArgs, opts ...pulumi.InvokeOption) (*LookupVirtualNetworkRetrieveResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupVirtualNetworkRetrieveResult
	err := ctx.Invoke("azure-native:hybridcontainerservice/v20231115preview:getVirtualNetworkRetrieve", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVirtualNetworkRetrieveArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Parameter for the name of the virtual network
	VirtualNetworkName string `pulumi:"virtualNetworkName"`
}

// The virtualNetworks resource definition.
type LookupVirtualNetworkRetrieveResult struct {
	ExtendedLocation *VirtualNetworkResponseExtendedLocation `pulumi:"extendedLocation"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// HybridAKSNetworkSpec defines the desired state of HybridAKSNetwork
	Properties VirtualNetworkPropertiesResponse `pulumi:"properties"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupVirtualNetworkRetrieveOutput(ctx *pulumi.Context, args LookupVirtualNetworkRetrieveOutputArgs, opts ...pulumi.InvokeOption) LookupVirtualNetworkRetrieveResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVirtualNetworkRetrieveResult, error) {
			args := v.(LookupVirtualNetworkRetrieveArgs)
			r, err := LookupVirtualNetworkRetrieve(ctx, &args, opts...)
			var s LookupVirtualNetworkRetrieveResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVirtualNetworkRetrieveResultOutput)
}

type LookupVirtualNetworkRetrieveOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Parameter for the name of the virtual network
	VirtualNetworkName pulumi.StringInput `pulumi:"virtualNetworkName"`
}

func (LookupVirtualNetworkRetrieveOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualNetworkRetrieveArgs)(nil)).Elem()
}

// The virtualNetworks resource definition.
type LookupVirtualNetworkRetrieveResultOutput struct{ *pulumi.OutputState }

func (LookupVirtualNetworkRetrieveResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualNetworkRetrieveResult)(nil)).Elem()
}

func (o LookupVirtualNetworkRetrieveResultOutput) ToLookupVirtualNetworkRetrieveResultOutput() LookupVirtualNetworkRetrieveResultOutput {
	return o
}

func (o LookupVirtualNetworkRetrieveResultOutput) ToLookupVirtualNetworkRetrieveResultOutputWithContext(ctx context.Context) LookupVirtualNetworkRetrieveResultOutput {
	return o
}

func (o LookupVirtualNetworkRetrieveResultOutput) ExtendedLocation() VirtualNetworkResponseExtendedLocationPtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkRetrieveResult) *VirtualNetworkResponseExtendedLocation {
		return v.ExtendedLocation
	}).(VirtualNetworkResponseExtendedLocationPtrOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupVirtualNetworkRetrieveResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkRetrieveResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupVirtualNetworkRetrieveResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkRetrieveResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupVirtualNetworkRetrieveResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkRetrieveResult) string { return v.Name }).(pulumi.StringOutput)
}

// HybridAKSNetworkSpec defines the desired state of HybridAKSNetwork
func (o LookupVirtualNetworkRetrieveResultOutput) Properties() VirtualNetworkPropertiesResponseOutput {
	return o.ApplyT(func(v LookupVirtualNetworkRetrieveResult) VirtualNetworkPropertiesResponse { return v.Properties }).(VirtualNetworkPropertiesResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupVirtualNetworkRetrieveResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupVirtualNetworkRetrieveResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupVirtualNetworkRetrieveResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupVirtualNetworkRetrieveResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupVirtualNetworkRetrieveResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkRetrieveResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVirtualNetworkRetrieveResultOutput{})
}
