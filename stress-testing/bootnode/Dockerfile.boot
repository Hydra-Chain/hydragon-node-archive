FROM 0xpolygon/polygon-edge:latest

COPY ./bootnode/secrets data-dir
COPY genesis.json .
COPY run.sh .

ENTRYPOINT "./run.sh"
