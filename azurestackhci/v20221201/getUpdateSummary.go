// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20221201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get the update summaries for the cluster
func LookupUpdateSummary(ctx *pulumi.Context, args *LookupUpdateSummaryArgs, opts ...pulumi.InvokeOption) (*LookupUpdateSummaryResult, error) {
	var rv LookupUpdateSummaryResult
	err := ctx.Invoke("azure-native:azurestackhci/v20221201:getUpdateSummary", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupUpdateSummaryArgs struct {
	// The name of the cluster.
	ClusterName string `pulumi:"clusterName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Get the update summaries for the cluster
type LookupUpdateSummaryResult struct {
	// Current Solution Bundle version of the stamp.
	CurrentVersion *string `pulumi:"currentVersion"`
	// Name of the hardware model.
	HardwareModel *string `pulumi:"hardwareModel"`
	// Last time the package-specific checks were run.
	HealthCheckDate *string `pulumi:"healthCheckDate"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// Last time the update service successfully checked for updates
	LastChecked *string `pulumi:"lastChecked"`
	// Last time an update installation completed successfully.
	LastUpdated *string `pulumi:"lastUpdated"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// OEM family name.
	OemFamily *string `pulumi:"oemFamily"`
	// Provisioning state of the UpdateSummaries proxy resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Overall update state of the stamp.
	State *string `pulumi:"state"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupUpdateSummaryOutput(ctx *pulumi.Context, args LookupUpdateSummaryOutputArgs, opts ...pulumi.InvokeOption) LookupUpdateSummaryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupUpdateSummaryResult, error) {
			args := v.(LookupUpdateSummaryArgs)
			r, err := LookupUpdateSummary(ctx, &args, opts...)
			var s LookupUpdateSummaryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupUpdateSummaryResultOutput)
}

type LookupUpdateSummaryOutputArgs struct {
	// The name of the cluster.
	ClusterName pulumi.StringInput `pulumi:"clusterName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupUpdateSummaryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUpdateSummaryArgs)(nil)).Elem()
}

// Get the update summaries for the cluster
type LookupUpdateSummaryResultOutput struct{ *pulumi.OutputState }

func (LookupUpdateSummaryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUpdateSummaryResult)(nil)).Elem()
}

func (o LookupUpdateSummaryResultOutput) ToLookupUpdateSummaryResultOutput() LookupUpdateSummaryResultOutput {
	return o
}

func (o LookupUpdateSummaryResultOutput) ToLookupUpdateSummaryResultOutputWithContext(ctx context.Context) LookupUpdateSummaryResultOutput {
	return o
}

// Current Solution Bundle version of the stamp.
func (o LookupUpdateSummaryResultOutput) CurrentVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUpdateSummaryResult) *string { return v.CurrentVersion }).(pulumi.StringPtrOutput)
}

// Name of the hardware model.
func (o LookupUpdateSummaryResultOutput) HardwareModel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUpdateSummaryResult) *string { return v.HardwareModel }).(pulumi.StringPtrOutput)
}

// Last time the package-specific checks were run.
func (o LookupUpdateSummaryResultOutput) HealthCheckDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUpdateSummaryResult) *string { return v.HealthCheckDate }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupUpdateSummaryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUpdateSummaryResult) string { return v.Id }).(pulumi.StringOutput)
}

// Last time the update service successfully checked for updates
func (o LookupUpdateSummaryResultOutput) LastChecked() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUpdateSummaryResult) *string { return v.LastChecked }).(pulumi.StringPtrOutput)
}

// Last time an update installation completed successfully.
func (o LookupUpdateSummaryResultOutput) LastUpdated() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUpdateSummaryResult) *string { return v.LastUpdated }).(pulumi.StringPtrOutput)
}

// The geo-location where the resource lives
func (o LookupUpdateSummaryResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUpdateSummaryResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o LookupUpdateSummaryResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUpdateSummaryResult) string { return v.Name }).(pulumi.StringOutput)
}

// OEM family name.
func (o LookupUpdateSummaryResultOutput) OemFamily() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUpdateSummaryResult) *string { return v.OemFamily }).(pulumi.StringPtrOutput)
}

// Provisioning state of the UpdateSummaries proxy resource.
func (o LookupUpdateSummaryResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUpdateSummaryResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Overall update state of the stamp.
func (o LookupUpdateSummaryResultOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUpdateSummaryResult) *string { return v.State }).(pulumi.StringPtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupUpdateSummaryResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupUpdateSummaryResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupUpdateSummaryResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUpdateSummaryResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupUpdateSummaryResultOutput{})
}
