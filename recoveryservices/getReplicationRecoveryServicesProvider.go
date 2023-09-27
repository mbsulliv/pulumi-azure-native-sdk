// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package recoveryservices

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets the details of registered recovery services provider.
// Azure REST API version: 2023-04-01.
func LookupReplicationRecoveryServicesProvider(ctx *pulumi.Context, args *LookupReplicationRecoveryServicesProviderArgs, opts ...pulumi.InvokeOption) (*LookupReplicationRecoveryServicesProviderResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupReplicationRecoveryServicesProviderResult
	err := ctx.Invoke("azure-native:recoveryservices:getReplicationRecoveryServicesProvider", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupReplicationRecoveryServicesProviderArgs struct {
	// Fabric name.
	FabricName string `pulumi:"fabricName"`
	// Recovery services provider name.
	ProviderName string `pulumi:"providerName"`
	// The name of the resource group where the recovery services vault is present.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the recovery services vault.
	ResourceName string `pulumi:"resourceName"`
}

// Provider details.
type LookupReplicationRecoveryServicesProviderResult struct {
	// Resource Id
	Id string `pulumi:"id"`
	// Resource Location
	Location *string `pulumi:"location"`
	// Resource Name
	Name string `pulumi:"name"`
	// Provider properties.
	Properties RecoveryServicesProviderPropertiesResponse `pulumi:"properties"`
	// Resource Type
	Type string `pulumi:"type"`
}

func LookupReplicationRecoveryServicesProviderOutput(ctx *pulumi.Context, args LookupReplicationRecoveryServicesProviderOutputArgs, opts ...pulumi.InvokeOption) LookupReplicationRecoveryServicesProviderResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupReplicationRecoveryServicesProviderResult, error) {
			args := v.(LookupReplicationRecoveryServicesProviderArgs)
			r, err := LookupReplicationRecoveryServicesProvider(ctx, &args, opts...)
			var s LookupReplicationRecoveryServicesProviderResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupReplicationRecoveryServicesProviderResultOutput)
}

type LookupReplicationRecoveryServicesProviderOutputArgs struct {
	// Fabric name.
	FabricName pulumi.StringInput `pulumi:"fabricName"`
	// Recovery services provider name.
	ProviderName pulumi.StringInput `pulumi:"providerName"`
	// The name of the resource group where the recovery services vault is present.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the recovery services vault.
	ResourceName pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupReplicationRecoveryServicesProviderOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReplicationRecoveryServicesProviderArgs)(nil)).Elem()
}

// Provider details.
type LookupReplicationRecoveryServicesProviderResultOutput struct{ *pulumi.OutputState }

func (LookupReplicationRecoveryServicesProviderResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReplicationRecoveryServicesProviderResult)(nil)).Elem()
}

func (o LookupReplicationRecoveryServicesProviderResultOutput) ToLookupReplicationRecoveryServicesProviderResultOutput() LookupReplicationRecoveryServicesProviderResultOutput {
	return o
}

func (o LookupReplicationRecoveryServicesProviderResultOutput) ToLookupReplicationRecoveryServicesProviderResultOutputWithContext(ctx context.Context) LookupReplicationRecoveryServicesProviderResultOutput {
	return o
}

func (o LookupReplicationRecoveryServicesProviderResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupReplicationRecoveryServicesProviderResult] {
	return pulumix.Output[LookupReplicationRecoveryServicesProviderResult]{
		OutputState: o.OutputState,
	}
}

// Resource Id
func (o LookupReplicationRecoveryServicesProviderResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReplicationRecoveryServicesProviderResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource Location
func (o LookupReplicationRecoveryServicesProviderResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupReplicationRecoveryServicesProviderResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource Name
func (o LookupReplicationRecoveryServicesProviderResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReplicationRecoveryServicesProviderResult) string { return v.Name }).(pulumi.StringOutput)
}

// Provider properties.
func (o LookupReplicationRecoveryServicesProviderResultOutput) Properties() RecoveryServicesProviderPropertiesResponseOutput {
	return o.ApplyT(func(v LookupReplicationRecoveryServicesProviderResult) RecoveryServicesProviderPropertiesResponse {
		return v.Properties
	}).(RecoveryServicesProviderPropertiesResponseOutput)
}

// Resource Type
func (o LookupReplicationRecoveryServicesProviderResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReplicationRecoveryServicesProviderResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupReplicationRecoveryServicesProviderResultOutput{})
}
