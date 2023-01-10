FROM rsantev/polygon-edge:latest

ARG BOOTNODE_FOLDER

COPY ./bootnodes/${BOOTNODE_FOLDER} data-dir
COPY genesis.json .
COPY ./bootnodes/run.sh .

EXPOSE 10000 10001 10002

ENTRYPOINT "./run.sh"
