// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20240101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets an incident task.
func LookupIncidentTask(ctx *pulumi.Context, args *LookupIncidentTaskArgs, opts ...pulumi.InvokeOption) (*LookupIncidentTaskResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupIncidentTaskResult
	err := ctx.Invoke("azure-native:securityinsights/v20240101preview:getIncidentTask", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIncidentTaskArgs struct {
	// Incident ID
	IncidentId string `pulumi:"incidentId"`
	// Incident task ID
	IncidentTaskId string `pulumi:"incidentTaskId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

type LookupIncidentTaskResult struct {
	// Information on the client (user or application) that made some action
	CreatedBy *ClientInfoResponse `pulumi:"createdBy"`
	// The time the task was created
	CreatedTimeUtc string `pulumi:"createdTimeUtc"`
	// The description of the task
	Description *string `pulumi:"description"`
	// Etag of the azure resource
	Etag *string `pulumi:"etag"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// Information on the client (user or application) that made some action
	LastModifiedBy *ClientInfoResponse `pulumi:"lastModifiedBy"`
	// The last time the task was updated
	LastModifiedTimeUtc string `pulumi:"lastModifiedTimeUtc"`
	// The name of the resource
	Name   string `pulumi:"name"`
	Status string `pulumi:"status"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The title of the task
	Title string `pulumi:"title"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupIncidentTaskOutput(ctx *pulumi.Context, args LookupIncidentTaskOutputArgs, opts ...pulumi.InvokeOption) LookupIncidentTaskResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIncidentTaskResult, error) {
			args := v.(LookupIncidentTaskArgs)
			r, err := LookupIncidentTask(ctx, &args, opts...)
			var s LookupIncidentTaskResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIncidentTaskResultOutput)
}

type LookupIncidentTaskOutputArgs struct {
	// Incident ID
	IncidentId pulumi.StringInput `pulumi:"incidentId"`
	// Incident task ID
	IncidentTaskId pulumi.StringInput `pulumi:"incidentTaskId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupIncidentTaskOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIncidentTaskArgs)(nil)).Elem()
}

type LookupIncidentTaskResultOutput struct{ *pulumi.OutputState }

func (LookupIncidentTaskResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIncidentTaskResult)(nil)).Elem()
}

func (o LookupIncidentTaskResultOutput) ToLookupIncidentTaskResultOutput() LookupIncidentTaskResultOutput {
	return o
}

func (o LookupIncidentTaskResultOutput) ToLookupIncidentTaskResultOutputWithContext(ctx context.Context) LookupIncidentTaskResultOutput {
	return o
}

// Information on the client (user or application) that made some action
func (o LookupIncidentTaskResultOutput) CreatedBy() ClientInfoResponsePtrOutput {
	return o.ApplyT(func(v LookupIncidentTaskResult) *ClientInfoResponse { return v.CreatedBy }).(ClientInfoResponsePtrOutput)
}

// The time the task was created
func (o LookupIncidentTaskResultOutput) CreatedTimeUtc() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIncidentTaskResult) string { return v.CreatedTimeUtc }).(pulumi.StringOutput)
}

// The description of the task
func (o LookupIncidentTaskResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIncidentTaskResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Etag of the azure resource
func (o LookupIncidentTaskResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIncidentTaskResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupIncidentTaskResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIncidentTaskResult) string { return v.Id }).(pulumi.StringOutput)
}

// Information on the client (user or application) that made some action
func (o LookupIncidentTaskResultOutput) LastModifiedBy() ClientInfoResponsePtrOutput {
	return o.ApplyT(func(v LookupIncidentTaskResult) *ClientInfoResponse { return v.LastModifiedBy }).(ClientInfoResponsePtrOutput)
}

// The last time the task was updated
func (o LookupIncidentTaskResultOutput) LastModifiedTimeUtc() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIncidentTaskResult) string { return v.LastModifiedTimeUtc }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupIncidentTaskResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIncidentTaskResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupIncidentTaskResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIncidentTaskResult) string { return v.Status }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupIncidentTaskResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupIncidentTaskResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The title of the task
func (o LookupIncidentTaskResultOutput) Title() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIncidentTaskResult) string { return v.Title }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupIncidentTaskResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIncidentTaskResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIncidentTaskResultOutput{})
}
