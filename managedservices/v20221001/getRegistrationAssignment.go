// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20221001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the details of the specified registration assignment.
func LookupRegistrationAssignment(ctx *pulumi.Context, args *LookupRegistrationAssignmentArgs, opts ...pulumi.InvokeOption) (*LookupRegistrationAssignmentResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupRegistrationAssignmentResult
	err := ctx.Invoke("azure-native:managedservices/v20221001:getRegistrationAssignment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRegistrationAssignmentArgs struct {
	// The flag indicating whether to return the registration definition details along with the registration assignment details.
	ExpandRegistrationDefinition *bool `pulumi:"expandRegistrationDefinition"`
	// The GUID of the registration assignment.
	RegistrationAssignmentId string `pulumi:"registrationAssignmentId"`
	// The scope of the resource.
	Scope string `pulumi:"scope"`
}

// The registration assignment.
type LookupRegistrationAssignmentResult struct {
	// The fully qualified path of the registration assignment.
	Id string `pulumi:"id"`
	// The name of the registration assignment.
	Name string `pulumi:"name"`
	// The properties of a registration assignment.
	Properties RegistrationAssignmentPropertiesResponse `pulumi:"properties"`
	// The metadata for the registration assignment resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the Azure resource (Microsoft.ManagedServices/registrationAssignments).
	Type string `pulumi:"type"`
}

func LookupRegistrationAssignmentOutput(ctx *pulumi.Context, args LookupRegistrationAssignmentOutputArgs, opts ...pulumi.InvokeOption) LookupRegistrationAssignmentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRegistrationAssignmentResult, error) {
			args := v.(LookupRegistrationAssignmentArgs)
			r, err := LookupRegistrationAssignment(ctx, &args, opts...)
			var s LookupRegistrationAssignmentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRegistrationAssignmentResultOutput)
}

type LookupRegistrationAssignmentOutputArgs struct {
	// The flag indicating whether to return the registration definition details along with the registration assignment details.
	ExpandRegistrationDefinition pulumi.BoolPtrInput `pulumi:"expandRegistrationDefinition"`
	// The GUID of the registration assignment.
	RegistrationAssignmentId pulumi.StringInput `pulumi:"registrationAssignmentId"`
	// The scope of the resource.
	Scope pulumi.StringInput `pulumi:"scope"`
}

func (LookupRegistrationAssignmentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRegistrationAssignmentArgs)(nil)).Elem()
}

// The registration assignment.
type LookupRegistrationAssignmentResultOutput struct{ *pulumi.OutputState }

func (LookupRegistrationAssignmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRegistrationAssignmentResult)(nil)).Elem()
}

func (o LookupRegistrationAssignmentResultOutput) ToLookupRegistrationAssignmentResultOutput() LookupRegistrationAssignmentResultOutput {
	return o
}

func (o LookupRegistrationAssignmentResultOutput) ToLookupRegistrationAssignmentResultOutputWithContext(ctx context.Context) LookupRegistrationAssignmentResultOutput {
	return o
}

// The fully qualified path of the registration assignment.
func (o LookupRegistrationAssignmentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistrationAssignmentResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the registration assignment.
func (o LookupRegistrationAssignmentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistrationAssignmentResult) string { return v.Name }).(pulumi.StringOutput)
}

// The properties of a registration assignment.
func (o LookupRegistrationAssignmentResultOutput) Properties() RegistrationAssignmentPropertiesResponseOutput {
	return o.ApplyT(func(v LookupRegistrationAssignmentResult) RegistrationAssignmentPropertiesResponse {
		return v.Properties
	}).(RegistrationAssignmentPropertiesResponseOutput)
}

// The metadata for the registration assignment resource.
func (o LookupRegistrationAssignmentResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupRegistrationAssignmentResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the Azure resource (Microsoft.ManagedServices/registrationAssignments).
func (o LookupRegistrationAssignmentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistrationAssignmentResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRegistrationAssignmentResultOutput{})
}
