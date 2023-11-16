// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20140401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns a server communication link.
func LookupServerCommunicationLink(ctx *pulumi.Context, args *LookupServerCommunicationLinkArgs, opts ...pulumi.InvokeOption) (*LookupServerCommunicationLinkResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupServerCommunicationLinkResult
	err := ctx.Invoke("azure-native:sql/v20140401:getServerCommunicationLink", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServerCommunicationLinkArgs struct {
	// The name of the server communication link.
	CommunicationLinkName string `pulumi:"communicationLinkName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
}

// Server communication link.
type LookupServerCommunicationLinkResult struct {
	// Resource ID.
	Id string `pulumi:"id"`
	// Communication link kind.  This property is used for Azure Portal metadata.
	Kind string `pulumi:"kind"`
	// Communication link location.
	Location string `pulumi:"location"`
	// Resource name.
	Name string `pulumi:"name"`
	// The name of the partner server.
	PartnerServer string `pulumi:"partnerServer"`
	// The state.
	State string `pulumi:"state"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupServerCommunicationLinkOutput(ctx *pulumi.Context, args LookupServerCommunicationLinkOutputArgs, opts ...pulumi.InvokeOption) LookupServerCommunicationLinkResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupServerCommunicationLinkResult, error) {
			args := v.(LookupServerCommunicationLinkArgs)
			r, err := LookupServerCommunicationLink(ctx, &args, opts...)
			var s LookupServerCommunicationLinkResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupServerCommunicationLinkResultOutput)
}

type LookupServerCommunicationLinkOutputArgs struct {
	// The name of the server communication link.
	CommunicationLinkName pulumi.StringInput `pulumi:"communicationLinkName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName pulumi.StringInput `pulumi:"serverName"`
}

func (LookupServerCommunicationLinkOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServerCommunicationLinkArgs)(nil)).Elem()
}

// Server communication link.
type LookupServerCommunicationLinkResultOutput struct{ *pulumi.OutputState }

func (LookupServerCommunicationLinkResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServerCommunicationLinkResult)(nil)).Elem()
}

func (o LookupServerCommunicationLinkResultOutput) ToLookupServerCommunicationLinkResultOutput() LookupServerCommunicationLinkResultOutput {
	return o
}

func (o LookupServerCommunicationLinkResultOutput) ToLookupServerCommunicationLinkResultOutputWithContext(ctx context.Context) LookupServerCommunicationLinkResultOutput {
	return o
}

// Resource ID.
func (o LookupServerCommunicationLinkResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerCommunicationLinkResult) string { return v.Id }).(pulumi.StringOutput)
}

// Communication link kind.  This property is used for Azure Portal metadata.
func (o LookupServerCommunicationLinkResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerCommunicationLinkResult) string { return v.Kind }).(pulumi.StringOutput)
}

// Communication link location.
func (o LookupServerCommunicationLinkResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerCommunicationLinkResult) string { return v.Location }).(pulumi.StringOutput)
}

// Resource name.
func (o LookupServerCommunicationLinkResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerCommunicationLinkResult) string { return v.Name }).(pulumi.StringOutput)
}

// The name of the partner server.
func (o LookupServerCommunicationLinkResultOutput) PartnerServer() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerCommunicationLinkResult) string { return v.PartnerServer }).(pulumi.StringOutput)
}

// The state.
func (o LookupServerCommunicationLinkResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerCommunicationLinkResult) string { return v.State }).(pulumi.StringOutput)
}

// Resource type.
func (o LookupServerCommunicationLinkResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerCommunicationLinkResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupServerCommunicationLinkResultOutput{})
}
