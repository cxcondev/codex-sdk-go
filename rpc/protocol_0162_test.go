package rpc

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/pmenglund/codex-sdk-go/protocol"
)

func TestProtocol0162AttachmentOwnerLookup(t *testing.T) {
	transport := newScriptedTransport()
	client := NewClient(transport, ClientOptions{})
	defer client.Close()
	transport.enqueueResult(json.RawMessage(`{"data":[{"threadId":"thread","archived":false}],"nextCursor":"next"}`))
	archived := false
	page, err := client.ThreadAttachmentOwnerList(context.Background(), protocol.ThreadAttachmentOwnerListParams{
		AttachmentType: "link", IdentityKey: "identity", Archived: &archived,
	})
	if err != nil {
		t.Fatal(err)
	}
	if page == nil || len(page.Data) != 1 || page.Data[0].ThreadID != "thread" || page.Data[0].Archived || page.NextCursor == nil || *page.NextCursor != "next" {
		t.Fatalf("owner page: %#v", page)
	}
	requests := transport.writtenRequests()
	if len(requests) != 1 || requests[0].Method != "thread/attachmentOwner/list" || string(requests[0].Params) != `{"archived":false,"attachmentType":"link","identityKey":"identity"}` {
		t.Fatalf("owner lookup requests: %#v", requests)
	}
}

func TestProtocol0162PredictionNotification(t *testing.T) {
	for _, result := range []string{`{"type":"completed","text":"next"}`, `{"type":"failed"}`, `{"type":"future","count":9007199254740993}`} {
		payload := `{"threadId":"thread","sourceTurnId":"turn","result":` + result + `}`
		note, err := parseServerNotification("thread/prediction/updated", json.RawMessage(payload))
		if err != nil {
			t.Fatal(err)
		}
		prediction, ok := note.Params.(protocol.ThreadPredictionUpdatedNotification)
		if !ok || prediction.ThreadID != "thread" || prediction.SourceTurnID != "turn" {
			t.Fatalf("prediction notification: %#v", note)
		}
		encoded, err := json.Marshal(prediction.Result)
		if err != nil || string(encoded) != result {
			t.Fatalf("prediction result changed: %s, %v", encoded, err)
		}
		if prediction.Result.IsKnown() != (string(prediction.Result.Kind()) != "future") {
			t.Fatalf("incorrect known variant: %s", prediction.Result.Kind())
		}
	}
	if _, err := parseServerNotification("thread/prediction/updated", json.RawMessage(`{"threadId":"thread","sourceTurnId":"turn","result":{}}`)); err == nil {
		t.Fatal("accepted prediction without discriminator")
	}
}

func TestProtocol0162ErrorNotificationPreservesFuturePayload(t *testing.T) {
	info := `{"futureError":{"count":9007199254740993}}`
	note, err := parseServerNotification("error", json.RawMessage(`{"threadId":"thread","turnId":"turn","willRetry":false,"error":{"message":"error","codexErrorInfo":`+info+`}}`))
	if err != nil {
		t.Fatal(err)
	}
	value, ok := note.Params.(protocol.ErrorNotification)
	if !ok || string(value.Error.CodexErrorInfo) != info {
		t.Fatalf("error notification lost payload: %#v", note)
	}
}
