# RootPowerPerf
A fun performance challenge! Try to complete the challenge as fast as you can in any language!

## How It Works
Clone this project, and use `generate.py` to create a data set. It will be saved in the same directory (or folder) as `generate.py`.
Everybody else's solution will be in the `src` folder.
Now you need to add your solution!

1. Go into `src`.
2. If there isn't a directory for your language, create one! (simply named the language).
3. In that directory create another directory with your GitHub name, or any other name you would like to go by.
4. Read down into section "Setting Up" to find out what extra files you have to create.
5. Start coding!

## What You Must Do
So what do you have to code?
Firstly, you may notice that the `data.txt` file has many lines, and each line has integers seperated by commas.
__NOTE__ that lines are seperated with `\r`, but __you are permitted to change this by giving `generate.py` a different character__.
Your job is to __sum__ those number in one of those lines.
You then have to squish the result into uint16 range (0 -> 65,535).
Then, with the total, you must find `y` such that `y = x^(3/2)`.
For example, to get this solution, I would take x (which is the total), then raise it to the power of 3, and then take the square root of that.
This could be a great part for performance gain, how would you get `y`?
`y` should be rounded __down__ to __0__ decimal places.
Lastly, find the min, max, and __mean__ of the dataset and either show it in standard out or write the answer to a file.
The mean also has to be rounded __down__ to __0__ decimal places.

## Setting Up
There is an extra file you have to create, which is `ME.md`.
This file is about you!
Firstly, a link to you GitHub, or any more of your connections would be great for people to know about you.
Add as little or as much as you'd like, just nothing other than that.
__Compulsory__ is your GitHub account.
Also add in this file your compiler/interpreter, and the version, and preferably a link to this compiler/interpreter.
Instructions on how to build/run the program would also be helpful.

## For Timing Your Code
Since this is transcending programming languages, we can't use any language specific timers.
In this way, I will be running your program through another program found in `runner`.
All this will do is start a timer, run your code, finish the time, and show the result.
This means it should tax your program, and everybody has that same disadvantage.
This code will be run on my laptop, with an R7, 16GB of ram, and 8 cores (Windows 11).
If your language has extra compilation flags for extra performance, make sure to write this either in `ME.md` or right at the top of your code file.
Once you've built, you can head into your terminal, head into the `runner` directory, and run `runner "pathToYourCode/code"`.

## Baseline
I have put in my baseline implementation in Go, which may help you with any other questions you may have.

Now _Go!_

## Leaderboard
| Place | Username | Language | Time (ms) |
| ----- | -------- | -------- | --------- |
| 1     | ChigBeef | Go       | 51,038    |