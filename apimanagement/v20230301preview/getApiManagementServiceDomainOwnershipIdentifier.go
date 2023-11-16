// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230301preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get the custom domain ownership identifier for an API Management service.
func GetApiManagementServiceDomainOwnershipIdentifier(ctx *pulumi.Context, args *GetApiManagementServiceDomainOwnershipIdentifierArgs, opts ...pulumi.InvokeOption) (*GetApiManagementServiceDomainOwnershipIdentifierResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv GetApiManagementServiceDomainOwnershipIdentifierResult
	err := ctx.Invoke("azure-native:apimanagement/v20230301preview:getApiManagementServiceDomainOwnershipIdentifier", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetApiManagementServiceDomainOwnershipIdentifierArgs struct {
}

// Response of the GetDomainOwnershipIdentifier operation.
type GetApiManagementServiceDomainOwnershipIdentifierResult struct {
	// The domain ownership identifier value.
	DomainOwnershipIdentifier string `pulumi:"domainOwnershipIdentifier"`
}

func GetApiManagementServiceDomainOwnershipIdentifierOutput(ctx *pulumi.Context, args GetApiManagementServiceDomainOwnershipIdentifierOutputArgs, opts ...pulumi.InvokeOption) GetApiManagementServiceDomainOwnershipIdentifierResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetApiManagementServiceDomainOwnershipIdentifierResult, error) {
			args := v.(GetApiManagementServiceDomainOwnershipIdentifierArgs)
			r, err := GetApiManagementServiceDomainOwnershipIdentifier(ctx, &args, opts...)
			var s GetApiManagementServiceDomainOwnershipIdentifierResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetApiManagementServiceDomainOwnershipIdentifierResultOutput)
}

type GetApiManagementServiceDomainOwnershipIdentifierOutputArgs struct {
}

func (GetApiManagementServiceDomainOwnershipIdentifierOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetApiManagementServiceDomainOwnershipIdentifierArgs)(nil)).Elem()
}

// Response of the GetDomainOwnershipIdentifier operation.
type GetApiManagementServiceDomainOwnershipIdentifierResultOutput struct{ *pulumi.OutputState }

func (GetApiManagementServiceDomainOwnershipIdentifierResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetApiManagementServiceDomainOwnershipIdentifierResult)(nil)).Elem()
}

func (o GetApiManagementServiceDomainOwnershipIdentifierResultOutput) ToGetApiManagementServiceDomainOwnershipIdentifierResultOutput() GetApiManagementServiceDomainOwnershipIdentifierResultOutput {
	return o
}

func (o GetApiManagementServiceDomainOwnershipIdentifierResultOutput) ToGetApiManagementServiceDomainOwnershipIdentifierResultOutputWithContext(ctx context.Context) GetApiManagementServiceDomainOwnershipIdentifierResultOutput {
	return o
}

// The domain ownership identifier value.
func (o GetApiManagementServiceDomainOwnershipIdentifierResultOutput) DomainOwnershipIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v GetApiManagementServiceDomainOwnershipIdentifierResult) string {
		return v.DomainOwnershipIdentifier
	}).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetApiManagementServiceDomainOwnershipIdentifierResultOutput{})
}
