// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20190401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

var _ = utilities.GetEnvOrDefault

type ControllerConnectionDetailsResponse struct {
	// Base class for types that supply values used to connect to container orchestrators
	OrchestratorSpecificConnectionDetails *KubernetesConnectionDetailsResponse `pulumi:"orchestratorSpecificConnectionDetails"`
}

type ControllerConnectionDetailsResponseOutput struct{ *pulumi.OutputState }

func (ControllerConnectionDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ControllerConnectionDetailsResponse)(nil)).Elem()
}

func (o ControllerConnectionDetailsResponseOutput) ToControllerConnectionDetailsResponseOutput() ControllerConnectionDetailsResponseOutput {
	return o
}

func (o ControllerConnectionDetailsResponseOutput) ToControllerConnectionDetailsResponseOutputWithContext(ctx context.Context) ControllerConnectionDetailsResponseOutput {
	return o
}

func (o ControllerConnectionDetailsResponseOutput) ToOutput(ctx context.Context) pulumix.Output[ControllerConnectionDetailsResponse] {
	return pulumix.Output[ControllerConnectionDetailsResponse]{
		OutputState: o.OutputState,
	}
}

// Base class for types that supply values used to connect to container orchestrators
func (o ControllerConnectionDetailsResponseOutput) OrchestratorSpecificConnectionDetails() KubernetesConnectionDetailsResponsePtrOutput {
	return o.ApplyT(func(v ControllerConnectionDetailsResponse) *KubernetesConnectionDetailsResponse {
		return v.OrchestratorSpecificConnectionDetails
	}).(KubernetesConnectionDetailsResponsePtrOutput)
}

type ControllerConnectionDetailsResponseArrayOutput struct{ *pulumi.OutputState }

func (ControllerConnectionDetailsResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ControllerConnectionDetailsResponse)(nil)).Elem()
}

func (o ControllerConnectionDetailsResponseArrayOutput) ToControllerConnectionDetailsResponseArrayOutput() ControllerConnectionDetailsResponseArrayOutput {
	return o
}

func (o ControllerConnectionDetailsResponseArrayOutput) ToControllerConnectionDetailsResponseArrayOutputWithContext(ctx context.Context) ControllerConnectionDetailsResponseArrayOutput {
	return o
}

func (o ControllerConnectionDetailsResponseArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]ControllerConnectionDetailsResponse] {
	return pulumix.Output[[]ControllerConnectionDetailsResponse]{
		OutputState: o.OutputState,
	}
}

func (o ControllerConnectionDetailsResponseArrayOutput) Index(i pulumi.IntInput) ControllerConnectionDetailsResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ControllerConnectionDetailsResponse {
		return vs[0].([]ControllerConnectionDetailsResponse)[vs[1].(int)]
	}).(ControllerConnectionDetailsResponseOutput)
}

// Contains information used to connect to a Kubernetes cluster
type KubernetesConnectionDetailsResponse struct {
	// Gets the Instance type.
	// Expected value is 'Kubernetes'.
	InstanceType string `pulumi:"instanceType"`
	// Gets the kubeconfig for the cluster.
	KubeConfig *string `pulumi:"kubeConfig"`
}

// Contains information used to connect to a Kubernetes cluster
type KubernetesConnectionDetailsResponseOutput struct{ *pulumi.OutputState }

func (KubernetesConnectionDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KubernetesConnectionDetailsResponse)(nil)).Elem()
}

func (o KubernetesConnectionDetailsResponseOutput) ToKubernetesConnectionDetailsResponseOutput() KubernetesConnectionDetailsResponseOutput {
	return o
}

func (o KubernetesConnectionDetailsResponseOutput) ToKubernetesConnectionDetailsResponseOutputWithContext(ctx context.Context) KubernetesConnectionDetailsResponseOutput {
	return o
}

func (o KubernetesConnectionDetailsResponseOutput) ToOutput(ctx context.Context) pulumix.Output[KubernetesConnectionDetailsResponse] {
	return pulumix.Output[KubernetesConnectionDetailsResponse]{
		OutputState: o.OutputState,
	}
}

// Gets the Instance type.
// Expected value is 'Kubernetes'.
func (o KubernetesConnectionDetailsResponseOutput) InstanceType() pulumi.StringOutput {
	return o.ApplyT(func(v KubernetesConnectionDetailsResponse) string { return v.InstanceType }).(pulumi.StringOutput)
}

// Gets the kubeconfig for the cluster.
func (o KubernetesConnectionDetailsResponseOutput) KubeConfig() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KubernetesConnectionDetailsResponse) *string { return v.KubeConfig }).(pulumi.StringPtrOutput)
}

type KubernetesConnectionDetailsResponsePtrOutput struct{ *pulumi.OutputState }

func (KubernetesConnectionDetailsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KubernetesConnectionDetailsResponse)(nil)).Elem()
}

func (o KubernetesConnectionDetailsResponsePtrOutput) ToKubernetesConnectionDetailsResponsePtrOutput() KubernetesConnectionDetailsResponsePtrOutput {
	return o
}

func (o KubernetesConnectionDetailsResponsePtrOutput) ToKubernetesConnectionDetailsResponsePtrOutputWithContext(ctx context.Context) KubernetesConnectionDetailsResponsePtrOutput {
	return o
}

func (o KubernetesConnectionDetailsResponsePtrOutput) ToOutput(ctx context.Context) pulumix.Output[*KubernetesConnectionDetailsResponse] {
	return pulumix.Output[*KubernetesConnectionDetailsResponse]{
		OutputState: o.OutputState,
	}
}

func (o KubernetesConnectionDetailsResponsePtrOutput) Elem() KubernetesConnectionDetailsResponseOutput {
	return o.ApplyT(func(v *KubernetesConnectionDetailsResponse) KubernetesConnectionDetailsResponse {
		if v != nil {
			return *v
		}
		var ret KubernetesConnectionDetailsResponse
		return ret
	}).(KubernetesConnectionDetailsResponseOutput)
}

// Gets the Instance type.
// Expected value is 'Kubernetes'.
func (o KubernetesConnectionDetailsResponsePtrOutput) InstanceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KubernetesConnectionDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.InstanceType
	}).(pulumi.StringPtrOutput)
}

// Gets the kubeconfig for the cluster.
func (o KubernetesConnectionDetailsResponsePtrOutput) KubeConfig() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KubernetesConnectionDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return v.KubeConfig
	}).(pulumi.StringPtrOutput)
}

// Model representing SKU for Azure Dev Spaces Controller.
type Sku struct {
	// The name of the SKU for Azure Dev Spaces Controller.
	Name string `pulumi:"name"`
	// The tier of the SKU for Azure Dev Spaces Controller.
	Tier *string `pulumi:"tier"`
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

// Model representing SKU for Azure Dev Spaces Controller.
type SkuArgs struct {
	// The name of the SKU for Azure Dev Spaces Controller.
	Name pulumi.StringInput `pulumi:"name"`
	// The tier of the SKU for Azure Dev Spaces Controller.
	Tier pulumi.StringPtrInput `pulumi:"tier"`
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

func (i SkuArgs) ToOutput(ctx context.Context) pulumix.Output[Sku] {
	return pulumix.Output[Sku]{
		OutputState: i.ToSkuOutputWithContext(ctx).OutputState,
	}
}

// Model representing SKU for Azure Dev Spaces Controller.
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

func (o SkuOutput) ToOutput(ctx context.Context) pulumix.Output[Sku] {
	return pulumix.Output[Sku]{
		OutputState: o.OutputState,
	}
}

// The name of the SKU for Azure Dev Spaces Controller.
func (o SkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v Sku) string { return v.Name }).(pulumi.StringOutput)
}

// The tier of the SKU for Azure Dev Spaces Controller.
func (o SkuOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Sku) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

// Model representing SKU for Azure Dev Spaces Controller.
type SkuResponse struct {
	// The name of the SKU for Azure Dev Spaces Controller.
	Name string `pulumi:"name"`
	// The tier of the SKU for Azure Dev Spaces Controller.
	Tier *string `pulumi:"tier"`
}

// Model representing SKU for Azure Dev Spaces Controller.
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

func (o SkuResponseOutput) ToOutput(ctx context.Context) pulumix.Output[SkuResponse] {
	return pulumix.Output[SkuResponse]{
		OutputState: o.OutputState,
	}
}

// The name of the SKU for Azure Dev Spaces Controller.
func (o SkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

// The tier of the SKU for Azure Dev Spaces Controller.
func (o SkuResponseOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuResponse) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ControllerConnectionDetailsResponseOutput{})
	pulumi.RegisterOutputType(ControllerConnectionDetailsResponseArrayOutput{})
	pulumi.RegisterOutputType(KubernetesConnectionDetailsResponseOutput{})
	pulumi.RegisterOutputType(KubernetesConnectionDetailsResponsePtrOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
}
