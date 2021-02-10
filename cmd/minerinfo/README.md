# minerinfo utility
Toy binary for interacting with the API to pull miner info from a given wallet address.

## Build + Usage
```
go build
./minerinfo -address "0x..."
```

You must give an address of a valid wallet that exists on the pool.

## Example Output
Note: Some info redacted.
```
Flexpool Miner '0x...' Stats
-

Unpaid Balance: 0.04026680 eth
Min Payout Threshold: 0.0500 eth         Donation Percent: 0.0100%       Round Share: 0.00010372%
Estimated Daily Eth: 0.0116875 eth       Total Paid: 0.06905678 eth      Total Donated: 0.00110340 eth

Workers:
	 mainpc (effective hashrate: 53MH/s) 	 (valid: 1318, stale: 0, invalid: 0)

Last 10 payments:
	 Txn: 0x... (amount: 0.06905068 eth) 	 2021-02-06 ..:..:.. -0500 EST

Last 10 blocks mined: 
	 No blocks mined yet.
```

