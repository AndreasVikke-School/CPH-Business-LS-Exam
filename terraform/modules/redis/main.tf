
resource "kubernetes_namespace" "redis" {
  metadata {
    name = "redis"
  }
}

resource "kubernetes_config_map" "redis_cluster" {
  metadata {
    name      = "redis-cluster"
    namespace = kubernetes_namespace.redis.metadata.0.name
  }

  data = {
    "update-node.sh" = <<-EOF
        #!/bin/sh
        REDIS_NODES="/data/nodes.conf"
        sed -i -e "/myself/ s/[0-9]\{1,3\}\.[0-9]\{1,3\}\.[0-9]\{1,3\}\.[0-9]\{1,3\}/$${POD_IP}/" $${REDIS_NODES}
        exec "$@"
    EOF
    "redis.conf"     = <<-EOF
        cluster-enabled yes
        cluster-require-full-coverage no
        cluster-node-timeout 15000
        cluster-config-file /data/nodes.conf
        cluster-migration-barrier 1
        appendonly yes
        protected-mode no
    EOF
  }
}

resource "kubernetes_stateful_set" "redis_cluster" {
  metadata {
    name = "redis-cluster"
  }

  spec {
    replicas     = 6
    service_name = "redis-cluster"

    selector {
      match_labels = {
        app = "redis-cluster"
      }
    }

    template {
      metadata {
        labels = {
          app = "redis-cluster"
        }
      }
      spec {
        container {
          name    = "redis"
          image   = "redis:6-alpine"
          command = ["/conf/update-node.sh", "redis-server", "/conf/redis.conf"]

          port {
            container_port = 6379
            name           = "client"
          }
          port {
            container_port = 16379
            name           = "gossip"
          }

          env {
            name = "POD_IP"
            value_from {
              field_ref {
                field_path = "status.podIP"
              }
            }
          }

          volume_mount {
            name       = "conf"
            mount_path = "/conf"
            read_only  = false
          }
          volume_mount {
            name       = "data"
            mount_path = "/data"
            read_only  = false
          }
        }
        volume {
          name = "conf"

          config_map {
            name = "redis-cluster"
          }
        }
      }
    }

    volume_claim_template {
      metadata {
        name = "data"
      }

      spec {
        access_modes = ["ReadWriteOnce"]

        resources {
          requests = {
            storage = "1Gi"
          }
        }
      }
    }
  }
}

resource "kubernetes_service" "redis_cluster" {
  metadata {
    name      = "redis-cluster"
    namespace = kubernetes_namespace.redis.metadata.0.name
  }
  spec {
    selector = {
      app = "redis-cluster"
    }
    type = "LoadBalancer"
    port {
      port = 6379
      name = "client"
    }
    port {
      port = 16379
      name = "gossip"
    }
  }
}

# # ==== REDIS SERVICE ====
# resource "kubernetes_deployment" "redis" {
#   metadata {
#     name      = "redis"
#     namespace = kubernetes_namespace.redis.metadata.0.name
#   }
#   spec {
#     replicas = 1
#     selector {
#       match_labels = {
#         app = "redis"
#       }
#     }
#     template {
#       metadata {
#         labels = {
#           app = "redis"
#         }
#       }
#       spec {
#         container {
#           image = "redis:6"
#           name  = "redis"
#           port {
#             container_port = 6379
#           }
#           env {
#             name  = "POSTGRES_USER"
#             value = var.postgress_username
#           }
#           env {
#             name  = "POSTGRES_PASSWORD"
#             value = var.postgress_password
#           }
#         }
#       }
#     }
#   }
# }

# resource "kubernetes_service" "postgres" {
#   metadata {
#     name      = "postgres"
#     namespace = kubernetes_namespace.postgres.metadata.0.name
#   }
#   spec {
#     selector = {
#       app = "postgres"
#     }
#     type = "ClusterIP"
#     port {
#       port = 5432
#     }
#   }
# }
# # ==== POSTGRESS SERVICE END ====

# # ==== PGADMIN SERVICE ====
# resource "kubernetes_deployment" "pgadmin" {
#   metadata {
#     name      = "pgadmin"
#     namespace = kubernetes_namespace.postgres.metadata.0.name
#   }
#   spec {
#     replicas = 1
#     selector {
#       match_labels = {
#         app = "pgadmin"
#       }
#     }
#     template {
#       metadata {
#         labels = {
#           app = "pgadmin"
#         }
#       }
#       spec {
#         container {
#           image = "dpage/pgadmin4"
#           name  = "postgres"
#           port {
#             container_port = 80
#           }
#           env {
#             name  = "PGADMIN_DEFAULT_EMAIL"
#             value = var.postgress_email
#           }
#           env {
#             name  = "PGADMIN_DEFAULT_PASSWORD"
#             value = var.postgress_password
#           }
#         }
#       }
#     }
#   }
# }

# resource "kubernetes_service" "pgadmin" {
#   metadata {
#     name      = "pgadmin"
#     namespace = kubernetes_namespace.postgres.metadata.0.name
#   }
#   spec {
#     selector = {
#       app = "pgadmin"
#     }
#     type = "LoadBalancer"
#     port {
#       port        = 5000
#       target_port = 80
#     }
#   }
# }
# # ==== PGADMIN END ====
