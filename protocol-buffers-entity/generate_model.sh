#!/bin/bash

function usage()
{
cat << _EOT_

 generate_model
------------------- author: xshoji

This script generates models.

Usage:
  ./$(basename "$0") --model-output-directory usersapi [ --proto ./proto/user.proto ]

Required:
  -m, --model-output-directory usersapi : Output directory for models.

Optional:
  -p, --proto ./proto/user.proto : File path of .proto

Helper options:
  --help, --debug

_EOT_
  [[ "${1+x}" != "" ]] && { exit "${1}"; }
  exit 1
}
function printColored() { local B="\033[0;"; local C=""; case "${1}" in "red") C="31m";; "green") C="32m";; "yellow") C="33m";; "blue") C="34m";; esac; printf "%b%b\033[0m" "${B}${C}" "${2}"; }
# [ keep-starter-parameters ] : curl -sf https://raw.githubusercontent.com/xshoji/bash-script-starter/master/ScriptStarter.sh |bash -s -  -n "generate_model" -a "xshoji" -d "This script generates models." -r "model-output-directory,usersapi,Output directory for models." -o "proto,./proto/user.proto,File path of .proto" -s -k -p



#------------------------------------------
# Preparation
#------------------------------------------
set -eu

# Parse parameters
readonly ARGS=("$@")
for ARG in "$@"
do
    SHIFT="true"
    [[ "${ARG}" == "--debug" ]] && { shift 1; set -eux; SHIFT="false"; }
    { [[ "${ARG}" == "--model-output-directory" ]] || [[ "${ARG}" == "-m" ]]; } && { shift 1; MODEL_OUTPUT_DIRECTORY="${1}"; SHIFT="false"; }
    { [[ "${ARG}" == "--proto" ]] || [[ "${ARG}" == "-p" ]]; } && { shift 1; PROTO="${1}"; SHIFT="false"; }
    { [[ "${ARG}" == "--help" ]] || [[ "${ARG}" == "-h" ]]; } && { shift 1; HELP="true"; SHIFT="false"; }
    { [[ "${SHIFT}" == "true" ]] && [[ "$#" -gt 0 ]]; } && { shift 1; }
done
[[ -n "${HELP+x}" ]] && { usage 0; }
# Check required parameters
[[ -z "${MODEL_OUTPUT_DIRECTORY+x}" ]] && { printColored yellow "[!] --model-output-directory is required.\n"; INVALID_STATE="true"; }
# Check invalid state and display usage
[[ -n "${INVALID_STATE+x}" ]] && { usage; }
# Initialize optional variables
[[ -z "${PROTO+x}" ]] && { PROTO=""; }
# To readonly variables
readonly MODEL_OUTPUT_DIRECTORY
readonly PROTO



#------------------------------------------
# Main
#------------------------------------------

# const
readonly SCRIPT_DIR="$(cd $(dirname "${BASH_SOURCE:-$0}") && pwd)"
readonly PROTO_DIR="${SCRIPT_DIR}/proto"
export SCRIPT_DIR
export PROTO_DIR

cat << __EOT__

[ Required parameters ]
model-output-directory: ${MODEL_OUTPUT_DIRECTORY}

[ Optional parameters ]
proto: ${PROTO}

__EOT__


function generate() {
  local SCRIPT_DIR="${1}"
  local MODEL_OUTPUT_DIR="${2}"
  local PROTO_FILE_PATH="${3}"
  protoc -I/usr/local/include -I${SCRIPT_DIR} \
    -I$GOPATH/src \
    -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
    --go_out=plugins=grpc:${MODEL_OUTPUT_DIR} \
    ${PROTO_FILE_PATH}
}

export -f generate

if [[ "${PROTO}" != "" ]]; then
  echo "Single mode. [ ${PROTO} ]"
  generate "${SCRIPT_DIR}" "${MODEL_OUTPUT_DIRECTORY}" "${PROTO}"
else
  echo "All mode."
  find "${PROTO_DIR}"/*.proto -type f |xargs -IXXX bash -c "generate ${SCRIPT_DIR} ${MODEL_OUTPUT_DIRECTORY} XXX"
fi
