package main

import (
	tf "github.com/tensorflow/tensorflow/tensorflow/go"
	"io/ioutil"
	"os"
	"bufio"
	"fmt"
	"errors"
	"log"
	"reflect"
)

const (
	graphFilePath  = "./tmp/speech_recognition_graph.pb"
	wavFilePath    = "./tmp/speech_dataset/go/9fac5701_nohash_2.wav"
	labelsFilePath = "./tmp/speech_commands_train/conv_labels.txt"
)

func main() {
	wavData, err := readWavDataFromFile(wavFilePath)
	if err != nil {
		log.Fatal(err)
	}
	labels, err := readLabelsFromFile(labelsFilePath)
	if err != nil {
		log.Fatal(err)
	}

	graph, err := importGraph(graphFilePath)
	if err != nil {
		log.Fatal(err)
	}
	res, err := runGraph(graph, wavData, labels)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}

// Read the entire wav file using ioutil.ReadFile (https://golang.org/pkg/io/ioutil/#ReadFile)
// return the bytes slice and error
func readWavDataFromFile(wavFilePath string) ([]byte, error) {
	return ioutil.ReadFile(wavFilePath)
}

// Open the labels file using os.Open (https://golang.org/pkg/os/#Open)
// return an error if necessary.
// Pass the returned file to bufio.NewScanner to read the lines (https://golang.org/pkg/bufio/#Scanner
// --checkout the lines example) return the error is necessary.
// don't forget to close the file.
func readLabelsFromFile(labelsFilePath string) ([]string, error) {
	lfile, err := os.Open(labelsFilePath)
	if err != nil {
		return nil, err
	}
	defer lfile.Close()

	//https://golang.org/pkg/bufio/#example_Scanner_lines
	s := bufio.NewScanner(lfile)
	var labels []string
	for s.Scan() {
		labels = append(labels, s.Text())
	}
	return labels, s.Err()
}

// Read the entire graph file using ioutil.ReadFile  (https://golang.org/pkg/io/ioutil/#ReadFile)
// What is a graph? https://www.tensorflow.org/api_docs/python/tf/Graph
// instantiate a new Graph using tf.NewGraph (https://godoc.org/github.com/tensorflow/tensorflow/tensorflow/go#NewGraph)
// Import the bytes read to the graph using the method
// Import with empty prefix (https://godoc.org/github.com/tensorflow/tensorflow/tensorflow/go#Graph.Import)
func importGraph(graphFilePath string) (*tf.Graph, error) {
	return nil, errors.New("not implemented")
}

//1, Create a new session using tf.NewSession (https://godoc.org/github.com/tensorflow/tensorflow/tensorflow/go#Session)
//   What is a TensorFlow session? Read here https://danijar.com/what-is-a-tensorflow-session/
//
//2. Define a tensor from our "stringified" wavData (https://godoc.org/github.com/tensorflow/tensorflow/tensorflow/go#NewTensor)
//   What's a Tensor? https://www.tensorflow.org/api_docs/python/tf/Tensor.
//	 Let's name it simply "tensor". Don't forget to return an error if necessary.
//
//3. Using the outputOperationName get the Operation from the graph (https://godoc.org/github.com/tensorflow/tensorflow/tensorflow/go#Graph.Operation)
//   What's an Operation? https://www.tensorflow.org/api_docs/python/tf/Operation
//   From the operation let's take the 0th output (https://godoc.org/github.com/tensorflow/tensorflow/tensorflow/go#Operation.Output)
//   let's call our result "softmaxTensor".
//
//4. Define another tenstor from graph operation with output(similar to 4) this time using the inputOperationName this time
//   (and the 0th output), let's name our result "inputOperation".
//
//5. Define a map from an Output type to a *Tensor type and instantiate it with one key value pair:
//   our inputOperation variable and our "tensor" variable (the one with the wavData)
//
//6. Run the session (https://godoc.org/github.com/tensorflow/tensorflow/tensorflow/go#Session.Run)
//   with the map we've instantiated above and an Output slice containing exactly one item - our softmaxTensor.
//   and use fmtOutput to format the result.
//   Return the respective error if necessary.
//
//7. Run the tests.
func runGraph(graph *tf.Graph, wavData []byte, labels []string) (string, error) {
	const (
		inputOperationName  = "wav_data"       //Name of WAVE data input node in model.
		outputOperationName = "labels_softmax" //Name of node outputting a prediction in the model
	)

	//YOUR CODE GOES HERE!

	// Uncomment this code when the output variable is defined:
	return "", nil
}

func fmtOutput(output []*tf.Tensor, labels []string) string {
	var str string
	for i := 0; i < len(output); i++ {
		if reflect.TypeOf(output[i].Value()).Kind() == reflect.Slice {
			s := reflect.ValueOf(output[i].Value())

			for j := 0; j < s.Len(); j++ {
				r := s.Index(j)

				if r.Kind() == reflect.Slice {
					for k := 0; k < r.Len(); k++ {
						q := r.Index(k)
						humanString := labels[k] + ":\t"
						str += fmt.Sprint(humanString, q, "\n")
					}
				}
			}
		}
	}
	return str
}
