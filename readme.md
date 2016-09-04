# Nginx Ip Access eXtractor

This is a simple command line application, written in go, to parse default format nginx access logs and group the entires per day per IP.

## Usage

To run, download or build the application then execute it via the command line, passing in nginx access log files as parameters.

**Example:**

```bash

nginx-ipax access.log access-example.com.log
```

## Output

This app will output to stdout in the following csv format:

```csv
Date, IP, Access Count
2016-06-09, 100.70.20.1, 40
2016-06-09, 8.8.8.8, 20
2016-07-22, 8.8.4.4, 5
2016-08-11, 192.168.0.1, 22
```

To send the output to a file you can redirect the output like so:

```bash
nginx-ipax access.log > ip_report.csv
```