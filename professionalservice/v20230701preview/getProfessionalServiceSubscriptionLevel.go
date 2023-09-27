// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230701preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets information about the specified Subscription Level ProfessionalService.
func LookupProfessionalServiceSubscriptionLevel(ctx *pulumi.Context, args *LookupProfessionalServiceSubscriptionLevelArgs, opts ...pulumi.InvokeOption) (*LookupProfessionalServiceSubscriptionLevelResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupProfessionalServiceSubscriptionLevelResult
	err := ctx.Invoke("azure-native:professionalservice/v20230701preview:getProfessionalServiceSubscriptionLevel", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupProfessionalServiceSubscriptionLevelArgs struct {
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the resource.
	ResourceName string `pulumi:"resourceName"`
	// The Azure subscription ID. This is a GUID-formatted string (e.g. 00000000-0000-0000-0000-000000000000)
	SubscriptionId *string `pulumi:"subscriptionId"`
}

// ProfessionalService REST API resource definition.
type LookupProfessionalServiceSubscriptionLevelResult struct {
	// The resource uri
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// professionalService properties
	Properties ProfessionalServiceResourceResponseProperties `pulumi:"properties"`
	// the resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupProfessionalServiceSubscriptionLevelOutput(ctx *pulumi.Context, args LookupProfessionalServiceSubscriptionLevelOutputArgs, opts ...pulumi.InvokeOption) LookupProfessionalServiceSubscriptionLevelResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupProfessionalServiceSubscriptionLevelResult, error) {
			args := v.(LookupProfessionalServiceSubscriptionLevelArgs)
			r, err := LookupProfessionalServiceSubscriptionLevel(ctx, &args, opts...)
			var s LookupProfessionalServiceSubscriptionLevelResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupProfessionalServiceSubscriptionLevelResultOutput)
}

type LookupProfessionalServiceSubscriptionLevelOutputArgs struct {
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the resource.
	ResourceName pulumi.StringInput `pulumi:"resourceName"`
	// The Azure subscription ID. This is a GUID-formatted string (e.g. 00000000-0000-0000-0000-000000000000)
	SubscriptionId pulumi.StringPtrInput `pulumi:"subscriptionId"`
}

func (LookupProfessionalServiceSubscriptionLevelOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProfessionalServiceSubscriptionLevelArgs)(nil)).Elem()
}

// ProfessionalService REST API resource definition.
type LookupProfessionalServiceSubscriptionLevelResultOutput struct{ *pulumi.OutputState }

func (LookupProfessionalServiceSubscriptionLevelResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProfessionalServiceSubscriptionLevelResult)(nil)).Elem()
}

func (o LookupProfessionalServiceSubscriptionLevelResultOutput) ToLookupProfessionalServiceSubscriptionLevelResultOutput() LookupProfessionalServiceSubscriptionLevelResultOutput {
	return o
}

func (o LookupProfessionalServiceSubscriptionLevelResultOutput) ToLookupProfessionalServiceSubscriptionLevelResultOutputWithContext(ctx context.Context) LookupProfessionalServiceSubscriptionLevelResultOutput {
	return o
}

func (o LookupProfessionalServiceSubscriptionLevelResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupProfessionalServiceSubscriptionLevelResult] {
	return pulumix.Output[LookupProfessionalServiceSubscriptionLevelResult]{
		OutputState: o.OutputState,
	}
}

// The resource uri
func (o LookupProfessionalServiceSubscriptionLevelResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProfessionalServiceSubscriptionLevelResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupProfessionalServiceSubscriptionLevelResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProfessionalServiceSubscriptionLevelResult) string { return v.Name }).(pulumi.StringOutput)
}

// professionalService properties
func (o LookupProfessionalServiceSubscriptionLevelResultOutput) Properties() ProfessionalServiceResourceResponsePropertiesOutput {
	return o.ApplyT(func(v LookupProfessionalServiceSubscriptionLevelResult) ProfessionalServiceResourceResponseProperties {
		return v.Properties
	}).(ProfessionalServiceResourceResponsePropertiesOutput)
}

// the resource tags.
func (o LookupProfessionalServiceSubscriptionLevelResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupProfessionalServiceSubscriptionLevelResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o LookupProfessionalServiceSubscriptionLevelResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProfessionalServiceSubscriptionLevelResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupProfessionalServiceSubscriptionLevelResultOutput{})
}
