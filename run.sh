#!/bin/bash

for i in 1 2 3; do

  docker run --rm jst-time-test$i

done
