// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets details about the specified dedicated capacity.
func LookupCapacityDetails(ctx *pulumi.Context, args *LookupCapacityDetailsArgs, opts ...pulumi.InvokeOption) (*LookupCapacityDetailsResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupCapacityDetailsResult
	err := ctx.Invoke("azure-native:powerbidedicated/v20210101:getCapacityDetails", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCapacityDetailsArgs struct {
	// The name of the dedicated capacity. It must be a minimum of 3 characters, and a maximum of 63.
	DedicatedCapacityName string `pulumi:"dedicatedCapacityName"`
	// The name of the Azure Resource group of which a given PowerBIDedicated capacity is part. This name must be at least 1 character in length, and no more than 90.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Represents an instance of a Dedicated Capacity resource.
type LookupCapacityDetailsResult struct {
	// A collection of Dedicated capacity administrators
	Administration *DedicatedCapacityAdministratorsResponse `pulumi:"administration"`
	// Capacity name
	FriendlyName string `pulumi:"friendlyName"`
	// An identifier that represents the PowerBI Dedicated resource.
	Id string `pulumi:"id"`
	// Location of the PowerBI Dedicated resource.
	Location string `pulumi:"location"`
	// Specifies the generation of the Power BI Embedded capacity. If no value is specified, the default value 'Gen2' is used. [Learn More](https://docs.microsoft.com/power-bi/developer/embedded/power-bi-embedded-generation-2)
	Mode *string `pulumi:"mode"`
	// The name of the PowerBI Dedicated resource.
	Name string `pulumi:"name"`
	// The current deployment state of PowerBI Dedicated resource. The provisioningState is to indicate states for resource provisioning.
	ProvisioningState string `pulumi:"provisioningState"`
	// The SKU of the PowerBI Dedicated capacity resource.
	Sku CapacitySkuResponse `pulumi:"sku"`
	// The current state of PowerBI Dedicated resource. The state is to indicate more states outside of resource provisioning.
	State string `pulumi:"state"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData *SystemDataResponse `pulumi:"systemData"`
	// Key-value pairs of additional resource provisioning properties.
	Tags map[string]string `pulumi:"tags"`
	// Tenant ID for the capacity. Used for creating Pro Plus capacity.
	TenantId string `pulumi:"tenantId"`
	// The type of the PowerBI Dedicated resource.
	Type string `pulumi:"type"`
}

func LookupCapacityDetailsOutput(ctx *pulumi.Context, args LookupCapacityDetailsOutputArgs, opts ...pulumi.InvokeOption) LookupCapacityDetailsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCapacityDetailsResult, error) {
			args := v.(LookupCapacityDetailsArgs)
			r, err := LookupCapacityDetails(ctx, &args, opts...)
			var s LookupCapacityDetailsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCapacityDetailsResultOutput)
}

type LookupCapacityDetailsOutputArgs struct {
	// The name of the dedicated capacity. It must be a minimum of 3 characters, and a maximum of 63.
	DedicatedCapacityName pulumi.StringInput `pulumi:"dedicatedCapacityName"`
	// The name of the Azure Resource group of which a given PowerBIDedicated capacity is part. This name must be at least 1 character in length, and no more than 90.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupCapacityDetailsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCapacityDetailsArgs)(nil)).Elem()
}

// Represents an instance of a Dedicated Capacity resource.
type LookupCapacityDetailsResultOutput struct{ *pulumi.OutputState }

func (LookupCapacityDetailsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCapacityDetailsResult)(nil)).Elem()
}

func (o LookupCapacityDetailsResultOutput) ToLookupCapacityDetailsResultOutput() LookupCapacityDetailsResultOutput {
	return o
}

func (o LookupCapacityDetailsResultOutput) ToLookupCapacityDetailsResultOutputWithContext(ctx context.Context) LookupCapacityDetailsResultOutput {
	return o
}

func (o LookupCapacityDetailsResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupCapacityDetailsResult] {
	return pulumix.Output[LookupCapacityDetailsResult]{
		OutputState: o.OutputState,
	}
}

// A collection of Dedicated capacity administrators
func (o LookupCapacityDetailsResultOutput) Administration() DedicatedCapacityAdministratorsResponsePtrOutput {
	return o.ApplyT(func(v LookupCapacityDetailsResult) *DedicatedCapacityAdministratorsResponse { return v.Administration }).(DedicatedCapacityAdministratorsResponsePtrOutput)
}

// Capacity name
func (o LookupCapacityDetailsResultOutput) FriendlyName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCapacityDetailsResult) string { return v.FriendlyName }).(pulumi.StringOutput)
}

// An identifier that represents the PowerBI Dedicated resource.
func (o LookupCapacityDetailsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCapacityDetailsResult) string { return v.Id }).(pulumi.StringOutput)
}

// Location of the PowerBI Dedicated resource.
func (o LookupCapacityDetailsResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCapacityDetailsResult) string { return v.Location }).(pulumi.StringOutput)
}

// Specifies the generation of the Power BI Embedded capacity. If no value is specified, the default value 'Gen2' is used. [Learn More](https://docs.microsoft.com/power-bi/developer/embedded/power-bi-embedded-generation-2)
func (o LookupCapacityDetailsResultOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCapacityDetailsResult) *string { return v.Mode }).(pulumi.StringPtrOutput)
}

// The name of the PowerBI Dedicated resource.
func (o LookupCapacityDetailsResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCapacityDetailsResult) string { return v.Name }).(pulumi.StringOutput)
}

// The current deployment state of PowerBI Dedicated resource. The provisioningState is to indicate states for resource provisioning.
func (o LookupCapacityDetailsResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCapacityDetailsResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The SKU of the PowerBI Dedicated capacity resource.
func (o LookupCapacityDetailsResultOutput) Sku() CapacitySkuResponseOutput {
	return o.ApplyT(func(v LookupCapacityDetailsResult) CapacitySkuResponse { return v.Sku }).(CapacitySkuResponseOutput)
}

// The current state of PowerBI Dedicated resource. The state is to indicate more states outside of resource provisioning.
func (o LookupCapacityDetailsResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCapacityDetailsResult) string { return v.State }).(pulumi.StringOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupCapacityDetailsResultOutput) SystemData() SystemDataResponsePtrOutput {
	return o.ApplyT(func(v LookupCapacityDetailsResult) *SystemDataResponse { return v.SystemData }).(SystemDataResponsePtrOutput)
}

// Key-value pairs of additional resource provisioning properties.
func (o LookupCapacityDetailsResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupCapacityDetailsResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Tenant ID for the capacity. Used for creating Pro Plus capacity.
func (o LookupCapacityDetailsResultOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCapacityDetailsResult) string { return v.TenantId }).(pulumi.StringOutput)
}

// The type of the PowerBI Dedicated resource.
func (o LookupCapacityDetailsResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCapacityDetailsResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCapacityDetailsResultOutput{})
}
