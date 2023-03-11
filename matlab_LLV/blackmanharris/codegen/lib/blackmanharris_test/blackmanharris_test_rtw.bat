@echo off

set MATLAB=D:\ProgramData\ML

cd .

if "%1"=="" ("D:\ProgramData\ML\bin\win64\gmake"  -f blackmanharris_test_rtw.mk all) else ("D:\ProgramData\ML\bin\win64\gmake"  -f blackmanharris_test_rtw.mk %1)
@if errorlevel 1 goto error_exit

exit 0

:error_exit
echo The make command returned an error of %errorlevel%
exit 1
