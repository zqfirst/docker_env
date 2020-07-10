@echo off

set binDir=%~dp0
set match=false
set command=%1
set container=%2
set project=%2
set operate=%2
set current=%CD%
set port=%2
::需要传递windows ssh key的镜像名称
set SSH_CONTAINER=php go

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
    echo Start building containers
    ::输出空行
    echo.
    docker-compose up -d --force-recreate --remove-orphans
	::加载ssh key
    call :sshkey
    ::恢复执行命令之前的目录
	call :cddir
    goto:EOF
)

if "%command%" == "start" (
    :: 启动所有容器
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
	docker-compose restart
	:: 启动swoole
    docker exec -it php /bin/bash -c "php /data/www/icenter/bin/start.php start"
	call :cddir
    goto:EOF
)

if "%command%" == "clear" (
    docker system prune --all
    call :cddir
    goto:EOF
)

if "%command%" == "utf8" (
    chcp 65001
    call :cddir
    goto:EOF
)

if "%command%" == "comp" (
    docker exec -it php /bin/bash -c "composer config -g secure-http false"
    docker exec -it php /bin/bash -c "composer config -g repositories.xeslib composer http://packagist.xesv5.com:8093"
    docker exec -it php /bin/bash -c "composer config -g repo.packagist composer https://mirrors.aliyun.com/composer/"
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
    docker logs --tail 50 --follow --timestamps %container%
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

if "%command%" == "sw" (
    echo Swoole service %operate%...
    if "%operate%" == "stop" (
        set operate=kill
    )
    docker exec -it php /bin/bash -c "php /data/www/icenter/bin/start.php %operate%"
    echo success
    call :cddir
    goto:EOF
)

if "%command%" == "ip" (
    docker inspect %container%
    call :cddir
    goto:EOF
)

if "%command%" == "s" (
    echo login %operate%...
    docker exec -it %operate% /bin/bash
    call :cddir
    goto:EOF
)

call :cddir
call :printHelpMsg
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
echo    sw;start    start swoole serve
echo    sw;stop     stop swoole serve
echo    rebuild     rebuild all containers from local Dockerfile
echo    s;ng        login shell on nginx docker conatainer
echo    s;php       login shell on php docker conatainer
echo    s;go      login shell on go docker conatainer
echo.
goto:EOF
