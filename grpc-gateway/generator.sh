#!/bin/bash

function usage()
{
cat << _EOT_

 generator
--------------

This is generator.

Usage:
  ./$(basename "$0") [ --proto sample.proto ]

Optional:
  -p, --proto sample.proto : File name of .proto.

Helper options:
  --help, --debug

_EOT_
  [[ "${1+x}" != "" ]] && { exit "${1}"; }
  exit 1
}




#------------------------------------------
# Preparation
#------------------------------------------
set -eu

# Parse parameters
for ARG in "$@"
do
    SHIFT="true"
    [[ "${ARG}" == "--debug" ]] && { shift 1; set -eux; SHIFT="false"; }
    { [[ "${ARG}" == "--proto" ]] || [[ "${ARG}" == "-p" ]]; } && { shift 1; PROTO="${1}"; SHIFT="false"; }
    { [[ "${ARG}" == "--help" ]] || [[ "${ARG}" == "-h" ]]; } && { shift 1; HELP="true"; SHIFT="false"; }
    { [[ "${SHIFT}" == "true" ]] && [[ "$#" -gt 0 ]]; } && { shift 1; }
done
[[ -n "${HELP+x}" ]] && { usage 0; }
# Check invalid state and display usage
[[ -n "${INVALID_STATE+x}" ]] && { usage; }
# Initialize optional variables
[[ -z "${PROTO+x}" ]] && { PROTO=""; }



#------------------------------------------
# Main
#------------------------------------------
function generate() {
  protoc -I/usr/local/include -I${1} \
    -I$GOPATH/src \
    -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
    --go_out=plugins=grpc:${1} \
    --grpc-gateway_out=logtostderr=true:${1} \
    --swagger_out=logtostderr=true:${1} \
    ${1}/proto/${2}
}
export -f generate

SCRIPT_DIR="$(cd $(dirname "${BASH_SOURCE:-$0}") && pwd)"
if [[ "${PROTO}" != "" ]]; then
  echo "Single mode. [ ${PROTO} ]"
  generate "${SCRIPT_DIR}" "${PROTO}"
else
  echo "All mode."
  find "${SCRIPT_DIR}"/proto/*.proto -type f -exec basename {} \; |xargs -I{} bash -c "generate ${SCRIPT_DIR} {}"
fi
