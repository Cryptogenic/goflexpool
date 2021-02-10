# poolinfo utility
Toy binary for interacting with the API to pull pool info.

## Build + Usage
```
go build
./poolinfo
```

## Example Output
```
Flexpool Stats
-

Miners: 2702 (Workers: 6883)

Hashrate: 1297GH/s (total)
	As: 36GH/s
	Au: 31GH/s
	Eu: 607GH/s
	Sa: 17GH/s
	Us: 605GH/s

PPLNS share window: 01:42:43 (hh:mm:ss)
Uncle rate: 6.00%
Average blocks per day: 20 (average reward: 4.41365273 eth)
	* Averages and uncle rate are over a 100 block period
```

