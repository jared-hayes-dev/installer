// Code generated by smithy-go-codegen DO NOT EDIT.

package comprehend

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/comprehend/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Inspects the input text and returns a sentiment analysis for each entity
// identified in the text. For more information about targeted sentiment, see
// Targeted sentiment
// (https://docs.aws.amazon.com/comprehend/latest/dg/how-targeted-sentiment.html).
func (c *Client) DetectTargetedSentiment(ctx context.Context, params *DetectTargetedSentimentInput, optFns ...func(*Options)) (*DetectTargetedSentimentOutput, error) {
	if params == nil {
		params = &DetectTargetedSentimentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DetectTargetedSentiment", params, optFns, c.addOperationDetectTargetedSentimentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DetectTargetedSentimentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DetectTargetedSentimentInput struct {

	// The language of the input documents. Currently, English is the only supported
	// language.
	//
	// This member is required.
	LanguageCode types.LanguageCode

	// A UTF-8 text string. The maximum string length is 5 KB.
	//
	// This member is required.
	Text *string

	noSmithyDocumentSerde
}

type DetectTargetedSentimentOutput struct {

	// Targeted sentiment analysis for each of the entities identified in the input
	// text.
	Entities []types.TargetedSentimentEntity

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDetectTargetedSentimentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDetectTargetedSentiment{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDetectTargetedSentiment{}, middleware.After)
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
	if err = addOpDetectTargetedSentimentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDetectTargetedSentiment(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDetectTargetedSentiment(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "comprehend",
		OperationName: "DetectTargetedSentiment",
	}
}
