<script type="application/typescript">
  import Item from './lib/Item.svelte';
  import Block from "./lib/Block.svelte";

  // Load all the stores from stores.js
  import { classic } from './lib/stores.js'
  import { shorty } from './lib/stores.js'
  import { frenzy } from './lib/stores.js'
  import { ghost } from './lib/stores.js'
  import { sheriff } from './lib/stores.js'
  import { stinger } from './lib/stores.js'
  import { spectre } from './lib/stores.js'
  import { bucky } from './lib/stores.js'
  import { judge } from './lib/stores.js'
  import { bulldog } from './lib/stores.js'
  import { guardian } from './lib/stores.js'
  import { phantom } from './lib/stores.js'
  import { vandal } from './lib/stores.js'
  import { marshal } from './lib/stores.js'
  import { operator } from './lib/stores.js'
  import { ares } from './lib/stores.js'
  import { odin } from './lib/stores.js'
  import { melee } from './lib/stores.js'
  import { outlaw } from './lib/stores.js'
</script>

<div class="weaponsGrid">
    <div class="sidearms">
        <Item weapon="classic" href={$classic} type="pistol"/>
        <Item weapon="shorty" href={$shorty} type="pistol"/>
        <Item weapon="frenzy" href={$frenzy} type="pistol"/>
        <Item weapon="ghost" href={$ghost} type="pistol"/>
        <Item weapon="sheriff" href={$sheriff} type="pistol"/>
    </div>

    <div class="smgs-shotguns">
        <Item weapon="stinger" href={$stinger} type="smg"/>
        <Item weapon="spectre" href={$spectre} type="smg"/>
        <p/>
        <Item weapon="bucky" href={$bucky} type="shotgun"/>
        <Item weapon="judge" href={$judge} type="shotgun"/>
    </div>

    <div class="rifles-melee">
        <Item weapon="bulldog" href={$bulldog} type="rifle"/>
        <Item weapon="guardian" href={$guardian} type="rifle"/>
        <Item weapon="phantom" href={$phantom} type="rifle"/>
        <Item weapon="vandal" href={$vandal} type="rifle"/>
        <p/>
        <Item weapon="melee" href={$melee} type="knife"/>
    </div>

    <div class="snipers-lmgs">
        <Item weapon="marshal" href={$marshal} type="sniper"/>
        <Item weapon="outlaw" href={$outlaw} type="sniper"/>
        <Item weapon="operator" href={$operator} type="sniper"/>
        <p/>
        <Item weapon="ares" href={$ares} type="lmg"/>
        <Item weapon="odin" href={$odin} type="lmg"/>
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
        overflow: hidden;
        background-attachment: fixed;
        background-image: url("./assets/bg_main.png");
    }
</style>