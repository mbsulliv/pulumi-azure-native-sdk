// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210301preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Azure Resource Manager resource envelope.
func LookupEnvironmentSpecificationVersion(ctx *pulumi.Context, args *LookupEnvironmentSpecificationVersionArgs, opts ...pulumi.InvokeOption) (*LookupEnvironmentSpecificationVersionResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupEnvironmentSpecificationVersionResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20210301preview:getEnvironmentSpecificationVersion", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEnvironmentSpecificationVersionArgs struct {
	// Container name.
	Name string `pulumi:"name"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Version identifier.
	Version string `pulumi:"version"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// Azure Resource Manager resource envelope.
type LookupEnvironmentSpecificationVersionResult struct {
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// [Required] Additional attributes of the entity.
	Properties EnvironmentSpecificationVersionResponse `pulumi:"properties"`
	// System data associated with resource provider
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupEnvironmentSpecificationVersionOutput(ctx *pulumi.Context, args LookupEnvironmentSpecificationVersionOutputArgs, opts ...pulumi.InvokeOption) LookupEnvironmentSpecificationVersionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEnvironmentSpecificationVersionResult, error) {
			args := v.(LookupEnvironmentSpecificationVersionArgs)
			r, err := LookupEnvironmentSpecificationVersion(ctx, &args, opts...)
			var s LookupEnvironmentSpecificationVersionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupEnvironmentSpecificationVersionResultOutput)
}

type LookupEnvironmentSpecificationVersionOutputArgs struct {
	// Container name.
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Version identifier.
	Version pulumi.StringInput `pulumi:"version"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupEnvironmentSpecificationVersionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEnvironmentSpecificationVersionArgs)(nil)).Elem()
}

// Azure Resource Manager resource envelope.
type LookupEnvironmentSpecificationVersionResultOutput struct{ *pulumi.OutputState }

func (LookupEnvironmentSpecificationVersionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEnvironmentSpecificationVersionResult)(nil)).Elem()
}

func (o LookupEnvironmentSpecificationVersionResultOutput) ToLookupEnvironmentSpecificationVersionResultOutput() LookupEnvironmentSpecificationVersionResultOutput {
	return o
}

func (o LookupEnvironmentSpecificationVersionResultOutput) ToLookupEnvironmentSpecificationVersionResultOutputWithContext(ctx context.Context) LookupEnvironmentSpecificationVersionResultOutput {
	return o
}

func (o LookupEnvironmentSpecificationVersionResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupEnvironmentSpecificationVersionResult] {
	return pulumix.Output[LookupEnvironmentSpecificationVersionResult]{
		OutputState: o.OutputState,
	}
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupEnvironmentSpecificationVersionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvironmentSpecificationVersionResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupEnvironmentSpecificationVersionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvironmentSpecificationVersionResult) string { return v.Name }).(pulumi.StringOutput)
}

// [Required] Additional attributes of the entity.
func (o LookupEnvironmentSpecificationVersionResultOutput) Properties() EnvironmentSpecificationVersionResponseOutput {
	return o.ApplyT(func(v LookupEnvironmentSpecificationVersionResult) EnvironmentSpecificationVersionResponse {
		return v.Properties
	}).(EnvironmentSpecificationVersionResponseOutput)
}

// System data associated with resource provider
func (o LookupEnvironmentSpecificationVersionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupEnvironmentSpecificationVersionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupEnvironmentSpecificationVersionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvironmentSpecificationVersionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEnvironmentSpecificationVersionResultOutput{})
}
