package modules
/* Data Structure to maintain zip information per recipient */


/* Data Structure to manage zip data for recipients */
/* =========================================================================== */
type ZipTracker struct{
    zipMap      *StreamMap      /* map of zips to streams */
}
/* =========================================================================== */


/* ZipTracker methods and functions */
/* =========================================================================== */

/* function to make a zip tracker */
func MakeZipTracker() *ZipTracker{
    newZip := new(ZipTracker)
    /* setup the maps */
    newZip.zipMap = MakeStreamMap()
    
    return newZip
}


/* method to add entry to zip tracker */
func (zt *ZipTracker) AddZipEntry(entry *Entry){
    /* add entry and immediately print information */
    entry.zipResult = ""
    if entry.ignore || entry.ignoreZip{
        return
    }
    /* use one stream map to save space instead of stream map per recipient by
    using recipient id + ! + zip as key */
    entry.zipResult = entry.cmte_id + "|"
    entry.zipResult += entry.zip + "|"
    entry.zipResult += zt.zipMap.PutRetStr(entry.cmte_id + "!" + entry.zip, entry.amount, "|")
}

/* =========================================================================== */
