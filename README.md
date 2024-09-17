# GO: Desafio do Cinturão

**GO: Desafio do Cinturão** é um jogo 2D em que você pilota uma nave através de uma perigosa região espacial, destruindo meteoros e tentando alcançar a maior pontuação possível. Um único impacto com um meteoro encerra a sua corrida.

![Imagem do Jogo](https://raw.githubusercontent.com/mirianrosa/go-game/main/assets/readme/screenshot.png)

## História

Em uma região remota da galáxia, existe um setor espacial inexplorado, conhecido apenas como o Cinturão da Destruição. Esse setor é um caminho cheio de meteoros imprevisíveis e perigosos que bloqueiam o acesso a vastos recursos e antigos segredos cósmicos. Exploradores e aventureiros do universo todo tentam atravessar o cinturão em busca de glória e fortuna, mas poucos retornam com sucesso.

Você é o piloto de uma nave especial, designada para enfrentar os desafios do Cinturão da Destruição. Com lasers avançados e agilidade sem igual, sua missão é simples: destruir o máximo possível de meteoros e sobreviver o maior tempo que puder. Cada meteoro dissipado rende um ponto de reconhecimento, aumentando sua reputação como o melhor piloto da galáxia.

Mas cuidado! Um único impacto com um meteoro é o fim da linha. Sua nave será destruída, e a tentativa reiniciada, com apenas seu recorde de pontuação para provar até onde você foi capaz de chegar.

Agora, todas as atenções da galáxia estão em você. Até onde consegue ir antes que os meteoros interrompam sua jornada?

## Como Jogar
- Use as setas **esquerda e direita** do teclado para mover a nave.
- Dispare lasers com a tecla **Espaço** para destruir meteoros.
- Ganhe pontos ao destruir meteoros, mas evite qualquer colisão!

## Jogar
### Online
Experimente o jogo diretamente no seu navegador clicando no link abaixo:

[**Jogar no Navegador**](http://seulink.aqui) 

### Localmente

1. Certifique-se de que o Go está instalado no seu sistema. Você pode baixá-lo [aqui](https://golang.org/doc/install).

2. Clone este repositório
```bash
git clone https://github.com/mirianrosa/go-game
```

3. Navegue até o diretório do projeto:
```bash
cd go-game
```

4. Instale as dependências necessárias usando:
```bash
go mod tidy
```

5. Inicie o jogo com o comando:
```bash
go run main.go
```
O jogo será executado em uma janela que se abrirá automaticamente.

## Fontes

- **Imagens 2D**: As imagens dos meteoros, naves, lasers, estrelas, planetas e ícones foram retiradas dos assets grátis do site [Kenney](https://kenney.nl/assets/category:2D). Os pacotes específicos utilizados incluem [Space Shooter Redux](https://kenney.nl/assets/space-shooter-redux), [Planets](https://kenney.nl/assets/planets) e [Input Prompts](https://kenney.nl/assets/input-prompts). A imagem do Gopher personalizada foi criada utilizando [Gopherize.me](https://gopherize.me/).

- **Engine**: Desenvolvido com [Ebitengine](https://github.com/hajimehoshi/ebiten), uma engine open source para criação de jogos 2D em Go.
