// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package servicelinker

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// get a dryrun job
// Azure REST API version: 2022-11-01-preview.
//
// Other available API versions: 2023-04-01-preview.
func LookupConnectorDryrun(ctx *pulumi.Context, args *LookupConnectorDryrunArgs, opts ...pulumi.InvokeOption) (*LookupConnectorDryrunResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupConnectorDryrunResult
	err := ctx.Invoke("azure-native:servicelinker:getConnectorDryrun", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConnectorDryrunArgs struct {
	// The name of dryrun.
	DryrunName string `pulumi:"dryrunName"`
	// The name of Azure region.
	Location string `pulumi:"location"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The ID of the target subscription.
	SubscriptionId *string `pulumi:"subscriptionId"`
}

// a dryrun job resource
type LookupConnectorDryrunResult struct {
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// the preview of the operations for creation
	OperationPreviews []DryrunOperationPreviewResponse `pulumi:"operationPreviews"`
	// The parameters of the dryrun
	Parameters *CreateOrUpdateDryrunParametersResponse `pulumi:"parameters"`
	// the result of the dryrun
	PrerequisiteResults []interface{} `pulumi:"prerequisiteResults"`
	// The provisioning state.
	ProvisioningState string `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupConnectorDryrunOutput(ctx *pulumi.Context, args LookupConnectorDryrunOutputArgs, opts ...pulumi.InvokeOption) LookupConnectorDryrunResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupConnectorDryrunResult, error) {
			args := v.(LookupConnectorDryrunArgs)
			r, err := LookupConnectorDryrun(ctx, &args, opts...)
			var s LookupConnectorDryrunResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupConnectorDryrunResultOutput)
}

type LookupConnectorDryrunOutputArgs struct {
	// The name of dryrun.
	DryrunName pulumi.StringInput `pulumi:"dryrunName"`
	// The name of Azure region.
	Location pulumi.StringInput `pulumi:"location"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The ID of the target subscription.
	SubscriptionId pulumi.StringPtrInput `pulumi:"subscriptionId"`
}

func (LookupConnectorDryrunOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConnectorDryrunArgs)(nil)).Elem()
}

// a dryrun job resource
type LookupConnectorDryrunResultOutput struct{ *pulumi.OutputState }

func (LookupConnectorDryrunResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConnectorDryrunResult)(nil)).Elem()
}

func (o LookupConnectorDryrunResultOutput) ToLookupConnectorDryrunResultOutput() LookupConnectorDryrunResultOutput {
	return o
}

func (o LookupConnectorDryrunResultOutput) ToLookupConnectorDryrunResultOutputWithContext(ctx context.Context) LookupConnectorDryrunResultOutput {
	return o
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupConnectorDryrunResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectorDryrunResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupConnectorDryrunResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectorDryrunResult) string { return v.Name }).(pulumi.StringOutput)
}

// the preview of the operations for creation
func (o LookupConnectorDryrunResultOutput) OperationPreviews() DryrunOperationPreviewResponseArrayOutput {
	return o.ApplyT(func(v LookupConnectorDryrunResult) []DryrunOperationPreviewResponse { return v.OperationPreviews }).(DryrunOperationPreviewResponseArrayOutput)
}

// The parameters of the dryrun
func (o LookupConnectorDryrunResultOutput) Parameters() CreateOrUpdateDryrunParametersResponsePtrOutput {
	return o.ApplyT(func(v LookupConnectorDryrunResult) *CreateOrUpdateDryrunParametersResponse { return v.Parameters }).(CreateOrUpdateDryrunParametersResponsePtrOutput)
}

// the result of the dryrun
func (o LookupConnectorDryrunResultOutput) PrerequisiteResults() pulumi.ArrayOutput {
	return o.ApplyT(func(v LookupConnectorDryrunResult) []interface{} { return v.PrerequisiteResults }).(pulumi.ArrayOutput)
}

// The provisioning state.
func (o LookupConnectorDryrunResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectorDryrunResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupConnectorDryrunResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupConnectorDryrunResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupConnectorDryrunResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectorDryrunResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupConnectorDryrunResultOutput{})
}
