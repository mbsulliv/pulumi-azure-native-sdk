// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package confluent

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// List cluster success response
// Azure REST API version: 2023-08-22.
//
// Other available API versions: 2024-02-13.
func ListAccessClusters(ctx *pulumi.Context, args *ListAccessClustersArgs, opts ...pulumi.InvokeOption) (*ListAccessClustersResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListAccessClustersResult
	err := ctx.Invoke("azure-native:confluent:listAccessClusters", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListAccessClustersArgs struct {
	// Organization resource name
	OrganizationName string `pulumi:"organizationName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Search filters for the request
	SearchFilters map[string]string `pulumi:"searchFilters"`
}

// List cluster success response
type ListAccessClustersResult struct {
	// Data of the environments list
	Data []ClusterRecordResponse `pulumi:"data"`
	// Type of response
	Kind *string `pulumi:"kind"`
	// Metadata of the list
	Metadata *ConfluentListMetadataResponse `pulumi:"metadata"`
}

func ListAccessClustersOutput(ctx *pulumi.Context, args ListAccessClustersOutputArgs, opts ...pulumi.InvokeOption) ListAccessClustersResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListAccessClustersResult, error) {
			args := v.(ListAccessClustersArgs)
			r, err := ListAccessClusters(ctx, &args, opts...)
			var s ListAccessClustersResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListAccessClustersResultOutput)
}

type ListAccessClustersOutputArgs struct {
	// Organization resource name
	OrganizationName pulumi.StringInput `pulumi:"organizationName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Search filters for the request
	SearchFilters pulumi.StringMapInput `pulumi:"searchFilters"`
}

func (ListAccessClustersOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListAccessClustersArgs)(nil)).Elem()
}

// List cluster success response
type ListAccessClustersResultOutput struct{ *pulumi.OutputState }

func (ListAccessClustersResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListAccessClustersResult)(nil)).Elem()
}

func (o ListAccessClustersResultOutput) ToListAccessClustersResultOutput() ListAccessClustersResultOutput {
	return o
}

func (o ListAccessClustersResultOutput) ToListAccessClustersResultOutputWithContext(ctx context.Context) ListAccessClustersResultOutput {
	return o
}

// Data of the environments list
func (o ListAccessClustersResultOutput) Data() ClusterRecordResponseArrayOutput {
	return o.ApplyT(func(v ListAccessClustersResult) []ClusterRecordResponse { return v.Data }).(ClusterRecordResponseArrayOutput)
}

// Type of response
func (o ListAccessClustersResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListAccessClustersResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// Metadata of the list
func (o ListAccessClustersResultOutput) Metadata() ConfluentListMetadataResponsePtrOutput {
	return o.ApplyT(func(v ListAccessClustersResult) *ConfluentListMetadataResponse { return v.Metadata }).(ConfluentListMetadataResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListAccessClustersResultOutput{})
}
