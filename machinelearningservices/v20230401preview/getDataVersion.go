// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Azure Resource Manager resource envelope.
func LookupDataVersion(ctx *pulumi.Context, args *LookupDataVersionArgs, opts ...pulumi.InvokeOption) (*LookupDataVersionResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupDataVersionResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20230401preview:getDataVersion", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDataVersionArgs struct {
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
type LookupDataVersionResult struct {
	// [Required] Additional attributes of the entity.
	DataVersionBaseProperties interface{} `pulumi:"dataVersionBaseProperties"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupDataVersionOutput(ctx *pulumi.Context, args LookupDataVersionOutputArgs, opts ...pulumi.InvokeOption) LookupDataVersionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDataVersionResult, error) {
			args := v.(LookupDataVersionArgs)
			r, err := LookupDataVersion(ctx, &args, opts...)
			var s LookupDataVersionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDataVersionResultOutput)
}

type LookupDataVersionOutputArgs struct {
	// Container name.
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Version identifier.
	Version pulumi.StringInput `pulumi:"version"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupDataVersionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDataVersionArgs)(nil)).Elem()
}

// Azure Resource Manager resource envelope.
type LookupDataVersionResultOutput struct{ *pulumi.OutputState }

func (LookupDataVersionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDataVersionResult)(nil)).Elem()
}

func (o LookupDataVersionResultOutput) ToLookupDataVersionResultOutput() LookupDataVersionResultOutput {
	return o
}

func (o LookupDataVersionResultOutput) ToLookupDataVersionResultOutputWithContext(ctx context.Context) LookupDataVersionResultOutput {
	return o
}

func (o LookupDataVersionResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupDataVersionResult] {
	return pulumix.Output[LookupDataVersionResult]{
		OutputState: o.OutputState,
	}
}

// [Required] Additional attributes of the entity.
func (o LookupDataVersionResultOutput) DataVersionBaseProperties() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupDataVersionResult) interface{} { return v.DataVersionBaseProperties }).(pulumi.AnyOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupDataVersionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataVersionResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupDataVersionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataVersionResult) string { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupDataVersionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupDataVersionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupDataVersionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataVersionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDataVersionResultOutput{})
}
