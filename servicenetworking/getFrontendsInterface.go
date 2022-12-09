// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package servicenetworking

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Frontend Subresource of Traffic Controller.
// API Version: 2022-10-01-preview.
func LookupFrontendsInterface(ctx *pulumi.Context, args *LookupFrontendsInterfaceArgs, opts ...pulumi.InvokeOption) (*LookupFrontendsInterfaceResult, error) {
	var rv LookupFrontendsInterfaceResult
	err := ctx.Invoke("azure-native:servicenetworking:getFrontendsInterface", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFrontendsInterfaceArgs struct {
	// Frontends
	FrontendName string `pulumi:"frontendName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// traffic controller name for path
	TrafficControllerName string `pulumi:"trafficControllerName"`
}

// Frontend Subresource of Traffic Controller.
type LookupFrontendsInterfaceResult struct {
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// Frontend IP Address Version (Optional).
	IpAddressVersion *string `pulumi:"ipAddressVersion"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// Frontend Mode (Optional).
	Mode *string `pulumi:"mode"`
	// The name of the resource
	Name string `pulumi:"name"`
	// test doc
	ProvisioningState string `pulumi:"provisioningState"`
	// Frontend Public IP Address (Optional).
	PublicIPAddress *FrontendPropertiesIPAddressResponse `pulumi:"publicIPAddress"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupFrontendsInterfaceOutput(ctx *pulumi.Context, args LookupFrontendsInterfaceOutputArgs, opts ...pulumi.InvokeOption) LookupFrontendsInterfaceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFrontendsInterfaceResult, error) {
			args := v.(LookupFrontendsInterfaceArgs)
			r, err := LookupFrontendsInterface(ctx, &args, opts...)
			var s LookupFrontendsInterfaceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFrontendsInterfaceResultOutput)
}

type LookupFrontendsInterfaceOutputArgs struct {
	// Frontends
	FrontendName pulumi.StringInput `pulumi:"frontendName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// traffic controller name for path
	TrafficControllerName pulumi.StringInput `pulumi:"trafficControllerName"`
}

func (LookupFrontendsInterfaceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFrontendsInterfaceArgs)(nil)).Elem()
}

// Frontend Subresource of Traffic Controller.
type LookupFrontendsInterfaceResultOutput struct{ *pulumi.OutputState }

func (LookupFrontendsInterfaceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFrontendsInterfaceResult)(nil)).Elem()
}

func (o LookupFrontendsInterfaceResultOutput) ToLookupFrontendsInterfaceResultOutput() LookupFrontendsInterfaceResultOutput {
	return o
}

func (o LookupFrontendsInterfaceResultOutput) ToLookupFrontendsInterfaceResultOutputWithContext(ctx context.Context) LookupFrontendsInterfaceResultOutput {
	return o
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupFrontendsInterfaceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFrontendsInterfaceResult) string { return v.Id }).(pulumi.StringOutput)
}

// Frontend IP Address Version (Optional).
func (o LookupFrontendsInterfaceResultOutput) IpAddressVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFrontendsInterfaceResult) *string { return v.IpAddressVersion }).(pulumi.StringPtrOutput)
}

// The geo-location where the resource lives
func (o LookupFrontendsInterfaceResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFrontendsInterfaceResult) string { return v.Location }).(pulumi.StringOutput)
}

// Frontend Mode (Optional).
func (o LookupFrontendsInterfaceResultOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFrontendsInterfaceResult) *string { return v.Mode }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o LookupFrontendsInterfaceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFrontendsInterfaceResult) string { return v.Name }).(pulumi.StringOutput)
}

// test doc
func (o LookupFrontendsInterfaceResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFrontendsInterfaceResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Frontend Public IP Address (Optional).
func (o LookupFrontendsInterfaceResultOutput) PublicIPAddress() FrontendPropertiesIPAddressResponsePtrOutput {
	return o.ApplyT(func(v LookupFrontendsInterfaceResult) *FrontendPropertiesIPAddressResponse { return v.PublicIPAddress }).(FrontendPropertiesIPAddressResponsePtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupFrontendsInterfaceResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupFrontendsInterfaceResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupFrontendsInterfaceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupFrontendsInterfaceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupFrontendsInterfaceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFrontendsInterfaceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFrontendsInterfaceResultOutput{})
}
