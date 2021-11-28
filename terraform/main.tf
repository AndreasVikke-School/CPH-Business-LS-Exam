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
resource "kubernetes_namespace" "test" {
  metadata {
    name = "test"
  }
}

resource "kubernetes_deployment" "redis_service" {
  metadata {
    name      = "redis-service"
    # namespace = kubernetes_namespace.test.metadata.0.name
    namespace = "redis"
  }
  spec {
    replicas = 1
    selector {
      match_labels = {
        app = "redis-service"
      }
    }
    template {
      metadata {
        labels = {
          app = "redis-service"
        }
      }
      spec {
        container {
          image = "ghcr.io/andreasvikke/cph-business-ls-exam/redis_service:${var.redis_service_image}"
          name  = "redis-service-container"
          port {
            container_port = 50051
          }
        }
      }
    }
  }
}

resource "kubernetes_deployment" "api" {
  metadata {
    name      = "api"
    namespace = kubernetes_namespace.test.metadata.0.name
  }
  spec {
    replicas = 1
    selector {
      match_labels = {
        app = "api"
      }
    }
    template {
      metadata {
        labels = {
          app = "api"
        }
      }
      spec {
        container {
          image = "ghcr.io/andreasvikke/cph-business-ls-exam/api_service:${var.api_service_image}"
          name  = "api-container"
          port {
            container_port = 8080
          }
        }
      }
    }
  }
}

resource "kubernetes_service" "service_redis" {
  metadata {
    name      = "redis-service"
    namespace = "redis"
  }
  spec {
    selector = {
      app = "redis-service"
    }
    type = "ClusterIP"
    port {
      port        = 50051
      target_port = 50051
    }
  }
}

resource "kubernetes_service" "api" {
  metadata {
    name      = "api"
    namespace = kubernetes_namespace.test.metadata.0.name
  }
  spec {
    selector = {
      app = kubernetes_deployment.api.spec.0.template.0.metadata.0.labels.app
    }
    type = "LoadBalancer"
    port {
      port        = 80
      target_port = 8080
    }
  }
}