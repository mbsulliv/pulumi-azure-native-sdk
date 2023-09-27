// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20200701privatepreview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets information about a configuration of server.
func LookupConfiguration(ctx *pulumi.Context, args *LookupConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupConfigurationResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupConfigurationResult
	err := ctx.Invoke("azure-native:dbformysql/v20200701privatepreview:getConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConfigurationArgs struct {
	// The name of the server configuration.
	ConfigurationName string `pulumi:"configurationName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
}

// Represents a Configuration.
type LookupConfigurationResult struct {
	// Allowed values of the configuration.
	AllowedValues string `pulumi:"allowedValues"`
	// Data type of the configuration.
	DataType string `pulumi:"dataType"`
	// Default value of the configuration.
	DefaultValue string `pulumi:"defaultValue"`
	// Description of the configuration.
	Description string `pulumi:"description"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// If is the configuration pending restart or not.
	IsConfigPendingRestart string `pulumi:"isConfigPendingRestart"`
	// If is the configuration dynamic.
	IsDynamicConfig string `pulumi:"isDynamicConfig"`
	// If is the configuration read only.
	IsReadOnly string `pulumi:"isReadOnly"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Source of the configuration.
	Source *string `pulumi:"source"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// Value of the configuration.
	Value *string `pulumi:"value"`
}

func LookupConfigurationOutput(ctx *pulumi.Context, args LookupConfigurationOutputArgs, opts ...pulumi.InvokeOption) LookupConfigurationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupConfigurationResult, error) {
			args := v.(LookupConfigurationArgs)
			r, err := LookupConfiguration(ctx, &args, opts...)
			var s LookupConfigurationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupConfigurationResultOutput)
}

type LookupConfigurationOutputArgs struct {
	// The name of the server configuration.
	ConfigurationName pulumi.StringInput `pulumi:"configurationName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName pulumi.StringInput `pulumi:"serverName"`
}

func (LookupConfigurationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConfigurationArgs)(nil)).Elem()
}

// Represents a Configuration.
type LookupConfigurationResultOutput struct{ *pulumi.OutputState }

func (LookupConfigurationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConfigurationResult)(nil)).Elem()
}

func (o LookupConfigurationResultOutput) ToLookupConfigurationResultOutput() LookupConfigurationResultOutput {
	return o
}

func (o LookupConfigurationResultOutput) ToLookupConfigurationResultOutputWithContext(ctx context.Context) LookupConfigurationResultOutput {
	return o
}

func (o LookupConfigurationResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupConfigurationResult] {
	return pulumix.Output[LookupConfigurationResult]{
		OutputState: o.OutputState,
	}
}

// Allowed values of the configuration.
func (o LookupConfigurationResultOutput) AllowedValues() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigurationResult) string { return v.AllowedValues }).(pulumi.StringOutput)
}

// Data type of the configuration.
func (o LookupConfigurationResultOutput) DataType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigurationResult) string { return v.DataType }).(pulumi.StringOutput)
}

// Default value of the configuration.
func (o LookupConfigurationResultOutput) DefaultValue() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigurationResult) string { return v.DefaultValue }).(pulumi.StringOutput)
}

// Description of the configuration.
func (o LookupConfigurationResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigurationResult) string { return v.Description }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupConfigurationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigurationResult) string { return v.Id }).(pulumi.StringOutput)
}

// If is the configuration pending restart or not.
func (o LookupConfigurationResultOutput) IsConfigPendingRestart() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigurationResult) string { return v.IsConfigPendingRestart }).(pulumi.StringOutput)
}

// If is the configuration dynamic.
func (o LookupConfigurationResultOutput) IsDynamicConfig() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigurationResult) string { return v.IsDynamicConfig }).(pulumi.StringOutput)
}

// If is the configuration read only.
func (o LookupConfigurationResultOutput) IsReadOnly() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigurationResult) string { return v.IsReadOnly }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupConfigurationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigurationResult) string { return v.Name }).(pulumi.StringOutput)
}

// Source of the configuration.
func (o LookupConfigurationResultOutput) Source() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConfigurationResult) *string { return v.Source }).(pulumi.StringPtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupConfigurationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigurationResult) string { return v.Type }).(pulumi.StringOutput)
}

// Value of the configuration.
func (o LookupConfigurationResultOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConfigurationResult) *string { return v.Value }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupConfigurationResultOutput{})
}
