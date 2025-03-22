package main

import (
	"bufio" 
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Please enter a Domian name:\n")

	// We scan multiple emails, one by one
	for scanner.Scan() { 
		checkDomain(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("error: Could not read from input: %v", err)
	}
}

// This function checks if the domain is valid
func checkDomain(domain string) {
	var hasMx, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	/* An MX (Mail Exchange) record is a type of DNS (Domain Name System) record that specifies which mail servers      are responsible for receiving emails for a given domain.
		 It returns slice of *net.MX struct, with Host - mail server hostname  & Priority - mail server priority.
	*/
	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
	if len(mxRecords) > 0 {
		hasMx = true
	}

	/* A TXT record is a DNS entry that holds text-based information for a domain. An SPF record is a specific          type of TXT record that lists the mail servers allowed to send emails on behalf of the domain.
	*/
	txtRecords, err := net.LookupTXT(domain) // fetching text Records
	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") { // in text records, spf records will be represented like this
			hasSPF = true
			spfRecord = record
			break
		}
	}

	/* A DMARC (DMARC (Domain-based Message Authentication, Reporting, and Conformance) record is a type of DNS         TXT record that helps prevent email spoofing and phishing by specifying how email servers should            handle SPF and DKIM failures.
  */
	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}

	fmt.Printf("| domain: %v\n| hasMx: %v\n| hasSPF: %v\n| spfRecord: %v\n| hasDMARC: %v\n| dmarcRecords: %v\n\n\n", domain, hasMx, hasSPF, spfRecord, hasDMARC, dmarcRecord)
}
