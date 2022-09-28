GOLANG_DIR=.
TYPESCRIPT_DIR=../../frontend/app/src/proto

mkdir -p $GOLANG_DIR
mkdir -p $TYPESCRIPT_DIR

cmd="protoc helloWorld.proto \
--plugin=protoc-gen-ts=../../frontend/app/node_modules/.bin/protoc-gen-ts \
--ts_out=service=grpc-web:$TYPESCRIPT_DIR \
-I. \
--go_out=$GOLANG_DIR --go-grpc_out=$GOLANG_DIR "


if ! eval "$cmd" ; then
    echo "Fail generating grpc"
    exit
fi

echo "Success generating grpc"
# eval "$cmd"

# echo "${cmd}"

