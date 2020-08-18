FROM centos:7

RUN set -e \
    && yum install -y openssh-server \
    && yum install -y git \
    && yum install -y make \
    && yum install -y glibc-headers \
    && yum install -y gcc-c++ \
    && mkdir /var/run/sshd  \
    && ssh-keygen -t rsa -f /etc/ssh/ssh_host_rsa_key  \
    && ssh-keygen -t dsa -f /etc/ssh/ssh_host_dsa_key  \
    && ssh-keygen -t ed25519 -f /etc/ssh/ssh_host_ed25519_key  \
    && ssh-keygen -t ecdsa -f /etc/ssh/ssh_host_ecdsa_key \
#    && /bin/echo 'root:123' |chpasswd \
#    && useradd tpythoner  \
#    && /bin/echo 'tpythoner:123' |chpasswd   \
    && /bin/sed -i 's/.*session.*required.*pam_loginuid.so.*/session optional pam_loginuid.so/g' /etc/pam.d/sshd   \
    && /bin/echo -e "LANG=\"en_US.UTF-8\"" > /etc/default/local   \
    && yum install -y wget  \
    && cd /tmp \
    && wget -O go.tar.gz https://dl.google.com/go/go1.12.2.linux-amd64.tar.gz   \
    && tar zxvf go.tar.gz -C /usr/local/   \
    && echo "export GOROOT=/usr/local/go" >> /root/.bashrc   \
    && echo "export GO111MODULE=on" >> /root/.bashrc   \
    && echo "export GOPROXY=https://goproxy.io" >> /root/.bashrc   \
    && echo "export GOBIN=\$GOROOT/bin" >> /root/.bashrc   \
    && echo "export PATH=\$PATH:\$GOBIN" >> /root/.bashrc   \
    && echo "export GOPATH=/home/golang" >> /root/.bashrc \
    && echo "export TZ=\"Asia/Shanghai\"" >> /root/.bashrc \
    && wget -O protobuf.tar.gz https://github.com/google/protobuf/releases/download/v3.6.1/protobuf-all-3.6.1.tar.gz \
    && tar zvxf  protobuf.tar.gz -C . \
    && cd protobuf-3.6.1 && ./configure && make && make install \
    && yum remove -y wget \
    && rm -rf /tmp/* /var/cache/* /usr/share/man /usr/share/go \
    && echo -e "\n\033[42;37m Build Completed :).\033[0m\n"

#开发端口
EXPOSE  22
EXPOSE  9888

#启动sshd服务
CMD     /usr/sbin/sshd -D
