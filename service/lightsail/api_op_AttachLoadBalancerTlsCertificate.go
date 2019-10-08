// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lightsail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type AttachLoadBalancerTlsCertificateInput struct {
	_ struct{} `type:"structure"`

	// The name of your SSL/TLS certificate.
	//
	// CertificateName is a required field
	CertificateName *string `locationName:"certificateName" type:"string" required:"true"`

	// The name of the load balancer to which you want to associate the SSL/TLS
	// certificate.
	//
	// LoadBalancerName is a required field
	LoadBalancerName *string `locationName:"loadBalancerName" type:"string" required:"true"`
}

// String returns the string representation
func (s AttachLoadBalancerTlsCertificateInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AttachLoadBalancerTlsCertificateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AttachLoadBalancerTlsCertificateInput"}

	if s.CertificateName == nil {
		invalidParams.Add(aws.NewErrParamRequired("CertificateName"))
	}

	if s.LoadBalancerName == nil {
		invalidParams.Add(aws.NewErrParamRequired("LoadBalancerName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type AttachLoadBalancerTlsCertificateOutput struct {
	_ struct{} `type:"structure"`

	// An object representing the API operations.
	//
	// These SSL/TLS certificates are only usable by Lightsail load balancers. You
	// can't get the certificate and use it for another purpose.
	Operations []Operation `locationName:"operations" type:"list"`
}

// String returns the string representation
func (s AttachLoadBalancerTlsCertificateOutput) String() string {
	return awsutil.Prettify(s)
}

const opAttachLoadBalancerTlsCertificate = "AttachLoadBalancerTlsCertificate"

// AttachLoadBalancerTlsCertificateRequest returns a request value for making API operation for
// Amazon Lightsail.
//
// Attaches a Transport Layer Security (TLS) certificate to your load balancer.
// TLS is just an updated, more secure version of Secure Socket Layer (SSL).
//
// Once you create and validate your certificate, you can attach it to your
// load balancer. You can also use this API to rotate the certificates on your
// account. Use the attach load balancer tls certificate operation with the
// non-attached certificate, and it will replace the existing one and become
// the attached certificate.
//
// The attach load balancer tls certificate operation supports tag-based access
// control via resource tags applied to the resource identified by load balancer
// name. For more information, see the Lightsail Dev Guide (https://lightsail.aws.amazon.com/ls/docs/en/articles/amazon-lightsail-controlling-access-using-tags).
//
//    // Example sending a request using AttachLoadBalancerTlsCertificateRequest.
//    req := client.AttachLoadBalancerTlsCertificateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/AttachLoadBalancerTlsCertificate
func (c *Client) AttachLoadBalancerTlsCertificateRequest(input *AttachLoadBalancerTlsCertificateInput) AttachLoadBalancerTlsCertificateRequest {
	op := &aws.Operation{
		Name:       opAttachLoadBalancerTlsCertificate,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AttachLoadBalancerTlsCertificateInput{}
	}

	req := c.newRequest(op, input, &AttachLoadBalancerTlsCertificateOutput{})
	return AttachLoadBalancerTlsCertificateRequest{Request: req, Input: input, Copy: c.AttachLoadBalancerTlsCertificateRequest}
}

// AttachLoadBalancerTlsCertificateRequest is the request type for the
// AttachLoadBalancerTlsCertificate API operation.
type AttachLoadBalancerTlsCertificateRequest struct {
	*aws.Request
	Input *AttachLoadBalancerTlsCertificateInput
	Copy  func(*AttachLoadBalancerTlsCertificateInput) AttachLoadBalancerTlsCertificateRequest
}

// Send marshals and sends the AttachLoadBalancerTlsCertificate API request.
func (r AttachLoadBalancerTlsCertificateRequest) Send(ctx context.Context) (*AttachLoadBalancerTlsCertificateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AttachLoadBalancerTlsCertificateResponse{
		AttachLoadBalancerTlsCertificateOutput: r.Request.Data.(*AttachLoadBalancerTlsCertificateOutput),
		response:                               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AttachLoadBalancerTlsCertificateResponse is the response type for the
// AttachLoadBalancerTlsCertificate API operation.
type AttachLoadBalancerTlsCertificateResponse struct {
	*AttachLoadBalancerTlsCertificateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AttachLoadBalancerTlsCertificate request.
func (r *AttachLoadBalancerTlsCertificateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
