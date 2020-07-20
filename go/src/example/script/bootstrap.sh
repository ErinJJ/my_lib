#!/bin/bash

CURDIR=$(cd $(dirname $0); pwd)
if [ "X$1" != "X" ]; then
    RUNTIME_ROOT=$1
else
    RUNTIME_ROOT=${CURDIR}
fi

if [ "X$RUNTIME_ROOT" == "X" ]; then
    echo "There is no RUNTIME_ROOT support."
    echo "Usage: ./bootstrap.sh $RUNTIME_ROOT"
    exit -1
fi

PORT=$2
RUNTIME_CONF_ROOT=$RUNTIME_ROOT/conf
RUNTIME_LOG_ROOT=$RUNTIME_ROOT/log

# 服务日志路径: $RUNTIME_LOG_ROOT/app/${svc_name}.log
if [ ! -d $RUNTIME_LOG_ROOT/app ]; then
    mkdir -p $RUNTIME_LOG_ROOT/app
fi

# RPC日志路径：$RUNTIME_LOG_ROOT/rpc/go_rpc.log
if [ ! -d $RUNTIME_LOG_ROOT/rpc ]; then
    mkdir -p $RUNTIME_LOG_ROOT/rpc
fi

if [ ! -f $CURDIR/settings.py ]; then
    echo "there is no settings.py exist."
    exit -1
fi

PRODUCT=$(cd $CURDIR; python -c "import settings; print(settings.PRODUCT)")
SUBSYS=$(cd $CURDIR; python -c "import settings; print(settings.SUBSYS)")
MODULE=$(cd $CURDIR; python -c "import settings; print(settings.MODULE)")
if [ -z "$PRODUCT" ] || [ -z "$SUBSYS" ] || [ -z "$MODULE" ]; then
    echo "Support PRODUCT SUBSYS MODULE PORT in settings.py"
    exit -1
fi

SVC_NAME=${PRODUCT}.${SUBSYS}.${MODULE}
CONF_FILE=$CURDIR/conf/${PRODUCT}_${SUBSYS}_${MODULE}.yml
RPC_CONF_DIR=$RUNTIME_CONF_ROOT
LOG_DIR=$RUNTIME_LOG_ROOT
BinaryName=${PRODUCT}.${SUBSYS}.${MODULE}



if [ "$IS_SYSTEM_TEST_ENV" != "1" ]; then
    echo "$CURDIR/bin/${BinaryName} -conf="$CONF_FILE" -rpc="$RPC_CONF_DIR" -log="$LOG_DIR" -svc="$SVC_NAME" -port="${PORT}""
    exec $CURDIR/bin/${BinaryName} -conf="$CONF_FILE" -rpc="$RPC_CONF_DIR" -log="$LOG_DIR" -svc="$SVC_NAME" -port="${PORT}"
else
    echo "$CURDIR/bin/${BinaryName} -conf="$CONF_FILE" -rpc="$RPC_CONF_DIR" -log="$LOG_DIR" -svc="$SVC_NAME" -systemTest -test.coverprofile=systest.cov.out"
    exec $CURDIR/bin/${BinaryName} -conf="$CONF_FILE" -rpc="$RPC_CONF_DIR" -log="$LOG_DIR" -svc="$SVC_NAME" -systemTest -test.coverprofile=systest.cov.out
fi
