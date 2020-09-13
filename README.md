# Simple cli tool to query oeis.org

`oeis` let's query [oeis.org](http://oeis.org/) database from terminal.
Tool can be downloaded from [releases](https://github.com/PotOfTea/oeis/releases). After download don't forget to set executable permision.

```sh
chmod +x oeis
```

![terminal output example](/oeis_terminal_output.png)

Example:

```sh
oeis 1 1 3 5 7
```


To build from source-code run

```sh
make build
```

To test run

```sh
make test
```