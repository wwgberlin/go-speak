# go-speak

Speech recognition challenge with TensorFlow and Go.

Your challenge, should you choose to accept it is to use the TensorFlow Go package to import a trained model and apply it to a wav file.

### Prerequisites:
- [Docker](https://www.docker.com/)
- optional: [Go](https://golang.org/doc/install) - we will have go running inside the docker container

**We will _NOT_ be installing TensorFlow locally in this exercise but use a docker container instead.* 

#### Setup
In your terminal (from inside the root directory of this repository):
```bash
docker-compose run speech_recognition /bin/bash  
```
Your terminal should now be running bash from inside the docker container we've setup for this task.

#### Train the model
To train our model we will be following the first part of the [TensorFlow tutorial for audio recognition](https://missinglink.ai/guides/tensorflow/tensorflow-speech-recognition-two-quick-tutorials/).

The Python scripts are already inside the container.
Run: 

```bash
cd /go/src/github.com/tensorflow/tensorflow/tensorflow/examples/speech_commands
python train.py \
--data_dir=$APP_DIR/tmp/speech_dataset/ \
--summaries_dir=$APP_DIR/tmp/retrain_logs \
--train_dir=$APP_DIR/tmp/speech_commands_train
```

If the training stopped in the middle you can check your 
speech_commands_train directory for the latest checkpoint and 
rerun the train.py script with `--start_checkpoint=$APP_DIR/tmp/speech_commands_train/conv.ckpt-[last-checkpoint-number]`
 
####Save the model (graph)
```bash
python freeze.py \
--start_checkpoint=$APP_DIR/tmp/speech_commands_train/conv.ckpt-1000 \
--output_file=$APP_DIR/tmp/speech_recognition_graph.pb
```

####Run the (failing) tests
```bash
cd $APP_DIR
go test
```
Implement the code until tests pass


####Run your code to see the results
```bash
go run main.go
```
