# Generate private keys for node
polygon-edge secrets init --data-dir data-dir

# Add node env variables to the staking contract repo
key=`cat data-dir/consensus/validator.key`
printf "PRIVATE_KEYS=%s\n" $key >> ./staking/staking-contracts/.env

output=$(polygon-edge secrets output --data-dir data-dir)
bls_key=$(echo $output| cut -d' ' -f 12)
printf "BLS_PUBLIC_KEY=%s\n" $bls_key >> ./staking/staking-contracts/.env
address=$(echo $output| cut -d' ' -f 7)
body="{\"address\": \"${address}\"}"
echo $body

# wait  5 minutes so the tx-executor would be started
sleep 300
# Fund node request
curl -v -XPOST -H 'Content-type: application/json' -d "$body" 'http://tx-executor:5000/fund-node'

# Setup staking for node
cd ./staking/staking-contracts && npm run build && npm run stake && npm run register-blskey
cd ./../../

polygon-edge server --data-dir ./data-dir --chain genesis.json --grpc-address :10000 --libp2p 0.0.0.0:10001 --jsonrpc :10002
