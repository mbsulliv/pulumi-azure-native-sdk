// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20211201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Define the SAP Application Server Instance resource.
func LookupSAPApplicationServerInstance(ctx *pulumi.Context, args *LookupSAPApplicationServerInstanceArgs, opts ...pulumi.InvokeOption) (*LookupSAPApplicationServerInstanceResult, error) {
	var rv LookupSAPApplicationServerInstanceResult
	err := ctx.Invoke("azure-native:workloads/v20211201preview:getSAPApplicationServerInstance", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSAPApplicationServerInstanceArgs struct {
	// The name of SAP Application Server instance resource.
	ApplicationInstanceName string `pulumi:"applicationInstanceName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Virtual Instances for SAP solutions resource
	SapVirtualInstanceName string `pulumi:"sapVirtualInstanceName"`
}

// Define the SAP Application Server Instance resource.
type LookupSAPApplicationServerInstanceResult struct {
	// Defines the Application Instance errors.
	Errors SAPVirtualInstanceErrorResponse `pulumi:"errors"`
	// Application server instance gateway Port.
	GatewayPort float64 `pulumi:"gatewayPort"`
	// Defines the health of SAP Instances.
	Health string `pulumi:"health"`
	// Application server instance SAP hostname.
	Hostname string `pulumi:"hostname"`
	// Application server instance ICM HTTP Port.
	IcmHttpPort float64 `pulumi:"icmHttpPort"`
	// Application server instance ICM HTTPS Port.
	IcmHttpsPort float64 `pulumi:"icmHttpsPort"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// Application server Instance Number.
	InstanceNo string `pulumi:"instanceNo"`
	//  Application server instance SAP IP Address.
	IpAddress string `pulumi:"ipAddress"`
	// Application server instance SAP Kernel Patch level.
	KernelPatch string `pulumi:"kernelPatch"`
	//  Application server instance SAP Kernel Version.
	KernelVersion string `pulumi:"kernelVersion"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Defines the provisioning states.
	ProvisioningState string `pulumi:"provisioningState"`
	// Defines the SAP Instance status.
	Status string `pulumi:"status"`
	// Storage details of all the Storage Accounts attached to the App Virtual Machine. For e.g. NFS on AFS Shared Storage.
	StorageDetails []StorageInformationResponse `pulumi:"storageDetails"`
	// Application server Subnet.
	Subnet string `pulumi:"subnet"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// The virtual machine.
	VirtualMachineId string `pulumi:"virtualMachineId"`
}

func LookupSAPApplicationServerInstanceOutput(ctx *pulumi.Context, args LookupSAPApplicationServerInstanceOutputArgs, opts ...pulumi.InvokeOption) LookupSAPApplicationServerInstanceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSAPApplicationServerInstanceResult, error) {
			args := v.(LookupSAPApplicationServerInstanceArgs)
			r, err := LookupSAPApplicationServerInstance(ctx, &args, opts...)
			var s LookupSAPApplicationServerInstanceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSAPApplicationServerInstanceResultOutput)
}

type LookupSAPApplicationServerInstanceOutputArgs struct {
	// The name of SAP Application Server instance resource.
	ApplicationInstanceName pulumi.StringInput `pulumi:"applicationInstanceName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the Virtual Instances for SAP solutions resource
	SapVirtualInstanceName pulumi.StringInput `pulumi:"sapVirtualInstanceName"`
}

func (LookupSAPApplicationServerInstanceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSAPApplicationServerInstanceArgs)(nil)).Elem()
}

// Define the SAP Application Server Instance resource.
type LookupSAPApplicationServerInstanceResultOutput struct{ *pulumi.OutputState }

func (LookupSAPApplicationServerInstanceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSAPApplicationServerInstanceResult)(nil)).Elem()
}

func (o LookupSAPApplicationServerInstanceResultOutput) ToLookupSAPApplicationServerInstanceResultOutput() LookupSAPApplicationServerInstanceResultOutput {
	return o
}

func (o LookupSAPApplicationServerInstanceResultOutput) ToLookupSAPApplicationServerInstanceResultOutputWithContext(ctx context.Context) LookupSAPApplicationServerInstanceResultOutput {
	return o
}

// Defines the Application Instance errors.
func (o LookupSAPApplicationServerInstanceResultOutput) Errors() SAPVirtualInstanceErrorResponseOutput {
	return o.ApplyT(func(v LookupSAPApplicationServerInstanceResult) SAPVirtualInstanceErrorResponse { return v.Errors }).(SAPVirtualInstanceErrorResponseOutput)
}

// Application server instance gateway Port.
func (o LookupSAPApplicationServerInstanceResultOutput) GatewayPort() pulumi.Float64Output {
	return o.ApplyT(func(v LookupSAPApplicationServerInstanceResult) float64 { return v.GatewayPort }).(pulumi.Float64Output)
}

// Defines the health of SAP Instances.
func (o LookupSAPApplicationServerInstanceResultOutput) Health() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSAPApplicationServerInstanceResult) string { return v.Health }).(pulumi.StringOutput)
}

// Application server instance SAP hostname.
func (o LookupSAPApplicationServerInstanceResultOutput) Hostname() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSAPApplicationServerInstanceResult) string { return v.Hostname }).(pulumi.StringOutput)
}

// Application server instance ICM HTTP Port.
func (o LookupSAPApplicationServerInstanceResultOutput) IcmHttpPort() pulumi.Float64Output {
	return o.ApplyT(func(v LookupSAPApplicationServerInstanceResult) float64 { return v.IcmHttpPort }).(pulumi.Float64Output)
}

// Application server instance ICM HTTPS Port.
func (o LookupSAPApplicationServerInstanceResultOutput) IcmHttpsPort() pulumi.Float64Output {
	return o.ApplyT(func(v LookupSAPApplicationServerInstanceResult) float64 { return v.IcmHttpsPort }).(pulumi.Float64Output)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupSAPApplicationServerInstanceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSAPApplicationServerInstanceResult) string { return v.Id }).(pulumi.StringOutput)
}

// Application server Instance Number.
func (o LookupSAPApplicationServerInstanceResultOutput) InstanceNo() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSAPApplicationServerInstanceResult) string { return v.InstanceNo }).(pulumi.StringOutput)
}

// Application server instance SAP IP Address.
func (o LookupSAPApplicationServerInstanceResultOutput) IpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSAPApplicationServerInstanceResult) string { return v.IpAddress }).(pulumi.StringOutput)
}

// Application server instance SAP Kernel Patch level.
func (o LookupSAPApplicationServerInstanceResultOutput) KernelPatch() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSAPApplicationServerInstanceResult) string { return v.KernelPatch }).(pulumi.StringOutput)
}

// Application server instance SAP Kernel Version.
func (o LookupSAPApplicationServerInstanceResultOutput) KernelVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSAPApplicationServerInstanceResult) string { return v.KernelVersion }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupSAPApplicationServerInstanceResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSAPApplicationServerInstanceResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupSAPApplicationServerInstanceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSAPApplicationServerInstanceResult) string { return v.Name }).(pulumi.StringOutput)
}

// Defines the provisioning states.
func (o LookupSAPApplicationServerInstanceResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSAPApplicationServerInstanceResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Defines the SAP Instance status.
func (o LookupSAPApplicationServerInstanceResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSAPApplicationServerInstanceResult) string { return v.Status }).(pulumi.StringOutput)
}

// Storage details of all the Storage Accounts attached to the App Virtual Machine. For e.g. NFS on AFS Shared Storage.
func (o LookupSAPApplicationServerInstanceResultOutput) StorageDetails() StorageInformationResponseArrayOutput {
	return o.ApplyT(func(v LookupSAPApplicationServerInstanceResult) []StorageInformationResponse { return v.StorageDetails }).(StorageInformationResponseArrayOutput)
}

// Application server Subnet.
func (o LookupSAPApplicationServerInstanceResultOutput) Subnet() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSAPApplicationServerInstanceResult) string { return v.Subnet }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupSAPApplicationServerInstanceResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupSAPApplicationServerInstanceResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupSAPApplicationServerInstanceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSAPApplicationServerInstanceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupSAPApplicationServerInstanceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSAPApplicationServerInstanceResult) string { return v.Type }).(pulumi.StringOutput)
}

// The virtual machine.
func (o LookupSAPApplicationServerInstanceResultOutput) VirtualMachineId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSAPApplicationServerInstanceResult) string { return v.VirtualMachineId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSAPApplicationServerInstanceResultOutput{})
}
