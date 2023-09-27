// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package security

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Details of a specific cloud account connector
// Azure REST API version: 2020-01-01-preview.
func LookupConnector(ctx *pulumi.Context, args *LookupConnectorArgs, opts ...pulumi.InvokeOption) (*LookupConnectorResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupConnectorResult
	err := ctx.Invoke("azure-native:security:getConnector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConnectorArgs struct {
	// Name of the cloud account connector
	ConnectorName string `pulumi:"connectorName"`
}

// The connector setting
type LookupConnectorResult struct {
	// Settings for authentication management, these settings are relevant only for the cloud connector.
	AuthenticationDetails interface{} `pulumi:"authenticationDetails"`
	// Settings for hybrid compute management. These settings are relevant only for Arc autoProvision (Hybrid Compute).
	HybridComputeSettings *HybridComputeSettingsPropertiesResponse `pulumi:"hybridComputeSettings"`
	// Resource Id
	Id string `pulumi:"id"`
	// Resource name
	Name string `pulumi:"name"`
	// Resource type
	Type string `pulumi:"type"`
}

func LookupConnectorOutput(ctx *pulumi.Context, args LookupConnectorOutputArgs, opts ...pulumi.InvokeOption) LookupConnectorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupConnectorResult, error) {
			args := v.(LookupConnectorArgs)
			r, err := LookupConnector(ctx, &args, opts...)
			var s LookupConnectorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupConnectorResultOutput)
}

type LookupConnectorOutputArgs struct {
	// Name of the cloud account connector
	ConnectorName pulumi.StringInput `pulumi:"connectorName"`
}

func (LookupConnectorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConnectorArgs)(nil)).Elem()
}

// The connector setting
type LookupConnectorResultOutput struct{ *pulumi.OutputState }

func (LookupConnectorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConnectorResult)(nil)).Elem()
}

func (o LookupConnectorResultOutput) ToLookupConnectorResultOutput() LookupConnectorResultOutput {
	return o
}

func (o LookupConnectorResultOutput) ToLookupConnectorResultOutputWithContext(ctx context.Context) LookupConnectorResultOutput {
	return o
}

func (o LookupConnectorResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupConnectorResult] {
	return pulumix.Output[LookupConnectorResult]{
		OutputState: o.OutputState,
	}
}

// Settings for authentication management, these settings are relevant only for the cloud connector.
func (o LookupConnectorResultOutput) AuthenticationDetails() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupConnectorResult) interface{} { return v.AuthenticationDetails }).(pulumi.AnyOutput)
}

// Settings for hybrid compute management. These settings are relevant only for Arc autoProvision (Hybrid Compute).
func (o LookupConnectorResultOutput) HybridComputeSettings() HybridComputeSettingsPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupConnectorResult) *HybridComputeSettingsPropertiesResponse { return v.HybridComputeSettings }).(HybridComputeSettingsPropertiesResponsePtrOutput)
}

// Resource Id
func (o LookupConnectorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectorResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource name
func (o LookupConnectorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectorResult) string { return v.Name }).(pulumi.StringOutput)
}

// Resource type
func (o LookupConnectorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectorResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupConnectorResultOutput{})
}
