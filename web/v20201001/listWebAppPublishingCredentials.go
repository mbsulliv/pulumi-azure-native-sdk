// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20201001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the Git/FTP publishing credentials of an app.
func ListWebAppPublishingCredentials(ctx *pulumi.Context, args *ListWebAppPublishingCredentialsArgs, opts ...pulumi.InvokeOption) (*ListWebAppPublishingCredentialsResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListWebAppPublishingCredentialsResult
	err := ctx.Invoke("azure-native:web/v20201001:listWebAppPublishingCredentials", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWebAppPublishingCredentialsArgs struct {
	// Name of the app.
	Name string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// User credentials used for publishing activity.
type ListWebAppPublishingCredentialsResult struct {
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
	// The system metadata relating to this resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource type.
	Type string `pulumi:"type"`
}

func ListWebAppPublishingCredentialsOutput(ctx *pulumi.Context, args ListWebAppPublishingCredentialsOutputArgs, opts ...pulumi.InvokeOption) ListWebAppPublishingCredentialsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListWebAppPublishingCredentialsResult, error) {
			args := v.(ListWebAppPublishingCredentialsArgs)
			r, err := ListWebAppPublishingCredentials(ctx, &args, opts...)
			var s ListWebAppPublishingCredentialsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListWebAppPublishingCredentialsResultOutput)
}

type ListWebAppPublishingCredentialsOutputArgs struct {
	// Name of the app.
	Name pulumi.StringInput `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListWebAppPublishingCredentialsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppPublishingCredentialsArgs)(nil)).Elem()
}

// User credentials used for publishing activity.
type ListWebAppPublishingCredentialsResultOutput struct{ *pulumi.OutputState }

func (ListWebAppPublishingCredentialsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppPublishingCredentialsResult)(nil)).Elem()
}

func (o ListWebAppPublishingCredentialsResultOutput) ToListWebAppPublishingCredentialsResultOutput() ListWebAppPublishingCredentialsResultOutput {
	return o
}

func (o ListWebAppPublishingCredentialsResultOutput) ToListWebAppPublishingCredentialsResultOutputWithContext(ctx context.Context) ListWebAppPublishingCredentialsResultOutput {
	return o
}

// Resource Id.
func (o ListWebAppPublishingCredentialsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppPublishingCredentialsResult) string { return v.Id }).(pulumi.StringOutput)
}

// Kind of resource.
func (o ListWebAppPublishingCredentialsResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppPublishingCredentialsResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// Resource Name.
func (o ListWebAppPublishingCredentialsResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppPublishingCredentialsResult) string { return v.Name }).(pulumi.StringOutput)
}

// Password used for publishing.
func (o ListWebAppPublishingCredentialsResultOutput) PublishingPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppPublishingCredentialsResult) *string { return v.PublishingPassword }).(pulumi.StringPtrOutput)
}

// Password hash used for publishing.
func (o ListWebAppPublishingCredentialsResultOutput) PublishingPasswordHash() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppPublishingCredentialsResult) *string { return v.PublishingPasswordHash }).(pulumi.StringPtrOutput)
}

// Password hash salt used for publishing.
func (o ListWebAppPublishingCredentialsResultOutput) PublishingPasswordHashSalt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppPublishingCredentialsResult) *string { return v.PublishingPasswordHashSalt }).(pulumi.StringPtrOutput)
}

// Username used for publishing.
func (o ListWebAppPublishingCredentialsResultOutput) PublishingUserName() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppPublishingCredentialsResult) string { return v.PublishingUserName }).(pulumi.StringOutput)
}

// Url of SCM site.
func (o ListWebAppPublishingCredentialsResultOutput) ScmUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppPublishingCredentialsResult) *string { return v.ScmUri }).(pulumi.StringPtrOutput)
}

// The system metadata relating to this resource.
func (o ListWebAppPublishingCredentialsResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v ListWebAppPublishingCredentialsResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource type.
func (o ListWebAppPublishingCredentialsResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppPublishingCredentialsResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListWebAppPublishingCredentialsResultOutput{})
}
