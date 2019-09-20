#!/bin/bash

Help(){

    cat <<-EOF

Usage: ${0##*/} COMMAND

  Commands:
    init        build all containers from origin image
    start       start all containers
    stop        stop all containers
    restart     restart all container
    status      list all container
    sw-start    start swoole serve
    sw-restart  restart swoole serve
    sw-stop     stop swoole serve
    rebuild     rebuild all containers from local Dockerfile
    remove      remove all containers
    s-nginx     login shell on nginx docker conatainer
    s-php       login shell on php docker conatainer
EOF
    exit 0;

}

if [ -z "$*" ]; then
    Help
fi

BIN_DIR=`DIRNAME $0`
PRO_DIR=`DIRNAME ${BIN_DIR}`

source ${PRO_DIR}/.env

case "$1" in
    init)
            # 启动容器服务
            printf "\033[33;49;1m 开始构建容器 \033[39;49;0m \n"
            if [ `uname -s | grep -i nt` ]; then
                docker-compose up -d --force-recreate
            else
                sudo -b docker-compose up -d --force-recreate
            fi
        ;;
    start)
            sudo docker-compose start
        ;;
    stop)
            sudo docker-compose stop
        ;;
    restart)
            sudo docker-compose restart
        ;;
    remove)
        sudo docker-compose down
        ;;
    status)
            docker-compose ps
        ;;
    sw-start)
            printf "\033[33;49;1m Swoole 服务开启... \033[39;49;0m \n"
            docker exec -it fend-php /bin/bash -c "php /home/www/fend/Bin/start.php -c /home/www/fend/App/Config/Swoole.php start"
            printf "\033[32;49;1m Swoole 服务已开启 \033[39;49;0m \n"
        ;;
    sw-restart)
            printf "\033[33;49;1m Swoole 服务重启... \033[39;49;0m \n"
            docker exec -it fend-php /bin/bash -c "php /home/www/fend/Bin/start.php -c /home/www/fend/App/Config/Swoole.php restart"
            printf "\033[32;49;1m Swoole 服务已开启 \033[39;49;0m \n"
        ;;
    sw-stop)
            printf "\033[33;49;1m Swoole 服务停止... \033[39;49;0m \n"
            docker exec -it fend-php /bin/bash -c "php /home/www/fend/Bin/start.php -c /home/www/fend/App/Config/Swoole.php kill"
            printf "\033[32;49;1m Swoole 服务已停止 \033[39;49;0m \n"
        ;;
    s-php)
            docker exec -it php-fpm /bin/bash
        ;;
    s-nginx)
            docker exec -it ng /bin/bash
        ;;
    rebuild)
        # 启动容器服务
        printf "\033[33;49;1m 开始构建容器 \033[39;49;0m \n"
        if [ `uname -s | grep -i nt` ]; then
            docker-compose -f docker-compose-build.yml up -d --build
        else
            sudo -b docker-compose -f docker-compose-build.yml up -d --build
        fi

        # 结尾清理中间镜像
        middle_images_num=`docker images -f "dangling=true" -q | awk '{print $1":"$2}' | wc -l`
        if [ $middle_images_num -ne 0 ]; then
            docker rmi $(docker images -f "dangling=true" -q)
            printf "\033[32;49;1m 中间镜像清理完成 \033[39;49;0m \n"
        fi
        ;;

    clearAll)
        #清除所有镜像 和 容器
        docker ps -a | awk '{print $1}' | xargs docker container rm;
         docker images -a | awk '{print $3}' | xargs docker rmi --force;
        ;;

    *)
            Help
        ;;
esac