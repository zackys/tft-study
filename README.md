tfft
====
Text File Filter Tools.

Usage
-----

Put input file name and output filename after `tfft` like this:
```
tfft input.txt output.txt
```
this works like copy comannd, the contents of output.txt is same as input.txt.

If you put filter option `addln` to apply input.txt like this:
```
tfft -filter "addln -start 200000 -step 20" input.txt output.txt
```
each line in input.txt is copied and attached Line-Number like COBOL source code which starts from 200000 incresing by 20 at the head, then written to output.txt.

If you put another filter option `rmln` like this,
```
tfft -filter "rmln -col 6" input.txt output.txt
```
first 6 charactors of each lines in input.txt are trimmed, and then copied to output.txt.

