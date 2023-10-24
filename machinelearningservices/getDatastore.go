// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package machinelearningservices

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Azure Resource Manager resource envelope.
// Azure REST API version: 2023-04-01.
//
// Other available API versions: 2021-03-01-preview, 2022-02-01-preview, 2023-04-01-preview, 2023-06-01-preview, 2023-08-01-preview, 2023-10-01.
func LookupDatastore(ctx *pulumi.Context, args *LookupDatastoreArgs, opts ...pulumi.InvokeOption) (*LookupDatastoreResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupDatastoreResult
	err := ctx.Invoke("azure-native:machinelearningservices:getDatastore", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDatastoreArgs struct {
	// Datastore name.
	Name string `pulumi:"name"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// Azure Resource Manager resource envelope.
type LookupDatastoreResult struct {
	// [Required] Additional attributes of the entity.
	DatastoreProperties interface{} `pulumi:"datastoreProperties"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupDatastoreOutput(ctx *pulumi.Context, args LookupDatastoreOutputArgs, opts ...pulumi.InvokeOption) LookupDatastoreResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDatastoreResult, error) {
			args := v.(LookupDatastoreArgs)
			r, err := LookupDatastore(ctx, &args, opts...)
			var s LookupDatastoreResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDatastoreResultOutput)
}

type LookupDatastoreOutputArgs struct {
	// Datastore name.
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupDatastoreOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatastoreArgs)(nil)).Elem()
}

// Azure Resource Manager resource envelope.
type LookupDatastoreResultOutput struct{ *pulumi.OutputState }

func (LookupDatastoreResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatastoreResult)(nil)).Elem()
}

func (o LookupDatastoreResultOutput) ToLookupDatastoreResultOutput() LookupDatastoreResultOutput {
	return o
}

func (o LookupDatastoreResultOutput) ToLookupDatastoreResultOutputWithContext(ctx context.Context) LookupDatastoreResultOutput {
	return o
}

func (o LookupDatastoreResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupDatastoreResult] {
	return pulumix.Output[LookupDatastoreResult]{
		OutputState: o.OutputState,
	}
}

// [Required] Additional attributes of the entity.
func (o LookupDatastoreResultOutput) DatastoreProperties() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupDatastoreResult) interface{} { return v.DatastoreProperties }).(pulumi.AnyOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupDatastoreResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatastoreResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupDatastoreResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatastoreResult) string { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupDatastoreResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupDatastoreResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupDatastoreResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatastoreResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDatastoreResultOutput{})
}
