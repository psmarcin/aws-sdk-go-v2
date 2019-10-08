// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateGroupInput struct {
	_ struct{} `type:"structure"`

	// A string containing the description of the group.
	Description *string `type:"string"`

	// The name of the group. Must be unique.
	//
	// GroupName is a required field
	GroupName *string `min:"1" type:"string" required:"true"`

	// A nonnegative integer value that specifies the precedence of this group relative
	// to the other groups that a user can belong to in the user pool. Zero is the
	// highest precedence value. Groups with lower Precedence values take precedence
	// over groups with higher or null Precedence values. If a user belongs to two
	// or more groups, it is the group with the lowest precedence value whose role
	// ARN will be used in the cognito:roles and cognito:preferred_role claims in
	// the user's tokens.
	//
	// Two groups can have the same Precedence value. If this happens, neither group
	// takes precedence over the other. If two groups with the same Precedence have
	// the same role ARN, that role is used in the cognito:preferred_role claim
	// in tokens for users in each group. If the two groups have different role
	// ARNs, the cognito:preferred_role claim is not set in users' tokens.
	//
	// The default Precedence value is null.
	Precedence *int64 `type:"integer"`

	// The role ARN for the group.
	RoleArn *string `min:"20" type:"string"`

	// The user pool ID for the user pool.
	//
	// UserPoolId is a required field
	UserPoolId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateGroupInput"}

	if s.GroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("GroupName"))
	}
	if s.GroupName != nil && len(*s.GroupName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("GroupName", 1))
	}
	if s.RoleArn != nil && len(*s.RoleArn) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("RoleArn", 20))
	}

	if s.UserPoolId == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserPoolId"))
	}
	if s.UserPoolId != nil && len(*s.UserPoolId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("UserPoolId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateGroupOutput struct {
	_ struct{} `type:"structure"`

	// The group object for the group.
	Group *GroupType `type:"structure"`
}

// String returns the string representation
func (s CreateGroupOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateGroup = "CreateGroup"

// CreateGroupRequest returns a request value for making API operation for
// Amazon Cognito Identity Provider.
//
// Creates a new group in the specified user pool.
//
// Calling this action requires developer credentials.
//
//    // Example sending a request using CreateGroupRequest.
//    req := client.CreateGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-idp-2016-04-18/CreateGroup
func (c *Client) CreateGroupRequest(input *CreateGroupInput) CreateGroupRequest {
	op := &aws.Operation{
		Name:       opCreateGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateGroupInput{}
	}

	req := c.newRequest(op, input, &CreateGroupOutput{})
	return CreateGroupRequest{Request: req, Input: input, Copy: c.CreateGroupRequest}
}

// CreateGroupRequest is the request type for the
// CreateGroup API operation.
type CreateGroupRequest struct {
	*aws.Request
	Input *CreateGroupInput
	Copy  func(*CreateGroupInput) CreateGroupRequest
}

// Send marshals and sends the CreateGroup API request.
func (r CreateGroupRequest) Send(ctx context.Context) (*CreateGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateGroupResponse{
		CreateGroupOutput: r.Request.Data.(*CreateGroupOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateGroupResponse is the response type for the
// CreateGroup API operation.
type CreateGroupResponse struct {
	*CreateGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateGroup request.
func (r *CreateGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
