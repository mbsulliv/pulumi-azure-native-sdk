// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = utilities.GetEnvOrDefault

// Additional attribute or filter to allow subscriptions meeting the requirements to be part of the GroupQuota.
type AdditionalAttributes struct {
	Environment interface{} `pulumi:"environment"`
	// The grouping Id for the group quota. It can be Billing Id or ServiceTreeId if applicable.
	GroupId GroupingId `pulumi:"groupId"`
}

// AdditionalAttributesInput is an input type that accepts AdditionalAttributesArgs and AdditionalAttributesOutput values.
// You can construct a concrete instance of `AdditionalAttributesInput` via:
//
//	AdditionalAttributesArgs{...}
type AdditionalAttributesInput interface {
	pulumi.Input

	ToAdditionalAttributesOutput() AdditionalAttributesOutput
	ToAdditionalAttributesOutputWithContext(context.Context) AdditionalAttributesOutput
}

// Additional attribute or filter to allow subscriptions meeting the requirements to be part of the GroupQuota.
type AdditionalAttributesArgs struct {
	Environment pulumi.Input `pulumi:"environment"`
	// The grouping Id for the group quota. It can be Billing Id or ServiceTreeId if applicable.
	GroupId GroupingIdInput `pulumi:"groupId"`
}

func (AdditionalAttributesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AdditionalAttributes)(nil)).Elem()
}

func (i AdditionalAttributesArgs) ToAdditionalAttributesOutput() AdditionalAttributesOutput {
	return i.ToAdditionalAttributesOutputWithContext(context.Background())
}

func (i AdditionalAttributesArgs) ToAdditionalAttributesOutputWithContext(ctx context.Context) AdditionalAttributesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdditionalAttributesOutput)
}

func (i AdditionalAttributesArgs) ToAdditionalAttributesPtrOutput() AdditionalAttributesPtrOutput {
	return i.ToAdditionalAttributesPtrOutputWithContext(context.Background())
}

func (i AdditionalAttributesArgs) ToAdditionalAttributesPtrOutputWithContext(ctx context.Context) AdditionalAttributesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdditionalAttributesOutput).ToAdditionalAttributesPtrOutputWithContext(ctx)
}

// AdditionalAttributesPtrInput is an input type that accepts AdditionalAttributesArgs, AdditionalAttributesPtr and AdditionalAttributesPtrOutput values.
// You can construct a concrete instance of `AdditionalAttributesPtrInput` via:
//
//	        AdditionalAttributesArgs{...}
//
//	or:
//
//	        nil
type AdditionalAttributesPtrInput interface {
	pulumi.Input

	ToAdditionalAttributesPtrOutput() AdditionalAttributesPtrOutput
	ToAdditionalAttributesPtrOutputWithContext(context.Context) AdditionalAttributesPtrOutput
}

type additionalAttributesPtrType AdditionalAttributesArgs

func AdditionalAttributesPtr(v *AdditionalAttributesArgs) AdditionalAttributesPtrInput {
	return (*additionalAttributesPtrType)(v)
}

func (*additionalAttributesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AdditionalAttributes)(nil)).Elem()
}

func (i *additionalAttributesPtrType) ToAdditionalAttributesPtrOutput() AdditionalAttributesPtrOutput {
	return i.ToAdditionalAttributesPtrOutputWithContext(context.Background())
}

func (i *additionalAttributesPtrType) ToAdditionalAttributesPtrOutputWithContext(ctx context.Context) AdditionalAttributesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdditionalAttributesPtrOutput)
}

// Additional attribute or filter to allow subscriptions meeting the requirements to be part of the GroupQuota.
type AdditionalAttributesOutput struct{ *pulumi.OutputState }

func (AdditionalAttributesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AdditionalAttributes)(nil)).Elem()
}

func (o AdditionalAttributesOutput) ToAdditionalAttributesOutput() AdditionalAttributesOutput {
	return o
}

func (o AdditionalAttributesOutput) ToAdditionalAttributesOutputWithContext(ctx context.Context) AdditionalAttributesOutput {
	return o
}

func (o AdditionalAttributesOutput) ToAdditionalAttributesPtrOutput() AdditionalAttributesPtrOutput {
	return o.ToAdditionalAttributesPtrOutputWithContext(context.Background())
}

func (o AdditionalAttributesOutput) ToAdditionalAttributesPtrOutputWithContext(ctx context.Context) AdditionalAttributesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AdditionalAttributes) *AdditionalAttributes {
		return &v
	}).(AdditionalAttributesPtrOutput)
}

func (o AdditionalAttributesOutput) Environment() pulumi.AnyOutput {
	return o.ApplyT(func(v AdditionalAttributes) interface{} { return v.Environment }).(pulumi.AnyOutput)
}

// The grouping Id for the group quota. It can be Billing Id or ServiceTreeId if applicable.
func (o AdditionalAttributesOutput) GroupId() GroupingIdOutput {
	return o.ApplyT(func(v AdditionalAttributes) GroupingId { return v.GroupId }).(GroupingIdOutput)
}

type AdditionalAttributesPtrOutput struct{ *pulumi.OutputState }

func (AdditionalAttributesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AdditionalAttributes)(nil)).Elem()
}

func (o AdditionalAttributesPtrOutput) ToAdditionalAttributesPtrOutput() AdditionalAttributesPtrOutput {
	return o
}

func (o AdditionalAttributesPtrOutput) ToAdditionalAttributesPtrOutputWithContext(ctx context.Context) AdditionalAttributesPtrOutput {
	return o
}

func (o AdditionalAttributesPtrOutput) Elem() AdditionalAttributesOutput {
	return o.ApplyT(func(v *AdditionalAttributes) AdditionalAttributes {
		if v != nil {
			return *v
		}
		var ret AdditionalAttributes
		return ret
	}).(AdditionalAttributesOutput)
}

func (o AdditionalAttributesPtrOutput) Environment() pulumi.AnyOutput {
	return o.ApplyT(func(v *AdditionalAttributes) interface{} {
		if v == nil {
			return nil
		}
		return v.Environment
	}).(pulumi.AnyOutput)
}

// The grouping Id for the group quota. It can be Billing Id or ServiceTreeId if applicable.
func (o AdditionalAttributesPtrOutput) GroupId() GroupingIdPtrOutput {
	return o.ApplyT(func(v *AdditionalAttributes) *GroupingId {
		if v == nil {
			return nil
		}
		return &v.GroupId
	}).(GroupingIdPtrOutput)
}

// Additional attribute or filter to allow subscriptions meeting the requirements to be part of the GroupQuota.
type AdditionalAttributesResponse struct {
	Environment interface{} `pulumi:"environment"`
	// The grouping Id for the group quota. It can be Billing Id or ServiceTreeId if applicable.
	GroupId GroupingIdResponse `pulumi:"groupId"`
}

// Additional attribute or filter to allow subscriptions meeting the requirements to be part of the GroupQuota.
type AdditionalAttributesResponseOutput struct{ *pulumi.OutputState }

func (AdditionalAttributesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AdditionalAttributesResponse)(nil)).Elem()
}

func (o AdditionalAttributesResponseOutput) ToAdditionalAttributesResponseOutput() AdditionalAttributesResponseOutput {
	return o
}

func (o AdditionalAttributesResponseOutput) ToAdditionalAttributesResponseOutputWithContext(ctx context.Context) AdditionalAttributesResponseOutput {
	return o
}

func (o AdditionalAttributesResponseOutput) Environment() pulumi.AnyOutput {
	return o.ApplyT(func(v AdditionalAttributesResponse) interface{} { return v.Environment }).(pulumi.AnyOutput)
}

// The grouping Id for the group quota. It can be Billing Id or ServiceTreeId if applicable.
func (o AdditionalAttributesResponseOutput) GroupId() GroupingIdResponseOutput {
	return o.ApplyT(func(v AdditionalAttributesResponse) GroupingIdResponse { return v.GroupId }).(GroupingIdResponseOutput)
}

type AdditionalAttributesResponsePtrOutput struct{ *pulumi.OutputState }

func (AdditionalAttributesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AdditionalAttributesResponse)(nil)).Elem()
}

func (o AdditionalAttributesResponsePtrOutput) ToAdditionalAttributesResponsePtrOutput() AdditionalAttributesResponsePtrOutput {
	return o
}

func (o AdditionalAttributesResponsePtrOutput) ToAdditionalAttributesResponsePtrOutputWithContext(ctx context.Context) AdditionalAttributesResponsePtrOutput {
	return o
}

func (o AdditionalAttributesResponsePtrOutput) Elem() AdditionalAttributesResponseOutput {
	return o.ApplyT(func(v *AdditionalAttributesResponse) AdditionalAttributesResponse {
		if v != nil {
			return *v
		}
		var ret AdditionalAttributesResponse
		return ret
	}).(AdditionalAttributesResponseOutput)
}

func (o AdditionalAttributesResponsePtrOutput) Environment() pulumi.AnyOutput {
	return o.ApplyT(func(v *AdditionalAttributesResponse) interface{} {
		if v == nil {
			return nil
		}
		return v.Environment
	}).(pulumi.AnyOutput)
}

// The grouping Id for the group quota. It can be Billing Id or ServiceTreeId if applicable.
func (o AdditionalAttributesResponsePtrOutput) GroupId() GroupingIdResponsePtrOutput {
	return o.ApplyT(func(v *AdditionalAttributesResponse) *GroupingIdResponse {
		if v == nil {
			return nil
		}
		return &v.GroupId
	}).(GroupingIdResponsePtrOutput)
}

type GroupQuotaSubscriptionIdResponseProperties struct {
	// Status of this subscriptionId being associated with the GroupQuotasEntity.
	ProvisioningState string `pulumi:"provisioningState"`
	// An Azure subscriptionId.
	SubscriptionId string `pulumi:"subscriptionId"`
}

type GroupQuotaSubscriptionIdResponsePropertiesOutput struct{ *pulumi.OutputState }

func (GroupQuotaSubscriptionIdResponsePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GroupQuotaSubscriptionIdResponseProperties)(nil)).Elem()
}

func (o GroupQuotaSubscriptionIdResponsePropertiesOutput) ToGroupQuotaSubscriptionIdResponsePropertiesOutput() GroupQuotaSubscriptionIdResponsePropertiesOutput {
	return o
}

func (o GroupQuotaSubscriptionIdResponsePropertiesOutput) ToGroupQuotaSubscriptionIdResponsePropertiesOutputWithContext(ctx context.Context) GroupQuotaSubscriptionIdResponsePropertiesOutput {
	return o
}

// Status of this subscriptionId being associated with the GroupQuotasEntity.
func (o GroupQuotaSubscriptionIdResponsePropertiesOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v GroupQuotaSubscriptionIdResponseProperties) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// An Azure subscriptionId.
func (o GroupQuotaSubscriptionIdResponsePropertiesOutput) SubscriptionId() pulumi.StringOutput {
	return o.ApplyT(func(v GroupQuotaSubscriptionIdResponseProperties) string { return v.SubscriptionId }).(pulumi.StringOutput)
}

// Properties and filters for ShareQuota. The request parameter is optional, if there are no filters specified.
type GroupQuotasEntityBase struct {
	// Additional attributes to filter/restrict the subscriptions, which can be added to the subscriptionIds.
	AdditionalAttributes *AdditionalAttributes `pulumi:"additionalAttributes"`
	// Display name of the GroupQuota entity.
	DisplayName *string `pulumi:"displayName"`
}

// GroupQuotasEntityBaseInput is an input type that accepts GroupQuotasEntityBaseArgs and GroupQuotasEntityBaseOutput values.
// You can construct a concrete instance of `GroupQuotasEntityBaseInput` via:
//
//	GroupQuotasEntityBaseArgs{...}
type GroupQuotasEntityBaseInput interface {
	pulumi.Input

	ToGroupQuotasEntityBaseOutput() GroupQuotasEntityBaseOutput
	ToGroupQuotasEntityBaseOutputWithContext(context.Context) GroupQuotasEntityBaseOutput
}

// Properties and filters for ShareQuota. The request parameter is optional, if there are no filters specified.
type GroupQuotasEntityBaseArgs struct {
	// Additional attributes to filter/restrict the subscriptions, which can be added to the subscriptionIds.
	AdditionalAttributes AdditionalAttributesPtrInput `pulumi:"additionalAttributes"`
	// Display name of the GroupQuota entity.
	DisplayName pulumi.StringPtrInput `pulumi:"displayName"`
}

func (GroupQuotasEntityBaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GroupQuotasEntityBase)(nil)).Elem()
}

func (i GroupQuotasEntityBaseArgs) ToGroupQuotasEntityBaseOutput() GroupQuotasEntityBaseOutput {
	return i.ToGroupQuotasEntityBaseOutputWithContext(context.Background())
}

func (i GroupQuotasEntityBaseArgs) ToGroupQuotasEntityBaseOutputWithContext(ctx context.Context) GroupQuotasEntityBaseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupQuotasEntityBaseOutput)
}

func (i GroupQuotasEntityBaseArgs) ToGroupQuotasEntityBasePtrOutput() GroupQuotasEntityBasePtrOutput {
	return i.ToGroupQuotasEntityBasePtrOutputWithContext(context.Background())
}

func (i GroupQuotasEntityBaseArgs) ToGroupQuotasEntityBasePtrOutputWithContext(ctx context.Context) GroupQuotasEntityBasePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupQuotasEntityBaseOutput).ToGroupQuotasEntityBasePtrOutputWithContext(ctx)
}

// GroupQuotasEntityBasePtrInput is an input type that accepts GroupQuotasEntityBaseArgs, GroupQuotasEntityBasePtr and GroupQuotasEntityBasePtrOutput values.
// You can construct a concrete instance of `GroupQuotasEntityBasePtrInput` via:
//
//	        GroupQuotasEntityBaseArgs{...}
//
//	or:
//
//	        nil
type GroupQuotasEntityBasePtrInput interface {
	pulumi.Input

	ToGroupQuotasEntityBasePtrOutput() GroupQuotasEntityBasePtrOutput
	ToGroupQuotasEntityBasePtrOutputWithContext(context.Context) GroupQuotasEntityBasePtrOutput
}

type groupQuotasEntityBasePtrType GroupQuotasEntityBaseArgs

func GroupQuotasEntityBasePtr(v *GroupQuotasEntityBaseArgs) GroupQuotasEntityBasePtrInput {
	return (*groupQuotasEntityBasePtrType)(v)
}

func (*groupQuotasEntityBasePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupQuotasEntityBase)(nil)).Elem()
}

func (i *groupQuotasEntityBasePtrType) ToGroupQuotasEntityBasePtrOutput() GroupQuotasEntityBasePtrOutput {
	return i.ToGroupQuotasEntityBasePtrOutputWithContext(context.Background())
}

func (i *groupQuotasEntityBasePtrType) ToGroupQuotasEntityBasePtrOutputWithContext(ctx context.Context) GroupQuotasEntityBasePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupQuotasEntityBasePtrOutput)
}

// Properties and filters for ShareQuota. The request parameter is optional, if there are no filters specified.
type GroupQuotasEntityBaseOutput struct{ *pulumi.OutputState }

func (GroupQuotasEntityBaseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GroupQuotasEntityBase)(nil)).Elem()
}

func (o GroupQuotasEntityBaseOutput) ToGroupQuotasEntityBaseOutput() GroupQuotasEntityBaseOutput {
	return o
}

func (o GroupQuotasEntityBaseOutput) ToGroupQuotasEntityBaseOutputWithContext(ctx context.Context) GroupQuotasEntityBaseOutput {
	return o
}

func (o GroupQuotasEntityBaseOutput) ToGroupQuotasEntityBasePtrOutput() GroupQuotasEntityBasePtrOutput {
	return o.ToGroupQuotasEntityBasePtrOutputWithContext(context.Background())
}

func (o GroupQuotasEntityBaseOutput) ToGroupQuotasEntityBasePtrOutputWithContext(ctx context.Context) GroupQuotasEntityBasePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GroupQuotasEntityBase) *GroupQuotasEntityBase {
		return &v
	}).(GroupQuotasEntityBasePtrOutput)
}

// Additional attributes to filter/restrict the subscriptions, which can be added to the subscriptionIds.
func (o GroupQuotasEntityBaseOutput) AdditionalAttributes() AdditionalAttributesPtrOutput {
	return o.ApplyT(func(v GroupQuotasEntityBase) *AdditionalAttributes { return v.AdditionalAttributes }).(AdditionalAttributesPtrOutput)
}

// Display name of the GroupQuota entity.
func (o GroupQuotasEntityBaseOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GroupQuotasEntityBase) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

type GroupQuotasEntityBasePtrOutput struct{ *pulumi.OutputState }

func (GroupQuotasEntityBasePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupQuotasEntityBase)(nil)).Elem()
}

func (o GroupQuotasEntityBasePtrOutput) ToGroupQuotasEntityBasePtrOutput() GroupQuotasEntityBasePtrOutput {
	return o
}

func (o GroupQuotasEntityBasePtrOutput) ToGroupQuotasEntityBasePtrOutputWithContext(ctx context.Context) GroupQuotasEntityBasePtrOutput {
	return o
}

func (o GroupQuotasEntityBasePtrOutput) Elem() GroupQuotasEntityBaseOutput {
	return o.ApplyT(func(v *GroupQuotasEntityBase) GroupQuotasEntityBase {
		if v != nil {
			return *v
		}
		var ret GroupQuotasEntityBase
		return ret
	}).(GroupQuotasEntityBaseOutput)
}

// Additional attributes to filter/restrict the subscriptions, which can be added to the subscriptionIds.
func (o GroupQuotasEntityBasePtrOutput) AdditionalAttributes() AdditionalAttributesPtrOutput {
	return o.ApplyT(func(v *GroupQuotasEntityBase) *AdditionalAttributes {
		if v == nil {
			return nil
		}
		return v.AdditionalAttributes
	}).(AdditionalAttributesPtrOutput)
}

// Display name of the GroupQuota entity.
func (o GroupQuotasEntityBasePtrOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GroupQuotasEntityBase) *string {
		if v == nil {
			return nil
		}
		return v.DisplayName
	}).(pulumi.StringPtrOutput)
}

// Properties and filters for ShareQuota. The request parameter is optional, if there are no filters specified.
type GroupQuotasEntityBaseResponse struct {
	// Additional attributes to filter/restrict the subscriptions, which can be added to the subscriptionIds.
	AdditionalAttributes *AdditionalAttributesResponse `pulumi:"additionalAttributes"`
	// Display name of the GroupQuota entity.
	DisplayName *string `pulumi:"displayName"`
	// Provisioning state of the operation.
	ProvisioningState string `pulumi:"provisioningState"`
}

// Properties and filters for ShareQuota. The request parameter is optional, if there are no filters specified.
type GroupQuotasEntityBaseResponseOutput struct{ *pulumi.OutputState }

func (GroupQuotasEntityBaseResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GroupQuotasEntityBaseResponse)(nil)).Elem()
}

func (o GroupQuotasEntityBaseResponseOutput) ToGroupQuotasEntityBaseResponseOutput() GroupQuotasEntityBaseResponseOutput {
	return o
}

func (o GroupQuotasEntityBaseResponseOutput) ToGroupQuotasEntityBaseResponseOutputWithContext(ctx context.Context) GroupQuotasEntityBaseResponseOutput {
	return o
}

// Additional attributes to filter/restrict the subscriptions, which can be added to the subscriptionIds.
func (o GroupQuotasEntityBaseResponseOutput) AdditionalAttributes() AdditionalAttributesResponsePtrOutput {
	return o.ApplyT(func(v GroupQuotasEntityBaseResponse) *AdditionalAttributesResponse { return v.AdditionalAttributes }).(AdditionalAttributesResponsePtrOutput)
}

// Display name of the GroupQuota entity.
func (o GroupQuotasEntityBaseResponseOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GroupQuotasEntityBaseResponse) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// Provisioning state of the operation.
func (o GroupQuotasEntityBaseResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v GroupQuotasEntityBaseResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The grouping Id for the group quota. It can be Billing Id or ServiceTreeId if applicable.
type GroupingId struct {
	// GroupingId type. It is a required property. More types of groupIds can be supported in future.
	GroupingIdType *string `pulumi:"groupingIdType"`
	// GroupId value based on the groupingType selected - Billing Id or ServiceTreeId.
	Value *string `pulumi:"value"`
}

// GroupingIdInput is an input type that accepts GroupingIdArgs and GroupingIdOutput values.
// You can construct a concrete instance of `GroupingIdInput` via:
//
//	GroupingIdArgs{...}
type GroupingIdInput interface {
	pulumi.Input

	ToGroupingIdOutput() GroupingIdOutput
	ToGroupingIdOutputWithContext(context.Context) GroupingIdOutput
}

// The grouping Id for the group quota. It can be Billing Id or ServiceTreeId if applicable.
type GroupingIdArgs struct {
	// GroupingId type. It is a required property. More types of groupIds can be supported in future.
	GroupingIdType pulumi.StringPtrInput `pulumi:"groupingIdType"`
	// GroupId value based on the groupingType selected - Billing Id or ServiceTreeId.
	Value pulumi.StringPtrInput `pulumi:"value"`
}

func (GroupingIdArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GroupingId)(nil)).Elem()
}

func (i GroupingIdArgs) ToGroupingIdOutput() GroupingIdOutput {
	return i.ToGroupingIdOutputWithContext(context.Background())
}

func (i GroupingIdArgs) ToGroupingIdOutputWithContext(ctx context.Context) GroupingIdOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupingIdOutput)
}

func (i GroupingIdArgs) ToGroupingIdPtrOutput() GroupingIdPtrOutput {
	return i.ToGroupingIdPtrOutputWithContext(context.Background())
}

func (i GroupingIdArgs) ToGroupingIdPtrOutputWithContext(ctx context.Context) GroupingIdPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupingIdOutput).ToGroupingIdPtrOutputWithContext(ctx)
}

// GroupingIdPtrInput is an input type that accepts GroupingIdArgs, GroupingIdPtr and GroupingIdPtrOutput values.
// You can construct a concrete instance of `GroupingIdPtrInput` via:
//
//	        GroupingIdArgs{...}
//
//	or:
//
//	        nil
type GroupingIdPtrInput interface {
	pulumi.Input

	ToGroupingIdPtrOutput() GroupingIdPtrOutput
	ToGroupingIdPtrOutputWithContext(context.Context) GroupingIdPtrOutput
}

type groupingIdPtrType GroupingIdArgs

func GroupingIdPtr(v *GroupingIdArgs) GroupingIdPtrInput {
	return (*groupingIdPtrType)(v)
}

func (*groupingIdPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupingId)(nil)).Elem()
}

func (i *groupingIdPtrType) ToGroupingIdPtrOutput() GroupingIdPtrOutput {
	return i.ToGroupingIdPtrOutputWithContext(context.Background())
}

func (i *groupingIdPtrType) ToGroupingIdPtrOutputWithContext(ctx context.Context) GroupingIdPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupingIdPtrOutput)
}

// The grouping Id for the group quota. It can be Billing Id or ServiceTreeId if applicable.
type GroupingIdOutput struct{ *pulumi.OutputState }

func (GroupingIdOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GroupingId)(nil)).Elem()
}

func (o GroupingIdOutput) ToGroupingIdOutput() GroupingIdOutput {
	return o
}

func (o GroupingIdOutput) ToGroupingIdOutputWithContext(ctx context.Context) GroupingIdOutput {
	return o
}

func (o GroupingIdOutput) ToGroupingIdPtrOutput() GroupingIdPtrOutput {
	return o.ToGroupingIdPtrOutputWithContext(context.Background())
}

func (o GroupingIdOutput) ToGroupingIdPtrOutputWithContext(ctx context.Context) GroupingIdPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GroupingId) *GroupingId {
		return &v
	}).(GroupingIdPtrOutput)
}

// GroupingId type. It is a required property. More types of groupIds can be supported in future.
func (o GroupingIdOutput) GroupingIdType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GroupingId) *string { return v.GroupingIdType }).(pulumi.StringPtrOutput)
}

// GroupId value based on the groupingType selected - Billing Id or ServiceTreeId.
func (o GroupingIdOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GroupingId) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type GroupingIdPtrOutput struct{ *pulumi.OutputState }

func (GroupingIdPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupingId)(nil)).Elem()
}

func (o GroupingIdPtrOutput) ToGroupingIdPtrOutput() GroupingIdPtrOutput {
	return o
}

func (o GroupingIdPtrOutput) ToGroupingIdPtrOutputWithContext(ctx context.Context) GroupingIdPtrOutput {
	return o
}

func (o GroupingIdPtrOutput) Elem() GroupingIdOutput {
	return o.ApplyT(func(v *GroupingId) GroupingId {
		if v != nil {
			return *v
		}
		var ret GroupingId
		return ret
	}).(GroupingIdOutput)
}

// GroupingId type. It is a required property. More types of groupIds can be supported in future.
func (o GroupingIdPtrOutput) GroupingIdType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GroupingId) *string {
		if v == nil {
			return nil
		}
		return v.GroupingIdType
	}).(pulumi.StringPtrOutput)
}

// GroupId value based on the groupingType selected - Billing Id or ServiceTreeId.
func (o GroupingIdPtrOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GroupingId) *string {
		if v == nil {
			return nil
		}
		return v.Value
	}).(pulumi.StringPtrOutput)
}

// The grouping Id for the group quota. It can be Billing Id or ServiceTreeId if applicable.
type GroupingIdResponse struct {
	// GroupingId type. It is a required property. More types of groupIds can be supported in future.
	GroupingIdType *string `pulumi:"groupingIdType"`
	// GroupId value based on the groupingType selected - Billing Id or ServiceTreeId.
	Value *string `pulumi:"value"`
}

// The grouping Id for the group quota. It can be Billing Id or ServiceTreeId if applicable.
type GroupingIdResponseOutput struct{ *pulumi.OutputState }

func (GroupingIdResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GroupingIdResponse)(nil)).Elem()
}

func (o GroupingIdResponseOutput) ToGroupingIdResponseOutput() GroupingIdResponseOutput {
	return o
}

func (o GroupingIdResponseOutput) ToGroupingIdResponseOutputWithContext(ctx context.Context) GroupingIdResponseOutput {
	return o
}

// GroupingId type. It is a required property. More types of groupIds can be supported in future.
func (o GroupingIdResponseOutput) GroupingIdType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GroupingIdResponse) *string { return v.GroupingIdType }).(pulumi.StringPtrOutput)
}

// GroupId value based on the groupingType selected - Billing Id or ServiceTreeId.
func (o GroupingIdResponseOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GroupingIdResponse) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type GroupingIdResponsePtrOutput struct{ *pulumi.OutputState }

func (GroupingIdResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupingIdResponse)(nil)).Elem()
}

func (o GroupingIdResponsePtrOutput) ToGroupingIdResponsePtrOutput() GroupingIdResponsePtrOutput {
	return o
}

func (o GroupingIdResponsePtrOutput) ToGroupingIdResponsePtrOutputWithContext(ctx context.Context) GroupingIdResponsePtrOutput {
	return o
}

func (o GroupingIdResponsePtrOutput) Elem() GroupingIdResponseOutput {
	return o.ApplyT(func(v *GroupingIdResponse) GroupingIdResponse {
		if v != nil {
			return *v
		}
		var ret GroupingIdResponse
		return ret
	}).(GroupingIdResponseOutput)
}

// GroupingId type. It is a required property. More types of groupIds can be supported in future.
func (o GroupingIdResponsePtrOutput) GroupingIdType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GroupingIdResponse) *string {
		if v == nil {
			return nil
		}
		return v.GroupingIdType
	}).(pulumi.StringPtrOutput)
}

// GroupId value based on the groupingType selected - Billing Id or ServiceTreeId.
func (o GroupingIdResponsePtrOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GroupingIdResponse) *string {
		if v == nil {
			return nil
		}
		return v.Value
	}).(pulumi.StringPtrOutput)
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
	pulumi.RegisterOutputType(AdditionalAttributesOutput{})
	pulumi.RegisterOutputType(AdditionalAttributesPtrOutput{})
	pulumi.RegisterOutputType(AdditionalAttributesResponseOutput{})
	pulumi.RegisterOutputType(AdditionalAttributesResponsePtrOutput{})
	pulumi.RegisterOutputType(GroupQuotaSubscriptionIdResponsePropertiesOutput{})
	pulumi.RegisterOutputType(GroupQuotasEntityBaseOutput{})
	pulumi.RegisterOutputType(GroupQuotasEntityBasePtrOutput{})
	pulumi.RegisterOutputType(GroupQuotasEntityBaseResponseOutput{})
	pulumi.RegisterOutputType(GroupingIdOutput{})
	pulumi.RegisterOutputType(GroupingIdPtrOutput{})
	pulumi.RegisterOutputType(GroupingIdResponseOutput{})
	pulumi.RegisterOutputType(GroupingIdResponsePtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
}
