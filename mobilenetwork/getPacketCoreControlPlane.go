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

// Gets information about the specified packet core control plane.
// Azure REST API version: 2023-06-01.
//
// Other available API versions: 2022-03-01-preview, 2022-04-01-preview, 2022-11-01, 2023-09-01.
func LookupPacketCoreControlPlane(ctx *pulumi.Context, args *LookupPacketCoreControlPlaneArgs, opts ...pulumi.InvokeOption) (*LookupPacketCoreControlPlaneResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupPacketCoreControlPlaneResult
	err := ctx.Invoke("azure-native:mobilenetwork:getPacketCoreControlPlane", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupPacketCoreControlPlaneArgs struct {
	// The name of the packet core control plane.
	PacketCoreControlPlaneName string `pulumi:"packetCoreControlPlaneName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Packet core control plane resource.
type LookupPacketCoreControlPlaneResult struct {
	// The control plane interface on the access network. For 5G networks, this is the N2 interface. For 4G networks, this is the S1-MME interface.
	ControlPlaneAccessInterface InterfacePropertiesResponse `pulumi:"controlPlaneAccessInterface"`
	// The core network technology generation (5G core or EPC / 4G core).
	CoreNetworkTechnology *string `pulumi:"coreNetworkTechnology"`
	// Configuration for uploading packet core diagnostics
	DiagnosticsUpload *DiagnosticsUploadConfigurationResponse `pulumi:"diagnosticsUpload"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// The identity used to retrieve the ingress certificate from Azure key vault.
	Identity *ManagedServiceIdentityResponse `pulumi:"identity"`
	// The installation state of the packet core control plane resource.
	Installation *InstallationResponse `pulumi:"installation"`
	// The currently installed version of the packet core software.
	InstalledVersion string `pulumi:"installedVersion"`
	// Settings to allow interoperability with third party components e.g. RANs and UEs.
	InteropSettings interface{} `pulumi:"interopSettings"`
	// The kubernetes ingress configuration to control access to packet core diagnostics over local APIs.
	LocalDiagnosticsAccess LocalDiagnosticsAccessConfigurationResponse `pulumi:"localDiagnosticsAccess"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The platform where the packet core is deployed.
	Platform PlatformConfigurationResponse `pulumi:"platform"`
	// The provisioning state of the packet core control plane resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// The previous version of the packet core software that was deployed. Used when performing the rollback action.
	RollbackVersion string `pulumi:"rollbackVersion"`
	// Site(s) under which this packet core control plane should be deployed. The sites must be in the same location as the packet core control plane.
	Sites []SiteResourceIdResponse `pulumi:"sites"`
	// The SKU defining the throughput and SIM allowances for this packet core control plane deployment.
	Sku string `pulumi:"sku"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// The MTU (in bytes) signaled to the UE. The same MTU is set on the user plane data links for all data networks. The MTU set on the user plane access link is calculated to be 60 bytes greater than this value to allow for GTP encapsulation.
	UeMtu *int `pulumi:"ueMtu"`
	// The desired version of the packet core software.
	Version *string `pulumi:"version"`
}

// Defaults sets the appropriate defaults for LookupPacketCoreControlPlaneResult
func (val *LookupPacketCoreControlPlaneResult) Defaults() *LookupPacketCoreControlPlaneResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Installation = tmp.Installation.Defaults()

	if tmp.UeMtu == nil {
		ueMtu_ := 1440
		tmp.UeMtu = &ueMtu_
	}
	return &tmp
}

func LookupPacketCoreControlPlaneOutput(ctx *pulumi.Context, args LookupPacketCoreControlPlaneOutputArgs, opts ...pulumi.InvokeOption) LookupPacketCoreControlPlaneResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPacketCoreControlPlaneResult, error) {
			args := v.(LookupPacketCoreControlPlaneArgs)
			r, err := LookupPacketCoreControlPlane(ctx, &args, opts...)
			var s LookupPacketCoreControlPlaneResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPacketCoreControlPlaneResultOutput)
}

type LookupPacketCoreControlPlaneOutputArgs struct {
	// The name of the packet core control plane.
	PacketCoreControlPlaneName pulumi.StringInput `pulumi:"packetCoreControlPlaneName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupPacketCoreControlPlaneOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPacketCoreControlPlaneArgs)(nil)).Elem()
}

// Packet core control plane resource.
type LookupPacketCoreControlPlaneResultOutput struct{ *pulumi.OutputState }

func (LookupPacketCoreControlPlaneResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPacketCoreControlPlaneResult)(nil)).Elem()
}

func (o LookupPacketCoreControlPlaneResultOutput) ToLookupPacketCoreControlPlaneResultOutput() LookupPacketCoreControlPlaneResultOutput {
	return o
}

func (o LookupPacketCoreControlPlaneResultOutput) ToLookupPacketCoreControlPlaneResultOutputWithContext(ctx context.Context) LookupPacketCoreControlPlaneResultOutput {
	return o
}

func (o LookupPacketCoreControlPlaneResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupPacketCoreControlPlaneResult] {
	return pulumix.Output[LookupPacketCoreControlPlaneResult]{
		OutputState: o.OutputState,
	}
}

// The control plane interface on the access network. For 5G networks, this is the N2 interface. For 4G networks, this is the S1-MME interface.
func (o LookupPacketCoreControlPlaneResultOutput) ControlPlaneAccessInterface() InterfacePropertiesResponseOutput {
	return o.ApplyT(func(v LookupPacketCoreControlPlaneResult) InterfacePropertiesResponse {
		return v.ControlPlaneAccessInterface
	}).(InterfacePropertiesResponseOutput)
}

// The core network technology generation (5G core or EPC / 4G core).
func (o LookupPacketCoreControlPlaneResultOutput) CoreNetworkTechnology() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPacketCoreControlPlaneResult) *string { return v.CoreNetworkTechnology }).(pulumi.StringPtrOutput)
}

// Configuration for uploading packet core diagnostics
func (o LookupPacketCoreControlPlaneResultOutput) DiagnosticsUpload() DiagnosticsUploadConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupPacketCoreControlPlaneResult) *DiagnosticsUploadConfigurationResponse {
		return v.DiagnosticsUpload
	}).(DiagnosticsUploadConfigurationResponsePtrOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupPacketCoreControlPlaneResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPacketCoreControlPlaneResult) string { return v.Id }).(pulumi.StringOutput)
}

// The identity used to retrieve the ingress certificate from Azure key vault.
func (o LookupPacketCoreControlPlaneResultOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupPacketCoreControlPlaneResult) *ManagedServiceIdentityResponse { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

// The installation state of the packet core control plane resource.
func (o LookupPacketCoreControlPlaneResultOutput) Installation() InstallationResponsePtrOutput {
	return o.ApplyT(func(v LookupPacketCoreControlPlaneResult) *InstallationResponse { return v.Installation }).(InstallationResponsePtrOutput)
}

// The currently installed version of the packet core software.
func (o LookupPacketCoreControlPlaneResultOutput) InstalledVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPacketCoreControlPlaneResult) string { return v.InstalledVersion }).(pulumi.StringOutput)
}

// Settings to allow interoperability with third party components e.g. RANs and UEs.
func (o LookupPacketCoreControlPlaneResultOutput) InteropSettings() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupPacketCoreControlPlaneResult) interface{} { return v.InteropSettings }).(pulumi.AnyOutput)
}

// The kubernetes ingress configuration to control access to packet core diagnostics over local APIs.
func (o LookupPacketCoreControlPlaneResultOutput) LocalDiagnosticsAccess() LocalDiagnosticsAccessConfigurationResponseOutput {
	return o.ApplyT(func(v LookupPacketCoreControlPlaneResult) LocalDiagnosticsAccessConfigurationResponse {
		return v.LocalDiagnosticsAccess
	}).(LocalDiagnosticsAccessConfigurationResponseOutput)
}

// The geo-location where the resource lives
func (o LookupPacketCoreControlPlaneResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPacketCoreControlPlaneResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupPacketCoreControlPlaneResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPacketCoreControlPlaneResult) string { return v.Name }).(pulumi.StringOutput)
}

// The platform where the packet core is deployed.
func (o LookupPacketCoreControlPlaneResultOutput) Platform() PlatformConfigurationResponseOutput {
	return o.ApplyT(func(v LookupPacketCoreControlPlaneResult) PlatformConfigurationResponse { return v.Platform }).(PlatformConfigurationResponseOutput)
}

// The provisioning state of the packet core control plane resource.
func (o LookupPacketCoreControlPlaneResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPacketCoreControlPlaneResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The previous version of the packet core software that was deployed. Used when performing the rollback action.
func (o LookupPacketCoreControlPlaneResultOutput) RollbackVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPacketCoreControlPlaneResult) string { return v.RollbackVersion }).(pulumi.StringOutput)
}

// Site(s) under which this packet core control plane should be deployed. The sites must be in the same location as the packet core control plane.
func (o LookupPacketCoreControlPlaneResultOutput) Sites() SiteResourceIdResponseArrayOutput {
	return o.ApplyT(func(v LookupPacketCoreControlPlaneResult) []SiteResourceIdResponse { return v.Sites }).(SiteResourceIdResponseArrayOutput)
}

// The SKU defining the throughput and SIM allowances for this packet core control plane deployment.
func (o LookupPacketCoreControlPlaneResultOutput) Sku() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPacketCoreControlPlaneResult) string { return v.Sku }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupPacketCoreControlPlaneResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupPacketCoreControlPlaneResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupPacketCoreControlPlaneResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupPacketCoreControlPlaneResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupPacketCoreControlPlaneResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPacketCoreControlPlaneResult) string { return v.Type }).(pulumi.StringOutput)
}

// The MTU (in bytes) signaled to the UE. The same MTU is set on the user plane data links for all data networks. The MTU set on the user plane access link is calculated to be 60 bytes greater than this value to allow for GTP encapsulation.
func (o LookupPacketCoreControlPlaneResultOutput) UeMtu() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupPacketCoreControlPlaneResult) *int { return v.UeMtu }).(pulumi.IntPtrOutput)
}

// The desired version of the packet core software.
func (o LookupPacketCoreControlPlaneResultOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPacketCoreControlPlaneResult) *string { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPacketCoreControlPlaneResultOutput{})
}
