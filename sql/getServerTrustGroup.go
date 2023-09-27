// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sql

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets a server trust group.
// Azure REST API version: 2021-11-01.
func LookupServerTrustGroup(ctx *pulumi.Context, args *LookupServerTrustGroupArgs, opts ...pulumi.InvokeOption) (*LookupServerTrustGroupResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupServerTrustGroupResult
	err := ctx.Invoke("azure-native:sql:getServerTrustGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServerTrustGroupArgs struct {
	// The name of the region where the resource is located.
	LocationName string `pulumi:"locationName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the server trust group.
	ServerTrustGroupName string `pulumi:"serverTrustGroupName"`
}

// A server trust group.
type LookupServerTrustGroupResult struct {
	// Group members information for the server trust group.
	GroupMembers []ServerInfoResponse `pulumi:"groupMembers"`
	// Resource ID.
	Id string `pulumi:"id"`
	// Resource name.
	Name string `pulumi:"name"`
	// Trust scope of the server trust group.
	TrustScopes []string `pulumi:"trustScopes"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupServerTrustGroupOutput(ctx *pulumi.Context, args LookupServerTrustGroupOutputArgs, opts ...pulumi.InvokeOption) LookupServerTrustGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupServerTrustGroupResult, error) {
			args := v.(LookupServerTrustGroupArgs)
			r, err := LookupServerTrustGroup(ctx, &args, opts...)
			var s LookupServerTrustGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupServerTrustGroupResultOutput)
}

type LookupServerTrustGroupOutputArgs struct {
	// The name of the region where the resource is located.
	LocationName pulumi.StringInput `pulumi:"locationName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the server trust group.
	ServerTrustGroupName pulumi.StringInput `pulumi:"serverTrustGroupName"`
}

func (LookupServerTrustGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServerTrustGroupArgs)(nil)).Elem()
}

// A server trust group.
type LookupServerTrustGroupResultOutput struct{ *pulumi.OutputState }

func (LookupServerTrustGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServerTrustGroupResult)(nil)).Elem()
}

func (o LookupServerTrustGroupResultOutput) ToLookupServerTrustGroupResultOutput() LookupServerTrustGroupResultOutput {
	return o
}

func (o LookupServerTrustGroupResultOutput) ToLookupServerTrustGroupResultOutputWithContext(ctx context.Context) LookupServerTrustGroupResultOutput {
	return o
}

func (o LookupServerTrustGroupResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupServerTrustGroupResult] {
	return pulumix.Output[LookupServerTrustGroupResult]{
		OutputState: o.OutputState,
	}
}

// Group members information for the server trust group.
func (o LookupServerTrustGroupResultOutput) GroupMembers() ServerInfoResponseArrayOutput {
	return o.ApplyT(func(v LookupServerTrustGroupResult) []ServerInfoResponse { return v.GroupMembers }).(ServerInfoResponseArrayOutput)
}

// Resource ID.
func (o LookupServerTrustGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerTrustGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource name.
func (o LookupServerTrustGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerTrustGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

// Trust scope of the server trust group.
func (o LookupServerTrustGroupResultOutput) TrustScopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupServerTrustGroupResult) []string { return v.TrustScopes }).(pulumi.StringArrayOutput)
}

// Resource type.
func (o LookupServerTrustGroupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerTrustGroupResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupServerTrustGroupResultOutput{})
}
