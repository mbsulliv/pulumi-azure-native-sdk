// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package web

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Workflow properties definition.
// Azure REST API version: 2022-09-01.
func ListWebAppWorkflowsConnectionsSlot(ctx *pulumi.Context, args *ListWebAppWorkflowsConnectionsSlotArgs, opts ...pulumi.InvokeOption) (*ListWebAppWorkflowsConnectionsSlotResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListWebAppWorkflowsConnectionsSlotResult
	err := ctx.Invoke("azure-native:web:listWebAppWorkflowsConnectionsSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWebAppWorkflowsConnectionsSlotArgs struct {
	// Site name.
	Name string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the deployment slot.
	Slot string `pulumi:"slot"`
}

// Workflow properties definition.
type ListWebAppWorkflowsConnectionsSlotResult struct {
	// The resource id.
	Id string `pulumi:"id"`
	// The resource kind.
	Kind *string `pulumi:"kind"`
	// The resource location.
	Location *string `pulumi:"location"`
	// Gets the resource name.
	Name string `pulumi:"name"`
	// Additional workflow properties.
	Properties WorkflowEnvelopeResponseProperties `pulumi:"properties"`
	// Gets the resource type.
	Type string `pulumi:"type"`
}

func ListWebAppWorkflowsConnectionsSlotOutput(ctx *pulumi.Context, args ListWebAppWorkflowsConnectionsSlotOutputArgs, opts ...pulumi.InvokeOption) ListWebAppWorkflowsConnectionsSlotResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListWebAppWorkflowsConnectionsSlotResult, error) {
			args := v.(ListWebAppWorkflowsConnectionsSlotArgs)
			r, err := ListWebAppWorkflowsConnectionsSlot(ctx, &args, opts...)
			var s ListWebAppWorkflowsConnectionsSlotResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListWebAppWorkflowsConnectionsSlotResultOutput)
}

type ListWebAppWorkflowsConnectionsSlotOutputArgs struct {
	// Site name.
	Name pulumi.StringInput `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name of the deployment slot.
	Slot pulumi.StringInput `pulumi:"slot"`
}

func (ListWebAppWorkflowsConnectionsSlotOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppWorkflowsConnectionsSlotArgs)(nil)).Elem()
}

// Workflow properties definition.
type ListWebAppWorkflowsConnectionsSlotResultOutput struct{ *pulumi.OutputState }

func (ListWebAppWorkflowsConnectionsSlotResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppWorkflowsConnectionsSlotResult)(nil)).Elem()
}

func (o ListWebAppWorkflowsConnectionsSlotResultOutput) ToListWebAppWorkflowsConnectionsSlotResultOutput() ListWebAppWorkflowsConnectionsSlotResultOutput {
	return o
}

func (o ListWebAppWorkflowsConnectionsSlotResultOutput) ToListWebAppWorkflowsConnectionsSlotResultOutputWithContext(ctx context.Context) ListWebAppWorkflowsConnectionsSlotResultOutput {
	return o
}

func (o ListWebAppWorkflowsConnectionsSlotResultOutput) ToOutput(ctx context.Context) pulumix.Output[ListWebAppWorkflowsConnectionsSlotResult] {
	return pulumix.Output[ListWebAppWorkflowsConnectionsSlotResult]{
		OutputState: o.OutputState,
	}
}

// The resource id.
func (o ListWebAppWorkflowsConnectionsSlotResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppWorkflowsConnectionsSlotResult) string { return v.Id }).(pulumi.StringOutput)
}

// The resource kind.
func (o ListWebAppWorkflowsConnectionsSlotResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppWorkflowsConnectionsSlotResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// The resource location.
func (o ListWebAppWorkflowsConnectionsSlotResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppWorkflowsConnectionsSlotResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// Gets the resource name.
func (o ListWebAppWorkflowsConnectionsSlotResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppWorkflowsConnectionsSlotResult) string { return v.Name }).(pulumi.StringOutput)
}

// Additional workflow properties.
func (o ListWebAppWorkflowsConnectionsSlotResultOutput) Properties() WorkflowEnvelopeResponsePropertiesOutput {
	return o.ApplyT(func(v ListWebAppWorkflowsConnectionsSlotResult) WorkflowEnvelopeResponseProperties {
		return v.Properties
	}).(WorkflowEnvelopeResponsePropertiesOutput)
}

// Gets the resource type.
func (o ListWebAppWorkflowsConnectionsSlotResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppWorkflowsConnectionsSlotResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListWebAppWorkflowsConnectionsSlotResultOutput{})
}
