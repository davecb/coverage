package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

// PDQ constants (you'll need to implement or import these)
//const (
//	TERM  = 0
//	CEN   = 1
//	FCFS  = 2
//	EXACT = 3
//)
//
//var (
//	nodes   int
//	streams int
//)

func usage(progName string) {
	fmt.Fprintf(os.Stderr, "Usage: %s [-z think][-s service][-vd] from to by\n", progName)
}

func main() {
	var (
		from          = 1.0
		to            = 0.0
		by            = 0.0
		think         = 0.0
		serviceDemand = 0.0
		verbose       = false
		debug         = false
	)

	progName := filepath.Base(os.Args[0])

	if len(os.Args) == 1 {
		usage(progName)
		os.Exit(1)
	}

	// Parse command line arguments
	i := 1
	for i < len(os.Args) && os.Args[i][0] == '-' {
		switch os.Args[i][1] {
		case 'h':
			usage(progName)
			os.Exit(0)
		case 'z', 't':
			i++
			think, _ = strconv.ParseFloat(os.Args[i], 64)
		case 'd', 'x':
			debug = true
		case 's':
			i++
			serviceDemand, _ = strconv.ParseFloat(os.Args[i], 64)
		case 'v':
			verbose = true
		default:
			fmt.Fprintf(os.Stderr, "%s: unknown option -%c, ignored.\n", progName, os.Args[i][1])
		}
		i++
	}

	// Parse remaining arguments FIXME, these are floats
	if i < len(os.Args) {
		from, _ = strconv.ParseFloat(os.Args[i], 64)
		i++
	}
	if i < len(os.Args) {
		to, _ = strconv.ParseFloat(os.Args[i], 64)
		i++
	}
	if i < len(os.Args) {
		by, _ = strconv.ParseFloat(os.Args[i], 64)
	}
	pdq(progName, think, serviceDemand, from, to, by, verbose, debug) // FIXME return value is err
}

// pdq is the code that calls the pdq library
func pdq(progName string, thinkTime, serviceTime, from, to, by float64, verbose, debug bool) error {
	// Check parameters
	if thinkTime <= 0.0 {
		return fmt.Errorf("%s: thinkTime == %g, which is negative or zero and not valid", progName, thinkTime) // FIXME halting should be reported by caller
	}
	if serviceTime <= 0.0 {
		return fmt.Errorf("%s: serviceTime == %g, which is negative or zero and not valid", progName, serviceTime)
	}

	if from <= 0.0 {
		return fmt.Errorf("%s: from == %g, which is zero or negative, and not valid", progName, from)
	}
	if to <= 0.0 {
		// FIXME also check that we have less than 1000 users
		return fmt.Errorf("%s: to == %g, which is zero or negative and not valid", progName, to)
	}
	if by <= 0.0 {
		return fmt.Errorf("%s: by == %g, which is zero or negative, and not valid", progName, by)
	}

	// Print headers
	//fmt.Printf("General closed solution from PDQ where serviceTime = %g thinkTime = %g, progression = %g, %g, %g\n",
	//	serviceTime, thinkTime, from, to, by)
	//
	//if verbose {
	//	fmt.Printf("Load\tThroughput\tUtilization\tQueueLen\tResidence\tResponse\n")
	//} else {
	//	fmt.Printf("\"# Load,\" Response\n")
	//}

	for load := from; load <= to; load += by {
		doOneStep(load, thinkTime, serviceTime, verbose)
	}
	return nil
}

func doOneStep(load, thinkTime, serviceTime float64, verbose bool) {

	//s := "closed uniserver"
	//w.Init(s)
	//
	//modelName := "work"
	//CreateClosed(modelName, load, thinkTime) {
	//
	//	nodeName := "server0"
	//	CreateNode(nodeName)

	//SetDemand(nodeName, modelName, serviceTime)
	//
	//Solve() // FIXME exact?
	//r := Result(load)
	//fmt.Printf("%d\t%f\t%f\t%f\t%f\t%f\n",
	//	int(load),
	//	PDQ_GetThruput(TERM, "work"),
	//	PDQ_GetUtilization("server0", "work", TERM),
	//	PDQ_GetQueueLength("server0", "work", TERM),
	//	PDQ_GetResidenceTime("server0", "work", TERM),
	//	PDQ_GetResponse(TERM, "work"))
}
