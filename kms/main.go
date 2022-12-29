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
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	_ "crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials/stscreds"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

// Usage:
// go run sts_assume_role.go
func main() {
	// Initialize a session in us-west-2 that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials.
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile("default"))
	if err != nil {
		panic(err)
	}

	stsSvc := sts.NewFromConfig(cfg)

	stsCredProvider := stscreds.NewAssumeRoleProvider(stsSvc, "arn:aws:iam::131891458378:role/Rohit-Tiwari-1-dev-Target2-CloudRanger-GXJ2XQ6UVETC", func(o *stscreds.AssumeRoleOptions) {
		o.ExternalID = aws.String("cloudranger_pCJHUXi3")
	})
	//sess, err := session.NewSession(&cfg)
	// value, err := stsCredProvider.Retrieve(context.TODO())
	// if err != nil {
	// 	// handle error
	// 	fmt.Println(err, "aray ")
	// }
	cfg.Credentials = aws.NewCredentialsCache(stsCredProvider)

	// Create service client value configured for credentials
	// from assumed role.
	svc := kms.NewFromConfig(cfg, func(o *kms.Options) {
		o.Region = "ap-south-1"
	})

	// keyInput := &kms.CreateKeyInput{
	// 	Description: aws.String("This is sample Syymetric key"),
	// 	KeySpec:     types.KeySpecSymmetricDefault,
	// 	KeyUsage:    types.KeyUsageTypeEncryptDecrypt,
	// 	Tags: []types.Tag{
	// 		types.Tag{TagKey: aws.String("Name"), TagValue: aws.String("Unity")},
	// 	},
	// }
	// kmsCreateKeyOutput, err := svc.CreateKey(context.TODO(), keyInput)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("SymetricKey", kmsCreateKeyOutput.KeyMetadata.KeyId)
	// fmt.Println(kmsCreateKeyOutput.KeyMetadata.KeyId)

	// keyInput1 := &kms.CreateKeyInput{
	// 	Description: aws.String("This is sample assymetric  key"),
	// 	KeySpec:     types.KeySpecRsa4096,
	// 	KeyUsage:    types.KeyUsageTypeEncryptDecrypt,
	// 	Tags: []types.Tag{
	// 		types.Tag{TagKey: aws.String("Name"), TagValue: aws.String("Unity1")},
	// 	},
	// }
	// kmsCreateKeyOutput1, err := svc.CreateKey(context.TODO(), keyInput1)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("AssymetricKey", kmsCreateKeyOutput1.KeyMetadata.KeyId)

	keyID := aws.String("def7d8c6-3e3e-41df-8c2a-fa86a29039bb")

	//keyID := aws.String("64578409-321f-4685-bfbb-625075b9ccc6")
	//pub, err := svc.GetPublicKey(context.TODO(), &kms.GetPublicKeyInput{KeyId: keyID})
	// fmt.Println("Print Public key")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//fmt.Println(string(pub.PublicKey[:]))
	//fmt.Printf("%v %T", pub.PublicKey, pub.PublicKey)
	// sEnc, err := b64.StdEncoding.DecodeString(string(pub.PublicKey[:]))
	// fmt.Println(sEnc)
	// fmt.Println("End")
	plainTextToEncrypt := []byte("This is Rohits Secrte")
	generateDataKeyInput := &kms.GenerateDataKeyInput{
		KeyId:   keyID,
		KeySpec: types.DataKeySpecAes256,
	}
	GenerateDataKeyOutput, err := svc.GenerateDataKey(context.TODO(), generateDataKeyInput)
	if err != nil {
		log.Fatal(err.Error())
	}

	block, err := aes.NewCipher(GenerateDataKeyOutput.Plaintext)
	if err != nil {
		log.Fatal(err.Error())
	}

	iv := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("%x\n", iv)
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Fatal(err.Error())
	}
	ciphertext := aesgcm.Seal(nil, iv, plainTextToEncrypt, nil)
	encryptedText := base64.StdEncoding.EncodeToString(ciphertext)
	fmt.Println("Printed Data")
	fmt.Println(encryptedText)

	// decryption

	ciphertext1, err := base64.StdEncoding.DecodeString(encryptedText)

	if err != nil {
		log.Fatal(err.Error())
	}

	block, err = aes.NewCipher(GenerateDataKeyOutput.Plaintext)
	if err != nil {
		panic(err.Error())
	}

	aesgcm, err = cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	plaintext, err := aesgcm.Open(nil, iv, ciphertext1, nil)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Decrypted using datakeyPair")
	fmt.Printf("%s", ciphertext1)
	fmt.Printf("%s", plaintext)

	// encrypt := kms.EncryptInput{KeyId: keyID, Plaintext: plainTextToEncrypt, EncryptionAlgorithm: types.EncryptionAlgorithmSpecSymmetricDefault}
	// encryptOut, err := svc.Encrypt(context.TODO(), &encrypt)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Entypted Data")
	// fmt.Println(encryptOut.CiphertextBlob)

	// cipherText := encryptOut.CiphertextBlob

	// var decryptParam kms.DecryptInput = kms.DecryptInput{CiphertextBlob: GenerateDataKeyOutput.CiphertextBlob, EncryptionAlgorithm: types.EncryptionAlgorithmSpecSymmetricDefault, KeyId: keyID}

	// decryptData, err := svc.Decrypt(context.TODO(), &decryptParam)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(string(decryptData.Plaintext))

	// ciphertext2, err := base64.StdEncoding.DecodeString(encryptedText)

	// if err != nil {
	// 	log.Fatal(err.Error())
	// }

	// block, err = aes.NewCipher(decryptData.Plaintext)
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }

	// if len(ciphertext2) < aes.BlockSize {
	// 	log.Fatal(err.Error())
	// }

	// iv = ciphertext2[:aes.BlockSize]
	// ciphertext2 = ciphertext2[aes.BlockSize:]
	// stream = cipher.NewCFBDecrypter(block, iv)

	// stream.XORKeyStream(ciphertext2, ciphertext2)
	// fmt.Println("Decrypted using datakeyPair from aws")
	// fmt.Printf("%s", ciphertext2)
	// var decryptParam kms.DecryptInput = kms.DecryptInput{CiphertextBlob: cipherText, EncryptionAlgorithm: types.EncryptionAlgorithmSpecSymmetricDefault, KeyId: keyID}
	// decryptData, err := svc.Decrypt(context.TODO(), &decryptParam)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("decrypted Data")
	// fmt.Println(string(decryptData.Plaintext))

	// Get the first page of results for ListObjectsV2 for a bucket
	// var limit int32 = 10
	// output, err := svc.ListKeys(context.TODO(), &kms.ListKeysInput{Limit: &limit})
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// //fmt.Println(output)
	// // log.Println("first page results:")
	// for _, object := range output.Keys {

	// 	fmt.Println(*object.KeyId, *object.KeyArn)
	// }

	// create key
	// result, err := svc.CreateKey(context.TODO(),
	// 	&kms.CreateKeyInput{Description: aws.String("Rohit first key creation"),
	// 		KeySpec:  types.KeySpecRsa2048,
	// 		KeyUsage: types.KeyUsageTypeEncryptDecrypt,
	// 		Tags: []types.Tag{
	// 			{
	// 				TagKey:   aws.String("AuthorC"),
	// 				TagValue: aws.String("RohitTiwariBhai"),
	// 			},
	// 		},
	// 	},
	// )

	// if err != nil {
	// 	fmt.Println("Got error creating key:")
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println("The Key Was created Bro")
	// fmt.Println(*result.KeyMetadata.KeyId)
}
