// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220831preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Sku family of the Cloud HSM Cluster
type CloudHsmClusterSkuFamily string

const (
	CloudHsmClusterSkuFamilyB = CloudHsmClusterSkuFamily("B")
)

// Sku name of the Cloud HSM Cluster
type CloudHsmClusterSkuName string

const (
	CloudHsmClusterSkuName_Standard_B1  = CloudHsmClusterSkuName("Standard_B1")
	CloudHsmClusterSkuName_Standard_B10 = CloudHsmClusterSkuName("Standard B10")
)

func (CloudHsmClusterSkuName) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudHsmClusterSkuName)(nil)).Elem()
}

func (e CloudHsmClusterSkuName) ToCloudHsmClusterSkuNameOutput() CloudHsmClusterSkuNameOutput {
	return pulumi.ToOutput(e).(CloudHsmClusterSkuNameOutput)
}

func (e CloudHsmClusterSkuName) ToCloudHsmClusterSkuNameOutputWithContext(ctx context.Context) CloudHsmClusterSkuNameOutput {
	return pulumi.ToOutputWithContext(ctx, e).(CloudHsmClusterSkuNameOutput)
}

func (e CloudHsmClusterSkuName) ToCloudHsmClusterSkuNamePtrOutput() CloudHsmClusterSkuNamePtrOutput {
	return e.ToCloudHsmClusterSkuNamePtrOutputWithContext(context.Background())
}

func (e CloudHsmClusterSkuName) ToCloudHsmClusterSkuNamePtrOutputWithContext(ctx context.Context) CloudHsmClusterSkuNamePtrOutput {
	return CloudHsmClusterSkuName(e).ToCloudHsmClusterSkuNameOutputWithContext(ctx).ToCloudHsmClusterSkuNamePtrOutputWithContext(ctx)
}

func (e CloudHsmClusterSkuName) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e CloudHsmClusterSkuName) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e CloudHsmClusterSkuName) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e CloudHsmClusterSkuName) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type CloudHsmClusterSkuNameOutput struct{ *pulumi.OutputState }

func (CloudHsmClusterSkuNameOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudHsmClusterSkuName)(nil)).Elem()
}

func (o CloudHsmClusterSkuNameOutput) ToCloudHsmClusterSkuNameOutput() CloudHsmClusterSkuNameOutput {
	return o
}

func (o CloudHsmClusterSkuNameOutput) ToCloudHsmClusterSkuNameOutputWithContext(ctx context.Context) CloudHsmClusterSkuNameOutput {
	return o
}

func (o CloudHsmClusterSkuNameOutput) ToCloudHsmClusterSkuNamePtrOutput() CloudHsmClusterSkuNamePtrOutput {
	return o.ToCloudHsmClusterSkuNamePtrOutputWithContext(context.Background())
}

func (o CloudHsmClusterSkuNameOutput) ToCloudHsmClusterSkuNamePtrOutputWithContext(ctx context.Context) CloudHsmClusterSkuNamePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CloudHsmClusterSkuName) *CloudHsmClusterSkuName {
		return &v
	}).(CloudHsmClusterSkuNamePtrOutput)
}

func (o CloudHsmClusterSkuNameOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o CloudHsmClusterSkuNameOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CloudHsmClusterSkuName) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o CloudHsmClusterSkuNameOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CloudHsmClusterSkuNameOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CloudHsmClusterSkuName) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type CloudHsmClusterSkuNamePtrOutput struct{ *pulumi.OutputState }

func (CloudHsmClusterSkuNamePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudHsmClusterSkuName)(nil)).Elem()
}

func (o CloudHsmClusterSkuNamePtrOutput) ToCloudHsmClusterSkuNamePtrOutput() CloudHsmClusterSkuNamePtrOutput {
	return o
}

func (o CloudHsmClusterSkuNamePtrOutput) ToCloudHsmClusterSkuNamePtrOutputWithContext(ctx context.Context) CloudHsmClusterSkuNamePtrOutput {
	return o
}

func (o CloudHsmClusterSkuNamePtrOutput) Elem() CloudHsmClusterSkuNameOutput {
	return o.ApplyT(func(v *CloudHsmClusterSkuName) CloudHsmClusterSkuName {
		if v != nil {
			return *v
		}
		var ret CloudHsmClusterSkuName
		return ret
	}).(CloudHsmClusterSkuNameOutput)
}

func (o CloudHsmClusterSkuNamePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CloudHsmClusterSkuNamePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *CloudHsmClusterSkuName) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// CloudHsmClusterSkuNameInput is an input type that accepts CloudHsmClusterSkuNameArgs and CloudHsmClusterSkuNameOutput values.
// You can construct a concrete instance of `CloudHsmClusterSkuNameInput` via:
//
//	CloudHsmClusterSkuNameArgs{...}
type CloudHsmClusterSkuNameInput interface {
	pulumi.Input

	ToCloudHsmClusterSkuNameOutput() CloudHsmClusterSkuNameOutput
	ToCloudHsmClusterSkuNameOutputWithContext(context.Context) CloudHsmClusterSkuNameOutput
}

var cloudHsmClusterSkuNamePtrType = reflect.TypeOf((**CloudHsmClusterSkuName)(nil)).Elem()

type CloudHsmClusterSkuNamePtrInput interface {
	pulumi.Input

	ToCloudHsmClusterSkuNamePtrOutput() CloudHsmClusterSkuNamePtrOutput
	ToCloudHsmClusterSkuNamePtrOutputWithContext(context.Context) CloudHsmClusterSkuNamePtrOutput
}

type cloudHsmClusterSkuNamePtr string

func CloudHsmClusterSkuNamePtr(v string) CloudHsmClusterSkuNamePtrInput {
	return (*cloudHsmClusterSkuNamePtr)(&v)
}

func (*cloudHsmClusterSkuNamePtr) ElementType() reflect.Type {
	return cloudHsmClusterSkuNamePtrType
}

func (in *cloudHsmClusterSkuNamePtr) ToCloudHsmClusterSkuNamePtrOutput() CloudHsmClusterSkuNamePtrOutput {
	return pulumi.ToOutput(in).(CloudHsmClusterSkuNamePtrOutput)
}

func (in *cloudHsmClusterSkuNamePtr) ToCloudHsmClusterSkuNamePtrOutputWithContext(ctx context.Context) CloudHsmClusterSkuNamePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(CloudHsmClusterSkuNamePtrOutput)
}

// Indicates whether the connection has been Approved/Rejected/Removed by the owner of the service.
type PrivateEndpointServiceConnectionStatus string

const (
	PrivateEndpointServiceConnectionStatusPending  = PrivateEndpointServiceConnectionStatus("Pending")
	PrivateEndpointServiceConnectionStatusApproved = PrivateEndpointServiceConnectionStatus("Approved")
	PrivateEndpointServiceConnectionStatusRejected = PrivateEndpointServiceConnectionStatus("Rejected")
)

// The Cloud HSM Cluster's provisioningState
type ProvisioningState string

const (
	ProvisioningStateProvisioning = ProvisioningState("Provisioning")
	ProvisioningStateSucceeded    = ProvisioningState("Succeeded")
	ProvisioningStateFailed       = ProvisioningState("Failed")
	ProvisioningStateDeleting     = ProvisioningState("Deleting")
	ProvisioningStateCanceled     = ProvisioningState("Canceled")
)

func init() {
	pulumi.RegisterOutputType(CloudHsmClusterSkuNameOutput{})
	pulumi.RegisterOutputType(CloudHsmClusterSkuNamePtrOutput{})
}