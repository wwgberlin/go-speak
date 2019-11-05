package main

import (
	tf "github.com/tensorflow/tensorflow/tensorflow/go"
	"io/ioutil"
	"os"
	"bufio"
	"fmt"
	"reflect"
)

var (
	graph  *tf.Graph
	labels []string
)

func main() {
	labelWav("/tmp/speech_dataset/go/9fac5701_nohash_2.wav",
		"/tmp/speech_commands_train/conv_labels.txt",
		"/tmp/my_frozen_graph.pb")
}

func labelWav(waveFile string, labelsFile string, graphFile string) {
	wavData, err := ioutil.ReadFile(waveFile)
	if err != nil {
		panic(err)
	}

	lfile, err := os.Open(labelsFile)
	if err != nil {
		panic(err)
	}

	//https://golang.org/pkg/bufio/#example_Scanner_lines
	s := bufio.NewScanner(lfile)
	var labels []string
	for s.Scan() {
		labels = append(labels, s.Text())
	}

	graph, err = loadGraph(graphFile)
	if err != nil {
		panic(err)
	}
	runGraph(wavData, labels)

}
func loadGraph(graphFile string) (*tf.Graph, error) {
	// Load inception model
	model, err := ioutil.ReadFile(graphFile)
	if err != nil {
		return nil, err
	}
	graph := tf.NewGraph()
	if err := graph.Import(model, ""); err != nil {
		return nil, err
	}
	return graph, nil
}

func runGraph(wavData []byte, labels []string) {
	const (
		inputName  = "wav_data"       //Name of WAVE data input node in model.
		outputName = "labels_softmax" //Name of node outputting a prediction in the model
		numLables  = 3
	)
	session, err := tf.NewSession(graph, &tf.SessionOptions{})
	if err != nil {
		panic(err)
	}
	defer session.Close()

	//tensor := graph.Operation(inputName)
	softmaxTensor := graph.Operation(outputName).Output(0)
	tensor, err := tf.NewTensor(string(wavData))
	if err != nil {
		panic(err)
	}

	output, err := session.Run(
		map[tf.Output]*tf.Tensor{
			graph.Operation(inputName).Output(0): tensor,
		}, []tf.Output{softmaxTensor},
		nil)
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(output); i++ {
		if reflect.TypeOf(output[i].Value()).Kind() == reflect.Slice {
			s := reflect.ValueOf(output[i].Value())

			for j := 0; j < s.Len(); j++ {
				r := s.Index(j)

				if r.Kind() == reflect.Slice {
					for k := 0; k < r.Len(); k++ {
						q := r.Index(k)
						humanString := labels[k] + ":\t"
						fmt.Println(humanString, q)
					}
				}
			}
		}
	}

}
