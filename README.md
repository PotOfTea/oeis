# Simple cli tool to query oeis.org

oeis let's query [oeis](http://oeis.org/) database from terminal.
Tool can be downloaded from [releases](https://github.com/PotOfTea/oeis/releases) after download don't forget to `chmod +x`

```sh
chmod +x oeis
```

Example:

```sh
oeis 1 1 3 5 7
```

![terminal output example](/oeis_terminal_output.png)



To build from source-code run

```sh
make build
```

To test run

```sh
make test
```