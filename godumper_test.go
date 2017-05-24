package godumper

import (
	"math/rand"
	"testing"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func GenerateRandomID() int {
	return 100000 + r.Intn(999999-100000)
}

type Message struct {
	Id       int `dump:"index"`
	UserFrom int
	UserTo   int
	Text     string
	Time     int64
	IsReaded bool
	ByteTest byte
}

var slice = []Message{
	Message{
		Id:       GenerateRandomID(),
		Time:     time.Now().Unix() + int64(GenerateRandomID()),
		Text:     "Hi dude, Are you okay ?",
		UserFrom: GenerateRandomID(),
		UserTo:   GenerateRandomID(),
		IsReaded: true,
		ByteTest: byte(1),
	},
	Message{
		Id:       GenerateRandomID(),
		UserFrom: GenerateRandomID(),
		UserTo:   GenerateRandomID(),
		Text:     "Hey buddy,I'm fine",
		Time:     time.Now().Unix() + int64(GenerateRandomID()),
		IsReaded: false,
		ByteTest: byte(2),
	},
	Message{
		Id:       GenerateRandomID(),
		UserFrom: GenerateRandomID(),
		UserTo:   GenerateRandomID(),
		Text:     "😆😆😆 😛",
		Time:     time.Now().Unix() + int64(GenerateRandomID()),
		IsReaded: true,
		ByteTest: byte(3),
	},
}

var array = [2]Message{
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
}

func TestNewDumper(t *testing.T) {
	dumper, err := New(Message{})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(dumper.Header)
}

func TestDumpSlice(t *testing.T) {
	dumper, err := New(Message{})
	if err != nil {
		t.Fatal(err)
	}

	err = dumper.Dump(slice)
	if err != nil {
		t.Fatal(err)
	}

	if len(dumper.Body) == 0 {
		t.Fatal("Empty body")
	}

	t.Log(dumper.Body)
}

func TestDumpArray(t *testing.T) {
	dumper, err := New(Message{})
	if err != nil {
		t.Fatal(err)
	}

	err = dumper.Dump(array)
	if err != nil {
		t.Fatal(err)
	}

	if len(dumper.Body) == 0 {
		t.Fatal("Empty body")
	}

	t.Log(dumper.Body)
}

func TestSave(t *testing.T) {
	dumper, err := New(Message{})
	if err != nil {
		t.Fatal(err)
	}

	err = dumper.Dump(slice)
	if err != nil {
		t.Fatal(err)
	}
	err = dumper.Save("result.csv")
	if err != nil {
		t.Fatal(err)
	}
}

func TestDumpAndSave(t *testing.T) {
	dumper, err := New(Message{})
	if err != nil {
		t.Fatal(err)
	}

	err = dumper.DumpAndSave(slice, "dumpAndSave.csv")
	if err != nil {
		t.Fatal(err)
	}
}
