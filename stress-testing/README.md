### Create a testing Cluster
#### Setup Main nodes
1. Download Polygon Edge
```
git clone https://github.com/R-Santev/polygon-edge.git 
git checkout testing-stuff
cd polygon-edge/  
go build -o polygon-edge main.go  
sudo mv polygon-edge /usr/local/bin
```
2. Generate two bootnodes
   ```
   polygon-edge secrets init --data-dir bootnode-one
   polygon-edge secrets init --data-dir bootnode-two
   ```
4. Genereate genesis file
```
polygon-edge genesis --pos --ibft-validator <bootnode 1 address>:<BLS Public key> --ibft-validator <bootnode 2 address>:<BLS Public key> --bootnode /ip4/127.0.0.1/tcp/<libp2p port>/p2p/<Node ID> --bootnode /ip4/127.0.0.1/tcp/<libp2p port>/p2p/<Node ID>  --premine=<bootnode 1 address>:100000000000000000000000 --max-validator-count 500
```
5. Add bootnodes to staking
They are automatically pre-added in the staking contract as stakers because we set them in the genesis file with the --ibft-validator flag.
6. Run bootnodes manually
```
polygon-edge server --data-dir ./bootnode-one --chain genesis.json --grpc-address :10000 --libp2p :10001 --jsonrpc :10002 --seal

polygon-edge server --data-dir ./bootnode-two --chain genesis.json --grpc-address :20000 --libp2p :20001 --jsonrpc :20002 --seal
```

#### Run additional node from docker image
1. Copy genesis.json file from the main nodes folder and paste it in polygon-edge/stress-testing directory (where is the Dockerfile.test file)
2. Build image from Dockerfile.test
Open a terminal in the Polygon-Edge/stress-testing folder and paste
```
sudo docker build -t stress-test:latest -f ./Dockerfile.test --network host --build-arg MAIN_NODE_KEY_ARG=<main node private key> --no-cache .
```
3. Use the created image to run new nodes
```
sudo docker run --network host  stress-test:latest
```

## Additional Notes

### Run a Node Setup

#### Prerequesits:

1. Download Polygon Edge
2. We need a genesis file
```
polygon-edge genesis --pos --ibft-validator <Address>:<BLS Public key> --ibft-validator <Address>:<BLS Public key> --bootnode /ip4/<url>/tcp/<port>/p2p/<Node ID> --bootnode /ip4/<url>/tcp/<port>/p2p/<Node ID>
```

#### Script Steps
1. Install Polygon Edge
2. Initialize Data Folder and Generate Private key
   ```
   polygon-edge secrets init --data-dir data-dir
   ```

3. Add Genesis file
   Get copy of it and paste it;
4. Run Client
```
polygon-edge server --data-dir ./data-dir --chain genesis.json --grpc-address :10000 --libp2p :10001 --jsonrpc :10002 --seal
```
#### Stake with Node
   1. Download staking contract repo
   2. Install it
   3.  Set env variables
   4. Send Coins to this Staker from a specific one
   5.  Run stake script
   6.  Run set BLS public key