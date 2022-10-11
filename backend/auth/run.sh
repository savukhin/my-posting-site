SCRIPT_DIR=$( cd -- "$( dirname -- "${0}" )" &> /dev/null && pwd )
cd ${SCRIPT_DIR}/../..

go run backend/auth/main.go
