output "public_ip" {
  value = aws_instance.dev.public_ip
}

output "tags" {
  value = aws_instance.dev.tags
}

output "instance_id"{
value = aws_instance.dev.id
}

