// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20190501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get a batch configuration for an integration account.
func LookupIntegrationAccountBatchConfiguration(ctx *pulumi.Context, args *LookupIntegrationAccountBatchConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupIntegrationAccountBatchConfigurationResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupIntegrationAccountBatchConfigurationResult
	err := ctx.Invoke("azure-native:logic/v20190501:getIntegrationAccountBatchConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIntegrationAccountBatchConfigurationArgs struct {
	// The batch configuration name.
	BatchConfigurationName string `pulumi:"batchConfigurationName"`
	// The integration account name.
	IntegrationAccountName string `pulumi:"integrationAccountName"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The batch configuration resource definition.
type LookupIntegrationAccountBatchConfigurationResult struct {
	// The resource id.
	Id string `pulumi:"id"`
	// The resource location.
	Location *string `pulumi:"location"`
	// Gets the resource name.
	Name string `pulumi:"name"`
	// The batch configuration properties.
	Properties BatchConfigurationPropertiesResponse `pulumi:"properties"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Gets the resource type.
	Type string `pulumi:"type"`
}

func LookupIntegrationAccountBatchConfigurationOutput(ctx *pulumi.Context, args LookupIntegrationAccountBatchConfigurationOutputArgs, opts ...pulumi.InvokeOption) LookupIntegrationAccountBatchConfigurationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIntegrationAccountBatchConfigurationResult, error) {
			args := v.(LookupIntegrationAccountBatchConfigurationArgs)
			r, err := LookupIntegrationAccountBatchConfiguration(ctx, &args, opts...)
			var s LookupIntegrationAccountBatchConfigurationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIntegrationAccountBatchConfigurationResultOutput)
}

type LookupIntegrationAccountBatchConfigurationOutputArgs struct {
	// The batch configuration name.
	BatchConfigurationName pulumi.StringInput `pulumi:"batchConfigurationName"`
	// The integration account name.
	IntegrationAccountName pulumi.StringInput `pulumi:"integrationAccountName"`
	// The resource group name.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupIntegrationAccountBatchConfigurationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIntegrationAccountBatchConfigurationArgs)(nil)).Elem()
}

// The batch configuration resource definition.
type LookupIntegrationAccountBatchConfigurationResultOutput struct{ *pulumi.OutputState }

func (LookupIntegrationAccountBatchConfigurationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIntegrationAccountBatchConfigurationResult)(nil)).Elem()
}

func (o LookupIntegrationAccountBatchConfigurationResultOutput) ToLookupIntegrationAccountBatchConfigurationResultOutput() LookupIntegrationAccountBatchConfigurationResultOutput {
	return o
}

func (o LookupIntegrationAccountBatchConfigurationResultOutput) ToLookupIntegrationAccountBatchConfigurationResultOutputWithContext(ctx context.Context) LookupIntegrationAccountBatchConfigurationResultOutput {
	return o
}

// The resource id.
func (o LookupIntegrationAccountBatchConfigurationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIntegrationAccountBatchConfigurationResult) string { return v.Id }).(pulumi.StringOutput)
}

// The resource location.
func (o LookupIntegrationAccountBatchConfigurationResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIntegrationAccountBatchConfigurationResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// Gets the resource name.
func (o LookupIntegrationAccountBatchConfigurationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIntegrationAccountBatchConfigurationResult) string { return v.Name }).(pulumi.StringOutput)
}

// The batch configuration properties.
func (o LookupIntegrationAccountBatchConfigurationResultOutput) Properties() BatchConfigurationPropertiesResponseOutput {
	return o.ApplyT(func(v LookupIntegrationAccountBatchConfigurationResult) BatchConfigurationPropertiesResponse {
		return v.Properties
	}).(BatchConfigurationPropertiesResponseOutput)
}

// The resource tags.
func (o LookupIntegrationAccountBatchConfigurationResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupIntegrationAccountBatchConfigurationResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Gets the resource type.
func (o LookupIntegrationAccountBatchConfigurationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIntegrationAccountBatchConfigurationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIntegrationAccountBatchConfigurationResultOutput{})
}
