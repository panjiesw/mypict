package model_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"panjiesw.com/mypict/server/model"
)

func TestNull(t *testing.T) {
	var tnull model.ImageS
	tjsnull := []byte(`{"id":"a","title":null}`)

	var tundef model.ImageS
	tjsundef := []byte(`{"id":"b"}`)

	var gnull model.ImageS
	gjsnull := []byte(`{"id":"a","gallery":"null"}`)

	var gundef model.ImageS
	gjsundef := []byte(`{"id":"b"}`)

	if err := json.Unmarshal(tjsnull, &tnull); err != nil {
		t.Fatal(err)
	}

	if err := json.Unmarshal(tjsundef, &tundef); err != nil {
		t.Fatal(err)
	}

	if err := json.Unmarshal(gjsnull, &gnull); err != nil {
		t.Fatal(err)
	}

	if err := json.Unmarshal(gjsundef, &gundef); err != nil {
		t.Fatal(err)
	}

	fmt.Printf("tnull %v\n", tnull)
	fmt.Printf("tundef %v\n", tundef)
	fmt.Printf("gnull %v\n", gnull)
	fmt.Printf("gundef %v\n", gundef)

	if btnull, err := json.Marshal(tnull); err != nil {
		t.Fatal(err)
	} else {
		fmt.Printf("tjsonnull %s\n", btnull)
	}

	if bgnull, err := json.Marshal(gnull); err != nil {
		t.Fatal(err)
	} else {
		fmt.Printf("gjsonnull %s\n", bgnull)
	}
}
