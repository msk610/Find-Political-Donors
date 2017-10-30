package modules
/* Data Structure to keep track of:
    median, total, and count of transactions in stream */

import(
    "math"
)


/* Data Structure to Keep Track Median, Total, and Transactions */
/* =========================================================================== */
type StreamData struct {
    lowerHalf       *MaxHeap            /* Store lower half using max heap */
    upperHalf       *MinHeap            /* Store upper half using min heap */
    median          float64             /* Current Median */
    transactions    int                 /* total number of transactions */
    total           int                 /* sum of stream data */
}
/* =========================================================================== */


/* StreamData methods and functions */
/* =========================================================================== */

/* function to make StreamData */
func MakeStream() *StreamData{
    /* make a new stream */
    newStream := new(StreamData)
    /* initialize max and min heaps */
    newStream.lowerHalf = MakeMaxHeap()
    newStream.upperHalf = MakeMinHeap()
    /* setup tracking variables */
    newStream.median = -999.0
    newStream.transactions = 0
    newStream.total = 0

    return newStream
}


/* function to return closest int */
func Round(x float64) int {
    /* round the number down */
    floorValue := math.Floor(x)
    if x - floorValue >= 0.5 {
        /* if higher than .5 then round up */
        return int(floorValue + 1)
    }
    /* otherwise round down */
    return int(floorValue)
}


/* method to push data to stream */
func (stream *StreamData) Push(x int) int {
    /* add to transaction and total */
    stream.transactions++
    stream.total += x
    val := float64(x)
    
    /* if no value added */
    if stream.median == -999.0 {
        stream.median = val
        PushMinHeap(stream.upperHalf, x)
    } else {
        /* otherwise */
        leftSize := len((*stream.lowerHalf))
        rightSize := len((*stream.upperHalf))

        /* balance sides */
        if leftSize < rightSize {
            /* case 1: leftsize is smaller so balance */
            if val < stream.median{
                PushMaxHeap(stream.lowerHalf, x)
            } else {
                PushMaxHeap(stream.lowerHalf, (*stream.upperHalf)[0])
                PopMinHeap(stream.upperHalf)
                PushMinHeap(stream.upperHalf, x)
            }
        } else if rightSize > leftSize {
            /* case 2: rightsize is smaller so balance */
            if val < stream.median{
                PushMinHeap(stream.upperHalf, (*stream.lowerHalf)[0])
                PopMaxHeap(stream.lowerHalf)
                PushMaxHeap(stream.lowerHalf, x)
            }
        } else {
            /* case 3: same size both so add to appropriate */
            if val < stream.median{
                PushMaxHeap(stream.lowerHalf, x)
            } else {
                PushMinHeap(stream.upperHalf, x)
            }
        }

        /* find new median */
        newLeftSize := len((*stream.lowerHalf))
        newRightSize := len((*stream.upperHalf))
        /* top of left and top of right divide by 2 */
        stream.median = (float64((*stream.lowerHalf)[0]) + float64((*stream.upperHalf)[0])) / 2.0
        /* if size mismatch then adjust newMedian */
        if newLeftSize < newRightSize {
            stream.median = float64((*stream.upperHalf)[0])
        } else if newLeftSize > newRightSize {
            stream.median = float64((*stream.lowerHalf)[0])
        }
    }

    return Round(stream.median)
}


/* method to add transaction */
func (stream *StreamData) AddTransaction() int {
    stream.transactions++
    return stream.transactions
}


/* method to add total */
func (stream *StreamData) AddTotal(x int) int {
    stream.total += x
    return stream.total
}

/* =========================================================================== */
