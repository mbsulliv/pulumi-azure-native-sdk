// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20170426

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets information about the specified relationship.
func LookupRelationship(ctx *pulumi.Context, args *LookupRelationshipArgs, opts ...pulumi.InvokeOption) (*LookupRelationshipResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupRelationshipResult
	err := ctx.Invoke("azure-native:customerinsights/v20170426:getRelationship", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRelationshipArgs struct {
	// The name of the hub.
	HubName string `pulumi:"hubName"`
	// The name of the relationship.
	RelationshipName string `pulumi:"relationshipName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The relationship resource format.
type LookupRelationshipResult struct {
	// The Relationship Cardinality.
	Cardinality *string `pulumi:"cardinality"`
	// Localized descriptions for the Relationship.
	Description map[string]string `pulumi:"description"`
	// Localized display name for the Relationship.
	DisplayName map[string]string `pulumi:"displayName"`
	// The expiry date time in UTC.
	ExpiryDateTimeUtc *string `pulumi:"expiryDateTimeUtc"`
	// The properties of the Relationship.
	Fields []PropertyDefinitionResponse `pulumi:"fields"`
	// Resource ID.
	Id string `pulumi:"id"`
	// Optional property to be used to map fields in profile to their strong ids in related profile.
	LookupMappings []RelationshipTypeMappingResponse `pulumi:"lookupMappings"`
	// Resource name.
	Name string `pulumi:"name"`
	// Profile type.
	ProfileType string `pulumi:"profileType"`
	// Provisioning state.
	ProvisioningState string `pulumi:"provisioningState"`
	// Related profile being referenced.
	RelatedProfileType string `pulumi:"relatedProfileType"`
	// The relationship guid id.
	RelationshipGuidId string `pulumi:"relationshipGuidId"`
	// The Relationship name.
	RelationshipName string `pulumi:"relationshipName"`
	// The hub name.
	TenantId string `pulumi:"tenantId"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupRelationshipOutput(ctx *pulumi.Context, args LookupRelationshipOutputArgs, opts ...pulumi.InvokeOption) LookupRelationshipResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRelationshipResult, error) {
			args := v.(LookupRelationshipArgs)
			r, err := LookupRelationship(ctx, &args, opts...)
			var s LookupRelationshipResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRelationshipResultOutput)
}

type LookupRelationshipOutputArgs struct {
	// The name of the hub.
	HubName pulumi.StringInput `pulumi:"hubName"`
	// The name of the relationship.
	RelationshipName pulumi.StringInput `pulumi:"relationshipName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupRelationshipOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRelationshipArgs)(nil)).Elem()
}

// The relationship resource format.
type LookupRelationshipResultOutput struct{ *pulumi.OutputState }

func (LookupRelationshipResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRelationshipResult)(nil)).Elem()
}

func (o LookupRelationshipResultOutput) ToLookupRelationshipResultOutput() LookupRelationshipResultOutput {
	return o
}

func (o LookupRelationshipResultOutput) ToLookupRelationshipResultOutputWithContext(ctx context.Context) LookupRelationshipResultOutput {
	return o
}

// The Relationship Cardinality.
func (o LookupRelationshipResultOutput) Cardinality() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRelationshipResult) *string { return v.Cardinality }).(pulumi.StringPtrOutput)
}

// Localized descriptions for the Relationship.
func (o LookupRelationshipResultOutput) Description() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupRelationshipResult) map[string]string { return v.Description }).(pulumi.StringMapOutput)
}

// Localized display name for the Relationship.
func (o LookupRelationshipResultOutput) DisplayName() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupRelationshipResult) map[string]string { return v.DisplayName }).(pulumi.StringMapOutput)
}

// The expiry date time in UTC.
func (o LookupRelationshipResultOutput) ExpiryDateTimeUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRelationshipResult) *string { return v.ExpiryDateTimeUtc }).(pulumi.StringPtrOutput)
}

// The properties of the Relationship.
func (o LookupRelationshipResultOutput) Fields() PropertyDefinitionResponseArrayOutput {
	return o.ApplyT(func(v LookupRelationshipResult) []PropertyDefinitionResponse { return v.Fields }).(PropertyDefinitionResponseArrayOutput)
}

// Resource ID.
func (o LookupRelationshipResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRelationshipResult) string { return v.Id }).(pulumi.StringOutput)
}

// Optional property to be used to map fields in profile to their strong ids in related profile.
func (o LookupRelationshipResultOutput) LookupMappings() RelationshipTypeMappingResponseArrayOutput {
	return o.ApplyT(func(v LookupRelationshipResult) []RelationshipTypeMappingResponse { return v.LookupMappings }).(RelationshipTypeMappingResponseArrayOutput)
}

// Resource name.
func (o LookupRelationshipResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRelationshipResult) string { return v.Name }).(pulumi.StringOutput)
}

// Profile type.
func (o LookupRelationshipResultOutput) ProfileType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRelationshipResult) string { return v.ProfileType }).(pulumi.StringOutput)
}

// Provisioning state.
func (o LookupRelationshipResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRelationshipResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Related profile being referenced.
func (o LookupRelationshipResultOutput) RelatedProfileType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRelationshipResult) string { return v.RelatedProfileType }).(pulumi.StringOutput)
}

// The relationship guid id.
func (o LookupRelationshipResultOutput) RelationshipGuidId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRelationshipResult) string { return v.RelationshipGuidId }).(pulumi.StringOutput)
}

// The Relationship name.
func (o LookupRelationshipResultOutput) RelationshipName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRelationshipResult) string { return v.RelationshipName }).(pulumi.StringOutput)
}

// The hub name.
func (o LookupRelationshipResultOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRelationshipResult) string { return v.TenantId }).(pulumi.StringOutput)
}

// Resource type.
func (o LookupRelationshipResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRelationshipResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRelationshipResultOutput{})
}
