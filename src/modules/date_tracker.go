package modules
/* Data Structure to maintain date information per recipient */


/* Data Structure to manage date data for recipients */
/* =========================================================================== */
type DateTracker struct{
    dateMap      *StreamMap         /* map of dates to streams */
    dateShown    map[string]bool    /* map of dates already shown */
}
/* =========================================================================== */


/* ZipTracker methods and functions */
/* =========================================================================== */

/* function to make a date tracker */
func MakeDateTracker() *DateTracker{
    newDate := new(DateTracker)
    /* setup the maps */
    newDate.dateMap = MakeStreamMap()
    newDate.dateShown = make(map[string]bool)
    
    return newDate
}


/* method to add entry to date tracker */
func (dt *DateTracker) AddDateEntry(entry *Entry) {
    /* add entry information */
    if !entry.ignoreDate && !entry.ignore{
        dt.dateMap.Put(entry.cmte_id + "!" + entry.date, entry.amount) /* look to ziptracker for why ! */
    }
}

/* method to show information about recipient in a date given entry*/
func (dt *DateTracker) UpdateDateInfo(entry *Entry) {
    entry.dateResult = ""
    /* if ignore then don't update */
    if entry.ignore || entry.ignoreDate || dt.dateShown[entry.cmte_id + "!" + entry.date]{
        return
    }
    /* otherwise update entry */
    dt.dateShown[entry.cmte_id + "!" + entry.date] = true
    entry.dateResult = entry.cmte_id + "|" + entry.date + "|"
    entry.dateResult += dt.dateMap.GetInfoStr(entry.cmte_id + "!" + entry.date, "|")
}

/* =========================================================================== */
