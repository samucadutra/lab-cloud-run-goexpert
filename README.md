# Lab Cloud Run GoExpert

## Executando o Projeto com Docker Compose

Siga estas etapas para executar o projeto usando `docker-compose`:

1. Certifique-se de que o Docker e o Docker Compose estão instalados no seu sistema.

2. Defina a variável de ambiente `WEATHER_API_KEY` com sua chave de API:
   ```bash
   export WEATHER_API_KEY=sua_chave_de_api_aqui
   ```

3. Construa e inicie a aplicação usando `docker-compose`:
   ```bash
   docker-compose up --build
   ```

4. Faça um request para a aplicação em `http://localhost:8080/weather/current/<zipcode>`.

5. Para parar a aplicação, pressione `Ctrl+C` e execute:
   ```bash
   docker-compose down
   ```

6. Você pode acessar a aplicação diretamente no Google Cloud Run pelo endereço:
   ```
   https://lab-cloud-run-goexpert-j2xraeslja-uc.a.run.app/weather/current/<zipcode>
   ```
