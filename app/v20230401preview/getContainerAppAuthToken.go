// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Container App Auth Token.
func GetContainerAppAuthToken(ctx *pulumi.Context, args *GetContainerAppAuthTokenArgs, opts ...pulumi.InvokeOption) (*GetContainerAppAuthTokenResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv GetContainerAppAuthTokenResult
	err := ctx.Invoke("azure-native:app/v20230401preview:getContainerAppAuthToken", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetContainerAppAuthTokenArgs struct {
	// Name of the Container App.
	ContainerAppName string `pulumi:"containerAppName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Container App Auth Token.
type GetContainerAppAuthTokenResult struct {
	// Token expiration date.
	Expires string `pulumi:"expires"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Auth token value.
	Token string `pulumi:"token"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func GetContainerAppAuthTokenOutput(ctx *pulumi.Context, args GetContainerAppAuthTokenOutputArgs, opts ...pulumi.InvokeOption) GetContainerAppAuthTokenResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetContainerAppAuthTokenResult, error) {
			args := v.(GetContainerAppAuthTokenArgs)
			r, err := GetContainerAppAuthToken(ctx, &args, opts...)
			var s GetContainerAppAuthTokenResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetContainerAppAuthTokenResultOutput)
}

type GetContainerAppAuthTokenOutputArgs struct {
	// Name of the Container App.
	ContainerAppName pulumi.StringInput `pulumi:"containerAppName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (GetContainerAppAuthTokenOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetContainerAppAuthTokenArgs)(nil)).Elem()
}

// Container App Auth Token.
type GetContainerAppAuthTokenResultOutput struct{ *pulumi.OutputState }

func (GetContainerAppAuthTokenResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetContainerAppAuthTokenResult)(nil)).Elem()
}

func (o GetContainerAppAuthTokenResultOutput) ToGetContainerAppAuthTokenResultOutput() GetContainerAppAuthTokenResultOutput {
	return o
}

func (o GetContainerAppAuthTokenResultOutput) ToGetContainerAppAuthTokenResultOutputWithContext(ctx context.Context) GetContainerAppAuthTokenResultOutput {
	return o
}

// Token expiration date.
func (o GetContainerAppAuthTokenResultOutput) Expires() pulumi.StringOutput {
	return o.ApplyT(func(v GetContainerAppAuthTokenResult) string { return v.Expires }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o GetContainerAppAuthTokenResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetContainerAppAuthTokenResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o GetContainerAppAuthTokenResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v GetContainerAppAuthTokenResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o GetContainerAppAuthTokenResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetContainerAppAuthTokenResult) string { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o GetContainerAppAuthTokenResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v GetContainerAppAuthTokenResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o GetContainerAppAuthTokenResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetContainerAppAuthTokenResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Auth token value.
func (o GetContainerAppAuthTokenResultOutput) Token() pulumi.StringOutput {
	return o.ApplyT(func(v GetContainerAppAuthTokenResult) string { return v.Token }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o GetContainerAppAuthTokenResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GetContainerAppAuthTokenResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetContainerAppAuthTokenResultOutput{})
}
