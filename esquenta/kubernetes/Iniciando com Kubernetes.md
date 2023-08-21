# Iniciando com Kubernetes

live 04 - 15/08 - 20h

# Kubernetes

    - capacidade de entender como as aplicacoes escalam
    - containers

    - servico de pagamento/container
    - usuario vai cair no servico de pagamento
    - se container sai do ar
        - parar de direcionar o fluxo
            - alguem que gerencie
        - usuario para de ter resposta
        * pode colocar mais containers rodando
            - evita ponto unico de falhas
            - consigo destribuir a carga

    ## Kubernetes (K8S) gerencia
        * considerado Sistema Operacional da nuvem
        - virou o jogo de como gerencia aplicacao em producao
        - como gerenciar milhares de containers
        - quantas replicas de um servico estou rodando
        - como saber quantos containers preciso escalar
        - se o trafego aumenta, cria mais containers
        - diversos containers na mesma maquina

        - necessidades:
            - variaveis de ambiente
            - gerenciamento de senhas / secrets
            - escolher os recursos computacionais da aplicacao
            - health check
            - load balancing
            - SSL / TLS
            - dominio (www) -> determinado servico
            - estrategias de deploy (blue/green, canary, etc)
            - storage
            - service discovery / DNS

        - estrutura de uma aplicacao que roda no K8S
            - deployment
                - replica set -> n de pods
                    - POD - instancia de kubernetes
                        - aplicacao (container)
                        - label -> qntd pods com determinada label

        - cluster -> conjunto de maquinas

        - simula kubernetes
            - kind -> roda kubernetes atraves do docker
            - kind create cluster
                - kubectl cluster-info --context kind-kind
            - kubectl get nodes
            - criar manifesto
                - pod.yaml
                    apiVersion: 1
                    kind: Pod
                    metadata:
                      - name: nginx
                        image: nginx:latest
                        ports:
                            - containerPort: 80

            - kubectl apply -f pod.yaml
            - kubectl get pods
            - kubectl delete pod nginx
            - replicaset.yaml
                apiVersoin: apps/v1
                kind: ReplicaSet
                metadata:
                    name: nginx
                spec:
                    selector:
                        matchLabels:
                            app: nginx
                    template:
                        metadata:
                        - name: nginx
                            image: nginx:latest
                            ports:
                                - containerPort: 80
                        spec:
                            containers:
                              - name: my app
                                image: <Image>
                                resources:
                                    limits:
                                        memory: "128Mi"
                                        cpu:
                                            "500m"
                                        ports:
            - kubectl apply -f replicaset.yaml

        - acessar containers no kubernetes
            - kubectl port-foward pod/nginx-asodaosd 8080:80

        - mudar a versao do servico
            - mudar a imagem no replicaset
                - com versao especifica

            - deletar pod e ele cria novo com versao nova
            - deployment
                mesma config do replicaset so muda kind: Deployment
                ele gerencia mudancas de versao
                    apaga replicaset antigo
                    cria replicaset com nova versao

        ## Balanceamento de carga
            - service
                - usuario cai no service
                    -service tem instrucao
                        label: app: xxxxx
                    - pega todos o pods com aquela label
                    - faz o balanceamento de carga
