package lorem

import (
	"errors"
	"io/ioutil"
	"strings"
	. "testing"
)

func TestOneLorem(t *T) {
	buf, err := ioutil.ReadAll(NewLoremN(len(Lorem)))
	if err != nil {
		t.Error(err)
	}

	if string(buf) != Lorem {
		t.Error(ErrNotALorem)
	}
}

func TestOneLoremIsLorem(t *T) {
	if !IsLorem([]byte(Lorem)) {
		t.Error(ErrNotALorem)
	}
}

func TestSomeLoremIsLorem(t *T) {
	ok, err := IsLoremReader(NewLoremN(3141592))
	if err != nil {
		t.Error("unexpected error", err)
	}
	if !ok {
		t.Error(ErrNotALorem)
	}
}

func TestSomeMerolIsNotALorem(t *T) {
	var merol string
	for _, r := range Lorem {
		merol = string(r) + merol
	}

	r := strings.NewReader(strings.Repeat(merol, 3))

	ok, err := IsLoremReader(r)
	if err != nil {
		t.Error("unexpected error", err)
	}
	if ok {
		t.Error(errors.New("Merol must not be lorem"))
	}
}
