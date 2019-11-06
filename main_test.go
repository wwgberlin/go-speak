package main

import (
	"testing"
	"os"
	"reflect"
)

func Test_readWavDataFromFile(t *testing.T) {
	//failure
	b, err := readWavDataFromFile("./fixtures/doesnt-exist.txt")
	if err != nil {
		if !os.IsNotExist(err) {
			t.Fatalf("readWavDataFromFile was expected to return the error returned from ioutil.ReadFile")
		}
	} else {
		t.Fatalf("readWavDataFromFile was expected to fail")
	}
	if b != nil {
		t.Fatalf("readWavDataFromFile was expected return a nil byte slice")
	}

	//success
	b, err = readWavDataFromFile("./fixtures/test.txt")
	if err != nil {
		t.Fatalf("readWavDataFromFile was expected to success, returned error: %s", err)
	}
	if string(b) != "I am a test" {
		t.Fatalf("readWavDataFromFile was expected to return the contents of the given file but instead returned '%v'", b)
	}
}

func Test_readLabelsFromFile(t *testing.T) {
	//failure
	s, err := readLabelsFromFile("./fixtures/doesntexist.txt")
	if err != nil {
		if !os.IsNotExist(err) {
			t.Fatalf("readLabelsFromFile was expected to return the error returned from ioutil.ReadFile")
		}
	} else {
		t.Fatalf("readLabelsFromFile was expected to fail")
	}
	if s != nil {
		t.Fatalf("readLabelsFromFile was expected return a nil string slice")
	}

	//success
	s, err = readLabelsFromFile("./fixtures/labels.txt")
	if err != nil {
		t.Fatal("readLabelsFromFile was expected to succeed")
	}
	if len(s) != 2 {
		t.Fatal("readLabelsFromFile was expected to return a slice containing the lines of the given file")
	}
	if !reflect.DeepEqual(s, []string{"label 1", "label 2"}) {
		t.Fatalf("unexpected return value from readLabelsFromFile: '%v'", s)
	}
}

func Test_importGraph(t *testing.T) {
	//failure
	g, err := importGraph("./fixtures/doesnt_exist.pb")
	if err != nil {
		if !os.IsNotExist(err) {
			t.Fatalf("importGraph was expected to return the error returned from ioutil.ReadFile")
		}
	} else {
		t.Fatalf("importGraph was expected to fail")
	}
	if g != nil {
		t.Fatalf("importGraph was expected return a nil string slice")
	}

	//bad file
	g, err = importGraph("./fixtures/labels.txt")
	if err == nil {
		t.Fatalf("importGraph was expected to fail")
	}
	if g != nil {
		t.Fatalf("importGraph was expected return a nil graph")
	}

	//success
	g, err = importGraph("./fixtures/graph.pb")
	if err != nil {
		t.Fatal("importGraph was expected to succeed")
	}
	if g == nil {
		t.Fatal("importGraph was expected to return a valid Graph object but returned nil")
	}
	if g.Operation("wav_data") == nil {
		t.Fatal("importGraph did not import graph object properly using the file path")
	}

}

func Test_runGraph(t *testing.T) {
	const expected = "_silence_:\t0.0014208348\n" +
		"_unknown_:\t0.0832541\n" +
		"yes:\t0.016642008\n" +
		"no:\t0.083147325\n" +
		"up:\t0.059981856\n" +
		"down:\t0.13438535\n" +
		"left:\t0.054391935\n" +
		"right:\t0.043705765\n" +
		"on:\t0.26739722\n" +
		"off:\t0.10459701\n" +
		"stop:\t0.101431474\n" +
		"go:\t0.049645092\n"

	graph, _ := importGraph("./fixtures/graph.pb")
	wavData, _ := readWavDataFromFile("./tmp/speech_dataset/on/0a9f9af7_nohash_1.wav")
	labels, _ := readLabelsFromFile(labelsFilePath)
	if str, _ := runGraph(graph, wavData, labels); str != expected {
		t.Fatalf("runGraph was expected to return: '%s' but returned: '%s'", expected, str)
	}
}
