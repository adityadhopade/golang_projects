package main

import (
	"bufio" // it is buffer package used to parse whatever needed in the terminal
	"fmt"   // to print out stuff
	"log"
	"net" // to make request
	"os"
	"strings"
)

// func main calls the checkDomain
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("domain, hasMX, hasSPF,sprRecord,hasDMARC,dmarcRecord\n")

	// to scan the information
	// we are proving here multiple checkings but not the bulk checking like 1000's of records in the csv files.
	for scanner.Scan() {
		checkDomain(scanner.Text()) // whatever is inputted is now sent to the check domain function here
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error: could not read from input: %v\n", err)
	}

}

// passing domain as a string
// check domina will cehk all the records passed to it
func checkDomain(domain string) {
	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	// net.LookupMX(domain) // using the net package and using lookup function in it alos passing the Domain name in it.

	// we use a variable to capture it named mxRecords

	mxRecords, err := net.LookupMX(domain)

	if err != nil {
		log.Printf("Error:%v\n", err)
	}

	if len(mxRecords) > 0 {
		hasMX = true
	}

	txtRecords, err := net.LookupTXT(domain)

	// if error found
	if err != nil {
		log.Printf("Error:%v\n,err")
	}

	//if not found

	//range is used for the for loop and here we are not using the variables in the for loop like i and j so we can make use of the "_" blank address here

	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}

	// check for DMARK record

	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)

	if err != nil {
		log.Printf("Error%v\n", err)
	}

	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}

	fmt.Printf("%v ,%v, %v, %v, %v, %v", domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord)

}
