// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20221101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// list source configurations for a Linker.
func ListLinkerConfigurations(ctx *pulumi.Context, args *ListLinkerConfigurationsArgs, opts ...pulumi.InvokeOption) (*ListLinkerConfigurationsResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListLinkerConfigurationsResult
	err := ctx.Invoke("azure-native:servicelinker/v20221101preview:listLinkerConfigurations", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListLinkerConfigurationsArgs struct {
	// The name Linker resource.
	LinkerName string `pulumi:"linkerName"`
	// The fully qualified Azure Resource manager identifier of the resource to be connected.
	ResourceUri string `pulumi:"resourceUri"`
}

// Configurations for source resource, include appSettings, connectionString and serviceBindings
type ListLinkerConfigurationsResult struct {
	// The configuration properties for source resource.
	Configurations []SourceConfigurationResponse `pulumi:"configurations"`
}

func ListLinkerConfigurationsOutput(ctx *pulumi.Context, args ListLinkerConfigurationsOutputArgs, opts ...pulumi.InvokeOption) ListLinkerConfigurationsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListLinkerConfigurationsResult, error) {
			args := v.(ListLinkerConfigurationsArgs)
			r, err := ListLinkerConfigurations(ctx, &args, opts...)
			var s ListLinkerConfigurationsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListLinkerConfigurationsResultOutput)
}

type ListLinkerConfigurationsOutputArgs struct {
	// The name Linker resource.
	LinkerName pulumi.StringInput `pulumi:"linkerName"`
	// The fully qualified Azure Resource manager identifier of the resource to be connected.
	ResourceUri pulumi.StringInput `pulumi:"resourceUri"`
}

func (ListLinkerConfigurationsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListLinkerConfigurationsArgs)(nil)).Elem()
}

// Configurations for source resource, include appSettings, connectionString and serviceBindings
type ListLinkerConfigurationsResultOutput struct{ *pulumi.OutputState }

func (ListLinkerConfigurationsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListLinkerConfigurationsResult)(nil)).Elem()
}

func (o ListLinkerConfigurationsResultOutput) ToListLinkerConfigurationsResultOutput() ListLinkerConfigurationsResultOutput {
	return o
}

func (o ListLinkerConfigurationsResultOutput) ToListLinkerConfigurationsResultOutputWithContext(ctx context.Context) ListLinkerConfigurationsResultOutput {
	return o
}

func (o ListLinkerConfigurationsResultOutput) ToOutput(ctx context.Context) pulumix.Output[ListLinkerConfigurationsResult] {
	return pulumix.Output[ListLinkerConfigurationsResult]{
		OutputState: o.OutputState,
	}
}

// The configuration properties for source resource.
func (o ListLinkerConfigurationsResultOutput) Configurations() SourceConfigurationResponseArrayOutput {
	return o.ApplyT(func(v ListLinkerConfigurationsResult) []SourceConfigurationResponse { return v.Configurations }).(SourceConfigurationResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListLinkerConfigurationsResultOutput{})
}
