// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get a Service Fabric application type version resource created or in the process of being created in the Service Fabric application type name resource.
func LookupApplicationTypeVersion(ctx *pulumi.Context, args *LookupApplicationTypeVersionArgs, opts ...pulumi.InvokeOption) (*LookupApplicationTypeVersionResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupApplicationTypeVersionResult
	err := ctx.Invoke("azure-native:servicefabric/v20210601:getApplicationTypeVersion", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApplicationTypeVersionArgs struct {
	// The name of the application type name resource.
	ApplicationTypeName string `pulumi:"applicationTypeName"`
	// The name of the cluster resource.
	ClusterName string `pulumi:"clusterName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The application type version.
	Version string `pulumi:"version"`
}

// An application type version resource for the specified application type name resource.
type LookupApplicationTypeVersionResult struct {
	// The URL to the application package
	AppPackageUrl string `pulumi:"appPackageUrl"`
	// List of application type parameters that can be overridden when creating or updating the application.
	DefaultParameterList map[string]string `pulumi:"defaultParameterList"`
	// Azure resource etag.
	Etag string `pulumi:"etag"`
	// Azure resource identifier.
	Id string `pulumi:"id"`
	// It will be deprecated in New API, resource location depends on the parent resource.
	Location *string `pulumi:"location"`
	// Azure resource name.
	Name string `pulumi:"name"`
	// The current deployment or provisioning state, which only appears in the response
	ProvisioningState string `pulumi:"provisioningState"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Azure resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Azure resource type.
	Type string `pulumi:"type"`
}

func LookupApplicationTypeVersionOutput(ctx *pulumi.Context, args LookupApplicationTypeVersionOutputArgs, opts ...pulumi.InvokeOption) LookupApplicationTypeVersionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupApplicationTypeVersionResult, error) {
			args := v.(LookupApplicationTypeVersionArgs)
			r, err := LookupApplicationTypeVersion(ctx, &args, opts...)
			var s LookupApplicationTypeVersionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupApplicationTypeVersionResultOutput)
}

type LookupApplicationTypeVersionOutputArgs struct {
	// The name of the application type name resource.
	ApplicationTypeName pulumi.StringInput `pulumi:"applicationTypeName"`
	// The name of the cluster resource.
	ClusterName pulumi.StringInput `pulumi:"clusterName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The application type version.
	Version pulumi.StringInput `pulumi:"version"`
}

func (LookupApplicationTypeVersionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplicationTypeVersionArgs)(nil)).Elem()
}

// An application type version resource for the specified application type name resource.
type LookupApplicationTypeVersionResultOutput struct{ *pulumi.OutputState }

func (LookupApplicationTypeVersionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplicationTypeVersionResult)(nil)).Elem()
}

func (o LookupApplicationTypeVersionResultOutput) ToLookupApplicationTypeVersionResultOutput() LookupApplicationTypeVersionResultOutput {
	return o
}

func (o LookupApplicationTypeVersionResultOutput) ToLookupApplicationTypeVersionResultOutputWithContext(ctx context.Context) LookupApplicationTypeVersionResultOutput {
	return o
}

func (o LookupApplicationTypeVersionResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupApplicationTypeVersionResult] {
	return pulumix.Output[LookupApplicationTypeVersionResult]{
		OutputState: o.OutputState,
	}
}

// The URL to the application package
func (o LookupApplicationTypeVersionResultOutput) AppPackageUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationTypeVersionResult) string { return v.AppPackageUrl }).(pulumi.StringOutput)
}

// List of application type parameters that can be overridden when creating or updating the application.
func (o LookupApplicationTypeVersionResultOutput) DefaultParameterList() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupApplicationTypeVersionResult) map[string]string { return v.DefaultParameterList }).(pulumi.StringMapOutput)
}

// Azure resource etag.
func (o LookupApplicationTypeVersionResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationTypeVersionResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Azure resource identifier.
func (o LookupApplicationTypeVersionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationTypeVersionResult) string { return v.Id }).(pulumi.StringOutput)
}

// It will be deprecated in New API, resource location depends on the parent resource.
func (o LookupApplicationTypeVersionResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationTypeVersionResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// Azure resource name.
func (o LookupApplicationTypeVersionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationTypeVersionResult) string { return v.Name }).(pulumi.StringOutput)
}

// The current deployment or provisioning state, which only appears in the response
func (o LookupApplicationTypeVersionResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationTypeVersionResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupApplicationTypeVersionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupApplicationTypeVersionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Azure resource tags.
func (o LookupApplicationTypeVersionResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupApplicationTypeVersionResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Azure resource type.
func (o LookupApplicationTypeVersionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationTypeVersionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApplicationTypeVersionResultOutput{})
}
