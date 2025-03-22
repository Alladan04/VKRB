# Используем базовый образ Debian
FROM --platform=linux/amd64 debian:latest

# Устанавливаем необходимые пакеты
RUN apt-get update && apt-get install -y \
    gdebi-core nano systemd \
    && rm -rf /var/lib/apt/lists/* && apt-get install dpkg

# Копируем deb-файл в контейнер
COPY WiMi-0.1.7.deb /tmp/WiMi-0.1.7.deb

# Устанавливаем WiMi
RUN dpkg --add-architecture amd64 && dpkg -i /tmp/WiMi-0.1.7.deb || apt-get -f install -y

# Создаём директорию и добавляем конфиг MivREST.ini
RUN mkdir -p /usr/local/bin/WiMi/ \
    && echo "[listener]\nport=8092\nminThreads=10\nmaxThreads=100" > /usr/local/bin/WiMi/MivREST.ini

# Создаём environment-файл
RUN echo 'LD_LIBRARY_PATH="$LD_LIBRARY_PATH:/usr/local/bin/WiMi/libs"' > /usr/local/bin/WiMi/environment

COPY WiMi-0.1.7.deb /tmp/WiMi-0.1.7.deb

#Добавить файл Mivar.Service (кажется можно завести на хосте и просто скопировать в директорию в контйнере)
COPY WiMi.service /usr/local/bin/WiMi/

RUN ln /usr/local/bin/WiMi/WiMi.service /etc/systemd/system/
# Добавляем права на запуск
RUN chmod +x /usr/local/bin/WiMi/WiMi.run
#
## Открываем порт 8092
EXPOSE 8092
#
## Устанавливаем рабочую директорию
WORKDIR /usr/local/bin/WiMi

RUN apt update && apt-get install -y libglib2.0-0 libsm6 libxrender1 libxext6

## Запуск WiMi при старте контейнера
CMD ["./WiMi.run"]
