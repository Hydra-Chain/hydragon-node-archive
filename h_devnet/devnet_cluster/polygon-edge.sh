#!/bin/sh

set -e

# Check if jq is installed. If not exit and inform user.
if ! command -v jq >/dev/null 2>&1; then
  echo "The jq utility is not installed or is not in the PATH. Please install it and run the script again."
  exit 1
fi

POLYGON_EDGE_BIN=polygon-edge
CHAIN_CUSTOM_OPTIONS=$(
  tr "\n" " " <<EOL
--block-gas-limit 10000000
--epoch-size 10
--chain-id 51001
--name polygon-edge-docker
--premine 0x0000000000000000000000000000000000000000
--premine 0x211881Bb4893dd733825A2D97e48bFc38cc70a0c:0xD3C21BCECCEDA1000000
--premine 0x8c293C5b70b6493856CF4C7419E1Fb137b97B25d:0xD3C21BCECCEDA1000000
--burn-contract 0:0x0000000000000000000000000000000000000000
--proxy-contracts-admin 0x211881Bb4893dd733825A2D97e48bFc38cc70a0c
EOL
)

case "$1" in
"init")
  echo "Generating PolyBFT secrets..."
  secrets=$("$POLYGON_EDGE_BIN" polybft-secrets init --insecure --chain-id 51001 --num 5 --data-dir /data/data- --json)
  echo "Secrets have been successfully generated"

  rm -f /data/genesis.json

  echo "Generating PolyBFT genesis file..."
  "$POLYGON_EDGE_BIN" genesis $CHAIN_CUSTOM_OPTIONS \
    --dir /data/genesis.json \
    --consensus polybft \
    --validators-path /data \
    --validators-prefix data- \
    --reward-wallet 0xDEADBEEF:1000000 \
    --native-token-config "Hydra Token:HYDRA:18:true:$(echo "$secrets" | jq -r '.[0] | .address')" \
    --bootnode "/dns4/node-1/tcp/1478/p2p/$(echo "$secrets" | jq -r '.[0] | .node_id')" \
    --bootnode "/dns4/node-2/tcp/1478/p2p/$(echo "$secrets" | jq -r '.[1] | .node_id')" \
    --bootnode "/dns4/node-3/tcp/1478/p2p/$(echo "$secrets" | jq -r '.[2] | .node_id')" \
    --bootnode "/dns4/node-4/tcp/1478/p2p/$(echo "$secrets" | jq -r '.[3] | .node_id')" \
    --bootnode "/dns4/node-5/tcp/1478/p2p/$(echo "$secrets" | jq -r '.[4] | .node_id')"

  supernetID=$(cat /data/genesis.json | jq -r '.params.engine.polybft.supernetID')
  addresses="$(echo "$secrets" | jq -r '.[0] | .address'),$(echo "$secrets" | jq -r '.[1] | .address'),$(echo "$secrets" | jq -r '.[2] | .address'),$(echo "$secrets" | jq -r '.[3] | .address')"
  ;;
*)
  echo "Executing polygon-edge..."
  exec "$POLYGON_EDGE_BIN" "$@"
  ;;
esac
