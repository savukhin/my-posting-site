if ! (sh ./api/run.sh & sh ./auth/run.sh) ; then
    echo "Fail: Running backend"
    exit
fi

echo "Success: Running backend"
