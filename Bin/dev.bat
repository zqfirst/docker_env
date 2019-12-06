@echo off

set binDir=%~dp0
set match=false
set command=%1
set container=%2


if "%command%" == "" (
    call :printHelpMsg
    goto:EOF
)

:: Checks if the configuration file exists
set sourceFile=%binDir%/../.env
if not exist %sourceFile% (
    echo The configuration file .env is missing. Refer to .env.example to create the configuration file
    goto:EOF
)

cd %binDir%/../

if "%command%" == "init" (
	set match=true

    echo Start building containers
    echo.
	docker-compose up -d --force-recreate
    goto:EOF
)

if "%command%" == "start" (
    set match=true

    docker-compose start
    goto:EOF
)

if "%command%" == "stop" (
	set match=true

	docker-compose stop
    goto:EOF
)

if "%command%" == "restart" (
	set match=true

	docker-compose restart
    goto:EOF
)

if "%command%" == "clear" (
    set match=true

    docker system prune --all
    goto:EOF
)

if "%command%" == "log" (
    set match=true
    docker logs --tail 50 --follow --timestamps %container%
    goto:EOF
)

if "%command%" == "sw-start" (
    set match=true

    echo Swoole service start...
    echo.
    docker exec -it php /bin/bash -c "php /data/www/icenter/bin/start.php start"
    echo Swoole service is up and running
    goto:EOF
)

if "%command%" == "sw-rstart" (
    set match=true

    echo Swoole service start...
    echo.
    docker exec -it php /bin/bash -c "php /data/www/icenter/bin/start.php rstart"
    echo Swoole service is up and running
    goto:EOF
)

if "%command%" == "sw-stop" (
    set match=true

    echo Swoole service stop...
    echo.
    docker exec -it php /bin/bash -c "php /data/www/icenter/bin/start.php  kill"
    echo Swoole service is stopped
    goto:EOF
)

if "%command%" == "s-php" (
    set match=true

    docker exec -it php /bin/bash
    goto:EOF
)

if "%command%" == "s-go" (
    set match=true

    docker exec -it mygo /bin/bash
    goto:EOF
)

if "%command%" == "shell-nginx" (
    set match=true

    docker exec -it fend-nginx /bin/bash
    goto:EOF
)

if "%command%" == "remove" (
    set match=true

    docker-compose down
    goto:EOF
)

if "%command%" == "status" (
    set match=true

    docker-compose ps
    goto:EOF
)

if "%command%" == "rebuild" (
    set match=true

    echo Start building containers...
    echo.
    docker-compose -f docker-compose-build.yml up -d --build

    goto:EOF
)

if %match% == false (
    call :printHelpMsg
    goto:EOF
)

goto:EOF

:: Print help information to console
:printHelpMsg
echo.
echo Usage: fendServe COMMAND
echo.
echo  Commands:
echo    init        build all containers from origin image
echo    start       start all containers
echo    stop        stop all containers
echo    restart     restart all container
echo    status      list all container
echo    sw-start    start swoole serve
echo    sw-stop     stop swoole serve
echo    rebuild     rebuild all containers from local Dockerfile
echo    remove      remove all containers
echo    shell-nginx login shell on nginx docker conatainer
echo    shell-php   login shell on php docker conatainer
echo.
goto:EOF
