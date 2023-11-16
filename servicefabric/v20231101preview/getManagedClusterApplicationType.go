// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20231101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get a Service Fabric application type name resource created or in the process of being created in the Service Fabric managed cluster resource.
func LookupManagedClusterApplicationType(ctx *pulumi.Context, args *LookupManagedClusterApplicationTypeArgs, opts ...pulumi.InvokeOption) (*LookupManagedClusterApplicationTypeResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupManagedClusterApplicationTypeResult
	err := ctx.Invoke("azure-native:servicefabric/v20231101preview:getManagedClusterApplicationType", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupManagedClusterApplicationTypeArgs struct {
	// The name of the application type name resource.
	ApplicationTypeName string `pulumi:"applicationTypeName"`
	// The name of the cluster resource.
	ClusterName string `pulumi:"clusterName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The application type name resource
type LookupManagedClusterApplicationTypeResult struct {
	// Azure resource identifier.
	Id string `pulumi:"id"`
	// Resource location depends on the parent resource.
	Location *string `pulumi:"location"`
	// Azure resource name.
	Name string `pulumi:"name"`
	// The current deployment or provisioning state, which only appears in the response.
	ProvisioningState string `pulumi:"provisioningState"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Azure resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Azure resource type.
	Type string `pulumi:"type"`
}

func LookupManagedClusterApplicationTypeOutput(ctx *pulumi.Context, args LookupManagedClusterApplicationTypeOutputArgs, opts ...pulumi.InvokeOption) LookupManagedClusterApplicationTypeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupManagedClusterApplicationTypeResult, error) {
			args := v.(LookupManagedClusterApplicationTypeArgs)
			r, err := LookupManagedClusterApplicationType(ctx, &args, opts...)
			var s LookupManagedClusterApplicationTypeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupManagedClusterApplicationTypeResultOutput)
}

type LookupManagedClusterApplicationTypeOutputArgs struct {
	// The name of the application type name resource.
	ApplicationTypeName pulumi.StringInput `pulumi:"applicationTypeName"`
	// The name of the cluster resource.
	ClusterName pulumi.StringInput `pulumi:"clusterName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupManagedClusterApplicationTypeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagedClusterApplicationTypeArgs)(nil)).Elem()
}

// The application type name resource
type LookupManagedClusterApplicationTypeResultOutput struct{ *pulumi.OutputState }

func (LookupManagedClusterApplicationTypeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagedClusterApplicationTypeResult)(nil)).Elem()
}

func (o LookupManagedClusterApplicationTypeResultOutput) ToLookupManagedClusterApplicationTypeResultOutput() LookupManagedClusterApplicationTypeResultOutput {
	return o
}

func (o LookupManagedClusterApplicationTypeResultOutput) ToLookupManagedClusterApplicationTypeResultOutputWithContext(ctx context.Context) LookupManagedClusterApplicationTypeResultOutput {
	return o
}

// Azure resource identifier.
func (o LookupManagedClusterApplicationTypeResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedClusterApplicationTypeResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource location depends on the parent resource.
func (o LookupManagedClusterApplicationTypeResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagedClusterApplicationTypeResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// Azure resource name.
func (o LookupManagedClusterApplicationTypeResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedClusterApplicationTypeResult) string { return v.Name }).(pulumi.StringOutput)
}

// The current deployment or provisioning state, which only appears in the response.
func (o LookupManagedClusterApplicationTypeResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedClusterApplicationTypeResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupManagedClusterApplicationTypeResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupManagedClusterApplicationTypeResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Azure resource tags.
func (o LookupManagedClusterApplicationTypeResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupManagedClusterApplicationTypeResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Azure resource type.
func (o LookupManagedClusterApplicationTypeResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagedClusterApplicationTypeResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupManagedClusterApplicationTypeResultOutput{})
}
