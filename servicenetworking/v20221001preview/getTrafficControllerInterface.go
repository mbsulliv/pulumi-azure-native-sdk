// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20221001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get a TrafficController
func LookupTrafficControllerInterface(ctx *pulumi.Context, args *LookupTrafficControllerInterfaceArgs, opts ...pulumi.InvokeOption) (*LookupTrafficControllerInterfaceResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupTrafficControllerInterfaceResult
	err := ctx.Invoke("azure-native:servicenetworking/v20221001preview:getTrafficControllerInterface", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTrafficControllerInterfaceArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// traffic controller name for path
	TrafficControllerName string `pulumi:"trafficControllerName"`
}

// Concrete tracked resource types can be created by aliasing this type using a specific property type.
type LookupTrafficControllerInterfaceResult struct {
	// Associations References List
	Associations []ResourceIDResponse `pulumi:"associations"`
	// Configuration Endpoints.
	ConfigurationEndpoints []string `pulumi:"configurationEndpoints"`
	// Frontends References List
	Frontends []ResourceIDResponse `pulumi:"frontends"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The status of the last operation.
	ProvisioningState string `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupTrafficControllerInterfaceOutput(ctx *pulumi.Context, args LookupTrafficControllerInterfaceOutputArgs, opts ...pulumi.InvokeOption) LookupTrafficControllerInterfaceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTrafficControllerInterfaceResult, error) {
			args := v.(LookupTrafficControllerInterfaceArgs)
			r, err := LookupTrafficControllerInterface(ctx, &args, opts...)
			var s LookupTrafficControllerInterfaceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTrafficControllerInterfaceResultOutput)
}

type LookupTrafficControllerInterfaceOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// traffic controller name for path
	TrafficControllerName pulumi.StringInput `pulumi:"trafficControllerName"`
}

func (LookupTrafficControllerInterfaceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTrafficControllerInterfaceArgs)(nil)).Elem()
}

// Concrete tracked resource types can be created by aliasing this type using a specific property type.
type LookupTrafficControllerInterfaceResultOutput struct{ *pulumi.OutputState }

func (LookupTrafficControllerInterfaceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTrafficControllerInterfaceResult)(nil)).Elem()
}

func (o LookupTrafficControllerInterfaceResultOutput) ToLookupTrafficControllerInterfaceResultOutput() LookupTrafficControllerInterfaceResultOutput {
	return o
}

func (o LookupTrafficControllerInterfaceResultOutput) ToLookupTrafficControllerInterfaceResultOutputWithContext(ctx context.Context) LookupTrafficControllerInterfaceResultOutput {
	return o
}

// Associations References List
func (o LookupTrafficControllerInterfaceResultOutput) Associations() ResourceIDResponseArrayOutput {
	return o.ApplyT(func(v LookupTrafficControllerInterfaceResult) []ResourceIDResponse { return v.Associations }).(ResourceIDResponseArrayOutput)
}

// Configuration Endpoints.
func (o LookupTrafficControllerInterfaceResultOutput) ConfigurationEndpoints() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupTrafficControllerInterfaceResult) []string { return v.ConfigurationEndpoints }).(pulumi.StringArrayOutput)
}

// Frontends References List
func (o LookupTrafficControllerInterfaceResultOutput) Frontends() ResourceIDResponseArrayOutput {
	return o.ApplyT(func(v LookupTrafficControllerInterfaceResult) []ResourceIDResponse { return v.Frontends }).(ResourceIDResponseArrayOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupTrafficControllerInterfaceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTrafficControllerInterfaceResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupTrafficControllerInterfaceResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTrafficControllerInterfaceResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupTrafficControllerInterfaceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTrafficControllerInterfaceResult) string { return v.Name }).(pulumi.StringOutput)
}

// The status of the last operation.
func (o LookupTrafficControllerInterfaceResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTrafficControllerInterfaceResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupTrafficControllerInterfaceResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupTrafficControllerInterfaceResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupTrafficControllerInterfaceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupTrafficControllerInterfaceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupTrafficControllerInterfaceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTrafficControllerInterfaceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTrafficControllerInterfaceResultOutput{})
}
