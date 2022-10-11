# RED="\033[0;31m"
# YELLOW="\033[1;33m"
# GREEN="\033[0;32m"
# NC="\033[0m" # No Color

# if ! cd ../../.. ; then
#     echo "${RED}Fail${NC} in cd command"
#     exit
# fi

# SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
# # export GOPATH=$SCRIPT_DIR

# echo "Set local GOPATH to" ${GOPATH}

# if ! cd src/backend/api; then
#     echo "${RED}Fail${NC} in cd command"
#     exit
# fi

# if ! go mod tidy; then
#     echo "${RED}Fail${NC}: go mod tidy"
#     exit
# fi
# echo "${YELLOW}Sucess${NC}: go mod tidy"

# if ! go install; then
#     echo "${RED}Fail${NC}: go install"
#     exit
# fi
# echo "${YELLOW}Sucess${NC}: go install"

# if ! cd ../../.. ; then
#     echo "${RED}Fail${NC} in cd command"
#     exit
# fi

# echo "${GREEN}Running app...${NC}"
# echo 
# $GOPATH/bin/api

SCRIPT_DIR=$( cd -- "$( dirname -- "${0}" )" &> /dev/null && pwd )
cd ${SCRIPT_DIR}/../..

go run backend/api/main.go
