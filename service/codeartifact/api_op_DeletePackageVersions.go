// Code generated by smithy-go-codegen DO NOT EDIT.

package codeartifact

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codeartifact/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes one or more versions of a package. A deleted package version cannot be
// restored in your repository. If you want to remove a package version from your
// repository and be able to restore it later, set its status to Archived. Archived
// packages cannot be downloaded from a repository and don't show up with list
// package APIs (for example, ListackageVersions
// (https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_ListPackageVersions.html)),
// but you can restore them using UpdatePackageVersionsStatus
// (https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_UpdatePackageVersionsStatus.html).
func (c *Client) DeletePackageVersions(ctx context.Context, params *DeletePackageVersionsInput, optFns ...func(*Options)) (*DeletePackageVersionsOutput, error) {
	if params == nil {
		params = &DeletePackageVersionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeletePackageVersions", params, optFns, c.addOperationDeletePackageVersionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeletePackageVersionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeletePackageVersionsInput struct {

	// The name of the domain that contains the package to delete.
	//
	// This member is required.
	Domain *string

	// The format of the package versions to delete.
	//
	// This member is required.
	Format types.PackageFormat

	// The name of the package with the versions to delete.
	//
	// This member is required.
	Package *string

	// The name of the repository that contains the package versions to delete.
	//
	// This member is required.
	Repository *string

	// An array of strings that specify the versions of the package to delete.
	//
	// This member is required.
	Versions []string

	// The 12-digit account number of the Amazon Web Services account that owns the
	// domain. It does not include dashes or spaces.
	DomainOwner *string

	// The expected status of the package version to delete.
	ExpectedStatus types.PackageVersionStatus

	// The namespace of the package versions to be deleted. The package version
	// component that specifies its namespace depends on its type. For example:
	//
	// * The
	// namespace of a Maven package version is its groupId. The namespace is required
	// when deleting Maven package versions.
	//
	// * The namespace of an npm package version
	// is its scope.
	//
	// * Python and NuGet package versions do not contain a
	// corresponding component, package versions of those formats do not have a
	// namespace.
	Namespace *string

	noSmithyDocumentSerde
}

type DeletePackageVersionsOutput struct {

	// A PackageVersionError object that contains a map of errors codes for the deleted
	// package that failed. The possible error codes are:
	//
	// * ALREADY_EXISTS
	//
	// *
	// MISMATCHED_REVISION
	//
	// * MISMATCHED_STATUS
	//
	// * NOT_ALLOWED
	//
	// * NOT_FOUND
	//
	// * SKIPPED
	FailedVersions map[string]types.PackageVersionError

	// A list of the package versions that were successfully deleted. The status of
	// every successful version will be Deleted.
	SuccessfulVersions map[string]types.SuccessfulPackageVersionInfo

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeletePackageVersionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeletePackageVersions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeletePackageVersions{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpDeletePackageVersionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeletePackageVersions(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opDeletePackageVersions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codeartifact",
		OperationName: "DeletePackageVersions",
	}
}
