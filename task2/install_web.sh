#! /bin/bash
sudo apt update
sudo apt install -y apache2 software-properties-common
sudo add-apt-repository ppa:deadsnakes/ppa -y
sudo apt update
sudo apt install -y python3
{
        echo 'f = open("index.html","w")'
        echo 'message = "Name : Flugel"'
#        echo 'message = """<html>'
#        echo '<body><p>Name : Flugel</p>'
#        echo '<p> Owner: InfraTeam</p>'  
#        echo '</body>'
#        echo '</html>"""'
        echo 'f.write(message)'
        echo 'f.close() '
} >> html.py
python3 html.py
sudo cp index.html /var/www/html/index.html
sudo systemctl restart apache2
sudo systemctl enable apache2
