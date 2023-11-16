// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230701preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a hunt relation
func LookupHuntRelation(ctx *pulumi.Context, args *LookupHuntRelationArgs, opts ...pulumi.InvokeOption) (*LookupHuntRelationResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupHuntRelationResult
	err := ctx.Invoke("azure-native:securityinsights/v20230701preview:getHuntRelation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupHuntRelationArgs struct {
	// The hunt id (GUID)
	HuntId string `pulumi:"huntId"`
	// The hunt relation id (GUID)
	HuntRelationId string `pulumi:"huntRelationId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// Represents a Hunt Relation in Azure Security Insights.
type LookupHuntRelationResult struct {
	// Etag of the azure resource
	Etag *string `pulumi:"etag"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// List of labels relevant to this hunt
	Labels []string `pulumi:"labels"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The id of the related resource
	RelatedResourceId string `pulumi:"relatedResourceId"`
	// The resource that the relation is related to
	RelatedResourceKind string `pulumi:"relatedResourceKind"`
	// The name of the related resource
	RelatedResourceName string `pulumi:"relatedResourceName"`
	// The type of the hunt relation
	RelationType string `pulumi:"relationType"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupHuntRelationOutput(ctx *pulumi.Context, args LookupHuntRelationOutputArgs, opts ...pulumi.InvokeOption) LookupHuntRelationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupHuntRelationResult, error) {
			args := v.(LookupHuntRelationArgs)
			r, err := LookupHuntRelation(ctx, &args, opts...)
			var s LookupHuntRelationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupHuntRelationResultOutput)
}

type LookupHuntRelationOutputArgs struct {
	// The hunt id (GUID)
	HuntId pulumi.StringInput `pulumi:"huntId"`
	// The hunt relation id (GUID)
	HuntRelationId pulumi.StringInput `pulumi:"huntRelationId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupHuntRelationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHuntRelationArgs)(nil)).Elem()
}

// Represents a Hunt Relation in Azure Security Insights.
type LookupHuntRelationResultOutput struct{ *pulumi.OutputState }

func (LookupHuntRelationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHuntRelationResult)(nil)).Elem()
}

func (o LookupHuntRelationResultOutput) ToLookupHuntRelationResultOutput() LookupHuntRelationResultOutput {
	return o
}

func (o LookupHuntRelationResultOutput) ToLookupHuntRelationResultOutputWithContext(ctx context.Context) LookupHuntRelationResultOutput {
	return o
}

// Etag of the azure resource
func (o LookupHuntRelationResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHuntRelationResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupHuntRelationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHuntRelationResult) string { return v.Id }).(pulumi.StringOutput)
}

// List of labels relevant to this hunt
func (o LookupHuntRelationResultOutput) Labels() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupHuntRelationResult) []string { return v.Labels }).(pulumi.StringArrayOutput)
}

// The name of the resource
func (o LookupHuntRelationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHuntRelationResult) string { return v.Name }).(pulumi.StringOutput)
}

// The id of the related resource
func (o LookupHuntRelationResultOutput) RelatedResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHuntRelationResult) string { return v.RelatedResourceId }).(pulumi.StringOutput)
}

// The resource that the relation is related to
func (o LookupHuntRelationResultOutput) RelatedResourceKind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHuntRelationResult) string { return v.RelatedResourceKind }).(pulumi.StringOutput)
}

// The name of the related resource
func (o LookupHuntRelationResultOutput) RelatedResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHuntRelationResult) string { return v.RelatedResourceName }).(pulumi.StringOutput)
}

// The type of the hunt relation
func (o LookupHuntRelationResultOutput) RelationType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHuntRelationResult) string { return v.RelationType }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupHuntRelationResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupHuntRelationResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupHuntRelationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHuntRelationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupHuntRelationResultOutput{})
}
