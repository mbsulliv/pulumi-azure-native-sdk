// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package hybridnetwork

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets information about the specified sku.
// Azure REST API version: 2022-01-01-preview.
func LookupVendorSkus(ctx *pulumi.Context, args *LookupVendorSkusArgs, opts ...pulumi.InvokeOption) (*LookupVendorSkusResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupVendorSkusResult
	err := ctx.Invoke("azure-native:hybridnetwork:getVendorSkus", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVendorSkusArgs struct {
	// The name of the sku.
	SkuName string `pulumi:"skuName"`
	// The name of the vendor.
	VendorName string `pulumi:"vendorName"`
}

// Sku sub resource.
type LookupVendorSkusResult struct {
	// The sku deployment mode.
	DeploymentMode *string `pulumi:"deploymentMode"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The parameters for the managed application to be supplied by the vendor.
	ManagedApplicationParameters interface{} `pulumi:"managedApplicationParameters"`
	// The template for the managed application deployment.
	ManagedApplicationTemplate interface{} `pulumi:"managedApplicationTemplate"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The template definition of the network function.
	NetworkFunctionTemplate *NetworkFunctionTemplateResponse `pulumi:"networkFunctionTemplate"`
	// The network function type.
	NetworkFunctionType *string `pulumi:"networkFunctionType"`
	// Indicates if the vendor sku is in preview mode.
	Preview *bool `pulumi:"preview"`
	// The provisioning state of the vendor sku sub resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// The sku type.
	SkuType *string `pulumi:"skuType"`
	// The system meta data relating to this resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupVendorSkusOutput(ctx *pulumi.Context, args LookupVendorSkusOutputArgs, opts ...pulumi.InvokeOption) LookupVendorSkusResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVendorSkusResult, error) {
			args := v.(LookupVendorSkusArgs)
			r, err := LookupVendorSkus(ctx, &args, opts...)
			var s LookupVendorSkusResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVendorSkusResultOutput)
}

type LookupVendorSkusOutputArgs struct {
	// The name of the sku.
	SkuName pulumi.StringInput `pulumi:"skuName"`
	// The name of the vendor.
	VendorName pulumi.StringInput `pulumi:"vendorName"`
}

func (LookupVendorSkusOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVendorSkusArgs)(nil)).Elem()
}

// Sku sub resource.
type LookupVendorSkusResultOutput struct{ *pulumi.OutputState }

func (LookupVendorSkusResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVendorSkusResult)(nil)).Elem()
}

func (o LookupVendorSkusResultOutput) ToLookupVendorSkusResultOutput() LookupVendorSkusResultOutput {
	return o
}

func (o LookupVendorSkusResultOutput) ToLookupVendorSkusResultOutputWithContext(ctx context.Context) LookupVendorSkusResultOutput {
	return o
}

// The sku deployment mode.
func (o LookupVendorSkusResultOutput) DeploymentMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVendorSkusResult) *string { return v.DeploymentMode }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupVendorSkusResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVendorSkusResult) string { return v.Id }).(pulumi.StringOutput)
}

// The parameters for the managed application to be supplied by the vendor.
func (o LookupVendorSkusResultOutput) ManagedApplicationParameters() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupVendorSkusResult) interface{} { return v.ManagedApplicationParameters }).(pulumi.AnyOutput)
}

// The template for the managed application deployment.
func (o LookupVendorSkusResultOutput) ManagedApplicationTemplate() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupVendorSkusResult) interface{} { return v.ManagedApplicationTemplate }).(pulumi.AnyOutput)
}

// The name of the resource
func (o LookupVendorSkusResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVendorSkusResult) string { return v.Name }).(pulumi.StringOutput)
}

// The template definition of the network function.
func (o LookupVendorSkusResultOutput) NetworkFunctionTemplate() NetworkFunctionTemplateResponsePtrOutput {
	return o.ApplyT(func(v LookupVendorSkusResult) *NetworkFunctionTemplateResponse { return v.NetworkFunctionTemplate }).(NetworkFunctionTemplateResponsePtrOutput)
}

// The network function type.
func (o LookupVendorSkusResultOutput) NetworkFunctionType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVendorSkusResult) *string { return v.NetworkFunctionType }).(pulumi.StringPtrOutput)
}

// Indicates if the vendor sku is in preview mode.
func (o LookupVendorSkusResultOutput) Preview() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVendorSkusResult) *bool { return v.Preview }).(pulumi.BoolPtrOutput)
}

// The provisioning state of the vendor sku sub resource.
func (o LookupVendorSkusResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVendorSkusResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The sku type.
func (o LookupVendorSkusResultOutput) SkuType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVendorSkusResult) *string { return v.SkuType }).(pulumi.StringPtrOutput)
}

// The system meta data relating to this resource.
func (o LookupVendorSkusResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupVendorSkusResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupVendorSkusResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVendorSkusResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVendorSkusResultOutput{})
}
