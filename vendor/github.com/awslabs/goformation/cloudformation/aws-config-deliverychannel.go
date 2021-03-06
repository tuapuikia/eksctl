package cloudformation

import (
	"encoding/json"
	"errors"
	"fmt"
)

// AWSConfigDeliveryChannel AWS CloudFormation Resource (AWS::Config::DeliveryChannel)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-deliverychannel.html
type AWSConfigDeliveryChannel struct {

	// ConfigSnapshotDeliveryProperties AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-deliverychannel.html#cfn-config-deliverychannel-configsnapshotdeliveryproperties
	ConfigSnapshotDeliveryProperties *AWSConfigDeliveryChannel_ConfigSnapshotDeliveryProperties `json:"ConfigSnapshotDeliveryProperties,omitempty"`

	// Name AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-deliverychannel.html#cfn-config-deliverychannel-name
	Name *Value `json:"Name,omitempty"`

	// S3BucketName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-deliverychannel.html#cfn-config-deliverychannel-s3bucketname
	S3BucketName *Value `json:"S3BucketName,omitempty"`

	// S3KeyPrefix AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-deliverychannel.html#cfn-config-deliverychannel-s3keyprefix
	S3KeyPrefix *Value `json:"S3KeyPrefix,omitempty"`

	// SnsTopicARN AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-deliverychannel.html#cfn-config-deliverychannel-snstopicarn
	SnsTopicARN *Value `json:"SnsTopicARN,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSConfigDeliveryChannel) AWSCloudFormationType() string {
	return "AWS::Config::DeliveryChannel"
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.
func (r *AWSConfigDeliveryChannel) MarshalJSON() ([]byte, error) {
	type Properties AWSConfigDeliveryChannel
	return json.Marshal(&struct {
		Type       string
		Properties Properties
	}{
		Type:       r.AWSCloudFormationType(),
		Properties: (Properties)(*r),
	})
}

// UnmarshalJSON is a custom JSON unmarshalling hook that strips the outer
// AWS CloudFormation resource object, and just keeps the 'Properties' field.
func (r *AWSConfigDeliveryChannel) UnmarshalJSON(b []byte) error {
	type Properties AWSConfigDeliveryChannel
	res := &struct {
		Type       string
		Properties *Properties
	}{}
	if err := json.Unmarshal(b, &res); err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return err
	}

	// If the resource has no Properties set, it could be nil
	if res.Properties != nil {
		*r = AWSConfigDeliveryChannel(*res.Properties)
	}

	return nil
}

// GetAllAWSConfigDeliveryChannelResources retrieves all AWSConfigDeliveryChannel items from an AWS CloudFormation template
func (t *Template) GetAllAWSConfigDeliveryChannelResources() map[string]AWSConfigDeliveryChannel {
	results := map[string]AWSConfigDeliveryChannel{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case AWSConfigDeliveryChannel:
			// We found a strongly typed resource of the correct type; use it
			results[name] = resource
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::Config::DeliveryChannel" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						result := &AWSConfigDeliveryChannel{}
						if err := result.UnmarshalJSON(b); err == nil {
							results[name] = *result
						}
					}
				}
			}
		}
	}
	return results
}

// GetAWSConfigDeliveryChannelWithName retrieves all AWSConfigDeliveryChannel items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSConfigDeliveryChannelWithName(name string) (AWSConfigDeliveryChannel, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case AWSConfigDeliveryChannel:
			// We found a strongly typed resource of the correct type; use it
			return resource, nil
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::Config::DeliveryChannel" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						result := &AWSConfigDeliveryChannel{}
						if err := result.UnmarshalJSON(b); err == nil {
							return *result, nil
						}
					}
				}
			}
		}
	}
	return AWSConfigDeliveryChannel{}, errors.New("resource not found")
}
