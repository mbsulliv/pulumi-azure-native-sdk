// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package migrate

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets the details of the modernize project.
// Azure REST API version: 2022-05-01-preview.
func LookupModernizeProject(ctx *pulumi.Context, args *LookupModernizeProjectArgs, opts ...pulumi.InvokeOption) (*LookupModernizeProjectResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupModernizeProjectResult
	err := ctx.Invoke("azure-native:migrate:getModernizeProject", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupModernizeProjectArgs struct {
	// Modernize project name.
	ModernizeProjectName string `pulumi:"modernizeProjectName"`
	// Name of the Azure Resource Group that project is part of.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Azure Subscription Id in which project was created.
	SubscriptionId *string `pulumi:"subscriptionId"`
}

// ModernizeProject model.
type LookupModernizeProjectResult struct {
	// Gets or sets the Id of the resource.
	Id       string                    `pulumi:"id"`
	Identity *ResourceIdentityResponse `pulumi:"identity"`
	// Gets or sets the location of the modernizeProject.
	Location *string `pulumi:"location"`
	// Gets or sets the name of the resource.
	Name string `pulumi:"name"`
	// ModernizeProject properties.
	Properties ModernizeProjectModelPropertiesResponse `pulumi:"properties"`
	SystemData ModernizeProjectModelResponseSystemData `pulumi:"systemData"`
	// Gets or sets the resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Gets or sets the type of the resource.
	Type string `pulumi:"type"`
}

func LookupModernizeProjectOutput(ctx *pulumi.Context, args LookupModernizeProjectOutputArgs, opts ...pulumi.InvokeOption) LookupModernizeProjectResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupModernizeProjectResult, error) {
			args := v.(LookupModernizeProjectArgs)
			r, err := LookupModernizeProject(ctx, &args, opts...)
			var s LookupModernizeProjectResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupModernizeProjectResultOutput)
}

type LookupModernizeProjectOutputArgs struct {
	// Modernize project name.
	ModernizeProjectName pulumi.StringInput `pulumi:"modernizeProjectName"`
	// Name of the Azure Resource Group that project is part of.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Azure Subscription Id in which project was created.
	SubscriptionId pulumi.StringPtrInput `pulumi:"subscriptionId"`
}

func (LookupModernizeProjectOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupModernizeProjectArgs)(nil)).Elem()
}

// ModernizeProject model.
type LookupModernizeProjectResultOutput struct{ *pulumi.OutputState }

func (LookupModernizeProjectResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupModernizeProjectResult)(nil)).Elem()
}

func (o LookupModernizeProjectResultOutput) ToLookupModernizeProjectResultOutput() LookupModernizeProjectResultOutput {
	return o
}

func (o LookupModernizeProjectResultOutput) ToLookupModernizeProjectResultOutputWithContext(ctx context.Context) LookupModernizeProjectResultOutput {
	return o
}

func (o LookupModernizeProjectResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupModernizeProjectResult] {
	return pulumix.Output[LookupModernizeProjectResult]{
		OutputState: o.OutputState,
	}
}

// Gets or sets the Id of the resource.
func (o LookupModernizeProjectResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupModernizeProjectResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupModernizeProjectResultOutput) Identity() ResourceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupModernizeProjectResult) *ResourceIdentityResponse { return v.Identity }).(ResourceIdentityResponsePtrOutput)
}

// Gets or sets the location of the modernizeProject.
func (o LookupModernizeProjectResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupModernizeProjectResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// Gets or sets the name of the resource.
func (o LookupModernizeProjectResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupModernizeProjectResult) string { return v.Name }).(pulumi.StringOutput)
}

// ModernizeProject properties.
func (o LookupModernizeProjectResultOutput) Properties() ModernizeProjectModelPropertiesResponseOutput {
	return o.ApplyT(func(v LookupModernizeProjectResult) ModernizeProjectModelPropertiesResponse { return v.Properties }).(ModernizeProjectModelPropertiesResponseOutput)
}

func (o LookupModernizeProjectResultOutput) SystemData() ModernizeProjectModelResponseSystemDataOutput {
	return o.ApplyT(func(v LookupModernizeProjectResult) ModernizeProjectModelResponseSystemData { return v.SystemData }).(ModernizeProjectModelResponseSystemDataOutput)
}

// Gets or sets the resource tags.
func (o LookupModernizeProjectResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupModernizeProjectResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Gets or sets the type of the resource.
func (o LookupModernizeProjectResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupModernizeProjectResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupModernizeProjectResultOutput{})
}
