/*
Copyright 2010-2019 Amazon.com, Inc. or its affiliates. All Rights Reserved.
This file is licensed under the Apache License, Version 2.0 (the "License").
You may not use this file except in compliance with the License. A copy of
the License is located at

	http://aws.amazon.com/apache2.0/

This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
CONDITIONS OF ANY KIND, either express or implied. See the License for the
specific language governing permissions and limitations under the License.
*/
package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials/stscreds"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

// Usage:
// go run sts_assume_role.go
func main1() {
	// Initialize a session in us-west-2 that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials.
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic(err)
	}

	//roleToAssumeArn := "arn:aws:iam::131891458378:role/Rohit-Tiwari-1-dev-Target2-CloudRanger-GXJ2XQ6UVETC"

	stsSvc := sts.NewFromConfig(cfg)
	//creds := stscreds.NewAssumeRoleProvider(stsSvc, "myRoleArn")

	stsCredProvider := stscreds.NewAssumeRoleProvider(stsSvc, "arn:aws:iam::131891458378:role/Rohit-Tiwari-1-dev-Target2-CloudRanger-GXJ2XQ6UVETC")
	//sess, err := session.NewSession(&cfg)
	value, err := stsCredProvider.Retrieve(context.TODO())
	if err != nil {
		// handle error
		fmt.Println(err, "aray ")
	}
	fmt.Println(value)
	// cfg.Credentials = aws.NewCredentialsCache(stsCredProvider)

	// // Create service client value configured for credentials
	// // from assumed role.
	// svc := s3.NewFromConfig(cfg)

	// // Get the first page of results for ListObjectsV2 for a bucket
	// output, err := svc.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
	// 	Bucket: aws.String("cloudranger-backup-feij22ip-6faaukhi-ap-south-1"),
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(output)
	// log.Println("first page results:")
	// for _, object := range output.Contents {
	// 	log.Printf("key=%s size=%d", aws.ToString(object.Key), object.Size)
	// }
}
