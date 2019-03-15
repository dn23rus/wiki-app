@ECHO off

IF [%1] == [] GOTO usage

@ECHO Connect to container %1
docker container exec -i %1 sh < %~dp0\build-in-container.sh
if %ERRORLEVEL% == 0 (
    @ECHO Restart container %1
    docker container restart %1
    @ECHO Done.
)

GOTO :eof

:usage
@ECHO Usage:
@ECHO %0 ^<container^>
@ECHO.
@ECHO Used to rebuild go app and then restart the container
@ECHO.
exit /B 1
