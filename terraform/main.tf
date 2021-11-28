provider "kubernetes" {
  config_path    = "~/.kube/config"
  config_context = "minikube"
}

module "kafka_module" {
  source = "./modules/kafka"
}
module "postgres_module" {
  source = "./modules/postgres"
}

module "redis_module" {
  source = "./modules/redis"
}
resource "kubernetes_namespace" "services" {
  metadata {
    name = "services"
  }
}

module "redis_service" {
  source = "./modules/service"

  name_prefix            = "redis-"
  namespace              = kubernetes_namespace.services.metadata.0.name
  image_name             = "redis_service"
  image_version          = var.redis_service_image_version
  container_port         = 50051
  container_replications = 2
  service_type           = "ClusterIP"
  service_ports = {
    server = {
      port        = 50051,
      target_port = 50051
    }
  }
}

module "api_service" {
  source = "./modules/service"

  name_prefix            = "api-"
  namespace              = kubernetes_namespace.services.metadata.0.name
  image_name             = "api_service"
  image_version          = var.api_service_image_version
  container_port         = 8080
  container_replications = 2
  service_type           = "LoadBalancer"
  service_ports = {
    server = {
      port        = 80,
      target_port = 8080
    }
  }
}

# resource "kubernetes_deployment" "redis_service" {
#   metadata {
#     name      = "redis-service"
#     # namespace = kubernetes_namespace.test.metadata.0.name
#     namespace = "redis"
#   }
#   spec {
#     replicas = 1
#     selector {
#       match_labels = {
#         app = "redis-service"
#       }
#     }
#     template {
#       metadata {
#         labels = {
#           app = "redis-service"
#         }
#       }
#       spec {
#         container {
#           image = "ghcr.io/andreasvikke/cph-business-ls-exam/redis_service:${var.redis_service_image_version}"
#           name  = "redis-service-container"
#           port {
#             container_port = 50051
#           }
#         }
#       }
#     }
#   }
# }

# resource "kubernetes_deployment" "api" {
#   metadata {
#     name      = "api"
#     namespace = kubernetes_namespace.test.metadata.0.name
#   }
#   spec {
#     replicas = 1
#     selector {
#       match_labels = {
#         app = "api"
#       }
#     }
#     template {
#       metadata {
#         labels = {
#           app = "api"
#         }
#       }
#       spec {
#         container {
#           image = "ghcr.io/andreasvikke/cph-business-ls-exam/api_service:${var.api_service_image_version}"
#           name  = "api-container"
#           port {
#             container_port = 8080
#           }
#         }
#       }
#     }
#   }
# }

# resource "kubernetes_service" "service_redis" {
#   metadata {
#     name      = "redis-service"
#     namespace = "redis"
#   }
#   spec {
#     selector = {
#       app = "redis-service"
#     }
#     type = "ClusterIP"
#     port {
#       port        = 50051
#       target_port = 50051
#     }
#   }
# }

# resource "kubernetes_service" "api" {
#   metadata {
#     name      = "api"
#     namespace = kubernetes_namespace.test.metadata.0.name
#   }
#   spec {
#     selector = {
#       app = kubernetes_deployment.api.spec.0.template.0.metadata.0.labels.app
#     }
#     type = "LoadBalancer"
#     port {
#       port        = 80
#       target_port = 8080
#     }
#   }
# }