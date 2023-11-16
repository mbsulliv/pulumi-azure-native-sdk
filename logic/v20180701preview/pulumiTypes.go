// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20180701preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = utilities.GetEnvOrDefault

// The azure resource error info.
type AzureResourceErrorInfoResponse struct {
	// The error code.
	Code string `pulumi:"code"`
	// The error details.
	Details []AzureResourceErrorInfoResponse `pulumi:"details"`
	// The error message.
	Message string `pulumi:"message"`
}

// The azure resource error info.
type AzureResourceErrorInfoResponseOutput struct{ *pulumi.OutputState }

func (AzureResourceErrorInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureResourceErrorInfoResponse)(nil)).Elem()
}

func (o AzureResourceErrorInfoResponseOutput) ToAzureResourceErrorInfoResponseOutput() AzureResourceErrorInfoResponseOutput {
	return o
}

func (o AzureResourceErrorInfoResponseOutput) ToAzureResourceErrorInfoResponseOutputWithContext(ctx context.Context) AzureResourceErrorInfoResponseOutput {
	return o
}

// The error code.
func (o AzureResourceErrorInfoResponseOutput) Code() pulumi.StringOutput {
	return o.ApplyT(func(v AzureResourceErrorInfoResponse) string { return v.Code }).(pulumi.StringOutput)
}

// The error details.
func (o AzureResourceErrorInfoResponseOutput) Details() AzureResourceErrorInfoResponseArrayOutput {
	return o.ApplyT(func(v AzureResourceErrorInfoResponse) []AzureResourceErrorInfoResponse { return v.Details }).(AzureResourceErrorInfoResponseArrayOutput)
}

// The error message.
func (o AzureResourceErrorInfoResponseOutput) Message() pulumi.StringOutput {
	return o.ApplyT(func(v AzureResourceErrorInfoResponse) string { return v.Message }).(pulumi.StringOutput)
}

type AzureResourceErrorInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (AzureResourceErrorInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureResourceErrorInfoResponse)(nil)).Elem()
}

func (o AzureResourceErrorInfoResponsePtrOutput) ToAzureResourceErrorInfoResponsePtrOutput() AzureResourceErrorInfoResponsePtrOutput {
	return o
}

func (o AzureResourceErrorInfoResponsePtrOutput) ToAzureResourceErrorInfoResponsePtrOutputWithContext(ctx context.Context) AzureResourceErrorInfoResponsePtrOutput {
	return o
}

func (o AzureResourceErrorInfoResponsePtrOutput) Elem() AzureResourceErrorInfoResponseOutput {
	return o.ApplyT(func(v *AzureResourceErrorInfoResponse) AzureResourceErrorInfoResponse {
		if v != nil {
			return *v
		}
		var ret AzureResourceErrorInfoResponse
		return ret
	}).(AzureResourceErrorInfoResponseOutput)
}

// The error code.
func (o AzureResourceErrorInfoResponsePtrOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureResourceErrorInfoResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Code
	}).(pulumi.StringPtrOutput)
}

// The error details.
func (o AzureResourceErrorInfoResponsePtrOutput) Details() AzureResourceErrorInfoResponseArrayOutput {
	return o.ApplyT(func(v *AzureResourceErrorInfoResponse) []AzureResourceErrorInfoResponse {
		if v == nil {
			return nil
		}
		return v.Details
	}).(AzureResourceErrorInfoResponseArrayOutput)
}

// The error message.
func (o AzureResourceErrorInfoResponsePtrOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureResourceErrorInfoResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Message
	}).(pulumi.StringPtrOutput)
}

type AzureResourceErrorInfoResponseArrayOutput struct{ *pulumi.OutputState }

func (AzureResourceErrorInfoResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AzureResourceErrorInfoResponse)(nil)).Elem()
}

func (o AzureResourceErrorInfoResponseArrayOutput) ToAzureResourceErrorInfoResponseArrayOutput() AzureResourceErrorInfoResponseArrayOutput {
	return o
}

func (o AzureResourceErrorInfoResponseArrayOutput) ToAzureResourceErrorInfoResponseArrayOutputWithContext(ctx context.Context) AzureResourceErrorInfoResponseArrayOutput {
	return o
}

func (o AzureResourceErrorInfoResponseArrayOutput) Index(i pulumi.IntInput) AzureResourceErrorInfoResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AzureResourceErrorInfoResponse {
		return vs[0].([]AzureResourceErrorInfoResponse)[vs[1].(int)]
	}).(AzureResourceErrorInfoResponseOutput)
}

type ExpressionResponse struct {
	// The azure resource error info.
	Error          *AzureResourceErrorInfoResponse `pulumi:"error"`
	Subexpressions []ExpressionResponse            `pulumi:"subexpressions"`
	Text           *string                         `pulumi:"text"`
	Value          interface{}                     `pulumi:"value"`
}

type ExpressionResponseOutput struct{ *pulumi.OutputState }

func (ExpressionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExpressionResponse)(nil)).Elem()
}

func (o ExpressionResponseOutput) ToExpressionResponseOutput() ExpressionResponseOutput {
	return o
}

func (o ExpressionResponseOutput) ToExpressionResponseOutputWithContext(ctx context.Context) ExpressionResponseOutput {
	return o
}

// The azure resource error info.
func (o ExpressionResponseOutput) Error() AzureResourceErrorInfoResponsePtrOutput {
	return o.ApplyT(func(v ExpressionResponse) *AzureResourceErrorInfoResponse { return v.Error }).(AzureResourceErrorInfoResponsePtrOutput)
}

func (o ExpressionResponseOutput) Subexpressions() ExpressionResponseArrayOutput {
	return o.ApplyT(func(v ExpressionResponse) []ExpressionResponse { return v.Subexpressions }).(ExpressionResponseArrayOutput)
}

func (o ExpressionResponseOutput) Text() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressionResponse) *string { return v.Text }).(pulumi.StringPtrOutput)
}

func (o ExpressionResponseOutput) Value() pulumi.AnyOutput {
	return o.ApplyT(func(v ExpressionResponse) interface{} { return v.Value }).(pulumi.AnyOutput)
}

type ExpressionResponseArrayOutput struct{ *pulumi.OutputState }

func (ExpressionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ExpressionResponse)(nil)).Elem()
}

func (o ExpressionResponseArrayOutput) ToExpressionResponseArrayOutput() ExpressionResponseArrayOutput {
	return o
}

func (o ExpressionResponseArrayOutput) ToExpressionResponseArrayOutputWithContext(ctx context.Context) ExpressionResponseArrayOutput {
	return o
}

func (o ExpressionResponseArrayOutput) Index(i pulumi.IntInput) ExpressionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ExpressionResponse {
		return vs[0].([]ExpressionResponse)[vs[1].(int)]
	}).(ExpressionResponseOutput)
}

type ExpressionRootResponse struct {
	// The azure resource error info.
	Error *AzureResourceErrorInfoResponse `pulumi:"error"`
	// The path.
	Path           *string              `pulumi:"path"`
	Subexpressions []ExpressionResponse `pulumi:"subexpressions"`
	Text           *string              `pulumi:"text"`
	Value          interface{}          `pulumi:"value"`
}

type ExpressionRootResponseOutput struct{ *pulumi.OutputState }

func (ExpressionRootResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExpressionRootResponse)(nil)).Elem()
}

func (o ExpressionRootResponseOutput) ToExpressionRootResponseOutput() ExpressionRootResponseOutput {
	return o
}

func (o ExpressionRootResponseOutput) ToExpressionRootResponseOutputWithContext(ctx context.Context) ExpressionRootResponseOutput {
	return o
}

// The azure resource error info.
func (o ExpressionRootResponseOutput) Error() AzureResourceErrorInfoResponsePtrOutput {
	return o.ApplyT(func(v ExpressionRootResponse) *AzureResourceErrorInfoResponse { return v.Error }).(AzureResourceErrorInfoResponsePtrOutput)
}

// The path.
func (o ExpressionRootResponseOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressionRootResponse) *string { return v.Path }).(pulumi.StringPtrOutput)
}

func (o ExpressionRootResponseOutput) Subexpressions() ExpressionResponseArrayOutput {
	return o.ApplyT(func(v ExpressionRootResponse) []ExpressionResponse { return v.Subexpressions }).(ExpressionResponseArrayOutput)
}

func (o ExpressionRootResponseOutput) Text() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressionRootResponse) *string { return v.Text }).(pulumi.StringPtrOutput)
}

func (o ExpressionRootResponseOutput) Value() pulumi.AnyOutput {
	return o.ApplyT(func(v ExpressionRootResponse) interface{} { return v.Value }).(pulumi.AnyOutput)
}

type ExpressionRootResponseArrayOutput struct{ *pulumi.OutputState }

func (ExpressionRootResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ExpressionRootResponse)(nil)).Elem()
}

func (o ExpressionRootResponseArrayOutput) ToExpressionRootResponseArrayOutput() ExpressionRootResponseArrayOutput {
	return o
}

func (o ExpressionRootResponseArrayOutput) ToExpressionRootResponseArrayOutputWithContext(ctx context.Context) ExpressionRootResponseArrayOutput {
	return o
}

func (o ExpressionRootResponseArrayOutput) Index(i pulumi.IntInput) ExpressionRootResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ExpressionRootResponse {
		return vs[0].([]ExpressionRootResponse)[vs[1].(int)]
	}).(ExpressionRootResponseOutput)
}

// The resource reference.
type ResourceReference struct {
	// The resource id.
	Id *string `pulumi:"id"`
}

// ResourceReferenceInput is an input type that accepts ResourceReferenceArgs and ResourceReferenceOutput values.
// You can construct a concrete instance of `ResourceReferenceInput` via:
//
//	ResourceReferenceArgs{...}
type ResourceReferenceInput interface {
	pulumi.Input

	ToResourceReferenceOutput() ResourceReferenceOutput
	ToResourceReferenceOutputWithContext(context.Context) ResourceReferenceOutput
}

// The resource reference.
type ResourceReferenceArgs struct {
	// The resource id.
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (ResourceReferenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceReference)(nil)).Elem()
}

func (i ResourceReferenceArgs) ToResourceReferenceOutput() ResourceReferenceOutput {
	return i.ToResourceReferenceOutputWithContext(context.Background())
}

func (i ResourceReferenceArgs) ToResourceReferenceOutputWithContext(ctx context.Context) ResourceReferenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceReferenceOutput)
}

func (i ResourceReferenceArgs) ToResourceReferencePtrOutput() ResourceReferencePtrOutput {
	return i.ToResourceReferencePtrOutputWithContext(context.Background())
}

func (i ResourceReferenceArgs) ToResourceReferencePtrOutputWithContext(ctx context.Context) ResourceReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceReferenceOutput).ToResourceReferencePtrOutputWithContext(ctx)
}

// ResourceReferencePtrInput is an input type that accepts ResourceReferenceArgs, ResourceReferencePtr and ResourceReferencePtrOutput values.
// You can construct a concrete instance of `ResourceReferencePtrInput` via:
//
//	        ResourceReferenceArgs{...}
//
//	or:
//
//	        nil
type ResourceReferencePtrInput interface {
	pulumi.Input

	ToResourceReferencePtrOutput() ResourceReferencePtrOutput
	ToResourceReferencePtrOutputWithContext(context.Context) ResourceReferencePtrOutput
}

type resourceReferencePtrType ResourceReferenceArgs

func ResourceReferencePtr(v *ResourceReferenceArgs) ResourceReferencePtrInput {
	return (*resourceReferencePtrType)(v)
}

func (*resourceReferencePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceReference)(nil)).Elem()
}

func (i *resourceReferencePtrType) ToResourceReferencePtrOutput() ResourceReferencePtrOutput {
	return i.ToResourceReferencePtrOutputWithContext(context.Background())
}

func (i *resourceReferencePtrType) ToResourceReferencePtrOutputWithContext(ctx context.Context) ResourceReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceReferencePtrOutput)
}

// The resource reference.
type ResourceReferenceOutput struct{ *pulumi.OutputState }

func (ResourceReferenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceReference)(nil)).Elem()
}

func (o ResourceReferenceOutput) ToResourceReferenceOutput() ResourceReferenceOutput {
	return o
}

func (o ResourceReferenceOutput) ToResourceReferenceOutputWithContext(ctx context.Context) ResourceReferenceOutput {
	return o
}

func (o ResourceReferenceOutput) ToResourceReferencePtrOutput() ResourceReferencePtrOutput {
	return o.ToResourceReferencePtrOutputWithContext(context.Background())
}

func (o ResourceReferenceOutput) ToResourceReferencePtrOutputWithContext(ctx context.Context) ResourceReferencePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceReference) *ResourceReference {
		return &v
	}).(ResourceReferencePtrOutput)
}

// The resource id.
func (o ResourceReferenceOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceReference) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type ResourceReferencePtrOutput struct{ *pulumi.OutputState }

func (ResourceReferencePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceReference)(nil)).Elem()
}

func (o ResourceReferencePtrOutput) ToResourceReferencePtrOutput() ResourceReferencePtrOutput {
	return o
}

func (o ResourceReferencePtrOutput) ToResourceReferencePtrOutputWithContext(ctx context.Context) ResourceReferencePtrOutput {
	return o
}

func (o ResourceReferencePtrOutput) Elem() ResourceReferenceOutput {
	return o.ApplyT(func(v *ResourceReference) ResourceReference {
		if v != nil {
			return *v
		}
		var ret ResourceReference
		return ret
	}).(ResourceReferenceOutput)
}

// The resource id.
func (o ResourceReferencePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceReference) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

// The resource reference.
type ResourceReferenceResponse struct {
	// The resource id.
	Id *string `pulumi:"id"`
	// Gets the resource name.
	Name string `pulumi:"name"`
	// Gets the resource type.
	Type string `pulumi:"type"`
}

// The resource reference.
type ResourceReferenceResponseOutput struct{ *pulumi.OutputState }

func (ResourceReferenceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceReferenceResponse)(nil)).Elem()
}

func (o ResourceReferenceResponseOutput) ToResourceReferenceResponseOutput() ResourceReferenceResponseOutput {
	return o
}

func (o ResourceReferenceResponseOutput) ToResourceReferenceResponseOutputWithContext(ctx context.Context) ResourceReferenceResponseOutput {
	return o
}

// The resource id.
func (o ResourceReferenceResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceReferenceResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// Gets the resource name.
func (o ResourceReferenceResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceReferenceResponse) string { return v.Name }).(pulumi.StringOutput)
}

// Gets the resource type.
func (o ResourceReferenceResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceReferenceResponse) string { return v.Type }).(pulumi.StringOutput)
}

type ResourceReferenceResponsePtrOutput struct{ *pulumi.OutputState }

func (ResourceReferenceResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceReferenceResponse)(nil)).Elem()
}

func (o ResourceReferenceResponsePtrOutput) ToResourceReferenceResponsePtrOutput() ResourceReferenceResponsePtrOutput {
	return o
}

func (o ResourceReferenceResponsePtrOutput) ToResourceReferenceResponsePtrOutputWithContext(ctx context.Context) ResourceReferenceResponsePtrOutput {
	return o
}

func (o ResourceReferenceResponsePtrOutput) Elem() ResourceReferenceResponseOutput {
	return o.ApplyT(func(v *ResourceReferenceResponse) ResourceReferenceResponse {
		if v != nil {
			return *v
		}
		var ret ResourceReferenceResponse
		return ret
	}).(ResourceReferenceResponseOutput)
}

// The resource id.
func (o ResourceReferenceResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

// Gets the resource name.
func (o ResourceReferenceResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

// Gets the resource type.
func (o ResourceReferenceResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

// The sku type.
type Sku struct {
	// The name.
	Name string `pulumi:"name"`
	// The reference to plan.
	Plan *ResourceReference `pulumi:"plan"`
}

// SkuInput is an input type that accepts SkuArgs and SkuOutput values.
// You can construct a concrete instance of `SkuInput` via:
//
//	SkuArgs{...}
type SkuInput interface {
	pulumi.Input

	ToSkuOutput() SkuOutput
	ToSkuOutputWithContext(context.Context) SkuOutput
}

// The sku type.
type SkuArgs struct {
	// The name.
	Name pulumi.StringInput `pulumi:"name"`
	// The reference to plan.
	Plan ResourceReferencePtrInput `pulumi:"plan"`
}

func (SkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Sku)(nil)).Elem()
}

func (i SkuArgs) ToSkuOutput() SkuOutput {
	return i.ToSkuOutputWithContext(context.Background())
}

func (i SkuArgs) ToSkuOutputWithContext(ctx context.Context) SkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuOutput)
}

func (i SkuArgs) ToSkuPtrOutput() SkuPtrOutput {
	return i.ToSkuPtrOutputWithContext(context.Background())
}

func (i SkuArgs) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuOutput).ToSkuPtrOutputWithContext(ctx)
}

// SkuPtrInput is an input type that accepts SkuArgs, SkuPtr and SkuPtrOutput values.
// You can construct a concrete instance of `SkuPtrInput` via:
//
//	        SkuArgs{...}
//
//	or:
//
//	        nil
type SkuPtrInput interface {
	pulumi.Input

	ToSkuPtrOutput() SkuPtrOutput
	ToSkuPtrOutputWithContext(context.Context) SkuPtrOutput
}

type skuPtrType SkuArgs

func SkuPtr(v *SkuArgs) SkuPtrInput {
	return (*skuPtrType)(v)
}

func (*skuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Sku)(nil)).Elem()
}

func (i *skuPtrType) ToSkuPtrOutput() SkuPtrOutput {
	return i.ToSkuPtrOutputWithContext(context.Background())
}

func (i *skuPtrType) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuPtrOutput)
}

// The sku type.
type SkuOutput struct{ *pulumi.OutputState }

func (SkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Sku)(nil)).Elem()
}

func (o SkuOutput) ToSkuOutput() SkuOutput {
	return o
}

func (o SkuOutput) ToSkuOutputWithContext(ctx context.Context) SkuOutput {
	return o
}

func (o SkuOutput) ToSkuPtrOutput() SkuPtrOutput {
	return o.ToSkuPtrOutputWithContext(context.Background())
}

func (o SkuOutput) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Sku) *Sku {
		return &v
	}).(SkuPtrOutput)
}

// The name.
func (o SkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v Sku) string { return v.Name }).(pulumi.StringOutput)
}

// The reference to plan.
func (o SkuOutput) Plan() ResourceReferencePtrOutput {
	return o.ApplyT(func(v Sku) *ResourceReference { return v.Plan }).(ResourceReferencePtrOutput)
}

type SkuPtrOutput struct{ *pulumi.OutputState }

func (SkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Sku)(nil)).Elem()
}

func (o SkuPtrOutput) ToSkuPtrOutput() SkuPtrOutput {
	return o
}

func (o SkuPtrOutput) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return o
}

func (o SkuPtrOutput) Elem() SkuOutput {
	return o.ApplyT(func(v *Sku) Sku {
		if v != nil {
			return *v
		}
		var ret Sku
		return ret
	}).(SkuOutput)
}

// The name.
func (o SkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

// The reference to plan.
func (o SkuPtrOutput) Plan() ResourceReferencePtrOutput {
	return o.ApplyT(func(v *Sku) *ResourceReference {
		if v == nil {
			return nil
		}
		return v.Plan
	}).(ResourceReferencePtrOutput)
}

// The sku type.
type SkuResponse struct {
	// The name.
	Name string `pulumi:"name"`
	// The reference to plan.
	Plan *ResourceReferenceResponse `pulumi:"plan"`
}

// The sku type.
type SkuResponseOutput struct{ *pulumi.OutputState }

func (SkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuResponse)(nil)).Elem()
}

func (o SkuResponseOutput) ToSkuResponseOutput() SkuResponseOutput {
	return o
}

func (o SkuResponseOutput) ToSkuResponseOutputWithContext(ctx context.Context) SkuResponseOutput {
	return o
}

// The name.
func (o SkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

// The reference to plan.
func (o SkuResponseOutput) Plan() ResourceReferenceResponsePtrOutput {
	return o.ApplyT(func(v SkuResponse) *ResourceReferenceResponse { return v.Plan }).(ResourceReferenceResponsePtrOutput)
}

type SkuResponsePtrOutput struct{ *pulumi.OutputState }

func (SkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuResponse)(nil)).Elem()
}

func (o SkuResponsePtrOutput) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return o
}

func (o SkuResponsePtrOutput) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return o
}

func (o SkuResponsePtrOutput) Elem() SkuResponseOutput {
	return o.ApplyT(func(v *SkuResponse) SkuResponse {
		if v != nil {
			return *v
		}
		var ret SkuResponse
		return ret
	}).(SkuResponseOutput)
}

// The name.
func (o SkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

// The reference to plan.
func (o SkuResponsePtrOutput) Plan() ResourceReferenceResponsePtrOutput {
	return o.ApplyT(func(v *SkuResponse) *ResourceReferenceResponse {
		if v == nil {
			return nil
		}
		return v.Plan
	}).(ResourceReferenceResponsePtrOutput)
}

// The workflow parameters.
type WorkflowParameter struct {
	// The description.
	Description *string `pulumi:"description"`
	// The metadata.
	Metadata interface{} `pulumi:"metadata"`
	// The type.
	Type *string `pulumi:"type"`
	// The value.
	Value interface{} `pulumi:"value"`
}

// WorkflowParameterInput is an input type that accepts WorkflowParameterArgs and WorkflowParameterOutput values.
// You can construct a concrete instance of `WorkflowParameterInput` via:
//
//	WorkflowParameterArgs{...}
type WorkflowParameterInput interface {
	pulumi.Input

	ToWorkflowParameterOutput() WorkflowParameterOutput
	ToWorkflowParameterOutputWithContext(context.Context) WorkflowParameterOutput
}

// The workflow parameters.
type WorkflowParameterArgs struct {
	// The description.
	Description pulumi.StringPtrInput `pulumi:"description"`
	// The metadata.
	Metadata pulumi.Input `pulumi:"metadata"`
	// The type.
	Type pulumi.StringPtrInput `pulumi:"type"`
	// The value.
	Value pulumi.Input `pulumi:"value"`
}

func (WorkflowParameterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkflowParameter)(nil)).Elem()
}

func (i WorkflowParameterArgs) ToWorkflowParameterOutput() WorkflowParameterOutput {
	return i.ToWorkflowParameterOutputWithContext(context.Background())
}

func (i WorkflowParameterArgs) ToWorkflowParameterOutputWithContext(ctx context.Context) WorkflowParameterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkflowParameterOutput)
}

// WorkflowParameterMapInput is an input type that accepts WorkflowParameterMap and WorkflowParameterMapOutput values.
// You can construct a concrete instance of `WorkflowParameterMapInput` via:
//
//	WorkflowParameterMap{ "key": WorkflowParameterArgs{...} }
type WorkflowParameterMapInput interface {
	pulumi.Input

	ToWorkflowParameterMapOutput() WorkflowParameterMapOutput
	ToWorkflowParameterMapOutputWithContext(context.Context) WorkflowParameterMapOutput
}

type WorkflowParameterMap map[string]WorkflowParameterInput

func (WorkflowParameterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]WorkflowParameter)(nil)).Elem()
}

func (i WorkflowParameterMap) ToWorkflowParameterMapOutput() WorkflowParameterMapOutput {
	return i.ToWorkflowParameterMapOutputWithContext(context.Background())
}

func (i WorkflowParameterMap) ToWorkflowParameterMapOutputWithContext(ctx context.Context) WorkflowParameterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkflowParameterMapOutput)
}

// The workflow parameters.
type WorkflowParameterOutput struct{ *pulumi.OutputState }

func (WorkflowParameterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkflowParameter)(nil)).Elem()
}

func (o WorkflowParameterOutput) ToWorkflowParameterOutput() WorkflowParameterOutput {
	return o
}

func (o WorkflowParameterOutput) ToWorkflowParameterOutputWithContext(ctx context.Context) WorkflowParameterOutput {
	return o
}

// The description.
func (o WorkflowParameterOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkflowParameter) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The metadata.
func (o WorkflowParameterOutput) Metadata() pulumi.AnyOutput {
	return o.ApplyT(func(v WorkflowParameter) interface{} { return v.Metadata }).(pulumi.AnyOutput)
}

// The type.
func (o WorkflowParameterOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkflowParameter) *string { return v.Type }).(pulumi.StringPtrOutput)
}

// The value.
func (o WorkflowParameterOutput) Value() pulumi.AnyOutput {
	return o.ApplyT(func(v WorkflowParameter) interface{} { return v.Value }).(pulumi.AnyOutput)
}

type WorkflowParameterMapOutput struct{ *pulumi.OutputState }

func (WorkflowParameterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]WorkflowParameter)(nil)).Elem()
}

func (o WorkflowParameterMapOutput) ToWorkflowParameterMapOutput() WorkflowParameterMapOutput {
	return o
}

func (o WorkflowParameterMapOutput) ToWorkflowParameterMapOutputWithContext(ctx context.Context) WorkflowParameterMapOutput {
	return o
}

func (o WorkflowParameterMapOutput) MapIndex(k pulumi.StringInput) WorkflowParameterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) WorkflowParameter {
		return vs[0].(map[string]WorkflowParameter)[vs[1].(string)]
	}).(WorkflowParameterOutput)
}

// The workflow parameters.
type WorkflowParameterResponse struct {
	// The description.
	Description *string `pulumi:"description"`
	// The metadata.
	Metadata interface{} `pulumi:"metadata"`
	// The type.
	Type *string `pulumi:"type"`
	// The value.
	Value interface{} `pulumi:"value"`
}

// The workflow parameters.
type WorkflowParameterResponseOutput struct{ *pulumi.OutputState }

func (WorkflowParameterResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkflowParameterResponse)(nil)).Elem()
}

func (o WorkflowParameterResponseOutput) ToWorkflowParameterResponseOutput() WorkflowParameterResponseOutput {
	return o
}

func (o WorkflowParameterResponseOutput) ToWorkflowParameterResponseOutputWithContext(ctx context.Context) WorkflowParameterResponseOutput {
	return o
}

// The description.
func (o WorkflowParameterResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkflowParameterResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The metadata.
func (o WorkflowParameterResponseOutput) Metadata() pulumi.AnyOutput {
	return o.ApplyT(func(v WorkflowParameterResponse) interface{} { return v.Metadata }).(pulumi.AnyOutput)
}

// The type.
func (o WorkflowParameterResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkflowParameterResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

// The value.
func (o WorkflowParameterResponseOutput) Value() pulumi.AnyOutput {
	return o.ApplyT(func(v WorkflowParameterResponse) interface{} { return v.Value }).(pulumi.AnyOutput)
}

type WorkflowParameterResponseMapOutput struct{ *pulumi.OutputState }

func (WorkflowParameterResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]WorkflowParameterResponse)(nil)).Elem()
}

func (o WorkflowParameterResponseMapOutput) ToWorkflowParameterResponseMapOutput() WorkflowParameterResponseMapOutput {
	return o
}

func (o WorkflowParameterResponseMapOutput) ToWorkflowParameterResponseMapOutputWithContext(ctx context.Context) WorkflowParameterResponseMapOutput {
	return o
}

func (o WorkflowParameterResponseMapOutput) MapIndex(k pulumi.StringInput) WorkflowParameterResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) WorkflowParameterResponse {
		return vs[0].(map[string]WorkflowParameterResponse)[vs[1].(string)]
	}).(WorkflowParameterResponseOutput)
}

// Gets the workflow trigger callback URL query parameters.
type WorkflowTriggerListCallbackUrlQueriesResponse struct {
	// The api version.
	ApiVersion *string `pulumi:"apiVersion"`
	// The SAS timestamp.
	Se *string `pulumi:"se"`
	// The SAS signature.
	Sig *string `pulumi:"sig"`
	// The SAS permissions.
	Sp *string `pulumi:"sp"`
	// The SAS version.
	Sv *string `pulumi:"sv"`
}

// Gets the workflow trigger callback URL query parameters.
type WorkflowTriggerListCallbackUrlQueriesResponseOutput struct{ *pulumi.OutputState }

func (WorkflowTriggerListCallbackUrlQueriesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkflowTriggerListCallbackUrlQueriesResponse)(nil)).Elem()
}

func (o WorkflowTriggerListCallbackUrlQueriesResponseOutput) ToWorkflowTriggerListCallbackUrlQueriesResponseOutput() WorkflowTriggerListCallbackUrlQueriesResponseOutput {
	return o
}

func (o WorkflowTriggerListCallbackUrlQueriesResponseOutput) ToWorkflowTriggerListCallbackUrlQueriesResponseOutputWithContext(ctx context.Context) WorkflowTriggerListCallbackUrlQueriesResponseOutput {
	return o
}

// The api version.
func (o WorkflowTriggerListCallbackUrlQueriesResponseOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkflowTriggerListCallbackUrlQueriesResponse) *string { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

// The SAS timestamp.
func (o WorkflowTriggerListCallbackUrlQueriesResponseOutput) Se() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkflowTriggerListCallbackUrlQueriesResponse) *string { return v.Se }).(pulumi.StringPtrOutput)
}

// The SAS signature.
func (o WorkflowTriggerListCallbackUrlQueriesResponseOutput) Sig() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkflowTriggerListCallbackUrlQueriesResponse) *string { return v.Sig }).(pulumi.StringPtrOutput)
}

// The SAS permissions.
func (o WorkflowTriggerListCallbackUrlQueriesResponseOutput) Sp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkflowTriggerListCallbackUrlQueriesResponse) *string { return v.Sp }).(pulumi.StringPtrOutput)
}

// The SAS version.
func (o WorkflowTriggerListCallbackUrlQueriesResponseOutput) Sv() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkflowTriggerListCallbackUrlQueriesResponse) *string { return v.Sv }).(pulumi.StringPtrOutput)
}

type WorkflowTriggerListCallbackUrlQueriesResponsePtrOutput struct{ *pulumi.OutputState }

func (WorkflowTriggerListCallbackUrlQueriesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkflowTriggerListCallbackUrlQueriesResponse)(nil)).Elem()
}

func (o WorkflowTriggerListCallbackUrlQueriesResponsePtrOutput) ToWorkflowTriggerListCallbackUrlQueriesResponsePtrOutput() WorkflowTriggerListCallbackUrlQueriesResponsePtrOutput {
	return o
}

func (o WorkflowTriggerListCallbackUrlQueriesResponsePtrOutput) ToWorkflowTriggerListCallbackUrlQueriesResponsePtrOutputWithContext(ctx context.Context) WorkflowTriggerListCallbackUrlQueriesResponsePtrOutput {
	return o
}

func (o WorkflowTriggerListCallbackUrlQueriesResponsePtrOutput) Elem() WorkflowTriggerListCallbackUrlQueriesResponseOutput {
	return o.ApplyT(func(v *WorkflowTriggerListCallbackUrlQueriesResponse) WorkflowTriggerListCallbackUrlQueriesResponse {
		if v != nil {
			return *v
		}
		var ret WorkflowTriggerListCallbackUrlQueriesResponse
		return ret
	}).(WorkflowTriggerListCallbackUrlQueriesResponseOutput)
}

// The api version.
func (o WorkflowTriggerListCallbackUrlQueriesResponsePtrOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkflowTriggerListCallbackUrlQueriesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ApiVersion
	}).(pulumi.StringPtrOutput)
}

// The SAS timestamp.
func (o WorkflowTriggerListCallbackUrlQueriesResponsePtrOutput) Se() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkflowTriggerListCallbackUrlQueriesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Se
	}).(pulumi.StringPtrOutput)
}

// The SAS signature.
func (o WorkflowTriggerListCallbackUrlQueriesResponsePtrOutput) Sig() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkflowTriggerListCallbackUrlQueriesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Sig
	}).(pulumi.StringPtrOutput)
}

// The SAS permissions.
func (o WorkflowTriggerListCallbackUrlQueriesResponsePtrOutput) Sp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkflowTriggerListCallbackUrlQueriesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Sp
	}).(pulumi.StringPtrOutput)
}

// The SAS version.
func (o WorkflowTriggerListCallbackUrlQueriesResponsePtrOutput) Sv() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkflowTriggerListCallbackUrlQueriesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Sv
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AzureResourceErrorInfoResponseOutput{})
	pulumi.RegisterOutputType(AzureResourceErrorInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(AzureResourceErrorInfoResponseArrayOutput{})
	pulumi.RegisterOutputType(ExpressionResponseOutput{})
	pulumi.RegisterOutputType(ExpressionResponseArrayOutput{})
	pulumi.RegisterOutputType(ExpressionRootResponseOutput{})
	pulumi.RegisterOutputType(ExpressionRootResponseArrayOutput{})
	pulumi.RegisterOutputType(ResourceReferenceOutput{})
	pulumi.RegisterOutputType(ResourceReferencePtrOutput{})
	pulumi.RegisterOutputType(ResourceReferenceResponseOutput{})
	pulumi.RegisterOutputType(ResourceReferenceResponsePtrOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuPtrOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
	pulumi.RegisterOutputType(WorkflowParameterOutput{})
	pulumi.RegisterOutputType(WorkflowParameterMapOutput{})
	pulumi.RegisterOutputType(WorkflowParameterResponseOutput{})
	pulumi.RegisterOutputType(WorkflowParameterResponseMapOutput{})
	pulumi.RegisterOutputType(WorkflowTriggerListCallbackUrlQueriesResponseOutput{})
	pulumi.RegisterOutputType(WorkflowTriggerListCallbackUrlQueriesResponsePtrOutput{})
}
