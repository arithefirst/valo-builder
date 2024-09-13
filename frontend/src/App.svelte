<script>
  import Item from './lib/Item.svelte';
  import Block from "./lib/Block.svelte";
</script>

<div class="weaponsGrid">
    <div class="sidearms">
        <Item weapon="classic" href="" type="pistol"/>
        <Item weapon="shorty" href="" type="pistol"/>
        <Item weapon="frenzy" href="" type="pistol"/>
        <Item weapon="ghost" href="" type="pistol"/>
        <Item weapon="sheriff" href="" type="pistol"/>
    </div>

    <div class="smgs-shotguns">
        <Item weapon="stinger" href="" type="smg"/>
        <Item weapon="spectre" href="" type="smg"/>
        <p/>
        <Item weapon="bucky" href="" type="shotgun"/>
        <Item weapon="judge" href="" type="shotgun"/>
    </div>

    <div class="rifles-melee">
        <Item weapon="bulldog" href="" type="rifle"/>
        <Item weapon="guardian" href="" type="rifle"/>
        <Item weapon="phantom" href="" type="rifle"/>
        <Item weapon="vandal" href="" type="rifle"/>
        <p/>
        <Item weapon="melee" href="" type="knife"/>
    </div>

    <div class="snipers-lmgs">
        <Item weapon="marshal" href="" type="sniper"/>
        <Item weapon="outlaw" href="" type="sniper"/>
        <Item weapon="operator" href="" type="sniper"/>
        <p/>
        <Item weapon="ares" href="" type="lmg"/>
        <Item weapon="odin" href="" type="lmg"/>
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

    .weaponsGrid{display: flex;}

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