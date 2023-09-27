// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230315preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get a FleetMember
func LookupFleetMember(ctx *pulumi.Context, args *LookupFleetMemberArgs, opts ...pulumi.InvokeOption) (*LookupFleetMemberResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupFleetMemberResult
	err := ctx.Invoke("azure-native:containerservice/v20230315preview:getFleetMember", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFleetMemberArgs struct {
	// The name of the Fleet member resource.
	FleetMemberName string `pulumi:"fleetMemberName"`
	// The name of the Fleet resource.
	FleetName string `pulumi:"fleetName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A member of the Fleet. It contains a reference to an existing Kubernetes cluster on Azure.
type LookupFleetMemberResult struct {
	// The ARM resource id of the cluster that joins the Fleet. Must be a valid Azure resource id. e.g.: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{clusterName}'.
	ClusterResourceId string `pulumi:"clusterResourceId"`
	// If eTag is provided in the response body, it may also be provided as a header per the normal etag convention.  Entity tags are used for comparing two or more entities from the same requested resource. HTTP/1.1 uses entity tags in the etag (section 14.19), If-Match (section 14.24), If-None-Match (section 14.26), and If-Range (section 14.27) header fields.
	ETag string `pulumi:"eTag"`
	// The group this member belongs to for multi-cluster update management.
	Group *string `pulumi:"group"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The status of the last operation.
	ProvisioningState string `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupFleetMemberOutput(ctx *pulumi.Context, args LookupFleetMemberOutputArgs, opts ...pulumi.InvokeOption) LookupFleetMemberResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFleetMemberResult, error) {
			args := v.(LookupFleetMemberArgs)
			r, err := LookupFleetMember(ctx, &args, opts...)
			var s LookupFleetMemberResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFleetMemberResultOutput)
}

type LookupFleetMemberOutputArgs struct {
	// The name of the Fleet member resource.
	FleetMemberName pulumi.StringInput `pulumi:"fleetMemberName"`
	// The name of the Fleet resource.
	FleetName pulumi.StringInput `pulumi:"fleetName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupFleetMemberOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFleetMemberArgs)(nil)).Elem()
}

// A member of the Fleet. It contains a reference to an existing Kubernetes cluster on Azure.
type LookupFleetMemberResultOutput struct{ *pulumi.OutputState }

func (LookupFleetMemberResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFleetMemberResult)(nil)).Elem()
}

func (o LookupFleetMemberResultOutput) ToLookupFleetMemberResultOutput() LookupFleetMemberResultOutput {
	return o
}

func (o LookupFleetMemberResultOutput) ToLookupFleetMemberResultOutputWithContext(ctx context.Context) LookupFleetMemberResultOutput {
	return o
}

func (o LookupFleetMemberResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupFleetMemberResult] {
	return pulumix.Output[LookupFleetMemberResult]{
		OutputState: o.OutputState,
	}
}

// The ARM resource id of the cluster that joins the Fleet. Must be a valid Azure resource id. e.g.: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{clusterName}'.
func (o LookupFleetMemberResultOutput) ClusterResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFleetMemberResult) string { return v.ClusterResourceId }).(pulumi.StringOutput)
}

// If eTag is provided in the response body, it may also be provided as a header per the normal etag convention.  Entity tags are used for comparing two or more entities from the same requested resource. HTTP/1.1 uses entity tags in the etag (section 14.19), If-Match (section 14.24), If-None-Match (section 14.26), and If-Range (section 14.27) header fields.
func (o LookupFleetMemberResultOutput) ETag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFleetMemberResult) string { return v.ETag }).(pulumi.StringOutput)
}

// The group this member belongs to for multi-cluster update management.
func (o LookupFleetMemberResultOutput) Group() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFleetMemberResult) *string { return v.Group }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupFleetMemberResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFleetMemberResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupFleetMemberResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFleetMemberResult) string { return v.Name }).(pulumi.StringOutput)
}

// The status of the last operation.
func (o LookupFleetMemberResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFleetMemberResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupFleetMemberResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupFleetMemberResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupFleetMemberResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFleetMemberResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFleetMemberResultOutput{})
}
