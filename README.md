# Find Political Donors with Go

## A Go (Golang) Implementation Using Go Routines, Channels and Maps
To tackle the find political donors competition problem, I thought about using Go, which is now one of the most
utilized language to deal with system programming and processing massive volumes of data by many open source softwares.

This implementation provides the use of master worker architecture, where master maps tasks to workers, who processes the
data and updates the zip and date trackers, which uses Maps. The number of workers is set to 5, but can be updated for larger files. Using this master to worker architecture allows concurrent data processing, which improves latency.

## Use of two heaps
To track medians, two binary heaps (min binary heap) and (max binary heap) were utilized. Using this allows faster
run time to track medians compared to storing in list and sorting every iteration. The max heap was used for all 
values lower than median and min heap was used for all values higher than median. The purpose of medians is to
split the data in half, and using these two heaps properly does that and helps track the median by tracking highest value of left side and lowest value on right side.

## Go Channels
To send jobs and process entries in proper manner, channels were utilized to communicate with all the go routines
running. Each worker routine had a channel to which the master sent jobs to and the master kept track of order entry
by sending it to a seperate go routine, which was responsible for writing out the data

## How data is tracked
The data was tracked via maps, where the key was entry + zip or date and value was pointer to stream tracker, which uses the two heaps to track median, and tracked total and number of transactions. 

## Possible Improvements
Just like how the zip file is written as the data is being processed, the date file writes could also be improved, especially for larger files, where it may not be optimal to write it all out at the end. Make more tests in the test suite to get a better understanding of performance.
