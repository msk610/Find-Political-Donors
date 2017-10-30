package modules
/* Data Structure to manage entries */

import (
    "time"
    "strconv"
    "strings"
)


const CMTE_ID = 0
const ZIP = 10
const TRANSACTION_DT = 13
const TRANSACTION_AMT = 14
const OTHER_ID = 15


/* Data Structure to manage entries */
/* =========================================================================== */
type Entry struct{
    cmte_id     string      /* cmte_id */
    zip         string      /* zip code */
    date        string      /* transaction date */
    amount      int         /* transaction amount */
    other       string      /* other id */
    ignoreZip   bool        /* whether to ignore zip info */
    ignoreDate  bool        /* whether to ignore date info */
    ignore      bool        /* whether to ignore completely */
    zipResult   string      /* zip result for entry */
    dateResult  string      /* date result for entry */
    id          int         /* entry id */
}
/* =========================================================================== */


/* Entry methods and functions */
/* =========================================================================== */


/* function to format date string into proper format */
func Dateformat(s string) string {
    var dateform string
    var total int = len(s) - 1
    /* add the year first */
    for i := 4; i <= total; i++ {
        dateform += string(s[i])
    }
    /* then add the month */
    for i := 0; i < 2; i++{
    dateform += string(s[i])
    }
    /* then add the date */
    for i := 2; i < 4; i++{
    dateform += string(s[i])
    }
    return dateform
}


/* function to check if string is valid */
func ValidDate(d string) bool {
    formatted := Dateformat(d)
    _, err := time.Parse("20060102", formatted)
    return err == nil
}


/* function to check if valid zip */
func ValidZip(z string) bool {
    sLen := len(z)
    if sLen < 5 {
        return false
    }
    _, err := strconv.Atoi(z[:5])
    return err == nil

}

/* function to check if input is valid */
func ValidInput(cid string, tAmt string, other string) bool {
    _, err := strconv.Atoi(tAmt)
    return (cid != "" && tAmt != "" && other == "" && err == nil)
}

/* function to make an entry */
func MakeEntry(line string, id int) *Entry{
    newEntry := new(Entry)
    /* setup vars */
    data := strings.Split(line, "|")
    newEntry.cmte_id = data[CMTE_ID]
    newEntry.zip = data[ZIP]
    newEntry.date = data[TRANSACTION_DT]
    newEntry.other = data[OTHER_ID]
    newEntry.ignore = true
    newEntry.ignoreDate = true
    newEntry.ignoreZip = true
    newEntry.id = id

    /* check validity of variables */
    if ValidInput(newEntry.cmte_id, data[TRANSACTION_AMT], newEntry.other) {
        newEntry.ignore = false
        newEntry.amount, _ = strconv.Atoi(data[TRANSACTION_AMT])
    }

    if ValidZip(newEntry.zip){
        newEntry.ignoreZip = false
        newEntry.zip = newEntry.zip[:5]
    }

    if ValidDate(newEntry.date){
        newEntry.ignoreDate = false
    }
    
    return newEntry
}


/* method to get zip result */
func (e *Entry) ZipRes() string {
    return e.zipResult
}


/* method to get date result */
func (e *Entry) DateRes() string {
    return e.dateResult
}


/* method to get entry id */
func (e *Entry) GetID() int {
    return e.id
}

/* =========================================================================== */
