resource "aws_lb_listener" "flugel_alb_listener" {
  load_balancer_arn = aws_lb.flugel_alb.arn
  port              = "80"
  protocol          = "HTTP"
  default_action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.flugel.arn
  }
}

#resource "aws_lb_listener_rule" "alb_listener_rule" {
#  listener_arn = aws_lb_listener.flugel_alb_listener.arn  
#  action {    
#    type             = "forward"    
#    target_group_arn = aws_lb_target_group.flugel.arn  
#  }   
#  condition {    
#    path_pattern{    
#    values = ["/*"]  
#  }
#}
#}
