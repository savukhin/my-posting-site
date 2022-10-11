SCRIPT_DIR=$( cd -- "$( dirname -- "${0}" )" &> /dev/null && pwd )
cd ${SCRIPT_DIR}/..

cmd="sh ./backend/auth/run.sh \
& sh ./backend/user/run.sh \
& sh ./backend/post/run.sh \
& sh ./backend/post-processor/run.sh \
& sh ./backend/api/run.sh"
# echo backend SCRIPT_DIR is $SCRIPT_DIR

# cmd="sh ./backend/auth/run.sh"

if ! eval $cmd ; then
    echo "Fail: Running backend"
    exit
fi

echo "Success: Running backend"
