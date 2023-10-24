// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apicenter

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// API entity.
// Azure REST API version: 2024-03-01.
type Api struct {
	pulumi.CustomResourceState

	Contacts ContactResponseArrayOutput `pulumi:"contacts"`
	// The custom metadata defined for API catalog entities.
	CustomProperties pulumi.AnyOutput `pulumi:"customProperties"`
	// Description of the API.
	Description           pulumi.StringPtrOutput                   `pulumi:"description"`
	ExternalDocumentation ExternalDocumentationResponseArrayOutput `pulumi:"externalDocumentation"`
	// Kind of API. For example, REST or GraphQL.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// The license information for the API.
	License LicenseResponsePtrOutput `pulumi:"license"`
	// Current lifecycle stage of the API.
	LifecycleStage pulumi.StringOutput `pulumi:"lifecycleStage"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Short description of the API.
	Summary pulumi.StringPtrOutput `pulumi:"summary"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Terms of service for the API.
	TermsOfService TermsOfServiceResponsePtrOutput `pulumi:"termsOfService"`
	// API title.
	Title pulumi.StringOutput `pulumi:"title"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewApi registers a new resource with the given unique name, arguments, and options.
func NewApi(ctx *pulumi.Context,
	name string, args *ApiArgs, opts ...pulumi.ResourceOption) (*Api, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	if args.Title == nil {
		return nil, errors.New("invalid value for required argument 'Title'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:apicenter/v20240301:Api"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Api
	err := ctx.RegisterResource("azure-native:apicenter:Api", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApi gets an existing Api resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApi(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApiState, opts ...pulumi.ResourceOption) (*Api, error) {
	var resource Api
	err := ctx.ReadResource("azure-native:apicenter:Api", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Api resources.
type apiState struct {
}

type ApiState struct {
}

func (ApiState) ElementType() reflect.Type {
	return reflect.TypeOf((*apiState)(nil)).Elem()
}

type apiArgs struct {
	// The name of the API.
	ApiName  *string   `pulumi:"apiName"`
	Contacts []Contact `pulumi:"contacts"`
	// The custom metadata defined for API catalog entities.
	CustomProperties interface{} `pulumi:"customProperties"`
	// Description of the API.
	Description           *string                 `pulumi:"description"`
	ExternalDocumentation []ExternalDocumentation `pulumi:"externalDocumentation"`
	// Kind of API. For example, REST or GraphQL.
	Kind string `pulumi:"kind"`
	// The license information for the API.
	License *License `pulumi:"license"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of Azure API Center service.
	ServiceName string `pulumi:"serviceName"`
	// Short description of the API.
	Summary *string `pulumi:"summary"`
	// Terms of service for the API.
	TermsOfService *TermsOfService `pulumi:"termsOfService"`
	// API title.
	Title string `pulumi:"title"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a Api resource.
type ApiArgs struct {
	// The name of the API.
	ApiName  pulumi.StringPtrInput
	Contacts ContactArrayInput
	// The custom metadata defined for API catalog entities.
	CustomProperties pulumi.Input
	// Description of the API.
	Description           pulumi.StringPtrInput
	ExternalDocumentation ExternalDocumentationArrayInput
	// Kind of API. For example, REST or GraphQL.
	Kind pulumi.StringInput
	// The license information for the API.
	License LicensePtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of Azure API Center service.
	ServiceName pulumi.StringInput
	// Short description of the API.
	Summary pulumi.StringPtrInput
	// Terms of service for the API.
	TermsOfService TermsOfServicePtrInput
	// API title.
	Title pulumi.StringInput
	// The name of the workspace.
	WorkspaceName pulumi.StringInput
}

func (ApiArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apiArgs)(nil)).Elem()
}

type ApiInput interface {
	pulumi.Input

	ToApiOutput() ApiOutput
	ToApiOutputWithContext(ctx context.Context) ApiOutput
}

func (*Api) ElementType() reflect.Type {
	return reflect.TypeOf((**Api)(nil)).Elem()
}

func (i *Api) ToApiOutput() ApiOutput {
	return i.ToApiOutputWithContext(context.Background())
}

func (i *Api) ToApiOutputWithContext(ctx context.Context) ApiOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiOutput)
}

func (i *Api) ToOutput(ctx context.Context) pulumix.Output[*Api] {
	return pulumix.Output[*Api]{
		OutputState: i.ToApiOutputWithContext(ctx).OutputState,
	}
}

type ApiOutput struct{ *pulumi.OutputState }

func (ApiOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Api)(nil)).Elem()
}

func (o ApiOutput) ToApiOutput() ApiOutput {
	return o
}

func (o ApiOutput) ToApiOutputWithContext(ctx context.Context) ApiOutput {
	return o
}

func (o ApiOutput) ToOutput(ctx context.Context) pulumix.Output[*Api] {
	return pulumix.Output[*Api]{
		OutputState: o.OutputState,
	}
}

func (o ApiOutput) Contacts() ContactResponseArrayOutput {
	return o.ApplyT(func(v *Api) ContactResponseArrayOutput { return v.Contacts }).(ContactResponseArrayOutput)
}

// The custom metadata defined for API catalog entities.
func (o ApiOutput) CustomProperties() pulumi.AnyOutput {
	return o.ApplyT(func(v *Api) pulumi.AnyOutput { return v.CustomProperties }).(pulumi.AnyOutput)
}

// Description of the API.
func (o ApiOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Api) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ApiOutput) ExternalDocumentation() ExternalDocumentationResponseArrayOutput {
	return o.ApplyT(func(v *Api) ExternalDocumentationResponseArrayOutput { return v.ExternalDocumentation }).(ExternalDocumentationResponseArrayOutput)
}

// Kind of API. For example, REST or GraphQL.
func (o ApiOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *Api) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// The license information for the API.
func (o ApiOutput) License() LicenseResponsePtrOutput {
	return o.ApplyT(func(v *Api) LicenseResponsePtrOutput { return v.License }).(LicenseResponsePtrOutput)
}

// Current lifecycle stage of the API.
func (o ApiOutput) LifecycleStage() pulumi.StringOutput {
	return o.ApplyT(func(v *Api) pulumi.StringOutput { return v.LifecycleStage }).(pulumi.StringOutput)
}

// The name of the resource
func (o ApiOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Api) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Short description of the API.
func (o ApiOutput) Summary() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Api) pulumi.StringPtrOutput { return v.Summary }).(pulumi.StringPtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o ApiOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Api) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Terms of service for the API.
func (o ApiOutput) TermsOfService() TermsOfServiceResponsePtrOutput {
	return o.ApplyT(func(v *Api) TermsOfServiceResponsePtrOutput { return v.TermsOfService }).(TermsOfServiceResponsePtrOutput)
}

// API title.
func (o ApiOutput) Title() pulumi.StringOutput {
	return o.ApplyT(func(v *Api) pulumi.StringOutput { return v.Title }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ApiOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Api) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ApiOutput{})
}
