// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package main

import (
	"io/ioutil"
	"os"
	"testing"
)

const cred = `
{
    "tenantId": "72f988bf-86f1-41af-91ab-2d7cd011db47",
    "subscriptionId": "11122233-4444-5555-6666-777888999000",
    "aadClientId": "123",
    "aadClientSecret": "456",
    "resourceGroup": "mykeyvaultrg",
    "location": "eastus",
    "providerVaultName": "k8skv",
	"providerKeyName": "mykey",
	"providerKeyVersion": "bd497c644699475d9fb22dbbc15ba286"
}`

func TestCreateInstance(t *testing.T) {
	file, err := ioutil.TempFile("", "kms_test")
	if err != nil {
		t.Error(err)
	}

	defer os.Remove(file.Name())

	if _, err := file.Write([]byte(cred)); err != nil {
		t.Error(err)
	}

	cred, err := GetAzureAuthConfig(file.Name())
	if err != nil {
		t.Error(err)
	}

	KVTestName, KVTestKeyName, KVTestVersion, RGTest, err := GetKMSProvider(file.Name())
	if err != nil {
		t.Error(err)
	}

	keyManagementServiceServer := new(KeyManagementServiceServer)
	keyManagementServiceServer.pathToUnixSocket = "/tmp/azurekms.socket"

	wanted := "72f988bf-86f1-41af-91ab-2d7cd011db47"
	if cred.TenantID != wanted {
		t.Errorf("Wanted %s, got %s.", wanted, cred.TenantID)
	}

	wanted = "11122233-4444-5555-6666-777888999000"
	if cred.SubscriptionID != wanted {
		t.Errorf("Wanted %s, got %s.", wanted, cred.SubscriptionID)
	}

	wanted = "123"
	if cred.AADClientID != wanted {
		t.Errorf("Wanted %s, got %s.", wanted, cred.AADClientID)
	}

	wanted = "456"
	if cred.AADClientSecret != wanted {
		t.Errorf("Wanted %s, got %s.", wanted, cred.AADClientSecret)
	}

	wanted = "mykeyvaultrg"
	if *RGTest != wanted {
		t.Errorf("Wanted %s, got %s.", wanted, *RGTest)
	}

	wanted = "k8skv"
	if *KVTestName != wanted {
		t.Errorf("Wanted %s, got %s.", wanted, *KVTestName)
	}

	wanted = "mykey"
	if *KVTestKeyName != wanted {
		t.Errorf("Wanted %s, got %s.", wanted, *KVTestKeyName)
	}

	wanted = "bd497c644699475d9fb22dbbc15ba286"
	if *KVTestVersion != wanted {
		t.Errorf("Wanted %s, got %s.", wanted, *KVTestVersion)
	}
}

func TestCreateInstanceNoCredentials(t *testing.T) {
	file, err := ioutil.TempFile("", "kms_test")
	if err != nil {
		t.Error(err)
	}

	fileName := file.Name()

	if err := file.Close(); err != nil {
		t.Error(err)
	}

	os.Remove(fileName)

	if _, err := GetAzureAuthConfig(file.Name()); err == nil {
		t.Fatal("expected to fail with bad json")
	}
}

const badCred = `
{
    "tenantId": "72f988bf-86f1-41af-91ab-2d7cd011db47",
    "subscriptionId": "11122233-4444-5555-6666-777888999000",
    "aadClientId": "123",
    "aadClientSecret": "456",
    "resourceGroup": "mykeyvaultrg",
    "location": "eastus",
    "providerVaultName": "k8skv",
	"providerKeyName": "mykey",
	"providerKeyVersion": "bd497c644699475d9fb22dbbc15ba286",`

func TestCreateInstanceBadCredentials(t *testing.T) {
	file, err := ioutil.TempFile("", "kms_test")
	if err != nil {
		t.Error(err)
	}

	defer os.Remove(file.Name())

	if _, err := file.Write([]byte(badCred)); err != nil {
		t.Error(err)
	}

	if _, err := GetAzureAuthConfig(file.Name()); err == nil {
		t.Fatal("expected to fail with bad json")
	}

}
