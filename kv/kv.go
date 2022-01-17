// https://github.com/Azure/azure-sdk-for-go/tree/main/sdk/keyvault/azkeys
package kv

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/keyvault/azkeys"
	"github.com/Azure/azure-sdk-for-go/sdk/keyvault/azsecrets"
	"github.com/davecgh/go-spew/spew"
)

var KeyVaultName string

func ListKeys(credential azcore.TokenCredential) {
	client, err := azkeys.NewClient(fmt.Sprintf("https://%s.vault.azure.net/", KeyVaultName), credential, nil)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%v", client)

	pager := client.ListKeys(nil)

	for pager.NextPage(context.TODO()) {
		for _, key := range pager.PageResponse().Keys {
			log.Println(*key.KID)
		}
	}

	if pager.Err() != nil {
		panic(pager.Err())
	}

}

func ListSecrets(credential azcore.TokenCredential) {
	client, err := azsecrets.NewClient(fmt.Sprintf("https://%s.vault.azure.net/", KeyVaultName), credential, nil)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%v", client)

	pager := client.ListSecrets(nil)

	for pager.NextPage(context.TODO()) {
		for _, key := range pager.PageResponse().Secrets {
			log.Println(*key.ID)
		}
	}

	if pager.Err() != nil {
		panic(pager.Err())
	}

}

func GetSecrets(credential azcore.TokenCredential, secretName string) ([]byte, error) {
	client, err := azsecrets.NewClient(fmt.Sprintf("https://%s.vault.azure.net/", KeyVaultName), credential, nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%v", client)

	res, err := client.GetSecret(context.Background(), secretName, nil)
	if err != nil {
		return nil, err
	}

	spew.Dump(res.Secret)

	if res.ContentType == nil {
		return nil, fmt.Errorf("content type of the secret (%q) is null. 'application/x-pkcs12' is expected", secretName)
	}
	if *res.ContentType != "application/x-pkcs12" {
		return nil, fmt.Errorf("content type of the secret (%q) is %q. 'application/x-pkcs12' is expected", secretName, *res.ContentType)
	}

	b, err := base64.StdEncoding.DecodeString(*res.Value)
	if err != nil {
		log.Fatal(err)
	}

	return b, nil
}
