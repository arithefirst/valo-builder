<script>
  import Item from './lib/Item.svelte';

  const defaults = new Map()
  defaults.set("phantom", "")

</script>

<div class="outer">
    <div class="sidearms">
        <Item weapon="classic" url="blank"/>
        <Item weapon="shorty" url="blank"/>
        <Item weapon="frenzy" url="blank"/>
        <Item weapon="ghost" url="blank"/>
        <Item weapon="sheriff" url="blank"/>
    </div>

    <div class="smgs-shotguns">
        <Item weapon="stinger" url="blank"/>
        <Item weapon="spectre" url="blank"/>
        <p/>
        <Item weapon="bucky" url="blank"/>
        <Item weapon="judge" url="blank"/>
    </div>

    <div class="rifles-melee">
        <Item weapon="bulldog" url="blank"/>
        <Item weapon="guardian" url="blank"/>
        <Item weapon="phantom" url="blank"/>
        <Item weapon="vandal" url="blank"/>
        <p/>
        <Item weapon="melee" url="blank"/>
    </div>

    <div class="snipers-lmgs">
        <Item weapon="marshal" url="blank"/>
        <Item weapon="outlaw" url="blank"/>
        <Item weapon="operator" url="blank"/>
        <p/>
        <Item weapon="ares" url="blank"/>
        <Item weapon="odin" url="blank"/>
    </div>
</div>
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

    .outer{display: flex;}

    .sidearms {
        margin-top: 2.75330396476vw;
        margin-left: 4.40528634361vw;
        width: 16.519823788vw;
    }

    .smgs-shotguns, .rifles-melee, .snipers-lmgs{
        margin-top: 2.75330396476vw;
        margin-left: 2.20264317181vw;
        width: 16.519823788vw;
    }

    p {
        margin-bottom: 3.74449339207vw
    }

    :root {
        background-attachment: fixed;
        background-image: url("./assets/bg_main.png");
    }
</style>