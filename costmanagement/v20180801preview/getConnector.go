// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20180801preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get a connector definition
func LookupConnector(ctx *pulumi.Context, args *LookupConnectorArgs, opts ...pulumi.InvokeOption) (*LookupConnectorResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupConnectorResult
	err := ctx.Invoke("azure-native:costmanagement/v20180801preview:getConnector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConnectorArgs struct {
	// Connector Name.
	ConnectorName string `pulumi:"connectorName"`
	// Azure Resource Group Name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The Connector model definition
type LookupConnectorResult struct {
	// Collection information
	Collection ConnectorCollectionInfoResponse `pulumi:"collection"`
	// Connector definition creation datetime
	CreatedOn string `pulumi:"createdOn"`
	// Credentials authentication key (eg AWS ARN)
	CredentialsKey *string `pulumi:"credentialsKey"`
	// Connector DisplayName (defaults to Name)
	DisplayName *string `pulumi:"displayName"`
	// Connector id
	Id string `pulumi:"id"`
	// Connector kind (eg aws)
	Kind *string `pulumi:"kind"`
	// Connector location
	Location *string `pulumi:"location"`
	// Connector last modified datetime
	ModifiedOn string `pulumi:"modifiedOn"`
	// Connector name
	Name string `pulumi:"name"`
	// Connector providerAccountId (determined from credentials)
	ProviderAccountId string `pulumi:"providerAccountId"`
	// Identifying source report. (For AWS this is a CUR report name, defined with Daily and with Resources)
	ReportId *string `pulumi:"reportId"`
	// Connector status
	Status *string `pulumi:"status"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Connector type
	Type string `pulumi:"type"`
}

func LookupConnectorOutput(ctx *pulumi.Context, args LookupConnectorOutputArgs, opts ...pulumi.InvokeOption) LookupConnectorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupConnectorResult, error) {
			args := v.(LookupConnectorArgs)
			r, err := LookupConnector(ctx, &args, opts...)
			var s LookupConnectorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupConnectorResultOutput)
}

type LookupConnectorOutputArgs struct {
	// Connector Name.
	ConnectorName pulumi.StringInput `pulumi:"connectorName"`
	// Azure Resource Group Name.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupConnectorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConnectorArgs)(nil)).Elem()
}

// The Connector model definition
type LookupConnectorResultOutput struct{ *pulumi.OutputState }

func (LookupConnectorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConnectorResult)(nil)).Elem()
}

func (o LookupConnectorResultOutput) ToLookupConnectorResultOutput() LookupConnectorResultOutput {
	return o
}

func (o LookupConnectorResultOutput) ToLookupConnectorResultOutputWithContext(ctx context.Context) LookupConnectorResultOutput {
	return o
}

// Collection information
func (o LookupConnectorResultOutput) Collection() ConnectorCollectionInfoResponseOutput {
	return o.ApplyT(func(v LookupConnectorResult) ConnectorCollectionInfoResponse { return v.Collection }).(ConnectorCollectionInfoResponseOutput)
}

// Connector definition creation datetime
func (o LookupConnectorResultOutput) CreatedOn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectorResult) string { return v.CreatedOn }).(pulumi.StringOutput)
}

// Credentials authentication key (eg AWS ARN)
func (o LookupConnectorResultOutput) CredentialsKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConnectorResult) *string { return v.CredentialsKey }).(pulumi.StringPtrOutput)
}

// Connector DisplayName (defaults to Name)
func (o LookupConnectorResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConnectorResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// Connector id
func (o LookupConnectorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectorResult) string { return v.Id }).(pulumi.StringOutput)
}

// Connector kind (eg aws)
func (o LookupConnectorResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConnectorResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// Connector location
func (o LookupConnectorResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConnectorResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// Connector last modified datetime
func (o LookupConnectorResultOutput) ModifiedOn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectorResult) string { return v.ModifiedOn }).(pulumi.StringOutput)
}

// Connector name
func (o LookupConnectorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectorResult) string { return v.Name }).(pulumi.StringOutput)
}

// Connector providerAccountId (determined from credentials)
func (o LookupConnectorResultOutput) ProviderAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectorResult) string { return v.ProviderAccountId }).(pulumi.StringOutput)
}

// Identifying source report. (For AWS this is a CUR report name, defined with Daily and with Resources)
func (o LookupConnectorResultOutput) ReportId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConnectorResult) *string { return v.ReportId }).(pulumi.StringPtrOutput)
}

// Connector status
func (o LookupConnectorResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConnectorResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

// Resource tags.
func (o LookupConnectorResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupConnectorResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Connector type
func (o LookupConnectorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectorResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupConnectorResultOutput{})
}
