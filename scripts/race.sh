#!/bin/bash
go test -race -v -run TestDataRace1 ../file1.go
go test -race -v -run TestDataRace2 ../file2.go