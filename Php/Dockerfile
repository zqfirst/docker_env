# Fend php image

FROM centos:7

ARG php_version
ARG swoole

##
# ---------- env settings ----------
# --build-arg swoole=4.4.5
##
ENV PHP_VERSION=${php_version:-"php73"}
ENV SWOOLE_VERSION=${swoole:-"4.4.12"}
ENV FPM_CONF_DIR="/etc/opt/remi/${PHP_VERSION}"
ENV FPM_CONFD_DIR="/etc/opt/remi/${PHP_VERSION}/php-fpm.d"

RUN set -ex \
    && rpm --import https://rpms.remirepo.net/RPM-GPG-KEY-remi \
#    && rpm -ivh http://dl.fedoraproject.org/pub/epel/epel-release-latest-7.noarch.rpm \
#    && rpm -ivh https://mirrors.tuna.tsinghua.edu.cn/remi/enterprise/remi-release-7.rpm \
    && rpm -Uvh https://dl.fedoraproject.org/pub/epel/epel-release-latest-7.noarch.rpm \
    && rpm -Uvh http://rpms.remirepo.net/enterprise/remi-release-7.rpm \
    && yum -y update \
    && yum install -y gcc \
        gcc-c++ \
        ca-certificates \
        curl \
        tar \
        xz \
        tzdata \
        pcre \
        autoconf \
        automake \
        make \
        zlib \
        zlib-devel \
        openssl \
        openssl-devel \
        automake \
        make \
        git \
    && yum --enablerepo=remi -y install \
    ${PHP_VERSION}-php \
    ${PHP_VERSION}-php-pear \
    ${PHP_VERSION}-php-bcmath \
    ${PHP_VERSION}-php-pecl-jsond-devel \
    ${PHP_VERSION}-php-mysqlnd \
    ${PHP_VERSION}-php-gd \
    ${PHP_VERSION}-php-common \
    ${PHP_VERSION}-php-fpm \
    ${PHP_VERSION}-php-intl \
    ${PHP_VERSION}-php-cli \
    ${PHP_VERSION}-php \
    ${PHP_VERSION}-php-xml \
    ${PHP_VERSION}-php-opcache \
    ${PHP_VERSION}-php-pecl-apcu \
    ${PHP_VERSION}-php-pdo \
    ${PHP_VERSION}-php-gmp \
    ${PHP_VERSION}-php-process \
    ${PHP_VERSION}-php-pecl-imagick \
    ${PHP_VERSION}-php-devel \
    ${PHP_VERSION}-php-mbstring \
    ${PHP_VERSION}-php-zip \
    ${PHP_VERSION}-php-ldap \
    ${PHP_VERSION}-php-imap \
    ${PHP_VERSION}-php-pecl-mcrypt \
    ${PHP_VERSION}-php-bcmath \
    ${PHP_VERSION}-php-fileinfo \
    ${PHP_VERSION}-php-memcached \
    ${PHP_VERSION}-php-sysvsem \
    ${PHP_VERSION}-php-mcrypt \
    ${PHP_VERSION}-php-soap \
    ${PHP_VERSION}-php-openssl \
    ${PHP_VERSION}-php-pcntl \
    ${PHP_VERSION}-php-sockets \
    ${PHP_VERSION}-php-curl \
    ${PHP_VERSION}-php-xml \
    ${PHP_VERSION}-php-zip \
    ${PHP_VERSION}-php-zlib \
    && ln -s /opt/remi/${PHP_VERSION}/root/bin/php /usr/bin/php \
    && ln -s /opt/remi/${PHP_VERSION}/root/bin/phpize /usr/local/bin/phpize \
    && ln -s /opt/remi/${PHP_VERSION}/root/bin/php-config /usr/local/bin/php-config \
    && ln -s /opt/remi/${PHP_VERSION}/root/bin/pecl /usr/local/bin/pecl \
    && ln -s /opt/remi/${PHP_VERSION}/root/sbin/php-fpm /usr/local/sbin/php-fpm \
    && sed -e 's/daemonize = yes/daemonize = no/g' -i ${FPM_CONF_DIR}/php-fpm.conf \
    && sed -e 's/listen.allowed_clients/;listen.allowed_clients/g' -i ${FPM_CONF_DIR}/php-fpm.d/www.conf \
    && sed -e 's/listen = 127.0.0.1:9000/listen = 0.0.0.0:9000/g' -i ${FPM_CONF_DIR}/php-fpm.d/www.conf \
    && cd /tmp \
    && curl -SL "https://github.com/swoole/swoole-src/archive/v${SWOOLE_VERSION}.tar.gz" -o swoole.tar.gz \
    # php extension:swoole
    && mkdir -p swoole \
    && tar -xf swoole.tar.gz -C swoole --strip-components=1 \
    && ( \
        cd swoole \
        && phpize \
        && ./configure --enable-mysqlnd --enable-sockets --enable-openssl --enable-http2  \
        && make -s -j$(nproc) && make install \
       ) \
    && echo "extension=swoole.so" > /etc/opt/remi/${PHP_VERSION}/php.d/swoole.ini \
    # show swoole info
    && php -v \
    && php -m \
    && php --ri swoole \
    # install composer
    && ( \
        cd /tmp \
        && curl -SL "http://m.94coder.cn/composer.phar" -o composer.phar | php \
        && chmod 755 composer.phar \
        && mv composer.phar /usr/local/bin/composer \
        && composer config -g repo.packagist composer https://packagist.phpcomposer.com \
        && composer self-update \
    ) \
    # ---------- clear works ----------
    && rm -rf /tmp/* /var/cache/*  /usr/share/man /usr/share/php7 \
    && echo -e "\n\033[42;37m Build Completed :).\033[0m\n"

WORKDIR /home/www

# Override stop signal to stop process gracefully
# https://github.com/php/php-src/blob/17baa87faddc2550def3ae7314236826bc1b1398/sapi/fpm/php-fpm.8.in#L163
STOPSIGNAL SIGKILL

EXPOSE 9000

CMD ["php-fpm"]