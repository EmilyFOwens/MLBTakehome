# MLB Take Home

Creator: Emily Owens

## Run Instructions

To build and execute the program via CLI, first run:
```
go build main.go
```
Main has two input flags, "date" and "teamId". These are the direct inputs into the sorting function.
ex. if you want the results for 11/02/2016, and your favorite team is the Cubs, you would enter:
```
./main -date="2016-11-02" -teamId=112 
```

The program will then execute, and the resulting sorted json will be written to the file ```schedule.json```
