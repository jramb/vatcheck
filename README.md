# vatcheck

Check VAT number against VIES: http://ec.europa.eu/taxation_customs/vies/

The Program is written in Go, I tested it on Linux and Windows

## Installation

    go install github.com/jramb/vatcheck
    
or you can clone or download this repository and do `go install` or `go build`.

## Usage (example):
    
    ./vatcheck SE556950473001
    Result: SE556950473001 = true
    Name:  Navigate Approval Solutions AB
           c/o KISTA SCIENCE TOWER 
    FÄRÖGATAN 33 
    164 51 KISTA
    Date:  2016-02-20+01:00
    

on Windows use `vatcheck.exe`

## Disclaimer
This program is not endorsed by the VIES service. Please
see the VIES disclaimer and FAQ on the VIES homepage
http://ec.europa.eu/taxation_customs/vies/viesdisc.do
before using it and comply with the terms of usage outlined there.

This program is put here as-is. If it stops working: not really my problem.

And always pay your taxes!
