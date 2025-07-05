
1) Atualize os pacotes
sudo apt update
sudo apt upgrade -y

2) Instale dependências
sudo apt install -y ca-certificates curl gnupg

3) Adicione a chave oficial do Docker
sudo install -m 0755 -d /etc/apt/keyrings
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | \
  sudo gpg --dearmor -o /etc/apt/keyrings/docker.gpg

4) Adicione o repositório do Docker
echo \
  "deb [arch=$(dpkg --print-architecture) \
  signed-by=/etc/apt/keyrings/docker.gpg] \
  https://download.docker.com/linux/ubuntu \
  $(lsb_release -cs) stable" | \
  sudo tee /etc/apt/sources.list.d/docker.list > /dev/null

5) Instale o Docker Engine + Compose
sudo apt update
sudo apt install -y docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin

6) Verifique se o Docker funciona
docker --version
docker compose version

7) (Opcional) Use Docker sem sudo
sudo usermod -aG docker $USER
newgrp docker  # ou faça logout/login depois

8) Teste com um container
docker run hello-world

