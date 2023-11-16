// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210701preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves details of a specific security connector
func LookupSecurityConnector(ctx *pulumi.Context, args *LookupSecurityConnectorArgs, opts ...pulumi.InvokeOption) (*LookupSecurityConnectorResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupSecurityConnectorResult
	err := ctx.Invoke("azure-native:security/v20210701preview:getSecurityConnector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSecurityConnectorArgs struct {
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The security connector name.
	SecurityConnectorName string `pulumi:"securityConnectorName"`
}

// The security connector resource.
type LookupSecurityConnectorResult struct {
	// The multi cloud resource's cloud name.
	CloudName *string `pulumi:"cloudName"`
	// Entity tag is used for comparing two or more entities from the same requested resource.
	Etag *string `pulumi:"etag"`
	// The multi cloud resource identifier (account id in case of AWS connector).
	HierarchyIdentifier *string `pulumi:"hierarchyIdentifier"`
	// Resource Id
	Id string `pulumi:"id"`
	// Kind of the resource
	Kind *string `pulumi:"kind"`
	// Location where the resource is stored
	Location *string `pulumi:"location"`
	// Resource name
	Name string `pulumi:"name"`
	// A collection of offerings for the security connector.
	Offerings []interface{} `pulumi:"offerings"`
	// The multi cloud account's organizational data
	OrganizationalData *SecurityConnectorPropertiesResponseOrganizationalData `pulumi:"organizationalData"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// A list of key value pairs that describe the resource.
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type string `pulumi:"type"`
}

func LookupSecurityConnectorOutput(ctx *pulumi.Context, args LookupSecurityConnectorOutputArgs, opts ...pulumi.InvokeOption) LookupSecurityConnectorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSecurityConnectorResult, error) {
			args := v.(LookupSecurityConnectorArgs)
			r, err := LookupSecurityConnector(ctx, &args, opts...)
			var s LookupSecurityConnectorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSecurityConnectorResultOutput)
}

type LookupSecurityConnectorOutputArgs struct {
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The security connector name.
	SecurityConnectorName pulumi.StringInput `pulumi:"securityConnectorName"`
}

func (LookupSecurityConnectorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSecurityConnectorArgs)(nil)).Elem()
}

// The security connector resource.
type LookupSecurityConnectorResultOutput struct{ *pulumi.OutputState }

func (LookupSecurityConnectorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSecurityConnectorResult)(nil)).Elem()
}

func (o LookupSecurityConnectorResultOutput) ToLookupSecurityConnectorResultOutput() LookupSecurityConnectorResultOutput {
	return o
}

func (o LookupSecurityConnectorResultOutput) ToLookupSecurityConnectorResultOutputWithContext(ctx context.Context) LookupSecurityConnectorResultOutput {
	return o
}

// The multi cloud resource's cloud name.
func (o LookupSecurityConnectorResultOutput) CloudName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSecurityConnectorResult) *string { return v.CloudName }).(pulumi.StringPtrOutput)
}

// Entity tag is used for comparing two or more entities from the same requested resource.
func (o LookupSecurityConnectorResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSecurityConnectorResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// The multi cloud resource identifier (account id in case of AWS connector).
func (o LookupSecurityConnectorResultOutput) HierarchyIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSecurityConnectorResult) *string { return v.HierarchyIdentifier }).(pulumi.StringPtrOutput)
}

// Resource Id
func (o LookupSecurityConnectorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecurityConnectorResult) string { return v.Id }).(pulumi.StringOutput)
}

// Kind of the resource
func (o LookupSecurityConnectorResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSecurityConnectorResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// Location where the resource is stored
func (o LookupSecurityConnectorResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSecurityConnectorResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource name
func (o LookupSecurityConnectorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecurityConnectorResult) string { return v.Name }).(pulumi.StringOutput)
}

// A collection of offerings for the security connector.
func (o LookupSecurityConnectorResultOutput) Offerings() pulumi.ArrayOutput {
	return o.ApplyT(func(v LookupSecurityConnectorResult) []interface{} { return v.Offerings }).(pulumi.ArrayOutput)
}

// The multi cloud account's organizational data
func (o LookupSecurityConnectorResultOutput) OrganizationalData() SecurityConnectorPropertiesResponseOrganizationalDataPtrOutput {
	return o.ApplyT(func(v LookupSecurityConnectorResult) *SecurityConnectorPropertiesResponseOrganizationalData {
		return v.OrganizationalData
	}).(SecurityConnectorPropertiesResponseOrganizationalDataPtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupSecurityConnectorResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupSecurityConnectorResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// A list of key value pairs that describe the resource.
func (o LookupSecurityConnectorResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSecurityConnectorResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type
func (o LookupSecurityConnectorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecurityConnectorResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSecurityConnectorResultOutput{})
}
