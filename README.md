Email Verifier Tool is a simple CLI application written in Go that checks if a given domain is valid by verifying its DNS records. The tool performs the following checks:

1) MX (Mail Exchange) Record – Ensures the domain can receive emails.
2) SPF (Sender Policy Framework) Record – Verifies the authorized mail servers for sending emails.
3) DMARC (Domain-based Message Authentication, Reporting & Conformance) Record – Checks the domain's email authentication policy.

How to run :

go run main.go 
Please enter a Domian name:
hotmail.com
| domain: hotmail.com
| hasMx: true
| hasSPF: true
| spfRecord: v=spf1 ip4:157.55.9.128/25 include:spf-a.outlook.com include:spf-b.hotmail.com include:spf-b.outlook.com include:spf-a.hotmail.com include:_spf-ssg-b.microsoft.com include:_spf-ssg-c.microsoft.com -all
| hasDMARC: true
| dmarcRecords: v=DMARC1; p=none; rua=mailto:rua@dmarc.microsoft;ruf=mailto:ruf@dmarc.microsoft;fo=1:s:d

