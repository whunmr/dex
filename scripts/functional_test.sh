#!/usr/bin/env bash

set -e

mkdir func_test

echo "$DPW" | docker login -u "$DUN" --password-stdin
docker pull matrixport/walle
docker run --rm -v $(pwd)/func_test:/test:Z matrixport/walle /data/script/cp_data.sh

mkdir func_test/run
cd func_test
bash script/init.sh

echo "Test begin"

pipenv run behave ./features/

echo "Test end"