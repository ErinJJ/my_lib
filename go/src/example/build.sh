RUN_NAME="webarch.kite.example"
KITE_PATH="code.byted.org/kite"
REPO_PATH="example"

mkdir -p output/bin output/conf
cp script/bootstrap.sh script/settings.py output
chmod +x output/bootstrap.sh
cp conf/webarch_kite_example.yml output/conf
find conf/ -type f ! -name "*_local.*" | xargs -I{} cp {} output/conf/

export GO15VENDOREXPERIMENT="1"

GIT_SHA=$(git rev-parse --short HEAD 2>/dev/null || echo "GitNotFoundOrNoCommitFound")
DATE=$(date '+%Y%m%d%H%M%S')
VERSION=${GIT_SHA}-${DATE}

LINK_OPERATOR="="

if [ "$IS_SYSTEM_TEST_ENV" != "1" ]; then
    go build -o output/bin/${RUN_NAME}
else
    go test -c -covermode=set -o output/bin/${RUN_NAME} -coverpkg=./...
fi


