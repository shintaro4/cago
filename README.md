# cago
An [elementary cellular automata](https://en.wikipedia.org/wiki/Elementary_cellular_automaton) with go

## Usage
```bash
$ go build
$ ./cago -help
```

```bash
Usage of ./cago:
  -n int
    	the number of the generations (default 16)
  -rule uint
    	a rule number (default 110)
```

## Examples

```bash
$ ./cago -rule 30
```

```bash
                o
               ooo
              oo  o
             oo oooo
            oo  o   o
           oo oooo ooo
          oo  o    o  o
         oo oooo  oooooo
        oo  o   ooo     o
       oo oooo oo  o   ooo
      oo  o    o oooo oo  o
     oo oooo  oo o    o oooo
    oo  o   ooo  oo  oo o   o
   oo oooo oo  ooo ooo  oo ooo
  oo  o    o ooo   o  ooo  o  o
 oo oooo  oo o  o ooooo  ooooooo
```

```bash
$ ./cago -rule 110 -n 16
```

```bash
                o
               oo
              ooo
             oo o
            ooooo
           oo   o
          ooo  oo
         oo o ooo
        ooooooo o
       oo     ooo
      ooo    oo o
     oo o   ooooo
    ooooo  oo   o
   oo   o ooo  oo
  ooo  oooo o ooo
 oo o oo  ooooo o
```

## Licence
[MIT](https://github.com/tcnksm/tool/blob/master/LICENCE)
