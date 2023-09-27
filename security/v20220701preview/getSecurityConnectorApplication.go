// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220701preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get a specific application for the requested scope by applicationId
func LookupSecurityConnectorApplication(ctx *pulumi.Context, args *LookupSecurityConnectorApplicationArgs, opts ...pulumi.InvokeOption) (*LookupSecurityConnectorApplicationResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupSecurityConnectorApplicationResult
	err := ctx.Invoke("azure-native:security/v20220701preview:getSecurityConnectorApplication", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSecurityConnectorApplicationArgs struct {
	// The security Application key - unique key for the standard application
	ApplicationId string `pulumi:"applicationId"`
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The security connector name.
	SecurityConnectorName string `pulumi:"securityConnectorName"`
}

// Security Application over a given scope
type LookupSecurityConnectorApplicationResult struct {
	// description of the application
	Description *string `pulumi:"description"`
	// display name of the application
	DisplayName *string `pulumi:"displayName"`
	// Resource Id
	Id string `pulumi:"id"`
	// Resource name
	Name string `pulumi:"name"`
	// The application source, what it affects, e.g. Assessments
	SourceResourceType string `pulumi:"sourceResourceType"`
	// Resource type
	Type string `pulumi:"type"`
}

func LookupSecurityConnectorApplicationOutput(ctx *pulumi.Context, args LookupSecurityConnectorApplicationOutputArgs, opts ...pulumi.InvokeOption) LookupSecurityConnectorApplicationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSecurityConnectorApplicationResult, error) {
			args := v.(LookupSecurityConnectorApplicationArgs)
			r, err := LookupSecurityConnectorApplication(ctx, &args, opts...)
			var s LookupSecurityConnectorApplicationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSecurityConnectorApplicationResultOutput)
}

type LookupSecurityConnectorApplicationOutputArgs struct {
	// The security Application key - unique key for the standard application
	ApplicationId pulumi.StringInput `pulumi:"applicationId"`
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The security connector name.
	SecurityConnectorName pulumi.StringInput `pulumi:"securityConnectorName"`
}

func (LookupSecurityConnectorApplicationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSecurityConnectorApplicationArgs)(nil)).Elem()
}

// Security Application over a given scope
type LookupSecurityConnectorApplicationResultOutput struct{ *pulumi.OutputState }

func (LookupSecurityConnectorApplicationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSecurityConnectorApplicationResult)(nil)).Elem()
}

func (o LookupSecurityConnectorApplicationResultOutput) ToLookupSecurityConnectorApplicationResultOutput() LookupSecurityConnectorApplicationResultOutput {
	return o
}

func (o LookupSecurityConnectorApplicationResultOutput) ToLookupSecurityConnectorApplicationResultOutputWithContext(ctx context.Context) LookupSecurityConnectorApplicationResultOutput {
	return o
}

func (o LookupSecurityConnectorApplicationResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupSecurityConnectorApplicationResult] {
	return pulumix.Output[LookupSecurityConnectorApplicationResult]{
		OutputState: o.OutputState,
	}
}

// description of the application
func (o LookupSecurityConnectorApplicationResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSecurityConnectorApplicationResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// display name of the application
func (o LookupSecurityConnectorApplicationResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSecurityConnectorApplicationResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// Resource Id
func (o LookupSecurityConnectorApplicationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecurityConnectorApplicationResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource name
func (o LookupSecurityConnectorApplicationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecurityConnectorApplicationResult) string { return v.Name }).(pulumi.StringOutput)
}

// The application source, what it affects, e.g. Assessments
func (o LookupSecurityConnectorApplicationResultOutput) SourceResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecurityConnectorApplicationResult) string { return v.SourceResourceType }).(pulumi.StringOutput)
}

// Resource type
func (o LookupSecurityConnectorApplicationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecurityConnectorApplicationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSecurityConnectorApplicationResultOutput{})
}
