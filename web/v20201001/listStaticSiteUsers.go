// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20201001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets the list of users of a static site.
func ListStaticSiteUsers(ctx *pulumi.Context, args *ListStaticSiteUsersArgs, opts ...pulumi.InvokeOption) (*ListStaticSiteUsersResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListStaticSiteUsersResult
	err := ctx.Invoke("azure-native:web/v20201001:listStaticSiteUsers", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListStaticSiteUsersArgs struct {
	// The auth provider for the users.
	Authprovider string `pulumi:"authprovider"`
	// Name of the static site.
	Name string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Collection of static site custom users.
type ListStaticSiteUsersResult struct {
	// Link to next page of resources.
	NextLink string `pulumi:"nextLink"`
	// Collection of resources.
	Value []StaticSiteUserARMResourceResponse `pulumi:"value"`
}

func ListStaticSiteUsersOutput(ctx *pulumi.Context, args ListStaticSiteUsersOutputArgs, opts ...pulumi.InvokeOption) ListStaticSiteUsersResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListStaticSiteUsersResult, error) {
			args := v.(ListStaticSiteUsersArgs)
			r, err := ListStaticSiteUsers(ctx, &args, opts...)
			var s ListStaticSiteUsersResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListStaticSiteUsersResultOutput)
}

type ListStaticSiteUsersOutputArgs struct {
	// The auth provider for the users.
	Authprovider pulumi.StringInput `pulumi:"authprovider"`
	// Name of the static site.
	Name pulumi.StringInput `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListStaticSiteUsersOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListStaticSiteUsersArgs)(nil)).Elem()
}

// Collection of static site custom users.
type ListStaticSiteUsersResultOutput struct{ *pulumi.OutputState }

func (ListStaticSiteUsersResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListStaticSiteUsersResult)(nil)).Elem()
}

func (o ListStaticSiteUsersResultOutput) ToListStaticSiteUsersResultOutput() ListStaticSiteUsersResultOutput {
	return o
}

func (o ListStaticSiteUsersResultOutput) ToListStaticSiteUsersResultOutputWithContext(ctx context.Context) ListStaticSiteUsersResultOutput {
	return o
}

func (o ListStaticSiteUsersResultOutput) ToOutput(ctx context.Context) pulumix.Output[ListStaticSiteUsersResult] {
	return pulumix.Output[ListStaticSiteUsersResult]{
		OutputState: o.OutputState,
	}
}

// Link to next page of resources.
func (o ListStaticSiteUsersResultOutput) NextLink() pulumi.StringOutput {
	return o.ApplyT(func(v ListStaticSiteUsersResult) string { return v.NextLink }).(pulumi.StringOutput)
}

// Collection of resources.
func (o ListStaticSiteUsersResultOutput) Value() StaticSiteUserARMResourceResponseArrayOutput {
	return o.ApplyT(func(v ListStaticSiteUsersResult) []StaticSiteUserARMResourceResponse { return v.Value }).(StaticSiteUserARMResourceResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListStaticSiteUsersResultOutput{})
}
