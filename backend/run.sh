SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
cd ${SCRIPT_DIR}/..

cmd = "sh ./backend/auth/run.sh \
& sh ./backend/user/run.sh \
& sh ./backend/post/run.sh \
& sh ./backend/post-processor/run.sh \
& sh ./backend/api/run.sh"

if ! ($cmd) ; then
    echo "Fail: Running backend"
    exit
fi

echo "Success: Running backend"
