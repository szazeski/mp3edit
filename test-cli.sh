#!/usr/bin/env bash

rm mp3edit > /dev/null 2>&1 # if it exists

set -e # exit on any errors
#set -x # debug logging
set -u # treat unset variables as errors

trap "git checkout sample-file.mp3" EXIT

git checkout sample-file.mp3

go build .

./mp3edit -version

echo ""
echo "Starting File Values:"
./mp3edit sample-file.mp3

TEST_INDEX=1
TEST_TOTAL=3

function testecho {
  echo ""$TEST_INDEX / $TEST_TOTAL - $1""
}

testecho "Testing title change"
./mp3edit -title="New Title" sample-file.mp3
./mp3edit sample-file.mp3 | grep "New Title" || echo "FAIL"

testecho "Testing artist change"
./mp3edit -artist="New Artist" sample-file.mp3
./mp3edit sample-file.mp3 | grep "New Artist" || echo "FAIL"

testecho "testing album change"
./mp3edit -album="New Album" sample-file.mp3
./mp3edit sample-file.mp3 | grep "New Album" || echo "FAIL"

echo "[PASS]"
./mp3edit sample-file.mp3
