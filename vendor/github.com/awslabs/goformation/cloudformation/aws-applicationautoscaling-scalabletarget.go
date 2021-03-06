package cloudformation

import (
	"encoding/json"
	"errors"
	"fmt"
)

// AWSApplicationAutoScalingScalableTarget AWS CloudFormation Resource (AWS::ApplicationAutoScaling::ScalableTarget)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalabletarget.html
type AWSApplicationAutoScalingScalableTarget struct {

	// MaxCapacity AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalabletarget.html#cfn-applicationautoscaling-scalabletarget-maxcapacity
	MaxCapacity *Value `json:"MaxCapacity,omitempty"`

	// MinCapacity AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalabletarget.html#cfn-applicationautoscaling-scalabletarget-mincapacity
	MinCapacity *Value `json:"MinCapacity,omitempty"`

	// ResourceId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalabletarget.html#cfn-applicationautoscaling-scalabletarget-resourceid
	ResourceId *Value `json:"ResourceId,omitempty"`

	// RoleARN AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalabletarget.html#cfn-applicationautoscaling-scalabletarget-rolearn
	RoleARN *Value `json:"RoleARN,omitempty"`

	// ScalableDimension AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalabletarget.html#cfn-applicationautoscaling-scalabletarget-scalabledimension
	ScalableDimension *Value `json:"ScalableDimension,omitempty"`

	// ScheduledActions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalabletarget.html#cfn-applicationautoscaling-scalabletarget-scheduledactions
	ScheduledActions []AWSApplicationAutoScalingScalableTarget_ScheduledAction `json:"ScheduledActions,omitempty"`

	// ServiceNamespace AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalabletarget.html#cfn-applicationautoscaling-scalabletarget-servicenamespace
	ServiceNamespace *Value `json:"ServiceNamespace,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSApplicationAutoScalingScalableTarget) AWSCloudFormationType() string {
	return "AWS::ApplicationAutoScaling::ScalableTarget"
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.
func (r *AWSApplicationAutoScalingScalableTarget) MarshalJSON() ([]byte, error) {
	type Properties AWSApplicationAutoScalingScalableTarget
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
func (r *AWSApplicationAutoScalingScalableTarget) UnmarshalJSON(b []byte) error {
	type Properties AWSApplicationAutoScalingScalableTarget
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
		*r = AWSApplicationAutoScalingScalableTarget(*res.Properties)
	}

	return nil
}

// GetAllAWSApplicationAutoScalingScalableTargetResources retrieves all AWSApplicationAutoScalingScalableTarget items from an AWS CloudFormation template
func (t *Template) GetAllAWSApplicationAutoScalingScalableTargetResources() map[string]AWSApplicationAutoScalingScalableTarget {
	results := map[string]AWSApplicationAutoScalingScalableTarget{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case AWSApplicationAutoScalingScalableTarget:
			// We found a strongly typed resource of the correct type; use it
			results[name] = resource
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::ApplicationAutoScaling::ScalableTarget" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						result := &AWSApplicationAutoScalingScalableTarget{}
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

// GetAWSApplicationAutoScalingScalableTargetWithName retrieves all AWSApplicationAutoScalingScalableTarget items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSApplicationAutoScalingScalableTargetWithName(name string) (AWSApplicationAutoScalingScalableTarget, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case AWSApplicationAutoScalingScalableTarget:
			// We found a strongly typed resource of the correct type; use it
			return resource, nil
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::ApplicationAutoScaling::ScalableTarget" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						result := &AWSApplicationAutoScalingScalableTarget{}
						if err := result.UnmarshalJSON(b); err == nil {
							return *result, nil
						}
					}
				}
			}
		}
	}
	return AWSApplicationAutoScalingScalableTarget{}, errors.New("resource not found")
}
