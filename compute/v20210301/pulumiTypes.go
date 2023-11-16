// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = utilities.GetEnvOrDefault

// LogAnalytics output properties
type LogAnalyticsOutputResponse struct {
	// Output file Uri path to blob container.
	Output string `pulumi:"output"`
}

// LogAnalytics output properties
type LogAnalyticsOutputResponseOutput struct{ *pulumi.OutputState }

func (LogAnalyticsOutputResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LogAnalyticsOutputResponse)(nil)).Elem()
}

func (o LogAnalyticsOutputResponseOutput) ToLogAnalyticsOutputResponseOutput() LogAnalyticsOutputResponseOutput {
	return o
}

func (o LogAnalyticsOutputResponseOutput) ToLogAnalyticsOutputResponseOutputWithContext(ctx context.Context) LogAnalyticsOutputResponseOutput {
	return o
}

// Output file Uri path to blob container.
func (o LogAnalyticsOutputResponseOutput) Output() pulumi.StringOutput {
	return o.ApplyT(func(v LogAnalyticsOutputResponse) string { return v.Output }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LogAnalyticsOutputResponseOutput{})
}
