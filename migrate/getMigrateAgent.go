// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package migrate

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets the details of the modernizeProject agent.
// Azure REST API version: 2022-05-01-preview.
func LookupMigrateAgent(ctx *pulumi.Context, args *LookupMigrateAgentArgs, opts ...pulumi.InvokeOption) (*LookupMigrateAgentResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupMigrateAgentResult
	err := ctx.Invoke("azure-native:migrate:getMigrateAgent", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMigrateAgentArgs struct {
	// MigrateAgent name.
	AgentName string `pulumi:"agentName"`
	// ModernizeProject name.
	ModernizeProjectName string `pulumi:"modernizeProjectName"`
	// Name of the Azure Resource Group that project is part of.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Azure Subscription Id in which project was created.
	SubscriptionId *string `pulumi:"subscriptionId"`
}

// MigrateAgent model.
type LookupMigrateAgentResult struct {
	// Gets or sets the Id of the resource.
	Id string `pulumi:"id"`
	// Gets or sets the name of the resource.
	Name string `pulumi:"name"`
	// MigrateAgent model properties.
	Properties MigrateAgentModelPropertiesResponse `pulumi:"properties"`
	SystemData MigrateAgentModelResponseSystemData `pulumi:"systemData"`
	// Gets or sets the resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Gets or sets the type of the resource.
	Type string `pulumi:"type"`
}

func LookupMigrateAgentOutput(ctx *pulumi.Context, args LookupMigrateAgentOutputArgs, opts ...pulumi.InvokeOption) LookupMigrateAgentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMigrateAgentResult, error) {
			args := v.(LookupMigrateAgentArgs)
			r, err := LookupMigrateAgent(ctx, &args, opts...)
			var s LookupMigrateAgentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMigrateAgentResultOutput)
}

type LookupMigrateAgentOutputArgs struct {
	// MigrateAgent name.
	AgentName pulumi.StringInput `pulumi:"agentName"`
	// ModernizeProject name.
	ModernizeProjectName pulumi.StringInput `pulumi:"modernizeProjectName"`
	// Name of the Azure Resource Group that project is part of.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Azure Subscription Id in which project was created.
	SubscriptionId pulumi.StringPtrInput `pulumi:"subscriptionId"`
}

func (LookupMigrateAgentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMigrateAgentArgs)(nil)).Elem()
}

// MigrateAgent model.
type LookupMigrateAgentResultOutput struct{ *pulumi.OutputState }

func (LookupMigrateAgentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMigrateAgentResult)(nil)).Elem()
}

func (o LookupMigrateAgentResultOutput) ToLookupMigrateAgentResultOutput() LookupMigrateAgentResultOutput {
	return o
}

func (o LookupMigrateAgentResultOutput) ToLookupMigrateAgentResultOutputWithContext(ctx context.Context) LookupMigrateAgentResultOutput {
	return o
}

func (o LookupMigrateAgentResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupMigrateAgentResult] {
	return pulumix.Output[LookupMigrateAgentResult]{
		OutputState: o.OutputState,
	}
}

// Gets or sets the Id of the resource.
func (o LookupMigrateAgentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMigrateAgentResult) string { return v.Id }).(pulumi.StringOutput)
}

// Gets or sets the name of the resource.
func (o LookupMigrateAgentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMigrateAgentResult) string { return v.Name }).(pulumi.StringOutput)
}

// MigrateAgent model properties.
func (o LookupMigrateAgentResultOutput) Properties() MigrateAgentModelPropertiesResponseOutput {
	return o.ApplyT(func(v LookupMigrateAgentResult) MigrateAgentModelPropertiesResponse { return v.Properties }).(MigrateAgentModelPropertiesResponseOutput)
}

func (o LookupMigrateAgentResultOutput) SystemData() MigrateAgentModelResponseSystemDataOutput {
	return o.ApplyT(func(v LookupMigrateAgentResult) MigrateAgentModelResponseSystemData { return v.SystemData }).(MigrateAgentModelResponseSystemDataOutput)
}

// Gets or sets the resource tags.
func (o LookupMigrateAgentResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupMigrateAgentResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Gets or sets the type of the resource.
func (o LookupMigrateAgentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMigrateAgentResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMigrateAgentResultOutput{})
}
