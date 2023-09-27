// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210115

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// List all apps that are assigned to a hostname.
func ListSiteIdentifiersAssignedToHostName(ctx *pulumi.Context, args *ListSiteIdentifiersAssignedToHostNameArgs, opts ...pulumi.InvokeOption) (*ListSiteIdentifiersAssignedToHostNameResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListSiteIdentifiersAssignedToHostNameResult
	err := ctx.Invoke("azure-native:web/v20210115:listSiteIdentifiersAssignedToHostName", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListSiteIdentifiersAssignedToHostNameArgs struct {
	// Name of the object.
	Name *string `pulumi:"name"`
}

// Collection of identifiers.
type ListSiteIdentifiersAssignedToHostNameResult struct {
	// Link to next page of resources.
	NextLink string `pulumi:"nextLink"`
	// Collection of resources.
	Value []IdentifierResponse `pulumi:"value"`
}

func ListSiteIdentifiersAssignedToHostNameOutput(ctx *pulumi.Context, args ListSiteIdentifiersAssignedToHostNameOutputArgs, opts ...pulumi.InvokeOption) ListSiteIdentifiersAssignedToHostNameResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListSiteIdentifiersAssignedToHostNameResult, error) {
			args := v.(ListSiteIdentifiersAssignedToHostNameArgs)
			r, err := ListSiteIdentifiersAssignedToHostName(ctx, &args, opts...)
			var s ListSiteIdentifiersAssignedToHostNameResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListSiteIdentifiersAssignedToHostNameResultOutput)
}

type ListSiteIdentifiersAssignedToHostNameOutputArgs struct {
	// Name of the object.
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (ListSiteIdentifiersAssignedToHostNameOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListSiteIdentifiersAssignedToHostNameArgs)(nil)).Elem()
}

// Collection of identifiers.
type ListSiteIdentifiersAssignedToHostNameResultOutput struct{ *pulumi.OutputState }

func (ListSiteIdentifiersAssignedToHostNameResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListSiteIdentifiersAssignedToHostNameResult)(nil)).Elem()
}

func (o ListSiteIdentifiersAssignedToHostNameResultOutput) ToListSiteIdentifiersAssignedToHostNameResultOutput() ListSiteIdentifiersAssignedToHostNameResultOutput {
	return o
}

func (o ListSiteIdentifiersAssignedToHostNameResultOutput) ToListSiteIdentifiersAssignedToHostNameResultOutputWithContext(ctx context.Context) ListSiteIdentifiersAssignedToHostNameResultOutput {
	return o
}

func (o ListSiteIdentifiersAssignedToHostNameResultOutput) ToOutput(ctx context.Context) pulumix.Output[ListSiteIdentifiersAssignedToHostNameResult] {
	return pulumix.Output[ListSiteIdentifiersAssignedToHostNameResult]{
		OutputState: o.OutputState,
	}
}

// Link to next page of resources.
func (o ListSiteIdentifiersAssignedToHostNameResultOutput) NextLink() pulumi.StringOutput {
	return o.ApplyT(func(v ListSiteIdentifiersAssignedToHostNameResult) string { return v.NextLink }).(pulumi.StringOutput)
}

// Collection of resources.
func (o ListSiteIdentifiersAssignedToHostNameResultOutput) Value() IdentifierResponseArrayOutput {
	return o.ApplyT(func(v ListSiteIdentifiersAssignedToHostNameResult) []IdentifierResponse { return v.Value }).(IdentifierResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListSiteIdentifiersAssignedToHostNameResultOutput{})
}
