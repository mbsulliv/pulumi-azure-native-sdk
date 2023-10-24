// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package customerinsights

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets information about the specified relationship Link.
// Azure REST API version: 2017-04-26.
//
// Other available API versions: 2017-01-01.
func LookupRelationshipLink(ctx *pulumi.Context, args *LookupRelationshipLinkArgs, opts ...pulumi.InvokeOption) (*LookupRelationshipLinkResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupRelationshipLinkResult
	err := ctx.Invoke("azure-native:customerinsights:getRelationshipLink", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRelationshipLinkArgs struct {
	// The name of the hub.
	HubName string `pulumi:"hubName"`
	// The name of the relationship link.
	RelationshipLinkName string `pulumi:"relationshipLinkName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The relationship link resource format.
type LookupRelationshipLinkResult struct {
	// Localized descriptions for the Relationship Link.
	Description map[string]string `pulumi:"description"`
	// Localized display name for the Relationship Link.
	DisplayName map[string]string `pulumi:"displayName"`
	// Resource ID.
	Id string `pulumi:"id"`
	// The InteractionType associated with the Relationship Link.
	InteractionType string `pulumi:"interactionType"`
	// The name of the Relationship Link.
	LinkName string `pulumi:"linkName"`
	// The mappings between Interaction and Relationship fields.
	Mappings []RelationshipLinkFieldMappingResponse `pulumi:"mappings"`
	// Resource name.
	Name string `pulumi:"name"`
	// The property references for the Profile of the Relationship.
	ProfilePropertyReferences []ParticipantProfilePropertyReferenceResponse `pulumi:"profilePropertyReferences"`
	// Provisioning state.
	ProvisioningState string `pulumi:"provisioningState"`
	// The property references for the Related Profile of the Relationship.
	RelatedProfilePropertyReferences []ParticipantProfilePropertyReferenceResponse `pulumi:"relatedProfilePropertyReferences"`
	// The relationship guid id.
	RelationshipGuidId string `pulumi:"relationshipGuidId"`
	// The Relationship associated with the Link.
	RelationshipName string `pulumi:"relationshipName"`
	// The hub name.
	TenantId string `pulumi:"tenantId"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupRelationshipLinkOutput(ctx *pulumi.Context, args LookupRelationshipLinkOutputArgs, opts ...pulumi.InvokeOption) LookupRelationshipLinkResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRelationshipLinkResult, error) {
			args := v.(LookupRelationshipLinkArgs)
			r, err := LookupRelationshipLink(ctx, &args, opts...)
			var s LookupRelationshipLinkResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRelationshipLinkResultOutput)
}

type LookupRelationshipLinkOutputArgs struct {
	// The name of the hub.
	HubName pulumi.StringInput `pulumi:"hubName"`
	// The name of the relationship link.
	RelationshipLinkName pulumi.StringInput `pulumi:"relationshipLinkName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupRelationshipLinkOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRelationshipLinkArgs)(nil)).Elem()
}

// The relationship link resource format.
type LookupRelationshipLinkResultOutput struct{ *pulumi.OutputState }

func (LookupRelationshipLinkResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRelationshipLinkResult)(nil)).Elem()
}

func (o LookupRelationshipLinkResultOutput) ToLookupRelationshipLinkResultOutput() LookupRelationshipLinkResultOutput {
	return o
}

func (o LookupRelationshipLinkResultOutput) ToLookupRelationshipLinkResultOutputWithContext(ctx context.Context) LookupRelationshipLinkResultOutput {
	return o
}

func (o LookupRelationshipLinkResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupRelationshipLinkResult] {
	return pulumix.Output[LookupRelationshipLinkResult]{
		OutputState: o.OutputState,
	}
}

// Localized descriptions for the Relationship Link.
func (o LookupRelationshipLinkResultOutput) Description() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupRelationshipLinkResult) map[string]string { return v.Description }).(pulumi.StringMapOutput)
}

// Localized display name for the Relationship Link.
func (o LookupRelationshipLinkResultOutput) DisplayName() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupRelationshipLinkResult) map[string]string { return v.DisplayName }).(pulumi.StringMapOutput)
}

// Resource ID.
func (o LookupRelationshipLinkResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRelationshipLinkResult) string { return v.Id }).(pulumi.StringOutput)
}

// The InteractionType associated with the Relationship Link.
func (o LookupRelationshipLinkResultOutput) InteractionType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRelationshipLinkResult) string { return v.InteractionType }).(pulumi.StringOutput)
}

// The name of the Relationship Link.
func (o LookupRelationshipLinkResultOutput) LinkName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRelationshipLinkResult) string { return v.LinkName }).(pulumi.StringOutput)
}

// The mappings between Interaction and Relationship fields.
func (o LookupRelationshipLinkResultOutput) Mappings() RelationshipLinkFieldMappingResponseArrayOutput {
	return o.ApplyT(func(v LookupRelationshipLinkResult) []RelationshipLinkFieldMappingResponse { return v.Mappings }).(RelationshipLinkFieldMappingResponseArrayOutput)
}

// Resource name.
func (o LookupRelationshipLinkResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRelationshipLinkResult) string { return v.Name }).(pulumi.StringOutput)
}

// The property references for the Profile of the Relationship.
func (o LookupRelationshipLinkResultOutput) ProfilePropertyReferences() ParticipantProfilePropertyReferenceResponseArrayOutput {
	return o.ApplyT(func(v LookupRelationshipLinkResult) []ParticipantProfilePropertyReferenceResponse {
		return v.ProfilePropertyReferences
	}).(ParticipantProfilePropertyReferenceResponseArrayOutput)
}

// Provisioning state.
func (o LookupRelationshipLinkResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRelationshipLinkResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The property references for the Related Profile of the Relationship.
func (o LookupRelationshipLinkResultOutput) RelatedProfilePropertyReferences() ParticipantProfilePropertyReferenceResponseArrayOutput {
	return o.ApplyT(func(v LookupRelationshipLinkResult) []ParticipantProfilePropertyReferenceResponse {
		return v.RelatedProfilePropertyReferences
	}).(ParticipantProfilePropertyReferenceResponseArrayOutput)
}

// The relationship guid id.
func (o LookupRelationshipLinkResultOutput) RelationshipGuidId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRelationshipLinkResult) string { return v.RelationshipGuidId }).(pulumi.StringOutput)
}

// The Relationship associated with the Link.
func (o LookupRelationshipLinkResultOutput) RelationshipName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRelationshipLinkResult) string { return v.RelationshipName }).(pulumi.StringOutput)
}

// The hub name.
func (o LookupRelationshipLinkResultOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRelationshipLinkResult) string { return v.TenantId }).(pulumi.StringOutput)
}

// Resource type.
func (o LookupRelationshipLinkResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRelationshipLinkResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRelationshipLinkResultOutput{})
}
