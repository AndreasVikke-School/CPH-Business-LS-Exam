# locals {
#   broker_env = {
#     "KAFKA_LISTENERS"                        = "INTERNAL://kafka-broker:29092,EXTERNAL://kafka-broker:9092"
#     "KAFKA_ADVERTISED_LISTENERS"             = "INTERNAL://kafka-broker:29092,EXTERNAL://kafka-broker:9092"
#     "KAFKA_ZOOKEEPER_CONNECT"                = "kafka-zookeeper:2181"
#     "KAFKA_LISTENER_SECURITY_PROTOCOL_MAP"   = "INTERNAL:PLAINTEXT,EXTERNAL:PLAINTEXT"
#     "KAFKA_INTER_BROKER_LISTENER_NAME"       = "INTERNAL"
#     "KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR" = "1"
#     "KAFKA_AUTO_CREATE_TOPICS_ENABLE"        = "true"
#   }
# }


resource "kubectl_manifest" "zookeeper" {
    yaml_body = file("${path.module}/zoo.yaml")
}

resource "kubectl_manifest" "kafka" {
    yaml_body = file("${path.module}/kafka.yaml")

    depends_on = [
      kubectl_manifest.zookeeper
    ]
}

resource "kubernetes_service" "kafka" {
  metadata {
    name      = "kafka"
    namespace = kubernetes_namespace.kafka.metadata.0.name
  }
  spec {
    selector = {
      app = "kafka"
    }
    type = "LoadBalancer"
    port {
      name = "kafka"
      port = 9094
    }
  }
}

resource "kubernetes_service" "zookeeper" {
  metadata {
    name      = "zookeeper"
    namespace = kubernetes_namespace.kafka.metadata.0.name
  }
  spec {
    selector = {
      app = "zookeeper"
    }
    type = "ClusterIP"
    port {
      name = "client"
      port = 2181
    }
    port {
      name = "server"
      port = 2888
    }
    port {
      name = "leader-election"
      port = 3888
    }
  }
}


resource "kubernetes_namespace" "kafka" {
  metadata {
    name = "kafka"
  }
}

# # ==== KAFKA ZOOKEEPER ====
# resource "kubernetes_deployment" "kafka_zookeeper" {
#   metadata {
#     name      = "kafka-zookeeper"
#     namespace = kubernetes_namespace.kafka.metadata.0.name
#   }
#   spec {
#     replicas = 2
#     selector {
#       match_labels = {
#         app = "kafka-zookeeper"
#       }
#     }
#     template {
#       metadata {
#         labels = {
#           app = "kafka-zookeeper"
#         }
#       }
#       spec {
#         container {
#           image = "confluentinc/cp-zookeeper"
#           name  = "kafka-zookeeper"
#           port {
#             container_port = 2181
#           }
#           env {
#             name  = "ZOOKEEPER_CLIENT_PORT"
#             value = "2181"
#           }
#         }
#       }
#     }
#   }
# }

# resource "kubernetes_service" "kafka_zookeeper" {
#   metadata {
#     name      = "kafka-zookeeper"
#     namespace = kubernetes_namespace.kafka.metadata.0.name
#   }
#   spec {
#     selector = {
#       app = "kafka-zookeeper"
#     }
#     type = "ClusterIP"
#     port {
#       port = 2181
#     }
#   }
# }
# # ==== KAFKA ZOOKEEPER END ====

# # ==== KAFKA BROKER ====
# resource "kubernetes_deployment" "kafka_broker" {
#   metadata {
#     name      = "kafka-broker"
#     namespace = kubernetes_namespace.kafka.metadata.0.name
#   }
#   spec {
#     replicas = 1
#     selector {
#       match_labels = {
#         app = "kafka-broker"
#       }
#     }
#     template {
#       metadata {
#         labels = {
#           app = "kafka-broker"
#         }
#       }
#       spec {
#         container {
#           image = "confluentinc/cp-kafka:7.0.0"
#           name  = "kafka-broker"
#           port {
#             container_port = 9092
#           }
#           dynamic "env" {
#             for_each = local.broker_env
#             content {
#               name  = env.key
#               value = env.value
#             }
#           }
#         }
#       }
#     }
#   }
# }

# resource "kubernetes_service" "kafka_broker" {
#   metadata {
#     name      = "kafka-broker"
#     namespace = kubernetes_namespace.kafka.metadata.0.name
#   }
#   spec {
#     selector = {
#       app = "kafka-broker"
#     }
#     type = "LoadBalancer"
#     port {
#       name = "kafka-external"
#       port = 9092
#     }
#     port {
#       name = "kafka-internal"
#       port = 19092
#     }
#   }
# }
# # ==== KAFKA BROKER END ====

# module "kafka" {
#   source  = "skydome/kafka/kubernetes"
#   version = "0.0.10"

#   cluster_size = 2
#   kafka_name   = "kafka"
#   namespace    = kubernetes_namespace.kafka.metadata.0.name
# }

# ==== KAFDROP ====
resource "kubernetes_deployment" "kafka_kafdrop" {
  metadata {
    name      = "kafka-kafdrop"
    namespace = kubernetes_namespace.kafka.metadata.0.name
  }
  spec {
    replicas = 2
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
            value = "10.111.221.127:9094"
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