// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package recoveryservices

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the details of a protection container mapping.
// Azure REST API version: 2023-04-01.
//
// Other available API versions: 2023-06-01.
func LookupReplicationProtectionContainerMapping(ctx *pulumi.Context, args *LookupReplicationProtectionContainerMappingArgs, opts ...pulumi.InvokeOption) (*LookupReplicationProtectionContainerMappingResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupReplicationProtectionContainerMappingResult
	err := ctx.Invoke("azure-native:recoveryservices:getReplicationProtectionContainerMapping", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupReplicationProtectionContainerMappingArgs struct {
	// Fabric name.
	FabricName string `pulumi:"fabricName"`
	// Protection Container mapping name.
	MappingName string `pulumi:"mappingName"`
	// Protection container name.
	ProtectionContainerName string `pulumi:"protectionContainerName"`
	// The name of the resource group where the recovery services vault is present.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the recovery services vault.
	ResourceName string `pulumi:"resourceName"`
}

// Protection container mapping object.
type LookupReplicationProtectionContainerMappingResult struct {
	// Resource Id
	Id string `pulumi:"id"`
	// Resource Location
	Location *string `pulumi:"location"`
	// Resource Name
	Name string `pulumi:"name"`
	// The custom data.
	Properties ProtectionContainerMappingPropertiesResponse `pulumi:"properties"`
	// Resource Type
	Type string `pulumi:"type"`
}

func LookupReplicationProtectionContainerMappingOutput(ctx *pulumi.Context, args LookupReplicationProtectionContainerMappingOutputArgs, opts ...pulumi.InvokeOption) LookupReplicationProtectionContainerMappingResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupReplicationProtectionContainerMappingResult, error) {
			args := v.(LookupReplicationProtectionContainerMappingArgs)
			r, err := LookupReplicationProtectionContainerMapping(ctx, &args, opts...)
			var s LookupReplicationProtectionContainerMappingResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupReplicationProtectionContainerMappingResultOutput)
}

type LookupReplicationProtectionContainerMappingOutputArgs struct {
	// Fabric name.
	FabricName pulumi.StringInput `pulumi:"fabricName"`
	// Protection Container mapping name.
	MappingName pulumi.StringInput `pulumi:"mappingName"`
	// Protection container name.
	ProtectionContainerName pulumi.StringInput `pulumi:"protectionContainerName"`
	// The name of the resource group where the recovery services vault is present.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the recovery services vault.
	ResourceName pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupReplicationProtectionContainerMappingOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReplicationProtectionContainerMappingArgs)(nil)).Elem()
}

// Protection container mapping object.
type LookupReplicationProtectionContainerMappingResultOutput struct{ *pulumi.OutputState }

func (LookupReplicationProtectionContainerMappingResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReplicationProtectionContainerMappingResult)(nil)).Elem()
}

func (o LookupReplicationProtectionContainerMappingResultOutput) ToLookupReplicationProtectionContainerMappingResultOutput() LookupReplicationProtectionContainerMappingResultOutput {
	return o
}

func (o LookupReplicationProtectionContainerMappingResultOutput) ToLookupReplicationProtectionContainerMappingResultOutputWithContext(ctx context.Context) LookupReplicationProtectionContainerMappingResultOutput {
	return o
}

// Resource Id
func (o LookupReplicationProtectionContainerMappingResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReplicationProtectionContainerMappingResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource Location
func (o LookupReplicationProtectionContainerMappingResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupReplicationProtectionContainerMappingResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource Name
func (o LookupReplicationProtectionContainerMappingResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReplicationProtectionContainerMappingResult) string { return v.Name }).(pulumi.StringOutput)
}

// The custom data.
func (o LookupReplicationProtectionContainerMappingResultOutput) Properties() ProtectionContainerMappingPropertiesResponseOutput {
	return o.ApplyT(func(v LookupReplicationProtectionContainerMappingResult) ProtectionContainerMappingPropertiesResponse {
		return v.Properties
	}).(ProtectionContainerMappingPropertiesResponseOutput)
}

// Resource Type
func (o LookupReplicationProtectionContainerMappingResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReplicationProtectionContainerMappingResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupReplicationProtectionContainerMappingResultOutput{})
}
