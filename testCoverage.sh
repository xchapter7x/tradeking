#!/bin/bash
WATERMARK=90;
COVEROUT=`go test -cover ./src/...`;
arrIN=${COVEROUT/// };
COVERLEVEL=`echo ${arrIN[4]} | tr "%" " "`;
COVERLEVEL=`awk -v vpct="$COVERLEVEL" 'BEGIN{print vpct * 1}'`;
echo "Coverage - ${COVERLEVEL}%";
echo "Watermark- ${WATERMARK}%";
if [[ $COVERLEVEL -lt $WATERMARK ]]; then exit 1; fi
