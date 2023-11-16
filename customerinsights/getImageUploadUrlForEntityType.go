// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package customerinsights

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets entity type (profile or interaction) image upload URL.
// Azure REST API version: 2017-04-26.
func GetImageUploadUrlForEntityType(ctx *pulumi.Context, args *GetImageUploadUrlForEntityTypeArgs, opts ...pulumi.InvokeOption) (*GetImageUploadUrlForEntityTypeResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv GetImageUploadUrlForEntityTypeResult
	err := ctx.Invoke("azure-native:customerinsights:getImageUploadUrlForEntityType", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetImageUploadUrlForEntityTypeArgs struct {
	// Type of entity. Can be Profile or Interaction.
	EntityType *string `pulumi:"entityType"`
	// Name of the entity type.
	EntityTypeName *string `pulumi:"entityTypeName"`
	// The name of the hub.
	HubName string `pulumi:"hubName"`
	// Relative path of the image.
	RelativePath *string `pulumi:"relativePath"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The image definition.
type GetImageUploadUrlForEntityTypeResult struct {
	// Content URL for the image blob.
	ContentUrl *string `pulumi:"contentUrl"`
	// Whether image exists already.
	ImageExists *bool `pulumi:"imageExists"`
	// Relative path of the image.
	RelativePath *string `pulumi:"relativePath"`
}

func GetImageUploadUrlForEntityTypeOutput(ctx *pulumi.Context, args GetImageUploadUrlForEntityTypeOutputArgs, opts ...pulumi.InvokeOption) GetImageUploadUrlForEntityTypeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetImageUploadUrlForEntityTypeResult, error) {
			args := v.(GetImageUploadUrlForEntityTypeArgs)
			r, err := GetImageUploadUrlForEntityType(ctx, &args, opts...)
			var s GetImageUploadUrlForEntityTypeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetImageUploadUrlForEntityTypeResultOutput)
}

type GetImageUploadUrlForEntityTypeOutputArgs struct {
	// Type of entity. Can be Profile or Interaction.
	EntityType pulumi.StringPtrInput `pulumi:"entityType"`
	// Name of the entity type.
	EntityTypeName pulumi.StringPtrInput `pulumi:"entityTypeName"`
	// The name of the hub.
	HubName pulumi.StringInput `pulumi:"hubName"`
	// Relative path of the image.
	RelativePath pulumi.StringPtrInput `pulumi:"relativePath"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (GetImageUploadUrlForEntityTypeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetImageUploadUrlForEntityTypeArgs)(nil)).Elem()
}

// The image definition.
type GetImageUploadUrlForEntityTypeResultOutput struct{ *pulumi.OutputState }

func (GetImageUploadUrlForEntityTypeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetImageUploadUrlForEntityTypeResult)(nil)).Elem()
}

func (o GetImageUploadUrlForEntityTypeResultOutput) ToGetImageUploadUrlForEntityTypeResultOutput() GetImageUploadUrlForEntityTypeResultOutput {
	return o
}

func (o GetImageUploadUrlForEntityTypeResultOutput) ToGetImageUploadUrlForEntityTypeResultOutputWithContext(ctx context.Context) GetImageUploadUrlForEntityTypeResultOutput {
	return o
}

// Content URL for the image blob.
func (o GetImageUploadUrlForEntityTypeResultOutput) ContentUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetImageUploadUrlForEntityTypeResult) *string { return v.ContentUrl }).(pulumi.StringPtrOutput)
}

// Whether image exists already.
func (o GetImageUploadUrlForEntityTypeResultOutput) ImageExists() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetImageUploadUrlForEntityTypeResult) *bool { return v.ImageExists }).(pulumi.BoolPtrOutput)
}

// Relative path of the image.
func (o GetImageUploadUrlForEntityTypeResultOutput) RelativePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetImageUploadUrlForEntityTypeResult) *string { return v.RelativePath }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetImageUploadUrlForEntityTypeResultOutput{})
}
