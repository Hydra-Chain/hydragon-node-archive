# General findings

## Validators update

Events for change in the balance are handled after every block and validators state is saved in the DB
At the end of the epoch the new state (lasly update at block endOfEpoch - 1) is applied.
PROBLEM: changes in the validators made in the last block of an peoch are not applied in the next epoch but in the next + 1 epoch.
So our contract would have different state of the current validators compared to the state in the node

// TODO: Modify our implementation of the contracts
