provider "aws" {
  region     = "eu-north-1"
  access_key = "xxx"
  secret_key = "xxx"
}

resource "aws_instance" "ec2_example" {
  ami = "ami-0b1ccb69e70ec2b9a"
  instance_type = "t3.micro"
  key_name= "DOCKER-KEY"
  tags = {
    Name = "Only-Posts"
  }
  connection {
    type        = "ssh"
    host        = self.public_ip
    user        = "ubuntu"
    private_key = file("path_to_file")
    timeout     = "4m"
  }
}
