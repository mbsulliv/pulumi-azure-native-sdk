// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Certificate used for Custom Domain bindings of Container Apps in a Managed Environment
func LookupConnectedEnvironmentsCertificate(ctx *pulumi.Context, args *LookupConnectedEnvironmentsCertificateArgs, opts ...pulumi.InvokeOption) (*LookupConnectedEnvironmentsCertificateResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupConnectedEnvironmentsCertificateResult
	err := ctx.Invoke("azure-native:app/v20230401preview:getConnectedEnvironmentsCertificate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConnectedEnvironmentsCertificateArgs struct {
	// Name of the Certificate.
	CertificateName string `pulumi:"certificateName"`
	// Name of the Connected Environment.
	ConnectedEnvironmentName string `pulumi:"connectedEnvironmentName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Certificate used for Custom Domain bindings of Container Apps in a Managed Environment
type LookupConnectedEnvironmentsCertificateResult struct {
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Certificate resource specific properties
	Properties CertificateResponseProperties `pulumi:"properties"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupConnectedEnvironmentsCertificateOutput(ctx *pulumi.Context, args LookupConnectedEnvironmentsCertificateOutputArgs, opts ...pulumi.InvokeOption) LookupConnectedEnvironmentsCertificateResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupConnectedEnvironmentsCertificateResult, error) {
			args := v.(LookupConnectedEnvironmentsCertificateArgs)
			r, err := LookupConnectedEnvironmentsCertificate(ctx, &args, opts...)
			var s LookupConnectedEnvironmentsCertificateResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupConnectedEnvironmentsCertificateResultOutput)
}

type LookupConnectedEnvironmentsCertificateOutputArgs struct {
	// Name of the Certificate.
	CertificateName pulumi.StringInput `pulumi:"certificateName"`
	// Name of the Connected Environment.
	ConnectedEnvironmentName pulumi.StringInput `pulumi:"connectedEnvironmentName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupConnectedEnvironmentsCertificateOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConnectedEnvironmentsCertificateArgs)(nil)).Elem()
}

// Certificate used for Custom Domain bindings of Container Apps in a Managed Environment
type LookupConnectedEnvironmentsCertificateResultOutput struct{ *pulumi.OutputState }

func (LookupConnectedEnvironmentsCertificateResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConnectedEnvironmentsCertificateResult)(nil)).Elem()
}

func (o LookupConnectedEnvironmentsCertificateResultOutput) ToLookupConnectedEnvironmentsCertificateResultOutput() LookupConnectedEnvironmentsCertificateResultOutput {
	return o
}

func (o LookupConnectedEnvironmentsCertificateResultOutput) ToLookupConnectedEnvironmentsCertificateResultOutputWithContext(ctx context.Context) LookupConnectedEnvironmentsCertificateResultOutput {
	return o
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupConnectedEnvironmentsCertificateResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectedEnvironmentsCertificateResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupConnectedEnvironmentsCertificateResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectedEnvironmentsCertificateResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupConnectedEnvironmentsCertificateResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectedEnvironmentsCertificateResult) string { return v.Name }).(pulumi.StringOutput)
}

// Certificate resource specific properties
func (o LookupConnectedEnvironmentsCertificateResultOutput) Properties() CertificateResponsePropertiesOutput {
	return o.ApplyT(func(v LookupConnectedEnvironmentsCertificateResult) CertificateResponseProperties {
		return v.Properties
	}).(CertificateResponsePropertiesOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupConnectedEnvironmentsCertificateResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupConnectedEnvironmentsCertificateResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupConnectedEnvironmentsCertificateResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupConnectedEnvironmentsCertificateResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupConnectedEnvironmentsCertificateResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectedEnvironmentsCertificateResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupConnectedEnvironmentsCertificateResultOutput{})
}
