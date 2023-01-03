key=`cat data-dir/consensus/validator.key`
printf "PRIVATE_KEYS=%s\n" $key >> ./staking/staking-contracts/.env

output=$(polygon-edge secrets output --data-dir data-dir)
bls_key=$(echo $output| cut -d' ' -f 12)
printf "BLS_PUBLIC_KEY=%s\n" $bls_key >> ./staking/staking-contracts/.env
