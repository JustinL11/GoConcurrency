>#### ***Purpose***

Code written in Go simulating how to calculate Pi
by throwing darts into a unit circle.

>#### ***Organization***
<br>
darts.go <br>
README.md <br>
DeltaResults.txt <br>
<br>

>#### ***How to use***
<br>

**darts.go** main program of the dart throwing.

Program can be called using Golang within the commandline
call "GO dart.go" and the program will return clocktimes based on the allocated threads currently on the program. 
Users may change how many threads being used by altering how many "go estimatePi()"s are within its own loop
As well as changing how many times the resulting loop "i" iterates accordingly
Both loops must be changed in parallel in order to display correct results.

Clocktime may change and vary depending on your own personal devices. 

...

>#### ***Design Notes***

    DeltaResults is data showing the amount of time it took to run the program depending on the size of darts thrown.

...
