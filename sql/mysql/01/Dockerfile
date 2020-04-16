from ubuntu
WORKDIR /
ADD https://github.com/datacharmer/test_db/archive/master.zip /
#ADD ./master.zip /
RUN apt-get update
RUN apt-get install zip -y
RUN unzip master.zip -d .

from mysql:5.7
ENV MYSQL_ALLOW_EMPTY_PASSWORD=yes
COPY --from=0 /test_db-master /test_db-master
EXPOSE 3306

#тут падает с ошибкой, пока не разобрался
#CMD ["echo", "use the 'mysql < /test_db-master/employees.sql' command to init test data"]
