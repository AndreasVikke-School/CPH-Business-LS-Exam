resource "kubernetes_namespace" "kafka" {
  metadata {
    name = "kafka"
  }
}

# ==== KAFKA ZOOKEEPER ====
resource "kubernetes_deployment" "kafka_zookeeper" {
  metadata {
    name      = "kafka-zookeeper"
    namespace = kubernetes_namespace.kafka.metadata.0.name
  }
  spec {
    replicas = 2
    selector {
      match_labels = {
        app = "kafka-zookeeper"
      }
    }
    template {
      metadata {
        labels = {
          app = "kafka-zookeeper"
        }
      }
      spec {
        container {
          image = "confluentinc/cp-zookeeper"
          name  = "kafka-zookeeper"
          port {
            container_port = 2181
          }
          env {
            name  = "ZOOKEEPER_CLIENT_PORT"
            value = "2181"
          }
        }
      }
    }
  }
}

resource "kubernetes_service" "kafka_zookeeper" {
  metadata {
    name      = "kafka-zookeeper"
    namespace = kubernetes_namespace.kafka.metadata.0.name
  }
  spec {
    selector = {
      app = "kafka-zookeeper"
    }
    type = "LoadBalancer"
    port {
      port = 2181
    }
  }
}
# ==== KAFKA ZOOKEEPER END ====

# ==== KAFKA BROKER ====
resource "kubernetes_deployment" "kafka_broker" {
  metadata {
    name      = "kafka-broker"
    namespace = kubernetes_namespace.kafka.metadata.0.name
  }
  spec {
    replicas = 1
    selector {
      match_labels = {
        app = "kafka-broker"
      }
    }
    template {
      metadata {
        labels = {
          app = "kafka-broker"
        }
      }
      spec {
        container {
          image = "confluentinc/cp-kafka:7.0.0"
          name  = "kafka-broker"
          port {
            container_port = 9092
          }
          env {
            name  = "KAFKA_ADVERTISED_LISTENERS"
            value = "INTERNAL://kafka-broker:19092,EXTERNAL://kafka-broker:9092"
          }
          env {
            name  = "KAFKA_ZOOKEEPER_CONNECT"
            value = "kafka-zookeeper:2181"
          }
          env {
            name  = "KAFKA_LISTENER_SECURITY_PROTOCOL_MAP"
            value = "INTERNAL:PLAINTEXT,EXTERNAL:PLAINTEXT"
          }
          env {
            name  = "KAFKA_INTER_BROKER_LISTENER_NAME"
            value = "INTERNAL"
          }
        }
      }
    }
  }
}

resource "kubernetes_service" "kafka_broker" {
  metadata {
    name      = "kafka-broker"
    namespace = kubernetes_namespace.kafka.metadata.0.name
  }
  spec {
    selector = {
      app = "kafka-broker"
    }
    type = "LoadBalancer"
    port {
      name      = "kafka-port"
      port      = 9092
    }
  }
}
# ==== KAFKA BROKER END ====

# ==== KAFDROP ====
resource "kubernetes_deployment" "kafka_kafdrop" {
  metadata {
    name      = "kafka-kafdrop"
    namespace = kubernetes_namespace.kafka.metadata.0.name
  }
  spec {
    replicas = 1
    selector {
      match_labels = {
        app = "kafka-kafdrop"
      }
    }
    template {
      metadata {
        labels = {
          app = "kafka-kafdrop"
        }
      }
      spec {
        container {
          image = "obsidiandynamics/kafdrop"
          name  = "kafka-broker"
          port {
            container_port = 9000
          }
          env {
            name  = "KAFKA_BROKERCONNECT"
            value = "kafka-broker:19092"
          }
        }
      }
    }
  }
}

resource "kubernetes_service" "kafka_kafdrop" {
  metadata {
    name      = "kafka-kafdrop"
    namespace = kubernetes_namespace.kafka.metadata.0.name
  }
  spec {
    selector = {
      app = "kafka-kafdrop"
    }
    type = "LoadBalancer"
    port {
      name = "kafdrop-port"
      port = 9000
    }
  }
}
# ==== KAFDROP END ====