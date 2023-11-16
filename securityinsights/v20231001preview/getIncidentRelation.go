// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20231001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets an incident relation.
func LookupIncidentRelation(ctx *pulumi.Context, args *LookupIncidentRelationArgs, opts ...pulumi.InvokeOption) (*LookupIncidentRelationResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupIncidentRelationResult
	err := ctx.Invoke("azure-native:securityinsights/v20231001preview:getIncidentRelation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIncidentRelationArgs struct {
	// Incident ID
	IncidentId string `pulumi:"incidentId"`
	// Relation Name
	RelationName string `pulumi:"relationName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// Represents a relation between two resources
type LookupIncidentRelationResult struct {
	// Etag of the azure resource
	Etag *string `pulumi:"etag"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The resource ID of the related resource
	RelatedResourceId string `pulumi:"relatedResourceId"`
	// The resource kind of the related resource
	RelatedResourceKind string `pulumi:"relatedResourceKind"`
	// The name of the related resource
	RelatedResourceName string `pulumi:"relatedResourceName"`
	// The resource type of the related resource
	RelatedResourceType string `pulumi:"relatedResourceType"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupIncidentRelationOutput(ctx *pulumi.Context, args LookupIncidentRelationOutputArgs, opts ...pulumi.InvokeOption) LookupIncidentRelationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIncidentRelationResult, error) {
			args := v.(LookupIncidentRelationArgs)
			r, err := LookupIncidentRelation(ctx, &args, opts...)
			var s LookupIncidentRelationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIncidentRelationResultOutput)
}

type LookupIncidentRelationOutputArgs struct {
	// Incident ID
	IncidentId pulumi.StringInput `pulumi:"incidentId"`
	// Relation Name
	RelationName pulumi.StringInput `pulumi:"relationName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupIncidentRelationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIncidentRelationArgs)(nil)).Elem()
}

// Represents a relation between two resources
type LookupIncidentRelationResultOutput struct{ *pulumi.OutputState }

func (LookupIncidentRelationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIncidentRelationResult)(nil)).Elem()
}

func (o LookupIncidentRelationResultOutput) ToLookupIncidentRelationResultOutput() LookupIncidentRelationResultOutput {
	return o
}

func (o LookupIncidentRelationResultOutput) ToLookupIncidentRelationResultOutputWithContext(ctx context.Context) LookupIncidentRelationResultOutput {
	return o
}

// Etag of the azure resource
func (o LookupIncidentRelationResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIncidentRelationResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupIncidentRelationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIncidentRelationResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupIncidentRelationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIncidentRelationResult) string { return v.Name }).(pulumi.StringOutput)
}

// The resource ID of the related resource
func (o LookupIncidentRelationResultOutput) RelatedResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIncidentRelationResult) string { return v.RelatedResourceId }).(pulumi.StringOutput)
}

// The resource kind of the related resource
func (o LookupIncidentRelationResultOutput) RelatedResourceKind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIncidentRelationResult) string { return v.RelatedResourceKind }).(pulumi.StringOutput)
}

// The name of the related resource
func (o LookupIncidentRelationResultOutput) RelatedResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIncidentRelationResult) string { return v.RelatedResourceName }).(pulumi.StringOutput)
}

// The resource type of the related resource
func (o LookupIncidentRelationResultOutput) RelatedResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIncidentRelationResult) string { return v.RelatedResourceType }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupIncidentRelationResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupIncidentRelationResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupIncidentRelationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIncidentRelationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIncidentRelationResultOutput{})
}
