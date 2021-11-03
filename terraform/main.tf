provider "kubernetes" {
  config_path    = "~/.kube/config"
  config_context = "minikube"
}

resource "kubernetes_namespace" "test" {
  metadata {
    name = "test"
  }
}

resource "kubernetes_deployment" "test" {
  metadata {
    name      = "test1"
    namespace = kubernetes_namespace.test.metadata.0.name
  }
  spec {
    replicas = 1
    selector {
      match_labels = {
        app = "test1"
      }
    }
    template {
      metadata {
        labels = {
          app = "test1"
        }
      }
      spec {
        container {
          image = "ghcr.io/andreasvikke/cph-business-ls-exam/test1:latest"
          name  = "test1-container"
          port {
            container_port = 8080
          }
        }
      }
    }
  }
}

resource "kubernetes_deployment" "test2" {
  metadata {
    name      = "test2"
    namespace = kubernetes_namespace.test.metadata.0.name
  }
  spec {
    replicas = 1
    selector {
      match_labels = {
        app = "test2"
      }
    }
    template {
      metadata {
        labels = {
          app = "test2"
        }
      }
      spec {
        container {
          image = "ghcr.io/andreasvikke/cph-business-ls-exam/test2:latest"
          name  = "test2-container"
          port {
            container_port = 8080
          }
        }
      }
    }
  }
}

resource "kubernetes_service" "test" {
  metadata {
    name      = "test1"
    namespace = kubernetes_namespace.test.metadata.0.name
  }
  spec {
    selector = {
      app = kubernetes_deployment.test.spec.0.template.0.metadata.0.labels.app
    }
    type = "ClusterIP"
    port {
      port        = 80
      target_port = 8080
    }
  }
}

resource "kubernetes_service" "test2" {
  metadata {
    name      = "test2"
    namespace = kubernetes_namespace.test.metadata.0.name
  }
  spec {
    selector = {
      app = kubernetes_deployment.test2.spec.0.template.0.metadata.0.labels.app
    }
    type = "LoadBalancer"
    port {
      port        = 80
      target_port = 8080
    }
  }
}