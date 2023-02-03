# Findings

## Reward Mechanism feature
### Edge accounts
Edge uses the same concept as Ethereum for implementing its accounts. It uses levelDB to store the storage (key-value pairs). keccak256 of an address is the key for its value which is bytes[]. The value can be serialized to object with nonce, balance, storageRoot (root hash of the evm storage of a smart contract; empty for EOA), codeHash - emv code of smart (empty for EOA).  

Coin balance is kept here.

## Default reward mechanism


## Polygon PoS

There is a limit of 100 active validators.
