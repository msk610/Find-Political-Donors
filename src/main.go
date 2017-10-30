package main
/* Main Entry Point */

import(
    "./modules"
    "fmt"
    "os"
    "bufio"
    "hash/fnv"
    "strings"
    "sync"
    "sync/atomic"
    "strconv"
)


/* function to hash string to int */
func Hash(s string) int {
    h := fnv.New32a()
    h.Write([]byte(s))
    return int(h.Sum32())
}


/* function to assign tasks to workers */
func AssignJob(line string, nWorkers int) int{
    /* get the cmte id */
    data := strings.Split(line, "|")
    cmteID := data[0]
    return Hash(cmteID) % nWorkers
}


/* function to print help instructions */
func PrintHelp() {
    fmt.Println("Help instructions: ")
    fmt.Println("================================")
    fmt.Println("usage: go run main.go  path/itcont.txt path/medianvals_by_zip.txt path/medianvals_by_date.txt numberOfWorkers")
    fmt.Println("================================\n")
}


func main() {
    /* check command line args passed */
    if len(os.Args) != 5 {
        PrintHelp()
        panic("main.go: improper usage!!!")
    }
    /* open source file */
    file, err := os.Open(os.Args[1])
    if err != nil{
        panic(err)
    }
    defer file.Close()

    /* setup out files */
    zipFile, err1 := os.Create(os.Args[2])
    if err1 != nil {
        panic(err1)
    }
    dateFile, err2 := os.Create(os.Args[3])
    if err2 != nil {
        panic(err2)
    }
    defer zipFile.Close()
    defer dateFile.Close()
    
    fmt.Println("main.go: initializing...")
    /* setup trackers and list of all entries and other master variables */
    zt := modules.MakeZipTracker() /* zip tracker */
    dt := modules.MakeDateTracker() /* date tracker */
    var entries []*modules.Entry  /* list of entries (used for writting date info) */
    status := make(map[int]bool) /* job status */
    done := make(chan bool) /* whether all jobs done */
    cond := &sync.Cond{L: &sync.Mutex{}} /* conditional var */
    allJobs := make(chan *modules.Entry) /* channel of all jobs in order */

    /* setup workers */
    argWorkers, werr := strconv.Atoi(os.Args[4])
    nWorkers := int(argWorkers)
    if werr != nil{
        panic("main.go: bad number of workers passed")
    }
    var workers []*modules.Worker
    /* loop and make worker */
    for i := 0; i < nWorkers; i++{
        workers = append(workers, modules.MakeWorker(i))
    }

    /* start all the workers */
    for i := 0; i < nWorkers; i++{
        go workers[i].Run(cond, dt, zt, status)
    }
    
    /* run a go routine to send jobs to workers */
    go func() {
        fmt.Println("main.go: sending jobs to workers...")
        var id uint64 = 0
        scanner := bufio.NewScanner(file)
        /* loop through file and send jobs */
        for scanner.Scan() {
            atomic.AddUint64(&id, 1)
            line := scanner.Text()
            wid := AssignJob(line, nWorkers)
            entry := modules.MakeEntry(line, int(id))
            entries = append(entries, entry)
            allJobs <- entry
            workers[wid].Consume(entry)
        }
        /* stop consumtions */
        for i := 0; i < nWorkers; i++{
            workers[i].StopConsume()
        }
        /* close all jobs */
        close(allJobs)
    }()

    /* run a go routine to write job outputs in order */
    go func() {
        fmt.Println("main.go: processing jobs...")
        for {
            /* get job in order */
            entry, more := <- allJobs
            /* if channel not closed */
            if more {
                cond.L.Lock()
                /* check if finished */
                for !status[entry.GetID()]{
                    cond.Wait()
                }
                cond.L.Unlock()
                /* if can then write zip result */
                if entry.ZipRes() != ""{
                    zipFile.WriteString(entry.ZipRes() + "\n")
                    zipFile.Sync()
                }
            } else {
                /* all jobs finished so update date info and write out*/
                for e := range entries {
                    dt.UpdateDateInfo(entries[e])
                    if entries[e].DateRes() != ""{
                        dateFile.WriteString(entries[e].DateRes() + "\n")
                        dateFile.Sync()
                    }
                }
                done <- true
                close(done)
                return
            }
        }
    }()

    /* wait till all jobs are done */
    <- done
    fmt.Println("main.go: finished processing file")
    
}
