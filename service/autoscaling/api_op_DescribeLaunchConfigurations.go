// Code generated by smithy-go-codegen DO NOT EDIT.

package autoscaling

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets information about the launch configurations in the account and Region.
func (c *Client) DescribeLaunchConfigurations(ctx context.Context, params *DescribeLaunchConfigurationsInput, optFns ...func(*Options)) (*DescribeLaunchConfigurationsOutput, error) {
	if params == nil {
		params = &DescribeLaunchConfigurationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeLaunchConfigurations", params, optFns, c.addOperationDescribeLaunchConfigurationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeLaunchConfigurationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeLaunchConfigurationsInput struct {

	// The launch configuration names. If you omit this property, all launch
	// configurations are described. Array Members: Maximum number of 50 items.
	LaunchConfigurationNames []string

	// The maximum number of items to return with this call. The default value is 50
	// and the maximum value is 100.
	MaxRecords *int32

	// The token for the next set of items to return. (You received this token from a
	// previous call.)
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeLaunchConfigurationsOutput struct {

	// The launch configurations.
	//
	// This member is required.
	LaunchConfigurations []types.LaunchConfiguration

	// A string that indicates that the response contains more items than can be
	// returned in a single response. To receive additional items, specify this string
	// for the NextToken value when requesting the next set of items. This value is
	// null when there are no more items to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeLaunchConfigurationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeLaunchConfigurations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeLaunchConfigurations{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeLaunchConfigurations(options.Region), middleware.Before); err != nil {
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

// DescribeLaunchConfigurationsAPIClient is a client that implements the
// DescribeLaunchConfigurations operation.
type DescribeLaunchConfigurationsAPIClient interface {
	DescribeLaunchConfigurations(context.Context, *DescribeLaunchConfigurationsInput, ...func(*Options)) (*DescribeLaunchConfigurationsOutput, error)
}

var _ DescribeLaunchConfigurationsAPIClient = (*Client)(nil)

// DescribeLaunchConfigurationsPaginatorOptions is the paginator options for
// DescribeLaunchConfigurations
type DescribeLaunchConfigurationsPaginatorOptions struct {
	// The maximum number of items to return with this call. The default value is 50
	// and the maximum value is 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeLaunchConfigurationsPaginator is a paginator for
// DescribeLaunchConfigurations
type DescribeLaunchConfigurationsPaginator struct {
	options   DescribeLaunchConfigurationsPaginatorOptions
	client    DescribeLaunchConfigurationsAPIClient
	params    *DescribeLaunchConfigurationsInput
	nextToken *string
	firstPage bool
}

// NewDescribeLaunchConfigurationsPaginator returns a new
// DescribeLaunchConfigurationsPaginator
func NewDescribeLaunchConfigurationsPaginator(client DescribeLaunchConfigurationsAPIClient, params *DescribeLaunchConfigurationsInput, optFns ...func(*DescribeLaunchConfigurationsPaginatorOptions)) *DescribeLaunchConfigurationsPaginator {
	if params == nil {
		params = &DescribeLaunchConfigurationsInput{}
	}

	options := DescribeLaunchConfigurationsPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeLaunchConfigurationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeLaunchConfigurationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeLaunchConfigurations page.
func (p *DescribeLaunchConfigurationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeLaunchConfigurationsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxRecords = limit

	result, err := p.client.DescribeLaunchConfigurations(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opDescribeLaunchConfigurations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "autoscaling",
		OperationName: "DescribeLaunchConfigurations",
	}
}
