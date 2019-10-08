// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the request to update the user's attributes as an administrator.
type AdminUpdateUserAttributesInput struct {
	_ struct{} `type:"structure"`

	// A map of custom key-value pairs that you can provide as input for any custom
	// workflows that this action triggers.
	//
	// You create custom workflows by assigning AWS Lambda functions to user pool
	// triggers. When you use the AdminUpdateUserAttributes API action, Amazon Cognito
	// invokes the function that is assigned to the custom message trigger. When
	// Amazon Cognito invokes this function, it passes a JSON payload, which the
	// function receives as input. This payload contains a clientMetadata attribute,
	// which provides the data that you assigned to the ClientMetadata parameter
	// in your AdminUpdateUserAttributes request. In your function code in AWS Lambda,
	// you can process the clientMetadata value to enhance your workflow for your
	// specific needs.
	//
	// For more information, see Customizing User Pool Workflows with Lambda Triggers
	// (https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-identity-pools-working-with-aws-lambda-triggers.html)
	// in the Amazon Cognito Developer Guide.
	//
	// Take the following limitations into consideration when you use the ClientMetadata
	// parameter:
	//
	//    * Amazon Cognito does not store the ClientMetadata value. This data is
	//    available only to AWS Lambda triggers that are assigned to a user pool
	//    to support custom workflows. If your user pool configuration does not
	//    include triggers, the ClientMetadata parameter serves no purpose.
	//
	//    * Amazon Cognito does not validate the ClientMetadata value.
	//
	//    * Amazon Cognito does not encrypt the the ClientMetadata value, so don't
	//    use it to provide sensitive information.
	ClientMetadata map[string]string `type:"map"`

	// An array of name-value pairs representing user attributes.
	//
	// For custom attributes, you must prepend the custom: prefix to the attribute
	// name.
	//
	// UserAttributes is a required field
	UserAttributes []AttributeType `type:"list" required:"true"`

	// The user pool ID for the user pool where you want to update user attributes.
	//
	// UserPoolId is a required field
	UserPoolId *string `min:"1" type:"string" required:"true"`

	// The user name of the user for whom you want to update user attributes.
	//
	// Username is a required field
	Username *string `min:"1" type:"string" required:"true" sensitive:"true"`
}

// String returns the string representation
func (s AdminUpdateUserAttributesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AdminUpdateUserAttributesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AdminUpdateUserAttributesInput"}

	if s.UserAttributes == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserAttributes"))
	}

	if s.UserPoolId == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserPoolId"))
	}
	if s.UserPoolId != nil && len(*s.UserPoolId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("UserPoolId", 1))
	}

	if s.Username == nil {
		invalidParams.Add(aws.NewErrParamRequired("Username"))
	}
	if s.Username != nil && len(*s.Username) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Username", 1))
	}
	if s.UserAttributes != nil {
		for i, v := range s.UserAttributes {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "UserAttributes", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the response from the server for the request to update user attributes
// as an administrator.
type AdminUpdateUserAttributesOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s AdminUpdateUserAttributesOutput) String() string {
	return awsutil.Prettify(s)
}

const opAdminUpdateUserAttributes = "AdminUpdateUserAttributes"

// AdminUpdateUserAttributesRequest returns a request value for making API operation for
// Amazon Cognito Identity Provider.
//
// Updates the specified user's attributes, including developer attributes,
// as an administrator. Works on any user.
//
// For custom attributes, you must prepend the custom: prefix to the attribute
// name.
//
// In addition to updating user attributes, this API can also be used to mark
// phone and email as verified.
//
// Calling this action requires developer credentials.
//
//    // Example sending a request using AdminUpdateUserAttributesRequest.
//    req := client.AdminUpdateUserAttributesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-idp-2016-04-18/AdminUpdateUserAttributes
func (c *Client) AdminUpdateUserAttributesRequest(input *AdminUpdateUserAttributesInput) AdminUpdateUserAttributesRequest {
	op := &aws.Operation{
		Name:       opAdminUpdateUserAttributes,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AdminUpdateUserAttributesInput{}
	}

	req := c.newRequest(op, input, &AdminUpdateUserAttributesOutput{})
	return AdminUpdateUserAttributesRequest{Request: req, Input: input, Copy: c.AdminUpdateUserAttributesRequest}
}

// AdminUpdateUserAttributesRequest is the request type for the
// AdminUpdateUserAttributes API operation.
type AdminUpdateUserAttributesRequest struct {
	*aws.Request
	Input *AdminUpdateUserAttributesInput
	Copy  func(*AdminUpdateUserAttributesInput) AdminUpdateUserAttributesRequest
}

// Send marshals and sends the AdminUpdateUserAttributes API request.
func (r AdminUpdateUserAttributesRequest) Send(ctx context.Context) (*AdminUpdateUserAttributesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AdminUpdateUserAttributesResponse{
		AdminUpdateUserAttributesOutput: r.Request.Data.(*AdminUpdateUserAttributesOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AdminUpdateUserAttributesResponse is the response type for the
// AdminUpdateUserAttributes API operation.
type AdminUpdateUserAttributesResponse struct {
	*AdminUpdateUserAttributesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AdminUpdateUserAttributes request.
func (r *AdminUpdateUserAttributesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
