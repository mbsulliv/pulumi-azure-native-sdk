// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220330preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The files resource is a nested, proxy-only resource representing a file stored under the project resource. This method retrieves information about a file.
func LookupFile(ctx *pulumi.Context, args *LookupFileArgs, opts ...pulumi.InvokeOption) (*LookupFileResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupFileResult
	err := ctx.Invoke("azure-native:datamigration/v20220330preview:getFile", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFileArgs struct {
	// Name of the File
	FileName string `pulumi:"fileName"`
	// Name of the resource group
	GroupName string `pulumi:"groupName"`
	// Name of the project
	ProjectName string `pulumi:"projectName"`
	// Name of the service
	ServiceName string `pulumi:"serviceName"`
}

// A file resource
type LookupFileResult struct {
	// HTTP strong entity tag value. This is ignored if submitted.
	Etag *string `pulumi:"etag"`
	// Resource ID.
	Id string `pulumi:"id"`
	// Resource name.
	Name string `pulumi:"name"`
	// Custom file properties
	Properties ProjectFilePropertiesResponse `pulumi:"properties"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupFileOutput(ctx *pulumi.Context, args LookupFileOutputArgs, opts ...pulumi.InvokeOption) LookupFileResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFileResult, error) {
			args := v.(LookupFileArgs)
			r, err := LookupFile(ctx, &args, opts...)
			var s LookupFileResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFileResultOutput)
}

type LookupFileOutputArgs struct {
	// Name of the File
	FileName pulumi.StringInput `pulumi:"fileName"`
	// Name of the resource group
	GroupName pulumi.StringInput `pulumi:"groupName"`
	// Name of the project
	ProjectName pulumi.StringInput `pulumi:"projectName"`
	// Name of the service
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupFileOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFileArgs)(nil)).Elem()
}

// A file resource
type LookupFileResultOutput struct{ *pulumi.OutputState }

func (LookupFileResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFileResult)(nil)).Elem()
}

func (o LookupFileResultOutput) ToLookupFileResultOutput() LookupFileResultOutput {
	return o
}

func (o LookupFileResultOutput) ToLookupFileResultOutputWithContext(ctx context.Context) LookupFileResultOutput {
	return o
}

// HTTP strong entity tag value. This is ignored if submitted.
func (o LookupFileResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFileResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// Resource ID.
func (o LookupFileResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFileResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource name.
func (o LookupFileResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFileResult) string { return v.Name }).(pulumi.StringOutput)
}

// Custom file properties
func (o LookupFileResultOutput) Properties() ProjectFilePropertiesResponseOutput {
	return o.ApplyT(func(v LookupFileResult) ProjectFilePropertiesResponse { return v.Properties }).(ProjectFilePropertiesResponseOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupFileResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupFileResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource type.
func (o LookupFileResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFileResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFileResultOutput{})
}
