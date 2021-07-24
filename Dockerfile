FROM alpine:3.14.0

LABEL maintainer="dreamcmi@openthing.cn"

###############################################################################
#                                INSTALLATION
###############################################################################

# 设置固定的项目路径
ENV WORKDIR /var/www/ipReq-http

# 添加应用可执行文件，并设置执行权限
ADD ./main   $WORKDIR/main
RUN chmod +x $WORKDIR/main

# 添加I18N多语言文件、静态文件、配置文件、模板文件
ADD i18n     $WORKDIR/i18n
ADD public   $WORKDIR/public
ADD config   $WORKDIR/config
ADD template $WORKDIR/template

###############################################################################
#                                   START
###############################################################################
WORKDIR $WORKDIR
EXPOSE 5555
CMD ./main
