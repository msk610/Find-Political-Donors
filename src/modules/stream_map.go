package modules
/* Using Golang Map with StreamData to make a mapper object that maps keys to streams */
import(
    "strconv"
)


/* Data Structure to map keys with streams */
/* =========================================================================== */
type StreamMap struct {
    sMap        map[string]*StreamData /* map of key to streams */
}
/* =========================================================================== */


/* StreamMap methods and functions */
/* =========================================================================== */

/* function to make a StreamMap */
func MakeStreamMap() *StreamMap {
    newMap := new(StreamMap)
    /* setup the map */
    newMap.sMap = make(map[string]*StreamData)
    return newMap
}


/* method to put key value in stream map */
func (sm *StreamMap) Put(key string, value int) (int, int, int) {
    /* if key doesn't exists */
    if sm.sMap[key] == nil {
        sm.sMap[key] = MakeStream()
    }
    /* get the updated trackers */
    newMedian := sm.sMap[key].Push(value)
    newTotal := sm.sMap[key].total
    newCount := sm.sMap[key].transactions

    return newMedian, newCount, newTotal
}

/* using put method to get string returns for tracker variables */
func (sm *StreamMap) PutRetStr(key string, value int, delim string) string {
    /* put the value */
    median, count, total := sm.Put(key, value)
    /* generate return str */
    retStr := strconv.Itoa(median) + delim
    retStr += strconv.Itoa(count) + delim
    retStr += strconv.Itoa(total)

    return retStr
}


/* method to get list of keys */
func (sm *StreamMap) GetKeys() []string {
    keys := make([]string, len(sm.sMap))
    /* loop and get key */
    var i int = 0
    for k := range sm.sMap {
        keys[i] = k
        i++
    }
    return keys
}


/* method to get key transactions */
func (sm *StreamMap) GetTransactions(key string) int {
    return sm.sMap[key].transactions
}

/* method to get key total */
func (sm *StreamMap) GetTotal(key string) int {
    return sm.sMap[key].total
}


/* method to add transactions */
func (sm *StreamMap) AddTransactions(key string) int {
    return sm.sMap[key].AddTransaction()
}

/* method to add total */
func (sm *StreamMap) AddTotal(key string, x int) int {
    return sm.sMap[key].AddTotal(x)
}


/* method to get median, count, total info in string*/
func (sm *StreamMap) GetInfoStr(key string, delim string) string {
    /* generate return str */
    retStr := strconv.Itoa(Round(sm.sMap[key].median)) + delim
    retStr += strconv.Itoa(sm.sMap[key].transactions) + delim
    retStr += strconv.Itoa(sm.sMap[key].total)

    return retStr
}

/* =========================================================================== */
