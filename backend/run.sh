SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
cd ${SCRIPT_DIR}/..

if ! (sh ./backend/api/run.sh & sh ./backend/auth/run.sh) ; then
    echo "Fail: Running backend"
    exit
fi

echo "Success: Running backend"
