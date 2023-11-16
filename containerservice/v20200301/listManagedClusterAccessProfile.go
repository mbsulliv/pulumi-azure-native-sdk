// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20200301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the accessProfile for the specified role name of the managed cluster with a specified resource group and name.
func ListManagedClusterAccessProfile(ctx *pulumi.Context, args *ListManagedClusterAccessProfileArgs, opts ...pulumi.InvokeOption) (*ListManagedClusterAccessProfileResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListManagedClusterAccessProfileResult
	err := ctx.Invoke("azure-native:containerservice/v20200301:listManagedClusterAccessProfile", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListManagedClusterAccessProfileArgs struct {
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the managed cluster resource.
	ResourceName string `pulumi:"resourceName"`
	// The name of the role for managed cluster accessProfile resource.
	RoleName string `pulumi:"roleName"`
}

// Managed cluster Access Profile.
type ListManagedClusterAccessProfileResult struct {
	// Resource Id
	Id string `pulumi:"id"`
	// Base64-encoded Kubernetes configuration file.
	KubeConfig *string `pulumi:"kubeConfig"`
	// Resource location
	Location string `pulumi:"location"`
	// Resource name
	Name string `pulumi:"name"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type string `pulumi:"type"`
}

func ListManagedClusterAccessProfileOutput(ctx *pulumi.Context, args ListManagedClusterAccessProfileOutputArgs, opts ...pulumi.InvokeOption) ListManagedClusterAccessProfileResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListManagedClusterAccessProfileResult, error) {
			args := v.(ListManagedClusterAccessProfileArgs)
			r, err := ListManagedClusterAccessProfile(ctx, &args, opts...)
			var s ListManagedClusterAccessProfileResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListManagedClusterAccessProfileResultOutput)
}

type ListManagedClusterAccessProfileOutputArgs struct {
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the managed cluster resource.
	ResourceName pulumi.StringInput `pulumi:"resourceName"`
	// The name of the role for managed cluster accessProfile resource.
	RoleName pulumi.StringInput `pulumi:"roleName"`
}

func (ListManagedClusterAccessProfileOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListManagedClusterAccessProfileArgs)(nil)).Elem()
}

// Managed cluster Access Profile.
type ListManagedClusterAccessProfileResultOutput struct{ *pulumi.OutputState }

func (ListManagedClusterAccessProfileResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListManagedClusterAccessProfileResult)(nil)).Elem()
}

func (o ListManagedClusterAccessProfileResultOutput) ToListManagedClusterAccessProfileResultOutput() ListManagedClusterAccessProfileResultOutput {
	return o
}

func (o ListManagedClusterAccessProfileResultOutput) ToListManagedClusterAccessProfileResultOutputWithContext(ctx context.Context) ListManagedClusterAccessProfileResultOutput {
	return o
}

// Resource Id
func (o ListManagedClusterAccessProfileResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ListManagedClusterAccessProfileResult) string { return v.Id }).(pulumi.StringOutput)
}

// Base64-encoded Kubernetes configuration file.
func (o ListManagedClusterAccessProfileResultOutput) KubeConfig() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListManagedClusterAccessProfileResult) *string { return v.KubeConfig }).(pulumi.StringPtrOutput)
}

// Resource location
func (o ListManagedClusterAccessProfileResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v ListManagedClusterAccessProfileResult) string { return v.Location }).(pulumi.StringOutput)
}

// Resource name
func (o ListManagedClusterAccessProfileResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ListManagedClusterAccessProfileResult) string { return v.Name }).(pulumi.StringOutput)
}

// Resource tags
func (o ListManagedClusterAccessProfileResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v ListManagedClusterAccessProfileResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type
func (o ListManagedClusterAccessProfileResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ListManagedClusterAccessProfileResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListManagedClusterAccessProfileResultOutput{})
}
