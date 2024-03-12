// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// get a dryrun job
func LookupLinkerDryrun(ctx *pulumi.Context, args *LookupLinkerDryrunArgs, opts ...pulumi.InvokeOption) (*LookupLinkerDryrunResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupLinkerDryrunResult
	err := ctx.Invoke("azure-native:servicelinker/v20230401preview:getLinkerDryrun", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLinkerDryrunArgs struct {
	// The name of dryrun.
	DryrunName string `pulumi:"dryrunName"`
	// The fully qualified Azure Resource manager identifier of the resource to be connected.
	ResourceUri string `pulumi:"resourceUri"`
}

// a dryrun job resource
type LookupLinkerDryrunResult struct {
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

func LookupLinkerDryrunOutput(ctx *pulumi.Context, args LookupLinkerDryrunOutputArgs, opts ...pulumi.InvokeOption) LookupLinkerDryrunResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupLinkerDryrunResult, error) {
			args := v.(LookupLinkerDryrunArgs)
			r, err := LookupLinkerDryrun(ctx, &args, opts...)
			var s LookupLinkerDryrunResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupLinkerDryrunResultOutput)
}

type LookupLinkerDryrunOutputArgs struct {
	// The name of dryrun.
	DryrunName pulumi.StringInput `pulumi:"dryrunName"`
	// The fully qualified Azure Resource manager identifier of the resource to be connected.
	ResourceUri pulumi.StringInput `pulumi:"resourceUri"`
}

func (LookupLinkerDryrunOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLinkerDryrunArgs)(nil)).Elem()
}

// a dryrun job resource
type LookupLinkerDryrunResultOutput struct{ *pulumi.OutputState }

func (LookupLinkerDryrunResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLinkerDryrunResult)(nil)).Elem()
}

func (o LookupLinkerDryrunResultOutput) ToLookupLinkerDryrunResultOutput() LookupLinkerDryrunResultOutput {
	return o
}

func (o LookupLinkerDryrunResultOutput) ToLookupLinkerDryrunResultOutputWithContext(ctx context.Context) LookupLinkerDryrunResultOutput {
	return o
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupLinkerDryrunResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLinkerDryrunResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupLinkerDryrunResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLinkerDryrunResult) string { return v.Name }).(pulumi.StringOutput)
}

// the preview of the operations for creation
func (o LookupLinkerDryrunResultOutput) OperationPreviews() DryrunOperationPreviewResponseArrayOutput {
	return o.ApplyT(func(v LookupLinkerDryrunResult) []DryrunOperationPreviewResponse { return v.OperationPreviews }).(DryrunOperationPreviewResponseArrayOutput)
}

// The parameters of the dryrun
func (o LookupLinkerDryrunResultOutput) Parameters() CreateOrUpdateDryrunParametersResponsePtrOutput {
	return o.ApplyT(func(v LookupLinkerDryrunResult) *CreateOrUpdateDryrunParametersResponse { return v.Parameters }).(CreateOrUpdateDryrunParametersResponsePtrOutput)
}

// the result of the dryrun
func (o LookupLinkerDryrunResultOutput) PrerequisiteResults() pulumi.ArrayOutput {
	return o.ApplyT(func(v LookupLinkerDryrunResult) []interface{} { return v.PrerequisiteResults }).(pulumi.ArrayOutput)
}

// The provisioning state.
func (o LookupLinkerDryrunResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLinkerDryrunResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupLinkerDryrunResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupLinkerDryrunResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupLinkerDryrunResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLinkerDryrunResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLinkerDryrunResultOutput{})
}
