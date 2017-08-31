/*
 * Copyright © 2017. TIBCO Software Inc.
 * This file is subject to the license terms contained
 * in the license file that is distributed with this file.
 */
package sqssendmessage

import (
	"io/ioutil"
	"testing"

	"github.com/TIBCOSoftware/flogo-contrib/action/flow/test"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/stretchr/testify/assert"
)

var activityMetadata *activity.Metadata
var connectionData = ``

func getActivityMetadata() *activity.Metadata {
	if activityMetadata == nil {
		jsonMetadataBytes, err := ioutil.ReadFile("activity.json")
		if err != nil {
			panic("No Json Metadata found for activity.json path")
		}
		activityMetadata = activity.NewMetadata(string(jsonMetadataBytes))
	}
	return activityMetadata
}

func TestActivityRegistration(t *testing.T) {
	act := NewActivity(getActivityMetadata())
	if act == nil {
		t.Error("Activity Not Registered")
		t.Fail()
		return
	}
}

func TestEval(t *testing.T) {
	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(act.Metadata())

	dummyConnectionData := make(map[string]interface{})
	dummyConnectionSettings := make(map[string]interface{}, 4)
	//Use your AWS information
	dummyConnectionSettings["accesskeyId"] = "<YOUR ACCESS KEY ID>"
	dummyConnectionSettings["secreteAccessKey"] = "<YOUR SECRETE ACCESS KEY>"
	dummyConnectionSettings["region"] = "<REGION NAME WHERE SQS IS RUNNING>"
	dummyConnectionSettings["name"] = "My SQS Connection"

	dummyConnectionData["settings"] = dummyConnectionSettings

	tc.SetInput(ivConnection, dummyConnectionData)
	tc.SetInput(ivQueueUrl, "<YOUR SQS QUEUE URL>")
	tc.SetInput(ivMessageBody, "Message from TIBCO")

	_, err := act.Eval(tc)
	assert.Nil(t, err)
}