// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package web

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Description for Gets the Git/FTP publishing credentials of an app.
// Azure REST API version: 2022-09-01.
//
// Other available API versions: 2016-08-01, 2018-11-01, 2020-10-01.
func ListWebAppPublishingCredentialsSlot(ctx *pulumi.Context, args *ListWebAppPublishingCredentialsSlotArgs, opts ...pulumi.InvokeOption) (*ListWebAppPublishingCredentialsSlotResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListWebAppPublishingCredentialsSlotResult
	err := ctx.Invoke("azure-native:web:listWebAppPublishingCredentialsSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWebAppPublishingCredentialsSlotArgs struct {
	// Name of the app.
	Name string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the deployment slot. If a slot is not specified, the API will get the publishing credentials for the production slot.
	Slot string `pulumi:"slot"`
}

// User credentials used for publishing activity.
type ListWebAppPublishingCredentialsSlotResult struct {
	// Resource Id.
	Id string `pulumi:"id"`
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Resource Name.
	Name string `pulumi:"name"`
	// Password used for publishing.
	PublishingPassword *string `pulumi:"publishingPassword"`
	// Password hash used for publishing.
	PublishingPasswordHash *string `pulumi:"publishingPasswordHash"`
	// Password hash salt used for publishing.
	PublishingPasswordHashSalt *string `pulumi:"publishingPasswordHashSalt"`
	// Username used for publishing.
	PublishingUserName string `pulumi:"publishingUserName"`
	// Url of SCM site.
	ScmUri *string `pulumi:"scmUri"`
	// Resource type.
	Type string `pulumi:"type"`
}

func ListWebAppPublishingCredentialsSlotOutput(ctx *pulumi.Context, args ListWebAppPublishingCredentialsSlotOutputArgs, opts ...pulumi.InvokeOption) ListWebAppPublishingCredentialsSlotResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListWebAppPublishingCredentialsSlotResult, error) {
			args := v.(ListWebAppPublishingCredentialsSlotArgs)
			r, err := ListWebAppPublishingCredentialsSlot(ctx, &args, opts...)
			var s ListWebAppPublishingCredentialsSlotResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListWebAppPublishingCredentialsSlotResultOutput)
}

type ListWebAppPublishingCredentialsSlotOutputArgs struct {
	// Name of the app.
	Name pulumi.StringInput `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name of the deployment slot. If a slot is not specified, the API will get the publishing credentials for the production slot.
	Slot pulumi.StringInput `pulumi:"slot"`
}

func (ListWebAppPublishingCredentialsSlotOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppPublishingCredentialsSlotArgs)(nil)).Elem()
}

// User credentials used for publishing activity.
type ListWebAppPublishingCredentialsSlotResultOutput struct{ *pulumi.OutputState }

func (ListWebAppPublishingCredentialsSlotResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppPublishingCredentialsSlotResult)(nil)).Elem()
}

func (o ListWebAppPublishingCredentialsSlotResultOutput) ToListWebAppPublishingCredentialsSlotResultOutput() ListWebAppPublishingCredentialsSlotResultOutput {
	return o
}

func (o ListWebAppPublishingCredentialsSlotResultOutput) ToListWebAppPublishingCredentialsSlotResultOutputWithContext(ctx context.Context) ListWebAppPublishingCredentialsSlotResultOutput {
	return o
}

func (o ListWebAppPublishingCredentialsSlotResultOutput) ToOutput(ctx context.Context) pulumix.Output[ListWebAppPublishingCredentialsSlotResult] {
	return pulumix.Output[ListWebAppPublishingCredentialsSlotResult]{
		OutputState: o.OutputState,
	}
}

// Resource Id.
func (o ListWebAppPublishingCredentialsSlotResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppPublishingCredentialsSlotResult) string { return v.Id }).(pulumi.StringOutput)
}

// Kind of resource.
func (o ListWebAppPublishingCredentialsSlotResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppPublishingCredentialsSlotResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// Resource Name.
func (o ListWebAppPublishingCredentialsSlotResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppPublishingCredentialsSlotResult) string { return v.Name }).(pulumi.StringOutput)
}

// Password used for publishing.
func (o ListWebAppPublishingCredentialsSlotResultOutput) PublishingPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppPublishingCredentialsSlotResult) *string { return v.PublishingPassword }).(pulumi.StringPtrOutput)
}

// Password hash used for publishing.
func (o ListWebAppPublishingCredentialsSlotResultOutput) PublishingPasswordHash() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppPublishingCredentialsSlotResult) *string { return v.PublishingPasswordHash }).(pulumi.StringPtrOutput)
}

// Password hash salt used for publishing.
func (o ListWebAppPublishingCredentialsSlotResultOutput) PublishingPasswordHashSalt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppPublishingCredentialsSlotResult) *string { return v.PublishingPasswordHashSalt }).(pulumi.StringPtrOutput)
}

// Username used for publishing.
func (o ListWebAppPublishingCredentialsSlotResultOutput) PublishingUserName() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppPublishingCredentialsSlotResult) string { return v.PublishingUserName }).(pulumi.StringOutput)
}

// Url of SCM site.
func (o ListWebAppPublishingCredentialsSlotResultOutput) ScmUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppPublishingCredentialsSlotResult) *string { return v.ScmUri }).(pulumi.StringPtrOutput)
}

// Resource type.
func (o ListWebAppPublishingCredentialsSlotResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppPublishingCredentialsSlotResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListWebAppPublishingCredentialsSlotResultOutput{})
}
