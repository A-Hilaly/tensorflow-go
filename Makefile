init:
	git clone https://github.com/tensorflow/serving && git checkout 2.0.0
	git clone https://github.com/tensorflow/tensorflow && git checkout v2.0.0


TF_ORG_PATH=~/source/github.com/tensorflow

gen:
	protoc -I $(TF_ORG_PATH)/serving \
		-I $(TF_ORG_PATH)/tensorflow \
		$(TF_ORG_PATH)/serving/tensorflow_serving/apis/*.proto \
		--go_out=plugins=grpc:vendor

	protoc -I $(TF_ORG_PATH)/serving \
		-I $(TF_ORG_PATH)/tensorflow \
		$(TF_ORG_PATH)/serving/tensorflow_serving/config/*.proto \
		--go_out=plugins=grpc:vendor

	protoc -I $(TF_ORG_PATH)/serving \
		-I $(TF_ORG_PATH)/tensorflow \
		$(TF_ORG_PATH)/serving/tensorflow_serving/util/*.proto \
		--go_out=plugins=grpc:vendor

	protoc -I $(TF_ORG_PATH)/serving \
		-I $(TF_ORG_PATH)/tensorflow \
		$(TF_ORG_PATH)/serving/tensorflow_serving/sources/storage_Path/*.proto \
		--go_out=plugins=grpc:vendor

	protoc -I $(TF_ORG_PATH)/serving \
		-I $(TF_ORG_PATH)/tensorflow \
		$(TF_ORG_PATH)/tensorflow/tensorflow/core/framework/*.proto \
		--go_out=plugins=grpc:vendor

all: init gen