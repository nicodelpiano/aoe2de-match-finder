# AOE2 DE Match Finder

AOE2 DE Match Finder is a Go library for fetching specific matches from the popular game Age of Empires 2 Definitive Edition.

It uses [aoe2.net API](https://aoe2.net/api) to make requests.

## Installation

## Usage

Run `go build` to compile the CLI executable.

To see the list of all the commands use `./aoe2de-match-finder --help` (powered by [Cobra](https://github.com/spf13/cobra)).

### Examples

```bash
❯ ./aoe2de-match-finder leaderboard -t 5
2021/10/16 19:00:27 ProfileID: 208611
Rank: 1
Rating: 2526
Name: Villese

2021/10/16 19:00:27 ProfileID: 409748
Rank: 2
Rating: 2520
Name: IMP | Capoch

2021/10/16 19:00:27 ProfileID: 197388
Rank: 3
Rating: 2517
Name: GL.TaToH

2021/10/16 19:00:27 ProfileID: 2858362
Rank: 4
Rating: 2509
Name: GL.JorDan_AoE

2021/10/16 19:00:27 ProfileID: 196310
Rank: 5
Rating: 2481
Name: [KGB] F1Re
```

```bash
❯ ./aoe2de-match-finder downloadMatch --match-id 122104697
Successfully downloaded match 122104697
```

## Contributing

Pull requests are welcome!

## License
[MIT](https://choosealicense.com/licenses/mit/)
