// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220521preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Implements Cloud GET method.
func LookupCloud(ctx *pulumi.Context, args *LookupCloudArgs, opts ...pulumi.InvokeOption) (*LookupCloudResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupCloudResult
	err := ctx.Invoke("azure-native:scvmm/v20220521preview:getCloud", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCloudArgs struct {
	// Name of the Cloud.
	CloudName string `pulumi:"cloudName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The Clouds resource definition.
type LookupCloudResult struct {
	// Capacity of the cloud.
	CloudCapacity CloudCapacityResponse `pulumi:"cloudCapacity"`
	// Name of the cloud in VMMServer.
	CloudName string `pulumi:"cloudName"`
	// The extended location.
	ExtendedLocation ExtendedLocationResponse `pulumi:"extendedLocation"`
	// Resource Id
	Id string `pulumi:"id"`
	// Gets or sets the inventory Item ID for the resource.
	InventoryItemId *string `pulumi:"inventoryItemId"`
	// Gets or sets the location.
	Location string `pulumi:"location"`
	// Resource Name
	Name string `pulumi:"name"`
	// Gets or sets the provisioning state.
	ProvisioningState string `pulumi:"provisioningState"`
	// List of QoS policies available for the cloud.
	StorageQoSPolicies []StorageQoSPolicyResponse `pulumi:"storageQoSPolicies"`
	// The system data.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource Type
	Type string `pulumi:"type"`
	// Unique ID of the cloud.
	Uuid *string `pulumi:"uuid"`
	// ARM Id of the vmmServer resource in which this resource resides.
	VmmServerId *string `pulumi:"vmmServerId"`
}

func LookupCloudOutput(ctx *pulumi.Context, args LookupCloudOutputArgs, opts ...pulumi.InvokeOption) LookupCloudResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCloudResult, error) {
			args := v.(LookupCloudArgs)
			r, err := LookupCloud(ctx, &args, opts...)
			var s LookupCloudResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCloudResultOutput)
}

type LookupCloudOutputArgs struct {
	// Name of the Cloud.
	CloudName pulumi.StringInput `pulumi:"cloudName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupCloudOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCloudArgs)(nil)).Elem()
}

// The Clouds resource definition.
type LookupCloudResultOutput struct{ *pulumi.OutputState }

func (LookupCloudResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCloudResult)(nil)).Elem()
}

func (o LookupCloudResultOutput) ToLookupCloudResultOutput() LookupCloudResultOutput {
	return o
}

func (o LookupCloudResultOutput) ToLookupCloudResultOutputWithContext(ctx context.Context) LookupCloudResultOutput {
	return o
}

func (o LookupCloudResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupCloudResult] {
	return pulumix.Output[LookupCloudResult]{
		OutputState: o.OutputState,
	}
}

// Capacity of the cloud.
func (o LookupCloudResultOutput) CloudCapacity() CloudCapacityResponseOutput {
	return o.ApplyT(func(v LookupCloudResult) CloudCapacityResponse { return v.CloudCapacity }).(CloudCapacityResponseOutput)
}

// Name of the cloud in VMMServer.
func (o LookupCloudResultOutput) CloudName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudResult) string { return v.CloudName }).(pulumi.StringOutput)
}

// The extended location.
func (o LookupCloudResultOutput) ExtendedLocation() ExtendedLocationResponseOutput {
	return o.ApplyT(func(v LookupCloudResult) ExtendedLocationResponse { return v.ExtendedLocation }).(ExtendedLocationResponseOutput)
}

// Resource Id
func (o LookupCloudResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudResult) string { return v.Id }).(pulumi.StringOutput)
}

// Gets or sets the inventory Item ID for the resource.
func (o LookupCloudResultOutput) InventoryItemId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCloudResult) *string { return v.InventoryItemId }).(pulumi.StringPtrOutput)
}

// Gets or sets the location.
func (o LookupCloudResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudResult) string { return v.Location }).(pulumi.StringOutput)
}

// Resource Name
func (o LookupCloudResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudResult) string { return v.Name }).(pulumi.StringOutput)
}

// Gets or sets the provisioning state.
func (o LookupCloudResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// List of QoS policies available for the cloud.
func (o LookupCloudResultOutput) StorageQoSPolicies() StorageQoSPolicyResponseArrayOutput {
	return o.ApplyT(func(v LookupCloudResult) []StorageQoSPolicyResponse { return v.StorageQoSPolicies }).(StorageQoSPolicyResponseArrayOutput)
}

// The system data.
func (o LookupCloudResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupCloudResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags
func (o LookupCloudResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupCloudResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource Type
func (o LookupCloudResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudResult) string { return v.Type }).(pulumi.StringOutput)
}

// Unique ID of the cloud.
func (o LookupCloudResultOutput) Uuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCloudResult) *string { return v.Uuid }).(pulumi.StringPtrOutput)
}

// ARM Id of the vmmServer resource in which this resource resides.
func (o LookupCloudResultOutput) VmmServerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCloudResult) *string { return v.VmmServerId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCloudResultOutput{})
}
