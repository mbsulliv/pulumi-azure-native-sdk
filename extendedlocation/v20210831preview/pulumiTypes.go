// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210831preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

var _ = utilities.GetEnvOrDefault

// This is optional input that contains the authentication that should be used to generate the namespace.
type CustomLocationPropertiesAuthentication struct {
	// The type of the Custom Locations authentication
	Type *string `pulumi:"type"`
	// The kubeconfig value.
	Value *string `pulumi:"value"`
}

// CustomLocationPropertiesAuthenticationInput is an input type that accepts CustomLocationPropertiesAuthenticationArgs and CustomLocationPropertiesAuthenticationOutput values.
// You can construct a concrete instance of `CustomLocationPropertiesAuthenticationInput` via:
//
//	CustomLocationPropertiesAuthenticationArgs{...}
type CustomLocationPropertiesAuthenticationInput interface {
	pulumi.Input

	ToCustomLocationPropertiesAuthenticationOutput() CustomLocationPropertiesAuthenticationOutput
	ToCustomLocationPropertiesAuthenticationOutputWithContext(context.Context) CustomLocationPropertiesAuthenticationOutput
}

// This is optional input that contains the authentication that should be used to generate the namespace.
type CustomLocationPropertiesAuthenticationArgs struct {
	// The type of the Custom Locations authentication
	Type pulumi.StringPtrInput `pulumi:"type"`
	// The kubeconfig value.
	Value pulumi.StringPtrInput `pulumi:"value"`
}

func (CustomLocationPropertiesAuthenticationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomLocationPropertiesAuthentication)(nil)).Elem()
}

func (i CustomLocationPropertiesAuthenticationArgs) ToCustomLocationPropertiesAuthenticationOutput() CustomLocationPropertiesAuthenticationOutput {
	return i.ToCustomLocationPropertiesAuthenticationOutputWithContext(context.Background())
}

func (i CustomLocationPropertiesAuthenticationArgs) ToCustomLocationPropertiesAuthenticationOutputWithContext(ctx context.Context) CustomLocationPropertiesAuthenticationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomLocationPropertiesAuthenticationOutput)
}

func (i CustomLocationPropertiesAuthenticationArgs) ToOutput(ctx context.Context) pulumix.Output[CustomLocationPropertiesAuthentication] {
	return pulumix.Output[CustomLocationPropertiesAuthentication]{
		OutputState: i.ToCustomLocationPropertiesAuthenticationOutputWithContext(ctx).OutputState,
	}
}

func (i CustomLocationPropertiesAuthenticationArgs) ToCustomLocationPropertiesAuthenticationPtrOutput() CustomLocationPropertiesAuthenticationPtrOutput {
	return i.ToCustomLocationPropertiesAuthenticationPtrOutputWithContext(context.Background())
}

func (i CustomLocationPropertiesAuthenticationArgs) ToCustomLocationPropertiesAuthenticationPtrOutputWithContext(ctx context.Context) CustomLocationPropertiesAuthenticationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomLocationPropertiesAuthenticationOutput).ToCustomLocationPropertiesAuthenticationPtrOutputWithContext(ctx)
}

// CustomLocationPropertiesAuthenticationPtrInput is an input type that accepts CustomLocationPropertiesAuthenticationArgs, CustomLocationPropertiesAuthenticationPtr and CustomLocationPropertiesAuthenticationPtrOutput values.
// You can construct a concrete instance of `CustomLocationPropertiesAuthenticationPtrInput` via:
//
//	        CustomLocationPropertiesAuthenticationArgs{...}
//
//	or:
//
//	        nil
type CustomLocationPropertiesAuthenticationPtrInput interface {
	pulumi.Input

	ToCustomLocationPropertiesAuthenticationPtrOutput() CustomLocationPropertiesAuthenticationPtrOutput
	ToCustomLocationPropertiesAuthenticationPtrOutputWithContext(context.Context) CustomLocationPropertiesAuthenticationPtrOutput
}

type customLocationPropertiesAuthenticationPtrType CustomLocationPropertiesAuthenticationArgs

func CustomLocationPropertiesAuthenticationPtr(v *CustomLocationPropertiesAuthenticationArgs) CustomLocationPropertiesAuthenticationPtrInput {
	return (*customLocationPropertiesAuthenticationPtrType)(v)
}

func (*customLocationPropertiesAuthenticationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomLocationPropertiesAuthentication)(nil)).Elem()
}

func (i *customLocationPropertiesAuthenticationPtrType) ToCustomLocationPropertiesAuthenticationPtrOutput() CustomLocationPropertiesAuthenticationPtrOutput {
	return i.ToCustomLocationPropertiesAuthenticationPtrOutputWithContext(context.Background())
}

func (i *customLocationPropertiesAuthenticationPtrType) ToCustomLocationPropertiesAuthenticationPtrOutputWithContext(ctx context.Context) CustomLocationPropertiesAuthenticationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomLocationPropertiesAuthenticationPtrOutput)
}

func (i *customLocationPropertiesAuthenticationPtrType) ToOutput(ctx context.Context) pulumix.Output[*CustomLocationPropertiesAuthentication] {
	return pulumix.Output[*CustomLocationPropertiesAuthentication]{
		OutputState: i.ToCustomLocationPropertiesAuthenticationPtrOutputWithContext(ctx).OutputState,
	}
}

// This is optional input that contains the authentication that should be used to generate the namespace.
type CustomLocationPropertiesAuthenticationOutput struct{ *pulumi.OutputState }

func (CustomLocationPropertiesAuthenticationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomLocationPropertiesAuthentication)(nil)).Elem()
}

func (o CustomLocationPropertiesAuthenticationOutput) ToCustomLocationPropertiesAuthenticationOutput() CustomLocationPropertiesAuthenticationOutput {
	return o
}

func (o CustomLocationPropertiesAuthenticationOutput) ToCustomLocationPropertiesAuthenticationOutputWithContext(ctx context.Context) CustomLocationPropertiesAuthenticationOutput {
	return o
}

func (o CustomLocationPropertiesAuthenticationOutput) ToCustomLocationPropertiesAuthenticationPtrOutput() CustomLocationPropertiesAuthenticationPtrOutput {
	return o.ToCustomLocationPropertiesAuthenticationPtrOutputWithContext(context.Background())
}

func (o CustomLocationPropertiesAuthenticationOutput) ToCustomLocationPropertiesAuthenticationPtrOutputWithContext(ctx context.Context) CustomLocationPropertiesAuthenticationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CustomLocationPropertiesAuthentication) *CustomLocationPropertiesAuthentication {
		return &v
	}).(CustomLocationPropertiesAuthenticationPtrOutput)
}

func (o CustomLocationPropertiesAuthenticationOutput) ToOutput(ctx context.Context) pulumix.Output[CustomLocationPropertiesAuthentication] {
	return pulumix.Output[CustomLocationPropertiesAuthentication]{
		OutputState: o.OutputState,
	}
}

// The type of the Custom Locations authentication
func (o CustomLocationPropertiesAuthenticationOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomLocationPropertiesAuthentication) *string { return v.Type }).(pulumi.StringPtrOutput)
}

// The kubeconfig value.
func (o CustomLocationPropertiesAuthenticationOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomLocationPropertiesAuthentication) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type CustomLocationPropertiesAuthenticationPtrOutput struct{ *pulumi.OutputState }

func (CustomLocationPropertiesAuthenticationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomLocationPropertiesAuthentication)(nil)).Elem()
}

func (o CustomLocationPropertiesAuthenticationPtrOutput) ToCustomLocationPropertiesAuthenticationPtrOutput() CustomLocationPropertiesAuthenticationPtrOutput {
	return o
}

func (o CustomLocationPropertiesAuthenticationPtrOutput) ToCustomLocationPropertiesAuthenticationPtrOutputWithContext(ctx context.Context) CustomLocationPropertiesAuthenticationPtrOutput {
	return o
}

func (o CustomLocationPropertiesAuthenticationPtrOutput) ToOutput(ctx context.Context) pulumix.Output[*CustomLocationPropertiesAuthentication] {
	return pulumix.Output[*CustomLocationPropertiesAuthentication]{
		OutputState: o.OutputState,
	}
}

func (o CustomLocationPropertiesAuthenticationPtrOutput) Elem() CustomLocationPropertiesAuthenticationOutput {
	return o.ApplyT(func(v *CustomLocationPropertiesAuthentication) CustomLocationPropertiesAuthentication {
		if v != nil {
			return *v
		}
		var ret CustomLocationPropertiesAuthentication
		return ret
	}).(CustomLocationPropertiesAuthenticationOutput)
}

// The type of the Custom Locations authentication
func (o CustomLocationPropertiesAuthenticationPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomLocationPropertiesAuthentication) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

// The kubeconfig value.
func (o CustomLocationPropertiesAuthenticationPtrOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomLocationPropertiesAuthentication) *string {
		if v == nil {
			return nil
		}
		return v.Value
	}).(pulumi.StringPtrOutput)
}

// This is optional input that contains the authentication that should be used to generate the namespace.
type CustomLocationPropertiesResponseAuthentication struct {
	// The type of the Custom Locations authentication
	Type *string `pulumi:"type"`
}

// This is optional input that contains the authentication that should be used to generate the namespace.
type CustomLocationPropertiesResponseAuthenticationOutput struct{ *pulumi.OutputState }

func (CustomLocationPropertiesResponseAuthenticationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomLocationPropertiesResponseAuthentication)(nil)).Elem()
}

func (o CustomLocationPropertiesResponseAuthenticationOutput) ToCustomLocationPropertiesResponseAuthenticationOutput() CustomLocationPropertiesResponseAuthenticationOutput {
	return o
}

func (o CustomLocationPropertiesResponseAuthenticationOutput) ToCustomLocationPropertiesResponseAuthenticationOutputWithContext(ctx context.Context) CustomLocationPropertiesResponseAuthenticationOutput {
	return o
}

func (o CustomLocationPropertiesResponseAuthenticationOutput) ToOutput(ctx context.Context) pulumix.Output[CustomLocationPropertiesResponseAuthentication] {
	return pulumix.Output[CustomLocationPropertiesResponseAuthentication]{
		OutputState: o.OutputState,
	}
}

// The type of the Custom Locations authentication
func (o CustomLocationPropertiesResponseAuthenticationOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomLocationPropertiesResponseAuthentication) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type CustomLocationPropertiesResponseAuthenticationPtrOutput struct{ *pulumi.OutputState }

func (CustomLocationPropertiesResponseAuthenticationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomLocationPropertiesResponseAuthentication)(nil)).Elem()
}

func (o CustomLocationPropertiesResponseAuthenticationPtrOutput) ToCustomLocationPropertiesResponseAuthenticationPtrOutput() CustomLocationPropertiesResponseAuthenticationPtrOutput {
	return o
}

func (o CustomLocationPropertiesResponseAuthenticationPtrOutput) ToCustomLocationPropertiesResponseAuthenticationPtrOutputWithContext(ctx context.Context) CustomLocationPropertiesResponseAuthenticationPtrOutput {
	return o
}

func (o CustomLocationPropertiesResponseAuthenticationPtrOutput) ToOutput(ctx context.Context) pulumix.Output[*CustomLocationPropertiesResponseAuthentication] {
	return pulumix.Output[*CustomLocationPropertiesResponseAuthentication]{
		OutputState: o.OutputState,
	}
}

func (o CustomLocationPropertiesResponseAuthenticationPtrOutput) Elem() CustomLocationPropertiesResponseAuthenticationOutput {
	return o.ApplyT(func(v *CustomLocationPropertiesResponseAuthentication) CustomLocationPropertiesResponseAuthentication {
		if v != nil {
			return *v
		}
		var ret CustomLocationPropertiesResponseAuthentication
		return ret
	}).(CustomLocationPropertiesResponseAuthenticationOutput)
}

// The type of the Custom Locations authentication
func (o CustomLocationPropertiesResponseAuthenticationPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomLocationPropertiesResponseAuthentication) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

// Identity for the resource.
type Identity struct {
	// The identity type.
	Type *string `pulumi:"type"`
}

// IdentityInput is an input type that accepts IdentityArgs and IdentityOutput values.
// You can construct a concrete instance of `IdentityInput` via:
//
//	IdentityArgs{...}
type IdentityInput interface {
	pulumi.Input

	ToIdentityOutput() IdentityOutput
	ToIdentityOutputWithContext(context.Context) IdentityOutput
}

// Identity for the resource.
type IdentityArgs struct {
	// The identity type.
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (IdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Identity)(nil)).Elem()
}

func (i IdentityArgs) ToIdentityOutput() IdentityOutput {
	return i.ToIdentityOutputWithContext(context.Background())
}

func (i IdentityArgs) ToIdentityOutputWithContext(ctx context.Context) IdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityOutput)
}

func (i IdentityArgs) ToOutput(ctx context.Context) pulumix.Output[Identity] {
	return pulumix.Output[Identity]{
		OutputState: i.ToIdentityOutputWithContext(ctx).OutputState,
	}
}

func (i IdentityArgs) ToIdentityPtrOutput() IdentityPtrOutput {
	return i.ToIdentityPtrOutputWithContext(context.Background())
}

func (i IdentityArgs) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityOutput).ToIdentityPtrOutputWithContext(ctx)
}

// IdentityPtrInput is an input type that accepts IdentityArgs, IdentityPtr and IdentityPtrOutput values.
// You can construct a concrete instance of `IdentityPtrInput` via:
//
//	        IdentityArgs{...}
//
//	or:
//
//	        nil
type IdentityPtrInput interface {
	pulumi.Input

	ToIdentityPtrOutput() IdentityPtrOutput
	ToIdentityPtrOutputWithContext(context.Context) IdentityPtrOutput
}

type identityPtrType IdentityArgs

func IdentityPtr(v *IdentityArgs) IdentityPtrInput {
	return (*identityPtrType)(v)
}

func (*identityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Identity)(nil)).Elem()
}

func (i *identityPtrType) ToIdentityPtrOutput() IdentityPtrOutput {
	return i.ToIdentityPtrOutputWithContext(context.Background())
}

func (i *identityPtrType) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityPtrOutput)
}

func (i *identityPtrType) ToOutput(ctx context.Context) pulumix.Output[*Identity] {
	return pulumix.Output[*Identity]{
		OutputState: i.ToIdentityPtrOutputWithContext(ctx).OutputState,
	}
}

// Identity for the resource.
type IdentityOutput struct{ *pulumi.OutputState }

func (IdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Identity)(nil)).Elem()
}

func (o IdentityOutput) ToIdentityOutput() IdentityOutput {
	return o
}

func (o IdentityOutput) ToIdentityOutputWithContext(ctx context.Context) IdentityOutput {
	return o
}

func (o IdentityOutput) ToIdentityPtrOutput() IdentityPtrOutput {
	return o.ToIdentityPtrOutputWithContext(context.Background())
}

func (o IdentityOutput) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Identity) *Identity {
		return &v
	}).(IdentityPtrOutput)
}

func (o IdentityOutput) ToOutput(ctx context.Context) pulumix.Output[Identity] {
	return pulumix.Output[Identity]{
		OutputState: o.OutputState,
	}
}

// The identity type.
func (o IdentityOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Identity) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type IdentityPtrOutput struct{ *pulumi.OutputState }

func (IdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Identity)(nil)).Elem()
}

func (o IdentityPtrOutput) ToIdentityPtrOutput() IdentityPtrOutput {
	return o
}

func (o IdentityPtrOutput) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return o
}

func (o IdentityPtrOutput) ToOutput(ctx context.Context) pulumix.Output[*Identity] {
	return pulumix.Output[*Identity]{
		OutputState: o.OutputState,
	}
}

func (o IdentityPtrOutput) Elem() IdentityOutput {
	return o.ApplyT(func(v *Identity) Identity {
		if v != nil {
			return *v
		}
		var ret Identity
		return ret
	}).(IdentityOutput)
}

// The identity type.
func (o IdentityPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Identity) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

// Identity for the resource.
type IdentityResponse struct {
	// The principal ID of resource identity.
	PrincipalId string `pulumi:"principalId"`
	// The tenant ID of resource.
	TenantId string `pulumi:"tenantId"`
	// The identity type.
	Type *string `pulumi:"type"`
}

// Identity for the resource.
type IdentityResponseOutput struct{ *pulumi.OutputState }

func (IdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityResponse)(nil)).Elem()
}

func (o IdentityResponseOutput) ToIdentityResponseOutput() IdentityResponseOutput {
	return o
}

func (o IdentityResponseOutput) ToIdentityResponseOutputWithContext(ctx context.Context) IdentityResponseOutput {
	return o
}

func (o IdentityResponseOutput) ToOutput(ctx context.Context) pulumix.Output[IdentityResponse] {
	return pulumix.Output[IdentityResponse]{
		OutputState: o.OutputState,
	}
}

// The principal ID of resource identity.
func (o IdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

// The tenant ID of resource.
func (o IdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

// The identity type.
func (o IdentityResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentityResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type IdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (IdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityResponse)(nil)).Elem()
}

func (o IdentityResponsePtrOutput) ToIdentityResponsePtrOutput() IdentityResponsePtrOutput {
	return o
}

func (o IdentityResponsePtrOutput) ToIdentityResponsePtrOutputWithContext(ctx context.Context) IdentityResponsePtrOutput {
	return o
}

func (o IdentityResponsePtrOutput) ToOutput(ctx context.Context) pulumix.Output[*IdentityResponse] {
	return pulumix.Output[*IdentityResponse]{
		OutputState: o.OutputState,
	}
}

func (o IdentityResponsePtrOutput) Elem() IdentityResponseOutput {
	return o.ApplyT(func(v *IdentityResponse) IdentityResponse {
		if v != nil {
			return *v
		}
		var ret IdentityResponse
		return ret
	}).(IdentityResponseOutput)
}

// The principal ID of resource identity.
func (o IdentityResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

// The tenant ID of resource.
func (o IdentityResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

// The identity type.
func (o IdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

// A label selector is composed of two parts, matchLabels and matchExpressions. The first part, matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The second part, matchExpressions is a list of resource selector requirements. Valid operators include In, NotIn, Exists, and DoesNotExist. The values set must be non-empty in the case of In and NotIn. The values set must be empty in the case of Exists and DoesNotExist. All of the requirements, from both matchLabels and matchExpressions must all be satisfied in order to match.
type ResourceSyncRulePropertiesResponseSelector struct {
	// MatchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'.
	MatchLabels map[string]string `pulumi:"matchLabels"`
}

// A label selector is composed of two parts, matchLabels and matchExpressions. The first part, matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The second part, matchExpressions is a list of resource selector requirements. Valid operators include In, NotIn, Exists, and DoesNotExist. The values set must be non-empty in the case of In and NotIn. The values set must be empty in the case of Exists and DoesNotExist. All of the requirements, from both matchLabels and matchExpressions must all be satisfied in order to match.
type ResourceSyncRulePropertiesResponseSelectorOutput struct{ *pulumi.OutputState }

func (ResourceSyncRulePropertiesResponseSelectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceSyncRulePropertiesResponseSelector)(nil)).Elem()
}

func (o ResourceSyncRulePropertiesResponseSelectorOutput) ToResourceSyncRulePropertiesResponseSelectorOutput() ResourceSyncRulePropertiesResponseSelectorOutput {
	return o
}

func (o ResourceSyncRulePropertiesResponseSelectorOutput) ToResourceSyncRulePropertiesResponseSelectorOutputWithContext(ctx context.Context) ResourceSyncRulePropertiesResponseSelectorOutput {
	return o
}

func (o ResourceSyncRulePropertiesResponseSelectorOutput) ToOutput(ctx context.Context) pulumix.Output[ResourceSyncRulePropertiesResponseSelector] {
	return pulumix.Output[ResourceSyncRulePropertiesResponseSelector]{
		OutputState: o.OutputState,
	}
}

// MatchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'.
func (o ResourceSyncRulePropertiesResponseSelectorOutput) MatchLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v ResourceSyncRulePropertiesResponseSelector) map[string]string { return v.MatchLabels }).(pulumi.StringMapOutput)
}

type ResourceSyncRulePropertiesResponseSelectorPtrOutput struct{ *pulumi.OutputState }

func (ResourceSyncRulePropertiesResponseSelectorPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceSyncRulePropertiesResponseSelector)(nil)).Elem()
}

func (o ResourceSyncRulePropertiesResponseSelectorPtrOutput) ToResourceSyncRulePropertiesResponseSelectorPtrOutput() ResourceSyncRulePropertiesResponseSelectorPtrOutput {
	return o
}

func (o ResourceSyncRulePropertiesResponseSelectorPtrOutput) ToResourceSyncRulePropertiesResponseSelectorPtrOutputWithContext(ctx context.Context) ResourceSyncRulePropertiesResponseSelectorPtrOutput {
	return o
}

func (o ResourceSyncRulePropertiesResponseSelectorPtrOutput) ToOutput(ctx context.Context) pulumix.Output[*ResourceSyncRulePropertiesResponseSelector] {
	return pulumix.Output[*ResourceSyncRulePropertiesResponseSelector]{
		OutputState: o.OutputState,
	}
}

func (o ResourceSyncRulePropertiesResponseSelectorPtrOutput) Elem() ResourceSyncRulePropertiesResponseSelectorOutput {
	return o.ApplyT(func(v *ResourceSyncRulePropertiesResponseSelector) ResourceSyncRulePropertiesResponseSelector {
		if v != nil {
			return *v
		}
		var ret ResourceSyncRulePropertiesResponseSelector
		return ret
	}).(ResourceSyncRulePropertiesResponseSelectorOutput)
}

// MatchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'.
func (o ResourceSyncRulePropertiesResponseSelectorPtrOutput) MatchLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ResourceSyncRulePropertiesResponseSelector) map[string]string {
		if v == nil {
			return nil
		}
		return v.MatchLabels
	}).(pulumi.StringMapOutput)
}

// A label selector is composed of two parts, matchLabels and matchExpressions. The first part, matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The second part, matchExpressions is a list of resource selector requirements. Valid operators include In, NotIn, Exists, and DoesNotExist. The values set must be non-empty in the case of In and NotIn. The values set must be empty in the case of Exists and DoesNotExist. All of the requirements, from both matchLabels and matchExpressions must all be satisfied in order to match.
type ResourceSyncRulePropertiesSelector struct {
	// MatchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'.
	MatchLabels map[string]string `pulumi:"matchLabels"`
}

// ResourceSyncRulePropertiesSelectorInput is an input type that accepts ResourceSyncRulePropertiesSelectorArgs and ResourceSyncRulePropertiesSelectorOutput values.
// You can construct a concrete instance of `ResourceSyncRulePropertiesSelectorInput` via:
//
//	ResourceSyncRulePropertiesSelectorArgs{...}
type ResourceSyncRulePropertiesSelectorInput interface {
	pulumi.Input

	ToResourceSyncRulePropertiesSelectorOutput() ResourceSyncRulePropertiesSelectorOutput
	ToResourceSyncRulePropertiesSelectorOutputWithContext(context.Context) ResourceSyncRulePropertiesSelectorOutput
}

// A label selector is composed of two parts, matchLabels and matchExpressions. The first part, matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The second part, matchExpressions is a list of resource selector requirements. Valid operators include In, NotIn, Exists, and DoesNotExist. The values set must be non-empty in the case of In and NotIn. The values set must be empty in the case of Exists and DoesNotExist. All of the requirements, from both matchLabels and matchExpressions must all be satisfied in order to match.
type ResourceSyncRulePropertiesSelectorArgs struct {
	// MatchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'.
	MatchLabels pulumi.StringMapInput `pulumi:"matchLabels"`
}

func (ResourceSyncRulePropertiesSelectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceSyncRulePropertiesSelector)(nil)).Elem()
}

func (i ResourceSyncRulePropertiesSelectorArgs) ToResourceSyncRulePropertiesSelectorOutput() ResourceSyncRulePropertiesSelectorOutput {
	return i.ToResourceSyncRulePropertiesSelectorOutputWithContext(context.Background())
}

func (i ResourceSyncRulePropertiesSelectorArgs) ToResourceSyncRulePropertiesSelectorOutputWithContext(ctx context.Context) ResourceSyncRulePropertiesSelectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceSyncRulePropertiesSelectorOutput)
}

func (i ResourceSyncRulePropertiesSelectorArgs) ToOutput(ctx context.Context) pulumix.Output[ResourceSyncRulePropertiesSelector] {
	return pulumix.Output[ResourceSyncRulePropertiesSelector]{
		OutputState: i.ToResourceSyncRulePropertiesSelectorOutputWithContext(ctx).OutputState,
	}
}

func (i ResourceSyncRulePropertiesSelectorArgs) ToResourceSyncRulePropertiesSelectorPtrOutput() ResourceSyncRulePropertiesSelectorPtrOutput {
	return i.ToResourceSyncRulePropertiesSelectorPtrOutputWithContext(context.Background())
}

func (i ResourceSyncRulePropertiesSelectorArgs) ToResourceSyncRulePropertiesSelectorPtrOutputWithContext(ctx context.Context) ResourceSyncRulePropertiesSelectorPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceSyncRulePropertiesSelectorOutput).ToResourceSyncRulePropertiesSelectorPtrOutputWithContext(ctx)
}

// ResourceSyncRulePropertiesSelectorPtrInput is an input type that accepts ResourceSyncRulePropertiesSelectorArgs, ResourceSyncRulePropertiesSelectorPtr and ResourceSyncRulePropertiesSelectorPtrOutput values.
// You can construct a concrete instance of `ResourceSyncRulePropertiesSelectorPtrInput` via:
//
//	        ResourceSyncRulePropertiesSelectorArgs{...}
//
//	or:
//
//	        nil
type ResourceSyncRulePropertiesSelectorPtrInput interface {
	pulumi.Input

	ToResourceSyncRulePropertiesSelectorPtrOutput() ResourceSyncRulePropertiesSelectorPtrOutput
	ToResourceSyncRulePropertiesSelectorPtrOutputWithContext(context.Context) ResourceSyncRulePropertiesSelectorPtrOutput
}

type resourceSyncRulePropertiesSelectorPtrType ResourceSyncRulePropertiesSelectorArgs

func ResourceSyncRulePropertiesSelectorPtr(v *ResourceSyncRulePropertiesSelectorArgs) ResourceSyncRulePropertiesSelectorPtrInput {
	return (*resourceSyncRulePropertiesSelectorPtrType)(v)
}

func (*resourceSyncRulePropertiesSelectorPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceSyncRulePropertiesSelector)(nil)).Elem()
}

func (i *resourceSyncRulePropertiesSelectorPtrType) ToResourceSyncRulePropertiesSelectorPtrOutput() ResourceSyncRulePropertiesSelectorPtrOutput {
	return i.ToResourceSyncRulePropertiesSelectorPtrOutputWithContext(context.Background())
}

func (i *resourceSyncRulePropertiesSelectorPtrType) ToResourceSyncRulePropertiesSelectorPtrOutputWithContext(ctx context.Context) ResourceSyncRulePropertiesSelectorPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceSyncRulePropertiesSelectorPtrOutput)
}

func (i *resourceSyncRulePropertiesSelectorPtrType) ToOutput(ctx context.Context) pulumix.Output[*ResourceSyncRulePropertiesSelector] {
	return pulumix.Output[*ResourceSyncRulePropertiesSelector]{
		OutputState: i.ToResourceSyncRulePropertiesSelectorPtrOutputWithContext(ctx).OutputState,
	}
}

// A label selector is composed of two parts, matchLabels and matchExpressions. The first part, matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The second part, matchExpressions is a list of resource selector requirements. Valid operators include In, NotIn, Exists, and DoesNotExist. The values set must be non-empty in the case of In and NotIn. The values set must be empty in the case of Exists and DoesNotExist. All of the requirements, from both matchLabels and matchExpressions must all be satisfied in order to match.
type ResourceSyncRulePropertiesSelectorOutput struct{ *pulumi.OutputState }

func (ResourceSyncRulePropertiesSelectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceSyncRulePropertiesSelector)(nil)).Elem()
}

func (o ResourceSyncRulePropertiesSelectorOutput) ToResourceSyncRulePropertiesSelectorOutput() ResourceSyncRulePropertiesSelectorOutput {
	return o
}

func (o ResourceSyncRulePropertiesSelectorOutput) ToResourceSyncRulePropertiesSelectorOutputWithContext(ctx context.Context) ResourceSyncRulePropertiesSelectorOutput {
	return o
}

func (o ResourceSyncRulePropertiesSelectorOutput) ToResourceSyncRulePropertiesSelectorPtrOutput() ResourceSyncRulePropertiesSelectorPtrOutput {
	return o.ToResourceSyncRulePropertiesSelectorPtrOutputWithContext(context.Background())
}

func (o ResourceSyncRulePropertiesSelectorOutput) ToResourceSyncRulePropertiesSelectorPtrOutputWithContext(ctx context.Context) ResourceSyncRulePropertiesSelectorPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceSyncRulePropertiesSelector) *ResourceSyncRulePropertiesSelector {
		return &v
	}).(ResourceSyncRulePropertiesSelectorPtrOutput)
}

func (o ResourceSyncRulePropertiesSelectorOutput) ToOutput(ctx context.Context) pulumix.Output[ResourceSyncRulePropertiesSelector] {
	return pulumix.Output[ResourceSyncRulePropertiesSelector]{
		OutputState: o.OutputState,
	}
}

// MatchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'.
func (o ResourceSyncRulePropertiesSelectorOutput) MatchLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v ResourceSyncRulePropertiesSelector) map[string]string { return v.MatchLabels }).(pulumi.StringMapOutput)
}

type ResourceSyncRulePropertiesSelectorPtrOutput struct{ *pulumi.OutputState }

func (ResourceSyncRulePropertiesSelectorPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceSyncRulePropertiesSelector)(nil)).Elem()
}

func (o ResourceSyncRulePropertiesSelectorPtrOutput) ToResourceSyncRulePropertiesSelectorPtrOutput() ResourceSyncRulePropertiesSelectorPtrOutput {
	return o
}

func (o ResourceSyncRulePropertiesSelectorPtrOutput) ToResourceSyncRulePropertiesSelectorPtrOutputWithContext(ctx context.Context) ResourceSyncRulePropertiesSelectorPtrOutput {
	return o
}

func (o ResourceSyncRulePropertiesSelectorPtrOutput) ToOutput(ctx context.Context) pulumix.Output[*ResourceSyncRulePropertiesSelector] {
	return pulumix.Output[*ResourceSyncRulePropertiesSelector]{
		OutputState: o.OutputState,
	}
}

func (o ResourceSyncRulePropertiesSelectorPtrOutput) Elem() ResourceSyncRulePropertiesSelectorOutput {
	return o.ApplyT(func(v *ResourceSyncRulePropertiesSelector) ResourceSyncRulePropertiesSelector {
		if v != nil {
			return *v
		}
		var ret ResourceSyncRulePropertiesSelector
		return ret
	}).(ResourceSyncRulePropertiesSelectorOutput)
}

// MatchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'.
func (o ResourceSyncRulePropertiesSelectorPtrOutput) MatchLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ResourceSyncRulePropertiesSelector) map[string]string {
		if v == nil {
			return nil
		}
		return v.MatchLabels
	}).(pulumi.StringMapOutput)
}

// Metadata pertaining to creation and last modification of the resource.
type SystemDataResponse struct {
	// The timestamp of resource creation (UTC).
	CreatedAt *string `pulumi:"createdAt"`
	// The identity that created the resource.
	CreatedBy *string `pulumi:"createdBy"`
	// The type of identity that created the resource.
	CreatedByType *string `pulumi:"createdByType"`
	// The timestamp of resource last modification (UTC)
	LastModifiedAt *string `pulumi:"lastModifiedAt"`
	// The identity that last modified the resource.
	LastModifiedBy *string `pulumi:"lastModifiedBy"`
	// The type of identity that last modified the resource.
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
}

// Metadata pertaining to creation and last modification of the resource.
type SystemDataResponseOutput struct{ *pulumi.OutputState }

func (SystemDataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemDataResponse)(nil)).Elem()
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutput() SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutputWithContext(ctx context.Context) SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) ToOutput(ctx context.Context) pulumix.Output[SystemDataResponse] {
	return pulumix.Output[SystemDataResponse]{
		OutputState: o.OutputState,
	}
}

// The timestamp of resource creation (UTC).
func (o SystemDataResponseOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

// The identity that created the resource.
func (o SystemDataResponseOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

// The type of identity that created the resource.
func (o SystemDataResponseOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

// The timestamp of resource last modification (UTC)
func (o SystemDataResponseOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedAt }).(pulumi.StringPtrOutput)
}

// The identity that last modified the resource.
func (o SystemDataResponseOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

// The type of identity that last modified the resource.
func (o SystemDataResponseOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(CustomLocationPropertiesAuthenticationOutput{})
	pulumi.RegisterOutputType(CustomLocationPropertiesAuthenticationPtrOutput{})
	pulumi.RegisterOutputType(CustomLocationPropertiesResponseAuthenticationOutput{})
	pulumi.RegisterOutputType(CustomLocationPropertiesResponseAuthenticationPtrOutput{})
	pulumi.RegisterOutputType(IdentityOutput{})
	pulumi.RegisterOutputType(IdentityPtrOutput{})
	pulumi.RegisterOutputType(IdentityResponseOutput{})
	pulumi.RegisterOutputType(IdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(ResourceSyncRulePropertiesResponseSelectorOutput{})
	pulumi.RegisterOutputType(ResourceSyncRulePropertiesResponseSelectorPtrOutput{})
	pulumi.RegisterOutputType(ResourceSyncRulePropertiesSelectorOutput{})
	pulumi.RegisterOutputType(ResourceSyncRulePropertiesSelectorPtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
}
