// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package powerbi

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

var _ = utilities.GetEnvOrDefault

type AzureSku struct {
	// SKU name
	Name string `pulumi:"name"`
	// SKU tier
	Tier string `pulumi:"tier"`
}

// AzureSkuInput is an input type that accepts AzureSkuArgs and AzureSkuOutput values.
// You can construct a concrete instance of `AzureSkuInput` via:
//
//	AzureSkuArgs{...}
type AzureSkuInput interface {
	pulumi.Input

	ToAzureSkuOutput() AzureSkuOutput
	ToAzureSkuOutputWithContext(context.Context) AzureSkuOutput
}

type AzureSkuArgs struct {
	// SKU name
	Name pulumi.StringInput `pulumi:"name"`
	// SKU tier
	Tier pulumi.StringInput `pulumi:"tier"`
}

func (AzureSkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureSku)(nil)).Elem()
}

func (i AzureSkuArgs) ToAzureSkuOutput() AzureSkuOutput {
	return i.ToAzureSkuOutputWithContext(context.Background())
}

func (i AzureSkuArgs) ToAzureSkuOutputWithContext(ctx context.Context) AzureSkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureSkuOutput)
}

func (i AzureSkuArgs) ToOutput(ctx context.Context) pulumix.Output[AzureSku] {
	return pulumix.Output[AzureSku]{
		OutputState: i.ToAzureSkuOutputWithContext(ctx).OutputState,
	}
}

func (i AzureSkuArgs) ToAzureSkuPtrOutput() AzureSkuPtrOutput {
	return i.ToAzureSkuPtrOutputWithContext(context.Background())
}

func (i AzureSkuArgs) ToAzureSkuPtrOutputWithContext(ctx context.Context) AzureSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureSkuOutput).ToAzureSkuPtrOutputWithContext(ctx)
}

// AzureSkuPtrInput is an input type that accepts AzureSkuArgs, AzureSkuPtr and AzureSkuPtrOutput values.
// You can construct a concrete instance of `AzureSkuPtrInput` via:
//
//	        AzureSkuArgs{...}
//
//	or:
//
//	        nil
type AzureSkuPtrInput interface {
	pulumi.Input

	ToAzureSkuPtrOutput() AzureSkuPtrOutput
	ToAzureSkuPtrOutputWithContext(context.Context) AzureSkuPtrOutput
}

type azureSkuPtrType AzureSkuArgs

func AzureSkuPtr(v *AzureSkuArgs) AzureSkuPtrInput {
	return (*azureSkuPtrType)(v)
}

func (*azureSkuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureSku)(nil)).Elem()
}

func (i *azureSkuPtrType) ToAzureSkuPtrOutput() AzureSkuPtrOutput {
	return i.ToAzureSkuPtrOutputWithContext(context.Background())
}

func (i *azureSkuPtrType) ToAzureSkuPtrOutputWithContext(ctx context.Context) AzureSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureSkuPtrOutput)
}

func (i *azureSkuPtrType) ToOutput(ctx context.Context) pulumix.Output[*AzureSku] {
	return pulumix.Output[*AzureSku]{
		OutputState: i.ToAzureSkuPtrOutputWithContext(ctx).OutputState,
	}
}

type AzureSkuOutput struct{ *pulumi.OutputState }

func (AzureSkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureSku)(nil)).Elem()
}

func (o AzureSkuOutput) ToAzureSkuOutput() AzureSkuOutput {
	return o
}

func (o AzureSkuOutput) ToAzureSkuOutputWithContext(ctx context.Context) AzureSkuOutput {
	return o
}

func (o AzureSkuOutput) ToAzureSkuPtrOutput() AzureSkuPtrOutput {
	return o.ToAzureSkuPtrOutputWithContext(context.Background())
}

func (o AzureSkuOutput) ToAzureSkuPtrOutputWithContext(ctx context.Context) AzureSkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AzureSku) *AzureSku {
		return &v
	}).(AzureSkuPtrOutput)
}

func (o AzureSkuOutput) ToOutput(ctx context.Context) pulumix.Output[AzureSku] {
	return pulumix.Output[AzureSku]{
		OutputState: o.OutputState,
	}
}

// SKU name
func (o AzureSkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v AzureSku) string { return v.Name }).(pulumi.StringOutput)
}

// SKU tier
func (o AzureSkuOutput) Tier() pulumi.StringOutput {
	return o.ApplyT(func(v AzureSku) string { return v.Tier }).(pulumi.StringOutput)
}

type AzureSkuPtrOutput struct{ *pulumi.OutputState }

func (AzureSkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureSku)(nil)).Elem()
}

func (o AzureSkuPtrOutput) ToAzureSkuPtrOutput() AzureSkuPtrOutput {
	return o
}

func (o AzureSkuPtrOutput) ToAzureSkuPtrOutputWithContext(ctx context.Context) AzureSkuPtrOutput {
	return o
}

func (o AzureSkuPtrOutput) ToOutput(ctx context.Context) pulumix.Output[*AzureSku] {
	return pulumix.Output[*AzureSku]{
		OutputState: o.OutputState,
	}
}

func (o AzureSkuPtrOutput) Elem() AzureSkuOutput {
	return o.ApplyT(func(v *AzureSku) AzureSku {
		if v != nil {
			return *v
		}
		var ret AzureSku
		return ret
	}).(AzureSkuOutput)
}

// SKU name
func (o AzureSkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureSku) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

// SKU tier
func (o AzureSkuPtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureSku) *string {
		if v == nil {
			return nil
		}
		return &v.Tier
	}).(pulumi.StringPtrOutput)
}

type AzureSkuResponse struct {
	// SKU name
	Name string `pulumi:"name"`
	// SKU tier
	Tier string `pulumi:"tier"`
}

type AzureSkuResponseOutput struct{ *pulumi.OutputState }

func (AzureSkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureSkuResponse)(nil)).Elem()
}

func (o AzureSkuResponseOutput) ToAzureSkuResponseOutput() AzureSkuResponseOutput {
	return o
}

func (o AzureSkuResponseOutput) ToAzureSkuResponseOutputWithContext(ctx context.Context) AzureSkuResponseOutput {
	return o
}

func (o AzureSkuResponseOutput) ToOutput(ctx context.Context) pulumix.Output[AzureSkuResponse] {
	return pulumix.Output[AzureSkuResponse]{
		OutputState: o.OutputState,
	}
}

// SKU name
func (o AzureSkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v AzureSkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

// SKU tier
func (o AzureSkuResponseOutput) Tier() pulumi.StringOutput {
	return o.ApplyT(func(v AzureSkuResponse) string { return v.Tier }).(pulumi.StringOutput)
}

type AzureSkuResponsePtrOutput struct{ *pulumi.OutputState }

func (AzureSkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureSkuResponse)(nil)).Elem()
}

func (o AzureSkuResponsePtrOutput) ToAzureSkuResponsePtrOutput() AzureSkuResponsePtrOutput {
	return o
}

func (o AzureSkuResponsePtrOutput) ToAzureSkuResponsePtrOutputWithContext(ctx context.Context) AzureSkuResponsePtrOutput {
	return o
}

func (o AzureSkuResponsePtrOutput) ToOutput(ctx context.Context) pulumix.Output[*AzureSkuResponse] {
	return pulumix.Output[*AzureSkuResponse]{
		OutputState: o.OutputState,
	}
}

func (o AzureSkuResponsePtrOutput) Elem() AzureSkuResponseOutput {
	return o.ApplyT(func(v *AzureSkuResponse) AzureSkuResponse {
		if v != nil {
			return *v
		}
		var ret AzureSkuResponse
		return ret
	}).(AzureSkuResponseOutput)
}

// SKU name
func (o AzureSkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureSkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

// SKU tier
func (o AzureSkuResponsePtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureSkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Tier
	}).(pulumi.StringPtrOutput)
}

// ConnectionState information.
type ConnectionState struct {
	// Actions required (if any).
	ActionsRequired *string `pulumi:"actionsRequired"`
	// Description of the connection state.
	Description *string `pulumi:"description"`
	// Status of the connection.
	Status *string `pulumi:"status"`
}

// ConnectionStateInput is an input type that accepts ConnectionStateArgs and ConnectionStateOutput values.
// You can construct a concrete instance of `ConnectionStateInput` via:
//
//	ConnectionStateArgs{...}
type ConnectionStateInput interface {
	pulumi.Input

	ToConnectionStateOutput() ConnectionStateOutput
	ToConnectionStateOutputWithContext(context.Context) ConnectionStateOutput
}

// ConnectionState information.
type ConnectionStateArgs struct {
	// Actions required (if any).
	ActionsRequired pulumi.StringPtrInput `pulumi:"actionsRequired"`
	// Description of the connection state.
	Description pulumi.StringPtrInput `pulumi:"description"`
	// Status of the connection.
	Status pulumi.StringPtrInput `pulumi:"status"`
}

func (ConnectionStateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionState)(nil)).Elem()
}

func (i ConnectionStateArgs) ToConnectionStateOutput() ConnectionStateOutput {
	return i.ToConnectionStateOutputWithContext(context.Background())
}

func (i ConnectionStateArgs) ToConnectionStateOutputWithContext(ctx context.Context) ConnectionStateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionStateOutput)
}

func (i ConnectionStateArgs) ToOutput(ctx context.Context) pulumix.Output[ConnectionState] {
	return pulumix.Output[ConnectionState]{
		OutputState: i.ToConnectionStateOutputWithContext(ctx).OutputState,
	}
}

func (i ConnectionStateArgs) ToConnectionStatePtrOutput() ConnectionStatePtrOutput {
	return i.ToConnectionStatePtrOutputWithContext(context.Background())
}

func (i ConnectionStateArgs) ToConnectionStatePtrOutputWithContext(ctx context.Context) ConnectionStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionStateOutput).ToConnectionStatePtrOutputWithContext(ctx)
}

// ConnectionStatePtrInput is an input type that accepts ConnectionStateArgs, ConnectionStatePtr and ConnectionStatePtrOutput values.
// You can construct a concrete instance of `ConnectionStatePtrInput` via:
//
//	        ConnectionStateArgs{...}
//
//	or:
//
//	        nil
type ConnectionStatePtrInput interface {
	pulumi.Input

	ToConnectionStatePtrOutput() ConnectionStatePtrOutput
	ToConnectionStatePtrOutputWithContext(context.Context) ConnectionStatePtrOutput
}

type connectionStatePtrType ConnectionStateArgs

func ConnectionStatePtr(v *ConnectionStateArgs) ConnectionStatePtrInput {
	return (*connectionStatePtrType)(v)
}

func (*connectionStatePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectionState)(nil)).Elem()
}

func (i *connectionStatePtrType) ToConnectionStatePtrOutput() ConnectionStatePtrOutput {
	return i.ToConnectionStatePtrOutputWithContext(context.Background())
}

func (i *connectionStatePtrType) ToConnectionStatePtrOutputWithContext(ctx context.Context) ConnectionStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionStatePtrOutput)
}

func (i *connectionStatePtrType) ToOutput(ctx context.Context) pulumix.Output[*ConnectionState] {
	return pulumix.Output[*ConnectionState]{
		OutputState: i.ToConnectionStatePtrOutputWithContext(ctx).OutputState,
	}
}

// ConnectionState information.
type ConnectionStateOutput struct{ *pulumi.OutputState }

func (ConnectionStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionState)(nil)).Elem()
}

func (o ConnectionStateOutput) ToConnectionStateOutput() ConnectionStateOutput {
	return o
}

func (o ConnectionStateOutput) ToConnectionStateOutputWithContext(ctx context.Context) ConnectionStateOutput {
	return o
}

func (o ConnectionStateOutput) ToConnectionStatePtrOutput() ConnectionStatePtrOutput {
	return o.ToConnectionStatePtrOutputWithContext(context.Background())
}

func (o ConnectionStateOutput) ToConnectionStatePtrOutputWithContext(ctx context.Context) ConnectionStatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConnectionState) *ConnectionState {
		return &v
	}).(ConnectionStatePtrOutput)
}

func (o ConnectionStateOutput) ToOutput(ctx context.Context) pulumix.Output[ConnectionState] {
	return pulumix.Output[ConnectionState]{
		OutputState: o.OutputState,
	}
}

// Actions required (if any).
func (o ConnectionStateOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionState) *string { return v.ActionsRequired }).(pulumi.StringPtrOutput)
}

// Description of the connection state.
func (o ConnectionStateOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionState) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Status of the connection.
func (o ConnectionStateOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionState) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type ConnectionStatePtrOutput struct{ *pulumi.OutputState }

func (ConnectionStatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectionState)(nil)).Elem()
}

func (o ConnectionStatePtrOutput) ToConnectionStatePtrOutput() ConnectionStatePtrOutput {
	return o
}

func (o ConnectionStatePtrOutput) ToConnectionStatePtrOutputWithContext(ctx context.Context) ConnectionStatePtrOutput {
	return o
}

func (o ConnectionStatePtrOutput) ToOutput(ctx context.Context) pulumix.Output[*ConnectionState] {
	return pulumix.Output[*ConnectionState]{
		OutputState: o.OutputState,
	}
}

func (o ConnectionStatePtrOutput) Elem() ConnectionStateOutput {
	return o.ApplyT(func(v *ConnectionState) ConnectionState {
		if v != nil {
			return *v
		}
		var ret ConnectionState
		return ret
	}).(ConnectionStateOutput)
}

// Actions required (if any).
func (o ConnectionStatePtrOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.ActionsRequired
	}).(pulumi.StringPtrOutput)
}

// Description of the connection state.
func (o ConnectionStatePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

// Status of the connection.
func (o ConnectionStatePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

// ConnectionState information.
type ConnectionStateResponse struct {
	// Actions required (if any).
	ActionsRequired *string `pulumi:"actionsRequired"`
	// Description of the connection state.
	Description *string `pulumi:"description"`
	// Status of the connection.
	Status *string `pulumi:"status"`
}

// ConnectionState information.
type ConnectionStateResponseOutput struct{ *pulumi.OutputState }

func (ConnectionStateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionStateResponse)(nil)).Elem()
}

func (o ConnectionStateResponseOutput) ToConnectionStateResponseOutput() ConnectionStateResponseOutput {
	return o
}

func (o ConnectionStateResponseOutput) ToConnectionStateResponseOutputWithContext(ctx context.Context) ConnectionStateResponseOutput {
	return o
}

func (o ConnectionStateResponseOutput) ToOutput(ctx context.Context) pulumix.Output[ConnectionStateResponse] {
	return pulumix.Output[ConnectionStateResponse]{
		OutputState: o.OutputState,
	}
}

// Actions required (if any).
func (o ConnectionStateResponseOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionStateResponse) *string { return v.ActionsRequired }).(pulumi.StringPtrOutput)
}

// Description of the connection state.
func (o ConnectionStateResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionStateResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Status of the connection.
func (o ConnectionStateResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionStateResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type ConnectionStateResponsePtrOutput struct{ *pulumi.OutputState }

func (ConnectionStateResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectionStateResponse)(nil)).Elem()
}

func (o ConnectionStateResponsePtrOutput) ToConnectionStateResponsePtrOutput() ConnectionStateResponsePtrOutput {
	return o
}

func (o ConnectionStateResponsePtrOutput) ToConnectionStateResponsePtrOutputWithContext(ctx context.Context) ConnectionStateResponsePtrOutput {
	return o
}

func (o ConnectionStateResponsePtrOutput) ToOutput(ctx context.Context) pulumix.Output[*ConnectionStateResponse] {
	return pulumix.Output[*ConnectionStateResponse]{
		OutputState: o.OutputState,
	}
}

func (o ConnectionStateResponsePtrOutput) Elem() ConnectionStateResponseOutput {
	return o.ApplyT(func(v *ConnectionStateResponse) ConnectionStateResponse {
		if v != nil {
			return *v
		}
		var ret ConnectionStateResponse
		return ret
	}).(ConnectionStateResponseOutput)
}

// Actions required (if any).
func (o ConnectionStateResponsePtrOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.ActionsRequired
	}).(pulumi.StringPtrOutput)
}

// Description of the connection state.
func (o ConnectionStateResponsePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

// Status of the connection.
func (o ConnectionStateResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type PrivateEndpoint struct {
	// Specifies the id of private endpoint.
	Id *string `pulumi:"id"`
}

// PrivateEndpointInput is an input type that accepts PrivateEndpointArgs and PrivateEndpointOutput values.
// You can construct a concrete instance of `PrivateEndpointInput` via:
//
//	PrivateEndpointArgs{...}
type PrivateEndpointInput interface {
	pulumi.Input

	ToPrivateEndpointOutput() PrivateEndpointOutput
	ToPrivateEndpointOutputWithContext(context.Context) PrivateEndpointOutput
}

type PrivateEndpointArgs struct {
	// Specifies the id of private endpoint.
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (PrivateEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpoint)(nil)).Elem()
}

func (i PrivateEndpointArgs) ToPrivateEndpointOutput() PrivateEndpointOutput {
	return i.ToPrivateEndpointOutputWithContext(context.Background())
}

func (i PrivateEndpointArgs) ToPrivateEndpointOutputWithContext(ctx context.Context) PrivateEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointOutput)
}

func (i PrivateEndpointArgs) ToOutput(ctx context.Context) pulumix.Output[PrivateEndpoint] {
	return pulumix.Output[PrivateEndpoint]{
		OutputState: i.ToPrivateEndpointOutputWithContext(ctx).OutputState,
	}
}

func (i PrivateEndpointArgs) ToPrivateEndpointPtrOutput() PrivateEndpointPtrOutput {
	return i.ToPrivateEndpointPtrOutputWithContext(context.Background())
}

func (i PrivateEndpointArgs) ToPrivateEndpointPtrOutputWithContext(ctx context.Context) PrivateEndpointPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointOutput).ToPrivateEndpointPtrOutputWithContext(ctx)
}

// PrivateEndpointPtrInput is an input type that accepts PrivateEndpointArgs, PrivateEndpointPtr and PrivateEndpointPtrOutput values.
// You can construct a concrete instance of `PrivateEndpointPtrInput` via:
//
//	        PrivateEndpointArgs{...}
//
//	or:
//
//	        nil
type PrivateEndpointPtrInput interface {
	pulumi.Input

	ToPrivateEndpointPtrOutput() PrivateEndpointPtrOutput
	ToPrivateEndpointPtrOutputWithContext(context.Context) PrivateEndpointPtrOutput
}

type privateEndpointPtrType PrivateEndpointArgs

func PrivateEndpointPtr(v *PrivateEndpointArgs) PrivateEndpointPtrInput {
	return (*privateEndpointPtrType)(v)
}

func (*privateEndpointPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpoint)(nil)).Elem()
}

func (i *privateEndpointPtrType) ToPrivateEndpointPtrOutput() PrivateEndpointPtrOutput {
	return i.ToPrivateEndpointPtrOutputWithContext(context.Background())
}

func (i *privateEndpointPtrType) ToPrivateEndpointPtrOutputWithContext(ctx context.Context) PrivateEndpointPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointPtrOutput)
}

func (i *privateEndpointPtrType) ToOutput(ctx context.Context) pulumix.Output[*PrivateEndpoint] {
	return pulumix.Output[*PrivateEndpoint]{
		OutputState: i.ToPrivateEndpointPtrOutputWithContext(ctx).OutputState,
	}
}

type PrivateEndpointOutput struct{ *pulumi.OutputState }

func (PrivateEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpoint)(nil)).Elem()
}

func (o PrivateEndpointOutput) ToPrivateEndpointOutput() PrivateEndpointOutput {
	return o
}

func (o PrivateEndpointOutput) ToPrivateEndpointOutputWithContext(ctx context.Context) PrivateEndpointOutput {
	return o
}

func (o PrivateEndpointOutput) ToPrivateEndpointPtrOutput() PrivateEndpointPtrOutput {
	return o.ToPrivateEndpointPtrOutputWithContext(context.Background())
}

func (o PrivateEndpointOutput) ToPrivateEndpointPtrOutputWithContext(ctx context.Context) PrivateEndpointPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateEndpoint) *PrivateEndpoint {
		return &v
	}).(PrivateEndpointPtrOutput)
}

func (o PrivateEndpointOutput) ToOutput(ctx context.Context) pulumix.Output[PrivateEndpoint] {
	return pulumix.Output[PrivateEndpoint]{
		OutputState: o.OutputState,
	}
}

// Specifies the id of private endpoint.
func (o PrivateEndpointOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateEndpoint) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type PrivateEndpointPtrOutput struct{ *pulumi.OutputState }

func (PrivateEndpointPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpoint)(nil)).Elem()
}

func (o PrivateEndpointPtrOutput) ToPrivateEndpointPtrOutput() PrivateEndpointPtrOutput {
	return o
}

func (o PrivateEndpointPtrOutput) ToPrivateEndpointPtrOutputWithContext(ctx context.Context) PrivateEndpointPtrOutput {
	return o
}

func (o PrivateEndpointPtrOutput) ToOutput(ctx context.Context) pulumix.Output[*PrivateEndpoint] {
	return pulumix.Output[*PrivateEndpoint]{
		OutputState: o.OutputState,
	}
}

func (o PrivateEndpointPtrOutput) Elem() PrivateEndpointOutput {
	return o.ApplyT(func(v *PrivateEndpoint) PrivateEndpoint {
		if v != nil {
			return *v
		}
		var ret PrivateEndpoint
		return ret
	}).(PrivateEndpointOutput)
}

// Specifies the id of private endpoint.
func (o PrivateEndpointPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateEndpoint) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type PrivateEndpointConnectionType struct {
	// Specifies the private endpoint.
	PrivateEndpoint *PrivateEndpoint `pulumi:"privateEndpoint"`
	// Specifies the connection state.
	PrivateLinkServiceConnectionState *ConnectionState `pulumi:"privateLinkServiceConnectionState"`
	// Provisioning state of the Private Endpoint Connection.
	ProvisioningState *string `pulumi:"provisioningState"`
}

// PrivateEndpointConnectionTypeInput is an input type that accepts PrivateEndpointConnectionTypeArgs and PrivateEndpointConnectionTypeOutput values.
// You can construct a concrete instance of `PrivateEndpointConnectionTypeInput` via:
//
//	PrivateEndpointConnectionTypeArgs{...}
type PrivateEndpointConnectionTypeInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionTypeOutput() PrivateEndpointConnectionTypeOutput
	ToPrivateEndpointConnectionTypeOutputWithContext(context.Context) PrivateEndpointConnectionTypeOutput
}

type PrivateEndpointConnectionTypeArgs struct {
	// Specifies the private endpoint.
	PrivateEndpoint PrivateEndpointPtrInput `pulumi:"privateEndpoint"`
	// Specifies the connection state.
	PrivateLinkServiceConnectionState ConnectionStatePtrInput `pulumi:"privateLinkServiceConnectionState"`
	// Provisioning state of the Private Endpoint Connection.
	ProvisioningState pulumi.StringPtrInput `pulumi:"provisioningState"`
}

func (PrivateEndpointConnectionTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionType)(nil)).Elem()
}

func (i PrivateEndpointConnectionTypeArgs) ToPrivateEndpointConnectionTypeOutput() PrivateEndpointConnectionTypeOutput {
	return i.ToPrivateEndpointConnectionTypeOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionTypeArgs) ToPrivateEndpointConnectionTypeOutputWithContext(ctx context.Context) PrivateEndpointConnectionTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionTypeOutput)
}

func (i PrivateEndpointConnectionTypeArgs) ToOutput(ctx context.Context) pulumix.Output[PrivateEndpointConnectionType] {
	return pulumix.Output[PrivateEndpointConnectionType]{
		OutputState: i.ToPrivateEndpointConnectionTypeOutputWithContext(ctx).OutputState,
	}
}

// PrivateEndpointConnectionTypeArrayInput is an input type that accepts PrivateEndpointConnectionTypeArray and PrivateEndpointConnectionTypeArrayOutput values.
// You can construct a concrete instance of `PrivateEndpointConnectionTypeArrayInput` via:
//
//	PrivateEndpointConnectionTypeArray{ PrivateEndpointConnectionTypeArgs{...} }
type PrivateEndpointConnectionTypeArrayInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionTypeArrayOutput() PrivateEndpointConnectionTypeArrayOutput
	ToPrivateEndpointConnectionTypeArrayOutputWithContext(context.Context) PrivateEndpointConnectionTypeArrayOutput
}

type PrivateEndpointConnectionTypeArray []PrivateEndpointConnectionTypeInput

func (PrivateEndpointConnectionTypeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateEndpointConnectionType)(nil)).Elem()
}

func (i PrivateEndpointConnectionTypeArray) ToPrivateEndpointConnectionTypeArrayOutput() PrivateEndpointConnectionTypeArrayOutput {
	return i.ToPrivateEndpointConnectionTypeArrayOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionTypeArray) ToPrivateEndpointConnectionTypeArrayOutputWithContext(ctx context.Context) PrivateEndpointConnectionTypeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionTypeArrayOutput)
}

func (i PrivateEndpointConnectionTypeArray) ToOutput(ctx context.Context) pulumix.Output[[]PrivateEndpointConnectionType] {
	return pulumix.Output[[]PrivateEndpointConnectionType]{
		OutputState: i.ToPrivateEndpointConnectionTypeArrayOutputWithContext(ctx).OutputState,
	}
}

type PrivateEndpointConnectionTypeOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionType)(nil)).Elem()
}

func (o PrivateEndpointConnectionTypeOutput) ToPrivateEndpointConnectionTypeOutput() PrivateEndpointConnectionTypeOutput {
	return o
}

func (o PrivateEndpointConnectionTypeOutput) ToPrivateEndpointConnectionTypeOutputWithContext(ctx context.Context) PrivateEndpointConnectionTypeOutput {
	return o
}

func (o PrivateEndpointConnectionTypeOutput) ToOutput(ctx context.Context) pulumix.Output[PrivateEndpointConnectionType] {
	return pulumix.Output[PrivateEndpointConnectionType]{
		OutputState: o.OutputState,
	}
}

// Specifies the private endpoint.
func (o PrivateEndpointConnectionTypeOutput) PrivateEndpoint() PrivateEndpointPtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionType) *PrivateEndpoint { return v.PrivateEndpoint }).(PrivateEndpointPtrOutput)
}

// Specifies the connection state.
func (o PrivateEndpointConnectionTypeOutput) PrivateLinkServiceConnectionState() ConnectionStatePtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionType) *ConnectionState { return v.PrivateLinkServiceConnectionState }).(ConnectionStatePtrOutput)
}

// Provisioning state of the Private Endpoint Connection.
func (o PrivateEndpointConnectionTypeOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionType) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

type PrivateEndpointConnectionTypeArrayOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionTypeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateEndpointConnectionType)(nil)).Elem()
}

func (o PrivateEndpointConnectionTypeArrayOutput) ToPrivateEndpointConnectionTypeArrayOutput() PrivateEndpointConnectionTypeArrayOutput {
	return o
}

func (o PrivateEndpointConnectionTypeArrayOutput) ToPrivateEndpointConnectionTypeArrayOutputWithContext(ctx context.Context) PrivateEndpointConnectionTypeArrayOutput {
	return o
}

func (o PrivateEndpointConnectionTypeArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]PrivateEndpointConnectionType] {
	return pulumix.Output[[]PrivateEndpointConnectionType]{
		OutputState: o.OutputState,
	}
}

func (o PrivateEndpointConnectionTypeArrayOutput) Index(i pulumi.IntInput) PrivateEndpointConnectionTypeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PrivateEndpointConnectionType {
		return vs[0].([]PrivateEndpointConnectionType)[vs[1].(int)]
	}).(PrivateEndpointConnectionTypeOutput)
}

type PrivateEndpointConnectionResponse struct {
	// Specifies the id of the resource.
	Id string `pulumi:"id"`
	// Specifies the name of the resource.
	Name string `pulumi:"name"`
	// Specifies the private endpoint.
	PrivateEndpoint *PrivateEndpointResponse `pulumi:"privateEndpoint"`
	// Specifies the connection state.
	PrivateLinkServiceConnectionState *ConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	// Provisioning state of the Private Endpoint Connection.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The system meta data relating to this resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Specifies the type of the resource.
	Type string `pulumi:"type"`
}

type PrivateEndpointConnectionResponseOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionResponse)(nil)).Elem()
}

func (o PrivateEndpointConnectionResponseOutput) ToPrivateEndpointConnectionResponseOutput() PrivateEndpointConnectionResponseOutput {
	return o
}

func (o PrivateEndpointConnectionResponseOutput) ToPrivateEndpointConnectionResponseOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponseOutput {
	return o
}

func (o PrivateEndpointConnectionResponseOutput) ToOutput(ctx context.Context) pulumix.Output[PrivateEndpointConnectionResponse] {
	return pulumix.Output[PrivateEndpointConnectionResponse]{
		OutputState: o.OutputState,
	}
}

// Specifies the id of the resource.
func (o PrivateEndpointConnectionResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Id }).(pulumi.StringOutput)
}

// Specifies the name of the resource.
func (o PrivateEndpointConnectionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Name }).(pulumi.StringOutput)
}

// Specifies the private endpoint.
func (o PrivateEndpointConnectionResponseOutput) PrivateEndpoint() PrivateEndpointResponsePtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) *PrivateEndpointResponse { return v.PrivateEndpoint }).(PrivateEndpointResponsePtrOutput)
}

// Specifies the connection state.
func (o PrivateEndpointConnectionResponseOutput) PrivateLinkServiceConnectionState() ConnectionStateResponsePtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) *ConnectionStateResponse {
		return v.PrivateLinkServiceConnectionState
	}).(ConnectionStateResponsePtrOutput)
}

// Provisioning state of the Private Endpoint Connection.
func (o PrivateEndpointConnectionResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

// The system meta data relating to this resource.
func (o PrivateEndpointConnectionResponseOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Specifies the type of the resource.
func (o PrivateEndpointConnectionResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Type }).(pulumi.StringOutput)
}

type PrivateEndpointConnectionResponseArrayOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateEndpointConnectionResponse)(nil)).Elem()
}

func (o PrivateEndpointConnectionResponseArrayOutput) ToPrivateEndpointConnectionResponseArrayOutput() PrivateEndpointConnectionResponseArrayOutput {
	return o
}

func (o PrivateEndpointConnectionResponseArrayOutput) ToPrivateEndpointConnectionResponseArrayOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponseArrayOutput {
	return o
}

func (o PrivateEndpointConnectionResponseArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]PrivateEndpointConnectionResponse] {
	return pulumix.Output[[]PrivateEndpointConnectionResponse]{
		OutputState: o.OutputState,
	}
}

func (o PrivateEndpointConnectionResponseArrayOutput) Index(i pulumi.IntInput) PrivateEndpointConnectionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PrivateEndpointConnectionResponse {
		return vs[0].([]PrivateEndpointConnectionResponse)[vs[1].(int)]
	}).(PrivateEndpointConnectionResponseOutput)
}

type PrivateEndpointResponse struct {
	// Specifies the id of private endpoint.
	Id *string `pulumi:"id"`
}

type PrivateEndpointResponseOutput struct{ *pulumi.OutputState }

func (PrivateEndpointResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointResponse)(nil)).Elem()
}

func (o PrivateEndpointResponseOutput) ToPrivateEndpointResponseOutput() PrivateEndpointResponseOutput {
	return o
}

func (o PrivateEndpointResponseOutput) ToPrivateEndpointResponseOutputWithContext(ctx context.Context) PrivateEndpointResponseOutput {
	return o
}

func (o PrivateEndpointResponseOutput) ToOutput(ctx context.Context) pulumix.Output[PrivateEndpointResponse] {
	return pulumix.Output[PrivateEndpointResponse]{
		OutputState: o.OutputState,
	}
}

// Specifies the id of private endpoint.
func (o PrivateEndpointResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateEndpointResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type PrivateEndpointResponsePtrOutput struct{ *pulumi.OutputState }

func (PrivateEndpointResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointResponse)(nil)).Elem()
}

func (o PrivateEndpointResponsePtrOutput) ToPrivateEndpointResponsePtrOutput() PrivateEndpointResponsePtrOutput {
	return o
}

func (o PrivateEndpointResponsePtrOutput) ToPrivateEndpointResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointResponsePtrOutput {
	return o
}

func (o PrivateEndpointResponsePtrOutput) ToOutput(ctx context.Context) pulumix.Output[*PrivateEndpointResponse] {
	return pulumix.Output[*PrivateEndpointResponse]{
		OutputState: o.OutputState,
	}
}

func (o PrivateEndpointResponsePtrOutput) Elem() PrivateEndpointResponseOutput {
	return o.ApplyT(func(v *PrivateEndpointResponse) PrivateEndpointResponse {
		if v != nil {
			return *v
		}
		var ret PrivateEndpointResponse
		return ret
	}).(PrivateEndpointResponseOutput)
}

// Specifies the id of private endpoint.
func (o PrivateEndpointResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateEndpointResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
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
	pulumi.RegisterOutputType(AzureSkuOutput{})
	pulumi.RegisterOutputType(AzureSkuPtrOutput{})
	pulumi.RegisterOutputType(AzureSkuResponseOutput{})
	pulumi.RegisterOutputType(AzureSkuResponsePtrOutput{})
	pulumi.RegisterOutputType(ConnectionStateOutput{})
	pulumi.RegisterOutputType(ConnectionStatePtrOutput{})
	pulumi.RegisterOutputType(ConnectionStateResponseOutput{})
	pulumi.RegisterOutputType(ConnectionStateResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointOutput{})
	pulumi.RegisterOutputType(PrivateEndpointPtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionTypeOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionTypeArrayOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseArrayOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponsePtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
}
