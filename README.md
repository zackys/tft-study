tfft
====
Text File Filter Tools.

Usage
-----

Put input file name and output file name after `tfft` like this:
```
tfft input.txt output.txt
```
this works like copy comannd, the contents of output.txt is same as input.txt.

If you put filter option `addln` like this:
```
tfft -filter "addln -start 200000 -step 20" input.txt output.txt
```
each line of input.txt is copied with Line-Number which starts from 200000 incresing by 20 at the head(like COBOL source code style), and then written to output.txt.

If you put another filter option `rmln` like this,
```
tfft -filter "rmln -col 6" input.txt output.txt
```
first 6 charactors of each lines in input.txt are trimmed, and then copied to output.txt.

You can put filters with `|`, and they works like unix commands with pipeline:
```
tfft -filter "rmln -col 6 | addln -start 200000 -step 20" input.txt output.txt
```

There are another Global Options:
* -Lu: the line separator of output.txt is Unix style '\n'.
* -Lw: the line separator of output.txt is Windows style '\r\n'.
* -Lm: the line separator of output.txt is Mac style '\r'.
