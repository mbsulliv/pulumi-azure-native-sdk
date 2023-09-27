// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20170426

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets data image upload URL.
func GetImageUploadUrlForData(ctx *pulumi.Context, args *GetImageUploadUrlForDataArgs, opts ...pulumi.InvokeOption) (*GetImageUploadUrlForDataResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv GetImageUploadUrlForDataResult
	err := ctx.Invoke("azure-native:customerinsights/v20170426:getImageUploadUrlForData", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetImageUploadUrlForDataArgs struct {
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
type GetImageUploadUrlForDataResult struct {
	// Content URL for the image blob.
	ContentUrl *string `pulumi:"contentUrl"`
	// Whether image exists already.
	ImageExists *bool `pulumi:"imageExists"`
	// Relative path of the image.
	RelativePath *string `pulumi:"relativePath"`
}

func GetImageUploadUrlForDataOutput(ctx *pulumi.Context, args GetImageUploadUrlForDataOutputArgs, opts ...pulumi.InvokeOption) GetImageUploadUrlForDataResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetImageUploadUrlForDataResult, error) {
			args := v.(GetImageUploadUrlForDataArgs)
			r, err := GetImageUploadUrlForData(ctx, &args, opts...)
			var s GetImageUploadUrlForDataResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetImageUploadUrlForDataResultOutput)
}

type GetImageUploadUrlForDataOutputArgs struct {
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

func (GetImageUploadUrlForDataOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetImageUploadUrlForDataArgs)(nil)).Elem()
}

// The image definition.
type GetImageUploadUrlForDataResultOutput struct{ *pulumi.OutputState }

func (GetImageUploadUrlForDataResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetImageUploadUrlForDataResult)(nil)).Elem()
}

func (o GetImageUploadUrlForDataResultOutput) ToGetImageUploadUrlForDataResultOutput() GetImageUploadUrlForDataResultOutput {
	return o
}

func (o GetImageUploadUrlForDataResultOutput) ToGetImageUploadUrlForDataResultOutputWithContext(ctx context.Context) GetImageUploadUrlForDataResultOutput {
	return o
}

func (o GetImageUploadUrlForDataResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetImageUploadUrlForDataResult] {
	return pulumix.Output[GetImageUploadUrlForDataResult]{
		OutputState: o.OutputState,
	}
}

// Content URL for the image blob.
func (o GetImageUploadUrlForDataResultOutput) ContentUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetImageUploadUrlForDataResult) *string { return v.ContentUrl }).(pulumi.StringPtrOutput)
}

// Whether image exists already.
func (o GetImageUploadUrlForDataResultOutput) ImageExists() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetImageUploadUrlForDataResult) *bool { return v.ImageExists }).(pulumi.BoolPtrOutput)
}

// Relative path of the image.
func (o GetImageUploadUrlForDataResultOutput) RelativePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetImageUploadUrlForDataResult) *string { return v.RelativePath }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetImageUploadUrlForDataResultOutput{})
}
