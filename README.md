# tf

```bash

docker build -t go-tens

docker run -it -p 6006:6006 go-tens /bin/bash

cd /go/src/github.com/tensorflow/tensorflow/tensorflow/examples/speech_commands

python train.py \
--data_dir=/go/src/github.com/wwgberlin/tf/tmp/speech_dataset/ \
--summaries_dir=/go/src/github.com/wwgberlin/tf/tmp/retrain_logs \
--train_dir=/go/src/github.com/wwgberlin/tf/tmp/speech_commands_train \

python freeze.py \
--start_checkpoint=/go/src/github.com/wwgberlin/tf/tmp/speech_commands_train/conv.ckpt-1000 \
--output_file=/go/src/github.com/wwgberlin/tf/tmp/my_frozen_graph.pb

python label_wav.py \
--graph=/go/src/github.com/wwgberlin/tf/tmp/my_frozen_graph.pb \
--labels=/go/src/github.com/wwgberlin/tf/tmp/speech_commands_train/conv_labels.txt \
--wav=/go/src/github.com/wwgberlin/tf/tmp/speech_dataset/go/9fac5701_nohash_2.wav

```
