# Project Description
Run is a CLI program written in Go. Its core function is to evaluate GitHub repositories and npm modules.
Execute ./run <url_string> and the program will output a NDJSON containing the repository URL, a NetScore,
and various metrics such as RampUp, Correctness, BusFactor, ResponsiveMaintainer, LicenseScore, etc.

# Developer instructions
To install dependancies `go install`

To build call `go build`


To run on Linux call `cli text.txt`

To run on Windows call `cli.exe test.txt`

---
Github token must be set in environment variables as `GITHUB_TOKEN`, not .env file
