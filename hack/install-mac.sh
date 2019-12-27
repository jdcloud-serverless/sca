mkdir ~/.sca
wget -P ~/.sca/sca https://github.com/jdcloud-serverless/sca/releases/download/v0.0.1/sca-mac

chmod +x ~/.sca/sca

echo PATH='$PATH':~/.sca >> ~/.bashrc
echo export PATH >> ~/.bashrc