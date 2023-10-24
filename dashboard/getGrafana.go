// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dashboard

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The grafana resource type.
// Azure REST API version: 2022-08-01.
//
// Other available API versions: 2021-09-01-preview, 2022-10-01-preview.
func LookupGrafana(ctx *pulumi.Context, args *LookupGrafanaArgs, opts ...pulumi.InvokeOption) (*LookupGrafanaResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupGrafanaResult
	err := ctx.Invoke("azure-native:dashboard:getGrafana", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGrafanaArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The workspace name of Azure Managed Grafana.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The grafana resource type.
type LookupGrafanaResult struct {
	// ARM id of the grafana resource
	Id string `pulumi:"id"`
	// The managed identity of the grafana resource.
	Identity *ManagedServiceIdentityResponse `pulumi:"identity"`
	// The geo-location where the grafana resource lives
	Location *string `pulumi:"location"`
	// Name of the grafana resource.
	Name string `pulumi:"name"`
	// Properties specific to the grafana resource.
	Properties ManagedGrafanaPropertiesResponse `pulumi:"properties"`
	// The Sku of the grafana resource.
	Sku *ResourceSkuResponse `pulumi:"sku"`
	// The system meta data relating to this grafana resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The tags for grafana resource.
	Tags map[string]string `pulumi:"tags"`
	// The type of the grafana resource.
	Type string `pulumi:"type"`
}

func LookupGrafanaOutput(ctx *pulumi.Context, args LookupGrafanaOutputArgs, opts ...pulumi.InvokeOption) LookupGrafanaResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGrafanaResult, error) {
			args := v.(LookupGrafanaArgs)
			r, err := LookupGrafana(ctx, &args, opts...)
			var s LookupGrafanaResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGrafanaResultOutput)
}

type LookupGrafanaOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The workspace name of Azure Managed Grafana.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupGrafanaOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGrafanaArgs)(nil)).Elem()
}

// The grafana resource type.
type LookupGrafanaResultOutput struct{ *pulumi.OutputState }

func (LookupGrafanaResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGrafanaResult)(nil)).Elem()
}

func (o LookupGrafanaResultOutput) ToLookupGrafanaResultOutput() LookupGrafanaResultOutput {
	return o
}

func (o LookupGrafanaResultOutput) ToLookupGrafanaResultOutputWithContext(ctx context.Context) LookupGrafanaResultOutput {
	return o
}

func (o LookupGrafanaResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupGrafanaResult] {
	return pulumix.Output[LookupGrafanaResult]{
		OutputState: o.OutputState,
	}
}

// ARM id of the grafana resource
func (o LookupGrafanaResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGrafanaResult) string { return v.Id }).(pulumi.StringOutput)
}

// The managed identity of the grafana resource.
func (o LookupGrafanaResultOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupGrafanaResult) *ManagedServiceIdentityResponse { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

// The geo-location where the grafana resource lives
func (o LookupGrafanaResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGrafanaResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// Name of the grafana resource.
func (o LookupGrafanaResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGrafanaResult) string { return v.Name }).(pulumi.StringOutput)
}

// Properties specific to the grafana resource.
func (o LookupGrafanaResultOutput) Properties() ManagedGrafanaPropertiesResponseOutput {
	return o.ApplyT(func(v LookupGrafanaResult) ManagedGrafanaPropertiesResponse { return v.Properties }).(ManagedGrafanaPropertiesResponseOutput)
}

// The Sku of the grafana resource.
func (o LookupGrafanaResultOutput) Sku() ResourceSkuResponsePtrOutput {
	return o.ApplyT(func(v LookupGrafanaResult) *ResourceSkuResponse { return v.Sku }).(ResourceSkuResponsePtrOutput)
}

// The system meta data relating to this grafana resource.
func (o LookupGrafanaResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupGrafanaResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The tags for grafana resource.
func (o LookupGrafanaResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupGrafanaResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the grafana resource.
func (o LookupGrafanaResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGrafanaResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGrafanaResultOutput{})
}
