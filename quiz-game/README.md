# Quiz Game

This program is a quiz game. It will read in a csv file containing a sequence of problems followed by answers in a format such as below

```
5+5,10
7+3,10
1+1,2
8+3,11
1+2,3
8+6,14
3+1,4
1+4,5
5+1,6
2+3,5
3+3,6
2+4,6
5+2,7
```

The file will be defaulted to `problems.csv` but may be overridden by using the `-f` flag. The quiz is timed. The default time is 30 seconds, but may also be overridden by entering the time desired (in seconds) into the `-t` flag. At the end of the quiz, whether timed out or finished it will output the final score.