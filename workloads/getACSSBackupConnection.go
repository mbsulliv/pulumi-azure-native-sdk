// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package workloads

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets the backup connection resource of virtual instance for SAP.
// Azure REST API version: 2023-10-01-preview.
func LookupACSSBackupConnection(ctx *pulumi.Context, args *LookupACSSBackupConnectionArgs, opts ...pulumi.InvokeOption) (*LookupACSSBackupConnectionResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupACSSBackupConnectionResult
	err := ctx.Invoke("azure-native:workloads:getACSSBackupConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupACSSBackupConnectionArgs struct {
	// The name of the backup connection resource of virtual instance for SAP.
	BackupName string `pulumi:"backupName"`
	// The name of the connector resource
	ConnectorName string `pulumi:"connectorName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Define the backup connection resource of virtual instance for SAP..
type LookupACSSBackupConnectionResult struct {
	// Information about the recovery services vault and backup policy used for backup.
	BackupData interface{} `pulumi:"backupData"`
	// Defines the errors related to backup connection resource of virtual instance for SAP.
	Errors ConnectorErrorDefinitionResponse `pulumi:"errors"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Defines the provisioning states.
	ProvisioningState string `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupACSSBackupConnectionOutput(ctx *pulumi.Context, args LookupACSSBackupConnectionOutputArgs, opts ...pulumi.InvokeOption) LookupACSSBackupConnectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupACSSBackupConnectionResult, error) {
			args := v.(LookupACSSBackupConnectionArgs)
			r, err := LookupACSSBackupConnection(ctx, &args, opts...)
			var s LookupACSSBackupConnectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupACSSBackupConnectionResultOutput)
}

type LookupACSSBackupConnectionOutputArgs struct {
	// The name of the backup connection resource of virtual instance for SAP.
	BackupName pulumi.StringInput `pulumi:"backupName"`
	// The name of the connector resource
	ConnectorName pulumi.StringInput `pulumi:"connectorName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupACSSBackupConnectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupACSSBackupConnectionArgs)(nil)).Elem()
}

// Define the backup connection resource of virtual instance for SAP..
type LookupACSSBackupConnectionResultOutput struct{ *pulumi.OutputState }

func (LookupACSSBackupConnectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupACSSBackupConnectionResult)(nil)).Elem()
}

func (o LookupACSSBackupConnectionResultOutput) ToLookupACSSBackupConnectionResultOutput() LookupACSSBackupConnectionResultOutput {
	return o
}

func (o LookupACSSBackupConnectionResultOutput) ToLookupACSSBackupConnectionResultOutputWithContext(ctx context.Context) LookupACSSBackupConnectionResultOutput {
	return o
}

func (o LookupACSSBackupConnectionResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupACSSBackupConnectionResult] {
	return pulumix.Output[LookupACSSBackupConnectionResult]{
		OutputState: o.OutputState,
	}
}

// Information about the recovery services vault and backup policy used for backup.
func (o LookupACSSBackupConnectionResultOutput) BackupData() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupACSSBackupConnectionResult) interface{} { return v.BackupData }).(pulumi.AnyOutput)
}

// Defines the errors related to backup connection resource of virtual instance for SAP.
func (o LookupACSSBackupConnectionResultOutput) Errors() ConnectorErrorDefinitionResponseOutput {
	return o.ApplyT(func(v LookupACSSBackupConnectionResult) ConnectorErrorDefinitionResponse { return v.Errors }).(ConnectorErrorDefinitionResponseOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupACSSBackupConnectionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupACSSBackupConnectionResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupACSSBackupConnectionResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupACSSBackupConnectionResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupACSSBackupConnectionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupACSSBackupConnectionResult) string { return v.Name }).(pulumi.StringOutput)
}

// Defines the provisioning states.
func (o LookupACSSBackupConnectionResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupACSSBackupConnectionResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupACSSBackupConnectionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupACSSBackupConnectionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupACSSBackupConnectionResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupACSSBackupConnectionResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupACSSBackupConnectionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupACSSBackupConnectionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupACSSBackupConnectionResultOutput{})
}
