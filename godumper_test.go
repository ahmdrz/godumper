package godumper

import (
	"math/rand"
	"testing"
	"time"
)

func GenerateRandomID() int {
	rand.Seed(time.Now().Unix())
	return 100000 + rand.Intn(999999-100000)
}

type Message struct {
	Id       int
	UserFrom int
	UserTo   int
	Text     string
	Time     int64
	IsReaded bool
}

var array = []Message{
	Message{
		Id:       GenerateRandomID(),
		Time:     time.Now().Unix() + int64(GenerateRandomID()),
		Text:     "Hi dude, Are you okay ?",
		UserFrom: GenerateRandomID(),
		UserTo:   GenerateRandomID(),
		IsReaded: true,
	},
	Message{
		Id:       GenerateRandomID(),
		UserFrom: GenerateRandomID(),
		UserTo:   GenerateRandomID(),
		Text:     "Hey buddy,I'm fine",
		Time:     time.Now().Unix() + int64(GenerateRandomID()),
		IsReaded: false,
	},
	Message{
		Id:       GenerateRandomID(),
		UserFrom: GenerateRandomID(),
		UserTo:   GenerateRandomID(),
		Text:     "😆😆😆 😛",
		Time:     time.Now().Unix() + int64(GenerateRandomID()),
		IsReaded: true,
	},
}

func TestNewDumper(t *testing.T) {
	dumper, err := New(Message{})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(dumper.Header)
}

func TestDump(t *testing.T) {
	dumper, err := New(Message{})
	if err != nil {
		t.Fatal(err)
	}

	dumper, err = dumper.Dump(array)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSave(t *testing.T) {
	dumper, err := New(Message{})
	if err != nil {
		t.Fatal(err)
	}

	dumper, err = dumper.Dump(array)
	if err != nil {
		t.Fatal(err)
	}
	err = dumper.Save("result.csv")
	if err != nil {
		t.Fatal(err)
	}
}
