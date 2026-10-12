package protocol

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestCodex0162TurnAttributionRoundTrip(t *testing.T) {
	for _, payload := range []string{
		`{"threadId":"thread","input":[],"parentTurnId":"parent","rootTurnId":"root"}`,
		`{"threadId":"thread","input":[],"parentTurnId":null,"rootTurnId":null}`,
		`{"threadId":"thread","input":[]}`,
	} {
		var params TurnStartParams
		if err := json.Unmarshal([]byte(payload), &params); err != nil {
			t.Fatal(err)
		}
		check0162Fields(t, payload, params, "parentTurnId", "rootTurnId")
	}
	var turn Turn
	payload := `{"id":"turn","items":[],"status":"completed","rootTurnId":"root"}`
	if err := json.Unmarshal([]byte(payload), &turn); err != nil {
		t.Fatal(err)
	}
	check0162Fields(t, payload, turn, "rootTurnId")
}

func TestCodex0162GoalOriginRoundTrip(t *testing.T) {
	for _, origin := range []string{`"user"`, `"automatic"`, `null`, ""} {
		payload := `{"threadId":"thread"`
		if origin != "" {
			payload += `,"origin":` + origin
		}
		payload += `}`
		var set ThreadGoalSetParams
		var clear ThreadGoalClearParams
		for _, value := range []any{&set, &clear} {
			if err := json.Unmarshal([]byte(payload), value); err != nil {
				t.Fatal(err)
			}
			check0162Fields(t, payload, value, "origin")
		}
	}
}

func TestCodex0162AdditionalToolsResponseItem(t *testing.T) {
	payload := []byte(`{"type":"additional_tools","role":"system","tools":[{"id":9007199254740993}],"future":9007199254740993}`)
	var item ResponseItem
	if err := json.Unmarshal(payload, &item); err != nil {
		t.Fatal(err)
	}
	if string(item.Kind()) != "additional_tools" || !item.IsKnown() {
		t.Fatalf("additional_tools not recognized: %s", item.Kind())
	}
	encoded, err := json.Marshal(item)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(encoded, payload) {
		t.Fatalf("tool payload changed: %s", encoded)
	}
	for _, invalid := range []string{`{"type":"additional_tools","tools":[]}`, `{"type":"additional_tools","role":"system"}`} {
		if err := json.Unmarshal([]byte(invalid), &item); err == nil {
			t.Fatalf("accepted incomplete tools item: %s", invalid)
		}
	}
}

func TestCodex0162ErrorInfoPreservesFuturePayload(t *testing.T) {
	for _, info := range []string{`"future_error"`, `{"futureError":{"count":9007199254740993}}`} {
		payload := `{"message":"error","codexErrorInfo":` + info + `}`
		var turnError TurnError
		if err := json.Unmarshal([]byte(payload), &turnError); err != nil {
			t.Fatal(err)
		}
		check0162Fields(t, payload, turnError, "codexErrorInfo")
	}
}

func check0162Fields(t *testing.T, payload string, value any, fields ...string) {
	t.Helper()
	encoded, err := json.Marshal(value)
	if err != nil {
		t.Fatal(err)
	}
	var before, after map[string]json.RawMessage
	if err := json.Unmarshal([]byte(payload), &before); err != nil {
		t.Fatal(err)
	}
	if err := json.Unmarshal(encoded, &after); err != nil {
		t.Fatal(err)
	}
	for _, field := range fields {
		// Optional null and absent values both encode as absent; neither supplies provenance.
		want := before[field]
		if bytes.Equal(want, []byte("null")) {
			want = nil
		}
		if !bytes.Equal(want, after[field]) {
			t.Errorf("%s: got %s, want %s (payload %s)", field, after[field], want, encoded)
		}
	}
}
