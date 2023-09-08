// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package security

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets an Azure API Management API if it has been onboarded to Defender for APIs. If an Azure API Management API is onboarded to Defender for APIs, the system will monitor the operations within the Azure API Management API for intrusive behaviors and provide alerts for attacks that have been detected.
// Azure REST API version: 2022-11-20-preview.
func LookupAPICollection(ctx *pulumi.Context, args *LookupAPICollectionArgs, opts ...pulumi.InvokeOption) (*LookupAPICollectionResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupAPICollectionResult
	err := ctx.Invoke("azure-native:security:getAPICollection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAPICollectionArgs struct {
	// A string representing the apiCollections resource within the Microsoft.Security provider namespace. This string matches the Azure API Management API name.
	ApiCollectionId string `pulumi:"apiCollectionId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
}

// An API collection as represented by Defender for APIs.
type LookupAPICollectionResult struct {
	// Additional data regarding the API collection.
	AdditionalData map[string]string `pulumi:"additionalData"`
	// The display name of the Azure API Management API.
	DisplayName *string `pulumi:"displayName"`
	// Resource Id
	Id string `pulumi:"id"`
	// Resource name
	Name string `pulumi:"name"`
	// Resource type
	Type string `pulumi:"type"`
}

func LookupAPICollectionOutput(ctx *pulumi.Context, args LookupAPICollectionOutputArgs, opts ...pulumi.InvokeOption) LookupAPICollectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAPICollectionResult, error) {
			args := v.(LookupAPICollectionArgs)
			r, err := LookupAPICollection(ctx, &args, opts...)
			var s LookupAPICollectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAPICollectionResultOutput)
}

type LookupAPICollectionOutputArgs struct {
	// A string representing the apiCollections resource within the Microsoft.Security provider namespace. This string matches the Azure API Management API name.
	ApiCollectionId pulumi.StringInput `pulumi:"apiCollectionId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupAPICollectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAPICollectionArgs)(nil)).Elem()
}

// An API collection as represented by Defender for APIs.
type LookupAPICollectionResultOutput struct{ *pulumi.OutputState }

func (LookupAPICollectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAPICollectionResult)(nil)).Elem()
}

func (o LookupAPICollectionResultOutput) ToLookupAPICollectionResultOutput() LookupAPICollectionResultOutput {
	return o
}

func (o LookupAPICollectionResultOutput) ToLookupAPICollectionResultOutputWithContext(ctx context.Context) LookupAPICollectionResultOutput {
	return o
}

// Additional data regarding the API collection.
func (o LookupAPICollectionResultOutput) AdditionalData() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupAPICollectionResult) map[string]string { return v.AdditionalData }).(pulumi.StringMapOutput)
}

// The display name of the Azure API Management API.
func (o LookupAPICollectionResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAPICollectionResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// Resource Id
func (o LookupAPICollectionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAPICollectionResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource name
func (o LookupAPICollectionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAPICollectionResult) string { return v.Name }).(pulumi.StringOutput)
}

// Resource type
func (o LookupAPICollectionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAPICollectionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAPICollectionResultOutput{})
}