// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Starts a job run using a job definition.
func (c *Client) StartJobRun(ctx context.Context, params *StartJobRunInput, optFns ...func(*Options)) (*StartJobRunOutput, error) {
	if params == nil {
		params = &StartJobRunInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartJobRun", params, optFns, c.addOperationStartJobRunMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartJobRunOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartJobRunInput struct {

	// The name of the job definition to use.
	//
	// This member is required.
	JobName *string

	// This field is deprecated. Use MaxCapacity instead. The number of Glue data
	// processing units (DPUs) to allocate to this JobRun. You can allocate a minimum
	// of 2 DPUs; the default is 10. A DPU is a relative measure of processing power
	// that consists of 4 vCPUs of compute capacity and 16 GB of memory. For more
	// information, see the Glue pricing page (https://aws.amazon.com/glue/pricing/).
	//
	// Deprecated: This property is deprecated, use MaxCapacity instead.
	AllocatedCapacity int32

	// The job arguments specifically for this run. For this job run, they replace the
	// default arguments set in the job definition itself. You can specify arguments
	// here that your own job-execution script consumes, as well as arguments that Glue
	// itself consumes. Job arguments may be logged. Do not pass plaintext secrets as
	// arguments. Retrieve secrets from a Glue Connection, Secrets Manager or other
	// secret management mechanism if you intend to keep them within the Job. For
	// information about how to specify and consume your own Job arguments, see the
	// Calling Glue APIs in Python
	// (https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-python-calling.html)
	// topic in the developer guide. For information about the key-value pairs that
	// Glue consumes to set up your job, see the Special Parameters Used by Glue
	// (https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html)
	// topic in the developer guide.
	Arguments map[string]string

	// The ID of a previous JobRun to retry.
	JobRunId *string

	// The number of Glue data processing units (DPUs) that can be allocated when this
	// job runs. A DPU is a relative measure of processing power that consists of 4
	// vCPUs of compute capacity and 16 GB of memory. For more information, see the
	// Glue pricing page (https://aws.amazon.com/glue/pricing/). Do not set Max
	// Capacity if using WorkerType and NumberOfWorkers. The value that can be
	// allocated for MaxCapacity depends on whether you are running a Python shell job,
	// or an Apache Spark ETL job:
	//
	// * When you specify a Python shell job
	// (JobCommand.Name="pythonshell"), you can allocate either 0.0625 or 1 DPU. The
	// default is 0.0625 DPU.
	//
	// * When you specify an Apache Spark ETL job
	// (JobCommand.Name="glueetl"), you can allocate a minimum of 2 DPUs. The default
	// is 10 DPUs. This job type cannot have a fractional DPU allocation.
	MaxCapacity *float64

	// Specifies configuration properties of a job run notification.
	NotificationProperty *types.NotificationProperty

	// The number of workers of a defined workerType that are allocated when a job
	// runs.
	NumberOfWorkers *int32

	// The name of the SecurityConfiguration structure to be used with this job run.
	SecurityConfiguration *string

	// The JobRun timeout in minutes. This is the maximum time that a job run can
	// consume resources before it is terminated and enters TIMEOUT status. The default
	// is 2,880 minutes (48 hours). This overrides the timeout value set in the parent
	// job.
	Timeout *int32

	// The type of predefined worker that is allocated when a job runs. Accepts a value
	// of Standard, G.1X, G.2X, or G.025X.
	//
	// * For the Standard worker type, each worker
	// provides 4 vCPU, 16 GB of memory and a 50GB disk, and 2 executors per worker.
	//
	// *
	// For the G.1X worker type, each worker provides 4 vCPU, 16 GB of memory and a
	// 64GB disk, and 1 executor per worker.
	//
	// * For the G.2X worker type, each worker
	// provides 8 vCPU, 32 GB of memory and a 128GB disk, and 1 executor per worker.
	//
	// *
	// For the G.025X worker type, each worker maps to 0.25 DPU (2 vCPU, 4 GB of
	// memory, 64 GB disk), and provides 1 executor per worker. We recommend this
	// worker type for low volume streaming jobs. This worker type is only available
	// for Glue version 3.0 streaming jobs.
	WorkerType types.WorkerType

	noSmithyDocumentSerde
}

type StartJobRunOutput struct {

	// The ID assigned to this job run.
	JobRunId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartJobRunMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStartJobRun{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartJobRun{}, middleware.After)
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
	if err = addOpStartJobRunValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartJobRun(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartJobRun(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "StartJobRun",
	}
}
