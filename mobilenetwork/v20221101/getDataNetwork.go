// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20221101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Data network resource. Must be created in the same location as its parent mobile network.
func LookupDataNetwork(ctx *pulumi.Context, args *LookupDataNetworkArgs, opts ...pulumi.InvokeOption) (*LookupDataNetworkResult, error) {
	var rv LookupDataNetworkResult
	err := ctx.Invoke("azure-native:mobilenetwork/v20221101:getDataNetwork", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDataNetworkArgs struct {
	// The name of the data network.
	DataNetworkName string `pulumi:"dataNetworkName"`
	// The name of the mobile network.
	MobileNetworkName string `pulumi:"mobileNetworkName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Data network resource. Must be created in the same location as its parent mobile network.
type LookupDataNetworkResult struct {
	// An optional description for this data network.
	Description *string `pulumi:"description"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The provisioning state of the data network resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupDataNetworkOutput(ctx *pulumi.Context, args LookupDataNetworkOutputArgs, opts ...pulumi.InvokeOption) LookupDataNetworkResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDataNetworkResult, error) {
			args := v.(LookupDataNetworkArgs)
			r, err := LookupDataNetwork(ctx, &args, opts...)
			var s LookupDataNetworkResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDataNetworkResultOutput)
}

type LookupDataNetworkOutputArgs struct {
	// The name of the data network.
	DataNetworkName pulumi.StringInput `pulumi:"dataNetworkName"`
	// The name of the mobile network.
	MobileNetworkName pulumi.StringInput `pulumi:"mobileNetworkName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupDataNetworkOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDataNetworkArgs)(nil)).Elem()
}

// Data network resource. Must be created in the same location as its parent mobile network.
type LookupDataNetworkResultOutput struct{ *pulumi.OutputState }

func (LookupDataNetworkResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDataNetworkResult)(nil)).Elem()
}

func (o LookupDataNetworkResultOutput) ToLookupDataNetworkResultOutput() LookupDataNetworkResultOutput {
	return o
}

func (o LookupDataNetworkResultOutput) ToLookupDataNetworkResultOutputWithContext(ctx context.Context) LookupDataNetworkResultOutput {
	return o
}

// An optional description for this data network.
func (o LookupDataNetworkResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDataNetworkResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupDataNetworkResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataNetworkResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupDataNetworkResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataNetworkResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupDataNetworkResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataNetworkResult) string { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the data network resource.
func (o LookupDataNetworkResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataNetworkResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupDataNetworkResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupDataNetworkResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupDataNetworkResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDataNetworkResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupDataNetworkResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataNetworkResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDataNetworkResultOutput{})
}
