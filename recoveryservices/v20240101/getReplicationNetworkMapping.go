// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20240101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the details of an ASR network mapping.
func LookupReplicationNetworkMapping(ctx *pulumi.Context, args *LookupReplicationNetworkMappingArgs, opts ...pulumi.InvokeOption) (*LookupReplicationNetworkMappingResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupReplicationNetworkMappingResult
	err := ctx.Invoke("azure-native:recoveryservices/v20240101:getReplicationNetworkMapping", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupReplicationNetworkMappingArgs struct {
	// Primary fabric name.
	FabricName string `pulumi:"fabricName"`
	// Network mapping name.
	NetworkMappingName string `pulumi:"networkMappingName"`
	// Primary network name.
	NetworkName string `pulumi:"networkName"`
	// The name of the resource group where the recovery services vault is present.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the recovery services vault.
	ResourceName string `pulumi:"resourceName"`
}

// Network Mapping model. Ideally it should have been possible to inherit this class from prev version in InheritedModels as long as there is no difference in structure or method signature. Since there were no base Models for certain fields and methods viz NetworkMappingProperties and Load with required return type, the class has been introduced in its entirety with references to base models to facilitate extensions in subsequent versions.
type LookupReplicationNetworkMappingResult struct {
	// Resource Id
	Id string `pulumi:"id"`
	// Resource Location
	Location *string `pulumi:"location"`
	// Resource Name
	Name string `pulumi:"name"`
	// The Network Mapping Properties.
	Properties NetworkMappingPropertiesResponse `pulumi:"properties"`
	// Resource Type
	Type string `pulumi:"type"`
}

func LookupReplicationNetworkMappingOutput(ctx *pulumi.Context, args LookupReplicationNetworkMappingOutputArgs, opts ...pulumi.InvokeOption) LookupReplicationNetworkMappingResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupReplicationNetworkMappingResult, error) {
			args := v.(LookupReplicationNetworkMappingArgs)
			r, err := LookupReplicationNetworkMapping(ctx, &args, opts...)
			var s LookupReplicationNetworkMappingResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupReplicationNetworkMappingResultOutput)
}

type LookupReplicationNetworkMappingOutputArgs struct {
	// Primary fabric name.
	FabricName pulumi.StringInput `pulumi:"fabricName"`
	// Network mapping name.
	NetworkMappingName pulumi.StringInput `pulumi:"networkMappingName"`
	// Primary network name.
	NetworkName pulumi.StringInput `pulumi:"networkName"`
	// The name of the resource group where the recovery services vault is present.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the recovery services vault.
	ResourceName pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupReplicationNetworkMappingOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReplicationNetworkMappingArgs)(nil)).Elem()
}

// Network Mapping model. Ideally it should have been possible to inherit this class from prev version in InheritedModels as long as there is no difference in structure or method signature. Since there were no base Models for certain fields and methods viz NetworkMappingProperties and Load with required return type, the class has been introduced in its entirety with references to base models to facilitate extensions in subsequent versions.
type LookupReplicationNetworkMappingResultOutput struct{ *pulumi.OutputState }

func (LookupReplicationNetworkMappingResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReplicationNetworkMappingResult)(nil)).Elem()
}

func (o LookupReplicationNetworkMappingResultOutput) ToLookupReplicationNetworkMappingResultOutput() LookupReplicationNetworkMappingResultOutput {
	return o
}

func (o LookupReplicationNetworkMappingResultOutput) ToLookupReplicationNetworkMappingResultOutputWithContext(ctx context.Context) LookupReplicationNetworkMappingResultOutput {
	return o
}

// Resource Id
func (o LookupReplicationNetworkMappingResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReplicationNetworkMappingResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource Location
func (o LookupReplicationNetworkMappingResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupReplicationNetworkMappingResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource Name
func (o LookupReplicationNetworkMappingResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReplicationNetworkMappingResult) string { return v.Name }).(pulumi.StringOutput)
}

// The Network Mapping Properties.
func (o LookupReplicationNetworkMappingResultOutput) Properties() NetworkMappingPropertiesResponseOutput {
	return o.ApplyT(func(v LookupReplicationNetworkMappingResult) NetworkMappingPropertiesResponse { return v.Properties }).(NetworkMappingPropertiesResponseOutput)
}

// Resource Type
func (o LookupReplicationNetworkMappingResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReplicationNetworkMappingResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupReplicationNetworkMappingResultOutput{})
}
