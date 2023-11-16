// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides the details of the protection intent up item. This is an asynchronous operation. To know the status of the operation,
// call the GetItemOperationResult API.
func LookupProtectionIntent(ctx *pulumi.Context, args *LookupProtectionIntentArgs, opts ...pulumi.InvokeOption) (*LookupProtectionIntentResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupProtectionIntentResult
	err := ctx.Invoke("azure-native:recoveryservices/v20230601:getProtectionIntent", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupProtectionIntentArgs struct {
	// Fabric name associated with the backed up item.
	FabricName string `pulumi:"fabricName"`
	// Backed up item name whose details are to be fetched.
	IntentObjectName string `pulumi:"intentObjectName"`
	// The name of the resource group where the recovery services vault is present.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the recovery services vault.
	VaultName string `pulumi:"vaultName"`
}

// Base class for backup ProtectionIntent.
type LookupProtectionIntentResult struct {
	// Optional ETag.
	ETag *string `pulumi:"eTag"`
	// Resource Id represents the complete path to the resource.
	Id string `pulumi:"id"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Resource name associated with the resource.
	Name string `pulumi:"name"`
	// ProtectionIntentResource properties
	Properties interface{} `pulumi:"properties"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type represents the complete path of the form Namespace/ResourceType/ResourceType/...
	Type string `pulumi:"type"`
}

func LookupProtectionIntentOutput(ctx *pulumi.Context, args LookupProtectionIntentOutputArgs, opts ...pulumi.InvokeOption) LookupProtectionIntentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupProtectionIntentResult, error) {
			args := v.(LookupProtectionIntentArgs)
			r, err := LookupProtectionIntent(ctx, &args, opts...)
			var s LookupProtectionIntentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupProtectionIntentResultOutput)
}

type LookupProtectionIntentOutputArgs struct {
	// Fabric name associated with the backed up item.
	FabricName pulumi.StringInput `pulumi:"fabricName"`
	// Backed up item name whose details are to be fetched.
	IntentObjectName pulumi.StringInput `pulumi:"intentObjectName"`
	// The name of the resource group where the recovery services vault is present.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the recovery services vault.
	VaultName pulumi.StringInput `pulumi:"vaultName"`
}

func (LookupProtectionIntentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProtectionIntentArgs)(nil)).Elem()
}

// Base class for backup ProtectionIntent.
type LookupProtectionIntentResultOutput struct{ *pulumi.OutputState }

func (LookupProtectionIntentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProtectionIntentResult)(nil)).Elem()
}

func (o LookupProtectionIntentResultOutput) ToLookupProtectionIntentResultOutput() LookupProtectionIntentResultOutput {
	return o
}

func (o LookupProtectionIntentResultOutput) ToLookupProtectionIntentResultOutputWithContext(ctx context.Context) LookupProtectionIntentResultOutput {
	return o
}

// Optional ETag.
func (o LookupProtectionIntentResultOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProtectionIntentResult) *string { return v.ETag }).(pulumi.StringPtrOutput)
}

// Resource Id represents the complete path to the resource.
func (o LookupProtectionIntentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProtectionIntentResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource location.
func (o LookupProtectionIntentResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProtectionIntentResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource name associated with the resource.
func (o LookupProtectionIntentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProtectionIntentResult) string { return v.Name }).(pulumi.StringOutput)
}

// ProtectionIntentResource properties
func (o LookupProtectionIntentResultOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupProtectionIntentResult) interface{} { return v.Properties }).(pulumi.AnyOutput)
}

// Resource tags.
func (o LookupProtectionIntentResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupProtectionIntentResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type represents the complete path of the form Namespace/ResourceType/ResourceType/...
func (o LookupProtectionIntentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProtectionIntentResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupProtectionIntentResultOutput{})
}
