@echo off

set binDir=%~dp0
set match=false
set command=%1
set container=%2
set project=%2
set current=%CD%

::需要传递windows ssh key的镜像名称
set SSH_CONTAINER=php mygo

::帮助菜单
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

::进入到根目录 需要用到.env 文件
cd %binDir%/../

if "%command%" == "init" (
	set match=true

    echo Start building containers
    ::输出空行
    echo.
    docker-compose up -d --force-recreate
	::加载ssh key
    call :sshkey
    ::恢复执行命令之前的目录
	call :cddir
    goto:EOF
)

if "%command%" == "start" (
    set match=true

    docker-compose start
    :: 启动swoole
    docker exec -it php /bin/bash -c "php /data/www/icenter/bin/start.php start"
    call :cddir
    goto:EOF
)

if "%command%" == "stop" (
	set match=true

	docker-compose stop
	call :cddir
    goto:EOF
)

if "%command%" == "restart" (
	set match=true

	docker-compose restart
	call :cddir
    goto:EOF
)

if "%command%" == "clear" (
    set match=true

    docker system prune --all
    call :cddir
    goto:EOF
)

if "%command%" == "comp" (
    set match=true
    docker exec -it php /bin/bash -c "cd %project% && composer update"
    call :cddir
    goto:EOF
)

if "%command%" == "compc" (
    echo "config success"
    call :cddir
    goto:EOF
)

if "%command%" == "log" (
    set match=true
    docker logs --tail 50 --follow --timestamps %container%
    call :cddir
    goto:EOF
)

if "%command%" == "sw-start" (
    set match=true

    echo Swoole service start...
    echo.
    docker exec -it php /bin/bash -c "php /data/www/icenter/bin/start.php start"
    echo Swoole service is up and running
    call :cddir
    goto:EOF
)

if "%command%" == "sw-restart" (
    set match=true

    echo Swoole service start...
    echo.
    docker exec -it php /bin/bash -c "php /data/www/icenter/bin/start.php restart"
    echo Swoole service is up and running
    call :cddir
    goto:EOF
)

if "%command%" == "sw-stop" (
    set match=true

    echo Swoole service stop...
    echo.
    docker exec -it php /bin/bash -c "php /data/www/icenter/bin/start.php  kill"
    echo Swoole service is stopped
    call :cddir
    goto:EOF
)

if "%command%" == "s-php" (
    set match=true

    docker exec -it php /bin/bash
    call :cddir
    goto:EOF
)

if "%command%" == "s-go" (
    set match=true

    docker exec -it mygo /bin/bash
    call :cddir
    goto:EOF
)

if "%command%" == "s-ng" (
    set match=true

    docker exec -it ng /bin/bash
    call :cddir
    goto:EOF
)

if "%command%" == "s-redis" (
    set match=true

    docker exec -it redis /bin/bash
    call :cddir
    goto:EOF
)

if "%command%" == "remove" (
    set match=true

    docker-compose down
    call :cddir
    goto:EOF
)

if "%command%" == "status" (
    set match=true

    docker-compose ps
    call :cddir
    goto:EOF
)

if "%command%" == "rebuild" (
    set match=true

    echo Start building containers...
    echo.
    docker-compose -f docker-compose-build.yml up -d --build
    call :cddir
    goto:EOF
)

if %match% == false (
    call :cddir
    call :printHelpMsg
    goto:EOF
)

goto:EOF

:sshkey
:: 初始化ssh key
(for %%a in (%SSH_CONTAINER%) do (
  docker exec -it %%a /bin/bash -c "sh /root/ssh.sh"
))
goto:EOF

:: return prev dir
:cddir
echo.
cd %current%
echo.
goto:EOF

:: Print help information to console
:printHelpMsg
echo.
echo Usage: DEV COMMAND
echo.
echo  Commands:
echo    init        build all containers from origin image
echo    start       start all containers
echo    stop        stop all containers
echo    clear       clear all images and containers
echo    comp        update composer depend
echo    compc       set composer to default config
echo    restart     restart all container
echo    status      list all container
echo    sw-start    start swoole serve
echo    sw-stop     stop swoole serve
echo    rebuild     rebuild all containers from local Dockerfile
echo    remove      remove all containers
echo    s-ng        login shell on nginx docker conatainer
echo    s-php       login shell on php docker conatainer
echo.
goto:EOF
