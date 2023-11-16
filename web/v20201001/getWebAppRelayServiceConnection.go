// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20201001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a hybrid connection configuration by its name.
func LookupWebAppRelayServiceConnection(ctx *pulumi.Context, args *LookupWebAppRelayServiceConnectionArgs, opts ...pulumi.InvokeOption) (*LookupWebAppRelayServiceConnectionResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupWebAppRelayServiceConnectionResult
	err := ctx.Invoke("azure-native:web/v20201001:getWebAppRelayServiceConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWebAppRelayServiceConnectionArgs struct {
	// Name of the hybrid connection.
	EntityName string `pulumi:"entityName"`
	// Name of the app.
	Name string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Hybrid Connection for an App Service app.
type LookupWebAppRelayServiceConnectionResult struct {
	BiztalkUri             *string `pulumi:"biztalkUri"`
	EntityConnectionString *string `pulumi:"entityConnectionString"`
	EntityName             *string `pulumi:"entityName"`
	Hostname               *string `pulumi:"hostname"`
	// Resource Id.
	Id string `pulumi:"id"`
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Resource Name.
	Name                     string  `pulumi:"name"`
	Port                     *int    `pulumi:"port"`
	ResourceConnectionString *string `pulumi:"resourceConnectionString"`
	ResourceType             *string `pulumi:"resourceType"`
	// The system metadata relating to this resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupWebAppRelayServiceConnectionOutput(ctx *pulumi.Context, args LookupWebAppRelayServiceConnectionOutputArgs, opts ...pulumi.InvokeOption) LookupWebAppRelayServiceConnectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWebAppRelayServiceConnectionResult, error) {
			args := v.(LookupWebAppRelayServiceConnectionArgs)
			r, err := LookupWebAppRelayServiceConnection(ctx, &args, opts...)
			var s LookupWebAppRelayServiceConnectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWebAppRelayServiceConnectionResultOutput)
}

type LookupWebAppRelayServiceConnectionOutputArgs struct {
	// Name of the hybrid connection.
	EntityName pulumi.StringInput `pulumi:"entityName"`
	// Name of the app.
	Name pulumi.StringInput `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupWebAppRelayServiceConnectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppRelayServiceConnectionArgs)(nil)).Elem()
}

// Hybrid Connection for an App Service app.
type LookupWebAppRelayServiceConnectionResultOutput struct{ *pulumi.OutputState }

func (LookupWebAppRelayServiceConnectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppRelayServiceConnectionResult)(nil)).Elem()
}

func (o LookupWebAppRelayServiceConnectionResultOutput) ToLookupWebAppRelayServiceConnectionResultOutput() LookupWebAppRelayServiceConnectionResultOutput {
	return o
}

func (o LookupWebAppRelayServiceConnectionResultOutput) ToLookupWebAppRelayServiceConnectionResultOutputWithContext(ctx context.Context) LookupWebAppRelayServiceConnectionResultOutput {
	return o
}

func (o LookupWebAppRelayServiceConnectionResultOutput) BiztalkUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppRelayServiceConnectionResult) *string { return v.BiztalkUri }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppRelayServiceConnectionResultOutput) EntityConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppRelayServiceConnectionResult) *string { return v.EntityConnectionString }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppRelayServiceConnectionResultOutput) EntityName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppRelayServiceConnectionResult) *string { return v.EntityName }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppRelayServiceConnectionResultOutput) Hostname() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppRelayServiceConnectionResult) *string { return v.Hostname }).(pulumi.StringPtrOutput)
}

// Resource Id.
func (o LookupWebAppRelayServiceConnectionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppRelayServiceConnectionResult) string { return v.Id }).(pulumi.StringOutput)
}

// Kind of resource.
func (o LookupWebAppRelayServiceConnectionResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppRelayServiceConnectionResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// Resource Name.
func (o LookupWebAppRelayServiceConnectionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppRelayServiceConnectionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupWebAppRelayServiceConnectionResultOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupWebAppRelayServiceConnectionResult) *int { return v.Port }).(pulumi.IntPtrOutput)
}

func (o LookupWebAppRelayServiceConnectionResultOutput) ResourceConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppRelayServiceConnectionResult) *string { return v.ResourceConnectionString }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppRelayServiceConnectionResultOutput) ResourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppRelayServiceConnectionResult) *string { return v.ResourceType }).(pulumi.StringPtrOutput)
}

// The system metadata relating to this resource.
func (o LookupWebAppRelayServiceConnectionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupWebAppRelayServiceConnectionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource type.
func (o LookupWebAppRelayServiceConnectionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppRelayServiceConnectionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWebAppRelayServiceConnectionResultOutput{})
}
