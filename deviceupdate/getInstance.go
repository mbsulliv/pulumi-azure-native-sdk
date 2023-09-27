// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package deviceupdate

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Returns instance details for the given instance and account name.
// Azure REST API version: 2023-07-01.
func LookupInstance(ctx *pulumi.Context, args *LookupInstanceArgs, opts ...pulumi.InvokeOption) (*LookupInstanceResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupInstanceResult
	err := ctx.Invoke("azure-native:deviceupdate:getInstance", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupInstanceArgs struct {
	// Account name.
	AccountName string `pulumi:"accountName"`
	// Instance name.
	InstanceName string `pulumi:"instanceName"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Device Update instance details.
type LookupInstanceResult struct {
	// Parent Device Update Account name which Instance belongs to.
	AccountName string `pulumi:"accountName"`
	// Customer-initiated diagnostic log collection storage properties
	DiagnosticStorageProperties *DiagnosticStoragePropertiesResponse `pulumi:"diagnosticStorageProperties"`
	// Enables or Disables the diagnostic logs collection
	EnableDiagnostics *bool `pulumi:"enableDiagnostics"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// List of IoT Hubs associated with the account.
	IotHubs []IotHubSettingsResponse `pulumi:"iotHubs"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Provisioning state.
	ProvisioningState string `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupInstanceOutput(ctx *pulumi.Context, args LookupInstanceOutputArgs, opts ...pulumi.InvokeOption) LookupInstanceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupInstanceResult, error) {
			args := v.(LookupInstanceArgs)
			r, err := LookupInstance(ctx, &args, opts...)
			var s LookupInstanceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupInstanceResultOutput)
}

type LookupInstanceOutputArgs struct {
	// Account name.
	AccountName pulumi.StringInput `pulumi:"accountName"`
	// Instance name.
	InstanceName pulumi.StringInput `pulumi:"instanceName"`
	// The resource group name.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupInstanceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInstanceArgs)(nil)).Elem()
}

// Device Update instance details.
type LookupInstanceResultOutput struct{ *pulumi.OutputState }

func (LookupInstanceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInstanceResult)(nil)).Elem()
}

func (o LookupInstanceResultOutput) ToLookupInstanceResultOutput() LookupInstanceResultOutput {
	return o
}

func (o LookupInstanceResultOutput) ToLookupInstanceResultOutputWithContext(ctx context.Context) LookupInstanceResultOutput {
	return o
}

func (o LookupInstanceResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupInstanceResult] {
	return pulumix.Output[LookupInstanceResult]{
		OutputState: o.OutputState,
	}
}

// Parent Device Update Account name which Instance belongs to.
func (o LookupInstanceResultOutput) AccountName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.AccountName }).(pulumi.StringOutput)
}

// Customer-initiated diagnostic log collection storage properties
func (o LookupInstanceResultOutput) DiagnosticStorageProperties() DiagnosticStoragePropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupInstanceResult) *DiagnosticStoragePropertiesResponse {
		return v.DiagnosticStorageProperties
	}).(DiagnosticStoragePropertiesResponsePtrOutput)
}

// Enables or Disables the diagnostic logs collection
func (o LookupInstanceResultOutput) EnableDiagnostics() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupInstanceResult) *bool { return v.EnableDiagnostics }).(pulumi.BoolPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupInstanceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.Id }).(pulumi.StringOutput)
}

// List of IoT Hubs associated with the account.
func (o LookupInstanceResultOutput) IotHubs() IotHubSettingsResponseArrayOutput {
	return o.ApplyT(func(v LookupInstanceResult) []IotHubSettingsResponse { return v.IotHubs }).(IotHubSettingsResponseArrayOutput)
}

// The geo-location where the resource lives
func (o LookupInstanceResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupInstanceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.Name }).(pulumi.StringOutput)
}

// Provisioning state.
func (o LookupInstanceResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupInstanceResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupInstanceResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupInstanceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupInstanceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupInstanceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupInstanceResultOutput{})
}
