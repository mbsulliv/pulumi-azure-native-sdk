// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20231001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the SAP Instance resource.
func LookupSapInstance(ctx *pulumi.Context, args *LookupSapInstanceArgs, opts ...pulumi.InvokeOption) (*LookupSapInstanceResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupSapInstanceResult
	err := ctx.Invoke("azure-native:workloads/v20231001preview:getSapInstance", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSapInstanceArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the discovery site resource for SAP Migration.
	SapDiscoverySiteName string `pulumi:"sapDiscoverySiteName"`
	// The name of SAP Instance resource for SAP Migration.
	SapInstanceName string `pulumi:"sapInstanceName"`
}

// Define the SAP Instance resource.
type LookupSapInstanceResult struct {
	// Enter a business function/department identifier to group multiple SIDs.
	Application string `pulumi:"application"`
	// The Environment; PRD, QA, DEV, etc to which SAP system belongs to. Select from the list of available dropdown values.
	Environment string `pulumi:"environment"`
	// Defines the errors related to SAP Instance resource.
	Errors SAPMigrateErrorResponse `pulumi:"errors"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// This is the SID of the production system in a landscape.  An SAP system could itself be a production SID or a part of a landscape with a different Production SID. This field can be used to relate non-prod SIDs, other components, SID (WEBDISP) to the prod SID. Enter the value of Production SID.
	LandscapeSid string `pulumi:"landscapeSid"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Defines the provisioning states.
	ProvisioningState string `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// This is the SID of SAP System. Keeping this not equal to ID as different landscapes can have repeated System SID IDs.
	SystemSid string `pulumi:"systemSid"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupSapInstanceOutput(ctx *pulumi.Context, args LookupSapInstanceOutputArgs, opts ...pulumi.InvokeOption) LookupSapInstanceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSapInstanceResult, error) {
			args := v.(LookupSapInstanceArgs)
			r, err := LookupSapInstance(ctx, &args, opts...)
			var s LookupSapInstanceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSapInstanceResultOutput)
}

type LookupSapInstanceOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the discovery site resource for SAP Migration.
	SapDiscoverySiteName pulumi.StringInput `pulumi:"sapDiscoverySiteName"`
	// The name of SAP Instance resource for SAP Migration.
	SapInstanceName pulumi.StringInput `pulumi:"sapInstanceName"`
}

func (LookupSapInstanceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSapInstanceArgs)(nil)).Elem()
}

// Define the SAP Instance resource.
type LookupSapInstanceResultOutput struct{ *pulumi.OutputState }

func (LookupSapInstanceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSapInstanceResult)(nil)).Elem()
}

func (o LookupSapInstanceResultOutput) ToLookupSapInstanceResultOutput() LookupSapInstanceResultOutput {
	return o
}

func (o LookupSapInstanceResultOutput) ToLookupSapInstanceResultOutputWithContext(ctx context.Context) LookupSapInstanceResultOutput {
	return o
}

// Enter a business function/department identifier to group multiple SIDs.
func (o LookupSapInstanceResultOutput) Application() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSapInstanceResult) string { return v.Application }).(pulumi.StringOutput)
}

// The Environment; PRD, QA, DEV, etc to which SAP system belongs to. Select from the list of available dropdown values.
func (o LookupSapInstanceResultOutput) Environment() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSapInstanceResult) string { return v.Environment }).(pulumi.StringOutput)
}

// Defines the errors related to SAP Instance resource.
func (o LookupSapInstanceResultOutput) Errors() SAPMigrateErrorResponseOutput {
	return o.ApplyT(func(v LookupSapInstanceResult) SAPMigrateErrorResponse { return v.Errors }).(SAPMigrateErrorResponseOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupSapInstanceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSapInstanceResult) string { return v.Id }).(pulumi.StringOutput)
}

// This is the SID of the production system in a landscape.  An SAP system could itself be a production SID or a part of a landscape with a different Production SID. This field can be used to relate non-prod SIDs, other components, SID (WEBDISP) to the prod SID. Enter the value of Production SID.
func (o LookupSapInstanceResultOutput) LandscapeSid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSapInstanceResult) string { return v.LandscapeSid }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupSapInstanceResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSapInstanceResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupSapInstanceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSapInstanceResult) string { return v.Name }).(pulumi.StringOutput)
}

// Defines the provisioning states.
func (o LookupSapInstanceResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSapInstanceResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupSapInstanceResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupSapInstanceResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// This is the SID of SAP System. Keeping this not equal to ID as different landscapes can have repeated System SID IDs.
func (o LookupSapInstanceResultOutput) SystemSid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSapInstanceResult) string { return v.SystemSid }).(pulumi.StringOutput)
}

// Resource tags.
func (o LookupSapInstanceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSapInstanceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupSapInstanceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSapInstanceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSapInstanceResultOutput{})
}
