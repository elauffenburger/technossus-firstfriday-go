package conc

type SlidingWindow struct {
    window           chan interface{}
    windowsize       int
    numjobs          int
    numcompletedjobs int
    Done             chan interface{}
}

func SlidingWindowFactory(windowsize int, numjobs int) *SlidingWindow {
    result := &SlidingWindow{windowsize: windowsize, numjobs: numjobs}
    result.window = make(chan interface{}, windowsize)
    result.Done = make(chan interface{})

    return result
}

func (s SlidingWindow) GetNumJobs() int {
    return s.numjobs
}

func (s SlidingWindow) GetWindowSize() int {
    return s.windowsize
}

func (s *SlidingWindow) AddJob() {
    s.numjobs++
}

func (s *SlidingWindow) CompleteJob() {
    // release token so next job can start
    <-s.window
    s.numcompletedjobs++

    if s.numcompletedjobs == s.numjobs {
        s.Done <- true
    }
}

func (s *SlidingWindow) GetToken() {
    // get token -- this blocks if the window is full
    s.window <- true
}
