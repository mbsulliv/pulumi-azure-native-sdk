// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Dapr Component.
func LookupConnectedEnvironmentsDaprComponent(ctx *pulumi.Context, args *LookupConnectedEnvironmentsDaprComponentArgs, opts ...pulumi.InvokeOption) (*LookupConnectedEnvironmentsDaprComponentResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupConnectedEnvironmentsDaprComponentResult
	err := ctx.Invoke("azure-native:app/v20230401preview:getConnectedEnvironmentsDaprComponent", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupConnectedEnvironmentsDaprComponentArgs struct {
	// Name of the Dapr Component.
	ComponentName string `pulumi:"componentName"`
	// Name of the connected environment.
	ConnectedEnvironmentName string `pulumi:"connectedEnvironmentName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Dapr Component.
type LookupConnectedEnvironmentsDaprComponentResult struct {
	// Component type
	ComponentType *string `pulumi:"componentType"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// Boolean describing if the component errors are ignores
	IgnoreErrors *bool `pulumi:"ignoreErrors"`
	// Initialization timeout
	InitTimeout *string `pulumi:"initTimeout"`
	// Component metadata
	Metadata []DaprMetadataResponse `pulumi:"metadata"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Names of container apps that can use this Dapr component
	Scopes []string `pulumi:"scopes"`
	// Name of a Dapr component to retrieve component secrets from
	SecretStoreComponent *string `pulumi:"secretStoreComponent"`
	// Collection of secrets used by a Dapr component
	Secrets []SecretResponse `pulumi:"secrets"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// Component version
	Version *string `pulumi:"version"`
}

// Defaults sets the appropriate defaults for LookupConnectedEnvironmentsDaprComponentResult
func (val *LookupConnectedEnvironmentsDaprComponentResult) Defaults() *LookupConnectedEnvironmentsDaprComponentResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.IgnoreErrors == nil {
		ignoreErrors_ := false
		tmp.IgnoreErrors = &ignoreErrors_
	}
	return &tmp
}

func LookupConnectedEnvironmentsDaprComponentOutput(ctx *pulumi.Context, args LookupConnectedEnvironmentsDaprComponentOutputArgs, opts ...pulumi.InvokeOption) LookupConnectedEnvironmentsDaprComponentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupConnectedEnvironmentsDaprComponentResult, error) {
			args := v.(LookupConnectedEnvironmentsDaprComponentArgs)
			r, err := LookupConnectedEnvironmentsDaprComponent(ctx, &args, opts...)
			var s LookupConnectedEnvironmentsDaprComponentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupConnectedEnvironmentsDaprComponentResultOutput)
}

type LookupConnectedEnvironmentsDaprComponentOutputArgs struct {
	// Name of the Dapr Component.
	ComponentName pulumi.StringInput `pulumi:"componentName"`
	// Name of the connected environment.
	ConnectedEnvironmentName pulumi.StringInput `pulumi:"connectedEnvironmentName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupConnectedEnvironmentsDaprComponentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConnectedEnvironmentsDaprComponentArgs)(nil)).Elem()
}

// Dapr Component.
type LookupConnectedEnvironmentsDaprComponentResultOutput struct{ *pulumi.OutputState }

func (LookupConnectedEnvironmentsDaprComponentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConnectedEnvironmentsDaprComponentResult)(nil)).Elem()
}

func (o LookupConnectedEnvironmentsDaprComponentResultOutput) ToLookupConnectedEnvironmentsDaprComponentResultOutput() LookupConnectedEnvironmentsDaprComponentResultOutput {
	return o
}

func (o LookupConnectedEnvironmentsDaprComponentResultOutput) ToLookupConnectedEnvironmentsDaprComponentResultOutputWithContext(ctx context.Context) LookupConnectedEnvironmentsDaprComponentResultOutput {
	return o
}

// Component type
func (o LookupConnectedEnvironmentsDaprComponentResultOutput) ComponentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConnectedEnvironmentsDaprComponentResult) *string { return v.ComponentType }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupConnectedEnvironmentsDaprComponentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectedEnvironmentsDaprComponentResult) string { return v.Id }).(pulumi.StringOutput)
}

// Boolean describing if the component errors are ignores
func (o LookupConnectedEnvironmentsDaprComponentResultOutput) IgnoreErrors() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupConnectedEnvironmentsDaprComponentResult) *bool { return v.IgnoreErrors }).(pulumi.BoolPtrOutput)
}

// Initialization timeout
func (o LookupConnectedEnvironmentsDaprComponentResultOutput) InitTimeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConnectedEnvironmentsDaprComponentResult) *string { return v.InitTimeout }).(pulumi.StringPtrOutput)
}

// Component metadata
func (o LookupConnectedEnvironmentsDaprComponentResultOutput) Metadata() DaprMetadataResponseArrayOutput {
	return o.ApplyT(func(v LookupConnectedEnvironmentsDaprComponentResult) []DaprMetadataResponse { return v.Metadata }).(DaprMetadataResponseArrayOutput)
}

// The name of the resource
func (o LookupConnectedEnvironmentsDaprComponentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectedEnvironmentsDaprComponentResult) string { return v.Name }).(pulumi.StringOutput)
}

// Names of container apps that can use this Dapr component
func (o LookupConnectedEnvironmentsDaprComponentResultOutput) Scopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupConnectedEnvironmentsDaprComponentResult) []string { return v.Scopes }).(pulumi.StringArrayOutput)
}

// Name of a Dapr component to retrieve component secrets from
func (o LookupConnectedEnvironmentsDaprComponentResultOutput) SecretStoreComponent() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConnectedEnvironmentsDaprComponentResult) *string { return v.SecretStoreComponent }).(pulumi.StringPtrOutput)
}

// Collection of secrets used by a Dapr component
func (o LookupConnectedEnvironmentsDaprComponentResultOutput) Secrets() SecretResponseArrayOutput {
	return o.ApplyT(func(v LookupConnectedEnvironmentsDaprComponentResult) []SecretResponse { return v.Secrets }).(SecretResponseArrayOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupConnectedEnvironmentsDaprComponentResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupConnectedEnvironmentsDaprComponentResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupConnectedEnvironmentsDaprComponentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectedEnvironmentsDaprComponentResult) string { return v.Type }).(pulumi.StringOutput)
}

// Component version
func (o LookupConnectedEnvironmentsDaprComponentResultOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConnectedEnvironmentsDaprComponentResult) *string { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupConnectedEnvironmentsDaprComponentResultOutput{})
}
