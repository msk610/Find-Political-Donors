package modules
/* Worker Structure to perform tasks provided by Master */

import(
    "sync"
)


/* Data Structure to manage jobs provided by master */
/* =========================================================================== */
type Worker struct{
    id          int                 /* worker id */
    jobs        chan *Entry         /* channel of entries to handle */
}
/* =========================================================================== */


/* Worker methods and functions */
/* =========================================================================== */

/* function to make a worker */
func MakeWorker(id int) *Worker {
    newWorker := new(Worker)
    /* set vars */
    newWorker.id = id
    newWorker.jobs = make(chan *Entry)
    return newWorker
}


/* method to run worker */
func (worker *Worker) Run(cond *sync.Cond, dt *DateTracker, zt *ZipTracker, status map[int]bool) {
    /* while true */
    for {
        /* see if worker channel still open */
        entry, more := <-worker.jobs
        if more {
            /* acquire a lock */
            cond.L.Lock()
            /* update trackers */
            zt.AddZipEntry(entry)
            dt.AddDateEntry(entry)
            /* update status */
            status[entry.id] = true
            cond.L.Unlock()
            /* signal */
            cond.Signal()
        } else {
            /* otherwise finish */
            return
        }
    }
}


/* method to consume a job */
func (worker *Worker) Consume(entry *Entry) {
    worker.jobs <- entry
}


/* method to stop consuming jobs */
func (worker *Worker) StopConsume() {
    close(worker.jobs)
}
