@echo off

set port=%1
netstat  -aon | findstr  %port%