@echo off
setlocal
set SDKPATH=%GOPATH%\src\dajour.christophe.org
set SDKBIN=%SDKPATH%\.bin
call :copy_tools
goto :eof

:copy_tools
cd %SDKBIN%
copy "pptc.exe" "%GOBIN%" >NUL
cd %SDKPATH%

endlocal
