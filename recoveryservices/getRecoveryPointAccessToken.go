// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package recoveryservices

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Azure REST API version: 2023-01-15.
//
// Other available API versions: 2018-12-20, 2021-11-15.
func GetRecoveryPointAccessToken(ctx *pulumi.Context, args *GetRecoveryPointAccessTokenArgs, opts ...pulumi.InvokeOption) (*GetRecoveryPointAccessTokenResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv GetRecoveryPointAccessTokenResult
	err := ctx.Invoke("azure-native:recoveryservices:getRecoveryPointAccessToken", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetRecoveryPointAccessTokenArgs struct {
	// Name of the container.
	ContainerName string `pulumi:"containerName"`
	// Optional ETag.
	ETag *string `pulumi:"eTag"`
	// Fabric name associated with the container.
	FabricName string `pulumi:"fabricName"`
	// Resource location.
	Location *string `pulumi:"location"`
	// AADPropertiesResource properties
	Properties *AADProperties `pulumi:"properties"`
	// Name of the Protected Item.
	ProtectedItemName string `pulumi:"protectedItemName"`
	// Recovery Point Id
	RecoveryPointId string `pulumi:"recoveryPointId"`
	// The name of the resource group where the recovery services vault is present.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The name of the recovery services vault.
	VaultName string `pulumi:"vaultName"`
}

type GetRecoveryPointAccessTokenResult struct {
	// Optional ETag.
	ETag *string `pulumi:"eTag"`
	// Resource Id represents the complete path to the resource.
	Id string `pulumi:"id"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Resource name associated with the resource.
	Name string `pulumi:"name"`
	// CrrAccessTokenResource properties
	Properties WorkloadCrrAccessTokenResponse `pulumi:"properties"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type represents the complete path of the form Namespace/ResourceType/ResourceType/...
	Type string `pulumi:"type"`
}

func GetRecoveryPointAccessTokenOutput(ctx *pulumi.Context, args GetRecoveryPointAccessTokenOutputArgs, opts ...pulumi.InvokeOption) GetRecoveryPointAccessTokenResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetRecoveryPointAccessTokenResult, error) {
			args := v.(GetRecoveryPointAccessTokenArgs)
			r, err := GetRecoveryPointAccessToken(ctx, &args, opts...)
			var s GetRecoveryPointAccessTokenResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetRecoveryPointAccessTokenResultOutput)
}

type GetRecoveryPointAccessTokenOutputArgs struct {
	// Name of the container.
	ContainerName pulumi.StringInput `pulumi:"containerName"`
	// Optional ETag.
	ETag pulumi.StringPtrInput `pulumi:"eTag"`
	// Fabric name associated with the container.
	FabricName pulumi.StringInput `pulumi:"fabricName"`
	// Resource location.
	Location pulumi.StringPtrInput `pulumi:"location"`
	// AADPropertiesResource properties
	Properties AADPropertiesPtrInput `pulumi:"properties"`
	// Name of the Protected Item.
	ProtectedItemName pulumi.StringInput `pulumi:"protectedItemName"`
	// Recovery Point Id
	RecoveryPointId pulumi.StringInput `pulumi:"recoveryPointId"`
	// The name of the resource group where the recovery services vault is present.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags pulumi.StringMapInput `pulumi:"tags"`
	// The name of the recovery services vault.
	VaultName pulumi.StringInput `pulumi:"vaultName"`
}

func (GetRecoveryPointAccessTokenOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRecoveryPointAccessTokenArgs)(nil)).Elem()
}

type GetRecoveryPointAccessTokenResultOutput struct{ *pulumi.OutputState }

func (GetRecoveryPointAccessTokenResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRecoveryPointAccessTokenResult)(nil)).Elem()
}

func (o GetRecoveryPointAccessTokenResultOutput) ToGetRecoveryPointAccessTokenResultOutput() GetRecoveryPointAccessTokenResultOutput {
	return o
}

func (o GetRecoveryPointAccessTokenResultOutput) ToGetRecoveryPointAccessTokenResultOutputWithContext(ctx context.Context) GetRecoveryPointAccessTokenResultOutput {
	return o
}

// Optional ETag.
func (o GetRecoveryPointAccessTokenResultOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRecoveryPointAccessTokenResult) *string { return v.ETag }).(pulumi.StringPtrOutput)
}

// Resource Id represents the complete path to the resource.
func (o GetRecoveryPointAccessTokenResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetRecoveryPointAccessTokenResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource location.
func (o GetRecoveryPointAccessTokenResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRecoveryPointAccessTokenResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource name associated with the resource.
func (o GetRecoveryPointAccessTokenResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetRecoveryPointAccessTokenResult) string { return v.Name }).(pulumi.StringOutput)
}

// CrrAccessTokenResource properties
func (o GetRecoveryPointAccessTokenResultOutput) Properties() WorkloadCrrAccessTokenResponseOutput {
	return o.ApplyT(func(v GetRecoveryPointAccessTokenResult) WorkloadCrrAccessTokenResponse { return v.Properties }).(WorkloadCrrAccessTokenResponseOutput)
}

// Resource tags.
func (o GetRecoveryPointAccessTokenResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetRecoveryPointAccessTokenResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type represents the complete path of the form Namespace/ResourceType/ResourceType/...
func (o GetRecoveryPointAccessTokenResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GetRecoveryPointAccessTokenResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetRecoveryPointAccessTokenResultOutput{})
}
