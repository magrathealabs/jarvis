FROM rabbitmq:3.6.1-management

RUN rabbitmq-plugins enable --offline rabbitmq_mqtt

EXPOSE 1883
