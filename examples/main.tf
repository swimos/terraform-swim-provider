terraform {
  required_providers {
    swim = {
      version = "0.1"
      source  = "swim.inc/swim/swim"
    }

    docker = {
      source = "kreuzwerker/docker"
      version = "~> 2.21.0"
    }
  }
}

# Docker
provider "docker" {}

resource "docker_image" "nginx" {
  name         = "nginx:latest"
  keep_locally = false
}

resource "docker_container" "nginx" {
  image = docker_image.nginx.image_id
  name  = "tutorial"
  ports {
    internal = 80
    external = 8000
  }
}

# Swim
provider "swim" {
  url = "ws://127.0.0.1:9001/"
}

resource "swim_value_downlink" "id" {
  node = "/container"
  lane = "id"
  value = docker_container.nginx.id
}

resource "swim_map_downlink" "ports" {
  node = "/container"
  lane = "ports"
  items = {
    "external" = docker_container.nginx.ports[0].external
    "internal" = docker_container.nginx.ports[0].internal
  }
}

output "id" {
  value = swim_value_downlink.id
}

output "ports" {
  value = swim_map_downlink.ports
}