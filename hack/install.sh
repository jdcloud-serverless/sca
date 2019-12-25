mkdir ~/.sca
wget -P ~/.sca/ https://github.com/jdcloud-serverless/sca/releases/download/v0.0.1/sca

chmod +777 ~/.sca/sca

echo "PATH=$PATH:~/.sca" >> ~/.bashrc
echo "export PATH" >> ~/.bashrc