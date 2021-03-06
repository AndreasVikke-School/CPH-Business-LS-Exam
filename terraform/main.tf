provider "kubernetes" {
  config_path    = "~/.kube/config"
  config_context = "minikube"
}

provider "kubectl" {
  config_path    = "~/.kube/config"
  config_context = "minikube"
}

provider "helm" {
  kubernetes {
    config_path    = "~/.kube/config"
    config_context = "minikube"
  }
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

module "api_service" {
  source = "./modules/service"

  name_prefix            = "api-"
  namespace              = kubernetes_namespace.services.metadata.0.name
  image_name             = "api_service"
  image_version          = var.api_service_image_version
  container_port         = 8081
  container_replications = 2
  service_type           = "LoadBalancer"
  service_ports = {
    server = {
      port        = 8080,
      target_port = 8081
    }
  }
}

module "frontend_service" {
  source = "./modules/service"

  name_prefix            = "frontend-"
  namespace              = kubernetes_namespace.services.metadata.0.name
  image_name             = "frontend_service"
  image_version          = var.frontend_service_image_version
  container_port         = 3000
  container_replications = 2
  service_type           = "LoadBalancer"
  container_env = {
    NEXT_PUBLIC_API_IP = module.api_service.service_ip
  }
  service_ports = {
    server = {
      port        = 8888,
      target_port = 3000
    }
  }
}

module "consumer_service" {
  source = "./modules/service"

  name_prefix = "consumer-"
  namespace = kubernetes_namespace.services.metadata.0.name
  image_name             = "consumer_service"
  image_version          = var.consumer_service_image_version
  container_port         = 5000
  container_replications = 2
  service_type           = "ClusterIP"
  service_ports = {
    server = {
      port        = 5000,
      target_port = 5000
    }
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

  depends_on = [
    module.redis_module
  ]
}

module "postgres_service" {
  source = "./modules/service"

  name_prefix            = "postgres-"
  namespace              = kubernetes_namespace.services.metadata.0.name
  image_name             = "postgres_service"
  image_version          = var.postgres_service_image_version
  container_port         = 50051
  container_replications = 2
  service_type           = "ClusterIP"
  service_ports = {
    server = {
      port        = 50051,
      target_port = 50051
    }
  }

  depends_on = [
    module.postgres_module
  ]
}
