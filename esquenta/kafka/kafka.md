# kafka

streaming de eventos
streaming analytics
data integration
mission critical applications

eventos
restaura historico de eventos
significa que consegue recriar o estado atual da aplicacao refazendo o historico de eventos

capacidades
alto throughput -> aguenta porrada
escalavel
historico permanente
alta disponibilidade

ecosistema

companies
nasce no linkedin
netflix
uber
twitter
dropbox
spotify
paypal
banks

events
notificacoes -> informa que algo mudou
streaming -> traz todas informacoes do eventos

persistencia em disco e nao em memoria

funcionamento geral
kafka -> conjunto de maquinas
brokers -> cada maquina que esta rodando no kafka
producer -> envia dados para kafka
zookeeper -> error handling/cluster manager/permissions manager \* nao vai mais ser usado

topic -> lugar onde as mensagens sao recebidas
particao ->
topico funciona com log
cada broker pode ter varias maquinas para garantir acessos

replication factor -> qtd de replicas distribuida pelos brokers
partition leadership -> quando uma particao for lider em um broker quando for ler tal particao sera lido no broker em que ela e lider
consumer -> sistema interessado em recber mensagens (ler todas particoes) 100 mensagens por segundo
consumer group -> adiciona mais replicas para receber mais mensagens por segundo
