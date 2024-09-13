<script>
  import Block from './lib/Block.svelte';
</script>

<Block weapon="phantom"></Block>
<footer>RIOT GAMES, VALORANT and any associated logos are trademarks, service marks, and/or registered trademarks of <a href="https://riotgames.com">Riot Games, Inc.</a></footer>

<style>
    footer {
        font-family: Inter, system-ui, Avenir, Helvetica, Arial, sans-serif;
        text-align: center;
        padding-top: 10px;
        color: white;
        width: 100%;
        bottom: 0;
    }

    a {
        color: inherit;
        text-decoration: underline;
    }

    :root {
        background-attachment: fixed;
        background-image: url("./assets/bg_main.png");
    }
</style>