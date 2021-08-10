resource "aws_lb" "flugel_alb" {
  name               = "flugel"
  internal           = false
  load_balancer_type = "application"
  security_groups    = [aws_security_group.flugel_web_sg.id]
#  subnets            = [aws_subnet.flugel.*.id]
subnet_mapping {
    subnet_id            = aws_subnet.flugel_public_a.id
  }
subnet_mapping {
    subnet_id            = aws_subnet.flugel_public_b.id
  }
  enable_deletion_protection = false

  tags = var.tags
}
output "dns_name" {
  value = aws_lb.flugel_alb.dns_name
}
