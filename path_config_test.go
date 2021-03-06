package ibmcloudauth

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/vault/sdk/logical"
)

func TestConfig_Write(t *testing.T) {
	b, s := testBackend(t)

	configData := map[string]interface{}{}
	if err := testConfigCreate(t, b, s, configData); err == nil {
		t.Fatal("expected error")
	}

	configData = map[string]interface{}{
		apiKeyField:    "theAPIKey",
		accountIDField: "theAccount",
	}
	if err := testConfigCreate(t, b, s, configData); err != nil {
		t.Fatalf("err: %v", err)
	}

	resp, err := b.HandleRequest(context.Background(), &logical.Request{
		Operation: logical.ReadOperation,
		Path:      "config",
		Storage:   s,
	})
	if err != nil {
		t.Fatal(err)
	}

	if resp != nil {
		keyVal, ok := resp.Data[apiKeyField]
		if !ok {
			t.Fatal("the api_key field was not found in the read config")
		}
		if keyVal != redacted {
			t.Fatal("the api_key value was not redacted")
		}
		_, ok = resp.Data[accountIDField]
		if !ok {
			t.Fatal("the account_id field was not found in the read config")
		}

	} else {
		t.Fatal("did not get a response from the read post-create")
	}
}

func TestConfigDelete(t *testing.T) {
	b, s := testBackend(t)

	configData := map[string]interface{}{
		apiKeyField:    "theAPIKey",
		accountIDField: "theAccount",
	}

	if err := testConfigCreate(t, b, s, configData); err != nil {
		t.Fatalf("err: %v", err)
	}

	resp, err := b.HandleRequest(context.Background(), &logical.Request{
		Operation: logical.DeleteOperation,
		Path:      fmt.Sprintf("config"),
		Storage:   s,
	})
	if err != nil {
		t.Fatal(err)
	}

	resp, err = b.HandleRequest(context.Background(), &logical.Request{
		Operation: logical.ReadOperation,
		Path:      "config",
		Storage:   s,
	})
	if err != nil {
		t.Fatal(err)
	}

	if resp != nil {
		t.Fatal("expected nil config after delete")
	}
}

func testConfigCreate(t *testing.T, b *ibmCloudAuthBackend, s logical.Storage, d map[string]interface{}) error {
	resp, err := b.HandleRequest(context.Background(), &logical.Request{
		Operation: logical.CreateOperation,
		Path:      fmt.Sprintf("config"),
		Data:      d,
		Storage:   s,
	})
	if err != nil {
		return err
	}
	if resp != nil && resp.IsError() {
		return resp.Error()
	}
	return nil
}
